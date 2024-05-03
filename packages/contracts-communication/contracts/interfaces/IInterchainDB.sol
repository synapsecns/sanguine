// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainBatch} from "../libs/InterchainBatch.sol";
import {InterchainEntry} from "../libs/InterchainEntry.sol";

interface IInterchainDB {
    error InterchainDB__BatchConflict(address module, bytes32 existingBatchRoot, InterchainBatch newBatch);
    error InterchainDB__BatchVersionMismatch(uint16 version, uint16 required);
    error InterchainDB__ChainIdNotRemote(uint64 chainId);
    error InterchainDB__EntryIndexOutOfRange(uint64 dbNonce, uint64 entryIndex, uint64 batchSize);
    error InterchainDB__EntryRangeInvalid(uint64 dbNonce, uint64 start, uint64 end);
    error InterchainDB__FeeAmountBelowMin(uint256 feeAmount, uint256 minRequired);
    error InterchainDB__ModulesNotProvided();

    function writeEntry(bytes32 dataHash) external returns (uint64 dbNonce, uint64 entryIndex);
    function requestBatchVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] memory srcModules
    )
        external
        payable;
    function writeEntryWithVerification(
        uint64 dstChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint64 dbNonce, uint64 entryIndex);

    function verifyRemoteBatch(bytes memory versionedBatch) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function getInterchainFee(uint64 dstChainId, address[] memory srcModules) external view returns (uint256);

    function getBatchLeafs(uint64 dbNonce) external view returns (bytes32[] memory);
    function getBatchLeafsPaginated(
        uint64 dbNonce,
        uint64 start,
        uint64 end
    )
        external
        view
        returns (bytes32[] memory);
    function getBatchSize(uint64 dbNonce) external view returns (uint64);
    function getBatch(uint64 dbNonce) external view returns (InterchainBatch memory);
    function getVersionedBatch(uint64 dbNonce) external view returns (bytes memory);

    function getEntryValue(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32);
    function getEntryProof(uint64 dbNonce, uint64 entryIndex) external view returns (bytes32[] memory proof);

    function getDBNonce() external view returns (uint64);
    function getNextEntryIndex() external view returns (uint64 dbNonce, uint64 entryIndex);

    function checkBatchVerification(
        address dstModule,
        InterchainBatch memory batch
    )
        external
        view
        returns (uint256 moduleVerifiedAt);

    function getBatchRoot(InterchainEntry memory entry, bytes32[] memory proof) external pure returns (bytes32);

    // solhint-disable-next-line func-name-mixedcase
    function DB_VERSION() external pure returns (uint16);
}
