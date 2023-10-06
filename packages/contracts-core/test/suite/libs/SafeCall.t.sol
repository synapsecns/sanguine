// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest} from "../../utils/SynapseLibraryTest.t.sol";

import {CalleeMock, CalleeReturnDataMock} from "../../mocks/CalleeMocks.t.sol";
import {SafeCallHarness} from "../../harnesses/libs/SafeCallHarness.t.sol";

// solhint-disable func-name-mixedcase
contract SafeCallTest is SynapseLibraryTest {
    SafeCallHarness internal libHarness;

    CalleeMock internal callee;
    CalleeReturnDataMock internal calleeReturnData;
    address internal eoa;

    function setUp() public {
        libHarness = new SafeCallHarness();
        callee = new CalleeMock();
        calleeReturnData = new CalleeReturnDataMock();
        eoa = makeAddr("EOA");
        vm.label(address(callee), "callee");
        vm.label(address(calleeReturnData), "calleeReturnData");
    }

    function test_safeCall() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // expectCall(address callee, uint256 msgValue, uint64 gas, bytes calldata data)
        vm.expectCall(address(callee), 0, 100_000, payload);
        // 100k gas should be enough for a simple call
        bool success = libHarness.safeCall(address(callee), 100_000, 0, payload);
        assertTrue(success);
        assertEq(callee.secret(), 42);
    }

    function test_safeCall_withReturnData() public {
        bytes memory payload = abi.encodeCall(CalleeReturnDataMock.setSecret, (42));
        // expectCall(address callee, uint256 msgValue, uint64 gas, bytes calldata data)
        vm.expectCall(address(calleeReturnData), 0, 100_000, payload);
        // 100k gas should be enough for a simple call
        bool success = libHarness.safeCall(address(calleeReturnData), 100_000, 0, payload);
        assertTrue(success);
        assertEq(calleeReturnData.secret(), 42);
    }

    function test_safeCall_withMsgValue() public {
        uint256 msgValue = 1337;
        deal(address(libHarness), msgValue);
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // expectCall(address callee, uint256 msgValue, uint64 gas, bytes calldata data)
        vm.expectCall(address(callee), msgValue, 100_000, payload);
        // 100k gas should be enough for a simple call
        bool success = libHarness.safeCall(address(callee), 100_000, msgValue, payload);
        assertTrue(success);
        assertEq(callee.secret(), 42);
        assertEq(address(callee).balance, msgValue);
    }

    function test_safeCall_withMsgValueAndReturnData() public {
        uint256 msgValue = 1337;
        deal(address(libHarness), msgValue);
        bytes memory payload = abi.encodeCall(CalleeReturnDataMock.setSecret, (42));
        // expectCall(address callee, uint256 msgValue, uint64 gas, bytes calldata data)
        vm.expectCall(address(calleeReturnData), msgValue, 100_000, payload);
        // 100k gas should be enough for a simple call
        bool success = libHarness.safeCall(address(calleeReturnData), 100_000, msgValue, payload);
        assertTrue(success);
        assertEq(calleeReturnData.secret(), 42);
        assertEq(address(calleeReturnData).balance, msgValue);
    }

    function test_safeCall_returnsFalse_onRecipientOutOfGas() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // 10k gas should not be enough for a storage write
        bool success = libHarness.safeCall(address(callee), 10_000, 0, payload);
        assertFalse(success);
    }

    function test_safeCall_returnsFalse_onRevert() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // Force the call to revert
        vm.mockCallRevert(address(callee), payload, "GM");
        bool success = libHarness.safeCall(address(callee), 100_000, 0, payload);
        assertFalse(success);
    }

    function test_safeCall_returnsFalse_recipientEOA() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        bool success = libHarness.safeCall(eoa, 100_000, 0, payload);
        assertFalse(success);
    }
}
