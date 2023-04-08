// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {EXECUTION_LENGTH} from "../../../contracts/libs/Constants.sol";

import {SynapseLibraryTest, TypedMemView} from "../../utils/SynapseLibraryTest.t.sol";
import {ExecutionHarness} from "../../harnesses/libs/ExecutionHarness.t.sol";

import {RawExecution} from "../../utils/libs/SynapseStructs.t.sol";

// solhint-disable func-name-mixedcase
contract ExecutionLibraryTest is SynapseLibraryTest {
    using TypedMemView for bytes;

    ExecutionHarness internal libHarness;

    function setUp() public {
        libHarness = new ExecutionHarness();
    }

    // ═════════════════════════════════════════════ TESTS: FORMATTING ═════════════════════════════════════════════════

    function test_formatExecution(RawExecution memory re) public {
        bytes memory tipsPayload = re.tips.formatTips();
        // Test formatting
        bytes memory payload = libHarness.formatExecution(
            re.origin, re.destination, re.messageHash, re.snapshotRoot, re.firstExecutor, re.finalExecutor, tipsPayload
        );
        assertEq(
            payload,
            abi.encodePacked(
                re.origin,
                re.destination,
                re.messageHash,
                re.snapshotRoot,
                re.firstExecutor,
                re.finalExecutor,
                tipsPayload
            ),
            "!formatExecution"
        );
        // Test formatting checker
        checkCastToExecution({payload: payload, isExecution: true});
        // Test getters
        assertEq(libHarness.origin(payload), re.origin, "!origin");
        assertEq(libHarness.destination(payload), re.destination, "!destination");
        assertEq(libHarness.messageHash(payload), re.messageHash, "!messageHash");
        assertEq(libHarness.snapshotRoot(payload), re.snapshotRoot, "!snapshotRoot");
        assertEq(libHarness.firstExecutor(payload), re.firstExecutor, "!firstExecutor");
        assertEq(libHarness.finalExecutor(payload), re.finalExecutor, "!finalExecutor");
        assertEq(libHarness.tips(payload), tipsPayload, "!tipsPayload");
    }

    function test_isExecution(uint8 length) public {
        bytes memory payload = new bytes(length);
        checkCastToExecution({payload: payload, isExecution: length == EXECUTION_LENGTH});
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    function checkCastToExecution(bytes memory payload, bool isExecution) public {
        if (isExecution) {
            assertTrue(libHarness.isExecution(payload), "!isExecution: when valid");
            assertEq(libHarness.castToExecution(payload), payload, "!castToExecution: when valid");
        } else {
            assertFalse(libHarness.isExecution(payload), "!isExecution: when valid");
            vm.expectRevert("Not a execution");
            libHarness.castToExecution(payload);
        }
    }
}
