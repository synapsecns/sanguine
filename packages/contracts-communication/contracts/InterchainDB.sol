// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IInterchainDB} from "./interfaces/IInterchainDB.sol";
import {IInterchainDBEvents} from "./interfaces/IInterchainDBEvents.sol";
import {IInterchainModule} from "./interfaces/IInterchainModule.sol";

import {TypeCasts} from "./libs/TypeCasts.sol";

contract InterchainDB is IInterchainDB, IInterchainDBEvents {
    using TypeCasts for address;

    mapping(address writer => bytes32[] dataHashes) internal _entries;

    // ═══════════════════════════════════════════════ WRITER-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce) {}

    /// @inheritdoc IInterchainDB
    function requestVerification(
        uint256 destChainId,
        address writer,
        uint256 writerNonce,
        address[] memory srcModules
    )
        external
        payable
    {}

    /// @inheritdoc IInterchainDB
    function writeEntryWithVerification(
        uint256 destChainId,
        bytes32 dataHash,
        address[] calldata srcModules
    )
        external
        payable
        returns (uint256 writerNonce)
    {}

    // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function verifyEntry(InterchainEntry memory entry) external {}

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function readEntry(
        InterchainEntry memory entry,
        address[] memory dstModules
    )
        external
        view
        returns (uint256[] memory moduleVerifiedAt)
    {}

    /// @inheritdoc IInterchainDB
    function getInterchainFee(uint256 destChainId, address[] calldata srcModules) public view returns (uint256 fee) {
        uint256 len = srcModules.length;
        if (len == 0) {
            revert InterchainDB__NoModulesSpecified();
        }
        for (uint256 i = 0; i < len; ++i) {
            fee += IInterchainModule(srcModules[i]).getModuleFee(destChainId);
        }
    }

    /// @inheritdoc IInterchainDB
    function getEntry(address writer, uint256 writerNonce) public view returns (InterchainEntry memory) {
        if (getWriterNonce(writer) <= writerNonce) {
            revert InterchainDB__EntryDoesNotExist(writer, writerNonce);
        }
        return InterchainEntry({
            srcChainId: block.chainid,
            srcWriter: writer.addressToBytes32(),
            writerNonce: writerNonce,
            dataHash: _entries[writer][writerNonce]
        });
    }

    /// @inheritdoc IInterchainDB
    function getWriterNonce(address writer) public view returns (uint256) {
        return _entries[writer].length;
    }
}
