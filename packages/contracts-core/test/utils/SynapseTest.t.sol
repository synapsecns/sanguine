// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {AgentFlag, BondingManager} from "../../contracts/manager/BondingManager.sol";
import {AgentStatus, LightManager} from "../../contracts/manager/LightManager.sol";
import {IAgentSecured} from "../../contracts/interfaces/IAgentSecured.sol";
import {Destination} from "../../contracts/Destination.sol";
import {Origin} from "../../contracts/Origin.sol";
import {Summit} from "../../contracts/Summit.sol";

import {BondingManagerHarness} from "../harnesses/manager/BondingManagerHarness.t.sol";
import {LightManagerHarness} from "../harnesses/manager/LightManagerHarness.t.sol";

import {DestinationMock} from "../mocks/DestinationMock.t.sol";
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

    address internal originSynapse;
    address internal summit; // Summit is Synapse Chain's Destination
    BondingManagerHarness internal bondingManager;

    address internal destination;
    address internal origin;
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
        deploySummit();
        deployOriginSynapse();
        // Setup agents in BondingManager
        initBondingManager();
        setupAgentsBM();
        // Deploy a single set of messaging contracts for local chain
        deployLightManager();
        deployDestination();
        deployOrigin();
        // Setup agents in LightManager
        initLightManager();
        setupAgentsLM();
        // Skip block
        skipBlock();
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
            Destination(destination).initialize(getAgentRoot());
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
        lightManager = new LightManagerHarness(DOMAIN_LOCAL);
        vm.label(address(lightManager), "LightManager");
    }

    function initLightManager() public virtual {
        lightManager.initialize(origin, destination);
    }

    function deployBondingManager() public virtual {
        bondingManager = new BondingManagerHarness(DOMAIN_SYNAPSE);
        vm.label(address(bondingManager), "BondingManager");
    }

    function initBondingManager() public virtual {
        bondingManager.initialize(originSynapse, summit);
    }

    function deployDestination() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_DESTINATION;
        if (option == DEPLOY_MOCK_DESTINATION) {
            destination = address(new DestinationMock());
        } else if (option == DEPLOY_PROD_DESTINATION) {
            destination = address(new Destination(DOMAIN_LOCAL, address(lightManager)));
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
            origin = address(new Origin(DOMAIN_LOCAL, address(lightManager)));
            Origin(origin).initialize();
        } else {
            revert("Unknown option: Origin");
        }
        vm.label(origin, "Origin Local");
    }

    function deployOriginSynapse() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_ORIGIN_SYNAPSE;
        if (option == DEPLOY_MOCK_ORIGIN_SYNAPSE) {
            originSynapse = address(new OriginMock());
        } else if (option == DEPLOY_PROD_ORIGIN_SYNAPSE) {
            originSynapse = address(new Origin(DOMAIN_LOCAL, address(bondingManager)));
            Origin(originSynapse).initialize();
        } else {
            revert("Unknown option: Origin");
        }
        vm.label(originSynapse, "Origin Synapse");
    }

    function deploySummit() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_SUMMIT;
        if (option == DEPLOY_MOCK_SUMMIT) {
            summit = address(new SummitMock());
        } else if (option == DEPLOY_PROD_SUMMIT) {
            summit = address(new Summit(DOMAIN_SYNAPSE, address(bondingManager)));
            Summit(summit).initialize();
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
