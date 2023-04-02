// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SystemEntity} from "../libs/Structures.sol";

interface InterfaceSystemRouter {
    /**
     * @notice Message recipient needs to implement this function in order to
     * receive cross-chain messages.
     * @dev Message recipient needs to ensure that merkle proof for the message
     * is at least as old as the optimistic period that the recipient is using.
     * Note: as this point it is checked that the "message optimistic period" has passed,
     * however the period value itself could be anything, and thus could differ from the one
     * that the recipient would like to enforce.
     * @param origin            Domain where message originated
     * @param nonce             Message nonce on the origin domain
     * @param rootSubmittedAt   Time when merkle root (used for proving this message) was submitted
     * @param body              Body of the system message
     */
    function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 rootSubmittedAt, bytes memory body) external;

    /**
     * @notice Call a System Contract on the destination chain with a given calldata.
     * Note: for system calls on the local chain
     * - use `destination = localDomain`
     * - `optimisticSeconds` value will be ignored
     *
     * @dev Only System contracts are allowed to call this function.
     * Note: knowledge of recipient address is not required, routing will be done by SystemRouter
     * on the destination chain. Following call will be made on destination chain:
     * - recipient.call(payload, callOrigin, systemCaller, rootSubmittedAt)
     * This allows recipient to check:
     * - callOrigin: domain where a system call originated (local domain in this case)
     * - systemCaller: system entity who initiated the call (msg.sender on local chain)
     * - rootSubmittedAt:
     *   - For cross-chain calls: timestamp when merkle root (used for executing the system call)
     *     was submitted to destination and its optimistic timer started ticking
     *   - For on-chain calls: timestamp of the current block
     *
     * @param destination           Domain of destination chain
     * @param optimisticPeriod      Optimistic period for the message
     * @param recipient             System entity to receive the call on destination chain
     * @param payload               Calldata payload for calling recipient on destination chain
     */
    function systemCall(uint32 destination, uint32 optimisticPeriod, SystemEntity recipient, bytes memory payload)
        external;
}
