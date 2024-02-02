pragma solidity 0.8.20;

import 'forge-std/Test.sol';
import 'forge-std/console.sol';
import '../contracts/mocks/MockL1CrossDomainMessenger.sol';
import '../contracts/mocks/MockL2CrossDomainMessenger.sol';
import '../contracts/modules/OPL1L2Module.sol';
import '../contracts/Interchain.sol';
import '../contracts/InterchainApp.sol';

contract OPL1L2MultiModuleE2ETest is Test {
  MockL1CrossDomainMessenger mockL1Messenger;
  MockL2CrossDomainMessenger mockL2Messenger;
  Interchain interchainL1;
  Interchain interchainL2;
  OPL1L2Module moduleL1_1;
  OPL1L2Module moduleL2_1;
  OPL1L2Module moduleL1_2;
  OPL1L2Module moduleL2_2;

  InterchainApp dstInterchainApp;
  InterchainApp originInterchainApp;

  address optimismRelayer = address(0x123);

  function setUp() public {
    // Deploying the mock L1 and L2 messengers
    mockL1Messenger = new MockL1CrossDomainMessenger();
    mockL2Messenger = new MockL2CrossDomainMessenger();

    // Deploying the InterchainModule contracts for L1 and L2
    moduleL1_1 = new OPL1L2Module(
      address(mockL1Messenger),
      address(moduleL1_1),
      address(mockL2Messenger),
      address(moduleL2_1),
      address(interchainL1)
    );

    moduleL2_1 = new OPL1L2Module(
      address(mockL1Messenger),
      address(moduleL1_1),
      address(mockL2Messenger),
      address(moduleL2_1),
      address(interchainL2)
    );

    moduleL1_2 = new OPL1L2Module(
      address(mockL1Messenger),
      address(moduleL1_2),
      address(mockL2Messenger),
      address(moduleL2_2),
      address(interchainL1)
    );

    moduleL2_2 = new OPL1L2Module(
      address(mockL1Messenger),
      address(moduleL1_2),
      address(mockL2Messenger),
      address(moduleL2_2),
      address(interchainL2)
    );

    interchainL1 = new Interchain();
    interchainL2 = new Interchain();

    moduleL1_1.setL2OPModule(address(moduleL2_1));

    moduleL2_1.setL1OPModule(address(moduleL1_1));

    moduleL1_2.setL2OPModule(address(moduleL2_2));

    moduleL2_2.setL1OPModule(address(moduleL1_2));

    moduleL1_1.setInterchain(address(interchainL1));

    moduleL2_1.setInterchain(address(interchainL2));

    moduleL1_2.setInterchain(address(interchainL1));

    moduleL2_2.setInterchain(address(interchainL2));

    address[] memory l1Modules = new address[](2);
    l1Modules[0] = address(moduleL1_1);
    l1Modules[1] = address(moduleL1_2);

    address[] memory l2Modules = new address[](2);
    l2Modules[0] = address(moduleL2_1);
    l2Modules[1] = address(moduleL2_2);

    originInterchainApp = new InterchainApp(
      address(interchainL1),
      l1Modules,
      l2Modules
    );
    dstInterchainApp = new InterchainApp(
      address(interchainL2),
      l2Modules,
      l1Modules
    );
  }

  function testSendAndReceiveMultiModuleMessage() public {
    bytes32 receiver = interchainL2.convertAddressToBytes32(
      address(dstInterchainApp)
    );
    uint256 dstChainId = 1;
    bytes memory message = abi.encodeWithSelector(
      InterchainApp.appReceive.selector
    );
    uint64 initialNonce = interchainL1.nonce();

    originInterchainApp.send(receiver, dstChainId, message);
    Interchain.InterchainTransaction memory interTransaction = Interchain
      .InterchainTransaction(
        address(this),
        block.chainid,
        receiver,
        dstChainId,
        message,
        initialNonce,
        keccak256(abi.encodePacked('transactionId')),
        new address[](2),
        2
      );
    interTransaction.modules[0] = address(moduleL1_1);
    interTransaction.modules[0] = address(moduleL1_2);
    bytes memory bytesInterTransaction = abi.encode(interTransaction);
    assertEq(
      interchainL1.nonce(),
      initialNonce + 1,
      'Nonce should increment by 1'
    );

    vm.startPrank(optimismRelayer);
    mockL2Messenger.relayMessage(
      0,
      address(moduleL1_1),
      address(moduleL2_1),
      0,
      1_000_000,
      abi.encodeWithSelector(
        bytes4(keccak256('receiveModuleMessage(bytes)')),
        bytesInterTransaction
      )
    );

    assertEq(
      interchainL2.verifiedTransactions(interTransaction.transactionId),
      false
    );

    uint completedModules = interchainL2.completedModuleResponses(
      interTransaction.transactionId
    );
    assertEq(completedModules, 1);

    mockL2Messenger.relayMessage(
      0,
      address(moduleL1_2),
      address(moduleL2_2),
      0,
      1_000_000,
      abi.encodeWithSelector(
        bytes4(keccak256('receiveModuleMessage(bytes)')),
        bytesInterTransaction
      )
    );

    assertEq(
      interchainL2.verifiedTransactions(interTransaction.transactionId),
      true
    );

    completedModules = interchainL2.completedModuleResponses(
      interTransaction.transactionId
    );
    assertEq(completedModules, 2);

    interchainL2.execute(interTransaction.transactionId);
  }
}
