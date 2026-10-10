// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC1155} from "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title GlintDough
 * @author Glint Protocol
 * @notice GlintDough is an ERC-1155 multi-token contract that issues
 *         collateral-backed crypto gift cards (vouchers). Each voucher type
 *         is backed 1:1 by a specific amount of a registered ERC-20 collateral
 *         token. When a voucher is minted, the corresponding collateral is
 *         locked inside this contract; when redeemed, the voucher is burned
 *         and the collateral is released to the redeemer.
 *
 * @dev This contract combines the ERC-1155 token standard with on-chain
 *      collateral management and role-based access control. It is the
 *      "Dough" variant of the Glint family, serving as the core, general-purpose
 *      implementation for collateral-backed gift cards.
 *
 *      Key features:
 *      - Collateral token registry: Admins can register ERC-20 tokens as
 *        collateral by providing a string key, the token contract address,
 *        symbol, and decimals. Tokens can be activated/deactivated.
 *      - Voucher type management: Admins create voucher types (ERC-1155 token IDs)
 *        by specifying a collateral token, a denomination (amount of collateral
 *        per voucher), and an optional creation fee (also in collateral tokens).
 *        Voucher types can be toggled active/inactive.
 *      - Minting: Accounts with MINTER_ROLE can mint single voucher types or
 *        batches of different voucher types. The minter must have approved this
 *        contract to transfer the total collateral plus creation fees. Fees are
 *        forwarded to a configurable fees address; collateral is locked here.
 *      - Redemption: Voucher holders can redeem their vouchers, burning them and
 *        receiving the underlying collateral. Redemption can be directed to any
 *        address.
 *      - Security: Inherits Pausable (admin can pause minting) and ReentrancyGuard
 *        to prevent reentrancy attacks. Uses SafeERC20 for safe ERC-20 transfers.
 *      - Access control: DEFAULT_ADMIN_ROLE manages collateral tokens, voucher
 *        types, fees address, pausing, and role grants. MINTER_ROLE is required
 *        for minting.
 *
 *      Collateral token metadata (decimals, address, symbol, active status)
 *      is stored directly in this contract via the `collateralTokens` mapping
 *      and the reverse lookup `keyBySmartContract`. Voucher type data is stored
 *      in the `vouchers` mapping.
 */
contract GlintDough is ERC1155, AccessControl, Pausable, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    // ------------------- Collateral Token Metadata  -------------------

    struct TokenData {
        uint8 decimals;
        address smartContract;
        string symbol;
        bool isActive;
    }

    struct Voucher {
        string token;
        address collateral_token;
        uint256 denomination;
        uint256 creationFees;
        bool is_active;
    }

    /// @notice Collateral token metadata by string key.
    mapping(string => TokenData) public collateralTokens;

    /// @notice Reverse lookup: ERC-20 address => collateral token key.
    mapping(address => string) private keyBySmartContract;

    /// @dev Recipient of creation fees collected on mint.
    address private feesAddress;

    /// @notice Auto-incrementing tracker for new voucher types.
    uint256 public nextTokenId = 1;

    /// @notice Voucher type data by ERC-1155 token ID.
    mapping(uint256 => Voucher) public vouchers;

    // ------------------- Events -------------------

    // Formerly from `Token`
    event TokenAdded(
        string indexed key,
        address indexed smartContract,
        string symbol,
        uint8 decimals
    );
    event TokenStatusChanged(string indexed key, bool isActive);

    event VoucherAdded(uint256 indexed tokenId, address indexed token, uint256 denomination);
    event VoucherToggled(uint256 indexed tokenId, bool status);
    event VouchersMintedBatch(
        address indexed minter,
        uint256[] tokenIds,
        uint256[] quantities,
        address indexed to
    );
    event VoucherMintedSingle(
        address indexed minter,
        uint256 indexed tokenId,
        uint256 quantity,
        address indexed to
    );
    event VoucherRedeemed(
        uint256 indexed tokenId,
        address indexed redeemer,
        address to,
        uint256 quantity
    );
    event PhysicalCardRegistered(bytes32 indexed serialHash, bytes32 checksum);
    event MinterRoleGranted(address minter);
    event FeesAddressUpdated(address indexed feesAddress);

    constructor(string memory uri_, address creationfeesAddress) ERC1155(uri_) {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        feesAddress = creationfeesAddress;
    }

    // ------------------- Admin: Fees Address -------------------

    /**
     * @notice Sets the address that receives creation fees.
     * @dev Only callable by DEFAULT_ADMIN_ROLE.
     */
    function setFeesAddress(
        address newFeesAddress
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(newFeesAddress != address(0), "Zero fees address");
        feesAddress = newFeesAddress;
        emit FeesAddressUpdated(newFeesAddress);
    }

    // ------------------- Admin: Collateral Token Management -------------------

    /**
     * @notice Registers a new collateral token.
     * @dev Only callable by DEFAULT_ADMIN_ROLE. Reverts if the key or the
     *      underlying ERC-20 address is already registered.
     */
    function addCollateralToken(
        string memory key,
        address smartContract,
        string memory symbol,
        uint8 decimals
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(smartContract != address(0), "Token: invalid address");
        require(
            collateralTokens[key].smartContract == address(0),
            "Token: key already exists"
        );
        require(
            bytes(keyBySmartContract[smartContract]).length == 0,
            "Token: address already registered"
        );

        collateralTokens[key] = TokenData({
            decimals: decimals,
            smartContract: smartContract,
            symbol: symbol,
            isActive: true
        });

        keyBySmartContract[smartContract] = key;

        emit TokenAdded(key, smartContract, symbol, decimals);
    }

    /**
     * @notice Deactivates a collateral token so it can no longer back new voucher types.
     * @dev Only callable by DEFAULT_ADMIN_ROLE.
     */
    function deactivateCollateralToken(
        string memory key
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        require(
            collateralTokens[key].smartContract != address(0),
            "Token: not found"
        );
        collateralTokens[key].isActive = false;
        emit TokenStatusChanged(key, false);
    }

    /**
     * @notice Reactivates a previously deactivated collateral token.
     * @dev Only callable by DEFAULT_ADMIN_ROLE.
     */
    function activateCollateralToken(
        string memory key
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        TokenData storage t = collateralTokens[key];
        require(t.smartContract != address(0), "Token: not found");
        require(!t.isActive, "Token: already active");
        t.isActive = true;
        emit TokenStatusChanged(key, true);
    }

    /**
     * @notice Retrieves metadata for a registered collateral token by key.
     */
    function getToken(
        string memory key
    ) external view returns (TokenData memory) {
        TokenData memory t = collateralTokens[key];
        require(t.smartContract != address(0), "Token: not found");
        return t;
    }

    /**
     * @notice Retrieves metadata for a registered collateral token by its
     *         smart contract address.
     */
    function getTokenByAddress(
        address smartContract
    ) external view returns (TokenData memory) {
        string memory key = keyBySmartContract[smartContract];
        TokenData memory t = collateralTokens[key];
        require(t.smartContract != address(0), "Token: not found");
        return t;
    }

    /**
     * @notice Checks whether a collateral token is registered by key.
     */
    function exists(string memory key) external view returns (bool) {
        return collateralTokens[key].smartContract != address(0);
    }

    /**
     * @notice Checks whether a collateral token is registered by its
     *         smart contract address.
     */
    function existsByAddress(
        address smartContract
    ) external view returns (bool) {
        return bytes(keyBySmartContract[smartContract]).length != 0;
    }

    // ------------------- Admin: Voucher Management -------------------

    /**
     * @notice Registers a new gift card voucher type and its underlying collateral.
     * @dev Only callable by DEFAULT_ADMIN_ROLE. The token must already be registered
     *      and active in the Token map via `addCollateralToken`. The tokenId is
     *      assigned automatically (incremental) and returned.
     * @param token        String key of the registered collateral token.
     * @param denomination Collateral amount per voucher (in collateral token decimals).
     * @param creationFees Fee per voucher charged on mint, paid in the collateral token
     *                     (in collateral token decimals). May be zero.
     */
    function addVoucher(
        string memory token,
        uint256 denomination,
        uint256 creationFees
    ) external onlyRole(DEFAULT_ADMIN_ROLE) returns (uint256) {
        require(denomination > 0, "Denomination must be greater zero");

        TokenData memory t = collateralTokens[token];
        require(t.smartContract != address(0), "Token: not found");
        require(t.isActive, "Collateral token is deactivated");

        uint256 tokenId = nextTokenId++;

        vouchers[tokenId] = Voucher({
            token: token,
            collateral_token: t.smartContract,
            denomination: denomination,
            creationFees: creationFees,
            is_active: true
        });

        emit VoucherAdded(tokenId, t.smartContract, denomination);

        return tokenId;
    }

    /**
     * @notice Activates or deactivates an existing voucher type.
     * @dev Only callable by DEFAULT_ADMIN_ROLE.
     * @param tokenId The ERC-1155 token ID of the voucher type.
     * @param status  True to activate, false to deactivate.
     */
    function toggleVoucher(
        uint256 tokenId,
        bool status
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        Voucher storage voucher = vouchers[tokenId];
        require(voucher.collateral_token != address(0), "Voucher does not exist");

        voucher.is_active = status;

        emit VoucherToggled(tokenId, status);
    }

    // ------------------- Minting (Single) -------------------

    /**
    * @notice Mints multiple copies of a SINGLE voucher type.
    * @param tokenId  The ERC-1155 voucher type ID to mint.
    * @param quantity Number of vouchers to mint (must be > 0).
    * @param to       Recipient address.
    * @param data     Optional mint hook data.
    */
    function mint(
        uint256 tokenId,
        uint256 quantity,
        address to,
        bytes memory data
    ) external onlyRole(MINTER_ROLE) whenNotPaused nonReentrant {
        require(tokenId > 0 && tokenId < nextTokenId, "Invalid Token ID");
        require(quantity > 0, "Quantity must be > 0");
        require(to != address(0), "Zero recipient");

        Voucher storage voucher = vouchers[tokenId];
        require(voucher.is_active, "Voucher not active");

        uint256 totalCollateral = voucher.denomination * quantity;
        uint256 totalFees = voucher.creationFees * quantity;
        IERC20 collateral = IERC20(voucher.collateral_token);

        // msg.sender must have approved collateral + fees
        require(
            collateral.allowance(msg.sender, address(this)) >= totalCollateral + totalFees,
            "Insufficient allowance"
        );

        // Creation fees go to the fees address
        if (totalFees > 0) {
            require(feesAddress != address(0), "Fees address not set");
            collateral.safeTransferFrom(msg.sender, feesAddress, totalFees);
        }

        // Collateral is locked in this contract
        collateral.safeTransferFrom(msg.sender, address(this), totalCollateral);

        _mint(to, tokenId, quantity, data);

        emit VoucherMintedSingle(msg.sender, tokenId, quantity, to);
    }

    // ------------------- Minting (Batch) -------------------

    /**
     * @notice Mints multiple vouchers of different types in a single transaction.
     * @dev Applies the same per-item logic as `mint`: checks allowance, forwards
     *      creation fees to the fees address, and locks collateral in this contract.
     */
    function mintBatch(
        uint256[] memory tokenIds,
        uint256[] memory quantities,
        address to,
        bytes memory data
    ) external onlyRole(MINTER_ROLE) whenNotPaused nonReentrant {
        require(tokenIds.length == quantities.length, "Length mismatch");
        require(tokenIds.length > 0, "Empty batch");
        require(to != address(0), "Zero recipient");

        for (uint256 i = 0; i < tokenIds.length; i++) {
            uint256 tokenId = tokenIds[i];
            uint256 quantity = quantities[i];

            require(tokenId > 0 && tokenId < nextTokenId, "Invalid Token ID");
            require(quantity > 0, "Quantity must be > 0");

            Voucher storage voucher = vouchers[tokenId];
            require(voucher.is_active, "Voucher not active");

            uint256 totalCollateral = voucher.denomination * quantity;
            uint256 totalFees = voucher.creationFees * quantity;
            IERC20 collateral = IERC20(voucher.collateral_token);

            // msg.sender must have approved collateral + fees for this line item
            require(
                collateral.allowance(msg.sender, address(this)) >= totalCollateral + totalFees,
                "Insufficient allowance"
            );

            // Creation fees go to the fees address
            if (totalFees > 0) {
                require(feesAddress != address(0), "Fees address not set");
                collateral.safeTransferFrom(msg.sender, feesAddress, totalFees);
            }

            // Collateral is locked in this contract
            collateral.safeTransferFrom(msg.sender, address(this), totalCollateral);
        }

        _mintBatch(to, tokenIds, quantities, data);

        emit VouchersMintedBatch(msg.sender, tokenIds, quantities, to);
    }

    // ------------------- Redemption -------------------

    /**
     * @notice Redeems vouchers and sends the underlying collateral to a specified address.
     */
    function redeemTo(
        uint256 tokenId,
        address to,
        uint256 quantity
    ) external nonReentrant {
        _redeemTo(msg.sender, tokenId, to, quantity);
    }

    /**
     * @notice Redeems vouchers and sends the underlying collateral to the caller.
     */
    function redeem(
        uint256 tokenId,
        uint256 quantity
    ) external nonReentrant {
        _redeemTo(msg.sender, tokenId, msg.sender, quantity);
    }

    function _redeemTo(
        address redeemer,
        uint256 tokenId,
        address to,
        uint256 quantity
    ) internal {
        require(quantity > 0, "Quantity must be > 0");
        require(to != address(0), "Cannot redeem to zero address");
        require(balanceOf(redeemer, tokenId) >= quantity, "Insufficient balance");

        Voucher memory voucher = vouchers[tokenId];
        require(voucher.is_active, "Voucher not active");
        require(voucher.collateral_token != address(0), "Invalid voucher type");

        uint256 totalCollateral = voucher.denomination * quantity;

        _burn(redeemer, tokenId, quantity);
        IERC20(voucher.collateral_token).safeTransfer(to, totalCollateral);

        emit VoucherRedeemed(tokenId, redeemer, to, quantity);
    }

    // ------------------- Admin Controls -------------------

    function pause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }

    function grantMinterRole(
        address minter
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        _grantRole(MINTER_ROLE, minter);
        emit MinterRoleGranted(minter);
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC1155, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}