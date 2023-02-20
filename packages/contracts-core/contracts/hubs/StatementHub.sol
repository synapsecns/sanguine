// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { Snapshot, SnapshotLib } from "../libs/Snapshot.sol";

import { AgentRegistry } from "../system/AgentRegistry.sol";

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
abstract contract StatementHub is AgentRegistry {
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
     * @dev Internal function to verify the signed snapshot payload.
     * Reverts if either of this is true:
     *  - Snapshot payload is not properly formatted.
     *  - Snapshot signer is not an active Agent.
     * @param _payload      Raw payload with snapshot data
     * @param _signature    Agent signature for the snapshot
     * @return snapshot     Typed memory view over snapshot payload
     * @return domain       Domain where the signed Agent is active
     * @return agent        Agent that signed the snapshot
     */
    function _verifySnapshot(bytes memory _payload, bytes memory _signature)
        internal
        view
        returns (
            Snapshot snapshot,
            uint32 domain,
            address agent
        )
    {
        // This will revert if payload is not a formatted snapshot
        snapshot = _payload.castToSnapshot();
        // This will revert if signer is not an active agent
        (domain, agent) = _recoverAgent(snapshot.hash(), _signature);
        // Guards and Notaries for all domains could sign Snapshots, no further checks are needed.
    }
}
