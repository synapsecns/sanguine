// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ICoreWriter} from "./interfaces/ICoreWriter.sol";
import {ISynapseHyperCoreComposer} from "./interfaces/ISynapseHyperCoreComposer.sol";

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

/// @notice Moves SYN received by SynapseBridgeAdapter from HyperEVM to its linked HIP-1 asset.
/// @dev This contract is intentionally fixed to SYN's 18 EVM decimals and 8 HyperCore wei decimals.
/// It must be activated on HyperCore before it can submit a CoreWriter spot-send action.
contract SynapseHyperCoreComposer is ISynapseHyperCoreComposer {
    using SafeERC20 for IERC20;

    uint8 public constant EVM_DECIMALS = 18;
    uint8 public constant CORE_DECIMALS = 8;
    uint256 public constant AMOUNT_SCALE = 10 ** (EVM_DECIMALS - CORE_DECIMALS);

    address public constant CORE_WRITER = 0x3333333333333333333333333333333333333333;
    address public constant SPOT_BALANCE_PRECOMPILE = 0x0000000000000000000000000000000000000801;
    address public constant CORE_USER_EXISTS_PRECOMPILE = 0x0000000000000000000000000000000000000810;
    address public constant BASE_ASSET_BRIDGE = 0x2000000000000000000000000000000000000000;

    bytes4 public constant SPOT_SEND_HEADER = 0x0100_0006;

    address public immutable override adapter;
    address public immutable override token;
    uint64 public immutable coreIndex;
    address public immutable assetBridge;

    uint256 private reservationBlock;
    uint64 private reservedCoreAmount;

    event HyperCoreTransferQueued(
        address indexed recipient, uint256 evmAmount, uint64 coreAmount, uint64 indexed coreIndex
    );

    error SHCC__AmountExceedsUint64(uint256 coreAmount);
    error SHCC__AmountNotRepresentable(uint256 amount, uint256 scale);
    error SHCC__AssetBridgeCapacityInsufficient(uint256 available, uint256 required);
    error SHCC__CoreUserNotActivated(address user);
    error SHCC__InvalidTokenDecimals(uint8 actual, uint8 expected);
    error SHCC__OnlyAdapter(address caller);
    error SHCC__PrecompileReadFailed(address precompile);
    error SHCC__ZeroAddress();
    error SHCC__ZeroAmount();

    constructor(address adapter_, address token_, uint64 coreIndex_) {
        if (adapter_ == address(0) || token_ == address(0)) revert SHCC__ZeroAddress();
        uint8 tokenDecimals = IERC20Metadata(token_).decimals();
        if (tokenDecimals != EVM_DECIMALS) revert SHCC__InvalidTokenDecimals(tokenDecimals, EVM_DECIMALS);

        adapter = adapter_;
        token = token_;
        coreIndex = coreIndex_;
        assetBridge = address(uint160(uint256(uint160(BASE_ASSET_BRIDGE)) + coreIndex_));
    }

    /// @inheritdoc ISynapseHyperCoreComposer
    function bridgeToHyperCore(address recipient, uint256 amount) external override {
        if (msg.sender != adapter) revert SHCC__OnlyAdapter(msg.sender);
        if (recipient == address(0)) revert SHCC__ZeroAddress();
        if (amount == 0) revert SHCC__ZeroAmount();
        if (amount % AMOUNT_SCALE != 0) revert SHCC__AmountNotRepresentable(amount, AMOUNT_SCALE);

        uint256 coreAmount256 = amount / AMOUNT_SCALE;
        if (coreAmount256 > type(uint64).max) revert SHCC__AmountExceedsUint64(coreAmount256);
        // forge-lint: disable-next-line(unsafe-typecast)
        uint64 coreAmount = uint64(coreAmount256);

        _requireCoreUser(address(this));
        _requireCoreUser(recipient);

        uint64 available = _spotBalance(assetBridge, coreIndex);
        // The equality deliberately selects this HyperEVM block's reservation epoch; it is not a value comparison.
        // slither-disable-next-line incorrect-equality
        uint64 reserved = block.number == reservationBlock ? reservedCoreAmount : 0;
        uint256 required = uint256(reserved) + coreAmount;
        if (required > available) revert SHCC__AssetBridgeCapacityInsufficient(available, required);

        reservationBlock = block.number;
        // `required` is bounded by the uint64 `available` value above.
        // forge-lint: disable-next-line(unsafe-typecast)
        reservedCoreAmount = uint64(required);

        // Hyperliquid credits this contract's HyperCore account from the ERC20 Transfer event before
        // processing the following CoreWriter action in the same L1 block.
        IERC20(token).safeTransfer(assetBridge, amount);
        ICoreWriter(CORE_WRITER)
            .sendRawAction(abi.encodePacked(SPOT_SEND_HEADER, abi.encode(recipient, coreIndex, coreAmount)));

        // The immutable adapter is the only caller, reservation state is committed above, and SYN has no callback.
        // slither-disable-next-line reentrancy-events
        emit HyperCoreTransferQueued(recipient, amount, coreAmount, coreIndex);
    }

    function _requireCoreUser(address user) internal view {
        // HyperEVM exposes this precompile without a Solidity ABI. Raw staticcall lets us fail closed on malformed
        // data.
        // slither-disable-next-line low-level-calls
        (bool success, bytes memory result) = CORE_USER_EXISTS_PRECOMPILE.staticcall(abi.encode(user));
        if (!success || result.length != 32) revert SHCC__PrecompileReadFailed(CORE_USER_EXISTS_PRECOMPILE);
        if (!abi.decode(result, (bool))) revert SHCC__CoreUserNotActivated(user);
    }

    function _spotBalance(address user, uint64 tokenIndex) internal view returns (uint64 total) {
        // HyperEVM exposes this precompile without a Solidity ABI. Raw staticcall lets us fail closed on malformed
        // data.
        // slither-disable-next-line low-level-calls
        (bool success, bytes memory result) = SPOT_BALANCE_PRECOMPILE.staticcall(abi.encode(user, tokenIndex));
        if (!success || result.length != 96) revert SHCC__PrecompileReadFailed(SPOT_BALANCE_PRECOMPILE);
        (total,,) = abi.decode(result, (uint64, uint64, uint64));
    }
}
