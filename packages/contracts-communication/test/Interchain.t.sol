pragma solidity 0.8.20;

import 'forge-std/Test.sol';
import '../contracts/Interchain.sol';
import '../contracts/InterchainModule.sol';

// contract InterchainTest is Test {
//   Interchain interchain;
//   InterchainModule module;
//   address defaultModuleAddress = address(0x123);

//   function setUp() public {
//     module = new InterchainModule();
//     interchain = new Interchain(address(module), address(0));
//   }

//   function testInterchainSend() public {
//     bytes32 receiver = bytes32('Receiver');
//     uint256 dstChainId = 1;
//     bytes memory message = 'Hello, Interchain!';
//     uint64 initialNonce = interchain.nonce();

//     interchain.interchainSend(receiver, dstChainId, message);

//     assertEq(
//       interchain.nonce(),
//       initialNonce + 1,
//       'Nonce should increment by 1'
//     );
//     // This is a simplistic way to check if an event was emitted.
//     // In a real test, you would use `expectEmit` to check the event's parameters.
//     //     vm.expectEmit(true, true, true, true);
//     //     emit interchain.InterchainTransactionSent(
//     //       address(this),
//     //       block.chainid,
//     //       receiver,
//     //       dstChainId,
//     //       message,
//     //       initialNonce,
//     //       keccak256(abi.encodePacked('messageId'))
//     //     );
//   }

//   function testInterchainReceive() public {
//     // This function would test the `interchainReceive` function.
//     // As `interchainReceive` is not implemented, we're leaving this as a placeholder.
//   }
// }
