// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SafeCall} from "../../../contracts/libs/SafeCall.sol";

contract SafeCallHarness {
    function safeCall(address recipient, uint256 gasLimit, uint256 msgValue, bytes memory payload)
        external
        returns (bool)
    {
        bool success = SafeCall.safeCall(recipient, gasLimit, msgValue, payload);
        return success;
    }
}
