// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB, InterchainEntry, InterchainBatch} from "../../contracts/interfaces/IInterchainDB.sol";

// solhint-disable no-empty-blocks
contract InterchainDBMock is IInterchainDB {
    function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce, uint64 entryIndex) {}

    function requestBatchVerification(
        uint256 dstChainId,
        uint256 dbNonce,
        address[] memory srcModules
    )
        external
        payable
    {}

    function writeEntryWithVerification(
        uint256 dstChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint256 writerNonce, uint64 entryIndex)
    {}

    function verifyRemoteBatch(bytes calldata versionedBatch) external {}

    function getInterchainFee(uint256 dstChainId, address[] memory srcModules) external view returns (uint256) {}

    function getBatchLeafs(uint256 dbNonce) external view returns (bytes32[] memory) {}

    function getBatchLeafsPaginated(
        uint256 dbNonce,
        uint64 start,
        uint64 end
    )
        external
        view
        returns (bytes32[] memory)
    {}

    function getBatchSize(uint256 dbNonce) external view returns (uint64) {}

    function getBatch(uint256 dbNonce) external view returns (InterchainBatch memory) {}

    function getEntryValue(uint256 dbNonce, uint64 entryIndex) external view returns (bytes32) {}

    function getEntryProof(uint256 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof) {}

    function getDBNonce() external view returns (uint256) {}

    function getNextEntryIndex() external view returns (uint256 dbNonce, uint64 entryIndex) {}

    function checkVerification(
        address dstModule,
        InterchainEntry memory entry,
        bytes32[] memory proof
    )
        external
        view
        returns (uint256 moduleVerifiedAt)
    {}
}
