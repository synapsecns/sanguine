// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {AbstractICApp} from "../../contracts/apps/AbstractICApp.sol";
import {AbstractICAppEvents} from "../../contracts/events/AbstractICAppEvents.sol";
import {InterchainAppV1Events} from "../../contracts/events/InterchainAppV1Events.sol";
import {IInterchainAppV1} from "../../contracts/interfaces/IInterchainAppV1.sol";
import {AppConfigV1} from "../../contracts/libs/AppConfig.sol";
import {TypeCasts} from "../../contracts/libs/TypeCasts.sol";

import {IInterchainAppV1Harness} from "../interfaces/IInterchainAppV1Harness.sol";
import {InterchainClientV1Mock} from "../mocks/InterchainClientV1Mock.sol";

import {Test} from "forge-std/Test.sol";

abstract contract InterchainAppV1Test is Test, AbstractICAppEvents, InterchainAppV1Events {
    bytes32 public constant IC_GOVERNOR_ROLE = keccak256("IC_GOVERNOR_ROLE");

    uint256 public constant LOCAL_CHAIN_ID = 1337;
    uint256 public constant REMOTE_CHAIN_ID = 7331;
    uint256 public constant UNKNOWN_CHAIN_ID = 420;
    uint256 public constant APP_OPTIMISTIC_PERIOD = 10 minutes;

    IInterchainAppV1Harness public appHarness;
    address public icClient;

    address public governor = makeAddr("Governor");
    address public moduleMock = makeAddr("Module Mock");
    address public execServiceMock = makeAddr("Execution Service Mock");
    address public linkedAppMock = makeAddr("Linked App Mock");
    bytes32 public linkedAppMockBytes32 = TypeCasts.addressToBytes32(linkedAppMock);

    AppConfigV1 public appConfig = AppConfigV1({requiredResponses: 1, optimisticPeriod: APP_OPTIMISTIC_PERIOD});

    function setUp() public virtual {
        vm.chainId(LOCAL_CHAIN_ID);
        appHarness = deployICAppV1();
        icClient = address(new InterchainClientV1Mock());
        vm.startPrank(governor);
        configureICAppV1();
        vm.stopPrank();
    }

    /// @dev This should deploy the Interchain App V1 contract and give `governor`
    /// privileges to setup its interchain configuration.
    function deployICAppV1() internal virtual returns (IInterchainAppV1Harness app);

    function configureICAppV1() internal virtual {}

    function assertEq(AppConfigV1 memory config, AppConfigV1 memory expected) internal {
        assertEq(config.requiredResponses, expected.requiredResponses);
        assertEq(config.optimisticPeriod, expected.optimisticPeriod);
    }

    function expectEventInterchainClientAdded(address client) internal {
        vm.expectEmit(address(appHarness));
        emit InterchainClientAdded(client);
    }

    function expectEventInterchainClientRemoved(address client) internal {
        vm.expectEmit(address(appHarness));
        emit InterchainClientRemoved(client);
    }

    function expectEventLatestClientSet(address client) internal {
        vm.expectEmit(address(appHarness));
        emit LatestClientSet(client);
    }

    function expectEventAppConfigV1Set(AppConfigV1 memory config) internal {
        vm.expectEmit(address(appHarness));
        emit AppConfigV1Set(config.requiredResponses, config.optimisticPeriod);
    }

    function expectEventAppLinked(uint256 chainId, bytes32 remoteApp) internal {
        vm.expectEmit(address(appHarness));
        emit AppLinked(chainId, remoteApp);
    }

    function expectEventExecutionServiceSet(address service) internal {
        vm.expectEmit(address(appHarness));
        emit ExecutionServiceSet(service);
    }

    function expectEventTrustedModuleAdded(address module) internal {
        vm.expectEmit(address(appHarness));
        emit TrustedModuleAdded(module);
    }

    function expectEventTrustedModuleRemoved(address module) internal {
        vm.expectEmit(address(appHarness));
        emit TrustedModuleRemoved(module);
    }

    function expectRevertAppZeroAddress() internal {
        vm.expectRevert(IInterchainAppV1.InterchainApp__AppZeroAddress.selector);
    }

    function expectRevertInvalidAppConfig(AppConfigV1 memory config) internal {
        vm.expectRevert(
            abi.encodeWithSelector(
                IInterchainAppV1.InterchainApp__InvalidAppConfig.selector,
                config.requiredResponses,
                config.optimisticPeriod
            )
        );
    }

    function expectRevertModuleAlreadyAdded(address module) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainAppV1.InterchainApp__ModuleAlreadyAdded.selector, module));
    }

    function expectRevertModuleNotAdded(address module) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainAppV1.InterchainApp__ModuleNotAdded.selector, module));
    }

    function expectRevertModuleZeroAddress() internal {
        vm.expectRevert(IInterchainAppV1.InterchainApp__ModuleZeroAddress.selector);
    }

    function expectRevertNotEVMLinkedApp(bytes32 linkedApp) internal {
        vm.expectRevert(abi.encodeWithSelector(IInterchainAppV1.InterchainApp__NotEVMLinkedApp.selector, linkedApp));
    }

    function expectRevertAlreadyLatestClient(address client) internal {
        vm.expectRevert(abi.encodeWithSelector(AbstractICApp.InterchainApp__AlreadyLatestClient.selector, client));
    }

    function expectRevertBalanceTooLow(uint256 actual, uint256 required) internal {
        vm.expectRevert(abi.encodeWithSelector(AbstractICApp.InterchainApp__BalanceTooLow.selector, actual, required));
    }

    function expectRevertClientAlreadyAdded(address client) internal {
        vm.expectRevert(abi.encodeWithSelector(AbstractICApp.InterchainApp__ClientAlreadyAdded.selector, client));
    }

    function expectRevertInterchainClientZeroAddress() internal {
        vm.expectRevert(AbstractICApp.InterchainApp__InterchainClientZeroAddress.selector);
    }

    function expectRevertNotInterchainClient(address account) internal {
        vm.expectRevert(abi.encodeWithSelector(AbstractICApp.InterchainApp__NotInterchainClient.selector, account));
    }

    function expectRevertReceiverNotSet(uint256 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(AbstractICApp.InterchainApp__ReceiverNotSet.selector, chainId));
    }

    function expectRevertSameChainId(uint256 chainId) internal {
        vm.expectRevert(abi.encodeWithSelector(AbstractICApp.InterchainApp__SameChainId.selector, chainId));
    }

    function expectRevertSenderNotAllowed(uint256 srcChainId, bytes32 sender) internal {
        vm.expectRevert(
            abi.encodeWithSelector(AbstractICApp.InterchainApp__SenderNotAllowed.selector, srcChainId, sender)
        );
    }

    function toArray(address a) internal pure returns (address[] memory arr) {
        arr = new address[](1);
        arr[0] = a;
    }

    function toArray(address a, address b) internal pure returns (address[] memory arr) {
        arr = new address[](2);
        arr[0] = a;
        arr[1] = b;
    }
}
