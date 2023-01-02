// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./ByteString.sol";

/// @dev Attestation is a memory view over a formatted attestation payload.
type Attestation is bytes29;
/// @dev AttestationData is a memory view over a formatted attestation data.
type AttestationData is bytes29;

library AttestationLib {
    // The goal of having custom types is to assign a type for every generic memory view (bytes29)
    // and enforce compile-time strict type checking for every operation with the memory views.
    // This will prevent a misuse of libraries, i.e. using Attestation functions on a Report view.
    // Every type is supposed to define a method to wrap a generic memory view into a given type,
    // while checking that the view is over a properly formatted payload: `castToAttestation()`.
    // Different types may define methods with the same name without any issues:
    //      Message msg;
    //      msg.nonce();    // gets a message nonce
    //      AttestationData data;
    //      data.nonce();   // gets an attestation nonce

    using ByteString for bytes;
    using ByteString for bytes29;

    using TypedMemView for bytes29;

    /**
     * @dev AttestationData memory layout
     * [000 .. 004): origin         uint32   4 bytes
     * [004 .. 008): destination    uint32   4 bytes
     * [008 .. 012): nonce          uint32   4 bytes
     * [012 .. 044): root           bytes32 32 bytes
     * [044 .. 049): blockNumber    uint40   5 bytes
     * [049 .. 054): timestamp      uint40   5 bytes
     *
     *      Attestation memory layout
     * [000 .. 054): attData        bytes   44 bytes (see above)
     * [054 .. 055): G = guardSigs  uint8    1 byte
     * [055 .. 056): N = notarySigs uint8    1 byte
     * [056 .. 121): guardSig[0]    bytes   65 bytes
     *      ..
     * [AAA .. BBB): guardSig[G-1]  bytes   65 bytes
     * [BBB .. CCC): notarySig[0]   bytes   65 bytes
     *      ..
     * [DDD .. END): notarySig[N-1] bytes   65 bytes
     *
     * END = 56 + 65 * (G + N)
     */

    uint256 internal constant OFFSET_ORIGIN = 0;
    uint256 internal constant OFFSET_DESTINATION = 4;
    uint256 internal constant OFFSET_NONCE = 8;
    uint256 internal constant OFFSET_ROOT = 12;
    uint256 internal constant OFFSET_BLOCK_NUMBER = 44;
    uint256 internal constant OFFSET_TIMESTAMP = 49;
    uint256 internal constant ATTESTATION_DATA_LENGTH = 54;

    uint256 internal constant OFFSET_AGENT_SIGS = ATTESTATION_DATA_LENGTH;
    uint256 internal constant OFFSET_FIRST_SIGNATURE = OFFSET_AGENT_SIGS + 2;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     FORMATTING: COMBINED FIELDS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Combines origin and destination fields into `domains`,
     * a unique ID for every (origin, destination) pair. Could be used to identify
     * Merkle trees on Origin, or Mirrors on Destination.
     */
    function packDomains(uint32 _origin, uint32 _destination)
        internal
        pure
        returns (uint64 _domains)
    {
        _domains = (uint64(_origin) << 32) | _destination;
    }

    /// @notice Unpacks a combined (origin, destination) field into `origin` and `destination`.
    function unpackDomains(uint64 _domains)
        internal
        pure
        returns (uint32 _origin, uint32 _destination)
    {
        // Shift out lower 32 bits
        _origin = uint32(_domains >> 32);
        // Use lower 32 bits
        _destination = uint32(_domains & type(uint32).max);
    }

    /**
     * @notice Combines origin, destination domains and message nonce into `key`,
     * a unique ID for every (origin, destination, nonce) tuple. Could be used to identify
     * any dispatched message.
     */
    function packKey(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce
    ) internal pure returns (uint96 _key) {
        _key = (uint96(_origin) << 64) | (uint96(_destination) << 32) | _nonce;
    }

    /// @notice Unpacks a combined (origin, destination, nonce) field into
    /// `domains = (origin, destination)` and `nonce`.
    function unpackKey(uint96 _key) internal pure returns (uint64 _domains, uint32 _nonce) {
        // Shift out lower 32 bits
        _domains = uint64(_key >> 32);
        // Use lower 32 bits
        _nonce = uint32(_key & type(uint32).max);
    }

    /// @notice Combines amounts of guard and notary signature into `agentSigs`.
    function packAgentsAmount(uint8 _guardsAmount, uint8 _notariesAmount)
        internal
        pure
        returns (uint16 _agentsAmount)
    {
        _agentsAmount = (uint16(_guardsAmount) << 8) | _notariesAmount;
    }

    /// @notice Unpacks a combined (guardSigs, notarySigs) field into `guardSigs` and `notarySigs`.
    function unpackAgentsAmount(uint16 _agentsAmount)
        internal
        pure
        returns (uint8 _guardsAmount, uint8 _notariesAmount)
    {
        // Shift out lower 8 bits
        _guardsAmount = uint8(_agentsAmount >> 8);
        // Use lower 8 bits
        _notariesAmount = uint8(_agentsAmount & type(uint8).max);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           ATTESTATION DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted AttestationData payload with provided fields
     * @param _origin       Domain of Origin's chain
     * @param _destination  Domain of Destination's chain
     * @param _root         New merkle root
     * @param _nonce        Nonce of the merkle root
     * @param _blockNumber  Block number when root was saved in Origin
     * @param _timestamp    Block timestamp when root was saved in Origin
     * @return Formatted attestation data
     **/
    function formatAttestationData(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root,
        uint40 _blockNumber,
        uint40 _timestamp
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(_origin, _destination, _nonce, _root, _blockNumber, _timestamp);
    }

    /**
     * @notice Returns an AttestationData view over the given payload.
     * @dev Will revert if the payload is not an attestation data.
     */
    function castToAttestationData(bytes memory _payload) internal pure returns (AttestationData) {
        return castToAttestationData(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to an AttestationData view.
     * @dev Will revert if the memory view is not over an attestation data.
     */
    function castToAttestationData(bytes29 _view) internal pure returns (AttestationData) {
        require(isAttestationData(_view), "Not an attestation data");
        return AttestationData.wrap(_view);
    }

    /// @notice Checks that a payload is a formatted Attestation Data.
    function isAttestationData(bytes29 _view) internal pure returns (bool) {
        return _view.len() == ATTESTATION_DATA_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(AttestationData _data) internal pure returns (bytes29) {
        return AttestationData.unwrap(_data);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       ATTESTATION DATA SLICING                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns domain of chain where the Origin contract is deployed.
    function origin(AttestationData _data) internal pure returns (uint32) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_data);
        return uint32(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 4 }));
    }

    /// @notice Returns domain of chain where the Destination contract is deployed.
    function destination(AttestationData _data) internal pure returns (uint32) {
        bytes29 _view = unwrap(_data);
        return uint32(_view.indexUint({ _index: OFFSET_DESTINATION, _bytes: 4 }));
    }

    /// @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
    function nonce(AttestationData _data) internal pure returns (uint32) {
        bytes29 _view = unwrap(_data);
        return uint32(_view.indexUint({ _index: OFFSET_NONCE, _bytes: 4 }));
    }

    /// @notice Returns a combined field for (origin, destination). See `attestationDomains()`.
    function domains(AttestationData _data) internal pure returns (uint64) {
        bytes29 _view = unwrap(_data);
        return uint64(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 8 }));
    }

    /// @notice Returns a combined field for (origin, destination, nonce). See `attestationKey()`.
    function key(AttestationData _data) internal pure returns (uint96) {
        bytes29 _view = unwrap(_data);
        return uint96(_view.indexUint({ _index: OFFSET_ORIGIN, _bytes: 12 }));
    }

    /// @notice Returns a historical Merkle root from the Origin contract.
    function root(AttestationData _data) internal pure returns (bytes32) {
        bytes29 _view = unwrap(_data);
        return _view.index({ _index: OFFSET_ROOT, _bytes: 32 });
    }

    /// @notice Returns a block number when `root` was saved in Origin.
    function blockNumber(AttestationData _data) internal pure returns (uint40) {
        bytes29 _view = unwrap(_data);
        return uint40(_view.indexUint({ _index: OFFSET_BLOCK_NUMBER, _bytes: 5 }));
    }

    /// @notice Returns a block timestamp when `root` was saved in Origin.
    /// @dev This is the timestamp according to the origin chain.
    function timestamp(AttestationData _data) internal pure returns (uint40) {
        bytes29 _view = unwrap(_data);
        return uint40(_view.indexUint({ _index: OFFSET_TIMESTAMP, _bytes: 5 }));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             ATTESTATION                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a formatted Attestation payload with provided fields
     * @dev `_guardSigsPayload` and `_notarySigsPayload` payloads could be empty.
     * They have to contain exactly 65 * N bytes, otherwise the execution will be reverted.
     * @param _dataPayload          Attestation Data payload (see above)
     * @param _guardSigsPayload     Payload with all Guard signatures on `_data`
     * @param _notarySigsPayload    Payload with all Notary signatures on `_data`
     * @return attestation Formatted attestation
     **/
    // solhint-disable-next-line ordering
    function formatAttestation(
        bytes memory _dataPayload,
        bytes memory _guardSigsPayload,
        bytes memory _notarySigsPayload
    ) internal view returns (bytes memory attestation) {
        attestation = formatAttestation({
            _data: castToAttestationData(_dataPayload),
            _guardSigs: _guardSigsPayload.castToRawBytes(),
            _notarySigs: _notarySigsPayload.castToRawBytes()
        });
    }

    /**
     * @notice Returns a formatted Attestation payload with provided fields
     * @dev `_guardSigs` and `_notarySigs` memory views could be empty.
     * They have to contain exactly 65 * N bytes, otherwise the execution will be reverted.
     * @param _data         Memory view over Attestation Data (see above)
     * @param _guardSigs    Memory view over the payload with all Guard signatures on `_data`
     * @param _notarySigs   Memory view over the payload with all Notary signatures on `_data`
     * @return attestation Formatted attestation
     **/
    function formatAttestation(
        AttestationData _data,
        bytes29 _guardSigs,
        bytes29 _notarySigs
    ) internal view returns (bytes memory attestation) {
        uint8 _guardsAmount = _signaturesAmount(_guardSigs);
        uint8 _notariesAmount = _signaturesAmount(_notarySigs);
        // Pack (_guardsAmount, _notariesAmount) into a single 16-byte value
        uint16 agents = packAgentsAmount(_guardsAmount, _notariesAmount);
        // We need to join: `_data`, `agents`, `_guardSigs`, `_notarySigs`
        bytes29[] memory allViews = new bytes29[](4);
        allViews[0] = unwrap(_data);
        allViews[1] = abi.encodePacked(agents).castToRawBytes();
        allViews[2] = _guardSigs;
        allViews[3] = _notarySigs;
        attestation = TypedMemView.join(allViews);
    }

    /**
     * @notice Returns an Attestation view over the given payload.
     * @dev Will revert if the payload is not an attestation.
     */
    function castToAttestation(bytes memory _payload) internal pure returns (Attestation) {
        return castToAttestation(_payload.castToRawBytes());
    }

    /**
     * @notice Casts a memory view to an Attestation view.
     * @dev Will revert if the memory view is not over an attestation.
     */
    function castToAttestation(bytes29 _view) internal pure returns (Attestation) {
        require(isAttestation(_view), "Not an attestation");
        return Attestation.wrap(_view);
    }

    /// @notice Checks that a payload is a formatted Attestation.
    function isAttestation(bytes29 _view) internal pure returns (bool) {
        uint256 length = _view.len();
        // (attData, guardSigs, notarySigs) need to exist
        if (length < OFFSET_FIRST_SIGNATURE) return false;
        (uint256 _guardsAmount, uint256 _notariesAmount) = _getAgentsAmount(_view);
        uint256 totalAgents = _guardsAmount + _notariesAmount;
        // There should be at least one signature
        if (totalAgents == 0) return false;
        // Every signature has length of exactly `ByteString.SIGNATURE_LENGTH`
        return length == OFFSET_FIRST_SIGNATURE + totalAgents * ByteString.SIGNATURE_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Attestation _attestation) internal pure returns (bytes29) {
        return Attestation.unwrap(_attestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         ATTESTATION SLICING                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns Attestation's Data, that is going to be signed by the Guards and Notaries.
    function data(Attestation _attestation) internal pure returns (AttestationData) {
        // Get the underlying memory view
        bytes29 _view = unwrap(_attestation);
        return
            castToAttestationData(
                _view.slice({ _index: OFFSET_ORIGIN, _len: ATTESTATION_DATA_LENGTH, newType: 0 })
            );
    }

    /// @notice Returns the amount of guard and notary signatures present in the Attestation.
    function agentsAmount(Attestation _attestation)
        internal
        pure
        returns (uint8 _guardsAmount, uint8 _notariesAmount)
    {
        bytes29 _view = unwrap(_attestation);
        (_guardsAmount, _notariesAmount) = _getAgentsAmount(_view);
    }

    /// @notice Returns the amount of guard signatures present in the Attestation.
    function guardsAmount(Attestation _attestation) internal pure returns (uint8 _guardsAmount) {
        bytes29 _view = unwrap(_attestation);
        (_guardsAmount, ) = _getAgentsAmount(_view);
    }

    /// @notice Returns the amount of notary signatures present in the Attestation.
    function notariesAmount(Attestation _attestation)
        internal
        pure
        returns (uint8 _notariesAmount)
    {
        bytes29 _view = unwrap(_attestation);
        (, _notariesAmount) = _getAgentsAmount(_view);
    }

    /**
     * @notice Returns signature of the i-th Guard on AttestationData,
     * @dev Will revert if index is out of range.
     */
    function guardSignature(Attestation _attestation, uint256 _guardIndex)
        internal
        pure
        returns (Signature)
    {
        bytes29 _view = unwrap(_attestation);
        (uint8 _guardsAmount, ) = _getAgentsAmount(_view);
        require(_guardIndex < _guardsAmount, "Out of range");
        // Determine the signature position in the payload
        uint256 sigPosition = OFFSET_FIRST_SIGNATURE + _guardIndex * ByteString.SIGNATURE_LENGTH;
        return
            _view
                .slice({ _index: sigPosition, _len: ByteString.SIGNATURE_LENGTH, newType: 0 })
                .castToSignature();
    }

    /**
     * @notice Returns signature of the i-th Notary on AttestationData,
     * @dev Will revert if index is out of range.
     */
    function notarySignature(Attestation _attestation, uint256 _notaryIndex)
        internal
        pure
        returns (Signature)
    {
        bytes29 _view = unwrap(_attestation);
        (uint8 _guardsAmount, uint8 _notariesAmount) = _getAgentsAmount(_view);
        require(_notaryIndex < _notariesAmount, "Out of range");
        // Determine the signature position in the payload
        _notaryIndex = _notaryIndex + _guardsAmount;
        uint256 sigPosition = OFFSET_FIRST_SIGNATURE + _notaryIndex * ByteString.SIGNATURE_LENGTH;
        return
            _view
                .slice({ _index: sigPosition, _len: ByteString.SIGNATURE_LENGTH, newType: 0 })
                .castToSignature();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PRIVATE HELPERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns the amount of guard and notary signatures present in the Attestation
    /// without checking if the payload is properly formatted.
    function _getAgentsAmount(bytes29 _view)
        private
        pure
        returns (uint8 _guardsAmount, uint8 _notariesAmount)
    {
        // Read both amounts at once
        uint16 agents = uint16(_view.indexUint({ _index: OFFSET_AGENT_SIGS, _bytes: 2 }));
        (_guardsAmount, _notariesAmount) = unpackAgentsAmount(agents);
    }

    /**
     * @dev Returns the amount of signatures in the "signatures" payload.
     * Reverts, if payload length is not exactly 65 * N bytes.
     * Reverts, if amount of signatures does not fit in `uint8`.
     */
    function _signaturesAmount(bytes29 _sigsView) private pure returns (uint8 amount) {
        uint256 length = _sigsView.len();
        uint256 _amount = length / ByteString.SIGNATURE_LENGTH;
        require(_amount * ByteString.SIGNATURE_LENGTH == length, "!signaturesLength");
        require(_amount < type(uint8).max, "Too many signatures");
        amount = uint8(_amount);
    }
}
