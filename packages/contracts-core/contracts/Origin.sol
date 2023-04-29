// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {BaseMessageLib} from "./libs/BaseMessage.sol";
import {MAX_CONTENT_BYTES} from "./libs/Constants.sol";
import {MemView, MemViewLib} from "./libs/MemView.sol";
import {Header, HeaderLib, MessageFlag} from "./libs/Message.sol";
import {Request, RequestLib} from "./libs/Request.sol";
import {StateReport} from "./libs/StateReport.sol";
import {State} from "./libs/State.sol";
import {Tips, TipsLib} from "./libs/Tips.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "./base/AgentSecured.sol";
import {OriginEvents} from "./events/OriginEvents.sol";
import {InterfaceGasOracle} from "./interfaces/InterfaceGasOracle.sol";
import {InterfaceOrigin} from "./interfaces/InterfaceOrigin.sol";
import {StateHub} from "./hubs/StateHub.sol";

contract Origin is StateHub, OriginEvents, InterfaceOrigin {
    using MemViewLib for bytes;
    using TipsLib for bytes;
    using TypeCasts for address;

    address public immutable gasOracle;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 domain, address agentManager_, address gasOracle_)
        AgentSecured("0.0.3", domain, agentManager_)
    {
        gasOracle = gasOracle_;
    }

    /// @notice Initializes Origin contract:
    /// - msg.sender is set as contract owner
    /// - State of "empty merkle tree" is saved
    function initialize() external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable_init();
        // Initialize "states": state of an "empty merkle tree" is saved
        _initializeStates();
    }

    // ═══════════════════════════════════════════════ SEND MESSAGES ═══════════════════════════════════════════════════

    /// @inheritdoc InterfaceOrigin
    function sendBaseMessage(
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticPeriod,
        uint256 paddedRequest,
        bytes memory content
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {
        // Check that content is not too large
        require(content.length <= MAX_CONTENT_BYTES, "content too long");
        // This will revert if msg.value is lower than value of minimum tips
        Tips tips = _getMinimumTips(destination, paddedRequest, content.length).matchValue(msg.value);
        Request request = RequestLib.wrapPadded(paddedRequest);
        // Format the BaseMessage body
        bytes memory body = BaseMessageLib.formatBaseMessage({
            sender_: msg.sender.addressToBytes32(),
            recipient_: recipient,
            tips_: tips,
            request_: request,
            content_: content
        });
        // Send the message
        return _sendMessage(destination, optimisticPeriod, MessageFlag.Base, body);
    }

    /// @inheritdoc InterfaceOrigin
    function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes memory payload)
        external
        onlyAgentManager
        returns (uint32 messageNonce, bytes32 messageHash)
    {
        // AgentManager (checked via modifier) is responsible for constructing the calldata payload correctly.
        return _sendMessage(destination, optimisticPeriod, MessageFlag.Manager, payload);
    }

    /// @inheritdoc InterfaceOrigin
    function withdrawTips(address recipient, uint256 amount) external onlyAgentManager {
        require(address(this).balance >= amount, "Insufficient balance");
        (bool success,) = recipient.call{value: amount}("");
        require(success, "Recipient reverted");
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceOrigin
    function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength)
        external
        view
        returns (uint256 tipsValue)
    {
        return _getMinimumTips(destination, paddedRequest, contentLength).value();
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Sends the given message to the specified destination. Message hash is inserted
    /// into the Origin Merkle Tree, which will enable message execution on destination chain.
    function _sendMessage(uint32 destination, uint32 optimisticPeriod, MessageFlag flag, bytes memory body)
        internal
        returns (uint32 messageNonce, bytes32 messageHash)
    {
        // Format the message header
        messageNonce = _nextNonce();
        Header header = HeaderLib.encodeHeader({
            origin_: localDomain,
            nonce_: messageNonce,
            destination_: destination,
            optimisticPeriod_: optimisticPeriod
        });
        // Format the full message payload
        bytes memory msgPayload = flag.formatMessage(header, body);
        // Insert new leaf into the Origin Merkle Tree and save the updated state
        messageHash = keccak256(msgPayload);
        _insertAndSave(messageHash);
        // Emit event with message information
        emit Sent(messageHash, messageNonce, destination, msgPayload);
    }

    /// @dev Returns the minimum tips for sending a message to the given destination with the given request and content.
    function _getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength)
        internal
        view
        returns (Tips)
    {
        return
            TipsLib.wrapPadded(InterfaceGasOracle(gasOracle).getMinimumTips(destination, paddedRequest, contentLength));
    }
}
