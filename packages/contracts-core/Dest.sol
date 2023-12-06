
// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════

// ══════════════════════════════ INVALID CALLER ═══════════════════════════════

error CallerNotAgentManager();
error CallerNotDestination();
error CallerNotInbox();
error CallerNotSummit();

// ══════════════════════════════ INCORRECT DATA ═══════════════════════════════

error IncorrectAttestation();
error IncorrectAgentDomain();
error IncorrectAgentIndex();
error IncorrectAgentProof();
error IncorrectAgentRoot();
error IncorrectDataHash();
error IncorrectDestinationDomain();
error IncorrectOriginDomain();
error IncorrectSnapshotProof();
error IncorrectSnapshotRoot();
error IncorrectState();
error IncorrectStatesAmount();
error IncorrectTipsProof();
error IncorrectVersionLength();

error IncorrectNonce();
error IncorrectSender();
error IncorrectRecipient();

error FlagOutOfRange();
error IndexOutOfRange();
error NonceOutOfRange();

error OutdatedNonce();

error UnformattedAttestation();
error UnformattedAttestationReport();
error UnformattedBaseMessage();
error UnformattedCallData();
error UnformattedCallDataPrefix();
error UnformattedMessage();
error UnformattedReceipt();
error UnformattedReceiptReport();
error UnformattedSignature();
error UnformattedSnapshot();
error UnformattedState();
error UnformattedStateReport();

// ═══════════════════════════════ MERKLE TREES ════════════════════════════════

error LeafNotProven();
error MerkleTreeFull();
error NotEnoughLeafs();
error TreeHeightTooLow();

// ═════════════════════════════ OPTIMISTIC PERIOD ═════════════════════════════

error BaseClientOptimisticPeriod();
error MessageOptimisticPeriod();
error SlashAgentOptimisticPeriod();
error WithdrawTipsOptimisticPeriod();
error ZeroProofMaturity();

// ═══════════════════════════════ AGENT MANAGER ═══════════════════════════════

error AgentNotGuard();
error AgentNotNotary();

error AgentCantBeAdded();
error AgentNotActive();
error AgentNotActiveNorUnstaking();
error AgentNotFraudulent();
error AgentNotUnstaking();
error AgentUnknown();

error AgentRootNotProposed();
error AgentRootTimeoutNotOver();

error NotStuck();

error DisputeAlreadyResolved();
error DisputeNotOpened();
error DisputeTimeoutNotOver();
error GuardInDispute();
error NotaryInDispute();

error MustBeSynapseDomain();
error SynapseDomainForbidden();

// ════════════════════════════════ DESTINATION ════════════════════════════════

error AlreadyExecuted();
error AlreadyFailed();
error DuplicatedSnapshotRoot();
error IncorrectMagicValue();
error GasLimitTooLow();
error GasSuppliedTooLow();

// ══════════════════════════════════ ORIGIN ═══════════════════════════════════

error ContentLengthTooBig();
error EthTransferFailed();
error InsufficientEthBalance();

// ════════════════════════════════ GAS ORACLE ═════════════════════════════════

error LocalGasDataNotSet();
error RemoteGasDataNotSet();

// ═══════════════════════════════════ TIPS ════════════════════════════════════

error SummitTipTooHigh();
error TipsClaimMoreThanEarned();
error TipsClaimZero();
error TipsOverflow();
error TipsValueTooLow();

// ════════════════════════════════ MEMORY VIEW ════════════════════════════════

error IndexedTooMuch();
error ViewOverrun();
error OccupiedMemory();
error UnallocatedMemory();
error PrecompileOutOfGas();

// ═════════════════════════════════ MULTICALL ═════════════════════════════════

error MulticallFailed();

/// @dev MemView is an untyped view over a portion of memory to be used instead of `bytes memory`
type MemView is uint256;

/// @dev Attach library functions to MemView
using MemViewLib for MemView global;

/// @notice Library for operations with the memory views.
/// Forked from https://github.com/summa-tx/memview-sol with several breaking changes:
/// - The codebase is ported to Solidity 0.8
/// - Custom errors are added
/// - The runtime type checking is replaced with compile-time check provided by User-Defined Value Types
///   https://docs.soliditylang.org/en/latest/types.html#user-defined-value-types
/// - uint256 is used as the underlying type for the "memory view" instead of bytes29.
///   It is wrapped into MemView custom type in order not to be confused with actual integers.
/// - Therefore the "type" field is discarded, allowing to allocate 16 bytes for both view location and length
/// - The documentation is expanded
/// - Library functions unused by the rest of the codebase are removed
//  - Very pretty code separators are added :)
library MemViewLib {
    /// @notice Stack layout for uint256 (from highest bits to lowest)
    /// (32 .. 16]      loc     16 bytes    Memory address of underlying bytes
    /// (16 .. 00]      len     16 bytes    Length of underlying bytes

    // ═══════════════════════════════════════════ BUILDING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Instantiate a new untyped memory view. This should generally not be called directly.
     * Prefer `ref` wherever possible.
     * @param loc_          The memory address
     * @param len_          The length
     * @return The new view with the specified location and length
     */
    function build(uint256 loc_, uint256 len_) internal pure returns (MemView) {
        uint256 end_ = loc_ + len_;
        // Make sure that a view is not constructed that points to unallocated memory
        // as this could be indicative of a buffer overflow attack
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            if gt(end_, mload(0x40)) { end_ := 0 }
        }
        if (end_ == 0) {
            revert UnallocatedMemory();
        }
        return _unsafeBuildUnchecked(loc_, len_);
    }

    /**
     * @notice Instantiate a memory view from a byte array.
     * @dev Note that due to Solidity memory representation, it is not possible to
     * implement a deref, as the `bytes` type stores its len in memory.
     * @param arr           The byte array
     * @return The memory view over the provided byte array
     */
    function ref(bytes memory arr) internal pure returns (MemView) {
        uint256 len_ = arr.length;
        // `bytes arr` is stored in memory in the following way
        // 1. First, uint256 arr.length is stored. That requires 32 bytes (0x20).
        // 2. Then, the array data is stored.
        uint256 loc_;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // We add 0x20, so that the view starts exactly where the array data starts
            loc_ := add(arr, 0x20)
        }
        return build(loc_, len_);
    }

    // ════════════════════════════════════════════ CLONING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Copies the referenced memory to a new loc in memory, returning a `bytes` pointing to the new memory.
     * @param memView       The memory view
     * @return arr          The cloned byte array
     */
    function clone(MemView memView) internal view returns (bytes memory arr) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
            // This is where the byte array will be stored
            arr := ptr
        }
        unchecked {
            _unsafeCopyTo(memView, ptr + 0x20);
        }
        // `bytes arr` is stored in memory in the following way
        // 1. First, uint256 arr.length is stored. That requires 32 bytes (0x20).
        // 2. Then, the array data is stored.
        uint256 len_ = memView.len();
        uint256 footprint_ = memView.footprint();
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Write new unused pointer: the old value + array footprint + 32 bytes to store the length
            mstore(0x40, add(add(ptr, footprint_), 0x20))
            // Write len of new array (in bytes)
            mstore(ptr, len_)
        }
    }

    /**
     * @notice Copies all views, joins them into a new bytearray.
     * @param memViews      The memory views
     * @return arr          The new byte array with joined data behind the given views
     */
    function join(MemView[] memory memViews) internal view returns (bytes memory arr) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
            // This is where the byte array will be stored
            arr := ptr
        }
        MemView newView;
        unchecked {
            newView = _unsafeJoin(memViews, ptr + 0x20);
        }
        uint256 len_ = newView.len();
        uint256 footprint_ = newView.footprint();
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Write new unused pointer: the old value + array footprint + 32 bytes to store the length
            mstore(0x40, add(add(ptr, footprint_), 0x20))
            // Write len of new array (in bytes)
            mstore(ptr, len_)
        }
    }

    // ══════════════════════════════════════════ INSPECTING MEMORY VIEW ═══════════════════════════════════════════════

    /**
     * @notice Returns the memory address of the underlying bytes.
     * @param memView       The memory view
     * @return loc_         The memory address
     */
    function loc(MemView memView) internal pure returns (uint256 loc_) {
        // loc is stored in the highest 16 bytes of the underlying uint256
        return MemView.unwrap(memView) >> 128;
    }

    /**
     * @notice Returns the number of bytes of the view.
     * @param memView       The memory view
     * @return len_         The length of the view
     */
    function len(MemView memView) internal pure returns (uint256 len_) {
        // len is stored in the lowest 16 bytes of the underlying uint256
        return MemView.unwrap(memView) & type(uint128).max;
    }

    /**
     * @notice Returns the endpoint of `memView`.
     * @param memView       The memory view
     * @return end_         The endpoint of `memView`
     */
    function end(MemView memView) internal pure returns (uint256 end_) {
        // The endpoint never overflows uint128, let alone uint256, so we could use unchecked math here
        unchecked {
            return memView.loc() + memView.len();
        }
    }

    /**
     * @notice Returns the number of memory words this memory view occupies, rounded up.
     * @param memView       The memory view
     * @return words_       The number of memory words
     */
    function words(MemView memView) internal pure returns (uint256 words_) {
        // returning ceil(length / 32.0)
        unchecked {
            return (memView.len() + 31) >> 5;
        }
    }

    /**
     * @notice Returns the in-memory footprint of a fresh copy of the view.
     * @param memView       The memory view
     * @return footprint_   The in-memory footprint of a fresh copy of the view.
     */
    function footprint(MemView memView) internal pure returns (uint256 footprint_) {
        // words() * 32
        return memView.words() << 5;
    }

    // ════════════════════════════════════════════ HASHING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Returns the keccak256 hash of the underlying memory
     * @param memView       The memory view
     * @return digest       The keccak256 hash of the underlying memory
     */
    function keccak(MemView memView) internal pure returns (bytes32 digest) {
        uint256 loc_ = memView.loc();
        uint256 len_ = memView.len();
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            digest := keccak256(loc_, len_)
        }
    }

    /**
     * @notice Adds a salt to the keccak256 hash of the underlying data and returns the keccak256 hash of the
     * resulting data.
     * @param memView       The memory view
     * @return digestSalted keccak256(salt, keccak256(memView))
     */
    function keccakSalted(MemView memView, bytes32 salt) internal pure returns (bytes32 digestSalted) {
        return keccak256(bytes.concat(salt, memView.keccak()));
    }

    // ════════════════════════════════════════════ SLICING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Safe slicing without memory modification.
     * @param memView       The memory view
     * @param index_        The start index
     * @param len_          The length
     * @return The new view for the slice of the given length starting from the given index
     */
    function slice(MemView memView, uint256 index_, uint256 len_) internal pure returns (MemView) {
        uint256 loc_ = memView.loc();
        // Ensure it doesn't overrun the view
        if (loc_ + index_ + len_ > memView.end()) {
            revert ViewOverrun();
        }
        // Build a view starting from index with the given length
        unchecked {
            // loc_ + index_ <= memView.end()
            return build({loc_: loc_ + index_, len_: len_});
        }
    }

    /**
     * @notice Shortcut to `slice`. Gets a view representing bytes from `index` to end(memView).
     * @param memView       The memory view
     * @param index_        The start index
     * @return The new view for the slice starting from the given index until the initial view endpoint
     */
    function sliceFrom(MemView memView, uint256 index_) internal pure returns (MemView) {
        uint256 len_ = memView.len();
        // Ensure it doesn't overrun the view
        if (index_ > len_) {
            revert ViewOverrun();
        }
        // Build a view starting from index with the given length
        unchecked {
            // index_ <= len_ => memView.loc() + index_ <= memView.loc() + memView.len() == memView.end()
            return build({loc_: memView.loc() + index_, len_: len_ - index_});
        }
    }

    /**
     * @notice Shortcut to `slice`. Gets a view representing the first `len` bytes.
     * @param memView       The memory view
     * @param len_          The length
     * @return The new view for the slice of the given length starting from the initial view beginning
     */
    function prefix(MemView memView, uint256 len_) internal pure returns (MemView) {
        return memView.slice({index_: 0, len_: len_});
    }

    /**
     * @notice Shortcut to `slice`. Gets a view representing the last `len` byte.
     * @param memView       The memory view
     * @param len_          The length
     * @return The new view for the slice of the given length until the initial view endpoint
     */
    function postfix(MemView memView, uint256 len_) internal pure returns (MemView) {
        uint256 viewLen = memView.len();
        // Ensure it doesn't overrun the view
        if (len_ > viewLen) {
            revert ViewOverrun();
        }
        // Could do the unchecked math due to the check above
        uint256 index_;
        unchecked {
            index_ = viewLen - len_;
        }
        // Build a view starting from index with the given length
        unchecked {
            // len_ <= memView.len() => memView.loc() <= loc_ <= memView.end()
            return build({loc_: memView.loc() + viewLen - len_, len_: len_});
        }
    }

    // ═══════════════════════════════════════════ INDEXING MEMORY VIEW ════════════════════════════════════════════════

    /**
     * @notice Load up to 32 bytes from the view onto the stack.
     * @dev Returns a bytes32 with only the `bytes_` HIGHEST bytes set.
     * This can be immediately cast to a smaller fixed-length byte array.
     * To automatically cast to an integer, use `indexUint`.
     * @param memView       The memory view
     * @param index_        The index
     * @param bytes_        The amount of bytes to load onto the stack
     * @return result       The 32 byte result having only `bytes_` highest bytes set
     */
    function index(MemView memView, uint256 index_, uint256 bytes_) internal pure returns (bytes32 result) {
        if (bytes_ == 0) {
            return bytes32(0);
        }
        // Can't load more than 32 bytes to the stack in one go
        if (bytes_ > 32) {
            revert IndexedTooMuch();
        }
        // The last indexed byte should be within view boundaries
        if (index_ + bytes_ > memView.len()) {
            revert ViewOverrun();
        }
        uint256 bitLength = bytes_ << 3; // bytes_ * 8
        uint256 loc_ = memView.loc();
        // Get a mask with `bitLength` highest bits set
        uint256 mask;
        // 0x800...00 binary representation is 100...00
        // sar stands for "signed arithmetic shift": https://en.wikipedia.org/wiki/Arithmetic_shift
        // sar(N-1, 100...00) = 11...100..00, with exactly N highest bits set to 1
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            mask := sar(sub(bitLength, 1), 0x8000000000000000000000000000000000000000000000000000000000000000)
        }
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load a full word using index offset, and apply mask to ignore non-relevant bytes
            result := and(mload(add(loc_, index_)), mask)
        }
    }

    /**
     * @notice Parse an unsigned integer from the view at `index`.
     * @dev Requires that the view have >= `bytes_` bytes following that index.
     * @param memView       The memory view
     * @param index_        The index
     * @param bytes_        The amount of bytes to load onto the stack
     * @return The unsigned integer
     */
    function indexUint(MemView memView, uint256 index_, uint256 bytes_) internal pure returns (uint256) {
        bytes32 indexedBytes = memView.index(index_, bytes_);
        // `index()` returns left-aligned `bytes_`, while integers are right-aligned
        // Shifting here to right-align with the full 32 bytes word: need to shift right `(32 - bytes_)` bytes
        unchecked {
            // memView.index() reverts when bytes_ > 32, thus unchecked math
            return uint256(indexedBytes) >> ((32 - bytes_) << 3);
        }
    }

    /**
     * @notice Parse an address from the view at `index`.
     * @dev Requires that the view have >= 20 bytes following that index.
     * @param memView       The memory view
     * @param index_        The index
     * @return The address
     */
    function indexAddress(MemView memView, uint256 index_) internal pure returns (address) {
        // index 20 bytes as `uint160`, and then cast to `address`
        return address(uint160(memView.indexUint(index_, 20)));
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns a memory view over the specified memory location
    /// without checking if it points to unallocated memory.
    function _unsafeBuildUnchecked(uint256 loc_, uint256 len_) private pure returns (MemView) {
        // There is no scenario where loc or len would overflow uint128, so we omit this check.
        // We use the highest 128 bits to encode the location and the lowest 128 bits to encode the length.
        return MemView.wrap((loc_ << 128) | len_);
    }

    /**
     * @notice Copy the view to a location, return an unsafe memory reference
     * @dev Super Dangerous direct memory access.
     * This reference can be overwritten if anything else modifies memory (!!!).
     * As such it MUST be consumed IMMEDIATELY. Update the free memory pointer to ensure the copied data
     * is not overwritten. This function is private to prevent unsafe usage by callers.
     * @param memView       The memory view
     * @param newLoc        The new location to copy the underlying view data
     * @return The memory view over the unsafe memory with the copied underlying data
     */
    function _unsafeCopyTo(MemView memView, uint256 newLoc) private view returns (MemView) {
        uint256 len_ = memView.len();
        uint256 oldLoc = memView.loc();

        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
        }
        // Revert if we're writing in occupied memory
        if (newLoc < ptr) {
            revert OccupiedMemory();
        }
        bool res;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // use the identity precompile (0x04) to copy
            res := staticcall(gas(), 0x04, oldLoc, len_, newLoc, len_)
        }
        if (!res) revert PrecompileOutOfGas();
        return _unsafeBuildUnchecked({loc_: newLoc, len_: len_});
    }

    /**
     * @notice Join the views in memory, return an unsafe reference to the memory.
     * @dev Super Dangerous direct memory access.
     * This reference can be overwritten if anything else modifies memory (!!!).
     * As such it MUST be consumed IMMEDIATELY. Update the free memory pointer to ensure the copied data
     * is not overwritten. This function is private to prevent unsafe usage by callers.
     * @param memViews      The memory views
     * @return The conjoined view pointing to the new memory
     */
    function _unsafeJoin(MemView[] memory memViews, uint256 location) private view returns (MemView) {
        uint256 ptr;
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            // Load unused memory pointer
            ptr := mload(0x40)
        }
        // Revert if we're writing in occupied memory
        if (location < ptr) {
            revert OccupiedMemory();
        }
        // Copy the views to the specified location one by one, by tracking the amount of copied bytes so far
        uint256 offset = 0;
        for (uint256 i = 0; i < memViews.length;) {
            MemView memView = memViews[i];
            // We can use the unchecked math here as location + sum(view.length) will never overflow uint256
            unchecked {
                _unsafeCopyTo(memView, location + offset);
                offset += memView.len();
                ++i;
            }
        }
        return _unsafeBuildUnchecked({loc_: location, len_: offset});
    }
}

// Here we define common constants to enable their easier reusing later.

// ══════════════════════════════════ MERKLE ═══════════════════════════════════
/// @dev Height of the Agent Merkle Tree
uint256 constant AGENT_TREE_HEIGHT = 32;
/// @dev Height of the Origin Merkle Tree
uint256 constant ORIGIN_TREE_HEIGHT = 32;
/// @dev Height of the Snapshot Merkle Tree. Allows up to 64 leafs, e.g. up to 32 states
uint256 constant SNAPSHOT_TREE_HEIGHT = 6;
// ══════════════════════════════════ STRUCTS ══════════════════════════════════
/// @dev See Attestation.sol: (bytes32,bytes32,uint32,uint40,uint40): 32+32+4+5+5
uint256 constant ATTESTATION_LENGTH = 78;
/// @dev See GasData.sol: (uint16,uint16,uint16,uint16,uint16,uint16): 2+2+2+2+2+2
uint256 constant GAS_DATA_LENGTH = 12;
/// @dev See Receipt.sol: (uint32,uint32,bytes32,bytes32,uint8,address,address,address): 4+4+32+32+1+20+20+20
uint256 constant RECEIPT_LENGTH = 133;
/// @dev See State.sol: (bytes32,uint32,uint32,uint40,uint40,GasData): 32+4+4+5+5+len(GasData)
uint256 constant STATE_LENGTH = 50 + GAS_DATA_LENGTH;
/// @dev Maximum amount of states in a single snapshot. Each state produces two leafs in the tree
uint256 constant SNAPSHOT_MAX_STATES = 1 << (SNAPSHOT_TREE_HEIGHT - 1);
// ══════════════════════════════════ MESSAGE ══════════════════════════════════
/// @dev See Header.sol: (uint8,uint32,uint32,uint32,uint32): 1+4+4+4+4
uint256 constant HEADER_LENGTH = 17;
/// @dev See Request.sol: (uint96,uint64,uint32): 12+8+4
uint256 constant REQUEST_LENGTH = 24;
/// @dev See Tips.sol: (uint64,uint64,uint64,uint64): 8+8+8+8
uint256 constant TIPS_LENGTH = 32;
/// @dev The amount of discarded last bits when encoding tip values
uint256 constant TIPS_GRANULARITY = 32;
/// @dev Tip values could be only the multiples of TIPS_MULTIPLIER
uint256 constant TIPS_MULTIPLIER = 1 << TIPS_GRANULARITY;
// ══════════════════════════════ STATEMENT SALTS ══════════════════════════════
/// @dev Salts for signing various statements
bytes32 constant ATTESTATION_VALID_SALT = keccak256("ATTESTATION_VALID_SALT");
bytes32 constant ATTESTATION_INVALID_SALT = keccak256("ATTESTATION_INVALID_SALT");
bytes32 constant RECEIPT_VALID_SALT = keccak256("RECEIPT_VALID_SALT");
bytes32 constant RECEIPT_INVALID_SALT = keccak256("RECEIPT_INVALID_SALT");
bytes32 constant SNAPSHOT_VALID_SALT = keccak256("SNAPSHOT_VALID_SALT");
bytes32 constant STATE_INVALID_SALT = keccak256("STATE_INVALID_SALT");
// ═════════════════════════════════ PROTOCOL ══════════════════════════════════
/// @dev Optimistic period for new agent roots in LightManager
uint32 constant AGENT_ROOT_OPTIMISTIC_PERIOD = 1 days;
/// @dev Timeout between the agent root could be proposed and resolved in LightManager
uint32 constant AGENT_ROOT_PROPOSAL_TIMEOUT = 12 hours;
uint32 constant BONDING_OPTIMISTIC_PERIOD = 1 days;
/// @dev Amount of time that the Notary will not be considered active after they won a dispute
uint32 constant DISPUTE_TIMEOUT_NOTARY = 12 hours;
/// @dev Amount of time without fresh data from Notaries before contract owner can resolve stuck disputes manually
uint256 constant FRESH_DATA_TIMEOUT = 4 hours;
/// @dev Maximum bytes per message = 2 KiB (somewhat arbitrarily set to begin)
uint256 constant MAX_CONTENT_BYTES = 2 * 2 ** 10;
/// @dev Maximum value for the summit tip that could be set in GasOracle
uint256 constant MAX_SUMMIT_TIP = 0.01 ether;

/// Attestation is a memory view over a formatted attestation payload.
type Attestation is uint256;

using AttestationLib for Attestation global;

/// # Attestation
/// Attestation structure represents the "Snapshot Merkle Tree" created from
/// every Notary snapshot accepted by the Summit contract. Attestation includes"
/// the root of the "Snapshot Merkle Tree", as well as additional metadata.
///
/// ## Steps for creation of "Snapshot Merkle Tree":
/// 1. The list of hashes is composed for states in the Notary snapshot.
/// 2. The list is padded with zero values until its length is 2**SNAPSHOT_TREE_HEIGHT.
/// 3. Values from the list are used as leafs and the merkle tree is constructed.
///
/// ## Differences between a State and Attestation
/// Similar to Origin, every derived Notary's "Snapshot Merkle Root" is saved in Summit contract.
/// The main difference is that Origin contract itself is keeping track of an incremental merkle tree,
/// by inserting the hash of the sent message and calculating the new "Origin Merkle Root".
/// While Summit relies on Guards and Notaries to provide snapshot data, which is used to calculate the
/// "Snapshot Merkle Root".
///
/// - Origin's State is "state of Origin Merkle Tree after N-th message was sent".
/// - Summit's Attestation is "data for the N-th accepted Notary Snapshot" + "agent merkle root at the
/// time snapshot was submitted" + "attestation metadata".
///
/// ## Attestation validity
/// - Attestation is considered "valid" in Summit contract, if it matches the N-th (nonce)
/// snapshot submitted by Notaries, as well as the historical agent merkle root.
/// - Attestation is considered "valid" in Origin contract, if its underlying Snapshot is "valid".
///
/// - This means that a snapshot could be "valid" in Summit contract and "invalid" in Origin, if the underlying
/// snapshot is invalid (i.e. one of the states in the list is invalid).
/// - The opposite could also be true. If a perfectly valid snapshot was never submitted to Summit, its attestation
/// would be valid in Origin, but invalid in Summit (it was never accepted, so the metadata would be incorrect).
///
/// - Attestation is considered "globally valid", if it is valid in the Summit and all the Origin contracts.
/// # Memory layout of Attestation fields
///
/// | Position   | Field       | Type    | Bytes | Description                                                    |
/// | ---------- | ----------- | ------- | ----- | -------------------------------------------------------------- |
/// | [000..032) | snapRoot    | bytes32 | 32    | Root for "Snapshot Merkle Tree" created from a Notary snapshot |
/// | [032..064) | dataHash    | bytes32 | 32    | Agent Root and SnapGasHash combined into a single hash         |
/// | [064..068) | nonce       | uint32  | 4     | Total amount of all accepted Notary snapshots                  |
/// | [068..073) | blockNumber | uint40  | 5     | Block when this Notary snapshot was accepted in Summit         |
/// | [073..078) | timestamp   | uint40  | 5     | Time when this Notary snapshot was accepted in Summit          |
///
/// @dev Attestation could be signed by a Notary and submitted to `Destination` in order to use if for proving
/// messages coming from origin chains that the initial snapshot refers to.
library AttestationLib {
    using MemViewLib for bytes;

    // TODO: compress three hashes into one?

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_SNAP_ROOT = 0;
    uint256 private constant OFFSET_DATA_HASH = 32;
    uint256 private constant OFFSET_NONCE = 64;
    uint256 private constant OFFSET_BLOCK_NUMBER = 68;
    uint256 private constant OFFSET_TIMESTAMP = 73;

    // ════════════════════════════════════════════════ ATTESTATION ════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Attestation payload with provided fields.
     * @param snapRoot_     Snapshot merkle tree's root
     * @param dataHash_     Agent Root and SnapGasHash combined into a single hash
     * @param nonce_        Attestation Nonce
     * @param blockNumber_  Block number when attestation was created in Summit
     * @param timestamp_    Block timestamp when attestation was created in Summit
     * @return Formatted attestation
     */
    function formatAttestation(
        bytes32 snapRoot_,
        bytes32 dataHash_,
        uint32 nonce_,
        uint40 blockNumber_,
        uint40 timestamp_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(snapRoot_, dataHash_, nonce_, blockNumber_, timestamp_);
    }

    /**
     * @notice Returns an Attestation view over the given payload.
     * @dev Will revert if the payload is not an attestation.
     */
    function castToAttestation(bytes memory payload) internal pure returns (Attestation) {
        return castToAttestation(payload.ref());
    }

    /**
     * @notice Casts a memory view to an Attestation view.
     * @dev Will revert if the memory view is not over an attestation.
     */
    function castToAttestation(MemView memView) internal pure returns (Attestation) {
        if (!isAttestation(memView)) revert UnformattedAttestation();
        return Attestation.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Attestation.
    function isAttestation(MemView memView) internal pure returns (bool) {
        return memView.len() == ATTESTATION_LENGTH;
    }

    /// @notice Returns the hash of an Attestation, that could be later signed by a Notary to signal
    /// that the attestation is valid.
    function hashValid(Attestation att) internal pure returns (bytes32) {
        // The final hash to sign is keccak(attestationSalt, keccak(attestation))
        return att.unwrap().keccakSalted(ATTESTATION_VALID_SALT);
    }

    /// @notice Returns the hash of an Attestation, that could be later signed by a Guard to signal
    /// that the attestation is invalid.
    function hashInvalid(Attestation att) internal pure returns (bytes32) {
        // The final hash to sign is keccak(attestationInvalidSalt, keccak(attestation))
        return att.unwrap().keccakSalted(ATTESTATION_INVALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Attestation att) internal pure returns (MemView) {
        return MemView.wrap(Attestation.unwrap(att));
    }

    // ════════════════════════════════════════════ ATTESTATION SLICING ════════════════════════════════════════════════

    /// @notice Returns root of the Snapshot merkle tree created in the Summit contract.
    function snapRoot(Attestation att) internal pure returns (bytes32) {
        return att.unwrap().index({index_: OFFSET_SNAP_ROOT, bytes_: 32});
    }

    /// @notice Returns hash of the Agent Root and SnapGasHash combined into a single hash.
    function dataHash(Attestation att) internal pure returns (bytes32) {
        return att.unwrap().index({index_: OFFSET_DATA_HASH, bytes_: 32});
    }

    /// @notice Returns hash of the Agent Root and SnapGasHash combined into a single hash.
    function dataHash(bytes32 agentRoot_, bytes32 snapGasHash_) internal pure returns (bytes32) {
        return keccak256(bytes.concat(agentRoot_, snapGasHash_));
    }

    /// @notice Returns nonce of Summit contract at the time, when attestation was created.
    function nonce(Attestation att) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(att.unwrap().indexUint({index_: OFFSET_NONCE, bytes_: 4}));
    }

    /// @notice Returns a block number when attestation was created in Summit.
    function blockNumber(Attestation att) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(att.unwrap().indexUint({index_: OFFSET_BLOCK_NUMBER, bytes_: 5}));
    }

    /// @notice Returns a block timestamp when attestation was created in Summit.
    /// @dev This is the timestamp according to the Synapse Chain.
    function timestamp(Attestation att) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(att.unwrap().indexUint({index_: OFFSET_TIMESTAMP, bytes_: 5}));
    }
}

/// @dev CallData is a memory view over the payload to be used for an external call, i.e.
/// recipient.call(callData). Its length is always (4 + 32 * N) bytes:
/// - First 4 bytes represent the function selector.
/// - 32 * N bytes represent N words that function arguments occupy.
type CallData is uint256;

/// @dev Attach library functions to CallData
using ByteString for CallData global;

/// @dev Signature is a memory view over a "65 bytes" array representing a ECDSA signature.
type Signature is uint256;

/// @dev Attach library functions to Signature
using ByteString for Signature global;

library ByteString {
    using MemViewLib for bytes;

    /**
     * @dev non-compact ECDSA signatures are enforced as of OZ 4.7.3
     *
     *      Signature payload memory layout
     * [000 .. 032) r   bytes32 32 bytes
     * [032 .. 064) s   bytes32 32 bytes
     * [064 .. 065) v   uint8    1 byte
     */
    uint256 internal constant SIGNATURE_LENGTH = 65;
    uint256 private constant OFFSET_R = 0;
    uint256 private constant OFFSET_S = 32;
    uint256 private constant OFFSET_V = 64;

    /**
     * @dev Calldata memory layout
     * [000 .. 004) selector    bytes4  4 bytes
     *      Optional: N function arguments
     * [004 .. 036) arg1        bytes32 32 bytes
     *      ..
     * [AAA .. END) argN        bytes32 32 bytes
     */
    uint256 internal constant SELECTOR_LENGTH = 4;
    uint256 private constant OFFSET_SELECTOR = 0;
    uint256 private constant OFFSET_ARGUMENTS = SELECTOR_LENGTH;

    // ═════════════════════════════════════════════════ SIGNATURE ═════════════════════════════════════════════════════

    /**
     * @notice Constructs the signature payload from the given values.
     * @dev Using ByteString.formatSignature({r: r, s: s, v: v}) will make sure
     * that params are given in the right order.
     */
    function formatSignature(bytes32 r, bytes32 s, uint8 v) internal pure returns (bytes memory) {
        return abi.encodePacked(r, s, v);
    }

    /**
     * @notice Returns a Signature view over for the given payload.
     * @dev Will revert if the payload is not a signature.
     */
    function castToSignature(bytes memory payload) internal pure returns (Signature) {
        return castToSignature(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Signature view.
     * @dev Will revert if the memory view is not over a signature.
     */
    function castToSignature(MemView memView) internal pure returns (Signature) {
        if (!isSignature(memView)) revert UnformattedSignature();
        return Signature.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a byte string is a signature
     */
    function isSignature(MemView memView) internal pure returns (bool) {
        return memView.len() == SIGNATURE_LENGTH;
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Signature signature) internal pure returns (MemView) {
        return MemView.wrap(Signature.unwrap(signature));
    }

    // ═════════════════════════════════════════════ SIGNATURE SLICING ═════════════════════════════════════════════════

    /// @notice Unpacks signature payload into (r, s, v) parameters.
    /// @dev Make sure to verify signature length with isSignature() beforehand.
    function toRSV(Signature signature) internal pure returns (bytes32 r, bytes32 s, uint8 v) {
        // Get the underlying memory view
        MemView memView = unwrap(signature);
        r = memView.index({index_: OFFSET_R, bytes_: 32});
        s = memView.index({index_: OFFSET_S, bytes_: 32});
        // Can be safely casted to uint8, since we index a single byte
        v = uint8(memView.indexUint({index_: OFFSET_V, bytes_: 1}));
    }

    // ═════════════════════════════════════════════════ CALLDATA ══════════════════════════════════════════════════════

    /**
     * @notice Constructs the calldata with the modified arguments:
     * the existing arguments are prepended with the arguments from the prefix.
     * @dev Given:
     *  - `calldata = abi.encodeWithSelector(foo.selector, d, e);`
     *  - `prefix = abi.encode(a, b, c);`
     *  - `a`, `b`, `c` are arguments of static type (i.e. not dynamically sized ones)
     *      Then:
     *  - Function will return abi.encodeWithSelector(foo.selector, a, c, c, d, e)
     *  - Returned calldata will trigger `foo(a, b, c, d, e)` when used for a contract call.
     * Note: for clarification as to what types are considered static, see
     * https://docs.soliditylang.org/en/latest/abi-spec.html#formal-specification-of-the-encoding
     * @param callData  Calldata that needs to be modified
     * @param prefix    ABI-encoded arguments to use as the first arguments in the new calldata
     * @return Modified calldata having prefix as the first arguments.
     */
    function addPrefix(CallData callData, bytes memory prefix) internal view returns (bytes memory) {
        // Prefix should occupy a whole amount of words in memory
        if (!_fullWords(prefix.length)) revert UnformattedCallDataPrefix();
        MemView[] memory views = new MemView[](3);
        // Use payload's function selector
        views[0] = abi.encodePacked(callData.callSelector()).ref();
        // Use prefix as the first arguments
        views[1] = prefix.ref();
        // Use payload's remaining arguments
        views[2] = callData.arguments();
        return MemViewLib.join(views);
    }

    /**
     * @notice Returns a CallData view over for the given payload.
     * @dev Will revert if the memory view is not over a calldata.
     */
    function castToCallData(bytes memory payload) internal pure returns (CallData) {
        return castToCallData(payload.ref());
    }

    /**
     * @notice Casts a memory view to a CallData view.
     * @dev Will revert if the memory view is not over a calldata.
     */
    function castToCallData(MemView memView) internal pure returns (CallData) {
        if (!isCallData(memView)) revert UnformattedCallData();
        return CallData.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a byte string is a valid calldata, i.e.
     * a function selector, followed by arbitrary amount of arguments.
     */
    function isCallData(MemView memView) internal pure returns (bool) {
        uint256 length = memView.len();
        // Calldata should at least have a function selector
        if (length < SELECTOR_LENGTH) return false;
        // The remainder of the calldata should be exactly N memory words (N >= 0)
        return _fullWords(length - SELECTOR_LENGTH);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(CallData callData) internal pure returns (MemView) {
        return MemView.wrap(CallData.unwrap(callData));
    }

    /// @notice Returns callData's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(CallData callData) internal pure returns (bytes32) {
        return callData.unwrap().keccak();
    }

    // ═════════════════════════════════════════════ CALLDATA SLICING ══════════════════════════════════════════════════

    /**
     * @notice Returns amount of memory words (32 byte chunks) the function arguments
     * occupy in the calldata.
     * @dev This might differ from amount of arguments supplied, if any of the arguments
     * occupies more than one memory slot. It is true, however, that argument part of the payload
     * occupies exactly N words, even for dynamic types like `bytes`
     */
    function argumentWords(CallData callData) internal pure returns (uint256) {
        // Get the underlying memory view
        MemView memView = unwrap(callData);
        // Equivalent of (length - SELECTOR_LENGTH) / 32
        return (memView.len() - SELECTOR_LENGTH) >> 5;
    }

    /// @notice Returns selector for the provided calldata.
    function callSelector(CallData callData) internal pure returns (bytes4) {
        // Get the underlying memory view
        MemView memView = unwrap(callData);
        return bytes4(memView.index({index_: OFFSET_SELECTOR, bytes_: SELECTOR_LENGTH}));
    }

    /// @notice Returns abi encoded arguments for the provided calldata.
    function arguments(CallData callData) internal pure returns (MemView) {
        // Get the underlying memory view
        MemView memView = unwrap(callData);
        return memView.sliceFrom({index_: OFFSET_ARGUMENTS});
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Checks if length is full amount of memory words (32 bytes).
    function _fullWords(uint256 length) internal pure returns (bool) {
        // The equivalent of length % 32 == 0
        return length & 31 == 0;
    }
}

/// Number is a compact representation of uint256, that is fit into 16 bits
/// with the maximum relative error under 0.4%.
type Number is uint16;

using NumberLib for Number global;

/// # Number
/// Library for compact representation of uint256 numbers.
/// - Number is stored using mantissa and exponent, each occupying 8 bits.
/// - Numbers under 2**8 are stored as `mantissa` with `exponent = 0xFF`.
/// - Numbers at least 2**8 are approximated as `(256 + mantissa) << exponent`
/// > - `0 <= mantissa < 256`
/// > - `0 <= exponent <= 247` (`256 * 2**248` doesn't fit into uint256)
/// # Number stack layout (from highest bits to lowest)
///
/// | Position   | Field    | Type  | Bytes |
/// | ---------- | -------- | ----- | ----- |
/// | (002..001] | mantissa | uint8 | 1     |
/// | (001..000] | exponent | uint8 | 1     |

library NumberLib {
    /// @dev Amount of bits to shift to mantissa field
    uint16 private constant SHIFT_MANTISSA = 8;

    /// @notice For bwad math (binary wad) we use 2**64 as "wad" unit.
    /// @dev We are using not using 10**18 as wad, because it is not stored precisely in NumberLib.
    uint256 internal constant BWAD_SHIFT = 64;
    uint256 internal constant BWAD = 1 << BWAD_SHIFT;
    /// @notice ~0.1% in bwad units.
    uint256 internal constant PER_MILLE_SHIFT = BWAD_SHIFT - 10;
    uint256 internal constant PER_MILLE = 1 << PER_MILLE_SHIFT;

    /// @notice Compresses uint256 number into 16 bits.
    function compress(uint256 value) internal pure returns (Number) {
        // Find `msb` such as `2**msb <= value < 2**(msb + 1)`
        uint256 msb = mostSignificantBit(value);
        // We want to preserve 9 bits of precision.
        // The highest bit is always 1, so we can skip it.
        // The remaining 8 highest bits are stored as mantissa.
        if (msb < 8) {
            // Value is less than 2**8, so we can use value as mantissa with "-1" exponent.
            return _encode(uint8(value), 0xFF);
        } else {
            // We use `msb - 8` as exponent otherwise. Note that `exponent >= 0`.
            unchecked {
                uint256 exponent = msb - 8;
                // Shifting right by `msb-8` bits will shift the "remaining 8 highest bits" into the 8 lowest bits.
                // uint8() will truncate the highest bit.
                return _encode(uint8(value >> exponent), uint8(exponent));
            }
        }
    }

    /// @notice Decompresses 16 bits number into uint256.
    /// @dev The outcome is an approximation of the original number: `(value - value / 256) < number <= value`.
    function decompress(Number number) internal pure returns (uint256 value) {
        // Isolate 8 highest bits as the mantissa.
        uint256 mantissa = Number.unwrap(number) >> SHIFT_MANTISSA;
        // This will truncate the highest bits, leaving only the exponent.
        uint256 exponent = uint8(Number.unwrap(number));
        if (exponent == 0xFF) {
            return mantissa;
        } else {
            unchecked {
                return (256 + mantissa) << (exponent);
            }
        }
    }

    /// @dev Returns the most significant bit of `x`
    /// https://solidity-by-example.org/bitwise/
    function mostSignificantBit(uint256 x) internal pure returns (uint256 msb) {
        // To find `msb` we determine it bit by bit, starting from the highest one.
        // `0 <= msb <= 255`, so we start from the highest bit, 1<<7 == 128.
        // If `x` is at least 2**128, then the highest bit of `x` is at least 128.
        // solhint-disable no-inline-assembly
        assembly {
            // `f` is set to 1<<7 if `x >= 2**128` and to 0 otherwise.
            let f := shl(7, gt(x, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF))
            // If `x >= 2**128` then set `msb` highest bit to 1 and shift `x` right by 128.
            // Otherwise, `msb` remains 0 and `x` remains unchanged.
            x := shr(f, x)
            msb := or(msb, f)
        }
        // `x` is now at most 2**128 - 1. Continue the same way, the next highest bit is 1<<6 == 64.
        assembly {
            // `f` is set to 1<<6 if `x >= 2**64` and to 0 otherwise.
            let f := shl(6, gt(x, 0xFFFFFFFFFFFFFFFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<5 if `x >= 2**32` and to 0 otherwise.
            let f := shl(5, gt(x, 0xFFFFFFFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<4 if `x >= 2**16` and to 0 otherwise.
            let f := shl(4, gt(x, 0xFFFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<3 if `x >= 2**8` and to 0 otherwise.
            let f := shl(3, gt(x, 0xFF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<2 if `x >= 2**4` and to 0 otherwise.
            let f := shl(2, gt(x, 0xF))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1<<1 if `x >= 2**2` and to 0 otherwise.
            let f := shl(1, gt(x, 0x3))
            x := shr(f, x)
            msb := or(msb, f)
        }
        assembly {
            // `f` is set to 1 if `x >= 2**1` and to 0 otherwise.
            let f := gt(x, 0x1)
            msb := or(msb, f)
        }
    }

    /// @dev Wraps (mantissa, exponent) pair into Number.
    function _encode(uint8 mantissa, uint8 exponent) private pure returns (Number) {
        // Casts below are upcasts, so they are safe.
        return Number.wrap(uint16(mantissa) << SHIFT_MANTISSA | uint16(exponent));
    }
}

/// GasData in encoded data with "basic information about gas prices" for some chain.
type GasData is uint96;

using GasDataLib for GasData global;

/// ChainGas is encoded data with given chain's "basic information about gas prices".
type ChainGas is uint128;

using GasDataLib for ChainGas global;

/// Library for encoding and decoding GasData and ChainGas structs.
/// # GasData
/// `GasData` is a struct to store the "basic information about gas prices", that could
/// be later used to approximate the cost of a message execution, and thus derive the
/// minimal tip values for sending a message to the chain.
/// > - `GasData` is supposed to be cached by `GasOracle` contract, allowing to store the
/// > approximates instead of the exact values, and thus save on storage costs.
/// > - For instance, if `GasOracle` only updates the values on +- 10% change, having an
/// > 0.4% error on the approximates would be acceptable.
/// `GasData` is supposed to be included in the Origin's state, which are synced across
/// chains using Agent-signed snapshots and attestations.
/// ## GasData stack layout (from highest bits to lowest)
///
/// | Position   | Field        | Type   | Bytes | Description                                         |
/// | ---------- | ------------ | ------ | ----- | --------------------------------------------------- |
/// | (012..010] | gasPrice     | uint16 | 2     | Gas price for the chain (in Wei per gas unit)       |
/// | (010..008] | dataPrice    | uint16 | 2     | Calldata price (in Wei per byte of content)         |
/// | (008..006] | execBuffer   | uint16 | 2     | Tx fee safety buffer for message execution (in Wei) |
/// | (006..004] | amortAttCost | uint16 | 2     | Amortized cost for attestation submission (in Wei)  |
/// | (004..002] | etherPrice   | uint16 | 2     | Chain's Ether Price / Mainnet Ether Price (in BWAD) |
/// | (002..000] | markup       | uint16 | 2     | Markup for the message execution (in BWAD)          |
/// > See Number.sol for more details on `Number` type and BWAD (binary WAD) math.
///
/// ## ChainGas stack layout (from highest bits to lowest)
///
/// | Position   | Field   | Type   | Bytes | Description      |
/// | ---------- | ------- | ------ | ----- | ---------------- |
/// | (016..004] | gasData | uint96 | 12    | Chain's gas data |
/// | (004..000] | domain  | uint32 | 4     | Chain's domain   |
library GasDataLib {
    /// @dev Amount of bits to shift to gasPrice field
    uint96 private constant SHIFT_GAS_PRICE = 10 * 8;
    /// @dev Amount of bits to shift to dataPrice field
    uint96 private constant SHIFT_DATA_PRICE = 8 * 8;
    /// @dev Amount of bits to shift to execBuffer field
    uint96 private constant SHIFT_EXEC_BUFFER = 6 * 8;
    /// @dev Amount of bits to shift to amortAttCost field
    uint96 private constant SHIFT_AMORT_ATT_COST = 4 * 8;
    /// @dev Amount of bits to shift to etherPrice field
    uint96 private constant SHIFT_ETHER_PRICE = 2 * 8;

    /// @dev Amount of bits to shift to gasData field
    uint128 private constant SHIFT_GAS_DATA = 4 * 8;

    // ═════════════════════════════════════════════════ GAS DATA ══════════════════════════════════════════════════════

    /// @notice Returns an encoded GasData struct with the given fields.
    /// @param gasPrice_        Gas price for the chain (in Wei per gas unit)
    /// @param dataPrice_       Calldata price (in Wei per byte of content)
    /// @param execBuffer_      Tx fee safety buffer for message execution (in Wei)
    /// @param amortAttCost_    Amortized cost for attestation submission (in Wei)
    /// @param etherPrice_      Ratio of Chain's Ether Price / Mainnet Ether Price (in BWAD)
    /// @param markup_          Markup for the message execution (in BWAD)
    function encodeGasData(
        Number gasPrice_,
        Number dataPrice_,
        Number execBuffer_,
        Number amortAttCost_,
        Number etherPrice_,
        Number markup_
    ) internal pure returns (GasData) {
        // Number type wraps uint16, so could safely be casted to uint96
        // forgefmt: disable-next-item
        return GasData.wrap(
            uint96(Number.unwrap(gasPrice_)) << SHIFT_GAS_PRICE |
            uint96(Number.unwrap(dataPrice_)) << SHIFT_DATA_PRICE |
            uint96(Number.unwrap(execBuffer_)) << SHIFT_EXEC_BUFFER |
            uint96(Number.unwrap(amortAttCost_)) << SHIFT_AMORT_ATT_COST |
            uint96(Number.unwrap(etherPrice_)) << SHIFT_ETHER_PRICE |
            uint96(Number.unwrap(markup_))
        );
    }

    /// @notice Wraps padded uint256 value into GasData struct.
    function wrapGasData(uint256 paddedGasData) internal pure returns (GasData) {
        // Casting to uint96 will truncate the highest bits, which is the behavior we want
        return GasData.wrap(uint96(paddedGasData));
    }

    /// @notice Returns the gas price, in Wei per gas unit.
    function gasPrice(GasData data) internal pure returns (Number) {
        // Casting to uint16 will truncate the highest bits, which is the behavior we want
        return Number.wrap(uint16(GasData.unwrap(data) >> SHIFT_GAS_PRICE));
    }

    /// @notice Returns the calldata price, in Wei per byte of content.
    function dataPrice(GasData data) internal pure returns (Number) {
        // Casting to uint16 will truncate the highest bits, which is the behavior we want
        return Number.wrap(uint16(GasData.unwrap(data) >> SHIFT_DATA_PRICE));
    }

    /// @notice Returns the tx fee safety buffer for message execution, in Wei.
    function execBuffer(GasData data) internal pure returns (Number) {
        // Casting to uint16 will truncate the highest bits, which is the behavior we want
        return Number.wrap(uint16(GasData.unwrap(data) >> SHIFT_EXEC_BUFFER));
    }

    /// @notice Returns the amortized cost for attestation submission, in Wei.
    function amortAttCost(GasData data) internal pure returns (Number) {
        // Casting to uint16 will truncate the highest bits, which is the behavior we want
        return Number.wrap(uint16(GasData.unwrap(data) >> SHIFT_AMORT_ATT_COST));
    }

    /// @notice Returns the ratio of Chain's Ether Price / Mainnet Ether Price, in BWAD math.
    function etherPrice(GasData data) internal pure returns (Number) {
        // Casting to uint16 will truncate the highest bits, which is the behavior we want
        return Number.wrap(uint16(GasData.unwrap(data) >> SHIFT_ETHER_PRICE));
    }

    /// @notice Returns the markup for the message execution, in BWAD math.
    function markup(GasData data) internal pure returns (Number) {
        // Casting to uint16 will truncate the highest bits, which is the behavior we want
        return Number.wrap(uint16(GasData.unwrap(data)));
    }

    // ════════════════════════════════════════════════ CHAIN DATA ═════════════════════════════════════════════════════

    /// @notice Returns an encoded ChainGas struct with the given fields.
    /// @param gasData_ Chain's gas data
    /// @param domain_  Chain's domain
    function encodeChainGas(GasData gasData_, uint32 domain_) internal pure returns (ChainGas) {
        // GasData type wraps uint96, so could safely be casted to uint128
        return ChainGas.wrap(uint128(GasData.unwrap(gasData_)) << SHIFT_GAS_DATA | uint128(domain_));
    }

    /// @notice Wraps padded uint256 value into ChainGas struct.
    function wrapChainGas(uint256 paddedChainGas) internal pure returns (ChainGas) {
        // Casting to uint128 will truncate the highest bits, which is the behavior we want
        return ChainGas.wrap(uint128(paddedChainGas));
    }

    /// @notice Returns the chain's gas data.
    function gasData(ChainGas data) internal pure returns (GasData) {
        // Casting to uint96 will truncate the highest bits, which is the behavior we want
        return GasData.wrap(uint96(ChainGas.unwrap(data) >> SHIFT_GAS_DATA));
    }

    /// @notice Returns the chain's domain.
    function domain(ChainGas data) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(ChainGas.unwrap(data));
    }

    /// @notice Returns the hash for the list of ChainGas structs.
    function snapGasHash(ChainGas[] memory snapGas) internal pure returns (bytes32 snapGasHash_) {
        // Use assembly to calculate the hash of the array without copying it
        // ChainGas takes a single word of storage, thus ChainGas[] is stored in the following way:
        // 0x00: length of the array, in words
        // 0x20: first ChainGas struct
        // 0x40: second ChainGas struct
        // And so on...
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Find the location where the array data starts, we add 0x20 to skip the length field
            let loc := add(snapGas, 0x20)
            // Load the length of the array (in words).
            // Shifting left 5 bits is equivalent to multiplying by 32: this converts from words to bytes.
            let len := shl(5, mload(snapGas))
            // Calculate the hash of the array
            snapGasHash_ := keccak256(loc, len)
        }
    }
}

// Here we define common enums and structures to enable their easier reusing later.

// ═══════════════════════════════ AGENT STATUS ════════════════════════════════

/// @dev Potential statuses for the off-chain bonded agent:
/// - Unknown: never provided a bond => signature not valid
/// - Active: has a bond in BondingManager => signature valid
/// - Unstaking: has a bond in BondingManager, initiated the unstaking => signature not valid
/// - Resting: used to have a bond in BondingManager, successfully unstaked => signature not valid
/// - Fraudulent: proven to commit fraud, value in Merkle Tree not updated => signature not valid
/// - Slashed: proven to commit fraud, value in Merkle Tree was updated => signature not valid
/// Unstaked agent could later be added back to THE SAME domain by staking a bond again.
/// Honest agent: Unknown -> Active -> unstaking -> Resting -> Active ...
/// Malicious agent: Unknown -> Active -> Fraudulent -> Slashed
/// Malicious agent: Unknown -> Active -> Unstaking -> Fraudulent -> Slashed
enum AgentFlag {
    Unknown,
    Active,
    Unstaking,
    Resting,
    Fraudulent,
    Slashed
}

/// @notice Struct for storing an agent in the BondingManager contract.
struct AgentStatus {
    AgentFlag flag;
    uint32 domain;
    uint32 index;
}
// 184 bits available for tight packing

using StructureUtils for AgentStatus global;

/// @notice Potential statuses of an agent in terms of being in dispute
/// - None: agent is not in dispute
/// - Pending: agent is in unresolved dispute
/// - Slashed: agent was in dispute that lead to agent being slashed
/// Note: agent who won the dispute has their status reset to None
enum DisputeFlag {
    None,
    Pending,
    Slashed
}

/// @notice Struct for storing dispute status of an agent.
/// @param flag         Dispute flag: None/Pending/Slashed.
/// @param openedAt     Timestamp when the latest agent dispute was opened (zero if not opened).
/// @param resolvedAt   Timestamp when the latest agent dispute was resolved (zero if not resolved).
struct DisputeStatus {
    DisputeFlag flag;
    uint40 openedAt;
    uint40 resolvedAt;
}

// ════════════════════════════════ DESTINATION ════════════════════════════════

/// @notice Struct representing the status of Destination contract.
/// @param snapRootTime     Timestamp when latest snapshot root was accepted
/// @param agentRootTime    Timestamp when latest agent root was accepted
/// @param notaryIndex      Index of Notary who signed the latest agent root
struct DestinationStatus {
    uint40 snapRootTime;
    uint40 agentRootTime;
    uint32 notaryIndex;
}

// ═══════════════════════════════ EXECUTION HUB ═══════════════════════════════

/// @notice Potential statuses of the message in Execution Hub.
/// - None: there hasn't been a valid attempt to execute the message yet
/// - Failed: there was a valid attempt to execute the message, but recipient reverted
/// - Success: there was a valid attempt to execute the message, and recipient did not revert
/// Note: message can be executed until its status is Success
enum MessageStatus {
    None,
    Failed,
    Success
}

library StructureUtils {
    /// @notice Checks that Agent is Active
    function verifyActive(AgentStatus memory status) internal pure {
        if (status.flag != AgentFlag.Active) {
            revert AgentNotActive();
        }
    }

    /// @notice Checks that Agent is Unstaking
    function verifyUnstaking(AgentStatus memory status) internal pure {
        if (status.flag != AgentFlag.Unstaking) {
            revert AgentNotUnstaking();
        }
    }

    /// @notice Checks that Agent is Active or Unstaking
    function verifyActiveUnstaking(AgentStatus memory status) internal pure {
        if (status.flag != AgentFlag.Active && status.flag != AgentFlag.Unstaking) {
            revert AgentNotActiveNorUnstaking();
        }
    }

    /// @notice Checks that Agent is Fraudulent
    function verifyFraudulent(AgentStatus memory status) internal pure {
        if (status.flag != AgentFlag.Fraudulent) {
            revert AgentNotFraudulent();
        }
    }

    /// @notice Checks that Agent is not Unknown
    function verifyKnown(AgentStatus memory status) internal pure {
        if (status.flag == AgentFlag.Unknown) {
            revert AgentUnknown();
        }
    }
}

// OpenZeppelin Contracts (last updated v4.8.0) (utils/math/SafeCast.sol)
// This file was procedurally generated from scripts/generate/templates/SafeCast.js.

/**
 * @dev Wrappers over Solidity's uintXX/intXX casting operators with added overflow
 * checks.
 *
 * Downcasting from uint256/int256 in Solidity does not revert on overflow. This can
 * easily result in undesired exploitation or bugs, since developers usually
 * assume that overflows raise errors. `SafeCast` restores this intuition by
 * reverting the transaction when such an operation overflows.
 *
 * Using this library instead of the unchecked operations eliminates an entire
 * class of bugs, so it's recommended to use it always.
 *
 * Can be combined with {SafeMath} and {SignedSafeMath} to extend it to smaller types, by performing
 * all math on `uint256` and `int256` and then downcasting.
 */
library SafeCast {
    /**
     * @dev Returns the downcasted uint248 from uint256, reverting on
     * overflow (when the input is greater than largest uint248).
     *
     * Counterpart to Solidity's `uint248` operator.
     *
     * Requirements:
     *
     * - input must fit into 248 bits
     *
     * _Available since v4.7._
     */
    function toUint248(uint256 value) internal pure returns (uint248) {
        require(value <= type(uint248).max, "SafeCast: value doesn't fit in 248 bits");
        return uint248(value);
    }

    /**
     * @dev Returns the downcasted uint240 from uint256, reverting on
     * overflow (when the input is greater than largest uint240).
     *
     * Counterpart to Solidity's `uint240` operator.
     *
     * Requirements:
     *
     * - input must fit into 240 bits
     *
     * _Available since v4.7._
     */
    function toUint240(uint256 value) internal pure returns (uint240) {
        require(value <= type(uint240).max, "SafeCast: value doesn't fit in 240 bits");
        return uint240(value);
    }

    /**
     * @dev Returns the downcasted uint232 from uint256, reverting on
     * overflow (when the input is greater than largest uint232).
     *
     * Counterpart to Solidity's `uint232` operator.
     *
     * Requirements:
     *
     * - input must fit into 232 bits
     *
     * _Available since v4.7._
     */
    function toUint232(uint256 value) internal pure returns (uint232) {
        require(value <= type(uint232).max, "SafeCast: value doesn't fit in 232 bits");
        return uint232(value);
    }

    /**
     * @dev Returns the downcasted uint224 from uint256, reverting on
     * overflow (when the input is greater than largest uint224).
     *
     * Counterpart to Solidity's `uint224` operator.
     *
     * Requirements:
     *
     * - input must fit into 224 bits
     *
     * _Available since v4.2._
     */
    function toUint224(uint256 value) internal pure returns (uint224) {
        require(value <= type(uint224).max, "SafeCast: value doesn't fit in 224 bits");
        return uint224(value);
    }

    /**
     * @dev Returns the downcasted uint216 from uint256, reverting on
     * overflow (when the input is greater than largest uint216).
     *
     * Counterpart to Solidity's `uint216` operator.
     *
     * Requirements:
     *
     * - input must fit into 216 bits
     *
     * _Available since v4.7._
     */
    function toUint216(uint256 value) internal pure returns (uint216) {
        require(value <= type(uint216).max, "SafeCast: value doesn't fit in 216 bits");
        return uint216(value);
    }

    /**
     * @dev Returns the downcasted uint208 from uint256, reverting on
     * overflow (when the input is greater than largest uint208).
     *
     * Counterpart to Solidity's `uint208` operator.
     *
     * Requirements:
     *
     * - input must fit into 208 bits
     *
     * _Available since v4.7._
     */
    function toUint208(uint256 value) internal pure returns (uint208) {
        require(value <= type(uint208).max, "SafeCast: value doesn't fit in 208 bits");
        return uint208(value);
    }

    /**
     * @dev Returns the downcasted uint200 from uint256, reverting on
     * overflow (when the input is greater than largest uint200).
     *
     * Counterpart to Solidity's `uint200` operator.
     *
     * Requirements:
     *
     * - input must fit into 200 bits
     *
     * _Available since v4.7._
     */
    function toUint200(uint256 value) internal pure returns (uint200) {
        require(value <= type(uint200).max, "SafeCast: value doesn't fit in 200 bits");
        return uint200(value);
    }

    /**
     * @dev Returns the downcasted uint192 from uint256, reverting on
     * overflow (when the input is greater than largest uint192).
     *
     * Counterpart to Solidity's `uint192` operator.
     *
     * Requirements:
     *
     * - input must fit into 192 bits
     *
     * _Available since v4.7._
     */
    function toUint192(uint256 value) internal pure returns (uint192) {
        require(value <= type(uint192).max, "SafeCast: value doesn't fit in 192 bits");
        return uint192(value);
    }

    /**
     * @dev Returns the downcasted uint184 from uint256, reverting on
     * overflow (when the input is greater than largest uint184).
     *
     * Counterpart to Solidity's `uint184` operator.
     *
     * Requirements:
     *
     * - input must fit into 184 bits
     *
     * _Available since v4.7._
     */
    function toUint184(uint256 value) internal pure returns (uint184) {
        require(value <= type(uint184).max, "SafeCast: value doesn't fit in 184 bits");
        return uint184(value);
    }

    /**
     * @dev Returns the downcasted uint176 from uint256, reverting on
     * overflow (when the input is greater than largest uint176).
     *
     * Counterpart to Solidity's `uint176` operator.
     *
     * Requirements:
     *
     * - input must fit into 176 bits
     *
     * _Available since v4.7._
     */
    function toUint176(uint256 value) internal pure returns (uint176) {
        require(value <= type(uint176).max, "SafeCast: value doesn't fit in 176 bits");
        return uint176(value);
    }

    /**
     * @dev Returns the downcasted uint168 from uint256, reverting on
     * overflow (when the input is greater than largest uint168).
     *
     * Counterpart to Solidity's `uint168` operator.
     *
     * Requirements:
     *
     * - input must fit into 168 bits
     *
     * _Available since v4.7._
     */
    function toUint168(uint256 value) internal pure returns (uint168) {
        require(value <= type(uint168).max, "SafeCast: value doesn't fit in 168 bits");
        return uint168(value);
    }

    /**
     * @dev Returns the downcasted uint160 from uint256, reverting on
     * overflow (when the input is greater than largest uint160).
     *
     * Counterpart to Solidity's `uint160` operator.
     *
     * Requirements:
     *
     * - input must fit into 160 bits
     *
     * _Available since v4.7._
     */
    function toUint160(uint256 value) internal pure returns (uint160) {
        require(value <= type(uint160).max, "SafeCast: value doesn't fit in 160 bits");
        return uint160(value);
    }

    /**
     * @dev Returns the downcasted uint152 from uint256, reverting on
     * overflow (when the input is greater than largest uint152).
     *
     * Counterpart to Solidity's `uint152` operator.
     *
     * Requirements:
     *
     * - input must fit into 152 bits
     *
     * _Available since v4.7._
     */
    function toUint152(uint256 value) internal pure returns (uint152) {
        require(value <= type(uint152).max, "SafeCast: value doesn't fit in 152 bits");
        return uint152(value);
    }

    /**
     * @dev Returns the downcasted uint144 from uint256, reverting on
     * overflow (when the input is greater than largest uint144).
     *
     * Counterpart to Solidity's `uint144` operator.
     *
     * Requirements:
     *
     * - input must fit into 144 bits
     *
     * _Available since v4.7._
     */
    function toUint144(uint256 value) internal pure returns (uint144) {
        require(value <= type(uint144).max, "SafeCast: value doesn't fit in 144 bits");
        return uint144(value);
    }

    /**
     * @dev Returns the downcasted uint136 from uint256, reverting on
     * overflow (when the input is greater than largest uint136).
     *
     * Counterpart to Solidity's `uint136` operator.
     *
     * Requirements:
     *
     * - input must fit into 136 bits
     *
     * _Available since v4.7._
     */
    function toUint136(uint256 value) internal pure returns (uint136) {
        require(value <= type(uint136).max, "SafeCast: value doesn't fit in 136 bits");
        return uint136(value);
    }

    /**
     * @dev Returns the downcasted uint128 from uint256, reverting on
     * overflow (when the input is greater than largest uint128).
     *
     * Counterpart to Solidity's `uint128` operator.
     *
     * Requirements:
     *
     * - input must fit into 128 bits
     *
     * _Available since v2.5._
     */
    function toUint128(uint256 value) internal pure returns (uint128) {
        require(value <= type(uint128).max, "SafeCast: value doesn't fit in 128 bits");
        return uint128(value);
    }

    /**
     * @dev Returns the downcasted uint120 from uint256, reverting on
     * overflow (when the input is greater than largest uint120).
     *
     * Counterpart to Solidity's `uint120` operator.
     *
     * Requirements:
     *
     * - input must fit into 120 bits
     *
     * _Available since v4.7._
     */
    function toUint120(uint256 value) internal pure returns (uint120) {
        require(value <= type(uint120).max, "SafeCast: value doesn't fit in 120 bits");
        return uint120(value);
    }

    /**
     * @dev Returns the downcasted uint112 from uint256, reverting on
     * overflow (when the input is greater than largest uint112).
     *
     * Counterpart to Solidity's `uint112` operator.
     *
     * Requirements:
     *
     * - input must fit into 112 bits
     *
     * _Available since v4.7._
     */
    function toUint112(uint256 value) internal pure returns (uint112) {
        require(value <= type(uint112).max, "SafeCast: value doesn't fit in 112 bits");
        return uint112(value);
    }

    /**
     * @dev Returns the downcasted uint104 from uint256, reverting on
     * overflow (when the input is greater than largest uint104).
     *
     * Counterpart to Solidity's `uint104` operator.
     *
     * Requirements:
     *
     * - input must fit into 104 bits
     *
     * _Available since v4.7._
     */
    function toUint104(uint256 value) internal pure returns (uint104) {
        require(value <= type(uint104).max, "SafeCast: value doesn't fit in 104 bits");
        return uint104(value);
    }

    /**
     * @dev Returns the downcasted uint96 from uint256, reverting on
     * overflow (when the input is greater than largest uint96).
     *
     * Counterpart to Solidity's `uint96` operator.
     *
     * Requirements:
     *
     * - input must fit into 96 bits
     *
     * _Available since v4.2._
     */
    function toUint96(uint256 value) internal pure returns (uint96) {
        require(value <= type(uint96).max, "SafeCast: value doesn't fit in 96 bits");
        return uint96(value);
    }

    /**
     * @dev Returns the downcasted uint88 from uint256, reverting on
     * overflow (when the input is greater than largest uint88).
     *
     * Counterpart to Solidity's `uint88` operator.
     *
     * Requirements:
     *
     * - input must fit into 88 bits
     *
     * _Available since v4.7._
     */
    function toUint88(uint256 value) internal pure returns (uint88) {
        require(value <= type(uint88).max, "SafeCast: value doesn't fit in 88 bits");
        return uint88(value);
    }

    /**
     * @dev Returns the downcasted uint80 from uint256, reverting on
     * overflow (when the input is greater than largest uint80).
     *
     * Counterpart to Solidity's `uint80` operator.
     *
     * Requirements:
     *
     * - input must fit into 80 bits
     *
     * _Available since v4.7._
     */
    function toUint80(uint256 value) internal pure returns (uint80) {
        require(value <= type(uint80).max, "SafeCast: value doesn't fit in 80 bits");
        return uint80(value);
    }

    /**
     * @dev Returns the downcasted uint72 from uint256, reverting on
     * overflow (when the input is greater than largest uint72).
     *
     * Counterpart to Solidity's `uint72` operator.
     *
     * Requirements:
     *
     * - input must fit into 72 bits
     *
     * _Available since v4.7._
     */
    function toUint72(uint256 value) internal pure returns (uint72) {
        require(value <= type(uint72).max, "SafeCast: value doesn't fit in 72 bits");
        return uint72(value);
    }

    /**
     * @dev Returns the downcasted uint64 from uint256, reverting on
     * overflow (when the input is greater than largest uint64).
     *
     * Counterpart to Solidity's `uint64` operator.
     *
     * Requirements:
     *
     * - input must fit into 64 bits
     *
     * _Available since v2.5._
     */
    function toUint64(uint256 value) internal pure returns (uint64) {
        require(value <= type(uint64).max, "SafeCast: value doesn't fit in 64 bits");
        return uint64(value);
    }

    /**
     * @dev Returns the downcasted uint56 from uint256, reverting on
     * overflow (when the input is greater than largest uint56).
     *
     * Counterpart to Solidity's `uint56` operator.
     *
     * Requirements:
     *
     * - input must fit into 56 bits
     *
     * _Available since v4.7._
     */
    function toUint56(uint256 value) internal pure returns (uint56) {
        require(value <= type(uint56).max, "SafeCast: value doesn't fit in 56 bits");
        return uint56(value);
    }

    /**
     * @dev Returns the downcasted uint48 from uint256, reverting on
     * overflow (when the input is greater than largest uint48).
     *
     * Counterpart to Solidity's `uint48` operator.
     *
     * Requirements:
     *
     * - input must fit into 48 bits
     *
     * _Available since v4.7._
     */
    function toUint48(uint256 value) internal pure returns (uint48) {
        require(value <= type(uint48).max, "SafeCast: value doesn't fit in 48 bits");
        return uint48(value);
    }

    /**
     * @dev Returns the downcasted uint40 from uint256, reverting on
     * overflow (when the input is greater than largest uint40).
     *
     * Counterpart to Solidity's `uint40` operator.
     *
     * Requirements:
     *
     * - input must fit into 40 bits
     *
     * _Available since v4.7._
     */
    function toUint40(uint256 value) internal pure returns (uint40) {
        require(value <= type(uint40).max, "SafeCast: value doesn't fit in 40 bits");
        return uint40(value);
    }

    /**
     * @dev Returns the downcasted uint32 from uint256, reverting on
     * overflow (when the input is greater than largest uint32).
     *
     * Counterpart to Solidity's `uint32` operator.
     *
     * Requirements:
     *
     * - input must fit into 32 bits
     *
     * _Available since v2.5._
     */
    function toUint32(uint256 value) internal pure returns (uint32) {
        require(value <= type(uint32).max, "SafeCast: value doesn't fit in 32 bits");
        return uint32(value);
    }

    /**
     * @dev Returns the downcasted uint24 from uint256, reverting on
     * overflow (when the input is greater than largest uint24).
     *
     * Counterpart to Solidity's `uint24` operator.
     *
     * Requirements:
     *
     * - input must fit into 24 bits
     *
     * _Available since v4.7._
     */
    function toUint24(uint256 value) internal pure returns (uint24) {
        require(value <= type(uint24).max, "SafeCast: value doesn't fit in 24 bits");
        return uint24(value);
    }

    /**
     * @dev Returns the downcasted uint16 from uint256, reverting on
     * overflow (when the input is greater than largest uint16).
     *
     * Counterpart to Solidity's `uint16` operator.
     *
     * Requirements:
     *
     * - input must fit into 16 bits
     *
     * _Available since v2.5._
     */
    function toUint16(uint256 value) internal pure returns (uint16) {
        require(value <= type(uint16).max, "SafeCast: value doesn't fit in 16 bits");
        return uint16(value);
    }

    /**
     * @dev Returns the downcasted uint8 from uint256, reverting on
     * overflow (when the input is greater than largest uint8).
     *
     * Counterpart to Solidity's `uint8` operator.
     *
     * Requirements:
     *
     * - input must fit into 8 bits
     *
     * _Available since v2.5._
     */
    function toUint8(uint256 value) internal pure returns (uint8) {
        require(value <= type(uint8).max, "SafeCast: value doesn't fit in 8 bits");
        return uint8(value);
    }

    /**
     * @dev Converts a signed int256 into an unsigned uint256.
     *
     * Requirements:
     *
     * - input must be greater than or equal to 0.
     *
     * _Available since v3.0._
     */
    function toUint256(int256 value) internal pure returns (uint256) {
        require(value >= 0, "SafeCast: value must be positive");
        return uint256(value);
    }

    /**
     * @dev Returns the downcasted int248 from int256, reverting on
     * overflow (when the input is less than smallest int248 or
     * greater than largest int248).
     *
     * Counterpart to Solidity's `int248` operator.
     *
     * Requirements:
     *
     * - input must fit into 248 bits
     *
     * _Available since v4.7._
     */
    function toInt248(int256 value) internal pure returns (int248 downcasted) {
        downcasted = int248(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 248 bits");
    }

    /**
     * @dev Returns the downcasted int240 from int256, reverting on
     * overflow (when the input is less than smallest int240 or
     * greater than largest int240).
     *
     * Counterpart to Solidity's `int240` operator.
     *
     * Requirements:
     *
     * - input must fit into 240 bits
     *
     * _Available since v4.7._
     */
    function toInt240(int256 value) internal pure returns (int240 downcasted) {
        downcasted = int240(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 240 bits");
    }

    /**
     * @dev Returns the downcasted int232 from int256, reverting on
     * overflow (when the input is less than smallest int232 or
     * greater than largest int232).
     *
     * Counterpart to Solidity's `int232` operator.
     *
     * Requirements:
     *
     * - input must fit into 232 bits
     *
     * _Available since v4.7._
     */
    function toInt232(int256 value) internal pure returns (int232 downcasted) {
        downcasted = int232(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 232 bits");
    }

    /**
     * @dev Returns the downcasted int224 from int256, reverting on
     * overflow (when the input is less than smallest int224 or
     * greater than largest int224).
     *
     * Counterpart to Solidity's `int224` operator.
     *
     * Requirements:
     *
     * - input must fit into 224 bits
     *
     * _Available since v4.7._
     */
    function toInt224(int256 value) internal pure returns (int224 downcasted) {
        downcasted = int224(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 224 bits");
    }

    /**
     * @dev Returns the downcasted int216 from int256, reverting on
     * overflow (when the input is less than smallest int216 or
     * greater than largest int216).
     *
     * Counterpart to Solidity's `int216` operator.
     *
     * Requirements:
     *
     * - input must fit into 216 bits
     *
     * _Available since v4.7._
     */
    function toInt216(int256 value) internal pure returns (int216 downcasted) {
        downcasted = int216(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 216 bits");
    }

    /**
     * @dev Returns the downcasted int208 from int256, reverting on
     * overflow (when the input is less than smallest int208 or
     * greater than largest int208).
     *
     * Counterpart to Solidity's `int208` operator.
     *
     * Requirements:
     *
     * - input must fit into 208 bits
     *
     * _Available since v4.7._
     */
    function toInt208(int256 value) internal pure returns (int208 downcasted) {
        downcasted = int208(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 208 bits");
    }

    /**
     * @dev Returns the downcasted int200 from int256, reverting on
     * overflow (when the input is less than smallest int200 or
     * greater than largest int200).
     *
     * Counterpart to Solidity's `int200` operator.
     *
     * Requirements:
     *
     * - input must fit into 200 bits
     *
     * _Available since v4.7._
     */
    function toInt200(int256 value) internal pure returns (int200 downcasted) {
        downcasted = int200(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 200 bits");
    }

    /**
     * @dev Returns the downcasted int192 from int256, reverting on
     * overflow (when the input is less than smallest int192 or
     * greater than largest int192).
     *
     * Counterpart to Solidity's `int192` operator.
     *
     * Requirements:
     *
     * - input must fit into 192 bits
     *
     * _Available since v4.7._
     */
    function toInt192(int256 value) internal pure returns (int192 downcasted) {
        downcasted = int192(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 192 bits");
    }

    /**
     * @dev Returns the downcasted int184 from int256, reverting on
     * overflow (when the input is less than smallest int184 or
     * greater than largest int184).
     *
     * Counterpart to Solidity's `int184` operator.
     *
     * Requirements:
     *
     * - input must fit into 184 bits
     *
     * _Available since v4.7._
     */
    function toInt184(int256 value) internal pure returns (int184 downcasted) {
        downcasted = int184(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 184 bits");
    }

    /**
     * @dev Returns the downcasted int176 from int256, reverting on
     * overflow (when the input is less than smallest int176 or
     * greater than largest int176).
     *
     * Counterpart to Solidity's `int176` operator.
     *
     * Requirements:
     *
     * - input must fit into 176 bits
     *
     * _Available since v4.7._
     */
    function toInt176(int256 value) internal pure returns (int176 downcasted) {
        downcasted = int176(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 176 bits");
    }

    /**
     * @dev Returns the downcasted int168 from int256, reverting on
     * overflow (when the input is less than smallest int168 or
     * greater than largest int168).
     *
     * Counterpart to Solidity's `int168` operator.
     *
     * Requirements:
     *
     * - input must fit into 168 bits
     *
     * _Available since v4.7._
     */
    function toInt168(int256 value) internal pure returns (int168 downcasted) {
        downcasted = int168(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 168 bits");
    }

    /**
     * @dev Returns the downcasted int160 from int256, reverting on
     * overflow (when the input is less than smallest int160 or
     * greater than largest int160).
     *
     * Counterpart to Solidity's `int160` operator.
     *
     * Requirements:
     *
     * - input must fit into 160 bits
     *
     * _Available since v4.7._
     */
    function toInt160(int256 value) internal pure returns (int160 downcasted) {
        downcasted = int160(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 160 bits");
    }

    /**
     * @dev Returns the downcasted int152 from int256, reverting on
     * overflow (when the input is less than smallest int152 or
     * greater than largest int152).
     *
     * Counterpart to Solidity's `int152` operator.
     *
     * Requirements:
     *
     * - input must fit into 152 bits
     *
     * _Available since v4.7._
     */
    function toInt152(int256 value) internal pure returns (int152 downcasted) {
        downcasted = int152(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 152 bits");
    }

    /**
     * @dev Returns the downcasted int144 from int256, reverting on
     * overflow (when the input is less than smallest int144 or
     * greater than largest int144).
     *
     * Counterpart to Solidity's `int144` operator.
     *
     * Requirements:
     *
     * - input must fit into 144 bits
     *
     * _Available since v4.7._
     */
    function toInt144(int256 value) internal pure returns (int144 downcasted) {
        downcasted = int144(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 144 bits");
    }

    /**
     * @dev Returns the downcasted int136 from int256, reverting on
     * overflow (when the input is less than smallest int136 or
     * greater than largest int136).
     *
     * Counterpart to Solidity's `int136` operator.
     *
     * Requirements:
     *
     * - input must fit into 136 bits
     *
     * _Available since v4.7._
     */
    function toInt136(int256 value) internal pure returns (int136 downcasted) {
        downcasted = int136(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 136 bits");
    }

    /**
     * @dev Returns the downcasted int128 from int256, reverting on
     * overflow (when the input is less than smallest int128 or
     * greater than largest int128).
     *
     * Counterpart to Solidity's `int128` operator.
     *
     * Requirements:
     *
     * - input must fit into 128 bits
     *
     * _Available since v3.1._
     */
    function toInt128(int256 value) internal pure returns (int128 downcasted) {
        downcasted = int128(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 128 bits");
    }

    /**
     * @dev Returns the downcasted int120 from int256, reverting on
     * overflow (when the input is less than smallest int120 or
     * greater than largest int120).
     *
     * Counterpart to Solidity's `int120` operator.
     *
     * Requirements:
     *
     * - input must fit into 120 bits
     *
     * _Available since v4.7._
     */
    function toInt120(int256 value) internal pure returns (int120 downcasted) {
        downcasted = int120(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 120 bits");
    }

    /**
     * @dev Returns the downcasted int112 from int256, reverting on
     * overflow (when the input is less than smallest int112 or
     * greater than largest int112).
     *
     * Counterpart to Solidity's `int112` operator.
     *
     * Requirements:
     *
     * - input must fit into 112 bits
     *
     * _Available since v4.7._
     */
    function toInt112(int256 value) internal pure returns (int112 downcasted) {
        downcasted = int112(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 112 bits");
    }

    /**
     * @dev Returns the downcasted int104 from int256, reverting on
     * overflow (when the input is less than smallest int104 or
     * greater than largest int104).
     *
     * Counterpart to Solidity's `int104` operator.
     *
     * Requirements:
     *
     * - input must fit into 104 bits
     *
     * _Available since v4.7._
     */
    function toInt104(int256 value) internal pure returns (int104 downcasted) {
        downcasted = int104(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 104 bits");
    }

    /**
     * @dev Returns the downcasted int96 from int256, reverting on
     * overflow (when the input is less than smallest int96 or
     * greater than largest int96).
     *
     * Counterpart to Solidity's `int96` operator.
     *
     * Requirements:
     *
     * - input must fit into 96 bits
     *
     * _Available since v4.7._
     */
    function toInt96(int256 value) internal pure returns (int96 downcasted) {
        downcasted = int96(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 96 bits");
    }

    /**
     * @dev Returns the downcasted int88 from int256, reverting on
     * overflow (when the input is less than smallest int88 or
     * greater than largest int88).
     *
     * Counterpart to Solidity's `int88` operator.
     *
     * Requirements:
     *
     * - input must fit into 88 bits
     *
     * _Available since v4.7._
     */
    function toInt88(int256 value) internal pure returns (int88 downcasted) {
        downcasted = int88(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 88 bits");
    }

    /**
     * @dev Returns the downcasted int80 from int256, reverting on
     * overflow (when the input is less than smallest int80 or
     * greater than largest int80).
     *
     * Counterpart to Solidity's `int80` operator.
     *
     * Requirements:
     *
     * - input must fit into 80 bits
     *
     * _Available since v4.7._
     */
    function toInt80(int256 value) internal pure returns (int80 downcasted) {
        downcasted = int80(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 80 bits");
    }

    /**
     * @dev Returns the downcasted int72 from int256, reverting on
     * overflow (when the input is less than smallest int72 or
     * greater than largest int72).
     *
     * Counterpart to Solidity's `int72` operator.
     *
     * Requirements:
     *
     * - input must fit into 72 bits
     *
     * _Available since v4.7._
     */
    function toInt72(int256 value) internal pure returns (int72 downcasted) {
        downcasted = int72(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 72 bits");
    }

    /**
     * @dev Returns the downcasted int64 from int256, reverting on
     * overflow (when the input is less than smallest int64 or
     * greater than largest int64).
     *
     * Counterpart to Solidity's `int64` operator.
     *
     * Requirements:
     *
     * - input must fit into 64 bits
     *
     * _Available since v3.1._
     */
    function toInt64(int256 value) internal pure returns (int64 downcasted) {
        downcasted = int64(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 64 bits");
    }

    /**
     * @dev Returns the downcasted int56 from int256, reverting on
     * overflow (when the input is less than smallest int56 or
     * greater than largest int56).
     *
     * Counterpart to Solidity's `int56` operator.
     *
     * Requirements:
     *
     * - input must fit into 56 bits
     *
     * _Available since v4.7._
     */
    function toInt56(int256 value) internal pure returns (int56 downcasted) {
        downcasted = int56(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 56 bits");
    }

    /**
     * @dev Returns the downcasted int48 from int256, reverting on
     * overflow (when the input is less than smallest int48 or
     * greater than largest int48).
     *
     * Counterpart to Solidity's `int48` operator.
     *
     * Requirements:
     *
     * - input must fit into 48 bits
     *
     * _Available since v4.7._
     */
    function toInt48(int256 value) internal pure returns (int48 downcasted) {
        downcasted = int48(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 48 bits");
    }

    /**
     * @dev Returns the downcasted int40 from int256, reverting on
     * overflow (when the input is less than smallest int40 or
     * greater than largest int40).
     *
     * Counterpart to Solidity's `int40` operator.
     *
     * Requirements:
     *
     * - input must fit into 40 bits
     *
     * _Available since v4.7._
     */
    function toInt40(int256 value) internal pure returns (int40 downcasted) {
        downcasted = int40(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 40 bits");
    }

    /**
     * @dev Returns the downcasted int32 from int256, reverting on
     * overflow (when the input is less than smallest int32 or
     * greater than largest int32).
     *
     * Counterpart to Solidity's `int32` operator.
     *
     * Requirements:
     *
     * - input must fit into 32 bits
     *
     * _Available since v3.1._
     */
    function toInt32(int256 value) internal pure returns (int32 downcasted) {
        downcasted = int32(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 32 bits");
    }

    /**
     * @dev Returns the downcasted int24 from int256, reverting on
     * overflow (when the input is less than smallest int24 or
     * greater than largest int24).
     *
     * Counterpart to Solidity's `int24` operator.
     *
     * Requirements:
     *
     * - input must fit into 24 bits
     *
     * _Available since v4.7._
     */
    function toInt24(int256 value) internal pure returns (int24 downcasted) {
        downcasted = int24(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 24 bits");
    }

    /**
     * @dev Returns the downcasted int16 from int256, reverting on
     * overflow (when the input is less than smallest int16 or
     * greater than largest int16).
     *
     * Counterpart to Solidity's `int16` operator.
     *
     * Requirements:
     *
     * - input must fit into 16 bits
     *
     * _Available since v3.1._
     */
    function toInt16(int256 value) internal pure returns (int16 downcasted) {
        downcasted = int16(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 16 bits");
    }

    /**
     * @dev Returns the downcasted int8 from int256, reverting on
     * overflow (when the input is less than smallest int8 or
     * greater than largest int8).
     *
     * Counterpart to Solidity's `int8` operator.
     *
     * Requirements:
     *
     * - input must fit into 8 bits
     *
     * _Available since v3.1._
     */
    function toInt8(int256 value) internal pure returns (int8 downcasted) {
        downcasted = int8(value);
        require(downcasted == value, "SafeCast: value doesn't fit in 8 bits");
    }

    /**
     * @dev Converts an unsigned uint256 into a signed int256.
     *
     * Requirements:
     *
     * - input must be less than or equal to maxInt256.
     *
     * _Available since v3.0._
     */
    function toInt256(uint256 value) internal pure returns (int256) {
        // Note: Unsafe cast below is okay because `type(int256).max` is guaranteed to be positive
        require(value <= uint256(type(int256).max), "SafeCast: value doesn't fit in an int256");
        return int256(value);
    }
}

/// @notice Library for accessing chain context variables as tightly packed integers.
/// Messaging contracts should rely on this library for accessing chain context variables
/// instead of doing the casting themselves.
library ChainContext {
    using SafeCast for uint256;

    /// @notice Returns the current block number as uint40.
    /// @dev Reverts if block number is greater than 40 bits, which is not supposed to happen
    /// until the block.timestamp overflows (assuming block time is at least 1 second).
    function blockNumber() internal view returns (uint40) {
        return block.number.toUint40();
    }

    /// @notice Returns the current block timestamp as uint40.
    /// @dev Reverts if block timestamp is greater than 40 bits, which is
    /// supposed to happen approximately in year 36835.
    function blockTimestamp() internal view returns (uint40) {
        return block.timestamp.toUint40();
    }

    /// @notice Returns the chain id as uint32.
    /// @dev Reverts if chain id is greater than 32 bits, which is not
    /// supposed to happen in production.
    function chainId() internal view returns (uint32) {
        return block.chainid.toUint32();
    }
}

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════

interface IAgentManager {
    /**
     * @notice Allows Inbox to open a Dispute between a Guard and a Notary, if they are both not in Dispute already.
     * > Will revert if any of these is true:
     * > - Caller is not Inbox.
     * > - Guard or Notary is already in Dispute.
     * @param guardIndex    Index of the Guard in the Agent Merkle Tree
     * @param notaryIndex   Index of the Notary in the Agent Merkle Tree
     */
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external;

    /**
     * @notice Allows Inbox to slash an agent, if their fraud was proven.
     * > Will revert if any of these is true:
     * > - Caller is not Inbox.
     * > - Domain doesn't match the saved agent domain.
     * @param domain    Domain where the Agent is active
     * @param agent     Address of the Agent
     * @param prover    Address that initially provided fraud proof
     */
    function slashAgent(uint32 domain, address agent, address prover) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the latest known root of the Agent Merkle Tree.
     */
    function agentRoot() external view returns (bytes32);

    /**
     * @notice Returns (flag, domain, index) for a given agent. See Structures.sol for details.
     * @dev Will return AgentFlag.Fraudulent for agents that have been proven to commit fraud,
     * but their status is not updated to Slashed yet.
     * @param agent     Agent address
     * @return          Status for the given agent: (flag, domain, index).
     */
    function agentStatus(address agent) external view returns (AgentStatus memory);

    /**
     * @notice Returns agent address and their current status for a given agent index.
     * @dev Will return empty values if agent with given index doesn't exist.
     * @param index     Agent index in the Agent Merkle Tree
     * @return agent    Agent address
     * @return status   Status for the given agent: (flag, domain, index)
     */
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status);

    /**
     * @notice Returns the number of opened Disputes.
     * @dev This includes the Disputes that have been resolved already.
     */
    function getDisputesAmount() external view returns (uint256);

    /**
     * @notice Returns information about the dispute with the given index.
     * @dev Will revert if dispute with given index hasn't been opened yet.
     * @param index             Dispute index
     * @return guard            Address of the Guard in the Dispute
     * @return notary           Address of the Notary in the Dispute
     * @return slashedAgent     Address of the Agent who was slashed when Dispute was resolved
     * @return fraudProver      Address who provided fraud proof to resolve the Dispute
     * @return reportPayload    Raw payload with report data that led to the Dispute
     * @return reportSignature  Guard signature for the report payload
     */
    function getDispute(uint256 index)
        external
        view
        returns (
            address guard,
            address notary,
            address slashedAgent,
            address fraudProver,
            bytes memory reportPayload,
            bytes memory reportSignature
        );

    /**
     * @notice Returns the current Dispute status of a given agent. See Structures.sol for details.
     * @dev Every returned value will be set to zero if agent was not slashed and is not in Dispute.
     * `rival` and `disputePtr` will be set to zero if the agent was slashed without being in Dispute.
     * @param agent         Agent address
     * @return flag         Flag describing the current Dispute status for the agent: None/Pending/Slashed
     * @return rival        Address of the rival agent in the Dispute
     * @return fraudProver  Address who provided fraud proof to resolve the Dispute
     * @return disputePtr   Index of the opened Dispute PLUS ONE. Zero if agent is not in Dispute.
     */
    function disputeStatus(address agent)
        external
        view
        returns (DisputeFlag flag, address rival, address fraudProver, uint256 disputePtr);
}

interface IAgentSecured {
    /**
     * @notice Local AgentManager should call this function to indicate that a dispute
     * between a Guard and a Notary has been opened.
     * @param guardIndex    Index of the Guard in the Agent Merkle Tree
     * @param notaryIndex   Index of the Notary in the Agent Merkle Tree
     */
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external;

    /**
     * @notice Local AgentManager should call this function to indicate that a dispute
     * has been resolved due to one of the agents being slashed.
     * > `rivalIndex` will be ZERO, if the slashed agent was not in the Dispute.
     * @param slashedIndex  Index of the slashed agent in the Agent Merkle Tree
     * @param rivalIndex    Index of the their Dispute Rival in the Agent Merkle Tree
     */
    function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the address of the local AgentManager contract, which is treated as
     * the "source of truth" for agent statuses.
     */
    function agentManager() external view returns (address);

    /**
     * @notice Returns the address of the local Inbox contract, which is treated as
     * the "source of truth" for agent-signed statements.
     * @dev Inbox passes verified agent statements to `IAgentSecured` contract.
     */
    function inbox() external view returns (address);

    /**
     * @notice Returns (flag, domain, index) for a given agent. See Structures.sol for details.
     * @dev Will return AgentFlag.Fraudulent for agents that have been proven to commit fraud,
     * but their status is not updated to Slashed yet.
     * @param agent     Agent address
     * @return          Status for the given agent: (flag, domain, index).
     */
    function agentStatus(address agent) external view returns (AgentStatus memory);

    /**
     * @notice Returns agent address and their current status for a given agent index.
     * @dev Will return empty values if agent with given index doesn't exist.
     * @param index     Agent index in the Agent Merkle Tree
     * @return agent    Agent address
     * @return status   Status for the given agent: (flag, domain, index)
     */
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status);

    /**
     * @notice Returns (flag, openedAt, resolvedAt) that describes the latest status of
     * the latest dispute for an agent with a given index.
     * @dev Will return empty values if agent with given index doesn't exist.
     * @param agentIndex    Agent index in the Agent Merkle Tree
     * @return              Latest dispute status for the given agent: (flag, openedAt, resolvedAt)
     */
    function latestDisputeStatus(uint32 agentIndex) external view returns (DisputeStatus memory);
}

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════

/// @notice Collection of Multicall utilities. Fork of Multicall3:
/// https://github.com/mds1/multicall/blob/master/src/Multicall3.sol
abstract contract MultiCallable {
    struct Call {
        bool allowFailure;
        bytes callData;
    }

    struct Result {
        bool success;
        bytes returnData;
    }

    /// @notice Aggregates a few calls to this contract into one multicall without modifying `msg.sender`.
    function multicall(Call[] calldata calls) external returns (Result[] memory callResults) {
        uint256 amount = calls.length;
        callResults = new Result[](amount);
        Call calldata call_;
        for (uint256 i = 0; i < amount;) {
            call_ = calls[i];
            Result memory result = callResults[i];
            // We perform a delegate call to ourselves here. Delegate call does not modify `msg.sender`, so
            // this will have the same effect as if `msg.sender` performed all the calls themselves one by one.
            // solhint-disable-next-line avoid-low-level-calls
            (result.success, result.returnData) = address(this).delegatecall(call_.callData);
            // solhint-disable-next-line no-inline-assembly
            assembly {
                // Revert if the call fails and failure is not allowed
                // `allowFailure := calldataload(call_)` and `success := mload(result)`
                if iszero(or(calldataload(call_), mload(result))) {
                    // Revert with `0x4d6a2328` (function selector for `MulticallFailed()`)
                    mstore(0x00, 0x4d6a232800000000000000000000000000000000000000000000000000000000)
                    revert(0x00, 0x04)
                }
            }
            unchecked {
                ++i;
            }
        }
    }
}

/**
 * @title Versioned
 * @notice Version getter for contracts. Doesn't use any storage slots, meaning
 * it will never cause any troubles with the upgradeable contracts. For instance, this contract
 * can be added or removed from the inheritance chain without shifting the storage layout.
 */
abstract contract Versioned {
    /**
     * @notice Struct that is mimicking the storage layout of a string with 32 bytes or less.
     * Length is limited by 32, so the whole string payload takes two memory words:
     * @param length    String length
     * @param data      String characters
     */
    struct _ShortString {
        uint256 length;
        bytes32 data;
    }

    /// @dev Length of the "version string"
    uint256 private immutable _length;
    /// @dev Bytes representation of the "version string".
    /// Strings with length over 32 are not supported!
    bytes32 private immutable _data;

    constructor(string memory version_) {
        _length = bytes(version_).length;
        if (_length > 32) revert IncorrectVersionLength();
        // bytes32 is left-aligned => this will store the byte representation of the string
        // with the trailing zeroes to complete the 32-byte word
        _data = bytes32(bytes(version_));
    }

    function version() external view returns (string memory versionString) {
        // Load the immutable values to form the version string
        _ShortString memory str = _ShortString(_length, _data);
        // The only way to do this cast is doing some dirty assembly
        assembly {
            // solhint-disable-previous-line no-inline-assembly
            versionString := str
        }
    }
}

// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════

// OpenZeppelin Contracts (last updated v4.9.0) (access/Ownable2Step.sol)

// OpenZeppelin Contracts (last updated v4.9.0) (access/Ownable.sol)

// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

// OpenZeppelin Contracts (last updated v4.9.0) (proxy/utils/Initializable.sol)

// OpenZeppelin Contracts (last updated v4.9.0) (utils/Address.sol)

/**
 * @dev Collection of functions related to the address type
 */
library AddressUpgradeable {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     *
     * Furthermore, `isContract` will also return true if the target contract within
     * the same transaction is already scheduled for destruction by `SELFDESTRUCT`,
     * which only has an effect at the end of a transaction.
     * ====
     *
     * [IMPORTANT]
     * ====
     * You shouldn't rely on `isContract` to protect against flash loan attacks!
     *
     * Preventing calls from contracts is highly discouraged. It breaks composability, breaks support for smart wallets
     * like Gnosis Safe, and does not provide security since it can be circumvented by calling from a contract
     * constructor.
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://consensys.net/diligence/blog/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.8.0/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain `call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(address target, bytes memory data, uint256 value) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionDelegateCall(target, data, "Address: low-level delegate call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        (bool success, bytes memory returndata) = target.delegatecall(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verify that a low level call to smart-contract was successful, and revert (either by bubbling
     * the revert reason or using the provided one) in case of unsuccessful call or if target was not a contract.
     *
     * _Available since v4.8._
     */
    function verifyCallResultFromTarget(
        address target,
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        if (success) {
            if (returndata.length == 0) {
                // only check isContract if the call was successful and the return data is empty
                // otherwise we already know that it was a contract
                require(isContract(target), "Address: call to non-contract");
            }
            return returndata;
        } else {
            _revert(returndata, errorMessage);
        }
    }

    /**
     * @dev Tool to verify that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason or using the provided one.
     *
     * _Available since v4.3._
     */
    function verifyCallResult(
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        if (success) {
            return returndata;
        } else {
            _revert(returndata, errorMessage);
        }
    }

    function _revert(bytes memory returndata, string memory errorMessage) private pure {
        // Look for revert reason and bubble it up if present
        if (returndata.length > 0) {
            // The easiest way to bubble the revert reason is using memory via assembly
            /// @solidity memory-safe-assembly
            assembly {
                let returndata_size := mload(returndata)
                revert(add(32, returndata), returndata_size)
            }
        } else {
            revert(errorMessage);
        }
    }
}

/**
 * @dev This is a base contract to aid in writing upgradeable contracts, or any kind of contract that will be deployed
 * behind a proxy. Since proxied contracts do not make use of a constructor, it's common to move constructor logic to an
 * external initializer function, usually called `initialize`. It then becomes necessary to protect this initializer
 * function so it can only be called once. The {initializer} modifier provided by this contract will have this effect.
 *
 * The initialization functions use a version number. Once a version number is used, it is consumed and cannot be
 * reused. This mechanism prevents re-execution of each "step" but allows the creation of new initialization steps in
 * case an upgrade adds a module that needs to be initialized.
 *
 * For example:
 *
 * [.hljs-theme-light.nopadding]
 * ```solidity
 * contract MyToken is ERC20Upgradeable {
 *     function initialize() initializer public {
 *         __ERC20_init("MyToken", "MTK");
 *     }
 * }
 *
 * contract MyTokenV2 is MyToken, ERC20PermitUpgradeable {
 *     function initializeV2() reinitializer(2) public {
 *         __ERC20Permit_init("MyToken");
 *     }
 * }
 * ```
 *
 * TIP: To avoid leaving the proxy in an uninitialized state, the initializer function should be called as early as
 * possible by providing the encoded function call as the `_data` argument to {ERC1967Proxy-constructor}.
 *
 * CAUTION: When used with inheritance, manual care must be taken to not invoke a parent initializer twice, or to ensure
 * that all initializers are idempotent. This is not verified automatically as constructors are by Solidity.
 *
 * [CAUTION]
 * ====
 * Avoid leaving a contract uninitialized.
 *
 * An uninitialized contract can be taken over by an attacker. This applies to both a proxy and its implementation
 * contract, which may impact the proxy. To prevent the implementation contract from being used, you should invoke
 * the {_disableInitializers} function in the constructor to automatically lock it when it is deployed:
 *
 * [.hljs-theme-light.nopadding]
 * ```
 * /// @custom:oz-upgrades-unsafe-allow constructor
 * constructor() {
 *     _disableInitializers();
 * }
 * ```
 * ====
 */
abstract contract Initializable {
    /**
     * @dev Indicates that the contract has been initialized.
     * @custom:oz-retyped-from bool
     */
    uint8 private _initialized;

    /**
     * @dev Indicates that the contract is in the process of being initialized.
     */
    bool private _initializing;

    /**
     * @dev Triggered when the contract has been initialized or reinitialized.
     */
    event Initialized(uint8 version);

    /**
     * @dev A modifier that defines a protected initializer function that can be invoked at most once. In its scope,
     * `onlyInitializing` functions can be used to initialize parent contracts.
     *
     * Similar to `reinitializer(1)`, except that functions marked with `initializer` can be nested in the context of a
     * constructor.
     *
     * Emits an {Initialized} event.
     */
    modifier initializer() {
        bool isTopLevelCall = !_initializing;
        require(
            (isTopLevelCall && _initialized < 1) || (!AddressUpgradeable.isContract(address(this)) && _initialized == 1),
            "Initializable: contract is already initialized"
        );
        _initialized = 1;
        if (isTopLevelCall) {
            _initializing = true;
        }
        _;
        if (isTopLevelCall) {
            _initializing = false;
            emit Initialized(1);
        }
    }

    /**
     * @dev A modifier that defines a protected reinitializer function that can be invoked at most once, and only if the
     * contract hasn't been initialized to a greater version before. In its scope, `onlyInitializing` functions can be
     * used to initialize parent contracts.
     *
     * A reinitializer may be used after the original initialization step. This is essential to configure modules that
     * are added through upgrades and that require initialization.
     *
     * When `version` is 1, this modifier is similar to `initializer`, except that functions marked with `reinitializer`
     * cannot be nested. If one is invoked in the context of another, execution will revert.
     *
     * Note that versions can jump in increments greater than 1; this implies that if multiple reinitializers coexist in
     * a contract, executing them in the right order is up to the developer or operator.
     *
     * WARNING: setting the version to 255 will prevent any future reinitialization.
     *
     * Emits an {Initialized} event.
     */
    modifier reinitializer(uint8 version) {
        require(!_initializing && _initialized < version, "Initializable: contract is already initialized");
        _initialized = version;
        _initializing = true;
        _;
        _initializing = false;
        emit Initialized(version);
    }

    /**
     * @dev Modifier to protect an initialization function so that it can only be invoked by functions with the
     * {initializer} and {reinitializer} modifiers, directly or indirectly.
     */
    modifier onlyInitializing() {
        require(_initializing, "Initializable: contract is not initializing");
        _;
    }

    /**
     * @dev Locks the contract, preventing any future reinitialization. This cannot be part of an initializer call.
     * Calling this in the constructor of a contract will prevent that contract from being initialized or reinitialized
     * to any version. It is recommended to use this to lock implementation contracts that are designed to be called
     * through proxies.
     *
     * Emits an {Initialized} event the first time it is successfully executed.
     */
    function _disableInitializers() internal virtual {
        require(!_initializing, "Initializable: contract is initializing");
        if (_initialized != type(uint8).max) {
            _initialized = type(uint8).max;
            emit Initialized(type(uint8).max);
        }
    }

    /**
     * @dev Returns the highest version that has been initialized. See {reinitializer}.
     */
    function _getInitializedVersion() internal view returns (uint8) {
        return _initialized;
    }

    /**
     * @dev Returns `true` if the contract is currently initializing. See {onlyInitializing}.
     */
    function _isInitializing() internal view returns (bool) {
        return _initializing;
    }
}

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract ContextUpgradeable is Initializable {
    function __Context_init() internal onlyInitializing {
    }

    function __Context_init_unchained() internal onlyInitializing {
    }
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract OwnableUpgradeable is Initializable, ContextUpgradeable {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    function __Ownable_init() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    function __Ownable_init_unchained() internal onlyInitializing {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby disabling any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}

/**
 * @dev Contract module which provides access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership} and {acceptOwnership}.
 *
 * This module is used through inheritance. It will make available all functions
 * from parent (Ownable).
 */
abstract contract Ownable2StepUpgradeable is Initializable, OwnableUpgradeable {
    function __Ownable2Step_init() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    function __Ownable2Step_init_unchained() internal onlyInitializing {
    }
    address private _pendingOwner;

    event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Returns the address of the pending owner.
     */
    function pendingOwner() public view virtual returns (address) {
        return _pendingOwner;
    }

    /**
     * @dev Starts the ownership transfer of the contract to a new account. Replaces the pending transfer if there is one.
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual override onlyOwner {
        _pendingOwner = newOwner;
        emit OwnershipTransferStarted(owner(), newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`) and deletes any pending owner.
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual override {
        delete _pendingOwner;
        super._transferOwnership(newOwner);
    }

    /**
     * @dev The new owner accepts the ownership transfer.
     */
    function acceptOwnership() public virtual {
        address sender = _msgSender();
        require(pendingOwner() == sender, "Ownable2Step: caller is not the new owner");
        _transferOwnership(sender);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}

/**
 * @notice Base contract for all messaging contracts.
 * - Provides context on the local chain's domain.
 * - Provides ownership functionality.
 * - Will be providing pausing functionality when it is implemented.
 */
abstract contract MessagingBase is MultiCallable, Versioned, Ownable2StepUpgradeable {
    // ════════════════════════════════════════════════ IMMUTABLES ═════════════════════════════════════════════════════

    /// @notice Domain of the local chain, set once upon contract creation
    uint32 public immutable localDomain;
    // @notice Domain of the Synapse chain, set once upon contract creation
    uint32 public immutable synapseDomain;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev gap for upgrade safety
    uint256[50] private __GAP; // solhint-disable-line var-name-mixedcase

    constructor(string memory version_, uint32 synapseDomain_) Versioned(version_) {
        localDomain = ChainContext.chainId();
        synapseDomain = synapseDomain_;
    }

    // TODO: Implement pausing

    /**
     * @dev Should be impossible to renounce ownership;
     * we override OpenZeppelin OwnableUpgradeable's
     * implementation of renounceOwnership to make it a no-op
     */
    function renounceOwnership() public override onlyOwner {} //solhint-disable-line no-empty-blocks
}

/**
 * @notice Base contract for messaging contracts that are secured by the agent manager.
 * `AgentSecured` relies on `AgentManager` to provide the following functionality:
 * - Keep track of agents and their statuses.
 * - Pass agent-signed statements that were verified by the agent manager.
 * - These statements are considered valid indefinitely, unless the agent is disputed.
 * - Disputes are opened and resolved by the agent manager.
 * > `AgentSecured` implementation should never use statements signed by agents that are disputed.
 */
abstract contract AgentSecured is MessagingBase, IAgentSecured {
    // ════════════════════════════════════════════════ IMMUTABLES ═════════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    address public immutable agentManager;

    /// @inheritdoc IAgentSecured
    address public immutable inbox;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    // (agent index => their dispute status: flag, openedAt, resolvedAt)
    mapping(uint32 => DisputeStatus) internal _disputes;

    /// @dev gap for upgrade safety
    uint256[49] private __GAP; // solhint-disable-line var-name-mixedcase

    modifier onlyAgentManager() {
        if (msg.sender != agentManager) revert CallerNotAgentManager();
        _;
    }

    modifier onlyInbox() {
        if (msg.sender != inbox) revert CallerNotInbox();
        _;
    }

    constructor(string memory version_, uint32 synapseDomain_, address agentManager_, address inbox_)
        MessagingBase(version_, synapseDomain_)
    {
        agentManager = agentManager_;
        inbox = inbox_;
    }

    // ════════════════════════════════════════════ ONLY AGENT MANAGER ═════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    function openDispute(uint32 guardIndex, uint32 notaryIndex) external onlyAgentManager {
        uint40 openedAt = ChainContext.blockTimestamp();
        DisputeStatus memory status = DisputeStatus({flag: DisputeFlag.Pending, openedAt: openedAt, resolvedAt: 0});
        _disputes[guardIndex] = status;
        _disputes[notaryIndex] = status;
    }

    /// @inheritdoc IAgentSecured
    function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) external onlyAgentManager {
        // Update the dispute status of the slashed agent first.
        uint40 resolvedAt = ChainContext.blockTimestamp();
        _disputes[slashedIndex].flag = DisputeFlag.Slashed;
        _disputes[slashedIndex].resolvedAt = resolvedAt;
        // Mark the rival agent as not disputed, if there was an ongoing dispute.
        if (rivalIndex != 0) {
            _disputes[rivalIndex].flag = DisputeFlag.None;
            _disputes[rivalIndex].resolvedAt = resolvedAt;
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IAgentSecured
    function agentStatus(address agent) external view returns (AgentStatus memory) {
        return _agentStatus(agent);
    }

    /// @inheritdoc IAgentSecured
    function getAgent(uint256 index) external view returns (address agent, AgentStatus memory status) {
        return _getAgent(index);
    }

    /// @inheritdoc IAgentSecured
    function latestDisputeStatus(uint32 agentIndex) external view returns (DisputeStatus memory) {
        return _disputes[agentIndex];
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns status of the given agent: (flag, domain, index).
    function _agentStatus(address agent) internal view returns (AgentStatus memory) {
        return IAgentManager(agentManager).agentStatus(agent);
    }

    /// @dev Returns agent and their status for a given agent index. Returns zero values for non existing indexes.
    function _getAgent(uint256 index) internal view returns (address agent, AgentStatus memory status) {
        return IAgentManager(agentManager).getAgent(index);
    }

    /// @dev Checks if a Dispute exists for the given Notary. This function returns true, if
    /// Notary is in ongoing Dispute, or if Dispute was resolved not in Notary's favor.
    /// In both cases we can't trust Notary's data.
    /// Note: Agent-Secured contracts can trust Notary data only if both `_notaryDisputeExists` and
    /// `_notaryDisputeTimeout` return false.
    function _notaryDisputeExists(uint32 notaryIndex) internal view returns (bool) {
        return _disputes[notaryIndex].flag != DisputeFlag.None;
    }

    /// @dev Checks if a Notary recently won a Dispute and is still in the "post-dispute" timeout period.
    /// In this period we still can't trust Notary's data, though we can optimistically assume that
    /// that the data will be correct after the timeout (assuming no new Disputes are opened).
    /// Note: Agent-Secured contracts can trust Notary data only if both `_notaryDisputeExists` and
    /// `_notaryDisputeTimeout` return false.
    function _notaryDisputeTimeout(uint32 notaryIndex) internal view returns (bool) {
        DisputeStatus memory status = _disputes[notaryIndex];
        // Exit early if Notary is in ongoing Dispute / slashed.
        if (status.flag != DisputeFlag.None) return false;
        // Check if Notary has been in any Dispute at all.
        if (status.openedAt == 0) return false;
        // Otherwise check if the Dispute timeout is still active.
        return block.timestamp < status.resolvedAt + DISPUTE_TIMEOUT_NOTARY;
    }
}

/// @notice A collection of events emitted by the Destination contract
abstract contract DestinationEvents {
    event AgentRootAccepted(bytes32 agentRoot);
}

interface InterfaceDestination {
    /**
     * @notice Attempts to pass a quarantined Agent Merkle Root to a local Light Manager.
     * @dev Will do nothing, if root optimistic period is not over.
     * @return rootPending  Whether there is a pending agent merkle root left
     */
    function passAgentRoot() external returns (bool rootPending);

    /**
     * @notice Accepts an attestation, which local `AgentManager` verified to have been signed
     * by an active Notary for this chain.
     * > Attestation is created whenever a Notary-signed snapshot is saved in Summit on Synapse Chain.
     * - Saved Attestation could be later used to prove the inclusion of message in the Origin Merkle Tree.
     * - Messages coming from chains included in the Attestation's snapshot could be proven.
     * - Proof only exists for messages that were sent prior to when the Attestation's snapshot was taken.
     * > Will revert if any of these is true:
     * > - Called by anyone other than local `AgentManager`.
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is in Dispute.
     * > - Attestation's snapshot root has been previously submitted.
     * Note: agentRoot and snapGas have been verified by the local `AgentManager`.
     * @param notaryIndex       Index of Attestation Notary in Agent Merkle Tree
     * @param sigIndex          Index of stored Notary signature
     * @param attPayload        Raw payload with Attestation data
     * @param agentRoot         Agent Merkle Root from the Attestation
     * @param snapGas           Gas data for each chain in the Attestation's snapshot
     * @return wasAccepted      Whether the Attestation was accepted
     */
    function acceptAttestation(
        uint32 notaryIndex,
        uint256 sigIndex,
        bytes memory attPayload,
        bytes32 agentRoot,
        ChainGas[] memory snapGas
    ) external returns (bool wasAccepted);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the total amount of Notaries attestations that have been accepted.
     */
    function attestationsAmount() external view returns (uint256);

    /**
     * @notice Returns a Notary-signed attestation with a given index.
     * > Index refers to the list of all attestations accepted by this contract.
     * @dev Attestations are created on Synapse Chain whenever a Notary-signed snapshot is accepted by Summit.
     * Will return an empty signature if this contract is deployed on Synapse Chain.
     * @param index             Attestation index
     * @return attPayload       Raw payload with Attestation data
     * @return attSignature     Notary signature for the reported attestation
     */
    function getAttestation(uint256 index) external view returns (bytes memory attPayload, bytes memory attSignature);

    /**
     * @notice Returns the gas data for a given chain from the latest accepted attestation with that chain.
     * @dev Will return empty values if there is no data for the domain,
     * or if the notary who provided the data is in dispute.
     * @param domain            Domain for the chain
     * @return gasData          Gas data for the chain
     * @return dataMaturity     Gas data age in seconds
     */
    function getGasData(uint32 domain) external view returns (GasData gasData, uint256 dataMaturity);

    /**
     * Returns status of Destination contract as far as snapshot/agent roots are concerned
     * @return snapRootTime     Timestamp when latest snapshot root was accepted
     * @return agentRootTime    Timestamp when latest agent root was accepted
     * @return notaryIndex      Index of Notary who signed the latest agent root
     */
    function destStatus() external view returns (uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex);

    /**
     * Returns Agent Merkle Root to be passed to LightManager once its optimistic period is over.
     */
    function nextAgentRoot() external view returns (bytes32);

    /**
     * @notice Returns the nonce of the last attestation submitted by a Notary with a given agent index.
     * @dev Will return zero if the Notary hasn't submitted any attestations yet.
     */
    function lastAttestationNonce(uint32 notaryIndex) external view returns (uint32);
}

interface InterfaceLightManager {
    /**
     * @notice Updates agent status, using a proof against the latest known Agent Merkle Root.
     * @dev Will revert if the provided proof doesn't match the latest merkle root.
     * @param agent     Agent address
     * @param status    Structure specifying agent status: (flag, domain, index)
     * @param proof     Merkle proof of Active status for the agent
     */
    function updateAgentStatus(address agent, AgentStatus memory status, bytes32[] memory proof) external;

    /**
     * @notice Updates the root of Agent Merkle Tree that the Light Manager is tracking.
     * Could be only called by a local Destination contract, which is supposed to
     * verify the attested Agent Merkle Roots.
     * @param agentRoot_    New Agent Merkle Root
     */
    function setAgentRoot(bytes32 agentRoot_) external;

    /**
     * @notice Allows contract owner to set the agent root to resolve the "stuck" chain
     * by proposing the new agent root. The contract owner will be able to resolve the proposed
     * agent root after a certain period of time.
     * Note: this function could be called multiple times, each time the timer will be reset.
     * This could only be called if no fresh data has been submitted by the Notaries to the Inbox,
     * indicating that the chain is stuck for one of the reasons:
     * - All active Notaries are in Dispute.
     * - No active Notaries exist under the current agent root.
     * @dev Will revert if any of the following conditions is met:
     * - Caller is not the contract owner.
     * - Agent root is empty.
     * - The chain is not in a stuck state (has recently received a fresh data from the Notaries).
     * @param agentRoot_    New Agent Merkle Root that is proposed to be set
     */
    function proposeAgentRootWhenStuck(bytes32 agentRoot_) external;

    /**
     * @notice Allows contract owner to cancel the previously proposed agent root.
     * @dev Will revert if any of the following conditions is met:
     * - Caller is not the contract owner.
     * - No agent root was proposed.
     */
    function cancelProposedAgentRoot() external;

    /**
     * @notice Allows contract owner to resolve the previously proposed agent root.
     * This will update the agent root, allowing the agents to update their status, effectively
     * resolving the "stuck" chain.
     * @dev Will revert if any of the following conditions is met:
     * - Caller is not the contract owner.
     * - No agent root was proposed.
     * - Not enough time has passed since the agent root was proposed.
     */
    function resolveProposedAgentRoot() external;

    /**
     * @notice Withdraws locked base message tips from local Origin to the recipient.
     * @dev Could only be remote-called by BondingManager contract on Synapse Chain.
     * Note: as an extra security check this function returns its own selector, so that
     * Destination could verify that a "remote" function was called when executing a manager message.
     * @param recipient     Address to withdraw tips to
     * @param amount        Tips value to withdraw
     */
    function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount)
        external
        returns (bytes4 magicValue);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the latest proposed agent root and the timestamp when it was proposed.
     * @dev Will return zero values if no agent root was proposed, or if the proposed agent root
     * was already resolved.
     */
    function proposedAgentRootData() external view returns (bytes32 agentRoot_, uint256 proposedAt_);
}

interface IStatementInbox {
    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Accepts a Guard's state report signature, a Snapshot containing the reported State,
     * as well as Notary signature for the Snapshot.
     * > StateReport is a Guard statement saying "Reported state is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a StateReport and use other values from
     * `verifyStateWithSnapshot()` successful call that led to Notary being slashed in remote Origin.
     * > Will revert if any of these is true:
     * > - State Report signer is not an active Guard.
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot signer is not an active Notary.
     * > - State index is out of range.
     * > - The Guard or the Notary are already in a Dispute
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param srSignature       Guard signature for the report
     * @param snapPayload       Raw payload with Snapshot data
     * @param snapSignature     Notary signature for the Snapshot
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithSnapshot(
        uint8 stateIndex,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory snapSignature
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a Guard's state report signature, a Snapshot containing the reported State,
     * as well as Notary signature for the Attestation created from this Snapshot.
     * > StateReport is a Guard statement saying "Reported state is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a StateReport and use other values from
     * `verifyStateWithAttestation()` successful call that led to Notary being slashed in remote Origin.
     * > Will revert if any of these is true:
     * > - State Report signer is not an active Guard.
     * > - Snapshot payload is not properly formatted.
     * > - State index is out of range.
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from the Snapshot.
     * > - The Guard or the Notary are already in a Dispute
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param srSignature       Guard signature for the report
     * @param snapPayload       Raw payload with Snapshot data
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the Attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithAttestation(
        uint8 stateIndex,
        bytes memory srSignature,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a Guard's state report signature, a proof of inclusion of the reported State in an Attestation,
     * as well as Notary signature for the Attestation.
     * > StateReport is a Guard statement saying "Reported state is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a StateReport and use other values from
     * `verifyStateWithSnapshotProof()` successful call that led to Notary being slashed in remote Origin.
     * > Will revert if any of these is true:
     * > - State payload is not properly formatted.
     * > - State Report signer is not an active Guard.
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from State and Snapshot Proof.
     * > - Snapshot Proof's first element does not match the State metadata.
     * > - Snapshot Proof length exceeds Snapshot Tree Height.
     * > - State index is out of range.
     * > - The Guard or the Notary are already in a Dispute
     * @param stateIndex        Index of the reported State in the Snapshot
     * @param statePayload      Raw payload with State data that Guard reports as invalid
     * @param srSignature       Guard signature for the report
     * @param snapProof         Proof of inclusion of reported State's Left Leaf into Snapshot Merkle Tree
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the Attestation
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitStateReportWithSnapshotProof(
        uint8 stateIndex,
        bytes memory statePayload,
        bytes memory srSignature,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool wasAccepted);

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Verifies a message receipt signed by the Notary.
     * - Does nothing, if the receipt is valid (matches the saved receipt data for the referenced message).
     * - Slashes the Notary, if the receipt is invalid.
     * > Will revert if any of these is true:
     * > - Receipt payload is not properly formatted.
     * > - Receipt signer is not an active Notary.
     * > - Receipt's destination chain does not refer to this chain.
     * @param rcptPayload       Raw payload with Receipt data
     * @param rcptSignature     Notary signature for the receipt
     * @return isValidReceipt   Whether the provided receipt is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyReceipt(bytes memory rcptPayload, bytes memory rcptSignature)
        external
        returns (bool isValidReceipt);

    /**
     * @notice Verifies a Guard's receipt report signature.
     * - Does nothing, if the report is valid (if the reported receipt is invalid).
     * - Slashes the Guard, if the report is invalid (if the reported receipt is valid).
     * > Will revert if any of these is true:
     * > - Receipt payload is not properly formatted.
     * > - Receipt Report signer is not an active Guard.
     * > - Receipt does not refer to this chain.
     * @param rcptPayload       Raw payload with Receipt data that Guard reports as invalid
     * @param rrSignature       Guard signature for the report
     * @return isValidReport    Whether the provided report is valid.
     *                          Guard is slashed, if return value is FALSE.
     */
    function verifyReceiptReport(bytes memory rcptPayload, bytes memory rrSignature)
        external
        returns (bool isValidReport);

    /**
     * @notice Verifies a state from the snapshot, that was used for the Notary-signed attestation.
     * - Does nothing, if the state is valid (matches the historical state of this contract).
     * - Slashes the Notary, if the state is invalid.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from the Snapshot.
     * > - Snapshot payload is not properly formatted.
     * > - State index is out of range.
     * > - State does not refer to this chain.
     * @param stateIndex        State index to check
     * @param snapPayload       Raw payload with snapshot data
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValidState     Whether the provided state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyStateWithAttestation(
        uint8 stateIndex,
        bytes memory snapPayload,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState);

    /**
     * @notice Verifies a state from the snapshot, that was used for the Notary-signed attestation.
     * - Does nothing, if the state is valid (matches the historical state of this contract).
     * - Slashes the Notary, if the state is invalid.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * > - Attestation's snapshot root is not equal to Merkle Root derived from State and Snapshot Proof.
     * > - Snapshot Proof's first element does not match the State metadata.
     * > - Snapshot Proof length exceeds Snapshot Tree Height.
     * > - State payload is not properly formatted.
     * > - State index is out of range.
     * > - State does not refer to this chain.
     * @param stateIndex        Index of state in the snapshot
     * @param statePayload      Raw payload with State data to check
     * @param snapProof         Proof of inclusion of provided State's Left Leaf into Snapshot Merkle Tree
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValidState     Whether the provided state is valid.
     *                          Notary is slashed, if return value is FALSE.
     */
    function verifyStateWithSnapshotProof(
        uint8 stateIndex,
        bytes memory statePayload,
        bytes32[] memory snapProof,
        bytes memory attPayload,
        bytes memory attSignature
    ) external returns (bool isValidState);

    /**
     * @notice Verifies a state from the snapshot (a list of states) signed by a Guard or a Notary.
     * - Does nothing, if the state is valid (matches the historical state of this contract).
     * - Slashes the Agent, if the state is invalid.
     * > Will revert if any of these is true:
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot signer is not an active Agent.
     * > - State index is out of range.
     * > - State does not refer to this chain.
     * @param stateIndex        State index to check
     * @param snapPayload       Raw payload with snapshot data
     * @param snapSignature     Agent signature for the snapshot
     * @return isValidState     Whether the provided state is valid.
     *                          Agent is slashed, if return value is FALSE.
     */
    function verifyStateWithSnapshot(uint8 stateIndex, bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bool isValidState);

    /**
     * @notice Verifies a Guard's state report signature.
     *  - Does nothing, if the report is valid (if the reported state is invalid).
     *  - Slashes the Guard, if the report is invalid (if the reported state is valid).
     * > Will revert if any of these is true:
     * > - State payload is not properly formatted.
     * > - State Report signer is not an active Guard.
     * > - Reported State does not refer to this chain.
     * @param statePayload      Raw payload with State data that Guard reports as invalid
     * @param srSignature       Guard signature for the report
     * @return isValidReport    Whether the provided report is valid.
     *                          Guard is slashed, if return value is FALSE.
     */
    function verifyStateReport(bytes memory statePayload, bytes memory srSignature)
        external
        returns (bool isValidReport);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns the amount of Guard Reports stored in StatementInbox.
     * > Only reports that led to opening a Dispute are stored.
     */
    function getReportsAmount() external view returns (uint256);

    /**
     * @notice Returns the Guard report with the given index stored in StatementInbox.
     * > Only reports that led to opening a Dispute are stored.
     * @dev Will revert if report with given index doesn't exist.
     * @param index             Report index
     * @return statementPayload Raw payload with statement that Guard reported as invalid
     * @return reportSignature  Guard signature for the report
     */
    function getGuardReport(uint256 index)
        external
        view
        returns (bytes memory statementPayload, bytes memory reportSignature);

    /**
     * @notice Returns the signature with the given index stored in StatementInbox.
     * @dev Will revert if signature with given index doesn't exist.
     * @param index     Signature index
     * @return          Raw payload with signature
     */
    function getStoredSignature(uint256 index) external view returns (bytes memory);
}

// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════

library MerkleMath {
    // ═════════════════════════════════════════ BASIC MERKLE CALCULATIONS ═════════════════════════════════════════════

    /**
     * @notice Calculates the merkle root for the given leaf and merkle proof.
     * @dev Will revert if proof length exceeds the tree height.
     * @param index     Index of `leaf` in tree
     * @param leaf      Leaf of the merkle tree
     * @param proof     Proof of inclusion of `leaf` in the tree
     * @param height    Height of the merkle tree
     * @return root_    Calculated Merkle Root
     */
    function proofRoot(uint256 index, bytes32 leaf, bytes32[] memory proof, uint256 height)
        internal
        pure
        returns (bytes32 root_)
    {
        // Proof length could not exceed the tree height
        uint256 proofLen = proof.length;
        if (proofLen > height) revert TreeHeightTooLow();
        root_ = leaf;
        /// @dev Apply unchecked to all ++h operations
        unchecked {
            // Go up the tree levels from the leaf following the proof
            for (uint256 h = 0; h < proofLen; ++h) {
                // Get a sibling node on current level: this is proof[h]
                root_ = getParent(root_, proof[h], index, h);
            }
            // Go up to the root: the remaining siblings are EMPTY
            for (uint256 h = proofLen; h < height; ++h) {
                root_ = getParent(root_, bytes32(0), index, h);
            }
        }
    }

    /**
     * @notice Calculates the parent of a node on the path from one of the leafs to root.
     * @param node          Node on a path from tree leaf to root
     * @param sibling       Sibling for a given node
     * @param leafIndex     Index of the tree leaf
     * @param nodeHeight    "Level height" for `node` (ZERO for leafs, ORIGIN_TREE_HEIGHT for root)
     */
    function getParent(bytes32 node, bytes32 sibling, uint256 leafIndex, uint256 nodeHeight)
        internal
        pure
        returns (bytes32 parent)
    {
        // Index for `node` on its "tree level" is (leafIndex / 2**height)
        // "Left child" has even index, "right child" has odd index
        if ((leafIndex >> nodeHeight) & 1 == 0) {
            // Left child
            return getParent(node, sibling);
        } else {
            // Right child
            return getParent(sibling, node);
        }
    }

    /// @notice Calculates the parent of tow nodes in the merkle tree.
    /// @dev We use implementation with H(0,0) = 0
    /// This makes EVERY empty node in the tree equal to ZERO,
    /// saving us from storing H(0,0), H(H(0,0), H(0, 0)), and so on
    /// @param leftChild    Left child of the calculated node
    /// @param rightChild   Right child of the calculated node
    /// @return parent      Value for the node having above mentioned children
    function getParent(bytes32 leftChild, bytes32 rightChild) internal pure returns (bytes32 parent) {
        if (leftChild == bytes32(0) && rightChild == bytes32(0)) {
            return 0;
        } else {
            return keccak256(bytes.concat(leftChild, rightChild));
        }
    }

    // ════════════════════════════════ ROOT/PROOF CALCULATION FOR A LIST OF LEAFS ═════════════════════════════════════

    /**
     * @notice Calculates merkle root for a list of given leafs.
     * Merkle Tree is constructed by padding the list with ZERO values for leafs until list length is `2**height`.
     * Merkle Root is calculated for the constructed tree, and then saved in `leafs[0]`.
     * > Note:
     * > - `leafs` values are overwritten in the process to avoid excessive memory allocations.
     * > - Caller is expected not to reuse `hashes` list after the call, and only use `leafs[0]` value,
     * which is guaranteed to contain the calculated merkle root.
     * > - root is calculated using the `H(0,0) = 0` Merkle Tree implementation. See MerkleTree.sol for details.
     * @dev Amount of leaves should be at most `2**height`
     * @param hashes    List of leafs for the merkle tree (to be overwritten)
     * @param height    Height of the Merkle Tree to construct
     */
    function calculateRoot(bytes32[] memory hashes, uint256 height) internal pure {
        uint256 levelLength = hashes.length;
        // Amount of hashes could not exceed amount of leafs in tree with the given height
        if (levelLength > (1 << height)) revert TreeHeightTooLow();
        /// @dev h, leftIndex, rightIndex and levelLength never overflow
        unchecked {
            // Iterate `height` levels up from the leaf level
            // For every level we will only record "significant values", i.e. not equal to ZERO
            for (uint256 h = 0; h < height; ++h) {
                // Let H be the height of the "current level". H = 0 for the "leafs level".
                // Invariant: a total of 2**(HEIGHT-H) nodes are on the current level
                // Invariant: hashes[0 .. length) are "significant values" for the "current level" nodes
                // Invariant: bytes32(0) is the value for nodes with indexes [length .. 2**(HEIGHT-H))

                // Iterate over every pair of (leftChild, rightChild) on the current level
                for (uint256 leftIndex = 0; leftIndex < levelLength; leftIndex += 2) {
                    uint256 rightIndex = leftIndex + 1;
                    bytes32 leftChild = hashes[leftIndex];
                    // Note: rightChild might be ZERO
                    bytes32 rightChild = rightIndex < levelLength ? hashes[rightIndex] : bytes32(0);
                    // Record the parent hash in the same array. This will not affect
                    // further calculations for the same level: (leftIndex >> 1) <= leftIndex.
                    hashes[leftIndex >> 1] = getParent(leftChild, rightChild);
                }
                // Set length for the "parent level": the amount of iterations for the for loop above.
                levelLength = (levelLength + 1) >> 1;
            }
        }
    }

    /**
     * @notice Generates a proof of inclusion of a leaf in the list. If the requested index is outside
     * of the list range, generates a proof of inclusion for an empty leaf (proof of non-inclusion).
     * The Merkle Tree is constructed by padding the list with ZERO values until list length is a power of two
     * __AND__ index is in the extended list range. For example:
     *  - `hashes.length == 6` and `0 <= index <= 7` will "extend" the list to 8 entries.
     *  - `hashes.length == 6` and `7 < index <= 15` will "extend" the list to 16 entries.
     * > Note: `leafs` values are overwritten in the process to avoid excessive memory allocations.
     * Caller is expected not to reuse `hashes` list after the call.
     * @param hashes    List of leafs for the merkle tree (to be overwritten)
     * @param index     Leaf index to generate the proof for
     * @return proof    Generated merkle proof
     */
    function calculateProof(bytes32[] memory hashes, uint256 index) internal pure returns (bytes32[] memory proof) {
        // Use only meaningful values for the shortened proof
        // Check if index is within the list range (we want to generates proofs for outside leafs as well)
        uint256 height = getHeight(index < hashes.length ? hashes.length : (index + 1));
        proof = new bytes32[](height);
        uint256 levelLength = hashes.length;
        /// @dev h, leftIndex, rightIndex and levelLength never overflow
        unchecked {
            // Iterate `height` levels up from the leaf level
            // For every level we will only record "significant values", i.e. not equal to ZERO
            for (uint256 h = 0; h < height; ++h) {
                // Use sibling for the merkle proof; `index^1` is index of our sibling
                proof[h] = (index ^ 1 < levelLength) ? hashes[index ^ 1] : bytes32(0);

                // Let H be the height of the "current level". H = 0 for the "leafs level".
                // Invariant: a total of 2**(HEIGHT-H) nodes are on the current level
                // Invariant: hashes[0 .. length) are "significant values" for the "current level" nodes
                // Invariant: bytes32(0) is the value for nodes with indexes [length .. 2**(HEIGHT-H))

                // Iterate over every pair of (leftChild, rightChild) on the current level
                for (uint256 leftIndex = 0; leftIndex < levelLength; leftIndex += 2) {
                    uint256 rightIndex = leftIndex + 1;
                    bytes32 leftChild = hashes[leftIndex];
                    // Note: rightChild might be ZERO
                    bytes32 rightChild = rightIndex < levelLength ? hashes[rightIndex] : bytes32(0);
                    // Record the parent hash in the same array. This will not affect
                    // further calculations for the same level: (leftIndex >> 1) <= leftIndex.
                    hashes[leftIndex >> 1] = getParent(leftChild, rightChild);
                }
                // Set length for the "parent level"
                levelLength = (levelLength + 1) >> 1;
                // Traverse to parent node
                index >>= 1;
            }
        }
    }

    /// @notice Returns the height of the tree having a given amount of leafs.
    function getHeight(uint256 leafs) internal pure returns (uint256 height) {
        uint256 amount = 1;
        while (amount < leafs) {
            unchecked {
                ++height;
            }
            amount <<= 1;
        }
    }
}

/// Request is encoded data with "message execution request".
type Request is uint192;

using RequestLib for Request global;

/// Library for formatting _the request part_ of _the base messages_.
/// - Request represents a message sender requirements for the message execution on the destination chain.
/// - Request occupies a single storage word, and thus is stored on stack instead of being stored in memory.
/// > gasDrop field is included for future compatibility and is ignored at the moment.
///
/// # Request stack layout (from highest bits to lowest)
///
/// | Position   | Field    | Type   | Bytes | Description                                          |
/// | ---------- | -------- | ------ | ----- | ---------------------------------------------------- |
/// | (024..012] | gasDrop  | uint96 | 12    | Minimum amount of gas token to drop to the recipient |
/// | (012..004] | gasLimit | uint64 | 8     | Minimum amount of gas units to supply for execution  |
/// | (004..000] | version  | uint32 | 4     | Base message version to pass to the recipient        |

library RequestLib {
    /// @dev Amount of bits to shift to gasDrop field
    uint192 private constant SHIFT_GAS_DROP = 12 * 8;
    /// @dev Amount of bits to shift to gasLimit field
    uint192 private constant SHIFT_GAS_LIMIT = 4 * 8;

    /// @notice Returns an encoded request with the given fields
    /// @param gasDrop_     Minimum amount of gas token to drop to the recipient (ignored at the moment)
    /// @param gasLimit_    Minimum amount of gas units to supply for execution
    /// @param version_     Base message version to pass to the recipient
    function encodeRequest(uint96 gasDrop_, uint64 gasLimit_, uint32 version_) internal pure returns (Request) {
        // Casts below are upcasts, so they are safe
        return Request.wrap(uint192(gasDrop_) << SHIFT_GAS_DROP | uint192(gasLimit_) << SHIFT_GAS_LIMIT | version_);
    }

    /// @notice Wraps the padded encoded request into a Request-typed value.
    /// @dev The "padded" request is simply an encoded request casted to uint256 (highest bits are set to zero).
    /// Casting to uint256 is done automatically in Solidity, so no extra actions from consumers are needed.
    /// The highest bits are discarded, so that the contracts dealing with encoded requests
    /// don't need to be updated, if a new field is added.
    function wrapPadded(uint256 paddedRequest) internal pure returns (Request) {
        // Casting to uint192 will truncate the highest bits, which is the behavior we want
        return Request.wrap(uint192(paddedRequest));
    }

    /// @notice Returns the requested of gas token to drop to the recipient.
    function gasDrop(Request request) internal pure returns (uint96) {
        // Casting to uint96 will truncate the highest bits, which is the behavior we want
        return uint96(Request.unwrap(request) >> SHIFT_GAS_DROP);
    }

    /// @notice Returns the requested minimum amount of gas units to supply for execution.
    function gasLimit(Request request) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Request.unwrap(request) >> SHIFT_GAS_LIMIT);
    }

    /// @notice Returns the requested base message version to pass to the recipient.
    function version(Request request) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Request.unwrap(request));
    }
}

/// Tips is encoded data with "tips paid for sending a base message".
/// Note: even though uint256 is also an underlying type for MemView, Tips is stored ON STACK.
type Tips is uint256;

using TipsLib for Tips global;

/// # Tips
/// Library for formatting _the tips part_ of _the base messages_.
///
/// ## How the tips are awarded
/// Tips are paid for sending a base message, and are split across all the agents that
/// made the message execution on destination chain possible.
/// ### Summit tips
/// Split between:
///     - Guard posting a snapshot with state ST_G for the origin chain.
///     - Notary posting a snapshot SN_N using ST_G. This creates attestation A.
///     - Notary posting a message receipt after it is executed on destination chain.
/// ### Attestation tips
/// Paid to:
///     - Notary posting attestation A to destination chain.
/// ### Execution tips
/// Paid to:
///     - First executor performing a valid execution attempt (correct proofs, optimistic period over),
///      using attestation A to prove message inclusion on origin chain, whether the recipient reverted or not.
/// ### Delivery tips.
/// Paid to:
///     - Executor who successfully executed the message on destination chain.
///
/// ## Tips encoding
/// - Tips occupy a single storage word, and thus are stored on stack instead of being stored in memory.
/// - The actual tip values should be determined by multiplying stored values by divided by TIPS_MULTIPLIER=2**32.
/// - Tips are packed into a single word of storage, while allowing real values up to ~8*10**28 for every tip category.
/// > The only downside is that the "real tip values" are now multiplies of ~4*10**9, which should be fine even for
/// the chains with the most expensive gas currency.
/// # Tips stack layout (from highest bits to lowest)
///
/// | Position   | Field          | Type   | Bytes | Description                                                |
/// | ---------- | -------------- | ------ | ----- | ---------------------------------------------------------- |
/// | (032..024] | summitTip      | uint64 | 8     | Tip for agents interacting with Summit contract            |
/// | (024..016] | attestationTip | uint64 | 8     | Tip for Notary posting attestation to Destination contract |
/// | (016..008] | executionTip   | uint64 | 8     | Tip for valid execution attempt on destination chain       |
/// | (008..000] | deliveryTip    | uint64 | 8     | Tip for successful message delivery on destination chain   |

library TipsLib {
    using SafeCast for uint256;

    /// @dev Amount of bits to shift to summitTip field
    uint256 private constant SHIFT_SUMMIT_TIP = 24 * 8;
    /// @dev Amount of bits to shift to attestationTip field
    uint256 private constant SHIFT_ATTESTATION_TIP = 16 * 8;
    /// @dev Amount of bits to shift to executionTip field
    uint256 private constant SHIFT_EXECUTION_TIP = 8 * 8;

    // ═══════════════════════════════════════════════════ TIPS ════════════════════════════════════════════════════════

    /// @notice Returns encoded tips with the given fields
    /// @param summitTip_        Tip for agents interacting with Summit contract, divided by TIPS_MULTIPLIER
    /// @param attestationTip_   Tip for Notary posting attestation to Destination contract, divided by TIPS_MULTIPLIER
    /// @param executionTip_     Tip for valid execution attempt on destination chain, divided by TIPS_MULTIPLIER
    /// @param deliveryTip_      Tip for successful message delivery on destination chain, divided by TIPS_MULTIPLIER
    function encodeTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_)
        internal
        pure
        returns (Tips)
    {
        // forgefmt: disable-next-item
        return Tips.wrap(
            uint256(summitTip_) << SHIFT_SUMMIT_TIP |
            uint256(attestationTip_) << SHIFT_ATTESTATION_TIP |
            uint256(executionTip_) << SHIFT_EXECUTION_TIP |
            uint256(deliveryTip_)
        );
    }

    /// @notice Convenience function to encode tips with uint256 values.
    function encodeTips256(uint256 summitTip_, uint256 attestationTip_, uint256 executionTip_, uint256 deliveryTip_)
        internal
        pure
        returns (Tips)
    {
        // In practice, the tips amounts are not supposed to be higher than 2**96, and with 32 bits of granularity
        // using uint64 is enough to store the values. However, we still check for overflow just in case.
        // TODO: consider using Number type to store the tips values.
        return encodeTips({
            summitTip_: (summitTip_ >> TIPS_GRANULARITY).toUint64(),
            attestationTip_: (attestationTip_ >> TIPS_GRANULARITY).toUint64(),
            executionTip_: (executionTip_ >> TIPS_GRANULARITY).toUint64(),
            deliveryTip_: (deliveryTip_ >> TIPS_GRANULARITY).toUint64()
        });
    }

    /// @notice Wraps the padded encoded tips into a Tips-typed value.
    /// @dev There is no actual padding here, as the underlying type is already uint256,
    /// but we include this function for consistency and to be future-proof, if tips will eventually use anything
    /// smaller than uint256.
    function wrapPadded(uint256 paddedTips) internal pure returns (Tips) {
        return Tips.wrap(paddedTips);
    }

    /**
     * @notice Returns a formatted Tips payload specifying empty tips.
     * @return Formatted tips
     */
    function emptyTips() internal pure returns (Tips) {
        return Tips.wrap(0);
    }

    /// @notice Returns tips's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(Tips tips) internal pure returns (bytes32 hashedTips) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Store tips in scratch space
            mstore(0, tips)
            // Compute hash of tips padded to 32 bytes
            hashedTips := keccak256(0, 32)
        }
    }

    // ═══════════════════════════════════════════════ TIPS SLICING ════════════════════════════════════════════════════

    /// @notice Returns summitTip field
    function summitTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips) >> SHIFT_SUMMIT_TIP);
    }

    /// @notice Returns attestationTip field
    function attestationTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips) >> SHIFT_ATTESTATION_TIP);
    }

    /// @notice Returns executionTip field
    function executionTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips) >> SHIFT_EXECUTION_TIP);
    }

    /// @notice Returns deliveryTip field
    function deliveryTip(Tips tips) internal pure returns (uint64) {
        // Casting to uint64 will truncate the highest bits, which is the behavior we want
        return uint64(Tips.unwrap(tips));
    }

    // ════════════════════════════════════════════════ TIPS VALUE ═════════════════════════════════════════════════════

    /// @notice Returns total value of the tips payload.
    /// This is the sum of the encoded values, scaled up by TIPS_MULTIPLIER
    function value(Tips tips) internal pure returns (uint256 value_) {
        value_ = uint256(tips.summitTip()) + tips.attestationTip() + tips.executionTip() + tips.deliveryTip();
        value_ <<= TIPS_GRANULARITY;
    }

    /// @notice Increases the delivery tip to match the new value.
    function matchValue(Tips tips, uint256 newValue) internal pure returns (Tips newTips) {
        uint256 oldValue = tips.value();
        if (newValue < oldValue) revert TipsValueTooLow();
        // We want to increase the delivery tip, while keeping the other tips the same
        unchecked {
            uint256 delta = (newValue - oldValue) >> TIPS_GRANULARITY;
            // `delta` fits into uint224, as TIPS_GRANULARITY is 32, so this never overflows uint256.
            // In practice, this will never overflow uint64 as well, but we still check it just in case.
            if (delta + tips.deliveryTip() > type(uint64).max) revert TipsOverflow();
            // Delivery tips occupy lowest 8 bytes, so we can just add delta to the tips value
            // to effectively increase the delivery tip (knowing that delta fits into uint64).
            newTips = Tips.wrap(Tips.unwrap(tips) + delta);
        }
    }
}

/// BaseMessage is a memory view over the base message supported by Origin-Destination
type BaseMessage is uint256;

using BaseMessageLib for BaseMessage global;

/// BaseMessage structure represents a base message sent via the Origin-Destination contracts.
/// - It only contains data relevant to the base message, the rest of data is encoded in the message header.
/// - `sender` and `recipient` for EVM chains are EVM addresses casted to bytes32, while preserving left-alignment.
/// - `tips` and `request` parameters are specified by a message sender
/// > Origin will calculate minimum tips for given request and content length, and will reject messages with tips
/// lower than that.
///
/// # Memory layout of BaseMessage fields
///
/// | Position   | Field     | Type    | Bytes | Description                            |
/// | ---------- | --------- | ------- | ----- | -------------------------------------- |
/// | [000..032) | tips      | uint256 | 32    | Encoded tips paid on origin chain      |
/// | [032..064) | sender    | bytes32 | 32    | Sender address on origin chain         |
/// | [064..096) | recipient | bytes32 | 32    | Recipient address on destination chain |
/// | [096..116) | request   | uint160 | 20    | Encoded request for message execution  |
/// | [104..AAA) | content   | bytes   | ??    | Content to be passed to recipient      |
library BaseMessageLib {
    using MemViewLib for bytes;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_TIPS = 0;
    uint256 private constant OFFSET_SENDER = 32;
    uint256 private constant OFFSET_RECIPIENT = 64;
    uint256 private constant OFFSET_REQUEST = OFFSET_RECIPIENT + TIPS_LENGTH;
    uint256 private constant OFFSET_CONTENT = OFFSET_REQUEST + REQUEST_LENGTH;

    // ═══════════════════════════════════════════════ BASE MESSAGE ════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted BaseMessage payload with provided fields.
     * @param tips_         Encoded tips information
     * @param sender_       Sender address on origin chain
     * @param recipient_    Recipient address on destination chain
     * @param request_      Encoded request for message execution
     * @param content_      Raw content to be passed to recipient on destination chain
     * @return Formatted base message
     */
    function formatBaseMessage(Tips tips_, bytes32 sender_, bytes32 recipient_, Request request_, bytes memory content_)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(tips_, sender_, recipient_, request_, content_);
    }

    /**
     * @notice Returns a BaseMessage view over the given payload.
     * @dev Will revert if the payload is not a base message.
     */
    function castToBaseMessage(bytes memory payload) internal pure returns (BaseMessage) {
        return castToBaseMessage(payload.ref());
    }

    /**
     * @notice Casts a memory view to a BaseMessage view.
     * @dev Will revert if the memory view is not over a base message payload.
     */
    function castToBaseMessage(MemView memView) internal pure returns (BaseMessage) {
        if (!isBaseMessage(memView)) revert UnformattedBaseMessage();
        return BaseMessage.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted BaseMessage.
    function isBaseMessage(MemView memView) internal pure returns (bool) {
        // Check if sender, recipient, tips fields exist
        return (memView.len() >= OFFSET_CONTENT);
        // Content could be empty, so we don't check that
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(BaseMessage baseMessage) internal pure returns (MemView) {
        return MemView.wrap(BaseMessage.unwrap(baseMessage));
    }

    /// @notice Returns baseMessage's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(BaseMessage baseMessage) internal pure returns (bytes32) {
        // We hash "tips" and "everything but tips" to make tips proofs easier to verify
        return MerkleMath.getParent(baseMessage.tips().leaf(), baseMessage.bodyLeaf());
    }

    /// @notice Returns hash for the "everything but tips" part of the base message.
    function bodyLeaf(BaseMessage baseMessage) internal pure returns (bytes32) {
        return baseMessage.unwrap().sliceFrom({index_: OFFSET_SENDER}).keccak();
    }

    // ═══════════════════════════════════════════ BASE MESSAGE SLICING ════════════════════════════════════════════════

    /// @notice Returns encoded tips paid on origin chain.
    function tips(BaseMessage baseMessage) internal pure returns (Tips) {
        return TipsLib.wrapPadded((baseMessage.unwrap().indexUint({index_: OFFSET_TIPS, bytes_: TIPS_LENGTH})));
    }

    /// @notice Returns sender address on origin chain.
    function sender(BaseMessage baseMessage) internal pure returns (bytes32) {
        return baseMessage.unwrap().index({index_: OFFSET_SENDER, bytes_: 32});
    }

    /// @notice Returns recipient address on destination chain.
    function recipient(BaseMessage baseMessage) internal pure returns (bytes32) {
        return baseMessage.unwrap().index({index_: OFFSET_RECIPIENT, bytes_: 32});
    }

    /// @notice Returns an encoded request for message execution on destination chain.
    function request(BaseMessage baseMessage) internal pure returns (Request) {
        return RequestLib.wrapPadded((baseMessage.unwrap().indexUint({index_: OFFSET_REQUEST, bytes_: REQUEST_LENGTH})));
    }

    /// @notice Returns an untyped memory view over the content to be passed to recipient.
    function content(BaseMessage baseMessage) internal pure returns (MemView) {
        return baseMessage.unwrap().sliceFrom({index_: OFFSET_CONTENT});
    }
}

/// @notice Library for performing safe calls to the recipients.
/// Inspired by https://github.com/nomad-xyz/ExcessivelySafeCall
library SafeCall {
    /// @notice Performs a call to the recipient using the provided gas limit, msg.value and payload
    /// in a safe way:
    /// - If the recipient is not a contract, false is returned instead of reverting.
    /// - If the recipient call reverts for any reason, false is returned instead of reverting.
    /// - If the recipient call succeeds, true is returned, and any returned data is ignored.
    /// @param recipient        The recipient of the call
    /// @param gasLimit         The gas limit to use for the call
    /// @param msgValue         The msg.value to use for the call
    /// @param payload          The payload to use for the call
    /// @return success         True if the call succeeded, false otherwise
    function safeCall(address recipient, uint256 gasLimit, uint256 msgValue, bytes memory payload)
        internal
        returns (bool success)
    {
        // Exit early if the recipient is not a contract
        if (recipient.code.length == 0) return false;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Perform the call to the recipient, while ignoring any returned data
            // call(g, a, v, in, insize, out, outsize) => returns 0 on error, 1 on success
            success := call(gasLimit, recipient, msgValue, add(payload, 0x20), mload(payload), 0, 0)
        }
    }
}

/// Header is encoded data with "general routing information".
type Header is uint136;

using HeaderLib for Header global;

/// Types of messages supported by Origin-Destination
/// - Base: message sent by protocol user, contains tips
/// - Manager: message sent between AgentManager contracts located on different chains, no tips
enum MessageFlag {
    Base,
    Manager
}

using HeaderLib for MessageFlag global;

/// Library for formatting _the header part_ of _the messages used by Origin and Destination_.
/// - Header represents general information for routing a Message for Origin and Destination.
/// - Header occupies a single storage word, and thus is stored on stack instead of being stored in memory.
///
/// # Header stack layout (from highest bits to lowest)
///
/// | Position   | Field            | Type   | Bytes | Description                             |
/// | ---------- | ---------------- | ------ | ----- | --------------------------------------- |
/// | (017..016] | flag             | uint8  | 1     | Flag specifying the type of message     |
/// | (016..012] | origin           | uint32 | 4     | Domain where message originated         |
/// | (012..008] | nonce            | uint32 | 4     | Message nonce on the origin domain      |
/// | (008..004] | destination      | uint32 | 4     | Domain where message will be executed   |
/// | (004..000] | optimisticPeriod | uint32 | 4     | Optimistic period that will be enforced |
library HeaderLib {
    /// @dev Amount of bits to shift to flag field
    uint136 private constant SHIFT_FLAG = 16 * 8;
    /// @dev Amount of bits to shift to origin field
    uint136 private constant SHIFT_ORIGIN = 12 * 8;
    /// @dev Amount of bits to shift to nonce field
    uint136 private constant SHIFT_NONCE = 8 * 8;
    /// @dev Amount of bits to shift to destination field
    uint136 private constant SHIFT_DESTINATION = 4 * 8;

    /// @notice Returns an encoded header with provided fields
    /// @param origin_              Domain of origin chain
    /// @param nonce_               Message nonce on origin chain
    /// @param destination_         Domain of destination chain
    /// @param optimisticPeriod_    Optimistic period for message execution
    function encodeHeader(
        MessageFlag flag_,
        uint32 origin_,
        uint32 nonce_,
        uint32 destination_,
        uint32 optimisticPeriod_
    ) internal pure returns (Header) {
        // All casts below are upcasts, so they are safe
        // forgefmt: disable-next-item
        return Header.wrap(
            uint136(uint8(flag_)) << SHIFT_FLAG |
            uint136(origin_) << SHIFT_ORIGIN |
            uint136(nonce_) << SHIFT_NONCE |
            uint136(destination_) << SHIFT_DESTINATION |
            uint136(optimisticPeriod_)
        );
    }

    /// @notice Checks that the header is a valid encoded header.
    function isHeader(uint256 paddedHeader) internal pure returns (bool) {
        // Check that flag is within range
        return _flag(paddedHeader) <= uint8(type(MessageFlag).max);
    }

    /// @notice Wraps the padded encoded request into a Header-typed value.
    /// @dev The "padded" header is simply an encoded header casted to uint256 (highest bits are set to zero).
    /// Casting to uint256 is done automatically in Solidity, so no extra actions from consumers are needed.
    /// The highest bits are discarded, so that the contracts dealing with encoded headers
    /// don't need to be updated, if a new field is added.
    function wrapPadded(uint256 paddedHeader) internal pure returns (Header) {
        // Check that flag is within range
        if (!isHeader(paddedHeader)) revert FlagOutOfRange();
        // Casting to uint136 will truncate the highest bits, which is the behavior we want
        return Header.wrap(uint136(paddedHeader));
    }

    /// @notice Returns header's hash: a leaf to be inserted in the "Message mini-Merkle tree".
    function leaf(Header header) internal pure returns (bytes32 hashedHeader) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Store header in scratch space
            mstore(0, header)
            // Compute hash of header padded to 32 bytes
            hashedHeader := keccak256(0, 32)
        }
    }

    // ══════════════════════════════════════════════ HEADER SLICING ═══════════════════════════════════════════════════

    /// @notice Returns header's flag field
    function flag(Header header) internal pure returns (MessageFlag) {
        // We check that flag is within range when wrapping the header, so this cast is safe
        return MessageFlag(_flag(Header.unwrap(header)));
    }

    /// @notice Returns header's origin field
    function origin(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header) >> SHIFT_ORIGIN);
    }

    /// @notice Returns header's nonce field
    function nonce(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header) >> SHIFT_NONCE);
    }

    /// @notice Returns header's destination field
    function destination(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header) >> SHIFT_DESTINATION);
    }

    /// @notice Returns header's optimistic seconds field
    function optimisticPeriod(Header header) internal pure returns (uint32) {
        // Casting to uint32 will truncate the highest bits, which is the behavior we want
        return uint32(Header.unwrap(header));
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns header's flag field without casting to MessageFlag
    function _flag(uint256 paddedHeader) private pure returns (uint8) {
        // Casting to uint8 will truncate the highest bits, which is the behavior we want
        return uint8(paddedHeader >> SHIFT_FLAG);
    }
}

/// Message is a memory over over a formatted message payload.
type Message is uint256;

using MessageLib for Message global;

/// Library for formatting the various messages supported by Origin and Destination.
///
/// # Message memory layout
///
/// | Position   | Field  | Type    | Bytes | Description                                             |
/// | ---------- | ------ | ------- | ----- | ------------------------------------------------------- |
/// | [000..017) | header | uint136 | 17    | Encoded general routing information for the message     |
/// | [017..AAA) | body   | bytes   | ??    | Formatted payload (according to flag) with message body |
library MessageLib {
    using BaseMessageLib for MemView;
    using ByteString for MemView;
    using MemViewLib for bytes;
    using HeaderLib for MemView;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_HEADER = 0;
    uint256 private constant OFFSET_BODY = OFFSET_HEADER + HEADER_LENGTH;

    // ══════════════════════════════════════════════════ MESSAGE ══════════════════════════════════════════════════════

    /**
     * @notice Returns formatted message with provided fields.
     * @param header_   Encoded general routing information for the message
     * @param body_     Formatted payload (according to flag) with message body
     * @return Formatted message
     */
    function formatMessage(Header header_, bytes memory body_) internal pure returns (bytes memory) {
        return abi.encodePacked(header_, body_);
    }

    /**
     * @notice Returns a Message view over for the given payload.
     * @dev Will revert if the payload is not a message payload.
     */
    function castToMessage(bytes memory payload) internal pure returns (Message) {
        return castToMessage(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Message view.
     * @dev Will revert if the memory view is not over a message payload.
     */
    function castToMessage(MemView memView) internal pure returns (Message) {
        if (!isMessage(memView)) revert UnformattedMessage();
        return Message.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a payload is a formatted Message.
     */
    function isMessage(MemView memView) internal pure returns (bool) {
        uint256 length = memView.len();
        // Check if headers exist in the payload
        if (length < OFFSET_BODY) return false;
        // Check that Header is valid
        uint256 paddedHeader = _header(memView);
        if (!HeaderLib.isHeader(paddedHeader)) return false;
        // Check that body is formatted according to the flag
        // Only Base/Manager message flags exist
        if (HeaderLib.wrapPadded(paddedHeader).flag() == MessageFlag.Base) {
            // Check if body is a formatted base message
            return _body(memView).isBaseMessage();
        } else {
            // Check if body is a formatted calldata for AgentManager call
            return _body(memView).isCallData();
        }
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Message message) internal pure returns (MemView) {
        return MemView.wrap(Message.unwrap(message));
    }

    /// @notice Returns message's hash: a leaf to be inserted in the Merkle tree.
    function leaf(Message message) internal pure returns (bytes32) {
        // We hash header and body separately to make message proofs easier to verify
        Header header_ = message.header();
        // Only Base/Manager message flags exist
        if (header_.flag() == MessageFlag.Base) {
            return MerkleMath.getParent(header_.leaf(), message.body().castToBaseMessage().leaf());
        } else {
            return MerkleMath.getParent(header_.leaf(), message.body().castToCallData().leaf());
        }
    }

    // ══════════════════════════════════════════════ MESSAGE SLICING ══════════════════════════════════════════════════

    /// @notice Returns message's encoded header field.
    function header(Message message) internal pure returns (Header) {
        return HeaderLib.wrapPadded((message.unwrap().indexUint({index_: OFFSET_HEADER, bytes_: HEADER_LENGTH})));
    }

    /// @notice Returns message's body field as an untyped memory view.
    function body(Message message) internal pure returns (MemView) {
        MemView memView = message.unwrap();
        return _body(memView);
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Returns message's padded header without checking that it is a valid header.
    function _header(MemView memView) private pure returns (uint256) {
        return memView.indexUint({index_: OFFSET_HEADER, bytes_: HEADER_LENGTH});
    }

    /// @dev Returns an untyped memory view over the body field without checking
    /// if the whole payload or the body are properly formatted.
    function _body(MemView memView) private pure returns (MemView) {
        return memView.sliceFrom({index_: OFFSET_BODY});
    }
}

/// Receipt is a memory view over a formatted "full receipt" payload.
type Receipt is uint256;

using ReceiptLib for Receipt global;

/// Receipt structure represents a Notary statement that a certain message has been executed in `ExecutionHub`.
/// - It is possible to prove the correctness of the tips payload using the message hash, therefore tips are not
///   included in the receipt.
/// - Receipt is signed by a Notary and submitted to `Summit` in order to initiate the tips distribution for an
///   executed message.
/// - If a message execution fails the first time, the `finalExecutor` field will be set to zero address. In this
///   case, when the message is finally executed successfully, the `finalExecutor` field will be updated. Both
///   receipts will be considered valid.
/// # Memory layout of Receipt fields
///
/// | Position   | Field         | Type    | Bytes | Description                                      |
/// | ---------- | ------------- | ------- | ----- | ------------------------------------------------ |
/// | [000..004) | origin        | uint32  | 4     | Domain where message originated                  |
/// | [004..008) | destination   | uint32  | 4     | Domain where message was executed                |
/// | [008..040) | messageHash   | bytes32 | 32    | Hash of the message                              |
/// | [040..072) | snapshotRoot  | bytes32 | 32    | Snapshot root used for proving the message       |
/// | [072..073) | stateIndex    | uint8   | 1     | Index of state used for the snapshot proof       |
/// | [073..093) | attNotary     | address | 20    | Notary who posted attestation with snapshot root |
/// | [093..113) | firstExecutor | address | 20    | Executor who performed first valid execution     |
/// | [113..133) | finalExecutor | address | 20    | Executor who successfully executed the message   |
library ReceiptLib {
    using MemViewLib for bytes;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_ORIGIN = 0;
    uint256 private constant OFFSET_DESTINATION = 4;
    uint256 private constant OFFSET_MESSAGE_HASH = 8;
    uint256 private constant OFFSET_SNAPSHOT_ROOT = 40;
    uint256 private constant OFFSET_STATE_INDEX = 72;
    uint256 private constant OFFSET_ATT_NOTARY = 73;
    uint256 private constant OFFSET_FIRST_EXECUTOR = 93;
    uint256 private constant OFFSET_FINAL_EXECUTOR = 113;

    // ═════════════════════════════════════════════════ RECEIPT ═════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Receipt payload with provided fields.
     * @param origin_           Domain where message originated
     * @param destination_      Domain where message was executed
     * @param messageHash_      Hash of the message
     * @param snapshotRoot_     Snapshot root used for proving the message
     * @param stateIndex_       Index of state used for the snapshot proof
     * @param attNotary_        Notary who posted attestation with snapshot root
     * @param firstExecutor_    Executor who performed first valid execution attempt
     * @param finalExecutor_    Executor who successfully executed the message
     * @return Formatted receipt
     */
    function formatReceipt(
        uint32 origin_,
        uint32 destination_,
        bytes32 messageHash_,
        bytes32 snapshotRoot_,
        uint8 stateIndex_,
        address attNotary_,
        address firstExecutor_,
        address finalExecutor_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            origin_, destination_, messageHash_, snapshotRoot_, stateIndex_, attNotary_, firstExecutor_, finalExecutor_
        );
    }

    /**
     * @notice Returns a Receipt view over the given payload.
     * @dev Will revert if the payload is not a receipt.
     */
    function castToReceipt(bytes memory payload) internal pure returns (Receipt) {
        return castToReceipt(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Receipt view.
     * @dev Will revert if the memory view is not over a receipt.
     */
    function castToReceipt(MemView memView) internal pure returns (Receipt) {
        if (!isReceipt(memView)) revert UnformattedReceipt();
        return Receipt.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted Receipt.
    function isReceipt(MemView memView) internal pure returns (bool) {
        // Check payload length
        return memView.len() == RECEIPT_LENGTH;
    }

    /// @notice Returns the hash of an Receipt, that could be later signed by a Notary to signal
    /// that the receipt is valid.
    function hashValid(Receipt receipt) internal pure returns (bytes32) {
        // The final hash to sign is keccak(receiptSalt, keccak(receipt))
        return receipt.unwrap().keccakSalted(RECEIPT_VALID_SALT);
    }

    /// @notice Returns the hash of a Receipt, that could be later signed by a Guard to signal
    /// that the receipt is invalid.
    function hashInvalid(Receipt receipt) internal pure returns (bytes32) {
        // The final hash to sign is keccak(receiptBodyInvalidSalt, keccak(receipt))
        return receipt.unwrap().keccakSalted(RECEIPT_INVALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Receipt receipt) internal pure returns (MemView) {
        return MemView.wrap(Receipt.unwrap(receipt));
    }

    /// @notice Compares two Receipt structures.
    function equals(Receipt a, Receipt b) internal pure returns (bool) {
        // Length of a Receipt payload is fixed, so we just need to compare the hashes
        return a.unwrap().keccak() == b.unwrap().keccak();
    }

    // ═════════════════════════════════════════════ RECEIPT SLICING ═════════════════════════════════════════════════

    /// @notice Returns receipt's origin field
    function origin(Receipt receipt) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(receipt.unwrap().indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns receipt's destination field
    function destination(Receipt receipt) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(receipt.unwrap().indexUint({index_: OFFSET_DESTINATION, bytes_: 4}));
    }

    /// @notice Returns receipt's "message hash" field
    function messageHash(Receipt receipt) internal pure returns (bytes32) {
        return receipt.unwrap().index({index_: OFFSET_MESSAGE_HASH, bytes_: 32});
    }

    /// @notice Returns receipt's "snapshot root" field
    function snapshotRoot(Receipt receipt) internal pure returns (bytes32) {
        return receipt.unwrap().index({index_: OFFSET_SNAPSHOT_ROOT, bytes_: 32});
    }

    /// @notice Returns receipt's "state index" field
    function stateIndex(Receipt receipt) internal pure returns (uint8) {
        // Can be safely casted to uint8, since we index a single byte
        return uint8(receipt.unwrap().indexUint({index_: OFFSET_STATE_INDEX, bytes_: 1}));
    }

    /// @notice Returns receipt's "attestation notary" field
    function attNotary(Receipt receipt) internal pure returns (address) {
        return receipt.unwrap().indexAddress({index_: OFFSET_ATT_NOTARY});
    }

    /// @notice Returns receipt's "first executor" field
    function firstExecutor(Receipt receipt) internal pure returns (address) {
        return receipt.unwrap().indexAddress({index_: OFFSET_FIRST_EXECUTOR});
    }

    /// @notice Returns receipt's "final executor" field
    function finalExecutor(Receipt receipt) internal pure returns (address) {
        return receipt.unwrap().indexAddress({index_: OFFSET_FINAL_EXECUTOR});
    }
}

/// State is a memory view over a formatted state payload.
type State is uint256;

using StateLib for State global;

/// # State
/// State structure represents the state of Origin contract at some point of time.
/// - State is structured in a way to track the updates of the Origin Merkle Tree.
/// - State includes root of the Origin Merkle Tree, origin domain and some additional metadata.
/// ## Origin Merkle Tree
/// Hash of every sent message is inserted in the Origin Merkle Tree, which changes
/// the value of Origin Merkle Root (which is the root for the mentioned tree).
/// - Origin has a single Merkle Tree for all messages, regardless of their destination domain.
/// - This leads to Origin state being updated if and only if a message was sent in a block.
/// - Origin contract is a "source of truth" for states: a state is considered "valid" in its Origin,
/// if it matches the state of the Origin contract after the N-th (nonce) message was sent.
///
/// # Memory layout of State fields
///
/// | Position   | Field       | Type    | Bytes | Description                    |
/// | ---------- | ----------- | ------- | ----- | ------------------------------ |
/// | [000..032) | root        | bytes32 | 32    | Root of the Origin Merkle Tree |
/// | [032..036) | origin      | uint32  | 4     | Domain where Origin is located |
/// | [036..040) | nonce       | uint32  | 4     | Amount of sent messages        |
/// | [040..045) | blockNumber | uint40  | 5     | Block of last sent message     |
/// | [045..050) | timestamp   | uint40  | 5     | Time of last sent message      |
/// | [050..062) | gasData     | uint96  | 12    | Gas data for the chain         |
///
/// @dev State could be used to form a Snapshot to be signed by a Guard or a Notary.
library StateLib {
    using MemViewLib for bytes;

    /// @dev The variables below are not supposed to be used outside of the library directly.
    uint256 private constant OFFSET_ROOT = 0;
    uint256 private constant OFFSET_ORIGIN = 32;
    uint256 private constant OFFSET_NONCE = 36;
    uint256 private constant OFFSET_BLOCK_NUMBER = 40;
    uint256 private constant OFFSET_TIMESTAMP = 45;
    uint256 private constant OFFSET_GAS_DATA = 50;

    // ═══════════════════════════════════════════════════ STATE ═══════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted State payload with provided fields
     * @param root_         New merkle root
     * @param origin_       Domain of Origin's chain
     * @param nonce_        Nonce of the merkle root
     * @param blockNumber_  Block number when root was saved in Origin
     * @param timestamp_    Block timestamp when root was saved in Origin
     * @param gasData_      Gas data for the chain
     * @return Formatted state
     */
    function formatState(
        bytes32 root_,
        uint32 origin_,
        uint32 nonce_,
        uint40 blockNumber_,
        uint40 timestamp_,
        GasData gasData_
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(root_, origin_, nonce_, blockNumber_, timestamp_, gasData_);
    }

    /**
     * @notice Returns a State view over the given payload.
     * @dev Will revert if the payload is not a state.
     */
    function castToState(bytes memory payload) internal pure returns (State) {
        return castToState(payload.ref());
    }

    /**
     * @notice Casts a memory view to a State view.
     * @dev Will revert if the memory view is not over a state.
     */
    function castToState(MemView memView) internal pure returns (State) {
        if (!isState(memView)) revert UnformattedState();
        return State.wrap(MemView.unwrap(memView));
    }

    /// @notice Checks that a payload is a formatted State.
    function isState(MemView memView) internal pure returns (bool) {
        return memView.len() == STATE_LENGTH;
    }

    /// @notice Returns the hash of a State, that could be later signed by a Guard to signal
    /// that the state is invalid.
    function hashInvalid(State state) internal pure returns (bytes32) {
        // The final hash to sign is keccak(stateInvalidSalt, keccak(state))
        return state.unwrap().keccakSalted(STATE_INVALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(State state) internal pure returns (MemView) {
        return MemView.wrap(State.unwrap(state));
    }

    /// @notice Compares two State structures.
    function equals(State a, State b) internal pure returns (bool) {
        // Length of a State payload is fixed, so we just need to compare the hashes
        return a.unwrap().keccak() == b.unwrap().keccak();
    }

    // ═══════════════════════════════════════════════ STATE HASHING ═══════════════════════════════════════════════════

    /// @notice Returns the hash of the State.
    /// @dev We are using the Merkle Root of a tree with two leafs (see below) as state hash.
    function leaf(State state) internal pure returns (bytes32) {
        (bytes32 leftLeaf_, bytes32 rightLeaf_) = state.subLeafs();
        // Final hash is the parent of these leafs
        return keccak256(bytes.concat(leftLeaf_, rightLeaf_));
    }

    /// @notice Returns "sub-leafs" of the State. Hash of these "sub leafs" is going to be used
    /// as a "state leaf" in the "Snapshot Merkle Tree".
    /// This enables proving that leftLeaf = (root, origin) was a part of the "Snapshot Merkle Tree",
    /// by combining `rightLeaf` with the remainder of the "Snapshot Merkle Proof".
    function subLeafs(State state) internal pure returns (bytes32 leftLeaf_, bytes32 rightLeaf_) {
        MemView memView = state.unwrap();
        // Left leaf is (root, origin)
        leftLeaf_ = memView.prefix({len_: OFFSET_NONCE}).keccak();
        // Right leaf is (metadata), or (nonce, blockNumber, timestamp)
        rightLeaf_ = memView.sliceFrom({index_: OFFSET_NONCE}).keccak();
    }

    /// @notice Returns the left "sub-leaf" of the State.
    function leftLeaf(bytes32 root_, uint32 origin_) internal pure returns (bytes32) {
        // We use encodePacked here to simulate the State memory layout
        return keccak256(abi.encodePacked(root_, origin_));
    }

    /// @notice Returns the right "sub-leaf" of the State.
    function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, GasData gasData_)
        internal
        pure
        returns (bytes32)
    {
        // We use encodePacked here to simulate the State memory layout
        return keccak256(abi.encodePacked(nonce_, blockNumber_, timestamp_, gasData_));
    }

    // ═══════════════════════════════════════════════ STATE SLICING ═══════════════════════════════════════════════════

    /// @notice Returns a historical Merkle root from the Origin contract.
    function root(State state) internal pure returns (bytes32) {
        return state.unwrap().index({index_: OFFSET_ROOT, bytes_: 32});
    }

    /// @notice Returns domain of chain where the Origin contract is deployed.
    function origin(State state) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(state.unwrap().indexUint({index_: OFFSET_ORIGIN, bytes_: 4}));
    }

    /// @notice Returns nonce of Origin contract at the time, when `root` was the Merkle root.
    function nonce(State state) internal pure returns (uint32) {
        // Can be safely casted to uint32, since we index 4 bytes
        return uint32(state.unwrap().indexUint({index_: OFFSET_NONCE, bytes_: 4}));
    }

    /// @notice Returns a block number when `root` was saved in Origin.
    function blockNumber(State state) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(state.unwrap().indexUint({index_: OFFSET_BLOCK_NUMBER, bytes_: 5}));
    }

    /// @notice Returns a block timestamp when `root` was saved in Origin.
    /// @dev This is the timestamp according to the origin chain.
    function timestamp(State state) internal pure returns (uint40) {
        // Can be safely casted to uint40, since we index 5 bytes
        return uint40(state.unwrap().indexUint({index_: OFFSET_TIMESTAMP, bytes_: 5}));
    }

    /// @notice Returns gas data for the chain.
    function gasData(State state) internal pure returns (GasData) {
        return GasDataLib.wrapGasData(state.unwrap().indexUint({index_: OFFSET_GAS_DATA, bytes_: GAS_DATA_LENGTH}));
    }
}

/// Snapshot is a memory view over a formatted snapshot payload: a list of states.
type Snapshot is uint256;

using SnapshotLib for Snapshot global;

/// # Snapshot
/// Snapshot structure represents the state of multiple Origin contracts deployed on multiple chains.
/// In short, snapshot is a list of "State" structs. See State.sol for details about the "State" structs.
///
/// ## Snapshot usage
/// - Both Guards and Notaries are supposed to form snapshots and sign `snapshot.hash()` to verify its validity.
/// - Each Guard should be monitoring a set of Origin contracts chosen as they see fit.
///   - They are expected to form snapshots with Origin states for this set of chains,
///   sign and submit them to Summit contract.
/// - Notaries are expected to monitor the Summit contract for new snapshots submitted by the Guards.
///   - They should be forming their own snapshots using states from snapshots of any of the Guards.
///   - The states for the Notary snapshots don't have to come from the same Guard snapshot,
///   or don't even have to be submitted by the same Guard.
/// - With their signature, Notary effectively "notarizes" the work that some Guards have done in Summit contract.
///   - Notary signature on a snapshot doesn't only verify the validity of the Origins, but also serves as
///   a proof of liveliness for Guards monitoring these Origins.
///
/// ## Snapshot validity
/// - Snapshot is considered "valid" in Origin, if every state referring to that Origin is valid there.
/// - Snapshot is considered "globally valid", if it is "valid" in every Origin contract.
///
/// # Snapshot memory layout
///
/// | Position   | Field       | Type  | Bytes | Description                  |
/// | ---------- | ----------- | ----- | ----- | ---------------------------- |
/// | [000..050) | states[0]   | bytes | 50    | Origin State with index==0   |
/// | [050..100) | states[1]   | bytes | 50    | Origin State with index==1   |
/// | ...        | ...         | ...   | 50    | ...                          |
/// | [AAA..BBB) | states[N-1] | bytes | 50    | Origin State with index==N-1 |
///
/// @dev Snapshot could be signed by both Guards and Notaries and submitted to `Summit` in order to produce Attestations
/// that could be used in ExecutionHub for proving the messages coming from origin chains that the snapshot refers to.
library SnapshotLib {
    using MemViewLib for bytes;
    using StateLib for MemView;

    // ═════════════════════════════════════════════════ SNAPSHOT ══════════════════════════════════════════════════════

    /**
     * @notice Returns a formatted Snapshot payload using a list of States.
     * @param states    Arrays of State-typed memory views over Origin states
     * @return Formatted snapshot
     */
    function formatSnapshot(State[] memory states) internal view returns (bytes memory) {
        if (!_isValidAmount(states.length)) revert IncorrectStatesAmount();
        // First we unwrap State-typed views into untyped memory views
        uint256 length = states.length;
        MemView[] memory views = new MemView[](length);
        for (uint256 i = 0; i < length; ++i) {
            views[i] = states[i].unwrap();
        }
        // Finally, we join them in a single payload. This avoids doing unnecessary copies in the process.
        return MemViewLib.join(views);
    }

    /**
     * @notice Returns a Snapshot view over for the given payload.
     * @dev Will revert if the payload is not a snapshot payload.
     */
    function castToSnapshot(bytes memory payload) internal pure returns (Snapshot) {
        return castToSnapshot(payload.ref());
    }

    /**
     * @notice Casts a memory view to a Snapshot view.
     * @dev Will revert if the memory view is not over a snapshot payload.
     */
    function castToSnapshot(MemView memView) internal pure returns (Snapshot) {
        if (!isSnapshot(memView)) revert UnformattedSnapshot();
        return Snapshot.wrap(MemView.unwrap(memView));
    }

    /**
     * @notice Checks that a payload is a formatted Snapshot.
     */
    function isSnapshot(MemView memView) internal pure returns (bool) {
        // Snapshot needs to have exactly N * STATE_LENGTH bytes length
        // N needs to be in [1 .. SNAPSHOT_MAX_STATES] range
        uint256 length = memView.len();
        uint256 statesAmount_ = length / STATE_LENGTH;
        return statesAmount_ * STATE_LENGTH == length && _isValidAmount(statesAmount_);
    }

    /// @notice Returns the hash of a Snapshot, that could be later signed by an Agent  to signal
    /// that the snapshot is valid.
    function hashValid(Snapshot snapshot) internal pure returns (bytes32 hashedSnapshot) {
        // The final hash to sign is keccak(snapshotSalt, keccak(snapshot))
        return snapshot.unwrap().keccakSalted(SNAPSHOT_VALID_SALT);
    }

    /// @notice Convenience shortcut for unwrapping a view.
    function unwrap(Snapshot snapshot) internal pure returns (MemView) {
        return MemView.wrap(Snapshot.unwrap(snapshot));
    }

    // ═════════════════════════════════════════════ SNAPSHOT SLICING ══════════════════════════════════════════════════

    /// @notice Returns a state with a given index from the snapshot.
    function state(Snapshot snapshot, uint256 stateIndex) internal pure returns (State) {
        MemView memView = snapshot.unwrap();
        uint256 indexFrom = stateIndex * STATE_LENGTH;
        if (indexFrom >= memView.len()) revert IndexOutOfRange();
        return memView.slice({index_: indexFrom, len_: STATE_LENGTH}).castToState();
    }

    /// @notice Returns the amount of states in the snapshot.
    function statesAmount(Snapshot snapshot) internal pure returns (uint256) {
        // Each state occupies exactly `STATE_LENGTH` bytes
        return snapshot.unwrap().len() / STATE_LENGTH;
    }

    /// @notice Extracts the list of ChainGas structs from the snapshot.
    function snapGas(Snapshot snapshot) internal pure returns (ChainGas[] memory snapGas_) {
        uint256 statesAmount_ = snapshot.statesAmount();
        snapGas_ = new ChainGas[](statesAmount_);
        for (uint256 i = 0; i < statesAmount_; ++i) {
            State state_ = snapshot.state(i);
            snapGas_[i] = GasDataLib.encodeChainGas(state_.gasData(), state_.origin());
        }
    }

    // ═════════════════════════════════════════ SNAPSHOT ROOT CALCULATION ═════════════════════════════════════════════

    /// @notice Returns the root for the "Snapshot Merkle Tree" composed of state leafs from the snapshot.
    function calculateRoot(Snapshot snapshot) internal pure returns (bytes32) {
        uint256 statesAmount_ = snapshot.statesAmount();
        bytes32[] memory hashes = new bytes32[](statesAmount_);
        for (uint256 i = 0; i < statesAmount_; ++i) {
            // Each State has two sub-leafs, which are used as the "leafs" in "Snapshot Merkle Tree"
            // We save their parent in order to calculate the root for the whole tree later
            hashes[i] = snapshot.state(i).leaf();
        }
        // We are subtracting one here, as we already calculated the hashes
        // for the tree level above the "leaf level".
        MerkleMath.calculateRoot(hashes, SNAPSHOT_TREE_HEIGHT - 1);
        // hashes[0] now stores the value for the Merkle Root of the list
        return hashes[0];
    }

    /// @notice Reconstructs Snapshot merkle Root from State Merkle Data (root + origin domain)
    /// and proof of inclusion of State Merkle Data (aka State "left sub-leaf") in Snapshot Merkle Tree.
    /// > Reverts if any of these is true:
    /// > - State index is out of range.
    /// > - Snapshot Proof length exceeds Snapshot tree Height.
    /// @param originRoot    Root of Origin Merkle Tree
    /// @param domain        Domain of Origin chain
    /// @param snapProof     Proof of inclusion of State Merkle Data into Snapshot Merkle Tree
    /// @param stateIndex    Index of Origin State in the Snapshot
    function proofSnapRoot(bytes32 originRoot, uint32 domain, bytes32[] memory snapProof, uint8 stateIndex)
        internal
        pure
        returns (bytes32)
    {
        // Index of "leftLeaf" is twice the state position in the snapshot
        // This is because each state is represented by two leaves in the Snapshot Merkle Tree:
        // - leftLeaf is a hash of (originRoot, originDomain)
        // - rightLeaf is a hash of (nonce, blockNumber, timestamp, gasData)
        uint256 leftLeafIndex = uint256(stateIndex) << 1;
        // Check that "leftLeaf" index fits into Snapshot Merkle Tree
        if (leftLeafIndex >= (1 << SNAPSHOT_TREE_HEIGHT)) revert IndexOutOfRange();
        bytes32 leftLeaf = StateLib.leftLeaf(originRoot, domain);
        // Reconstruct snapshot root using proof of inclusion
        // This will revert if snapshot proof length exceeds Snapshot Tree Height
        return MerkleMath.proofRoot(leftLeafIndex, leftLeaf, snapProof, SNAPSHOT_TREE_HEIGHT);
    }

    // ══════════════════════════════════════════════ PRIVATE HELPERS ══════════════════════════════════════════════════

    /// @dev Checks if snapshot's states amount is valid.
    function _isValidAmount(uint256 statesAmount_) internal pure returns (bool) {
        // Need to have at least one state in a snapshot.
        // Also need to have no more than `SNAPSHOT_MAX_STATES` states in a snapshot.
        return statesAmount_ != 0 && statesAmount_ <= SNAPSHOT_MAX_STATES;
    }
}

library TypeCasts {
    // alignment preserving cast
    function addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    // alignment preserving cast
    function bytes32ToAddress(bytes32 buf) internal pure returns (address) {
        return address(uint160(uint256(buf)));
    }
}

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════

/// @notice A collection of events emitted by the ExecutionHub contract
abstract contract ExecutionHubEvents {
    /**
     * @notice Emitted when message is executed.
     * @param remoteDomain  Remote domain where message originated
     * @param messageHash   The keccak256 hash of the message that was executed
     * @param success       Whether the message was executed successfully
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash, bool success);

    /**
     * @notice Emitted when message tips are recorded.
     * @param messageHash   The keccak256 hash of the message that was executed
     * @param paddedTips    Padded encoded paid tips information
     */
    event TipsRecorded(bytes32 messageHash, uint256 paddedTips);
}

interface InterfaceInbox {
    // ══════════════════════════════════════════ SUBMIT AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Accepts a snapshot signed by a Guard or a Notary and passes it to Summit contract to save.
     * > Snapshot is a list of states for a set of Origin contracts residing on any of the chains.
     * - Guard-signed snapshots: all the states in the snapshot become available for Notary signing.
     * - Notary-signed snapshots: Snapshot Merkle Root is saved for valid snapshots, i.e.
     * snapshots which are only using states previously submitted by any of the Guards.
     * - Notary doesn't have to use states submitted by a single Guard in their snapshot.
     * - Notary could then proceed to sign the attestation for their submitted snapshot.
     * > Will revert if any of these is true:
     * > - Snapshot payload is not properly formatted.
     * > - Snapshot signer is not an active Agent.
     * > - Agent snapshot contains a state with a nonce smaller or equal then they have previously submitted.
     * > - Notary snapshot contains a state that hasn't been previously submitted by any of the Guards.
     * > - Note: Agent will NOT be slashed for submitting such a snapshot.
     * @dev Notary will need to provide both `agentRoot` and `snapGas` when submitting an attestation on
     * the remote chain (the attestation contains only their merged hash). These are returned by this function,
     * and could be also obtained by calling `getAttestation(nonce)` or `getLatestNotaryAttestation(notary)`.
     * @param snapPayload       Raw payload with snapshot data
     * @param snapSignature     Agent signature for the snapshot
     * @return attPayload       Raw payload with data for attestation derived from Notary snapshot.
     *                          Empty payload, if a Guard snapshot was submitted.
     * @return agentRoot        Current root of the Agent Merkle Tree (zero, if a Guard snapshot was submitted)
     * @return snapGas          Gas data for each chain in the snapshot
     *                          Empty list, if a Guard snapshot was submitted.
     */
    function submitSnapshot(bytes memory snapPayload, bytes memory snapSignature)
        external
        returns (bytes memory attPayload, bytes32 agentRoot, uint256[] memory snapGas);

    /**
     * @notice Accepts a receipt signed by a Notary and passes it to Summit contract to save.
     * > Receipt is a statement about message execution status on the remote chain.
     * - This will distribute the message tips across the off-chain actors once the receipt optimistic period is over.
     * > Will revert if any of these is true:
     * > - Receipt payload is not properly formatted.
     * > - Receipt signer is not an active Notary.
     * > - Receipt signer is in Dispute.
     * > - Receipt's snapshot root is unknown.
     * > - Provided tips could not be proven against the message hash.
     * @param rcptPayload       Raw payload with receipt data
     * @param rcptSignature     Notary signature for the receipt
     * @param paddedTips        Tips for the message execution
     * @param headerHash        Hash of the message header
     * @param bodyHash          Hash of the message body excluding the tips
     * @return wasAccepted      Whether the receipt was accepted
     */
    function submitReceipt(
        bytes memory rcptPayload,
        bytes memory rcptSignature,
        uint256 paddedTips,
        bytes32 headerHash,
        bytes32 bodyHash
    ) external returns (bool wasAccepted);

    /**
     * @notice Accepts a Guard's receipt report signature, as well as Notary signature
     * for the reported Receipt.
     * > ReceiptReport is a Guard statement saying "Reported receipt is invalid".
     * - This results in an opened Dispute between the Guard and the Notary.
     * - Note: Guard could (but doesn't have to) form a ReceiptReport and use receipt signature from
     * `verifyReceipt()` successful call that led to Notary being slashed in Summit on Synapse Chain.
     * > Will revert if any of these is true:
     * > - Receipt payload is not properly formatted.
     * > - Receipt Report signer is not an active Guard.
     * > - Receipt signer is not an active Notary.
     * @param rcptPayload       Raw payload with Receipt data that Guard reports as invalid
     * @param rcptSignature     Notary signature for the reported receipt
     * @param rrSignature       Guard signature for the report
     * @return wasAccepted      Whether the Report was accepted (resulting in Dispute between the agents)
     */
    function submitReceiptReport(bytes memory rcptPayload, bytes memory rcptSignature, bytes memory rrSignature)
        external
        returns (bool wasAccepted);

    /**
     * @notice Passes the message execution receipt from Destination to the Summit contract to save.
     * > Will revert if any of these is true:
     * > - Called by anyone other than Destination.
     * @dev If a receipt is not accepted, any of the Notaries can submit it later using `submitReceipt`.
     * @param attNotaryIndex    Index of the Notary who signed the attestation
     * @param attNonce          Nonce of the attestation used for proving the executed message
     * @param paddedTips        Tips for the message execution
     * @param rcptPayload       Raw payload with message execution receipt
     * @return wasAccepted      Whether the receipt was accepted
     */
    function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes memory rcptPayload)
        external
        returns (bool wasAccepted);

    // ══════════════════════════════════════════ VERIFY AGENT STATEMENTS ══════════════════════════════════════════════

    /**
     * @notice Verifies an attestation signed by a Notary.
     *  - Does nothing, if the attestation is valid (was submitted by this Notary as a snapshot).
     *  - Slashes the Notary, if the attestation is invalid.
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation signer is not an active Notary.
     * @param attPayload        Raw payload with Attestation data
     * @param attSignature      Notary signature for the attestation
     * @return isValidAttestation   Whether the provided attestation is valid.
     *                              Notary is slashed, if return value is FALSE.
     */
    function verifyAttestation(bytes memory attPayload, bytes memory attSignature)
        external
        returns (bool isValidAttestation);

    /**
     * @notice Verifies a Guard's attestation report signature.
     *  - Does nothing, if the report is valid (if the reported attestation is invalid).
     *  - Slashes the Guard, if the report is invalid (if the reported attestation is valid).
     * > Will revert if any of these is true:
     * > - Attestation payload is not properly formatted.
     * > - Attestation Report signer is not an active Guard.
     * @param attPayload        Raw payload with Attestation data that Guard reports as invalid
     * @param arSignature       Guard signature for the report
     * @return isValidReport    Whether the provided report is valid.
     *                          Guard is slashed, if return value is FALSE.
     */
    function verifyAttestationReport(bytes memory attPayload, bytes memory arSignature)
        external
        returns (bool isValidReport);
}

interface IExecutionHub {
    /**
     * @notice Attempts to prove inclusion of message into one of Snapshot Merkle Trees,
     * previously submitted to this contract in a form of a signed Attestation.
     * Proven message is immediately executed by passing its contents to the specified recipient.
     * @dev Will revert if any of these is true:
     *  - Message is not meant to be executed on this chain
     *  - Message was sent from this chain
     *  - Message payload is not properly formatted.
     *  - Snapshot root (reconstructed from message hash and proofs) is unknown
     *  - Snapshot root is known, but was submitted by an inactive Notary
     *  - Snapshot root is known, but optimistic period for a message hasn't passed
     *  - Provided gas limit is lower than the one requested in the message
     *  - Recipient doesn't implement a `handle` method (refer to IMessageRecipient.sol)
     *  - Recipient reverted upon receiving a message
     * Note: refer to libs/memory/State.sol for details about Origin State's sub-leafs.
     * @param msgPayload    Raw payload with a formatted message to execute
     * @param originProof   Proof of inclusion of message in the Origin Merkle Tree
     * @param snapProof     Proof of inclusion of Origin State's Left Leaf into Snapshot Merkle Tree
     * @param stateIndex    Index of Origin State in the Snapshot
     * @param gasLimit      Gas limit for message execution
     */
    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint8 stateIndex,
        uint64 gasLimit
    ) external;

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /**
     * @notice Returns attestation nonce for a given snapshot root.
     * @dev Will return 0 if the root is unknown.
     */
    function getAttestationNonce(bytes32 snapRoot) external view returns (uint32 attNonce);

    /**
     * @notice Checks the validity of the unsigned message receipt.
     * @dev Will revert if any of these is true:
     *  - Receipt payload is not properly formatted.
     *  - Receipt signer is not an active Notary.
     *  - Receipt destination chain does not refer to this chain.
     * @param rcptPayload       Raw payload with Receipt data
     * @return isValid          Whether the requested receipt is valid.
     */
    function isValidReceipt(bytes memory rcptPayload) external view returns (bool isValid);

    /**
     * @notice Returns message execution status: None/Failed/Success.
     * @param messageHash       Hash of the message payload
     * @return status           Message execution status
     */
    function messageStatus(bytes32 messageHash) external view returns (MessageStatus status);

    /**
     * @notice Returns a formatted payload with the message receipt.
     * @dev Notaries could derive the tips, and the tips proof using the message payload, and submit
     * the signed receipt with the proof of tips to `Summit` in order to initiate tips distribution.
     * @param messageHash       Hash of the message payload
     * @return data             Formatted payload with the message execution receipt
     */
    function messageReceipt(bytes32 messageHash) external view returns (bytes memory data);
}

interface IMessageRecipient {
    /**
     * @notice Message recipient needs to implement this function in order to
     * receive cross-chain messages.
     * @dev Message recipient needs to ensure that merkle proof for the message
     * is at least as old as the optimistic period that the recipient is using.
     * Note: as this point it is checked that the "message optimistic period" has passed,
     * however the period value itself could be anything, and thus could differ from the one
     * that the recipient would like to enforce.
     * @param origin            Domain where message originated
     * @param nonce             Message nonce on the origin domain
     * @param sender            Sender address on origin chain
     * @param proofMaturity     Message's merkle proof age in seconds
     * @param version           Message version specified by sender
     * @param content           Raw bytes content of message
     */
    function receiveBaseMessage(
        uint32 origin,
        uint32 nonce,
        bytes32 sender,
        uint256 proofMaturity,
        uint32 version,
        bytes memory content
    ) external payable;
}

// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════

// OpenZeppelin Contracts (last updated v4.9.0) (utils/Address.sol)

/**
 * @dev Collection of functions related to the address type
 */
library Address {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     *
     * Furthermore, `isContract` will also return true if the target contract within
     * the same transaction is already scheduled for destruction by `SELFDESTRUCT`,
     * which only has an effect at the end of a transaction.
     * ====
     *
     * [IMPORTANT]
     * ====
     * You shouldn't rely on `isContract` to protect against flash loan attacks!
     *
     * Preventing calls from contracts is highly discouraged. It breaks composability, breaks support for smart wallets
     * like Gnosis Safe, and does not provide security since it can be circumvented by calling from a contract
     * constructor.
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize/address.code.length, which returns 0
        // for contracts in construction, since the code is only stored at the end
        // of the constructor execution.

        return account.code.length > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://consensys.net/diligence/blog/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.8.0/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain `call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(address target, bytes memory data, uint256 value) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value,
        string memory errorMessage
    ) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionDelegateCall(target, data, "Address: low-level delegate call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.4._
     */
    function functionDelegateCall(
        address target,
        bytes memory data,
        string memory errorMessage
    ) internal returns (bytes memory) {
        (bool success, bytes memory returndata) = target.delegatecall(data);
        return verifyCallResultFromTarget(target, success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verify that a low level call to smart-contract was successful, and revert (either by bubbling
     * the revert reason or using the provided one) in case of unsuccessful call or if target was not a contract.
     *
     * _Available since v4.8._
     */
    function verifyCallResultFromTarget(
        address target,
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        if (success) {
            if (returndata.length == 0) {
                // only check isContract if the call was successful and the return data is empty
                // otherwise we already know that it was a contract
                require(isContract(target), "Address: call to non-contract");
            }
            return returndata;
        } else {
            _revert(returndata, errorMessage);
        }
    }

    /**
     * @dev Tool to verify that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason or using the provided one.
     *
     * _Available since v4.3._
     */
    function verifyCallResult(
        bool success,
        bytes memory returndata,
        string memory errorMessage
    ) internal pure returns (bytes memory) {
        if (success) {
            return returndata;
        } else {
            _revert(returndata, errorMessage);
        }
    }

    function _revert(bytes memory returndata, string memory errorMessage) private pure {
        // Look for revert reason and bubble it up if present
        if (returndata.length > 0) {
            // The easiest way to bubble the revert reason is using memory via assembly
            /// @solidity memory-safe-assembly
            assembly {
                let returndata_size := mload(returndata)
                revert(add(32, returndata), returndata_size)
            }
        } else {
            revert(errorMessage);
        }
    }
}

// OpenZeppelin Contracts (last updated v4.9.0) (security/ReentrancyGuard.sol)

/**
 * @dev Contract module that helps prevent reentrant calls to a function.
 *
 * Inheriting from `ReentrancyGuard` will make the {nonReentrant} modifier
 * available, which can be applied to functions to make sure there are no nested
 * (reentrant) calls to them.
 *
 * Note that because there is a single `nonReentrant` guard, functions marked as
 * `nonReentrant` may not call one another. This can be worked around by making
 * those functions `private`, and then adding `external` `nonReentrant` entry
 * points to them.
 *
 * TIP: If you would like to learn more about reentrancy and alternative ways
 * to protect against it, check out our blog post
 * https://blog.openzeppelin.com/reentrancy-after-istanbul/[Reentrancy After Istanbul].
 */
abstract contract ReentrancyGuardUpgradeable is Initializable {
    // Booleans are more expensive than uint256 or any type that takes up a full
    // word because each write operation emits an extra SLOAD to first read the
    // slot's contents, replace the bits taken up by the boolean, and then write
    // back. This is the compiler's defense against contract upgrades and
    // pointer aliasing, and it cannot be disabled.

    // The values being non-zero value makes deployment a bit more expensive,
    // but in exchange the refund on every call to nonReentrant will be lower in
    // amount. Since refunds are capped to a percentage of the total
    // transaction's gas, it is best to keep them low in cases like this one, to
    // increase the likelihood of the full refund coming into effect.
    uint256 private constant _NOT_ENTERED = 1;
    uint256 private constant _ENTERED = 2;

    uint256 private _status;

    function __ReentrancyGuard_init() internal onlyInitializing {
        __ReentrancyGuard_init_unchained();
    }

    function __ReentrancyGuard_init_unchained() internal onlyInitializing {
        _status = _NOT_ENTERED;
    }

    /**
     * @dev Prevents a contract from calling itself, directly or indirectly.
     * Calling a `nonReentrant` function from another `nonReentrant`
     * function is not supported. It is possible to prevent this from happening
     * by making the `nonReentrant` function external, and making it call a
     * `private` function that does the actual work.
     */
    modifier nonReentrant() {
        _nonReentrantBefore();
        _;
        _nonReentrantAfter();
    }

    function _nonReentrantBefore() private {
        // On the first call to nonReentrant, _status will be _NOT_ENTERED
        require(_status != _ENTERED, "ReentrancyGuard: reentrant call");

        // Any calls to nonReentrant after this point will fail
        _status = _ENTERED;
    }

    function _nonReentrantAfter() private {
        // By storing the original value once again, a refund is triggered (see
        // https://eips.ethereum.org/EIPS/eip-2200)
        _status = _NOT_ENTERED;
    }

    /**
     * @dev Returns true if the reentrancy guard is currently set to "entered", which indicates there is a
     * `nonReentrant` function in the call stack.
     */
    function _reentrancyGuardEntered() internal view returns (bool) {
        return _status == _ENTERED;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}

/// @notice `ExecutionHub` is a parent contract for `Destination`. It is responsible for the following:
/// - Executing the messages that are proven against the saved Snapshot Merkle Roots.
/// - Base messages are forwarded to the specified message recipient, ensuring that the original
///   execution request is fulfilled correctly.
/// - Manager messages are forwarded to the local `AgentManager` contract.
/// - Keeping track of the saved Snapshot Merkle Roots (which are accepted in `Destination`).
/// - Keeping track of message execution Receipts, as well as verify their validity.
abstract contract ExecutionHub is AgentSecured, ReentrancyGuardUpgradeable, ExecutionHubEvents, IExecutionHub {
    using Address for address;
    using BaseMessageLib for MemView;
    using ByteString for MemView;
    using MessageLib for bytes;
    using ReceiptLib for bytes;
    using SafeCall for address;
    using SafeCast for uint256;
    using TypeCasts for bytes32;

    /// @notice Struct representing stored data for the snapshot root
    /// @param notaryIndex  Index of Notary who submitted the statement with the snapshot root
    /// @param attNonce     Nonce of the attestation for this snapshot root
    /// @param attBN        Summit block number of the attestation for this snapshot root
    /// @param attTS        Summit timestamp of the attestation for this snapshot root
    /// @param index        Index of snapshot root in `_roots`
    /// @param submittedAt  Timestamp when the statement with the snapshot root was submitted
    /// @param notaryV      V-value from the Notary signature for the attestation
    // TODO: tight pack this
    struct SnapRootData {
        uint32 notaryIndex;
        uint32 attNonce;
        uint40 attBN;
        uint40 attTS;
        uint32 index;
        uint40 submittedAt;
        uint256 sigIndex;
    }

    /// @notice Struct representing stored receipt data for the message in Execution Hub.
    /// @param origin       Domain where message originated
    /// @param rootIndex    Index of snapshot root used for proving the message
    /// @param stateIndex   Index of state used for the snapshot proof
    /// @param executor     Executor who successfully executed the message
    struct ReceiptData {
        uint32 origin;
        uint32 rootIndex;
        uint8 stateIndex;
        address executor;
    }
    // TODO: include nonce?
    // 24 bits available for tight packing

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @notice (messageHash => status)
    /// @dev Messages coming from different origins will always have a different hash
    /// as origin domain is encoded into the formatted message.
    /// Thus we can use hash as a key instead of an (origin, hash) tuple.
    mapping(bytes32 => ReceiptData) private _receiptData;

    /// @notice First executor who made a valid attempt of executing a message.
    /// Note: stored only for messages that had Failed status at some point of time
    mapping(bytes32 => address) private _firstExecutor;

    /// @dev All saved snapshot roots
    bytes32[] internal _roots;

    /// @dev Tracks data for all saved snapshot roots
    mapping(bytes32 => SnapRootData) internal _rootData;

    /// @dev gap for upgrade safety
    uint256[46] private __GAP; // solhint-disable-line var-name-mixedcase

    // ═════════════════════════════════════════════ MESSAGE EXECUTION ═════════════════════════════════════════════════

    /// @inheritdoc IExecutionHub
    function execute(
        bytes memory msgPayload,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint8 stateIndex,
        uint64 gasLimit
    ) external nonReentrant {
        // This will revert if payload is not a formatted message payload
        Message message = msgPayload.castToMessage();
        Header header = message.header();
        bytes32 msgLeaf = message.leaf();
        // Ensure message was meant for this domain
        if (header.destination() != localDomain) revert IncorrectDestinationDomain();
        // Ensure message was not sent from this domain
        if (header.origin() == localDomain) revert IncorrectOriginDomain();
        // Check that message has not been executed before
        ReceiptData memory rcptData = _receiptData[msgLeaf];
        if (rcptData.executor != address(0)) revert AlreadyExecuted();
        // Check proofs validity
        SnapRootData memory rootData = _proveAttestation(header, msgLeaf, originProof, snapProof, stateIndex);
        // Check if optimistic period has passed
        uint256 proofMaturity = block.timestamp - rootData.submittedAt;
        if (proofMaturity < header.optimisticPeriod()) revert MessageOptimisticPeriod();
        uint256 paddedTips;
        bool success;
        // Only Base/Manager message flags exist
        if (header.flag() == MessageFlag.Base) {
            // This will revert if message body is not a formatted BaseMessage payload
            BaseMessage baseMessage = message.body().castToBaseMessage();
            success = _executeBaseMessage(header, proofMaturity, gasLimit, baseMessage);
            paddedTips = Tips.unwrap(baseMessage.tips());
        } else {
            // gasLimit is ignored when executing manager messages
            success = _executeManagerMessage(header, proofMaturity, message.body());
        }
        if (rcptData.origin == 0) {
            // This is the first valid attempt to execute the message => save origin and snapshot proof
            rcptData.origin = header.origin();
            rcptData.rootIndex = rootData.index;
            rcptData.stateIndex = stateIndex;
            if (success) {
                // This is the successful attempt to execute the message => save the executor
                rcptData.executor = msg.sender;
            } else {
                // Save as the "first executor", if execution failed
                _firstExecutor[msgLeaf] = msg.sender;
            }
            _receiptData[msgLeaf] = rcptData;
        } else {
            if (!success) revert AlreadyFailed();
            // There has been a failed attempt to execute the message before => don't touch origin and snapshot root
            // This is the successful attempt to execute the message => save the executor
            rcptData.executor = msg.sender;
            _receiptData[msgLeaf] = rcptData;
        }
        emit Executed(header.origin(), msgLeaf, success);
        if (!_passReceipt(rootData.notaryIndex, rootData.attNonce, msgLeaf, paddedTips, rcptData)) {
            // Emit event with the recorded tips so that Notaries could form a receipt to submit to Summit
            emit TipsRecorded(msgLeaf, paddedTips);
        }
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc IExecutionHub
    function getAttestationNonce(bytes32 snapRoot) external view returns (uint32 attNonce) {
        return _rootData[snapRoot].attNonce;
    }

    /// @inheritdoc IExecutionHub
    function isValidReceipt(bytes memory rcptPayload) external view returns (bool isValid) {
        // This will revert if payload is not a receipt
        // This will revert if receipt refers to another domain
        return _isValidReceipt(rcptPayload.castToReceipt());
    }

    /// @inheritdoc IExecutionHub
    function messageStatus(bytes32 messageHash) external view returns (MessageStatus status) {
        ReceiptData memory rcptData = _receiptData[messageHash];
        if (rcptData.executor != address(0)) {
            return MessageStatus.Success;
        } else if (_firstExecutor[messageHash] != address(0)) {
            return MessageStatus.Failed;
        } else {
            return MessageStatus.None;
        }
    }

    /// @inheritdoc IExecutionHub
    function messageReceipt(bytes32 messageHash) external view returns (bytes memory rcptPayload) {
        ReceiptData memory rcptData = _receiptData[messageHash];
        // Return empty payload if there has been no attempt to execute the message
        if (rcptData.origin == 0) return "";
        return _messageReceipt(messageHash, rcptData);
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Passes message content to recipient that conforms to IMessageRecipient interface.
    function _executeBaseMessage(Header header, uint256 proofMaturity, uint64 gasLimit, BaseMessage baseMessage)
        internal
        returns (bool)
    {
        // Check that gas limit covers the one requested by the sender.
        // We let the executor specify gas limit higher than requested to guarantee the execution of
        // messages with gas limit set too low.
        Request request = baseMessage.request();
        if (gasLimit < request.gasLimit()) revert GasLimitTooLow();
        // TODO: check that the discarded bits are empty
        address recipient = baseMessage.recipient().bytes32ToAddress();
        // Forward message content to the recipient, and limit the amount of forwarded gas
        if (gasleft() <= gasLimit) revert GasSuppliedTooLow();
        // receiveBaseMessage(origin, nonce, sender, proofMaturity, version, content)
        bytes memory payload = abi.encodeCall(
            IMessageRecipient.receiveBaseMessage,
            (
                header.origin(),
                header.nonce(),
                baseMessage.sender(),
                proofMaturity,
                request.version(),
                baseMessage.content().clone()
            )
        );
        // Pass the base message to the recipient, return the success status of the call
        return recipient.safeCall({gasLimit: gasLimit, msgValue: 0, payload: payload});
    }

    /// @dev Uses message body for a call to AgentManager, and checks the returned magic value to ensure that
    /// only "remoteX" functions could be called this way.
    function _executeManagerMessage(Header header, uint256 proofMaturity, MemView body) internal returns (bool) {
        // TODO: introduce incentives for executing Manager Messages?
        CallData callData = body.castToCallData();
        // Add the (origin, proofMaturity) values to the calldata
        bytes memory payload = callData.addPrefix(abi.encode(header.origin(), proofMaturity));
        // functionCall() calls AgentManager and bubbles the revert from the external call
        bytes memory magicValue = address(agentManager).functionCall(payload);
        // We check the returned value here to ensure that only "remoteX" functions could be called this way.
        // This is done to prevent an attack by a malicious Notary trying to force Destination to call an arbitrary
        // function in a local AgentManager. Any other function will not return the required selector,
        // while the "remoteX" functions will perform the proofMaturity check that will make impossible to
        // submit an attestation and execute a malicious Manager Message immediately, preventing this attack vector.
        if (magicValue.length != 32 || bytes32(magicValue) != callData.callSelector()) revert IncorrectMagicValue();
        return true;
    }

    /// @dev Passes the message receipt to the Inbox contract, if it is deployed on Synapse Chain.
    /// This ensures that the message receipts for the messages executed on Synapse Chain are passed to Summit
    /// without a Notary having to sign them.
    function _passReceipt(
        uint32 attNotaryIndex,
        uint32 attNonce,
        bytes32 messageHash,
        uint256 paddedTips,
        ReceiptData memory rcptData
    ) internal returns (bool) {
        // Do nothing if contract is not deployed on Synapse Chain
        if (localDomain != synapseDomain) return false;
        // Do nothing for messages with no tips (TODO: introduce incentives for manager messages?)
        if (paddedTips == 0) return false;
        return InterfaceInbox(inbox).passReceipt({
            attNotaryIndex: attNotaryIndex,
            attNonce: attNonce,
            paddedTips: paddedTips,
            rcptPayload: _messageReceipt(messageHash, rcptData)
        });
    }

    /// @dev Saves a snapshot root with the attestation data provided by a Notary.
    /// It is assumed that the Notary signature has been checked outside of this contract.
    function _saveAttestation(Attestation att, uint32 notaryIndex, uint256 sigIndex) internal {
        bytes32 root = att.snapRoot();
        if (_rootData[root].submittedAt != 0) revert DuplicatedSnapshotRoot();
        // TODO: consider using more than 32 bits for the root index
        _rootData[root] = SnapRootData({
            notaryIndex: notaryIndex,
            attNonce: att.nonce(),
            attBN: att.blockNumber(),
            attTS: att.timestamp(),
            index: _roots.length.toUint32(),
            submittedAt: ChainContext.blockTimestamp(),
            sigIndex: sigIndex
        });
        _roots.push(root);
    }

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Checks if receipt body matches the saved data for the referenced message.
    /// Reverts if destination domain doesn't match the local domain.
    function _isValidReceipt(Receipt rcpt) internal view returns (bool) {
        // Check if receipt refers to this chain
        if (rcpt.destination() != localDomain) revert IncorrectDestinationDomain();
        bytes32 messageHash = rcpt.messageHash();
        ReceiptData memory rcptData = _receiptData[messageHash];
        // Check if there has been a single attempt to execute the message
        if (rcptData.origin == 0) return false;
        // Check that origin and state index fields match
        if (rcpt.origin() != rcptData.origin || rcpt.stateIndex() != rcptData.stateIndex) return false;
        // Check that snapshot root and notary who submitted it match in the Receipt
        bytes32 snapRoot = rcpt.snapshotRoot();
        (address attNotary,) = _getAgent(_rootData[snapRoot].notaryIndex);
        if (snapRoot != _roots[rcptData.rootIndex] || rcpt.attNotary() != attNotary) return false;
        // Check if message was executed from the first attempt
        address firstExecutor = _firstExecutor[messageHash];
        if (firstExecutor == address(0)) {
            // Both first and final executors are saved in receipt data
            return rcpt.firstExecutor() == rcptData.executor && rcpt.finalExecutor() == rcptData.executor;
        } else {
            // Message was Failed at some point of time, so both receipts are valid:
            // "Failed": finalExecutor is ZERO
            // "Success": finalExecutor matches executor from saved receipt data
            address finalExecutor = rcpt.finalExecutor();
            return rcpt.firstExecutor() == firstExecutor
                && (finalExecutor == address(0) || finalExecutor == rcptData.executor);
        }
    }

    /**
     * @notice Attempts to prove the validity of the cross-chain message.
     * First, the origin Merkle Root is reconstructed using the origin proof.
     * Then the origin state's "left leaf" is reconstructed using the origin domain.
     * After that the snapshot Merkle Root is reconstructed using the snapshot proof.
     * The snapshot root needs to have been submitted by an undisputed Notary.
     * @dev Reverts if any of the checks fail.
     * @param header        Memory view over the message header
     * @param msgLeaf       Message Leaf that was inserted in the Origin Merkle Tree
     * @param originProof   Proof of inclusion of Message Leaf in the Origin Merkle Tree
     * @param snapProof     Proof of inclusion of Origin State Left Leaf into Snapshot Merkle Tree
     * @param stateIndex    Index of Origin State in the Snapshot
     * @return rootData     Data for the derived snapshot root
     */
    function _proveAttestation(
        Header header,
        bytes32 msgLeaf,
        bytes32[] calldata originProof,
        bytes32[] calldata snapProof,
        uint8 stateIndex
    ) internal view returns (SnapRootData memory rootData) {
        // Reconstruct Origin Merkle Root using the origin proof
        // Message index in the tree is (nonce - 1), as nonce starts from 1
        // This will revert if origin proof length exceeds Origin Tree height
        bytes32 originRoot = MerkleMath.proofRoot(header.nonce() - 1, msgLeaf, originProof, ORIGIN_TREE_HEIGHT);
        // Reconstruct Snapshot Merkle Root using the snapshot proof
        // This will revert if:
        //  - State index is out of range.
        //  - Snapshot Proof length exceeds Snapshot tree Height.
        bytes32 snapshotRoot = SnapshotLib.proofSnapRoot(originRoot, header.origin(), snapProof, stateIndex);
        // Fetch the attestation data for the snapshot root
        rootData = _rootData[snapshotRoot];
        // Check if snapshot root has been submitted
        if (rootData.submittedAt == 0) revert IncorrectSnapshotRoot();
        // Check that Notary who submitted the attestation is not in dispute
        if (_notaryDisputeExists(rootData.notaryIndex)) revert NotaryInDispute();
        // Check that Notary who submitted the attestation isn't in post-dispute timeout
        if (_notaryDisputeTimeout(rootData.notaryIndex)) revert DisputeTimeoutNotOver();
    }

    /// @dev Formats the message execution receipt payload for the given hash and receipt data.
    function _messageReceipt(bytes32 messageHash, ReceiptData memory rcptData)
        internal
        view
        returns (bytes memory rcptPayload)
    {
        // Determine the first executor who tried to execute the message
        address firstExecutor = _firstExecutor[messageHash];
        if (firstExecutor == address(0)) firstExecutor = rcptData.executor;
        // Determine the snapshot root that was used for proving the message
        bytes32 snapRoot = _roots[rcptData.rootIndex];
        (address attNotary,) = _getAgent(_rootData[snapRoot].notaryIndex);
        // ExecutionHub does not store the tips,
        // the Notary will have to derive the proof of tips from the message payload.
        return ReceiptLib.formatReceipt({
            origin_: rcptData.origin,
            destination_: localDomain,
            messageHash_: messageHash,
            snapshotRoot_: snapRoot,
            stateIndex_: rcptData.stateIndex,
            attNotary_: attNotary,
            firstExecutor_: firstExecutor,
            finalExecutor_: rcptData.executor
        });
    }
}

/// @notice `Destination` contract is used for receiving messages from other chains. It relies on
/// Notary-signed statements to get the truthful states of the remote chains. These states are then
/// used to verify the validity of the messages sent from the remote chains.
/// `Destination` is responsible for the following:
/// - Accepting the Attestations from the local Inbox contract.
/// - Using these Attestations to execute the messages (see parent `ExecutionHub`).
/// - Passing the Agent Merkle Roots from the Attestations to the local LightManager contract,
///   if deployed on a non-Synapse chain.
/// - Keeping track of the remote domains GasData submitted by Notaries, that could be later consumed
///   by the local `GasOracle` contract.
contract Destination is ExecutionHub, DestinationEvents, InterfaceDestination {
    using AttestationLib for bytes;
    using ByteString for bytes;

    // TODO: this could be further optimized in terms of storage
    struct StoredAttData {
        bytes32 agentRoot;
        bytes32 dataHash;
    }

    struct StoredGasData {
        GasData gasData;
        uint32 notaryIndex;
        uint40 submittedAt;
    }

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev Invariant: this is either current LightManager root,
    /// or the pending root to be passed to LightManager once its optimistic period is over.
    bytes32 internal _nextAgentRoot;

    /// @inheritdoc InterfaceDestination
    DestinationStatus public destStatus;

    /// @inheritdoc InterfaceDestination
    mapping(uint32 => uint32) public lastAttestationNonce;

    /// @dev Stored lookup data for all accepted Notary Attestations
    StoredAttData[] internal _storedAttestations;

    /// @dev Remote domains GasData submitted by Notaries
    mapping(uint32 => StoredGasData) internal _storedGasData;

    // ═════════════════════════════════════════ CONSTRUCTOR & INITIALIZER ═════════════════════════════════════════════

    constructor(uint32 synapseDomain_, address agentManager_, address inbox_)
        AgentSecured("0.0.3", synapseDomain_, agentManager_, inbox_)
    {} // solhint-disable-line no-empty-blocks

    /// @notice Initializes Destination contract:
    /// - msg.sender is set as contract owner
    function initialize(bytes32 agentRoot) external initializer {
        // Initialize Ownable: msg.sender is set as "owner"
        __Ownable2Step_init();
        // Initialize ReeentrancyGuard
        __ReentrancyGuard_init();
        // Set Agent Merkle Root in Light Manager
        if (localDomain != synapseDomain) {
            _nextAgentRoot = agentRoot;
            InterfaceLightManager(address(agentManager)).setAgentRoot(agentRoot);
            destStatus.agentRootTime = ChainContext.blockTimestamp();
        }
        // No need to do anything on Synapse Chain, as the agent root is set in BondingManager
    }

    // ═════════════════════════════════════════════ ACCEPT STATEMENTS ═════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    function acceptAttestation(
        uint32 notaryIndex,
        uint256 sigIndex,
        bytes memory attPayload,
        bytes32 agentRoot,
        ChainGas[] memory snapGas
    ) external onlyInbox returns (bool wasAccepted) {
        // Check that we can trust the Notary data: they are not in dispute, and the dispute timeout is over (if any)
        if (_notaryDisputeExists(notaryIndex)) revert NotaryInDispute();
        if (_notaryDisputeTimeout(notaryIndex)) revert DisputeTimeoutNotOver();
        // First, try passing current agent merkle root
        // This will revert if payload is not an attestation
        Attestation att = attPayload.castToAttestation();
        // Check that this Notary hasn't used a more fresh nonce
        uint32 attNonce = att.nonce();
        if (attNonce <= lastAttestationNonce[notaryIndex]) revert OutdatedNonce();
        lastAttestationNonce[notaryIndex] = attNonce;
        // This will revert if snapshot root has been previously submitted
        _saveAttestation(att, notaryIndex, sigIndex);
        _storedAttestations.push(StoredAttData({agentRoot: agentRoot, dataHash: att.dataHash()}));
        // Save Agent Root if required, and update the Destination's Status
        bool rootPending = passAgentRoot();
        destStatus = _saveAgentRoot(rootPending, agentRoot, notaryIndex);
        _saveGasData(snapGas, notaryIndex);
        return true;
    }

    // ═══════════════════════════════════════════ AGENT ROOT QUARANTINE ═══════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    function passAgentRoot() public returns (bool rootPending) {
        // Agent root is not passed on Synapse Chain, as it could be accessed via BondingManager
        if (localDomain == synapseDomain) return false;
        bytes32 oldRoot = IAgentManager(agentManager).agentRoot();
        bytes32 newRoot = _nextAgentRoot;
        // Check if agent root differs from the current one in LightManager
        if (oldRoot == newRoot) return false;
        DestinationStatus memory status = destStatus;
        // Invariant: Notary who supplied `newRoot` was registered as active against `oldRoot`
        // So we just need to check the Dispute status of the Notary
        if (_notaryDisputeExists(status.notaryIndex)) {
            // Remove the pending agent merkle root, as its signer is in dispute
            _nextAgentRoot = oldRoot;
            return false;
        }
        // If Notary recently won a Dispute, we can optimistically assume that their passed root is valid.
        // However, we need to wait until the Dispute timeout is over, before passing the new root to LightManager.
        if (_notaryDisputeTimeout(status.notaryIndex)) {
            // We didn't pass anything, but there is a pending root
            return true;
        }
        // Check if agent root optimistic period is over
        if (status.agentRootTime + AGENT_ROOT_OPTIMISTIC_PERIOD > block.timestamp) {
            // We didn't pass anything, but there is a pending root
            return true;
        }
        // `newRoot` signer was not disputed, and the root optimistic period is over.
        // Finally, pass the Agent Merkle Root to LightManager
        InterfaceLightManager(address(agentManager)).setAgentRoot(newRoot);
        return false;
    }

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    /// @inheritdoc InterfaceDestination
    // solhint-disable-next-line ordering
    function attestationsAmount() external view returns (uint256) {
        return _roots.length;
    }

    /// @inheritdoc InterfaceDestination
    function getAttestation(uint256 index) external view returns (bytes memory attPayload, bytes memory attSignature) {
        if (index >= _roots.length) revert IndexOutOfRange();
        bytes32 snapRoot = _roots[index];
        SnapRootData memory rootData = _rootData[snapRoot];
        StoredAttData memory storedAtt = _storedAttestations[index];
        attPayload = AttestationLib.formatAttestation({
            snapRoot_: snapRoot,
            dataHash_: storedAtt.dataHash,
            nonce_: rootData.attNonce,
            blockNumber_: rootData.attBN,
            timestamp_: rootData.attTS
        });
        // Attestation signatures are not required on Synapse Chain, as the attestations could be accessed via Summit.
        if (localDomain != synapseDomain) {
            attSignature = IStatementInbox(inbox).getStoredSignature(rootData.sigIndex);
        }
    }

    /// @inheritdoc InterfaceDestination
    function getGasData(uint32 domain) external view returns (GasData gasData, uint256 dataMaturity) {
        StoredGasData memory storedGasData = _storedGasData[domain];
        // Form the data to return only if it exists and we can trust it:
        // - There is stored gas data for the domain
        // - Notary who provided the data is not in dispute
        // - Notary who provided the data is not in post-dispute timeout period
        // forgefmt: disable-next-item
        if (
            storedGasData.submittedAt != 0 &&
            !_notaryDisputeExists(storedGasData.notaryIndex) &&
            !_notaryDisputeTimeout(storedGasData.notaryIndex)
        ) {
            gasData = storedGasData.gasData;
            dataMaturity = block.timestamp - storedGasData.submittedAt;
        }
        // Return empty values if there is no data for the domain, or the notary who provided the data can't be trusted.
    }

    /// @inheritdoc InterfaceDestination
    function nextAgentRoot() external view returns (bytes32) {
        // Return current agent root on Synapse Chain for consistency
        return localDomain == synapseDomain ? IAgentManager(agentManager).agentRoot() : _nextAgentRoot;
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Saves Agent Merkle Root from the accepted attestation, if there is
    /// no pending root to be passed to LightManager.
    /// Returns the updated "last snapshot root / last agent root" status struct.
    function _saveAgentRoot(bool rootPending, bytes32 agentRoot, uint32 notaryIndex)
        internal
        returns (DestinationStatus memory status)
    {
        status = destStatus;
        // Update the timestamp for the latest snapshot root
        status.snapRootTime = ChainContext.blockTimestamp();
        // No need to save agent roots on Synapse Chain, as they could be accessed via BondingManager
        // Don't update agent root, if there is already a pending one
        // Update the data for latest agent root only if it differs from the saved one
        if (localDomain != synapseDomain && !rootPending && _nextAgentRoot != agentRoot) {
            status.agentRootTime = ChainContext.blockTimestamp();
            status.notaryIndex = notaryIndex;
            _nextAgentRoot = agentRoot;
            emit AgentRootAccepted(agentRoot);
        }
    }

    /// @dev Saves updated values from the snapshot's gas data list.
    function _saveGasData(ChainGas[] memory snapGas, uint32 notaryIndex) internal {
        uint256 statesAmount = snapGas.length;
        for (uint256 i = 0; i < statesAmount; i++) {
            ChainGas chainGas = snapGas[i];
            uint32 domain = chainGas.domain();
            // Don't save gas data for the local domain
            if (domain == localDomain) continue;
            StoredGasData memory storedGasData = _storedGasData[domain];
            // Check that the gas data is not already saved
            GasData gasData = chainGas.gasData();
            if (GasData.unwrap(gasData) == GasData.unwrap(storedGasData.gasData)) continue;
            // Save the gas data
            _storedGasData[domain] =
                StoredGasData({gasData: gasData, notaryIndex: notaryIndex, submittedAt: ChainContext.blockTimestamp()});
        }
    }
}
