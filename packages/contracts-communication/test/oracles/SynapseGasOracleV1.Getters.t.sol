// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SynapseGasOracleV1, ISynapseGasOracleV1} from "../../contracts/oracles/SynapseGasOracleV1.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseGasOracleV1GettersTest is Test {
    uint256 public constant LOCAL_CHAIN_ID = 1337;
    uint256 public constant REMOTE_CHAIN_ID = 7331;

    // Per byte of calldata
    uint256 public constant REMOTE_CALLDATA_PRICE = 4844 gwei;
    uint256 public constant REMOTE_GAS_PRICE = 1 gwei;
    // Remote native token is worth 2 mainnet ETH
    uint256 public constant REMOTE_NATIVE_PRICE = 2 ether;

    // Local native token is worth 0.5 mainnet ETH
    uint256 public constant LOCAL_NATIVE_PRICE = 0.5 ether;

    uint256 public constant MOCK_GAS_LIMIT = 100_000;
    uint256 public constant MOCK_CALLDATA_SIZE = 100;

    uint256 public constant REMOTE_TX_COST =
        MOCK_GAS_LIMIT * REMOTE_GAS_PRICE + MOCK_CALLDATA_SIZE * REMOTE_CALLDATA_PRICE;

    SynapseGasOracleV1 public oracle;

    modifier withLocalSetup() {
        oracle.setLocalNativePrice(LOCAL_NATIVE_PRICE);
        _;
    }

    modifier withRemoteSetup() {
        oracle.setRemoteGasData(
            REMOTE_CHAIN_ID,
            ISynapseGasOracleV1.RemoteGasData({
                calldataPrice: REMOTE_CALLDATA_PRICE,
                gasPrice: REMOTE_GAS_PRICE,
                nativePrice: REMOTE_NATIVE_PRICE
            })
        );
        _;
    }

    function setUp() public {
        vm.chainId(LOCAL_CHAIN_ID);
        oracle = new SynapseGasOracleV1(address(this));
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

    function test_convertRemoteValueToLocalUnits() public withLocalSetup withRemoteSetup {
        // Remote native token is worth 4x local native token
        assertEq(oracle.convertRemoteValueToLocalUnits(REMOTE_CHAIN_ID, 2 gwei), 8 gwei);
        assertEq(oracle.convertRemoteValueToLocalUnits(REMOTE_CHAIN_ID, 1 ether), 4 ether);
        assertEq(oracle.convertRemoteValueToLocalUnits(REMOTE_CHAIN_ID, 3 ether), 12 ether);
    }

    function test_convertRemoteValueToLocalUnits_revert_notRemoteChainId() public withLocalSetup withRemoteSetup {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        oracle.convertRemoteValueToLocalUnits(LOCAL_CHAIN_ID, 1 ether);
    }

    function test_convertRemoteValueToLocalUnits_revert_localNativePriceNotSet() public withRemoteSetup {
        expectRevertNativePriceNotSet(LOCAL_CHAIN_ID);
        oracle.convertRemoteValueToLocalUnits(REMOTE_CHAIN_ID, 1 ether);
    }

    function test_convertRemoteValueToLocalUnits_revert_remoteNativePriceNotSet() public withLocalSetup {
        expectRevertNativePriceNotSet(REMOTE_CHAIN_ID);
        oracle.convertRemoteValueToLocalUnits(REMOTE_CHAIN_ID, 1 ether);
    }

    function test_estimateTxCostInRemoteUnits() public withRemoteSetup {
        assertEq(
            oracle.estimateTxCostInRemoteUnits(REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE), REMOTE_TX_COST
        );
    }

    function test_estimateTxCostInRemoteUnits_revert_notRemoteChainId() public withLocalSetup withRemoteSetup {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        oracle.estimateTxCostInRemoteUnits(LOCAL_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE);
    }

    function test_estimateTxCostInRemoteUnits_revert_remoteNativePriceNotSet() public withLocalSetup {
        expectRevertNativePriceNotSet(REMOTE_CHAIN_ID);
        oracle.estimateTxCostInRemoteUnits(REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE);
    }

    function test_estimateTxCostInLocalUnits() public withLocalSetup withRemoteSetup {
        // Remote native token is worth 4x local native token
        assertEq(
            oracle.estimateTxCostInLocalUnits(REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE), 4 * REMOTE_TX_COST
        );
    }

    function test_estimateTxCostInLocalUnits_revert_notRemoteChainId() public withLocalSetup withRemoteSetup {
        expectRevertNotRemoteChainId(LOCAL_CHAIN_ID);
        oracle.estimateTxCostInLocalUnits(LOCAL_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE);
    }

    function test_estimateTxCostInLocalUnits_revert_localNativePriceNotSet() public withRemoteSetup {
        expectRevertNativePriceNotSet(LOCAL_CHAIN_ID);
        oracle.estimateTxCostInLocalUnits(REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE);
    }

    function test_estimateTxCostInLocalUnits_revert_remoteNativePriceNotSet() public withLocalSetup {
        expectRevertNativePriceNotSet(REMOTE_CHAIN_ID);
        oracle.estimateTxCostInLocalUnits(REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE);
    }
}
