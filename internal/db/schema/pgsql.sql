-- ============================================================
-- GLINT ECOSYSTEM — DYNAMIC POSTGRES SCHEMA (v4)
--
-- Flavour-agnostic. glint_flavours is the registry that drives everything:
--   * flavour_event_types   : events a flavour emits + JSON Schema for decoded `data`
--   * flavour_entity_types  : derived state objects + JSON Schema for `state`
--   * glint_flavours.config_schema : JSON Schema for contracts.config
--
-- The application validates writes against these JSON Schemas before
-- inserting (Postgres does not validate JSON Schema natively).
--
-- v4: no created_at / updated_at columns. Row history is recorded in
-- schema_logs.logs (CREATE / UPDATE / DELETE) by an audit trigger.
--
-- Requires PostgreSQL 13+. Idempotent. Addresses/hashes stored lowercase.
-- ============================================================

CREATE SCHEMA IF NOT EXISTS blockchain;
CREATE SCHEMA IF NOT EXISTS schema_logs;
CREATE SCHEMA IF NOT EXISTS public;

-- ============================================================
-- DOMAINS
-- ============================================================
DO $$ BEGIN
    CREATE DOMAIN blockchain.evm_address AS VARCHAR(42)
        CHECK (VALUE ~ '^0x[0-9a-f]{40}$');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

DO $$ BEGIN
    CREATE DOMAIN blockchain.evm_hash32 AS VARCHAR(66)
        CHECK (VALUE ~ '^0x[0-9a-f]{64}$');
EXCEPTION WHEN duplicate_object THEN NULL; END $$;

-- ============================================================
-- LOGS (audit trail)
-- One row per INSERT / UPDATE / DELETE on an audited table.
--   CREATE : old_data NULL,     new_data = row
--   UPDATE : old_data = before, new_data = after
--   DELETE : old_data = row,    new_data NULL
-- `record_pk` holds the primary-key columns of the affected row,
-- so composite keys work too. Set the actor per transaction with:
--   SET LOCAL app.current_user = '<issuer id or service name>';
-- ============================================================
CREATE TABLE IF NOT EXISTS schema_logs.logs (
    id          BIGSERIAL PRIMARY KEY,
    action      VARCHAR(10) NOT NULL CHECK (action IN ('CREATE', 'UPDATE', 'DELETE')),
    table_name  VARCHAR(100) NOT NULL,
    record_pk   JSONB NOT NULL,
    old_data    JSONB,
    new_data    JSONB,
    changed_by  TEXT DEFAULT NULLIF(current_setting('app.current_user', TRUE), ''),
    logged_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT logs_payload_matches_action CHECK (
        (action = 'CREATE' AND old_data IS NULL     AND new_data IS NOT NULL) OR
        (action = 'UPDATE' AND old_data IS NOT NULL AND new_data IS NOT NULL) OR
        (action = 'DELETE' AND old_data IS NOT NULL AND new_data IS NULL)
    )
);

CREATE INDEX IF NOT EXISTS idx_logs_table_time ON schema_logs.logs (table_name, logged_at DESC);
CREATE INDEX IF NOT EXISTS idx_logs_action     ON schema_logs.logs (action);
CREATE INDEX IF NOT EXISTS idx_logs_record_pk  ON schema_logs.logs USING GIN (record_pk jsonb_path_ops);

-- ============================================================
-- AUDIT TRIGGER FUNCTION
-- Usage: EXECUTE FUNCTION blockchain.log_row_change('<pk cols, comma-separated>'[, '<cols to exclude from the log>'])
-- Skips UPDATEs that change nothing.
-- ============================================================
CREATE OR REPLACE FUNCTION blockchain.log_row_change()
RETURNS TRIGGER AS $$
DECLARE
    pk_cols  TEXT[] := string_to_array(COALESCE(TG_ARGV[0], 'id'), ',');
    excluded TEXT[] := CASE WHEN TG_NARGS > 1 THEN string_to_array(TG_ARGV[1], ',') ELSE ARRAY[]::TEXT[] END;
    old_j    JSONB;
    new_j    JSONB;
    pk       JSONB;
    act      VARCHAR(10);
BEGIN
    IF TG_OP = 'INSERT' THEN
        act   := 'CREATE';
        new_j := to_jsonb(NEW) - excluded;
        pk    := (SELECT jsonb_object_agg(c, to_jsonb(NEW) -> c) FROM unnest(pk_cols) AS c);
    ELSIF TG_OP = 'UPDATE' THEN
        act   := 'UPDATE';
        old_j := to_jsonb(OLD) - excluded;
        new_j := to_jsonb(NEW) - excluded;
        IF old_j = new_j THEN
            RETURN NULL;
        END IF;
        pk    := (SELECT jsonb_object_agg(c, to_jsonb(NEW) -> c) FROM unnest(pk_cols) AS c);
    ELSE
        act   := 'DELETE';
        old_j := to_jsonb(OLD) - excluded;
        pk    := (SELECT jsonb_object_agg(c, to_jsonb(OLD) -> c) FROM unnest(pk_cols) AS c);
    END IF;

    INSERT INTO schema_logs.logs (action, table_name, record_pk, old_data, new_data)
    VALUES (act, TG_TABLE_NAME, pk, old_j, new_j);

    RETURN NULL;  -- AFTER trigger: return value is ignored
END;
$$ LANGUAGE plpgsql;

-- ============================================================
-- GLINT FLAVOURS (the registry root)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.glint_flavours (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    flavour_key    VARCHAR(50)  UNIQUE NOT NULL,
    display_name   VARCHAR(100) UNIQUE NOT NULL,
    description    TEXT,
    abi            JSONB,                                   -- full ABI, used by the program to decode logs
    config_schema  JSONB NOT NULL DEFAULT '{}'::JSONB,      -- JSON Schema for contracts.config
    schema_version INTEGER NOT NULL DEFAULT 1 CHECK (schema_version > 0),
    is_active      BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT flavours_config_schema_is_object CHECK (jsonb_typeof(config_schema) = 'object')
);

DROP TRIGGER IF EXISTS log_glint_flavours ON blockchain.glint_flavours;
CREATE TRIGGER log_glint_flavours
AFTER INSERT OR UPDATE OR DELETE ON blockchain.glint_flavours
FOR EACH ROW EXECUTE FUNCTION blockchain.log_row_change('id');

-- ============================================================
-- FLAVOUR EVENT TYPES
-- Every event a flavour can emit. `category` is free-form and
-- flavour-defined. `data_schema` validates contract_events.data.
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.flavour_event_types (
    flavour_id     UUID NOT NULL REFERENCES blockchain.glint_flavours(id) ON DELETE CASCADE,
    event_type     VARCHAR(100) NOT NULL,
    category       VARCHAR(60)  NOT NULL,
    topic0         blockchain.evm_hash32,                   -- keccak256(signature); routes logs to a type
    abi_fragment   JSONB,
    data_schema    JSONB NOT NULL DEFAULT '{}'::JSONB,
    schema_version INTEGER NOT NULL DEFAULT 1 CHECK (schema_version > 0),
    is_active      BOOLEAN NOT NULL DEFAULT TRUE,
    PRIMARY KEY (flavour_id, event_type),
    CONSTRAINT event_types_data_schema_is_object CHECK (jsonb_typeof(data_schema) = 'object')
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_flavour_event_types_topic0
    ON blockchain.flavour_event_types (flavour_id, topic0) WHERE topic0 IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_flavour_event_types_category
    ON blockchain.flavour_event_types (flavour_id, category);

DROP TRIGGER IF EXISTS log_flavour_event_types ON blockchain.flavour_event_types;
CREATE TRIGGER log_flavour_event_types
AFTER INSERT OR UPDATE OR DELETE ON blockchain.flavour_event_types
FOR EACH ROW EXECUTE FUNCTION blockchain.log_row_change('flavour_id,event_type');

-- ============================================================
-- FLAVOUR ENTITY TYPES
-- Derived state objects the program maintains from events.
-- `state_schema` validates contract_entities.state.
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.flavour_entity_types (
    flavour_id     UUID NOT NULL REFERENCES blockchain.glint_flavours(id) ON DELETE CASCADE,
    entity_type    VARCHAR(100) NOT NULL,
    description    TEXT,
    state_schema   JSONB NOT NULL DEFAULT '{}'::JSONB,
    schema_version INTEGER NOT NULL DEFAULT 1 CHECK (schema_version > 0),
    is_active      BOOLEAN NOT NULL DEFAULT TRUE,
    PRIMARY KEY (flavour_id, entity_type),
    CONSTRAINT entity_types_state_schema_is_object CHECK (jsonb_typeof(state_schema) = 'object')
);

DROP TRIGGER IF EXISTS log_flavour_entity_types ON blockchain.flavour_entity_types;
CREATE TRIGGER log_flavour_entity_types
AFTER INSERT OR UPDATE OR DELETE ON blockchain.flavour_entity_types
FOR EACH ROW EXECUTE FUNCTION blockchain.log_row_change('flavour_id,entity_type');

-- ============================================================
-- ISSUERS
-- password_hash is excluded from the audit log.
-- ============================================================
CREATE TABLE IF NOT EXISTS issuers (
    id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    display_name      VARCHAR(255) UNIQUE,
    email             VARCHAR(255) NOT NULL,
    password_hash     TEXT NOT NULL,
    organization_name VARCHAR(255) NOT NULL,
    description       TEXT,
    website           VARCHAR(255)
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_issuers_email_lower ON issuers (LOWER(email));

DROP TRIGGER IF EXISTS log_issuers ON issuers;
CREATE TRIGGER log_issuers
AFTER INSERT OR UPDATE OR DELETE ON issuers
FOR EACH ROW EXECUTE FUNCTION blockchain.log_row_change('id', 'password_hash');

-- ============================================================
-- TOKENS (optional generic ERC-20 reference, scoped per network)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.tokens (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    evm_network    VARCHAR(255) NOT NULL,
    smart_contract blockchain.evm_address NOT NULL,
    symbol         VARCHAR(255) NOT NULL,
    decimals       INTEGER NOT NULL CHECK (decimals >= 0 AND decimals <= 77),
    is_active      BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT unique_token_network_address UNIQUE (evm_network, smart_contract)
);

CREATE INDEX IF NOT EXISTS idx_tokens_is_active ON blockchain.tokens (is_active) WHERE is_active = TRUE;

DROP TRIGGER IF EXISTS log_tokens ON blockchain.tokens;
CREATE TRIGGER log_tokens
AFTER INSERT OR UPDATE OR DELETE ON blockchain.tokens
FOR EACH ROW EXECUTE FUNCTION blockchain.log_row_change('id');

-- ============================================================
-- CONTRACTS (deployment metadata; flavour-specific data in `config`)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.contracts (
    id                    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    issuer_id             UUID NOT NULL REFERENCES issuers(id) ON DELETE RESTRICT,
    flavour_id            UUID NOT NULL REFERENCES blockchain.glint_flavours(id) ON DELETE RESTRICT,
    name                  VARCHAR(255) NOT NULL,
    evm_network           VARCHAR(255) NOT NULL,
    evm_contract_address  blockchain.evm_address NOT NULL,
    deployment_tx_hash    blockchain.evm_hash32,
    deployed_block        BIGINT,
    config                JSONB NOT NULL DEFAULT '{}'::JSONB,
    config_schema_version INTEGER,                          -- flavour schema_version used to validate `config`
    CONSTRAINT unique_contract_network_address UNIQUE (evm_network, evm_contract_address),
    CONSTRAINT unique_contract_id_flavour UNIQUE (id, flavour_id),
    CONSTRAINT contracts_config_is_object CHECK (jsonb_typeof(config) = 'object')
);

CREATE INDEX IF NOT EXISTS idx_contracts_issuer_id  ON blockchain.contracts (issuer_id);
CREATE INDEX IF NOT EXISTS idx_contracts_flavour_id ON blockchain.contracts (flavour_id);

DROP TRIGGER IF EXISTS log_contracts ON blockchain.contracts;
CREATE TRIGGER log_contracts
AFTER INSERT OR UPDATE OR DELETE ON blockchain.contracts
FOR EACH ROW EXECUTE FUNCTION blockchain.log_row_change('id');

-- ============================================================
-- CONTRACT EVENTS (raw + decoded, one row per log)
-- Append-only chain data, so it is NOT audited in `logs`
-- (block_timestamp is the chain's own time).
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.contract_events (
    id              BIGSERIAL PRIMARY KEY,
    contract_id     UUID NOT NULL,
    flavour_id      UUID NOT NULL,
    event_type      VARCHAR(100) NOT NULL,
    schema_version  INTEGER NOT NULL DEFAULT 1,
    tx_hash         blockchain.evm_hash32 NOT NULL,
    log_index       INTEGER NOT NULL CHECK (log_index >= 0),
    block_number    BIGINT NOT NULL,
    block_hash      blockchain.evm_hash32,
    block_timestamp TIMESTAMPTZ,
    emitter         blockchain.evm_address NOT NULL,
    topics          JSONB NOT NULL DEFAULT '[]'::JSONB,
    data            JSONB NOT NULL DEFAULT '{}'::JSONB,
    FOREIGN KEY (contract_id, flavour_id)
        REFERENCES blockchain.contracts (id, flavour_id) ON DELETE CASCADE,
    FOREIGN KEY (flavour_id, event_type)
        REFERENCES blockchain.flavour_event_types (flavour_id, event_type) ON DELETE RESTRICT,
    CONSTRAINT unique_event_log UNIQUE (contract_id, tx_hash, log_index),
    CONSTRAINT contract_events_topics_is_array CHECK (jsonb_typeof(topics) = 'array'),
    CONSTRAINT contract_events_data_is_object  CHECK (jsonb_typeof(data) = 'object')
);

CREATE INDEX IF NOT EXISTS idx_contract_events_contract_block
    ON blockchain.contract_events (contract_id, block_number, log_index);
CREATE INDEX IF NOT EXISTS idx_contract_events_contract_type_block
    ON blockchain.contract_events (contract_id, event_type, block_number DESC);
CREATE INDEX IF NOT EXISTS idx_contract_events_flavour_type
    ON blockchain.contract_events (flavour_id, event_type);
CREATE INDEX IF NOT EXISTS idx_contract_events_tx
    ON blockchain.contract_events (tx_hash);
CREATE INDEX IF NOT EXISTS idx_contract_events_timestamp
    ON blockchain.contract_events (block_timestamp);
CREATE INDEX IF NOT EXISTS idx_contract_events_data_gin
    ON blockchain.contract_events USING GIN (data jsonb_path_ops);

-- ============================================================
-- CONTRACT ENTITIES (generic derived state, one row per object)
-- Derived and high-churn, so NOT audited in `logs`; it can be
-- rebuilt by replaying contract_events. To audit it anyway, add:
--   CREATE TRIGGER log_contract_entities
--   AFTER INSERT OR UPDATE OR DELETE ON blockchain.contract_entities
--   FOR EACH ROW EXECUTE FUNCTION blockchain.log_row_change('contract_id,entity_type,entity_key');
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.contract_entities (
    contract_id    UUID NOT NULL,
    flavour_id     UUID NOT NULL,
    entity_type    VARCHAR(100) NOT NULL,
    entity_key     VARCHAR(255) NOT NULL,
    state          JSONB NOT NULL DEFAULT '{}'::JSONB,
    schema_version INTEGER NOT NULL DEFAULT 1,
    last_event_id  BIGINT REFERENCES blockchain.contract_events(id) ON DELETE SET NULL,
    last_block     BIGINT,
    PRIMARY KEY (contract_id, entity_type, entity_key),
    FOREIGN KEY (contract_id, flavour_id)
        REFERENCES blockchain.contracts (id, flavour_id) ON DELETE CASCADE,
    FOREIGN KEY (flavour_id, entity_type)
        REFERENCES blockchain.flavour_entity_types (flavour_id, entity_type) ON DELETE RESTRICT,
    CONSTRAINT contract_entities_state_is_object CHECK (jsonb_typeof(state) = 'object')
);

CREATE INDEX IF NOT EXISTS idx_contract_entities_flavour_type
    ON blockchain.contract_entities (flavour_id, entity_type);
CREATE INDEX IF NOT EXISTS idx_contract_entities_state_gin
    ON blockchain.contract_entities USING GIN (state jsonb_path_ops);

-- ============================================================
-- INGESTION CURSOR (not audited; written constantly by the indexer)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.event_ingestion_cursors (
    contract_id       UUID NOT NULL REFERENCES blockchain.contracts(id) ON DELETE CASCADE,
    indexer_name      VARCHAR(100) NOT NULL,
    last_block_number BIGINT NOT NULL DEFAULT 0 CHECK (last_block_number >= 0),
    last_block_hash   blockchain.evm_hash32,
    PRIMARY KEY (contract_id, indexer_name)
);
