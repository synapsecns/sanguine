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

    function getGasData() external view returns (uint256 paddedGasData) {
        return mockedGasData;
    }

    function getMinimumTips(uint32, uint256, uint256) external view returns (uint256 paddedTips) {
        return mockedMinimumTips;
    }
}
