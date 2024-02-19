// SPDX-License-Identifier: MIT

pragma solidity 0.8.20;

import {InterchainClientV1} from "../../contracts/InterchainClientV1.sol";
import { InterchainEntry } from "../../contracts/libs/InterchainEntry.sol";
contract InterchainClientV1Harness is InterchainClientV1 {
    constructor() InterchainClientV1() {}

    /**
     * @dev Harness for testing _generateTransactionId function
     */
    function generateTransactionIdHarness(
        bytes32 srcSender,
        uint256 srcChainId,
        bytes32 dstReceiver,
        uint256 dstChainId,
        bytes memory message,
        uint64 nonce,
        bytes memory options
    )
        public
        pure
        returns (bytes32)
    {
        return _generateTransactionId(srcSender, srcChainId, dstReceiver, dstChainId, message, nonce, options);
    }

    /**
     * @dev Harness for testing _getApprovedResponses function
     */
    function getApprovedResponsesHarness(
        address[] memory approvedModules,
        InterchainEntry memory icEntry
    )
        public
        view
        returns (uint256[] memory)
    {
        return _getApprovedResponses(approvedModules, icEntry);
    }

    /**
     * @dev Harness for testing _getAppConfig function
     */
    function getAppConfigHarness(address receiverApp)
        public
        view
        returns (uint256 requiredResponses, uint256 optimisticTimePeriod, address[] memory approvedDstModules)
    {
        return _getAppConfig(receiverApp);
    }

    /**
     * @dev Harness for testing _getFinalizedResponsesCount function
     */
    function getFinalizedResponsesCountHarness(
        uint256[] memory approvedResponses,
        uint256 optimisticTimePeriod
    )
        public
        view
        returns (uint256)
    {
        return _getFinalizedResponsesCount(approvedResponses, optimisticTimePeriod);
    }

}
