// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "forge-std/Script.sol";

import { DeployerUtils } from "./utils/DeployerUtils.sol";

import { BondingManager, BondingMVP } from "../contracts/bonding/BondingMVP.sol";
import { Destination } from "../contracts/Destination.sol";
import { Origin } from "../contracts/Origin.sol";
import { SystemRouter } from "../contracts/system/SystemRouter.sol";

contract DeployMessagingMVPScript is DeployerUtils {
    string public constant BONDING_MANAGER_NAME = "BondingMVP";
    string public constant DESTINATION_NAME = "Destination";
    string public constant ORIGIN_NAME = "Origin";
    string public constant SYSTEM_ROUTER_NAME = "SystemRouter";

    BondingMVP public bondingMVP;
    Destination public destination;
    Origin public origin;
    SystemRouter public systemRouter;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
    }

    /// @notice Main function with the deploy logic
    /// @dev To deploy contracts on $chainName:
    /// Make sure ./deployments/$chainName exists, then call
    /// forge script script/DeployMessagingMVP.s.sol -f chainName --ffi --broadcast --verify
    function run() external {
        _deploy(true);
    }

    /// @notice Function to simulate the deployment procedure
    /// @dev To simulate deployment on {chainName}
    /// forge script script/DeployMessagingMVP.s.sol -f chainName --sig "runDry()"
    function runDry() external {
        _deploy(false);
    }

    function checkDeployments() public {
        vm.startPrank(broadcasterAddress);
        _checkAgent({
            domain: 0,
            agent: address(1),
            agentName: "Guard",
            originAffected: true,
            destinationAffected: true
        });
        _checkAgent({
            domain: uint32(block.chainid),
            agent: address(2),
            agentName: "Local Notary",
            originAffected: false,
            destinationAffected: true
        });
        _checkAgent({
            domain: uint32(block.chainid ^ 1),
            agent: address(3),
            agentName: "Remote Notary",
            originAffected: true,
            destinationAffected: false
        });
        vm.stopPrank();
    }

    function _deploy(bool _isBroadcasted) internal {
        startBroadcast(_isBroadcasted);
        _deployBondingMVP();
        _deployDestination();
        _deployOrigin();
        _deploySystemRouter();
        stopBroadcast();
        // Test add and remove agents without broadcasting using the deployed contracts
        checkDeployments();
    }

    function _deployBondingMVP() internal {
        bondingMVP = new BondingMVP(uint32(block.chainid));
        // Initialize to take ownership
        bondingMVP.initialize();
        saveDeployment(BONDING_MANAGER_NAME, address(bondingMVP));
    }

    function _deployDestination() internal {
        destination = new Destination(uint32(block.chainid));
        // Initialize to take ownership
        destination.initialize();
        saveDeployment(DESTINATION_NAME, address(destination));
    }

    function _deployOrigin() internal {
        origin = new Origin(uint32(block.chainid));
        // Initialize to take ownership
        origin.initialize();
        saveDeployment(ORIGIN_NAME, address(origin));
    }

    function _deploySystemRouter() internal {
        systemRouter = new SystemRouter({
            _domain: uint32(block.chainid),
            _origin: address(origin),
            _destination: address(destination),
            _bondingManager: address(bondingMVP)
        });
        // SystemRouter is unowned immutable contract, no further setup is required
        saveDeployment(SYSTEM_ROUTER_NAME, address(systemRouter));
        // Setup SystemRouter on deployed messaging contracts
        bondingMVP.setSystemRouter(systemRouter);
        destination.setSystemRouter(systemRouter);
        origin.setSystemRouter(systemRouter);
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
        bondingMVP.removeAgent(domain, agent);
        require(
            !origin.isActiveAgent(domain, agent),
            string.concat("Origin: failed to remove ", agentName)
        );
        require(
            !destination.isActiveAgent(domain, agent),
            string.concat("Destination: failed to remove ", agentName)
        );
    }
}
