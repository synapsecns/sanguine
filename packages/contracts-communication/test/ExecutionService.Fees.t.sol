// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ExecutionService, ExecutionServiceEvents} from "../contracts/ExecutionService.sol";
import {OptionsLib, OptionsV1} from "../contracts/libs/Options.sol";

import {SynapseGasOracleMock} from "./mocks/SynapseGasOracleMock.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract ExecutionServiceFeesTest is ExecutionServiceEvents, Test {
    bytes32 public constant MOCK_TX_ID = keccak256("mock-tx-id");

    uint256 public constant MOCK_GAS_LIMIT = 100_000;
    uint256 public constant MOCK_GAS_AIRDROP = 0.01 ether;
    uint256 public constant MOCK_CALLDATA_SIZE = 1024;

    // Execution of tx costs 0.1 remote ETH
    uint256 public constant MOCK_REMOTE_EXEC_COST = 0.1 ether;
    // 1 remote ETH = 2 local ETH
    // Execution cost in local units
    uint256 public constant MOCK_LOCAL_EXEC_COST = MOCK_REMOTE_EXEC_COST * 2;
    uint256 public constant MOCK_LOCAL_AIRDROP_COST = MOCK_GAS_AIRDROP * 2;

    uint256 public constant MOCK_FEE_NO_AIRDROP = MOCK_LOCAL_EXEC_COST;
    uint256 public constant MOCK_FEE_WITH_AIRDROP = MOCK_LOCAL_EXEC_COST + MOCK_LOCAL_AIRDROP_COST;

    uint256 public constant LOCAL_CHAIN_ID = 1;
    uint256 public constant REMOTE_CHAIN_ID = 1337;

    OptionsV1 public optionsNoAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: 0});
    OptionsV1 public optionsWithAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: MOCK_GAS_AIRDROP});

    ExecutionService public service;
    SynapseGasOracleMock public gasOracle;

    address public icClient = makeAddr("InterchainClient");
    address public executorEOA = makeAddr("ExecutorEOA");
    address public owner = makeAddr("Owner");

    function setUp() public {
        vm.chainId(LOCAL_CHAIN_ID);
        gasOracle = new SynapseGasOracleMock();
        service = new ExecutionService(address(this));
        service.setInterchainClient(icClient);
        service.setExecutorEOA(executorEOA);
        service.setGasOracle(address(gasOracle));
        service.transferOwnership(owner);
        mockGasPricing();
    }

    function mockGasPricing() internal {
        // Mock: execution cost in local units
        vm.mockCall(
            address(gasOracle),
            abi.encodeCall(
                SynapseGasOracleMock.estimateTxCostInLocalUnits, (REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE)
            ),
            abi.encode(MOCK_LOCAL_EXEC_COST)
        );
        // Mock: execution cost in remote units + its conversion to local units
        vm.mockCall(
            address(gasOracle),
            abi.encodeCall(
                SynapseGasOracleMock.estimateTxCostInRemoteUnits, (REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE)
            ),
            abi.encode(MOCK_REMOTE_EXEC_COST)
        );
        vm.mockCall(
            address(gasOracle),
            abi.encodeCall(
                SynapseGasOracleMock.convertRemoteValueToLocalUnits, (REMOTE_CHAIN_ID, MOCK_REMOTE_EXEC_COST)
            ),
            abi.encode(MOCK_LOCAL_EXEC_COST)
        );
        // Mock: airdrop conversion to local units
        vm.mockCall(
            address(gasOracle),
            abi.encodeCall(SynapseGasOracleMock.convertRemoteValueToLocalUnits, (REMOTE_CHAIN_ID, MOCK_GAS_AIRDROP)),
            abi.encode(MOCK_LOCAL_AIRDROP_COST)
        );
    }

    function getMockOptionsV2() internal view returns (bytes memory) {
        bytes memory optionsV1 = optionsNoAirdrop.encodeOptionsV1();
        (, bytes memory data) = OptionsLib.decodeVersionedOptions(optionsV1);
        return OptionsLib.encodeVersionedOptions(2, data);
    }

    function expectExecutionFeeNotHighEnoughError() internal {
        // TODO: custom error
        vm.expectRevert("ExecutionService: execution fee is not high enough");
    }

    function expectUnsupportedOptionsVersionError() internal {
        // TODO: custom error
        vm.expectRevert("Unsupported options version: version must be OPTIONS_V1");
    }

    function expectCallerNotInterchainClientError() internal {
        // TODO: custom error
        vm.expectRevert("ExecutionService: caller is not the InterchainClient");
    }

    function test_getExecutionFee_noAirdrop() public {
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_NO_AIRDROP);
    }

    function test_getExecutionFee_withAirdrop() public {
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_WITH_AIRDROP);
    }

    function test_getExecutionFee_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectUnsupportedOptionsVersionError();
        service.getExecutionFee({dstChainId: REMOTE_CHAIN_ID, txPayloadSize: MOCK_CALLDATA_SIZE, options: optionsV2});
    }

    function test_requestExecution_noAirdrop_exactFee() public {
        vm.expectEmit(address(service));
        emit ExecutionRequested(MOCK_TX_ID, icClient);
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
    }

    function test_requestExecution_noAirdrop_higherFee() public {
        vm.expectEmit(address(service));
        emit ExecutionRequested(MOCK_TX_ID, icClient);
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP + 1,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
    }

    function test_requestExecution_noAirdrop_lowerFee() public {
        expectExecutionFeeNotHighEnoughError();
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP - 1,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
    }

    function test_requestExecution_noAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectUnsupportedOptionsVersionError();
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: optionsV2
        });
    }

    function test_requestExecution_noAirdrop_revert_callerNotInterchainClient(address caller) public {
        vm.assume(caller != icClient);
        expectCallerNotInterchainClientError();
        vm.prank(caller);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
    }

    function test_requestExecution_withAirdrop_exactFee() public {
        vm.expectEmit(address(service));
        emit ExecutionRequested(MOCK_TX_ID, icClient);
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
    }

    function test_requestExecution_withAirdrop_higherFee() public {
        vm.expectEmit(address(service));
        emit ExecutionRequested(MOCK_TX_ID, icClient);
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP + 1,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
    }

    function test_requestExecution_withAirdrop_lowerFee() public {
        expectExecutionFeeNotHighEnoughError();
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP - 1,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
    }

    function test_requestExecution_withAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectUnsupportedOptionsVersionError();
        vm.prank(icClient);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: optionsV2
        });
    }

    function test_requestExecution_withAirdrop_revert_callerNotInterchainClient(address caller) public {
        vm.assume(caller != icClient);
        expectCallerNotInterchainClientError();
        vm.prank(caller);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
    }
}
