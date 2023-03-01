// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { AgentRegistry } from "../system/AgentRegistry.sol";
import "../Version.sol";
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "../libs/Attestation.sol";
import "../libs/Snapshot.sol";
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
abstract contract StatementHub is AgentRegistry, Version0_0_2 {
    using AttestationLib for bytes;
    using SnapshotLib for bytes;

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
     * Reverts if either of this is true:
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
        // This will revert if signer is not an active agent
        (domain, notary) = _recoverAgent(attestation.hash(), _attSignature);
        // Attestation signer needs to be a Notary, not a Guard
        require(domain != 0, "Signer is not a Notary");
    }

    /**
     * @dev Internal function to verify the signed snapshot payload.
     * Reverts if either of this is true:
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
        (domain, agent) = _recoverAgent(snapshot.hash(), _snapSignature);
        // Guards and Notaries for all domains could sign Snapshots, no further checks are needed.
    }

    /**
     * @dev Internal function to verify that snapshot root matches the root from Attestation.
     * Reverts if either of this is true:
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
}
