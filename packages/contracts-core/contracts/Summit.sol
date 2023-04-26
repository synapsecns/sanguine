// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {AttestationLib} from "./libs/Attestation.sol";
import {ByteString} from "./libs/ByteString.sol";
import {BONDING_OPTIMISTIC_PERIOD, SYNAPSE_DOMAIN} from "./libs/Constants.sol";
import {Receipt, ReceiptLib} from "./libs/Receipt.sol";
import {Snapshot, SnapshotLib} from "./libs/Snapshot.sol";
import {AgentFlag, AgentStatus, DisputeFlag} from "./libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "./base/AgentSecured.sol";
import {SummitEvents} from "./events/SummitEvents.sol";
import {IAgentManager} from "./interfaces/IAgentManager.sol";
import {InterfaceBondingManager} from "./interfaces/InterfaceBondingManager.sol";
import {InterfaceSummit} from "./interfaces/InterfaceSummit.sol";
import {ExecutionHub, MessageStatus, ReceiptBody, Tips} from "./hubs/ExecutionHub.sol";
import {SnapshotHub} from "./hubs/SnapshotHub.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {DoubleEndedQueue} from "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

contract Summit is ExecutionHub, SnapshotHub, SummitEvents, InterfaceSummit {
    using AttestationLib for bytes;
    using ByteString for bytes;
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using ReceiptLib for bytes;
    using SnapshotLib for bytes;

    struct StoredSnapData {
        bytes32 r;
        bytes32 s;
    }

    // TODO: write docs, pack values
    struct SummitReceipt {
        uint32 origin;
        uint32 destination;
        uint32 snapRootIndex;
        uint8 stateIndex;
        uint32 attNotaryIndex;
        address firstExecutor;
        address finalExecutor;
    }

    struct ReceiptStatus {
        MessageStatus status;
        bool pending;
        bool tipsAwarded;
        uint32 receiptNotaryIndex;
        uint40 submittedAt;
    }

    struct ReceiptTips {
        uint64 summitTip;
        uint64 attestationTip;
        uint64 executionTip;
        uint64 deliveryTip;
    }

    struct ActorTips {
        uint128 earned;
        uint128 claimed;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // (message hash => receipt data)
    mapping(bytes32 => SummitReceipt) private _receipts;

    // (message hash => receipt status)
    mapping(bytes32 => ReceiptStatus) private _receiptStatus;

    // (message hash => receipt tips)
    mapping(bytes32 => ReceiptTips) private _receiptTips;

    // Quarantine queue for message hashes
    DoubleEndedQueue.Bytes32Deque private _receiptQueue;

    /// @inheritdoc InterfaceSummit
    mapping(address => mapping(uint32 => ActorTips)) public actorTips;

    /// @dev Stored lookup data for all accepted Notary Snapshots
    StoredSnapData[] internal _storedSnapshots;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain, address agentManager_) AgentSecured("0.0.3", domain, agentManager_) {
        require(domain == SYNAPSE_DOMAIN, "Only deployed on SynChain");
    }

    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        _initializeAttestations();
    }

    // ═════════════════════════════════════════════ ACCEPT STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceSummit
    function acceptReceipt(
        address notary,
        AgentStatus memory status,
        bytes memory rcptPayload,
        bytes memory rcptSignature
    ) external returns (bool wasAccepted) {
        // This will revert if payload is not an receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // Receipt needs to be signed by a destination chain Notary
        ReceiptBody rcptBody = rcpt.body();
        // TODO: remove this restriction
        require(rcptBody.destination() == status.domain, "Wrong Notary domain");
        wasAccepted = _saveReceipt(rcptBody, rcpt.tips(), status.index);
        if (wasAccepted) {
            // TODO: save signature
            emit ReceiptAccepted(status.domain, notary, rcptPayload, rcptSignature);
        }
    }

    /// @inheritdoc InterfaceSummit
    function acceptSnapshot(
        address agent,
        AgentStatus memory status,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bytes memory attPayload) {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        if (status.domain == 0) {
            /// @dev We don't check if Guard is in dispute for accepting the snapshots.
            /// Guard could only be in Dispute, if they submitted a Report on a Notary.
            /// This should not strip away their ability to post snapshots, as they require
            /// a Notary signature in order to be used / gain tips anyway.

            // This will revert if Guard has previously submitted
            // a fresher state than one in the snapshot.
            _acceptGuardSnapshot(snapshot, agent, status.index);
        } else {
            // Fetch current Agent Root from BondingManager
            bytes32 agentRoot = IAgentManager(agentManager).agentRoot();
            // This will revert if any of the states from the Notary snapshot
            // haven't been submitted by any of the Guards before.
            attPayload = _acceptNotarySnapshot(snapshot, agentRoot, agent, status.index);
            // Save attestation derived from Notary snapshot.
            (bytes32 r, bytes32 s, uint8 v) = snapSignature.castToSignature().toRSV();
            _saveAttestation(attPayload.castToAttestation(), status.index, v);
            _storedSnapshots.push(StoredSnapData({r: r, s: s}));
        }
        emit SnapshotAccepted(status.domain, agent, snapPayload, snapSignature);
    }

    // ════════════════════════════════════════════════ TIPS LOGIC ═════════════════════════════════════════════════════

    /// @inheritdoc InterfaceSummit
    function distributeTips() public returns (bool queuePopped) {
        // Check message that is first in the "quarantine queue"
        if (_receiptQueue.empty()) return false;
        bytes32 messageHash = _receiptQueue.front();
        ReceiptStatus memory rcptStatus = _receiptStatus[messageHash];
        // Check if optimistic period for the receipt is over
        if (block.timestamp < uint256(rcptStatus.submittedAt) + BONDING_OPTIMISTIC_PERIOD) return false;
        // Fetch Notary who signed the receipt. If they are Slashed or in Dispute, exit early.
        if (_checkNotaryDisputed(messageHash, rcptStatus.receiptNotaryIndex)) return true;
        SummitReceipt memory summitRcpt = _receipts[messageHash];
        // Fetch Notary who signed the statement with snapshot root. If they are Slashed or in Dispute, exit early.
        if (_checkNotaryDisputed(messageHash, summitRcpt.attNotaryIndex)) return true;
        // At this point Receipt is optimistically verified to be correct, as well as the receipt's attestation
        // Meaning we can go ahead and distribute the tip values among the tipped actors.
        _awardTips(rcptStatus.receiptNotaryIndex, summitRcpt.attNotaryIndex, messageHash, summitRcpt, rcptStatus);
        // Save new receipt status
        rcptStatus.pending = false;
        rcptStatus.tipsAwarded = true;
        _receiptStatus[messageHash] = rcptStatus;
        // Remove the receipt from the queue
        _receiptQueue.popFront();
        return true;
    }

    /// @inheritdoc InterfaceSummit
    // solhint-disable-next-line ordering
    function withdrawTips(uint32 origin, uint256 amount) external {
        require(amount != 0, "Amount is zero");
        ActorTips memory tips = actorTips[msg.sender][origin];
        require(tips.earned >= amount + tips.claimed, "Tips balance too low");
        // Guaranteed to fit into uint128, as the sum is lower than `earned`
        actorTips[msg.sender][origin].claimed = uint128(tips.claimed + amount);
        InterfaceBondingManager(address(agentManager)).withdrawTips(msg.sender, origin, amount);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceSummit
    // solhint-disable-next-line ordering
    function receiptQueueLength() external view returns (uint256) {
        return _receiptQueue.length();
    }

    /// @inheritdoc InterfaceSummit
    function getLatestState(uint32 origin) external view returns (bytes memory statePayload) {
        // Get a list of currently active guards
        address[] memory guards = InterfaceBondingManager(address(agentManager)).getActiveAgents(0);
        SummitState memory latestState;
        for (uint256 i = 0; i < guards.length; ++i) {
            SummitState memory state = _latestState(origin, guards[i]);
            if (state.nonce > latestState.nonce) latestState = state;
        }
        // Check if we found anything
        if (latestState.nonce != 0) {
            statePayload = _formatSummitState(latestState);
        }
    }

    /// @inheritdoc InterfaceSummit
    function getSignedSnapshot(uint256 nonce)
        external
        view
        returns (bytes memory snapPayload, bytes memory snapSignature)
    {
        // This will revert if nonce is out of range
        snapPayload = getNotarySnapshot(nonce);
        StoredSnapData memory storedSnap = _storedSnapshots[nonce - 1];
        SnapRootData memory rootData = _rootData[_roots[nonce - 1]];
        snapSignature = ByteString.formatSignature({r: storedSnap.r, s: storedSnap.s, v: rootData.notaryV});
    }

    // ═══════════════════════════════════════════ INTERNAL LOGIC: QUEUE ═══════════════════════════════════════════════

    /// @dev Checks if the given Notary has been disputed.
    /// - Notary was slashed => receipt is invalided and deleted
    /// - Notary is in Dispute => receipt handling is postponed
    function _checkNotaryDisputed(bytes32 messageHash, uint32 notaryIndex) internal returns (bool queuePopped) {
        DisputeFlag flag = _disputes[notaryIndex];
        if (flag == DisputeFlag.Slashed) {
            // Notary has been slashed, so we can't trust their statement.
            // Honest Notaries are incentivized to resubmit the Receipt or Attestation if it was in fact valid.
            _deleteFromQueue(messageHash);
            queuePopped = true;
        } else if (flag == DisputeFlag.Pending) {
            // Notary is not slashed, but is in Dispute. To keep the tips flow going we add the receipt to the back of
            // the queue, hoping that by the next interaction the dispute will have been resolved.
            _moveToBack();
            queuePopped = true;
        }
    }

    /// @dev Deletes all stored receipt data and removes it from the queue.
    function _deleteFromQueue(bytes32 messageHash) internal {
        delete _receipts[messageHash];
        delete _receiptStatus[messageHash];
        delete _receiptTips[messageHash];
        _receiptQueue.popFront();
    }

    /// @dev Moves the front element of the queue to its back.
    function _moveToBack() internal {
        bytes32 popped = _receiptQueue.popFront();
        _receiptQueue.pushBack(popped);
    }

    /// @dev Saves the message from the receipt into the "quarantine queue". Once message leaves the queue,
    /// tips associated with the message are distributed across off-chain actors.
    function _saveReceipt(ReceiptBody rcptBody, Tips tips, uint32 rcptNotaryIndex) internal returns (bool) {
        bytes32 snapRoot = rcptBody.snapshotRoot();
        SnapRootData memory rootData = _rootData[snapRoot];
        require(rootData.submittedAt != 0, "Unknown snapshot root");
        // Attestation Notary needs to be known and not slashed
        address attNotary = rcptBody.attNotary();
        AgentStatus memory attNotaryStatus = _agentStatus(attNotary);
        attNotaryStatus.verifyKnown();
        attNotaryStatus.verifyNotSlashed();
        // Check if tip values are non-zero
        if (tips.value() == 0) return false;
        // Check if there already exists receipt for the message
        bytes32 messageHash = rcptBody.messageHash();
        ReceiptStatus memory savedRcpt = _receiptStatus[messageHash];
        // Don't save if receipt is already in the queue
        if (savedRcpt.pending) return false;
        // Get the status from the provided receipt
        MessageStatus msgStatus = rcptBody.finalExecutor() == address(0) ? MessageStatus.Failed : MessageStatus.Success;
        // Don't save if we already have the receipt with at least this status
        if (savedRcpt.status >= msgStatus) return false;
        // Save information from the receipt
        _receipts[messageHash] = SummitReceipt({
            origin: rcptBody.origin(),
            destination: rcptBody.destination(),
            snapRootIndex: rootData.index,
            stateIndex: rcptBody.stateIndex(),
            attNotaryIndex: attNotaryStatus.index,
            firstExecutor: rcptBody.firstExecutor(),
            finalExecutor: rcptBody.finalExecutor()
        });
        // Save receipt status: transfer tipsAwarded field (whether we paid tips for Failed Receipt before)
        _receiptStatus[messageHash] = ReceiptStatus({
            status: msgStatus,
            pending: true,
            tipsAwarded: savedRcpt.tipsAwarded,
            receiptNotaryIndex: rcptNotaryIndex,
            submittedAt: uint40(block.timestamp)
        });
        // Save receipt tips
        _receiptTips[messageHash] = ReceiptTips({
            summitTip: tips.summitTip(),
            attestationTip: tips.attestationTip(),
            executionTip: tips.executionTip(),
            deliveryTip: tips.deliveryTip()
        });
        // Add message hash to the quarantine queue
        _receiptQueue.pushBack(messageHash);
        return true;
    }

    // ══════════════════════════════════════ INTERNAL LOGIC: TIPS ACCOUNTING ══════════════════════════════════════════

    /// @dev Awards tips to the agent/actors that participated in message lifecycle
    function _awardTips(
        uint32 rcptNotaryIndex,
        uint32 attNotaryIndex,
        bytes32 messageHash,
        SummitReceipt memory summitRcpt,
        ReceiptStatus memory rcptStatus
    ) internal {
        ReceiptTips memory tips = _receiptTips[messageHash];
        // Check if we awarded tips for this message earlier
        bool awardFirst = !rcptStatus.tipsAwarded;
        // Check if this is the final tips distribution
        bool awardFinal = rcptStatus.status == MessageStatus.Success;
        if (awardFirst) {
            // There has been a valid attempt to execute the message
            _awardSnapshotTip(
                _roots[summitRcpt.snapRootIndex], summitRcpt.stateIndex, summitRcpt.origin, tips.summitTip
            );
            _awardAgentTip(attNotaryIndex, summitRcpt.origin, tips.attestationTip);
            _awardActorTip(summitRcpt.firstExecutor, summitRcpt.origin, tips.executionTip);
        }
        _awardReceiptTip(rcptNotaryIndex, awardFirst, awardFinal, summitRcpt.origin, tips.summitTip);
        if (awardFinal) {
            // Message has been executed successfully
            _awardActorTip(summitRcpt.finalExecutor, summitRcpt.origin, tips.deliveryTip);
        }
    }

    /// @dev Award tip to the bonded agent
    function _awardAgentTip(uint32 agentIndex, uint32 origin, uint64 tip) internal {
        (address agent, AgentStatus memory status) = _getAgent(agentIndex);
        // If agent has been slashed, their earned tips go to treasury
        if (status.flag == AgentFlag.Fraudulent || status.flag == AgentFlag.Slashed) {
            agent = address(0);
        }
        _awardActorTip(agent, origin, tip);
    }

    /// @dev Award tip to any actor whether bonded or unbonded
    function _awardActorTip(address actor, uint32 origin, uint64 tip) internal {
        actorTips[actor][origin].earned += tip;
        emit TipAwarded(actor, origin, tip);
    }

    /// @dev Award tip for posting Receipt to Summit contract.
    function _awardReceiptTip(uint32 rcptNotaryIndex, bool awardFirst, bool awardFinal, uint32 origin, uint64 summitTip)
        internal
    {
        uint64 receiptTip = _receiptTip(summitTip);
        uint64 receiptTipAwarded;
        if (awardFirst && awardFinal) {
            receiptTipAwarded = receiptTip;
        } else if (awardFirst) {
            // Tip for posting Receipt with status >= MessageStatus.Failed
            receiptTipAwarded = receiptTip / 2;
        } else if (awardFinal) {
            // Tip for posting Receipt with status == MessageStatus.Success
            receiptTipAwarded = receiptTip - receiptTip / 2;
        }
        _awardAgentTip(rcptNotaryIndex, origin, receiptTipAwarded);
    }

    /// @dev Award tip for posting Snapshot to Summit contract.
    function _awardSnapshotTip(bytes32 snapRoot, uint8 stateIndex, uint32 origin, uint64 summitTip) internal {
        uint64 snapshotTip = _snapshotTip(summitTip);
        // Get the attestation nonce for the snapshot root
        uint32 attNonce = _rootData[snapRoot].attNonce;
        // Get the agents who submitted the given state for the attestation's snapshot
        (uint32 guardIndex, uint32 notaryIndex) = _stateAgents(attNonce, stateIndex);
        _awardAgentTip(guardIndex, origin, snapshotTip);
        _awardAgentTip(notaryIndex, origin, snapshotTip);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns "snapshot part" of the summit tip.
    function _snapshotTip(uint64 summitTip) internal pure returns (uint64) {
        return summitTip / 3;
    }

    /// @dev Returns "receipt part" of the summit tip.
    function _receiptTip(uint64 summitTip) internal pure returns (uint64) {
        return summitTip - 2 * _snapshotTip(summitTip);
    }
}
