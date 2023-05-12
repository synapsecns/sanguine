// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {CallerNotDestination} from "../../../contracts/libs/Errors.sol";
import {AgentManager} from "../../../contracts/manager/AgentManager.sol";

abstract contract AgentManagerHarness is AgentManager {
    /// @notice Exposes _slashAgent for testing.
    function slashAgentExposed(uint32 domain, address agent, address prover) external {
        _slashAgent(domain, agent, prover);
    }

    // ══════════════════════════════════════════ REMOTE CALLED FUNCTIONS ══════════════════════════════════════════════

    /// @notice Function that should NOT be callable by a Manager Message.
    /// Note: first two arguments could be of other types than (msgOrigin, proofMaturity)
    function sensitiveMockFunc(address, uint8, bytes32 data) external view returns (bytes32) {
        if (msg.sender != destination) revert CallerNotDestination();
        if (data == bytes32(0)) data = "GM";
        // Data is not zero, so this will differ from the required magic value
        return this.sensitiveMockFunc.selector ^ data;
    }

    /// @notice Function that should NOT be callable by a Manager Message.
    /// Note: first two arguments could be of other types than (msgOrigin, proofMaturity)
    function sensitiveMockFuncVoid(uint16, bytes4, bytes32) external view {
        if (msg.sender != destination) revert CallerNotDestination();
        // Doesn't return anything
    }

    /// @notice Function that should NOT be callable by a Manager Message.
    /// Note: first two arguments could be of other types than (msgOrigin, proofMaturity)
    function sensitiveMockFuncOver32Bytes(uint16, bytes4, bytes32 data) external view returns (bytes4, bytes32) {
        if (msg.sender != destination) revert CallerNotDestination();
        // Returning over 32 bytes should also fail the magic value test
        return (this.sensitiveMockFuncOver32Bytes.selector, data);
    }

    function remoteMockFunc(uint32, uint256, bytes32) external view returns (bytes4) {
        if (msg.sender != destination) revert CallerNotDestination();
        return this.remoteMockFunc.selector;
    }
}
