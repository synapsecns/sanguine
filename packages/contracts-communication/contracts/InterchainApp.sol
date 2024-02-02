pragma solidity 0.8.20;

import './IInterchain.sol';

contract InterchainApp {
  // What properties should Interchain be pulling from InterchainApp?
  // 1. Which modules to use, and how many are required?

  IInterchain public interchain;

  address[] private sendingModules;
  address[] private receivingModules;

  function getSendingModules(
    bytes32 receiver,
    uint256 dstChainId
  ) external view returns (address[] memory) {
    return sendingModules;
  }

  function getReceivingModules(
    bytes32 transactionId
  ) external view returns (address[] memory) {
    return receivingModules;
  }

  constructor(
    address _interchain,
    address[] memory _sendingModules,
    address[] memory _receivingModules
  ) {
    interchain = IInterchain(_interchain);
    sendingModules = _sendingModules;
    receivingModules = _receivingModules;
  }

  event AppMessageRecieve();
  event AppMessageSent();

  function send(
    bytes32 receiver,
    uint256 dstChainId,
    bytes calldata message
  ) external {
    interchain.interchainSend(receiver, dstChainId, message, sendingModules);
    emit AppMessageSent();
  }

  // TODO: Auth checks based on incoming message
  function appReceive() external {
    emit AppMessageRecieve();
  }
}
