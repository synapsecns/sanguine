// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Random, MessagingBaseTest} from "./base/MessagingBase.t.sol";
import {GasOracle, SynapseTest} from "../utils/SynapseTest.t.sol";

import {RawGasData256} from "../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleTest is MessagingBaseTest {
    // Deploy Production version of GasOracle and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_GAS_ORACLE) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        vm.chainId(domain);
        address caller = random.nextAddress();
        address destination_ = random.nextAddress();
        GasOracle cleanContract = new GasOracle(DOMAIN_SYNAPSE, destination_);
        vm.prank(caller);
        cleanContract.initialize();
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.destination(), destination_, "!destination");
    }

    function test_constructor_revert_chainIdOverflow() public {
        vm.chainId(2 ** 32);
        vm.expectRevert("SafeCast: value doesn't fit in 32 bits");
        new GasOracle({synapseDomain_: 1, destination_: address(2)});
    }

    function initializeLocalContract() public override {
        testedGO().initialize();
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function setGasData(uint32 domain, RawGasData256 memory rgd256) public {
        testedGO().setGasData({
            domain: domain,
            gasPrice: rgd256.gasPrice,
            dataPrice: rgd256.dataPrice,
            execBuffer: rgd256.execBuffer,
            amortAttCost: rgd256.amortAttCost,
            etherPrice: rgd256.etherPrice,
            markup: rgd256.markup
        });
    }

    function checkGasData(uint32 domain, RawGasData256 memory expected) public {
        (
            uint256 gasPrice,
            uint256 dataPrice,
            uint256 execBuffer,
            uint256 amortAttCost,
            uint256 etherPrice,
            uint256 markup
        ) = testedGO().getDecodedGasData(domain);
        assertEq(gasPrice, expected.gasPrice, "!gasPrice");
        assertEq(dataPrice, expected.dataPrice, "!dataPrice");
        assertEq(execBuffer, expected.execBuffer, "!execBuffer");
        assertEq(amortAttCost, expected.amortAttCost, "!amortAttCost");
        assertEq(etherPrice, expected.etherPrice, "!etherPrice");
        assertEq(markup, expected.markup, "!markup");
    }

    function max(uint256 a, uint256 b) public pure returns (uint256) {
        return a > b ? a : b;
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return gasOracle;
    }

    function testedGO() public view returns (GasOracle) {
        return GasOracle(gasOracle);
    }
}
