// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { ISystemRouter } from "../contracts/interfaces/ISystemRouter.sol";
import { SystemMessage } from "../contracts/libs/SystemMessage.sol";
import { Message } from "../contracts/libs/Message.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Tips } from "../contracts/libs/Tips.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { DestinationHarness } from "./harnesses/DestinationHarness.sol";
import { SystemRouterHarness } from "./harnesses/SystemRouterHarness.sol";

interface ISystemMockContract {
    function sensitiveValue() external view returns (uint256);
}

// solhint-disable func-name-mixedcase
contract SystemRouterTest is SynapseTestWithNotaryManager {
    struct MessageContext {
        uint32 origin;
        bytes32 sender;
        uint32 destination;
        bytes32 recipient;
    }

    SystemRouterHarness internal systemRouter;
    OriginHarness internal origin;
    DestinationHarness internal destination;

    MessageContext internal sentSystemMessage;
    MessageContext internal receivedSystemMessage;
    MessageContext internal receivedUsualMessage;

    uint32[] internal domains;

    uint32 internal nonce = 1;
    uint32 internal optimisticSeconds = OPTIMISTIC_PERIOD;

    uint8 internal systemCaller = uint8(ISystemRouter.SystemContracts.Origin);
    uint8 internal systemRecipient = uint8(ISystemRouter.SystemContracts.Destination);
    uint256 internal secretValue = 1337;
    bytes internal data;

    bytes32 internal constant SYSTEM_ROUTER =
        0xFFFFFFFF_FFFFFFFF_FFFFFFFF_00000000_00000000_00000000_00000000_00000000;

    uint32 internal constant OPTIMISTIC_PERIOD = 420;
    uint32 internal constant NONCE = 69;
    bytes32 internal constant ROOT = "root";

    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes tips,
        bytes message
    );

    event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt);

    function setUp() public override {
        super.setUp();

        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));

        destination = new DestinationHarness(localDomain);
        destination.initialize(remoteDomain, notary);

        systemRouter = new SystemRouterHarness(localDomain, address(origin), address(destination));
        origin.setSystemRouter(systemRouter);
        destination.setSystemRouter(systemRouter);

        sentSystemMessage = MessageContext({
            origin: localDomain,
            sender: SYSTEM_ROUTER,
            destination: remoteDomain,
            recipient: SYSTEM_ROUTER
        });
        receivedSystemMessage = MessageContext({
            origin: remoteDomain,
            sender: SYSTEM_ROUTER,
            destination: localDomain,
            recipient: SYSTEM_ROUTER
        });
        receivedUsualMessage = MessageContext({
            origin: remoteDomain,
            sender: addressToBytes32(address(this)),
            destination: localDomain,
            recipient: addressToBytes32(address(systemRouter))
        });
        domains = new uint32[](2);
        domains[0] = localDomain;
        domains[1] = remoteDomain;
        data = _createData(secretValue);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TEST: SETUP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_constructor() public {
        assertEq(systemRouter.origin(), address(origin));
        assertEq(systemRouter.destination(), address(destination));
    }

    function test_trustedSender() public {
        assertEq(systemRouter.trustedSender(0), SYSTEM_ROUTER);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TEST: SYSTEM CALL                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_systemCall_local_toOrigin() public {
        systemCaller = uint8(ISystemRouter.SystemContracts.Destination);
        systemRecipient = uint8(ISystemRouter.SystemContracts.Origin);
        // Destination calls Origin
        _checkLocalSystemCall();
    }

    function test_systemCall_local_toDestination() public {
        // Origin calls Destination (default setup)
        _checkLocalSystemCall();
    }

    function test_systemCall_remote_origin() public {
        _checkRemoteSystemCall();
    }

    function test_systemCall_remote_destination() public {
        systemCaller = uint8(ISystemRouter.SystemContracts.Destination);
        _checkRemoteSystemCall();
    }

    function test_systemCall_notSystemContract() public {
        for (uint256 i = 0; i < domains.length; ++i) {
            vm.expectRevert("Unauthorized caller");
            systemRouter.systemCall(domains[i], 0, ISystemRouter.SystemContracts.Origin, data);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TEST: RECEIVE SYSTEM MESSAGE                     ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_receiveSystemMessage_origin() public {
        // remote Destination -> local Origin
        systemCaller = uint8(ISystemRouter.SystemContracts.Destination);
        systemRecipient = uint8(ISystemRouter.SystemContracts.Origin);
        _checkReceiveSystemMessage();
    }

    function test_receiveSystemMessage_destination() public {
        // remote Origin -> local Destination (default setup)
        _checkReceiveSystemMessage();
    }

    function test_receiveSystemMessage_optimisticPeriodZero() public {
        // Test harnesses don't block system messages with a small (or even zero)
        // optimistic period.
        // TODO: introduce and test "force optimistic period" modifiers in system contracts
        optimisticSeconds = 0;
        _checkReceiveSystemMessage();
    }

    function test_receiveSystemMessage_optimisticPeriodNotOver() public {
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds - 1);
        vm.expectRevert("!optimisticSeconds");
        destination.execute(message);
    }

    function test_receiveSystemMessage_unknownRecipient() public {
        // recipient = 2 does not exist
        systemRecipient = 2;
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds);
        vm.expectRevert("Unknown recipient");
        destination.execute(message);
    }

    /**
     * Anyone can send a "usual message" to SystemRouter, using its address.
     * Such messages should be rejected by SystemRouter upon receiving.
     */
    function test_rejectUsualReceivedMessage() public {
        bytes memory message = _prepareReceiveTest(receivedUsualMessage);
        skip(optimisticSeconds);
        vm.expectRevert("BasicClient: !trustedSender");
        destination.execute(message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _checkLocalSystemCall() internal {
        address sender = _getSystemAddress(systemCaller);
        ISystemRouter.SystemContracts recipient = ISystemRouter.SystemContracts(systemRecipient);
        ISystemMockContract recipientMock = ISystemMockContract(_getSystemAddress(systemRecipient));
        // Sanity check
        assertFalse(recipientMock.sensitiveValue() == secretValue);
        vm.prank(sender);
        // Send system call to update sensitive value
        systemRouter.systemCall(localDomain, 0, recipient, data);
        // Check for success
        assertTrue(recipientMock.sensitiveValue() == secretValue);
    }

    function _checkRemoteSystemCall() internal {
        address sender = _getSystemAddress(systemCaller);
        // Send messages from sender to every system contract on remote chain
        for (systemRecipient = 0; systemRecipient <= 1; (++systemRecipient, ++nonce)) {
            ISystemRouter.SystemContracts recipient = ISystemRouter.SystemContracts(
                systemRecipient
            );
            bytes memory message = _createSystemMessage(sentSystemMessage);
            vm.expectEmit(true, true, true, true);
            emit Dispatch(
                keccak256(message),
                nonce - 1,
                (uint64(remoteDomain) << 32) | nonce,
                Tips.emptyTips(),
                message
            );
            vm.prank(sender);
            systemRouter.systemCall(remoteDomain, optimisticSeconds, recipient, data);
        }
    }

    function _checkReceiveSystemMessage() internal {
        uint256 rootSubmittedAt = block.timestamp;
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds);
        vm.expectEmit(true, true, true, true);
        emit LogSystemCall(remoteDomain, systemCaller, rootSubmittedAt);
        destination.execute(message);
        assertEq(
            ISystemMockContract(_getSystemAddress(systemRecipient)).sensitiveValue(),
            secretValue
        );
    }

    function _prepareReceiveTest(MessageContext memory context)
        internal
        returns (bytes memory message)
    {
        message = _createSystemMessage(context);
        // Mark message as proved against ROOT
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, ROOT);
        destination.submitAttestation(attestation);
        destination.setMessageStatus(remoteDomain, keccak256(message), ROOT);
        // Sanity check
        assert(origin.sensitiveValue() != secretValue);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _createData(uint256 _secretValue) internal view returns (bytes memory) {
        return abi.encodeWithSelector(origin.setSensitiveValue.selector, _secretValue);
    }

    function _createSystemMessage(MessageContext memory context)
        internal
        view
        returns (bytes memory)
    {
        // Reconstruct payload in the brute force way
        bytes memory payload = abi.encodePacked(
            data,
            uint256(context.origin),
            uint256(systemCaller)
        );
        bytes[] memory systemCalls = new bytes[](1);
        systemCalls[0] = SystemMessage.formatSystemCall(systemRecipient, payload);
        return
            Message.formatMessage(
                Header.formatHeader(
                    context.origin,
                    context.sender,
                    nonce,
                    context.destination,
                    context.recipient,
                    optimisticSeconds
                ),
                Tips.emptyTips(),
                abi.encode(systemCalls)
            );
    }

    function _getSystemContract(address _account)
        internal
        view
        returns (ISystemRouter.SystemContracts systemContract)
    {
        if (_account == address(origin)) {
            systemContract = ISystemRouter.SystemContracts.Origin;
        } else if (_account == address(destination)) {
            systemContract = ISystemRouter.SystemContracts.Destination;
        } else {
            revert("Unknown caller");
        }
    }

    function _getSystemAddress(uint8 _systemContract) internal view returns (address account) {
        if (_systemContract == uint8(ISystemRouter.SystemContracts.Origin)) {
            account = address(origin);
        } else if (_systemContract == uint8(ISystemRouter.SystemContracts.Destination)) {
            account = address(destination);
        } else {
            // Sanity check
            assert(false);
        }
    }
}
