// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {BridgeTransactionV2Lib} from "./libs/BridgeTransactionV2.sol";

import {AdminV2} from "./AdminV2.sol";
import {IFastBridge} from "./interfaces/IFastBridge.sol";
import {IFastBridgeV2} from "./interfaces/IFastBridgeV2.sol";
import {IFastBridgeV2Errors} from "./interfaces/IFastBridgeV2Errors.sol";
import {IZapRecipient} from "./interfaces/IZapRecipient.sol";

import {MulticallTarget} from "./utils/MulticallTarget.sol";

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/// @notice FastBridgeV2 is a contract for bridging tokens across chains.
contract FastBridgeV2 is AdminV2, MulticallTarget, IFastBridgeV2, IFastBridgeV2Errors {
    using BridgeTransactionV2Lib for bytes;
    using SafeERC20 for IERC20;

    /// @notice Dispute period for relayed transactions
    uint256 public constant DISPUTE_PERIOD = 30 minutes;

    /// @notice Minimum deadline period to relay a requested bridge transaction
    uint256 public constant MIN_DEADLINE_PERIOD = 30 minutes;

    /// @notice Maximum length of accepted zapData
    uint256 public constant MAX_ZAP_DATA_LENGTH = 2 ** 16 - 1;

    /// @notice Status of the bridge tx on origin chain
    mapping(bytes32 => BridgeTxDetails) public bridgeTxDetails;
    /// @notice Relay details on destination chain
    mapping(bytes32 => BridgeRelay) public bridgeRelayDetails;
    /// @notice Unique bridge nonces tracked per originSender
    mapping(address => uint256) public senderNonces;

    /// @notice This is deprecated and should not be used.
    /// @dev Replaced by senderNonces
    uint256 public immutable nonce = 0;
    /// @notice the block the contract was deployed at
    uint256 public immutable deployBlock;

    constructor(address _owner) AdminV2(_owner) {
        deployBlock = block.number;
    }

    /// @inheritdoc IFastBridge
    function bridge(BridgeParams memory params) external payable {
        bridge({
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

    /// @inheritdoc IFastBridge
    function relay(bytes calldata request) external payable {
        // relay override will validate the request
        relay({request: request, relayer: msg.sender});
    }

    /// @inheritdoc IFastBridge
    function prove(bytes calldata request, bytes32 destTxHash) external {
        request.validateV2();
        prove({transactionId: keccak256(request), destTxHash: destTxHash, relayer: msg.sender});
    }

    /// @inheritdoc IFastBridgeV2
    function claim(bytes calldata request) external {
        // claim override will validate the request
        claim({request: request, to: address(0)});
    }

    /// @inheritdoc IFastBridge
    function dispute(bytes32 transactionId) external onlyRole(GUARD_ROLE) {
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        // Aggregate the read operations from the same storage slot
        address disputedRelayer = $.proofRelayer;
        BridgeStatus status = $.status;
        uint56 proofBlockTimestamp = $.proofBlockTimestamp;
        // Can only dispute a RELAYER_PROVED transaction within the dispute period
        if (status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (_timeSince(proofBlockTimestamp) > DISPUTE_PERIOD) {
            revert DisputePeriodPassed();
        }
        // Update status to REQUESTED and delete the disputed proof details
        // Note: these are storage writes
        $.status = BridgeStatus.REQUESTED;
        $.proofRelayer = address(0);
        $.proofBlockTimestamp = 0;

        emit BridgeProofDisputed(transactionId, disputedRelayer);
    }

    /// Note: this function is deprecated and will be removed in a future version.
    /// @inheritdoc IFastBridge
    function refund(bytes calldata request) external {
        cancel(request);
    }

    /// @inheritdoc IFastBridge
    function canClaim(bytes32 transactionId, address relayer) external view returns (bool) {
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        // The correct relayer can only claim a RELAYER_PROVED transaction after the dispute period
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
        // Try decoding into V2 struct first. This will revert if V1 struct is passed
        try this.getBridgeTransactionV2(request) returns (BridgeTransactionV2 memory txV2) {
            // Note: we entirely ignore the zapData field, as it was not present in V1
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
            // Fallback to V1 struct
            return abi.decode(request, (BridgeTransaction));
        }
    }

    /// @inheritdoc IFastBridgeV2
    function getBridgeTransactionV2(bytes calldata request) external pure returns (BridgeTransactionV2 memory) {
        request.validateV2();
        return BridgeTransactionV2Lib.decodeV2(request);
    }

    /// @inheritdoc IFastBridgeV2
    function bridge(BridgeParams memory params, BridgeParamsV2 memory paramsV2) public payable {
        int256 exclusivityEndTime = 0;
        // if relayer exclusivity is not intended for this bridge, set exclusivityEndTime to static zero
        // otherwise, set exclusivity to expire at the current block ts offset by quoteExclusivitySeconds
        if (paramsV2.quoteRelayer != address(0)) {
            exclusivityEndTime = int256(block.timestamp) + paramsV2.quoteExclusivitySeconds;
        }
        _validateBridgeParams(params, paramsV2, exclusivityEndTime);

        // transfer tokens to bridge contract
        /// @dev use returned originAmount in request in case of transfer fees
        uint256 originAmount = _takeBridgedUserAsset(params.originToken, params.originAmount);

        // track amount of origin token owed to protocol
        uint256 originFeeAmount;
        if (protocolFeeRate > 0) {
            originFeeAmount = (originAmount * protocolFeeRate) / FEE_BPS;
            originAmount -= originFeeAmount; // remove from amount used in request as not relevant for relayers
        }

        // set status to requested
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
                originFeeAmount: originFeeAmount,
                deadline: params.deadline,
                nonce: senderNonces[params.sender]++, // increment nonce on every bridge
                exclusivityRelayer: paramsV2.quoteRelayer,
                // We checked exclusivityEndTime to be in range [0 .. params.deadline] above, so can safely cast
                exclusivityEndTime: uint256(exclusivityEndTime),
                zapNative: paramsV2.zapNative,
                zapData: paramsV2.zapData
            })
        );
        bytes32 transactionId = keccak256(request);
        // Note: the tx status will be updated throughout the tx lifecycle, while destChainId is set once here
        bridgeTxDetails[transactionId].status = BridgeStatus.REQUESTED;
        bridgeTxDetails[transactionId].destChainId = params.dstChainId;

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
    }

    /// @inheritdoc IFastBridgeV2
    function relay(bytes calldata request, address relayer) public payable {
        request.validateV2();
        bytes32 transactionId = keccak256(request);
        _validateRelayParams(request, transactionId, relayer);
        // mark bridge transaction as relayed
        bridgeRelayDetails[transactionId] =
            BridgeRelay({blockNumber: uint48(block.number), blockTimestamp: uint48(block.timestamp), relayer: relayer});

        // transfer tokens to recipient on destination chain and trigger Zap if requested
        address to = request.destRecipient();
        address token = request.destToken();
        uint256 amount = request.destAmount();
        uint256 zapNative = request.zapNative();

        // Emit the event before any external calls
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
            // For the native gas token, additional zapNative is not allowed
            if (zapNative != 0) revert ZapNativeNotSupported();
            // Check that the correct msg.value was sent
            if (msg.value != amount) revert MsgValueIncorrect();
            // Don't do a native transfer yet: we will handle it alongside the Zap below
        } else {
            // For ERC20s, we check that the correct msg.value was sent
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

            // Note: if token has a fee on transfers, the recipient will have received less than `amount`.
            // This is a very niche edge case and should be handled by the recipient contract.
            _triggerZapWithChecks({recipient: to, token: token, amount: amount, zapData: zapData});
        } else if (msg.value != 0) {
            // Zap Data is missing, but msg.value was sent. This could happen in two different cases:
            // - Relay with the native gas token is happening.
            // - Relay with ERC20 is happening, with a `zapNative > 0` request.
            // In both cases, we need to transfer the full msg.value to the recipient.
            Address.sendValue(payable(to), msg.value);
        }
    }

    /// @inheritdoc IFastBridgeV2
    function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) public onlyRole(RELAYER_ROLE) {
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];

        // Can only prove a REQUESTED transaction
        if ($.status != BridgeStatus.REQUESTED) revert StatusIncorrect();
        // Update status to RELAYER_PROVED and store the proof details
        // Note: these are storage writes
        $.status = BridgeStatus.RELAYER_PROVED;
        $.proofBlockTimestamp = uint56(block.timestamp);
        $.proofRelayer = relayer;

        emit BridgeProofProvided(transactionId, relayer, destTxHash);
    }

    /// @inheritdoc IFastBridge
    function claim(bytes calldata request, address to) public {
        request.validateV2();
        bytes32 transactionId = keccak256(request);
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        // Aggregate the read operations from the same storage slot
        address proofRelayer = $.proofRelayer;
        BridgeStatus status = $.status;
        uint56 proofBlockTimestamp = $.proofBlockTimestamp;

        // Can only claim a RELAYER_PROVED transaction after the dispute period
        if (status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (_timeSince(proofBlockTimestamp) <= DISPUTE_PERIOD) {
            revert DisputePeriodNotPassed();
        }

        if (to == address(0)) {
            // Anyone could claim the funds to the proven relayer on their behalf
            to = proofRelayer;
        } else if (proofRelayer != msg.sender) {
            // Only the proven relayer could specify an address to claim the funds to
            revert SenderIncorrect();
        }

        // Update status to RELAYER_CLAIMED and transfer the origin collateral to the specified claim address
        // Note: this is a storage write
        $.status = BridgeStatus.RELAYER_CLAIMED;

        address token = request.originToken();
        uint256 amount = request.originAmount();
        // Update protocol fees if origin fee amount exists
        uint256 originFeeAmount = request.originFeeAmount();
        if (originFeeAmount > 0) protocolFees[token] += originFeeAmount;
        // Emit the event before any external calls
        emit BridgeDepositClaimed(transactionId, proofRelayer, to, token, amount);
        // Complete the relayer claim as the last transaction action
        if (token == NATIVE_GAS_TOKEN) {
            Address.sendValue(payable(to), amount);
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

    /// @inheritdoc IFastBridgeV2
    function cancel(bytes calldata request) public {
        request.validateV2();
        bytes32 transactionId = keccak256(request);
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        // Can only cancel a REQUESTED transaction after its deadline expires
        if ($.status != BridgeStatus.REQUESTED) revert StatusIncorrect();
        uint256 deadline = request.deadline();
        // Permissionless cancel is only allowed after `cancelDelay` on top of the deadline
        if (!hasRole(CANCELER_ROLE, msg.sender)) deadline += cancelDelay;
        if (block.timestamp <= deadline) revert DeadlineNotExceeded();
        // Update status to REFUNDED and return the full amount (collateral + protocol fees) to the original sender.
        // The protocol fees are only updated when the transaction is claimed, so we don't need to update them here.
        // Note: this is a storage write
        $.status = BridgeStatus.REFUNDED;

        address to = request.originSender();
        address token = request.originToken();
        uint256 amount = request.originAmount() + request.originFeeAmount();
        // Emit the event before any external calls
        emit BridgeDepositRefunded(transactionId, to, token, amount);
        // Complete the user cancel as the last transaction action
        if (token == NATIVE_GAS_TOKEN) {
            Address.sendValue(payable(to), amount);
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

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
        // has this transactionId been relayed?
        return bridgeRelayDetails[transactionId].relayer != address(0);
    }

    /// @notice Takes the bridged asset from the user into FastBridgeV2 custody. It will be later
    /// claimed by the relayer who completed the relay on destination chain, or transferred back to the user
    /// via the cancel function should no one complete the relay.
    function _takeBridgedUserAsset(address token, uint256 amount) internal returns (uint256 amountTaken) {
        if (token == NATIVE_GAS_TOKEN) {
            // For the native gas token, we just need to check that the supplied msg.value is correct.
            // Supplied `msg.value` is already in FastBridgeV2 custody.
            if (amount != msg.value) revert MsgValueIncorrect();
            amountTaken = msg.value;
        } else {
            // For ERC20s, token is explicitly transferred from the user to FastBridgeV2.
            // We don't allow non-zero `msg.value` to avoid extra funds from being stuck in FastBridgeV2.
            if (msg.value != 0) revert MsgValueIncorrect();
            // Throw an explicit error if the provided token address is not a contract
            if (token.code.length == 0) revert TokenNotContract();
            amountTaken = IERC20(token).balanceOf(address(this));
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
            // Use the balance difference as the amount taken in case of fee on transfer tokens.
            amountTaken = IERC20(token).balanceOf(address(this)) - amountTaken;
        }
    }

    /// @notice Calls the Recipient's hook function with the specified zapData and performs
    /// all the necessary checks for the returned value.
    function _triggerZapWithChecks(address recipient, address token, uint256 amount, bytes calldata zapData) internal {
        // This will bubble any revert messages from the hook function
        bytes memory returnData = Address.functionCallWithValue({
            target: recipient,
            data: abi.encodeCall(IZapRecipient.zap, (token, amount, zapData)),
            // Note: see `relay()` for reasoning behind passing msg.value
            value: msg.value
        });
        // Explicit revert if no return data at all
        if (returnData.length == 0) revert RecipientNoReturnValue();
        // Check that exactly a single return value was returned
        if (returnData.length != 32) revert RecipientIncorrectReturnValue();
        // Return value should be abi-encoded hook function selector
        if (bytes32(returnData) != bytes32(IZapRecipient.zap.selector)) {
            revert RecipientIncorrectReturnValue();
        }
    }

    /// @notice Calculates time since proof submitted
    /// @dev proof.timestamp stores casted uint56(block.timestamp) block timestamps for gas optimization
    ///      _timeSince(proof) can accomodate rollover case when block.timestamp > type(uint56).max but
    ///      proof.timestamp < type(uint56).max via unchecked statement
    /// @param proofBlockTimestamp The bridge proof block timestamp
    /// @return delta Time delta since proof submitted
    function _timeSince(uint56 proofBlockTimestamp) internal view returns (uint256 delta) {
        unchecked {
            delta = uint56(block.timestamp) - proofBlockTimestamp;
        }
    }

    /// @notice Performs all the necessary checks for a bridge to happen.
    /// @dev There's no good way to refactor this function to reduce cyclomatic complexity due to
    /// the number of checks that need to be performed, so we skip the code-complexity rule here.
    // solhint-disable-next-line code-complexity
    function _validateBridgeParams(
        BridgeParams memory params,
        BridgeParamsV2 memory paramsV2,
        int256 exclusivityEndTime
    )
        internal
        view
    {
        // Check V1 (legacy) params
        if (params.dstChainId == block.chainid) revert ChainIncorrect();
        if (params.originAmount == 0 || params.destAmount == 0) revert AmountIncorrect();
        if (params.sender == address(0) || params.to == address(0)) revert ZeroAddress();
        if (params.originToken == address(0) || params.destToken == address(0)) revert ZeroAddress();
        if (params.deadline < block.timestamp + MIN_DEADLINE_PERIOD) revert DeadlineTooShort();
        // Check V2 params
        if (paramsV2.zapData.length > MAX_ZAP_DATA_LENGTH) revert ZapDataLengthAboveMax();
        if (paramsV2.zapNative != 0 && params.destToken == NATIVE_GAS_TOKEN) {
            revert ZapNativeNotSupported();
        }
        // exclusivityEndTime must be in range [0 .. params.deadline]
        if (exclusivityEndTime < 0 || exclusivityEndTime > int256(params.deadline)) {
            revert ExclusivityParamsIncorrect();
        }
    }

    /// @notice Performs all the necessary checks for a relay to happen.
    function _validateRelayParams(bytes calldata request, bytes32 transactionId, address relayer) internal view {
        if (relayer == address(0)) revert ZeroAddress();
        // Check if the transaction has already been relayed
        if (bridgeRelays(transactionId)) revert TransactionRelayed();
        if (request.destChainId() != block.chainid) revert ChainIncorrect();
        // Check the deadline for relay to happen
        if (block.timestamp > request.deadline()) revert DeadlineExceeded();
        // Check the exclusivity period, if it is still ongoing
        address exclRelayer = request.exclusivityRelayer();
        if (exclRelayer != address(0) && exclRelayer != relayer && block.timestamp <= request.exclusivityEndTime()) {
            revert ExclusivityPeriodNotPassed();
        }
    }
}
