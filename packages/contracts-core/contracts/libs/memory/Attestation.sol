// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {MemView, MemViewLib} from "./MemView.sol";
import {ATTESTATION_LENGTH, ATTESTATION_VALID_SALT, ATTESTATION_INVALID_SALT} from "../Constants.sol";
import {UnformattedAttestation} from "../Errors.sol";

/// Attestation is a memory view over a formatted attestation payload.
type Attestation is uint256;

using AttestationLib for Attestation global;

/// # Attestation
/// Attestation structure represents the "Snapshot Merkle Tree" created from
/// every Notary snapshot accepted by the Summit contract. Attestation includes"
/// the root of the "Snapshot Merkle Tree", as well as additional metadata.
///
/// ## Steps for creation of "Snapshot Merkle Tree":
/// 1. The list of hashes is composed for states in the Notary snapshot.
/// 2. The list is padded with zero values until its length is 2**SNAPSHOT_TREE_HEIGHT.
/// 3. Values from the list are used as leafs and the merkle tree is constructed.
///
/// ## Differences between a State and Attestation
/// Similar to Origin, every derived Notary's "Snapshot Merkle Root" is saved in Summit contract.
/// The main difference is that Origin contract itself is keeping track of an incremental merkle tree,
/// by inserting the hash of the sent message and calculating the new "Origin Merkle Root".
/// While Summit relies on Guards and Notaries to provide snapshot data, which is used to calculate the
/// "Snapshot Merkle Root".
///
/// - Origin's State is "state of Origin Merkle Tree after N-th message was sent".
/// - Summit's Attestation is "data for the N-th accepted Notary Snapshot" + "agent merkle root at the
/// time snapshot was submitted" + "attestation metadata".
///
/// ## Attestation validity
/// - Attestation is considered "valid" in Summit contract, if it matches the N-th (nonce)
/// snapshot submitted by Notaries, as well as the historical agent merkle root.
/// - Attestation is considered "valid" in Origin contract, if its underlying Snapshot is "valid".
///
/// - This means that a snapshot could be "valid" in Summit contract and "invalid" in Origin, if the underlying
/// snapshot is invalid (i.e. one of the states in the list is invalid).
/// - The opposite could also be true. If a perfectly valid snapshot was never submitted to Summit, its attestation
/// would be valid in Origin, but invalid in Summit (it was never accepted, so the metadata would be incorrect).
///
/// - Attestation is considered "globally valid", if it is valid in the Summit and all the Origin contracts.
/// # Memory layout of Attestation fields
///
/// | Position   | Field       | Type    | Bytes | Description                                                    |
/// | ---------- | ----------- | ------- | ----- | -------------------------------------------------------------- |
/// | [000..032) | snapRoot    | bytes32 | 32    | Root for "Snapshot Merkle Tree" created from a Notary snapshot |
/// | [032..064) | dataHash    | bytes32 | 32    | Agent Root and SnapGasHash combined into a single hash         |
/// | [064..068) | nonce       | uint32  | 4     | Total amount of all accepted Notary snapshots                  |
/// | [068..073) | blockNumber | uint40  | 5     | Block when this Notary snapshot was accepted in Summit         |
/// | [073..078) | timestamp   | uint40  | 5     | Time when this Notary snapshot was accepted in Summit          |
///
/// @dev Attestation could be signed by a Notary and submitted to `Destination` in order to use if for proving
/// messages coming from origin chains that the initial snapshot refers to.
library AttestationLib {
    using MemViewLib for bytes;

    // TODO: compress three hashes into one?

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_SNAP_ROOT = 0;
    uint256 private constant OFFSET_DATA_HASH = 32;
    uint256 private constant OFFSET_NONCE = 64;
    uint256 private constant OFFSET_BLOCK_NUMBER = 68;
    uint256 private constant OFFSET_TIMESTAMP = 73;

    // ════════════════════════════════════════════════ ATTESTATION ════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Attestation payload with provided fields.
     * @param snapRoot_     Snapshot merkle tree's root
     * @param dataHash_     Agent Root and SnapGasHash combined into a single hash
     * @param nonce_        Attestation Nonce
     * @param blockNumber_  Block number when attestation was created in Summit
     * @param timestamp_    Block timestamp when attestation was created in Summit
     * @return Formatted attestation
     */
    function formatAttestation(
        bytes32 snapRoot_,
        bytes32 dataHash_,
        uint32 nonce_,
        uint40 blockNumber_,
        uint40 timestamp_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(snapRoot_, dataHash_, nonce_, blockNumber_, timestamp_);
    }

    /**
     * @notice Returns an Attestation view over the given payload.
     * @dev Will revert if the payload is not an attestation.
     */
    function castToAttestation(bytes memory payload) internal pure returns (Attestation) {
        return castToAttestation(payload.ref());
    }

    /**
     * @notice Casts a memory view to an Attestation view.
     * @dev Will revert if the memory view is not over an attestation.
     */
    function castToAttestation(MemView memView) internal pure returns (Attestation) {
        if (!isAttestation(memView)) revert UnformattedAttestation();
        return Attestation.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Attestation.
    function isAttestation(MemView memView) internal pure returns (bool) {
        return memView.len() == ATTESTATION_LENGTH;
    }

    /// @notice Returns the hash of an Attestation, that could be later signed by a Notary to signal
    /// that the attestation is valid.
    function hashValid(Attestation att) internal pure returns (bytes32) {
        // The final hash to sign is keccak(attestationSalt, keccak(attestation))
        return att.unwrap().keccakSalted(ATTESTATION_VALID_SALT);
    }

    /// @notice Returns the hash of an Attestation, that could be later signed by a Guard to signal
    /// that the attestation is invalid.
    function hashInvalid(Attestation att) internal pure returns (bytes32) {
        // The final hash to sign is keccak(attestationInvalidSalt, keccak(attestation))
        return att.unwrap().keccakSalted(ATTESTATION_INVALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Attestation att) internal pure returns (MemView) {
        return MemView.wrap(Attestation.unwrap(att));
    }

    // ════════════════════════════════════════════ ATTESTATION SLICING ════════════════════════════════════════════════

    /// @notice Returns root of the Snapshot merkle tree created in the Summit contract.
    function snapRoot(Attestation att) internal pure returns (bytes32) {
        return att.unwrap().index({index_: OFFSET_SNAP_ROOT, bytes_: 32});
    }

    /// @notice Returns hash of the Agent Root and SnapGasHash combined into a single hash.
    function dataHash(Attestation att) internal pure returns (bytes32) {
        return att.unwrap().index({index_: OFFSET_DATA_HASH, bytes_: 32});
    }

    /// @notice Returns hash of the Agent Root and SnapGasHash combined into a single hash.
    function dataHash(bytes32 agentRoot_, bytes32 snapGasHash_) internal pure returns (bytes32) {
        return keccak256(bytes.concat(agentRoot_, snapGasHash_));
    }

    /// @notice Returns nonce of Summit contract at the time, when attestation was created.
    function nonce(Attestation att) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(att.unwrap().indexUint({index_: OFFSET_NONCE, bytes_: 4}));
    }

    /// @notice Returns a block number when attestation was created in Summit.
    function blockNumber(Attestation att) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(att.unwrap().indexUint({index_: OFFSET_BLOCK_NUMBER, bytes_: 5}));
    }

    /// @notice Returns a block timestamp when attestation was created in Summit.
    /// @dev This is the timestamp according to the Synapse Chain.
    function timestamp(Attestation att) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(att.unwrap().indexUint({index_: OFFSET_TIMESTAMP, bytes_: 5}));
    }
}
