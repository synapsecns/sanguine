pragma solidity 0.8.20;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {SynapseGasService} from "./SynapseGasService.sol";
import "../IInterchain.sol";
import "forge-std/console.sol";
import {IInterchainDB} from "../interfaces/IInterchainDB.sol";
import {IInterchainModule} from "../interfaces/IInterchainModule.sol";
import {InterchainEntry} from "../libs/InterchainEntry.sol";

import {ISynapseModuleEvents} from "../interfaces/ISynapseModuleEvents.sol";

contract SynapseModule is Ownable, SynapseGasService, ISynapseModuleEvents {
    address[] public verifiers;
    uint256 public requiredThreshold;
    address public interchainDB;

    constructor() public Ownable(msg.sender) {}

    function setInterchainDB(address _interchainDB) public onlyOwner {
        interchainDB = _interchainDB;
    }

    function setRequiredThreshold(uint256 _threshold) public onlyOwner {
        requiredThreshold = _threshold;
    }

    function setVerifiers(address[] calldata _verifiers) public onlyOwner {
        verifiers = _verifiers;
    }

    function requestVerification(uint256 destChainId, InterchainEntry memory entry) external payable {
        require(msg.sender == interchainDB, "Only InterchainDB can request verification");

        require(msg.value >= getModuleFee(destChainId), "Insufficient fee to request verification");

        _payFeesForExecution(msg.value);
        emit VerificationRequested(destChainId, entry);
    }

    function verifyEntry(InterchainEntry memory entry, bytes[] calldata signatures) external {
        bytes32 messageHashToCheck = keccak256(abi.encode(entry));

        require(signatures.length >= requiredThreshold, "Not enough signatures to meet the threshold");

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

        require(validSignatures >= requiredThreshold, "Not enough valid signatures to meet the threshold");

        IInterchainDB(interchainDB).verifyEntry(entry);

        emit EntryVerified(entry);
    }
}
