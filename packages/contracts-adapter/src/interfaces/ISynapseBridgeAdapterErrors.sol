// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseBridgeAdapterErrors {
    error SBA__BridgeAlreadySet();
    error SBA__BridgeNotSet();
    error SBA__GasLimitBelowMinimum();
    error SBA__RemoteTokenAlreadyAssigned(uint32 eid, address localAddr);
    error SBA__RemoteTokenAlreadyUsed(uint32 eid, address remoteAddr);
    error SBA__RemoteTokenNotAssigned(uint32 eid, address localAddr);
    error SBA__RemoteTokenUnknown(uint32 eid, address remoteAddr);
    error SBA__TokenAlreadyAdded(address token);
    error SBA__TokenTypeUnknown();
    error SBA__TokenUnknown(address token);
    error SBA__ZeroAddress();
    error SBA__ZeroAmount();
}
