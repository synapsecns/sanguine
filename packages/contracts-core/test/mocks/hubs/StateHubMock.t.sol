// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IStateHub } from "../../../contracts/interfaces/IStateHub.sol";

// solhint-disable no-empty-blocks
contract StateHubMock is IStateHub {
    /// @notice Prevents this contract from being included in the coverage report
    function testStateHubMock() external {}

    function isValidState(bytes memory statePayload) external view returns (bool isValid) {}

    function statesAmount() external view returns (uint256) {}

    function suggestLatestState() external view returns (bytes memory statePayload) {}

    function suggestState(uint32 nonce) external view returns (bytes memory statePayload) {}
}
