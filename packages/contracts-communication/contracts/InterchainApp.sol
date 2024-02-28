pragma solidity 0.8.20;

import "./interfaces/IInterchainClientV1.sol";
import "./interfaces/IInterchainApp.sol";
import {OptionsLib, OptionsV1} from "./libs/Options.sol";
import {AppConfigLib, AppConfigV1} from "./libs/AppConfig.sol";

contract InterchainApp is IInterchainApp {
    using AppConfigLib for bytes;
    // What properties should Interchain be pulling from InterchainApp?
    // 1. Which modules to use, and how many are required?

    IInterchainClientV1 public interchain;

    address[] private sendingModules;
    address[] private receivingModules;

    struct AppConfig {
        // ChainID -> Linked IApps
        mapping(uint64 => address) linkedIApps;
        // Sends message to be verified through all modules
        address[] sendingModules;
        // Accepts messages from these destination chain modules
        address[] receivingModules;
        AppConfigV1 bytesAppConfig;
    }

    AppConfig private localAppConfig;

    // Set the application configuration
    function setAppConfig(
        uint64[] memory chainIDs,
        address[] memory linkedIApps,
        address[] memory _sendingModules,
        address[] memory _receivingModules,
        uint256 _requiredResponses,
        uint64 _optimisticTimePeriod
    )
        public
    {
        // TODO: Add access control or ownership checks
        require(chainIDs.length == linkedIApps.length, "ChainIDs and IApps length mismatch");

        for (uint256 i = 0; i < chainIDs.length; i++) {
            localAppConfig.linkedIApps[chainIDs[i]] = linkedIApps[i];
        }

        localAppConfig.bytesAppConfig =
            AppConfigV1({requiredResponses: _requiredResponses, optimisticPeriod: _optimisticTimePeriod});

        localAppConfig.sendingModules = _sendingModules;
        localAppConfig.receivingModules = _receivingModules;
    }

    // Getters for the application configuration
    function getLinkedIApp(uint64 chainID) external view returns (address) {
        return localAppConfig.linkedIApps[chainID];
    }

    // TODO: Is a receiving module the same as a sending module?
    function getSendingModules() external view returns (address[] memory) {
        return localAppConfig.sendingModules;
    }

    function getReceivingModules() external view returns (address[] memory) {
        return localAppConfig.receivingModules;
    }

    function getRequiredResponses() external view returns (uint256) {
        return localAppConfig.bytesAppConfig.requiredResponses;
    }

    function getOptimisticTimePeriod() external view returns (uint256) {
        return localAppConfig.bytesAppConfig.optimisticPeriod;
    }

    function getSendingModules(bytes32 receiver, uint256 dstChainId) external view returns (address[] memory) {
        return sendingModules;
    }

    function getReceivingConfig() external view returns (bytes memory, address[] memory) {
        return (AppConfigLib.encodeAppConfigV1(localAppConfig.bytesAppConfig), localAppConfig.receivingModules);
    }

    constructor(address _interchain, address[] memory _sendingModules, address[] memory _receivingModules) {
        interchain = IInterchainClientV1(_interchain);
        localAppConfig.sendingModules = _sendingModules;
        localAppConfig.receivingModules = _receivingModules;
    }

    event AppMessageRecieve();
    event AppMessageSent();

    function send(bytes32 receiver, uint256 dstChainId, bytes calldata message) external payable {
        bytes memory options = OptionsV1(200_000, 0).encodeOptionsV1();
        // TODO: Currently, we forward all gas to Interchain, this may not be expected behavior, and the real abstract contract shouldn't do this
        interchain.interchainSend{value: msg.value}(
            dstChainId, receiver, address(0), localAppConfig.sendingModules, options, message
        );
        emit AppMessageSent();
    }

    // TODO: Auth checks based on incoming message
    function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, bytes calldata message) external payable {
        emit AppMessageRecieve();
    }
}
