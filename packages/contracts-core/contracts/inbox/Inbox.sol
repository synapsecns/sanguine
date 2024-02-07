// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/memory/Attestation.sol";
import {
    CallerNotDestination,
    IncorrectAgentDomain,
    IncorrectSnapshotRoot,
    IncorrectTipsProof,
    MustBeSynapseDomain
} from "../libs/Errors.sol";
import {ChainGas} from "../libs/stack/GasData.sol";
import {MerkleMath} from "../libs/merkle/MerkleMath.sol";
import {Receipt, ReceiptLib} from "../libs/memory/Receipt.sol";
import {Snapshot, SnapshotLib} from "../libs/memory/Snapshot.sol";
import {AgentStatus} from "../libs/Structures.sol";
import {Tips, TipsLib} from "../libs/stack/Tips.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {StatementInbox} from "./StatementInbox.sol";
import {MessagingBase} from "../base/MessagingBase.sol";
import {InboxEvents} from "../events/InboxEvents.sol";
import {IAgentManager} from "../interfaces/IAgentManager.sol";
import {InterfaceDestination} from "../interfaces/InterfaceDestination.sol";
import {IExecutionHub} from "../interfaces/IExecutionHub.sol";
import {InterfaceInbox} from "../interfaces/InterfaceInbox.sol";
import {ISnapshotHub} from "../interfaces/ISnapshotHub.sol";
import {InterfaceSummit} from "../interfaces/InterfaceSummit.sol";

/// @notice `Inbox` is the child of `StatementInbox` contract, that is used on Synapse Chain.
/// In addition to the functionality of `StatementInbox`, it also:
/// - Accepts Guard and Notary Snapshots and passes them to `Summit` contract.
/// - Accepts Notary-signed Receipts and passes them to `Summit` contract.
/// - Accepts Receipt Reports to initiate a dispute between Guard and Notary.
/// - Verifies Attestations and Attestation Reports, and slashes the signer if they are invalid.
contract Inbox is StatementInbox, InboxEvents, InterfaceInbox {
    using AttestationLib for bytes;
    using ReceiptLib for bytes;
    using SnapshotLib for bytes;

    // Struct to get around stack too deep error. TODO: revisit this
    struct ReceiptInfo {
        AgentStatus rcptNotaryStatus;
        address notary;
        uint32 attNonce;
        AgentStatus attNotaryStatus;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // The address of the Summit contract.
    address public summit;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 synapseDomain_) MessagingBase("0.0.3", synapseDomain_) {
        if (localDomain != synapseDomain) revert MustBeSynapseDomain();
    }

    /// @notice Initializes `Inbox` contract:
    /// - Sets `msg.sender` as the owner of the contract
    /// - Sets `agentManager`, `origin`, `destination` and `summit` addresses
    function initialize(address agentManager_, address origin_, address destination_, address summit_)
        external
        initializer
    {
        __StatementInbox_init(agentManager_, origin_, destination_);
        summit = summit_;
    }

    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc InterfaceInbox
    function submitSnapshot(bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bytes memory attPayload, bytes32 agentRoot_, uint256[] memory snapGas)
    {
        // This will revert if payload is not a snapshot
        Snapshot snapshot = snapPayload.castToSnapshot();
        // This will revert if the signer is not a known Guard/Notary
        (AgentStatus memory status, address agent) =
            _verifySnapshot({snapshot: snapshot, snapSignature: snapSignature, verifyNotary: false});
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
            // Get current agentRoot from AgentManager
            agentRoot_ = IAgentManager(agentManager).agentRoot();
            // This will revert if Notary is in Dispute
            attPayload = InterfaceSummit(summit).acceptNotarySnapshot({
                notaryIndex: status.index,
                sigIndex: sigIndex,
                agentRoot: agentRoot_,
                snapPayload: snapPayload
            });
            ChainGas[] memory snapGas_ = snapshot.snapGas();
            // Pass created attestation to Destination to enable executing messages coming to Synapse Chain
            InterfaceDestination(destination).acceptAttestation(
                status.index, type(uint256).max, attPayload, agentRoot_, snapGas_
            );
            // Use assembly to cast ChainGas[] to uint256[] without copying. Highest bits are left zeroed.
            // solhint-disable-next-line no-inline-assembly
            assembly {
                snapGas := snapGas_
            }
        }
        emit SnapshotAccepted(status.domain, agent, snapPayload, snapSignature);
    }

    /// @inheritdoc InterfaceInbox
    function submitReceipt(
        bytes memory rcptPayload,
        bytes memory rcptSignature,
        uint256 paddedTips,
        bytes32 headerHash,
        bytes32 bodyHash
    ) external returns (bool wasAccepted) {
        // Struct to get around stack too deep error.
        ReceiptInfo memory info;
        // This will revert if payload is not a receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the receipt signer is not a known Notary
        (info.rcptNotaryStatus, info.notary) = _verifyReceipt(rcpt, rcptSignature);
        // Receipt Notary needs to be Active
        info.rcptNotaryStatus.verifyActive();
        info.attNonce = IExecutionHub(destination).getAttestationNonce(rcpt.snapshotRoot());
        if (info.attNonce == 0) revert IncorrectSnapshotRoot();
        // Attestation Notary domain needs to match the destination domain
        info.attNotaryStatus = IAgentManager(agentManager).agentStatus(rcpt.attNotary());
        if (info.attNotaryStatus.domain != rcpt.destination()) revert IncorrectAgentDomain();
        // Check that the correct tip values for the message were provided
        _verifyReceiptTips(rcpt.messageHash(), paddedTips, headerHash, bodyHash);
        // Store Notary signature for the Receipt
        uint256 sigIndex = _saveSignature(rcptSignature);
        // This will revert if Receipt Notary is in Dispute
        wasAccepted = InterfaceSummit(summit).acceptReceipt({
            rcptNotaryIndex: info.rcptNotaryStatus.index,
            attNotaryIndex: info.attNotaryStatus.index,
            sigIndex: sigIndex,
            attNonce: info.attNonce,
            paddedTips: paddedTips,
            rcptPayload: rcptPayload
        });
        if (wasAccepted) {
            emit ReceiptAccepted(info.rcptNotaryStatus.domain, info.notary, rcptPayload, rcptSignature);
        }
    }

    /// @inheritdoc InterfaceInbox
    function submitReceiptReport(bytes memory rcptPayload, bytes memory rcptSignature, bytes memory rrSignature)
        external
        returns (bool wasAccepted)
    {
        // This will revert if payload is not a receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the receipt signer is not a known Guard
        (AgentStatus memory guardStatus,) = _verifyReceiptReport(rcpt, rrSignature);
        // Guard needs to be Active
        guardStatus.verifyActive();
        // This will revert if report signer is not a known Notary
        (AgentStatus memory notaryStatus,) = _verifyReceipt(rcpt, rcptSignature);
        // Notary needs to be Active/Unstaking
        notaryStatus.verifyActiveUnstaking();
        _saveReport(rcptPayload, rrSignature);
        // This will revert if either actor is already in dispute
        IAgentManager(agentManager).openDispute(guardStatus.index, notaryStatus.index);
        return true;
    }

    /// @inheritdoc InterfaceInbox
    function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes memory rcptPayload)
        external
        returns (bool wasAccepted)
    {
        // Only Destination can pass receipts
        if (msg.sender != destination) revert CallerNotDestination();
        return InterfaceSummit(summit).acceptReceipt({
            rcptNotaryIndex: attNotaryIndex,
            attNotaryIndex: attNotaryIndex,
            sigIndex: type(uint256).max,
            attNonce: attNonce,
            paddedTips: paddedTips,
            rcptPayload: rcptPayload
        });
    }

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /// @inheritdoc InterfaceInbox
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
            IAgentManager(agentManager).slashAgent(status.domain, notary, msg.sender);
        }
    }

    /// @inheritdoc InterfaceInbox
    function verifyAttestationReport(bytes memory attPayload, bytes memory arSignature)
        external
        returns (bool isValidReport)
    {
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // This will revert if the report signer is not a known Guard
        (AgentStatus memory status, address guard) = _verifyAttestationReport(att, arSignature);
        // Guard needs to be Active/Unstaking
        status.verifyActiveUnstaking();
        // Report is valid IF AND ONLY IF the reported attestation in invalid
        isValidReport = !ISnapshotHub(summit).isValidAttestation(attPayload);
        if (!isValidReport) {
            emit InvalidAttestationReport(attPayload, arSignature);
            IAgentManager(agentManager).slashAgent(status.domain, guard, msg.sender);
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Verifies that tips proof matches the message hash.
    function _verifyReceiptTips(bytes32 msgHash, uint256 paddedTips, bytes32 headerHash, bytes32 bodyHash)
        internal
        pure
    {
        Tips tips = TipsLib.wrapPadded(paddedTips);
        // full message leaf is (header, baseMessage), while base message leaf is (tips, remainingBody).
        if (MerkleMath.getParent(headerHash, MerkleMath.getParent(tips.leaf(), bodyHash)) != msgHash) {
            revert IncorrectTipsProof();
        }
    }
}
