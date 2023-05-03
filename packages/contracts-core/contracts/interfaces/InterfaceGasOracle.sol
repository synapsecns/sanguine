// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceGasOracle {
    /**
     * @notice Returns the gas data for the local chain.
     */
    function getGasData() external view returns (uint256 paddedGasData);

    /**
     * @notice Returns the minimum tips for sending a message to a given destination.
     * @param destination       Domain of destination chain
     * @param paddedRequest     Padded encoded message execution request on destination chain
     * @param contentLength     The length of the message content
     * @return paddedTips       Padded encoded minimum tips information
     */
    function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength)
        external
        view
        returns (uint256 paddedTips);
}
