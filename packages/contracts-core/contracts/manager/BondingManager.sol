// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/Attestation.sol";
import {AttestationReport, AttestationReportLib} from "../libs/AttestationReport.sol";
import {BONDING_OPTIMISTIC_PERIOD, SYNAPSE_DOMAIN} from "../libs/Constants.sol";
import {DynamicTree, MerkleMath} from "../libs/MerkleTree.sol";
import {Receipt, ReceiptBody, ReceiptLib} from "../libs/Receipt.sol";
import {Snapshot, SnapshotLib} from "../libs/Snapshot.sol";
import {AgentFlag, AgentStatus, DisputeFlag} from "../libs/Structures.sol";
import {Tips} from "../libs/Tips.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentManager, IAgentManager, IAgentSecured} from "./AgentManager.sol";
import {MessagingBase} from "../base/MessagingBase.sol";
import {BondingManagerEvents} from "../events/BondingManagerEvents.sol";
import {InterfaceBondingManager} from "../interfaces/InterfaceBondingManager.sol";
import {InterfaceDestination} from "../interfaces/InterfaceDestination.sol";
import {IExecutionHub} from "../interfaces/IExecutionHub.sol";
import {InterfaceLightManager} from "../interfaces/InterfaceLightManager.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";
import {ISnapshotHub} from "../interfaces/ISnapshotHub.sol";
import {InterfaceSummit} from "../interfaces/InterfaceSummit.sol";

/// @notice BondingManager keeps track of all existing _agents.
/// Used on the Synapse Chain, serves as the "source of truth" for LightManagers on remote chains.
contract BondingManager is AgentManager, BondingManagerEvents, InterfaceBondingManager {
    using AttestationLib for bytes;
    using AttestationReportLib for bytes;
    using ReceiptLib for bytes;
    using SnapshotLib for bytes;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // The address of the Summit contract.
    address public summit;

    // (agent => their status)
    mapping(address => AgentStatus) private _agentMap;

    // (domain => past and current agents for domain)
    mapping(uint32 => address[]) private _domainAgents;

    // A list of all agent accounts. First entry is address(0) to make agent indexes start from 1.
    address[] private _agents;

    // Merkle Tree for Agents.
    // leafs[0] = 0
    // leafs[index > 0] = keccak(agentFlag, domain, _agents[index])
    DynamicTree private _agentTree;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain) MessagingBase("0.0.3", domain) {
        require(domain == SYNAPSE_DOMAIN, "Only deployed on SynChain");
    }

    function initialize(address origin_, address destination_, address summit_) external initializer {
        __AgentManager_init(origin_, destination_);
        summit = summit_;
        __Ownable_init();
        // Insert a zero address to make indexes for Agents start from 1.
        // Zeroed index is supposed to be used as a sentinel value meaning "no agent".
        _agents.push(address(0));
    }

    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function submitSnapshot(bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bytes memory attPayload)
    {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the signer is not a known Agent
        (AgentStatus memory status, address agent) = _verifySnapshot(snapshot, snapSignature);
        // Check that Agent is active
        status.verifyActive();
        // Store Agent signature for the Snapshot
        uint256 sigIndex = _saveSignature(snapSignature);
        if (status.domain == 0) {
            // Guard that is in Dispute could still submit new snapshots, so we don't check that
            InterfaceSummit(summit).acceptGuardSnapshot({
                guardIndex: status.index,
                sigIndex: sigIndex,
                snapPayload: snapPayload
            });
        } else {
            // Check that Notary is not in dispute
            require(_disputes[agent].flag == DisputeFlag.None, "Notary is in dispute");
            attPayload = InterfaceSummit(summit).acceptNotarySnapshot({
                notaryIndex: status.index,
                sigIndex: sigIndex,
                agentRoot: _agentTree.root,
                snapPayload: snapPayload
            });
            // Pass created attestation to Destination to enable executing messages coming to Synapse Chain
            InterfaceDestination(destination).acceptAttestation(status.index, type(uint256).max, attPayload);
        }
        emit SnapshotAccepted(status.domain, agent, snapPayload, snapSignature);
    }

    /// @inheritdoc InterfaceBondingManager
    function submitReceipt(bytes memory rcptPayload, bytes memory rcptSignature) external returns (bool wasAccepted) {
        // This will revert if payload is not a receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the receipt signer is not a known Notary
        (AgentStatus memory rcptNotaryStatus, address notary) = _verifyReceipt(rcpt, rcptSignature);
        // Receipt Notary needs to be Active and not in dispute
        rcptNotaryStatus.verifyActive();
        require(_disputes[notary].flag == DisputeFlag.None, "Notary is in dispute");
        // Check that receipt's snapshot root exists in Summit
        ReceiptBody rcptBody = rcpt.body();
        uint32 attNonce = IExecutionHub(destination).getAttestationNonce(rcptBody.snapshotRoot());
        require(attNonce != 0, "Unknown snapshot root");
        // Attestation Notary domain needs to match the destination domain
        AgentStatus memory attNotaryStatus = agentStatus(rcptBody.attNotary());
        require(attNotaryStatus.domain == rcptBody.destination(), "Wrong attestation Notary domain");
        // Store Notary signature for the Receipt
        uint256 sigIndex = _saveSignature(rcptSignature);
        wasAccepted = InterfaceSummit(summit).acceptReceipt({
            rcptNotaryIndex: rcptNotaryStatus.index,
            attNotaryIndex: attNotaryStatus.index,
            sigIndex: sigIndex,
            attNonce: attNonce,
            paddedTips: Tips.unwrap(rcpt.tips()),
            rcptBodyPayload: rcptBody.unwrap().clone()
        });
        if (wasAccepted) {
            emit ReceiptAccepted(rcptNotaryStatus.domain, notary, rcptPayload, rcptSignature);
        }
    }

    /// @inheritdoc InterfaceBondingManager
    function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes memory rcptBodyPayload)
        external
        returns (bool wasAccepted)
    {
        require(msg.sender == destination, "Only Destination passes receipts");
        return InterfaceSummit(summit).acceptReceipt({
            rcptNotaryIndex: attNotaryIndex,
            attNotaryIndex: attNotaryIndex,
            sigIndex: type(uint256).max,
            attNonce: attNonce,
            paddedTips: paddedTips,
            rcptBodyPayload: rcptBodyPayload
        });
    }

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function verifyAttestation(bytes memory attPayload, bytes memory attSignature)
        external
        returns (bool isValidAttestation)
    {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        isValidAttestation = ISnapshotHub(summit).isValidAttestation(attPayload);
        if (!isValidAttestation) {
            emit InvalidAttestation(attPayload, attSignature);
            _slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc InterfaceBondingManager
    function verifyAttestationReport(bytes memory arPayload, bytes memory arSignature)
        external
        returns (bool isValidReport)
    {
        // This will revert if payload is not an attestation report
        AttestationReport report = arPayload.castToAttestationReport();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyAttestationReport(report, arSignature);
        // Guard needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // Report is valid IF AND ONLY IF the reported attestation in invalid
        isValidReport = !ISnapshotHub(summit).isValidAttestation(report.attestation().unwrap().clone());
        if (!isValidReport) {
            emit InvalidAttestationReport(arPayload, arSignature);
            _slashAgent(status.domain, guard, msg.sender);
        }
    }

    // ════════════════════════════════════════════ AGENTS LOGIC (MVP) ═════════════════════════════════════════════════

    // TODO: remove these MVP functions once token staking is implemented

    /// @inheritdoc InterfaceBondingManager
    function addAgent(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        // Check the STORED status of the added agent in the merkle tree
        AgentStatus memory status = _storedAgentStatus(agent);
        // Agent index in `_agents`
        uint32 index;
        // Leaf representing currently saved agent information in the tree
        bytes32 oldValue;
        if (status.flag == AgentFlag.Unknown) {
            // Unknown address could be added to any domain
            // New agent will need to be added to `_agents` list
            require(_agents.length < type(uint32).max, "Agents list if full");
            index = uint32(_agents.length);
            // Current leaf for index is bytes32(0), which is already assigned to `leaf`
            _agents.push(agent);
            _domainAgents[domain].push(agent);
        } else if (status.flag == AgentFlag.Resting && status.domain == domain) {
            // Resting agent could be only added back to the same domain
            // Agent is already in `_agents`, fetch the saved index
            index = status.index;
            // Generate the current leaf for the agent
            // oldValue includes the domain information, so we didn't had to check it above.
            // However, we are still doing this check to have a more appropriate revert string,
            // if a resting agent is requesting to be added to another domain.
            oldValue = _agentLeaf(AgentFlag.Resting, domain, agent);
        } else {
            // Any other flag indicates that agent could not be added
            revert("Agent could not be added");
        }
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Active, domain, index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function initiateUnstaking(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        // Check the CURRENT status of the unstaking agent
        AgentStatus memory status = agentStatus(agent);
        // Could only initiate the unstaking for the active agent for the domain
        require(status.flag == AgentFlag.Active && status.domain == domain, "Unstaking could not be initiated");
        // Leaf representing currently saved agent information in the tree.
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if an agent is initiating the unstaking, but specifies incorrect domain.
        bytes32 oldValue = _agentLeaf(AgentFlag.Active, domain, agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Unstaking, domain, status.index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function completeUnstaking(uint32 domain, address agent, bytes32[] memory proof) external onlyOwner {
        // Check the CURRENT status of the unstaking agent
        AgentStatus memory status = agentStatus(agent);
        // Could only complete the unstaking, if it was previously initiated
        // TODO: add more checks (time-based, possibly collecting info from other chains)
        require(status.flag == AgentFlag.Unstaking && status.domain == domain, "Unstaking could not be completed");
        // Leaf representing currently saved agent information in the tree
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if an agent is completing the unstaking, but specifies incorrect domain.
        bytes32 oldValue = _agentLeaf(AgentFlag.Unstaking, domain, agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Resting, domain, status.index), agent);
    }

    // ══════════════════════════════════════════════ SLASHING LOGIC ═══════════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function completeSlashing(uint32 domain, address agent, bytes32[] memory proof) external {
        // Check that slashing was previously initiated in AgentManager
        require(_disputes[agent].flag == DisputeFlag.Slashed, "Slashing not initiated");
        // Check that the STORED status is Active/Unstaking in the merkle tree and that the domains match
        AgentStatus memory status = _storedAgentStatus(agent);
        require(
            (status.flag == AgentFlag.Active || status.flag == AgentFlag.Unstaking) && status.domain == domain,
            "Slashing could not be completed"
        );
        // Leaf representing currently saved agent information in the tree
        // oldValue includes the domain information, so we didn't had to check it above.
        // However, we are still doing this check to have a more appropriate revert string,
        // if anyone is completing the slashing, but specifies incorrect domain.
        bytes32 oldValue = _agentLeaf(status.flag, domain, agent);
        // This will revert if the proof for the old value is incorrect
        _updateLeaf(oldValue, proof, AgentStatus(AgentFlag.Slashed, domain, status.index), agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover)
        external
        returns (bytes4 magicValue)
    {
        // Only destination can pass Manager Messages
        require(msg.sender == destination, "!destination");
        // Check that merkle proof is mature enough
        require(proofMaturity >= BONDING_OPTIMISTIC_PERIOD, "!optimisticPeriod");
        // TODO: do we need to save this?
        msgOrigin;
        // Slash agent and notify local AgentSecured contracts
        _slashAgent(domain, agent, prover);
        // Magic value to return is selector of the called function
        return this.remoteSlashAgent.selector;
    }

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceBondingManager
    function withdrawTips(address recipient, uint32 origin_, uint256 amount) external {
        require(msg.sender == summit, "Only Summit withdraws tips");
        if (origin_ == localDomain) {
            // Call local Origin to withdraw tips
            InterfaceOrigin(address(origin)).withdrawTips(recipient, amount);
        } else {
            // For remote chains: send a manager message to remote LightManager to handle the withdrawal
            // remoteWithdrawTips(msgOrigin, proofMaturity, recipient, amount) with the first two security args omitted
            InterfaceOrigin(origin).sendManagerMessage({
                destination: origin_,
                optimisticPeriod: BONDING_OPTIMISTIC_PERIOD,
                payload: abi.encodeWithSelector(InterfaceLightManager.remoteWithdrawTips.selector, recipient, amount)
            });
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentManager
    function agentRoot() external view override returns (bytes32) {
        return _agentTree.root;
    }

    /// @inheritdoc InterfaceBondingManager
    function getActiveAgents(uint32 domain) external view returns (address[] memory agents) {
        uint256 amount = _domainAgents[domain].length;
        agents = new address[](amount);
        uint256 activeAgents = 0;
        for (uint256 i = 0; i < amount; ++i) {
            address agent = _domainAgents[domain][i];
            if (agentStatus(agent).flag == AgentFlag.Active) {
                agents[activeAgents++] = agent;
            }
        }
        if (activeAgents != amount) {
            // Shrink the returned array by storing the required length in memory
            // solhint-disable-next-line no-inline-assembly
            assembly {
                mstore(agents, activeAgents)
            }
        }
    }

    /// @inheritdoc InterfaceBondingManager
    function agentLeaf(address agent) external view returns (bytes32 leaf) {
        return _getLeaf(agent);
    }

    /// @inheritdoc InterfaceBondingManager
    function leafsAmount() external view returns (uint256 amount) {
        return _agents.length;
    }

    /// @inheritdoc InterfaceBondingManager
    function getProof(address agent) external view returns (bytes32[] memory proof) {
        bytes32[] memory leafs = allLeafs();
        // Use the STORED agent status from the merkle tree
        AgentStatus memory status = _storedAgentStatus(agent);
        // Use next available index for unknown agents
        uint256 index = status.flag == AgentFlag.Unknown ? _agents.length : status.index;
        return MerkleMath.calculateProof(leafs, index);
    }

    /// @inheritdoc InterfaceBondingManager
    function allLeafs() public view returns (bytes32[] memory leafs) {
        return getLeafs(0, _agents.length);
    }

    /// @inheritdoc InterfaceBondingManager
    function getLeafs(uint256 indexFrom, uint256 amount) public view returns (bytes32[] memory leafs) {
        uint256 amountTotal = _agents.length;
        require(indexFrom < amountTotal, "Out of range");
        if (indexFrom + amount > amountTotal) {
            amount = amountTotal - indexFrom;
        }
        leafs = new bytes32[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            leafs[i] = _getLeaf(indexFrom + i);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Updates value in the Agent Merkle Tree to reflect the `newStatus`.
    /// Will revert, if supplied proof for the old value is incorrect.
    function _updateLeaf(bytes32 oldValue, bytes32[] memory proof, AgentStatus memory newStatus, address agent)
        internal
    {
        // New leaf value for the agent in the Agent Merkle Tree
        bytes32 newValue = _agentLeaf(newStatus.flag, newStatus.domain, agent);
        // This will revert if the proof for the old value is incorrect
        bytes32 newRoot = _agentTree.update(newStatus.index, oldValue, proof, newValue);
        _agentMap[agent] = newStatus;
        emit StatusUpdated(newStatus.flag, newStatus.domain, agent);
        emit RootUpdated(newRoot);
    }

    /// @dev Notify local AgentSecured contracts about the opened dispute.
    function _notifyDisputeOpened(uint32 guardIndex, uint32 notaryIndex) internal override {
        IAgentSecured(destination).openDispute(guardIndex, notaryIndex);
        IAgentSecured(summit).openDispute(guardIndex, notaryIndex);
    }

    /// @dev Notify local AgentSecured contracts about the resolved dispute.
    function _notifyDisputeResolved(uint32 slashedIndex, uint32 rivalIndex) internal override {
        IAgentSecured(destination).resolveDispute(slashedIndex, rivalIndex);
        IAgentSecured(summit).resolveDispute(slashedIndex, rivalIndex);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the status of the agent.
    function _storedAgentStatus(address agent) internal view override returns (AgentStatus memory) {
        return _agentMap[agent];
    }

    /// @dev Returns agent address for the given index. Returns zero for non existing indexes.
    function _getAgent(uint256 index) internal view override returns (address agent) {
        if (index < _agents.length) {
            agent = _agents[index];
        }
    }

    /// @dev Returns the current leaf representing agent in the Agent Merkle Tree.
    function _getLeaf(address agent) internal view returns (bytes32 leaf) {
        // Get the agent status STORED in the merkle tree
        AgentStatus memory status = _storedAgentStatus(agent);
        if (status.flag != AgentFlag.Unknown) {
            return _agentLeaf(status.flag, status.domain, agent);
        }
        // Return empty leaf for unknown _agents
    }

    /// @dev Returns a leaf from the Agent Merkle Tree with a given index.
    function _getLeaf(uint256 index) internal view returns (bytes32 leaf) {
        if (index != 0) {
            return _getLeaf(_agents[index]);
        }
        // Return empty leaf for a zero index
    }
}
