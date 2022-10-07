// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Attestation } from "../libs/Attestation.sol";
import { Report } from "../libs/Report.sol";
import { TypedMemView } from "../libs/TypedMemView.sol";

import { DomainNotaryRegistry } from "../registry/DomainNotaryRegistry.sol";
import { GuardRegistry } from "../registry/GuardRegistry.sol";
import { OriginHubEvents } from "../events/OriginHubEvents.sol";
import { AttestationHub } from "./AttestationHub.sol";
import { ReportHub } from "./ReportHub.sol";

import { MerkleLib } from "../libs/Merkle.sol";

/**
 * @notice Inserts new message hashes into the Merkle Tree,
 * and keeps track of the historical merkle state.
 * Keeps track of this domain's Notaries and all Guards: accepts
 * and checks their attestations/reports related to Origin.
 */
abstract contract OriginHub is
    OriginHubEvents,
    AttestationHub,
    ReportHub,
    DomainNotaryRegistry,
    GuardRegistry
{
    using Attestation for bytes29;
    using Report for bytes29;
    using TypedMemView for bytes29;

    using MerkleLib for MerkleLib.Tree;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Merkle Tree containing all hashes of sent messages
    MerkleLib.Tree internal tree;
    // Merkle tree roots after inserting a sent message
    bytes32[] public historicalRoots;

    // gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Suggest an attestation for the off-chain actors to sign and submit.
     * @dev If no messages have been sent, following values are returned:
     * - nonce = 0
     * - root = 0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757
     * Which is the merkle root for an empty sparse merkle tree.
     * @return latestNonce Current nonce
     * @return latestRoot  Current merkle root
     */
    function suggestAttestation() external view returns (uint32 latestNonce, bytes32 latestRoot) {
        latestNonce = nonce();
        latestRoot = historicalRoots[latestNonce];
    }

    /**
     * @notice Returns nonce of the last inserted Merkle root, which is also
     * the number of inserted leaves in the tree (current index).
     */
    function nonce() public view returns (uint32 latestNonce) {
        latestNonce = uint32(_getTreeCount());
    }

    /**
     * @notice Calculates and returns tree's current root.
     */
    function root() public view returns (bytes32) {
        return tree.root(_getTreeCount());
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Checks if a submitted Attestation is a valid Attestation.
     * Attestation Flag can be either "Fraud" or "Valid".
     * A "Fraud" Attestation is a (_nonce, _root) attestation that doesn't correspond with
     * the historical state of Origin contract. Either of these needs to be true:
     * - _nonce is higher than current nonce (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce
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
        uint32 attestedNonce = _attestationView.attestedNonce();
        bytes32 attestedRoot = _attestationView.attestedRoot();
        isValid = _isValidAttestation(attestedNonce, attestedRoot);
        if (!isValid) {
            emit FraudAttestation(_notary, _attestation);
            // Guard doesn't receive anything, as Notary wasn't slashed using the Fraud Report
            _slashNotary(_notary, address(0));
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
     * A Fraud Attestation is a (_nonce, _root) attestation that doesn't correspond with
     * the historical state of Origin contract. Either of those needs to be true:
     * - _nonce is higher than current nonce (no root exists for this nonce)
     * - _root is not equal to the historical root of _nonce
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
        uint32 attestedNonce = _attestationView.attestedNonce();
        bytes32 attestedRoot = _attestationView.attestedRoot();
        if (_isValidAttestation(attestedNonce, attestedRoot)) {
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
                _slashNotary(_notary, _guard);
                return true;
            } else {
                // Flag: Valid
                // Report is incorrect, slash the Guard
                emit IncorrectReport(_guard, _report);
                _slashGuard(_guard);
                emit FraudAttestation(_notary, _attestationView.clone());
                // Guard doesn't receive anything due to Valid flag on the Report
                _slashNotary(_notary, address(0));
                return false;
            }
        }
    }

    /**
     * @notice Inserts a merkle root for an empty sparse merkle tree
     * into the historical roots array.
     * @dev This enables:
     * - Counting nonces from 1 (nonce=0 meaning no messages have been sent).
     * - Not slashing the Notaries for signing an attestation for an empty tree
     * (assuming they sign the correct root outlined below).
     */
    function _initializeHistoricalRoots() internal {
        // This function should only be called only if array is empty
        assert(historicalRoots.length == 0);
        // Insert a historical root so nonces start at 1 rather then 0.
        // Here we insert the default root of a sparse merkle tree
        historicalRoots.push(hex"27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757");
    }

    /**
     * @notice Inserts new message into the Merkle tree and stores the new merkle root.
     * @param _messageNonce Nonce of the dispatched message
     * @param _messageHash  Hash of the dispatched message
     */
    function _insertMessage(uint32 _messageNonce, bytes32 _messageHash) internal {
        /// @dev _messageNonce == tree.count() + 1
        // tree.insert() requires amount of leaves AFTER the leaf insertion (i.e. tree.count() + 1)
        tree.insert(_messageNonce, _messageHash);
        /// @dev leaf is inserted => _messageNonce == tree.count()
        // tree.root() requires current amount of leaves (i.e. tree.count())
        historicalRoots.push(tree.root(_messageNonce));
    }

    /**
     * @notice Child contract should implement the slashing logic for Notaries
     * with all the required system calls.
     * @dev Called when fraud is proven (Fraud Attestation).
     * @param _notary   Notary to slash
     * @param _guard    Guard who reported fraudulent Notary [address(0) if not a Guard report]
     */
    function _slashNotary(address _notary, address _guard) internal virtual;

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

    /**
     * @notice Returns whether (_nonce, _root) matches the historical state
     * of the Merkle Tree.
     */
    function _isValidAttestation(uint32 _nonce, bytes32 _root) internal view returns (bool) {
        // Check if nonce is valid, if not => attestation is fraud
        // Check if root the same as the historical one, if not => attestation is fraud
        return (_nonce < historicalRoots.length && _root == historicalRoots[_nonce]);
    }

    /**
     * @notice Returns amount of leaves in the merkle tree.
     * @dev Every inserted leaf leads to adding a historical root,
     * removing the necessity to store amount of leaves separately.
     * Historical roots array is initialized with a root of an empty Sparse Merkle tree,
     * thus actual amount of leaves is lower by one.
     */
    function _getTreeCount() internal view returns (uint256) {
        // historicalRoots has length of 1 upon initializing,
        // so this never underflows assuming contract was initialized
        return historicalRoots.length - 1;
    }
}
