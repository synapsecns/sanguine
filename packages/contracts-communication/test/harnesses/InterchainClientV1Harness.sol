// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {InterchainClientV1} from "../../contracts/InterchainClientV1.sol";
import {InterchainEntry} from "../../contracts/libs/InterchainEntry.sol";

contract InterchainClientV1Harness is InterchainClientV1 {
    constructor(address interchainDB, address owner_) InterchainClientV1(interchainDB, owner_) {}

    /**
     * @dev Harness for testing _getFinalizedResponsesCount function
     */
    function getFinalizedResponsesCountHarness(
        address[] memory approvedModules,
        InterchainEntry memory icEntry,
        uint256 optimisticPeriod
    )
        public
        view
        returns (uint256)
    {
        return _getFinalizedResponsesCount(approvedModules, icEntry, optimisticPeriod);
    }
}
