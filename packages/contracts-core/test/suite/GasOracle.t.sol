// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {GasOracle, SynapseTest} from "../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleTest is SynapseTest {
    // Deploy mocks for every contract
    constructor() SynapseTest(0) {}

    function test_initializer(uint32 domain, address caller) public {
        GasOracle oracle = new GasOracle(domain);
        vm.prank(caller);
        oracle.initialize();
        assertEq(oracle.owner(), caller, "!owner");
        assertEq(oracle.localDomain(), domain, "!localDomain");
    }

    function test_initializer_revert_alreadyInitialized() public {
        expectRevertAlreadyInitialized();
        GasOracle(gasOracle).initialize();
    }
}
