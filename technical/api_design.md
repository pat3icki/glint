## API Design for Decentralized Crypto Gift Card Platform

Based on the smart contract (`GiftCardMinter`) and PostgreSQL schema, this document defines the **backend API** for issuers (organizations) and public card resolution.

### Core Concepts

- **Issuer**: Organization that creates and sells gift cards.
- **Contract**: Deployment of `GiftCardMinter` on a specific EVM network, with its pool configurations.
- **Batch**: A logical group of physical cards (e.g., 500 cards of $25 USDC) linked to a single blockchain pool (`poolId`).
- **Card**: Individual physical card with a unique serial, QR code, and scratch-off code.
- **Resolution**: Public endpoint that verifies card authenticity and returns detailed card information (denomination, issuer, redemption status, etc.).

---

### QR Code Content

The QR code printed on each physical card encodes a **JSON object** (or compact URI-encoded string) containing:

```json
{
  "serial": 1234567890123456,          // int64 (unique serial number)
  "checksum": "0x...",                // bytes32 (hash of all fields except checksum)
  "cryptoAddress": "0x...",           // public address of the card’s embedded wallet
  "smartContract": "0x...",           // contract address
  "network": "ethereum-mainnet",      // chain identifier
  "issuer": "Glint Cards",            // issuer display name
  "timestamp": "2026-08-02T10:00:00Z" // issuance timestamp (ISO 8601)
}
```

This data is used by the `resolve` endpoint to verify authenticity via the contract’s `verifyCardAuthenticity(serial, checksum)`.

---

### General API Conventions

- Base URL: `https://api.example.com/api/v1`
- Content-Type: `application/json`
- Authentication: Bearer JWT token (required for issuer endpoints)
- Error Responses: Standard HTTP status codes with `{ "error": "message" }`

---

### Issuer Endpoints (Authenticated)

#### 1. Register Issuer

**POST** `/issuers/register`

Registers a new issuer organization.

**Request Body**

```json
{
  "displayName": "Glint Cards",
  "email": "admin@glint.cards",
  "password": "secure-password",
  "organizationName": "Glint Gift Cards Inc.",
  "description": "Premium crypto gift cards",
  "website": "https://glint.cards"
}
```

**Response (201 Created)**

```json
{
  "id": "uuid",
  "displayName": "Glint Cards",
  "email": "admin@glint.cards",
  "organizationName": "Glint Gift Cards Inc.",
  "description": "Premium crypto gift cards",
  "website": "https://glint.cards",
  "createdAt": "2026-08-02T10:00:00Z"
}
```

---

#### 2. Login

**POST** `/issuers/login`

**Request Body**

```json
{
  "email": "admin@glint.cards",
  "password": "secure-password"
}
```

**Response (200 OK)**

```json
{
  "token": "jwt-token",
  "issuer": { ... } // same as profile
}
```

---

#### 3. Get Issuer Profile

**GET** `/issuers/me`

**Response (200 OK)**

```json
{
  "id": "uuid",
  "displayName": "Glint Cards",
  "email": "admin@glint.cards",
  "organizationName": "Glint Gift Cards Inc.",
  "description": "Premium crypto gift cards",
  "website": "https://glint.cards",
  "createdAt": "2026-08-02T10:00:00Z",
  "updatedAt": "2026-08-02T10:00:00Z"
}
```

---

#### 4. Register a Smart Contract Deployment

**POST** `/contracts`

Records a deployed `GiftCardMinter` contract.

**Request Body**

```json
{
  "name": "USDC Gift Card Minter",
  "evmNetwork": "ethereum-mainnet",
  "evmContractAddress": "0x1234...",
  "erc20TokenAddress": "0xUSDC...",
  "erc20TokenDecimal": 6,
  "trustedForwarder": "0xForwarder...",
  "maxBatchSize": 100,
  "baseUri": "https://metadata.glint.cards/{id}.json",
  "pool": { // optional JSON describing existing pools
    "1": 25,
    "2": 50
  }
}
```

**Response (201 Created)**

```json
{
  "id": "uuid",
  "name": "USDC Gift Card Minter",
  "evmNetwork": "ethereum-mainnet",
  "evmContractAddress": "0x1234...",
  "erc20TokenAddress": "0xUSDC...",
  "erc20TokenDecimal": 6,
  "trustedForwarder": "0xForwarder...",
  "maxBatchSize": 100,
  "baseUri": "https://metadata.glint.cards/{id}.json",
  "pool": { "1": 25, "2": 50 },
  "createdAt": "...",
  "updatedAt": "..."
}
```

---

#### 5. List Contracts

**GET** `/contracts`

**Response (200 OK)**

```json
{
  "contracts": [ ... ]
}
```

---

#### 6. Create a Card Batch

**POST** `/batches`

Creates a batch of physical cards linked to an existing blockchain pool. This step typically precedes minting on-chain.

**Request Body**

```json
{
  "contractId": "uuid-of-contract",
  "poolId": 1,                    // numeric pool ID from contract
  "version": "v1.0",
  "denomination": 25,             // face value in whole tokens
  "tokenDecimals": 6,
  "merkleRoot": "0x...",          // optional merkle root
  "merkleRootSibling": "0x...",
  "manufacturer": "CardFactory Inc.",
  "issuerId": "uuid-of-issuer",   // or inferred from auth
  "cardCount": 500                // number of cards in this batch
}
```

**Response (201 Created)**

```json
{
  "batchId": "uuid",
  "contractId": "...",
  "blockchainPoolId": 1,
  "denomination": 25,
  "cardCount": 500,
  "createdAt": "...",
  "issuerId": "...",
  "cards": {} // initially empty JSON object
}
```

---

#### 7. Register Physical Cards (Bulk)

**POST** `/batches/{batchId}/cards`

Registers each physical card’s serial and content hash on-chain (`registerPhysicalCards`) and stores individual card records in the database.

**Request Body**

```json
{
  "cards": [
    {
      "serial": 1234567890123456,
      "contentHash": "0x...",      // bytes32
      "cardIndex": 0,              // index within batch
      "walletAddress": "0x...",    // embedded wallet public key
      "qrCode": "base64...",       // optional QR image bytes
      "scratchCode": "A2B4C6D8E0F1..." // encrypted private key or plain 16-digit code
    },
    // ... up to max batch size
  ]
}
```

> **Note**: The `contentHash` is the SHA-256 checksum of the concatenated card fields (excluding the checksum itself). The smart contract stores `serial → contentHash` for authenticity verification.

**Response (201 Created)**

```json
{
  "batchId": "uuid",
  "registeredCount": 500,
  "failedSerial": [] // any serials that failed (e.g., duplicate)
}
```

---

#### 8. List Batches

**GET** `/batches`

Query params: `issuerId`, `contractId`, `limit`, `offset`

**Response (200 OK)**

```json
{
  "batches": [
    {
      "batchId": "uuid",
      "version": "v1.0",
      "denomination": 25,
      "blockchainPoolId": 1,
      "network": "ethereum-mainnet",
      "smartContractAddress": "0x...",
      "tokenAddress": "0x...",
      "issuedAt": "...",
      "manufacturer": "...",
      "issuerId": "...",
      "cardCount": 500,
      "createdAt": "..."
    }
  ],
  "total": 10,
  "limit": 20,
  "offset": 0
}
```

---

#### 9. Get Batch Details

**GET** `/batches/{batchId}`

**Response (200 OK)**

Includes batch metadata and a list of cards (paginated).

```json
{
  "batchId": "uuid",
  "version": "v1.0",
  "denomination": 25,
  "blockchainPoolId": 1,
  "network": "ethereum-mainnet",
  "smartContractAddress": "0x...",
  "tokenAddress": "0x...",
  "issuedAt": "...",
  "manufacturer": "...",
  "issuerId": "...",
  "createdAt": "...",
  "cards": [
    {
      "serial": 1234567890123456,
      "cardIndex": 0,
      "walletAddress": "0x...",
      "redeemed": false,
      "createdAt": "..."
    }
    // ... paginated
  ],
  "totalCards": 500
}
```

---

### Public Endpoints (No Authentication)

#### 1. Resolve Card Information

**GET** `/cards/resolve`

This endpoint is called by the user’s wallet/dApp after scanning the QR code (or manually entering serial and checksum). It verifies card authenticity and returns full card details.

**Query Parameters** (either via QR code data or manual input)

- `serial` (required) – int64
- `checksum` (required) – bytes32 (hex string)
- Optionally, the client can send the full QR payload as JSON in a POST request.

**GET Example**

```
GET /api/v1/cards/resolve?serial=1234567890123456&checksum=0xabcd...
```

**Response (200 OK)**

```json
{
  "serial": 1234567890123456,
  "checksum": "0xabcd...",
  "authentic": true,                // result of contract.verifyCardAuthenticity()
  "redeemed": false,
  "batch": {
    "batchId": "uuid",
    "denomination": 25,             // in whole tokens (e.g., 25 USDC)
    "tokenDecimals": 6,
    "network": "ethereum-mainnet",
    "smartContractAddress": "0x...",
    "tokenAddress": "0x...",
    "issuer": {
      "displayName": "Glint Cards",
      "organizationName": "Glint Gift Cards Inc.",
      "website": "https://glint.cards"
    }
  },
  "cardIndex": 0,
  "walletAddress": "0x...",         // embedded wallet (public)
  "issuedAt": "2026-08-02T10:00:00Z",
  "qrCodeData": {                  // optional, original QR payload
    "cryptoAddress": "0x...",
    "smartContract": "0x...",
    "network": "ethereum-mainnet",
    "issuer": "Glint Cards",
    "timestamp": "2026-08-02T10:00:00Z"
  }
}
```

If the card is not authentic (checksum mismatch) or not found:

```json
{
  "serial": 1234567890123456,
  "authentic": false,
  "error": "Card not registered or checksum invalid"
}
```

If the card is redeemed:

```json
{
  "serial": 1234567890123456,
  "authentic": true,
  "redeemed": true,
  "redeemedAt": "2026-08-02T12:00:00Z",
  "redeemedBy": "0x...",
  "redemptionTxHash": "0x..."
}
```

**POST Alternative** (for large payloads)

```
POST /api/v1/cards/resolve
Content-Type: application/json

{
  "serial": 1234567890123456,
  "checksum": "0xabcd..."
}
```

Same response.

---

### Card Redemption Flow (On-Chain)

Redemption is performed directly via the smart contract’s `burnAndWithdraw` function. The backend **does not** handle redemption; it only listens for on-chain events (`CardRedeemed`, `BatchCardsRedeemed`) and updates the database `redeemed` flag and records the transaction hash.

**Webhook/Event Listener** (optional)  
The backend can subscribe to contract events and call:

```
POST /internal/events/card-redeemed
{
  "serial": 1234567890123456,
  "redeemer": "0x...",
  "poolId": 1,
  "tokenAmount": "1500000000000000000", // in wei
  "txHash": "0x..."
}
```

Then update `cards.redeemed = true`, `redeem_card_records`.

---

### Summary of API Paths

| Method | Endpoint                  | Description                          | Auth |
|--------|---------------------------|--------------------------------------|------|
| POST   | `/issuers/register`       | Register issuer                      | No   |
| POST   | `/issuers/login`          | Login                                | No   |
| GET    | `/issuers/me`             | Get profile                          | Yes  |
| POST   | `/contracts`              | Register contract deployment         | Yes  |
| GET    | `/contracts`              | List contracts                       | Yes  |
| POST   | `/batches`                | Create a card batch                  | Yes  |
| GET    | `/batches`                | List batches                         | Yes  |
| GET    | `/batches/{batchId}`      | Get batch + cards                    | Yes  |
| POST   | `/batches/{batchId}/cards`| Register physical cards              | Yes  |
| GET    | `/cards/resolve`          | Verify and get card info (public)   | No   |
| POST   | `/cards/resolve`          | Alternative resolve (public)         | No   |

---

### Security Considerations

- JWT tokens should have short expiry and refresh mechanism.
- Issuer passwords must be hashed (bcrypt/Argon2) – the schema shows `password_hash TEXT`.
- The `resolve` endpoint must call the contract’s `verifyCardAuthenticity` to prevent off-chain tampering.
- Rate limiting on public endpoints to avoid brute‑force serial enumeration.
- The QR code content should be signed or include a MAC to prevent forgery; the checksum serves as a hash but not a signature – the contract stores the checksum, so only the issuer can register valid checksums.

---

This API design covers the full lifecycle: issuer onboarding, contract registration, batch creation, physical card registration, and public card resolution – enabling a fully decentralized gift card platform with on-chain solvency verification and off-chain metadata. 