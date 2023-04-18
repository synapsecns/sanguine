// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {
    ATTESTATION_SALT,
    ATTESTATION_REPORT_SALT,
    RECEIPT_SALT,
    RECEIPT_REPORT_SALT,
    SNAPSHOT_SALT,
    STATE_REPORT_SALT
} from "../../contracts/libs/Constants.sol";

import {
    SnapshotLib,
    State,
    RawAttestation,
    RawAttestationReport,
    RawExecReceipt,
    RawReceiptReport,
    RawSnapshot,
    RawStateReport
} from "./libs/SynapseStructs.t.sol";

import {SynapseUtilities} from "./SynapseUtilities.t.sol";

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

// solhint-disable no-empty-blocks
// solhint-disable ordering
abstract contract SynapseAgents is SynapseUtilities {
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

    // ═══════════════════════════════════════════════════ SETUP ═══════════════════════════════════════════════════════

    function setupDomain(uint32 domain, string memory name) public virtual {
        allDomains.push(domain);
        domains[domain].name = name;
        string memory baseAgentName = domain == 0 ? "Guard" : string.concat("Notary(", name, ")");
        domains[domain].agents = new address[](DOMAIN_AGENTS);
        for (uint256 i = 0; i < DOMAIN_AGENTS; ++i) {
            domains[domain].agents[i] = createAgent(string.concat(baseAgentName, " ", Strings.toString(i)));
        }
        domains[domain].agent = domains[domain].agents[0];
    }

    // ══════════════════════════════════════════════════ AGENTS ═══════════════════════════════════════════════════════

    function createAgent(string memory name) public returns (address agent) {
        uint256 privKey;
        (agent, privKey) = makeAddrAndKey(name);
        agentPK[agent] = privKey;
    }

    function getAgent(uint256 domainId, uint256 agentId) public view returns (uint32 domain, address agent) {
        domain = allDomains[domainId % allDomains.length];
        agent = domains[domain].agents[agentId % DOMAIN_AGENTS];
    }

    /// @dev Private to enforce using salt-specific signing
    function signMessage(address agent, bytes32 hashedMsg) private view returns (bytes memory signature) {
        uint256 privKey = agentPK[agent];
        require(privKey != 0, "Unknown agent");
        return signMessage(privKey, hashedMsg);
    }

    /// @notice Signs a salted hash of a message
    function signMessage(address agent, bytes32 salt, bytes32 hashedMsg) public view returns (bytes memory signature) {
        return signMessage(agent, keccak256(bytes.concat(salt, hashedMsg)));
    }

    /// @notice Signs hashed message, by using a requested salt.
    function signMessage(address agent, bytes32 salt, bytes memory message)
        public
        view
        returns (bytes memory signature)
    {
        return signMessage(agent, salt, keccak256(message));
    }

    // ════════════════════════════════════════════ SIGNING STATEMENTS ═════════════════════════════════════════════════

    function signAttestation(address agent, bytes memory attestation) public view returns (bytes memory signature) {
        return signMessage(agent, ATTESTATION_SALT, attestation);
    }

    function signAttestation(address agent, RawAttestation memory ra)
        public
        view
        returns (bytes memory attestation, bytes memory signature)
    {
        attestation = ra.formatAttestation();
        signature = signAttestation(agent, attestation);
    }

    function signAttestationReport(address agent, bytes memory arPayload)
        public
        view
        returns (bytes memory signature)
    {
        return signMessage(agent, ATTESTATION_REPORT_SALT, arPayload);
    }

    function signAttestationReport(address agent, RawAttestationReport memory rawAR)
        public
        view
        returns (bytes memory attestationReport, bytes memory signature)
    {
        attestationReport = rawAR.formatAttestationReport();
        signature = signAttestationReport(agent, attestationReport);
    }

    function signReceipt(address agent, bytes memory receipt) public view returns (bytes memory signature) {
        return signMessage(agent, RECEIPT_SALT, receipt);
    }

    function signReceipt(address agent, RawExecReceipt memory re)
        public
        view
        returns (bytes memory receipt, bytes memory signature)
    {
        receipt = re.formatReceipt();
        signature = signReceipt(agent, receipt);
    }

    function signReceiptReport(address agent, bytes memory rrPayload) public view returns (bytes memory signature) {
        return signMessage(agent, RECEIPT_REPORT_SALT, rrPayload);
    }

    function signReceiptReport(address agent, RawReceiptReport memory rawRR)
        public
        view
        returns (bytes memory rrPayload, bytes memory signature)
    {
        rrPayload = rawRR.formatReceiptReport();
        signature = signReceiptReport(agent, rrPayload);
    }

    function signSnapshot(address agent, bytes memory snapshot) public view returns (bytes memory signature) {
        return signMessage(agent, SNAPSHOT_SALT, snapshot);
    }

    function signSnapshot(address agent, RawSnapshot memory rawSnap)
        public
        view
        returns (bytes memory snapshot, bytes memory signature)
    {
        snapshot = rawSnap.formatSnapshot();
        signature = signSnapshot(agent, snapshot);
    }

    function signSnapshot(address agent, State[] memory states)
        public
        view
        returns (bytes memory snapshot, bytes memory signature)
    {
        snapshot = SnapshotLib.formatSnapshot(states);
        signature = signSnapshot(agent, snapshot);
    }

    function signStateReport(address agent, bytes memory srPayload) public view returns (bytes memory signature) {
        return signMessage(agent, STATE_REPORT_SALT, srPayload);
    }

    function signStateReport(address agent, RawStateReport memory rawSR)
        public
        view
        returns (bytes memory stateReport, bytes memory signature)
    {
        stateReport = rawSR.formatStateReport();
        signature = signStateReport(agent, stateReport);
    }
}
