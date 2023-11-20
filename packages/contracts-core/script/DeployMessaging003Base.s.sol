// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ CONTRACT IMPORTS ══════════════════════════════
import {AgentSecured, MessagingBase} from "../contracts/base/AgentSecured.sol";
import {BondingManager} from "../contracts/manager/BondingManager.sol";
import {Destination} from "../contracts/Destination.sol";
import {GasOracle} from "../contracts/GasOracle.sol";
import {Origin} from "../contracts/Origin.sol";
import {Summit} from "../contracts/Summit.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DeployerUtils} from "./utils/DeployerUtils.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {console, stdJson} from "forge-std/Script.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

// solhint-disable no-console
// solhint-disable ordering
abstract contract DeployMessaging003BaseScript is DeployerUtils {
    using stdJson for string;
    using Strings for uint256;

    string public constant DESTINATION = "Destination";
    string public constant GAS_ORACLE = "GasOracle";
    string public constant ORIGIN = "Origin";
    string public constant SUMMIT = "Summit";

    /// @notice BondingManager or LightManager address
    address public agentManager;
    /// @notice Inbox or LightInbox address
    address public statementInbox;

    // Common deployments for all chains
    address public destination;
    address public gasOracle;
    address public origin;

    // Only present on SynChain
    address public summit;

    uint32 public localDomain;
    uint256 public synapseDomain;
    string public globalConfig;

    constructor() {
        setupPK("MESSAGING_DEPLOYER_PRIVATE_KEY");
        localDomain = uint32(block.chainid);
        deploymentSalt = keccak256("Messaging003");
        setupDevnetIfEnabled();
    }

    /// @dev Function to exclude script from coverage report
    function testScript() external {}

    /// @notice Main function with the deploy logic.
    /// @dev To deploy contracts on $chainName
    /// Make sure "./script/configs/Messaging003.dc.json" exists, then call
    /// forge script script/DeployMessaging003<?>.s.sol -f chainName --ffi --broadcast
    /// Note: Verification will not natively work, need to verify externally
    function run() external {
        _deploy(true);
    }

    /// @notice Function to simulate the deployment procedure.
    /// @dev To simulate deployment on $chainName
    /// forge script script/DeployMessaging003<?>.s.sol -f chainName --sig "runDry()"
    function runDry() external {
        _deploy(false);
    }

    function isSynapseChain() public returns (bool) {
        return globalConfig.readUint(".chainidSummit") == localDomain;
    }

    function agentManagerName() public returns (string memory) {
        return isSynapseChain() ? "BondingManager" : "LightManager";
    }

    function statementInboxName() public returns (string memory) {
        return isSynapseChain() ? "Inbox" : "LightInbox";
    }

    /// @dev Deploys Messaging contracts, transfer ownership and sanity check the new deployments.
    /// Will save the deployments, if script is being broadcasted.
    function _deploy(bool _isBroadcasted) internal {
        globalConfig = loadGlobalDeployConfig("Messaging003");
        synapseDomain = globalConfig.readUint(".chainidSummit");
        startBroadcast(_isBroadcasted);
        // assert this is the first thing deployed
        getFactory();
        // Predict deployments
        agentManager = predictFactoryDeployment(agentManagerName());
        statementInbox = predictFactoryDeployment(statementInboxName());
        destination = predictFactoryDeployment(DESTINATION);
        gasOracle = predictFactoryDeployment(GAS_ORACLE);
        origin = predictFactoryDeployment(ORIGIN);
        if (isSynapseChain()) {
            summit = predictFactoryDeployment(SUMMIT);
        }
        // Deploy and initialize contracts
        _deployAndCheckAddress(agentManagerName(), _deployAgentManager, _initializeAgentManager, agentManager);
        _deployAndCheckAddress(statementInboxName(), _deployStatementInbox, _initializeStatementInbox, statementInbox);
        _deployAndCheckAddress(DESTINATION, _deployDestination, _initializeDestination, destination);
        _deployAndCheckAddress(GAS_ORACLE, _deployGasOracle, _initializeGasOracle, gasOracle);
        _deployAndCheckAddress(ORIGIN, _deployOrigin, _initializeOrigin, origin);
        if (isSynapseChain()) {
            _deployAndCheckAddress(SUMMIT, _deploySummit, _initializeSummit, summit);
        }
        // Add agents to BondingManager
        _addAgents();
        // Transfer ownership of contracts
        _transferOwnership();
        stopBroadcast();
        // Do some sanity checks
        _checkDeployments();
        _checkAgents();
    }

    function _deployAndCheckAddress(
        string memory contractName,
        function() internal returns (address, bytes memory) deployFunc,
        function(address) internal initFunc,
        address predictedDeployment
    ) internal {
        (address deployment,) = deployContract(contractName, deployFunc, initFunc);
        require(deployment == predictedDeployment, string.concat(contractName, ": wrong address"));
    }

    /// @dev Deploys BondingManager or LightManager
    /// Note: requires Origin, Destination, StatementInbox (and Summit for BondingManager) addresses to be set
    function _deployAgentManager() internal virtual returns (address deployment, bytes memory constructorArgs);

    /// @dev Initializes BondingManager or LightManager
    function _initializeAgentManager(address deployment) internal virtual;

    /// @dev Deploys Inbox or LightInbox
    /// Note: requires AgentManager, Origin, Destination (and Summit for Inbox) addresses to be set
    function _deployStatementInbox() internal virtual returns (address deployment, bytes memory constructorArgs);

    /// @dev Initializes Inbox or LightInbox
    function _initializeStatementInbox(address deployment) internal virtual;

    /// @dev Adds agents to BondingManager (no-op for LightManager)
    function _addAgents() internal virtual;

    // ═══════════════════════════════════════ DEPLOY AND INITIALIZE ROUTINE ═══════════════════════════════════════════

    /// @dev Deploys Destination.
    /// Note: requires AgentManager and StatementInbox addresses to have been set.
    /// Note: requires AgentManager to have been deployed.
    function _deployDestination() internal returns (address deployment, bytes memory constructorArgs) {
        // new Destination(domain, agentManager, statementInbox)
        require(agentManager != address(0), "Agent Manager not set");
        require(statementInbox != address(0), "Statement Inbox not set");
        require(agentManager.code.length > 0, "Agent Manager not deployed");
        constructorArgs = abi.encode(synapseDomain, agentManager, statementInbox);
        deployment = factoryDeploy(DESTINATION, type(Destination).creationCode, constructorArgs);
    }

    /// @dev Initializes Destination.
    function _initializeDestination(address deployment) internal {
        if (Destination(deployment).owner() == address(0)) {
            console.log("   %s: initializing", DESTINATION);
            Destination(deployment).initialize(_getInitialAgentRoot(), broadcasterAddress);
        } else {
            console.log("   %s: already initialized", DESTINATION);
        }
    }

    /// @dev Deploys GasOracle.
    /// Note: requires Destination address to have been set.
    function _deployGasOracle() internal returns (address deployment, bytes memory constructorArgs) {
        // new GasOracle(domain, destination)
        require(destination != address(0), "Destination not set");
        constructorArgs = abi.encode(synapseDomain, destination);
        deployment = factoryDeploy(GAS_ORACLE, type(GasOracle).creationCode, constructorArgs);
    }

    /// @dev Initializes GasOracle.
    function _initializeGasOracle(address deployment) internal {
        if (GasOracle(deployment).owner() == address(0)) {
            console.log("   %s: initializing", GAS_ORACLE);
            GasOracle(deployment).initialize(broadcasterAddress);
        } else {
            console.log("   %s: already initialized", GAS_ORACLE);
        }
    }

    /// @dev Deploys Origin.
    /// Note: requires AgentManager, StatementInbox and GasOracle addresses to have been set.
    function _deployOrigin() internal returns (address deployment, bytes memory constructorArgs) {
        // new Origin(domain, agentManager, statementInbox, gasOracle)
        require(agentManager != address(0), "Agent Manager not set");
        require(statementInbox != address(0), "Statement Inbox not set");
        require(gasOracle != address(0), "Gas Oracle not set");
        constructorArgs = abi.encode(synapseDomain, agentManager, statementInbox, gasOracle);
        deployment = factoryDeploy(ORIGIN, type(Origin).creationCode, constructorArgs);
    }

    /// @dev Initializes Origin.
    function _initializeOrigin(address deployment) internal {
        if (Origin(deployment).owner() == address(0)) {
            console.log("   %s: initializing", ORIGIN);
            Origin(deployment).initialize(broadcasterAddress);
        } else {
            console.log("   %s: already initialized", ORIGIN);
        }
    }

    /// @dev Deploys Summit.
    /// Note: requires AgentManager and StatementInbox addresses to have been set.
    function _deploySummit() internal returns (address deployment, bytes memory constructorArgs) {
        // new Summit(domain, agentManager, statementInbox)
        require(agentManager != address(0), "Agent Manager not set");
        require(statementInbox != address(0), "Statement Inbox not set");
        constructorArgs = abi.encode(synapseDomain, agentManager, statementInbox);
        deployment = factoryDeploy(SUMMIT, type(Summit).creationCode, constructorArgs);
    }

    /// @dev Initializes Summit.
    function _initializeSummit(address deployment) internal {
        if (Summit(deployment).owner() == address(0)) {
            console.log("   %s: initializing", SUMMIT);
            Summit(deployment).initialize(broadcasterAddress);
        } else {
            console.log("   %s: already initialized", SUMMIT);
        }
    }

    // ═════════════════════════════════════════════════ OWNERSHIP ═════════════════════════════════════════════════════

    /// @dev Transfers ownership of contracts to the owner defined in the global config.
    function _transferOwnership() internal {
        address owner = globalConfig.readAddress(".owner");
        console.log("Transferring ownership to %s", owner);
        _transferOwnership(agentManagerName(), agentManager, owner);
        _transferOwnership(statementInboxName(), statementInbox, owner);
        _transferOwnership(DESTINATION, destination, owner);
        _transferOwnership(GAS_ORACLE, gasOracle, owner);
        _transferOwnership(ORIGIN, origin, owner);
        _transferOwnership(SUMMIT, summit, owner);
    }

    /// @dev Transfers ownership of a given contract to a new owner.
    function _transferOwnership(string memory name, address deployment, address newOwner) internal {
        if (deployment == address(0)) {
            console.log("   %s: skipped (not deployed)", name);
            return;
        }
        if (Ownable(deployment).owner() == newOwner) {
            console.log("   %s: skipped (already owned)", name);
            return;
        }
        console.log("   %s: transferring ownership to %s", name, newOwner);
        Ownable(deployment).transferOwnership(newOwner);
    }

    // ══════════════════════════════════════════ POST DEPLOYMENT CHECKS ═══════════════════════════════════════════════

    /// @dev Checks that all agents have been added correctly to BondingManager
    /// or that they could be added to LightManager.
    function _checkAgents() internal virtual;

    /// @dev Checks that all contracts have been deployed and initialized correctly
    function _checkDeployments() internal {
        console.log("Checking deployments");
        // Check AgentSecured contracts
        console.log("   Checking AgentSecured contracts");
        _checkAgentSecured(DESTINATION, destination);
        _checkAgentSecured(ORIGIN, origin);
        if (isSynapseChain()) _checkAgentSecured(SUMMIT, summit);
        // Check MessagingBase contracts
        console.log("   Checking MessagingBase contracts");
        _checkMessagingBase(agentManagerName(), agentManager);
        _checkMessagingBase(statementInboxName(), statementInbox);
        _checkMessagingBase(GAS_ORACLE, gasOracle);
        // Check MessagingBase having .agentManager() getter
        console.log("   Checking contracts having .agentManager() getter");
        _checkGetterAgentManager(statementInboxName(), statementInbox);
        // Check MessagingBase having .inbox() getter
        console.log("   Checking contracts having .inbox() getter");
        _checkGetterInbox(agentManagerName(), agentManager);
        // Check contracts having .destination() getter
        console.log("   Checking contracts having .destination() getter");
        _checkGetterDestination(agentManagerName(), agentManager);
        _checkGetterDestination(statementInboxName(), statementInbox);
        _checkGetterDestination(GAS_ORACLE, gasOracle);
        // Check contracts having .gasOracle() getter
        console.log("   Checking contracts having .gasOracle() getter");
        _checkGetterGasOracle(ORIGIN, origin);
        // Check contracts having .origin() getter
        console.log("   Checking contracts having .origin() getter");
        _checkGetterOrigin(agentManagerName(), agentManager);
        _checkGetterOrigin(statementInboxName(), statementInbox);
        // Check contracts having .summit() getter
        if (isSynapseChain()) {
            console.log("   Checking contracts having .summit() getter");
            _checkGetterSummit(agentManagerName(), agentManager);
            _checkGetterSummit(statementInboxName(), statementInbox);
        }
    }

    /// @dev Checks AgentSecured getters:
    /// - .localDomain()
    /// - .owner()
    /// - .agentManager()
    /// - .inbox()
    function _checkAgentSecured(string memory name, address deployment) internal {
        _checkGetterAgentManager(name, deployment);
        _checkGetterInbox(name, deployment);
        _checkMessagingBase(name, deployment);
    }

    /// @dev Checks MessagingBase getters:
    /// - .localDomain()
    /// - .owner()
    function _checkMessagingBase(string memory name, address deployment) internal {
        _logCheckedGetter(name, "localDomain", deployment, localDomain);
        require(MessagingBase(deployment).localDomain() == localDomain, string.concat("Wrong localDomain: ", name));
        address owner = globalConfig.readAddress(".owner");
        _logCheckedGetter(name, "owner", deployment, owner);
        require(Ownable(deployment).owner() == owner, string.concat("Wrong owner: ", name));
    }

    /// @dev Checks .agentManager() getter
    function _checkGetterAgentManager(string memory name, address deployment) internal view {
        _logCheckedGetter(name, "agentManager", deployment, agentManager);
        require(AgentSecured(deployment).agentManager() == agentManager, string.concat("Wrong agentManager: ", name));
    }

    /// @dev Checks .inbox() getter
    function _checkGetterInbox(string memory name, address deployment) internal view {
        _logCheckedGetter(name, "inbox", deployment, statementInbox);
        require(AgentSecured(deployment).inbox() == statementInbox, string.concat("Wrong statementInbox: ", name));
    }

    /// @dev Checks .destination() getter
    function _checkGetterDestination(string memory name, address deployment) internal view {
        _logCheckedGetter(name, "destination", deployment, destination);
        require(GasOracle(deployment).destination() == destination, string.concat("Wrong destination: ", name));
    }

    /// @dev Checks .gasOracle() getter
    function _checkGetterGasOracle(string memory name, address deployment) internal view {
        _logCheckedGetter(name, "gasOracle", deployment, gasOracle);
        require(Origin(deployment).gasOracle() == gasOracle, string.concat("Wrong gasOracle: ", name));
    }

    /// @dev Checks .origin() getter
    function _checkGetterOrigin(string memory name, address deployment) internal view {
        _logCheckedGetter(name, "origin", deployment, origin);
        require(BondingManager(deployment).origin() == origin, string.concat("Wrong origin: ", name));
    }

    /// @dev Checks .summit() getter
    function _checkGetterSummit(string memory name, address deployment) internal view {
        _logCheckedGetter(name, "summit", deployment, summit);
        require(BondingManager(deployment).summit() == summit, string.concat("Wrong summit: ", name));
    }

    /// @dev Shortcut for logging
    function _logCheckedGetter(string memory name, string memory getter, address deployment, address expectedValue)
        internal
        view
    {
        console.log("       [%s][%s]: checking .%s()", name, deployment, getter);
        console.log("           expecting value: %s", expectedValue);
    }

    /// @dev Shortcut for logging
    function _logCheckedGetter(string memory name, string memory getter, address deployment, uint32 expectedValue)
        internal
        view
    {
        console.log("       [%s][%s]: checking .%s()", name, deployment, getter);
        console.log("           expecting value: %s", expectedValue);
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @dev Reads initial agent root
    function _getInitialAgentRoot() internal returns (bytes32) {
        // No saved agent root exists when deploying to SynChain
        if (isSynapseChain()) return 0;
        string memory agentRootConfig = loadGlobalDeployConfig("Messaging003AgentRoot");
        return agentRootConfig.readBytes32(".initialAgentRoot");
    }
}
