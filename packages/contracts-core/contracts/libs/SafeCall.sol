// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/// @notice Library for performing safe calls to the recipients.
/// Inspired by https://github.com/nomad-xyz/ExcessivelySafeCall
library SafeCall {
    /// @notice Performs a call to the recipient using the provided gas limit, msg.value and payload
    /// in a safe way:
    /// - If the recipient is not a contract, false is returned instead of reverting.
    /// - If the recipient call reverts for any reason, false is returned instead of reverting.
    /// - If the recipient call succeeds, true is returned, and any returned data is ignored.
    /// @param recipient        The recipient of the call
    /// @param gasLimit         The gas limit to use for the call
    /// @param msgValue         The msg.value to use for the call
    /// @param payload          The payload to use for the call
    /// @return success         True if the call succeeded, false otherwise
    function safeCall(address recipient, uint256 gasLimit, uint256 msgValue, bytes memory payload)
        internal
        returns (bool success)
    {
        // Exit early if the recipient is not a contract
        if (recipient.code.length == 0) return false;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Perform the call to the recipient, while ignoring any returned data
            // call(g, a, v, in, insize, out, outsize) => returns 0 on error, 1 on success
            success := call(gasLimit, recipient, msgValue, add(payload, 0x20), mload(payload), 0, 0)
        }
    }
}
