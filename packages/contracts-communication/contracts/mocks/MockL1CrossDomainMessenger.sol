pragma solidity 0.8.20;

contract MockL1CrossDomainMessenger {
    event MessageSent(address target, bytes message, uint32 gasLimit);

    function sendMessage(address target, bytes calldata message, uint32 gasLimit) external payable {
        emit MessageSent(target, message, gasLimit);
    }

    function xDomainMessageSender(address msgSender) public view returns (address) {
        return msgSender;
    }
}
