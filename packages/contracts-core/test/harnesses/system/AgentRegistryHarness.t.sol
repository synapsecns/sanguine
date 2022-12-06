// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import "../events/AgentRegistryHarnessEvents.sol";
import { AgentRegistryExtended } from "./AgentRegistryExtended.t.sol";
import { ByteString } from "../../../contracts/libs/ByteString.sol";

// solhint-disable no-empty-blocks
contract AgentRegistryHarness is AgentRegistryHarnessEvents, AgentRegistryExtended {
    using ByteString for bytes;

    bool internal ignoreMode;
    uint32 internal ignoredDomain;
    address internal ignoredAddress;

    function toggleIgnoreMode(bool _newValue) external {
        ignoreMode = _newValue;
    }

    function setIgnoredAgent(uint32 _domain, address _account) external {
        ignoredDomain = _domain;
        ignoredAddress = _account;
    }

    function checkAgentAuth(
        uint32 _domain,
        bytes32 _digest,
        bytes memory _signature
    ) external view returns (address) {
        return _checkAgentAuth(_domain, _digest, _signature.castToSignature());
    }

    function currentEpoch() external view returns (uint256) {
        return _currentEpoch();
    }

    /// @notice Exposes haveActiveGuard modifier for testing
    function onlyActiveGuard() external view haveActiveGuard {}

    /// @notice Exposes haveActiveNotary modifier for testing
    function onlyActiveNotary(uint32 _domain) external view haveActiveNotary(_domain) {}

    /**
     * @notice Hook that is called right after a Agent is added for specified domain.
     */
    function _afterAgentAdded(uint32 _domain, address _account) internal virtual override {
        require(_isActiveAgent(_domain, _account), "!afterAgentAdded");
        emit AfterAgentAdded(_domain, _account);
    }

    /**
     * @notice Hook that is called right after a Agent is removed from specified domain.
     */
    function _afterAgentRemoved(uint32 _domain, address _account) internal virtual override {
        require(!_isActiveAgent(_domain, _account), "!afterAgentRemoved");
        emit AfterAgentRemoved(_domain, _account);
    }

    function _isIgnoredAgent(uint32 _domain, address _account)
        internal
        view
        override
        returns (bool)
    {
        return ignoreMode && ignoredDomain == _domain && ignoredAddress == _account;
    }
}
