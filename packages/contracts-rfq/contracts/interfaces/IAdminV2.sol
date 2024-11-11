// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface IAdminV2 {
    event CancelDelayUpdated(uint256 oldCancelDelay, uint256 newCancelDelay);
    event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
    event FeesSwept(address token, address recipient, uint256 amount);

    event ChainGasAmountUpdated(uint256 oldChainGasAmount, uint256 newChainGasAmount);

    function setCancelDelay(uint256 newCancelDelay) external;

    function setProtocolFeeRate(uint256 newFeeRate) external;

    function sweepProtocolFees(address token, address recipient) external;

    function setChainGasAmount(uint256 newChainGasAmount) external;
}
