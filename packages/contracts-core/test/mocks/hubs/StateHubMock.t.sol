// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { IStateHub } from "../../../contracts/interfaces/IStateHub.sol";
import { ExcludeCoverage } from "../ExcludeCoverage.sol";

// solhint-disable no-empty-blocks
contract StateHubMock is ExcludeCoverage, IStateHub {
    function isValidState(bytes memory _statePayload) external view returns (bool isValid) {}

    function statesAmount() external view returns (uint256) {}

    function suggestLatestState() external view returns (bytes memory statePayload) {}

    function suggestState(uint32 _nonce) external view returns (bytes memory statePayload) {}
}
