// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

abstract contract InterchainAppV1Events {
    /// @notice Emitted when the app configuration V1 is set.
    /// The V1 version of app config requests at least N confirmations that are at least T seconds old
    /// from trusted modules to execute a transaction.
    /// @param requiredResponses    The number of module responses required for a transaction to be executed.
    /// @param optimisticPeriod     The time period after which a module response is considered final.
    event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod);

    /// @notice Emitted when a remote instance of the app is linked.
    /// This instance is the only one that can send messages to this app from the remote chain.
    /// @param chainId              The remote chain ID.
    /// @param remoteApp            The address of the remote app on that chain.
    event AppLinked(uint64 chainId, bytes32 remoteApp);

    /// @notice Emitted when the execution service is set.
    /// This service will be used for requesting the execution of transactions on the remote chain.
    /// @param executionService     The address of the execution service.
    event ExecutionServiceSet(address executionService);

    /// @notice Emitted when a trusted module is added.
    /// The trusted modules will be used to verify the messages coming from the remote chains,
    /// as well as request the verification of the sent messages on the remote chains.
    /// @param module               The address of the trusted module that was added.
    event TrustedModuleAdded(address module);

    /// @notice Emitted when a trusted module is removed.
    /// @param module               The address of the trusted module that was removed.
    event TrustedModuleRemoved(address module);
}
