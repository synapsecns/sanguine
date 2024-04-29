// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ISynapseGasOracle} from "../../contracts/interfaces/ISynapseGasOracle.sol";

// solhint-disable no-empty-blocks
contract SynapseGasOracleMock is ISynapseGasOracle {
    function receiveRemoteGasData(uint64 srcChainId, bytes calldata data) external {}

    function getLocalGasData() external view returns (bytes memory) {}

    function convertRemoteValueToLocalUnits(
        uint64 remoteChainId,
        uint256 value
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInLocalUnits(
        uint64 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInRemoteUnits(
        uint64 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}
}
