// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { Attestation, AttestationLib } from "../libs/Attestation.sol";
import { Snapshot, SnapshotLib, State, StateLib } from "../libs/Snapshot.sol";
import { AttestationReport, AttestationReportLib } from "../libs/AttestationReport.sol";
import { MerkleLib } from "../libs/Merkle.sol";
import { StateReport, StateReportLib } from "../libs/StateReport.sol";
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
     *  - Report signer is not an active Guard.
     * @param _report           Typed memory view over report payload
     * @param _arSignature      Guard signature for the report
     * @return guard            Guard that signed the report
     */
    function _verifyAttestationReport(AttestationReport _report, bytes memory _arSignature)
        internal
        view
        returns (address guard)
    {
        // This will revert if signer is not an active agent
        uint32 domain;
        (domain, guard) = _recoverAgent(_report.hash(), _arSignature);
        // Report signer needs to be a Guard, not a Notary
        require(domain == 0, "Signer is not a Guard");
    }

    /**
     * @dev Internal function to verify the signed snapshot report payload.
     * Reverts if any of these is true:
     *  - Report signer is not an active Guard.
     * @param _report           Typed memory view over report payload
     * @param _srSignature      Guard signature for the report
     * @return guard            Guard that signed the report
     */
    function _verifyStateReport(StateReport _report, bytes memory _srSignature)
        internal
        view
        returns (address guard)
    {
        // This will revert if signer is not an active agent
        uint32 domain;
        (domain, guard) = _recoverAgent(_report.hash(), _srSignature);
        // Report signer needs to be a Guard, not a Notary
        require(domain == 0, "Signer is not a Guard");
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
     * @dev Internal function to verify that snapshot and attestation has the same Merkle Data.
     * Reverts if any of these is true:
     *  - Attestation root is not equal to root derived from the snapshot.
     *  - Attestation height is not equal to snapshot's height.
     * @param _att          Typed memory view over Attestation
     * @param _snapshot     Typed memory view over snapshot payload
     */
    function _verifySnapshotMerkle(Attestation _att, Snapshot _snapshot) internal pure {
        require(_att.root() == _snapshot.root(), "Incorrect snapshot root");
        require(_att.height() == _snapshot.height(), "Incorrect snapshot height");
    }

    /**
     * @dev Internal function to verify that snapshot roots match.
     * Reverts if any of these is true:
     *  - Attestation root is not equal to Merkle Root derived from State and Snapshot Proof.
     *  - Snapshot Proof has length different to Attestation height.
     *  - Snapshot Proof's first element does not match the State metadata.
     *  - State index is out of range.
     * @param _att              Typed memory view over Attestation
     * @param _stateIndex       Index of state in the snapshot
     * @param _state            Typed memory view over the provided state payload
     * @param _snapProof        Raw payload with snapshot data
     */
    function _verifySnapshotMerkle(
        Attestation _att,
        uint256 _stateIndex,
        State _state,
        bytes32[] memory _snapProof
    ) internal pure {
        // Snapshot proof length should match attestation height (and should be non-zero)
        require(
            _snapProof.length == _att.height() && _snapProof.length != 0,
            "Incorrect proof length"
        );
        // Snapshot proof first element should match State metadata (aka "right sub-leaf")
        (, bytes32 rightSubLeaf) = _state.subLeafs();
        require(_snapProof[0] == rightSubLeaf, "Incorrect proof[0]");
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if state index is out of range
        bytes32 snapshotRoot = _snapshotRoot(
            _state.root(),
            _state.origin(),
            _snapProof,
            _stateIndex
        );
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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          STATEMENT WRAPPERS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // These functions are implemented to reduce the amount of imports in the child contracts.

    /// @dev Wraps Attestation payload into a typed memory view. Reverts if not properly formatted.
    function _wrapAttestation(bytes memory _attPayload) internal pure returns (Attestation) {
        return _attPayload.castToAttestation();
    }

    /// @dev Wraps AttestationReport payload into a typed memory view. Reverts if not properly formatted.
    function _wrapAttestationReport(bytes memory _arPayload)
        internal
        pure
        returns (AttestationReport)
    {
        return _arPayload.castToAttestationReport();
    }

    /// @dev Wraps Snapshot payload into a typed memory view. Reverts if not properly formatted.
    function _wrapSnapshot(bytes memory _snapPayload) internal pure returns (Snapshot) {
        return _snapPayload.castToSnapshot();
    }

    /// @dev Wraps State payload into a typed memory view. Reverts if not properly formatted.
    function _wrapState(bytes memory _statePayload) internal pure returns (State) {
        return _statePayload.castToState();
    }

    /// @dev Wraps StateReport payload into a typed memory view. Reverts if not properly formatted.
    function _wrapStateReport(bytes memory _srPayload) internal pure returns (StateReport) {
        return _srPayload.castToStateReport();
    }
}
