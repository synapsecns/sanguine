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
    SystemRouterHarness internal systemRouter;
    OriginHarness internal origin;
    DestinationHarness internal destination;

    uint32 internal optimisticPeriod = 420;
    uint256 internal secretValue = 1337;
    bytes internal payload = abi.encodeWithSelector(origin.setSensitiveValue.selector, secretValue);

    bytes32 internal constant SYSTEM_ROUTER =
        0xFFFFFFFF_FFFFFFFF_FFFFFFFF_00000000_00000000_00000000_00000000_00000000;

    uint32 internal constant NONCE = 420;
    bytes32 internal constant ROOT = "root";

    event Dispatch(
        bytes32 indexed messageHash,
        uint256 indexed leafIndex,
        uint64 indexed destinationAndNonce,
        bytes tips,
        bytes message
    );

    function setUp() public override {
        super.setUp();

        origin = new OriginHarness(localDomain);
        origin.initialize(INotaryManager(notaryManager));
        notaryManager.setOrigin(address(origin));

        destination = new DestinationHarness(localDomain);
        destination.initialize(remoteDomain, notary);

        systemRouter = new SystemRouterHarness(
            address(origin),
            address(destination),
            optimisticPeriod
        );
        origin.setSystemRouter(systemRouter);
        destination.setSystemRouter(systemRouter);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TEST: SETUP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_constructor() public {
        assertEq(systemRouter.origin(), address(origin));
        assertEq(systemRouter.destination(), address(destination));
        assertEq(systemRouter.optimisticSeconds(), optimisticPeriod);
    }

    function test_trustedSender() public {
        assertEq(systemRouter.trustedSender(0), SYSTEM_ROUTER);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TEST: LOCAL SYSTEM CALL                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_systemCall_toOrigin() public {
        // Destination calls Origin
        _checkSystemCall(address(destination), ISystemRouter.SystemContracts.Origin, true, "");
    }

    function test_systemCall_toOrigin_notSystemContract() public {
        // Impostor calls Origin -> should revert
        _checkSystemCall(
            address(this),
            ISystemRouter.SystemContracts.Origin,
            false,
            "Unauthorized caller"
        );
    }

    function test_systemCall_toDestination() public {
        // Origin calls Destination
        _checkSystemCall(address(origin), ISystemRouter.SystemContracts.Destination, true, "");
    }

    function test_systemCall_toDestination_notSystemContract() public {
        // Impostor calls Destination -> should revert
        _checkSystemCall(
            address(this),
            ISystemRouter.SystemContracts.Destination,
            false,
            "Unauthorized caller"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TEST: SEND SYSTEM MESSAGE                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_sendSystemMessage_origin() public {
        _testSendSystemMessage(address(origin));
    }

    function test_sendSystemMessage_destination() public {
        _testSendSystemMessage(address(destination));
    }

    function test_sendSystemMessage_notSystemContract() public {
        vm.expectRevert("Unauthorized caller");
        systemRouter.sendSystemMessage(remoteDomain, ISystemRouter.SystemContracts(0), payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TEST: RECEIVE SYSTEM MESSAGE                     ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_receiveSystemMessage_origin() public {
        bytes memory message = _prepareReceiveTest(
            optimisticPeriod,
            0,
            _createReceivedSystemMessage
        );
        skip(optimisticPeriod);
        destination.execute(message);
        assertEq(origin.sensitiveValue(), secretValue);
    }

    function test_receiveSystemMessage_destination() public {
        bytes memory message = _prepareReceiveTest(
            optimisticPeriod,
            1,
            _createReceivedSystemMessage
        );
        skip(optimisticPeriod);
        destination.execute(message);
        assertEq(destination.sensitiveValue(), secretValue);
    }

    function test_receiveSystemMessage_optimisticPeriodNotOver() public {
        bytes memory message = _prepareReceiveTest(
            optimisticPeriod,
            0,
            _createReceivedSystemMessage
        );
        skip(optimisticPeriod - 1);
        vm.expectRevert("!optimisticSeconds");
        destination.execute(message);
    }

    function test_receiveSystemMessage_optimisticPeriodForged() public {
        uint32 fakePeriod = 1;
        bytes memory message = _prepareReceiveTest(fakePeriod, 0, _createReceivedSystemMessage);
        skip(fakePeriod);
        vm.expectRevert("Client: !optimisticSeconds");
        destination.execute(message);
    }

    function test_receiveSystemMessage_unknownRecipient() public {
        bytes memory message = _prepareReceiveTest(
            optimisticPeriod,
            2,
            _createReceivedSystemMessage
        );
        skip(optimisticPeriod);
        vm.expectRevert("Unknown recipient");
        destination.execute(message);
    }

    /**
     * Anyone can send a "usual message" to SystemRouter, using its address.
     * Such messages should be rejected by SystemRouter upon receiving.
     */
    function test_rejectUsualReceivedMessage() public {
        bytes memory message = _prepareReceiveTest(
            optimisticPeriod,
            0,
            _createUsualReceivedMessage
        );
        skip(optimisticPeriod);
        vm.expectRevert("BasicClient: !trustedSender");
        destination.execute(message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _checkSystemCall(
        address _sender,
        ISystemRouter.SystemContracts _recipient,
        bool _expectSuccess,
        bytes memory _revertMessage
    ) internal {
        ISystemMockContract recipient = ISystemMockContract(_getRecipient(_recipient));
        // Sanity check
        assertFalse(recipient.sensitiveValue() == secretValue);
        if (!_expectSuccess && _revertMessage.length > 0) {
            vm.expectRevert(_revertMessage);
        }
        vm.prank(_sender);
        // Send system call to update sensitive value
        systemRouter.systemCall(_recipient, payload);
        // Check for success
        assertEq(recipient.sensitiveValue() == secretValue, _expectSuccess);
    }

    function _testSendSystemMessage(address _sender) internal {
        for (uint8 t = 0; t <= 1; ++t) {
            bytes memory message = _createSentSystemMessage(t + 1, t);
            bytes32 messageHash = keccak256(message);

            vm.expectEmit(true, true, true, true);
            emit Dispatch(
                messageHash,
                t,
                (uint64(remoteDomain) << 32) | (t + 1),
                Tips.emptyTips(),
                message
            );
            vm.prank(_sender);
            systemRouter.sendSystemMessage(remoteDomain, ISystemRouter.SystemContracts(t), payload);
        }
    }

    function _prepareReceiveTest(
        uint32 _optimisticSeconds,
        uint8 _recipient,
        function(uint32, uint32, uint8) internal returns (bytes memory) _createReceivedMessage
    ) internal returns (bytes memory message) {
        (bytes memory attestation, ) = signRemoteAttestation(notaryPK, NONCE, ROOT);
        destination.submitAttestation(attestation);

        message = _createReceivedMessage(69, _optimisticSeconds, _recipient);
        bytes32 messageHash = keccak256(message);
        destination.setMessageStatus(remoteDomain, messageHash, ROOT);

        assert(origin.sensitiveValue() != secretValue);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _createSentSystemMessage(uint32 _nonce, uint8 _recipient)
        internal
        view
        returns (bytes memory)
    {
        return
            _createSystemMessage(
                localDomain,
                SYSTEM_ROUTER,
                _nonce,
                remoteDomain,
                SYSTEM_ROUTER,
                optimisticPeriod,
                _recipient
            );
    }

    function _createReceivedSystemMessage(
        uint32 _nonce,
        uint32 _optimisticSeconds,
        uint8 _recipient
    ) internal view returns (bytes memory) {
        return
            _createSystemMessage(
                remoteDomain,
                SYSTEM_ROUTER,
                _nonce,
                localDomain,
                SYSTEM_ROUTER,
                _optimisticSeconds,
                _recipient
            );
    }

    function _createUsualReceivedMessage(
        uint32 _optimisticSeconds,
        uint32 _nonce,
        uint8 _recipient
    ) internal view returns (bytes memory) {
        return
            _createSystemMessage(
                remoteDomain,
                addressToBytes32(fakeGuard),
                _nonce,
                localDomain,
                addressToBytes32(address(systemRouter)),
                _optimisticSeconds,
                _recipient
            );
    }

    function _createSystemMessage(
        uint32 _origin,
        bytes32 _sender,
        uint32 _nonce,
        uint32 _destination,
        bytes32 _receiver,
        uint32 _optimisticSeconds,
        uint8 _recipient
    ) internal view returns (bytes memory) {
        return
            Message.formatMessage(
                Header.formatHeader(
                    _origin,
                    _sender,
                    _nonce,
                    _destination,
                    _receiver,
                    _optimisticSeconds
                ),
                Tips.emptyTips(),
                SystemMessage.formatSystemCall(_recipient, payload)
            );
    }

    function _getRecipient(ISystemRouter.SystemContracts _recipient)
        internal
        view
        returns (address recipient)
    {
        if (_recipient == ISystemRouter.SystemContracts.Origin) {
            recipient = address(origin);
        } else if (_recipient == ISystemRouter.SystemContracts.Destination) {
            recipient = address(destination);
        } else {
            revert("Unknown recipient");
        }
    }
}
