pragma solidity 0.8.20;

import 'forge-std/Test.sol';
import '../contracts/modules/SynapseModule.sol';
import '../contracts/Interchain.sol';
import '../contracts/InterchainApp.sol';

import {ECDSA} from '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';

contract SynapseModuleE2ETest is Test {
  SynapseModule moduleL1;
  SynapseModule moduleL2;
  Vm.Wallet verifier1;
  Vm.Wallet verifier2;
  Vm.Wallet verifier3;
  address[] verifierAddresses;
  uint256[] verifierPrivateKeys;

  Interchain interchainL1;
  Interchain interchainL2;

  InterchainApp dstInterchainApp;
  InterchainApp originInterchainApp;

  function setUp() public {
    moduleL1 = new SynapseModule();
    moduleL2 = new SynapseModule();

    verifier1 = vm.createWallet(
      0x4a80c7a42367b3f50bc561c298c52233aca8e93dd511a9116281eb529e57ba14
    );
    verifier2 = vm.createWallet(
      0x1dea3ab3c2dddaea367a91af68d2e5097669417df0a195867db0b9f6079ddeeb
    );
    verifier3 = vm.createWallet(
      0xed292910d48741c4072b504ce6127157911b4eb56cefeb87a57e63217e0b7ee3
    );
    verifierAddresses = new address[](3);
    verifierAddresses[0] = verifier1.addr;
    verifierAddresses[1] = verifier2.addr;
    verifierAddresses[2] = verifier3.addr;
    verifierPrivateKeys = new uint256[](3);
    verifierPrivateKeys[0] = verifier1.privateKey;
    verifierPrivateKeys[1] = verifier2.privateKey;
    verifierPrivateKeys[
      2
    ] = 0xed292910d48741c4072b504ce6127157911b4eb56cefeb87a57e63217e0b7ee3;

    interchainL1 = new Interchain();
    interchainL2 = new Interchain();
    moduleL1.setInterchain(address(interchainL1));
    moduleL1.setVerifiers(verifierAddresses);
    moduleL2.setInterchain(address(interchainL2));
    moduleL2.setVerifiers(verifierAddresses);
    moduleL1.setRequiredThreshold(3);
    moduleL2.setRequiredThreshold(3);
    address[] memory l1Modules = new address[](1);
    l1Modules[0] = address(moduleL1);
    address[] memory l2Modules = new address[](1);
    l2Modules[0] = address(moduleL2);
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

  function testSendAndReceiveSynapseModuleMessage() public {
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
        new address[](1),
        1
      );
    interTransaction.modules[0] = address(moduleL1);
    bytes memory bytesInterTransaction = abi.encode(interTransaction);
    assertEq(
      interchainL1.nonce(),
      initialNonce + 1,
      'Nonce should increment by 1'
    );

    // Now message was sent
    // Mock out the verifiers responses
    bytes memory proof;
    bytes[] memory signatures = new bytes[](verifierAddresses.length);
    bytes32 messageHashToSign = keccak256(bytesInterTransaction);
    for (uint256 i = 0; i < verifierAddresses.length; i++) {
      (uint8 v, bytes32 r, bytes32 s) = vm.sign(
        verifierPrivateKeys[i],
        messageHashToSign
      );
      bytes memory signature = abi.encodePacked(r, s, v);
      (address recovered, ECDSA.RecoverError error, ) = ECDSA.tryRecover(
        messageHashToSign,
        signature
      );
      require(error == ECDSA.RecoverError.NoError, 'Invalid signature');
      console.log(recovered);

      signatures[i] = signature;
    }

    moduleL2.receiveModuleMessage(bytesInterTransaction, signatures);
  }
}
