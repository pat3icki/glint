# `blockchain` — Plugin-based EVM contract indexer

A framework for deterministically indexing EVM smart-contract events into
Postgres, built around a plugin interface that isolates each contract family
into its own self-contained sub-package.

`glintdough` is the reference implementation.

---

## Package layout

```
pkg/
├── blockchain/            # Root framework — Plugin interface, Registry, shared types, sentinel errors
│   ├── chain/             # ethclient wrapper (the only authorised path to EVM RPC)
│   ├── store/             # pgx/v5 wrapper (the only authorised path to SQL)
│   ├── schema/            # JSON Schema draft 2020-12 compiler/validator
│   └── ingest/            # Ordered ingest pipeline (decode → insert → project → apply → advance)
└── glintdough/            # GlintDough ERC-1155 plugin (reference implementation)
    └── schemas/           # Embedded JSON Schemas for every event and entity type
```

### Import rules

| Package | May import | Must NOT import |
|---|---|---|
| `blockchain` | stdlib, go-ethereum types | Any sub-package |
| `blockchain/chain` | `blockchain`, go-ethereum | `store`, `ingest`, `schema`, sibling plugins |
| `blockchain/store` | `blockchain`, pgx/v5 | `chain`, `ingest`, `schema`, sibling plugins |
| `blockchain/schema` | stdlib, jsonschema/v6 | Any other framework package |
| `blockchain/ingest` | `blockchain`, `chain`, `store` | Sibling plugins |
| `glintdough` | `blockchain`, go-ethereum | Any sibling plugin |

The root package never imports sub-packages. Sub-packages never import siblings.

---

## Core concepts

### Plugin

Every contract family implements `blockchain.Plugin`:

```go
type Plugin interface {
    FlavourKey()    string          // stable DB key, e.g. "glintdough"
    DisplayName()   string
    Description()   string
    SchemaVersion() int

    ConfigSchema() json.RawMessage  // JSON Schema for contracts.config
    ABI()          json.RawMessage  // full ABI JSON

    EventTypes()  []EventType       // catalogue of events the contract emits
    EntityTypes() []EntityType      // catalogue of derived-state objects maintained

    DecodeEvent(ctx, log) (*DecodedEvent, error)
    Project(ctx, event)  ([]EntityMutation, error)
}
```

Plugins register themselves into `blockchain.GlobalRegistry` via their
package-level `init()` function. Importing the plugin package is enough to
activate it — no manual wiring required.

```go
import _ "github.com/pat3icki/glint/pkg/glintdough" // side-effect: registers plugin
```

### Registry

```go
// Register a plugin (called from init() inside each plugin package).
blockchain.Register(plugin)

// Look up a plugin at runtime.
p, err := blockchain.Lookup("glintdough")

// Iterate all registered plugins (e.g. to seed the DB at startup).
for _, p := range blockchain.GlobalRegistry.All() { ... }
```

`Register` panics on duplicate keys — consistent with `database/sql` driver
registration.

### Ingest pipeline

For every EVM log, the `ingest.Indexer` executes these steps **in order**,
each inside its own serialisable Postgres transaction:

```
FilterLogs (batch)
  └── for each log (sorted by block, then log index):
        1. Decode   — plugin.DecodeEvent()        → *DecodedEvent
        2. Insert   — store.InsertEvent()          → event ID  (idempotent)
        3. Project  — plugin.Project()             → []EntityMutation
        4. Apply    — store.UpsertEntity() × N     (same tx as step 2)
        5. Advance  — store.AdvanceCursor()        (after batch commits)
```

Poisoned logs (decode errors) are logged via `slog` and skipped without
failing the batch. Already-ingested logs (`ErrEventAlreadyIngested`) are
silently skipped.

---

## Sentinel errors

| Sentinel | When returned |
|---|---|
| `ErrPluginNotFound` | `Lookup` can't find a flavour key |
| `ErrPluginAlreadyRegistered` | `Register` called twice with the same key |
| `ErrUnknownEvent` | `DecodeEvent` can't match `topic0` |
| `ErrInvalidData` | Decoded data fails JSON-Schema validation |
| `ErrEventAlreadyIngested` | Unique constraint on `(contract_id, tx_hash, log_index)` fired |
| `ErrCursorNotFound` | No cursor row exists for a given `contract_id + indexer_name` |

All errors are designed for use with `errors.Is`.

---

## Sub-packages

### `blockchain/chain`

Wraps `ethclient.Client`. No other package may import `ethclient` directly.

```go
c, err := chain.Dial(ctx, "wss://mainnet.example.com")
defer c.Close()

latest, err := c.LatestBlockNumber(ctx)
logs, err   := c.FilterLogs(ctx, ethereum.FilterQuery{...})
```

### `blockchain/store`

Wraps `pgx/v5`. All SQL lives exclusively in this package.

```go
s, err := store.New(ctx, "postgres://...")
defer s.Close()

// Idempotent — safe to call at every startup.
err = s.SeedFlavour(ctx, plugin)

// Idempotent event insert.
eventID, err := s.InsertEvent(ctx, decodedEvent)

// Entity projection.
err = s.UpsertEntity(ctx, &mutation)

// Cursor management.
err = s.AdvanceCursor(ctx, contractID, indexerName, blockNum, blockHash)
from, hash, err := s.GetCursor(ctx, contractID, indexerName)
```

`store.WithTx` wraps a callback in a serialisable transaction and rolls back
on any error.

### `blockchain/schema`

Compiles and validates JSON Schemas (draft 2020-12).

```go
s := schema.MustCompile(rawSchemaJSON) // panics on bad schema — safe at init
err := s.Validate(myData)              // returns descriptive error on failure
```

### `blockchain/ingest`

Poll-based indexer. Call `RunOnce` on a ticker; it fetches all new logs since
the last cursor position and runs the full pipeline.

```go
idx := ingest.New(ingest.Config{
    ContractID:      "uuid-of-contract-row",
    FlavourID:       "uuid-of-flavour-row",
    ContractAddress: "0xabc...",
    IndexerName:     "main",
    StartBlock:      18_000_000,
    BlockBatchSize:  2000,       // defaults to 2000 if zero
}, plugin, chainClient, storeClient)

err := idx.RunOnce(ctx)
```

---

## Reference plugin: `glintdough`

Implements the `GlintDough` ERC-1155 collateral-backed gift card contract.

### Events decoded (18 total)

| Event | Category | Entity mutations |
|---|---|---|
| `TokenAdded` | `collateral_token` | Upserts `collateral_token` |
| `TokenStatusChanged` | `collateral_token` | Updates `is_active` on `collateral_token` |
| `VoucherAdded` | `voucher` | Upserts `voucher`; updates `contract_config.next_token_id` |
| `VoucherToggled` | `voucher` | Updates `is_active` on `voucher` |
| `VoucherMintedSingle` | `mint` | Updates `total_minted` on `voucher` |
| `VouchersMintedBatch` | `mint` | Updates `total_minted` on N `voucher` rows |
| `VoucherRedeemed` | `redeem` | Updates `total_redeemed` on `voucher` |
| `TransferSingle` | `transfer` | Upserts `voucher_balance` for sender and recipient |
| `TransferBatch` | `transfer` | Upserts `voucher_balance` for sender and recipient × N |
| `PhysicalCardRegistered` | `physical_card` | Upserts `physical_card` |
| `MinterRoleGranted` | `role` | Upserts `role_assignment` (MINTER_ROLE) |
| `RoleGranted` | `role` | Upserts `role_assignment` |
| `RoleRevoked` | `role` | Updates `granted=false` on `role_assignment` |
| `RoleAdminChanged` | `role` | No entity mutations |
| `FeesAddressUpdated` | `config` | Updates `fees_address` on `contract_config` |
| `Paused` | `status` | Sets `contract_status.paused = true` |
| `Unpaused` | `status` | Sets `contract_status.paused = false` |
| `URI` | `uri` | Upserts `uri` |

### Entity types (8 total)

| Entity type | Key | Description |
|---|---|---|
| `collateral_token` | token string key | Registered ERC-20 collateral tokens |
| `voucher` | ERC-1155 token ID | Gift card denominations |
| `voucher_balance` | `<holder>:<token_id>` | Per-holder balances (source of truth: TransferSingle/Batch) |
| `role_assignment` | `<role_bytes32>:<account>` | AccessControl role membership |
| `contract_status` | `singleton` | Paused / unpaused state |
| `contract_config` | `singleton` | `next_token_id` and `fees_address` |
| `physical_card` | `keccak256(serial)` | Registered physical gift cards |
| `uri` | ERC-1155 token ID | Metadata URIs |

### Registration

```go
import _ "github.com/pat3icki/glint/pkg/glintdough"
```

The plugin's `init()` calls `blockchain.Register(New())` automatically.

---

## Writing a new plugin

1. Create `pkg/<name>/` — never import `glintdough` or any sibling.
2. Implement `blockchain.Plugin` (all 9 methods).
3. Add a `register.go` with:
   ```go
   func init() { blockchain.Register(New()) }
   ```
4. Embed JSON Schemas for every event's `DataSchema` and each entity's
   `StateSchema`. Use `schema.MustCompile` at init to catch bad schemas early.
5. Call `store.SeedFlavour(ctx, plugin)` at application startup to idempotently
   write the plugin's metadata to the database.

---

## Conventions

- All EVM addresses and hashes stored/returned as **lowercase `0x`-prefixed hex**.
- `context.Context` is the first argument of every method that may block.
- `log/slog` for structured logging — no `fmt.Println` or `log.Printf`.
- Sentinel errors only — no ad-hoc `errors.New` outside `errors.go`.
- No raw SQL outside `blockchain/store`. No `ethclient` outside `blockchain/chain`.
- No `panic` outside `Register` (duplicate key) and ABI parsing at `init`.

---

## License

MIT — see [SPDX-License-Identifier](https://spdx.org/licenses/MIT.html) headers.
