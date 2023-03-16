// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    ATTESTATION_SALT,
    ATTESTATION_REPORT_SALT,
    SNAPSHOT_SALT,
    STATE_REPORT_SALT
} from "../../contracts/libs/Constants.sol";

import { SynapseTestConstants } from "./SynapseTestConstants.t.sol";

import { Test } from "forge-std/Test.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";

abstract contract SynapseAgents is SynapseTestConstants, Test {
    struct Domain {
        string name;
        address agent;
        address[] agents;
    }

    // domain => Domain's name
    uint32[] internal allDomains;
    mapping(uint32 => Domain) internal domains;
    mapping(address => uint256) internal agentPK;

    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseAgents() external {}

    function setUp() public virtual {
        // Setup domains and create agents for them
        setupDomain(0, "Guards");
        setupDomain(DOMAIN_LOCAL, "Local");
        setupDomain(DOMAIN_REMOTE, "Remote");
        setupDomain(DOMAIN_SYNAPSE, "Synapse");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SETUP DOMAINS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function setupDomain(uint32 domain, string memory name) public virtual {
        allDomains.push(domain);
        domains[domain].name = name;
        string memory baseAgentName = domain == 0 ? "Guard" : string.concat("Notary(", name, ")");
        domains[domain].agents = new address[](DOMAIN_AGENTS);
        for (uint256 i = 0; i < DOMAIN_AGENTS; ++i) {
            domains[domain].agents[i] = createAgent(
                string.concat(baseAgentName, " ", Strings.toString(i))
            );
        }
        domains[domain].agent = domains[domain].agents[0];
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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SIGNING STATEMENTS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

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
}
