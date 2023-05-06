// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ INVALID CALLER ═══════════════════════════════

error CallerNotAgentManager();

// ═══════════════════════════════ AGENT MANAGER ═══════════════════════════════

error AgentNotGuard();
error AgentNotNotary();

error AgentCantBeAdded();
error AgentDomainIncorrect();
error AgentNotActive();
error AgentNotActiveNorUnstaking();
error AgentNotFraudulent();
error AgentNotUnstaking();
error AgentUnknown();

error DisputeAlreadyResolved();
error GuardInDispute();
error NotaryInDispute();

error MustBeSynapseDomain();
error SynapseDomainForbidden();

error SnapshotRootMismatch();
error StateMismatch();

// ════════════════════════════════ MEMORY VIEW ════════════════════════════════

error IndexedTooMuch();
error ViewOverrun();
error OccupiedMemory();
error UnallocatedMemory();
error PrecompileOutOfGas();

// ═════════════════════════════════ MULTICALL ═════════════════════════════════

error MulticallFailed();
