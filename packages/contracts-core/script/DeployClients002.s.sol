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

contract DeployClients002Script is DeployerUtils {
    using stdJson for string;

    string public constant PING_PONG_CLIENT_NAME = "PingPongClient";
    string public constant TEST_CLIENT_NAME = "TestClient";

    address internal origin;
    address internal destination;

    PingPongClient public pingPongClient;
    TestClient public testClient;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
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
        // Deploy clients
        pingPongClient = PingPongClient(
            deployContract(PING_PONG_CLIENT_NAME, _deployPingPongClient)
        );
        testClient = TestClient(deployContract(TEST_CLIENT_NAME, _deployTestClient));
        stopBroadcast();
        // Check clients without broadcasting
        _checkClients();
    }

    function _deployPingPongClient() internal returns (address) {
        // (origin, destination)
        bytes memory constructorArgs = abi.encode(origin, destination);
        return
            factoryDeploy(
                PING_PONG_CLIENT_NAME,
                type(PingPongClient).creationCode,
                constructorArgs
            );
    }

    function _deployTestClient() internal returns (address) {
        // (origin, destination)
        bytes memory constructorArgs = abi.encode(origin, destination);
        return factoryDeploy(TEST_CLIENT_NAME, type(TestClient).creationCode, constructorArgs);
    }

    function _checkClients() internal {
        console.log("Checking Clients");
        uint256 initialStates = IStateHub(origin).statesAmount();
        uint32 remoteDomain = block.chainid == 10 ? 137 : 10;
        // Check Test Client
        testClient.sendMessage(remoteDomain, address(testClient), 0, "test message");
        require(IStateHub(origin).statesAmount() == initialStates + 1, "TestClient didn't send");
        console.log(unicode"   TestClient: ✅");
        // Check Ping Pong Client
        pingPongClient.doPing(remoteDomain, address(pingPongClient), 0);
        require(
            IStateHub(origin).statesAmount() == initialStates + 2,
            "PingPongClient didn't send"
        );
        console.log(unicode"   PingPongClient: ✅");
    }
}
