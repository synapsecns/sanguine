// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntry} from "../libs/InterchainEntry.sol";

interface ISynapseModule {
    /// @notice Sets the address of the InterchainDB contract to be used for verifying entries.
    /// @dev This function can only be called by the contract owner.
    /// @param _interchainDB The address of the InterchainDB contract.
    function setInterchainDB(address _interchainDB) external;

    /// @notice Sets the required threshold for verification.
    /// @dev This function updates the threshold value that determines the minimum number of verifications required for an entry to be considered valid. Can only be called by the contract owner.
    /// @param _threshold The new threshold value.
    function setRequiredThreshold(uint256 _threshold) external;

    /// @notice Updates the list of verifier addresses.
    /// @dev This function sets the addresses that are allowed to act as verifiers for entries. Can only be called by the contract owner.
    /// @param _verifiers An array of addresses to be set as verifiers.
    function setVerifiers(address[] calldata _verifiers) external;

    /// @notice Requests off-chain verification of an interchain entry for a specified destination chain. This function requires a fee.
    /// @dev This function can only be called by the InterchainDB contract. It checks if the sent value covers the module fee for the requested destination chain, then proceeds to pay the fee for execution. Emits a VerificationRequested event upon success.
    /// @param destChainId The ID of the destination chain where the entry needs to be verified.
    /// @param entry The interchain entry to be verified.
    function requestVerification(uint256 destChainId, InterchainEntry memory entry) external payable;

    /// @notice Verifies an interchain entry using a set of verifier signatures.
    /// @dev This function checks if the provided signatures meet the required threshold for verification.
    /// It then calls the InterchainDB contract to verify the entry. Requires that the number of valid signatures meets or exceeds the required threshold.
    /// @param entry The interchain entry to be verified.
    /// @param signatures An array of signatures used to verify the entry.
    function verifyEntry(InterchainEntry memory entry, bytes[] calldata signatures) external;
}
