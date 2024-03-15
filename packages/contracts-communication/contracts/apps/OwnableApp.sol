// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainAppV1, AppConfigV1} from "./InterchainAppV1.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

abstract contract OwnableApp is InterchainAppV1, Ownable {
    constructor(address owner_) Ownable(owner_) {}

    /// @notice Allows the owner to link the remote app for the given chain ID.
    /// - This address will be used as the receiver for the messages sent from this chain.
    /// - This address will be the only trusted sender for the messages sent to this chain.
    function linkRemoteApp(uint256 chainId, bytes32 remoteApp) external onlyOwner {
        _linkRemoteApp(chainId, remoteApp);
    }

    /// @notice This wrapper for `linkRemoteApp` to accept EVM address as a parameter.
    function linkRemoteAppEVM(uint256 chainId, address remoteApp) external onlyOwner {
        _linkRemoteAppEVM(chainId, remoteApp);
    }

    /// @notice Allows the owner to add the module to the trusted modules set.
    /// - This set of modules will be used to verify message sent from this chain.
    /// - This set of modules will be used to verify message sent to this chain.
    function addTrustedModule(address module) external onlyOwner {
        _addTrustedModule(module);
    }

    /// @notice Allows the owner to remove the module from the trusted modules set.
    function removeTrustedModule(address module) external onlyOwner {
        _removeTrustedModule(module);
    }

    /// @notice Allows the owner to set the app config for the current app. App config includes:
    /// - requiredResponses: the number of module responses required for accepting the message
    /// - optimisticPeriod: the minimum time after which the module responses are considered final
    function setAppConfigV1(AppConfigV1 memory appConfig) external onlyOwner {
        _setAppConfigV1(appConfig);
    }

    /// @notice Allows the owner to set the address of the Execution Service.
    /// This address will be used to request execution of the messages sent from this chain,
    /// by supplying the Service's execution fee.
    function setExecutionService(address executionService) external onlyOwner {
        _setExecutionService(executionService);
    }

    /// @notice Allows the owner to set the address of the InterchainClient contract.
    function setInterchainClient(address interchain_) external onlyOwner {
        _setInterchainClient(interchain_);
    }
}
