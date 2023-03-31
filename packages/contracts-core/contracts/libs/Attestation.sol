// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ByteString } from "./ByteString.sol";
import { ATTESTATION_LENGTH, ATTESTATION_SALT } from "./Constants.sol";
import { TypedMemView } from "./TypedMemView.sol";

/// @dev Attestation is a memory view over a formatted attestation payload.
type Attestation is bytes29;
/// @dev Attach library functions to Attestation
using {
    AttestationLib.unwrap,
    AttestationLib.equalToSummit,
    AttestationLib.toExecutionAttestation,
    AttestationLib.hash,
    AttestationLib.snapRoot,
    AttestationLib.agentRoot,
    AttestationLib.nonce,
    AttestationLib.blockNumber,
    AttestationLib.timestamp
} for Attestation global;

/// @dev Struct representing Attestation, as it is stored in the Summit contract.
struct SummitAttestation {
    bytes32 snapRoot;
    bytes32 agentRoot;
    uint40 blockNumber;
    uint40 timestamp;
}
/// @dev Attach library functions to SummitAttestation
using { AttestationLib.formatSummitAttestation } for SummitAttestation global;

/// @dev Struct representing Attestation, as it is stored in the ExecutionHub contract.
/// mapping (bytes32 root => ExecutionAttestation) is supposed to be used
struct ExecutionAttestation {
    address notary;
    uint32 nonce;
    uint40 submittedAt;
    // 24 bits left for tight packing
}
/// @dev Attach library functions to ExecutionAttestation
using { AttestationLib.isEmpty } for ExecutionAttestation global;

library AttestationLib {
    using ByteString for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev Attestation structure represents the "Snapshot Merkle Tree" created from
     * every Notary snapshot accepted by the Summit contract. Attestation includes
     * the root of the "Snapshot Merkle Tree", as well as additional metadata.
     *
     * Steps for creation of "Snapshot Merkle Tree":
     * 1. The list of hashes is composed for states in the Notary snapshot.
     * 2. The list is padded with zero values until its length is 2**SNAPSHOT_TREE_HEIGHT.
     * 3. Values from the list are used as leafs and the merkle tree is constructed.
     *
     * Similar to Origin, every derived Notary's "Snapshot Merkle Root" is saved in Summit contract.
     * The main difference is that Origin contract itself is keeping track of an incremental merkle tree,
     * by inserting the hash of the dispatched message and calculating the new "Origin Merkle Root".
     * While Summit relies on Guards and Notaries to provide snapshot data, which is used to calculate the
     * "Snapshot Merkle Root".
     *
     * Origin's State is "state of Origin Merkle Tree after N-th message was dispatched".
     * Summit's Attestation is "data for the N-th accepted Notary Snapshot".
     *
     * Attestation is considered "valid" in Summit contract, if it matches the N-th (nonce)
     * snapshot submitted by Notaries.
     * Attestation is considered "valid" in Origin contract, if its underlying Snapshot is "valid".
     *
     * This means that a snapshot could be "valid" in Summit contract and "invalid" in Origin, if the underlying
     * snapshot is invalid (i.e. one of the states in the list is invalid).
     * The opposite could also be true. If a perfectly valid snapshot was never submitted to Summit, its attestation
     * would be valid in Origin, but invalid in Summit (it was never accepted, so the metadata would be incorrect).
     *
     * Attestation is considered "globally valid", if it is valid in the Summit and all the Origin contracts.
     *
     * @dev Memory layout of Attestation fields
     * [000 .. 032): snapRoot       bytes32 32 bytes    Root for "Snapshot Merkle Tree" created from a Notary snapshot
     * [032 .. 064): agentRoot      bytes32 32 bytes    Root for "Agent Merkle Tree" tracked by BondingManager
     * [064 .. 068): nonce          uint32   4 bytes    Total amount of all accepted Notary snapshots
     * [068 .. 073): blockNumber    uint40   5 bytes    Block when this Notary snapshot was accepted in Summit
     * [073 .. 078): timestamp      uint40   5 bytes    Time when this Notary snapshot was accepted in Summit
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_SNAP_ROOT = 0;
    uint256 private constant OFFSET_AGENT_ROOT = 32;
    uint256 private constant OFFSET_NONCE = 64;
    uint256 private constant OFFSET_BLOCK_NUMBER = 68;
    uint256 private constant OFFSET_TIMESTAMP = 73;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ATTESTATION                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted Attestation payload with provided fields.
     * @param snapRoot_     Snapshot merkle tree's root
     * @param agentRoot_    Agent merkle tree's root
     * @param nonce_        Attestation Nonce
     * @param blockNumber_  Block number when attestation was created in Summit
     * @param timestamp_    Block timestamp when attestation was created in Summit
     * @return Formatted attestation
     **/
    function formatAttestation(
        bytes32 snapRoot_,
        bytes32 agentRoot_,
        uint32 nonce_,
        uint40 blockNumber_,
        uint40 timestamp_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(snapRoot_, agentRoot_, nonce_, blockNumber_, timestamp_);
    }

    /**
     * @notice Returns an Attestation view over the given payload.
     * @dev Will revert if the payload is not an attestation.
     */
    function castToAttestation(bytes memory payload) internal pure returns (Attestation) {
        return castToAttestation(payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to an Attestation view.
     * @dev Will revert if the memory view is not over an attestation.
     */
    function castToAttestation(bytes29 view_) internal pure returns (Attestation) {
        require(isAttestation(view_), "Not an attestation");
        return Attestation.wrap(view_);
    }

    /// @notice Checks that a payload is a formatted Attestation.
    function isAttestation(bytes29 view_) internal pure returns (bool) {
        return view_.len() == ATTESTATION_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Attestation att) internal pure returns (bytes29) {
        return Attestation.unwrap(att);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SUMMIT ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted Attestation payload with provided fields.
     * @param summitAtt     Attestation struct as it stored in Summit contract
     * @param nonce_        Attestation nonce
     * @return Formatted attestation
     */
    function formatSummitAttestation(SummitAttestation memory summitAtt, uint32 nonce_)
        internal
        pure
        returns (bytes memory)
    {
        return
            formatAttestation({
                snapRoot_: summitAtt.snapRoot,
                agentRoot_: summitAtt.agentRoot,
                nonce_: nonce_,
                blockNumber_: summitAtt.blockNumber,
                timestamp_: summitAtt.timestamp
            });
    }

    /// @notice Returns an empty struct to save in Summit contract upon initialization.
    // solhint-disable-next-line ordering
    function emptySummitAttestation() internal view returns (SummitAttestation memory) {
        return summitAttestation(bytes32(0), bytes32(0));
    }

    /// @notice Returns a struct to save in the Summit contract for the given root and height.
    function summitAttestation(bytes32 snapRoot_, bytes32 agentRoot_)
        internal
        view
        returns (SummitAttestation memory summitAtt)
    {
        summitAtt.snapRoot = snapRoot_;
        summitAtt.agentRoot = agentRoot_;
        summitAtt.blockNumber = uint40(block.number);
        summitAtt.timestamp = uint40(block.timestamp);
    }

    /// @notice Checks that an Attestation and its Summit representation are equal.
    function equalToSummit(Attestation att, SummitAttestation memory summitAtt)
        internal
        pure
        returns (bool)
    {
        return
            att.snapRoot() == summitAtt.snapRoot &&
            att.agentRoot() == summitAtt.agentRoot &&
            att.blockNumber() == summitAtt.blockNumber &&
            att.timestamp() == summitAtt.timestamp;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       DESTINATION ATTESTATION                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function toExecutionAttestation(Attestation att, address notary)
        internal
        view
        returns (ExecutionAttestation memory attestation)
    {
        attestation.notary = notary;
        attestation.nonce = att.nonce();
        // We need to store the timestamp when attestation was submitted to Destination
        attestation.submittedAt = uint40(block.timestamp);
    }

    function isEmpty(ExecutionAttestation memory execAtt) internal pure returns (bool) {
        return execAtt.notary == address(0);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION HASHING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns the hash of an Attestation, that could be later signed by a Notary.
    function hash(Attestation att) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 view_ = att.unwrap();
        // The final hash to sign is keccak(attestationSalt, keccak(attestation))
        return keccak256(bytes.concat(ATTESTATION_SALT, view_.keccak()));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns root of the Snapshot merkle tree created in the Summit contract.
    function snapRoot(Attestation att) internal pure returns (bytes32) {
        bytes29 view_ = att.unwrap();
        return view_.index({ index_: OFFSET_SNAP_ROOT, bytes_: 32 });
    }

    /// @notice Returns root of the Agent merkle tree tracked by BondingManager.
    function agentRoot(Attestation att) internal pure returns (bytes32) {
        bytes29 view_ = att.unwrap();
        return view_.index({ index_: OFFSET_AGENT_ROOT, bytes_: 32 });
    }

    /// @notice Returns nonce of Summit contract at the time, when attestation was created.
    function nonce(Attestation att) internal pure returns (uint32) {
        bytes29 view_ = att.unwrap();
        return uint32(view_.indexUint({ index_: OFFSET_NONCE, bytes_: 4 }));
    }

    /// @notice Returns a block number when attestation was created in Summit.
    function blockNumber(Attestation att) internal pure returns (uint40) {
        bytes29 view_ = att.unwrap();
        return uint40(view_.indexUint({ index_: OFFSET_BLOCK_NUMBER, bytes_: 5 }));
    }

    /// @notice Returns a block timestamp when attestation was created in Summit.
    /// @dev This is the timestamp according to the Synapse Chain.
    function timestamp(Attestation att) internal pure returns (uint40) {
        bytes29 view_ = att.unwrap();
        return uint40(view_.indexUint({ index_: OFFSET_TIMESTAMP, bytes_: 5 }));
    }
}
