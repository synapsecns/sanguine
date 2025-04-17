// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseBridgeAdapterErrors {
    error SBA__AmountZero();
    error SBA__BridgeAlreadySet();
    error SBA__BridgeNotSet();
    error SBA__GasLimitBelowMinimum();
    error SBA__RecipientZeroAddress();
    error SBA__SymbolAlreadyAdded(bytes31 symbol);
    error SBA__SymbolUnknown(bytes31 symbol);
    error SBA__TokenAlreadyAdded(address token);
    error SBA__TokenUnknown(address token);
}
