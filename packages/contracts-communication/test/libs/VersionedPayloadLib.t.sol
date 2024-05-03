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

    function test_encodeVersionedPayload_roundtrip(uint16 version, bytes memory payload) public view {
        bytes memory versionedPayload = libHarness.encodeVersionedPayload(version, payload);
        assertEq(libHarness.getVersion(versionedPayload), version);
        assertEq(libHarness.getPayload(versionedPayload), payload);
    }

    function test_encodeVersionedPayloadFromMemory_roundtrip(uint16 version, bytes memory payload) public view {
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

    // Tests to check that surrounding memory is preserved, and has no impact on the payload.
    function test_decodeVersionedPayload_roundtrip_surrounded(
        bytes memory a,
        bytes memory b,
        uint16 version,
        bytes memory payload,
        bytes memory c,
        bytes memory d
    )
        public
        view
    {
        bytes memory versionedPayload = libHarness.encodeVersionedPayload(version, payload);
        (bytes memory a_, bytes memory b_, uint16 version_, bytes memory payload_, bytes memory c_, bytes memory d_) =
            libHarness.decodePayloadSurrounded(a, b, versionedPayload, c, d);
        assertEq(a_, a);
        assertEq(b_, b);
        assertEq(version_, version);
        assertEq(payload_, payload);
        assertEq(c_, c);
        assertEq(d_, d);
    }

    function test_decodeVersionedPayloadFromMemory_roundtrip_surrounded(
        bytes memory a,
        bytes memory b,
        uint16 version,
        bytes memory payload,
        bytes memory c,
        bytes memory d
    )
        public
        view
    {
        bytes memory versionedPayload = libHarness.encodeVersionedPayload(version, payload);
        (bytes memory a_, bytes memory b_, uint16 version_, bytes memory payload_, bytes memory c_, bytes memory d_) =
            libHarness.decodePayloadFromMemorySurrounded(a, b, versionedPayload, c, d);
        assertEq(a_, a);
        assertEq(b_, b);
        assertEq(version_, version);
        assertEq(payload_, payload);
        assertEq(c_, c);
        assertEq(d_, d);
    }
}
