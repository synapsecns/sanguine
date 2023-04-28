// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/Attestation.sol";
import {AttestationReport, AttestationReportLib} from "../libs/AttestationReport.sol";
import {AGENT_TREE_HEIGHT, BONDING_OPTIMISTIC_PERIOD, SYNAPSE_DOMAIN} from "../libs/Constants.sol";
import {MerkleMath} from "../libs/MerkleMath.sol";
import {AgentFlag, AgentStatus, DisputeFlag} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager, IAgentSecured} from "./AgentManager.sol";
import {MessagingBase} from "../base/MessagingBase.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceDestination} from "../interfaces/InterfaceDestination.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";

/// @notice LightManager keeps track of all agents, staying in sync with the BondingManager.
/// Used on chains other than Synapse Chain, serves as "light client" for BondingManager.
contract LightManager is AgentManager, InterfaceLightManager {
    using AttestationLib for bytes;
    using AttestationReportLib for bytes;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════
    /// @inheritdoc IAgentManager
    bytes32 public agentRoot;

    // (agentRoot => (agent => status))
    mapping(bytes32 => mapping(address => AgentStatus)) private _agentMap;

    // (index => agent)
    mapping(uint256 => address) private _agents;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain) MessagingBase("0.0.3", domain) {
        require(domain != SYNAPSE_DOMAIN, "Can't be deployed on SynChain");
    }

    function initialize(address origin_, address destination_) external initializer {
        __AgentManager_init(origin_, destination_);
        __Ownable_init();
    }

    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function submitAttestation(bytes memory attPayload, bytes memory attSignature)
        external
        returns (bool wasAccepted)
    {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not an known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Check that Notary is active
        status.verifyActive();
        // Check that Notary domain is local domain
        require(status.domain == localDomain, "Wrong Notary domain");
        // Notary needs to be not in dispute
        require(_disputes[notary].flag == DisputeFlag.None, "Notary is in dispute");
        // Store Notary signature for the attestation
        uint256 sigIndex = _saveSignature(attSignature);
        wasAccepted = InterfaceDestination(destination).acceptAttestation(status.index, sigIndex, attPayload);
        if (wasAccepted) {
            emit AttestationAccepted(status.domain, notary, attPayload, attSignature);
        }
    }

    /// @inheritdoc InterfaceLightManager
    function submitAttestationReport(bytes memory arPayload, bytes memory arSignature, bytes memory attSignature)
        external
        returns (bool wasAccepted)
    {
        // This will revert if payload is not an attestation report
        AttestationReport report = arPayload.castToAttestationReport();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory guardStatus, address guard) = _verifyAttestationReport(report, arSignature);
        // Check that Guard is active
        guardStatus.verifyActive();
        // This will revert if attestation signer is not a known Notary
        (AgentStatus memory notaryStatus, address notary) = _verifyAttestation(report.attestation(), attSignature);
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // This will revert if either actor is already in dispute
        _openDispute(guard, guardStatus.index, notary, notaryStatus.index);
        return true;
    }

    // ═══════════════════════════════════════════════ AGENTS LOGIC ════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function updateAgentStatus(address agent, AgentStatus memory status, bytes32[] memory proof) external {
        address storedAgent = _agents[status.index];
        require(storedAgent == address(0) || storedAgent == agent, "Invalid agent index");
        // Reconstruct the agent leaf: flag should be Active
        bytes32 leaf = _agentLeaf(status.flag, status.domain, agent);
        bytes32 root = agentRoot;
        // Check that proof matches the latest merkle root
        require(MerkleMath.proofRoot(status.index, leaf, proof, AGENT_TREE_HEIGHT) == root, "Invalid proof");
        // Save index => agent in the map
        if (storedAgent == address(0)) _agents[status.index] = agent;
        // Update the agent status against this root
        _agentMap[root][agent] = status;
        emit StatusUpdated(status.flag, status.domain, agent);
        // Notify local AgentSecured contracts, if agent flag is Slashed
        if (status.flag == AgentFlag.Slashed) {
            // This will revert if the agent has been slashed earlier
            _resolveDispute(agent, status.index, msg.sender);
        }
    }

    /// @inheritdoc InterfaceLightManager
    function setAgentRoot(bytes32 agentRoot_) external {
        require(msg.sender == destination, "Only Destination sets agent root");
        _setAgentRoot(agentRoot_);
    }

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceLightManager
    function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount)
        external
        returns (bytes4 magicValue)
    {
        // Only destination can pass Manager Messages
        require(msg.sender == destination, "!destination");
        // Only AgentManager on Synapse Chain can give instructions to withdraw tips
        require(msgOrigin == SYNAPSE_DOMAIN, "!synapseDomain");
        // Check that merkle proof is mature enough
        require(proofMaturity >= BONDING_OPTIMISTIC_PERIOD, "!optimisticPeriod");
        InterfaceOrigin(origin).withdrawTips(recipient, amount);
        // Magic value to return is selector of the called function
        return this.remoteWithdrawTips.selector;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    function _afterAgentSlashed(uint32 domain, address agent, address prover) internal virtual override {
        // Send a manager message to BondingManager on SynChain
        // remoteSlashAgent(msgOrigin, proofMaturity, domain, agent, prover) with the first two security args omitted
        InterfaceOrigin(origin).sendManagerMessage({
            destination: SYNAPSE_DOMAIN,
            optimisticPeriod: BONDING_OPTIMISTIC_PERIOD,
            payload: abi.encodeWithSelector(InterfaceBondingManager.remoteSlashAgent.selector, domain, agent, prover)
        });
    }

    /// @dev Notify local AgentSecured contracts about the opened dispute.
    function _notifyDisputeOpened(uint32 guardIndex, uint32 notaryIndex) internal override {
        // Origin contract doesn't need to know about the dispute
        IAgentSecured(destination).openDispute(guardIndex, notaryIndex);
    }

    /// @dev Notify local AgentSecured contracts about the resolved dispute.
    function _notifyDisputeResolved(uint32 slashedIndex, uint32 rivalIndex) internal override {
        // Origin contract doesn't need to know about the dispute
        IAgentSecured(destination).resolveDispute(slashedIndex, rivalIndex);
    }

    /// @dev Updates the Agent Merkle Root that Light Manager is tracking.
    function _setAgentRoot(bytes32 _agentRoot) internal {
        if (agentRoot != _agentRoot) {
            agentRoot = _agentRoot;
            emit RootUpdated(_agentRoot);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the stored status for the agent: whether or not they have been added
    /// using latest Agent merkle Root.
    function _storedAgentStatus(address agent) internal view override returns (AgentStatus memory) {
        return _agentMap[agentRoot][agent];
    }

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes.
    function _getAgent(uint256 index) internal view override returns (address agent) {
        return _agents[index];
    }
}
