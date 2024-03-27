// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainAppV1Events {
    event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod);
    event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp);
    event ExecutionServiceSet(address executionService);
    event TrustedModuleAdded(address module);
    event TrustedModuleRemoved(address module);
}
