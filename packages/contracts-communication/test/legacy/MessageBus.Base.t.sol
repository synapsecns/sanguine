// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {MessageBusEvents} from "../../contracts/legacy/events/MessageBusEvents.sol";
import {LegacyMessageLib} from "../../contracts/legacy/libs/LegacyMessage.sol";
import {LegacyOptionsLib} from "../../contracts/legacy/libs/LegacyOptions.sol";
import {MessageBus, IMessageBus} from "../../contracts/legacy/MessageBus.sol";
import {OptionsV1} from "../../contracts/libs/Options.sol";
import {TypeCasts} from "../../contracts/libs/TypeCasts.sol";

import {LegacyMessage} from "./libs/LegacyMessage.t.sol";
import {LegacyReceiverMock} from "./mocks/LegacyReceiverMock.sol";
import {ExecutionServiceMock} from "../mocks/ExecutionServiceMock.sol";
import {InterchainClientV1Mock} from "../mocks/InterchainClientV1Mock.sol";
import {InterchainModuleMock} from "../mocks/InterchainModuleMock.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {Test} from "forge-std/Test.sol";

// solhint-disable max-states-count
// solhint-disable ordering
abstract contract MessageBusBaseTest is MessageBusEvents, Test {
    bytes32 public constant IC_GOVERNOR_ROLE = keccak256("IC_GOVERNOR_ROLE");

    uint64 public constant LOCAL_CHAIN_ID = 1337;
    uint64 public constant REMOTE_CHAIN_ID = 7331;
    uint256 public constant BUS_OPTIMISTIC_PERIOD = 1 minutes;
    uint256 public constant MOCK_FEE = 0.001 ether;
    uint256 public constant MOCK_GAS_LIMIT = 400_000;
    uint256 public constant GAS_BUFFER = 20_000;
    uint256 public constant IC_GAS_LIMIT = MOCK_GAS_LIMIT + GAS_BUFFER;

    uint64 public constant MOCK_NONCE = 42;
    bytes public constant MESSAGE = "One small step for man, one giant leap for mankind.";

    bytes public legacyOptions = LegacyOptionsLib.encodeLegacyOptions(MOCK_GAS_LIMIT);
    bytes public icOptions = OptionsV1({gasLimit: IC_GAS_LIMIT, gasAirdrop: 0}).encodeOptionsV1();

    MessageBus public messageBus;
    address public remoteMessageBus = makeAddr("RemoteMessageBus");
    bytes32 public remoteMessageBusBytes32 = TypeCasts.addressToBytes32(remoteMessageBus);
    address public icClient;
    address public srcSender;
    bytes32 public srcSenderBytes32;
    address public dstReceiver;
    bytes32 public dstReceiverBytes32;
    address public execService;
    address public icModule;
    address[] public icModules;

    address public admin = makeAddr("Admin");
    address public governor = makeAddr("Governor");

    function setUp() public virtual {
        vm.chainId(LOCAL_CHAIN_ID);
        messageBus = new MessageBus(admin);
        icClient = address(new InterchainClientV1Mock());
        icModule = address(new InterchainModuleMock());
        icModules.push(icModule);
        execService = address(new ExecutionServiceMock());
        srcSender = address(new LegacyReceiverMock());
        dstReceiver = address(new LegacyReceiverMock());
        srcSenderBytes32 = TypeCasts.addressToBytes32(srcSender);
        dstReceiverBytes32 = TypeCasts.addressToBytes32(dstReceiver);
        configureMessageBus();
    }

    function configureMessageBus() internal virtual {
        bytes32 icGovernorRole = messageBus.IC_GOVERNOR_ROLE();
        vm.prank(admin);
        messageBus.grantRole(icGovernorRole, governor);
        vm.startPrank(governor);
        messageBus.addInterchainClient({client: icClient, updateLatest: true});
        messageBus.linkRemoteAppEVM({chainId: REMOTE_CHAIN_ID, remoteApp: remoteMessageBus});
        messageBus.addTrustedModule(icModule);
        messageBus.setAppConfigV1({requiredResponses: 1, optimisticPeriod: BUS_OPTIMISTIC_PERIOD});
        messageBus.setExecutionService(execService);
        vm.stopPrank();
    }

    function encodeLegacyMessage(LegacyMessage memory legacyMsg) internal pure returns (bytes memory) {
        return LegacyMessageLib.encodeLegacyMessage(
            legacyMsg.srcSender, legacyMsg.dstReceiver, legacyMsg.srcNonce, legacyMsg.message
        );
    }

    function expectEventExecuted(uint256 srcChainId, LegacyMessage memory legacyMsg) internal {
        bytes32 messageId = keccak256(encodeLegacyMessage(legacyMsg));
        vm.expectEmit(address(messageBus));
        emit Executed({
            messageId: messageId,
            status: TxStatus.Success,
            dstAddress: legacyMsg.dstReceiver,
            srcChainId: uint64(srcChainId),
            srcNonce: legacyMsg.srcNonce
        });
    }

    function expectEventMessageSent(uint256 srcChainId, uint256 dstChainId, LegacyMessage memory legacyMsg) internal {
        bytes32 messageId = keccak256(encodeLegacyMessage(legacyMsg));
        vm.expectEmit(address(messageBus));
        emit MessageSent({
            sender: legacyMsg.srcSender,
            srcChainID: srcChainId,
            receiver: TypeCasts.addressToBytes32(legacyMsg.dstReceiver),
            dstChainId: dstChainId,
            message: legacyMsg.message,
            nonce: legacyMsg.srcNonce,
            options: legacyOptions,
            fee: MOCK_FEE,
            messageId: messageId
        });
    }

    function expectEventGasBufferSet(uint64 gasBuffer) internal {
        vm.expectEmit(address(messageBus));
        emit GasBufferSet(gasBuffer);
    }

    function expectEventMessageLengthEstimateSet(uint256 length) internal {
        vm.expectEmit(address(messageBus));
        emit MessageLengthEstimateSet(length);
    }

    function expectRevertNotEVMReceiver(bytes32 receiver) internal {
        vm.expectRevert(abi.encodeWithSelector(IMessageBus.MessageBus__NotEVMReceiver.selector, receiver));
    }

    function expectRevertInvalidOptions(bytes memory legacyOpts) internal {
        vm.expectRevert(abi.encodeWithSelector(LegacyOptionsLib.LegacyOptionsLib__InvalidOptions.selector, legacyOpts));
    }

    function expectRevertUnauthorizedGovernor(address caller) internal {
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, IC_GOVERNOR_ROLE)
        );
    }
}
