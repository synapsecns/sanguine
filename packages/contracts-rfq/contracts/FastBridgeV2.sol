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

    /// @inheritdoc IFastBridge
    function relay(bytes memory request) external payable onlyRelayer {
        bytes32 transactionId = keccak256(request);
        BridgeTransaction memory transaction = getBridgeTransaction(request);
        if (transaction.destChainId != uint32(block.chainid)) revert FastBridge__ChainIncorrect();

        // check haven't exceeded deadline for relay to happen
        if (block.timestamp > transaction.deadline) revert FastBridge__DeadlineExceeded();

        // mark bridge transaction as relayed
        if (_destinationRelayer[transactionId] != address(0)) revert FastBridge__TransactionRelayed();
        _destinationRelayer[transactionId] = msg.sender;

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
            msg.sender,
            to,
            transaction.originChainId,
            transaction.originToken,
            transaction.destToken,
            transaction.originAmount,
            transaction.destAmount,
            rebate
        );
    }

    /// @inheritdoc IFastBridge
    function prove(bytes memory request, bytes32 destTxHash) external onlyRelayer {
        bytes32 transactionId = keccak256(request);
        BridgeTransaction memory transaction = getBridgeTransaction(request);

        // check haven't exceeded deadline for prove to happen
        if (block.timestamp > transaction.deadline + PROVE_PERIOD) revert FastBridge__DeadlineExceeded();
        _verifyAndUpdateStatus({
            transactionId: transactionId,
            expectedOldValue: BridgeStatusV2.REQUESTED,
            newValue: BridgeStatusV2.RELAYER_PROVED
        });
        bridgeProofs[transactionId] = BridgeProof({timestamp: uint96(block.timestamp), relayer: msg.sender}); // overflow ok

        emit BridgeProofProvided(transactionId, msg.sender, destTxHash);
    }

    /// @inheritdoc IFastBridge
    function claim(bytes memory request, address to) external onlyRelayer {
        bytes32 transactionId = keccak256(request);
        BridgeProof memory proof = bridgeProofs[transactionId];
        if (proof.relayer != msg.sender) revert FastBridge__SenderIncorrect();
        if (_timeSince(proof) <= DISPUTE_PERIOD) revert FastBridge__DisputePeriodNotPassed();
        _verifyAndUpdateStatus({
            transactionId: transactionId,
            expectedOldValue: BridgeStatusV2.RELAYER_PROVED,
            newValue: BridgeStatusV2.RELAYER_CLAIMED
        });
        // update protocol fees if origin fee amount exists
        BridgeTransaction memory transaction = getBridgeTransaction(request);
        if (transaction.originFeeAmount > 0) protocolFees[transaction.originToken] += transaction.originFeeAmount;

        // transfer origin collateral less fee to specified address
        address token = transaction.originToken;
        uint256 amount = transaction.originAmount;
        token.universalTransfer(to, amount);

        emit BridgeDepositClaimed(transactionId, msg.sender, to, token, amount);
    }

    // ════════════════════════════════════════════════ ONLY GUARD ═════════════════════════════════════════════════════

    /// @inheritdoc IFastBridge
    function dispute(bytes32 transactionId) external onlyGuard {
        // TODO: this should be a two-step process:
        // 1. dispute() marks transaction as disputed - this stops relayer from claiming
        // 2. resolve() resolves a disputed transaction - this uses the external source of truth
        if (_timeSince(bridgeProofs[transactionId]) > DISPUTE_PERIOD) revert FastBridge__DisputePeriodPassed();
        // @dev relayer gets slashed effectively if dest relay has gone thru
        _verifyAndUpdateStatus({
            transactionId: transactionId,
            expectedOldValue: BridgeStatusV2.RELAYER_PROVED,
            newValue: BridgeStatusV2.REQUESTED
        });
        delete bridgeProofs[transactionId];

        emit BridgeProofDisputed(transactionId, msg.sender);
    }

    /// @inheritdoc IFastBridgeV2
    function resolve(bytes32 transactionId, address destRelayer) external onlyGuard {
        // TODO: implement
    }

    // ════════════════════════════════════════════════ USER-FACING ════════════════════════════════════════════════════

    /// @inheritdoc IFastBridge
    function bridge(BridgeParams memory params) external payable {
        // check bridge params
        if (params.dstChainId == block.chainid) revert FastBridge__ChainIncorrect();
        if (params.originAmount == 0 || params.destAmount == 0) revert FastBridge__AmountIncorrect();
        if (params.originToken == address(0) || params.destToken == address(0)) revert FastBridge__ZeroAddress();
        if (params.deadline < block.timestamp + MIN_DEADLINE_PERIOD) revert FastBridge__DeadlineTooShort();

        // transfer tokens to bridge contract
        // @dev use returned originAmount in request in case of transfer fees
        uint256 originAmount = _pullToken(address(this), params.originToken, params.originAmount);

        // track amount of origin token owed to protocol
        uint256 originFeeAmount;
        if (protocolFeeRate > 0) originFeeAmount = (originAmount * protocolFeeRate) / FEE_BPS;
        originAmount -= originFeeAmount; // remove from amount used in request as not relevant for relayers

        // set status to requested
        bytes memory request = abi.encode(
            BridgeTransaction({
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
                nonce: nonce++ // increment nonce on every bridge
            })
        );
        bytes32 transactionId = keccak256(request);
        _verifyAndUpdateStatus({
            transactionId: transactionId,
            expectedOldValue: BridgeStatusV2.NULL,
            newValue: BridgeStatusV2.REQUESTED
        });
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
    }

    /// @inheritdoc IFastBridge
    function refund(bytes memory request) external {
        bytes32 transactionId = keccak256(request);
        BridgeTransaction memory transaction = getBridgeTransaction(request);

        // check exceeded deadline for prove to happen
        if (block.timestamp <= transaction.deadline + PROVE_PERIOD) revert FastBridge__DeadlineNotExceeded();
        _verifyAndUpdateStatus({
            transactionId: transactionId,
            expectedOldValue: BridgeStatusV2.REQUESTED,
            newValue: BridgeStatusV2.REFUNDED
        });
        // transfer origin collateral back to original sender
        address to = transaction.originSender;
        address token = transaction.originToken;
        uint256 amount = transaction.originAmount + transaction.originFeeAmount;
        token.universalTransfer(to, amount);

        emit BridgeDepositRefunded(transactionId, to, token, amount);
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IFastBridgeV2
    function bridgeRelays(bytes32 transactionId) external view returns (bool) {
        return _destinationRelayer[transactionId] != address(0);
    }

    /// @inheritdoc IFastBridgeV2
    function getDestinationRelayer(bytes32 transactionId) external view returns (address) {
        return _destinationRelayer[transactionId];
    }

    /// @inheritdoc IFastBridge
    function canClaim(bytes32 transactionId, address relayer) external view returns (bool) {
        if (bridgeStatuses[transactionId] != BridgeStatusV2.RELAYER_PROVED) revert FastBridge__StatusIncorrect();
        BridgeProof memory proof = bridgeProofs[transactionId];
        if (proof.relayer != relayer) revert FastBridge__SenderIncorrect();
        return _timeSince(proof) > DISPUTE_PERIOD;
    }

    /// @inheritdoc IFastBridge
    function getBridgeTransaction(bytes memory request) public pure returns (BridgeTransaction memory) {
        return abi.decode(request, (BridgeTransaction));
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Changes the status of a bridge transaction, but only if the current status matches the expected value.
    /// Note: this is the only function that can change the status of a bridge transaction.
    function _verifyAndUpdateStatus(
        bytes32 transactionId,
        BridgeStatusV2 expectedOldValue,
        BridgeStatusV2 newValue
    )
        internal
    {
        if (bridgeStatuses[transactionId] != expectedOldValue) revert FastBridge__StatusIncorrect();
        bridgeStatuses[transactionId] = newValue;
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
