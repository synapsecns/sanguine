// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsV1} from "../../contracts/libs/Options.sol";
import {VersionedPayloadLib} from "../../contracts/libs/VersionedPayload.sol";

import {SynapseExecutionServiceV1, SynapseExecutionServiceV1Test} from "./SynapseExecutionServiceV1.t.sol";

import {SynapseGasOracleMock} from "../mocks/SynapseGasOracleMock.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseExecutionServiceV1ExecutionTest is SynapseExecutionServiceV1Test {
    bytes32 public constant MOCK_TX_ID = keccak256("mock-tx-id");
    uint64 public constant LOCAL_CHAIN_ID = 1337;
    uint64 public constant REMOTE_CHAIN_ID = 7331;

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

    // 10% markup
    uint256 public constant MOCK_MARKUP = 0.1e18;
    uint256 public constant MOCK_GAS_AIRDROP_MARKUP = 0.011 ether;
    uint256 public constant MOCK_REMOTE_EXEC_COST_MARKUP = 0.11 ether;
    uint256 public constant MOCK_LOCAL_EXEC_COST_MARKUP = MOCK_REMOTE_EXEC_COST_MARKUP * 2;
    uint256 public constant MOCK_LOCAL_AIRDROP_COST_MARKUP = MOCK_GAS_AIRDROP_MARKUP * 2;

    uint256 public constant MOCK_FEE_NO_AIRDROP_MARKUP = MOCK_LOCAL_EXEC_COST_MARKUP;
    uint256 public constant MOCK_FEE_WITH_AIRDROP_MARKUP = MOCK_LOCAL_EXEC_COST_MARKUP + MOCK_LOCAL_AIRDROP_COST_MARKUP;

    OptionsV1 public optionsNoAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: 0});
    OptionsV1 public optionsWithAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: MOCK_GAS_AIRDROP});
    bytes public encodedOptionsNoAirdrop = optionsNoAirdrop.encodeOptionsV1();
    bytes public encodedOptionsWithAirdrop = optionsWithAirdrop.encodeOptionsV1();

    address public icClientA = makeAddr("ICClientA");
    address public icClientB = makeAddr("ICClientB");
    address public executorEOA = makeAddr("ExecutorEOA");
    address public gasOracle;

    function setUp() public override {
        super.setUp();
        gasOracle = address(new SynapseGasOracleMock());
        configureService();
        mockGasPricing();
    }

    function configureService() internal {
        service.initialize(address(this));
        service.grantRole(GOVERNOR_ROLE, address(this));
        service.grantRole(IC_CLIENT_ROLE, icClientA);
        service.grantRole(IC_CLIENT_ROLE, icClientB);
        service.setExecutorEOA(executorEOA);
        service.setGasOracle(gasOracle);
    }

    function mockGasPricing() internal {
        // Mock: execution cost in local units
        vm.mockCall(
            gasOracle,
            abi.encodeCall(
                SynapseGasOracleMock.estimateTxCostInLocalUnits, (REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE)
            ),
            abi.encode(MOCK_LOCAL_EXEC_COST)
        );
        // Mock: execution cost in remote units + its conversion to local units
        vm.mockCall(
            gasOracle,
            abi.encodeCall(
                SynapseGasOracleMock.estimateTxCostInRemoteUnits, (REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE)
            ),
            abi.encode(MOCK_REMOTE_EXEC_COST)
        );
        vm.mockCall(
            gasOracle,
            abi.encodeCall(
                SynapseGasOracleMock.convertRemoteValueToLocalUnits, (REMOTE_CHAIN_ID, MOCK_REMOTE_EXEC_COST)
            ),
            abi.encode(MOCK_LOCAL_EXEC_COST)
        );
        // Mock: airdrop conversion to local units
        vm.mockCall(
            gasOracle,
            abi.encodeCall(SynapseGasOracleMock.convertRemoteValueToLocalUnits, (REMOTE_CHAIN_ID, MOCK_GAS_AIRDROP)),
            abi.encode(MOCK_LOCAL_AIRDROP_COST)
        );
    }

    function getMockOptionsV2() internal view returns (bytes memory) {
        bytes memory optionsV1 = optionsNoAirdrop.encodeOptionsV1();
        bytes memory data = VersionedPayloadLib.getPayloadFromMemory(optionsV1);
        return VersionedPayloadLib.encodeVersionedPayload(2, data);
    }

    function requestTxExecution(address caller, uint256 executionFee, bytes memory options) internal {
        deal(caller, executionFee);
        vm.prank(caller);
        service.requestTxExecution{value: executionFee}({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            options: options
        });
    }

    function test_claimFees() public {
        address caller = makeAddr("Random Caller");
        uint256 amount = 1 ether;
        deal(address(service), amount);
        vm.expectEmit(address(service));
        emit FeesClaimed(executorEOA, amount);
        vm.prank(caller);
        service.claimFees();
        assertEq(executorEOA.balance, amount);
    }

    function test_claimFees_revert_zeroAmount() public {
        address caller = makeAddr("Random Caller");
        expectRevertZeroAmount();
        vm.prank(caller);
        service.claimFees();
    }

    function test_claimFees_revert_zeroExecutorEOA() public {
        SynapseExecutionServiceV1 freshService = SynapseExecutionServiceV1(deployProxy(address(implementation)));
        address caller = makeAddr("Random Caller");
        deal(address(freshService), 1 ether);
        expectRevertZeroAddress();
        vm.prank(caller);
        freshService.claimFees();
    }

    function test_getExecutionFee_noAirdrop() public {
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_NO_AIRDROP);
    }

    function test_getExecutionFee_noAirdrop_withMarkup() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_NO_AIRDROP_MARKUP);
    }

    function test_getExecutionFee_withAirdrop() public {
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_WITH_AIRDROP);
    }

    function test_getExecutionFee_withAirdrop_withMarkup() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_WITH_AIRDROP_MARKUP);
    }

    function test_getExecutionFee_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        service.getExecutionFee({dstChainId: REMOTE_CHAIN_ID, txPayloadSize: MOCK_CALLDATA_SIZE, options: optionsV2});
    }

    function test_requestTxExecution_clientA_noAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_NO_AIRDROP);
        requestTxExecution(icClientA, MOCK_FEE_NO_AIRDROP, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientA_noAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_NO_AIRDROP + 1);
        requestTxExecution(icClientA, MOCK_FEE_NO_AIRDROP + 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientA_noAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_NO_AIRDROP - 1, MOCK_FEE_NO_AIRDROP);
        requestTxExecution(icClientA, MOCK_FEE_NO_AIRDROP - 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientA_noAirdrop_withMarkup_exactFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_NO_AIRDROP_MARKUP);
        requestTxExecution(icClientA, MOCK_FEE_NO_AIRDROP_MARKUP, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientA_noAirdrop_withMarkup_higherFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_NO_AIRDROP_MARKUP + 1);
        requestTxExecution(icClientA, MOCK_FEE_NO_AIRDROP_MARKUP + 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientA_noAirdrop_withMarkup_revert_lowerFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectRevertFeeAmountTooLow(MOCK_FEE_NO_AIRDROP_MARKUP - 1, MOCK_FEE_NO_AIRDROP_MARKUP);
        requestTxExecution(icClientA, MOCK_FEE_NO_AIRDROP_MARKUP - 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientA_noAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        requestTxExecution(icClientA, MOCK_FEE_NO_AIRDROP, optionsV2);
    }

    function test_requestTxExecution_clientA_withAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_WITH_AIRDROP);
        requestTxExecution(icClientA, MOCK_FEE_WITH_AIRDROP, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientA_withAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_WITH_AIRDROP + 1);
        requestTxExecution(icClientA, MOCK_FEE_WITH_AIRDROP + 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientA_withAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_WITH_AIRDROP - 1, MOCK_FEE_WITH_AIRDROP);
        requestTxExecution(icClientA, MOCK_FEE_WITH_AIRDROP - 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientA_withAirdrop_withMarkup_exactFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_WITH_AIRDROP_MARKUP);
        requestTxExecution(icClientA, MOCK_FEE_WITH_AIRDROP_MARKUP, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientA_withAirdrop_withMarkup_higherFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientA, MOCK_FEE_WITH_AIRDROP_MARKUP + 1);
        requestTxExecution(icClientA, MOCK_FEE_WITH_AIRDROP_MARKUP + 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientA_withAirdrop_withMarkup_revert_lowerFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectRevertFeeAmountTooLow(MOCK_FEE_WITH_AIRDROP_MARKUP - 1, MOCK_FEE_WITH_AIRDROP_MARKUP);
        requestTxExecution(icClientA, MOCK_FEE_WITH_AIRDROP_MARKUP - 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientA_withAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        requestTxExecution(icClientA, MOCK_FEE_WITH_AIRDROP, optionsV2);
    }

    function test_requestTxExecution_clientB_noAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_NO_AIRDROP);
        requestTxExecution(icClientB, MOCK_FEE_NO_AIRDROP, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientB_noAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_NO_AIRDROP + 1);
        requestTxExecution(icClientB, MOCK_FEE_NO_AIRDROP + 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientB_noAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_NO_AIRDROP - 1, MOCK_FEE_NO_AIRDROP);
        requestTxExecution(icClientB, MOCK_FEE_NO_AIRDROP - 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientB_noAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        requestTxExecution(icClientB, MOCK_FEE_NO_AIRDROP, optionsV2);
    }

    function test_requestTxExecution_clientB_noAirdrop_withMarkup_exactFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_NO_AIRDROP_MARKUP);
        requestTxExecution(icClientB, MOCK_FEE_NO_AIRDROP_MARKUP, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientB_noAirdrop_withMarkup_higherFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_NO_AIRDROP_MARKUP + 1);
        requestTxExecution(icClientB, MOCK_FEE_NO_AIRDROP_MARKUP + 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientB_noAirdrop_withMarkup_revert_lowerFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectRevertFeeAmountTooLow(MOCK_FEE_NO_AIRDROP_MARKUP - 1, MOCK_FEE_NO_AIRDROP_MARKUP);
        requestTxExecution(icClientB, MOCK_FEE_NO_AIRDROP_MARKUP - 1, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_clientB_withAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_WITH_AIRDROP);
        requestTxExecution(icClientB, MOCK_FEE_WITH_AIRDROP, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientB_withAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_WITH_AIRDROP + 1);
        requestTxExecution(icClientB, MOCK_FEE_WITH_AIRDROP + 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientB_withAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_WITH_AIRDROP - 1, MOCK_FEE_WITH_AIRDROP);
        requestTxExecution(icClientB, MOCK_FEE_WITH_AIRDROP - 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientB_withAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        requestTxExecution(icClientB, MOCK_FEE_WITH_AIRDROP, optionsV2);
    }

    function test_requestTxExecution_clientB_withAirdrop_withMarkup_exactFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_WITH_AIRDROP_MARKUP);
        requestTxExecution(icClientB, MOCK_FEE_WITH_AIRDROP_MARKUP, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientB_withAirdrop_withMarkup_higherFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectEventExecutionRequested(MOCK_TX_ID, icClientB, MOCK_FEE_WITH_AIRDROP_MARKUP + 1);
        requestTxExecution(icClientB, MOCK_FEE_WITH_AIRDROP_MARKUP + 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_clientB_withAirdrop_withMarkup_revert_lowerFee() public {
        service.setGlobalMarkup(MOCK_MARKUP);
        expectRevertFeeAmountTooLow(MOCK_FEE_WITH_AIRDROP_MARKUP - 1, MOCK_FEE_WITH_AIRDROP_MARKUP);
        requestTxExecution(icClientB, MOCK_FEE_WITH_AIRDROP_MARKUP - 1, encodedOptionsWithAirdrop);
    }

    function test_requestTxExecution_noAirdrop_revert_notInterchainClient(address caller) public {
        assumeNotProxyAdmin({target: address(service), caller: caller});
        vm.assume(caller != icClientA && caller != icClientB);
        expectRevertNotInterchainClient(caller);
        requestTxExecution(caller, MOCK_FEE_NO_AIRDROP, encodedOptionsNoAirdrop);
    }

    function test_requestTxExecution_withAirdrop_revert_notInterchainClient(address caller) public {
        assumeNotProxyAdmin({target: address(service), caller: caller});
        vm.assume(caller != icClientA && caller != icClientB);
        expectRevertNotInterchainClient(caller);
        requestTxExecution(caller, MOCK_FEE_WITH_AIRDROP, encodedOptionsWithAirdrop);
    }
}
