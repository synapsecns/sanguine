pragma solidity 0.8.20;

interface IInterchainModule {
  function sendModuleMessage(bytes calldata transaction) external;

  function receiveModuleMessage() external;
}
