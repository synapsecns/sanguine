// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IMessageRecipient } from "../interfaces/IMessageRecipient.sol";
import { IOrigin } from "../interfaces/IOrigin.sol";
import { Tips } from "../libs/Tips.sol";
import { TypeCasts } from "../libs/TypeCasts.sol";

contract PingPongClient is IMessageRecipient {
    using TypeCasts for address;

    struct PingMessage {
        uint256 pingId;
        uint16 pongsLeft;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              IMMUTABLES                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // local chain Origin: used for sending messages
    address public immutable origin;

    // local chain Destination: used for receiving messages
    address public immutable destination;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256 public random;

    uint256 public totalSent;

    uint256 public totalReceived;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Emitted when "Ping" is sent
    event Ping(uint256 indexed pingId, uint16 pongsLeft);

    // Emitted when "Ping" is received
    event Pong(uint256 indexed pingId, uint16 pongsLeft);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(address _origin, address _destination) {
        origin = _origin;
        destination = _destination;
        // Initiate "random" value
        random = uint256(keccak256(abi.encode(block.number)));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          RECEIVING MESSAGES                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Called by Destination upon executing the message.
    function handle(
        uint32 _origin,
        uint32,
        bytes32 _sender,
        uint256,
        bytes memory _message
    ) external {
        require(msg.sender == destination, "TestClient: !destination");
        PingMessage memory _msg = abi.decode(_message, (PingMessage));
        ++totalReceived;
        emit Pong(_msg.pingId, _msg.pongsLeft);
        // Send the message back, if there are pongs left to do
        if (_msg.pongsLeft != 0) {
            _sendPing(_origin, _sender, PingMessage(_msg.pingId, _msg.pongsLeft - 1));
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           SENDING MESSAGES                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function doPings(
        uint16 _pingCount,
        uint32 _destination,
        address _recipient,
        uint16 _pongsTotal
    ) external {
        for (uint256 i = 0; i < _pingCount; ++i) {
            doPing(_destination, _recipient, _pongsTotal);
        }
    }

    /// @notice Send a ping to destination chain. Upon receiving a ping,
    /// a pong will be performed until the specified amount of pongs is reached.
    function doPing(
        uint32 _destination,
        address _recipient,
        uint16 _pongsTotal
    ) public {
        uint256 pingId = totalSent++;
        _sendPing(_destination, _recipient.addressToBytes32(), PingMessage(pingId, _pongsTotal));
    }

    function nextOptimisticPeriod() public view returns (uint32 period) {
        // Use random optimistic period up to one minute
        return uint32(random % 1 minutes);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @dev Returns a random optimistic period value from 0 to 59 seconds.
    function _optimisticPeriod() internal returns (uint32 period) {
        // Use random optimistic period up to one minute
        period = nextOptimisticPeriod();
        // Adjust "random" value
        random = uint256(keccak256(abi.encode(random)));
    }

    /**
     * @dev Send a "Ping" message.
     * @param _destination  Domain of destination chain
     * @param _recipient    Message recipient on destination chain
     * @param _msg          Ping-pong message
     */
    function _sendPing(
        uint32 _destination,
        bytes32 _recipient,
        PingMessage memory _msg
    ) internal {
        bytes memory tips = Tips.emptyTips();
        bytes memory message = abi.encode(_msg);
        IOrigin(origin).dispatch(_destination, _recipient, _optimisticPeriod(), tips, message);
        emit Ping(_msg.pingId, _msg.pongsLeft);
    }
}
