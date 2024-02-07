// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";
import {SummitTipTooHigh} from "../../contracts/libs/Errors.sol";

import {RawGasData, RawGasData256} from "../utils/libs/SynapseStructs.t.sol";
import {Random, GasOracle, GasOracleTest} from "./GasOracle.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract GasOracleGasDataTest is GasOracleTest {
    function test_setGasData_revert_notOwner(address caller) public {
        vm.assume(caller != GasOracle(gasOracle).owner());
        expectRevertNotOwner();
        vm.prank(caller);
        GasOracle(gasOracle).setGasData(0, 0, 0, 0, 0, 0, 0);
    }

    function test_setSummitTip_revert_notOwner(address caller) public {
        vm.assume(caller != GasOracle(gasOracle).owner());
        expectRevertNotOwner();
        vm.prank(caller);
        GasOracle(gasOracle).setSummitTip(0);
    }

    function test_setSummitTip_allowsUpperBound() public {
        GasOracle(gasOracle).setSummitTip(0.01 ether);
        assertEq(GasOracle(gasOracle).summitTipWei(), 0.01 ether);
    }

    function test_setSummitTip_emitsEvent() public {
        vm.expectEmit(gasOracle);
        emit SummitTipUpdated(1337);
        GasOracle(gasOracle).setSummitTip(1337);
    }

    function test_setSummitTip_revert_higherThanUpperBound() public {
        vm.expectRevert(SummitTipTooHigh.selector);
        GasOracle(gasOracle).setSummitTip(0.01 ether + 1 wei);
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
}
