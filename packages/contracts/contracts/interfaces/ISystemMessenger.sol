// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

interface ISystemMessenger {
    /// @dev Potential senders/recipients of a system message
    enum SystemContracts {
        Origin,
        ReplicaManager
    }

    /**
     * @notice  Send System Message to one of the System Contracts on origin chain
     * @dev     Note that knowledge of recipient address is not required,
     *          routing will be done by SystemMessenger on destination chain.
     * @param _destDomain   Domain of destination chain
     * @param _recipient    System contract type of the recipient
     * @param _payload      Data for calling recipient on destination chain
     */
    function sendSystemMessage(
        uint32 _destDomain,
        SystemContracts _recipient,
        bytes memory _payload
    ) external;
}
