// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

import {BridgeTransactionV2Lib} from "./libs/BridgeTransactionV2.sol";
import {UniversalTokenLib} from "./libs/UniversalToken.sol";

import {Admin} from "./Admin.sol";
import {IFastBridge} from "./interfaces/IFastBridge.sol";
import {IFastBridgeV2} from "./interfaces/IFastBridgeV2.sol";
import {IFastBridgeV2Errors} from "./interfaces/IFastBridgeV2Errors.sol";
import {IFastBridgeRecipient} from "./interfaces/IFastBridgeRecipient.sol";

/// @notice FastBridgeV2 is a contract for bridging tokens across chains.
contract FastBridgeV2 is Admin, IFastBridgeV2, IFastBridgeV2Errors {
    using BridgeTransactionV2Lib for bytes;
    using SafeERC20 for IERC20;
    using UniversalTokenLib for address;

    /// @notice Dispute period for relayed transactions
    uint256 public constant DISPUTE_PERIOD = 30 minutes;

    /// @notice Delay for a transaction after which it could be permisionlessly refunded
    uint256 public constant REFUND_DELAY = 7 days;

    /// @notice Minimum deadline period to relay a requested bridge transaction
    uint256 public constant MIN_DEADLINE_PERIOD = 30 minutes;

    /// @notice Maximum length of accepted callParams
    uint256 public constant MAX_CALL_PARAMS_LENGTH = 2 ** 16 - 1;

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

    constructor(address _owner) Admin(_owner) {
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
                callValue: 0,
                callParams: bytes("")
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
        uint40 proofBlockTimestamp = $.proofBlockTimestamp;
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
        $.proofBlockNumber = 0;

        emit BridgeProofDisputed(transactionId, disputedRelayer);
    }

    /// @inheritdoc IFastBridge
    function refund(bytes calldata request) external {
        request.validateV2();
        bytes32 transactionId = keccak256(request);
        BridgeTxDetails storage $ = bridgeTxDetails[transactionId];
        // Can only refund a REQUESTED transaction after its deadline expires
        if ($.status != BridgeStatus.REQUESTED) revert StatusIncorrect();
        uint256 deadline = request.deadline();
        // Permissionless refund is only allowed after REFUND_DELAY on top of the deadline
        if (!hasRole(REFUNDER_ROLE, msg.sender)) deadline += REFUND_DELAY;
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
        // Complete the user refund as the last transaction action
        if (token == UniversalTokenLib.ETH_ADDRESS) {
            Address.sendValue(payable(to), amount);
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
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
    /// - `callValue` is partially reported as a zero/non-zero flag
    /// - `callParams` is ignored
    /// In order to process all kinds of requests use getBridgeTransactionV2 instead.
    function getBridgeTransaction(bytes calldata request) external view returns (BridgeTransaction memory) {
        // Try decoding into V2 struct first. This will revert if V1 struct is passed
        try this.getBridgeTransactionV2(request) returns (BridgeTransactionV2 memory txV2) {
            // Note: we entirely ignore the callParams field, as it was not present in V1
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
                sendChainGas: txV2.callValue != 0,
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
        int256 exclusivityEndTime = paramsV2.quoteRelayer != address(0)
            // prettier-ignore
            ? int256(block.timestamp) + paramsV2.quoteExclusivitySeconds
            : int256(0);
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
                // We checked exclusivityEndTime to be in range (0 .. params.deadline] above, so can safely cast
                exclusivityEndTime: uint256(exclusivityEndTime),
                callValue: paramsV2.callValue,
                callParams: paramsV2.callParams
            })
        );
        bytes32 transactionId = keccak256(request);
        bridgeTxDetails[transactionId].status = BridgeStatus.REQUESTED;

        emit BridgeRequested({
            transactionId: transactionId,
            sender: params.sender,
            request: request,
            destChainId: params.dstChainId,
            originToken: params.originToken,
            destToken: params.destToken,
            originAmount: originAmount,
            destAmount: params.destAmount,
            sendChainGas: paramsV2.callValue != 0
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

        // transfer tokens to recipient on destination chain and do an arbitrary call if requested
        address to = request.destRecipient();
        address token = request.destToken();
        uint256 amount = request.destAmount();
        uint256 callValue = request.callValue();

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
            chainGasAmount: callValue
        });

        // All state changes have been done at this point, can proceed to the external calls.
        // This follows the checks-effects-interactions pattern to mitigate potential reentrancy attacks.
        if (token == UniversalTokenLib.ETH_ADDRESS) {
            // For ETH non-zero callValue is not allowed
            if (callValue != 0) revert NativeTokenCallValueNotSupported();
            // Check that the correct msg.value was sent
            if (msg.value != amount) revert MsgValueIncorrect();
        } else {
            // For ERC20s, we check that the correct msg.value was sent
            if (msg.value != callValue) revert MsgValueIncorrect();
            // We need to transfer the tokens from the Relayer to the recipient first before performing an
            // optional post-transfer arbitrary call.
            IERC20(token).safeTransferFrom(msg.sender, to, amount);
        }

        bytes calldata callParams = request.callParams();
        if (callParams.length != 0) {
            // Arbitrary call requested, perform it while supplying full msg.value to the recipient
            // Note: if token has a fee on transfers, the recipient will have received less than `amount`.
            // This is a very niche edge case and should be handled by the recipient contract.
            _checkedCallRecipient({recipient: to, token: token, amount: amount, callParams: callParams});
        } else if (msg.value != 0) {
            // No arbitrary call requested, but msg.value was sent. This is either a relay with ETH,
            // or a non-zero callValue request with an ERC20. In both cases, transfer the ETH to the recipient.
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
        $.proofBlockTimestamp = uint40(block.timestamp);
        $.proofBlockNumber = uint48(block.number);
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
        uint40 proofBlockTimestamp = $.proofBlockTimestamp;

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
        if (token == UniversalTokenLib.ETH_ADDRESS) {
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
    /// claimed by the relayer who completed the relay on destination chain, or refunded back to the user,
    /// should no one complete the relay.
    function _takeBridgedUserAsset(address token, uint256 amount) internal returns (uint256 amountTaken) {
        if (token == UniversalTokenLib.ETH_ADDRESS) {
            // For ETH we just need to check that the supplied msg.value is correct.
            // Supplied `msg.value` is already in FastBridgeV2 custody.
            if (amount != msg.value) revert MsgValueIncorrect();
            amountTaken = msg.value;
        } else {
            // For ERC20s, token is explicitly transferred from the user to FastBridgeV2.
            // We don't allow non-zero `msg.value` to avoid extra funds from being stuck in FastBridgeV2.
            token.assertIsContract();
            if (msg.value != 0) revert MsgValueIncorrect();
            amountTaken = IERC20(token).balanceOf(address(this));
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
            // Use the balance difference as the amount taken in case of fee on transfer tokens.
            amountTaken = IERC20(token).balanceOf(address(this)) - amountTaken;
        }
    }

    /// @notice Calls the Recipient's hook function with the specified callParams and performs
    /// all the necessary checks for the returned value.
    function _checkedCallRecipient(
        address recipient,
        address token,
        uint256 amount,
        bytes calldata callParams
    )
        internal
    {
        bytes memory hookData =
            abi.encodeCall(IFastBridgeRecipient.fastBridgeTransferReceived, (token, amount, callParams));
        // This will bubble any revert messages from the hook function
        bytes memory returnData = Address.functionCallWithValue({target: recipient, data: hookData, value: msg.value});
        // Explicit revert if no return data at all
        if (returnData.length == 0) revert RecipientNoReturnValue();
        // Check that exactly a single return value was returned
        if (returnData.length != 32) revert RecipientIncorrectReturnValue();
        // Return value should be abi-encoded hook function selector
        if (bytes32(returnData) != bytes32(IFastBridgeRecipient.fastBridgeTransferReceived.selector)) {
            revert RecipientIncorrectReturnValue();
        }
    }

    /// @notice Calculates time since proof submitted
    /// @dev proof.timestamp stores casted uint40(block.timestamp) block timestamps for gas optimization
    ///      _timeSince(proof) can accomodate rollover case when block.timestamp > type(uint40).max but
    ///      proof.timestamp < type(uint40).max via unchecked statement
    /// @param proofBlockTimestamp The bridge proof block timestamp
    /// @return delta Time delta since proof submitted
    function _timeSince(uint40 proofBlockTimestamp) internal view returns (uint256 delta) {
        unchecked {
            delta = uint40(block.timestamp) - proofBlockTimestamp;
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
        if (paramsV2.callParams.length > MAX_CALL_PARAMS_LENGTH) revert CallParamsLengthAboveMax();
        if (paramsV2.callValue != 0 && params.destToken == UniversalTokenLib.ETH_ADDRESS) {
            revert NativeTokenCallValueNotSupported();
        }
        // exclusivityEndTime must be in range (0 .. params.deadline]
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
