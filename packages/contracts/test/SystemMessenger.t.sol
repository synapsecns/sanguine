// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { INotaryManager } from "../contracts/interfaces/INotaryManager.sol";
import { ISystemMessenger } from "../contracts/interfaces/ISystemMessenger.sol";
import { SystemMessage } from "../contracts/system/SystemMessage.sol";
import { Message } from "../contracts/libs/Message.sol";
import { Header } from "../contracts/libs/Header.sol";
import { Tips } from "../contracts/libs/Tips.sol";
import { SynapseTestWithNotaryManager } from "./utils/SynapseTest.sol";

import { OriginHarness } from "./harnesses/OriginHarness.sol";
import { DestinationHarness } from "./harnesses/DestinationHarness.sol";
import { SystemMessengerHarness } from "./harnesses/SystemMessengerHarness.sol";

contract SystemMessengerTest is SynapseTestWithNotaryManager {
    SystemMessengerHarness internal systemMessenger;
    OriginHarness internal origin;
    DestinationHarness internal destination;

    uint32 internal optimisticPeriod = 420;
    uint256 internal secretValue = 1337;
    bytes payload = abi.encodeWithSelector(origin.setSensitiveValue.selector, secretValue);

    bytes32 internal constant SYSTEM_SENDER =
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

        systemMessenger = new SystemMessengerHarness(
            address(origin),
            address(destination),
            optimisticPeriod
        );
        origin.setSystemMessenger(systemMessenger);
        destination.setSystemMessenger(systemMessenger);
    }

    function test_constructor() public {
        assertEq(systemMessenger.origin(), address(origin));
        assertEq(systemMessenger.destination(), address(destination));
        assertEq(systemMessenger.optimisticSeconds(), optimisticPeriod);
    }

    function test_trustedSender() public {
        assertEq(systemMessenger.trustedSender(0), SYSTEM_SENDER);
    }

    function test_sendSystemMessage_origin() public {
        _testSendSystemMessage(address(origin));
    }

    function test_sendSystemMessage_destination() public {
        _testSendSystemMessage(address(destination));
    }

    function test_sendSystemMessage_notSystemSender() public {
        vm.expectRevert("Unauthorized caller");
        systemMessenger.sendSystemMessage(
            remoteDomain,
            ISystemMessenger.SystemContracts(0),
            payload
        );
    }

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
     * Anyone can send a "usual message" to SystemMessenger, using its address.
     * Such messages should be rejected by SystemMessenger upon receiving.
     */
    function test_rejectUsualReceivedMessage() public {
        bytes memory message = _prepareReceiveTest(
            optimisticPeriod,
            0,
            _createUsualReceivedMessage
        );
        skip(optimisticPeriod);
        vm.expectRevert("Client: !trustedSender");
        destination.execute(message);
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
            systemMessenger.sendSystemMessage(
                remoteDomain,
                ISystemMessenger.SystemContracts(t),
                payload
            );
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

    function _createSentSystemMessage(uint32 _nonce, uint8 _recipient)
        internal
        view
        returns (bytes memory)
    {
        return
            _createSystemMessage(
                localDomain,
                SYSTEM_SENDER,
                _nonce,
                remoteDomain,
                SYSTEM_SENDER,
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
                SYSTEM_SENDER,
                _nonce,
                localDomain,
                SYSTEM_SENDER,
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
                addressToBytes32(fakeSigner),
                _nonce,
                localDomain,
                addressToBytes32(address(systemMessenger)),
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
}
