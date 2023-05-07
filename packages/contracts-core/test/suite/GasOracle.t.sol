// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {GasData, InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";
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
        setGasData(localDomain(), rgd256);
        RawGasData memory rgd = rgd256.compress();
        assertEq(GasOracle(gasOracle).getGasData(), rgd.encodeGasData());
    }

    function test_getDecodedData(uint32 domain, RawGasData256 memory rgd256) public {
        setGasData(domain, rgd256);
        RawGasData256 memory expected = rgd256.compress().decompress();
        checkGasData(domain, expected);
    }

    function test_updateGasData_beforeIncreaseOptimisticPeriod(Random memory random, uint256 timePassed) public {
        uint32 domain = random.nextUint32();
        vm.assume(domain != 0 && domain != localDomain());
        timePassed = timePassed % GasOracle(gasOracle).GAS_DATA_INCREASED_OPTIMISTIC_PERIOD();
        RawGasData memory current = random.nextGasData();
        setGasData(domain, current.decompress());
        RawGasData memory updated = random.nextGasData();
        // Force getGasData(domain) to return (updated, timePassed)
        vm.mockCall(
            destination,
            abi.encodeWithSelector(InterfaceDestination.getGasData.selector, domain),
            abi.encode(updated.castToGasData(), timePassed)
        );
        address caller = random.nextAddress();
        vm.prank(caller);
        GasOracle(gasOracle).updateGasData(domain);
        checkGasData(domain, current.decompress());
    }

    function test_updateGasData_parametersOnlyIncreased(Random memory random, uint256 timePassed) public {
        uint32 domain = random.nextUint32();
        vm.assume(domain != 0 && domain != localDomain());
        timePassed = bound(
            timePassed,
            GasOracle(gasOracle).GAS_DATA_INCREASED_OPTIMISTIC_PERIOD(),
            GasOracle(gasOracle).GAS_DATA_DECREASED_OPTIMISTIC_PERIOD() - 1
        );
        RawGasData memory current = random.nextGasData();
        setGasData(domain, current.decompress());
        RawGasData memory updated = random.nextGasData();
        // Force getGasData(domain) to return (updated, timePassed)
        vm.mockCall(
            destination,
            abi.encodeWithSelector(InterfaceDestination.getGasData.selector, domain),
            abi.encode(updated.castToGasData(), timePassed)
        );
        address caller = random.nextAddress();
        vm.prank(caller);
        GasOracle(gasOracle).updateGasData(domain);
        RawGasData256 memory cur256 = current.decompress();
        RawGasData256 memory upd256 = updated.decompress();
        RawGasData256 memory expected = RawGasData256({
            gasPrice: max(cur256.gasPrice, upd256.gasPrice),
            dataPrice: max(cur256.dataPrice, upd256.dataPrice),
            execBuffer: max(cur256.execBuffer, upd256.execBuffer),
            amortAttCost: max(cur256.amortAttCost, upd256.amortAttCost),
            etherPrice: max(cur256.etherPrice, upd256.etherPrice),
            markup: max(cur256.markup, upd256.markup)
        });
        checkGasData(domain, expected);
    }

    function test_updateGasData_fullyUpdated(Random memory random, uint256 timePassed) public {
        uint32 domain = random.nextUint32();
        vm.assume(domain != 0 && domain != localDomain());
        timePassed = bound(timePassed, GasOracle(gasOracle).GAS_DATA_DECREASED_OPTIMISTIC_PERIOD(), 30 days);
        RawGasData memory current = random.nextGasData();
        setGasData(domain, current.decompress());
        RawGasData memory updated = random.nextGasData();
        // Force getGasData(domain) to return (updated, timePassed)
        vm.mockCall(
            destination,
            abi.encodeWithSelector(InterfaceDestination.getGasData.selector, domain),
            abi.encode(updated.castToGasData(), timePassed)
        );
        address caller = random.nextAddress();
        vm.prank(caller);
        GasOracle(gasOracle).updateGasData(domain);
        checkGasData(domain, updated.decompress());
    }

    function test_updateGasData_localDomain(Random memory random, uint256 timePassed) public {
        uint32 domain = localDomain();
        timePassed = timePassed % (2 * GasOracle(gasOracle).GAS_DATA_DECREASED_OPTIMISTIC_PERIOD());
        RawGasData memory current = random.nextGasData();
        setGasData(domain, current.decompress());
        RawGasData memory updated = random.nextGasData();
        // Force getGasData(domain) to return (updated, timePassed). Note that the current Destination
        // implementation will always return zero values for local domain, but it's good to have the
        // extra check in gas oracle as well.
        vm.mockCall(
            destination,
            abi.encodeWithSelector(InterfaceDestination.getGasData.selector, domain),
            abi.encode(updated.castToGasData(), timePassed)
        );
        address caller = random.nextAddress();
        vm.prank(caller);
        GasOracle(gasOracle).updateGasData(domain);
        // Should always return current values
        checkGasData(domain, current.decompress());
    }

    function setGasData(uint32 domain, RawGasData256 memory rgd256) public {
        GasOracle(gasOracle).setGasData({
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
        ) = GasOracle(gasOracle).getDecodedGasData(domain);
        assertEq(gasPrice, expected.gasPrice, "!gasPrice");
        assertEq(dataPrice, expected.dataPrice, "!dataPrice");
        assertEq(execBuffer, expected.execBuffer, "!execBuffer");
        assertEq(amortAttCost, expected.amortAttCost, "!amortAttCost");
        assertEq(etherPrice, expected.etherPrice, "!etherPrice");
        assertEq(markup, expected.markup, "!markup");
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

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
}
