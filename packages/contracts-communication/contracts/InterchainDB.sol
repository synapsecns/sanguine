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
    function writeEntry(bytes32 dataHash) external returns (uint256 writerNonce) {
        return _writeEntry(dataHash);
    }

    /// @inheritdoc IInterchainDB
    function requestVerification(
        uint256 destChainId,
        address writer,
        uint256 writerNonce,
        address[] calldata srcModules
    )
        external
        payable
    {
        InterchainEntry memory entry = getEntry(writer, writerNonce);
        _requestVerification(destChainId, entry, srcModules);
    }

    /// @inheritdoc IInterchainDB
    function writeEntryWithVerification(
        uint256 destChainId,
        bytes32 dataHash,
        address[] calldata srcModules
    )
        external
        payable
        returns (uint256 writerNonce)
    {
        writerNonce = _writeEntry(dataHash);
        InterchainEntry memory entry = _constructEntry(msg.sender, writerNonce, dataHash);
        _requestVerification(destChainId, entry, srcModules);
    }

    // ═══════════════════════════════════════════════ MODULE-FACING ═══════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function verifyEntry(InterchainEntry memory entry) external {}

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IInterchainDB
    function readEntry(
        InterchainEntry memory entry,
        address[] calldata dstModules
    )
        external
        view
        returns (uint256[] memory moduleVerifiedAt)
    {}

    /// @inheritdoc IInterchainDB
    function getInterchainFee(uint256 destChainId, address[] calldata srcModules) external view returns (uint256 fee) {
        (, fee) = _getModuleFees(destChainId, srcModules);
    }

    /// @inheritdoc IInterchainDB
    function getEntry(address writer, uint256 writerNonce) public view returns (InterchainEntry memory) {
        if (getWriterNonce(writer) <= writerNonce) {
            revert InterchainDB__EntryDoesNotExist(writer, writerNonce);
        }
        return _constructEntry(writer, writerNonce, _entries[writer][writerNonce]);
    }

    /// @inheritdoc IInterchainDB
    function getWriterNonce(address writer) public view returns (uint256) {
        return _entries[writer].length;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Write the entry to the database and emit the event.
    function _writeEntry(bytes32 dataHash) internal returns (uint256 writerNonce) {
        writerNonce = _entries[msg.sender].length;
        _entries[msg.sender].push(dataHash);
        emit InterchainEntryWritten(block.chainid, msg.sender.addressToBytes32(), writerNonce, dataHash);
    }

    /// @dev Request the verification of the entry by the modules, and emit the event.
    /// Note: the validity of the passed entry is enforced in the calling function.
    function _requestVerification(
        uint256 destChainId,
        InterchainEntry memory entry,
        address[] calldata srcModules
    )
        internal
    {
        (uint256[] memory fees, uint256 totalFee) = _getModuleFees(destChainId, srcModules);
        if (msg.value != totalFee) {
            revert InterchainDB__IncorrectFeeAmount(msg.value, totalFee);
        }
        uint256 len = srcModules.length;
        for (uint256 i = 0; i < len; ++i) {
            IInterchainModule(srcModules[i]).requestVerification{value: fees[i]}(destChainId, entry);
        }
        emit InterchainVerificationRequested(destChainId, entry.srcWriter, entry.writerNonce, srcModules);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Construct the entry struct from the given parameters
    function _constructEntry(
        address writer,
        uint256 writerNonce,
        bytes32 dataHash
    )
        internal
        view
        returns (InterchainEntry memory)
    {
        return InterchainEntry({
            srcChainId: block.chainid,
            srcWriter: writer.addressToBytes32(),
            writerNonce: writerNonce,
            dataHash: dataHash
        });
    }

    /// @dev Get the verification fees for the modules
    function _getModuleFees(
        uint256 destChainId,
        address[] calldata srcModules
    )
        internal
        view
        returns (uint256[] memory fees, uint256 totalFee)
    {
        uint256 len = srcModules.length;
        if (len == 0) {
            revert InterchainDB__NoModulesSpecified();
        }
        fees = new uint256[](len);
        for (uint256 i = 0; i < len; ++i) {
            fees[i] = IInterchainModule(srcModules[i]).getModuleFee(destChainId);
            totalFee += fees[i];
        }
    }
}
