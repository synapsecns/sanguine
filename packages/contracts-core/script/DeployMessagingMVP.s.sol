// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Script.sol";

import { DeployerUtils } from "./utils/DeployerUtils.sol";

import { AttestationCollector } from "../contracts/AttestationCollector.sol";
import { BondingManager, BondingMVP, SystemContract } from "../contracts/bonding/BondingMVP.sol";
import { Destination } from "../contracts/Destination.sol";
import { Origin } from "../contracts/Origin.sol";
import { SystemRouter } from "../contracts/system/SystemRouter.sol";
import { TestClient } from "../contracts/client/TestClient.sol";

import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";

contract DeployMessagingMVPScript is DeployerUtils {
    using stdJson for string;
    using Strings for uint256;

    string public constant ATTESTATION_COLLECTOR_NAME = "AttestationCollector";
    string public constant BONDING_MANAGER_NAME = "BondingMVP";
    string public constant DESTINATION_NAME = "Destination";
    string public constant ORIGIN_NAME = "Origin";
    string public constant SYSTEM_ROUTER_NAME = "SystemRouter";

    string public constant CLIENT_NAME = "TestClient";

    string public constant MESSAGING = "MessagingMVP";

    AttestationCollector public collector;
    BondingMVP public bondingMVP;
    Destination public destination;
    Origin public origin;
    SystemRouter public systemRouter;

    TestClient public testClient;

    address public owner;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
    }

    /// @notice Main function with the deploy logic.
    /// @dev To deploy contracts on $chainName
    /// Make sure "./script/configs/MessagingMVP.dc.json" exists, then call
    /// forge script script/DeployMessagingMVP.s.sol -f chainName --ffi --broadcast --verify
    function run() external {
        _deploy(true);
    }

    /// @notice Function to simulate the deployment procedure.
    /// @dev To simulate deployment on $chainName
    /// forge script script/DeployMessagingMVP.s.sol -f chainName --sig "runDry()"
    function runDry() external {
        _deploy(false);
    }

    /// @dev Deploys Messaging contracts, transfer ownership and sanity check the new deployments.
    /// Will save the deployments, if script is being broadcasted.
    function _deploy(bool _isBroadcasted) internal {
        startBroadcast(_isBroadcasted);
        // TODO: setup actual address in .dc.json files
        string memory config = loadGlobalDeployConfig(MESSAGING);
        owner = config.readAddress("owner");
        // Deploy System Contracts
        bondingMVP = BondingMVP(deployContract(BONDING_MANAGER_NAME, _deployBondingMVP));
        destination = Destination(deployContract(DESTINATION_NAME, _deployDestination));
        origin = Origin(deployContract(ORIGIN_NAME, _deployOrigin));
        // Deploy System Router
        systemRouter = SystemRouter(deployContract(SYSTEM_ROUTER_NAME, _deploySystemRouter));
        // Deploy Test Client
        testClient = TestClient(deployContract(CLIENT_NAME, _deployTestClient));
        // Deploy AttestationCollector, if requested for the current chain
        if (config.readUint("chainidAC") == block.chainid) {
            collector = AttestationCollector(
                deployContract(ATTESTATION_COLLECTOR_NAME, _deployAttestationCollector)
            );
            if (bondingMVP.attestationCollector() != address(collector)) {
                bondingMVP.setAttestationCollector(address(collector));
            }
        }
        // Setup System Contracts
        _setSystemRouter(bondingMVP);
        _setSystemRouter(destination);
        _setSystemRouter(origin);
        // Add preset agents from the config
        _addAgents(config);
        // Transfer ownership
        _transferOwnership(bondingMVP);
        _transferOwnership(destination);
        _transferOwnership(origin);
        // Stop broadcasting before testing the deployed contracts
        stopBroadcast();
        // Test: check if all agents from config were added
        _checkAgents(config);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            DEPLOY HELPERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _deployAttestationCollector() internal returns (address) {
        AttestationCollector _collector = new AttestationCollector();
        // Initialize to take ownership
        _collector.initialize();
        require(address(bondingMVP) != address(0), "BodingMVP is not yet deployed");
        _collector.transferOwnership(address(bondingMVP));
        return address(_collector);
    }

    /// @dev Callback function to deploy BondingManager
    function _deployBondingMVP() internal returns (address) {
        BondingMVP _bondingMVP = new BondingMVP(uint32(block.chainid));
        // Initialize to take ownership
        _bondingMVP.initialize();
        return address(_bondingMVP);
    }

    /// @dev Callback function to deploy Destination
    function _deployDestination() internal returns (address) {
        Destination _destination = new Destination(uint32(block.chainid));
        // Initialize to take ownership
        _destination.initialize();
        return address(_destination);
    }

    /// @dev Callback function to deploy Origin
    function _deployOrigin() internal returns (address) {
        Origin _origin = new Origin(uint32(block.chainid));
        // Initialize to take ownership
        _origin.initialize();
        return address(_origin);
    }

    /// @dev Callback function to deploy SystemRouter
    function _deploySystemRouter() internal returns (address) {
        SystemRouter _systemRouter = new SystemRouter({
            _domain: uint32(block.chainid),
            _origin: address(origin),
            _destination: address(destination),
            _bondingManager: address(bondingMVP)
        });
        // SystemRouter is unowned
        return address(_systemRouter);
    }

    function _deployTestClient() internal returns (address) {
        TestClient _testClient = new TestClient({
            _origin: address(origin),
            _destination: address(destination)
        });
        // TestClient is unowned
        return address(_testClient);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SETUP HELPERS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _addAgents(string memory config) internal {
        uint256[] memory domains = config.readUintArray("domains");
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            // Key is "agents.0: for Guards, "agents.10" for Optimism Notaries, etc
            address[] memory agents = config.readAddressArray(
                string.concat("agents.", domain.toString())
            );
            for (uint256 j = 0; j < agents.length; ++j) {
                bondingMVP.addAgent(uint32(domain), agents[j]);
                console.log("Adding Agent: %s on domain [%s]", agents[j], domain);
            }
        }
    }

    function _setSystemRouter(SystemContract sc) internal {
        // Check if broadcaster is the owner
        if (sc.owner() == broadcasterAddress) {
            // Setup systemRouter, if needed
            if (sc.systemRouter() != systemRouter) {
                sc.setSystemRouter(systemRouter);
                console.log("%s: systemRouter set to %s", address(sc), address(systemRouter));
            }
        }
    }

    function _transferOwnership(SystemContract sc) internal {
        // Check if broadcaster is the owner
        if (sc.owner() == broadcasterAddress) {
            // Transfer ownership
            sc.transferOwnership(owner);
            console.log("%s: ownership transferred to %s", address(sc), owner);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                TESTS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _checkAgents(string memory config) internal {
        uint256[] memory domains = config.readUintArray("domains");
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            bool isGuard = domain == 0;
            bool isLocalNotary = domain == block.chainid;
            bool isRemoteNotary = !isGuard && !isLocalNotary;
            // Key is "agents.0: for Guards, "agents.10" for Optimism Notaries, etc
            address[] memory agents = config.readAddressArray(
                string.concat("agents.", domain.toString())
            );
            for (uint256 j = 0; j < agents.length; ++j) {
                address agent = agents[j];
                // Destination needs to know about Guards and local Notaries
                require(
                    destination.isActiveAgent(uint32(domain), agent) == isGuard || isLocalNotary,
                    string.concat("!destination: ", domain.toString())
                );
                // Origin needs to know about Guards and remote Notaries
                require(
                    origin.isActiveAgent(uint32(domain), agent) == isGuard || isRemoteNotary,
                    string.concat("!origin: ", domain.toString())
                );
                // AttestationCollector needs to know about everything (if deployed)
                if (address(collector) != address(0)) {
                    require(
                        collector.isActiveAgent(uint32(domain), agent),
                        string.concat("!collector: ", domain.toString())
                    );
                }
            }
        }
    }
}
