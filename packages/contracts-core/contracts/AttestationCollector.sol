// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./Version.sol";
import { Attestation } from "./libs/Attestation.sol";
import { AttestationHub } from "./hubs/AttestationHub.sol";
import { AttestationCollectorEvents } from "./events/AttestationCollectorEvents.sol";

import { ByteString } from "./libs/Attestation.sol";

import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract AttestationCollector is
    AttestationCollectorEvents,
    AttestationHub,
    OwnableUpgradeable,
    Version0_0_1
{
    using Attestation for bytes29;
    using ByteString for bytes;
    using ByteString for bytes29;

    /**
     * @notice Contains a merkle root and existing agent signatures for this root.
     * @dev We are storing indexes for agent signatures. They are stored as
     * "position of signature in `savedSignatures` array plus 1". The default value of 0
     * means there is no agent signature for this root.
     * Enforced invariant: for every saved root, al lest one of the indexes is non-zero,
     * i.e. any root is saved with at least one agent (Guard/Notary) signature.
     * Note: "index plus 1" is abstracted away from the off-chain agents.
     * Events and getters containing `index` variable refer to conventional: 0 <= index < length
     * @param root              Merkle root for some given `(origin, destination, nonce)`
     * @param guardSigIndex     Guard signature's index in `savedSignatures` plus 1
     * @param notarySigIndex    Notary signature's index in `savedSignatures` plus 1
     */
    struct SignedRoot {
        bytes32 root;
        uint128 guardSigIndex;
        uint128 notarySigIndex;
    }

    /**
     * @notice Contains an agent signature and attestation key if refers to.
     * @dev We're storing saved merkle roots separately using {SignedRoot} struct.
     * At the moment, no conflicting roots are saved.
     * @param r             R-value of the signature payload
     * @param s             S-value of the signature payload
     * @param v             V-value of the signature payload
     * @param isGuard       Whether the signer is Guard or Notary
     * @param origin        Attestation origin domain
     * @param destination   Attestation destination domain
     * @param nonce         Attestation nonce
     * @param blockNumber   Attestation block number
     * @param timestamp     Attestation block timestamp
     */
    struct AgentSignature {
        bytes32 r;
        bytes32 s;
        uint8 v;
        bool isGuard;
        uint32 origin;
        uint32 destination;
        uint32 nonce;
        uint40 blockNumber;
        uint40 timestamp;
        // 64 bits available
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Attested root for every (origin, destination, nonce) tuple.
     * @dev At the moment, we only save one merkle root per (origin, destination, nonce) tuple.
     * Every conflicting root is discarded.
     *
     * attKey = (origin, destination, nonce)
     * signedRoots: attKey => (root with signatures)
     */
    mapping(uint96 => SignedRoot) internal signedRoots;

    /**
     * @notice All stored agent signatures.
     * @dev We save an signature only if the latest saved signature for that agent
     * precedes the new one, i.e. has a lower nonce.
     */
    AgentSignature[] internal savedSignatures;

    /**
     * @notice A list of signature indexes for every (origin, destination, agent) tuple.
     * @dev signatureIndex is the position of agent signature in `savedSignatures` list plus 1.
     * The default value of 0 indicates that signature is not in `savedSignatures`.
     * Invariant: signature indexes in `agentSigIndexes` are non-zero (refer to saved signature).
     *
     * attDomains = (origin, destination)
     * agentSigIndexes: attDomains => (agent => [signature indexes])
     */
    mapping(uint64 => mapping(address => uint256[])) internal agentSigIndexes;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __Ownable_init_unchained();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        ADDING AGENTS (MOCKS)                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    // TODO (Chi): add/remove agents via system calls from local BondingManager

    function addAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _addAgent(_domain, _account);
    }

    function removeAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _removeAgent(_domain, _account);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Get the amount of (origin, destination) attestations saved for a given agent.
     */
    function agentAttestations(
        uint32 _origin,
        uint32 _destination,
        address _agent
    ) external view returns (uint256) {
        uint64 attDomains = Attestation.attestationDomains(_origin, _destination);
        return agentSigIndexes[attDomains][_agent].length;
    }

    /**
     * @notice Get the total amount of saved attestations.
     */
    function savedAttestations() external view returns (uint256) {
        return savedSignatures.length;
    }

    /**
     * @notice Get i-th (origin, destination) Attestation for a given agent.
     * Will always contain exactly one agent signature.
     */
    function getAgentAttestation(
        uint32 _origin,
        uint32 _destination,
        address _agent,
        uint256 _index
    ) external view returns (bytes memory) {
        uint64 attDomains = Attestation.attestationDomains(_origin, _destination);
        require(_index < agentSigIndexes[attDomains][_agent].length, "Out of range");
        uint256 signatureIndex = agentSigIndexes[attDomains][_agent][_index];
        return _formatAgentAttestation(signatureIndex);
    }

    /**
     * @notice Get the latest known nonce for (origin, destination) signed by the given agent.
     * @dev Will return 0, if an agent hasn't submitted a single attestation yet.
     */
    function getLatestNonce(
        uint32 _origin,
        uint32 _destination,
        address _agent
    ) external view returns (uint32) {
        uint64 attestationDomains = Attestation.attestationDomains(_origin, _destination);
        uint32 latestNonce = _latestAgentNonce(attestationDomains, _agent);
        return latestNonce;
    }

    /**
     * @notice Get latest attestation for (origin, destination) signed by given agent.
     */
    function getLatestAttestation(
        uint32 _origin,
        uint32 _destination,
        address _agent
    ) external view returns (bytes memory) {
        uint64 attDomains = Attestation.attestationDomains(_origin, _destination);
        uint256 amount = agentSigIndexes[attDomains][_agent].length;
        require(amount != 0, "No attestations found");
        uint256 signatureIndex = agentSigIndexes[attDomains][_agent][amount - 1];
        return _formatAgentAttestation(signatureIndex);
    }

    /**
     * @notice Get Attestation for (origin, destination, nonce), if it was previously saved.
     * Will contain at least one agent signature.
     * Will contain a single guard signature, if it was previously saved.
     * Will contain a single notary signature, if it was previously saved.
     */
    function getAttestation(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce
    ) external view returns (bytes memory) {
        uint96 attKey = Attestation.attestationKey(_origin, _destination, _nonce);
        SignedRoot memory signedRoot = signedRoots[attKey];
        require(signedRoot.root != bytes32(0), "Unknown nonce");
        // Find an existing signature for the attestation
        uint256 agentSigIndex = signedRoot.guardSigIndex != 0
            ? signedRoot.guardSigIndex
            : signedRoot.notarySigIndex;
        // Get block number and timestamp from the signature
        // TODO: deal with potential conflicts here
        (uint40 blockNumber, uint40 timestamp) = (
            savedSignatures[agentSigIndex - 1].blockNumber,
            savedSignatures[agentSigIndex - 1].timestamp
        );
        bytes memory attData = Attestation.formatAttestationData({
            _origin: _origin,
            _destination: _destination,
            _nonce: _nonce,
            _root: signedRoot.root,
            _blockNumber: blockNumber,
            _timestamp: timestamp
        });
        return
            _formatDualAttestation({
                _attestationData: attData,
                _guardSignatureIndex: signedRoot.guardSigIndex,
                _notarySignatureIndex: signedRoot.notarySigIndex
            });
    }

    /**
     * @notice Get merkle root for (origin, destination, nonce), if it was previously saved.
     */
    function getRoot(
        uint32 _origin,
        uint32 _destination,
        uint32 _nonce
    ) external view returns (bytes32) {
        uint96 attKey = Attestation.attestationKey(_origin, _destination, _nonce);
        return signedRoots[attKey].root;
    }

    /**
     * @notice Get i-th saved Attestation from the global list of "all saved agents attestations"
     * Will always contain exactly one agent signature.
     */
    function getSavedAttestation(uint256 _index) external view returns (bytes memory) {
        require(_index < savedSignatures.length, "Out of range");
        return _formatAgentAttestation({ _signatureIndex: _index + 1 });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Saves the attestation data, if it doesn't contradict the existing data.
     * Saves all agent signatures that are not outdated.
     * @dev Guards and Notaries signatures and roles have been checked in AttestationHub.
     *
     * @param _guards           Guard addresses (signatures&roles already verified)
     * @param _notaries         Notary addresses (signatures&roles already verified)
     * @param _attestationView  Memory view over the Attestation for convenience
     * @return stored   TRUE if Attestation was stored.
     */
    function _handleAttestation(
        address[] memory _guards,
        address[] memory _notaries,
        bytes29 _attestationView,
        bytes memory _attestation
    ) internal override returns (bool stored) {
        uint96 attKey = _attestationView.attestedKey();
        bytes32 root = _attestationView.attestedRoot();
        // TODO (Chi): to enforce "non-zero saved root" invariant by
        // checking if root is non-zero in Attestation.isAttestation()
        require(root != bytes32(0), "Root is zero");
        // Check what we have saved for this `attKey` previously
        SignedRoot memory existingRoot = signedRoots[attKey];
        if (existingRoot.root == bytes32(0)) {
            // Case 1: no root was saved for `attKey`.
            // Meaning no sig indexes were saved as well.
            // Attestation has at least one signature (enforced in Attestation.isAttestation).
            existingRoot.root = root;
        } else if (existingRoot.root != root) {
            // Case 2: another root was saved for `attKey`.
            // At the moment we don't do anything here
            // TODO (Chi): actually do something
            return false;
        }
        // Case 3: the same root was saved for `attKey`.
        // Need to go through attestation signatures and store the ones we don't have.
        // Track if at least one new signature was linked to the attested root
        // Save all new guard signatures
        bool linked;
        (stored, linked) = _handleSignatures({
            _attestationView: _attestationView,
            _existingRoot: existingRoot,
            _isGuard: true,
            _agents: _guards
        });
        // Save all new notary signatures
        (bool _stored, bool _linked) = _handleSignatures({
            _attestationView: _attestationView,
            _existingRoot: existingRoot,
            _isGuard: false,
            _agents: _notaries
        });
        // Check if at least one agent signature was stored / linked
        stored = stored || _stored;
        linked = linked || _linked;
        // Emit event only if at least one signature was stored
        if (stored) {
            emit AttestationAccepted(_guards, _notaries, _attestation);
        }
        // Update storage records if at least one signature was linked
        if (linked) {
            signedRoots[attKey] = existingRoot;
        }
    }

    /**
     * @notice Saves not-outdated signatures for either all guard or notary signers
     * form the attestation.
     * Signature is considered outdated, if the same signer has already submitted
     * an attestation with an equal or bigger nonce.
     */
    function _handleSignatures(
        bytes29 _attestationView,
        SignedRoot memory _existingRoot,
        bool _isGuard,
        address[] memory _agents
    ) internal returns (bool signatureStored, bool signatureLinked) {
        uint256 amount = _agents.length;
        for (uint256 i = 0; i < amount; ++i) {
            uint256 savedSigIndex = _insertAttestation({
                _attestationView: _attestationView,
                _agentIndex: i,
                _isGuard: _isGuard,
                _agent: _agents[i]
            });
            // Check if the signature was saved
            if (savedSigIndex != 0) {
                signatureStored = true;
                // TODO (Chi): link every saved signature to have fallback signatures
                if (_isGuard && _existingRoot.guardSigIndex == 0) {
                    // Link a guard signature only if no guard signatures have been linked before
                    _existingRoot.guardSigIndex = uint128(savedSigIndex);
                    signatureLinked = true;
                } else if (!_isGuard && _existingRoot.notarySigIndex == 0) {
                    // Link a notary signature only if no notary signatures have been linked before
                    _existingRoot.notarySigIndex = uint128(savedSigIndex);
                    signatureLinked = true;
                }
            }
        }
    }

    /**
     * @notice Saves signature of a given attestation signer, if it is not outdated.
     * Signature is considered outdated, if the same signer has already submitted
     * an attestation with an equal or bigger nonce.
     */
    function _insertAttestation(
        bytes29 _attestationView,
        uint256 _agentIndex,
        bool _isGuard,
        address _agent
    ) internal returns (uint256 signatureIndex) {
        uint64 attDomains = _attestationView.attestedDomains();
        uint32 nonce = _attestationView.attestedNonce();
        // Don't store outdated agent attestation
        if (nonce <= _latestAgentNonce(attDomains, _agent)) return 0;
        // Get the memory view over the agent's signature
        bytes29 signature = (
            _isGuard
                ? _attestationView.guardSignature(_agentIndex)
                : _attestationView.notarySignature(_agentIndex)
        );
        // Second agent signature will be left empty
        bytes29 emptySig = bytes("").castToSignature();
        // Construct the signature struct to save
        AgentSignature memory agentSig;
        (agentSig.r, agentSig.s, agentSig.v) = signature.toRSV();
        agentSig.isGuard = _isGuard;
        (agentSig.origin, agentSig.destination) = Attestation.unpackDomains(attDomains);
        agentSig.nonce = nonce;
        agentSig.blockNumber = _attestationView.attestedBlockNumber();
        agentSig.timestamp = _attestationView.attestedTimestamp();
        savedSignatures.push(agentSig);
        // The signature is stored at length-1, but we add 1 to all indexes
        // and use 0 as a sentinel value
        signatureIndex = uint128(savedSignatures.length);
        agentSigIndexes[attDomains][_agent].push(signatureIndex);
        // Construct attestation with a single signature of a given agent
        // Here we pass views over the existing byte arrays to reduce amount of copying into memory
        bytes memory agentAttestation = Attestation.formatAttestation({
            _dataView: _attestationView.attestationData(),
            _guardSigsView: _isGuard ? signature : emptySig,
            _notarySigsView: _isGuard ? emptySig : signature
        });
        // Use the actual signature position in `savedSignatures` for the event
        emit AttestationSaved(signatureIndex - 1, agentAttestation);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns a previously saved attestation.
     * @dev The default index value of 0 means there is no saved agent signature.
     * This is abstracted away from the off-chain agents.
     * @param _signatureIndex   Signature position in `savedSignatures` plus 1
     */
    function _getSignature(uint256 _signatureIndex) internal view returns (bytes memory signature) {
        if (_signatureIndex != 0) {
            AgentSignature memory agentSig = savedSignatures[_signatureIndex - 1];
            signature = ByteString.formatSignature({ r: agentSig.r, s: agentSig.s, v: agentSig.v });
        }
    }

    /**
     * @notice Forms a "single-agent" attestation using a previously saved signature.
     * @param _signatureIndex   Signature position in `savedSignatures` plus 1
     */
    function _formatAgentAttestation(uint256 _signatureIndex) internal view returns (bytes memory) {
        // Invariant: we always save "index in the array plus 1" as `_signatureIndex`
        assert(_signatureIndex != 0);
        // Read saved agent signature
        AgentSignature memory agentSig = savedSignatures[_signatureIndex - 1];
        uint96 attKey = Attestation.attestationKey({
            _origin: agentSig.origin,
            _destination: agentSig.destination,
            _nonce: agentSig.nonce
        });
        bytes32 root = signedRoots[attKey].root;
        // Invariant: Every saved signature refers to saved root
        assert(root != bytes32(0));
        // Reconstruct attestation data
        bytes memory attData = Attestation.formatAttestationData({
            _origin: agentSig.origin,
            _destination: agentSig.destination,
            _nonce: agentSig.nonce,
            _root: root,
            _blockNumber: agentSig.blockNumber,
            _timestamp: agentSig.timestamp
        });
        // Reconstruct agent signature on `attData`
        bytes memory signature = ByteString.formatSignature({
            r: agentSig.r,
            s: agentSig.s,
            v: agentSig.v
        });
        // Format attestation using a single signature
        return
            Attestation.formatAttestation({
                _data: attData,
                _guardSignatures: agentSig.isGuard ? signature : bytes(""),
                _notarySignatures: agentSig.isGuard ? bytes("") : signature
            });
    }

    /**
     * @notice Forms an attestation with one guard signature (if present)
     * and one notary signature (if present).
     */
    function _formatDualAttestation(
        bytes memory _attestationData,
        uint256 _guardSignatureIndex,
        uint256 _notarySignatureIndex
    ) internal view returns (bytes memory) {
        return
            Attestation.formatAttestation({
                _data: _attestationData,
                _guardSignatures: _getSignature(_guardSignatureIndex),
                _notarySignatures: _getSignature(_notarySignatureIndex)
            });
    }

    /**
     * @notice Returns the latest known nonce that was used in an attestation
     * by a given agent.
     * @dev Will return 0, if an agent hasn't submitted a single attestation yet.
     */
    function _latestAgentNonce(uint64 _attDomains, address _agent)
        internal
        view
        returns (uint32 nonce)
    {
        uint256 length = agentSigIndexes[_attDomains][_agent].length;
        if (length > 0) {
            uint256 sigIndex = agentSigIndexes[_attDomains][_agent][length - 1];
            nonce = savedSignatures[sigIndex - 1].nonce;
        }
    }

    function _isIgnoredAgent(uint32, address) internal pure override returns (bool) {
        // AttestationCollector doesn't ignore anything
        return false;
    }
}
