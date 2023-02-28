// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// TODO: rename to just "Attestation" once older version is deprecated?
import { ByteString, TypedMemView } from "./ByteString.sol";
import { Snapshot } from "./Snapshot.sol";
import { SNAP_ATTESTATION_LENGTH } from "./Structures.sol";

/// @dev SnapAttestation is a memory view over a formatted snapshot attestation payload.
type SnapAttestation is bytes29;
/// @dev Attach library functions to SnapAttestation
using {
    SnapAttestationLib.unwrap,
    SnapAttestationLib.equalToSummit,
    SnapAttestationLib.toDestinationAttestation,
    SnapAttestationLib.hash,
    SnapAttestationLib.root,
    SnapAttestationLib.height,
    SnapAttestationLib.nonce,
    SnapAttestationLib.blockNumber,
    SnapAttestationLib.timestamp
} for SnapAttestation global;

/// @dev Struct representing SnapAttestation, as it is stored in the Summit contract.
struct SummitAttestation {
    bytes32 root;
    uint8 height;
    uint40 blockNumber;
    uint40 timestamp;
}
/// @dev Attach library functions to SummitAttestation
using { SnapAttestationLib.formatSummitAttestation } for SummitAttestation global;

/// @dev Struct representing SnapAttestation, as it is stored in the Destination contract.
/// mapping (bytes32 root => DestinationAttestation) is supposed to be used
struct DestinationAttestation {
    address notary;
    uint8 height;
    uint32 nonce;
    uint40 destTimestamp;
    // 16 bits left for tight packing
}
/// @dev Attach library functions to DestinationAttestation
using { SnapAttestationLib.isEmpty } for DestinationAttestation global;

library SnapAttestationLib {
    using ByteString for bytes;
    using TypedMemView for bytes29;

    /**
     * @dev SnapAttestation structure represents the "Snapshot Merkle Tree" created from
     * every Notary snapshot accepted by the Summit contract. SnapAttestation includes
     * the root and height of "Snapshot Merkle Tree", as well as additional metadata.
     *
     * Steps for creation of "Snapshot Merkle Tree":
     * 1. The list of hashes is composed for states in the Notary snapshot.
     * 2. The list is padded with zero values until its length is a power of two.
     * 3. Values from the lists are used as leafs and the merkle tree is constructed.
     *
     * Similar to Origin, every derived Notary's "Snapshot Merkle Root" is saved in Summit contract.
     * The main difference is that Origin contract itself is keeping track of an incremental merkle tree,
     * by inserting the hash of the dispatched message and calculating the new "Origin Merkle Root".
     * While Summit relies on Guards and Notaries to provide snapshot data, which is used to calculate the
     * "Snapshot Merkle Root".
     *
     * Origin's State is "state of Origin Merkle Tree after N-th message was dispatched".
     * Summit's SnapAttestation is "data for the N-th accepted Notary Snapshot".
     *
     * SnapAttestation is considered "valid" in Summit contract, if it matches the N-th (nonce)
     * snapshot submitted by Notaries.
     * SnapAttestation is considered "valid" in Origin contract, if its underlying Snapshot is "valid".
     *
     * This means that a snapshot could be "valid" in Summit contract and "invalid" in Origin, if the underlying
     * snapshot is invalid (i.e. one of the states in the list is invalid).
     * The opposite could also be true. If a perfectly valid snapshot was never submitted to Summit, its attestation
     * would be valid in Origin, but invalid in Summit (it was never accepted, so the metadata would be incorrect).
     *
     * SnapAttestation is considered "globally valid", if it is valid in the Summit and all the Origin contracts.
     *
     * @dev Memory layout of SnapAttestation fields
     * [000 .. 032): root           bytes32 32 bytes    Root for "Snapshot Merkle Tree" created from a Notary snapshot
     * [032 .. 033): height         uint8    1 byte     Height of "Snapshot Merkle Tree" created from a Notary snapshot
     * [033 .. 037): nonce          uint32   4 bytes    Total amount of all accepted Notary snapshots
     * [037 .. 042): blockNumber    uint40   5 bytes    Block when this Notary snapshot was accepted in Summit
     * [042 .. 047): timestamp      uint40   5 bytes    Time when this Notary snapshot was accepted in Summit
     *
     * The variables below are not supposed to be used outside of the library directly.
     */

    uint256 private constant OFFSET_ROOT = 0;
    uint256 private constant OFFSET_DEPTH = 32;
    uint256 private constant OFFSET_NONCE = 33;
    uint256 private constant OFFSET_BLOCK_NUMBER = 37;
    uint256 private constant OFFSET_TIMESTAMP = 42;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ATTESTATION                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted Attestation payload with provided fields.
     * @param _root         Snapshot merkle tree's root
     * @param _height       Snapshot merkle tree's height
     * @param _nonce        Attestation Nonce
     * @param _blockNumber  Block number when attestation was created in Summit
     * @param _timestamp    Block timestamp when attestation was created in Summit
     * @return Formatted attestation
     **/
    function formatSnapAttestation(
        bytes32 _root,
        uint8 _height,
        uint32 _nonce,
        uint40 _blockNumber,
        uint40 _timestamp
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_root, _height, _nonce, _blockNumber, _timestamp);
    }

    /**
     * @notice Returns a Attestation view over the given payload.
     * @dev Will revert if the payload is not a attestation.
     */
    function castToSnapAttestation(bytes memory _payload) internal pure returns (SnapAttestation) {
        return castToSnapAttestation(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to a Attestation view.
     * @dev Will revert if the memory view is not over a attestation.
     */
    function castToSnapAttestation(bytes29 _view) internal pure returns (SnapAttestation) {
        require(isSnapAttestation(_view), "Not a snapshot attestation");
        return SnapAttestation.wrap(_view);
    }

    /// @notice Checks that a payload is a formatted SnapAttestation.
    function isSnapAttestation(bytes29 _view) internal pure returns (bool) {
        return _view.len() == SNAP_ATTESTATION_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(SnapAttestation _snapAtt) internal pure returns (bytes29) {
        return SnapAttestation.unwrap(_snapAtt);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SUMMIT ATTESTATION                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted Attestation payload with provided fields.
     * @param _summitAtt    SnapAttestation struct as it stored in Summit contract
     * @param _nonce        Attestation nonce
     * @return Formatted attestation
     */
    function formatSummitAttestation(SummitAttestation memory _summitAtt, uint32 _nonce)
        internal
        pure
        returns (bytes memory)
    {
        return
            formatSnapAttestation({
                _root: _summitAtt.root,
                _height: _summitAtt.height,
                _nonce: _nonce,
                _blockNumber: _summitAtt.blockNumber,
                _timestamp: _summitAtt.timestamp
            });
    }

    /// @notice Checks that a SnapshotAttestation and its Summit representation are equal.
    function equalToSummit(SnapAttestation _snapAtt, SummitAttestation memory _summitAtt)
        internal
        pure
        returns (bool)
    {
        return
            _snapAtt.root() == _summitAtt.root &&
            _snapAtt.height() == _summitAtt.height &&
            _snapAtt.blockNumber() == _summitAtt.blockNumber &&
            _snapAtt.timestamp() == _summitAtt.timestamp;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       DESTINATION ATTESTATION                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function toDestinationAttestation(SnapAttestation _snapAtt, address _notary)
        internal
        view
        returns (DestinationAttestation memory attestation)
    {
        attestation.notary = _notary;
        attestation.height = _snapAtt.height();
        attestation.nonce = _snapAtt.nonce();
        // We need to store the timestamp when attestation was submitted to Destination
        attestation.destTimestamp = uint40(block.timestamp);
    }

    function isEmpty(DestinationAttestation memory _destAtt) internal pure returns (bool) {
        return _destAtt.notary == address(0);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION HASHING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns the hash of a SnapAttestation, that could be later signed by a Notary.
    function hash(SnapAttestation _snapAtt) internal pure returns (bytes32) {
        // Get the underlying memory view
        bytes29 _view = _snapAtt.unwrap();
        // TODO: include Attestation-unique salt in the hash
        return _view.keccak();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns root of the Snapshot merkle tree created in the Summit contract.
    function root(SnapAttestation _snapAtt) internal pure returns (bytes32) {
        bytes29 _view = _snapAtt.unwrap();
        return _view.index({ _index: OFFSET_ROOT, _bytes: 32 });
    }

    /// @notice Returns height of the Snapshot merkle tree created in the Summit contract.
    function height(SnapAttestation _snapAtt) internal pure returns (uint8) {
        bytes29 _view = _snapAtt.unwrap();
        return uint8(_view.indexUint({ _index: OFFSET_DEPTH, _bytes: 1 }));
    }

    /// @notice Returns nonce of Summit contract at the time, when attestation was created.
    function nonce(SnapAttestation _snapAtt) internal pure returns (uint32) {
        bytes29 _view = _snapAtt.unwrap();
        return uint32(_view.indexUint({ _index: OFFSET_NONCE, _bytes: 4 }));
    }

    /// @notice Returns a block number when attestation was created in Summit.
    function blockNumber(SnapAttestation _snapAtt) internal pure returns (uint40) {
        bytes29 _view = _snapAtt.unwrap();
        return uint40(_view.indexUint({ _index: OFFSET_BLOCK_NUMBER, _bytes: 5 }));
    }

    /// @notice Returns a block timestamp when attestation was created in Summit.
    /// @dev This is the timestamp according to the Synapse Chain.
    function timestamp(SnapAttestation _snapAtt) internal pure returns (uint40) {
        bytes29 _view = _snapAtt.unwrap();
        return uint40(_view.indexUint({ _index: OFFSET_TIMESTAMP, _bytes: 5 }));
    }
}
