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

    function getReceivingModules() external view returns (address[] memory);

    function getRequiredResponses() external view returns (uint256);

    function getOptimisticTimePeriod() external view returns (uint64);

    function send(bytes32 receiver, uint256 dstChainId, bytes calldata message) external payable;

    function appReceive() external;
}
