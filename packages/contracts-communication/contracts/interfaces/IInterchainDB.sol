// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntry} from "../libs/InterchainEntry.sol";

interface IInterchainDB {
    error InterchainDB__ChainIdNotRemote(uint64 chainId);
    error InterchainDB__EntryConflict(address module, InterchainEntry newEntry);
    error InterchainDB__EntryVersionMismatch(uint16 version, uint16 required);
    error InterchainDB__FeeAmountBelowMin(uint256 feeAmount, uint256 minRequired);
    error InterchainDB__ModulesNotProvided();

    function writeEntry(bytes32 digest) external returns (uint64 dbNonce);

    function requestEntryVerification(
        uint64 dstChainId,
        uint64 dbNonce,
        address[] memory srcModules
    )
        external
        payable;

    function writeEntryRequestVerification(
        uint64 dstChainId,
        bytes32 digest,
        address[] memory srcModules
    )
        external
        payable
        returns (uint64 dbNonce);

    function verifyRemoteEntry(bytes memory encodedEntry) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function getInterchainFee(uint64 dstChainId, address[] memory srcModules) external view returns (uint256);

    function getEncodedEntry(uint64 dbNonce) external view returns (bytes memory);
    function getEntry(uint64 dbNonce) external view returns (InterchainEntry memory);
    function getEntryValue(uint64 dbNonce) external view returns (bytes32);

    function getDBNonce() external view returns (uint64);

    function checkEntryVerification(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        returns (uint256 moduleVerifiedAt);

    // solhint-disable-next-line func-name-mixedcase
    function DB_VERSION() external pure returns (uint16);
}
