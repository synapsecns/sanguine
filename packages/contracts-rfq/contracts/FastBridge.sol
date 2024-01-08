// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import "./libs/Errors.sol";
import {UniversalTokenLib} from "./libs/UniversalToken.sol";

import {Admin} from "./Admin.sol";
import {IFastBridge} from "./interfaces/IFastBridge.sol";

contract FastBridge is IFastBridge, Admin {
    using SafeERC20 for IERC20;
    using UniversalTokenLib for address;

    /// @notice Dispute period for relayed transactions
    uint256 public constant DISPUTE_PERIOD = 30 minutes;

    /// @notice Minimum deadline period to relay a requested bridge transaction
    uint256 public constant MIN_DEADLINE_PERIOD = 30 minutes;

    enum BridgeStatus {
        NULL, // doesn't exist yet
        REQUESTED,
        RELAYER_PROVED,
        RELAYER_CLAIMED,
        REFUNDED
    }

    /// @notice Status of the bridge tx on origin chain
    mapping(bytes32 => BridgeStatus) public bridgeStatuses;
    /// @notice Proof of relayed bridge tx on origin chain
    mapping(bytes32 => BridgeProof) public bridgeProofs;
    /// @notice Whether bridge has been relayed on destination chain
    mapping(bytes32 => bool) public bridgeRelays;

    /// @dev to prevent replays
    uint256 public nonce;
    // @dev the block the contract was deployed at
    uint256 public immutable deployBlock;

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
    function getBridgeTransaction(bytes memory request) public pure returns (BridgeTransaction memory) {
        return abi.decode(request, (BridgeTransaction));
    }

    /// @inheritdoc IFastBridge
    function bridge(BridgeParams memory params) external payable {
        // check bridge params
        if (params.dstChainId == block.chainid) revert ChainIncorrect();
        if (params.originAmount == 0 || params.destAmount == 0) revert AmountIncorrect();
        if (params.originToken == address(0) || params.destToken == address(0)) revert ZeroAddress();
        if (params.deadline < block.timestamp + MIN_DEADLINE_PERIOD) revert DeadlineTooShort();

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
        bridgeStatuses[transactionId] = BridgeStatus.REQUESTED;

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
    function relay(bytes memory request) external payable onlyRelayer {
        bytes32 transactionId = keccak256(request);
        BridgeTransaction memory transaction = getBridgeTransaction(request);
        if (transaction.destChainId != uint32(block.chainid)) revert ChainIncorrect();

        // mark bridge transaction as relayed
        if (bridgeRelays[transactionId]) revert TransactionRelayed();
        bridgeRelays[transactionId] = true;

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

        // check haven't exceeded deadline for relay to happen
        if (block.timestamp > transaction.deadline) revert DeadlineExceeded();

        // update bridge tx status given proof provided
        if (bridgeStatuses[transactionId] != BridgeStatus.REQUESTED) revert StatusIncorrect();
        bridgeStatuses[transactionId] = BridgeStatus.RELAYER_PROVED;
        bridgeProofs[transactionId] = BridgeProof({timestamp: uint96(block.timestamp), relayer: msg.sender}); // overflow ok

        emit BridgeProofProvided(transactionId, msg.sender, destTxHash);
    }

    /// @notice Calculates time since proof submitted
    /// @dev proof.timestamp stores casted uint96(block.timestamp) block timestamps for gas optimization
    ///      _timeSince(proof) can accomodate rollover case when block.timestamp > type(uint96).max but
    ///      proof.timestamp < type(uint96).max via unchecked statement
    /// @param proof The bridge proof
    /// @return delta Time delta since proof submitted
    function _timeSince(BridgeProof memory proof) internal view returns (uint256 delta) {
        unchecked {
            delta = uint96(block.timestamp) - proof.timestamp;
        }
    }

    /// @inheritdoc IFastBridge
    function canClaim(bytes32 transactionId, address relayer) external view returns (bool) {
        if (bridgeStatuses[transactionId] != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        BridgeProof memory proof = bridgeProofs[transactionId];
        if (proof.relayer != relayer) revert SenderIncorrect();
        return _timeSince(proof) > DISPUTE_PERIOD;
    }

    /// @inheritdoc IFastBridge
    function claim(bytes memory request, address to) external onlyRelayer {
        bytes32 transactionId = keccak256(request);
        BridgeTransaction memory transaction = getBridgeTransaction(request);

        // update bridge tx status if able to claim origin collateral
        if (bridgeStatuses[transactionId] != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();

        BridgeProof memory proof = bridgeProofs[transactionId];
        if (proof.relayer != msg.sender) revert SenderIncorrect();
        if (_timeSince(proof) <= DISPUTE_PERIOD) revert DisputePeriodNotPassed();

        bridgeStatuses[transactionId] = BridgeStatus.RELAYER_CLAIMED;

        // update protocol fees if origin fee amount exists
        if (transaction.originFeeAmount > 0) protocolFees[transaction.originToken] += transaction.originFeeAmount;

        // transfer origin collateral less fee to specified address
        address token = transaction.originToken;
        uint256 amount = transaction.originAmount;
        token.universalTransfer(to, amount);

        emit BridgeDepositClaimed(transactionId, msg.sender, to, token, amount);
    }

    /// @inheritdoc IFastBridge
    function dispute(bytes32 transactionId) external onlyGuard {
        if (bridgeStatuses[transactionId] != BridgeStatus.RELAYER_PROVED) revert StatusIncorrect();
        if (_timeSince(bridgeProofs[transactionId]) > DISPUTE_PERIOD) revert DisputePeriodPassed();

        // @dev relayer gets slashed effectively if dest relay has gone thru
        bridgeStatuses[transactionId] = BridgeStatus.REQUESTED;
        delete bridgeProofs[transactionId];

        emit BridgeProofDisputed(transactionId, msg.sender);
    }

    /// @inheritdoc IFastBridge
    function refund(bytes memory request, address to) external {
        bytes32 transactionId = keccak256(request);
        BridgeTransaction memory transaction = getBridgeTransaction(request);
        if (transaction.originSender != msg.sender) revert SenderIncorrect();
        if (block.timestamp <= transaction.deadline) revert DeadlineNotExceeded();

        // set status to refunded if still in requested state
        if (bridgeStatuses[transactionId] != BridgeStatus.REQUESTED) revert StatusIncorrect();
        bridgeStatuses[transactionId] = BridgeStatus.REFUNDED;

        // transfer origin collateral back to original sender's specified recipient
        address token = transaction.originToken;
        uint256 amount = transaction.originAmount + transaction.originFeeAmount;
        token.universalTransfer(to, amount);

        emit BridgeDepositRefunded(transactionId, to, token, amount);
    }
}
