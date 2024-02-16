pragma solidity 0.8.20;

import './IInterchain.sol';

contract InterchainApp {
  // What properties should Interchain be pulling from InterchainApp?
  // 1. Which modules to use, and how many are required?

  IInterchain public interchain;

  address[] private sendingModules;
  address[] private receivingModules;

  struct AppConfig {
    // ChainID -> Linked IApps
    mapping(uint64 => address) linkedIApps;
    // Sends message to be verified through all modules
    address[] sendingModules;
    // Accepts messages from these destination chain modules
    address[] receivingModules;
    // Threshold for execution
    uint requiredResponses;
    // Time period for optimistic execution
    uint64 optimisticTimePeriod; // in seconds
  }

  AppConfig private appConfig;

  // Set the application configuration
  function setAppConfig(
    uint64[] memory chainIDs,
    address[] memory linkedIApps,
    address[] memory _sendingModules,
    address[] memory _receivingModules,
    uint _requiredResponses,
    uint64 _optimisticTimePeriod
  ) public {
    // TODO: Add access control or ownership checks
    require(
      chainIDs.length == linkedIApps.length,
      'ChainIDs and IApps length mismatch'
    );

    for (uint i = 0; i < chainIDs.length; i++) {
      appConfig.linkedIApps[chainIDs[i]] = linkedIApps[i];
    }

    appConfig.sendingModules = _sendingModules;
    appConfig.receivingModules = _receivingModules;
    appConfig.requiredResponses = _requiredResponses;
    appConfig.optimisticTimePeriod = _optimisticTimePeriod;
  }

  // Getters for the application configuration
  function getLinkedIApp(uint64 chainID) external view returns (address) {
    return appConfig.linkedIApps[chainID];
  }

  // TODO: Is a receiving module the same as a sending module?
  function getSendingModules() external view returns (address[] memory) {
    return appConfig.sendingModules;
  }

  function getReceivingModules() external view returns (address[] memory) {
    return appConfig.receivingModules;
  }

  function getRequiredResponses() external view returns (uint) {
    return appConfig.requiredResponses;
  }

  function getOptimisticTimePeriod() external view returns (uint64) {
    return appConfig.optimisticTimePeriod;
  }

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
    appConfig.sendingModules = _sendingModules;
    appConfig.receivingModules = _receivingModules;
  }

  event AppMessageRecieve();
  event AppMessageSent();

  function send(
    bytes32 receiver,
    uint256 dstChainId,
    bytes calldata message
  ) external payable {
    // TODO: Currently, we forward all gas to Interchain, this may not be expected behavior, and the real abstract contract shouldn't do this
    interchain.interchainSend{value: msg.value}(
      receiver,
      dstChainId,
      message,
      appConfig.sendingModules
    );
    emit AppMessageSent();
  }

  // TODO: Auth checks based on incoming message
  function appReceive() external {
    emit AppMessageRecieve();
  }
}
