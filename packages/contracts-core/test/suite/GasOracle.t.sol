// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Random, MessagingBaseTest} from "./base/MessagingBase.t.sol";
import {GasOracle, SynapseTest} from "../utils/SynapseTest.t.sol";

import {RawGasData, RawGasData256} from "../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleTest is MessagingBaseTest {
    // Deploy Production version of GasOracle and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_GAS_ORACLE) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = random.nextUint32();
        address caller = random.nextAddress();
        address destination_ = random.nextAddress();
        GasOracle cleanContract = new GasOracle(domain, destination_);
        vm.prank(caller);
        cleanContract.initialize();
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.destination(), destination_, "!destination");
    }

    function initializeLocalContract() public override {
        GasOracle(localContract()).initialize();
    }

    function test_setGasData_revert_notOwner(address caller) public {
        vm.assume(caller != GasOracle(gasOracle).owner());
        expectRevertNotOwner();
        vm.prank(caller);
        GasOracle(gasOracle).setGasData(0, 0, 0, 0, 0, 0, 0);
    }

    function test_getGasData(RawGasData256 memory rgd256) public {
        GasOracle(gasOracle).setGasData({
            domain: localDomain(),
            gasPrice: rgd256.gasPrice,
            dataPrice: rgd256.dataPrice,
            execBuffer: rgd256.execBuffer,
            amortAttCost: rgd256.amortAttCost,
            etherPrice: rgd256.etherPrice,
            markup: rgd256.markup
        });
        RawGasData memory rgd = rgd256.compress();
        assertEq(GasOracle(gasOracle).getGasData(), rgd.encodeGasData());
    }

    function test_getDecodedData(uint32 domain, RawGasData256 memory rgd256) public {
        GasOracle(gasOracle).setGasData({
            domain: domain,
            gasPrice: rgd256.gasPrice,
            dataPrice: rgd256.dataPrice,
            execBuffer: rgd256.execBuffer,
            amortAttCost: rgd256.amortAttCost,
            etherPrice: rgd256.etherPrice,
            markup: rgd256.markup
        });
        RawGasData256 memory expected = rgd256.compress().decompress();
        (
            uint256 gasPrice,
            uint256 dataPrice,
            uint256 execBuffer,
            uint256 amortAttCost,
            uint256 etherPrice,
            uint256 markup
        ) = GasOracle(gasOracle).getDecodedGasData(domain);
        assertEq(gasPrice, expected.gasPrice, "!gasPrice");
        assertEq(dataPrice, expected.dataPrice, "!dataPrice");
        assertEq(execBuffer, expected.execBuffer, "!execBuffer");
        assertEq(amortAttCost, expected.amortAttCost, "!amortAttCost");
        assertEq(etherPrice, expected.etherPrice, "!etherPrice");
        assertEq(markup, expected.markup, "!markup");
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }

    /// @notice Returns address of the tested contract
    function localContract() public view override returns (address) {
        return gasOracle;
    }
}
