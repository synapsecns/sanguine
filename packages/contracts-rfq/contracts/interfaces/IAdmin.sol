// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IAdmin {
    // ============ Events ============

    event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
    event FeesSwept(address token, address recipient, uint256 amount);

    event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount);

    // ============ Methods ============

    function setProtocolFeeRate(uint256 newFeeRate) external;

    function sweepProtocolFees(address token, address recipient) external;

    function setChainGasAmount(uint256 newChainGasAmount) external;
}
