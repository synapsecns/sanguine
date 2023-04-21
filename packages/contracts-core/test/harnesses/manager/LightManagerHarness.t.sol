// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {LightManager} from "../../../contracts/manager/LightManager.sol";

// solhint-disable no-empty-blocks
contract LightManagerHarness is LightManager {
    constructor(uint32 domain) LightManager(domain) {}

    /// @notice Function that should NOT be callable by a Manager Message.
    /// Note: first two arguments could be of other types than (msgOrigin, proofMaturity)
    function sensitiveMockFunc(address, uint8, bytes32 data) external returns (bytes32) {
        require(msg.sender == destination, "!destination");
        if (data == bytes32(0)) data = "GM";
        _setAgentRoot(data);
        // Data is not zero, so this will differ from the required magic value
        return this.sensitiveMockFunc.selector ^ data;
    }

    /// @notice Function that should NOT be callable by a Manager Message.
    /// Note: first two arguments could be of other types than (msgOrigin, proofMaturity)
    function sensitiveMockFuncVoid(uint16, bytes4, bytes32 data) external {
        require(msg.sender == destination, "!destination");
        _setAgentRoot(data);
        // Doesn't return anything
    }

    /// @notice Function that should NOT be callable by a Manager Message.
    /// Note: first two arguments could be of other types than (msgOrigin, proofMaturity)
    function sensitiveMockFuncOver32Bytes(uint16, bytes4, bytes32 data) external returns (bytes4, bytes32) {
        require(msg.sender == destination, "!destination");
        _setAgentRoot(data);
        // Returning over 32 bytes should also fail the magic value test
        return (this.sensitiveMockFuncOver32Bytes.selector, data);
    }

    function remoteMockFunc(uint32, uint256, bytes32) external view returns (bytes4) {
        require(msg.sender == destination, "!destination");
        return this.remoteMockFunc.selector;
    }
}
