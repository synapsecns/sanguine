// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import { ByteString, CallData } from "../libs/ByteString.sol";
import { SYSTEM_ROUTER } from "../libs/Constants.sol";
import { SystemMessage, SystemMessageLib } from "../libs/SystemMessage.sol";
import { SystemEntity } from "../libs/Structures.sol";
import { TipsLib } from "../libs/Tips.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { BasicClient } from "../client/BasicClient.sol";
import { DomainContext } from "../context/DomainContext.sol";
import { InterfaceSystemRouter } from "../interfaces/InterfaceSystemRouter.sol";
import { Versioned } from "../Version.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @notice Router for calls between system contracts (aka "System Calls").
 * SystemRouter makes it possible to perform a call from one system contract to another
 * without knowing the recipient address. This works for both on-chain calls, when caller and
 * recipient are deployed on the same chain, and for cross-chain calls, when caller and
 * recipient are deployed on different chains.
 *
 * SystemRouter allows both calls and "multi calls". Multicall performs a series of calls,
 * calling requested recipients one by one, supplying the requested calldata. The whole multicall
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
 * 1. Such functions should have the same first three arguments:
 * - foo(uint256 rootSubmittedAt, uint32 callOrigin, SystemEntity systemCaller, <...>)
 * These arguments are filled by the System Routers on origin and destination chain. This allows
 * the recipient to set the restrictions for receiving the call in a very granular way.
 * To perform a call, use any values for `(rootSubmittedAt, callOrigin,systemCaller)`
 * i.e. `payload = abi.encodeWithSelector(foo.selector, 0, 0, 0, <...>);`
 * These values are overwritten, so using non-zero values will not achieve anything.
 * 2. Guard such function with `onlySystemRouter` modifier to prevent unauthorized direct calls.
 * Guard function with additional modifiers based on `rootSubmittedAt`, `origin` and `caller`.
 * `rootSubmittedAt` based modifier is a must for receiving cross-chain system calls. Any Notary
 * can potentially commit fraud, and try to execute an arbitrary message, including
 * a "message to System Router". By enforcing a minimum optimistic latency for the recipient this
 * attack can be mitigated, assuming there is at least one honest Guard willing to report the fraud.
 */
contract SystemRouter is DomainContext, BasicClient, InterfaceSystemRouter, Versioned {
    using Address for address;
    using ByteString for bytes;
    using SystemMessageLib for bytes;
    using SystemMessageLib for SystemMessage;

    /**
     * @dev System entity initiates a system call with given calldata.
     *      Entity provides calldata = (foo.selector, security arguments, remaining arguments).
     *      Provided security arguments are overwritten by System Routers with the correct ones.
     *      Full calldata for the performed call on destination chain is:
     * ============   GIVEN ENTITY DATA             ============
     * 1. Selector from given calldata.
     * ============   FILLED ON DESTINATION CHAIN   ============
     * 2. Root timestamp is the first security argument filled by SystemRouter:
     * - rootSubmittedAt: time when merkle root used for proving a message
     * with a system call was submitted.
     * For local system calls: rootSubmittedAt = block.timestamp
     * ============   FILLED ON ORIGIN CHAIN        ============
     * 3. (callOrigin, systemCaller) are second and third security arguments filled by SystemRouter:
     * - callOrigin: domain where system call originated
     * - systemCaller: entity that initiated the system call on origin chain
     * ============   GIVEN ENTITY DATA             ============
     * 4. Remaining arguments from given calldata.
     *
     * As the result, following call is performed: `recipient.foo(securityArgs, remainingArgs)`
     * - `securityArgs` part is filled collectively by System Routers on origin, destination chains
     * - `remainingArgs` part is provided by the system entity on origin chain
     */

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              IMMUTABLES                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    address public immutable agentManager;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    constructor(
        uint32 domain,
        address origin_,
        address destination_,
        address agentManager_
    ) BasicClient(origin_, destination_) DomainContext(domain) Versioned("0.0.3") {
        agentManager = agentManager_;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSystemRouter
    function systemCall(
        uint32 destination_,
        uint32 optimisticSeconds,
        SystemEntity recipient,
        bytes memory payload
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        // To generalize things, a system call is always a multicall.
        // In case of a "single system call", this is a multicall with exactly one call inside.
        SystemEntity[] memory recipients = new SystemEntity[](1);
        CallData[] memory callDataArray = new CallData[](1);
        recipients[0] = recipient;
        callDataArray[0] = payload.castToCallData();
        _multiCall(caller, destination_, optimisticSeconds, recipients, callDataArray);
    }

    /// @inheritdoc InterfaceSystemRouter
    function systemMultiCall(
        uint32 destination_,
        uint32 optimisticSeconds,
        SystemEntity[] memory recipients,
        bytes[] memory payloadArray
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        uint256 amount = recipients.length;
        CallData[] memory callDataArray = new CallData[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            // Do a cast to a CallData view for every element
            callDataArray[i] = payloadArray[i].castToCallData();
        }
        _multiCall(caller, destination_, optimisticSeconds, recipients, callDataArray);
    }

    /// @inheritdoc InterfaceSystemRouter
    function systemMultiCall(
        uint32 destination_,
        uint32 optimisticSeconds,
        SystemEntity[] memory recipients,
        bytes memory payload
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        uint256 amount = recipients.length;
        CallData[] memory callDataArray = new CallData[](amount);
        CallData callData = payload.castToCallData();
        for (uint256 i = 0; i < amount; ++i) {
            // `callData` is never modified, so we can reuse the same memory view here
            callDataArray[i] = callData;
        }
        _multiCall(caller, destination_, optimisticSeconds, recipients, callDataArray);
    }

    /// @inheritdoc InterfaceSystemRouter
    function systemMultiCall(
        uint32 destination_,
        uint32 optimisticSeconds,
        SystemEntity recipient,
        bytes[] memory payloadArray
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity caller = _getSystemEntity(msg.sender);
        uint256 amount = payloadArray.length;
        CallData[] memory callDataArray = new CallData[](amount);
        SystemEntity[] memory recipients = new SystemEntity[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            // Every call recipient is the same
            recipients[i] = recipient;
            // Do a cast to a CallData view for every element
            callDataArray[i] = payloadArray[i].castToCallData();
        }
        _multiCall(caller, destination_, optimisticSeconds, recipients, callDataArray);
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
        return SYSTEM_ROUTER;
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
        uint256 rootSubmittedAt,
        bytes memory content
    ) internal override {
        // TODO: use TypedMemView for encoding/decoding instead
        // Deserialize the message into a series of system calls to perform
        bytes[] memory systemMessages = abi.decode(content, (bytes[]));
        uint256 amount = systemMessages.length;
        // Received a message containing a remote system call, use the corresponding prefix
        bytes29 prefix = _prefixReceiveCall(rootSubmittedAt).castToRawBytes();
        for (uint256 i = 0; i < amount; ++i) {
            SystemMessage systemMessage = systemMessages[i].castToSystemMessage();
            // Route the system call to specified recipient
            _localSystemCall(systemMessage.callRecipient(), systemMessage.callData(), prefix);
        }
    }

    /**
     * @notice Routes a system call to a local System Contract, using provided
     * calldata, and abi-encoded arguments to add as the prefix.
     * @dev Suppose following values were passed:
     * - recipient: System Contract to call
     * - callData = abi.encodeWithSelector(foo.selector, a, b, c, d, e, f);
     * - prefix = abi.encode(x, y, z)
     * - (a, b, c) types match (x, y, z) types, and they are all static
     * Following call will be performed:
     * - recipient.foo(x, y, z, d, e, f);
     */
    function _localSystemCall(
        uint8 systemRecipient,
        CallData callData,
        bytes29 prefix
    ) internal {
        // We adjust the first arguments for the call using the given `prefix`.
        // For remote system calls:
        // - (rootSubmittedAt, callOrigin, systemCaller) are adjusted on origin chain
        // - (rootSubmittedAt) is readjusted on destination chain
        // For local system calls:
        // - (rootSubmittedAt, callOrigin, systemCaller) is adjusted
        // That gives us the following first three arguments for the system call (remote or local):
        // - (rootSubmittedAt, callOrigin, systemCaller)
        address recipient = _getSystemAddress(systemRecipient);
        require(recipient != address(0), "System Contract not set");
        // recipient.functionCall() calls recipient and bubbles the revert from the external call
        recipient.functionCall(SystemMessageLib.formatAdjustedCallData(callData, prefix));
    }

    /**
     * @notice Performs the "sending part" of a remote system multicall.
     * @param destination_          Destination domain where system multicall will be performed
     * @param optimisticSeconds     Optimistic period for the executing the system multicall
     * @param systemCalls           List of system calls to execute on destination chain
     */
    function _remoteSystemCall(
        uint32 destination_,
        uint32 optimisticSeconds,
        bytes[] memory systemCalls
    ) internal {
        // TODO: use TypedMemView for encoding/decoding instead
        // Serialize the series of system calls into a byte string
        bytes memory content = abi.encode(systemCalls);
        /**
         * @dev Origin will use SYSTEM_ROUTER as "sender" field for messages
         * sent by System Router.
         */
        _send(destination_, optimisticSeconds, TipsLib.emptyTips(), content);
    }

    /**
     * @notice Performs a system multicall with given parameters.
     * @dev `caller` is derived from msg.sender
     * @param caller                System entity that initiated the system multicall
     * @param destination_          Destination domain where system multicall will be performed
     * @param optimisticSeconds     Optimistic period for the executing the system multicall
     * @param recipients            List of system entities to route the system call to
     * @param callDataArray         List of memory views over calldata
     */
    function _multiCall(
        SystemEntity caller,
        uint32 destination_,
        uint32 optimisticSeconds,
        SystemEntity[] memory recipients,
        CallData[] memory callDataArray
    ) internal {
        uint256 amount = recipients.length;
        // Performing a system call on origin chain,
        // Get a prefix for performing the call on origin chain, use the corresponding prefix
        bytes29 prefix = _prefixPerformCall(caller).castToRawBytes();
        if (destination_ == localDomain) {
            // Performing a local system multicall
            for (uint256 i = 0; i < amount; ++i) {
                _localSystemCall(uint8(recipients[i]), callDataArray[i], prefix);
            }
        } else {
            // Performing a remote system multicall
            bytes[] memory systemMessages = new bytes[](amount);
            for (uint256 i = 0; i < amount; ++i) {
                systemMessages[i] = SystemMessageLib.formatSystemMessage({
                    systemRecipient: uint8(recipients[i]),
                    callData_: callDataArray[i],
                    prefix: prefix
                });
            }
            _remoteSystemCall(destination_, optimisticSeconds, systemMessages);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @notice Returns a corresponding System Entity for a given caller.
    function _getSystemEntity(address caller) internal view returns (SystemEntity) {
        if (caller == origin) return SystemEntity.Origin;
        if (caller == destination) return SystemEntity.Destination;
        if (caller == agentManager) return SystemEntity.AgentManager;
        revert("Unauthorized caller");
    }

    /// @notice Returns a corresponding address for a given system recipient.
    function _getSystemAddress(uint8 entity) internal view returns (address) {
        if (entity == uint8(SystemEntity.Origin)) return origin;
        if (entity == uint8(SystemEntity.Destination)) return destination;
        if (entity == uint8(SystemEntity.AgentManager)) return agentManager;
        revert("Unknown recipient");
    }

    /// @notice Returns prefix with the security arguments
    /// for making a system call on origin chain.
    function _prefixPerformCall(SystemEntity caller) internal view returns (bytes memory) {
        // Origin chain: adjust (rootSubmittedAt, callOrigin, systemCaller)
        return abi.encode(block.timestamp, localDomain, caller);
        // Passing current timestamp for consistency
        // For a cross-chain call (rootSubmittedAt) will be later adjusted on destination chain
    }

    /// @notice Returns prefix with the security arguments
    /// for receiving a remote system call on destination chain.
    function _prefixReceiveCall(uint256 rootSubmittedAt) internal pure returns (bytes memory) {
        // Destination chain: adjust (rootSubmittedAt)
        // (callOrigin, systemCaller) were adjusted on origin chain, no need to touch these
        return abi.encode(rootSubmittedAt);
    }
}
