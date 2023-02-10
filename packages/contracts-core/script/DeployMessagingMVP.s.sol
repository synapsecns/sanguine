// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Script.sol";

import { DeployerUtils } from "./utils/DeployerUtils.sol";

import { AttestationCollector } from "../contracts/AttestationCollector.sol";
import { BondingManager, BondingMVP, SystemContract } from "../contracts/bonding/BondingMVP.sol";
import { Destination } from "../contracts/Destination.sol";
import { Origin } from "../contracts/Origin.sol";
import { SystemRouter } from "../contracts/system/SystemRouter.sol";

contract DeployMessagingMVPScript is DeployerUtils {
    using stdJson for string;

    string public constant ATTESTATION_COLLECTOR_NAME = "AttestationCollector";
    string public constant BONDING_MANAGER_NAME = "BondingMVP";
    string public constant DESTINATION_NAME = "Destination";
    string public constant ORIGIN_NAME = "Origin";
    string public constant SYSTEM_ROUTER_NAME = "SystemRouter";

    string public constant MESSAGING = "MessagingMVP";

    AttestationCollector public collector;
    BondingMVP public bondingMVP;
    Destination public destination;
    Origin public origin;
    SystemRouter public systemRouter;

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

    function checkDeployments() public {
        vm.startPrank(owner);
        // Zero Domain refers to a Guard
        _checkAgent({
            domain: 0,
            agent: address(1),
            agentName: "Guard",
            originAffected: true,
            destinationAffected: true
        });
        // Check Notary for current chain
        _checkAgent({
            domain: uint32(block.chainid),
            agent: address(2),
            agentName: "Local Notary",
            originAffected: false,
            destinationAffected: true
        });
        // Check Notary for chain other than the current one
        _checkAgent({
            domain: uint32(block.chainid ^ 1),
            agent: address(3),
            agentName: "Remote Notary",
            originAffected: true,
            destinationAffected: false
        });
        vm.stopPrank();
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
        // Deploy AttestationCollector, if requested for the current chain
        if (config.readUint("chainidAC") == block.chainid) {
            collector = AttestationCollector(
                deployContract(ATTESTATION_COLLECTOR_NAME, _deployAttestationCollector)
            );
            if (bondingMVP.attestationCollector() != address(collector)) {
                bondingMVP.setAttestationCollector(address(collector));
            }
        }
        // Deploy System Router
        systemRouter = SystemRouter(deployContract(SYSTEM_ROUTER_NAME, _deploySystemRouter));
        // Setup System Contracts
        _setupSystemContract(bondingMVP);
        _setupSystemContract(destination);
        _setupSystemContract(origin);
        // Stop broadcasting before testing the deployed contracts
        stopBroadcast();
        // Test: add and remove agents using the deployed contracts (no broadcast)
        checkDeployments();
    }

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

    function _setupSystemContract(SystemContract sc) internal {
        // Check if broadcaster is the owner
        if (sc.owner() == broadcasterAddress) {
            // Setup systemRouter, if needed
            if (sc.systemRouter() != systemRouter) {
                sc.setSystemRouter(systemRouter);
                console.log("%s: systemRouter set to %s", address(sc), address(systemRouter));
            }
            // Transfer ownership
            sc.transferOwnership(owner);
            console.log("%s: ownership transferred to %s", address(sc), owner);
        }
    }

    function _checkAgent(
        uint32 domain,
        address agent,
        string memory agentName,
        bool originAffected,
        bool destinationAffected
    ) internal {
        // Add agent
        bondingMVP.addAgent(domain, agent);
        // Check if agent was added to Origin
        if (originAffected) {
            require(
                origin.isActiveAgent(domain, agent),
                string.concat("Origin: failed to add ", agentName)
            );
        } else {
            require(
                !origin.isActiveAgent(domain, agent),
                string.concat("Origin: added ", agentName)
            );
        }
        // Check if agent was added to Destination
        if (destinationAffected) {
            require(
                destination.isActiveAgent(domain, agent),
                string.concat("Destination: failed to add ", agentName)
            );
        } else {
            require(
                !destination.isActiveAgent(domain, agent),
                string.concat("Destination: added ", agentName)
            );
        }
        // Check if agent was added to AttestationCollector
        if (address(collector) != address(0)) {
            require(
                collector.isActiveAgent(domain, agent),
                string.concat("AttestationCollector: failed to add ", agentName)
            );
        }
        bondingMVP.removeAgent(domain, agent);
        require(
            !origin.isActiveAgent(domain, agent),
            string.concat("Origin: failed to remove ", agentName)
        );
        require(
            !destination.isActiveAgent(domain, agent),
            string.concat("Destination: failed to remove ", agentName)
        );
        if (address(collector) != address(0)) {
            require(
                !collector.isActiveAgent(domain, agent),
                string.concat("AttestationCollector: failed to remove ", agentName)
            );
        }
    }
}
