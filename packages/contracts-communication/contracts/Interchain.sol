pragma solidity 0.8.20;

import './IInterchainModule.sol';
import 'forge-std/console.sol';

contract Interchain {
  uint64 public nonce;

  mapping(bytes32 => InterchainTransaction) public queuedTransactions;
  mapping(bytes32 => bool) public verifiedTransactions;

  constructor() {}

  mapping(bytes32 => uint) public completedModuleResponses;

  event InterchainTransactionSent(
    address srcSender,
    uint256 srcChainId,
    bytes32 indexed dstReceiver,
    uint256 indexed dstChainId,
    bytes message,
    uint64 nonce,
    bytes32 indexed transactionId
  );

  struct InterchainTransaction {
    address srcSender;
    uint256 srcChainId;
    bytes32 dstReceiver;
    uint256 dstChainId;
    bytes message;
    uint64 nonce;
    bytes32 transactionId;
    address[] modules;
    uint requiredModuleResponses;
  }

  function estimateInterchainTransactionFee(
    uint256 dstChainId,
    // TODO: Should anyone be able to call this, or should this only be used via IApps?
    address[] calldata modules
  ) public view returns (uint256) {
    uint256 totalFee = 0;
    for (uint i = 0; i < modules.length; i++) {
      totalFee =
        totalFee +
        IInterchainModuleV1(modules[i]).estimateFee(dstChainId);
    }
    return totalFee;
  }

  // TODO: Calculate Gas Pricing per module and charge fees
  // TODO: Customizable Gas Limit for Execution
  function interchainSend(
    bytes32 receiver,
    uint256 dstChainId,
    bytes calldata message,
    address[] calldata modules // Add modules as a parameter
  ) public payable {
    // TODO: Pull modules to send through from app config
    // Right now, we will just use the default module, as if the app config does not specify it.
    // This is a temporary solution until we have a proper app interface.

    InterchainTransaction memory newTransaction = InterchainTransaction(
      msg.sender,
      block.chainid,
      receiver,
      dstChainId,
      message,
      nonce,
      keccak256(abi.encodePacked('transactionId')), // TODO: dynamic ID generation
      modules,
      // TODO: This needs to move to app config, we cannot pass in modules directly and achieve all functionality
      modules.length // The number of required module responses is the length of the modules array
    );

    for (uint i = 0; i < newTransaction.modules.length; i++) {
      // TODO: How to disperse fees per module?
      // TODO: This is required per module right now, it will fail without it
      uint256 estimatedModuleFee = IInterchainModuleV1(newTransaction.modules[i])
        .estimateFee(newTransaction.dstChainId);
      // TODO: Right now, this could drain the Interchain.sol contract of any ETH held, since estimateFee is an untrusted function
      IInterchainModuleV1(newTransaction.modules[i]).sendModuleMessage{
        value: estimatedModuleFee
      }(abi.encode(newTransaction));
    }

    emit InterchainTransactionSent(
      msg.sender,
      block.chainid,
      receiver,
      dstChainId,
      message,
      nonce,
      bytes32(keccak256('transactionId'))
    );

    nonce++;
  }

  event TransactionReceived(bytes transaction);

  // TODO: Auth checks for which module is sending to destination receiver contract?
  function interchainReceive(bytes calldata transaction) public {
    InterchainTransaction memory interTransaction = abi.decode(
      transaction,
      (InterchainTransaction)
    );

    bytes32 transactionId = interTransaction.transactionId;

    // Check if the transaction is new or a module response
    if (queuedTransactions[transactionId].transactionId == bytes32(0)) {
      // New transaction
      queuedTransactions[transactionId] = interTransaction;
      // Initialize the module response for this transaction
      completedModuleResponses[transactionId] = 1;

      emit TransactionReceived(transaction);

      // TODO: We cannot rely on the interTransaction.requiredModuleResponses to provide us the required amount of module verifications
      // This is because any one module could then attempt to set something into execution that is not yet verified by the other modules
      // This config must happen at the destination app level for it to be secure
      // What are other things like this in terms of auth?
      // In general, application level config checks are completely missing.
      // Each message received should be verified by the destination receivers logic to ensure the correct module is verifying

      // Immediately verify if only one module response is required
      if (interTransaction.requiredModuleResponses == 1) {
        verifiedTransactions[transactionId] = true;
        // Optionally, remove the transaction from the queue
        // emit an event or take other actions as needed
      }
    } else {
      // Module response for an existing transaction
      require(
        queuedTransactions[transactionId].transactionId != bytes32(0),
        'Transaction does not exist'
      );
      uint numOfResponses = completedModuleResponses[transactionId];
      completedModuleResponses[transactionId] = ++numOfResponses;

      // TODO: Perform verification checks that all modules provided the same data for the transaction ID?
      // Check if all modules have completed
      if (
        numOfResponses >=
        queuedTransactions[transactionId].requiredModuleResponses
      ) {
        verifiedTransactions[transactionId] = true;
        // Optionally, remove the transaction from the queue
        // emit an event or take other actions as needed
        emit TransactionReceived(transaction);
      }
    }
  }

  // TODO: Gas Fee Consideration that is paid to executor
  function execute(bytes32 transactionId) public {
    require(
      verifiedTransactions[transactionId] == true,
      'Transaction not verified'
    );
    InterchainTransaction memory interTransaction = queuedTransactions[
      transactionId
    ];
    address dstReceiver = convertBytes32ToAddress(interTransaction.dstReceiver);
    // TODO: Modifidable gas & values
    dstReceiver.call{value: 0, gas: 1_000_000}(interTransaction.message);
  }

  // TODO: Seperate out into utils
  /**
   * @dev Converts a bytes32 to an address
   * @param _bytes32 The bytes32 value to be converted.
   * @return address The address obtained from the bytes32 input.
   */
  function convertBytes32ToAddress(
    bytes32 _bytes32
  ) public pure returns (address) {
    return address(uint160(uint256(_bytes32)));
  }

  /**
   * @dev Converts an address to bytes32
   * @param _address The address to be converted.
   * @return bytes32 The bytes32 representation of the input address.
   */
  function convertAddressToBytes32(
    address _address
  ) public pure returns (bytes32) {
    return bytes32(uint256(uint160(_address)));
  }
}
