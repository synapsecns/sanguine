// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

struct AppConfigV1 {
    uint256 requiredResponses;
    uint256 optimisticPeriod;
}

using AppConfigLib for AppConfigV1 global;

library AppConfigLib {
    error AppConfigLib__IncorrectVersion(uint8 version);

    uint8 constant APP_CONFIG_V1 = 1;

    /// @notice Encodes versioned app config into a bytes format.
    /// @param version      The version of the app config.
    /// @param appConfig    The app config to encode.
    function encodeVersionedAppConfig(uint8 version, bytes memory appConfig) internal pure returns (bytes memory) {
        return abi.encode(version, appConfig);
    }

    /// @notice Decodes versioned app config from a bytes format back into a version and app config.
    /// @param data         The versioned app config data in bytes format.
    /// @return version     The version of the app config.
    /// @return appConfig   The app config as bytes.
    function decodeVersionedAppConfig(bytes memory data)
        internal
        pure
        returns (uint8 version, bytes memory appConfig)
    {
        (version, appConfig) = abi.decode(data, (uint8, bytes));
    }

    /// @notice Encodes V1 app config into a bytes format.
    /// @param appConfig    The AppConfigV1 to encode.
    function encodeAppConfigV1(AppConfigV1 memory appConfig) internal pure returns (bytes memory) {
        return encodeVersionedAppConfig(APP_CONFIG_V1, abi.encode(appConfig));
    }

    /// @notice Decodes app config (V1 or higher) from a bytes format back into an AppConfigV1 struct.
    /// @param data         The app config data in bytes format.
    function decodeAppConfigV1(bytes memory data) internal pure returns (AppConfigV1 memory) {
        (uint8 version, bytes memory appConfig) = decodeVersionedAppConfig(data);
        if (version < APP_CONFIG_V1) {
            revert AppConfigLib__IncorrectVersion(version);
        }
        // Structs of the same version will always be decoded correctly.
        // Following versions will be decoded correctly if they have the same fields as the previous version,
        // and new fields at the end: abi.decode ignores the extra bytes in the decoded payload.
        return abi.decode(appConfig, (AppConfigV1));
    }
}
