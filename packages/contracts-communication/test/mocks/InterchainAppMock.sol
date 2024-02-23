// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainApp} from "../../contracts/interfaces/IInterchainApp.sol";
import {AppConfigV1} from "../../contracts/libs/AppConfig.sol";

contract InterchainAppMock is IInterchainApp {
    address[] public receivingModules;
    AppConfigV1 public appConfig;

    function setReceivingModule(address _receivingModule) public {
        receivingModules.push(_receivingModule);
    }

    function setAppConfig(
        uint64[] memory chainIDs,
        address[] memory linkedIApps,
        address[] memory _sendingModules,
        address[] memory _receivingModules,
        uint256 _requiredResponses,
        uint64 _optimisticTimePeriod
    )
        external
        virtual
        override
    {
        // TODO: Add access control or ownership checks
        require(chainIDs.length == linkedIApps.length, "ChainIDs and IApps length mismatch");

        for (uint256 i = 0; i < chainIDs.length; i++) {
            appConfig.linkedIApps[chainIDs[i]] = linkedIApps[i];
        }

        appConfig.sendingModules = _sendingModules;
        appConfig.receivingModules = _receivingModules;
        appConfig.requiredResponses = _requiredResponses;
        appConfig.optimisticTimePeriod = _optimisticTimePeriod;
    }

    function getLinkedIApp(uint64 chainID) external view virtual override returns (address) {}

    function getSendingModules() external view virtual override returns (address[] memory) {}

    function getReceivingConfig() external override view returns (bytes memory retConfig, address[] memory modules) {
        return (abi.encode(appConfig), receivingModules);
    }

    function send(bytes32 receiver, uint256 dstChainId, bytes calldata message) external payable virtual override {}

    function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes calldata message) external payable {}
}
