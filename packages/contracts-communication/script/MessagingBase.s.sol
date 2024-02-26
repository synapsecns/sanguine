// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";

import {InterchainDB} from "../contracts/InterchainDB.sol";

import {InterchainClientV1} from "../contracts/InterchainClientV1.sol";

import {SynapseModule} from "../contracts/modules/SynapseModule.sol";

import {InterchainApp} from "../contracts/InterchainApp.sol";

import {GasOracleMock} from "../test/mocks/GasOracleMock.sol";

import {ExecutionFeesMock} from "../test/mocks/ExecutionFeesMock.sol";

contract MessagingBase is Script {
    InterchainDB public icDB;
    InterchainClientV1 public icClient;
    SynapseModule public synapseModule;
    GasOracleMock public gasOracleMock;
    InterchainApp public icApp;
    ExecutionFeesMock public executionFees;

    function run() external {
        vm.startBroadcast();
        icDB = new InterchainDB();
        // icClient deployment & config
        icClient = new InterchainClientV1();
        icClient.setInterchainDB(address(icDB));

        synapseModule = new SynapseModule(address(icDB), msg.sender);
        gasOracleMock = new GasOracleMock();
        executionFees = new ExecutionFeesMock();
        synapseModule.setGasOracle(address(gasOracleMock));
        icClient.setExecutionFees(address(executionFees));
        icApp = new InterchainApp(address(icClient), new address[](0), new address[](0));

        vm.stopBroadcast();
    }

    function newApp() external {
        vm.startBroadcast();
        icApp = new InterchainApp(0x007f3EC4A8E6DbaA84E50d0901F48966D83B7300, new address[](0), new address[](0));
        vm.stopBroadcast();
    }

    function newAppSepolia() public {
        vm.startBroadcast();
        icApp = new InterchainApp(0xFd8A9eDf272e54614426a4c5849851D26A57C644, new address[](0), new address[](0));
        vm.stopBroadcast();
    }

    function newAppConfigOP() external {
        vm.startBroadcast();
        icApp = InterchainApp(0x4a1f8D1378b614a59D0BB62EeeD811aaC1d22EC0);
        uint64[] memory chainIDs = new uint64[](1);
        chainIDs[0] = 11_155_111;
        address[] memory linkedIApps = new address[](1);
        linkedIApps[0] = 0x4a1f8D1378b614a59D0BB62EeeD811aaC1d22EC0;
        address[] memory _sendingModules = new address[](1);
        _sendingModules[0] = 0x48ADb7308f59d98657e779681aE6037902901918;
        address[] memory _receivingModules = new address[](1);
        _receivingModules[0] = 0x48ADb7308f59d98657e779681aE6037902901918;
        uint256 _requiredResponses = 1;
        uint64 _optimisticTimePeriod = 1;
        icApp.setAppConfig(
            chainIDs, linkedIApps, _sendingModules, _receivingModules, _requiredResponses, _optimisticTimePeriod
        );
        vm.stopBroadcast();
    }

    function newAppConfigSepolia() external {
        vm.startBroadcast();
        icApp = InterchainApp(0x4a1f8D1378b614a59D0BB62EeeD811aaC1d22EC0);
        uint64[] memory chainIDs = new uint64[](1);
        chainIDs[0] = 11_155_111;
        address[] memory linkedIApps = new address[](1);
        linkedIApps[0] = 0x4a1f8D1378b614a59D0BB62EeeD811aaC1d22EC0;
        address[] memory _sendingModules = new address[](1);
        _sendingModules[0] = 0x135189D37b0a734e4A339F0e9fd3219521729e7A;
        address[] memory _receivingModules = new address[](1);
        _receivingModules[0] = 0x135189D37b0a734e4A339F0e9fd3219521729e7A;
        uint256 _requiredResponses = 1;
        uint64 _optimisticTimePeriod = 1;
        icApp.setAppConfig(
            chainIDs, linkedIApps, _sendingModules, _receivingModules, _requiredResponses, _optimisticTimePeriod
        );
        vm.stopBroadcast();
    }
}
