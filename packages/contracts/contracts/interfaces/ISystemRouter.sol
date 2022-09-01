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
     * @dev     Note that knowledge of recipient address is not required,
     *          routing will be done by SystemRouter on destination chain.
     * @param _destination  Domain of destination chain
     * @param _recipient    System contract type of the recipient
     * @param _payload      Data for calling recipient on destination chain
     */
    function sendSystemMessage(
        uint32 _destination,
        SystemContracts _recipient,
        bytes memory _payload
    ) external;

    /**
     * @notice  Call a System Contract on the local chain.
     * @dev     Only System contracts are allowed to call this function.
     *          Note that knowledge of recipient address is not required,
     *          routing will be done by SystemRouter on the local chain.
     * @param _recipient    System contract type of the recipient
     * @param _payload      Data for calling recipient on destination chain
     */
    function systemCall(SystemContracts _recipient, bytes memory _payload) external;
}
