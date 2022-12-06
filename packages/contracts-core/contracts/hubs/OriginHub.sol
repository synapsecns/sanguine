// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Attestation } from "../libs/Attestation.sol";
import { Report } from "../libs/Report.sol";
import { TypedMemView } from "../libs/TypedMemView.sol";

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
    using Attestation for bytes29;
    using Report for bytes29;
    using TypedMemView for bytes29;

    using MerkleLib for MerkleLib.Tree;

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

    // gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

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
     * @notice Suggest attestation for the off-chain actors to sign for a specific destination.
     * Note: signing the suggested attestation will will never lead
     * to slashing of the actor, assuming they have confirmed that the block, where the merkle root
     * was updated, is not subject to reorganization (which is different for every observed chain).
     * @dev If no messages have been sent, following values are returned:
     * - nonce = 0
     * - root = 0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757
     * Which is the merkle root for an empty merkle tree.
     * @return attestationData Data for the suggested attestation
     */
    function suggestAttestation(uint32 _destination)
        public
        view
        returns (bytes memory attestationData)
    {
        uint32 latestNonce = nonce(_destination);
        return
            Attestation.formatAttestationData({
                _origin: _localDomain(),
                _destination: _destination,
                _nonce: latestNonce,
                _root: getHistoricalRoot(_destination, latestNonce)
            });
    }

    /**
     * @notice Returns a historical merkle root for the given destination.
     * Note: signing the attestation with the given historical root will never lead
     * to slashing of the actor, assuming they have confirmed that the block, where the merkle root
     * was updated, is not subject to reorganization (which is different for every observed chain).
     * @param _destination  Destination domain
     * @param _nonce        Historical nonce
     * @return Root for destination's merkle tree right after message to `_destination`
     * with `nonce = _nonce` was dispatched.
     */
    function getHistoricalRoot(uint32 _destination, uint32 _nonce) public view returns (bytes32) {
        // Check if destination is known
        if (historicalRoots[_destination].length > 0) {
            // Check if nonce exists
            require(_nonce < historicalRoots[_destination].length, "!nonce: existing destination");
            return historicalRoots[_destination][_nonce];
        } else {
            // If destination is unknown, we have the root of an empty merkle tree
            require(_nonce == 0, "!nonce: unknown destination");
            return EMPTY_TREE_ROOT;
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
     * @dev Notary signature and role have been checked in AttestationHub,
     * hence `_notary` is an active Notary at this point.
     *
     * @param _notary           Notary address (signature&role already verified)
     * @param _attestationView  Memory view over reported Attestation for convenience
     * @param _attestation      Payload with Attestation data and signature
     * @return isValid          TRUE if Attestation was valid (implying Notary was not slashed).
     */
    function _handleAttestation(
        address _notary,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal override returns (bool isValid) {
        uint32 _origin = _attestationView.attestedOrigin();
        uint32 _destination = _attestationView.attestedDestination();
        uint32 _nonce = _attestationView.attestedNonce();
        bytes32 _root = _attestationView.attestedRoot();
        isValid = _isValidAttestation(_origin, _destination, _nonce, _root);
        if (!isValid) {
            emit FraudAttestation(_notary, _attestation);
            // Guard doesn't receive anything, as Notary wasn't slashed using the Fraud Report
            _slashNotary({ _domain: _destination, _notary: _notary, _guard: address(0) });
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
    ) internal {
        // TODO: when Notary is active on Destination, initialize historical roots
        // upon adding a first Notary for given destination
        if (historicalRoots[_destination].length == 0) _initializeHistoricalRoots(_destination);
        /// @dev _messageNonce == tree.count() + 1
        // tree.insert() requires amount of leaves AFTER the leaf insertion (i.e. tree.count() + 1)
        trees[_destination].insert(_messageNonce, _messageHash);
        /// @dev leaf is inserted => _messageNonce == tree.count()
        // tree.root() requires current amount of leaves (i.e. tree.count())
        historicalRoots[_destination].push(trees[_destination].root(_messageNonce));
    }

    /**
     * @notice Child contract should implement the slashing logic for Notaries
     * with all the required system calls.
     * @dev Called when fraud is proven (Fraud Attestation).
     * @param _domain   Domain where the reported Notary is active
     * @param _notary   Notary to slash
     * @param _guard    Guard who reported fraudulent Notary [address(0) if not a Guard report]
     */
    function _slashNotary(
        uint32 _domain,
        address _notary,
        address _guard
    ) internal virtual;

    /**
     * @notice Child contract should implement the slashing logic for Guards
     * with all the required system calls.
     * @dev Called when guard misbehavior is proven (Incorrect Report).
     * @param _guard    Guard to slash
     */
    function _slashGuard(address _guard) internal virtual;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isIgnoredAgent(uint32 _domain, address) internal view override returns (bool) {
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
        bytes32 _root
    ) internal view returns (bool) {
        // Attestation with origin domain not matching local domain should be discarded
        require(_origin == _localDomain(), "!attestationOrigin: !local");
        if (_nonce < historicalRoots[_destination].length) {
            // If a nonce exists for a given destination,
            // a root should match the historical root
            return _root == historicalRoots[_destination][_nonce];
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
}
