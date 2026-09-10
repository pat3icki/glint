// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {BasePaymaster} from "@account-abstraction/contracts/core/BasePaymaster.sol";
import {PackedUserOperation} from "@account-abstraction/contracts/interfaces/PackedUserOperation.sol";
import {IEntryPoint} from "@account-abstraction/contracts/interfaces/IEntryPoint.sol";
import {_packValidationData} from "@account-abstraction/contracts/core/Helpers.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";


/**
 * @title GiftCardPaymaster
 * @dev An ERC-4337 Verifying Paymaster that sponsors transactions for gift card operations.
 *      Uses time-bound signatures signed by a trusted backend signer.
 */
contract GiftCardPaymaster is BasePaymaster {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    address public verifyingSigner;

    // Time offsets for validation data packing (as per ERC-4337)
    uint48 public constant VALID_UNTIL_OFFSET = 0;
    uint48 public constant VALID_AFTER_OFFSET = 48;
    uint256 public constant SIGNATURE_OFFSET = 96; // not used directly, kept for clarity

    event VerifyingSignerUpdated(address oldSigner, address newSigner);

    /**
     * @param _entryPoint The global ERC-4337 EntryPoint contract.
     * @param _verifyingSigner The backend address authorized to sign PaymasterData.
     */
    constructor(IEntryPoint _entryPoint, address _verifyingSigner) BasePaymaster(_entryPoint) {
        require(_verifyingSigner != address(0), "Invalid verifying signer");
        verifyingSigner = _verifyingSigner;
    }

    /**
     * @notice Updates the backend signer address in case of key rotation.
     */
    function setVerifyingSigner(address _newSigner) external onlyOwner {
        require(_newSigner != address(0), "Invalid verifying signer");
        emit VerifyingSignerUpdated(verifyingSigner, _newSigner);
        verifyingSigner = _newSigner;
    }

    /**
     * @notice Validates that the UserOperation was signed by the authorized backend.
     * @dev The backend MUST decode the userOp.callData off-chain to ensure it only interacts
     *      with the GiftCardNFT contract before generating this signature.
     *      The signature is over: keccak(userOpHash, validUntil, validAfter).
     *      Layout of paymasterAndData: [paymasterAddress(20)] [validUntil(6)] [validAfter(6)] [signature(...)]
     */
    function _validatePaymasterUserOp(
        PackedUserOperation calldata userOp,
        bytes32 userOpHash,
        uint256 /* requiredPreFund */
    ) internal view override returns (bytes memory context, uint256 validationData) {
        bytes calldata paymasterAndData = userOp.paymasterAndData;
        require(paymasterAndData.length >= 20 + 6 + 6 + 65, "Invalid paymasterAndData length");

        // Extract time bounds and signature
        uint48 validUntil = uint48(bytes6(paymasterAndData[20:26]));
        uint48 validAfter = uint48(bytes6(paymasterAndData[26:32]));
        bytes calldata signature = paymasterAndData[32:];

        // Hash the data that the signer signed
        bytes32 hash = keccak256(abi.encode(userOpHash, validUntil, validAfter));

        // Verify ECDSA signature
        bytes32 ethSignedMessageHash = hash.toEthSignedMessageHash();
        address recoveredSigner = ethSignedMessageHash.recover(signature);

        if (recoveredSigner != verifyingSigner) {
            // Return standard SIG_VALIDATION_FAILED (1)
            return ("", 1);
        }

        // Pack validation data: (authorizer == 0 for success, validUntil, validAfter)
        return ("", _packValidationData(false, validUntil, validAfter));
    }

    /**
     * @dev Accept ETH from EntryPoint or owner (e.g., for refunds or manual deposits).
     */
    receive() external payable {
        require(
            msg.sender == address(entryPoint) || msg.sender == owner(),
            "GiftCardPaymaster: unauthorized sender"
        );
    }
}