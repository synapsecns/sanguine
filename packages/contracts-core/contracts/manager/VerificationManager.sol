// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Attestation} from "../libs/Attestation.sol";
import {AttestationReport} from "../libs/AttestationReport.sol";
import {MerkleMath} from "../libs/MerkleMath.sol";
import {Receipt} from "../libs/Receipt.sol";
import {Snapshot, SNAPSHOT_TREE_HEIGHT} from "../libs/Snapshot.sol";
import {State, StateLib} from "../libs/State.sol";
import {StateReport} from "../libs/StateReport.sol";
import {AgentFlag, AgentStatus} from "../libs/Structures.sol";

/// @notice VerificationManager is a stateless contract responsible for verifying agent signatures,
/// as well as some common basic checks for the agent statements or the agent statuses.
abstract contract VerificationManager {
    /// @dev gap for upgrade safety
    uint256[50] private __GAP; // solhint-disable-line var-name-mixedcase

    /**
     * @dev Recovers a signer from a hashed message, and a EIP-191 signature for it.
     * Will revert, if the signer is not a known agent.
     * @dev Agent flag could be any of these: Active/Unstaking/Resting/Slashed
     * Further checks need to be performed in a caller function.
     * @param hashedStatement   Hash of the statement that was signed by an Agent
     * @param signature         Agent signature for the hashed statement
     * @return status   Struct representing agent status:
     *                  - flag      Unknown/Active/Unstaking/Resting/Slashed
     *                  - domain    Domain where agent is/was active
     *                  - index     Index of agent in the Agent Merkle Tree
     * @return agent    Agent that signed the statement
     */
    function _recoverAgent(bytes32 hashedStatement, bytes memory signature)
        internal
        view
        virtual
        returns (AgentStatus memory status, address agent);

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

    // ═══════════════════════════════════════ STATE/SNAPSHOT RELATED CHECKS ═══════════════════════════════════════════

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
}
