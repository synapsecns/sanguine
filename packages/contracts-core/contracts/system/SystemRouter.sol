// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { BasicClient } from "../client/BasicClient.sol";
import { LocalDomainContext } from "../context/LocalDomainContext.sol";
import { TypedMemView } from "../libs/TypedMemView.sol";
import { SystemCall } from "../libs/SystemCall.sol";
import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { Tips } from "../libs/Tips.sol";

import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @notice Router for calls between system contracts (aka "System Calls").
 * SystemRouter makes it possible to perform a call from one system contract to another
 * without knowing the recipient address. This works for both on-chain calls, when caller and
 * recipient are deployed on the same chain, and for cross-chain calls, when caller and
 * recipient are deployed on different chains.
 *
 * SystemRouter allows both calls and "multi calls". Multicall performs a series of calls,
 * calling requested recipients one by one, supplying the requested payloads. The whole multicall
 * will fail, if any of the calls reverts.
 *
 * SystemRouter keeps track of all system contracts deployed on current chain. This enables sending
 * calls "to Origin", "to Destination" without actually knowing the Origin/Destination address.
 *
 * For the on-chain calls, SystemRouter is simply forwarding the call to the requested recipient(s).
 *
 * For the cross-chain calls, SystemRouter sends a message to SystemRouter on destination chain,
 * in order for it to forward the call to the request recipient(s). SystemRouter doesn't need to be
 * aware of System Router address on destination chain. It uses a special value SYSTEM_ROUTER for
 * the recipient instead. Origin contract on origin chain enforces the invariant that only
 * the system router could specify such a recipient. For these messages Origin uses SYSTEM_ROUTER
 * for "sender" field, instead of actual System Router address (as it does for usual messages).
 * Destination contract routes messages, where SYSTEM_ROUTER is specified as recipient, to a local
 * System Router. System Router only accepts incoming messages with "sender = SYSTEM_ROUTER",
 * enforcing the invariant that only origin's System Router is able to send cross-chain messages to
 * destination's System Router, assuming no optimistic verification fraud (more on this below).
 *
 * @dev Security considerations
 * System Router only accepts calls from the system contracts, so it's not possible for the attacker
 * to initiate a system call through the system router. However, every system contract that wants
 * to expose one of its external functions for the system calls, should do the following:
 * 1. Such functions should have the same last three arguments:
 * - someFunction(<...>, uint32 origin, ISystemRouter.SystemEntity caller, uint256 rootSubmittedAt)
 * These arguments are filled by System Routers on origin and destination chain. This allows
 * the recipient to set the restrictions for receiving the call in a very granular way.
 * To perform a call, use payload = abi.encodeWithSelector(someFunction.selector, <...>);
 * 2. Guard such function with `onlySystemRouter` modifier to prevent unauthorized direct calls.
 * Guard function with additional modifiers based on `origin`, `caller`, `rootSubmittedAt`.
 * `rootSubmittedAt` based modifier is a must for receiving cross-chain system calls.  Any Notary
 * can potentially commit fraud, and try to execute an arbitrary message, including
 * a "message to System Router". By enforcing a minimum optimistic latency for the recipient this
 * attack can be militated, assuming there is at least one honest Guard willing to report the fraud.
 */
contract SystemRouter is LocalDomainContext, BasicClient, ISystemRouter {
    using Address for address;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using SystemCall for bytes;
    using SystemCall for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable no-empty-blocks
    constructor(
        uint32 _domain,
        address _origin,
        address _destination
    ) BasicClient(_origin, _destination) LocalDomainContext(_domain) {}

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
     * - recipient.call(_data, callOrigin, systemCaller, rootSubmittedAt)
     * This allows recipient to check:
     * - callOrigin: domain where a system call originated (local domain in this case)
     * - systemCaller: system entity who initiated the call (msg.sender on local chain)
     * - rootSubmittedAt:
     *   - For cross-chain calls: timestamp when merkle root (used for executing the system call)
     *     was submitted to destination and its optimistic timer started ticking
     *   - For on-chain calls: timestamp of the current block
     *
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
        // Append calldata with (callOrigin, systemCaller) to form the payload
        bytes memory payload = _formatCalldata(caller, _data);
        if (_destination == _localDomain()) {
            /// @dev Passing current timestamp for consistency
            /// Functions that could be called both from a local chain,
            /// as well as from a remote chain with an optimistic period
            /// will have to check `callOrigin` and `rootSubmittedAt` to ensure validity.
            _localSystemCall(uint8(_recipient), payload, block.timestamp);
        } else {
            bytes[] memory systemCalls = new bytes[](1);
            systemCalls[0] = SystemCall.formatSystemCall(uint8(_recipient), payload);
            // To generalize things, a remote system call is always a multicall.
            // In case of a "usual call", this is a multicall with exactly one call inside.
            _remoteSystemCall(_destination, _optimisticSeconds, systemCalls);
        }
    }

    /**
     * @notice Calls a few system contracts using the given calldata for each call.
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
            // Append calldata with (callOrigin, systemCaller) to form the payload
            payloads[i] = _formatCalldata(caller, _dataArray[i]);
        }
        _multiCall(_destination, _optimisticSeconds, _recipients, payloads);
    }

    /**
     * @notice Calls a few system contracts using the same calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity[] memory _recipients,
        bytes memory _data
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        uint256 amount = _recipients.length;
        bytes[] memory payloads = new bytes[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            // Append calldata with (callOrigin, systemCaller) to form the payload
            payloads[i] = _formatCalldata(caller, _data);
        }
        _multiCall(_destination, _optimisticSeconds, _recipients, payloads);
    }

    /**
     * @notice Calls a single system contract a few times using the given calldata for each call.
     * See `systemCall` for details on system calls.
     * Note: tx will revert if any of the calls revert, guaranteeing
     * that either all calls succeed or none.
     */
    function systemMultiCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        SystemEntity _recipient,
        bytes[] memory _dataArray
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        uint256 amount = _dataArray.length;
        bytes[] memory payloads = new bytes[](amount);
        SystemEntity[] memory recipients = new SystemEntity[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            // Append calldata with (callOrigin, systemCaller) to form the payload
            payloads[i] = _formatCalldata(caller, _dataArray[i]);
            recipients[i] = _recipient;
        }
        _multiCall(_destination, _optimisticSeconds, recipients, payloads);
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
        return SystemCall.SYSTEM_ROUTER;
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
        // Deserialize the message into a series of system calls to perform
        bytes[] memory systemMessages = abi.decode(_message, (bytes[]));
        uint256 amount = systemMessages.length;
        for (uint256 i = 0; i < amount; ++i) {
            bytes29 _view = systemMessages[i].castToSystemCall();
            // Check that payload in a properly formatted system call
            require(_view.isSystemCall(), "Not a system call");
            // Route the system call to specified recipient
            _localSystemCall(_view.callRecipient(), _view.callPayload().clone(), _rootSubmittedAt);
        }
    }

    function _localSystemCall(
        uint8 _recipient,
        bytes memory _payload,
        uint256 _rootSubmittedAt
    ) internal {
        address recipient = _getSystemAddress(_recipient);
        require(recipient != address(0), "System Contract not set");
        // recipient.functionCall() calls recipient and bubbles the revert from the external call
        // We add `rootSubmittedAt` as the last argument.
        // (callOrigin, systemCaller) were added by the System Router on origin chain.
        // So the last arguments for the call are: (callOrigin, systemCaller, rootSubmittedAt)
        recipient.functionCall(abi.encodePacked(_payload, _rootSubmittedAt));
        // uint256 takes the full word of storage, so we can use encodePacked here w/o casting
    }

    function _remoteSystemCall(
        uint32 _destination,
        uint32 _optimisticSeconds,
        bytes[] memory _systemCalls
    ) internal {
        // Serialize the series of system calls into a byte string
        bytes memory message = abi.encode(_systemCalls);
        /**
         * @dev Origin will use SYSTEM_ROUTER as "sender" field for messages
         * sent by System Router.
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
        if (_destination == _localDomain()) {
            // Perform an on-chain multicall
            for (uint256 i = 0; i < amount; ++i) {
                /// @dev Passing current timestamp for consistency, see systemCall() for details
                _localSystemCall(uint8(_recipients[i]), _payloads[i], block.timestamp);
            }
        } else {
            // Perform a cross-chain multicall
            bytes[] memory systemCalls = new bytes[](amount);
            for (uint256 i = 0; i < amount; ++i) {
                systemCalls[i] = SystemCall.formatSystemCall(uint8(_recipients[i]), _payloads[i]);
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
         * 2. (callOrigin, systemCaller) are the following two arguments:
         *    - callOrigin is local domain
         *    - systemCaller is `_caller`
         * ====== ENCODED ON REMOTE CHAIN ======
         * 3. Root timestamp is the last argument, and will be appended
         * before the call on destination chain, see _localSystemCall()
         */
        return abi.encodePacked(_data, abi.encode(_localDomain(), _caller));
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
