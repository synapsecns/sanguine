// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {VersionedPayloadLibHarness, VersionedPayloadLib} from "../harnesses/VersionedPayloadLibHarness.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract VersionedPayloadLibraryTest is Test {
    VersionedPayloadLibHarness public libHarness;

    function setUp() public {
        libHarness = new VersionedPayloadLibHarness();
    }

    function expectRevertTooShort(bytes memory versionedPayload) internal {
        vm.expectRevert(
            abi.encodeWithSelector(VersionedPayloadLib.VersionedPayload__TooShort.selector, versionedPayload)
        );
    }

    function makeInvalid(bytes calldata data) internal pure returns (bytes memory) {
        uint256 length = data.length % VersionedPayloadLib.VERSION_LENGTH;
        return data[:length];
    }

    function test_encodeVersionedPayload_roundtrip(uint16 version, bytes memory payload) public {
        bytes memory versionedPayload = libHarness.encodeVersionedPayload(version, payload);
        assertEq(libHarness.getVersion(versionedPayload), version);
        assertEq(libHarness.getPayload(versionedPayload), payload);
    }

    function test_encodeVersionedPayloadFromMemory_roundtrip(uint16 version, bytes memory payload) public {
        bytes memory versionedPayload = VersionedPayloadLib.encodeVersionedPayload(version, payload);
        assertEq(libHarness.getVersionFromMemory(versionedPayload), version);
        assertEq(libHarness.getPayloadFromMemory(versionedPayload), payload);
    }

    function test_getVersion_revert_tooShort(bytes calldata data) public {
        bytes memory invalidPayload = makeInvalid(data);
        expectRevertTooShort(invalidPayload);
        libHarness.getVersion(invalidPayload);
    }

    function test_getPayload_revert_tooShort(bytes calldata data) public {
        bytes memory invalidPayload = makeInvalid(data);
        expectRevertTooShort(invalidPayload);
        libHarness.getPayload(invalidPayload);
    }

    function test_getVersionFromMemory_revert_tooShort(bytes calldata data) public {
        bytes memory invalidPayload = makeInvalid(data);
        expectRevertTooShort(invalidPayload);
        libHarness.getVersionFromMemory(invalidPayload);
    }

    function test_getPayloadFromMemory_revert_tooShort(bytes calldata data) public {
        bytes memory invalidPayload = makeInvalid(data);
        expectRevertTooShort(invalidPayload);
        libHarness.getPayloadFromMemory(invalidPayload);
    }
}