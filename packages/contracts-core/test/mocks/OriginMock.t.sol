// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {InterfaceOrigin} from "../../contracts/interfaces/InterfaceOrigin.sol";
import {StateHubMock} from "./hubs/StateHubMock.t.sol";
import {SystemRegistryMock} from "./system/SystemRegistryMock.t.sol";

// solhint-disable no-empty-blocks
contract OriginMock is StateHubMock, SystemRegistryMock, InterfaceOrigin {
    /// @notice Prevents this contract from being included in the coverage report
    function testOriginMock() external {}

    function sendBaseMessage(
        uint32 destination,
        bytes32 recipient,
        uint32 optimisticPeriod,
        uint256 paddedTips,
        uint256 paddedRequest,
        bytes memory content
    ) external payable returns (uint32 messageNonce, bytes32 messageHash) {}

    function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes memory payload)
        external
        returns (uint32 messageNonce, bytes32 messageHash)
    {}

    function withdrawTips(address recipient, uint256 amount) external {}
}
