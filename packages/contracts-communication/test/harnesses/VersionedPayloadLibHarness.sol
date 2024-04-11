// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {VersionedPayloadLib} from "../../contracts/libs/VersionedPayload.sol";

// solhint-disable ordering
contract VersionedPayloadLibHarness {
    function encodeVersionedPayload(uint16 version, bytes memory payload) external pure returns (bytes memory) {
        return VersionedPayloadLib.encodeVersionedPayload(version, payload);
    }

    function getVersion(bytes calldata versionedPayload) external pure returns (uint16) {
        return VersionedPayloadLib.getVersion(versionedPayload);
    }

    function getPayload(bytes calldata versionedPayload) external pure returns (bytes calldata) {
        return VersionedPayloadLib.getPayload(versionedPayload);
    }

    function getVersionFromMemory(bytes memory versionedPayload) external pure returns (uint16) {
        return VersionedPayloadLib.getVersionFromMemory(versionedPayload);
    }

    function getPayloadFromMemory(bytes memory versionedPayload) external view returns (bytes memory) {
        return VersionedPayloadLib.getPayloadFromMemory(versionedPayload);
    }

    function decodePayloadSurrounded(
        bytes calldata a,
        bytes memory b,
        bytes calldata versionedPayload,
        bytes memory c,
        bytes calldata d
    )
        external
        pure
        returns (
            bytes memory a_,
            bytes memory b_,
            uint16 version,
            bytes memory payload,
            bytes memory c_,
            bytes memory d_
        )
    {
        a_ = a;
        b_ = b;
        version = VersionedPayloadLib.getVersion(versionedPayload);
        payload = VersionedPayloadLib.getPayload(versionedPayload);
        c_ = c;
        d_ = d;
    }

    function decodePayloadFromMemorySurrounded(
        bytes calldata a,
        bytes memory b,
        bytes memory versionedPayload,
        bytes memory c,
        bytes memory d
    )
        external
        view
        returns (
            bytes memory a_,
            bytes memory b_,
            uint16 version,
            bytes memory payload,
            bytes memory c_,
            bytes memory d_
        )
    {
        a_ = a;
        b_ = b;
        version = VersionedPayloadLib.getVersionFromMemory(versionedPayload);
        payload = VersionedPayloadLib.getPayloadFromMemory(versionedPayload);
        c_ = c;
        d_ = d;
    }
}
