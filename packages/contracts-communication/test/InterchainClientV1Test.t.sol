// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {AppConfigV1, InterchainTransaction} from "../contracts/InterchainClientV1.sol";
import {InterchainDB} from "../contracts/InterchainDB.sol";

import {InterchainEntry} from "../contracts/libs/InterchainEntry.sol";
import {OptionsV1} from "../contracts/libs/Options.sol";
import {TypeCasts} from "../contracts/libs/TypeCasts.sol";

import {InterchainClientV1Harness} from "./harnesses/InterchainClientV1Harness.sol";

import {ExecutionFeesMock} from "./mocks/ExecutionFeesMock.sol";
import {ExecutionServiceMock} from "./mocks/ExecutionServiceMock.sol";
import {InterchainAppMock} from "./mocks/InterchainAppMock.sol";
import {InterchainModuleMock} from "./mocks/InterchainModuleMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable max-line-length
// solhint-disable ordering
contract InterchainClientV1Test is Test {
    ExecutionFeesMock public executionFees;
    ExecutionServiceMock public executionService;

    InterchainClientV1Harness public icClient;
    InterchainDB public icDB;
    InterchainAppMock public icApp;
    InterchainModuleMock public icModule;

    address[] public mockApprovedModules;

    uint256 public constant GAS_AIRDROP = 0.5 ether;
    // Use default options of V1, 200k gas limit, 0.5 ETH gas airdrop
    bytes public options = OptionsV1(200_000, GAS_AIRDROP).encodeOptionsV1();

    uint256 public constant SRC_CHAIN_ID = 1337;
    uint256 public constant DST_CHAIN_ID = 7331;

    address public contractOwner = makeAddr("Contract Owner");

    function setUp() public {
        vm.startPrank(contractOwner);
        executionFees = new ExecutionFeesMock();
        executionService = new ExecutionServiceMock();
        icDB = new InterchainDB();
        icClient = new InterchainClientV1Harness(address(icDB));
        icClient.setExecutionFees(address(executionFees));

        icModule = new InterchainModuleMock();
        icApp = new InterchainAppMock();
        icApp.setReceivingModule(address(icModule));
        vm.stopPrank();
        mockModuleFee(icModule, 1);
        mockApprovedModules.push(address(icModule));
    }

    /// @dev Mocks a return value of module.getModuleFee(DST_CHAIN_ID)
    function mockModuleFee(InterchainModuleMock module, uint256 feeValue) internal {
        bytes memory callData = abi.encodeCall(module.getModuleFee, (DST_CHAIN_ID));
        bytes memory returnData = abi.encode(feeValue);
        vm.mockCall(address(module), callData, returnData);
    }

    // ══════════════════════════════════════════════ INTERNAL TESTS ══════════════════════════════════════════════════

    // TODO: split into multiple single-purpose tests
    /*
    function test_getFinalizedResponsesCount() public {
        vm.warp(10 days);
        uint256[] memory approvedResponses = new uint256[](3);
        approvedResponses[0] = block.timestamp + 1 days;
        // This should be counted as finalized because it's outside the optimistic time period
        approvedResponses[1] = block.timestamp - 20 minutes;
        // This should not be counted as finalized
        approvedResponses[2] = block.timestamp - 1 minutes;

        uint256 optimisticTimePeriod = 15 minutes; // Setting the optimistic time period to 15 minutes

        uint256 responses = icClient.getFinalizedResponsesCountHarness(approvedResponses, optimisticTimePeriod);

        assertEq(responses, 1, "Only 1 response should be finalized within the optimistic time period");

        // Test with all responses outside the optimistic time period
        approvedResponses[0] = block.timestamp + 30 minutes;
        approvedResponses[1] = block.timestamp + 40 minutes;
        approvedResponses[2] = block.timestamp + 50 minutes;

        responses = icClient.getFinalizedResponsesCountHarness(approvedResponses, optimisticTimePeriod);

        assertEq(responses, 0, "There should be 0 finalized responses outside the optimistic time period");

        // Test with empty responses array
        approvedResponses = new uint256[](0);

        responses = icClient.getFinalizedResponsesCountHarness(approvedResponses, optimisticTimePeriod);

        assertEq(responses, 0, "There should be 0 finalized responses with an empty responses array");
    }
    */
    function test_interchainSend() public {
        bytes32 receiver = TypeCasts.addressToBytes32(makeAddr("Receiver"));
        bytes memory message = "Hello World";
        address[] memory srcModules = new address[](1);
        srcModules[0] = address(icModule);
        uint256 totalModuleFees = 1;
        uint64 nonce = 1;
        bytes32 transactionID = keccak256(
            abi.encode(TypeCasts.addressToBytes32(msg.sender), block.chainid, receiver, DST_CHAIN_ID, message, nonce)
        );
        icClient.interchainSend{value: totalModuleFees}(
            DST_CHAIN_ID, receiver, address(executionService), srcModules, options, message
        );
        // TODO: should check the transaction ID?
        transactionID;
    }

    // TODO: more tests
    function test_interchainReceive() public {
        vm.chainId(DST_CHAIN_ID);
        bytes32 dstReceiver = TypeCasts.addressToBytes32(address(icApp));
        bytes memory message = "Hello World";
        bytes32 srcSender = TypeCasts.addressToBytes32(makeAddr("Sender"));
        vm.prank(contractOwner);
        icClient.setLinkedClient(SRC_CHAIN_ID, srcSender);
        uint64 nonce = 1;
        uint256 dbNonce = 2;
        InterchainTransaction memory transaction = InterchainTransaction({
            srcSender: srcSender,
            srcChainId: SRC_CHAIN_ID,
            dstReceiver: dstReceiver,
            dstChainId: DST_CHAIN_ID,
            nonce: nonce,
            dbNonce: dbNonce,
            options: options,
            message: message
        });
        bytes32 transactionID = keccak256(abi.encode(transaction));
        bytes memory expectedAppCalldata = abi.encodeCall(icApp.appReceive, (SRC_CHAIN_ID, srcSender, nonce, message));

        AppConfigV1 memory mockAppConfig = AppConfigV1({requiredResponses: 1, optimisticPeriod: 1 hours});
        vm.mockCall(
            address(icApp),
            abi.encodeCall(icApp.getReceivingConfig, ()),
            abi.encode(mockAppConfig.encodeAppConfigV1(), mockApprovedModules)
        );

        InterchainEntry memory entry =
            InterchainEntry({srcChainId: SRC_CHAIN_ID, srcWriter: srcSender, dbNonce: dbNonce, dataHash: transactionID});

        icModule.mockVerifyEntry(address(icDB), entry);

        deal(address(this), GAS_AIRDROP);
        // Skip ahead of optimistic period
        skip(mockAppConfig.optimisticPeriod + 1);
        // Expect App to be called with the message
        vm.expectCall({callee: address(icApp), msgValue: GAS_AIRDROP, gas: 200_000, data: expectedAppCalldata, count: 1});
        icClient.interchainExecute{value: GAS_AIRDROP}({gasLimit: 0, transaction: abi.encode(transaction)});
    }
}
