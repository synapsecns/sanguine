// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;


import { Script } from "forge-std/Script.sol";

import { InterchainDB } from "../contracts/InterchainDB.sol";

import { InterchainClientV1 } from "../contracts/InterchainClientV1.sol";

import { SynapseModule } from "../contracts/modules/SynapseModule.sol";

import { InterchainApp } from "../contracts/InterchainApp.sol";

import { GasOracleMock } from "../test/mocks/GasOracleMock.sol";



contract MessagingBase is Script {
    InterchainDB public icDB;
    InterchainClientV1 public icClient;
    SynapseModule public synapseModule;
    GasOracleMock public gasOracleMock;
    InterchainApp public icApp;

    function run() external {
        vm.startBroadcast();
        icDB = new InterchainDB();
        // icClient deployment & config
        icClient = new InterchainClientV1();
        icClient.setInterchainDB(address(icDB));


        synapseModule = new SynapseModule(address(icDB), msg.sender);
        gasOracleMock = new GasOracleMock();
        synapseModule.setGasOracle(address(gasOracleMock));
        icApp = new InterchainApp(address(icClient), new address[](0), new address[](0));



        vm.stopBroadcast();
    }
}
