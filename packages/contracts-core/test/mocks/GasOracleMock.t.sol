// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceGasOracle} from "../../contracts/interfaces/InterfaceGasOracle.sol";
import {BaseMock} from "./base/BaseMock.t.sol";

contract GasOracleMock is BaseMock, InterfaceGasOracle {
    function getGasData() external view returns (uint256 paddedGasData) {
        return getReturnValueUint();
    }

    function getMinimumTips(uint32, uint256, uint256) external view returns (uint256 paddedTips) {
        return getReturnValueUint();
    }
}
