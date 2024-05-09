// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {AppConfigV1} from "../libs/AppConfig.sol";

import {IInterchainApp} from "./IInterchainApp.sol";

interface IInterchainAppV1 is IInterchainApp {
    error InterchainApp__AppConfigInvalid(uint256 requiredResponses, uint256 optimisticPeriod);
    error InterchainApp__LinkedAppNotEVM(bytes32 linkedApp);
    error InterchainApp__ModuleAlreadyAdded(address module);
    error InterchainApp__ModuleNotAdded(address module);
    error InterchainApp__ModuleZeroAddress();
    error InterchainApp__RemoteAppZeroAddress();

    function addInterchainClient(address client, bool updateLatest) external;
    function removeInterchainClient(address client) external;
    function setLatestInterchainClient(address client) external;

    function linkRemoteApp(uint64 chainId, bytes32 remoteApp) external;
    function linkRemoteAppEVM(uint64 chainId, address remoteApp) external;

    function addTrustedModule(address module) external;
    function removeTrustedModule(address module) external;
    function setAppConfigV1(uint256 requiredResponses, uint256 optimisticPeriod) external;
    function setExecutionService(address executionService) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function getAppConfigV1() external view returns (AppConfigV1 memory);
    function getExecutionService() external view returns (address);
    function getInterchainClients() external view returns (address[] memory);
    function getLatestInterchainClient() external view returns (address);
    function getLinkedApp(uint64 chainId) external view returns (bytes32);
    function getLinkedAppEVM(uint64 chainId) external view returns (address);
    function getModules() external view returns (address[] memory);
}
