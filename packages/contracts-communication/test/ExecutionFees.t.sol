import "forge-std/Test.sol";
import "../contracts/ExecutionFees.sol";

contract ExecutionFeesTest is Test {
    ExecutionFees executionFees;
    address icClient = address(0x123);
    address executor = address(0x456);
    bytes32 transactionId = keccak256("transaction");
    uint256 dstChainId = 1;
    uint256 executionFee = 1 ether;

    function setUp() public {
        executionFees = new ExecutionFees(icClient);
    }

    function test_AddExecutionFee() public {
        vm.deal(address(this), executionFee);
        executionFees.addExecutionFee{value: executionFee}(dstChainId, transactionId);

        assertEq(executionFees.getAccumulatedRewards(executor), 0);
        assertEq(executionFees.getUnclaimedRewards(executor), 0);
    }

    function test_RecordExecutor() public {
        vm.deal(address(this), executionFee);
        executionFees.addExecutionFee{value: executionFee}(dstChainId, transactionId);

        vm.prank(address(icClient));
        executionFees.recordExecutor(dstChainId, transactionId, executor);

        assertEq(executionFees.getAccumulatedRewards(executor), executionFee);
        assertEq(executionFees.getUnclaimedRewards(executor), executionFee);
    }

    function test_ClaimExecutionFees() public {
        vm.deal(address(this), executionFee);
        executionFees.addExecutionFee{value: executionFee}(dstChainId, transactionId);

        vm.prank(address(icClient));
        executionFees.recordExecutor(dstChainId, transactionId, executor);

        uint256 beforeBalance = executor.balance;
        vm.prank(executor);
        executionFees.claimExecutionFees();
        uint256 afterBalance = executor.balance;

        assertEq(afterBalance - beforeBalance, executionFee);
        assertEq(executionFees.getUnclaimedRewards(executor), 0);
    }

    function test_AddExecutionFee_ExecutorAlreadyRecorded() public {
        vm.deal(address(this), executionFee);
        executionFees.addExecutionFee{value: executionFee}(dstChainId, transactionId);
        vm.prank(address(icClient));
        executionFees.recordExecutor(dstChainId, transactionId, executor);
        vm.deal(address(this), executionFee);
        vm.expectRevert("ExecutionFees: Executor already recorded");
        executionFees.addExecutionFee{value: executionFee}(dstChainId, transactionId);
    }

    function test_AddExecutionFee_ZeroFee() public {
        vm.expectRevert("ExecutionFees: Fee must be greater than 0");
        executionFees.addExecutionFee(dstChainId, transactionId);
    }

    function test_RecordExecutor_ExecutorAlreadyRecorded() public {
        vm.deal(address(this), executionFee);
        executionFees.addExecutionFee{value: executionFee}(dstChainId, transactionId);
        vm.startPrank(address(icClient));
        executionFees.recordExecutor(dstChainId, transactionId, executor);
        vm.expectRevert();
        executionFees.recordExecutor(dstChainId, transactionId, executor);
    }

    function test_RecordExecutor_NoExecutionFeeFound() public {
        vm.startPrank(address(icClient));
        vm.expectRevert();
        executionFees.recordExecutor(dstChainId, transactionId, executor);
    }

    function test_ClaimExecutionFees_NoUnclaimedRewards() public {
        vm.prank(executor);
        vm.expectRevert("ExecutionFees: No unclaimed rewards");
        executionFees.claimExecutionFees();
    }
}
