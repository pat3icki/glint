-- ============================================================
-- SCHEMA
-- ============================================================
CREATE SCHEMA IF NOT EXISTS blockchain;

-- ============================================================
-- ISSUERS (organizations that login)
-- ============================================================
CREATE TABLE IF NOT EXISTS issuers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    display_name VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    organization_name VARCHAR(255) NOT NULL,
    description TEXT,
    website VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_issuers_updated_at
BEFORE UPDATE ON issuers
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ============================================================
-- CONTRACTS (deployment metadata)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.contracts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    evm_network VARCHAR(255) NOT NULL,
    evm_contract_address VARCHAR(255) NOT NULL,
    erc20_token_address VARCHAR(255) NOT NULL,
    erc20_token_decimal INT NOT NULL,
    trusted_forwarder VARCHAR(255) NOT NULL,
    max_batch_size NUMERIC(78,0) NOT NULL,
    base_uri TEXT,
    pool JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_contract_network_address UNIQUE (evm_network, evm_contract_address)
);

CREATE TRIGGER update_contracts_updated_at
BEFORE UPDATE ON blockchain.contracts
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ============================================================
-- TOKEN TRANSFER RECORDS (mint / redeem events)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.contract_token_records (
    id BIGSERIAL PRIMARY KEY,
    contract_id UUID NOT NULL REFERENCES blockchain.contracts(id) ON DELETE CASCADE,
    tx_hash VARCHAR(66) NOT NULL,
    pool_id NUMERIC(78,0) NOT NULL,
    operation_type VARCHAR(20) NOT NULL CHECK (operation_type IN ('mint', 'redeem')),
    from_address VARCHAR(255) NOT NULL,
    to_address VARCHAR(255) NOT NULL,
    amount_tokens NUMERIC(78,0) NOT NULL,
    amount_cards NUMERIC(78,0) NOT NULL,
    block_number BIGINT,
    block_timestamp TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (tx_hash, pool_id, operation_type, from_address)
);

CREATE INDEX idx_token_records_contract ON blockchain.contract_token_records(contract_id);
CREATE INDEX idx_token_records_tx ON blockchain.contract_token_records(tx_hash);
CREATE INDEX idx_token_records_pool ON blockchain.contract_token_records(pool_id);

-- ============================================================
-- CARD BATCH METADATA (renamed to card_batches)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.card_batches (
    batch_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    version VARCHAR(50) NOT NULL,
    denomination NUMERIC(78,0) NOT NULL,
    token_decimals INTEGER NOT NULL,
    merkle_root VARCHAR(66) NOT NULL,
    merkle_root_sibling BYTEA,
    network VARCHAR(255) NOT NULL,
    blockchain_pool_id NUMERIC(78,0) NOT NULL,
    meta_hash BYTEA NOT NULL CHECK (LENGTH(meta_hash) = 32),
    smart_contract_address VARCHAR(255) NOT NULL,
    token_address VARCHAR(255) NOT NULL,
    issued_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    manufacturer VARCHAR(255) NOT NULL,
    issuer_id UUID NOT NULL REFERENCES issuers(id) ON DELETE RESTRICT,
    cards JSONB NOT NULL DEFAULT '{}'::JSONB,
    CONSTRAINT cards_is_object CHECK (jsonb_typeof(cards) = 'object'),
    CONSTRAINT unique_network_contract_pool UNIQUE (network, smart_contract_address, blockchain_pool_id)
);

CREATE INDEX idx_card_batches_network_pool ON blockchain.card_batches (network, blockchain_pool_id);
CREATE INDEX idx_card_batches_contract ON blockchain.card_batches (smart_contract_address);
CREATE INDEX idx_card_batches_issuer_id ON blockchain.card_batches (issuer_id);
CREATE INDEX idx_card_batches_created_at ON blockchain.card_batches (created_at);

-- ============================================================
-- INDIVIDUAL CARDS
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.cards (
    serial BIGINT PRIMARY KEY,
    batch_id UUID NOT NULL REFERENCES blockchain.card_batches(batch_id) ON DELETE CASCADE,
    card_index NUMERIC(78,0) NOT NULL,
    hash BYTEA,
    network VARCHAR(42) NOT NULL,
    wallet_address VARCHAR(255) NOT NULL,
    qr_code TEXT,
    redeemed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (batch_id, card_index)
);

CREATE INDEX idx_cards_batch ON blockchain.cards (batch_id);
CREATE INDEX idx_cards_serial ON blockchain.cards (serial);
CREATE INDEX idx_cards_wallet ON blockchain.cards (wallet_address);
CREATE INDEX idx_cards_redeemed ON blockchain.cards (redeemed) WHERE redeemed = false;

-- ============================================================
-- REDEEM CARD RECORDS (log of redemptions)
-- ============================================================
CREATE TABLE IF NOT EXISTS blockchain.redeem_card_records (
    serial BIGINT PRIMARY KEY,
    redeemed BOOLEAN NOT NULL DEFAULT FALSE,
    redeemed_at TIMESTAMPTZ,
    redeemed_by VARCHAR(255),
    redemption_tx_hash VARCHAR(66)
);