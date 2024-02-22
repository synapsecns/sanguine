// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IGasOracle } from "./interfaces/IGasOracle.sol";
contract GasOracle is Ownable, IGasOracle {
    // DstChainId => The estimated current gas price in wei of the destination chain
    mapping(uint256 => uint256) public dstGasPriceInWei;
    // DstChainId => USD gas ratio of dstGasToken / srcGasToken
    mapping(uint256 => uint256) public dstGasTokenRatio;

    constructor() {}

    /**
     * @notice Permissioned method to allow an off-chain party to set what each dstChain's
     * gas cost is priced in the srcChain's native gas currency.
     * Example: call on ETH, setCostPerChain(43114, 30000000000, 25180000000000000)
     * chain ID 43114
     * Average of 30 gwei cost to transaction on 43114
     * AVAX/ETH = 0.02518, scaled to gas in wei = 25180000000000000
     * @param _dstChainId The destination chain ID - typically, standard EVM chain ID, but differs on nonEVM chains
     * @param _gasUnitPrice The estimated current gas price in wei of the destination chain
     * @param _gasTokenPriceRatio USD gas ratio of dstGasToken / srcGasToken
     */
    // Example:
    // DstChainId = 1666600000
    // Harmony set gwei to 200000000000
    // ONE / JEWEL = 0.05 == 50000000000000000

    // DstChainId = 53935
    // DFK Chain set 1 gwei = 1000000000
    // JEWEL / ONE = 20000000000000000000

    // DstChainId = 8217
    // Klaytn Gwei set to 250000000000
    // KLAYTN / JEWEL = 1200000000000000000

    // DstchainId = 53935
    // DFK Chain set to 1 gwei
    // JEWEL / KLAYTN = 900000000000000000

    function setCostPerChain(
        uint256 _dstChainId,
        uint256 _gasUnitPrice,
        uint256 _gasTokenPriceRatio
    ) external onlyOwner {
        dstGasPriceInWei[_dstChainId] = _gasUnitPrice;
        dstGasTokenRatio[_dstChainId] = _gasTokenPriceRatio;
    }

    /**
     * @notice Returns srcGasToken fee to charge in wei for the cross-chain message based on the gas limit
     * @param _options Versioned struct used to instruct relayer on how to proceed with gas limits. Contains data on gas limit to submit tx with.
     */
    function estimateGasFee(uint256 _dstChainId, bytes memory _options) external view returns (uint256) {
        uint256 gasLimit;
        // temporary gas limit set
        if (_options.length != 0) {
            (uint16 _txType, uint256 _gasLimit, uint256 _dstAirdrop, bytes32 _dstAddress) = decodeOptions(_options);
            gasLimit = _gasLimit;
        } else {
            gasLimit = 200000;
        }

        uint256 minFee = ((dstGasPriceInWei[_dstChainId] * dstGasTokenRatio[_dstChainId] * gasLimit) / 10**18);

        return minFee;
    }

}
