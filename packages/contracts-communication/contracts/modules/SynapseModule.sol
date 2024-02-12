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
  }

  function setRequiredThreshold(uint256 _threshold) public onlyOwner {
    requiredThreshold = _threshold;
  }

  function setVerifiers(address[] calldata _verifiers) public onlyOwner {
    verifiers = _verifiers;
  }

  event ModuleMessageSent(uint256 dstChainId, bytes transaction);

  function sendModuleMessage(bytes calldata transaction) public {
    Interchain.InterchainTransaction memory decodedTransaction = abi.decode(
      transaction,
      (Interchain.InterchainTransaction)
    );

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
