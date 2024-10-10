// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

import {UniversalTokenLib} from "./libs/UniversalToken.sol";

import {Admin} from "./Admin.sol";
import {IFastBridge} from "./interfaces/IFastBridge.sol";
import {IFastBridgeV2} from "./interfaces/IFastBridgeV2.sol";
import {IFastBridgeV2Errors} from "./interfaces/IFastBridgeV2Errors.sol";
import {IFastBridgeRecipient} from "./interfaces/IFastBridgeRecipient.sol";

/// @notice FastBridgeV2 is a contract for bridging tokens across chains.
contract FastBridgeV2 is Admin, IFastBridgeV2, IFastBridgeV2Errors {
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
    function relay(bytes memory request) external payable {
        relay({request: request, relayer: msg.sender});
    }

    /// @inheritdoc IFastBridge
    function prove(bytes memory request, bytes32 destTxHash) external {
        prove({transactionId: keccak256(request), destTxHash: destTxHash, relayer: msg.sender});
    }

    /// @inheritdoc IFastBridgeV2
    function claim(bytes memory request) external {
        claim({request: request, to: address(0)});
    }

    /// @inheritdoc IFastBridge
    function dispute(bytes32 transactionId) external onlyRole(GUARD_ROLE) {
        if (bridgeTxDetails[transactionId].status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (_timeSince(bridgeTxDetails[transactionId].proofBlockTimestamp) > DISPUTE_PERIOD) {
            revert DisputePeriodPassed();
        }

        // @dev relayer gets slashed effectively if dest relay has gone thru
        bridgeTxDetails[transactionId].status = BridgeStatus.REQUESTED;
        bridgeTxDetails[transactionId].proofRelayer = address(0);
        bridgeTxDetails[transactionId].proofBlockTimestamp = 0;
        bridgeTxDetails[transactionId].proofBlockNumber = 0;

        emit BridgeProofDisputed(transactionId, msg.sender);
    }

    /// @inheritdoc IFastBridge
    function refund(bytes memory request) external {
        bytes32 transactionId = keccak256(request);

        BridgeTransactionV2 memory transaction = getBridgeTransactionV2(request);

        if (bridgeTxDetails[transactionId].status != BridgeStatus.REQUESTED) revert StatusIncorrect();

        if (hasRole(REFUNDER_ROLE, msg.sender)) {
            // Refunder can refund if deadline has passed
            if (block.timestamp <= transaction.deadline) revert DeadlineNotExceeded();
        } else {
            // Permissionless refund is allowed after REFUND_DELAY
            if (block.timestamp <= transaction.deadline + REFUND_DELAY) revert DeadlineNotExceeded();
        }

        // if all checks passed, set to REFUNDED status
        bridgeTxDetails[transactionId].status = BridgeStatus.REFUNDED;

        // transfer origin collateral back to original sender
        address to = transaction.originSender;
        address token = transaction.originToken;
        uint256 amount = transaction.originAmount + transaction.originFeeAmount;
        token.universalTransfer(to, amount);

        emit BridgeDepositRefunded(transactionId, to, token, amount);
    }

    /// @inheritdoc IFastBridge
    function canClaim(bytes32 transactionId, address relayer) external view returns (bool) {
        if (bridgeTxDetails[transactionId].status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (bridgeTxDetails[transactionId].proofRelayer != relayer) revert SenderIncorrect();
        return _timeSince(bridgeTxDetails[transactionId].proofBlockTimestamp) > DISPUTE_PERIOD;
    }

    /// @inheritdoc IFastBridge
    /// @dev This method is added to achieve backwards compatibility with decoding requests into V1 structs:
    /// - `callValue` is partially reported as a zero/non-zero flag
    /// - `callParams` is ignored
    /// In order to process all kinds of requests use getBridgeTransactionV2 instead.
    function getBridgeTransaction(bytes memory request) external view returns (BridgeTransaction memory) {
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
    function bridge(BridgeParams memory params, BridgeParamsV2 memory paramsV2) public payable {
        int256 exclusivityEndTime = int256(block.timestamp) + paramsV2.quoteExclusivitySeconds;
        _validateBridgeParams(params, paramsV2, exclusivityEndTime);

        // transfer tokens to bridge contract
        /// @dev use returned originAmount in request in case of transfer fees
        uint256 originAmount = _takeBridgedUserAsset(params.originToken, params.originAmount);

        // track amount of origin token owed to protocol
        uint256 originFeeAmount;
        if (protocolFeeRate > 0) originFeeAmount = (originAmount * protocolFeeRate) / FEE_BPS;
        originAmount -= originFeeAmount; // remove from amount used in request as not relevant for relayers

        // set status to requested
        bytes memory request = abi.encode(
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
                callValue: paramsV2.callValue,
                deadline: params.deadline,
                nonce: senderNonces[params.sender]++, // increment nonce on every bridge
                exclusivityRelayer: paramsV2.quoteRelayer,
                // We checked exclusivityEndTime to be in range (0 .. params.deadline] above, so can safely cast
                exclusivityEndTime: uint256(exclusivityEndTime),
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
    function relay(bytes memory request, address relayer) public payable {
        bytes32 transactionId = keccak256(request);
        BridgeTransactionV2 memory transaction = getBridgeTransactionV2(request);
        _validateRelayParams(transaction, transactionId, relayer);
        // mark bridge transaction as relayed
        bridgeRelayDetails[transactionId] =
            BridgeRelay({blockNumber: uint48(block.number), blockTimestamp: uint48(block.timestamp), relayer: relayer});

        // transfer tokens to recipient on destination chain and do an arbitrary call if requested
        address to = transaction.destRecipient;
        address token = transaction.destToken;
        uint256 amount = transaction.destAmount;
        uint256 callValue = transaction.callValue;

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

        if (transaction.callParams.length != 0) {
            // Arbitrary call requested, perform it while supplying full msg.value to the recipient
            // Note: if token has a fee on transfers, the recipient will have received less than `amount`.
            // This is a very niche edge case and should be handled by the recipient contract.
            _checkedCallRecipient({recipient: to, token: token, amount: amount, callParams: transaction.callParams});
        } else if (msg.value != 0) {
            // No arbitrary call requested, but msg.value was sent. This is either a relay with ETH,
            // or a non-zero callValue request with an ERC20. In both cases, transfer the ETH to the recipient.
            Address.sendValue(payable(to), msg.value);
        }

        emit BridgeRelayed({
            transactionId: transactionId,
            relayer: relayer,
            to: to,
            originChainId: transaction.originChainId,
            originToken: transaction.originToken,
            destToken: token,
            originAmount: transaction.originAmount,
            destAmount: amount,
            chainGasAmount: callValue
        });
    }

    /// @inheritdoc IFastBridgeV2
    function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) public onlyRole(RELAYER_ROLE) {
        // update bridge tx status given proof provided
        if (bridgeTxDetails[transactionId].status != BridgeStatus.REQUESTED) revert StatusIncorrect();
        bridgeTxDetails[transactionId].status = BridgeStatus.RELAYER_PROVED;
        bridgeTxDetails[transactionId].proofBlockTimestamp = uint40(block.timestamp);
        bridgeTxDetails[transactionId].proofBlockNumber = uint48(block.number);
        bridgeTxDetails[transactionId].proofRelayer = relayer;

        emit BridgeProofProvided(transactionId, relayer, destTxHash);
    }

    /// @inheritdoc IFastBridge
    function claim(bytes memory request, address to) public {
        bytes32 transactionId = keccak256(request);
        BridgeTransactionV2 memory transaction = getBridgeTransactionV2(request);

        // update bridge tx status if able to claim origin collateral
        if (bridgeTxDetails[transactionId].status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();

        // if "to" is zero addr, permissionlessly send funds to proven relayer
        if (to == address(0)) {
            to = bridgeTxDetails[transactionId].proofRelayer;
        } else if (bridgeTxDetails[transactionId].proofRelayer != msg.sender) {
            revert SenderIncorrect();
        }

        if (_timeSince(bridgeTxDetails[transactionId].proofBlockTimestamp) <= DISPUTE_PERIOD) {
            revert DisputePeriodNotPassed();
        }

        bridgeTxDetails[transactionId].status = BridgeStatus.RELAYER_CLAIMED;

        // update protocol fees if origin fee amount exists
        if (transaction.originFeeAmount > 0) protocolFees[transaction.originToken] += transaction.originFeeAmount;

        // transfer origin collateral less fee to specified address
        address token = transaction.originToken;
        uint256 amount = transaction.originAmount;
        token.universalTransfer(to, amount);

        emit BridgeDepositClaimed(transactionId, bridgeTxDetails[transactionId].proofRelayer, to, token, amount);
    }

    function bridgeStatuses(bytes32 transactionId) public view returns (BridgeStatus status) {
        return bridgeTxDetails[transactionId].status;
    }

    function bridgeProofs(bytes32 transactionId) public view returns (uint96 timestamp, address relayer) {
        timestamp = bridgeTxDetails[transactionId].proofBlockTimestamp;
        relayer = bridgeTxDetails[transactionId].proofRelayer;
    }

    /// @inheritdoc IFastBridgeV2
    function bridgeRelays(bytes32 transactionId) public view returns (bool) {
        // has this transactionId been relayed?
        return bridgeRelayDetails[transactionId].relayer != address(0);
    }

    /// @inheritdoc IFastBridgeV2
    function getBridgeTransactionV2(bytes memory request) public pure returns (BridgeTransactionV2 memory) {
        return abi.decode(request, (BridgeTransactionV2));
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
        bytes memory callParams
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
        if (exclusivityEndTime <= 0 || exclusivityEndTime > int256(params.deadline)) {
            revert ExclusivityParamsIncorrect();
        }
    }

    /// @notice Performs all the necessary checks for a relay to happen.
    function _validateRelayParams(
        BridgeTransactionV2 memory transaction,
        bytes32 transactionId,
        address relayer
    )
        internal
        view
    {
        if (relayer == address(0)) revert ZeroAddress();
        // Check if the transaction has already been relayed
        if (bridgeRelays(transactionId)) revert TransactionRelayed();
        if (transaction.destChainId != block.chainid) revert ChainIncorrect();
        // Check the deadline for relay to happen
        if (block.timestamp > transaction.deadline) revert DeadlineExceeded();
        // Check the exclusivity period, if it is still ongoing
        // forgefmt: disable-next-item
        if (
            transaction.exclusivityRelayer != address(0) &&
            transaction.exclusivityRelayer != relayer &&
            block.timestamp <= transaction.exclusivityEndTime
        ) {
            revert ExclusivityPeriodNotPassed();
        }
    }
}
