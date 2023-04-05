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
     * @param proofMaturity     Message's merkle proof age in seconds
     * @param body              Body of the system message
     */
    function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 proofMaturity, bytes memory body) external;

    /**
     * @notice Call a System Contract on the remote chain with a given calldata.
     * This is done by sending a system message to the System Router on the destination chain.
     * Note: knowledge of recipient address is not required, routing will be done by the System Router.
     * @dev Only System contracts are allowed to call this function.
     * System Entities should expose functions for cross-chain system calls using this template:
     *  - `function foo(uint256 proofMaturity, uint32 origin, SystemEntity sender, *args)`
     *  - `(proofMaturity, origin, sender)` are later referenced as "security arguments" filled by SystemRouter
     *  - `*args` is used to denote the non-security function arguments (that could be of any type).
     * Note: such function should be protected with onlySystemRouter modifier
     * @dev Assuming `payload = abi.encodeWithSelector(foo.selector, *args)`,
     * following call will be made on destination chain:
     *  - `recipient.foo(proofMaturity, origin, sender, *args)`
     * This allows recipient to check:
     * - `uint256 proofMaturity`: system message's merkle proof age in seconds
     * - `uint32 origin`: domain where a system call originated
     * - `SystemEntity `sender`: system entity who initiated the call on origin chain
     * @param destination           Domain of destination chain
     * @param optimisticPeriod      Optimistic period for the message
     * @param recipient             System entity to be called on destination chain
     * @param payload               Calldata payload without security arguments
     */
    function systemCall(uint32 destination, uint32 optimisticPeriod, SystemEntity recipient, bytes memory payload)
        external;
}
