// SPDX-License-Identifier: MIT
pragma solidity =0.8.20 ^0.8.0;

// contracts/interfaces/IInterchainApp.sol

/// @notice Minimal interface for the Interchain App to work with the Interchain Client.
interface IInterchainApp {
    function appReceive(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        external
        payable;

    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules);
}

// test/mocks/InterchainAppMock.sol

// solhint-disable no-empty-blocks
contract InterchainAppMock is IInterchainApp {
    address[] public receivingModules;

    function setReceivingModule(address _receivingModule) external {
        receivingModules.push(_receivingModule);
    }

    function appReceive(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        external
        payable
    {}

    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules) {}
}
