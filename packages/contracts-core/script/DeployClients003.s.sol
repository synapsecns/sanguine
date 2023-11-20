// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ CONTRACT IMPORTS ══════════════════════════════

import { PingPongClient } from "../contracts/client/PingPongClient.sol";
import { TestClient } from "../contracts/client/TestClient.sol";
import { IStateHub } from "../contracts/interfaces/IStateHub.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DeployerUtils } from "./utils/DeployerUtils.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import { console, stdJson } from "forge-std/Script.sol";

// solhint-disable no-console
// solhint-disable ordering
contract DeployClients003Script is DeployerUtils {
    using stdJson for string;

    string public constant PING_PONG_CLIENT_NAME = "PingPongClient";
    string public constant TEST_CLIENT_NAME = "TestClient";

    address internal origin;
    address internal destination;

    address public pingPongClient;
    address public testClient;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
        deploymentSalt = keccak256("Messaging003-08");
    }

    /// @dev Function to exclude script from coverage report
    function testScript() external {}

    /// @notice Main function with the deploy logic.
    function run() external {
        _deploy(true);
    }

    /// @notice Function to simulate the deployment procedure.
    function runDry() external {
        _deploy(false);
    }

    /// @dev Deploys PingPongClient and TestClient
    function _deploy(bool _isBroadcasted) internal {
        startBroadcast(_isBroadcasted);
        // Load Origin/Destination deployments
        origin = loadDeployment("Origin");
        destination = loadDeployment("Destination");
        // Predict deployments
        pingPongClient = predictFactoryDeployment(PING_PONG_CLIENT_NAME);
        testClient = predictFactoryDeployment(TEST_CLIENT_NAME);
        // Deploy clients
        _deployAndCheckAddress(
            PING_PONG_CLIENT_NAME,
            _deployPingPongClient,
            _initNoop,
            pingPongClient
        );
        _deployAndCheckAddress(TEST_CLIENT_NAME, _deployTestClient, _initNoop, testClient);
        stopBroadcast();
        // Check clients without broadcasting
        _checkClients();
    }

    function _deployAndCheckAddress(
        string memory contractName,
        function() internal returns (address, bytes memory) deployFunc,
        function(address) internal initFunc,
        address predictedDeployment
    ) internal {
        (address deployment, ) = deployContract(contractName, deployFunc, initFunc);
        require(deployment == predictedDeployment, string.concat(contractName, ": wrong address"));
    }

    function _deployPingPongClient()
        internal
        returns (address deployment, bytes memory constructorArgs)
    {
        // new PingPongClient(origin, destination)
        constructorArgs = abi.encode(origin, destination);
        deployment = factoryDeploy(
            PING_PONG_CLIENT_NAME,
            type(PingPongClient).creationCode,
            constructorArgs
        );
    }

    function _deployTestClient()
        internal
        returns (address deployment, bytes memory constructorArgs)
    {
        // new TestClient(origin, destination)
        constructorArgs = abi.encode(origin, destination);
        deployment = factoryDeploy(
            TEST_CLIENT_NAME,
            type(TestClient).creationCode,
            constructorArgs
        );
    }

    function _initNoop(address _deployment) internal {}

    function _checkClients() internal {
        console.log("Checking Clients");
        uint256 initialStates = IStateHub(origin).statesAmount();
        uint32 remoteDomain = block.chainid == 10 ? 137 : 10;
        // Check Test Client
        TestClient(testClient).sendMessage(
            remoteDomain,
            address(testClient),
            0,
            100_000,
            0,
            "test message"
        );
        require(IStateHub(origin).statesAmount() == initialStates + 1, "TestClient didn't send");
        console.log(unicode"   TestClient: ✅");
        // Check Ping Pong Client
        PingPongClient(pingPongClient).doPing(remoteDomain, address(pingPongClient), 0);
        require(
            IStateHub(origin).statesAmount() == initialStates + 2,
            "PingPongClient didn't send"
        );
        console.log(unicode"   PingPongClient: ✅");
    }
}
