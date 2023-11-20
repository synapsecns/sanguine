// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {BaseMessageLib} from "./libs/memory/BaseMessage.sol";
import {MAX_CONTENT_BYTES} from "./libs/Constants.sol";
import {
    ContentLengthTooBig,
    EthTransferFailed,
    IncorrectDestinationDomain,
    InsufficientEthBalance
} from "./libs/Errors.sol";
import {GasData, GasDataLib} from "./libs/stack/GasData.sol";
import {MemView, MemViewLib} from "./libs/memory/MemView.sol";
import {Header, MessageLib, MessageFlag} from "./libs/memory/Message.sol";
import {Request, RequestLib} from "./libs/stack/Request.sol";
import {Tips, TipsLib} from "./libs/stack/Tips.sol";
import {TypeCasts} from "./libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {AgentSecured} from "./base/AgentSecured.sol";
import {OriginEvents} from "./events/OriginEvents.sol";
import {InterfaceGasOracle} from "./interfaces/InterfaceGasOracle.sol";
import {InterfaceOrigin} from "./interfaces/InterfaceOrigin.sol";
import {StateHub} from "./hubs/StateHub.sol";

/// @notice `Origin` contract is used for sending messages to remote chains. It is done
/// by inserting the message hashes into the Origin Merkle, which makes it possible to
/// prove that message was sent using the Merkle proof against the Origin Merkle Root. This essentially
/// compresses the list of messages into a single 32-byte value that needs to be stored on the destination chain.
/// `Origin` is responsible for the following:
/// - Formatting the sent message payloads, and inserting their hashes into the Origin Merkle Tree.
/// - Keeping track of its own historical states (see parent contract `StateHub`).
/// - Enforcing minimum tip values for sent base messages based on the provided execution requests.
/// - Distributing the collected tips upon request from a local `AgentManager` contract.
contract Origin is StateHub, OriginEvents, InterfaceOrigin {
    using MemViewLib for bytes;
    using MessageLib for bytes;
    using TipsLib for bytes;
    using TypeCasts for address;

    address public immutable gasOracle;

    modifier onlyRemoteDestination(uint32 destination) {
        if (destination == localDomain) revert IncorrectDestinationDomain();
        _;
    }

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 synapseDomain_, address agentManager_, address inbox_, address gasOracle_)
        AgentSecured("0.0.3", synapseDomain_, agentManager_, inbox_)
    {
        gasOracle = gasOracle_;
    }

    /// @notice Initializes Origin contract:
    /// - `owner_` is set as contract owner
    /// - State of "empty merkle tree" is saved
    function initialize(address owner_) external initializer {
        __MessagingBase_init(owner_);
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
    ) external payable onlyRemoteDestination(destination) returns (uint32 messageNonce, bytes32 messageHash) {
        // Check that content is not too large
        if (content.length > MAX_CONTENT_BYTES) revert ContentLengthTooBig();
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
        onlyRemoteDestination(destination)
        returns (uint32 messageNonce, bytes32 messageHash)
    {
        // AgentManager (checked via modifier) is responsible for constructing the calldata payload correctly.
        return _sendMessage(destination, optimisticPeriod, MessageFlag.Manager, payload);
    }

    /// @inheritdoc InterfaceOrigin
    function withdrawTips(address recipient, uint256 amount) external onlyAgentManager {
        if (address(this).balance < amount) revert InsufficientEthBalance();
        (bool success,) = recipient.call{value: amount}("");
        if (!success) revert EthTransferFailed();
        emit TipWithdrawalCompleted(recipient, amount);
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
        Header header = flag.encodeHeader({
            origin_: localDomain,
            nonce_: messageNonce,
            destination_: destination,
            optimisticPeriod_: optimisticPeriod
        });
        // Format the full message payload
        bytes memory msgPayload = MessageLib.formatMessage(header, body);
        // Insert new leaf into the Origin Merkle Tree and save the updated state
        messageHash = msgPayload.castToMessage().leaf();
        _insertAndSave(messageHash);
        // Emit event with message information
        emit Sent(messageHash, messageNonce, destination, msgPayload);
        // Update the gas oracle data. We do this after the provided tips are checked so that
        // the provided message is sent in the event that the gas oracle data is updated to a higher value.
        InterfaceGasOracle(gasOracle).updateGasData(destination);
        // TODO: consider doing this before the message is sent, while adjusting GasOracle.getMinimumTips() to
        // use the pending gas oracle data.
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

    /// @dev Gets the current gas data from the gas oracle to be saved as part of the Origin State.
    function _fetchGasData() internal view override returns (GasData) {
        return GasDataLib.wrapGasData(InterfaceGasOracle(gasOracle).getGasData());
    }
}
