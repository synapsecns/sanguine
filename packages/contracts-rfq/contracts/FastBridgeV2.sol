// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

// ════════════════════════════════════════════════ INTERFACES ═════════════════════════════════════════════════════

import {IFastBridge} from "./interfaces/IFastBridge.sol";
import {IFastBridgeV2} from "./interfaces/IFastBridgeV2.sol";
import {IFastBridgeV2Errors} from "./interfaces/IFastBridgeV2Errors.sol";
import {IZapRecipient} from "./interfaces/IZapRecipient.sol";

// ═════════════════════════════════════════════ INTERNAL IMPORTS ══════════════════════════════════════════════════

import {AdminV2} from "./AdminV2.sol";
import {BridgeTransactionV2Lib} from "./libs/BridgeTransactionV2.sol";
import {MulticallTarget} from "./utils/MulticallTarget.sol";

// ═════════════════════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════════════════════════

import {IERC20, SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/// @title FastBridgeV2
/// @notice Core component of the SynapseRFQ protocol, enabling Relayers (Solvers) to fulfill bridge requests.
/// Supports ERC20 and native gas tokens, along with the Zap feature for executing actions on the destination chain.
/// Users interact with the off-chain Quoter API to obtain a current quote for a bridge transaction.
/// They then submit the bridge request with the quote to this contract, depositing their assets in escrow.
/// Relayers can fulfill requests by relaying them to the destination chain and must prove fulfillment to claim funds.
/// Guards monitor proofs and can dispute discrepancies.
/// Users can reclaim funds by cancelling their requests if it has not been fulfilled within the specified deadline.
contract FastBridgeV2 is AdminV2, MulticallTarget, IFastBridgeV2, IFastBridgeV2Errors {
    using BridgeTransactionV2Lib for bytes;
    using SafeERC20 for IERC20;

    /// @notice The duration of the dispute period for relayed transactions.
    uint256 public constant DISPUTE_PERIOD = 30 minutes;

    /// @notice The minimum required time between transaction request and deadline.
    uint256 public constant MIN_DEADLINE_PERIOD = 30 minutes;

    /// @notice The maximum allowed length for zapData.
    uint256 public constant MAX_ZAP_DATA_LENGTH = 2 ** 16 - 1;

    /// @notice Maps transaction IDs to bridge details (status, destination chain ID, proof timestamp, and relayer).
    /// Note: this is only stored for transactions having local chain as the origin chain.
    mapping(bytes32 => BridgeTxDetails) public bridgeTxDetails;
    /// @notice Maps transaction IDs to relay details (block number, block timestamp, and relayer).
    /// Note: this is only stored for transactions having local chain as the destination chain.
    mapping(bytes32 => BridgeRelay) public bridgeRelayDetails;
    /// @notice Maps sender addresses to their unique bridge nonce.
    mapping(address => uint256) public senderNonces;

    /// @notice This variable is deprecated and should not be used.
    /// @dev Replaced by senderNonces.
    uint256 public immutable nonce = 0;

    /// @notice Initializes the FastBridgeV2 contract with the provided default admin,
    /// sets the default cancel delay, and records the deploy block number.
    constructor(address defaultAdmin) AdminV2(defaultAdmin) {}

    // ══════════════════════════════════════ EXTERNAL MUTABLE (USER FACING) ═══════════════════════════════════════════

    /// @inheritdoc IFastBridge
    function bridge(BridgeParams memory params) external payable {
        bridgeV2({
            params: params,
            paramsV2: BridgeParamsV2({
                quoteRelayer: address(0),
                quoteExclusivitySeconds: 0,
                quoteId: bytes(""),
                zapNative: 0,
                zapData: bytes("")
            })
        });
    }

    /// Note: this function is deprecated and will be removed in a future version.
    /// @dev Replaced by `cancel`.
    /// @inheritdoc IFastBridge
    function refund(bytes calldata request) external {
        cancelV2(request);
    }

    // ══════════════════════════════════════ EXTERNAL MUTABLE (AGENT FACING) ══════════════════════════════════════════

    /// @inheritdoc IFastBridge
    function relay(bytes calldata request) external payable {
        // `relay` override will validate the request.
        relayV2({request: request, relayer: msg.sender});
    }

    /// @inheritdoc IFastBridge
    function prove(bytes calldata request, bytes32 destTxHash) external {
        request.validateV2();
        proveV2({transactionId: keccak256(request), destTxHash: destTxHash, relayer: msg.sender});
    }

    /// @inheritdoc IFastBridgeV2
    function claimV2(bytes calldata request) external {
        // `claim` override will validate the request.
        claim({request: request, to: address(0)});
    }

    /// @inheritdoc IFastBridge
    function dispute(bytes32 transactionId) external onlyRole(GUARD_ROLE) {
        // Aggregate the read operations from the same storage slot.
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        uint16 proverID = $.proverID;
        address disputedRelayer = $.proofRelayer;
        BridgeStatus status = $.status;
        uint40 proofBlockTimestamp = $.proofBlockTimestamp;

        // Can only dispute a RELAYER_PROVED transaction within the dispute period.
        if (status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (_timeSince(proofBlockTimestamp) > DISPUTE_PERIOD) {
            revert DisputePeriodPassed();
        }

        // Apply the timeout penalty to the prover that submitted the proof.
        // Note: this is a no-op if the prover has already been removed.
        _applyDisputePenaltyTime(proverID);

        // Update status to REQUESTED and delete the disputed proof details.
        // Note: these are storage writes.
        $.status = BridgeStatus.REQUESTED;
        $.proverID = 0;
        $.proofRelayer = address(0);
        $.proofBlockTimestamp = 0;

        emit BridgeProofDisputed(transactionId, disputedRelayer);
    }

    // ══════════════════════════════════════════════ EXTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @inheritdoc IFastBridge
    function canClaim(bytes32 transactionId, address relayer) external view returns (bool) {
        // The correct relayer can only claim a RELAYER_PROVED transaction after the dispute period.
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        if ($.status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if ($.proofRelayer != relayer) revert SenderIncorrect();

        return _timeSince($.proofBlockTimestamp) > DISPUTE_PERIOD;
    }

    /// @inheritdoc IFastBridge
    /// @dev This method is added to achieve backwards compatibility with decoding requests into V1 structs:
    /// - `zapNative` is partially reported as a zero/non-zero flag
    /// - `zapData` is ignored
    /// In order to process all kinds of requests use getBridgeTransactionV2 instead.
    function getBridgeTransaction(bytes calldata request) external view returns (BridgeTransaction memory) {
        // Try decoding into V2 struct first. This will revert if V1 struct is passed.
        try this.getBridgeTransactionV2(request) returns (BridgeTransactionV2 memory txV2) {
            // Note: we entirely ignore the zapData field, as it was not present in V1.
            return BridgeTransaction({
                originChainId: txV2.originChainId,
                destChainId: txV2.destChainId,
                originSender: txV2.originSender,
                destRecipient: txV2.destRecipient,
                originToken: txV2.originToken,
                destToken: txV2.destToken,
                originAmount: txV2.originAmount,
                destAmount: txV2.destAmount,
                originFeeAmount: txV2.originFeeAmount,
                sendChainGas: txV2.zapNative != 0,
                deadline: txV2.deadline,
                nonce: txV2.nonce
            });
        } catch {
            // Fallback to V1 struct.
            return abi.decode(request, (BridgeTransaction));
        }
    }

    /// @inheritdoc IFastBridgeV2
    function getBridgeTransactionV2(bytes calldata request) external pure returns (BridgeTransactionV2 memory) {
        request.validateV2();
        return BridgeTransactionV2Lib.decodeV2(request);
    }

    // ═══════════════════════════════════════ PUBLIC MUTABLE (USER FACING) ════════════════════════════════════════════

    /// @inheritdoc IFastBridgeV2
    function bridgeV2(BridgeParams memory params, BridgeParamsV2 memory paramsV2) public payable {
        // If relayer exclusivity is not intended for this bridge, set exclusivityEndTime to static zero.
        // Otherwise, set exclusivity to expire at the current block ts offset by quoteExclusivitySeconds.
        int256 exclusivityEndTime = 0;
        if (paramsV2.quoteRelayer != address(0)) {
            exclusivityEndTime = int256(block.timestamp) + paramsV2.quoteExclusivitySeconds;
        }
        _validateBridgeParams(params, paramsV2, exclusivityEndTime);

        // Track the amount of origin token owed to protocol.
        uint256 originAmount = params.originAmount;
        uint256 protocolFeeAmount = 0;
        if (protocolFeeRate > 0) {
            protocolFeeAmount = (originAmount * protocolFeeRate) / FEE_BPS;
            // The Relayer filling this request will be paid the originAmount after fees.
            // Note: the protocol fees will be accumulated only when the Relayer claims the origin collateral.
            originAmount -= protocolFeeAmount;
        }

        // Hash the bridge request and set the initial status to REQUESTED.
        bytes memory request = BridgeTransactionV2Lib.encodeV2(
            BridgeTransactionV2({
                originChainId: uint32(block.chainid),
                destChainId: params.dstChainId,
                originSender: params.sender,
                destRecipient: params.to,
                originToken: params.originToken,
                destToken: params.destToken,
                originAmount: originAmount,
                destAmount: params.destAmount,
                originFeeAmount: protocolFeeAmount,
                deadline: params.deadline,
                // Increment the sender's nonce on every bridge.
                nonce: senderNonces[params.sender]++,
                exclusivityRelayer: paramsV2.quoteRelayer,
                // We checked exclusivityEndTime to be in range [0 .. params.deadline] above, so can safely cast.
                exclusivityEndTime: uint256(exclusivityEndTime),
                zapNative: paramsV2.zapNative,
                zapData: paramsV2.zapData
            })
        );
        bytes32 transactionId = keccak256(request);
        // Note: the tx status will be updated throughout the tx lifecycle, while destChainId is set once here.
        bridgeTxDetails[transactionId].status = BridgeStatus.REQUESTED;
        bridgeTxDetails[transactionId].destChainId = params.dstChainId;

        // Emit the events before any external calls.
        emit BridgeRequested({
            transactionId: transactionId,
            sender: params.sender,
            request: request,
            destChainId: params.dstChainId,
            originToken: params.originToken,
            destToken: params.destToken,
            originAmount: originAmount,
            destAmount: params.destAmount,
            sendChainGas: paramsV2.zapNative != 0
        });
        emit BridgeQuoteDetails(transactionId, paramsV2.quoteId);

        // Transfer the tokens from the user as the last transaction action.
        address originToken = params.originToken;
        if (originToken != NATIVE_GAS_TOKEN) {
            // We need to take the full origin amount from the provided params (that includes `protocolFeeAmount`).
            uint256 amountToTake = params.originAmount;
            uint256 balanceBefore = IERC20(originToken).balanceOf(address(this));
            IERC20(originToken).safeTransferFrom(msg.sender, address(this), amountToTake);
            uint256 balanceAfter = IERC20(originToken).balanceOf(address(this));
            // Tokens with fees on transfer (or transferring more than requested) are not supported.
            if (balanceAfter != balanceBefore + amountToTake) revert AmountIncorrect();
        }
    }

    /// @inheritdoc IFastBridgeV2
    function cancelV2(bytes calldata request) public {
        // Decode the request and check that it could be cancelled.
        request.validateV2();
        bytes32 transactionId = keccak256(request);

        // Can only cancel a REQUESTED transaction after its deadline expires.
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        if ($.status != BridgeStatus.REQUESTED) revert StatusIncorrect();

        // Permissionless cancel is only allowed after `cancelDelay` on top of the deadline.
        uint256 deadline = request.deadline();
        if (!hasRole(CANCELER_ROLE, msg.sender)) deadline += cancelDelay;
        if (block.timestamp <= deadline) revert DeadlineNotExceeded();

        // Update status to REFUNDED.
        // Note: this is a storage write.
        $.status = BridgeStatus.REFUNDED;

        // Return the full amount (collateral + protocol fees) to the original sender.
        // The protocol fees are only accumulated when the transaction is claimed, so we don't need to update them here.
        address to = request.originSender();
        address token = request.originToken();
        uint256 amount = request.originAmount() + request.originFeeAmount();

        // Emit the event before any external calls.
        emit BridgeDepositRefunded(transactionId, to, token, amount);

        // Return the funds to the original sender as last transaction action.
        if (token == NATIVE_GAS_TOKEN) {
            Address.sendValue(payable(to), amount);
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

    // ═══════════════════════════════════════ PUBLIC MUTABLE (AGENT FACING) ═══════════════════════════════════════════

    /// @inheritdoc IFastBridgeV2
    function relayV2(bytes calldata request, address relayer) public payable {
        // Decode the request and check that it could be relayed.
        request.validateV2();
        bytes32 transactionId = keccak256(request);
        _validateRelayParams(request, transactionId, relayer);

        // Mark the bridge request as relayed by saving the relayer and the block details.
        bridgeRelayDetails[transactionId].blockNumber = uint48(block.number);
        bridgeRelayDetails[transactionId].blockTimestamp = uint48(block.timestamp);
        bridgeRelayDetails[transactionId].relayer = relayer;

        // Transfer tokens to recipient on destination chain and trigger Zap if requested.
        address to = request.destRecipient();
        address token = request.destToken();
        uint256 amount = request.destAmount();
        uint256 zapNative = request.zapNative();

        // Emit the event before any external calls.
        emit BridgeRelayed({
            transactionId: transactionId,
            relayer: relayer,
            to: to,
            originChainId: request.originChainId(),
            originToken: request.originToken(),
            destToken: token,
            originAmount: request.originAmount(),
            destAmount: amount,
            chainGasAmount: zapNative
        });

        // All state changes have been done at this point, can proceed to the external calls.
        // This follows the checks-effects-interactions pattern to mitigate potential reentrancy attacks.
        if (token == NATIVE_GAS_TOKEN) {
            // For the native gas token, additional zapNative is not allowed.
            if (zapNative != 0) revert ZapNativeNotSupported();
            // Check that the correct msg.value was sent.
            if (msg.value != amount) revert MsgValueIncorrect();
            // Don't do a native transfer yet: we will handle it alongside the Zap below.
        } else {
            // For ERC20s, we check that the correct msg.value was sent.
            if (msg.value != zapNative) revert MsgValueIncorrect();
            // We need to transfer the tokens from the Relayer to the recipient first before performing an
            // optional post-transfer Zap.
            IERC20(token).safeTransferFrom(msg.sender, to, amount);
        }
        // At this point we have done:
        // - Transferred the requested amount of ERC20 tokens to the recipient.
        // At this point we have confirmed:
        // - For ERC20s: msg.value matches the requested zapNative amount.
        // - For the native gas token: msg.value matches the requested destAmount.
        // Remaining optional things to do:
        // - Forward the full msg.value to the recipient (if non-zero).
        // - Trigger a Zap (if zapData is present).
        bytes calldata zapData = request.zapData();
        if (zapData.length != 0) {
            // Zap Data is present: Zap has been requested by the recipient. Trigger it forwarding the full msg.value.
            _triggerZapWithChecks({recipient: to, token: token, amount: amount, zapData: zapData});
            // Note: if token has a fee on transfers, the recipient will have received less than `amount`.
            // This is a very niche edge case and should be handled by the recipient contract.
        } else if (msg.value != 0) {
            // Zap Data is missing, but msg.value was sent. This could happen in two different cases:
            // - Relay with the native gas token is happening.
            // - Relay with ERC20 is happening, with a `zapNative > 0` request.
            // In both cases, we need to transfer the full msg.value to the recipient.
            Address.sendValue(payable(to), msg.value);
        }
    }

    /// @inheritdoc IFastBridgeV2
    function proveV2(bytes32 transactionId, bytes32 destTxHash, address relayer) public {
        uint16 proverID = getActiveProverID(msg.sender);
        if (proverID == 0) revert ProverNotActive();
        // Can only prove a REQUESTED transaction.
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        if ($.status != BridgeStatus.REQUESTED) revert StatusIncorrect();

        // Update status to RELAYER_PROVED and store the proof details.
        // Note: these are storage writes.
        $.status = BridgeStatus.RELAYER_PROVED;
        $.proverID = proverID;
        $.proofBlockTimestamp = uint40(block.timestamp);
        $.proofRelayer = relayer;

        emit BridgeProofProvided(transactionId, relayer, destTxHash);
    }

    /// @inheritdoc IFastBridge
    function claim(bytes calldata request, address to) public {
        // Decode the request and check that it could be claimed.
        request.validateV2();
        bytes32 transactionId = keccak256(request);

        // Aggregate the read operations from the same storage slot.
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        address proofRelayer = $.proofRelayer;
        BridgeStatus status = $.status;
        uint40 proofBlockTimestamp = $.proofBlockTimestamp;

        // Can only claim a RELAYER_PROVED transaction after the dispute period.
        if (status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (_timeSince(proofBlockTimestamp) <= DISPUTE_PERIOD) revert DisputePeriodNotPassed();
        if (to == address(0)) {
            // Anyone could claim the funds to the proven relayer on their behalf.
            to = proofRelayer;
        } else if (proofRelayer != msg.sender) {
            // Only the proven relayer could specify an address to claim the funds to.
            revert SenderIncorrect();
        }

        // Update status to RELAYER_CLAIMED and transfer the origin collateral to the specified claim address.
        // Note: this is a storage write.
        $.status = BridgeStatus.RELAYER_CLAIMED;

        // Accumulate protocol fees if origin fee amount exists.
        address token = request.originToken();
        uint256 amount = request.originAmount();
        uint256 protocolFeeAmount = request.originFeeAmount();
        if (protocolFeeAmount > 0) protocolFees[token] += protocolFeeAmount;

        // Emit the event before any external calls.
        emit BridgeDepositClaimed(transactionId, proofRelayer, to, token, amount);

        // Complete the relayer claim as the last transaction action.
        if (token == NATIVE_GAS_TOKEN) {
            Address.sendValue(payable(to), amount);
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

    // ═══════════════════════════════════════════════ PUBLIC VIEWS ════════════════════════════════════════════════════

    /// @inheritdoc IFastBridgeV2
    function bridgeStatuses(bytes32 transactionId) public view returns (BridgeStatus status) {
        return bridgeTxDetails[transactionId].status;
    }

    /// @inheritdoc IFastBridgeV2
    function bridgeProofs(bytes32 transactionId) public view returns (uint96 timestamp, address relayer) {
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        timestamp = $.proofBlockTimestamp;
        relayer = $.proofRelayer;
    }

    /// @inheritdoc IFastBridgeV2
    function bridgeRelays(bytes32 transactionId) public view returns (bool) {
        // This transaction has been relayed if the relayer address is recorded.
        return bridgeRelayDetails[transactionId].relayer != address(0);
    }

    // ═════════════════════════════════════════════ INTERNAL METHODS ══════════════════════════════════════════════════

    /// @notice Calls the recipient's hook function with the specified zapData and validates
    /// the returned value.
    function _triggerZapWithChecks(address recipient, address token, uint256 amount, bytes calldata zapData) internal {
        // Call the recipient's hook function with the specified zapData, bubbling any revert messages.
        bytes memory returnData = Address.functionCallWithValue({
            target: recipient,
            data: abi.encodeCall(IZapRecipient.zap, (token, amount, zapData)),
            // Note: see `relay()` for reasoning behind passing msg.value.
            value: msg.value
        });

        // Explicit revert if no return data at all.
        if (returnData.length == 0) revert RecipientNoReturnValue();
        // Check that exactly a single return value was returned.
        if (returnData.length != 32) revert RecipientIncorrectReturnValue();
        // Return value should be abi-encoded hook function selector.
        if (bytes32(returnData) != bytes32(IZapRecipient.zap.selector)) {
            revert RecipientIncorrectReturnValue();
        }
    }

    /// @notice Calculates the time elapsed since a proof was submitted.
    /// @dev The proof.timestamp stores block timestamps as uint40 for gas optimization.
    /// _timeSince(proof) handles timestamp rollover when block.timestamp > type(uint40).max but
    /// proof.timestamp < type(uint40).max via an unchecked statement.
    /// @param proofBlockTimestamp The block timestamp when the proof was submitted.
    /// @return delta The time elapsed since proof submission.
    function _timeSince(uint40 proofBlockTimestamp) internal view returns (uint256 delta) {
        unchecked {
            delta = uint40(block.timestamp) - proofBlockTimestamp;
        }
    }

    /// @notice Validates all parameters required for a bridge transaction.
    /// @dev This function's complexity cannot be reduced due to the number of required checks,
    /// so we disable the code-complexity rule.
    // solhint-disable-next-line code-complexity
    function _validateBridgeParams(
        BridgeParams memory params,
        BridgeParamsV2 memory paramsV2,
        int256 exclusivityEndTime
    )
        internal
        view
    {
        address originToken = params.originToken;
        uint256 originAmount = params.originAmount;
        // Check V1 (legacy) params.
        if (params.dstChainId == block.chainid) revert ChainIncorrect();
        if (originAmount == 0 || params.destAmount == 0) revert AmountIncorrect();
        if (params.sender == address(0) || params.to == address(0)) revert ZeroAddress();
        if (originToken == address(0) || params.destToken == address(0)) revert ZeroAddress();
        if (params.deadline < block.timestamp + MIN_DEADLINE_PERIOD) revert DeadlineTooShort();

        // Check V2 params.
        if (paramsV2.zapData.length > MAX_ZAP_DATA_LENGTH) revert ZapDataLengthAboveMax();
        if (paramsV2.zapNative != 0 && params.destToken == NATIVE_GAS_TOKEN) {
            revert ZapNativeNotSupported();
        }

        // exclusivityEndTime must be in range [0 .. params.deadline].
        if (exclusivityEndTime < 0 || exclusivityEndTime > int256(params.deadline)) {
            revert ExclusivityParamsIncorrect();
        }

        // Check supplied msg.value.
        if (originToken == NATIVE_GAS_TOKEN) {
            // For the native gas token, we just need to check that the supplied msg.value is correct.
            if (msg.value != originAmount) revert MsgValueIncorrect();
        } else {
            // We don't allow non-zero `msg.value` to avoid extra funds from being stuck in FastBridgeV2.
            if (msg.value != 0) revert MsgValueIncorrect();
            // Throw an explicit error if the provided token address is not a contract.
            if (originToken.code.length == 0) revert TokenNotContract();
        }
    }

    /// @notice Validates all parameters required for a relay transaction.
    function _validateRelayParams(bytes calldata request, bytes32 transactionId, address relayer) internal view {
        if (relayer == address(0)) revert ZeroAddress();
        // Check that the transaction has not been relayed yet and is for the current chain.
        if (bridgeRelays(transactionId)) revert TransactionRelayed();
        if (request.destChainId() != block.chainid) revert ChainIncorrect();
        // Check that the deadline for relay to happen has not passed yet.
        if (block.timestamp > request.deadline()) revert DeadlineExceeded();
        // Check the exclusivity period, if it was specified and is still ongoing.
        address exclRelayer = request.exclusivityRelayer();
        if (exclRelayer != address(0) && exclRelayer != relayer && block.timestamp <= request.exclusivityEndTime()) {
            revert ExclusivityPeriodNotPassed();
        }
    }
}
