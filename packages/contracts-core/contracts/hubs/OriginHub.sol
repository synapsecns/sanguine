// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../libs/Report.sol";

import { OriginHubEvents } from "../events/OriginHubEvents.sol";
import { ReportHub } from "./ReportHub.sol";
import { SystemRegistry } from "../system/SystemRegistry.sol";

import { MerkleLib } from "../libs/Merkle.sol";

/**
 * @notice Inserts new message hashes into the Merkle Trees by destination chain domain,
 * and keeps track of the historical merkle state for each destination.
 * Keeps track of this domain's Notaries and all Guards: accepts
 * and checks their attestations/reports related to Origin.
 */
abstract contract OriginHub is OriginHubEvents, SystemRegistry, ReportHub {
    using AttestationLib for Attestation;
    using AttestationLib for AttestationData;

    using MerkleLib for MerkleLib.Tree;

    /**
     * @notice Additional data stored for every saved merkle root.
     * @dev uint40 gives us the runway of ~35k years.
     * @param timestamp     Block timestamp when merkle root was saved
     * @param blockNumber   Block number when merkle root was saved
     */
    struct RootMetadata {
        uint40 blockNumber;
        uint40 timestamp;
        // 176 bits remaining
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Merkle root for an empty merkle tree.
    bytes32 internal constant EMPTY_TREE_ROOT =
        hex"27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757";

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // [destination domain] => [Merkle Tree containing all hashes of sent messages to that domain]
    mapping(uint32 => MerkleLib.Tree) internal trees;
    // [destination domain] => [Merkle tree roots after inserting a sent message to that domain]
    mapping(uint32 => bytes32[]) internal historicalRoots;
    // [destination domain] => [block numbers for each nonce written so far]
    mapping(uint32 => RootMetadata[]) internal historicalMetadata;

    // gap for upgrade safety
    uint256[47] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Suggest a list of attestations for the off-chain actors to sign.
     * Returns attestation data for every destination domain having at least one active Notary.
     */
    function suggestAttestations() external view returns (bytes[] memory attestationDataArray) {
        uint256 amount = amountDomains();
        attestationDataArray = new bytes[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            uint32 domain = getDomain(i);
            attestationDataArray[i] = suggestAttestation({ _destination: domain });
        }
    }

    /**
     * @notice Suggest the latest attestation data to sign for a specific destination.
     * Note: signing the suggested attestation data will will never lead to slashing of the actor,
     * assuming they have confirmed that the block, which number is included in the data,
     * is not subject to reorganization (which is different for every observed chain).
     * @dev If no messages to a given destination have been sent, the execution will be reverted.
     * @return attestationData Data for the suggested attestation
     */
    function suggestAttestation(uint32 _destination)
        public
        view
        returns (bytes memory attestationData)
    {
        uint32 latestNonce = nonce(_destination);
        require(latestNonce != 0, "No messages to destination");
        attestationData = suggestHistoricalAttestation(_destination, latestNonce);
    }

    /**
     * @notice Suggest the attestation data to sign for a specific destination and nonce.
     * Note: signing the suggested attestation data will will never lead to slashing of the actor,
     * assuming they have confirmed that the block, which number is included in the data,
     * is not subject to reorganization (which is different for every observed chain).
     * @dev If no messages to a given destination have been sent, the execution will be reverted.
     * @return attestationData Data for the suggested attestation
     */
    function suggestHistoricalAttestation(uint32 _destination, uint32 _nonce)
        public
        view
        returns (bytes memory attestationData)
    {
        // Check if nonce exists
        require(_nonce < historicalRoots[_destination].length, "!nonce");
        bytes32 historicalRoot = historicalRoots[_destination][_nonce];
        RootMetadata memory metadata = historicalMetadata[_destination][_nonce];
        attestationData = AttestationLib.formatAttestationData({
            _origin: _localDomain(),
            _destination: _destination,
            _nonce: _nonce,
            _root: historicalRoot,
            _blockNumber: metadata.blockNumber,
            _timestamp: metadata.timestamp
        });
    }

    /**
     * @notice Returns a historical merkle root for the given destination.
     * Note: signing the attestation with the given historical root will never lead
     * to slashing of the actor, assuming they have confirmed that the block, where the merkle root
     * was updated, is not subject to reorganization (which is different for every observed chain).
     * @param _destination  Destination domain
     * @param _nonce        Historical nonce
     * @return historicalRoot   Root for destination's merkle tree right after
     *                          message to `_destination` with `nonce = _nonce` was dispatched.
     * @return blockNumber      Block number when the above message was dispatched.
     * @return timestamp        Block timestamp when the above message was dispatched.
     */
    function getHistoricalRoot(uint32 _destination, uint32 _nonce)
        public
        view
        returns (
            bytes32 historicalRoot,
            uint40 blockNumber,
            uint40 timestamp
        )
    {
        // Check if destination is known
        if (historicalRoots[_destination].length > 0) {
            // Check if nonce exists
            require(_nonce < historicalRoots[_destination].length, "!nonce: existing destination");
            RootMetadata memory metadata = historicalMetadata[_destination][_nonce];
            historicalRoot = historicalRoots[_destination][_nonce];
            blockNumber = metadata.blockNumber;
            timestamp = metadata.timestamp;
        } else {
            // If destination is unknown, we have the root of an empty merkle tree
            require(_nonce == 0, "!nonce: unknown destination");
            historicalRoot = EMPTY_TREE_ROOT;
            // return (0, 0) for (blockNumber, timestamp)
        }
    }

    /**
     * @notice Returns nonce of the last inserted Merkle root for the given destination,
     * which is also the number of inserted leaves in the destination merkle tree (current index).
     */
    function nonce(uint32 _destination) public view returns (uint32 latestNonce) {
        latestNonce = uint32(_getTreeCount(_destination));
    }

    /**
     * @notice Calculates and returns tree's current root for the given destination.
     */
    function root(uint32 _destination) public view returns (bytes32) {
        return trees[_destination].root(_getTreeCount(_destination));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Checks if a submitted Attestation is a valid Attestation.
     * Attestation Flag can be either "Fraud" or "Valid".
     * A "Fraud" Attestation is a (_destination, _nonce, _root) attestation that doesn't correspond
     * with the historical state of Origin contract. Either of these needs to be true:
     * - _nonce is higher than current nonce for _destination (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce for _destination
     * This would mean that message(s) that were not truly
     * dispatched on Origin were falsely included in the signed root.
     *
     * A Fraud Attestation will only be accepted as valid by the Mirror.
     * If a Fraud Attestation is submitted to the Mirror, a Guard should
     * submit a Fraud Report using Origin.submitReport()
     * in order to slash the Notary with a Fraud Attestation.
     *
     * @dev Guards and Notaries signatures and roles have been checked in AttestationHub.
     *
     * @param _guards       Guard addresses (signatures&roles already verified)
     * @param _notaries     Notary addresses (signatures&roles already verified)
     * @param _att          Memory view over the Attestation for convenience
     * @param _attPayload   Payload with Attestation data and signature
     * @return isValid      TRUE if Attestation was valid (implying no agent was slashed).
     */
    function _handleAttestation(
        address[] memory _guards,
        address[] memory _notaries,
        Attestation _att,
        bytes memory _attPayload
    ) internal override returns (bool isValid) {
        AttestationData attData = _att.data();
        uint32 _origin = attData.origin();
        uint32 _destination = attData.destination();
        uint32 _nonce = attData.nonce();
        bytes32 _root = attData.root();
        RootMetadata memory metadata = RootMetadata({
            blockNumber: attData.blockNumber(),
            timestamp: attData.timestamp()
        });
        isValid = _isValidAttestation(_origin, _destination, _nonce, _root, metadata);
        if (!isValid) {
            emit FraudAttestation(_guards, _notaries, _attPayload);
            // Guard doesn't receive anything, as Agents weren't slashed using the Fraud Report
            _slashAgents({ _domain: 0, _accounts: _guards, _guard: address(0) });
            _slashAgents({ _domain: _destination, _accounts: _notaries, _guard: address(0) });
            /**
             * TODO: design incentives for the reporter in a way, where they get less
             * by reporting directly instead of using a correct Fraud Report.
             * That will allow Guards to focus on Report signing and don't worry
             * about submitReport (whether their own or outsourced) txs being frontrun.
             */
        }
    }

    /**
     * @notice Checks if a submitted Report is a correct Report. Reported Attestation
     * can be either valid or fraud. Report flag can also be either Valid or Fraud.
     * Report is correct if its flag matches the Attestation validity.
     * 1. Attestation: valid, Flag: Fraud.
     *      Report is deemed incorrect, Guard is slashed (if they haven't been already).
     * 2. Attestation: valid, Flag: Valid.
     *      Report is deemed correct, no action is done.
     * 3. Attestation: Fraud, Flag: Fraud.
     *      Report is deemed correct, Notary is slashed (if they haven't been already).
     * 4. Attestation: Fraud, Flag: Valid.
     *      Report is deemed incorrect, Guard is slashed (if they haven't been already).
     *      Notary is slashed (if they haven't been already), but Guard doesn't receive
     *      any rewards (as their report indicated that the attestation was valid).
     *
     * A "Fraud" Attestation is a (_destination, _nonce, _root) attestation that doesn't correspond
     * with the historical state of Origin contract. Either of these needs to be true:
     * - _nonce is higher than current nonce for _destination (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce for _destination
     * This would mean that message(s) that were not truly
     * dispatched on Origin were falsely included in the signed root.
     *
     * A Fraud Attestation will only be accepted as valid by the Mirror.
     * If a Fraud Attestation is submitted to the Mirror, a Guard should
     * submit a Fraud Report using Origin.submitReport()
     * in order to slash the Notary with a Fraud Attestation.
     *
     * @dev Both Notary and Guard signatures and roles have been checked in ReportHub,
     * hence `_notary` is an active Notary, `_guard` is an active Guard at this point.
     *
     * @param _guard            Guard address (signature&role already verified)
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation
     * @param _reportView       Memory view over Report
     * @param _report           Payload with Report data and signature
     * @return TRUE if Report was correct (implying Guard was not slashed)
     */
    function _handleReport(
        address _guard,
        address _notary,
        bytes29 _attestationView,
        bytes29 _reportView,
        bytes memory _report
    ) internal override returns (bool) {
        // TODO(Chi): enable reports once co-signed Attestation is implemented
        /*
        uint32 _origin = _attestationView.attestedOrigin();
        uint32 _destination = _attestationView.attestedDestination();
        uint32 _nonce = _attestationView.attestedNonce();
        bytes32 _root = _attestationView.attestedRoot();
        if (_isValidAttestation(_origin, _destination, _nonce, _root)) {
            // Attestation: Valid
            if (_reportView.reportedFraud()) {
                // Flag: Fraud
                // Report is incorrect, slash the Guard
                emit IncorrectReport(_guard, _report);
                _slashGuard(_guard);
                return false;
            } else {
                // Flag: Valid
                // Report is correct, no action needed
                return true;
            }
        } else {
            // Attestation: Fraud
            if (_reportView.reportedFraud()) {
                // Flag: Fraud
                // Report is correct, slash the Notary
                emit CorrectFraudReport(_guard, _report);
                emit FraudAttestation(_notary, _attestationView.clone());
                _slashNotary({ _domain: _destination, _notary: _notary, _guard: _guard });
                return true;
            } else {
                // Flag: Valid
                // Report is incorrect, slash the Guard
                emit IncorrectReport(_guard, _report);
                _slashGuard(_guard);
                emit FraudAttestation(_notary, _attestationView.clone());
                // Guard doesn't receive anything due to Valid flag on the Report
                _slashNotary({ _domain: _destination, _notary: _notary, _guard: address(0) });
                return false;
            }
        }
        */
    }

    /**
     * @notice Inserts a merkle root for an empty merkle tree into the historical roots array
     * for the given destination.
     * @dev This enables:
     * - Counting nonces from 1 (nonce=0 meaning no messages have been sent).
     * - Not slashing the Notaries for signing an attestation for an empty tree
     * (assuming they sign the correct root outlined below).
     */
    function _initializeHistoricalRoots(uint32 _destination) internal {
        // This function should only be called only if the array is empty
        assert(historicalRoots[_destination].length == 0);
        // Insert a historical root so nonces start at 1 rather then 0.
        // Here we insert the root of an empty merkle tree
        historicalRoots[_destination].push(EMPTY_TREE_ROOT);
        historicalMetadata[_destination].push(RootMetadata(0, 0));
    }

    /**
     * @notice Inserts new message into the Merkle tree for the given destination
     * and stores the new merkle root.
     * @param _destination  Destination domain of the dispatched message
     * @param _messageNonce Nonce of the dispatched message
     * @param _messageHash  Hash of the dispatched message
     */
    function _insertMessage(
        uint32 _destination,
        uint32 _messageNonce,
        bytes32 _messageHash
    ) internal returns (bytes32 newRoot) {
        // TODO: when Notary is active on Destination, initialize historical roots
        // upon adding a first Notary for given destination
        if (historicalRoots[_destination].length == 0) _initializeHistoricalRoots(_destination);
        /// @dev _messageNonce == tree.count() + 1
        // tree.insert() requires amount of leaves AFTER the leaf insertion (i.e. tree.count() + 1)
        trees[_destination].insert(_messageNonce, _messageHash);
        /// @dev leaf is inserted => _messageNonce == tree.count()
        // tree.root() requires current amount of leaves (i.e. tree.count())
        newRoot = trees[_destination].root(_messageNonce);
        historicalRoots[_destination].push(newRoot);
        historicalMetadata[_destination].push(
            RootMetadata({ blockNumber: uint40(block.number), timestamp: uint40(block.timestamp) })
        );
    }

    /**
     * @notice Slashes a few agents, that are active on the same domain.
     * @dev Called when agent fraud is proven.
     * @param _domain   Domain where the reported Agents are active
     * @param _accounts Addresses of the fraudulent Agents
     * @param _guard    Guard who reported the fraudulent Agents [address(0) if not a Guard report]
     */
    function _slashAgents(
        uint32 _domain,
        address[] memory _accounts,
        address _guard
    ) internal {
        uint256 amount = _accounts.length;
        for (uint256 i = 0; i < amount; ++i) {
            _slashAgent(_domain, _accounts[i], _guard);
        }
    }

    /**
     * @notice Child contract should implement the slashing logic for Agents
     * with all the required system calls.
     * @dev Called when agent fraud is proven.
     * @param _domain   Domain where the reported Agent is active
     * @param _account  Address of the fraudulent Agent
     * @param _guard    Guard who reported the fraudulent Agent [address(0) if not a Guard report]
     */
    function _slashAgent(
        uint32 _domain,
        address _account,
        address _guard
    ) internal virtual;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isIgnoredAgent(uint32 _domain, address)
        internal
        view
        virtual
        override
        returns (bool)
    {
        // Origin only keeps track of remote Notaries
        return _domain == _localDomain();
    }

    /**
     * @notice Returns whether (_destination, _nonce, _root) matches the historical state
     * of the Merkle Tree for that destination.
     * @dev For `_nonce == 0`: root has to match `EMPTY_TREE_ROOT` (root of an empty merkle tree)
     * For `_nonce != 0`:
     * - There has to be at least `_nonce` messages sent to `_destination`
     * - Merkle root after sending message with `nonce == _nonce` should match `_root`
     */
    function _isValidAttestation(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce,
        bytes32 _root,
        RootMetadata memory metadata
    ) internal view returns (bool) {
        // Attestation with origin domain not matching local domain should be discarded
        require(_origin == _localDomain(), "!attestationOrigin: !local");
        if (_nonce < historicalRoots[_destination].length) {
            RootMetadata memory histMetadata = historicalMetadata[_destination][_nonce];
            // If a nonce exists for a given destination,
            // a root should match the historical root,
            // and the metadata should match the historical one.
            return
                _root == historicalRoots[_destination][_nonce] &&
                _isSameMetadata(histMetadata, metadata);
        }
        // If a nonce doesn't exist for a given destination,
        // it should be a zero nonce with a root of an empty merkle tree
        return _nonce == 0 && _root == EMPTY_TREE_ROOT;
    }

    /**
     * @notice Returns amount of leaves in the merkle tree for the given destination.
     * @dev Every inserted leaf leads to adding a historical root,
     * removing the necessity to store amount of leaves separately.
     * Historical roots array is initialized with a root of an empty Merkle tree,
     * thus actual amount of leaves is lower by one.
     */
    function _getTreeCount(uint32 _destination) internal view returns (uint256) {
        // if no historical roots are saved, destination is unknown, and there were
        // no dispatched messages to that destination
        if (historicalRoots[_destination].length == 0) return 0;
        // We subtract 1, as the very first inserted root is EMPTY_TREE_ROOT
        return historicalRoots[_destination].length - 1;
    }

    /**
     * @notice Returns whether the given metadata structs have the same data.
     */
    function _isSameMetadata(RootMetadata memory a, RootMetadata memory b)
        internal
        pure
        returns (bool)
    {
        return (a.blockNumber == b.blockNumber && a.timestamp == b.timestamp);
    }
}
