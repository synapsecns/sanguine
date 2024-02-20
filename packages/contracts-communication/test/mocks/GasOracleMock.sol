// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IGasOracle} from "../../contracts/interfaces/IGasOracle.sol";

contract GasOracleMock is IGasOracle {
    function convertRemoteValueToLocalUnits(
        uint256 remoteChainId,
        uint256 value
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInLocalUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInRemoteUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}
}
