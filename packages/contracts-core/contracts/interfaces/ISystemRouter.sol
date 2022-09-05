// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

interface ISystemRouter {
    /// @dev Potential senders/recipients of a system message
    enum SystemContracts {
        Origin,
        Destination
    }

    /**
     * @notice  Call a System Contract on the destination chain.
     *          Note: use `localDomain` to call a contract on local chain.
     * @dev     Only System contracts are allowed to call this function.
     *          Note: knowledge of recipient address is not required,
     *          routing will be done by SystemRouter on the destination chain.
     *          Following call will be made on destination chain:
     *              recipient.call(_data, originDomain, originSender}
     *          Note: effectively extending payload with abi encoded (domain, sender)
     *          This allows recipient to check:
     *          - domain where system call originated (local domain in this case)
     *          - System contract type of the sender (msg.sender on local chain)
     * @param _destination  Domain of destination chain
     * @param _recipient    System contract type of the recipient
     * @param _data         Data for calling recipient on destination chain
     */
    function systemCall(
        uint32 _destination,
        SystemContracts _recipient,
        bytes memory _data
    ) external;
}
