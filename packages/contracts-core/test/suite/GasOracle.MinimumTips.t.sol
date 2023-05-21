// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Random, GasOracle, GasOracleTest} from "./GasOracle.t.sol";
import {RawGasData, RawGasData256} from "../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleMinimumTipsTest is GasOracleTest {
    function setUp() public override {
        super.setUp();
    }
}
