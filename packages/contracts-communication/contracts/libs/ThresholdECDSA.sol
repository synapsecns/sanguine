// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

struct ThresholdECDSA {
    uint256 _threshold;
    EnumerableSet.AddressSet _signers;
}

using ThresholdECDSALib for ThresholdECDSA global;

library ThresholdECDSALib {
    using EnumerableSet for EnumerableSet.AddressSet;

    error ThresholdECDSA__AlreadySigner(address account);
    error ThresholdECDSA__InvalidSignature(bytes signature);
    error ThresholdECDSA__NotEnoughSignatures(uint256 threshold);
    error ThresholdECDSA__NotSigner(address account);
    error ThresholdECDSA__RecoveredSignersNotSorted();
    error ThresholdECDSA__ZeroThreshold();

    /// @notice Adds a new signer to the list of signers.
    /// @dev Will revert if the account is already a signer.
    function addSigner(ThresholdECDSA storage self, address account) internal {
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
    /// - Any of the payloads is not a valid signature payload.
    /// - The number of signatures is less than the threshold.
    /// - The recovered list of signers is not sorted in the ascending order.
    function verifySignedHash(ThresholdECDSA storage self, bytes32 hash, bytes[] memory signatures) internal view {
        // TODO: Implement
    }
}
