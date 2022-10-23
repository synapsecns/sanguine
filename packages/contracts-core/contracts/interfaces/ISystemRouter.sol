// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface ISystemRouter {
    /// @dev Potential senders/recipients of a system message
    enum SystemEntity {
        Origin,
        Destination
    }

    /**
     * @notice Call a System Contract on the destination chain with a given data payload.
     * Note: for system calls on the local chain
     * - use `destination = localDomain`
     * - `_optimisticSeconds` value will be ignored
     *
     * @dev Only System contracts are allowed to call this function.
     * Note: knowledge of recipient address is not required, routing will be done by SystemRouter
     * on the destination chain. Following call will be made on destination chain:
     * - recipient.call(_data, originDomain, originSender, rootSubmittedAt)
     * Note: data payload is extended with abi encoded (domain, sender, rootTimestamp)
     * This allows recipient to check:
     * - domain where a system call originated (local domain in this case)
     * - system entity, who initiated the call (msg.sender on local chain)
     * - timestamp when merkle root was submitted and optimistic timer started ticking
     * @param _destination          Domain of destination chain
     * @param _optimisticSeconds    Optimistic period for the message
     * @param _recipient            System entity to receive the call on destination chain
     * @param _data                 Data for calling recipient on destination chain
     */
    function systemCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity _recipient,
        bytes memory _data
    ) external;

    /**
     * @notice Calls a few system contracts with the given calldata.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes[] memory _dataArray
    ) external;
}
