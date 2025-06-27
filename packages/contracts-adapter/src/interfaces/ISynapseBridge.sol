// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseBridge {
    function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) external;
    function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) external;
}
