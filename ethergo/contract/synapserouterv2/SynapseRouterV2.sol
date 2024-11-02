// SPDX-License-Identifier: MIT
pragma solidity =0.8.17 >=0.8.13 ^0.8.0 ^0.8.1;

// contracts/router/interfaces/IDefaultPool.sol

interface IDefaultPool {
    function swap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx,
        uint256 minDy,
        uint256 deadline
    ) external returns (uint256 amountOut);

    function calculateSwap(
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx
    ) external view returns (uint256 amountOut);

    function getToken(uint8 index) external view returns (address token);
}

// contracts/router/interfaces/IRouterAdapter.sol

interface IRouterAdapter {
    /// @notice Performs a tokenIn -> tokenOut swap, according to the provided params.
    /// If tokenIn is ETH_ADDRESS, this method should be invoked with `msg.value = amountIn`.
    /// If tokenIn is ERC20, the tokens should be already transferred to this contract (using `msg.value = 0`).
    /// If tokenOut is ETH_ADDRESS, native ETH will be sent to the recipient (be aware of potential reentrancy).
    /// If tokenOut is ERC20, the tokens will be transferred to the recipient.
    /// @dev Contracts implementing {IRouterAdapter} interface are required to enforce the above restrictions.
    /// On top of that, they must ensure that exactly `amountOut` worth of `tokenOut` is transferred to the recipient.
    /// Swap deadline and slippage is checked outside of this contract.
    /// @param recipient    Address to receive the swapped token
    /// @param tokenIn      Token to sell (use ETH_ADDRESS to start from native ETH)
    /// @param amountIn     Amount of tokens to sell
    /// @param tokenOut     Token to buy (use ETH_ADDRESS to end with native ETH)
    /// @param rawParams    Additional swap parameters
    /// @return amountOut   Amount of bought tokens
    function adapterSwap(
        address recipient,
        address tokenIn,
        uint256 amountIn,
        address tokenOut,
        bytes calldata rawParams
    ) external payable returns (uint256 amountOut);
}

// contracts/router/interfaces/IWETH9.sol

interface IWETH9 {
    function deposit() external payable;

    function withdraw(uint256 wad) external;
}

// contracts/router/libs/Errors.sol

    error DeadlineExceeded();
    error InsufficientOutputAmount();

    error MsgValueIncorrect();
    error PoolNotFound();
    error TokenAddressMismatch();
    error TokenNotContract();
    error TokenNotETH();
    error TokensIdentical();

// contracts/router/libs/Structs.sol

// "using A for B global" requires 0.8.13 or higher

// ══════════════════════════════════════════ TOKEN AND POOL DESCRIPTION ═══════════════════════════════════════════════

/// @notice Struct representing a bridge token. Used as the return value in view functions.
/// @param symbol   Bridge token symbol: unique token ID consistent among all chains
/// @param token    Bridge token address
    struct BridgeToken {
        string symbol;
        address token;
    }

/// @notice Struct used by IPoolHandler to represent a token in a pool
/// @param index    Token index in the pool
/// @param token    Token address
    struct IndexedToken {
        uint8 index;
        address token;
    }

/// @notice Struct representing a token, and the available Actions for performing a swap.
/// @param actionMask   Bitmask representing what actions (see ActionLib) are available for swapping a token
/// @param token        Token address
    struct LimitedToken {
        uint256 actionMask;
        address token;
    }

/// @notice Struct representing how pool tokens are stored by `SwapQuoter`.
/// @param isWeth   Whether the token represents Wrapped ETH.
/// @param token    Token address.
    struct PoolToken {
        bool isWeth;
        address token;
    }

/// @notice Struct representing a liquidity pool. Used as the return value in view functions.
/// @param pool         Pool address.
/// @param lpToken      Address of pool's LP token.
/// @param tokens       List of pool's tokens.
    struct Pool {
        address pool;
        address lpToken;
        PoolToken[] tokens;
    }

// ════════════════════════════════════════════════ ROUTER STRUCTS ═════════════════════════════════════════════════════

/// @notice Struct representing a quote request for swapping a bridge token.
/// Used in destination chain's SynapseRouter, hence the name "Destination Request".
/// @dev tokenOut is passed externally.
/// @param symbol   Bridge token symbol: unique token ID consistent among all chains
/// @param amountIn Amount of bridge token to start with, before the bridge fee is applied
    struct DestRequest {
        string symbol;
        uint256 amountIn;
    }

/// @notice Struct representing a swap request for SynapseRouter.
/// @dev tokenIn is supplied separately.
/// @param routerAdapter    Contract that will perform the swap for the Router. Address(0) specifies a "no swap" query.
/// @param tokenOut         Token address to swap to.
/// @param minAmountOut     Minimum amount of tokens to receive after the swap, or tx will be reverted.
/// @param deadline         Latest timestamp for when the transaction needs to be executed, or tx will be reverted.
/// @param rawParams        ABI-encoded params for the swap that will be passed to `routerAdapter`.
///                         Should be DefaultParams for swaps via DefaultAdapter.
    struct SwapQuery {
        address routerAdapter;
        address tokenOut;
        uint256 minAmountOut;
        uint256 deadline;
        bytes rawParams;
    }

    using SwapQueryLib for SwapQuery global;

library SwapQueryLib {
    /// @notice Checks whether the router adapter was specified in the query.
    /// Query without a router adapter specifies that no action needs to be taken.
    function hasAdapter(SwapQuery memory query) internal pure returns (bool) {
        return query.routerAdapter != address(0);
    }

    /// @notice Fills `routerAdapter` and `deadline` fields in query, if it specifies one of the supported Actions,
    /// and if a path for this action was found.
    function fillAdapterAndDeadline(SwapQuery memory query, address routerAdapter) internal pure {
        // Fill the fields only if some path was found.
        if (query.minAmountOut == 0) return;
        // Empty params indicates no action needs to be done, thus no adapter is needed.
        query.routerAdapter = query.rawParams.length == 0 ? address(0) : routerAdapter;
        // Set default deadline to infinity. Not using the value of 0,
        // which would lead to every swap to revert by default.
        query.deadline = type(uint256).max;
    }
}

// ════════════════════════════════════════════════ ADAPTER STRUCTS ════════════════════════════════════════════════════

/// @notice Struct representing parameters for swapping via DefaultAdapter.
/// @param action           Action that DefaultAdapter needs to perform.
/// @param pool             Liquidity pool that will be used for Swap/AddLiquidity/RemoveLiquidity actions.
/// @param tokenIndexFrom   Token index to swap from. Used for swap/addLiquidity actions.
/// @param tokenIndexTo     Token index to swap to. Used for swap/removeLiquidity actions.
    struct DefaultParams {
        Action action;
        address pool;
        uint8 tokenIndexFrom;
        uint8 tokenIndexTo;
    }

/// @notice All possible actions that DefaultAdapter could perform.
    enum Action {
        Swap, // swap between two pools tokens
        AddLiquidity, // add liquidity in a form of a single pool token
        RemoveLiquidity, // remove liquidity in a form of a single pool token
        HandleEth // ETH <> WETH interaction
    }

    using ActionLib for Action global;

/// @notice Library for dealing with bit masks which describe what set of Actions is available.
library ActionLib {
    /// @notice Returns a bitmask with all possible actions set to True.
    function allActions() internal pure returns (uint256 actionMask) {
        actionMask = type(uint256).max;
    }

    /// @notice Returns whether the given action is set to True in the bitmask.
    function isIncluded(Action action, uint256 actionMask) internal pure returns (bool) {
        return actionMask & mask(action) != 0;
    }

    /// @notice Returns a bitmask with only the given action set to True.
    function mask(Action action) internal pure returns (uint256) {
        return 1 << uint256(action);
    }

    /// @notice Returns a bitmask with only two given actions set to True.
    function mask(Action a, Action b) internal pure returns (uint256) {
        return mask(a) | mask(b);
    }
}

// node_modules/@openzeppelin/contracts-4.5.0/token/ERC20/IERC20.sol

// OpenZeppelin Contracts (last updated v4.5.0) (token/ERC20/IERC20.sol)

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `to`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `from` to `to` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);

    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

// node_modules/@openzeppelin/contracts-4.5.0/utils/Address.sol

// OpenZeppelin Contracts (last updated v4.5.0) (utils/Address.sol)

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
     * https://diligence.consensys.net/posts/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.5.11/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
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
        return functionCall(target, data, "Address: low-level call failed");
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
    function functionCallWithValue(
        address target,
        bytes memory data,
        uint256 value
    ) internal returns (bytes memory) {
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
        require(isContract(target), "Address: call to non-contract");

        (bool success, bytes memory returndata) = target.call{value: value}(data);
        return verifyCallResult(success, returndata, errorMessage);
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
        require(isContract(target), "Address: static call to non-contract");

        (bool success, bytes memory returndata) = target.staticcall(data);
        return verifyCallResult(success, returndata, errorMessage);
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
        require(isContract(target), "Address: delegate call to non-contract");

        (bool success, bytes memory returndata) = target.delegatecall(data);
        return verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Tool to verifies that a low level call was successful, and revert if it wasn't, either by bubbling the
     * revert reason using the provided one.
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
            // Look for revert reason and bubble it up if present
            if (returndata.length > 0) {
                // The easiest way to bubble the revert reason is using memory via assembly

                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert(errorMessage);
            }
        }
    }
}

// node_modules/@openzeppelin/contracts-4.5.0/utils/Context.sol

// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

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
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

// node_modules/@openzeppelin/contracts-4.5.0/utils/structs/EnumerableSet.sol

// OpenZeppelin Contracts v4.4.1 (utils/structs/EnumerableSet.sol)

/**
 * @dev Library for managing
 * https://en.wikipedia.org/wiki/Set_(abstract_data_type)[sets] of primitive
 * types.
 *
 * Sets have the following properties:
 *
 * - Elements are added, removed, and checked for existence in constant time
 * (O(1)).
 * - Elements are enumerated in O(n). No guarantees are made on the ordering.
 *
 * ```
 * contract Example {
 *     // Add the library methods
 *     using EnumerableSet for EnumerableSet.AddressSet;
 *
 *     // Declare a set state variable
 *     EnumerableSet.AddressSet private mySet;
 * }
 * ```
 *
 * As of v3.3.0, sets of type `bytes32` (`Bytes32Set`), `address` (`AddressSet`)
 * and `uint256` (`UintSet`) are supported.
 */
library EnumerableSet {
    // To implement this library for multiple types with as little code
    // repetition as possible, we write it in terms of a generic Set type with
    // bytes32 values.
    // The Set implementation uses private functions, and user-facing
    // implementations (such as AddressSet) are just wrappers around the
    // underlying Set.
    // This means that we can only create new EnumerableSets for types that fit
    // in bytes32.

    struct Set {
        // Storage of set values
        bytes32[] _values;
        // Position of the value in the `values` array, plus 1 because index 0
        // means a value is not in the set.
        mapping(bytes32 => uint256) _indexes;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function _add(Set storage set, bytes32 value) private returns (bool) {
        if (!_contains(set, value)) {
            set._values.push(value);
            // The value is stored at length-1, but we add 1 to all indexes
            // and use 0 as a sentinel value
            set._indexes[value] = set._values.length;
            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function _remove(Set storage set, bytes32 value) private returns (bool) {
        // We read and store the value's index to prevent multiple reads from the same storage slot
        uint256 valueIndex = set._indexes[value];

        if (valueIndex != 0) {
            // Equivalent to contains(set, value)
            // To delete an element from the _values array in O(1), we swap the element to delete with the last one in
            // the array, and then remove the last element (sometimes called as 'swap and pop').
            // This modifies the order of the array, as noted in {at}.

            uint256 toDeleteIndex = valueIndex - 1;
            uint256 lastIndex = set._values.length - 1;

            if (lastIndex != toDeleteIndex) {
                bytes32 lastvalue = set._values[lastIndex];

                // Move the last value to the index where the value to delete is
                set._values[toDeleteIndex] = lastvalue;
                // Update the index for the moved value
                set._indexes[lastvalue] = valueIndex; // Replace lastvalue's index to valueIndex
            }

            // Delete the slot where the moved value was stored
            set._values.pop();

            // Delete the index for the deleted slot
            delete set._indexes[value];

            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function _contains(Set storage set, bytes32 value) private view returns (bool) {
        return set._indexes[value] != 0;
    }

    /**
     * @dev Returns the number of values on the set. O(1).
     */
    function _length(Set storage set) private view returns (uint256) {
        return set._values.length;
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function _at(Set storage set, uint256 index) private view returns (bytes32) {
        return set._values[index];
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function _values(Set storage set) private view returns (bytes32[] memory) {
        return set._values;
    }

    // Bytes32Set

    struct Bytes32Set {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(Bytes32Set storage set, bytes32 value) internal returns (bool) {
        return _add(set._inner, value);
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(Bytes32Set storage set, bytes32 value) internal returns (bool) {
        return _remove(set._inner, value);
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(Bytes32Set storage set, bytes32 value) internal view returns (bool) {
        return _contains(set._inner, value);
    }

    /**
     * @dev Returns the number of values in the set. O(1).
     */
    function length(Bytes32Set storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(Bytes32Set storage set, uint256 index) internal view returns (bytes32) {
        return _at(set._inner, index);
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function values(Bytes32Set storage set) internal view returns (bytes32[] memory) {
        return _values(set._inner);
    }

    // AddressSet

    struct AddressSet {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(AddressSet storage set, address value) internal returns (bool) {
        return _add(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(AddressSet storage set, address value) internal returns (bool) {
        return _remove(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(AddressSet storage set, address value) internal view returns (bool) {
        return _contains(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Returns the number of values in the set. O(1).
     */
    function length(AddressSet storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(AddressSet storage set, uint256 index) internal view returns (address) {
        return address(uint160(uint256(_at(set._inner, index))));
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function values(AddressSet storage set) internal view returns (address[] memory) {
        bytes32[] memory store = _values(set._inner);
        address[] memory result;

        assembly {
            result := store
        }

        return result;
    }

    // UintSet

    struct UintSet {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(UintSet storage set, uint256 value) internal returns (bool) {
        return _add(set._inner, bytes32(value));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(UintSet storage set, uint256 value) internal returns (bool) {
        return _remove(set._inner, bytes32(value));
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(UintSet storage set, uint256 value) internal view returns (bool) {
        return _contains(set._inner, bytes32(value));
    }

    /**
     * @dev Returns the number of values on the set. O(1).
     */
    function length(UintSet storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

    /**
     * @dev Returns the value stored at position `index` in the set. O(1).
     *
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(UintSet storage set, uint256 index) internal view returns (uint256) {
        return uint256(_at(set._inner, index));
    }

    /**
     * @dev Return the entire set in an array
     *
     * WARNING: This operation will copy the entire storage to memory, which can be quite expensive. This is designed
     * to mostly be used by view accessors that are queried without any gas fees. Developers should keep in mind that
     * this function has an unbounded cost, and using it as part of a state-changing function may render the function
     * uncallable if the set grows to a point where copying to memory consumes too much gas to fit in a block.
     */
    function values(UintSet storage set) internal view returns (uint256[] memory) {
        bytes32[] memory store = _values(set._inner);
        uint256[] memory result;

        assembly {
            result := store
        }

        return result;
    }
}

// contracts/router/interfaces/IBridgeModule.sol

interface IBridgeModule {
    /// @notice Performs a bridging transaction on behalf of the sender, assuming they already have `token`.
    /// @dev This will be used via delegatecall from SynapseRouterV2, which will have custody over the bridge tokens.
    /// This will revert if delegatecall is not used.
    /// @param to            Address to receive tokens on destination chain
    /// @param chainId       Destination chain id
    /// @param token         Address of the bridge token
    /// @param amount        Amount of the tokens for the bridge transaction
    /// @param destQuery     Destination swap query. Empty struct indicates no swap is required
    function delegateBridge(
        address to,
        uint256 chainId,
        address token,
        uint256 amount,
        SwapQuery memory destQuery
    ) external payable;

    /// @notice Gets the maximum amount of tokens user can bridge from this chain.
    /// @param token        Address of the bridge token
    /// @return amount      Max amount of tokens user can bridge from this chain
    function getMaxBridgedAmount(address token) external view returns (uint256 amount);

    /// @notice Calculates the fee amount for bridging a token to this chain.
    /// @dev Will revert if the token is not supported.
    /// @param token        Address of the bridge token
    /// @param amount       Amount of tokens to be bridged
    /// @param isSwap       Whether the user provided swap details for converting the bridge token
    ///                     to the final token on this chain
    /// @return fee         Fee amount
    function calculateFeeAmount(
        address token,
        uint256 amount,
        bool isSwap
    ) external view returns (uint256 fee);

    /// @notice Returns the list of all supported bridge tokens and their bridge symbols.
    /// - Bridge symbol is consistent across all chains for a given token and their bridge.
    /// - Bridge symbol doesn't have to be the same as the token symbol on this chain.
    /// @return bridgeTokens Supported bridge tokens and their bridge symbols
    function getBridgeTokens() external view returns (BridgeToken[] memory bridgeTokens);

    /// @notice Returns the address of the bridge token for a given bridge symbol.
    /// - Bridge symbol is consistent across all chains for a given token and their bridge.
    /// - Bridge symbol doesn't have to be the same as the token symbol on this chain.
    /// @dev Will return address(0) if the token is not supported.
    /// @param symbol       Symbol of the supported bridge token used by the token's bridge
    /// @return token       Address of the bridge token
    function symbolToToken(string memory symbol) external view returns (address token);

    /// @notice Returns the bridge symbol of a given bridge token.
    /// - Bridge symbol is consistent across all chains for a given token and their bridge.
    /// - Bridge symbol doesn't have to be the same as the token symbol on this chain.
    /// @dev Will return empty string if the token is not supported.
    /// @param token        Address of the bridge token
    /// @return symbol      Symbol of the supported bridge token used by the token's bridge
    function tokenToSymbol(address token) external view returns (string memory symbol);

    /// @notice Returns the action mask associated with bridging a token to this chain.
    /// Action mask is a bitmask of the actions that could be performed with the token atomically with the
    /// incoming bridge transaction to this chain. See Structs.sol for the list of actions.
    /// @dev Will return 0 (empty mask) if the token is not supported.
    /// @param token        Address of the bridge token
    /// @return actionMask  Action mask for the bridge token
    function tokenToActionMask(address token) external view returns (uint256 actionMask);
}

// contracts/router/interfaces/IDefaultExtendedPool.sol

interface IDefaultExtendedPool is IDefaultPool {
    function addLiquidity(
        uint256[] calldata amounts,
        uint256 minToMint,
        uint256 deadline
    ) external returns (uint256);

    function removeLiquidityOneToken(
        uint256 tokenAmount,
        uint8 tokenIndex,
        uint256 minAmount,
        uint256 deadline
    ) external returns (uint256);

    // ═══════════════════════════════════════════════════ VIEWS ═══════════════════════════════════════════════════════

    function calculateRemoveLiquidity(uint256 amount) external view returns (uint256[] memory);

    function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex)
    external
    view
    returns (uint256 availableTokenAmount);

    function getAPrecise() external view returns (uint256);

    function getTokenBalance(uint8 index) external view returns (uint256);

    function swapStorage()
    external
    view
    returns (
        uint256 initialA,
        uint256 futureA,
        uint256 initialATime,
        uint256 futureATime,
        uint256 swapFee,
        uint256 adminFee,
        address lpToken
    );
}

// contracts/router/interfaces/ISwapQuoterV1.sol

/// @notice Interface for the SwapQuoterV1 version with updated pragma and enriched docs.
interface ISwapQuoterV1 {
    // ════════════════════════════════════════════════ IMMUTABLES ════════════════════════════════════════════════════

    /// @notice Address of deployed calculator contract for DefaultPool, which is able to calculate
    /// EXACT quotes for AddLiquidity action (something that DefaultPool contract itself is unable to do).
    function defaultPoolCalc() external view returns (address);

    /// @notice Address of WETH token used in the pools. Represents wrapped version of chain's native currency,
    /// e.g. WETH on Ethereum, WBNB on BSC, etc.
    function weth() external view returns (address);

    // ═══════════════════════════════════════════════ POOL GETTERS ════════════════════════════════════════════════════

    /// @notice Returns a list of all supported pools.
    function allPools() external view returns (Pool[] memory pools);

    /// @notice Returns the amount of supported pools.
    function poolsAmount() external view returns (uint256 amtPools);

    /// @notice Returns the number of tokens the given pool supports and the pool's LP token.
    function poolInfo(address pool) external view returns (uint256 numTokens, address lpToken);

    /// @notice Returns a list of pool tokens for the given pool.
    function poolTokens(address pool) external view returns (PoolToken[] memory tokens);

    // ══════════════════════════════════════════════ GENERAL QUOTES ═══════════════════════════════════════════════════

    /// @notice Checks if a swap is possible between every bridge token in the given list and tokenOut.
    /// Only the bridge token's whitelisted pool is considered for every `tokenIn -> tokenOut` swap.
    /// @param bridgeTokensIn   List of structs with following information:
    ///                         - actionMask    Bitmask of available actions for doing tokenIn -> tokenOut
    ///                         - token         Bridge token address to swap from
    /// @param tokenOut         Token address to swap to
    /// @return amountFound     Amount of tokens from the list that are swappable to tokenOut
    /// @return isConnected     List of bool values, specifying whether a token from the list is swappable to tokenOut
    function findConnectedTokens(LimitedToken[] memory bridgeTokensIn, address tokenOut)
    external
    view
    returns (uint256 amountFound, bool[] memory isConnected);

    /// @notice Finds the quote and the swap parameters for a tokenIn -> tokenOut swap from the list of supported pools.
    /// - If this is a request for the swap to be performed immediately (or the "origin swap" in the bridge workflow),
    /// `tokenIn.actionMask` needs to be set to bitmask of all possible actions (ActionLib.allActions()).
    /// - If this is a request for the swap to be performed as the "destination swap" in the bridge workflow,
    /// `tokenIn.actionMask` needs to be set to bitmask of possible actions for `tokenIn.token` as a bridge token,
    /// e.g. Action.Swap for minted tokens, or Action.RemoveLiquidity | Action.HandleEth for withdrawn tokens.
    /// > Returns the `SwapQuery` struct, that can be used on SynapseRouter.
    /// > minAmountOut and deadline fields will need to be adjusted based on the swap settings.
    /// @dev If tokenIn or tokenOut is ETH_ADDRESS, only the pools having WETH as a pool token will be considered.
    /// Three potential outcomes are available:
    /// 1. `tokenIn` and `tokenOut` represent the same token address (identical tokens).
    /// 2. `tokenIn` and `tokenOut` represent different addresses. No trade path from `tokenIn` to `tokenOut` is found.
    /// 3. `tokenIn` and `tokenOut` represent different addresses. Trade path from `tokenIn` to `tokenOut` is found.
    /// The exact composition of the returned struct for every case is documented in the return parameter documentation.
    /// @param tokenIn  Struct with following information:
    ///                 - actionMask    Bitmask of available actions for doing tokenIn -> tokenOut
    ///                 - token         Token address to swap from
    /// @param tokenOut Token address to swap to
    /// @param amountIn Amount of tokens to swap from
    /// @return query   Struct representing trade path between tokenIn and tokenOut:
    ///                 - swapAdapter: adapter address that would handle the swap. Address(0) if no path is found,
    ///                 or tokens are identical. Address of SynapseRouter otherwise.
    ///                 - tokenOut: always equals to the provided `tokenOut`, even if no path if found.
    ///                 - minAmountOut: amount of `tokenOut`, if swap was completed now. 0, if no path is found.
    ///                 - deadline: 2**256-1 if path was found, or tokens are identical. 0, if no path is found.
    ///                 - rawParams: ABI-encoded DefaultParams struct indicating the swap parameters. Empty string,
    ///                 if no path is found, or tokens are identical.
    function getAmountOut(
        LimitedToken memory tokenIn,
        address tokenOut,
        uint256 amountIn
    ) external view returns (SwapQuery memory query);

    // ═══════════════════════════════════════════ SPECIFIC POOL QUOTES ════════════════════════════════════════════════

    /// @notice Returns the exact quote for adding liquidity to a given pool in a form of a single token.
    /// @dev The only way to get a quote for adding liquidity would be `pool.calculateTokenAmount()`,
    /// which gives an ESTIMATE: it doesn't take the trade fees into account.
    /// We do need the exact quotes for (DAI/USDC/USDT) -> nUSD "swaps" on Mainnet, hence we do this.
    /// We also need the exact quotes for adding liquidity to the pools.
    /// Note: the function might revert instead of returning 0 for incorrect requests. Make sure
    /// to take that into account.
    function calculateAddLiquidity(address pool, uint256[] memory amounts) external view returns (uint256 amountOut);

    /// @notice Returns the exact quote for swapping between two given tokens.
    /// @dev Exposes IDefaultPool.calculateSwap(tokenIndexFrom, tokenIndexTo, dx);
    function calculateSwap(
        address pool,
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 dx
    ) external view returns (uint256 amountOut);

    /// @notice Returns the exact quote for withdrawing pools tokens in a balanced way.
    /// @dev Exposes IDefaultPool.calculateRemoveLiquidity(amount);
    function calculateRemoveLiquidity(address pool, uint256 amount) external view returns (uint256[] memory amountsOut);

    /// @notice Returns the exact quote for withdrawing a single pool token.
    /// @dev Exposes IDefaultPool.calculateRemoveLiquidityOneToken(tokenAmount, tokenIndex);
    function calculateWithdrawOneToken(
        address pool,
        uint256 tokenAmount,
        uint8 tokenIndex
    ) external view returns (uint256 amountOut);
}

// contracts/router/libs/Arrays.sol

/// @notice Arrays library offers helper functions for working with arrays and array of arrays
library Arrays {
    error ArrayLengthInvalid(uint256 count);

    /// @notice Flattens out a list of lists of bridge tokens into a list of bridge tokens
    /// @param unflattened The list of lists of bridge tokens
    /// @param count The total number of bridge tokens in unflattened
    /// @return flattened The flattened list of bridge tokens
    function flatten(BridgeToken[][] memory unflattened, uint256 count)
    internal
    pure
    returns (BridgeToken[] memory flattened)
    {
        flattened = new BridgeToken[](count);

        uint256 k;
        for (uint256 i = 0; i < unflattened.length; ++i) {
            for (uint256 j = 0; j < unflattened[i].length; ++j) {
                flattened[k] = unflattened[i][j];
                k++;
            }
        }

        if (k != count) revert ArrayLengthInvalid(k); // @dev should never happen in practice w router
    }

    /// @notice Flattens out a list of lists of addresses into a list of addresses
    /// @param unflattened The list of lists of addresses
    /// @param count The total number of addresses in unflattened
    /// @return flattened The flattened list of addresses
    function flatten(address[][] memory unflattened, uint256 count) internal pure returns (address[] memory flattened) {
        flattened = new address[](count);

        uint256 k;
        for (uint256 i = 0; i < unflattened.length; ++i) {
            for (uint256 j = 0; j < unflattened[i].length; ++j) {
                flattened[k] = unflattened[i][j];
                k++;
            }
        }

        if (k != count) revert ArrayLengthInvalid(k); // @dev should never happen in practice w router
    }

    /// @notice Converts a list of bridge tokens to a list of their token addresses
    /// @param b The list of bridge tokens
    /// @return t The list of token addresses associated with given bridge tokens
    function tokens(BridgeToken[] memory b) internal pure returns (address[] memory t) {
        t = new address[](b.length);
        for (uint256 i = 0; i < b.length; ++i) t[i] = b[i].token;
    }

    /// @notice Converts a list of bridge tokens to a list of their token symbols
    /// @param b The list of bridge tokens
    /// @return s The list of token symbols associated with given bridge tokens
    function symbols(BridgeToken[] memory b) internal pure returns (string[] memory s) {
        s = new string[](b.length);
        for (uint256 i = 0; i < b.length; ++i) s[i] = b[i].symbol;
    }

    /// @notice Filters out duplicates and zero addresses from given list of addresses
    /// @dev Removes zero addresses from list
    /// @param unfiltered The list of addresses with duplicates
    /// @return filtered The list of addresses without duplicates
    function unique(address[] memory unfiltered) internal pure returns (address[] memory filtered) {
        address[] memory intermediate = new address[](unfiltered.length);

        // add unique elements to intermediate
        uint256 count;
        for (uint256 i = 0; i < unfiltered.length; ++i) {
            address el = unfiltered[i];
            if (!contains(intermediate, el)) {
                intermediate[count] = el;
                count++;
            }
        }

        // remove the zero elements at the end if any duplicates
        filtered = new address[](count);
        for (uint256 i = 0; i < count; i++) {
            filtered[i] = intermediate[i];
        }
    }

    /// @notice Whether given element is in the list of addresses
    /// @param l The list of addresses
    /// @param el The element to search for
    /// @return does If given list does contain element
    function contains(address[] memory l, address el) internal pure returns (bool does) {
        for (uint256 j = 0; j < l.length; ++j) {
            does = (el == l[j]);
            if (does) break;
        }
    }

    /// @notice Appends a new element to the end of the list of addresses
    /// @param l The list of addresses
    /// @param el The element to append
    /// @param r The new list of addresses with appended element
    function append(address[] memory l, address el) internal pure returns (address[] memory r) {
        r = new address[](l.length + 1);
        for (uint256 i = 0; i < l.length; i++) r[i] = l[i];
        r[r.length - 1] = el;
    }
}

// node_modules/@openzeppelin/contracts-4.5.0/access/Ownable.sol

// OpenZeppelin Contracts v4.4.1 (access/Ownable.sol)

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
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
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
}

// node_modules/@openzeppelin/contracts-4.5.0/utils/structs/EnumerableMap.sol

// OpenZeppelin Contracts v4.4.1 (utils/structs/EnumerableMap.sol)

/**
 * @dev Library for managing an enumerable variant of Solidity's
 * https://solidity.readthedocs.io/en/latest/types.html#mapping-types[`mapping`]
 * type.
 *
 * Maps have the following properties:
 *
 * - Entries are added, removed, and checked for existence in constant time
 * (O(1)).
 * - Entries are enumerated in O(n). No guarantees are made on the ordering.
 *
 * ```
 * contract Example {
 *     // Add the library methods
 *     using EnumerableMap for EnumerableMap.UintToAddressMap;
 *
 *     // Declare a set state variable
 *     EnumerableMap.UintToAddressMap private myMap;
 * }
 * ```
 *
 * As of v3.0.0, only maps of type `uint256 -> address` (`UintToAddressMap`) are
 * supported.
 */
library EnumerableMap {
    using EnumerableSet for EnumerableSet.Bytes32Set;

    // To implement this library for multiple types with as little code
    // repetition as possible, we write it in terms of a generic Map type with
    // bytes32 keys and values.
    // The Map implementation uses private functions, and user-facing
    // implementations (such as Uint256ToAddressMap) are just wrappers around
    // the underlying Map.
    // This means that we can only create new EnumerableMaps for types that fit
    // in bytes32.

    struct Map {
        // Storage of keys
        EnumerableSet.Bytes32Set _keys;
        mapping(bytes32 => bytes32) _values;
    }

    /**
     * @dev Adds a key-value pair to a map, or updates the value for an existing
     * key. O(1).
     *
     * Returns true if the key was added to the map, that is if it was not
     * already present.
     */
    function _set(
        Map storage map,
        bytes32 key,
        bytes32 value
    ) private returns (bool) {
        map._values[key] = value;
        return map._keys.add(key);
    }

    /**
     * @dev Removes a key-value pair from a map. O(1).
     *
     * Returns true if the key was removed from the map, that is if it was present.
     */
    function _remove(Map storage map, bytes32 key) private returns (bool) {
        delete map._values[key];
        return map._keys.remove(key);
    }

    /**
     * @dev Returns true if the key is in the map. O(1).
     */
    function _contains(Map storage map, bytes32 key) private view returns (bool) {
        return map._keys.contains(key);
    }

    /**
     * @dev Returns the number of key-value pairs in the map. O(1).
     */
    function _length(Map storage map) private view returns (uint256) {
        return map._keys.length();
    }

    /**
     * @dev Returns the key-value pair stored at position `index` in the map. O(1).
     *
     * Note that there are no guarantees on the ordering of entries inside the
     * array, and it may change when more entries are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function _at(Map storage map, uint256 index) private view returns (bytes32, bytes32) {
        bytes32 key = map._keys.at(index);
        return (key, map._values[key]);
    }

    /**
     * @dev Tries to returns the value associated with `key`.  O(1).
     * Does not revert if `key` is not in the map.
     */
    function _tryGet(Map storage map, bytes32 key) private view returns (bool, bytes32) {
        bytes32 value = map._values[key];
        if (value == bytes32(0)) {
            return (_contains(map, key), bytes32(0));
        } else {
            return (true, value);
        }
    }

    /**
     * @dev Returns the value associated with `key`.  O(1).
     *
     * Requirements:
     *
     * - `key` must be in the map.
     */
    function _get(Map storage map, bytes32 key) private view returns (bytes32) {
        bytes32 value = map._values[key];
        require(value != 0 || _contains(map, key), "EnumerableMap: nonexistent key");
        return value;
    }

    /**
     * @dev Same as {_get}, with a custom error message when `key` is not in the map.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {_tryGet}.
     */
    function _get(
        Map storage map,
        bytes32 key,
        string memory errorMessage
    ) private view returns (bytes32) {
        bytes32 value = map._values[key];
        require(value != 0 || _contains(map, key), errorMessage);
        return value;
    }

    // UintToAddressMap

    struct UintToAddressMap {
        Map _inner;
    }

    /**
     * @dev Adds a key-value pair to a map, or updates the value for an existing
     * key. O(1).
     *
     * Returns true if the key was added to the map, that is if it was not
     * already present.
     */
    function set(
        UintToAddressMap storage map,
        uint256 key,
        address value
    ) internal returns (bool) {
        return _set(map._inner, bytes32(key), bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the key was removed from the map, that is if it was present.
     */
    function remove(UintToAddressMap storage map, uint256 key) internal returns (bool) {
        return _remove(map._inner, bytes32(key));
    }

    /**
     * @dev Returns true if the key is in the map. O(1).
     */
    function contains(UintToAddressMap storage map, uint256 key) internal view returns (bool) {
        return _contains(map._inner, bytes32(key));
    }

    /**
     * @dev Returns the number of elements in the map. O(1).
     */
    function length(UintToAddressMap storage map) internal view returns (uint256) {
        return _length(map._inner);
    }

    /**
     * @dev Returns the element stored at position `index` in the set. O(1).
     * Note that there are no guarantees on the ordering of values inside the
     * array, and it may change when more values are added or removed.
     *
     * Requirements:
     *
     * - `index` must be strictly less than {length}.
     */
    function at(UintToAddressMap storage map, uint256 index) internal view returns (uint256, address) {
        (bytes32 key, bytes32 value) = _at(map._inner, index);
        return (uint256(key), address(uint160(uint256(value))));
    }

    /**
     * @dev Tries to returns the value associated with `key`.  O(1).
     * Does not revert if `key` is not in the map.
     *
     * _Available since v3.4._
     */
    function tryGet(UintToAddressMap storage map, uint256 key) internal view returns (bool, address) {
        (bool success, bytes32 value) = _tryGet(map._inner, bytes32(key));
        return (success, address(uint160(uint256(value))));
    }

    /**
     * @dev Returns the value associated with `key`.  O(1).
     *
     * Requirements:
     *
     * - `key` must be in the map.
     */
    function get(UintToAddressMap storage map, uint256 key) internal view returns (address) {
        return address(uint160(uint256(_get(map._inner, bytes32(key)))));
    }

    /**
     * @dev Same as {get}, with a custom error message when `key` is not in the map.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {tryGet}.
     */
    function get(
        UintToAddressMap storage map,
        uint256 key,
        string memory errorMessage
    ) internal view returns (address) {
        return address(uint160(uint256(_get(map._inner, bytes32(key), errorMessage))));
    }
}

// contracts/router/interfaces/ISwapQuoterV2.sol

interface ISwapQuoterV2 is ISwapQuoterV1 {
    /// @notice Checks if tokenIn -> tokenOut swap is possible using the supported pools.
    /// Follows `getAmountOut()` convention when it comes to providing tokenIn.actionMask:
    /// - If this is a request for the swap to be performed immediately (or the "origin swap" in the bridge workflow),
    /// `tokenIn.actionMask` needs to be set to bitmask of all possible actions (ActionLib.allActions()).
    ///  For this case, all pools added to SwapQuoterV2 will be considered for the swap.
    /// - If this is a request for the swap to be performed as the "destination swap" in the bridge workflow,
    /// `tokenIn.actionMask` needs to be set to bitmask of possible actions for `tokenIn.token` as a bridge token,
    /// e.g. Action.Swap for minted tokens, or Action.RemoveLiquidity | Action.HandleEth for withdrawn tokens.
    ///
    /// As for the pools considered for the swap, there are two cases:
    /// - If this is a request for the swap to be performed immediately (or the "origin swap" in the bridge workflow),
    /// all pools added to SwapQuoterV2 will be considered for the swap.
    /// - If this is a request for the swap to be performed as the "destination swap" in the bridge workflow,
    /// only the whitelisted pool for tokenIn.token will be considered for the swap.
    function areConnectedTokens(LimitedToken memory tokenIn, address tokenOut) external view returns (bool);

    /// @notice Allows to set the SynapseRouter contract, which is used as "Router Adapter" for doing
    /// swaps through Default Pools (or handling ETH).
    /// Note: this will not affect the old SynapseRouter contract which still uses this Quoter, as the old SynapseRouter
    /// could handle the requests with the new SynapseRouter as external "Router Adapter".
    function setSynapseRouter(address synapseRouter_) external;
}

// node_modules/@openzeppelin/contracts-4.5.0/token/ERC20/utils/SafeERC20.sol

// OpenZeppelin Contracts v4.4.1 (token/ERC20/utils/SafeERC20.sol)

/**
 * @title SafeERC20
 * @dev Wrappers around ERC20 operations that throw on failure (when the token
 * contract returns false). Tokens that return no value (and instead revert or
 * throw on failure) are also supported, non-reverting calls are assumed to be
 * successful.
 * To use this library you can add a `using SafeERC20 for IERC20;` statement to your contract,
 * which allows you to call the safe operations as `token.safeTransfer(...)`, etc.
 */
library SafeERC20 {
    using Address for address;

    function safeTransfer(
        IERC20 token,
        address to,
        uint256 value
    ) internal {
        _callOptionalReturn(token, abi.encodeWithSelector(token.transfer.selector, to, value));
    }

    function safeTransferFrom(
        IERC20 token,
        address from,
        address to,
        uint256 value
    ) internal {
        _callOptionalReturn(token, abi.encodeWithSelector(token.transferFrom.selector, from, to, value));
    }

    /**
     * @dev Deprecated. This function has issues similar to the ones found in
     * {IERC20-approve}, and its usage is discouraged.
     *
     * Whenever possible, use {safeIncreaseAllowance} and
     * {safeDecreaseAllowance} instead.
     */
    function safeApprove(
        IERC20 token,
        address spender,
        uint256 value
    ) internal {
        // safeApprove should only be called when setting an initial allowance,
        // or when resetting it to zero. To increase and decrease it, use
        // 'safeIncreaseAllowance' and 'safeDecreaseAllowance'
        require(
            (value == 0) || (token.allowance(address(this), spender) == 0),
            "SafeERC20: approve from non-zero to non-zero allowance"
        );
        _callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, value));
    }

    function safeIncreaseAllowance(
        IERC20 token,
        address spender,
        uint256 value
    ) internal {
        uint256 newAllowance = token.allowance(address(this), spender) + value;
        _callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, newAllowance));
    }

    function safeDecreaseAllowance(
        IERC20 token,
        address spender,
        uint256 value
    ) internal {
        unchecked {
            uint256 oldAllowance = token.allowance(address(this), spender);
            require(oldAllowance >= value, "SafeERC20: decreased allowance below zero");
            uint256 newAllowance = oldAllowance - value;
            _callOptionalReturn(token, abi.encodeWithSelector(token.approve.selector, spender, newAllowance));
        }
    }

    /**
     * @dev Imitates a Solidity high-level call (i.e. a regular function call to a contract), relaxing the requirement
     * on the return value: the return value is optional (but if data is returned, it must not be false).
     * @param token The token targeted by the call.
     * @param data The call data (encoded using abi.encode or one of its variants).
     */
    function _callOptionalReturn(IERC20 token, bytes memory data) private {
        // We need to perform a low level call here, to bypass Solidity's return data size checking mechanism, since
        // we're implementing it ourselves. We use {Address.functionCall} to perform this call, which verifies that
        // the target address contains contract code and also asserts for success in the low-level call.

        bytes memory returndata = address(token).functionCall(data, "SafeERC20: low-level call failed");
        if (returndata.length > 0) {
            // Return data is optional
            require(abi.decode(returndata, (bool)), "SafeERC20: ERC20 operation did not succeed");
        }
    }
}

// contracts/router/interfaces/IRouterV2.sol

interface IRouterV2 {
    /// @notice Initiate a bridge transaction with an optional swap on both origin and destination chains.
    /// @dev Note that method is payable.
    /// If token is ETH_ADDRESS, this method should be invoked with `msg.value = amountIn`.
    /// If token is ERC20, the tokens will be pulled from msg.sender (use `msg.value = 0`).
    /// Make sure to approve this contract for spending `token` beforehand.
    /// originQuery.tokenOut should never be ETH_ADDRESS, bridge only works with ERC20 tokens.
    ///
    /// `token` is always a token user is sending. In case token requires a wrapper token to be bridge,
    /// use underlying address for `token` instead of the wrapper one.
    ///
    /// `originQuery` contains instructions for the swap on origin chain. As above, originQuery.tokenOut
    /// should always use the underlying address. In other words, the concept of wrapper token is fully
    /// abstracted away from the end user.
    ///
    /// `originQuery` is supposed to be fetched using SynapseRouter.getOriginAmountOut().
    /// Alternatively one could use an external adapter for more complex swaps on the origin chain.
    ///
    /// `destQuery` is supposed to be fetched using SynapseRouter.getDestinationAmountOut().
    /// Complex swaps on destination chain are not supported for the time being.
    /// Check contract description above for more details.
    ///
    /// @param to            Address to receive tokens on destination chain
    /// @param chainId       Destination chain id
    /// @param moduleId      Bridge module id to delegate bridge call
    /// @param token         Initial token for the bridge transaction to be pulled from the user
    /// @param amount        Amount of the initial tokens for the bridge transaction
    /// @param originQuery   Origin swap query. Empty struct indicates no swap is required
    /// @param destQuery     Destination swap query. Empty struct indicates no swap is required
    function bridgeViaSynapse(
        address to,
        uint256 chainId,
        bytes32 moduleId,
        address token,
        uint256 amount,
        SwapQuery memory originQuery,
        SwapQuery memory destQuery
    ) external payable;

    /// @notice Perform a swap using the supplied parameters.
    /// @dev Note that method is payable.
    /// If token is ETH_ADDRESS, this method should be invoked with `msg.value = amountIn`.
    /// If token is ERC20, the tokens will be pulled from msg.sender (use `msg.value = 0`).
    /// Make sure to approve this contract for spending `token` beforehand.
    /// If query.tokenOut is ETH_ADDRESS, native ETH will be sent to the recipient (be aware of potential reentrancy).
    /// If query.tokenOut is ERC20, the tokens will be transferred to the recipient.
    /// @param to            Address to receive swapped tokens
    /// @param token         Token to swap
    /// @param amount        Amount of tokens to swap
    /// @param query         Query with the swap parameters (see BridgeStructs.sol)
    /// @return amountOut    Amount of swapped tokens received by the user
    function swap(
        address to,
        address token,
        uint256 amount,
        SwapQuery memory query
    ) external payable returns (uint256 amountOut);

    /// @notice Sets the Swap Quoter address to get the swap quotes from.
    /// @param _swapQuoter Swap Quoter
    function setSwapQuoter(ISwapQuoterV2 _swapQuoter) external;

    /// @notice Sets a custom allowance for the given token.
    /// @dev Reverts if not router owner. To be used for the wrapper token setups.
    function setAllowance(
        address token,
        address spender,
        uint256 amount
    ) external;

    /// @notice Whitelists a new bridge module for users to route through
    /// @dev Reverts if not router owner
    /// @param moduleId Unique bridge module ID
    /// @param bridgeModule Bridge module address
    function connectBridgeModule(bytes32 moduleId, address bridgeModule) external;

    /// @notice Updates a whitelisted bridge module
    /// @dev Reverts if not router owner
    /// @param moduleId Unique bridge module ID
    /// @param bridgeModule New bridge module address to update to
    function updateBridgeModule(bytes32 moduleId, address bridgeModule) external;

    /// @notice Disconnects a whitelisted bridge module
    /// @dev Reverts if not router owner
    /// @param moduleId Unique bridge module ID
    function disconnectBridgeModule(bytes32 moduleId) external;

    /// @notice Gets the address associated with the given bridge module ID
    /// @param moduleId Unique bridge module ID
    /// @return bridgeModule Bridge module address
    function idToModule(bytes32 moduleId) external view returns (address bridgeModule);

    /// @notice Gets the module ID associated with the given bridge module address
    /// @param bridgeModule Bridge module address
    /// @return moduleId Unique bridge module ID
    function moduleToId(address bridgeModule) external view returns (bytes32 moduleId);

    /// @notice Gets all bridge tokens for supported bridge modules
    /// @dev Intended for off-chain queries.
    /// @return bridgeTokens List of structs with following information:
    ///                  - symbol: unique token ID consistent among all chains
    ///                  - token: bridge token address
    function getBridgeTokens() external view returns (BridgeToken[] memory bridgeTokens);

    /// @notice Gets the list of all bridge tokens (and their symbols), such that destination swap
    /// from a bridge token to `tokenOut` is possible.
    /// @dev Intended for off-chain queries.
    /// @param tokenOut  Token address to swap to on destination chain
    /// @return destTokens List of structs with following information:
    ///                  - symbol: unique token ID consistent among all chains
    ///                  - token: bridge token address
    function getDestinationBridgeTokens(address tokenOut) external view returns (BridgeToken[] memory destTokens);

    /// @notice Gets the list of all bridge tokens (and their symbols), such that origin swap
    /// from `tokenIn` to a bridge token is possible.
    /// @dev Intended for off-chain queries.
    /// @param tokenIn  Token address to swap from on origin chain
    /// @return originTokens List of structs with following information:
    ///                  - symbol: unique token ID consistent among all chains
    ///                  - token: bridge token address
    function getOriginBridgeTokens(address tokenIn) external view returns (BridgeToken[] memory originTokens);

    /// @notice Gets the list of all tokens that can be swapped into bridge tokens on this chain.
    /// @dev Supported tokens should include all bridge tokens and any pool tokens paired with a bridge token.
    /// @dev Intended for off-chain queries.
    /// @return supportedTokens Supported token addresses that can be swapped for bridge tokens
    function getSupportedTokens() external view returns (address[] memory supportedTokens);

    /// @notice Finds the best path between every supported bridge token from the given list and `tokenOut`,
    /// treating the swap as "destination swap", limiting possible actions to those available for every bridge token.
    /// @dev Intended for off-chain queries. Will NOT revert if any of the tokens are not supported, instead will return an empty query for that symbol.
    /// Note: it is NOT possible to form a SwapQuery off-chain using alternative SwapAdapter for the destination swap.
    /// For the time being, only swaps through the Synapse-supported pools are available on destination chain.
    /// @param request Struct with following information:
    ///                 - symbol: unique token ID consistent among all chains
    ///                 - amountIn: amount of bridge token to start with, before the bridge fee is applied
    /// @param tokenOut  Token user wants to receive on destination chain
    /// @return destQuery  Structs that could be used as `destQuery` in SynapseRouter.
    ///                      minAmountOut and deadline fields will need to be adjusted based on the user settings.
    function getDestinationAmountOut(DestRequest memory request, address tokenOut)
    external
    view
    returns (SwapQuery memory destQuery);

    /// @notice Finds the best path between `tokenIn` and every supported bridge token from the given list,
    /// treating the swap as "origin swap", without putting any restrictions on the swap.
    /// @dev Intended for off-chain queries. Will NOT revert if any of the tokens are not supported, instead will return an empty query for that symbol.
    /// Check (query.minAmountOut != 0): this is true only if the swap is possible and bridge token is supported.
    /// The returned queries with minAmountOut != 0 could be used as `originQuery` with SynapseRouter.
    /// Note: it is possible to form a SwapQuery off-chain using alternative SwapAdapter for the origin swap.
    /// @param tokenIn       Initial token that user wants to bridge/swap
    /// @param tokenSymbol  Symbol representing bridge tokens
    /// @param amountIn      Amount of tokens user wants to bridge/swap
    /// @return originQuery    Structs that could be used as `originQuery` in SynapseRouter.
    ///                          minAmountOut and deadline fields will need to be adjusted based on the user settings.
    function getOriginAmountOut(
        address tokenIn,
        string memory tokenSymbol,
        uint256 amountIn
    ) external view returns (SwapQuery memory originQuery);
}

// contracts/router/libs/UniversalToken.sol

library UniversalTokenLib {
    using SafeERC20 for IERC20;

    address internal constant ETH_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

    /// @notice Transfers tokens to the given account. Reverts if transfer is not successful.
    /// @dev This might trigger fallback, if ETH is transferred to the contract.
    /// Make sure this can not lead to reentrancy attacks.
    function universalTransfer(
        address token,
        address to,
        uint256 value
    ) internal {
        // Don't do anything, if need to send tokens to this address
        if (to == address(this)) return;
        if (token == ETH_ADDRESS) {
            /// @dev Note: this can potentially lead to executing code in `to`.
            // solhint-disable-next-line avoid-low-level-calls
            (bool success, ) = to.call{value: value}("");
            require(success, "ETH transfer failed");
        } else {
            IERC20(token).safeTransfer(to, value);
        }
    }

    /// @notice Issues an infinite allowance to the spender, if the current allowance is insufficient
    /// to spend the given amount.
    function universalApproveInfinity(
        address token,
        address spender,
        uint256 amountToSpend
    ) internal {
        // ETH Chad doesn't require your approval
        if (token == ETH_ADDRESS) return;
        // No-op if allowance is already sufficient
        uint256 allowance = IERC20(token).allowance(address(this), spender);
        if (allowance >= amountToSpend) return;
        // Otherwise, reset approval to 0 and set to max allowance
        if (allowance > 0) IERC20(token).safeApprove(spender, 0);
        IERC20(token).safeApprove(spender, type(uint256).max);
    }

    /// @notice Returns the balance of the given token (or native ETH) for the given account.
    function universalBalanceOf(address token, address account) internal view returns (uint256) {
        if (token == ETH_ADDRESS) {
            return account.balance;
        } else {
            return IERC20(token).balanceOf(account);
        }
    }

    /// @dev Checks that token is a contract and not ETH_ADDRESS.
    function assertIsContract(address token) internal view {
        // Check that ETH_ADDRESS was not used (in case this is a predeploy on any of the chains)
        if (token == UniversalTokenLib.ETH_ADDRESS) revert TokenNotContract();
        // Check that token is not an EOA
        if (token.code.length == 0) revert TokenNotContract();
    }
}

// contracts/router/adapters/DefaultAdapter.sol

contract DefaultAdapter is IRouterAdapter {
    using SafeERC20 for IERC20;
    using UniversalTokenLib for address;

    /// @notice Enable this contract to receive Ether when withdrawing from WETH.
    /// @dev Consider implementing rescue functions to withdraw Ether from this contract.
    receive() external payable {}

    /// @inheritdoc IRouterAdapter
    function adapterSwap(
        address recipient,
        address tokenIn,
        uint256 amountIn,
        address tokenOut,
        bytes memory rawParams
    ) external payable returns (uint256 amountOut) {
        return _adapterSwap(recipient, tokenIn, amountIn, tokenOut, rawParams);
    }

    /// @dev Internal logic for doing a tokenIn -> tokenOut swap.
    /// Note: `tokenIn` is assumed to have already been transferred to this contract.
    function _adapterSwap(
        address recipient,
        address tokenIn,
        uint256 amountIn,
        address tokenOut,
        bytes memory rawParams
    ) internal virtual returns (uint256 amountOut) {
        // We define a few phases for the whole Adapter's swap process.
        // (?) means the phase is optional.
        // (!) means the phase is mandatory.

        // PHASE 0(!): CHECK ALL THE PARAMS
        DefaultParams memory params = _checkParams(tokenIn, tokenOut, rawParams);

        // PHASE 1(?): WRAP RECEIVED ETH INTO WETH
        tokenIn = _wrapReceivedETH(tokenIn, amountIn, tokenOut, params);
        // After PHASE 1 this contract has `amountIn` worth of `tokenIn`, tokenIn != ETH_ADDRESS

        // PHASE 2(?): PREPARE TO UNWRAP SWAPPED WETH
        address tokenSwapTo = _deriveTokenSwapTo(tokenIn, tokenOut, params);
        // We need to perform tokenIn -> tokenSwapTo action in PHASE 3.
        // if tokenOut == ETH_ADDRESS, we need to unwrap WETH in PHASE 4.
        // Recipient will receive `tokenOut` in PHASE 5.

        // PHASE 3(?): PERFORM A REQUESTED SWAP
        amountOut = _performPoolAction(tokenIn, amountIn, tokenSwapTo, params);
        // After PHASE 3 this contract has `amountOut` worth of `tokenSwapTo`, tokenSwapTo != ETH_ADDRESS

        // PHASE 4(?): UNWRAP SWAPPED WETH
        // Check if the final token is native ETH
        if (tokenOut == UniversalTokenLib.ETH_ADDRESS) {
            // PHASE 2: WETH address was stored as `tokenSwapTo`
            _unwrapETH(tokenSwapTo, amountOut);
        }

        // PHASE 5(!): TRANSFER SWAPPED TOKENS TO RECIPIENT
        // Note: this will transfer native ETH, if tokenOut == ETH_ADDRESS
        // Note: this is a no-op if recipient == address(this)
        tokenOut.universalTransfer(recipient, amountOut);
    }

    /// @dev Checks the params and decodes them into a struct.
    function _checkParams(
        address tokenIn,
        address tokenOut,
        bytes memory rawParams
    ) internal pure returns (DefaultParams memory params) {
        if (tokenIn == tokenOut) revert TokensIdentical();
        // Decode params for swapping via a Default pool
        params = abi.decode(rawParams, (DefaultParams));
        // Swap pool should exist, if action other than HandleEth was requested
        if (params.pool == address(0) && params.action != Action.HandleEth) revert PoolNotFound();
    }

    /// @dev Wraps native ETH into WETH, if requested.
    /// Returns the address of the token this contract ends up with.
    function _wrapReceivedETH(
        address tokenIn,
        uint256 amountIn,
        address tokenOut,
        DefaultParams memory params
    ) internal returns (address wrappedTokenIn) {
        // tokenIn was already transferred to this contract, check if we start from native ETH
        if (tokenIn == UniversalTokenLib.ETH_ADDRESS) {
            // Determine WETH address: this is either tokenOut (if no swap is needed),
            // or a pool token with index `tokenIndexFrom` (if swap is needed).
            wrappedTokenIn = _deriveWethAddress({token: tokenOut, params: params, isTokenFromWeth: true});
            // Wrap ETH into WETH and leave it in this contract
            _wrapETH(wrappedTokenIn, amountIn);
        } else {
            wrappedTokenIn = tokenIn;
            // For ERC20 tokens msg.value should be zero
            if (msg.value != 0) revert MsgValueIncorrect();
        }
    }

    /// @dev Derives the address of token to be received after an action defined in `params`.
    function _deriveTokenSwapTo(
        address tokenIn,
        address tokenOut,
        DefaultParams memory params
    ) internal view returns (address tokenSwapTo) {
        // Check if swap to native ETH was requested
        if (tokenOut == UniversalTokenLib.ETH_ADDRESS) {
            // Determine WETH address: this is either tokenIn (if no swap is needed),
            // or a pool token with index `tokenIndexTo` (if swap is needed).
            tokenSwapTo = _deriveWethAddress({token: tokenIn, params: params, isTokenFromWeth: false});
        } else {
            tokenSwapTo = tokenOut;
        }
    }

    /// @dev Performs an action defined in `params` and returns the amount of `tokenSwapTo` received.
    function _performPoolAction(
        address tokenIn,
        uint256 amountIn,
        address tokenSwapTo,
        DefaultParams memory params
    ) internal returns (uint256 amountOut) {
        // Determine if we need to perform a swap
        if (params.action == Action.HandleEth) {
            // If no swap is required, amountOut doesn't change
            amountOut = amountIn;
        } else {
            // Record balance before the swap
            amountOut = IERC20(tokenSwapTo).balanceOf(address(this));
            // Approve the pool for spending exactly `amountIn` of `tokenIn`
            IERC20(tokenIn).safeIncreaseAllowance(params.pool, amountIn);
            if (params.action == Action.Swap) {
                _swap(params.pool, params, amountIn, tokenSwapTo);
            } else if (params.action == Action.AddLiquidity) {
                _addLiquidity(params.pool, params, amountIn, tokenSwapTo);
            } else {
                // The only remaining action is RemoveLiquidity
                _removeLiquidity(params.pool, params, amountIn, tokenSwapTo);
            }
            // Use the difference between the balance after the swap and the recorded balance as `amountOut`
            amountOut = IERC20(tokenSwapTo).balanceOf(address(this)) - amountOut;
        }
    }

    // ═══════════════════════════════════════ INTERNAL LOGIC: SWAP ACTIONS ════════════════════════════════════════════

    /// @dev Performs a swap through the given pool.
    /// Note: The pool should be already approved for spending `tokenIn`.
    function _swap(
        address pool,
        DefaultParams memory params,
        uint256 amountIn,
        address tokenOut
    ) internal {
        // tokenOut should match the "swap to" token
        if (IDefaultPool(pool).getToken(params.tokenIndexTo) != tokenOut) revert TokenAddressMismatch();
        // amountOut and deadline are not checked in RouterAdapter
        IDefaultPool(pool).swap({
            tokenIndexFrom: params.tokenIndexFrom,
            tokenIndexTo: params.tokenIndexTo,
            dx: amountIn,
            minDy: 0,
            deadline: type(uint256).max
        });
    }

    /// @dev Adds liquidity in a form of a single token to the given pool.
    /// Note: The pool should be already approved for spending `tokenIn`.
    function _addLiquidity(
        address pool,
        DefaultParams memory params,
        uint256 amountIn,
        address tokenOut
    ) internal {
        uint256 numTokens = _getPoolNumTokens(pool);
        address lpToken = _getPoolLPToken(pool);
        // tokenOut should match the LP token
        if (lpToken != tokenOut) revert TokenAddressMismatch();
        uint256[] memory amounts = new uint256[](numTokens);
        amounts[params.tokenIndexFrom] = amountIn;
        // amountOut and deadline are not checked in RouterAdapter
        IDefaultExtendedPool(pool).addLiquidity({amounts: amounts, minToMint: 0, deadline: type(uint256).max});
    }

    /// @dev Removes liquidity in a form of a single token from the given pool.
    /// Note: The pool should be already approved for spending `tokenIn`.
    function _removeLiquidity(
        address pool,
        DefaultParams memory params,
        uint256 amountIn,
        address tokenOut
    ) internal {
        // tokenOut should match the "swap to" token
        if (IDefaultPool(pool).getToken(params.tokenIndexTo) != tokenOut) revert TokenAddressMismatch();
        // amountOut and deadline are not checked in RouterAdapter
        IDefaultExtendedPool(pool).removeLiquidityOneToken({
            tokenAmount: amountIn,
            tokenIndex: params.tokenIndexTo,
            minAmount: 0,
            deadline: type(uint256).max
        });
    }

    // ═════════════════════════════════════════ INTERNAL LOGIC: POOL LENS ═════════════════════════════════════════════

    /// @dev Returns the LP token address of the given pool.
    function _getPoolLPToken(address pool) internal view returns (address lpToken) {
        (, , , , , , lpToken) = IDefaultExtendedPool(pool).swapStorage();
    }

    /// @dev Returns the number of tokens in the given pool.
    function _getPoolNumTokens(address pool) internal view returns (uint256 numTokens) {
        // Iterate over all tokens in the pool until the end is reached
        for (uint8 index = 0; ; ++index) {
            try IDefaultPool(pool).getToken(index) returns (address) {} catch {
                // End of pool reached
                numTokens = index;
                break;
            }
        }
    }

    /// @dev Returns the tokens in the given pool.
    function _getPoolTokens(address pool) internal view returns (address[] memory tokens) {
        uint256 numTokens = _getPoolNumTokens(pool);
        tokens = new address[](numTokens);
        for (uint8 i = 0; i < numTokens; ++i) {
            // This will not revert because we already know the number of tokens in the pool
            tokens[i] = IDefaultPool(pool).getToken(i);
        }
    }

    /// @dev Returns the quote for a swap through the given pool.
    /// Note: will return 0 on invalid swaps.
    function _getPoolSwapQuote(
        address pool,
        uint8 tokenIndexFrom,
        uint8 tokenIndexTo,
        uint256 amountIn
    ) internal view returns (uint256 amountOut) {
        try IDefaultPool(pool).calculateSwap(tokenIndexFrom, tokenIndexTo, amountIn) returns (uint256 dy) {
            amountOut = dy;
        } catch {
            // Return 0 instead of reverting
            amountOut = 0;
        }
    }

    // ════════════════════════════════════════ INTERNAL LOGIC: ETH <> WETH ════════════════════════════════════════════

    /// @dev Wraps ETH into WETH.
    function _wrapETH(address weth, uint256 amount) internal {
        if (amount != msg.value) revert MsgValueIncorrect();
        // Deposit in order to have WETH in this contract
        IWETH9(weth).deposit{value: amount}();
    }

    /// @dev Unwraps WETH into ETH.
    function _unwrapETH(address weth, uint256 amount) internal {
        // Withdraw ETH to this contract
        IWETH9(weth).withdraw(amount);
    }

    /// @dev Derives WETH address from swap parameters.
    function _deriveWethAddress(
        address token,
        DefaultParams memory params,
        bool isTokenFromWeth
    ) internal view returns (address weth) {
        if (params.action == Action.HandleEth) {
            // If we only need to wrap/unwrap ETH, WETH address should be specified as the other token
            weth = token;
        } else {
            // Otherwise, we need to get WETH address from the liquidity pool
            weth = address(
                IDefaultPool(params.pool).getToken(isTokenFromWeth ? params.tokenIndexFrom : params.tokenIndexTo)
            );
        }
    }
}

// contracts/router/DefaultRouter.sol

/// @title DefaultRouter
/// @notice Base contract for all Synapse Routers, that is able to natively work with Default Pools
/// due to the fact that it inherits from DefaultAdapter.
abstract contract DefaultRouter is DefaultAdapter {
    using SafeERC20 for IERC20;
    using UniversalTokenLib for address;

    /// @dev Performs a "swap from tokenIn" following instructions from `query`.
    /// `query` will include the router adapter to use, and the exact type of "tokenIn -> tokenOut swap"
    /// should be encoded in `query.rawParams`.
    function _doSwap(
        address recipient,
        address tokenIn,
        uint256 amountIn,
        SwapQuery memory query
    ) internal returns (address tokenOut, uint256 amountOut) {
        // First, check the deadline for the swap
        // solhint-disable-next-line not-rely-on-time
        if (block.timestamp > query.deadline) revert DeadlineExceeded();
        // Pull initial token from the user to specified router adapter
        amountIn = _pullToken(query.routerAdapter, tokenIn, amountIn);
        tokenOut = query.tokenOut;
        address routerAdapter = query.routerAdapter;
        if (routerAdapter == address(this)) {
            // If the router adapter is this contract, we can perform the swap directly and trust the result
            amountOut = _adapterSwap(recipient, tokenIn, amountIn, tokenOut, query.rawParams);
        } else {
            // Otherwise, we need to call the router adapter. Adapters are permissionless, so we verify the result
            // Record tokenOut balance before swap
            amountOut = tokenOut.universalBalanceOf(recipient);
            IRouterAdapter(routerAdapter).adapterSwap{value: msg.value}({
                recipient: recipient,
                tokenIn: tokenIn,
                amountIn: amountIn,
                tokenOut: tokenOut,
                rawParams: query.rawParams
            });
            // Use the difference between the recorded balance and the current balance as the amountOut
            amountOut = tokenOut.universalBalanceOf(recipient) - amountOut;
        }
        // Finally, check that the recipient received at least as much as they wanted
        if (amountOut < query.minAmountOut) revert InsufficientOutputAmount();
    }

    /// @dev Pulls a requested token from the user to the requested recipient.
    /// Or, if msg.value was provided, check that ETH_ADDRESS was used and msg.value is correct.
    function _pullToken(
        address recipient,
        address token,
        uint256 amount
    ) internal returns (uint256 amountPulled) {
        if (msg.value == 0) {
            token.assertIsContract();
            // Record token balance before transfer
            amountPulled = IERC20(token).balanceOf(recipient);
            // Token needs to be pulled only if msg.value is zero
            // This way user can specify WETH as the origin asset
            IERC20(token).safeTransferFrom(msg.sender, recipient, amount);
            // Use the difference between the recorded balance and the current balance as the amountPulled
            amountPulled = IERC20(token).balanceOf(recipient) - amountPulled;
        } else {
            // Otherwise, we need to check that ETH was specified
            if (token != UniversalTokenLib.ETH_ADDRESS) revert TokenNotETH();
            // And that amount matches msg.value
            if (amount != msg.value) revert MsgValueIncorrect();
            // We will forward msg.value in the external call later, if recipient is not this contract
            amountPulled = msg.value;
        }
    }
}

// contracts/router/SynapseRouterV2.sol

contract SynapseRouterV2 is IRouterV2, DefaultRouter, Ownable {
    using Address for address;
    using Arrays for BridgeToken[][];
    using Arrays for BridgeToken[];
    using Arrays for address[][];
    using Arrays for address[];
    using EnumerableMap for EnumerableMap.UintToAddressMap;
    using SafeERC20 for IERC20;

    /// @notice swap quoter
    ISwapQuoterV2 public swapQuoter;

    /// @notice Enumerable map of all connected bridge modules
    EnumerableMap.UintToAddressMap internal _bridgeModules;

    event QuoterSet(address oldSwapQuoter, address newSwapQuoter);
    event ModuleConnected(bytes32 indexed moduleId, address bridgeModule);
    event ModuleUpdated(bytes32 indexed moduleId, address oldBridgeModule, address newBridgeModule);
    event ModuleDisconnected(bytes32 indexed moduleId);

    error SynapseRouterV2__ModuleExists();
    error SynapseRouterV2__ModuleNotExists();
    error SynapseRouterV2__ModuleInvalid();
    error SynapseRouterV2__QueryEmpty();

    /// @inheritdoc IRouterV2
    function bridgeViaSynapse(
        address to,
        uint256 chainId,
        bytes32 moduleId,
        address token,
        uint256 amount,
        SwapQuery memory originQuery,
        SwapQuery memory destQuery
    ) external payable {
        address bridgeModule = idToModule(moduleId);

        // pull (and possibly swap) token into router
        if (originQuery.hasAdapter()) {
            (token, amount) = _doSwap(address(this), token, amount, originQuery);
        } else {
            amount = _pullToken(address(this), token, amount);
        }

        // delegate bridge call to module
        // @dev delegatecall should approve to spend
        bytes memory payload = abi.encodeWithSelector(
            IBridgeModule.delegateBridge.selector,
            to,
            chainId,
            token,
            amount,
            destQuery
        );
        bridgeModule.functionDelegateCall(payload); // bubbles up the error, but nothing to be returned
    }

    /// @inheritdoc IRouterV2
    function swap(
        address to,
        address token,
        uint256 amount,
        SwapQuery memory query
    ) external payable returns (uint256 amountOut) {
        if (!query.hasAdapter()) revert SynapseRouterV2__QueryEmpty();
        (, amountOut) = _doSwap(to, token, amount, query);
    }

    /// @inheritdoc IRouterV2
    function setSwapQuoter(ISwapQuoterV2 _swapQuoter) external onlyOwner {
        emit QuoterSet(address(swapQuoter), address(_swapQuoter));
        swapQuoter = _swapQuoter;
    }

    /// @inheritdoc IRouterV2
    function setAllowance(
        address token,
        address spender,
        uint256 amount
    ) external onlyOwner {
        IERC20(token).safeApprove(spender, amount);
    }

    /// @inheritdoc IRouterV2
    function connectBridgeModule(bytes32 moduleId, address bridgeModule) external onlyOwner {
        if (moduleId == bytes32(0) || bridgeModule == address(0)) revert SynapseRouterV2__ModuleInvalid();
        if (_hasModule(moduleId)) revert SynapseRouterV2__ModuleExists();

        _bridgeModules.set(uint256(moduleId), bridgeModule);
        emit ModuleConnected(moduleId, bridgeModule);
    }

    /// @inheritdoc IRouterV2
    function updateBridgeModule(bytes32 moduleId, address bridgeModule) external onlyOwner {
        if (bridgeModule == address(0)) revert SynapseRouterV2__ModuleInvalid();
        if (!_hasModule(moduleId)) revert SynapseRouterV2__ModuleNotExists();

        address module = _bridgeModules.get(uint256(moduleId));
        _bridgeModules.set(uint256(moduleId), bridgeModule);

        emit ModuleUpdated(moduleId, module, bridgeModule);
    }

    /// @inheritdoc IRouterV2
    function disconnectBridgeModule(bytes32 moduleId) external onlyOwner {
        if (!_hasModule(moduleId)) revert SynapseRouterV2__ModuleNotExists();

        _bridgeModules.remove(uint256(moduleId));
        emit ModuleDisconnected(moduleId);
    }

    /// @inheritdoc IRouterV2
    function idToModule(bytes32 moduleId) public view returns (address bridgeModule) {
        if (!_hasModule(moduleId)) revert SynapseRouterV2__ModuleNotExists();
        bridgeModule = _bridgeModules.get(uint256(moduleId));
    }

    /// @inheritdoc IRouterV2
    function moduleToId(address bridgeModule) public view returns (bytes32 moduleId) {
        uint256 len = _bridgeModules.length();
        for (uint256 i = 0; i < len; ++i) {
            (uint256 key, address module) = _bridgeModules.at(i);
            if (module == bridgeModule) {
                moduleId = bytes32(key);
                break;
            }
        }
        if (moduleId == bytes32(0)) revert SynapseRouterV2__ModuleNotExists();
    }

    /// @inheritdoc IRouterV2
    function getBridgeTokens() public view returns (BridgeToken[] memory) {
        uint256 len = _bridgeModules.length();
        BridgeToken[][] memory unflattened = new BridgeToken[][](len);

        uint256 count;
        for (uint256 i = 0; i < len; ++i) {
            (, address bridgeModule) = _bridgeModules.at(i);
            unflattened[i] = IBridgeModule(bridgeModule).getBridgeTokens();
            count += unflattened[i].length;
        }

        // flatten into bridge tokens array
        return unflattened.flatten(count);
    }

    /// @inheritdoc IRouterV2
    function getDestinationBridgeTokens(address tokenOut) external view returns (BridgeToken[] memory destTokens) {
        destTokens = _getConnectedBridgeTokens(tokenOut, false);
    }

    /// @inheritdoc IRouterV2
    function getOriginBridgeTokens(address tokenIn) external view returns (BridgeToken[] memory originTokens) {
        originTokens = _getConnectedBridgeTokens(tokenIn, true);
    }

    /// @inheritdoc IRouterV2
    function getSupportedTokens() external view returns (address[] memory supportedTokens) {
        // get tokens in each quoter pool
        Pool[] memory pools = swapQuoter.allPools();

        // init unflattened for aggregating supported tokens
        address[][] memory unflattened = new address[][](pools.length + 1);
        uint256 count;

        // last index of unflattened dedicated to bridge tokens
        unflattened[pools.length] = getBridgeTokens().tokens();
        count += unflattened[pools.length].length;

        // fill pool tokens up to but not including last index of unflattened
        for (uint256 i = 0; i < pools.length; ++i) {
            Pool memory pool = pools[i];
            unflattened[i] = new address[](pool.tokens.length);

            // whether pool.tokens does contain a bridge token
            // @dev need to initialize with contains pool LP token to catch edge case
            // @dev of Ethereum Stablepool having LP token as bridge token
            bool does = unflattened[pools.length].contains(pool.lpToken);
            for (uint256 j = 0; j < pool.tokens.length; ++j) {
                // optimistically add pool token to list
                unflattened[i][j] = pool.tokens[j].token;

                // check whether pool token is a bridge token if haven't found one prior
                // @dev remember last index of unflattened has all the addresses of the bridge tokens
                if (!does) does = unflattened[pools.length].contains(pool.tokens[j].token);
            }

            // zero out if no bridge token
            if (!does) delete unflattened[i];
            else count += pool.tokens.length;
        }

        // flatten into supported tokens and filter out duplicates
        supportedTokens = unflattened.flatten(count).unique();

        // add native weth if in supported list
        if (supportedTokens.contains(swapQuoter.weth())) return supportedTokens.append(UniversalTokenLib.ETH_ADDRESS);
    }

    /// @inheritdoc IRouterV2
    function getDestinationAmountOut(DestRequest memory request, address tokenOut)
    external
    view
    returns (SwapQuery memory destQuery)
    {
        (address token, uint256 actionMask, address bridgeModule) = _getTokenAndActionMaskFromSymbol(request.symbol);
        if (token == address(0)) return destQuery; // empty

        // account for bridge fees in amountIn
        bool isSwap = !(token == tokenOut || (tokenOut == UniversalTokenLib.ETH_ADDRESS && token == swapQuoter.weth()));
        uint256 amountIn = _calculateBridgeAmountIn(bridgeModule, token, request.amountIn, isSwap);
        if (amountIn == 0) return destQuery; // empty

        // query the quoter
        LimitedToken memory tokenIn = LimitedToken({actionMask: actionMask, token: token});
        destQuery = swapQuoter.getAmountOut(tokenIn, tokenOut, amountIn);
    }

    /// @inheritdoc IRouterV2
    function getOriginAmountOut(
        address tokenIn,
        string memory tokenSymbol,
        uint256 amountIn
    ) external view returns (SwapQuery memory originQuery) {
        (address tokenOut, , address bridgeModule) = _getTokenAndActionMaskFromSymbol(tokenSymbol);
        if (tokenOut == address(0)) return originQuery; // empty

        // query the quoter
        LimitedToken memory _tokenIn = LimitedToken({actionMask: ActionLib.allActions(), token: tokenIn});
        SwapQuery memory query = swapQuoter.getAmountOut(_tokenIn, tokenOut, amountIn);

        // check max amount can bridge
        uint256 maxAmountOut = IBridgeModule(bridgeModule).getMaxBridgedAmount(tokenOut);
        if (query.minAmountOut > maxAmountOut) return originQuery; // empty

        // set in return array
        originQuery = query;
    }

    /// @notice Checks whether module ID has already been connected to router
    function _hasModule(bytes32 moduleId) internal view returns (bool) {
        return _bridgeModules.contains(uint256(moduleId));
    }

    /// @notice Searches all bridge modules to get the token address from the unique bridge symbol
    /// @dev Returns zero value for unknown symbols (does not revert)
    /// @param symbol Symbol of the supported bridge token
    function _getTokenAndActionMaskFromSymbol(string memory symbol)
    internal
    view
    returns (
        address token,
        uint256 actionMask,
        address bridgeModule
    )
    {
        uint256 len = _bridgeModules.length();
        for (uint256 i = 0; i < len; ++i) {
            (, address _bridgeModule) = _bridgeModules.at(i);
            token = IBridgeModule(_bridgeModule).symbolToToken(symbol);
            if (token != address(0)) {
                actionMask = IBridgeModule(_bridgeModule).tokenToActionMask(token);
                bridgeModule = _bridgeModule;
                break;
            }
        }
    }

    /// @notice Gets all connected bridge tokens to the given token
    /// @param token The token to connect bridge tokens with
    /// @param origin Whether gathering on origin or destination chain
    /// @param connected The connected bridge tokens. If origin == True, then returns origin
    ///                  bridge tokens. If origin == False, returns dest bridge tokens.
    function _getConnectedBridgeTokens(address token, bool origin)
    internal
    view
    returns (BridgeToken[] memory connected)
    {
        uint256 len = _bridgeModules.length();
        BridgeToken[][] memory unflattened = new BridgeToken[][](len);

        uint256 count;
        for (uint256 i = 0; i < len; ++i) {
            (, address bridgeModule) = _bridgeModules.at(i);
            BridgeToken[] memory bridgeTokens = IBridgeModule(bridgeModule).getBridgeTokens();

            // assemble limited token format for quoter call
            uint256 amountFound;
            bool[] memory isConnected = new bool[](bridgeTokens.length);
            for (uint256 j = 0; j < bridgeTokens.length; ++j) {
                LimitedToken memory _tokenIn = origin
                    ? LimitedToken({actionMask: ActionLib.allActions(), token: token}) // origin bridge tokens
                    : LimitedToken({
                        actionMask: IBridgeModule(bridgeModule).tokenToActionMask(bridgeTokens[j].token),
                        token: bridgeTokens[j].token
                    }); // dest bridge tokens
                address _tokenOut = origin ? bridgeTokens[j].token : token;
                isConnected[j] = swapQuoter.areConnectedTokens(_tokenIn, _tokenOut);
                if (isConnected[j]) amountFound++;
            }

            // push to unflattened tokens if bridge token connected to given token
            unflattened[i] = new BridgeToken[](amountFound);
            count += amountFound;

            uint256 k;
            for (uint256 j = 0; j < bridgeTokens.length; ++j) {
                if (isConnected[j]) {
                    unflattened[i][k] = bridgeTokens[j];
                    k++;
                }
            }
        }

        // flatten into connected tokens
        connected = unflattened.flatten(count);
    }

    /// @notice Calculates amount of bridge token in accounting for bridge fees
    /// @dev Returns zero value if fee too big (does not revert)
    /// @param token    Address of the bridging token
    /// @param amount   Amount in before fees
    /// @param isSwap   Whether the user provided swap details for converting the bridge token
    ///                 to the final token on this chain
    function _calculateBridgeAmountIn(
        address bridgeModule,
        address token,
        uint256 amount,
        bool isSwap
    ) internal view returns (uint256 amount_) {
        uint256 feeAmount = IBridgeModule(bridgeModule).calculateFeeAmount(token, amount, isSwap);
        if (feeAmount < amount) amount_ = amount - feeAmount;
    }
}
