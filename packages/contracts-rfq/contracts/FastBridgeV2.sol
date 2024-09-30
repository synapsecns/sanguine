// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {UniversalTokenLib} from "./libs/UniversalToken.sol";

import {Admin} from "./Admin.sol";
import {IFastBridge} from "./interfaces/IFastBridge.sol";
import {IFastBridgeV2} from "./interfaces/IFastBridgeV2.sol";
import {IFastBridgeV2Errors} from "./interfaces/IFastBridgeV2Errors.sol";

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

    /// @notice Status of the bridge tx on origin chain
    mapping(bytes32 => BridgeTxDetails) public bridgeTxDetails;
    /// @notice Relay details on destination chain
    mapping(bytes32 => BridgeRelay) public bridgeRelayDetails;

    /// @dev to prevent replays
    uint256 public nonce;
    // @dev the block the contract was deployed at
    uint256 public immutable deployBlock;

    function bridgeStatuses(bytes32 transactionId) public view returns (BridgeStatus status) {
        return bridgeTxDetails[transactionId].status;
    }

    function bridgeProofs(bytes32 transactionId) public view returns (uint96 timestamp, address relayer) {
        timestamp = bridgeTxDetails[transactionId].proofBlockTimestamp;
        relayer = bridgeTxDetails[transactionId].proofRelayer;
    }

    constructor(address _owner) Admin(_owner) {
        deployBlock = block.number;
    }

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
            if (amount != msg.value) revert MsgValueIncorrect();
            // Transfer value to recipient if not this address
            if (recipient != address(this)) token.universalTransfer(recipient, amount);
            // We will forward msg.value in the external call later, if recipient is not this contract
            amountPulled = msg.value;
        }
    }

    /// @inheritdoc IFastBridge
    function getBridgeTransaction(bytes memory request) external pure returns (BridgeTransaction memory) {
        // Note: when passing V2 request, this will decode the V1 fields correctly since the new fields were
        // added as the last fields of the struct and hence the ABI decoder will simply ignore the extra data.
        return abi.decode(request, (BridgeTransaction));
    }

    /// @inheritdoc IFastBridgeV2
    function getBridgeTransactionV2(bytes memory request) public pure returns (BridgeTransactionV2 memory) {
        return abi.decode(request, (BridgeTransactionV2));
    }

    /// @inheritdoc IFastBridge
    function bridge(BridgeParams memory params) external payable {
        BridgeParamsV2 memory defaultParamsV2 =
            BridgeParamsV2({quoteRelayer: address(0), quoteExclusivitySeconds: 0, quoteId: bytes("")});
        bridge(params, defaultParamsV2);
    }

    /// @inheritdoc IFastBridgeV2
    // TODO: reduce cyclomatic complexity alongside arbitrary call
    // solhint-disable-next-line code-complexity
    function bridge(BridgeParams memory params, BridgeParamsV2 memory paramsV2) public payable {
        // check bridge params
        if (params.dstChainId == block.chainid) revert ChainIncorrect();
        if (params.originAmount == 0 || params.destAmount == 0) revert AmountIncorrect();
        if (params.sender == address(0) || params.to == address(0)) revert ZeroAddress();
        if (params.originToken == address(0) || params.destToken == address(0)) revert ZeroAddress();
        if (params.deadline < block.timestamp + MIN_DEADLINE_PERIOD) revert DeadlineTooShort();
        // quoteRelayer and quoteExclusivitySeconds must be both set or both not set
        if ((paramsV2.quoteRelayer == address(0)) != (paramsV2.quoteExclusivitySeconds == 0)) {
            revert ExclusivityParamsIncorrect();
        }

        // transfer tokens to bridge contract
        // @dev use returned originAmount in request in case of transfer fees
        uint256 originAmount = _pullToken(address(this), params.originToken, params.originAmount);

        // track amount of origin token owed to protocol
        uint256 originFeeAmount;
        if (protocolFeeRate > 0) originFeeAmount = (originAmount * protocolFeeRate) / FEE_BPS;
        originAmount -= originFeeAmount; // remove from amount used in request as not relevant for relayers

        // Calculate exclusivity end time only if exclusivity params are set
        uint256 exclusivityEndTime =
            paramsV2.quoteExclusivitySeconds == 0 ? 0 : block.timestamp + paramsV2.quoteExclusivitySeconds;
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
                sendChainGas: params.sendChainGas,
                deadline: params.deadline,
                nonce: nonce++, // increment nonce on every bridge
                exclusivityRelayer: paramsV2.quoteRelayer,
                exclusivityEndTime: exclusivityEndTime
            })
        );
        bytes32 transactionId = keccak256(request);
        bridgeTxDetails[transactionId].status = BridgeStatus.REQUESTED;

        emit BridgeRequested(
            transactionId,
            params.sender,
            request,
            params.dstChainId,
            params.originToken,
            params.destToken,
            originAmount,
            params.destAmount,
            params.sendChainGas
        );
        emit BridgeQuoteDetails(transactionId, paramsV2.quoteId);
    }

    /// @inheritdoc IFastBridge
    function relay(bytes memory request) external payable {
        relay(request, msg.sender);
    }

    /// @inheritdoc IFastBridgeV2
    // TODO: reduce cyclomatic complexity alongside arbitrary call
    // solhint-disable-next-line code-complexity
    function relay(bytes memory request, address relayer) public payable {
        if (relayer == address(0)) revert ZeroAddress();
        // Check if the transaction has already been relayed
        bytes32 transactionId = keccak256(request);
        if (bridgeRelays(transactionId)) revert TransactionRelayed();
        // Decode the transaction and check that it could be relayed on this chain
        BridgeTransactionV2 memory transaction = getBridgeTransactionV2(request);
        if (transaction.destChainId != uint32(block.chainid)) revert ChainIncorrect();
        // Check the deadline for relay to happen
        if (block.timestamp > transaction.deadline) revert DeadlineExceeded();
        // Check the exclusivity period, if it is still ongoing
        if (block.timestamp <= transaction.exclusivityEndTime && relayer != transaction.exclusivityRelayer) {
            revert ExclusivityPeriodNotPassed();
        }
        // mark bridge transaction as relayed
        bridgeRelayDetails[transactionId] =
            BridgeRelay({blockNumber: uint48(block.number), blockTimestamp: uint48(block.timestamp), relayer: relayer});

        // transfer tokens to recipient on destination chain and gas rebate if requested
        address to = transaction.destRecipient;
        address token = transaction.destToken;
        uint256 amount = transaction.destAmount;

        uint256 rebate = chainGasAmount;
        if (!transaction.sendChainGas) {
            // forward erc20
            rebate = 0;
            _pullToken(to, token, amount);
        } else if (token == UniversalTokenLib.ETH_ADDRESS) {
            // lump in gas rebate into amount in native gas token
            _pullToken(to, token, amount + rebate);
        } else {
            // forward erc20 then forward gas rebate in native gas token
            _pullToken(to, token, amount);
            _pullToken(to, UniversalTokenLib.ETH_ADDRESS, rebate);
        }

        emit BridgeRelayed(
            transactionId,
            relayer,
            to,
            transaction.originChainId,
            transaction.originToken,
            transaction.destToken,
            transaction.originAmount,
            transaction.destAmount,
            rebate
        );
    }

    /// @inheritdoc IFastBridgeV2
    function bridgeRelays(bytes32 transactionId) public view returns (bool) {
        // has this transactionId been relayed?
        return bridgeRelayDetails[transactionId].relayer != address(0);
    }

    /// @inheritdoc IFastBridge
    function prove(bytes memory request, bytes32 destTxHash) external {
        bytes32 transactionId = keccak256(request);
        prove(transactionId, destTxHash, msg.sender);
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

    /// @inheritdoc IFastBridge
    function canClaim(bytes32 transactionId, address relayer) external view returns (bool) {
        if (bridgeTxDetails[transactionId].status != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (bridgeTxDetails[transactionId].proofRelayer != relayer) revert SenderIncorrect();
        return _timeSince(bridgeTxDetails[transactionId].proofBlockTimestamp) > DISPUTE_PERIOD;
    }

    /// @inheritdoc IFastBridgeV2
    function claim(bytes memory request) external {
        claim(request, address(0));
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
}
