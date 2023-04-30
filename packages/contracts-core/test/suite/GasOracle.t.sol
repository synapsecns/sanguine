// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Random, MessagingBaseTest} from "./base/MessagingBase.t.sol";
import {GasOracle, SynapseTest} from "../utils/SynapseTest.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleTest is MessagingBaseTest {
    // Deploy mocks for every contract
    constructor() SynapseTest(0) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        address caller = random.nextAddress();
        GasOracle cleanContract = new GasOracle(domain);
        vm.prank(caller);
        cleanContract.initialize();
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
    }

    function initializeLocalContract() public override {
        GasOracle(localContract()).initialize();
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return gasOracle;
    }
}
