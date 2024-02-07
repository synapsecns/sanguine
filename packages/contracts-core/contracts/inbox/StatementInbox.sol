// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/memory/Attestation.sol";
import {
    AgentNotGuard,
    AgentNotNotary,
    IncorrectAgentDomain,
    IncorrectSnapshotProof,
    IncorrectSnapshotRoot,
    IncorrectState,
    IndexOutOfRange
} from "../libs/Errors.sol";
import {Receipt, ReceiptLib} from "../libs/memory/Receipt.sol";
import {Snapshot, SnapshotLib} from "../libs/memory/Snapshot.sol";
import {State, StateLib} from "../libs/memory/State.sol";
import {AgentStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {MessagingBase} from "../base/MessagingBase.sol";
import {StatementInboxEvents} from "../events/StatementInboxEvents.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {IExecutionHub} from "../interfaces/IExecutionHub.sol";
import {IStateHub} from "../interfaces/IStateHub.sol";
import {IStatementInbox} from "../interfaces/IStatementInbox.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/// @notice `StatementInbox` is the entry point for all agent-signed statements. It verifies the
/// agent signatures, and passes the unsigned statements to the contract to consume it via `acceptX` functions. Is is
/// also used to verify the agent-signed statements and initiate the agent slashing, should the statement be invalid.
/// `StatementInbox` is responsible for the following:
/// - Accepting State and Receipt Reports to initiate a dispute between Guard and Notary.
/// - Storing all the Guard Reports with the Guard signature leading to a dispute.
/// - Verifying State/State Reports referencing the local chain and slashing the signer if statement is invalid.
/// - Verifying Receipt/Receipt Reports referencing the local chain and slashing the signer if statement is invalid.
abstract contract StatementInbox is MessagingBase, StatementInboxEvents, IStatementInbox {
    using AttestationLib for bytes;
    using ReceiptLib for bytes;
    using StateLib for bytes;
    using SnapshotLib for bytes;

    struct StoredReport {
        uint256 sigIndex;
        bytes statementPayload;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    address public agentManager;
    address public origin;
    address public destination;

    // TODO: optimize this
    bytes[] internal _storedSignatures;
    StoredReport[] internal _storedReports;

    /// @dev gap for upgrade safety
    uint256[45] private __GAP; // solhint-disable-line var-name-mixedcase

    // ════════════════════════════════════════════════ INITIALIZER ════════════════════════════════════════════════════

    /// @dev Initializes the contract:
    /// - Sets up `msg.sender` as the owner of the contract.
    /// - Sets up `agentManager`, `origin`, and `destination`.
    // solhint-disable-next-line func-name-mixedcase
    function __StatementInbox_init(address agentManager_, address origin_, address destination_)
        internal
        onlyInitializing
    {
        agentManager = agentManager_;
        origin = origin_;
        destination = destination_;
        __Ownable2Step_init();
    }

    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc IStatementInbox
    // solhint-disable-next-line ordering
    function submitStateReportWithSnapshot(
        uint8 stateIndex,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the snapshot signer is not a known Notary
        (AgentStatus memory notaryStatus,) =
            _verifySnapshot({snapshot: snapshot, snapSignature: snapSignature, verifyNotary: true});
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(notaryStatus.domain);
        // This will revert if state index is out of range
        State state = snapshot.state(stateIndex);
        // This will revert if the report signer is not an known Guard
        (AgentStatus memory guardStatus,) = _verifyStateReport(state, srSignature);
        // Check that Guard is active
        guardStatus.verifyActive();
        _saveReport(state.unwrap().clone(), srSignature);
        // This will revert if either actor is already in dispute
        IAgentManager(agentManager).openDispute(guardStatus.index, notaryStatus.index);
        return true;
    }

    /// @inheritdoc IStatementInbox
    function submitStateReportWithAttestation(
        uint8 stateIndex,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if state index is out of range
        State state = snapshot.state(stateIndex);
        // This will revert if the report signer is not an known Guard
        (AgentStatus memory guardStatus,) = _verifyStateReport(state, srSignature);
        // Check that Guard is active
        guardStatus.verifyActive();
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not an known Notary
        (AgentStatus memory notaryStatus,) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(notaryStatus.domain);
        if (snapshot.calculateRoot() != att.snapRoot()) revert IncorrectSnapshotRoot();
        _saveReport(state.unwrap().clone(), srSignature);
        // This will revert if either actor is already in dispute
        IAgentManager(agentManager).openDispute(guardStatus.index, notaryStatus.index);
        return true;
    }

    /// @inheritdoc IStatementInbox
    function submitStateReportWithSnapshotProof(
        uint8 stateIndex,
        bytes memory statePayload,
        bytes memory srSignature,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not a state
        State state = statePayload.castToState();
        // This will revert if the report signer is not an known Guard
        (AgentStatus memory guardStatus,) = _verifyStateReport(state, srSignature);
        // Check that Guard is active
        guardStatus.verifyActive();
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if signer is not a known Notary
        (AgentStatus memory notaryStatus,) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        // Check if Notary is active on this chain
        _verifyNotaryDomain(notaryStatus.domain);
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, stateIndex, state, snapProof);
        _saveReport(statePayload, srSignature);
        // This will revert if either actor is already in dispute
        IAgentManager(agentManager).openDispute(guardStatus.index, notaryStatus.index);
        return true;
    }

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc IStatementInbox
    function verifyReceipt(bytes memory rcptPayload, bytes memory rcptSignature)
        external
        returns (bool isValidReceipt)
    {
        // This will revert if payload is not a receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyReceipt(rcpt, rcptSignature);
        // Notary needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        isValidReceipt = IExecutionHub(destination).isValidReceipt(rcptPayload);
        if (!isValidReceipt) {
            emit InvalidReceipt(rcptPayload, rcptSignature);
            IAgentManager(agentManager).slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc IStatementInbox
    function verifyReceiptReport(bytes memory rcptPayload, bytes memory rrSignature)
        external
        returns (bool isValidReport)
    {
        // This will revert if payload is not a receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyReceiptReport(rcpt, rrSignature);
        // Guard needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // Report is valid IF AND ONLY IF the reported receipt in invalid
        isValidReport = !IExecutionHub(destination).isValidReceipt(rcptPayload);
        if (!isValidReport) {
            emit InvalidReceiptReport(rcptPayload, rrSignature);
            IAgentManager(agentManager).slashAgent(status.domain, guard, msg.sender);
        }
    }

    /// @inheritdoc IStatementInbox
    function verifyStateWithAttestation(
        uint8 stateIndex,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState) {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        if (snapshot.calculateRoot() != att.snapRoot()) revert IncorrectSnapshotRoot();
        // This will revert if state does not refer to this chain
        bytes memory statePayload = snapshot.state(stateIndex).unwrap().clone();
        isValidState = IStateHub(origin).isValidState(statePayload);
        if (!isValidState) {
            emit InvalidStateWithAttestation(stateIndex, statePayload, attPayload, attSignature);
            IAgentManager(agentManager).slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc IStatementInbox
    function verifyStateWithSnapshotProof(
        uint8 stateIndex,
        bytes memory statePayload,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState) {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if the attestation signer is not a known Notary
        (AgentStatus memory status, address notary) = _verifyAttestation(att, attSignature);
        // Notary needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // This will revert if payload is not a state
        State state = statePayload.castToState();
        // This will revert if any of these is true:
        //  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
        //  - Snapshot Proof's first element does not match the State metadata.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        //  - State index is out of range.
        _verifySnapshotMerkle(att, stateIndex, state, snapProof);
        // This will revert if state does not refer to this chain
        isValidState = IStateHub(origin).isValidState(statePayload);
        if (!isValidState) {
            emit InvalidStateWithAttestation(stateIndex, statePayload, attPayload, attSignature);
            IAgentManager(agentManager).slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc IStatementInbox
    function verifyStateWithSnapshot(uint8 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValidState)
    {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the snapshot signer is not a known Guard/Notary
        (AgentStatus memory status, address agent) =
            _verifySnapshot({snapshot: snapshot, snapSignature: snapSignature, verifyNotary: false});
        // Agent needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // This will revert if state does not refer to this chain
        isValidState = IStateHub(origin).isValidState(snapshot.state(stateIndex).unwrap().clone());
        if (!isValidState) {
            emit InvalidStateWithSnapshot(stateIndex, snapPayload, snapSignature);
            IAgentManager(agentManager).slashAgent(status.domain, agent, msg.sender);
        }
    }

    /// @inheritdoc IStatementInbox
    function verifyStateReport(bytes memory statePayload, bytes memory srSignature)
        external
        returns (bool isValidReport)
    {
        // This will revert if payload is not a state
        State state = statePayload.castToState();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyStateReport(state, srSignature);
        // Guard needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // Report is valid IF AND ONLY IF the reported state in invalid
        // This will revert if the reported state does not refer to this chain
        isValidReport = !IStateHub(origin).isValidState(statePayload);
        if (!isValidReport) {
            emit InvalidStateReport(statePayload, srSignature);
            IAgentManager(agentManager).slashAgent(status.domain, guard, msg.sender);
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IStatementInbox
    function getReportsAmount() external view returns (uint256) {
        return _storedReports.length;
    }

    /// @inheritdoc IStatementInbox
    function getGuardReport(uint256 index)
        external
        view
        returns (bytes memory statementPayload, bytes memory reportSignature)
    {
        if (index >= _storedReports.length) revert IndexOutOfRange();
        StoredReport memory storedReport = _storedReports[index];
        statementPayload = storedReport.statementPayload;
        reportSignature = _storedSignatures[storedReport.sigIndex];
    }

    /// @inheritdoc IStatementInbox
    function getStoredSignature(uint256 index) external view returns (bytes memory) {
        return _storedSignatures[index];
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Saves the statement reported by Guard as invalid and the Guard Report signature.
    function _saveReport(bytes memory statementPayload, bytes memory reportSignature) internal {
        uint256 sigIndex = _saveSignature(reportSignature);
        _storedReports.push(StoredReport(sigIndex, statementPayload));
    }

    /// @dev Saves the signature and returns its index.
    function _saveSignature(bytes memory signature) internal returns (uint256 sigIndex) {
        sigIndex = _storedSignatures.length;
        _storedSignatures.push(signature);
    }

    // ═══════════════════════════════════════════════ AGENT CHECKS ════════════════════════════════════════════════════

    /**
     * @dev Recovers a signer from a hashed message, and a EIP-191 signature for it.
     * Will revert, if the signer is not a known agent.
     * @dev Agent flag could be any of these: Active/Unstaking/Resting/Fraudulent/Slashed
     * Further checks need to be performed in a caller function.
     * @param hashedStatement   Hash of the statement that was signed by an Agent
     * @param signature         Agent signature for the hashed statement
     * @return status   Struct representing agent status:
     *                  - flag      Unknown/Active/Unstaking/Resting/Fraudulent/Slashed
     *                  - domain    Domain where agent is/was active
     *                  - index     Index of agent in the Agent Merkle Tree
     * @return agent    Agent that signed the statement
     */
    function _recoverAgent(bytes32 hashedStatement, bytes memory signature)
        internal
        view
        returns (AgentStatus memory status, address agent)
    {
        bytes32 ethSignedMsg = ECDSA.toEthSignedMessageHash(hashedStatement);
        agent = ECDSA.recover(ethSignedMsg, signature);
        status = IAgentManager(agentManager).agentStatus(agent);
        // Discard signature of unknown agents.
        // Further flag checks are supposed to be performed in a caller function.
        status.verifyKnown();
    }

    /// @dev Verifies that Notary signature is active on local domain.
    function _verifyNotaryDomain(uint32 notaryDomain) internal view {
        // Notary needs to be from the local domain (if contract is not deployed on Synapse Chain).
        // Or Notary could be from any domain (if contract is deployed on Synapse Chain).
        if (notaryDomain != localDomain && localDomain != synapseDomain) revert IncorrectAgentDomain();
    }

    // ════════════════════════════════════════ ATTESTATION RELATED CHECKS ═════════════════════════════════════════════

    /**
     * @dev Internal function to verify the signed attestation payload.
     * Reverts if any of these is true:
     *  - Attestation signer is not a known Notary.
     * @param att               Typed memory view over attestation payload
     * @param attSignature      Notary signature for the attestation
     * @return status           Struct representing agent status, see {_recoverAgent}
     * @return notary           Notary that signed the snapshot
     */
    function _verifyAttestation(Attestation att, bytes memory attSignature)
        internal
        view
        returns (AgentStatus memory status, address notary)
    {
        // This will revert if signer is not a known agent
        (status, notary) = _recoverAgent(att.hashValid(), attSignature);
        // Attestation signer needs to be a Notary, not a Guard
        if (status.domain == 0) revert AgentNotNotary();
    }

    /**
     * @dev Internal function to verify the signed attestation report payload.
     * Reverts if any of these is true:
     *  - Report signer is not a known Guard.
     * @param att               Typed memory view over attestation payload that Guard reports as invalid
     * @param arSignature       Guard signature for the "invalid attestation" report
     * @return status           Struct representing guard status, see {_recoverAgent}
     * @return guard            Guard that signed the report
     */
    function _verifyAttestationReport(Attestation att, bytes memory arSignature)
        internal
        view
        returns (AgentStatus memory status, address guard)
    {
        // This will revert if signer is not a known agent
        (status, guard) = _recoverAgent(att.hashInvalid(), arSignature);
        // Report signer needs to be a Guard, not a Notary
        if (status.domain != 0) revert AgentNotGuard();
    }

    // ══════════════════════════════════════════ RECEIPT RELATED CHECKS ═══════════════════════════════════════════════

    /**
     * @dev Internal function to verify the signed receipt payload.
     * Reverts if any of these is true:
     *  - Receipt signer is not a known Notary.
     * @param rcpt              Typed memory view over receipt payload
     * @param rcptSignature     Notary signature for the receipt
     * @return status           Struct representing agent status, see {_recoverAgent}
     * @return notary           Notary that signed the snapshot
     */
    function _verifyReceipt(Receipt rcpt, bytes memory rcptSignature)
        internal
        view
        returns (AgentStatus memory status, address notary)
    {
        // This will revert if signer is not a known agent
        (status, notary) = _recoverAgent(rcpt.hashValid(), rcptSignature);
        // Receipt signer needs to be a Notary, not a Guard
        if (status.domain == 0) revert AgentNotNotary();
    }

    /**
     * @dev Internal function to verify the signed receipt report payload.
     * Reverts if any of these is true:
     * - Report signer is not a known Guard.
     * @param rcpt              Typed memory view over receipt payload that Guard reports as invalid
     * @param rrSignature       Guard signature for the "invalid receipt" report
     * @return status           Struct representing guard status, see {_recoverAgent}
     * @return guard            Guard that signed the report
     */
    function _verifyReceiptReport(Receipt rcpt, bytes memory rrSignature)
        internal
        view
        returns (AgentStatus memory status, address guard)
    {
        // This will revert if signer is not a known agent
        (status, guard) = _recoverAgent(rcpt.hashInvalid(), rrSignature);
        // Report signer needs to be a Guard, not a Notary
        if (status.domain != 0) revert AgentNotGuard();
    }

    // ═══════════════════════════════════════ STATE/SNAPSHOT RELATED CHECKS ═══════════════════════════════════════════

    /**
     * @dev Internal function to verify the signed snapshot report payload.
     * Reverts if any of these is true:
     *  - Report signer is not a known Guard.
     * @param state             Typed memory view over state payload that Guard reports as invalid
     * @param srSignature       Guard signature for the report
     * @return status           Struct representing guard status, see {_recoverAgent}
     * @return guard            Guard that signed the report
     */
    function _verifyStateReport(State state, bytes memory srSignature)
        internal
        view
        returns (AgentStatus memory status, address guard)
    {
        // This will revert if signer is not a known agent
        (status, guard) = _recoverAgent(state.hashInvalid(), srSignature);
        // Report signer needs to be a Guard, not a Notary
        if (status.domain != 0) revert AgentNotGuard();
    }

    /**
     * @dev Internal function to verify the signed snapshot payload.
     * Reverts if any of these is true:
     *  - Snapshot signer is not a known Agent.
     *  - Snapshot signer is not a Notary (if verifyNotary is true).
     * @param snapshot          Typed memory view over snapshot payload
     * @param snapSignature     Agent signature for the snapshot
     * @param verifyNotary      If true, snapshot signer needs to be a Notary, not a Guard
     * @return status           Struct representing agent status, see {_recoverAgent}
     * @return agent            Agent that signed the snapshot
     */
    function _verifySnapshot(Snapshot snapshot, bytes memory snapSignature, bool verifyNotary)
        internal
        view
        returns (AgentStatus memory status, address agent)
    {
        // This will revert if signer is not a known agent
        (status, agent) = _recoverAgent(snapshot.hashValid(), snapSignature);
        // If requested, snapshot signer needs to be a Notary, not a Guard
        if (verifyNotary && status.domain == 0) revert AgentNotNotary();
    }

    // ═══════════════════════════════════════════ MERKLE RELATED CHECKS ═══════════════════════════════════════════════

    /**
     * @dev Internal function to verify that snapshot roots match.
     * Reverts if any of these is true:
     *  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
     *  - Snapshot Proof's first element does not match the State metadata.
     *  - Snapshot Proof length exceeds Snapshot tree Height.
     *  - State index is out of range.
     * @param att               Typed memory view over Attestation
     * @param stateIndex        Index of state in the snapshot
     * @param state             Typed memory view over the provided state payload
     * @param snapProof         Raw payload with snapshot data
     */
    function _verifySnapshotMerkle(Attestation att, uint8 stateIndex, State state, bytes32[] memory snapProof)
        internal
        pure
    {
        // Snapshot proof first element should match State metadata (aka "right sub-leaf")
        (, bytes32 rightSubLeaf) = state.subLeafs();
        if (snapProof[0] != rightSubLeaf) revert IncorrectSnapshotProof();
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if:
        //  - State index is out of range.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        bytes32 snapshotRoot = SnapshotLib.proofSnapRoot(state.root(), state.origin(), snapProof, stateIndex);
        // Snapshot root should match the attestation root
        if (att.snapRoot() != snapshotRoot) revert IncorrectSnapshotRoot();
    }
}
