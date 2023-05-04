// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface InterfaceGasOracle {
    /**
     * @notice Fetches the latest gas data for the chain from `Destination` contract,
     * and uses it to update the oracle values for the requested chain.
     * @param domain    Domain to update the gas data for
     */
    function updateGasData(uint32 domain) external;

    /**
     * @notice Returns the gas data for the local chain.
     */
    function getGasData() external view returns (uint256 paddedGasData);

    /**
     * @notice Returns the gas data for the given domain, in the decoded format.
     * @param domain        Domain of chain to get gas data for
     * @return gasPrice     Gas price for the chain (in Wei per gas unit)
     * @return dataPrice    Calldata price (in Wei per byte of content)
     * @return execBuffer   Tx fee safety buffer for message execution (in Wei)
     * @return amortAttCost Amortized cost for attestation submission (in Wei)
     * @return etherPrice   Ratio of Chain's Ether Price / Mainnet Ether Price (in BWAD)
     * @return markup       Markup for the message execution (in BWAD)
     */
    function getDecodedGasData(uint32 domain)
        external
        view
        returns (
            uint256 gasPrice,
            uint256 dataPrice,
            uint256 execBuffer,
            uint256 amortAttCost,
            uint256 etherPrice,
            uint256 markup
        );

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
