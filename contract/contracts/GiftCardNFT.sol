// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC1155} from "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title GiftCardNFT
 * @dev ERC-1155 contract for collateral-backed crypto gift cards.
 * Supports batch minting, multiple ERC-20 collaterals, physical card registration,
 * and secure redemption with pausability and reentrancy protection.
 */
contract GiftCardNFT is ERC1155, AccessControl, Pausable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    string public constant name = "Decentralized Gift Card";
    string public constant symbol = "GCARD";
    string public constant version = "1.0.0";

    struct CardType {
        address collateral_token;
        string token_symbol;
        uint256 denomination; // in the token's native decimals
        bool isActive;
    }
    struct PhysicalCard{
        bytes32 Checksum;
        bytes QRCode;
    }

    // Mapping from Token ID to its collateral parameters
    mapping(uint256 => CardType) public cardTypes;

    // Physical card registry: keccak256(serialNumber) => checksum
    mapping(bytes32 => PhysicalCard) public physicalCardChecksums;


    // Events
    event CardTypeAdded(uint256 indexed tokenId, address indexed token, uint256 denomination);
    event CardTypeDeactivated(uint256 indexed tokenId);
    event CardsMinted(address indexed minter, uint256[] tokenIds, uint256[] amounts, address indexed to);
    event CardRedeemed(uint256 indexed tokenId, address indexed redeemer, address to, uint256 amount);
    event PhysicalCardRegistered(bytes32 indexed serialHash, bytes32 checksum);

    constructor(string memory uri) ERC1155(uri) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    // ------------------- Admin: Card Type Management -------------------

    /**
     * @notice Registers a new gift card denomination and underlying collateral token.
     * @dev Only admin can call, cannot be updated once set (use toggle to deactivate).
     */
    function addCardType(
        uint256 tokenId,
        address collateral_token,
        string calldata token_symbol,
        uint256 denomination
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(collateral_token != address(0), "Invalid token address");
        require(denomination > 0, "Denomination must be > 0");
        require(cardTypes[tokenId].collateral_token == address(0), "Card type already exists");

        cardTypes[tokenId] = CardType({
            collateral_token: collateral_token,
            token_symbol: token_symbol,
            denomination: denomination,
            isActive: true
        });

        emit CardTypeAdded(tokenId, collateral_token, denomination);
    }

    /**
     * @notice Activates or deactivates a card type (e.g., to deprecate an old collateral).
     */
    function toggleCardType(uint256 tokenId, bool status) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(cardTypes[tokenId].collateral_token != address(0), "Card type does not exist");
        cardTypes[tokenId].isActive = status;
        emit CardTypeDeactivated(tokenId);
    }

    // ------------------- Physical Card Registration -------------------

    /**
     * @notice Registers hashes and checksums for physical cards (batch).
     * @dev Only MINTER_ROLE can register.
     */
    function registerPhysicalCards(bytes32[] calldata serialHashes, bytes32[] calldata checksums)
        external
        onlyRole(MINTER_ROLE)
    {
        require(serialHashes.length == checksums.length, "Array length mismatch");
        for (uint256 i = 0; i < serialHashes.length; i++) {
            physicalCardChecksums[serialHashes[i]] = PhysicalCard({Checksum: checksums[i], QRCode: ""});
            emit PhysicalCardRegistered(serialHashes[i], checksums[i]);
        }
    }

    /**
     * @notice Verifies a physical card by comparing the provided checksum with the stored one.
     * @return bool True if the checksum matches.
     */
    function verifyPhysicalCard(bytes32 serialHash, bytes32 checksum) external view returns (bool) {
        return physicalCardChecksums[serialHash].Checksum == checksum;
    }

    // ------------------- Minting (Batch) -------------------

    /**
     * @notice Locks ERC-20 collateral for multiple token types and mints the corresponding NFTs.
     * @dev Caller (must have MINTER_ROLE) must approve this contract for the total collateral amount per token.
     * @param tokenIds Array of token IDs to mint.
     * @param amounts Array of quantities to mint for each tokenId.
     * @param to Recipient of the minted NFTs.
     * @param data Optional data for ERC-1155 mint.
     */
    function batchMintGiftCards(
        uint256[] calldata tokenIds,
        uint256[] calldata amounts,
        address to,
        bytes memory data
    ) external onlyRole(MINTER_ROLE) whenNotPaused nonReentrant {
        require(tokenIds.length == amounts.length, "Length mismatch");
        require(to != address(0), "Zero recipient");

        // Pull collateral for each token type
        for (uint256 i = 0; i < tokenIds.length; i++) {
            uint256 tokenId = tokenIds[i];
            uint256 amount = amounts[i];
            require(amount > 0, "Amount must be > 0");

            CardType memory card = cardTypes[tokenId];
            require(card.isActive, "Card type not active");
            require(card.collateral_token != address(0), "Card type not registered");

            uint256 totalCollateral = card.denomination * amount;
            IERC20(card.collateral_token).safeTransferFrom(msg.sender, address(this), totalCollateral);
        }

        // Mint all NFTs (batch mint)
        _mintBatch(to, tokenIds, amounts, data);

        emit CardsMinted(msg.sender, tokenIds, amounts, to);
    }

    // ------------------- Redemption -------------------

    /**
     * @notice Burns the ERC-1155 gift card(s) and transfers the underlying collateral to the specified address.
     * @param tokenId The token ID to redeem.
     * @param to Recipient of the collateral.
     * @param amount Number of NFTs to redeem.
     */
    function redeemCard(uint256 tokenId, address to, uint256 amount)
        external
        whenNotPaused
        nonReentrant
    {
        require(amount > 0, "Amount must be > 0");
        require(to != address(0), "Cannot redeem to zero address");
        require(balanceOf(msg.sender, tokenId) >= amount, "Insufficient balance");

        CardType memory card = cardTypes[tokenId];
        require(card.collateral_token != address(0), "Invalid card type");

        uint256 totalCollateral = card.denomination * amount;

        // Burn NFT(s) (Checks-Effects-Interactions pattern)
        _burn(msg.sender, tokenId, amount);

        // Transfer collateral to user
        IERC20(card.collateral_token).safeTransfer(to, totalCollateral);

        emit CardRedeemed(tokenId, msg.sender, to, amount);
    }

    // ------------------- Admin Controls -------------------

    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }

    // Override required by Solidity
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC1155, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}