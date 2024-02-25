// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainApp {
    function setAppConfig(
        uint64[] memory chainIDs,
        address[] memory linkedIApps,
        address[] memory sendingModules,
        address[] memory receivingModules,
        uint256 requiredResponses,
        uint64 optimisticTimePeriod
    )
        external;

    function getLinkedIApp(uint64 chainID) external view returns (address);

    function getSendingModules() external view returns (address[] memory);

    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules);

    function send(bytes32 receiver, uint256 dstChainId, bytes calldata message) external payable;

    function appReceive(uint256 srcChainId, bytes32 sender, uint64 nonce, bytes calldata message) external payable;
}
