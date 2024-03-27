// pragma solidity 0.8.20;

// import 'forge-std/Test.sol';
// import '../contracts/mocks/MockL1CrossDomainMessenger.sol';
// import '../contracts/mocks/MockL2CrossDomainMessenger.sol';
// import '../contracts/modules/OPL1L2Module.sol';
// import '../contracts/Interchain.sol';

// contract OPL1L2ModuleE2ETest is Test {
//   MockL1CrossDomainMessenger mockL1Messenger;
//   MockL2CrossDomainMessenger mockL2Messenger;
//   Interchain interchainL1;
//   Interchain interchainL2;
//   OPL1L2Module moduleL1;
//   OPL1L2Module moduleL2;

//   address optimismRelayer = address(0x123);

//   function setUp() public {
//     // Deploying the mock L1 and L2 messengers
//     mockL1Messenger = new MockL1CrossDomainMessenger();
//     mockL2Messenger = new MockL2CrossDomainMessenger();

//     // Deploying the InterchainModule contracts for L1 and L2
//     moduleL1 = new OPL1L2Module(
//       address(mockL1Messenger),
//       address(moduleL1),
//       address(mockL2Messenger),
//       address(moduleL2),
//       address(interchainL1)
//     );

//     moduleL2 = new OPL1L2Module(
//       address(mockL1Messenger),
//       address(moduleL1),
//       address(mockL2Messenger),
//       address(moduleL2),
//       address(interchainL2)
//     );

//     interchainL1 = new Interchain(address(moduleL1), address(0));
//     interchainL2 = new Interchain(address(moduleL2), address(0));

//     moduleL1.setL2OPModule(address(moduleL2));

//     moduleL2.setL1OPModule(address(moduleL1));

//     moduleL1.setInterchain(address(interchainL1));

//     moduleL2.setInterchain(address(interchainL2));
//   }

//   function testSendAndReceiveSingleModuleMessage() public {
//     bytes32 receiver = bytes32('Receiver');
//     uint256 dstChainId = 1;
//     bytes memory message = 'Hello, Interchain!';
//     uint64 initialNonce = interchainL1.nonce();

//     interchainL1.interchainSend(receiver, dstChainId, message);
//     Interchain.InterchainTransaction memory interTransaction = Interchain
//       .InterchainTransaction(
//         address(this),
//         block.chainid,
//         receiver,
//         dstChainId,
//         message,
//         initialNonce,
//         keccak256(abi.encodePacked('transactionId')),
//         new address[](1),
//         1
//       );
//     interTransaction.modules[0] = address(moduleL1);
//     bytes memory bytesInterTransaction = abi.encode(interTransaction);
//     assertEq(
//       interchainL1.nonce(),
//       initialNonce + 1,
//       'Nonce should increment by 1'
//     );

//     vm.startPrank(optimismRelayer);
//     mockL2Messenger.relayMessage(
//       0,
//       address(moduleL1),
//       address(moduleL2),
//       0,
//       1_000_000,
//       abi.encodeWithSelector(
//         bytes4(keccak256('receiveModuleMessage(bytes)')),
//         bytesInterTransaction
//       )
//     );

//     assertEq(
//       interchainL2.verifiedTransactions(interTransaction.transactionId),
//       true
//     );

//     // moduleL2.receiveModuleMessage(transaction);
//   }
// }
