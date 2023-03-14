// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { HeaderLib, MessageLib } from "../../../contracts/libs/Message.sol";

import { TipsTools } from "./TipsTools.t.sol";
import { SynapseUtilities } from "../../utils/SynapseUtilities.t.sol";

abstract contract MessageTools is SynapseUtilities, TipsTools {
    // Values specifying the need to mock the data instead of using the provided data
    bytes internal constant MOCK_BODY = "Mock the message body";
    bytes32 internal constant MOCK_RECIPIENT = "Mock the message recipient";
    uint32 internal constant MOCK_OPTIMISTIC_SECONDS = type(uint32).max - 1337;

    // Origin domain
    uint32 internal messageOrigin;
    address internal messageSenderAddress;
    // Sender of message
    bytes32 internal messageSender;
    // Nonce of message
    uint32 internal messageNonce;
    // Destination domain
    uint32 internal messageDestination;
    // Recipient of message
    bytes32 internal messageRecipient;
    // Message optimistic period
    uint32 internal messageOptimisticSeconds;
    // Message body
    bytes internal messageBody;
    // Full message payload (all above, formatted)
    bytes internal messageRaw;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Create message with all the given data
    // pass MOCK_X constant to mock field X instead
    function createMessage(
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination,
        bool mockTips,
        bytes memory body,
        bytes32 recipient,
        uint32 optimisticSeconds
    ) public {
        saveMessageData(origin, sender, nonce, destination);
        saveMockableMessageData(body, recipient, optimisticSeconds, mockTips);
        createMessage();
    }

    // Create message using all the saved data
    function createMessage() public {
        messageRaw = MessageLib.formatMessage(
            HeaderLib.formatHeader(
                messageOrigin,
                messageSender,
                messageNonce,
                messageDestination,
                messageRecipient,
                messageOptimisticSeconds
            ),
            tipsRaw,
            messageBody
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            SAVE TEST DATA                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Save the given data for later use
    function saveMessageData(
        uint32 origin,
        bytes32 sender,
        uint32 nonce,
        uint32 destination
    ) public {
        messageOrigin = origin;
        messageSender = sender;
        messageSenderAddress = bytes32ToAddress(sender);
        messageNonce = nonce;
        messageDestination = destination;
    }

    // Save the given data for later use. Use mocks if instructed:
    // i.e. passing MOCK_BODY as body will lead to using a mocked value instead
    function saveMockableMessageData(
        bytes memory body,
        bytes32 recipient,
        uint32 optimisticSeconds,
        bool mockTips
    ) public {
        messageBody = (keccak256(body) == keccak256(MOCK_BODY)) ? _createMockBody() : body;
        messageRecipient = (recipient == MOCK_RECIPIENT) ? _createMockRecipient() : recipient;
        messageOptimisticSeconds = (optimisticSeconds == MOCK_OPTIMISTIC_SECONDS)
            ? _createMockOptimisticSeconds()
            : optimisticSeconds;
        if (mockTips) {
            createMockTips({ nonce: messageNonce });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Create unique salt for mocking data
    function _createSalt() internal view returns (bytes memory) {
        return _createSalt(messageOrigin, messageDestination, messageNonce);
    }

    function _createMockBody() internal view returns (bytes memory) {
        return _createMockBody(messageOrigin, messageDestination, messageNonce);
    }

    function _createMockRecipient() internal view returns (bytes32) {
        return keccak256(_createSalt());
    }

    function _createMockOptimisticSeconds() internal view returns (uint32) {
        return messageNonce * 10;
    }

    function _createSalt(
        uint32 origin,
        uint32 destination,
        uint32 nonce
    ) internal pure returns (bytes memory) {
        return abi.encode(origin, destination, nonce);
    }

    function _createMockBody(
        uint32 origin,
        uint32 destination,
        uint32 nonce
    ) internal pure returns (bytes memory) {
        return abi.encode("message body", _createSalt(origin, destination, nonce));
    }
}
