// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { BasicClient } from "../client/BasicClient.sol";
import { TypedMemView } from "../libs/TypedMemView.sol";
import { SystemMessage } from "../libs/SystemMessage.sol";
import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { Tips } from "../libs/Tips.sol";

import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @notice Sends and receives system messages:
 * an internal messaging channel between system contracts.
 * This makes it possible to send message to any system contract (e.g. Origin) on another chain
 * without knowing its address in advance, making easy cross-chain setups possible.
 * What is more, knowing System Router address on destination chain is also
 * not required. Instead, a special SYSTEM_ROUTER value is used as sender/recipient
 * for the system messages (see SystemMessage.sol for more details).
 *
 * SystemRouter keeps track of all system contracts deployed on current chain,
 * and routes messages to/from them. System contracts are supposed to have
 * external methods, guarded by onlySystemRouter modifier. These methods could
 * be called cross-chain from any of the system contracts.
 */
contract SystemRouter is BasicClient, ISystemRouter {
    using Address for address;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using SystemMessage for bytes;
    using SystemMessage for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              IMMUTABLES                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint32 public immutable localDomain;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(
        uint32 _localDomain,
        address _origin,
        address _destination
    ) BasicClient(_origin, _destination) {
        localDomain = _localDomain;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Call a System Contract on the destination chain with a given data payload.
     * Note: for system calls on the local chain
     * - use `destination = localDomain`
     * - `_optimisticSeconds` value will be ignored
     *
     * @dev Only System contracts are allowed to call this function.
     * Note: knowledge of recipient address is not required, routing will be done by SystemRouter
     * on the destination chain. Following call will be made on destination chain:
     * - recipient.call(_data, originDomain, originSender, rootSubmittedAt)
     * Note: data payload is extended with abi encoded (domain, sender, rootTimestamp)
     * This allows recipient to check:
     * - domain where a system call originated (local domain in this case)
     * - system entity, who initiated the call (msg.sender on local chain)
     * - timestamp when merkle root was submitted and optimistic timer started ticking
     * @param _destination          Domain of destination chain
     * @param _optimisticSeconds    Optimistic period for the message
     * @param _recipient            System entity to receive the call on destination chain
     * @param _data                 Data for calling recipient on destination chain
     */
    function systemCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity _recipient,
        bytes memory _data
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        bytes memory payload = _formatCalldata(caller, _data);
        if (_destination == localDomain) {
            /// @dev Passing current timestamp for consistency
            /// Functions that could be called both from a local chain,
            /// as well as from a remote chain with an optimistic period
            /// will have to check `origin` and `rootSubmittedAt` to ensure validity.
            _localSystemCall(uint8(_recipient), payload, block.timestamp);
        } else {
            bytes[] memory systemCalls = new bytes[](1);
            systemCalls[0] = SystemMessage.formatSystemCall(uint8(_recipient), payload);
            _remoteSystemCall(_destination, _optimisticSeconds, systemCalls);
        }
    }

    /**
     * @notice Calls a few system contracts with the given calldata.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes[] memory _dataArray
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        uint256 amount = _recipients.length;
        bytes[] memory payloads = new bytes[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            payloads[i] = _formatCalldata(caller, _dataArray[i]);
        }
        _multiCall(_destination, _optimisticSeconds, _recipients, payloads);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PUBLIC FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns eligible address of sender/receiver on given remote chain.
     */
    function trustedSender(uint32) public pure override returns (bytes32) {
        /**
         * @dev SystemRouter will be sending messages to SYSTEM_ROUTER address,
         * and will only accept incoming messages from SYSTEM_ROUTER as well (see Client.sol).
         *
         * It's not possible for anyone but SystemRouter
         * to send messages "from SYSTEM_ROUTER" on other deployed chains.
         *
         * Destination is supposed to reject messages
         * from unknown chains, so we can skip origin check here.
         */
        return SystemMessage.SYSTEM_ROUTER;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Handles an incoming message. Security checks are done in BasicClient.handle()
     * Optimistic period could be anything at this point.
     */
    function _handleUnsafe(
        uint32,
        uint32,
        uint256 _rootSubmittedAt,
        bytes memory _message
    ) internal override {
        bytes[] memory systemMessages = abi.decode(_message, (bytes[]));
        uint256 amount = systemMessages.length;
        for (uint256 i = 0; i < amount; ++i) {
            bytes29 message = systemMessages[i].castToSystemMessage();
            require(message.isSystemMessage(), "Not a system message");
            (SystemMessage.MessageFlag messageType, bytes29 body) = message.unpackMessage();
            if (messageType == SystemMessage.MessageFlag.Call) {
                /// @dev System Contract has to check root submission time
                /// to confirm that optimistic period for given function has passed.
                _localSystemCall(
                    body.callRecipient(),
                    body.callPayload().clone(),
                    _rootSubmittedAt
                );
            } else {
                // Sanity check: no other MessageFlag values exist
                assert(false);
            }
        }
    }

    function _localSystemCall(
        uint8 _recipient,
        bytes memory _payload,
        uint256 _rootSubmittedAt
    ) internal {
        address recipient = _getSystemAddress(_recipient);
        require(recipient != address(0), "System Contract not set");
        // this will call recipient and bubble the revert from the external call
        recipient.functionCall(abi.encodePacked(_payload, _rootSubmittedAt));
        /// @dev add root timestamp as the last argument
        /// uint256 type allows us to use encodePacked w/o casting
    }

    function _remoteSystemCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        bytes[] memory _systemCalls
    ) internal {
        bytes memory message = abi.encode(_systemCalls);
        /**
         * @dev Origin should recognize SystemRouter as the "true sender"
         *      and use SYSTEM_ROUTER address as "sender" instead. This enables not
         *      knowing SystemRouter address on remote chain in advance.
         */
        _send(_destination, _optimisticSeconds, Tips.emptyTips(), message);
    }

    function _multiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes[] memory _payloads
    ) internal {
        uint256 amount = _recipients.length;
        if (_destination == localDomain) {
            for (uint256 i = 0; i < amount; ++i) {
                /// @dev Passing current timestamp for consistency, see systemCall() for details
                _localSystemCall(uint8(_recipients[i]), _payloads[i], block.timestamp);
            }
        } else {
            bytes[] memory systemCalls = new bytes[](amount);
            for (uint256 i = 0; i < amount; ++i) {
                systemCalls[i] = SystemMessage.formatSystemCall(
                    uint8(_recipients[i]),
                    _payloads[i]
                );
            }
            _remoteSystemCall(_destination, _optimisticSeconds, systemCalls);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _formatCalldata(SystemEntity _caller, bytes memory _data)
        internal
        view
        returns (bytes memory)
    {
        /**
         * @dev Payload for contract call is:
         * ====== ENCODED ON ORIGIN CHAIN ======
         * 1. Function selector and params (_data)
         * 2. (domain, caller) are the following two arguments
         * ====== ENCODED ON REMOTE CHAIN ======
         * 3. Root timestamp is the last argument,
         * and will be appended before the call on destination chain.
         */
        return abi.encodePacked(_data, abi.encode(localDomain, _caller));
    }

    function _getSystemEntity(address _caller) internal view returns (SystemEntity) {
        if (_caller == origin) return SystemEntity.Origin;
        if (_caller == destination) return SystemEntity.Destination;
        revert("Unauthorized caller");
    }

    function _getSystemAddress(uint8 _recipient) internal view returns (address) {
        if (_recipient == uint8(SystemEntity.Origin)) return origin;
        if (_recipient == uint8(SystemEntity.Destination)) return destination;
        revert("Unknown recipient");
    }
}
