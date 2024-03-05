// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB, InterchainEntry, InterchainBatch} from "../../contracts/interfaces/IInterchainDB.sol";

contract InterchainDBMock is IInterchainDB {
    function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce) {}

    function requestVerification(uint256 dstChainId, uint256 dbNonce, address[] memory srcModules) external payable {}

    function writeEntryWithVerification(
        uint256 dstChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint256 writerNonce)
    {}

    function verifyEntry(InterchainEntry memory entry) external {}

    function verifyRemoteBatch(InterchainBatch memory batch) external {}

    function getInterchainFee(uint256 dstChainId, address[] memory srcModules) external view returns (uint256) {}

    function getEntry(uint256 dbNonce) external view returns (InterchainEntry memory) {}

    function getDBNonce() external view returns (uint256) {}

    function getNextEntryIndex() external view returns (uint256 dbNonce, uint64 entryIndex) {}

    function readEntry(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        returns (uint256 moduleVerifiedAt)
    {}
}
