// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IMulticallTarget} from "../interfaces/IMulticallTarget.sol";

// solhint-disable avoid-low-level-calls
abstract contract MulticallTarget is IMulticallTarget {
    error MulticallTarget__UndeterminedRevert();

    function multicallNoResults(bytes[] calldata data, bool ignoreReverts) external {
        for (uint256 i = 0; i < data.length; ++i) {
            // We perform a delegate call to ourself to preserve the msg.sender. This is identical to `msg.sender`
            // calling the functions directly one by one, therefore doesn't add any security risks.
            // Note: msg.value is also preserved when doing a delegate call, but this function is not payable,
            // so it's always 0 and not a security risk.
            (bool success, bytes memory result) = address(this).delegatecall(data[i]);
            if (!success && !ignoreReverts) {
                revert(string(result));
            }
        }
    }

    function multicallWithResults(
        bytes[] calldata data,
        bool ignoreReverts
    )
        external
        returns (Result[] memory results)
    {
        results = new Result[](data.length);
        for (uint256 i = 0; i < data.length; ++i) {
            // We perform a delegate call to ourself to preserve the msg.sender. This is identical to `msg.sender`
            // calling the functions directly one by one, therefore doesn't add any security risks.
            // Note: msg.value is also preserved when doing a delegate call, but this function is not payable,
            // so it's always 0 and not a security risk.
            (results[i].success, results[i].returnData) = address(this).delegatecall(data[i]);
            if (!results[i].success && !ignoreReverts) {
                revert(string(results[i].returnData));
            }
        }
    }
}
