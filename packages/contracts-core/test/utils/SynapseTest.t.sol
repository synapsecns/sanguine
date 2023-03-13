// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    ATTESTATION_SALT,
    ATTESTATION_REPORT_SALT,
    SNAPSHOT_SALT,
    STATE_REPORT_SALT
} from "../../contracts/libs/Constants.sol";

import { BondingSecondary } from "../../contracts/bonding/BondingSecondary.sol";
import { ISystemContract } from "../../contracts/interfaces/ISystemContract.sol";
import { Destination } from "../../contracts/Destination.sol";
import { Origin } from "../../contracts/Origin.sol";
import { Summit } from "../../contracts/Summit.sol";

import { SystemRouterHarness } from "../harnesses/system/SystemRouterHarness.t.sol";

import { DestinationMock } from "../mocks/DestinationMock.t.sol";
import { OriginMock } from "../mocks/OriginMock.t.sol";
import { SummitMock } from "../mocks/SummitMock.t.sol";

import { ProductionEvents } from "./events/ProductionEvents.t.sol";
// import "./libs/SynapseUtilities.t.sol";
import { SynapseTestConstants } from "./SynapseTestConstants.t.sol";

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

    address internal destinationSynapse;
    address internal originSynapse;
    address internal summit;
    SystemRouterHarness internal systemRouterSynapse;

    address internal destination;
    address internal origin;
    BondingSecondary internal bondingManager;

    SystemRouterHarness internal systemRouter;

    // domain => Domain's name
    uint32[] internal allDomains;
    mapping(uint32 => Domain) internal domains;
    mapping(address => uint256) internal agentPK;

    constructor(uint256 _deployMask) {
        deployMask = _deployMask;
    }

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseTest() external {}

    function setUp() public virtual {
        // Setup domains and create agents for them
        setupDomain(0, "Guards");
        setupDomain(DOMAIN_LOCAL, "Local");
        setupDomain(DOMAIN_REMOTE, "Remote");
        setupDomain(DOMAIN_SYNAPSE, "Synapse");
        // Deploy a single set of messaging contracts for local chain
        deployBondingS();
        deployDestination();
        deployOrigin();
        deploySystemRouter();
        // Deploy a single set of messaging contracts for synapse chain
        deploySummit();
        deployDestinationSynapse();
        deployOriginSynapse();
        deploySystemRouterSynapse();
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
                address agent = domains[domain].agents[i];
                bondingManager.addAgent(domain, agent);
                Summit(summit).addAgent(domain, agent);
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

    /// @dev Private to enforce using salt-specific signing
    function signMessage(address agent, bytes32 hashedMsg)
        private
        view
        returns (bytes memory signature)
    {
        uint256 privKey = agentPK[agent];
        require(privKey != 0, "Unknown agent");
        return signMessage(privKey, hashedMsg);
    }

    function signAttestation(address agent, bytes memory attestation)
        public
        view
        returns (bytes memory signature)
    {
        bytes32 hashedAtt = keccak256(attestation);
        return signMessage(agent, keccak256(bytes.concat(ATTESTATION_SALT, hashedAtt)));
    }

    function signAttestationReport(address agent, bytes memory arPayload)
        public
        view
        returns (bytes memory signature)
    {
        bytes32 hashedAR = keccak256(arPayload);
        return signMessage(agent, keccak256(bytes.concat(ATTESTATION_REPORT_SALT, hashedAR)));
    }

    function signSnapshot(address agent, bytes memory snapshot)
        public
        view
        returns (bytes memory signature)
    {
        bytes32 hashedSnap = keccak256(snapshot);
        return signMessage(agent, keccak256(bytes.concat(SNAPSHOT_SALT, hashedSnap)));
    }

    function signStateReport(address agent, bytes memory srPayload)
        public
        view
        returns (bytes memory signature)
    {
        bytes32 hashedAR = keccak256(srPayload);
        return signMessage(agent, keccak256(bytes.concat(STATE_REPORT_SALT, hashedAR)));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           DEPLOY CONTRACTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function deployBondingS() public virtual {
        bondingManager = new BondingSecondary(DOMAIN_LOCAL);
        bondingManager.initialize();
        vm.label(address(bondingManager), "BondingSecondary");
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
        vm.label(destination, "Destination Local");
    }

    function deployDestinationSynapse() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_DESTINATION_SYNAPSE;
        if (option == DEPLOY_MOCK_DESTINATION_SYNAPSE) {
            destinationSynapse = address(new DestinationMock());
        } else if (option == DEPLOY_PROD_DESTINATION_SYNAPSE) {
            destinationSynapse = address(new Destination(DOMAIN_LOCAL));
            Destination(destinationSynapse).initialize();
        } else {
            revert("Unknown option: Destination");
        }
        vm.label(destinationSynapse, "Destination Synapse");
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
        vm.label(origin, "Origin Local");
    }

    function deployOriginSynapse() public virtual {
        uint256 option = deployMask & DEPLOY_MASK_ORIGIN_SYNAPSE;
        if (option == DEPLOY_MOCK_ORIGIN_SYNAPSE) {
            originSynapse = address(new OriginMock());
        } else if (option == DEPLOY_PROD_ORIGIN_SYNAPSE) {
            originSynapse = address(new Origin(DOMAIN_LOCAL));
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
            summit = address(new Summit(DOMAIN_SYNAPSE));
            Summit(summit).initialize();
        } else {
            revert("Unknown option: Summit");
        }
        vm.label(summit, "Summit");
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
        vm.label(address(systemRouter), "SystemRouter Local");
    }

    function deploySystemRouterSynapse() public virtual {
        systemRouterSynapse = new SystemRouterHarness(
            DOMAIN_SYNAPSE,
            address(originSynapse),
            address(destinationSynapse),
            address(summit)
        );
        ISystemContract(originSynapse).setSystemRouter(systemRouterSynapse);
        ISystemContract(destinationSynapse).setSystemRouter(systemRouterSynapse);
        ISystemContract(summit).setSystemRouter(systemRouterSynapse);
        vm.label(address(systemRouterSynapse), "SystemRouter Synapse");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               VM UTILS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectRevertNotOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
    }

    function skipBlock() public {
        skipBlocks(1);
    }

    function skipBlocks(uint256 blocks) public {
        vm.roll(block.number + blocks);
        skip(blocks * 12 seconds);
    }
}
