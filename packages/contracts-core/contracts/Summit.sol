// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {AttestationLib} from "./libs/memory/Attestation.sol";
import {ByteString} from "./libs/memory/ByteString.sol";
import {BONDING_OPTIMISTIC_PERIOD, TIPS_GRANULARITY} from "./libs/Constants.sol";
import {
    MustBeSynapseDomain,
    DisputeTimeoutNotOver,
    NotaryInDispute,
    TipsClaimMoreThanEarned,
    TipsClaimZero
} from "./libs/Errors.sol";
import {Receipt, ReceiptLib} from "./libs/memory/Receipt.sol";
import {Snapshot, SnapshotLib} from "./libs/memory/Snapshot.sol";
import {AgentFlag, AgentStatus, DisputeFlag, MessageStatus} from "./libs/Structures.sol";
import {Tips, TipsLib} from "./libs/stack/Tips.sol";
import {ChainContext} from "./libs/ChainContext.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "./base/AgentSecured.sol";
import {SummitEvents} from "./events/SummitEvents.sol";
import {IAgentManager} from "./interfaces/IAgentManager.sol";
import {InterfaceBondingManager} from "./interfaces/InterfaceBondingManager.sol";
import {InterfaceSummit} from "./interfaces/InterfaceSummit.sol";
import {SnapshotHub} from "./hubs/SnapshotHub.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {DoubleEndedQueue} from "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

/// @notice `Summit` contract is the cornerstone of the Synapse messaging protocol. This is where the
/// states of all the remote chains (provided collectively by the Guards and Notaries) are stored. This is
/// also the place where the tips are distributed among the off-chain actors.
/// `Summit` is responsible for the following:
/// - Accepting Guard and Notary snapshots from the local `Inbox` contract, and storing the states from these
///   snapshots (see parent contract `SnapshotHub`).
/// - Accepting Notary Receipts from the local `Inbox` contract, and using them to distribute tips among the
///   off-chain actors that participated in the message lifecycle.
contract Summit is SnapshotHub, SummitEvents, InterfaceSummit {
    using AttestationLib for bytes;
    using ByteString for bytes;
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;
    using ReceiptLib for bytes;
    using SnapshotLib for bytes;

    // TODO: write docs, pack values
    struct SummitReceipt {
        uint32 origin;
        uint32 destination;
        uint32 attNonce;
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

    /// @notice Struct for storing the actor tips for a given origin domain.
    /// @param earned   Total amount of tips earned by the actor, denominated in domain's wei
    /// @param claimed  Total amount of tips claimed by the actor, denominated in domain's wei
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

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 synapseDomain_, address agentManager_, address inbox_)
        AgentSecured("0.0.3", synapseDomain_, agentManager_, inbox_)
    {
        if (localDomain != synapseDomain) revert MustBeSynapseDomain();
    }

    /// @notice Initializes Summit contract:
    /// - `owner_` is set as contract owner
    /// - Empty snapshot and attestation are saved, so that the attestation nonce starts from 1
    function initialize(address owner_) external initializer {
        __MessagingBase_init(owner_);
        _initializeAttestations();
    }

    // ═════════════════════════════════════════════ ACCEPT STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceSummit
    function acceptReceipt(
        uint32 rcptNotaryIndex,
        uint32 attNotaryIndex,
        uint256 sigIndex,
        uint32 attNonce,
        uint256 paddedTips,
        bytes memory rcptPayload
    ) external onlyInbox returns (bool wasAccepted) {
        // Check that we can trust the receipt Notary data: they are not in dispute, and the dispute timeout is over.
        if (_notaryDisputeExists(rcptNotaryIndex)) revert NotaryInDispute();
        if (_notaryDisputeTimeout(rcptNotaryIndex)) revert DisputeTimeoutNotOver();
        // This will revert if payload is not a receipt body
        return _saveReceipt({
            rcpt: rcptPayload.castToReceipt(),
            tips: TipsLib.wrapPadded(paddedTips),
            rcptNotaryIndex: rcptNotaryIndex,
            attNotaryIndex: attNotaryIndex,
            sigIndex: sigIndex,
            attNonce: attNonce
        });
    }

    /// @inheritdoc InterfaceSummit
    function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes memory snapPayload) external onlyInbox {
        // Note: we don't check if Guard is in Dispute,
        // as the Guards could continue to submit snapshots after submitting a report.
        // This will revert if payload is not a snapshot
        _acceptGuardSnapshot(snapPayload.castToSnapshot(), guardIndex, sigIndex);
    }

    /// @inheritdoc InterfaceSummit
    function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes memory snapPayload)
        external
        onlyInbox
        returns (bytes memory attPayload)
    {
        // Check that we can trust the snapshot Notary data: they are not in dispute, and the dispute timeout is over.
        if (_notaryDisputeExists(notaryIndex)) revert NotaryInDispute();
        if (_notaryDisputeTimeout(notaryIndex)) revert DisputeTimeoutNotOver();
        // This will revert if payload is not a snapshot
        return _acceptNotarySnapshot(snapPayload.castToSnapshot(), agentRoot, notaryIndex, sigIndex);
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
        if (amount == 0) revert TipsClaimZero();
        ActorTips memory tips = actorTips[msg.sender][origin];
        if (tips.earned < amount + tips.claimed) revert TipsClaimMoreThanEarned();
        // Guaranteed to fit into uint128, as the sum is lower than `earned`
        actorTips[msg.sender][origin].claimed = uint128(tips.claimed + amount);
        InterfaceBondingManager(address(agentManager)).withdrawTips(msg.sender, origin, amount);
        emit TipWithdrawalInitiated(msg.sender, origin, amount);
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
            SummitState memory state = _latestState(origin, _agentStatus(guards[i]).index);
            if (state.nonce > latestState.nonce) latestState = state;
        }
        // Check if we found anything
        if (latestState.nonce != 0) {
            statePayload = _formatSummitState(latestState);
        }
    }

    // ═══════════════════════════════════════════ INTERNAL LOGIC: QUEUE ═══════════════════════════════════════════════

    /// @dev Checks if the given Notary has been disputed.
    /// - Notary was slashed => receipt is invalided and deleted
    /// - Notary is in Dispute => receipt handling is postponed
    function _checkNotaryDisputed(bytes32 messageHash, uint32 notaryIndex) internal returns (bool queuePopped) {
        // TODO: add timeout for Notaries that just won the dispute.
        DisputeFlag flag = _disputes[notaryIndex].flag;
        if (flag == DisputeFlag.Slashed) {
            // Notary has been slashed, so we can't trust their statement.
            // Honest Notaries are incentivized to resubmit the Receipt or Attestation if it was in fact valid.
            _deleteFromQueue(messageHash);
            queuePopped = true;
        } else if (flag == DisputeFlag.Pending || _notaryDisputeTimeout(notaryIndex)) {
            // Notary is in the ongoing Dispute, or has recently won one. We postpone the receipt handling.
            // To keep the tips flow going we add the receipt to the back of the queue,
            // hoping that by the next interaction the dispute will have been resolved.
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
    function _saveReceipt(
        Receipt rcpt,
        Tips tips,
        uint32 rcptNotaryIndex,
        uint32 attNotaryIndex,
        uint256 sigIndex,
        uint32 attNonce
    ) internal returns (bool) {
        // TODO: save signature index
        // Check if tip values are non-zero
        if (tips.value() == 0) return false;
        // Check if there already exists receipt for the message
        bytes32 messageHash = rcpt.messageHash();
        ReceiptStatus memory savedRcpt = _receiptStatus[messageHash];
        // Don't save if receipt is already in the queue
        if (savedRcpt.pending) return false;
        // Get the status from the provided receipt
        MessageStatus msgStatus = rcpt.finalExecutor() == address(0) ? MessageStatus.Failed : MessageStatus.Success;
        // Don't save if we already have the receipt with at least this status
        if (savedRcpt.status >= msgStatus) return false;
        // Save information from the receipt
        _receipts[messageHash] = SummitReceipt({
            origin: rcpt.origin(),
            destination: rcpt.destination(),
            attNonce: attNonce,
            stateIndex: rcpt.stateIndex(),
            attNotaryIndex: attNotaryIndex,
            firstExecutor: rcpt.firstExecutor(),
            finalExecutor: rcpt.finalExecutor()
        });
        // Save receipt status: transfer tipsAwarded field (whether we paid tips for Failed Receipt before)
        _receiptStatus[messageHash] = ReceiptStatus({
            status: msgStatus,
            pending: true,
            tipsAwarded: savedRcpt.tipsAwarded,
            receiptNotaryIndex: rcptNotaryIndex,
            submittedAt: ChainContext.blockTimestamp()
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
            _awardSnapshotTip(summitRcpt.attNonce, summitRcpt.stateIndex, summitRcpt.origin, tips.summitTip);
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
        // We need to do a shit here, as we operate with "scaled down" tips everywhere,
        // but Summit is supposed to store the "full tip value".
        // Tip fits into 64 bits, so it's safe to do a 32 bit shift without risk of overflow
        uint128 tipAwarded = uint128(tip) << uint128(TIPS_GRANULARITY);
        actorTips[actor][origin].earned += tipAwarded;
        emit TipAwarded(actor, origin, tipAwarded);
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
    function _awardSnapshotTip(uint32 attNonce, uint8 stateIndex, uint32 origin, uint64 summitTip) internal {
        uint64 snapshotTip = _snapshotTip(summitTip);
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
