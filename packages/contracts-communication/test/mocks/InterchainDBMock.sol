// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB, InterchainBatch, InterchainEntry} from "../../contracts/interfaces/IInterchainDB.sol";

// solhint-disable no-empty-blocks
contract InterchainDBMock is IInterchainDB {
    uint16 public constant DB_VERSION = 1;

    function writeEntry(bytes32 dataHash) external returns (uint64 writerNonce, uint64 entryIndex) {}

    function requestBatchVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] memory srcModules
    )
        external
        payable
    {}

    function writeEntryWithVerification(
        uint64 dstChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint64 writerNonce, uint64 entryIndex)
    {}

    function verifyRemoteBatch(bytes calldata versionedBatch) external {}

    function getInterchainFee(uint64 dstChainId, address[] memory srcModules) external view returns (uint256) {}

    function getBatchLeafs(uint64 dbNonce) external view returns (bytes32[] memory) {}

    function getBatchLeafsPaginated(
        uint64 dbNonce,
        uint64 start,
        uint64 end
    )
        external
        view
        returns (bytes32[] memory)
    {}

    function getBatchSize(uint64 dbNonce) external view returns (uint64) {}

    function getBatch(uint64 dbNonce) external view returns (InterchainBatch memory) {}

    function getEntryValue(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32) {}

    function getEntryProof(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {}

    function getDBNonce() external view returns (uint64) {}

    function getNextEntryIndex() external view returns (uint64 dbNonce, uint64 entryIndex) {}

    function checkBatchVerification(
        address dstModule,
        InterchainBatch memory batch
    )
        external
        view
        returns (uint256 moduleVerifiedAt)
    {}

    function getBatchRoot(InterchainEntry memory entry, bytes32[] memory proof) external pure returns (bytes32) {}
}
