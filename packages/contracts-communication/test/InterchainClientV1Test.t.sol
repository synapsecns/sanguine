pragma solidity 0.8.20;

import "forge-std/Test.sol";
import {InterchainClientV1} from "../contracts/InterchainClientV1.sol";
import "../contracts/InterchainDB.sol";
import {InterchainAppMock} from "./mocks/InterchainAppMock.sol";

import {InterchainModuleMock} from "./mocks/InterchainModuleMock.sol";
import "../contracts/modules/SynapseModule.sol";

import {InterchainEntry} from "../contracts/libs/InterchainEntry.sol";

import {TypeCasts} from "../contracts/libs/TypeCasts.sol";

contract InterchainClientV1Test is Test {
    InterchainClientV1 icClient;
    InterchainDB icDB;
    SynapseModule synapseModule;
    InterchainAppMock icApp;
    InterchainModuleMock icModule;

    uint256 public constant SRC_CHAIN_ID = 1337;
    uint256 public constant DST_CHAIN_ID = 7331;

    address public contractOwner = makeAddr("Contract Owner");

    function setUp() public {
        vm.startPrank(contractOwner);
        icClient = new InterchainClientV1();
        icDB = new InterchainDB();
        synapseModule = new SynapseModule();
        synapseModule.setInterchainDB(address(icDB));
        icClient.setInterchainDB(address(icDB));

        icModule = new InterchainModuleMock();
        icApp = new InterchainAppMock();
        icApp.setReceivingModule(address(icModule));
        vm.stopPrank();
    }

    function test_interchainSend() public {
        bytes32 receiver = TypeCasts.addressToBytes32(makeAddr("Receiver"));
        bytes memory message = "Hello World";
        address[] memory srcModules = new address[](1);
        srcModules[0] = address(synapseModule);
        uint256 totalModuleFees = 1;
        uint64 nonce = 1;
        bytes32 transactionID = keccak256(
            abi.encode(
                TypeCasts.addressToBytes32(msg.sender), block.chainid, receiver, DST_CHAIN_ID, message, nonce
            )
        );
        bytes memory options = "";
        icClient.interchainSend{value: 1}(receiver, DST_CHAIN_ID, message, options, srcModules);
    }

    function test_interchainReceive() public {
        bytes32 dstReceiver = TypeCasts.addressToBytes32(address(icApp));
        bytes memory message = "Hello World";
        bytes memory options = "";
        bytes32 srcSender = TypeCasts.addressToBytes32(makeAddr("Sender"));
        icClient.setLinkedClient(SRC_CHAIN_ID, srcSender);
        uint64 nonce = 1;
        bytes32 transactionID =
            keccak256(abi.encode(srcSender, SRC_CHAIN_ID, dstReceiver, DST_CHAIN_ID, message, nonce, options));

        InterchainEntry memory entry =
            InterchainEntry({srcChainId: SRC_CHAIN_ID, srcWriter: srcSender, writerNonce: 0, dataHash: transactionID});

        icModule.mockVerifyEntry(address(icDB), entry);

        InterchainClientV1.InterchainTransaction memory transaction = InterchainClientV1.InterchainTransaction({
            srcSender: srcSender,
            srcChainId: SRC_CHAIN_ID,
            dstReceiver: dstReceiver,
            dstChainId: DST_CHAIN_ID,
            message: message,
            nonce: nonce,
            options: options,
            transactionId: transactionID,
            dbWriterNonce: 0
        });

        icClient.interchainExecute(transactionID, abi.encode(transaction));
    }
}
