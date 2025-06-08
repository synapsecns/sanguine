// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridge} from "../../src/interfaces/ISynapseBridge.sol";

// solhint-disable no-empty-blocks
contract SynapseBridgeMock is ISynapseBridge {
    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testSynapseBridgeMock() external {}

    function mint(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) external override {}
    function withdraw(address to, address token, uint256 amount, uint256 fee, bytes32 kappa) external override {}
}
