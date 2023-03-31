// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SystemEntity} from "../libs/Structures.sol";

interface InterfaceSystemRouter {
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
     * @param optimisticSeconds     Optimistic period for the message
     * @param recipient             System entity to receive the call on destination chain
     * @param payload               Calldata payload for calling recipient on destination chain
     */
    function systemCall(uint32 destination, uint32 optimisticSeconds, SystemEntity recipient, bytes memory payload)
        external;

    /**
     * @notice Calls a few system contracts using the given calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 destination,
        uint32 optimisticSeconds,
        SystemEntity[] memory recipients,
        bytes[] memory payloadArray
    ) external;

    /**
     * @notice Calls a few system contracts using the same calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 destination,
        uint32 optimisticSeconds,
        SystemEntity[] memory recipients,
        bytes memory payload
    ) external;

    /**
     * @notice Calls a single system contract a few times using the given calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 destination,
        uint32 optimisticSeconds,
        SystemEntity recipient,
        bytes[] memory payloadArray
    ) external;
}
