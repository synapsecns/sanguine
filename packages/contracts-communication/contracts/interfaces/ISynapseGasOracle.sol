// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IGasOracle} from "./IGasOracle.sol";

interface ISynapseGasOracle is IGasOracle {
    /// @notice Allows Synapse Module to pass the gas data from a remote chain to the Gas Oracle.
    /// @dev Could only be called by Synapse Module.
    /// @param srcChainId        The chain id of the remote chain.
    /// @param data              The gas data from the remote chain.
    function receiveRemoteGasData(uint64 srcChainId, bytes calldata data) external;

    /// @notice Gets the gas data for the local chain.
    function getLocalGasData() external view returns (bytes memory);
}
