// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { Client } from "../client/Client.sol";
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
contract SystemRouter is Client, ISystemRouter {
    using Address for address;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;
    using SystemMessage for bytes;
    using SystemMessage for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Should be set to a reasonably high value to prevent forging of a system message.
     *      Optimistic period is enforced in the base contract: Client.handle()
     */
    uint32 internal _optimisticPeriod;

    // gap for upgrade safety
    uint256[49] private __GAP; //solhint-disable-line var-name-mixedcase

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
        address _destination,
        uint32 _optimisticSeconds
    ) Client(_origin, _destination) {
        localDomain = _localDomain;
        // TODO: Do we ever want to adjust this?
        // (the value should be the same across all chains)
        // Or could it be converted into immutable?

        // TODO: Do we ever want to have "faster" system messages
        // for "smaller" rapid adjustments?
        _optimisticPeriod = _optimisticSeconds;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Send System Message to one of the System Contracts on origin chain
     * @dev     Note: knowledge of recipient address is not required,
     *          routing will be done by SystemRouter on destination chain.
     *          Following call will be made on destination chain:
     *              recipient.func(originDomain, originSender, _data)
     *          Allowing recipient to check:
     *          - domain where remote system call originated
     *          - System contract type of the sender
     * @param _destination  Domain of destination chain
     * @param _recipient    System contract type of the recipient
     * @param _selector     Function to call on destination chain
     * @param _data         Data for calling recipient on destination chain
     */
    function remoteSystemCall(
        uint32 _destination,
        SystemContracts _recipient,
        bytes4 _selector,
        bytes memory _data
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemContracts caller = _getSystemCaller(msg.sender);
        bytes memory message = SystemMessage.formatSystemCall(
            uint8(_recipient),
            _formatCalldata(caller, _selector, _data)
        );
        /**
         * @dev Origin should recognize SystemRouter as the "true sender"
         *      and use SYSTEM_ROUTER address as "sender" instead. This enables not
         *      knowing SystemRouter address on remote chain in advance.
         */
        _send(_destination, Tips.emptyTips(), message);
    }

    /**
     * @notice  Call a System Contract on the local chain.
     * @dev     Only System contracts are allowed to call this function.
     *          Note: knowledge of recipient address is not required,
     *          routing will be done by SystemRouter on the local chain.
     *          Following call will be made on local chain:
     *              recipient.func(localDomain, localSender, _data)
     *          Allowing recipient to check:
     *          - domain where system call originated (local domain in this case)
     *          - System contract type of the sender
     * @param _recipient    System contract type of the recipient
     * @param _selector     Function to call on destination chain
     * @param _data         Data for calling recipient on destination chain
     */
    function localSystemCall(
        SystemContracts _recipient,
        bytes4 _selector,
        bytes memory _data
    ) external {
        /// @dev This will revert if msg.sender is not a system contract
        SystemContracts caller = _getSystemCaller(msg.sender);
        _localSystemCall(uint8(_recipient), _formatCalldata(caller, _selector, _data));
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           PUBLIC FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Returns optimistic period of the merkle root, used for
     *          proving messages to SystemRouter.
     *          All messages to remote chains will be sent with this period.
     *          Merkle root is checked to be at least this old (from time of submission)
     *          for all incoming messages: see Client.handle()
     */
    function optimisticSeconds() public view override returns (uint32) {
        return _optimisticPeriod;
    }

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
     * @dev Handles an incoming message. All security checks are done in Client.handle()
     */
    function _handle(
        uint32,
        uint32,
        bytes memory _message
    ) internal override {
        bytes29 message = _message.castToSystemMessage();
        require(message.isSystemMessage(), "Not a system message");
        (SystemMessage.MessageFlag messageType, bytes29 body) = message.unpackMessage();

        if (messageType == SystemMessage.MessageFlag.Call) {
            _localSystemCall(body.callRecipient(), body.callPayload().clone());
        } else if (messageType == SystemMessage.MessageFlag.Adjust) {
            // TODO: handle messages with instructions
            // to adjust some of the SystemRouter parameters
        }
    }

    function _localSystemCall(uint8 _recipient, bytes memory _payload) internal {
        address recipient = _getSystemRecipient(_recipient);
        require(recipient != address(0), "System Contract not set");
        // this will call recipient and bubble the revert from the external call
        recipient.functionCall(_payload);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _formatCalldata(
        SystemContracts _caller,
        bytes4 _selector,
        bytes memory _data
    ) internal view returns (bytes memory) {
        /**
         * @dev Payload for contract call is:
         * 1. Function selector
         * 2. (domain, caller) are first two arguments
         * 3. data is remaining arguments, abi-encoded
         */
        return abi.encodePacked(_selector, abi.encode(localDomain, _caller), _data);
    }

    function _getSystemCaller(address _caller) internal view returns (SystemContracts) {
        if (_caller == origin) return SystemContracts.Origin;
        if (_caller == destination) return SystemContracts.Destination;
        revert("Unauthorized caller");
    }

    function _getSystemRecipient(uint8 _recipient) internal view returns (address) {
        if (_recipient == uint8(SystemContracts.Origin)) return origin;
        if (_recipient == uint8(SystemContracts.Destination)) return destination;
        revert("Unknown recipient");
    }
}
