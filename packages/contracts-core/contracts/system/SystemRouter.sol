// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import {ByteString, CallData} from "../libs/ByteString.sol";
import {SYSTEM_ROUTER} from "../libs/Constants.sol";
import {SystemMessage, SystemMessageLib} from "../libs/SystemMessage.sol";
import {SystemEntity} from "../libs/Structures.sol";
import {TipsLib} from "../libs/Tips.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {DomainContext} from "../context/DomainContext.sol";
import {InterfaceOrigin} from "../interfaces/InterfaceOrigin.sol";
import {InterfaceSystemRouter} from "../interfaces/InterfaceSystemRouter.sol";
import {Versioned} from "../Version.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @notice Router for cross-chain calls between system contracts (aka "System Calls").
 * This contract enables system contracts to issue a cross-chain call to another system contract,
 * without knowing the recipient address. Instead both sender and recipient are encoded as SystemEntity:
 *  - AgentManager
 *  - Destination
 *  - Origin
 * The system calls are performed by sending a "system message" via Origin-Destination contracts.
 * Note: System Router should be only used for remote calls.
 * System Contracts are supposed to know the local addresses for on-chain calls.
 * @dev Security considerations.
 *  - System Router only accepts calls from the system contracts, so it's not possible for anyone but
 *  the system contracts to initiate a system call.
 *  - System Contracts should expose functions for receveing cross-chain system calls:
 *  `function foo(uint256 proofMaturity, uint32 origin, SystemEntity sender, *args)`
 *  - The first three arguments in these functions are the security arguments, which will be filled
 *  by SystemRouter on the destination chain.
 *  - Such functions should be protected with `onlySystemRouter` modifier.
 *  - The system contract should verify the security arguments before handling the remaining `*args`.
 *  - To perform a system call, contract should omit the security arguments when abi-encoding the payload
 *  `payload = abi.encodeWithSelector(foo.selector, *args);`
 */
contract SystemRouter is DomainContext, InterfaceSystemRouter, Versioned {
    using Address for address;
    using ByteString for bytes;
    using SystemMessageLib for bytes;

    address public immutable agentManager;
    address public immutable destination;
    address public immutable origin;

    // ════════════════════════════════════════════════ CONSTRUCTOR ════════════════════════════════════════════════════

    constructor(uint32 domain, address origin_, address destination_, address agentManager_)
        DomainContext(domain)
        Versioned("0.0.3")
    {
        origin = origin_;
        destination = destination_;
        agentManager = agentManager_;
    }

    // ════════════════════════════════════════════ EXTERNAL FUNCTIONS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceSystemRouter
    function receiveSystemMessage(uint32 origin_, uint32, uint256 proofMaturity, bytes memory body) external {
        // Only Destination can deliver messages
        require(msg.sender == destination, "SystemRouter: !destination");
        // TODO: figure out if we need nonce to be passed here
        // This will revert if message body is not a system message
        SystemMessage systemMessage = body.castToSystemMessage();
        // Destination chain: set (proofMaturity, origin, systemSender) values for system message
        bytes memory prefix = abi.encode(proofMaturity, origin_, systemMessage.sender());
        // Add the (proofMaturity, origin, systemSender) values to the calldata
        bytes memory payload = systemMessage.callData().addPrefix(prefix);
        _callSystemRecipient(systemMessage.recipient(), payload);
    }

    /// @inheritdoc InterfaceSystemRouter
    function systemCall(uint32 destination_, uint32 optimisticPeriod, SystemEntity recipient, bytes memory payload)
        external
    {
        require(destination_ != localDomain, "Must be a remote destination");
        /// @dev This will revert if msg.sender is not a system contract
        SystemEntity sender = _getSystemEntity(msg.sender);
        // Construct the System Message: no security arguments are added on origin chain
        bytes memory body = SystemMessageLib.formatSystemMessage(sender, recipient, payload);
        InterfaceOrigin(origin).sendSystemMessage(destination_, optimisticPeriod, body);
    }

    // ════════════════════════════════════════════ INTERNAL FUNCTIONS ═════════════════════════════════════════════════

    /// @dev Calls a local System Contract, using calldata from the received system message.
    function _callSystemRecipient(SystemEntity systemRecipient, bytes memory payload) internal {
        address recipient = _getSystemAddress(systemRecipient);
        require(recipient != address(0), "System Contract not set");
        // recipient.functionCall() calls recipient and bubbles the revert from the external call
        recipient.functionCall(payload);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns a corresponding System Entity for a given message sender.
    function _getSystemEntity(address sender) internal view returns (SystemEntity) {
        if (sender == origin) return SystemEntity.Origin;
        if (sender == destination) return SystemEntity.Destination;
        if (sender == agentManager) return SystemEntity.AgentManager;
        revert("Unauthorized caller");
    }

    /// @dev Returns a corresponding address for a given system recipient.
    function _getSystemAddress(SystemEntity entity) internal view returns (address) {
        // Possible SystemEntity values: AgentManager / Destination / Origin
        if (entity == SystemEntity.AgentManager) {
            return agentManager;
        } else if (entity == SystemEntity.Destination) {
            return destination;
        } else {
            return origin;
        }
    }
}
