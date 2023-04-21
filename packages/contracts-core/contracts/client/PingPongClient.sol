// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {Request, RequestLib} from "../libs/Request.sol";
import {Tips, TipsLib} from "../libs/Tips.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {IMessageRecipient} from "../interfaces/IMessageRecipient.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";

contract PingPongClient is IMessageRecipient {
    using TypeCasts for address;

    struct PingPongMessage {
        uint256 pingId;
        bool isPing;
        uint16 counter;
    }

    // ════════════════════════════════════════════════ IMMUTABLES ═════════════════════════════════════════════════════

    /// @notice Local chain Origin: used for sending messages
    address public immutable origin;

    /// @notice Local chain Destination: used for receiving messages
    address public immutable destination;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    uint256 public random;

    /// @notice Amount of "Ping" messages sent.
    uint256 public pingsSent;

    /// @notice Amount of "Ping" messages received.
    /// Every received Ping message leads to sending a Pong message back to initial sender.
    uint256 public pingsReceived;

    /// @notice Amount of "Pong" messages received.
    /// When all messages are delivered, should be equal to `pingsSent`
    uint256 public pongsReceived;

    // ══════════════════════════════════════════════════ EVENTS ═══════════════════════════════════════════════════════

    /// @notice Emitted when a Ping message is sent.
    /// Triggered externally, or by receveing a Pong message with instructions to do more pings.
    event PingSent(uint256 pingId);

    /// @notice Emitted when a Ping message is received.
    /// Will always send a Pong message back.
    event PingReceived(uint256 pingId);

    /// @notice Emitted when a Pong message is sent.
    /// Triggered whenever a Ping message is received.
    event PongSent(uint256 pingId);

    /// @notice Emitted when a Pong message is received.
    /// Will initiate a new Ping, if the counter in the message is non-zero.
    event PongReceived(uint256 pingId);

    // ════════════════════════════════════════════════ CONSTRUCTOR ════════════════════════════════════════════════════

    constructor(address origin_, address destination_) {
        origin = origin_;
        destination = destination_;
        // Initiate "random" value
        random = uint256(keccak256(abi.encode(block.number)));
    }

    // ═══════════════════════════════════════════════ MESSAGE LOGIC ═══════════════════════════════════════════════════

    /// @inheritdoc IMessageRecipient
    function receiveBaseMessage(uint32 origin_, uint32, bytes32 sender, uint256, bytes memory content)
        external
        payable
    {
        require(msg.sender == destination, "PingPongClient: !destination");
        PingPongMessage memory message = abi.decode(content, (PingPongMessage));
        if (message.isPing) {
            // Ping is received
            ++pingsReceived;
            emit PingReceived(message.pingId);
            // Send Pong back
            _pong(origin_, sender, message);
        } else {
            // Pong is received
            ++pongsReceived;
            emit PongReceived(message.pingId);
            // Send extra ping, if initially requested
            if (message.counter != 0) {
                _ping(origin_, sender, message.counter - 1);
            }
        }
    }

    function doPings(uint16 pingCount, uint32 destination_, address recipient, uint16 counter) external {
        for (uint256 i = 0; i < pingCount; ++i) {
            _ping(destination_, recipient.addressToBytes32(), counter);
        }
    }

    /// @notice Send a Ping message to destination chain.
    /// Upon receiving a Ping, a Pong message will be sent back.
    /// If `counter > 0`, this process will be repeated when the Pong message is received.
    /// @param destination_ Chain to send Ping message to
    /// @param recipient    Recipient of Ping message
    /// @param counter      Additional amount of Ping-Pong rounds to conclude
    function doPing(uint32 destination_, address recipient, uint16 counter) external {
        _ping(destination_, recipient.addressToBytes32(), counter);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function nextOptimisticPeriod() public view returns (uint32 period) {
        // Use random optimistic period up to one minute
        return uint32(random % 1 minutes);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Returns a random optimistic period value from 0 to 59 seconds.
    function optimisticPeriod() internal returns (uint32 period) {
        // Use random optimistic period up to one minute
        period = nextOptimisticPeriod();
        // Adjust "random" value
        random = uint256(keccak256(abi.encode(random)));
    }

    /**
     * @dev Send a "Ping" or "Pong" message.
     * @param destination_  Domain of destination chain
     * @param recipient     Message recipient on destination chain
     * @param message   Ping-pong message
     */
    function _sendMessage(uint32 destination_, bytes32 recipient, PingPongMessage memory message) internal {
        Tips tips = TipsLib.emptyTips();
        // TODO: figure out the logic for a ping-pong test
        Request request = RequestLib.encodeRequest(0, 0);
        bytes memory content = abi.encode(message);
        InterfaceOrigin(origin).sendBaseMessage(
            destination_, recipient, optimisticPeriod(), Tips.unwrap(tips), Request.unwrap(request), content
        );
    }

    /// @dev Initiate a new Ping-Pong round.
    function _ping(uint32 destination_, bytes32 recipient, uint16 counter) internal {
        uint256 pingId = pingsSent++;
        _sendMessage(destination_, recipient, PingPongMessage({pingId: pingId, isPing: true, counter: counter}));
        emit PingSent(pingId);
    }

    /// @dev Send a Pong message back.
    function _pong(uint32 destination_, bytes32 recipient, PingPongMessage memory message) internal {
        _sendMessage(
            destination_, recipient, PingPongMessage({pingId: message.pingId, isPing: false, counter: message.counter})
        );
        emit PongSent(message.pingId);
    }
}
