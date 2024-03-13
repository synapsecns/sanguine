pragma solidity =0.8.20;

// contracts/modules/IL1CrossDomainMessenger.sol

interface IL1CrossDomainMessenger {
    function sendMessage(address target, bytes calldata message, uint32 gasLimit) external payable;

    function xDomainMessageSender() external returns (address);
}
