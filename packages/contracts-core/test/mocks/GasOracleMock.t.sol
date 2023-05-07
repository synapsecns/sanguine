// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceGasOracle} from "../../contracts/interfaces/InterfaceGasOracle.sol";

contract GasOracleMock is InterfaceGasOracle {
    uint256 public mockedGasData;
    uint256 public mockedMinimumTips;

    function setMockedGasData(uint256 paddedGasData) external {
        mockedGasData = paddedGasData;
    }

    function setMockedMinimumTips(uint256 paddedTips) external {
        mockedMinimumTips = paddedTips;
    }

    function updateGasData(uint32 domain) external {}

    function getDecodedGasData(uint32 domain)
        external
        view
        returns (
            uint256 gasPrice,
            uint256 dataPrice,
            uint256 execBuffer,
            uint256 amortAttCost,
            uint256 etherPrice,
            uint256 markup
        )
    {}

    function getGasData() external view returns (uint256 paddedGasData) {
        return mockedGasData;
    }

    function getMinimumTips(uint32, uint256, uint256) external view returns (uint256 paddedTips) {
        return mockedMinimumTips;
    }
}
