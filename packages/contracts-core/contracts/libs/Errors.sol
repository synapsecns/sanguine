// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ INVALID CALLER ═══════════════════════════════

error CallerNotAgentManager();
error CallerNotDestination();
error CallerNotInbox();
error CallerNotSummit();

// ══════════════════════════════ INCORRECT DATA ═══════════════════════════════

error IncorrectAttestation();
error IncorrectAgentDomain();
error IncorrectAgentIndex();
error IncorrectAgentProof();
error IncorrectAgentRoot();
error IncorrectDataHash();
error IncorrectDestinationDomain();
error IncorrectOriginDomain();
error IncorrectSnapshotProof();
error IncorrectSnapshotRoot();
error IncorrectState();
error IncorrectStatesAmount();
error IncorrectTipsProof();
error IncorrectVersionLength();

error IncorrectNonce();
error IncorrectSender();
error IncorrectRecipient();

error FlagOutOfRange();
error IndexOutOfRange();
error NonceOutOfRange();

error OutdatedNonce();

error UnformattedAttestation();
error UnformattedAttestationReport();
error UnformattedBaseMessage();
error UnformattedCallData();
error UnformattedCallDataPrefix();
error UnformattedMessage();
error UnformattedReceipt();
error UnformattedReceiptReport();
error UnformattedSignature();
error UnformattedSnapshot();
error UnformattedState();
error UnformattedStateReport();

// ═══════════════════════════════ MERKLE TREES ════════════════════════════════

error LeafNotProven();
error MerkleTreeFull();
error NotEnoughLeafs();
error TreeHeightTooLow();

// ═════════════════════════════ OPTIMISTIC PERIOD ═════════════════════════════

error BaseClientOptimisticPeriod();
error MessageOptimisticPeriod();
error SlashAgentOptimisticPeriod();
error WithdrawTipsOptimisticPeriod();
error ZeroProofMaturity();

// ═══════════════════════════════ AGENT MANAGER ═══════════════════════════════

error AgentNotGuard();
error AgentNotNotary();

error AgentCantBeAdded();
error AgentNotActive();
error AgentNotActiveNorUnstaking();
error AgentNotFraudulent();
error AgentNotUnstaking();
error AgentUnknown();

error AgentRootNotProposed();
error AgentRootTimeoutNotOver();

error NotStuck();

error DisputeAlreadyResolved();
error DisputeNotOpened();
error DisputeTimeoutNotOver();
error GuardInDispute();
error NotaryInDispute();

error MustBeSynapseDomain();
error SynapseDomainForbidden();

// ════════════════════════════════ DESTINATION ════════════════════════════════

error AlreadyExecuted();
error AlreadyFailed();
error DuplicatedSnapshotRoot();
error IncorrectMagicValue();
error GasLimitTooLow();
error GasSuppliedTooLow();

// ══════════════════════════════════ ORIGIN ═══════════════════════════════════

error ContentLengthTooBig();
error EthTransferFailed();
error InsufficientEthBalance();

// ════════════════════════════════ GAS ORACLE ═════════════════════════════════

error LocalGasDataNotSet();
error RemoteGasDataNotSet();

// ═══════════════════════════════════ TIPS ════════════════════════════════════

error SummitTipTooHigh();
error TipsClaimMoreThanEarned();
error TipsClaimZero();
error TipsOverflow();
error TipsValueTooLow();

// ════════════════════════════════ MEMORY VIEW ════════════════════════════════

error IndexedTooMuch();
error ViewOverrun();
error OccupiedMemory();
error UnallocatedMemory();
error PrecompileOutOfGas();

// ═════════════════════════════════ MULTICALL ═════════════════════════════════

error MulticallFailed();
