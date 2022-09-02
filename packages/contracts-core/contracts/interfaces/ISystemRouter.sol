// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

interface ISystemRouter {
    /// @dev Potential senders/recipients of a system message
    enum SystemContracts {
        Origin,
        Destination
    }

    /**
     * @notice  Send System Message to one of the System Contracts on origin chain
     * @dev     Note: knowledge of recipient address is not required,
     *          routing will be done by SystemRouter on destination chain.
     *          Following call will be made on destination chain:
     *              recipient.func(originDomain, originSender, _data)
     *          Allowing recipient to check:
     *          - domain where remote system call originated
     *          - System contract type of the sender
     * @param _destination  Domain of destination chain
     * @param _recipient    System contract type of the recipient
     * @param _selector     Function to call on destination chain
     * @param _data         Data for calling recipient on destination chain
     */
    function remoteSystemCall(
        uint32 _destination,
        SystemContracts _recipient,
        bytes4 _selector,
        bytes memory _data
    ) external;

    /**
     * @notice  Call a System Contract on the local chain.
     * @dev     Only System contracts are allowed to call this function.
     *          Note: knowledge of recipient address is not required,
     *          routing will be done by SystemRouter on the local chain.
     *          Following call will be made on local chain:
     *              recipient.func(localDomain, localSender, _data)
     *          Allowing recipient to check:
     *          - domain where system call originated (local domain in this case)
     *          - System contract type of the sender
     * @param _recipient    System contract type of the recipient
     * @param _selector     Function to call on destination chain
     * @param _data         Data for calling recipient on destination chain
     */
    function localSystemCall(
        SystemContracts _recipient,
        bytes4 _selector,
        bytes memory _data
    ) external;
}
