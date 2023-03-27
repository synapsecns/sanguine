// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ CONTRACT IMPORTS ══════════════════════════════
import { BondingSecondary } from "../contracts/bonding/BondingSecondary.sol";
import { Destination } from "../contracts/Destination.sol";
import { Origin } from "../contracts/Origin.sol";
import { Summit } from "../contracts/Summit.sol";
import { SystemContract } from "../contracts/system/SystemContract.sol";
import { SystemRouter } from "../contracts/system/SystemRouter.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { DeployerUtils } from "./utils/DeployerUtils.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import { console, stdJson } from "forge-std/Script.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";

contract DeployMessagingTestnet002Script is DeployerUtils {
    using stdJson for string;
    using Strings for uint256;

    string public constant BONDING_MANAGER_NAME = "BondingSecondary";
    string public constant DESTINATION_NAME = "Destination";
    string public constant ORIGIN_NAME = "Origin";
    string public constant SYSTEM_ROUTER_NAME = "SystemRouter";
    string public constant SUMMIT_NAME = "Summit";

    string public constant MESSAGING_TESTNET_002 = "MessagingTestnet002";

    BondingSecondary public bondingManager;
    Destination public destination;
    Origin public origin;
    Summit public summit;
    SystemRouter public systemRouter;

    address public owner;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
    }

    /// @dev Function to exclude script from coverage report
    function testScript() external {}

    /// @notice Main function with the deploy logic.
    /// @dev To deploy contracts on $chainName
    /// Make sure "./script/configs/MessagingTestnet002.dc.json" exists, then call
    /// forge script script/DeployMessagingTestnet002.s.sol -f chainName --ffi --broadcast --verify
    function run() external {
        _deploy(true);
    }

    /// @notice Function to simulate the deployment procedure.
    /// @dev To simulate deployment on $chainName
    /// forge script script/DeployMessagingTestnet002.s.sol -f chainName --sig "runDry()"
    function runDry() external {
        _deploy(false);
    }

    /// @dev Deploys Messaging contracts, transfer ownership and sanity check the new deployments.
    /// Will save the deployments, if script is being broadcasted.
    function _deploy(bool _isBroadcasted) internal {
        startBroadcast(_isBroadcasted);
        string memory config = loadGlobalDeployConfig(MESSAGING_TESTNET_002);
        // Deploy System Contracts
        if (config.readUint(".chainidSummit") == block.chainid) {
            address deployed = deployContract(SUMMIT_NAME, _deploySummit);
            summit = Summit(deployed);
            // Summit is also Bonding Primary Manager
            bondingManager = BondingSecondary(deployed);
        } else {
            bondingManager = BondingSecondary(
                deployContract(BONDING_MANAGER_NAME, _deployBondingSecondary)
            );
        }
        // Now `bondingManager` points to local BondingManager, whether it is Summit or not
        destination = Destination(deployContract(DESTINATION_NAME, _deployDestination));
        origin = Origin(deployContract(ORIGIN_NAME, _deployOrigin));
        systemRouter = SystemRouter(deployContract(SYSTEM_ROUTER_NAME, _deploySystemRouter));
        // Setup System Contracts
        console.log("Setting SystemRouter");
        _setSystemRouter(bondingManager);
        _setSystemRouter(destination);
        _setSystemRouter(origin);
        // Add preset agents from the config
        _addAgents(config);
        // Transfer ownership
        owner = config.readAddress(".owner");
        console.log("Transferring ownership");
        _transferOwnership(bondingManager);
        _transferOwnership(destination);
        _transferOwnership(origin);
        // Stop broadcasting before testing the deployed contracts
        stopBroadcast();
        _checkAgents(config);
    }

    function _deployBondingSecondary() internal returns (address) {
        // (domain)
        bytes memory constructorArgs = abi.encode(block.chainid);
        BondingSecondary deployed = BondingSecondary(
            factoryDeploy(
                BONDING_MANAGER_NAME,
                type(BondingSecondary).creationCode,
                constructorArgs
            )
        );
        // Initialize to take ownership
        deployed.initialize();
        return address(deployed);
    }

    function _deployDestination() internal returns (address) {
        // (domain)
        bytes memory constructorArgs = abi.encode(block.chainid);
        Destination deployed = Destination(
            factoryDeploy(DESTINATION_NAME, type(Destination).creationCode, constructorArgs)
        );
        // Initialize to take ownership
        deployed.initialize();
        return address(deployed);
    }

    function _deployOrigin() internal returns (address) {
        // (domain)
        bytes memory constructorArgs = abi.encode(block.chainid);
        Origin deployed = Origin(
            factoryDeploy(ORIGIN_NAME, type(Origin).creationCode, constructorArgs)
        );
        // Initialize to take ownership
        deployed.initialize();
        return address(deployed);
    }

    function _deploySummit() internal returns (address) {
        // (domain)
        bytes memory constructorArgs = abi.encode(block.chainid);
        Summit deployed = Summit(
            factoryDeploy(SUMMIT_NAME, type(Summit).creationCode, constructorArgs)
        );
        // Initialize to take ownership
        deployed.initialize();
        return address(deployed);
    }

    function _deploySystemRouter() internal returns (address) {
        // (domain, origin, destination, bondingManager)
        bytes memory constructorArgs = abi.encode(
            block.chainid,
            origin,
            destination,
            bondingManager
        );
        // SystemRouter is unowned
        return factoryDeploy(SYSTEM_ROUTER_NAME, type(SystemRouter).creationCode, constructorArgs);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SETUP HELPERS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _addAgents(string memory config) internal {
        console.log("Adding Agents");
        uint256[] memory domains = config.readUintArray(".domains");
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            // Key is ".agents.0: for Guards, ".agents.10" for Optimism Notaries, etc
            address[] memory agents = config.readAddressArray(
                string.concat(".agents.", domain.toString())
            );
            for (uint256 j = 0; j < agents.length; ++j) {
                bondingManager.addAgent(uint32(domain), agents[j]);
                console.log("   %s on domain [%s]", agents[j], domain);
            }
        }
    }

    function _setSystemRouter(SystemContract sc) internal {
        // Check if broadcaster is the owner
        if (sc.owner() == broadcasterAddress) {
            // Setup systemRouter, if needed
            if (sc.systemRouter() != systemRouter) {
                sc.setSystemRouter(systemRouter);
                console.log("   %s: set to %s", address(sc), address(systemRouter));
            } else {
                console.log("   %s: already set", address(sc));
            }
        } else {
            console.log("   %s: broadcaster is not owner", address(sc));
        }
    }

    function _transferOwnership(SystemContract sc) internal {
        // Check if broadcaster is the owner
        if (sc.owner() == broadcasterAddress) {
            // Transfer ownership
            sc.transferOwnership(owner);
            console.log("   %s: new owner is %s", address(sc), owner);
        } else {
            console.log("   %s: broadcaster is not owner! Owner: %s", address(sc), sc.owner());
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                TESTS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _checkAgents(string memory config) internal {
        console.log("Checking Agents");
        uint256[] memory domains = config.readUintArray(".domains");
        for (uint256 i = 0; i < domains.length; ++i) {
            uint256 domain = domains[i];
            bool isGuard = domain == 0;
            bool isLocalNotary = domain == block.chainid;
            // Key is ".agents.0: for Guards, ".agents.10" for Optimism Notaries, etc
            address[] memory agents = config.readAddressArray(
                string.concat(".agents.", domain.toString())
            );
            for (uint256 j = 0; j < agents.length; ++j) {
                address agent = agents[j];
                // Destination needs to know about Guards and local Notaries
                require(
                    destination.isActiveAgent(uint32(domain), agent) == isGuard || isLocalNotary,
                    string.concat("!destination: ", domain.toString())
                );
                // Origin needs to know about every Agent
                require(
                    origin.isActiveAgent(uint32(domain), agent),
                    string.concat("!origin: ", domain.toString())
                );
                // BondingManager/Summit needs to know about every Agent
                require(
                    bondingManager.isActiveAgent(uint32(domain), agent),
                    string.concat("!bondingManager: ", domain.toString())
                );
                console.log("   %s on domain [%s]", agent, domain);
            }
        }
    }
}
