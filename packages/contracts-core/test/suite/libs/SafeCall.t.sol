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

    function test_safeCall_callsRecipient() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // 100k gas should be enough for a simple call
        assertTrue(libHarness.safeCall(address(callee), 100_000, 0, payload));
        assertEq(callee.secret(), 42);
    }

    function test_safeCall_callsRecipient_withReturnData() public {
        bytes memory payload = abi.encodeCall(CalleeReturnDataMock.setSecret, (42));
        // 100k gas should be enough for a simple call
        assertTrue(libHarness.safeCall(address(calleeReturnData), 100_000, 0, payload));
        assertEq(calleeReturnData.secret(), 42);
    }

    function test_safeCall_forwardsMsgValue() public {
        uint256 msgValue = 1337;
        deal(address(libHarness), msgValue);
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // 100k gas should be enough for a simple call
        assertTrue(libHarness.safeCall(address(callee), 100_000, msgValue, payload));
        assertEq(address(callee).balance, msgValue);
    }

    function test_safeCall_forwardsMsgValue_withReturnData() public {
        uint256 msgValue = 1337;
        deal(address(libHarness), msgValue);
        bytes memory payload = abi.encodeCall(CalleeReturnDataMock.setSecret, (42));
        // 100k gas should be enough for a simple call
        assertTrue(libHarness.safeCall(address(calleeReturnData), 100_000, msgValue, payload));
        assertEq(address(calleeReturnData).balance, msgValue);
    }

    function test_safeCall_setsGasLimit() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // 10k gas should not be enough for a storage write
        assertFalse(libHarness.safeCall(address(callee), 10_000, 0, payload));
    }

    function test_safeCall_returnsFalse_onRevert() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        // Force the call to revert
        vm.mockCallRevert(address(callee), payload, "GM");
        assertFalse(libHarness.safeCall(address(callee), 100_000, 0, payload));
    }

    function test_safeCall_returnsFalse_recipientEOA() public {
        bytes memory payload = abi.encodeCall(CalleeMock.setSecret, (42));
        assertFalse(libHarness.safeCall(eoa, 100_000, 0, payload));
    }
}
