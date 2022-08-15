// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { Client } from "../client/Client.sol";
import { TypedMemView } from "../libs/TypedMemView.sol";
import { SystemMessage } from "./SystemMessage.sol";
import { ISystemMessenger } from "../interfaces/ISystemMessenger.sol";
import { Tips } from "../libs/Tips.sol";

import { Address } from "@openzeppelin/contracts/utils/Address.sol";

contract SystemMessenger is Client, ISystemMessenger {
    using Address for address;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using SystemMessage for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Should be set to a reasonably high value to prevent forging of a system message.
     *      Optimistic period is enforced in the base contract: Client.handle()
     */
    uint32 internal _optimisticPeriod;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[49] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(
        address _origin,
        address _destination,
        uint32 _optimisticSeconds
    ) Client(_origin, _destination) {
        // TODO: Do we ever want to adjust this?
        // (the value should be the same across all chains)
        // Or could it be converted into immutable?
        _optimisticPeriod = _optimisticSeconds;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Allows calls only from any of the System Contracts
    modifier anySystem() {
        require(msg.sender == origin || msg.sender == destination, "Unauthorized caller");
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Send System Message to one of the System Contracts on origin chain
     * @dev     Only System contracts are allowed to call this function.
     *          Note that knowledge of recipient address is not required,
     *          routing will be done by SystemMessenger on the destination chain.
     * @param _destination  Domain of destination chain
     * @param _recipient    System contract type of the recipient
     * @param _payload      Data for calling recipient on destination chain
     */
    function sendSystemMessage(
        uint32 _destination,
        SystemContracts _recipient,
        bytes memory _payload
    ) external anySystem {
        bytes memory message = SystemMessage.formatCall(uint8(_recipient), _payload);
        /**
         * @dev Origin should recognize SystemMessenger as the "true sender"
         *      and use SYSTEM_SENDER address as "sender" instead. This enables not
         *      knowing SystemMessenger address on remote chain in advance.
         */
        _send(_destination, Tips.emptyTips(), message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PUBLIC FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Returns optimistic period of the merkle root, used for proving messages to SystemMessenger.
     *          All messages to remote chains will be sent with this period.
     *          Merkle root is checked to be at least this old on all incoming messages: see Client.handle()
     */
    function optimisticSeconds() public view override returns (uint32) {
        return _optimisticPeriod;
    }

    /**
     * @notice Returns eligible address of sender/receiver on given remote chain.
     */
    function trustedSender(uint32) public pure override returns (bytes32) {
        /**
         * @dev SystemMessenger will be sending messages to SYSTEM_SENDER address,
         * and will only accept incoming messages from SYSTEM_SENDER as well (see Client.sol).
         *
         * It's not possible for anyone but SystemMessenger
         * to send messages "from SYSTEM_SENDER" on other deployed chains.
         *
         * Destination is supposed to reject messages
         * from unknown chains, so we can skip origin check here.
         */
        return SystemMessage.SYSTEM_SENDER;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Handles an incoming message. All security checks are done in Client.handle()
     */
    function _handle(
        uint32,
        uint32,
        bytes32,
        bytes memory _message
    ) internal override {
        (SystemMessage.SystemMessageType messageType, bytes29 messageView) = _message
            .ref(0)
            .systemMessage();

        if (messageType == SystemMessage.SystemMessageType.Call) {
            address recipient = _getSystemRecipient(messageView.callRecipient());
            require(recipient != address(0), "System Contract not set");
            bytes29 payload = messageView.callPayload();
            // this will call recipient and bubble the revert from the external call
            recipient.functionCall(payload.clone());
        } else if (messageType == SystemMessage.SystemMessageType.Adjust) {
            // TODO: handle messages with instructions to adjust some of the SystemMessenger parameters
        }
    }

    function _getSystemRecipient(uint8 _recipient) internal view returns (address) {
        if (_recipient == uint8(SystemContracts.Origin)) return origin;
        if (_recipient == uint8(SystemContracts.Destination)) return destination;
        revert("Unknown recipient");
    }
}
