// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

struct ThresholdECDSA {
    uint256 _threshold;
    EnumerableSet.AddressSet _signers;
}

using ThresholdECDSALib for ThresholdECDSA global;

library ThresholdECDSALib {
    using EnumerableSet for EnumerableSet.AddressSet;

    error ThresholdECDSA__AlreadySigner(address account);
    error ThresholdECDSA__IncorrectSignaturesLength(uint256 length);
    error ThresholdECDSA__InvalidSignature(bytes signature);
    error ThresholdECDSA__NotEnoughSignatures(uint256 threshold);
    error ThresholdECDSA__NotSigner(address account);
    error ThresholdECDSA__RecoveredSignersNotSorted();
    error ThresholdECDSA__ZeroAddress();
    error ThresholdECDSA__ZeroThreshold();

    uint256 private constant SIGNATURE_LENGTH = 65;

    /// @notice Adds a new signer to the list of signers.
    /// @dev Will revert if the account is already a signer.
    function addSigner(ThresholdECDSA storage self, address account) internal {
        if (account == address(0)) revert ThresholdECDSA__ZeroAddress();
        bool added = self._signers.add(account);
        if (!added) {
            revert ThresholdECDSA__AlreadySigner(account);
        }
    }

    /// @notice Removes a signer from the list of signers.
    /// @dev Will revert if the account is not a signer.
    function removeSigner(ThresholdECDSA storage self, address account) internal {
        bool removed = self._signers.remove(account);
        if (!removed) {
            revert ThresholdECDSA__NotSigner(account);
        }
    }

    /// @notice Modifies the threshold of signatures required.
    function modifyThreshold(ThresholdECDSA storage self, uint256 threshold) internal {
        if (threshold == 0) {
            revert ThresholdECDSA__ZeroThreshold();
        }
        self._threshold = threshold;
    }

    /// @notice Checks if the account is a signer.
    function isSigner(ThresholdECDSA storage self, address account) internal view returns (bool) {
        return self._signers.contains(account);
    }

    /// @notice Gets the full list of signers.
    function getSigners(ThresholdECDSA storage self) internal view returns (address[] memory) {
        return self._signers.values();
    }

    /// @notice Gets the threshold of signatures required.
    function getThreshold(ThresholdECDSA storage self) internal view returns (uint256) {
        return self._threshold;
    }

    /// @notice Verifies that the number of signatures is greater than or equal to the threshold.
    /// Note: the list of signers recovered from the signatures is required to be sorted in ascending order.
    /// @dev Will revert if either of the conditions is met:
    /// - Threshold is not configured.
    /// - Any of the payloads is not a valid signature payload.
    /// - The number of signatures is less than the threshold.
    /// - The recovered list of signers is not sorted in the ascending order.
    function verifySignedHash(ThresholdECDSA storage self, bytes32 hash, bytes calldata signatures) internal view {
        // Figure out the signaturesAmount of signatures provided
        uint256 signaturesAmount = signatures.length / SIGNATURE_LENGTH;
        if (signaturesAmount == 0 || signaturesAmount * SIGNATURE_LENGTH != signatures.length) {
            revert ThresholdECDSA__IncorrectSignaturesLength(signatures.length);
        }
        // First, check that threshold is configured and enough signatures are provided
        uint256 threshold = self._threshold;
        if (threshold == 0) {
            revert ThresholdECDSA__ZeroThreshold();
        }
        if (signaturesAmount < threshold) {
            revert ThresholdECDSA__NotEnoughSignatures(threshold);
        }
        uint256 offset = 0;
        uint256 validSignatures = 0;
        address lastSigner = address(0);
        for (uint256 i = 0; i < signaturesAmount; ++i) {
            bytes memory signature = signatures[offset:offset + SIGNATURE_LENGTH];
            (address recovered, ECDSA.RecoverError error,) = ECDSA.tryRecover(hash, signature);
            if (error != ECDSA.RecoverError.NoError) {
                revert ThresholdECDSA__InvalidSignature(signature);
            }
            // Check that the recovered addresses list is strictly increasing
            if (recovered <= lastSigner) {
                revert ThresholdECDSA__RecoveredSignersNotSorted();
            }
            lastSigner = recovered;
            // Since the signers list is sorted, every time we find a valid signer it's not a duplicate
            if (isSigner(self, recovered)) {
                validSignatures += 1;
            }
            offset += SIGNATURE_LENGTH;
        }
        if (validSignatures < threshold) {
            revert ThresholdECDSA__NotEnoughSignatures(threshold);
        }
    }
}
