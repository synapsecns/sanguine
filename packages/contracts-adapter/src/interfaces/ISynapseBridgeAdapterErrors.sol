// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface ISynapseBridgeAdapterErrors {
    error SBA__BridgeAlreadySet();
    error SBA__BridgeNotSet();
    error SBA__GasLimitBelowMinimum();
    error SBA__LocalPairAlreadyExists(uint32 eid, address remoteAddr);
    error SBA__LocalPairNotFound(uint32 eid, address remoteAddr);
    error SBA__RemotePairAlreadySet(uint32 eid, address localAddr);
    error SBA__RemotePairNotSet(uint32 eid, address localAddr);
    error SBA__TokenAlreadyAdded(address token);
    error SBA__TokenTypeUnknown();
    error SBA__TokenUnknown(address token);
    error SBA__ZeroAddress();
    error SBA__ZeroAmount();
}
