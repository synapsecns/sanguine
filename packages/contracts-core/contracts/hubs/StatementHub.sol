// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation, AttestationLib} from "../libs/Attestation.sol";
import {Snapshot, SnapshotLib, SNAPSHOT_TREE_HEIGHT, State, StateLib} from "../libs/Snapshot.sol";
import {AttestationReport, AttestationReportLib} from "../libs/AttestationReport.sol";
import {Receipt, ReceiptLib} from "../libs/Receipt.sol";
import {MerkleMath} from "../libs/MerkleMath.sol";
import {StateReport, StateReportLib} from "../libs/StateReport.sol";
import {AgentFlag, AgentStatus} from "../libs/Structures.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {SystemRegistry} from "../system/SystemRegistry.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/**
 * @notice This abstract contract is used for verifying Guards and Notaries
 * signature over several type of statements such as:
 * - Attestations
 * - Snapshots
 * - Reports
 * Several checks are performed in StatementHub, abstracting it away from the child contracts:
 * - Statement being properly formatted
 * - Signer being an active agent
 * - Signer being allowed to sign the particular type of statement
 */
abstract contract StatementHub is SystemRegistry {
    using AttestationLib for bytes;
    using AttestationReportLib for bytes;
    using ReceiptLib for bytes;
    using SnapshotLib for bytes;
    using StateLib for bytes;
    using StateReportLib for bytes;

    /**
     * @dev Recovers a signer from a hashed message, and a EIP-191 signature for it.
     * Will revert, if the signer is not a known agent.
     * @dev Agent flag could ne eny of these: Active/unstaking/Resting/Slashed
     * Further checks need to be performed in a caller function.
     * @param hashedStatement   Hash of the statement that was signed by an Agent
     * @param signature         Agent signature for the hashed statement
     * @return status   Struct representing agent status:
     *                  - flag      Unknown/Active/unstaking/Resting/Slashed
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
        // TODO: ensure that Unstaking agents could be slashed,
        // but their signature is considered invalid for new statements
        status = _agentStatus(agent);
        // Discard signature of unknown agents.
        // Further flag checks are supposed to be performed in a caller function.
        require(status.flag != AgentFlag.Unknown, "Unknown agent");
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
        (status, notary) = _recoverAgent(att.hash(), attSignature);
        // Attestation signer needs to be a Notary, not a Guard
        require(status.domain != 0, "Signer is not a Notary");
    }

    /**
     * @dev Internal function to verify the signed attestation report payload.
     * Reverts if any of these is true:
     *  - Report signer is not a known Guard.
     * @param report            Typed memory view over report payload
     * @param arSignature       Guard signature for the report
     * @return status           Struct representing guard status, see {_recoverAgent}
     * @return guard            Guard that signed the report
     */
    function _verifyAttestationReport(AttestationReport report, bytes memory arSignature)
        internal
        view
        returns (AgentStatus memory status, address guard)
    {
        // This will revert if signer is not a known agent
        (status, guard) = _recoverAgent(report.hash(), arSignature);
        // Report signer needs to be a Guard, not a Notary
        require(status.domain == 0, "Signer is not a Guard");
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
        (status, notary) = _recoverAgent(rcpt.hash(), rcptSignature);
        // Receipt signer needs to be a Notary, not a Guard
        require(status.domain != 0, "Signer is not a Notary");
    }

    // ═════════════════════════════════════ STATE OR SNAPSHOT RELATED CHECKS ══════════════════════════════════════════

    /**
     * @dev Internal function to verify the signed snapshot report payload.
     * Reverts if any of these is true:
     *  - Report signer is not a known Guard.
     * @param report            Typed memory view over report payload
     * @param srSignature       Guard signature for the report
     * @return status           Struct representing guard status, see {_recoverAgent}
     * @return guard            Guard that signed the report
     */
    function _verifyStateReport(StateReport report, bytes memory srSignature)
        internal
        view
        returns (AgentStatus memory status, address guard)
    {
        // This will revert if signer is not a known agent
        (status, guard) = _recoverAgent(report.hash(), srSignature);
        // Report signer needs to be a Guard, not a Notary
        require(status.domain == 0, "Signer is not a Guard");
    }

    /**
     * @dev Internal function to verify the signed snapshot payload.
     * Reverts if any of these is true:
     *  - Snapshot signer is not a known Agent.
     * @param snapshot          Typed memory view over snapshot payload
     * @param snapSignature     Agent signature for the snapshot
     * @return status           Struct representing agent status, see {_recoverAgent}
     * @return agent            Agent that signed the snapshot
     */
    function _verifySnapshot(Snapshot snapshot, bytes memory snapSignature)
        internal
        view
        returns (AgentStatus memory status, address agent)
    {
        // This will revert if signer is not a known agent
        (status, agent) = _recoverAgent(snapshot.hash(), snapSignature);
        // Guards and Notaries for all domains could sign Snapshots, no further checks are needed.
    }

    /**
     * @dev Internal function to verify that snapshot and attestation has the same Merkle Data.
     * Reverts if any of these is true:
     *  - Attestation root is not equal to root derived from the snapshot.
     * @param att           Typed memory view over Attestation
     * @param snapshot      Typed memory view over snapshot payload
     */
    function _verifySnapshotMerkle(Attestation att, Snapshot snapshot) internal pure {
        require(att.snapRoot() == snapshot.root(), "Incorrect snapshot root");
    }

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
    function _verifySnapshotMerkle(Attestation att, uint256 stateIndex, State state, bytes32[] memory snapProof)
        internal
        pure
    {
        // Snapshot proof first element should match State metadata (aka "right sub-leaf")
        (, bytes32 rightSubLeaf) = state.subLeafs();
        require(snapProof[0] == rightSubLeaf, "Incorrect proof[0]");
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if:
        //  - State index is out of range.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        bytes32 snapshotRoot = _snapshotRoot(state.root(), state.origin(), snapProof, stateIndex);
        // Snapshot root should match the attestation root
        require(att.snapRoot() == snapshotRoot, "Incorrect snapshot root");
    }

    /**
     * @dev Reconstructs Snapshot merkle Root from State Merkle Data (root + origin domain)
     * and proof of inclusion of State Merkle Data (aka State "left sub-leaf") in Snapshot Merkle Tree.
     * Reverts if any of these is true:
     *  - State index is out of range.
     *  - Snapshot Proof length exceeds Snapshot tree Height.
     * @param originRoot    Root of Origin Merkle Tree
     * @param origin        Domain of Origin chain
     * @param snapProof     Proof of inclusion of State Merkle Data into Snapshot Merkle Tree
     * @param stateIndex    Index of Origin State in the Snapshot
     */
    function _snapshotRoot(bytes32 originRoot, uint32 origin, bytes32[] memory snapProof, uint256 stateIndex)
        internal
        pure
        returns (bytes32 snapshotRoot)
    {
        // Index of "leftLeaf" is twice the state position in the snapshot
        uint256 leftLeafIndex = stateIndex << 1;
        // Check that "leftLeaf" index fits into Snapshot Merkle Tree
        require(leftLeafIndex < (1 << SNAPSHOT_TREE_HEIGHT), "State index out of range");
        // Reconstruct left sub-leaf of the Origin State: (originRoot, originDomain)
        bytes32 leftLeaf = StateLib.leftLeaf(originRoot, origin);
        // Reconstruct snapshot root using proof of inclusion
        // This will revert if snapshot proof length exceeds Snapshot Tree Height
        return MerkleMath.proofRoot(leftLeafIndex, leftLeaf, snapProof, SNAPSHOT_TREE_HEIGHT);
    }

    // ════════════════════════════════════════════════ FLAG CHECKS ════════════════════════════════════════════════════

    /// @dev Checks that Agent is Active
    function _verifyActive(AgentStatus memory status) internal pure {
        require(status.flag == AgentFlag.Active, status.domain == 0 ? "Not an active guard" : "Not an active notary");
    }

    /// @dev Checks that Agent is Active or Unstaking
    function _verifyActiveUnstaking(AgentStatus memory status) internal pure {
        require(
            (status.flag == AgentFlag.Active || status.flag == AgentFlag.Unstaking),
            status.domain == 0 ? "Not an active guard" : "Not an active notary"
        );
    }

    /// @dev Checks that Agent is not Unknown
    function _verifyKnown(AgentStatus memory status) internal pure {
        require(status.flag != AgentFlag.Unknown, status.domain == 0 ? "Not a known guard" : "Not a known notary");
    }

    /// @dev Checks that Agent is not Fraudulent/Slashed
    function _verifyNotSlashed(AgentStatus memory status) internal pure {
        require(
            status.flag != AgentFlag.Fraudulent && status.flag != AgentFlag.Slashed,
            status.domain == 0 ? "Slashed guard" : "Slashed notary"
        );
    }

    // ════════════════════════════════════════════ STATEMENT WRAPPERS ═════════════════════════════════════════════════

    // These functions are implemented to reduce the amount of imports in the child contracts.

    /// @dev Wraps Attestation payload into a typed memory view. Reverts if not properly formatted.
    function _wrapAttestation(bytes memory attPayload) internal pure returns (Attestation) {
        return attPayload.castToAttestation();
    }

    /// @dev Wraps AttestationReport payload into a typed memory view. Reverts if not properly formatted.
    function _wrapAttestationReport(bytes memory arPayload) internal pure returns (AttestationReport) {
        return arPayload.castToAttestationReport();
    }

    /// @dev Wraps Receipt payload into a typed memory view. Reverts if not properly formatted.
    function _wrapReceipt(bytes memory rcptPayload) internal pure returns (Receipt) {
        return rcptPayload.castToReceipt();
    }

    /// @dev Wraps Snapshot payload into a typed memory view. Reverts if not properly formatted.
    function _wrapSnapshot(bytes memory snapPayload) internal pure returns (Snapshot) {
        return snapPayload.castToSnapshot();
    }

    /// @dev Wraps State payload into a typed memory view. Reverts if not properly formatted.
    function _wrapState(bytes memory statePayload) internal pure returns (State) {
        return statePayload.castToState();
    }

    /// @dev Wraps StateReport payload into a typed memory view. Reverts if not properly formatted.
    function _wrapStateReport(bytes memory srPayload) internal pure returns (StateReport) {
        return srPayload.castToStateReport();
    }
}
