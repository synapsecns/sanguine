// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IFastBridge, IFastBridgeV2} from "./interfaces/IFastBridgeV2.sol";
import {UniversalTokenLib} from "./libs/UniversalToken.sol";

import {Admin} from "./Admin.sol";

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract FastBridgeV2 is Admin, IFastBridgeV2 {
    using SafeERC20 for IERC20;
    using UniversalTokenLib for address;

    /// @notice Dispute period for relayed transactions
    uint256 public constant DISPUTE_PERIOD = 30 minutes;

    /// @notice Prove period added to deadline period for proven transactions
    uint256 public constant PROVE_PERIOD = 60 minutes;

    /// @notice Minimum deadline period to relay a requested bridge transaction
    uint256 public constant MIN_DEADLINE_PERIOD = 30 minutes;

    /// @notice Status of the bridge tx on origin chain
    mapping(bytes32 => BridgeStatusV2) public bridgeStatuses;
    /// @notice Proof of relayed bridge tx on origin chain
    mapping(bytes32 => BridgeProof) public bridgeProofs;

    /// @dev Destination relayer for a given transaction
    mapping(bytes32 => address) internal _destinationRelayer;

    /// @dev to prevent replays
    uint256 public nonce;
    // @dev the block the contract was deployed at
    uint256 public immutable deployBlock;

    constructor(address _owner) Admin(_owner) {
        deployBlock = block.number;
    }

    // ═══════════════════════════════════════════════ ONLY RELAYER ════════════════════════════════════════════════════

    // ════════════════════════════════════════════════ ONLY GUARD ═════════════════════════════════════════════════════

    // ════════════════════════════════════════════════ USER-FACING ════════════════════════════════════════════════════

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IFastBridgeV2
    function bridgeRelays(bytes32 transactionId) external view returns (bool) {
        return _destinationRelayer[transactionId] != address(0);
    }

    /// @inheritdoc IFastBridgeV2
    function getTransactionRelayer(bytes32 transactionId) external view returns (address) {
        return _destinationRelayer[transactionId];
    }

    /// @inheritdoc IFastBridge
    function getBridgeTransaction(bytes memory request) public pure returns (BridgeTransaction memory) {
        return abi.decode(request, (BridgeTransaction));
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @notice Pulls a requested token from the user to the requested recipient.
    /// @dev Be careful of re-entrancy issues when msg.value > 0 and recipient != address(this)
    function _pullToken(address recipient, address token, uint256 amount) internal returns (uint256 amountPulled) {
        if (token != UniversalTokenLib.ETH_ADDRESS) {
            token.assertIsContract();
            // Record token balance before transfer
            amountPulled = IERC20(token).balanceOf(recipient);
            // Token needs to be pulled only if msg.value is zero
            // This way user can specify WETH as the origin asset
            IERC20(token).safeTransferFrom(msg.sender, recipient, amount);
            // Use the difference between the recorded balance and the current balance as the amountPulled
            amountPulled = IERC20(token).balanceOf(recipient) - amountPulled;
        } else {
            // Otherwise, we need to check that ETH amount matches msg.value
            if (amount != msg.value) revert FastBridge__MsgValueIncorrect();
            // Transfer value to recipient if not this address
            if (recipient != address(this)) token.universalTransfer(recipient, amount);
            // We will forward msg.value in the external call later, if recipient is not this contract
            amountPulled = msg.value;
        }
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @notice Calculates time since proof submitted
    /// @dev proof.timestamp stores casted uint96(block.timestamp) block timestamps for gas optimization
    ///      _timeSince(proof) can accommodate rollover case when block.timestamp > type(uint96).max but
    ///      proof.timestamp < type(uint96).max via unchecked statement
    /// @param proof The bridge proof
    /// @return delta Time delta since proof submitted
    function _timeSince(BridgeProof memory proof) internal view returns (uint256 delta) {
        unchecked {
            delta = uint96(block.timestamp) - proof.timestamp;
        }
    }
}
