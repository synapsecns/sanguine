// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainClientV1 {
    /**
     * @notice Sets the linked client for a specific chain ID.
     * @dev Stores the address of the linked client in a mapping with the chain ID as the key.
     * @param chainId The chain ID for which the client is being set.
     * @param client The address of the client being linked.
     */
    function setLinkedClient(uint256 chainId, bytes32 client) external;

    /**
     * @notice Sets the address of the InterchainDB contract.
     * @dev Only callable by the contract owner or an authorized account.
     * @param _interchainDB The address of the InterchainDB contract.
     */
    function setInterchainDB(address _interchainDB) external;

    /**
     * @notice Sends a message to another chain via the Interchain Communication Protocol.
     * @dev Charges a fee for the message, which is payable upon calling this function.
     * @param receiver The address of the receiver on the destination chain.
     * @param dstChainId The chain ID of the destination chain.
     * @param message The message being sent.
     * @param srcModules The source modules involved in the message sending.
     */
    function interchainSend(
        bytes32 receiver,
        uint256 dstChainId,
        bytes calldata message,
        address[] calldata srcModules
    ) external payable;

    /**
     * @notice Executes a transaction that has been sent via the Interchain.
     * @dev The transaction must have been previously sent and recorded.
     * @param transactionID The ID of the transaction being executed.
     * @param transaction The transaction data.
     */
    function interchainExecute(bytes32 transactionID, bytes calldata transaction) external;

    /**
     * @notice Converts a bytes32 value to an address.
     * @dev Useful for converting blockchain-specific identifiers to Ethereum addresses.
     * @param _bytes32 The bytes32 value to convert.
     * @return address The address obtained from the bytes32 value.
     */
    function convertBytes32ToAddress(bytes32 _bytes32) external pure returns (address);

    /**
     * @notice Converts an address to a bytes32 value.
     * @dev Useful for converting Ethereum addresses to blockchain-specific identifiers.
     * @param _address The address to convert.
     * @return bytes32 The bytes32 representation of the address.
     */
    function convertAddressToBytes32(address _address) external pure returns (bytes32);
}
