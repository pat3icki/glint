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
 * @dev ERC-4337 Verifying Paymaster with time-bound backend signatures.
 *
 *      FIXES APPLIED:
 *      - Critical: Signature no longer uses EntryPoint's `userOpHash` (which includes
 *        `paymasterAndData` and therefore the signature itself). Instead, `getHash`
 *        builds a custom hash from all critical UserOperation fields, excluding the
 *        signature and paymasterAndData. This removes the circular dependency.
 *      - High: Exact length check for `paymasterAndData` (97 bytes) prevents
 *        malformed or extra data from being silently ignored.
 *      - Medium: Removed `receive()` to prevent ETH from being locked in the contract.
 *        Fund the paymaster via `entryPoint.depositTo{value: ...}(address(this))`.
 *      - Low: Added zero-address check inside `_validatePaymasterUserOp`.
 */
contract GiftCardPaymaster is BasePaymaster {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    address public verifyingSigner;

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
     * @notice Computes the hash that the backend signer must sign.
     * @dev Excludes `paymasterAndData` and `signature` to avoid circular dependency.
     *      Includes all critical UserOperation fields plus time bounds.
     *      The backend must sign `toEthSignedMessageHash(getHash(...))`.
     */
    function getHash(
        PackedUserOperation calldata userOp,
        uint48 validUntil,
        uint48 validAfter
    ) public view returns (bytes32) {
        return keccak256(
            abi.encode(
                userOp.sender,
                userOp.nonce,
                keccak256(userOp.initCode),
                keccak256(userOp.callData),
                userOp.accountGasLimits,
                userOp.preVerificationGas,
                userOp.gasFees,
                block.chainid,
                address(this),
                validUntil,
                validAfter
            )
        );
    }

    /**
     * @notice Validates that the UserOperation was signed by the authorized backend.
     * @dev Layout of paymasterAndData:
     *      [paymasterAddress(20)] [validUntil(6)] [validAfter(6)] [signature(65)]
     */
    function _validatePaymasterUserOp(
        PackedUserOperation calldata userOp,
        bytes32, /* userOpHash */
        uint256 /* requiredPreFund */
    ) internal view override returns (bytes memory context, uint256 validationData) {
        require(verifyingSigner != address(0), "Verifying signer not set");

        bytes calldata paymasterAndData = userOp.paymasterAndData;
        // Exact length: 20 + 6 + 6 + 65 = 97 bytes
        require(paymasterAndData.length == 97, "Invalid paymasterAndData length");

        uint48 validUntil = uint48(bytes6(paymasterAndData[20:26]));
        uint48 validAfter = uint48(bytes6(paymasterAndData[26:32]));
        bytes calldata signature = paymasterAndData[32:];

        bytes32 hash = getHash(userOp, validUntil, validAfter);
        bytes32 ethSignedMessageHash = hash.toEthSignedMessageHash();
        address recoveredSigner = ethSignedMessageHash.recover(signature);

        if (recoveredSigner != verifyingSigner) {
            return ("", 1);
        }

        return ("", _packValidationData(false, validUntil, validAfter));
    }

    // No receive() function. ETH must be deposited directly to EntryPoint:
    // entryPoint.depositTo{value: msg.value}(address(this));
}