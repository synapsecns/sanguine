// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { BondingMVP } from "../../contracts/bonding/BondingMVP.sol";
import { Destination } from "../../contracts/Destination.sol";
import { Origin } from "../../contracts/Origin.sol";
import { Summit } from "../../contracts/Summit.sol";

import "../harnesses/system/SystemRouterHarness.t.sol";

import "../mocks/DestinationMock.t.sol";
import "../mocks/OriginMock.t.sol";
import "../mocks/SummitMock.t.sol";

import "./events/ProductionEvents.t.sol";
import "./libs/SynapseUtilities.t.sol";
import "./SynapseTestConstants.t.sol";

import { Test } from "forge-std/Test.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";

// solhint-disable ordering
abstract contract SynapseTest is ProductionEvents, SynapseTestConstants, Test {
    struct Domain {
        string name;
        address agent;
        address[] agents;
    }

    uint256 private immutable deployMask;

    address internal destination;
    address internal origin;
    address internal summit;
    BondingMVP internal bondingManager;
    SystemRouterHarness internal systemRouter;

    // domain => Domain's name
    uint32[] internal allDomains;
    mapping(uint32 => Domain) internal domains;
    mapping(address => uint256) internal agentPK;

    constructor(uint256 _deployMask) {
        deployMask = _deployMask;
    }

    function setUp() public virtual {
        // Setup domains and create agents for them
        setupDomain(0, "Guards");
        setupDomain(DOMAIN_LOCAL, "Local");
        setupDomain(DOMAIN_REMOTE, "Remote");
        setupDomain(DOMAIN_SYNAPSE, "Synapse");
        // Deploy a single set of messaging contracts
        deployBondingMVP();
        deployDestination();
        deployOrigin();
        deploySummit();
        deploySystemRouter();
        // Setup agents on created contracts
        setupAgents();
        // Skip block
        skipBlock();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SETUP DOMAINS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function setupDomain(uint32 domain, string memory name) public virtual {
        allDomains.push(domain);
        domains[domain].name = name;
        string memory baseAgentName = domain == 0 ? "Guard" : "Notary";
        baseAgentName = string.concat(baseAgentName, "(", name, ") ");
        domains[domain].agents = new address[](DOMAIN_AGENTS);
        for (uint256 i = 0; i < DOMAIN_AGENTS; ++i) {
            domains[domain].agents[i] = createAgent(
                string.concat(baseAgentName, Strings.toString(i))
            );
        }
        domains[domain].agent = domains[domain].agents[0];
    }

    function setupAgents() public virtual {
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < DOMAIN_AGENTS; ++i) {
                bondingManager.addAgent(domain, domains[domain].agents[i]);
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                AGENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function createAgent(string memory name) public returns (address agent) {
        uint256 privKey;
        (agent, privKey) = makeAddrAndKey(name);
        agentPK[agent] = privKey;
    }

    function signMessage(uint256 privKey, bytes32 hashedMsg)
        public
        pure
        returns (bytes memory signature)
    {
        bytes32 digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hashedMsg));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, digest);
        signature = abi.encodePacked(r, s, v);
    }

    function signMessage(address agent, bytes32 hashedMsg)
        public
        view
        returns (bytes memory signature)
    {
        uint256 privKey = agentPK[agent];
        require(privKey != 0, "Unknown agent");
        return signMessage(privKey, hashedMsg);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           DEPLOY CONTRACTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function deployBondingMVP() public virtual {
        bondingManager = new BondingMVP(DOMAIN_LOCAL);
        bondingManager.initialize();
    }

    function deployDestination() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_DESTINATION;
        if (option == DEPLOY_MOCK_DESTINATION) {
            destination = address(new DestinationMock());
        } else if (option == DEPLOY_PROD_DESTINATION) {
            destination = address(new Destination(DOMAIN_LOCAL));
            Destination(destination).initialize();
        } else {
            revert("Unknown option: Destination");
        }
    }

    function deployOrigin() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_ORIGIN;
        if (option == DEPLOY_MOCK_ORIGIN) {
            origin = address(new OriginMock());
        } else if (option == DEPLOY_PROD_ORIGIN) {
            origin = address(new Origin(DOMAIN_LOCAL));
            Origin(origin).initialize();
        } else {
            revert("Unknown option: Origin");
        }
    }

    function deploySummit() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_SUMMIT;
        if (option == DEPLOY_MOCK_SUMMIT) {
            summit = address(new SummitMock());
        } else if (option == DEPLOY_PROD_SUMMIT) {
            summit = address(new Summit());
            Summit(summit).initialize();
            Summit(summit).transferOwnership(address(bondingManager));
        } else {
            revert("Unknown option: Summit");
        }
        // TODO: remove when Summit is merged with Bonding Primary
        bondingManager.setSummit(address(summit));
    }

    function deploySystemRouter() public virtual {
        systemRouter = new SystemRouterHarness(
            DOMAIN_LOCAL,
            address(origin),
            address(destination),
            address(bondingManager)
        );
        ISystemContract(origin).setSystemRouter(systemRouter);
        ISystemContract(destination).setSystemRouter(systemRouter);
        bondingManager.setSystemRouter(systemRouter);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               VM UTILS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function skipBlock() public {
        skipBlocks(1);
    }

    function skipBlocks(uint256 blocks) public {
        vm.roll(block.number + blocks);
        skip(blocks * 12 seconds);
    }
}
