// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12;
pragma experimental ABIEncoderV2;

import {ChainAwareness} from "../base/ChainAwareness.sol";
import {StringUtils} from "../libs/StringUtils.sol";
import {DataReader} from "../reader/DataReader.sol";

abstract contract ChainAwareReader is ChainAwareness, DataReader {
    using StringUtils for *;

    // ════════════════════════════════════════════ GENERIC DATA READS ═════════════════════════════════════════════════

    /// @notice Returns the saved deployment artifact JSON for a contract on the active chain.
    /// @dev Returns an empty string or reverts if the artifact is not found based on the flag.
    function readDeploymentArtifact(
        string memory contractName,
        bool revertIfNotFound
    )
        internal
        view
        returns (string memory artifact)
    {
        return readDeploymentArtifact(activeChain, contractName, revertIfNotFound);
    }

    /// @notice Returns the deployment configuration JSON for a contract on the active chain.
    /// @dev Returns an empty string or reverts if the artifact is not found based on the flag.
    function readDeployConfig(
        string memory contractName,
        bool revertIfNotFound
    )
        internal
        view
        returns (string memory deployConfig)
    {
        return readDeployConfig(activeChain, contractName, revertIfNotFound);
    }

    // ════════════════════════════════════════════ SPECIFIC DATA READS ════════════════════════════════════════════════

    /// @notice Returns the deployment address for a contract on the active chain.
    /// @dev Returns address(0) or reverts if the address is not found based on the flag.
    function getDeploymentAddress(
        string memory contractName,
        bool revertIfNotFound
    )
        internal
        view
        returns (address deploymentAddress)
    {
        return getDeploymentAddress(activeChain, contractName, revertIfNotFound);
    }

    /// @notice Checks if a contract is deployed on the active chain without reverting.
    function isDeployed(string memory contractName) internal view returns (bool) {
        return isDeployed(activeChain, contractName);
    }

    // ═══════════════════════════════════════════════════ UTILS ═══════════════════════════════════════════════════════

    /// @notice Asserts that a contract is NOT deployed on the active chain by checking its code size.
    function assertContractCodeEmpty(address contractAddr, string memory errMsg) internal view {
        require(getCodeSize(contractAddr) == 0, errMsg.concat(" at ", vm.toString(contractAddr), " on ", activeChain));
    }

    /// @notice Asserts that a contract is deployed on the active chain by checking its code size.
    function assertContractCodeExists(address contractAddr, string memory errMsg) internal view {
        require(getCodeSize(contractAddr) != 0, errMsg.concat(" at ", vm.toString(contractAddr), " on ", activeChain));
    }

    /// @notice Returns the code size for a given address on the active chain.
    function getCodeSize(address contractAddr) internal view returns (uint256 codeSize) {
        // address.code.length is only available in Solidity 0.8.0+
        // solhint-disable-next-line no-inline-assembly
        assembly {
            codeSize := extcodesize(contractAddr)
        }
    }
}
