pragma solidity 0.8.20;

contract InterchainModule {
  // TODO: All modules need to have an understanding of the gas limit needed to execute on the destination chain.

  function sendModuleMessage(bytes calldata transaction) public {
    // This function would send the transaction to the module.
    // As `sendModuleMessage` is not implemented, we're leaving this as a placeholder.
  }

  function receiveModuleMessage() public {}
}
