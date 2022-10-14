// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { ISystemRouter } from "../contracts/interfaces/ISystemRouter.sol";
import { SystemMessage } from "../contracts/libs/SystemMessage.sol";
import { Message } from "../contracts/libs/Message.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Tips } from "../contracts/libs/Tips.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";
import { ProofGenerator } from "./utils/ProofGenerator.sol";

import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { DestinationHarness } from "./harnesses/DestinationHarness.sol";
import { SystemRouterHarness } from "./harnesses/SystemRouterHarness.sol";

interface ISystemMockContract {
    function sensitiveValue() external view returns (uint256);
}

// solhint-disable func-name-mixedcase, max-states-count
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
    MessageContext internal receivedFakeSystemMessage;
    MessageContext internal receivedUsualMessage;
    MessageContext internal synapseSystemMessage;

    uint32[] internal domains;

    uint32 internal nonce = 1;
    uint32 internal optimisticSeconds = 420;

    /// @notice Default test setup is Origin calling Destination
    uint8 internal systemCaller = uint8(ISystemRouter.SystemEntity.Origin);
    uint8 internal systemRecipient = uint8(ISystemRouter.SystemEntity.Destination);
    uint256 internal secretValue = 1337;
    bytes internal data;

    bytes32 internal constant SYSTEM_ROUTER =
        0xFFFFFFFF_FFFFFFFF_FFFFFFFF_00000000_00000000_00000000_00000000_00000000;

    uint32 internal constant SYNAPSE_DOMAIN = 4269;

    ProofGenerator internal proofGen;
    uint32 internal constant NONCE = 1;
    uint32 internal constant INDEX = 0;
    bytes32[32] internal proof;
    bytes32 internal root;

    event Dispatch(
        bytes32 indexed messageHash,
        uint32 indexed nonce,
        uint32 indexed destination,
        bytes tips,
        bytes message
    );

    event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt);

    event UsualCall(address recipient, uint256 newValue);
    event OnlyLocalCall(address recipient, uint256 newValue);
    event OnlyOriginCall(address recipient, uint256 newValue);
    event OnlyDestinationCall(address recipient, uint256 newValue);
    event OnlyTwoHoursCall(address recipient, uint256 newValue);

    function setUp() public override {
        super.setUp();

        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));

        destination = new DestinationHarness(localDomain);
        destination.initialize();
        destination.addNotary(remoteDomain, notary);

        systemRouter = new SystemRouterHarness(localDomain, address(origin), address(destination));
        origin.setSystemRouter(systemRouter);
        destination.setSystemRouter(systemRouter);

        proofGen = new ProofGenerator();

        // System Router on local chain sends a system message to remote
        sentSystemMessage = MessageContext({
            origin: localDomain,
            sender: SYSTEM_ROUTER,
            destination: remoteDomain,
            recipient: SYSTEM_ROUTER
        });
        // System Router on remote chain sends a system message to local
        receivedSystemMessage = MessageContext({
            origin: remoteDomain,
            sender: SYSTEM_ROUTER,
            destination: localDomain,
            recipient: SYSTEM_ROUTER
        });
        // Arbitrary address on remote sends a message to local, specifying
        //  System Router address as the recipient
        receivedUsualMessage = MessageContext({
            origin: remoteDomain,
            sender: addressToBytes32(address(this)),
            destination: localDomain,
            recipient: addressToBytes32(address(systemRouter))
        });
        // Arbitrary address on remote sends a message to local, specifying
        // SYSTEM_ROUTER as the recipient
        receivedFakeSystemMessage = MessageContext({
            origin: remoteDomain,
            sender: addressToBytes32(address(this)),
            destination: localDomain,
            recipient: SYSTEM_ROUTER
        });
        // System Router on Synapse chain sends a system message to local
        synapseSystemMessage = MessageContext({
            origin: SYNAPSE_DOMAIN,
            sender: SYSTEM_ROUTER,
            destination: localDomain,
            recipient: SYSTEM_ROUTER
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

    /**
     * @notice System call on local chain: Destination to Origin.
     */
    function test_systemCall_local_toOrigin() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        systemRecipient = uint8(ISystemRouter.SystemEntity.Origin);
        _checkLocalSystemCall();
    }

    /**
     * @notice System call on local chain: Origin to Destination.
     */
    function test_systemCall_local_toDestination() public {
        _checkLocalSystemCall();
    }

    /**
     * @notice Send system call from local to remote chain: Origin to Destination.
     */
    function test_systemCall_remote_origin() public {
        _checkRemoteSystemCall();
    }

    /**
     * @notice Send system call from local to remote chain: Destination to Destination.
     */
    function test_systemCall_remote_destination() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        _checkRemoteSystemCall();
    }

    /**
     * @notice System call (both local and remote), when invoked by NOT a system contract.
     */
    function test_systemCall_notSystemContract() public {
        for (uint256 i = 0; i < domains.length; ++i) {
            vm.expectRevert("Unauthorized caller");
            systemRouter.systemCall(domains[i], 0, ISystemRouter.SystemEntity.Origin, data);
        }
    }

    /**
     * @notice System call to Destination, when Destination address is not set in the SystemRouter.
     */
    function test_systemCall_systemRouterNotConfigured() public {
        systemRouter = new SystemRouterHarness(localDomain, address(origin), address(0));
        vm.prank(address(origin));
        vm.expectRevert("System Contract not set");
        systemRouter.systemCall(localDomain, 0, ISystemRouter.SystemEntity.Destination, data);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TEST: RECEIVE SYSTEM MESSAGE                     ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive system call from remote to local chain: Destination to Origin.
     */
    function test_receiveSystemMessage_origin() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        systemRecipient = uint8(ISystemRouter.SystemEntity.Origin);
        _checkReceiveSystemMessage();
    }

    /**
     * @notice Receive system call from remote to local chain: Origin to Destination.
     */
    function test_receiveSystemMessage_destination() public {
        _checkReceiveSystemMessage();
    }

    /**
     * @notice Receive system call from remote to local chain: Origin to Destination.
     * Optimistic period is zero, unprotected function is called, meaning test is successful.
     */
    function test_receiveSystemMessage_optimisticPeriodZero() public {
        // Test harnesses don't block system messages with a small (or even zero)
        // optimistic period.
        optimisticSeconds = 0;
        _checkReceiveSystemMessage();
    }

    /**
     * @notice Receive system call from remote to local chain: Origin to Destination.
     * Called before optimistic period is over, should fail.
     */
    function test_receiveSystemMessage_optimisticPeriodNotOver() public {
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds - 1);
        vm.expectRevert("!optimisticSeconds");
        destination.execute(message, proof, INDEX);
    }

    /**
     * @notice Receive system call from remote to local chain: ??? to Destination.
     * Specified recipient does not exist, should fail.
     */
    function test_receiveSystemMessage_unknownRecipient() public {
        // recipient = 2 does not exist
        systemRecipient = 2;
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds);
        vm.expectRevert("Unknown recipient");
        destination.execute(message, proof, INDEX);
    }

    /**
     * @notice Receive system call from remote to local chain: Origin to Destination.
     * System call with empty payload, should be rejected as badly formatted system message.
     */
    function test_receiveSystemMessage_badlyFormattedMessage() public {
        bytes memory message = _createSystemMessage(receivedSystemMessage, "");
        _prepareReceiveTest(message);
        skip(optimisticSeconds);
        vm.expectRevert("Not a system message");
        destination.execute(message, proof, INDEX);
    }

    /**
     * @notice SystemRouter receives a plain message sent via Destination directly
     * on the remote chain, specifying SystemRouter address as the recipient (anyone could do that).
     * SystemRouter should reject such messages: only special value of SYSTEM_ROUTER
     * is accepted as sender for the system messages.
     */
    function test_rejectUsualReceivedMessage() public {
        bytes memory message = _prepareReceiveTest(receivedUsualMessage);
        skip(optimisticSeconds);
        vm.expectRevert("BasicClient: !trustedSender");
        destination.execute(message, proof, INDEX);
    }

    /**
     * @notice SystemRouter receives a plain message sent via Destination directly
     * on the remote chain, specifying SYSTEM_ROUTER value as the recipient (anyone could do that).
     * SystemRouter should reject such messages: only special value of SYSTEM_ROUTER
     * is accepted as sender for the system messages.
     */
    function test_rejectFakeSystemMessage() public {
        bytes memory message = _prepareReceiveTest(receivedFakeSystemMessage);
        skip(optimisticSeconds);
        vm.expectRevert("BasicClient: !trustedSender");
        destination.execute(message, proof, INDEX);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        TEST: LOCAL MULTICALL                         ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice System multicall on local chain. Origin invokes:
     * 1. origin.setSensitiveValue(1337)
     * 2. destination.setSensitiveValueOnlyLocal(1337 * 2)
     * 3. destination.setSensitiveValueOnlyOrigin(1337 * 3)
     */
    function test_systemMultiCall_local() public {
        (
            ISystemRouter.SystemEntity[] memory recipients,
            bytes[] memory dataArray
        ) = _prepareMultiCallTest(true, true);
        vm.prank(address(origin));
        systemRouter.systemMultiCall(localDomain, 0, recipients, dataArray);
        // TODO: check execution order
    }

    /**
     * @notice System multicall on local chain. Destination invokes:
     * 1. origin.setSensitiveValue(1337)
     * 2. destination.setSensitiveValueOnlyLocal(1337 * 2)
     * 3. destination.setSensitiveValueOnlyOrigin(1337 * 3)
     * Call #3 fails meaning the whole transaction is reverted.
     */
    function test_systemMultiCall_local_failedCall() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        (
            ISystemRouter.SystemEntity[] memory recipients,
            bytes[] memory dataArray
        ) = _prepareMultiCallTest(true, false);
        vm.prank(address(destination));
        vm.expectRevert("!allowedCaller");
        // Multicall includes onlyOrigin call, meaning multicall will fail
        systemRouter.systemMultiCall(localDomain, 0, recipients, dataArray);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                        TEST: REMOTE MULTICALL                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Send system multicall from local chain to remote chain. Destination invokes:
     * 1. origin.setSensitiveValue(1337)
     * 2. destination.setSensitiveValueOnlyTwoHours(1337 * 2)
     * 3. destination.setSensitiveValueOnlyDestination(1337 * 3)
     */
    function test_systemMultiCall_remote() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        (
            ISystemRouter.SystemEntity[] memory recipients,
            bytes[] memory dataArray
        ) = _prepareMultiCallTest(true, false);
        bytes memory message = _createSystemMessage(sentSystemMessage, recipients, dataArray);
        _expectSystemMessage(message);
        vm.prank(address(destination));
        systemRouter.systemMultiCall(remoteDomain, optimisticSeconds, recipients, dataArray);
    }

    /**
     * @notice Receive system multicall from remote chain to local chain.
     * Remote Destination invoked:
     * 1. origin.setSensitiveValue(1337)
     * 2. destination.setSensitiveValueOnlyTwoHours(1337 * 2)
     * 3. destination.setSensitiveValueOnlyDestination(1337 * 3)
     */
    function test_receiveSystemMessage_multicall() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        optimisticSeconds = 2 hours;
        (
            ISystemRouter.SystemEntity[] memory recipients,
            bytes[] memory dataArray
        ) = _prepareMultiCallTest(false, true);
        bytes memory message = _createSystemMessage(receivedSystemMessage, recipients, dataArray);
        _prepareReceiveTest(message);
        skip(optimisticSeconds);
        destination.execute(message, proof, INDEX);
        // TODO: check execution order
    }

    /**
     * @notice Receive system multicall from remote chain to local chain.
     * Remote Destination invoked:
     * 1. origin.setSensitiveValue(1337)
     * 2. destination.setSensitiveValueOnlyTwoHours(1337 * 2)
     * 3. destination.setSensitiveValueOnlyDestination(1337 * 3)
     * Optimistic period is set to (2 hours - 1 second)
     * Call #2 fails meaning the whole transaction is reverted.
     */
    function test_receiveSystemMessage_multicall_failedCall() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        optimisticSeconds = 2 hours - 1;
        (
            ISystemRouter.SystemEntity[] memory recipients,
            bytes[] memory dataArray
        ) = _prepareMultiCallTest(false, false);
        bytes memory message = _createSystemMessage(receivedSystemMessage, recipients, dataArray);
        _prepareReceiveTest(message);
        skip(optimisticSeconds);
        vm.expectRevert("!optimisticPeriod");
        destination.execute(message, proof, INDEX);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   TEST: SYSTEM CONTRACT MODIFIERS                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    // TODO: this should be in a SystemContract mock test.
    // Move over once a reusable testing suite is established.

    /**
     * @notice System call on local chain.
     * Origin invokes:
     *  destination.setSensitiveValueOnlyOrigin(1337)
     * Destination invokes:
     *  origin.setSensitiveValueOnlyDestination(1337)
     */
    function test_systemCall_local_onlyCaller() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyOrigin.selector,
            secretValue
        );
        _checkLocalSystemCall();

        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        systemRecipient = uint8(ISystemRouter.SystemEntity.Origin);
        data = abi.encodeWithSelector(
            origin.setSensitiveValueOnlyDestination.selector,
            secretValue
        );
        _checkLocalSystemCall();
    }

    /**
     * @notice System call on local chain. Destination invokes:
     *  origin.setSensitiveValueOnlyOrigin(1337)
     * Should fail due to the wrong caller.
     */
    function test_systemCall_local_onlyCaller_wrongCaller() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        systemRecipient = uint8(ISystemRouter.SystemEntity.Origin);
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyOrigin.selector,
            secretValue
        );
        vm.prank(address(destination));
        vm.expectRevert("!allowedCaller");
        systemRouter.systemCall(localDomain, 0, ISystemRouter.SystemEntity.Destination, data);
    }

    /**
     * @notice System call on local chain.
     * Origin invokes:
     *  destination.setSensitiveValueOnlyOriginDestination(1337)
     * Destination invokes:
     *  origin.setSensitiveValueOnlyOriginDestination(1337)
     */
    function test_systemCall_local_onlyCallers() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyOriginDestination.selector,
            secretValue
        );
        _checkLocalSystemCall();

        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        systemRecipient = uint8(ISystemRouter.SystemEntity.Origin);
        data = abi.encodeWithSelector(
            origin.setSensitiveValueOnlyOriginDestination.selector,
            secretValue
        );
        _checkLocalSystemCall();
    }

    /**
     * @notice System call on local chain. Origin invokes:
     *  destination.setSensitiveValueOnlyLocal(1337)
     */
    function test_systemCall_local_onlyLocal() public {
        data = abi.encodeWithSelector(destination.setSensitiveValueOnlyLocal.selector, secretValue);
        _checkLocalSystemCall();
    }

    /**
     * @notice System call on local chain. Origin invokes:
     *  destination.setSensitiveValueOnlyTwoHours(1337)
     * Should fail: optimistic period check always fails on local calls
     */
    function test_systemCall_local_onlyOptimisticPeriodOver() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyTwoHours.selector,
            secretValue
        );
        vm.prank(address(destination));
        vm.expectRevert("!optimisticPeriod");
        systemRouter.systemCall(localDomain, 0, ISystemRouter.SystemEntity.Destination, data);
    }

    /**
     * @notice Receive system call from remote chain to local chain.
     * Remote Origin invokes:
     *  destination.setSensitiveValueOnlyOrigin(1337)
     */
    function test_systemCall_remote_onlyCaller() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyOrigin.selector,
            secretValue
        );
        _checkReceiveSystemMessage();
    }

    /**
     * @notice Receive system call from remote chain to local chain.
     * Remote Destination invokes:
     *  destination.setSensitiveValueOnlyOrigin(1337)
     * Should fail due to the wrong caller.
     */
    function test_systemCall_remote_onlyCaller_wrongCaller() public {
        systemCaller = uint8(ISystemRouter.SystemEntity.Destination);
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyOrigin.selector,
            secretValue
        );
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds);
        vm.expectRevert("!allowedCaller");
        destination.execute(message, proof, INDEX);
    }

    /**
     * @notice Receive system call from remote chain to local chain.
     * Remote Origin invokes:
     *  destination.setSensitiveValueOnlyLocal(1337)
     * Should fail as this is a cross-chain call, not the local one.
     */
    function test_systemCall_remote_onlyLocal() public {
        data = abi.encodeWithSelector(destination.setSensitiveValueOnlyLocal.selector, secretValue);
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds);
        vm.expectRevert("!localDomain");
        destination.execute(message, proof, INDEX);
    }

    /**
     * @notice Receive system call from remote chain to local chain.
     * Remote Origin invokes:
     *  destination.setSensitiveValueOnlyTwoHours(1337)
     * Optimistic period is set to 2 hours.
     */
    function test_systemCall_remote_onlyOptimisticPeriodOver() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyTwoHours.selector,
            secretValue
        );

        optimisticSeconds = 2 hours;
        _checkReceiveSystemMessage();
    }

    /**
     * @notice Receive system call from remote chain to local chain.
     * Remote Origin invokes:
     *  destination.setSensitiveValueOnlyTwoHours(1337)
     * Optimistic period is set to (2 hours - 1 second) causing recipient to reject the call.
     */
    function test_systemCall_remote_onlyOptimisticPeriodOver_periodReduced() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlyTwoHours.selector,
            secretValue
        );
        optimisticSeconds = 2 hours - 1;
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds);
        vm.expectRevert("!optimisticPeriod");
        destination.execute(message, proof, INDEX);
    }

    /**
     * @notice Receive system call from Synapse chain to local chain.
     * Remote Origin invokes:
     *  destination.setSensitiveValueOnlySynapseChain(1337)
     */
    function test_systemCall_remote_onlySynapseChain() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlySynapseChain.selector,
            secretValue
        );
        _checkReceiveSystemMessage(synapseSystemMessage);
    }

    /**
     * @notice Receive system call from remote chain to local chain.
     * Remote Origin invokes:
     *  destination.setSensitiveValueOnlySynapseChain(1337)
     * Calls comes not from the Synapse chain making transaction fail
     */
    function test_systemCall_remote_onlySynapseChain_wrongChain() public {
        data = abi.encodeWithSelector(
            destination.setSensitiveValueOnlySynapseChain.selector,
            secretValue
        );
        bytes memory message = _prepareReceiveTest(receivedSystemMessage);
        skip(optimisticSeconds);
        vm.expectRevert("!synapseDomain");
        destination.execute(message, proof, INDEX);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Local system call is triggered.
     * `systemCaller` is calling `systemRecipient`:
     *      setSensitiveValue(1337)
     * The sensitive value is checked to have been updated.
     */
    function _checkLocalSystemCall() internal {
        address sender = _getSystemAddress(systemCaller);
        ISystemRouter.SystemEntity recipient = ISystemRouter.SystemEntity(systemRecipient);
        ISystemMockContract recipientMock = ISystemMockContract(_getSystemAddress(systemRecipient));
        // Sanity check
        assertFalse(recipientMock.sensitiveValue() == secretValue);
        vm.prank(sender);
        // Send system call to update sensitive value
        systemRouter.systemCall(localDomain, 0, recipient, data);
        // Check for success
        assertTrue(recipientMock.sensitiveValue() == secretValue);
    }

    /**
     * @dev Remote system call is triggered.
     * `systemCaller` is calling every possible recipient (one by one)
     *      setSensitiveValue(1337)
     * Every dispatched message is checked to contain needed information.
     */
    function _checkRemoteSystemCall() internal {
        address sender = _getSystemAddress(systemCaller);
        // Send messages from sender to every system contract on remote chain
        for (systemRecipient = 0; systemRecipient <= 1; (++systemRecipient, ++nonce)) {
            ISystemRouter.SystemEntity recipient = ISystemRouter.SystemEntity(systemRecipient);
            bytes memory message = _createSystemMessage(sentSystemMessage);
            _expectSystemMessage(message);
            vm.prank(sender);
            systemRouter.systemCall(remoteDomain, optimisticSeconds, recipient, data);
        }
    }

    /**
     * @dev Expect Dispatch event to remote chain for a given (nonce, message)
     */
    function _expectSystemMessage(bytes memory _message) internal {
        vm.expectEmit(true, true, true, true);
        emit Dispatch(keccak256(_message), nonce, remoteDomain, Tips.emptyTips(), _message);
    }

    /**
     * @dev System call message (from the remote chain) is executed on local chain.
     * (domain, caller, timestamp) are checked on the recipient side.
     * Secret value is checked to have been updated.
     */
    function _checkReceiveSystemMessage() internal {
        _checkReceiveSystemMessage(receivedSystemMessage);
    }

    function _checkReceiveSystemMessage(MessageContext memory context) internal {
        uint256 rootSubmittedAt = block.timestamp;
        bytes memory message = _prepareReceiveTest(context);
        skip(optimisticSeconds);
        vm.expectEmit(true, true, true, true);
        emit LogSystemCall(context.origin, systemCaller, rootSubmittedAt);
        destination.execute(message, proof, INDEX);
        assertEq(
            ISystemMockContract(_getSystemAddress(systemRecipient)).sensitiveValue(),
            secretValue
        );
    }

    /**
     * @dev Constructs a message for "receive message" test given the test context.
     * Then marks this message as being ready for the execution with the given optimistic period.
     */
    function _prepareReceiveTest(MessageContext memory context)
        internal
        returns (bytes memory message)
    {
        message = _createSystemMessage(context);
        if (context.origin != remoteDomain) {
            uint32 _prevRemoteDomainValue = remoteDomain;
            destination.removeNotary(remoteDomain, notary);
            remoteDomain = context.origin;
            destination.addNotary(context.origin, notary);
            _prepareReceiveTest(message);
            remoteDomain = _prevRemoteDomainValue;
        } else {
            _prepareReceiveTest(message);
        }
    }

    /**
     * @dev Marks given message as being ready for the execution with the given optimistic period.
     */
    function _prepareReceiveTest(bytes memory _message) internal {
        // Create a merkle tree with a lonely message
        bytes32[] memory leafs = new bytes32[](1);
        leafs[0] = keccak256(_message);
        proofGen.createTree(leafs);
        root = proofGen.getRoot();
        proof = proofGen.getProof(INDEX);
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, root);
        destination.submitAttestation(attestation);
    }

    /**
     * @dev Constructs recipient and data arrays for the multicall test.
     * Emits events if a test assumes successful message execution.
     */
    // solhint-disable code-complexity
    function _prepareMultiCallTest(bool _sentFromLocal, bool _emitEvents)
        internal
        returns (ISystemRouter.SystemEntity[] memory recipients, bytes[] memory dataArray)
    {
        uint256 amount = 3;
        recipients = new ISystemRouter.SystemEntity[](amount);
        dataArray = new bytes[](amount);
        {
            uint256 value = _getSecretValue(0);
            recipients[0] = ISystemRouter.SystemEntity.Origin;
            dataArray[0] = abi.encodeWithSelector(origin.setSensitiveValue.selector, value);
            if (_emitEvents) {
                vm.expectEmit(true, true, true, true);
                emit UsualCall(address(origin), value);
            }
        }
        {
            uint256 value = _getSecretValue(1);
            recipients[1] = ISystemRouter.SystemEntity.Destination;
            dataArray[1] = abi.encodeWithSelector(
                _sentFromLocal
                    ? destination.setSensitiveValueOnlyLocal.selector
                    : destination.setSensitiveValueOnlyTwoHours.selector,
                value
            );
            if (_emitEvents) {
                vm.expectEmit(true, true, true, true);
                if (_sentFromLocal) {
                    emit OnlyLocalCall(address(destination), value);
                } else {
                    emit OnlyTwoHoursCall(address(destination), value);
                }
            }
        }
        {
            uint256 value = _getSecretValue(2);
            recipients[2] = ISystemRouter.SystemEntity.Destination;
            dataArray[2] = abi.encodeWithSelector(
                _sentFromLocal
                    ? destination.setSensitiveValueOnlyOrigin.selector
                    : destination.setSensitiveValueOnlyDestination.selector,
                value
            );
            if (_emitEvents) {
                vm.expectEmit(true, true, true, true);
                if (_sentFromLocal) {
                    emit OnlyOriginCall(address(destination), value);
                } else {
                    emit OnlyDestinationCall(address(destination), value);
                }
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Creates payload for setSensitiveValue(_secretValue) call
     */
    function _createData(uint256 _secretValue) internal view returns (bytes memory) {
        return abi.encodeWithSelector(origin.setSensitiveValue.selector, _secretValue);
    }

    /**
     * @dev Constructs a system message for a system call given the test context.
     * `systemCaller` is calling `systemRecipient` with `data` payload.
     */
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
        return _createSystemMessage(context, payload);
    }

    /**
     * @dev Constructs a system message for a system call given the test context.
     * `systemCaller` is calling `systemRecipient` with a given payload.
     */
    function _createSystemMessage(MessageContext memory context, bytes memory payload)
        internal
        view
        returns (bytes memory)
    {
        bytes[] memory systemCalls = new bytes[](1);
        systemCalls[0] = SystemMessage.formatSystemCall(systemRecipient, payload);
        return _createSystemMessage(context, systemCalls);
    }

    /**
     * @dev Constructs a system message for a system multicall given the test context.
     * `systemCaller` is calling `recipients[i]` with `dataArray[i]` payload.
     */
    function _createSystemMessage(
        MessageContext memory context,
        ISystemRouter.SystemEntity[] memory recipients,
        bytes[] memory dataArray
    ) internal view returns (bytes memory) {
        uint256 amount = recipients.length;
        bytes[] memory systemCalls = new bytes[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            bytes memory payload = abi.encodePacked(
                dataArray[i],
                uint256(context.origin),
                uint256(systemCaller)
            );
            systemCalls[i] = SystemMessage.formatSystemCall(uint8(recipients[i]), payload);
        }
        return _createSystemMessage(context, systemCalls);
    }

    /**
     * @dev Constructs a system message for a system multicall given the test context.
     * `systemCaller` is doing following calls: `systemCalls`
     */
    function _createSystemMessage(MessageContext memory context, bytes[] memory systemCalls)
        internal
        view
        returns (bytes memory)
    {
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

    /**
     * @dev Gets secret value for a multicall test, given the system call index.
     * All calls are done with different values to check that needed payloads are used.
     */
    function _getSecretValue(uint256 _testIndex) internal view returns (uint256) {
        return secretValue * (_testIndex + 1);
    }

    /**
     * @dev Gets system contract's respective SystemEntity value.
     */
    function _getSystemContract(address _account)
        internal
        view
        returns (ISystemRouter.SystemEntity systemContract)
    {
        if (_account == address(origin)) {
            systemContract = ISystemRouter.SystemEntity.Origin;
        } else if (_account == address(destination)) {
            systemContract = ISystemRouter.SystemEntity.Destination;
        } else {
            revert("Unknown caller");
        }
    }

    /**
     * @dev Gets system contract's address by its SystemEntity value.
     */
    function _getSystemAddress(uint8 _systemContract) internal view returns (address account) {
        if (_systemContract == uint8(ISystemRouter.SystemEntity.Origin)) {
            account = address(origin);
        } else if (_systemContract == uint8(ISystemRouter.SystemEntity.Destination)) {
            account = address(destination);
        } else {
            // Sanity check
            assert(false);
        }
    }
}
