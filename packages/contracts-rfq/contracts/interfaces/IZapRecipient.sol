// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IZapRecipient {
    function zap(address token, uint256 amount, bytes memory zapData) external payable returns (bytes4);
}
