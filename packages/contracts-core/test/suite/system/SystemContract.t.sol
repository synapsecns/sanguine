// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseTest} from "../../utils/SynapseTest.t.sol";
import {RawStateIndex} from "../../utils/libs/SynapseStructs.t.sol";

abstract contract SystemContractTest is SynapseTest {
    modifier boundIndex(RawStateIndex memory rsi) {
        rsi.boundStateIndex();
        _;
    }

    modifier onlySupportedDomain() virtual {
        require(localDomain() == DOMAIN_LOCAL || localDomain() == DOMAIN_SYNAPSE, "Unsupported local domain");
        _;
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Returns local domain for the tested system contract
    function localDomain() public view virtual returns (uint32);

    /// @notice Returns address of the tested system contract
    function systemContract() public view virtual returns (address);

    /// @notice Returns address of Agent Manager on the tested domain
    function localAgentManager() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(lightManager) : address(bondingManager);
    }

    /// @notice Returns address of Destination on the tested domain
    function localDestination() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(destination) : address(summit);
    }

    /// @notice Returns address of Origin on the tested domain
    function localOrigin() public view virtual onlySupportedDomain returns (address) {
        return localDomain() == DOMAIN_LOCAL ? address(origin) : address(originSynapse);
    }

    /// @notice Checks if contract is a local SystemContract
    function isLocalSystemContract(address addr) public view returns (bool) {
        return addr == localAgentManager() || addr == localDestination() || addr == localOrigin();
    }

    /// @notice Checks if contract is a local SystemRegistry
    function isLocalSystemRegistry(address addr) public view returns (bool) {
        return addr == localDestination() || addr == localOrigin();
    }
}
