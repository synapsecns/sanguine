// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IMulticallTarget} from "../interfaces/IMulticallTarget.sol";

// solhint-disable avoid-low-level-calls
/// @notice Abstract contract template that supports batched calls while preserving msg.sender.
/// Only calls with msg.value of 0 can be batched.
abstract contract MulticallTarget is IMulticallTarget {
    error MulticallTarget__UndeterminedRevert();

    /// @inheritdoc IMulticallTarget
    function multicallNoResults(bytes[] calldata data, bool ignoreReverts) external {
        for (uint256 i = 0; i < data.length; ++i) {
            // We perform a delegate call to ourself to preserve the msg.sender. This is identical to `msg.sender`
            // calling the functions directly one by one, therefore doesn't add any security risks.
            // Note: msg.value is also preserved when doing a delegate call, but this function is not payable,
            // so it's always 0 and not a security risk.
            (bool success, bytes memory result) = address(this).delegatecall(data[i]);
            if (!success && !ignoreReverts) {
                _bubbleRevert(result);
            }
        }
    }

    /// @inheritdoc IMulticallTarget
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
                _bubbleRevert(results[i].returnData);
            }
        }
    }

    /// @dev Bubbles the revert message from the underlying call.
    /// Note: preserves the same custom error or revert string, if one was used.
    /// Source: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v5.0.2/contracts/utils/Address.sol#L143-L158
    function _bubbleRevert(bytes memory returnData) internal pure {
        // Look for revert reason and bubble it up if present
        if (returnData.length > 0) {
            // The easiest way to bubble the revert reason is using memory via assembly
            // solhint-disable-next-line no-inline-assembly
            assembly {
                let returndata_size := mload(returnData)
                revert(add(32, returnData), returndata_size)
            }
        } else {
            revert MulticallTarget__UndeterminedRevert();
        }
    }
}
