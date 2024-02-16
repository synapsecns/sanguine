pragma solidity 0.8.20;

import 'forge-std/Test.sol';
import '../contracts/InterchainClientV1.sol';
import '../contracts/InterchainDB.sol';
import '../contracts/modules/SynapseModule.sol';

contract InterchainClientV1Test is Test {
  InterchainClientV1 icClient;
  InterchainDB icDB;
  SynapseModule synapseModule;

  uint256 public constant SRC_CHAIN_ID = 1337;
  uint256 public constant DST_CHAIN_ID = 7331;

  address public contractOwner = makeAddr('Contract Owner');

  function setUp() public {
    vm.startPrank(contractOwner);
    icClient = new InterchainClientV1();
    icDB = new InterchainDB();
    synapseModule = new SynapseModule();
    synapseModule.setInterchainDB(address(icDB));
    icClient.setInterchainDB(address(icDB));
    vm.stopPrank();
  }

  function test_interchainSend() public {
    bytes32 receiver = icClient.convertAddressToBytes32(makeAddr('Receiver'));
    bytes memory message = 'Hello World';
    address[] memory srcModules = new address[](1);
    srcModules[0] = address(synapseModule);
    uint256 totalModuleFees = 1;
    uint64 nonce = 1;
    bytes32 transactionID = keccak256(
      abi.encode(
        icClient.convertAddressToBytes32(msg.sender),
        block.chainid,
        receiver,
        DST_CHAIN_ID,
        message,
        nonce
      )
    );
    icClient.interchainSend{value: 1}(
      receiver,
      DST_CHAIN_ID,
      message,
      srcModules
    );
  }
}
