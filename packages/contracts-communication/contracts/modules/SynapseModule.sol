pragma solidity 0.8.20;

import {Ownable} from '@openzeppelin/contracts/access/Ownable.sol';
import {ECDSA} from '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';
import {Interchain} from '../Interchain.sol';
import '../IInterchain.sol';
import 'forge-std/console.sol';

contract SynapseModule is Ownable {
  address[] public verifiers;
  uint256 requiredThreshold;
  address public interchain;

  constructor() public Ownable(msg.sender) {}

  function setInterchain(address _interchain) public onlyOwner {
    interchain = _interchain;
    emit ConfigUpdated();
  }

  function setRequiredThreshold(uint256 _threshold) public onlyOwner {
    requiredThreshold = _threshold;
    emit ConfigUpdated();
  }

  function setVerifiers(address[] calldata _verifiers) public onlyOwner {
    verifiers = _verifiers;
    emit ConfigUpdated();
  }

  event ConfigUpdated();
  event ModuleMessageSent(uint256 dstChainId, bytes transaction);

  function estimateFee(uint256 dstChainId) public view returns (uint256) {
    // Get Latest Posted Destination Gas Price from oracle
    // Requires: Access to origin USD Gas Price / Destination USD Gas PRice
    // Get current price of origin chain assets
    // Get current price of destination chain assets
    // Calculate the estiamted fee based on preset gas limit
    // return

    // TODO: Right now, we don't have all of the info needed to provide a real fee estimation - we will provide 1 wei to enable other functionality to be built.
    return 1;
  }

  function decodeInterchainTransaction(bytes calldata transaction) public view returns (Interchain.InterchainTransaction memory) {
    Interchain.InterchainTransaction memory decodedTransaction = abi.decode(
      transaction,
      (Interchain.InterchainTransaction)
    );
    return decodedTransaction;
  }

  function sendModuleMessage(bytes calldata transaction) public payable {
    Interchain.InterchainTransaction memory decodedTransaction = abi.decode(
      transaction,
      (Interchain.InterchainTransaction)
    );

    // TODO: Require fee is above estimation(?), does this make sense?
    uint256 currentEstimatedFee = estimateFee(decodedTransaction.dstChainId);
    require(
      msg.value >= currentEstimatedFee,
      'Insufficient fee to send transaction'
    );
    // Transfer fee to module executor

    emit ModuleMessageSent(decodedTransaction.dstChainId, transaction);
  }

  function receiveModuleMessage(
    bytes calldata transaction,
    bytes[] calldata signatures
  ) public {
    bytes32 messageHashToCheck = keccak256(transaction);

    require(
      verifiers.length >= requiredThreshold,
      'Not enough verifiers to meet the threshold'
    );

    uint256 validSignatures;
    for (uint256 i = 0; i < verifiers.length; i++) {
      // TODO: Use TryRecover for explicit error handling
      address signer = ECDSA.recover(messageHashToCheck, signatures[i]);
      for (uint256 j = 0; j < verifiers.length; j++) {
        if (signer == verifiers[j]) {
          validSignatures++;
          break;
        }
      }
    }
    bytes32 hash = keccak256(transaction);

    require(
      validSignatures >= requiredThreshold,
      'Not enough valid signatures to meet the threshold'
    );
    IInterchain(interchain).interchainReceive(transaction);
  }
}
