// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {VersionedPayloadLib} from "../../contracts/libs/VersionedPayload.sol";

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

    function getPayloadFromMemory(bytes memory versionedPayload) external pure returns (bytes memory) {
        return VersionedPayloadLib.getPayloadFromMemory(versionedPayload);
    }
}
