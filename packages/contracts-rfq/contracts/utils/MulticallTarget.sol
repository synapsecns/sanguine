// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IMulticallTarget} from "../interfaces/IMulticallTarget.sol";

abstract contract MulticallTarget is IMulticallTarget {
    error MulticallTarget__UndeterminedRevert();

    function multicallNoResults(bytes[] calldata data, bool ignoreReverts) external {
        // TODO: Implement
    }

    function multicallWithResults(
        bytes[] calldata data,
        bool ignoreReverts
    )
        external
        returns (Result[] memory results)
    {
        // TODO: Implement
    }
}
