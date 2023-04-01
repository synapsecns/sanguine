// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {STATE_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {RawStateReport} from "../../utils/libs/SynapseStructs.t.sol";

import {StateFlag, StateReportHarness} from "../../harnesses/libs/StateReportHarness.t.sol";

// solhint-disable func-name-mixedcase
contract StateReportLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    StateReportHarness internal libHarness;

    function setUp() public {
        libHarness = new StateReportHarness();
    }

    function test_formatStateReport(RawStateReport memory rawSR) public {
        // Make sure flag fits into enum
        rawSR.flag = uint8(bound(rawSR.flag, 0, uint8(type(StateFlag).max)));
        // This is tested in StateLibraryTest, we assume it's working here
        bytes memory state = rawSR.state.formatState();
        bytes memory payload = libHarness.formatStateReport(StateFlag(rawSR.flag), state);
        assertEq(payload, abi.encodePacked(rawSR.flag, state), "!formatStateReport");
        checkCastToStateReport({payload: payload, isStateReport: true});
        // Check getters
        assertEq(uint8(libHarness.flag(payload)), rawSR.flag, "!flag");
        assertEq(libHarness.state(payload), state, "!state");
        // Test hashing
        bytes32 stateReportSalt = keccak256("STATE_REPORT_SALT");
        bytes32 hashedStateReport = keccak256(abi.encodePacked(stateReportSalt, keccak256(payload)));
        assertEq(libHarness.hash(payload), hashedStateReport, "!hash");
    }

    function test_isStateReport(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToStateReport({payload: payload, isStateReport: length == 1 + STATE_LENGTH});
    }

    function checkCastToStateReport(bytes memory payload, bool isStateReport) public {
        if (isStateReport) {
            assertTrue(libHarness.isStateReport(payload), "!isStateReport: when valid");
            assertEq(libHarness.castToStateReport(payload), payload, "!castToStateReport: when valid");
        } else {
            assertFalse(libHarness.isStateReport(payload), "!isStateReport: when valid");
            vm.expectRevert("Not a state report");
            libHarness.castToStateReport(payload);
        }
    }
}
