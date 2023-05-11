// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/Attestation.sol";
import {
    CallerNotDestination, IncorrectAgentDomain, IncorrectSnapshotRoot, MustBeSynapseDomain
} from "../libs/Errors.sol";
import {SYNAPSE_DOMAIN} from "../libs/Constants.sol";
import {ChainGas} from "../libs/GasData.sol";
import {Receipt, ReceiptBody, ReceiptLib} from "../libs/Receipt.sol";
import {Snapshot, SnapshotLib} from "../libs/Snapshot.sol";
import {AgentStatus} from "../libs/Structures.sol";
import {Tips} from "../libs/Tips.sol";
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

contract Inbox is StatementInbox, InboxEvents, InterfaceInbox {
    using AttestationLib for bytes;
    using ReceiptLib for bytes;
    using SnapshotLib for bytes;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // The address of the Summit contract.
    address public summit;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 domain) MessagingBase("0.0.3", domain) {
        if (domain != SYNAPSE_DOMAIN) revert MustBeSynapseDomain();
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
    function submitReceipt(bytes memory rcptPayload, bytes memory rcptSignature) external returns (bool wasAccepted) {
        // This will revert if payload is not a receipt
        Receipt rcpt = rcptPayload.castToReceipt();
        // This will revert if the receipt signer is not a known Notary
        (AgentStatus memory rcptNotaryStatus, address notary) = _verifyReceipt(rcpt, rcptSignature);
        // Receipt Notary needs to be Active
        rcptNotaryStatus.verifyActive();
        // Check that receipt's snapshot root exists in Summit
        ReceiptBody rcptBody = rcpt.body();
        uint32 attNonce = IExecutionHub(destination).getAttestationNonce(rcptBody.snapshotRoot());
        if (attNonce == 0) revert IncorrectSnapshotRoot();
        // Attestation Notary domain needs to match the destination domain
        AgentStatus memory attNotaryStatus = IAgentManager(agentManager).agentStatus(rcptBody.attNotary());
        if (attNotaryStatus.domain != rcptBody.destination()) revert IncorrectAgentDomain();
        // Store Notary signature for the Receipt
        uint256 sigIndex = _saveSignature(rcptSignature);
        // This will revert if Receipt Notary is in Dispute
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

    /// @inheritdoc InterfaceInbox
    function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes memory rcptBodyPayload)
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
            rcptBodyPayload: rcptBodyPayload
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
}
