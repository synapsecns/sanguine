// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IFastBridgeV2} from "./interfaces/IFastBridgeV2.sol";
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

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════
}
