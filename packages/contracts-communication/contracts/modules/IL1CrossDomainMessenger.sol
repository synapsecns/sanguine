pragma solidity 0.8.20;

interface IL1CrossDomainMessenger {
    function sendMessage(address target, bytes calldata message, uint32 gasLimit) external payable;

    function xDomainMessageSender() external returns (address);
}
