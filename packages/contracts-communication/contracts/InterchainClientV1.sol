pragma solidity 0.8.20;

import {Ownable} from '@openzeppelin/contracts/access/Ownable.sol';
import 'forge-std/console.sol';
import {IInterchainDB} from './interfaces/IInterchainDB.sol';

contract InterchainClientV1 is Ownable {
  uint64 public clientNonce;
  address public interchainDB;

  mapping(bytes32 => InterchainTransaction) public queuedTransactions;

  constructor() public Ownable(msg.sender) {}

  function setInterchainDB(address _interchainDB) public onlyOwner {
    interchainDB = _interchainDB;
  }

  event InterchainTransactionSent(
    bytes32 srcSender,
    uint256 srcChainId,
    bytes32 indexed dstReceiver,
    uint256 indexed dstChainId,
    bytes message,
    uint64 nonce,
    bytes32 indexed transactionId,
    uint256 dbWriterNonce
  );

  struct InterchainTransaction {
    bytes32 srcSender;
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
    address[] calldata srcModules // Add modules as a parameter
  ) public payable {
    uint256 totalModuleFees = msg.value;
    bytes32 sender = convertAddressToBytes32(msg.sender);
    bytes32 transactionID = keccak256(
      abi.encode(
        sender,
        block.chainid,
        receiver,
        dstChainId,
        message,
        clientNonce
      )
    );

    uint256 dbWriterNonce = IInterchainDB(interchainDB)
      .writeEntryWithVerification{value: totalModuleFees}(
      dstChainId,
      transactionID,
      srcModules
    );

    emit InterchainTransactionSent(
      sender,
      block.chainid,
      receiver,
      dstChainId,
      message,
      clientNonce,
      transactionID,
      dbWriterNonce
    );
    // Increment nonce for next message
    clientNonce++;
  }

  // TODO: Gas Fee Consideration that is paid to executor
  function execute() public {}

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
