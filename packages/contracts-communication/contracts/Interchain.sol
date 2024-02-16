pragma solidity 0.8.20;

import 'forge-std/console.sol';

contract Interchain {
  uint64 public nonce;

  mapping(bytes32 => InterchainTransaction) public queuedTransactions;
  mapping(bytes32 => bool) public verifiedTransactions;

  constructor() {}

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
  }

  // TODO: Calculate Gas Pricing per module and charge fees
  // TODO: Customizable Gas Limit for Execution
  function interchainSend(
    bytes32 receiver,
    uint256 dstChainId,
    bytes calldata message,
    address[] calldata modules // Add modules as a parameter
  ) public payable {
    // InterchainTransaction memory newTransaction = InterchainTransaction(
    //   msg.sender,
    //   block.chainid,
    //   receiver,
    //   dstChainId,
    //   message,
    //   nonce,
    //   keccak256(abi.encodePacked('transactionId')), // TODO: dynamic ID generation
    // );
    // emit InterchainTransactionSent(
    //   msg.sender,
    //   block.chainid,
    //   receiver,
    //   dstChainId,
    //   message,
    //   nonce,
    //   bytes32(keccak256('transactionId'))
    // );
    // nonce++;
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
