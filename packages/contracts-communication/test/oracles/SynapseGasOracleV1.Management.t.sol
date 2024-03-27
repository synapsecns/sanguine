// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseGasOracleV1Events} from "../../contracts/events/SynapseGasOracleV1Events.sol";
import {SynapseGasOracleV1, ISynapseGasOracleV1} from "../../contracts/oracles/SynapseGasOracleV1.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseGasOracleV1ManagementTest is Test, SynapseGasOracleV1Events {
    uint256 public constant LOCAL_CHAIN_ID = 1337;
    uint256 public constant REMOTE_CHAIN_ID = 7331;

    uint256 public constant MOCK_CALLDATA_PRICE = 1000;
    uint256 public constant MOCK_GAS_PRICE = 2000;
    uint256 public constant MOCK_NATIVE_PRICE = 3000;

    ISynapseGasOracleV1.RemoteGasData public mockGasData = ISynapseGasOracleV1.RemoteGasData({
        calldataPrice: MOCK_CALLDATA_PRICE,
        gasPrice: MOCK_GAS_PRICE,
        nativePrice: MOCK_NATIVE_PRICE
    });

    SynapseGasOracleV1 public oracle;

    address public owner = makeAddr("Owner");

    modifier withRemoteMockGasData() {
        vm.prank(owner);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
        _;
    }

    modifier withLocalMockNativePrice() {
        vm.prank(owner);
        oracle.setLocalNativePrice(MOCK_NATIVE_PRICE);
        _;
    }

    function setUp() public {
        vm.chainId(LOCAL_CHAIN_ID);
        oracle = new SynapseGasOracleV1(owner);
    }

    function assertEq(
        ISynapseGasOracleV1.RemoteGasData memory a,
        ISynapseGasOracleV1.RemoteGasData memory b
    )
        internal
    {
        assertEq(a.calldataPrice, b.calldataPrice);
        assertEq(a.gasPrice, b.gasPrice);
        assertEq(a.nativePrice, b.nativePrice);
    }

    function expectEventCalldataPriceSet(uint256 chainId, uint256 calldataPrice) internal {
        vm.expectEmit(address(oracle));
        emit CalldataPriceSet(chainId, calldataPrice);
    }

    function expectEventGasPriceSet(uint256 chainId, uint256 gasPrice) internal {
        vm.expectEmit(address(oracle));
        emit GasPriceSet(chainId, gasPrice);
    }

    function expectEventNativePriceSet(uint256 chainId, uint256 nativePrice) internal {
        vm.expectEmit(address(oracle));
        emit NativePriceSet(chainId, nativePrice);
    }

    function expectRevertOwnableUnauthorizedAccount(address caller) internal {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, caller));
    }

    function expectRevertNotRemoteChainId(uint256 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(ISynapseGasOracleV1.SynapseGasOracleV1__NotRemoteChainId.selector, chainId)
        );
    }

    function expectRevertNativePriceNotSet(uint256 chainId) internal {
        vm.expectRevert(
            abi.encodeWithSelector(ISynapseGasOracleV1.SynapseGasOracleV1__NativePriceNotSet.selector, chainId)
        );
    }

    function expectRevertNativePriceZero() internal {
        vm.expectRevert(ISynapseGasOracleV1.SynapseGasOracleV1__NativePriceZero.selector);
    }

    function test_constructor() public {
        assertEq(oracle.owner(), owner);
    }

    function test_setLocalNativePrice_emitsEvent() public {
        expectEventNativePriceSet(LOCAL_CHAIN_ID, MOCK_NATIVE_PRICE);
        vm.prank(owner);
        oracle.setLocalNativePrice(MOCK_NATIVE_PRICE);
    }

    function test_setLocalNativePrice_setsPrice() public {
        vm.prank(owner);
        oracle.setLocalNativePrice(MOCK_NATIVE_PRICE);
        assertEq(oracle.getLocalNativePrice(), MOCK_NATIVE_PRICE);
    }

    function test_setLocalNativePrice_sameValue() public withLocalMockNativePrice {
        vm.recordLogs();
        vm.prank(owner);
        oracle.setLocalNativePrice(MOCK_NATIVE_PRICE);
        assertEq(vm.getRecordedLogs().length, 0);
        assertEq(oracle.getLocalNativePrice(), MOCK_NATIVE_PRICE);
    }

    function test_setLocalNativePrice_revert_callerNotOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        oracle.setLocalNativePrice(MOCK_NATIVE_PRICE);
    }

    function test_setLocalNativePrice_revert_priceZero() public {
        expectRevertNativePriceZero();
        vm.prank(owner);
        oracle.setLocalNativePrice(0);
    }

    function test_setRemoteCallDataPrice_emitsEvent() public withRemoteMockGasData {
        expectEventCalldataPriceSet(REMOTE_CHAIN_ID, MOCK_CALLDATA_PRICE + 1);
        vm.prank(owner);
        oracle.setRemoteCallDataPrice(REMOTE_CHAIN_ID, MOCK_CALLDATA_PRICE + 1);
    }

    function test_setRemoteCallDataPrice_setsPrice() public withRemoteMockGasData {
        vm.prank(owner);
        oracle.setRemoteCallDataPrice(REMOTE_CHAIN_ID, MOCK_CALLDATA_PRICE + 1);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID).calldataPrice, MOCK_CALLDATA_PRICE + 1);
    }

    function test_setRemoteCallDataPrice_sameValue() public withRemoteMockGasData {
        vm.recordLogs();
        vm.prank(owner);
        oracle.setRemoteCallDataPrice(REMOTE_CHAIN_ID, MOCK_CALLDATA_PRICE);
        assertEq(vm.getRecordedLogs().length, 0);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID).calldataPrice, MOCK_CALLDATA_PRICE);
    }

    function test_setRemoteCallDataPrice_revert_callerNotOwner(address caller) public withRemoteMockGasData {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        oracle.setRemoteCallDataPrice(REMOTE_CHAIN_ID, MOCK_CALLDATA_PRICE);
    }

    function test_setRemoteCallDataPrice_revert_notRemoteChainId() public withRemoteMockGasData {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        vm.prank(owner);
        oracle.setRemoteCallDataPrice(LOCAL_CHAIN_ID, MOCK_CALLDATA_PRICE);
    }

    function test_setRemoteCallDataPrice_revert_nativePriceNotSet() public {
        // Setting calldata price without setting native token price is not allowed.
        expectRevertNativePriceNotSet(REMOTE_CHAIN_ID);
        vm.prank(owner);
        oracle.setRemoteCallDataPrice(REMOTE_CHAIN_ID, MOCK_CALLDATA_PRICE);
    }

    function test_setRemoteGasPrice_emitsEvent() public withRemoteMockGasData {
        expectEventGasPriceSet(REMOTE_CHAIN_ID, MOCK_GAS_PRICE + 1);
        vm.prank(owner);
        oracle.setRemoteGasPrice(REMOTE_CHAIN_ID, MOCK_GAS_PRICE + 1);
    }

    function test_setRemoteGasPrice_setsPrice() public withRemoteMockGasData {
        vm.prank(owner);
        oracle.setRemoteGasPrice(REMOTE_CHAIN_ID, MOCK_GAS_PRICE + 1);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID).gasPrice, MOCK_GAS_PRICE + 1);
    }

    function test_setRemoteGasPrice_sameValue() public withRemoteMockGasData {
        vm.recordLogs();
        vm.prank(owner);
        oracle.setRemoteGasPrice(REMOTE_CHAIN_ID, MOCK_GAS_PRICE);
        assertEq(vm.getRecordedLogs().length, 0);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID).gasPrice, MOCK_GAS_PRICE);
    }

    function test_setRemoteGasPrice_revert_callerNotOwner(address caller) public withRemoteMockGasData {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        oracle.setRemoteGasPrice(REMOTE_CHAIN_ID, MOCK_GAS_PRICE);
    }

    function test_setRemoteGasPrice_revert_notRemoteChainId() public withRemoteMockGasData {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        vm.prank(owner);
        oracle.setRemoteGasPrice(LOCAL_CHAIN_ID, MOCK_GAS_PRICE);
    }

    function test_setRemoteGasPrice_revert_nativePriceNotSet() public {
        // Setting gas price without setting native token price is not allowed.
        expectRevertNativePriceNotSet(REMOTE_CHAIN_ID);
        vm.prank(owner);
        oracle.setRemoteGasPrice(REMOTE_CHAIN_ID, MOCK_GAS_PRICE);
    }

    function test_setRemoteNativePrice_emitsEvent() public {
        expectEventNativePriceSet(REMOTE_CHAIN_ID, MOCK_NATIVE_PRICE);
        vm.prank(owner);
        oracle.setRemoteNativePrice(REMOTE_CHAIN_ID, MOCK_NATIVE_PRICE);
    }

    function test_setRemoteNativePrice_setsPrice() public {
        vm.prank(owner);
        oracle.setRemoteNativePrice(REMOTE_CHAIN_ID, MOCK_NATIVE_PRICE);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID).nativePrice, MOCK_NATIVE_PRICE);
    }

    function test_setRemoteNativePrice_sameValue() public withRemoteMockGasData {
        vm.recordLogs();
        vm.prank(owner);
        oracle.setRemoteNativePrice(REMOTE_CHAIN_ID, MOCK_NATIVE_PRICE);
        assertEq(vm.getRecordedLogs().length, 0);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID).nativePrice, MOCK_NATIVE_PRICE);
    }

    function test_setRemoteNativePrice_revert_callerNotOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        oracle.setRemoteNativePrice(REMOTE_CHAIN_ID, MOCK_NATIVE_PRICE);
    }

    function test_setRemoteNativePrice_revert_notRemoteChainId() public {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        vm.prank(owner);
        oracle.setRemoteNativePrice(LOCAL_CHAIN_ID, MOCK_NATIVE_PRICE);
    }

    function test_setRemoteNativePrice_revert_priceZero() public {
        expectRevertNativePriceZero();
        vm.prank(owner);
        oracle.setRemoteNativePrice(REMOTE_CHAIN_ID, 0);
    }

    function test_setRemoteGasData_emitsEvents() public {
        expectEventCalldataPriceSet(REMOTE_CHAIN_ID, MOCK_CALLDATA_PRICE);
        expectEventGasPriceSet(REMOTE_CHAIN_ID, MOCK_GAS_PRICE);
        expectEventNativePriceSet(REMOTE_CHAIN_ID, MOCK_NATIVE_PRICE);
        vm.prank(owner);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
    }

    function test_setRemoteGasData_setsData() public {
        vm.prank(owner);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID), mockGasData);
    }

    function test_setRemoteGasData_allValuesSame() public withRemoteMockGasData {
        vm.recordLogs();
        vm.prank(owner);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
        assertEq(vm.getRecordedLogs().length, 0);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID), mockGasData);
    }

    function test_setRemoteGasData_oneDiffValue() public withRemoteMockGasData {
        mockGasData.calldataPrice += 1;
        expectEventCalldataPriceSet(REMOTE_CHAIN_ID, mockGasData.calldataPrice);
        vm.recordLogs();
        vm.prank(owner);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID), mockGasData);
    }

    function test_setRemoteGasData_twoDiffValues() public withRemoteMockGasData {
        mockGasData.calldataPrice += 1;
        mockGasData.gasPrice += 1;
        expectEventCalldataPriceSet(REMOTE_CHAIN_ID, mockGasData.calldataPrice);
        expectEventGasPriceSet(REMOTE_CHAIN_ID, mockGasData.gasPrice);
        vm.recordLogs();
        vm.prank(owner);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
        assertEq(vm.getRecordedLogs().length, 2);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID), mockGasData);
    }

    function test_setRemoteGasData_revert_callerNotOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertOwnableUnauthorizedAccount(caller);
        vm.prank(caller);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
    }

    function test_setRemoteGasData_revert_notRemoteChainId() public {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        vm.prank(owner);
        oracle.setRemoteGasData(LOCAL_CHAIN_ID, mockGasData);
    }

    function test_setRemoteGasData_revert_priceZero() public {
        mockGasData.nativePrice = 0;
        expectRevertNativePriceZero();
        vm.prank(owner);
        oracle.setRemoteGasData(REMOTE_CHAIN_ID, mockGasData);
    }

    function test_receiveRemoteGasData_noop() public withRemoteMockGasData {
        vm.recordLogs();
        oracle.receiveRemoteGasData(REMOTE_CHAIN_ID, "");
        assertEq(vm.getRecordedLogs().length, 0);
        assertEq(oracle.getRemoteGasData(REMOTE_CHAIN_ID), mockGasData);
    }

    function test_getLocalGasData_noop() public withRemoteMockGasData withLocalMockNativePrice {
        assertEq(oracle.getLocalGasData(), "");
    }
}
