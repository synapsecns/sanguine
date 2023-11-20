// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, BondingManager} from "../../contracts/manager/BondingManager.sol";
import {AgentStatus, LightManager} from "../../contracts/manager/LightManager.sol";
import {IAgentSecured} from "../../contracts/interfaces/IAgentSecured.sol";
import {Destination} from "../../contracts/Destination.sol";
import {GasOracle} from "../../contracts/GasOracle.sol";
import {Origin} from "../../contracts/Origin.sol";
import {Summit} from "../../contracts/Summit.sol";

import {Inbox} from "../../contracts/inbox/Inbox.sol";
import {LightInbox} from "../../contracts/inbox/LightInbox.sol";

import {BondingManagerHarness} from "../harnesses/manager/BondingManagerHarness.t.sol";
import {LightManagerHarness} from "../harnesses/manager/LightManagerHarness.t.sol";

import {DestinationMock} from "../mocks/DestinationMock.t.sol";
import {GasOracleMock} from "../mocks/GasOracleMock.t.sol";
import {OriginMock} from "../mocks/OriginMock.t.sol";
import {SummitMock} from "../mocks/SummitMock.t.sol";

import {ProductionEvents} from "./events/ProductionEvents.t.sol";
import {SuiteEvents} from "./events/SuiteEvents.t.sol";
import {SynapseAgents} from "./SynapseAgents.t.sol";
import {SynapseProofs} from "./SynapseProofs.t.sol";

// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract SynapseTest is ProductionEvents, SuiteEvents, SynapseAgents, SynapseProofs {
    uint256 private immutable deployMask;

    address internal gasOracleSynapse;
    address internal originSynapse;
    address internal destinationSynapse;
    address internal summit;
    Inbox internal inbox;
    BondingManagerHarness internal bondingManager;

    address internal gasOracle;
    address internal destination;
    address internal origin;
    LightInbox internal lightInbox;
    LightManagerHarness internal lightManager;

    constructor(uint256 deployMask_) {
        deployMask = deployMask_;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseTest() external {}

    function setUp() public virtual override {
        // Setup domains and create agents for them
        super.setUp();
        // Deploy a single set of messaging contracts for synapse chain
        deployBondingManager();
        deployInbox();
        deploySummit();
        deployDestinationSynapse();
        deployGasOracleSynapse();
        deployOriginSynapse();
        // Setup agents in BondingManager
        initBondingManager();
        initInbox();
        setupAgentsBM();
        // Deploy a single set of messaging contracts for local chain
        deployLightManager();
        deployLightInbox();
        deployDestination();
        deployGasOracle();
        deployOrigin();
        // Setup agents in LightManager
        initLightManager();
        initLightInbox();
        setupAgentsLM();
        // Skip block
        skipBlock();
        // cleanup the env.
        vm.chainId(DOMAIN_LOCAL);
    }

    function setupAgentsBM() public virtual {
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < DOMAIN_AGENTS; ++i) {
                address agent = domains[domain].agents[i];
                addAgentBM(domain, agent);
            }
        }
    }

    function setupAgentsLM() public virtual {
        if (deployMask & DEPLOY_MASK_DESTINATION == DEPLOY_PROD_DESTINATION) {
            // Set initial agent merkle root via production Destination
            Destination(destination).initialize(getAgentRoot(), address(this));
        } else {
            // Mock a call from destination instead
            bytes32 root = getAgentRoot();
            vm.prank(destination);
            lightManager.setAgentRoot(root);
        }
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < DOMAIN_AGENTS; ++i) {
                address agent = domains[domain].agents[i];
                updateAgentLM(agent);
            }
        }
    }

    // ═════════════════════════════════════════════ DEPLOY CONTRACTS ══════════════════════════════════════════════════

    function deployLightManager() public virtual {
        vm.chainId(DOMAIN_LOCAL);
        lightManager = new LightManagerHarness(DOMAIN_SYNAPSE);
        vm.label(address(lightManager), "LightManager");
    }

    function initLightManager() public virtual {
        lightManager.initialize(origin, destination, address(lightInbox), address(this));
    }

    function deployBondingManager() public virtual {
        vm.chainId(DOMAIN_SYNAPSE);
        bondingManager = new BondingManagerHarness(DOMAIN_SYNAPSE);
        vm.label(address(bondingManager), "BondingManager");
    }

    function initBondingManager() public virtual {
        bondingManager.initialize(originSynapse, destinationSynapse, address(inbox), summit, address(this));
    }

    function deployLightInbox() public virtual {
        vm.chainId(DOMAIN_LOCAL);
        lightInbox = new LightInbox(DOMAIN_SYNAPSE);
    }

    function initLightInbox() public virtual {
        LightInbox(lightInbox).initialize(address(lightManager), origin, destination, address(this));
    }

    function deployInbox() public virtual {
        vm.chainId(DOMAIN_SYNAPSE);
        inbox = new Inbox(DOMAIN_SYNAPSE);
    }

    function initInbox() public virtual {
        Inbox(inbox).initialize(address(bondingManager), originSynapse, destinationSynapse, summit, address(this));
    }

    function deployGasOracle() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_GAS_ORACLE;
        if (option == DEPLOY_MOCK_GAS_ORACLE) {
            gasOracle = address(new GasOracleMock());
        } else if (option == DEPLOY_PROD_GAS_ORACLE) {
            gasOracle = address(new GasOracle(DOMAIN_SYNAPSE, destination));
            GasOracle(gasOracle).initialize(address(this));
        } else {
            revert("Unknown option: GasOracle");
        }
        vm.label(gasOracle, "GasOracle Local");
    }

    function deployDestination() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_DESTINATION;
        if (option == DEPLOY_MOCK_DESTINATION) {
            destination = address(new DestinationMock());
        } else if (option == DEPLOY_PROD_DESTINATION) {
            destination = address(new Destination(DOMAIN_SYNAPSE, address(lightManager), address(lightInbox)));
            // Destination will be initialized once agents are setup
        } else {
            revert("Unknown option: Destination");
        }
        vm.label(destination, "Destination Local");
    }

    function deployOrigin() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_ORIGIN;
        if (option == DEPLOY_MOCK_ORIGIN) {
            origin = address(new OriginMock());
        } else if (option == DEPLOY_PROD_ORIGIN) {
            origin = address(new Origin(DOMAIN_SYNAPSE, address(lightManager), address(lightInbox), gasOracle));
            Origin(origin).initialize(address(this));
        } else {
            revert("Unknown option: Origin");
        }
        vm.label(origin, "Origin Local");
    }

    function deployGasOracleSynapse() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_GAS_ORACLE_SYNAPSE;
        if (option == DEPLOY_MOCK_GAS_ORACLE_SYNAPSE) {
            gasOracleSynapse = address(new GasOracleMock());
        } else if (option == DEPLOY_PROD_GAS_ORACLE_SYNAPSE) {
            gasOracleSynapse = address(new GasOracle(DOMAIN_SYNAPSE, destinationSynapse));
            GasOracle(gasOracleSynapse).initialize(address(this));
        } else {
            revert("Unknown option: GasOracle");
        }
        vm.label(gasOracleSynapse, "GasOracle Synapse");
    }

    function deployDestinationSynapse() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_DESTINATION_SYNAPSE;
        if (option == DEPLOY_MOCK_DESTINATION_SYNAPSE) {
            destinationSynapse = address(new DestinationMock());
        } else if (option == DEPLOY_PROD_DESTINATION_SYNAPSE) {
            destinationSynapse = address(new Destination(DOMAIN_SYNAPSE, address(bondingManager), address(inbox)));
            Destination(destinationSynapse).initialize(0, address(this));
        } else {
            revert("Unknown option: Destination");
        }
        vm.label(destinationSynapse, "Destination Synapse");
    }

    function deployOriginSynapse() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_ORIGIN_SYNAPSE;
        if (option == DEPLOY_MOCK_ORIGIN_SYNAPSE) {
            originSynapse = address(new OriginMock());
        } else if (option == DEPLOY_PROD_ORIGIN_SYNAPSE) {
            originSynapse =
                address(new Origin(DOMAIN_SYNAPSE, address(bondingManager), address(inbox), gasOracleSynapse));
            Origin(originSynapse).initialize(address(this));
        } else {
            revert("Unknown option: Origin");
        }
        vm.label(originSynapse, "Origin Synapse");
    }

    function deploySummit() public virtual {
        vm.chainId(DOMAIN_SYNAPSE);

        uint256 option = deployMask & DEPLOY_MASK_SUMMIT;
        if (option == DEPLOY_MOCK_SUMMIT) {
            summit = address(new SummitMock());
        } else if (option == DEPLOY_PROD_SUMMIT) {
            summit = address(new Summit(DOMAIN_SYNAPSE, address(bondingManager), address(inbox)));
            Summit(summit).initialize(address(this));
        } else {
            revert("Unknown option: Summit");
        }
        vm.label(summit, "Summit");
    }

    // ═══════════════════════════════════════════════ ADDING AGENTS ═══════════════════════════════════════════════════

    function addAgentBM(uint32 domain, address agent) public {
        bytes32[] memory proof = getZeroProof();
        bondingManager.addAgent(domain, agent, proof);
        addNewAgent(domain, agent);
    }

    function updateAgentLM(address agent) public {
        bytes32[] memory proof = getAgentProof(agent);
        lightManager.updateAgentStatus(agent, getAgentStatus(agent), proof);
    }

    function checkAgentStatus(address agent, AgentStatus memory status, AgentFlag flag) public virtual {
        assertEq(uint8(status.flag), uint8(flag), "!flag");
        assertEq(status.domain, agentDomain[agent], "!domain");
        assertEq(status.index, agentIndex[agent], "!index");
    }
}
