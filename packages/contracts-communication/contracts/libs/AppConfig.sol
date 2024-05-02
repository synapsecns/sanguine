// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {VersionedPayloadLib} from "./VersionedPayload.sol";

// TODO: all of these could fit into a single 32 bytes slot
struct AppConfigV1 {
    uint256 requiredResponses;
    uint256 optimisticPeriod;
    uint256 guardFlag;
    address guard;
}

using AppConfigLib for AppConfigV1 global;

/// @dev Signals that the app opted out of using any Guard module.
uint8 constant APP_CONFIG_GUARD_DISABLED = 0;
/// @dev Signals that the app uses the default Guard module provided by InterchainClient contract.
uint8 constant APP_CONFIG_GUARD_DEFAULT = 1;
/// @dev Signals that the app uses a custom Guard module.
uint8 constant APP_CONFIG_GUARD_CUSTOM = 2;

library AppConfigLib {
    using VersionedPayloadLib for bytes;

    uint16 internal constant APP_CONFIG_V1 = 1;

    error AppConfigLib__IncorrectVersion(uint16 version);

    /// @notice Decodes app config (V1 or higher) from a bytes format back into an AppConfigV1 struct.
    /// @param data         The app config data in bytes format.
    function decodeAppConfigV1(bytes memory data) internal view returns (AppConfigV1 memory) {
        uint16 version = data.getVersionFromMemory();
        if (version < APP_CONFIG_V1) {
            revert AppConfigLib__IncorrectVersion(version);
        }
        // Structs of the same version will always be decoded correctly.
        // Following versions will be decoded correctly if they have the same fields as the previous version,
        // and new fields at the end: abi.decode ignores the extra bytes in the decoded payload.
        return abi.decode(data.getPayloadFromMemory(), (AppConfigV1));
    }

    /// @notice Encodes V1 app config into a bytes format.
    /// @param appConfig    The AppConfigV1 to encode.
    function encodeAppConfigV1(AppConfigV1 memory appConfig) internal pure returns (bytes memory) {
        return VersionedPayloadLib.encodeVersionedPayload(APP_CONFIG_V1, abi.encode(appConfig));
    }
}
