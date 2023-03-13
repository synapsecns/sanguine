// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { Attestation, AttestationLib } from "../libs/Attestation.sol";
import { Snapshot, SnapshotLib } from "../libs/Snapshot.sol";
import { AttestationReport, AttestationReportLib } from "../libs/AttestationReport.sol";
import { MerkleLib } from "../libs/Merkle.sol";
import { State, StateLib, StateReport, StateReportLib } from "../libs/StateReport.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentRegistry } from "../system/AgentRegistry.sol";
import { Versioned } from "../Version.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

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
abstract contract StatementHub is AgentRegistry, Versioned {
    using AttestationLib for bytes;
    using AttestationReportLib for bytes;
    using SnapshotLib for bytes;
    using StateLib for bytes;
    using StateReportLib for bytes;

    // solhint-disable-next-line no-empty-blocks
    constructor() Versioned("0.0.3") {}

    /**
     * @dev Recovers a signer from a hashed message, and a EIP-191 signature for it.
     * Will revert, if the signer is not an active agent.
     * @param _hashedStatement  Hash of the statement that was signed by an Agent
     * @param _signature        Agent signature for the hashed statement
     * @return domain   Domain where the signed Agent is active
     * @return agent    Agent that signed the statement
     */
    function _recoverAgent(bytes32 _hashedStatement, bytes memory _signature)
        internal
        view
        returns (uint32 domain, address agent)
    {
        bytes32 ethSignedMsg = ECDSA.toEthSignedMessageHash(_hashedStatement);
        agent = ECDSA.recover(ethSignedMsg, _signature);
        bool isActive;
        (isActive, domain) = _isActiveAgent(agent);
        require(isActive, "Not an active agent");
    }

    /**
     * @dev Internal function to verify the signed attestation payload.
     * Reverts if any of these is true:
     *  - Attestation payload is not properly formatted.
     *  - Attestation signer is not an active Notary.
     * @param _attPayload       Raw payload with attestation data
     * @param _attSignature     Notary signature for the attestation
     * @return attestation      Typed memory view over attestation payload
     * @return domain           Domain where the signed Notary is active
     * @return notary           Notary that signed the snapshot
     */
    function _verifyAttestation(bytes memory _attPayload, bytes memory _attSignature)
        internal
        view
        returns (
            Attestation attestation,
            uint32 domain,
            address notary
        )
    {
        // This will revert if payload is not a formatted attestation
        attestation = _attPayload.castToAttestation();
        // This will revert if signer is not an active Notary
        (domain, notary) = _verifyAttestation(attestation, _attSignature);
    }

    /**
     * @dev Internal function to verify the signed attestation payload.
     * Reverts if any of these is true:
     *  - Attestation signer is not an active Notary.
     * @param _att              Typed memory view over attestation payload
     * @param _attSignature     Notary signature for the attestation
     * @return domain           Domain where the signed Notary is active
     * @return notary           Notary that signed the snapshot
     */
    function _verifyAttestation(Attestation _att, bytes memory _attSignature)
        internal
        view
        returns (uint32 domain, address notary)
    {
        // This will revert if signer is not an active agent
        (domain, notary) = _recoverAgent(_att.hash(), _attSignature);
        // Attestation signer needs to be a Notary, not a Guard
        require(domain != 0, "Signer is not a Notary");
    }

    /**
     * @dev Internal function to verify the signed attestation report payload.
     * Reverts if any of these is true:
     *  - Report payload is not properly formatted AttestationReport.
     *  - Report signer is not an active Guard.
     * @param _arPayload        Raw payload with report data
     * @param _arSignature      Guard signature for the report
     * @return report           Typed memory view over report payload
     * @return guard            Guard that signed the report
     */
    function _verifyAttestationReport(bytes memory _arPayload, bytes memory _arSignature)
        internal
        view
        returns (AttestationReport report, address guard)
    {
        // This will revert if payload is not a formatted attestation report
        report = _arPayload.castToAttestationReport();
        // This will revert if signer is not an active agent
        uint32 domain;
        (domain, guard) = _recoverAgent(report.hash(), _arSignature);
        // Report signer needs to be a Guard, not a Notary
        require(domain == 0, "Signer is not a Guard");
    }

    /**
     * @dev Internal function to verify the signed snapshot payload.
     * Reverts if any of these is true:
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Agent.
     * @param _snapPayload      Raw payload with snapshot data
     * @param _snapSignature    Agent signature for the snapshot
     * @return snapshot         Typed memory view over snapshot payload
     * @return domain           Domain where the signed Agent is active
     * @return agent            Agent that signed the snapshot
     */
    function _verifySnapshot(bytes memory _snapPayload, bytes memory _snapSignature)
        internal
        view
        returns (
            Snapshot snapshot,
            uint32 domain,
            address agent
        )
    {
        // This will revert if payload is not a formatted snapshot
        snapshot = _snapPayload.castToSnapshot();
        // This will revert if signer is not an active agent
        (domain, agent) = _verifySnapshot(snapshot, _snapSignature);
    }

    /**
     * @dev Internal function to verify the signed snapshot payload.
     * Reverts if any of these is true:
     *  - Snapshot signer is not an active Agent.
     * @param _snapshot         Typed memory view over snapshot payload
     * @param _snapSignature    Agent signature for the snapshot
     * @return domain           Domain where the signed Agent is active
     * @return agent            Agent that signed the snapshot
     */
    function _verifySnapshot(Snapshot _snapshot, bytes memory _snapSignature)
        internal
        view
        returns (uint32 domain, address agent)
    {
        // This will revert if signer is not an active agent
        (domain, agent) = _recoverAgent(_snapshot.hash(), _snapSignature);
        // Guards and Notaries for all domains could sign Snapshots, no further checks are needed.
    }

    /**
     * @dev Internal function to verify the signed snapshot report payload.
     * Reverts if any of these is true:
     *  - Report payload is not properly formatted StateReport.
     *  - Report signer is not an active Guard.
     * @param _srPayload        Raw payload with report data
     * @param _srSignature      Guard signature for the report
     * @return report           Typed memory view over report payload
     * @return guard            Guard that signed the report
     */
    function _verifyStateReport(bytes memory _srPayload, bytes memory _srSignature)
        internal
        view
        returns (StateReport report, address guard)
    {
        // This will revert if payload is not a formatted snapshot report
        report = _srPayload.castToStateReport();
        // This will revert if signer is not an active agent
        uint32 domain;
        (domain, guard) = _recoverAgent(report.hash(), _srSignature);
        // Report signer needs to be a Guard, not a Notary
        require(domain == 0, "Signer is not a Guard");
    }

    /**
     * @dev Internal function to verify that snapshot root matches the root from Attestation.
     * Reverts if any of these is true:
     *  - Snapshot payload is not properly formatted.
     *  - Attestation root is not equal to root derived from the snapshot.
     * @param _att          Typed memory view over Attestation
     * @param _snapPayload  Raw payload with snapshot data
     * @return snapshot     Typed memory view over snapshot payload
     */
    function _verifySnapshotRoot(Attestation _att, bytes memory _snapPayload)
        internal
        pure
        returns (Snapshot snapshot)
    {
        // This will revert if payload is not a formatted snapshot
        snapshot = _snapPayload.castToSnapshot();
        require(_att.root() == snapshot.root(), "Incorrect snapshot root");
    }

    /**
     * @dev Internal function to verify that snapshot roots match.
     * Reverts if any of these is true:
     *  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
     *  - Snapshot Proof has length different to Attestation height.
     *  - Snapshot Proof's first element does not match the State metadata.
     *  - State payload is not properly formatted.
     *  - State index is out of range.
     * @param _att              Typed memory view over Attestation
     * @param _stateIndex       Index of state in the snapshot
     * @param _statePayload     State to check
     * @param _snapProof        Raw payload with snapshot data
     * @return state            Typed memory view over the provided state payload
     */
    function _verifySnapshotRoot(
        Attestation _att,
        uint256 _stateIndex,
        bytes memory _statePayload,
        bytes32[] memory _snapProof
    ) internal pure returns (State state) {
        // This will revert if payload is not a state payload
        state = _statePayload.castToState();
        // Snapshot proof length should match attestation height (and should be non-zero)
        require(
            _snapProof.length == _att.height() && _snapProof.length != 0,
            "Incorrect proof length"
        );
        // Snapshot proof first element should match State metadata (aka "right sub-leaf")
        (, bytes32 rightSubLeaf) = state.subLeafs();
        require(_snapProof[0] == rightSubLeaf, "Incorrect proof[0]");
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if state index is out of range
        bytes32 snapshotRoot = _snapshotRoot(state.root(), state.origin(), _snapProof, _stateIndex);
        // Snapshot root should match the attestation root
        require(_att.root() == snapshotRoot, "Incorrect snapshot root");
    }

    /**
     * @dev Reconstructs Snapshot merkle Root from State Merkle Data (root + origin domain)
     * and proof of inclusion of State Merkle Data (aka State "left sub-leaf") in Snapshot Merkle Tree.
     * Reverts if any of these is true:
     *  - State index is out of range.
     * @param _originRoot   Root of Origin Merkle Tree
     * @param _origin       Domain of Origin chain
     * @param _snapProof    Proof of inclusion of State Merkle Data into Snapshot Merkle Tree
     * @param _stateIndex   Index of Origin State in the Snapshot
     */
    function _snapshotRoot(
        bytes32 _originRoot,
        uint32 _origin,
        bytes32[] memory _snapProof,
        uint256 _stateIndex
    ) internal pure returns (bytes32 snapshotRoot) {
        // Index of "leftLeaf" is twice the state position in the snapshot
        uint256 _leftLeafIndex = _stateIndex << 1;
        // Check that "leftLeaf" index fits into Merkle Tree with specified height
        require(_leftLeafIndex < (1 << _snapProof.length), "State index out of range");
        // Reconstruct left sub-leaf of the Origin State: (originRoot, originDomain)
        bytes32 leftLeaf = StateLib.leftLeaf(_originRoot, _origin);
        // Reconstruct snapshot root using proof of inclusion
        return MerkleLib.branchRoot(leftLeaf, _snapProof, _stateIndex << 1);
    }
}
