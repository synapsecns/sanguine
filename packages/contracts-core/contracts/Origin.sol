// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Internal Imports ============
import "./Version.sol";
import "./libs/Message.sol";
import "./libs/SystemMessage.sol";
import { LocalDomainContext } from "./context/LocalDomainContext.sol";
import { OriginEvents } from "./events/OriginEvents.sol";
import { OriginHub } from "./hubs/OriginHub.sol";
// ============ External Imports ============
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title Origin
 * @author Illusory Systems Inc.
 * @notice Accepts messages to be dispatched to remote chains and
 * constructs a Merkle tree of the messages.
 * Notaries are signing the attestations of the Merkle tree's root state (aka merkle state),
 * which are broadcasted to Destination, where the merkle root is used for proving that
 * the message has been indeed dispatched on Origin.
 * Origin accepts submissions of fraudulent signatures by the Notary,
 * directly or in the form of a Guard's Fraud report on such an attestation,
 * and slashes the Notary in this case.
 * Origin accepts submissions of fraudulent signatures by the Guard in the form
 * of a Guard's report with said signature and slashes Guard in that case.
 */
contract Origin is OriginEvents, OriginHub, LocalDomainContext, Version0_0_1 {
    using TipsLib for bytes;
    using TipsLib for Tips;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Maximum bytes per message = 2 KiB
    // (somewhat arbitrarily set to begin)
    uint256 public constant MAX_MESSAGE_BODY_BYTES = 2 * 2**10;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // gap for upgrade safety
    uint256[50] private __GAP; //solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             CONSTRUCTOR                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) LocalDomainContext(_domain) {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __SystemContract_initialize();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          EXTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Dispatch the message to the destination domain & recipient
     * @dev Format the message, insert its hash into Merkle tree,
     * enqueue the new Merkle root, and emit `Dispatch` event with message information.
     * @param _destination      Domain of destination chain
     * @param _recipient        Address of recipient on destination chain as bytes32
     * @param _messageBody      Raw bytes content of message
     */
    function dispatch(
        uint32 _destination,
        bytes32 _recipient,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    )
        external
        payable
        // TODO: enable Guards check once Go tests are updated
        // haveActiveGuard
        haveActiveNotary(_destination)
        returns (uint32 messageNonce, bytes32 messageHash)
    {
        // TODO: add unit tests covering return values
        require(_messageBody.length <= MAX_MESSAGE_BODY_BYTES, "msg too long");
        // Check: tips payload is correctly formatted
        Tips tips = _tips.castToTips();
        // Check: total tips value matches msg.value
        require(tips.totalTips() == msg.value, "!tips: totalTips");
        // Latest nonce (i.e. "last message" nonce) is current amount of leaves in the tree.
        // Message nonce is the amount of leaves after the leaf insertion
        messageNonce = nonce(_destination) + 1;
        // format the message into packed bytes
        bytes memory message = MessageLib.formatMessage({
            _origin: _localDomain(),
            _sender: _checkForSystemRouter(_recipient),
            _nonce: messageNonce,
            _destination: _destination,
            _recipient: _recipient,
            _optimisticSeconds: _optimisticSeconds,
            _tips: _tips,
            _messageBody: _messageBody
        });
        messageHash = keccak256(message);
        // insert the hashed message into the Merkle tree
        bytes32 newRoot = _insertMessage(_destination, messageNonce, messageHash);
        // Emit Dispatch event with message information
        // note: leaf index in the tree is messageNonce - 1, meaning we don't need to emit that
        emit Dispatch(messageHash, messageNonce, _destination, newRoot, _tips, message);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Slash the fraudulent Agent.
     * @dev Called when agent fraud is proven.
     * @param _domain   Domain where the reported Agent is active
     * @param _account  Address of the fraudulent Agent
     * @param _guard    Guard who reported the fraudulent Agent [address(0) if not a Guard report]
     */
    function _slashAgent(
        uint32 _domain,
        address _account,
        address _guard
    ) internal override {
        _removeAgent(_domain, _account);
        if (_domain == 0) {
            emit GuardSlashed({ guard: _account, reporter: msg.sender });
        } else {
            emit NotarySlashed({ notary: _account, guard: _guard, reporter: msg.sender });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns adjusted "sender" field.
     * @dev By default, "sender" field is msg.sender address casted to bytes32.
     * However, if SYSTEM_ROUTER is used for "recipient" field, and msg.sender is SystemRouter,
     * SYSTEM_ROUTER is also used as "sender" field.
     * Note: tx will revert if anyone but SystemRouter uses SYSTEM_ROUTER as the recipient.
     */
    function _checkForSystemRouter(bytes32 _recipient) internal view returns (bytes32 sender) {
        if (_recipient != SystemMessageLib.SYSTEM_ROUTER) {
            sender = TypeCasts.addressToBytes32(msg.sender);
            /**
             * @dev Note: SYSTEM_ROUTER has only the highest 12 bytes set,
             * whereas TypeCasts.addressToBytes32 sets only the lowest 20 bytes.
             * Thus, in this branch: sender != SystemMessageLib.SYSTEM_ROUTER
             */
        } else {
            // Check that SystemRouter specified SYSTEM_ROUTER as recipient, revert otherwise.
            _assertSystemRouter();
            // Adjust "sender" field for correct processing on remote chain.
            sender = SystemMessageLib.SYSTEM_ROUTER;
        }
    }
}
