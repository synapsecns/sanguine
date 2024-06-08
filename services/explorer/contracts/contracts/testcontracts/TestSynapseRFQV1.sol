// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IFastBridge {
  struct BridgeTransaction {
    uint32 originChainId;
    uint32 destChainId;
    address originSender; // user (origin)
    address destRecipient; // user (dest)
    address originToken;
    address destToken;
    uint256 originAmount; // amount in on origin bridge less originFeeAmount
    uint256 destAmount;
    uint256 originFeeAmount;
    bool sendChainGas;
    uint256 deadline; // user specified deadline for destination relay
    uint256 nonce;
  }

  struct BridgeProof {
    uint96 timestamp;
    address relayer;
  }

  // ============ Events ============

  event BridgeRequested(
    bytes32 indexed transactionId,
    address indexed sender,
    bytes request,
    uint32 destChainId,
    address originToken,
    address destToken,
    uint256 originAmount,
    uint256 destAmount,
    bool sendChainGas
  );
  event BridgeRelayed(
    bytes32 indexed transactionId,
    address indexed relayer,
    address indexed to,
    uint32 originChainId,
    address originToken,
    address destToken,
    uint256 originAmount,
    uint256 destAmount,
    uint256 chainGasAmount
  );
  event BridgeProofProvided(
    bytes32 indexed transactionId,
    address indexed relayer,
    bytes32 transactionHash
  );
  event BridgeProofDisputed(
    bytes32 indexed transactionId,
    address indexed relayer
  );
  event BridgeDepositClaimed(
    bytes32 indexed transactionId,
    address indexed relayer,
    address indexed to,
    address token,
    uint256 amount
  );
  event BridgeDepositRefunded(
    bytes32 indexed transactionId,
    address indexed to,
    address token,
    uint256 amount
  );

  // ============ Methods ============

  struct BridgeParams {
    uint32 dstChainId;
    address sender;
    address to;
    address originToken;
    address destToken;
    uint256 originAmount; // should include protocol fee (if any)
    uint256 destAmount; // should include relayer fee
    bool sendChainGas;
    uint256 deadline;
  }

  /// @notice Initiates bridge on origin chain to be relayed by off-chain relayer
  /// @param params The parameters required to bridge
  function bridge(BridgeParams memory params) external payable;

  /// @notice Relays destination side of bridge transaction by off-chain relayer
  /// @param request The encoded bridge transaction to relay on destination chain
  function relay(bytes memory request) external payable;

  /// @notice Provides proof on origin side that relayer provided funds on destination side of bridge transaction
  /// @param request The encoded bridge transaction to prove on origin chain
  /// @param destTxHash The destination tx hash proving bridge transaction was relayed
  function prove(bytes memory request, bytes32 destTxHash) external;

  /// @notice Completes bridge transaction on origin chain by claiming originally deposited capital
  /// @param request The encoded bridge transaction to claim on origin chain
  /// @param to The recipient address of the funds
  function claim(bytes memory request, address to) external;

  /// @notice Disputes an outstanding proof in case relayer provided dest chain tx is invalid
  /// @param transactionId The transaction id associated with the encoded bridge transaction to dispute
  function dispute(bytes32 transactionId) external;

  /// @notice Refunds an outstanding bridge transaction in case optimistic bridging failed
  /// @param request The encoded bridge transaction to refund
  function refund(bytes memory request) external;

  // ============ Views ============

  /// @notice Decodes bridge request into a bridge transaction
  /// @param request The bridge request to decode
  function getBridgeTransaction(
    bytes memory request
  ) external pure returns (BridgeTransaction memory);

  /// @notice Checks if the dispute period has passed so bridge deposit can be claimed
  /// @param transactionId The transaction id associated with the encoded bridge transaction to check
  /// @param relayer The address of the relayer attempting to claim
  function canClaim(
    bytes32 transactionId,
    address relayer
  ) external view returns (bool);
}

// File: mainnet/0x4983DB49336fD4f95e864aB6DA9135e057EF0be1/contracts/interfaces/IAdmin.sol

pragma solidity ^0.8.0;

interface IAdmin {
  // ============ Events ============

  event RelayerAdded(address relayer);
  event RelayerRemoved(address relayer);

  event GuardAdded(address guard);
  event GuardRemoved(address guard);

  event GovernorAdded(address governor);
  event GovernorRemoved(address governor);

  event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
  event FeesSwept(address token, address recipient, uint256 amount);

  event ChainGasAmountUpdated(
    uint256 oldChainGasAmount,
    uint256 newChainGasAmount
  );

  // ============ Methods ============

  function addRelayer(address _relayer) external;

  function removeRelayer(address _relayer) external;

  function addGuard(address _guard) external;

  function removeGuard(address _guard) external;

  function addGovernor(address _governor) external;

  function removeGovernor(address _governor) external;

  function setProtocolFeeRate(uint256 newFeeRate) external;

  function sweepProtocolFees(address token, address recipient) external;

  function setChainGasAmount(uint256 newChainGasAmount) external;
}

// File: @openzeppelin/contracts/utils/introspection/IERC165.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/introspection/IERC165.sol)

pragma solidity ^0.8.20;

/**
 * @dev Interface of the ERC165 standard, as defined in the
 * https://eips.ethereum.org/EIPS/eip-165[EIP].
 *
 * Implementers can declare support of contract interfaces, which can then be
 * queried by others ({ERC165Checker}).
 *
 * For an implementation, see {ERC165}.
 */
interface IERC165 {
  /**
   * @dev Returns true if this contract implements the interface defined by
   * `interfaceId`. See the corresponding
   * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified[EIP section]
   * to learn more about how these ids are created.
   *
   * This function call must use less than 30 000 gas.
   */
  function supportsInterface(bytes4 interfaceId) external view returns (bool);
}

// File: @openzeppelin/contracts/utils/introspection/ERC165.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/introspection/ERC165.sol)

pragma solidity ^0.8.20;

/**
 * @dev Implementation of the {IERC165} interface.
 *
 * Contracts that want to implement ERC165 should inherit from this contract and override {supportsInterface} to check
 * for the additional interface id that will be supported. For example:
 *
 * ```solidity
 * function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
 *     return interfaceId == type(MyInterface).interfaceId || super.supportsInterface(interfaceId);
 * }
 * ```
 */
abstract contract ERC165 is IERC165 {
  /**
   * @dev See {IERC165-supportsInterface}.
   */
  function supportsInterface(
    bytes4 interfaceId
  ) public view virtual returns (bool) {
    return interfaceId == type(IERC165).interfaceId;
  }
}

// File: @openzeppelin/contracts/utils/Context.sol

// OpenZeppelin Contracts (last updated v5.0.1) (utils/Context.sol)

pragma solidity ^0.8.20;

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

  function _contextSuffixLength() internal view virtual returns (uint256) {
    return 0;
  }
}

// File: @openzeppelin/contracts/access/IAccessControl.sol

// OpenZeppelin Contracts (last updated v5.0.0) (access/IAccessControl.sol)

pragma solidity ^0.8.20;

/**
 * @dev External interface of AccessControl declared to support ERC165 detection.
 */
interface IAccessControl {
  /**
   * @dev The `account` is missing a role.
   */
  error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

  /**
   * @dev The caller of a function is not the expected one.
   *
   * NOTE: Don't confuse with {AccessControlUnauthorizedAccount}.
   */
  error AccessControlBadConfirmation();

  /**
   * @dev Emitted when `newAdminRole` is set as ``role``'s admin role, replacing `previousAdminRole`
   *
   * `DEFAULT_ADMIN_ROLE` is the starting admin for all roles, despite
   * {RoleAdminChanged} not being emitted signaling this.
   */
  event RoleAdminChanged(
    bytes32 indexed role,
    bytes32 indexed previousAdminRole,
    bytes32 indexed newAdminRole
  );

  /**
   * @dev Emitted when `account` is granted `role`.
   *
   * `sender` is the account that originated the contract call, an admin role
   * bearer except when using {AccessControl-_setupRole}.
   */
  event RoleGranted(
    bytes32 indexed role,
    address indexed account,
    address indexed sender
  );

  /**
   * @dev Emitted when `account` is revoked `role`.
   *
   * `sender` is the account that originated the contract call:
   *   - if using `revokeRole`, it is the admin role bearer
   *   - if using `renounceRole`, it is the role bearer (i.e. `account`)
   */
  event RoleRevoked(
    bytes32 indexed role,
    address indexed account,
    address indexed sender
  );

  /**
   * @dev Returns `true` if `account` has been granted `role`.
   */
  function hasRole(bytes32 role, address account) external view returns (bool);

  /**
   * @dev Returns the admin role that controls `role`. See {grantRole} and
   * {revokeRole}.
   *
   * To change a role's admin, use {AccessControl-_setRoleAdmin}.
   */
  function getRoleAdmin(bytes32 role) external view returns (bytes32);

  /**
   * @dev Grants `role` to `account`.
   *
   * If `account` had not been already granted `role`, emits a {RoleGranted}
   * event.
   *
   * Requirements:
   *
   * - the caller must have ``role``'s admin role.
   */
  function grantRole(bytes32 role, address account) external;

  /**
   * @dev Revokes `role` from `account`.
   *
   * If `account` had been granted `role`, emits a {RoleRevoked} event.
   *
   * Requirements:
   *
   * - the caller must have ``role``'s admin role.
   */
  function revokeRole(bytes32 role, address account) external;

  /**
   * @dev Revokes `role` from the calling account.
   *
   * Roles are often managed via {grantRole} and {revokeRole}: this function's
   * purpose is to provide a mechanism for accounts to lose their privileges
   * if they are compromised (such as when a trusted device is misplaced).
   *
   * If the calling account had been granted `role`, emits a {RoleRevoked}
   * event.
   *
   * Requirements:
   *
   * - the caller must be `callerConfirmation`.
   */
  function renounceRole(bytes32 role, address callerConfirmation) external;
}

// File: @openzeppelin/contracts/access/AccessControl.sol

// OpenZeppelin Contracts (last updated v5.0.0) (access/AccessControl.sol)

pragma solidity ^0.8.20;

/**
 * @dev Contract module that allows children to implement role-based access
 * control mechanisms. This is a lightweight version that doesn't allow enumerating role
 * members except through off-chain means by accessing the contract event logs. Some
 * applications may benefit from on-chain enumerability, for those cases see
 * {AccessControlEnumerable}.
 *
 * Roles are referred to by their `bytes32` identifier. These should be exposed
 * in the external API and be unique. The best way to achieve this is by
 * using `public constant` hash digests:
 *
 * ```solidity
 * bytes32 public constant MY_ROLE = keccak256("MY_ROLE");
 * ```
 *
 * Roles can be used to represent a set of permissions. To restrict access to a
 * function call, use {hasRole}:
 *
 * ```solidity
 * function foo() public {
 *     require(hasRole(MY_ROLE, msg.sender));
 *     ...
 * }
 * ```
 *
 * Roles can be granted and revoked dynamically via the {grantRole} and
 * {revokeRole} functions. Each role has an associated admin role, and only
 * accounts that have a role's admin role can call {grantRole} and {revokeRole}.
 *
 * By default, the admin role for all roles is `DEFAULT_ADMIN_ROLE`, which means
 * that only accounts with this role will be able to grant or revoke other
 * roles. More complex role relationships can be created by using
 * {_setRoleAdmin}.
 *
 * WARNING: The `DEFAULT_ADMIN_ROLE` is also its own admin: it has permission to
 * grant and revoke this role. Extra precautions should be taken to secure
 * accounts that have been granted it. We recommend using {AccessControlDefaultAdminRules}
 * to enforce additional security measures for this role.
 */
abstract contract AccessControl is Context, IAccessControl, ERC165 {
  struct RoleData {
    mapping(address account => bool) hasRole;
    bytes32 adminRole;
  }

  mapping(bytes32 role => RoleData) private _roles;

  bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

  /**
   * @dev Modifier that checks that an account has a specific role. Reverts
   * with an {AccessControlUnauthorizedAccount} error including the required role.
   */
  modifier onlyRole(bytes32 role) {
    _checkRole(role);
    _;
  }

  /**
   * @dev See {IERC165-supportsInterface}.
   */
  function supportsInterface(
    bytes4 interfaceId
  ) public view virtual override returns (bool) {
    return
      interfaceId == type(IAccessControl).interfaceId ||
      super.supportsInterface(interfaceId);
  }

  /**
   * @dev Returns `true` if `account` has been granted `role`.
   */
  function hasRole(
    bytes32 role,
    address account
  ) public view virtual returns (bool) {
    return _roles[role].hasRole[account];
  }

  /**
   * @dev Reverts with an {AccessControlUnauthorizedAccount} error if `_msgSender()`
   * is missing `role`. Overriding this function changes the behavior of the {onlyRole} modifier.
   */
  function _checkRole(bytes32 role) internal view virtual {
    _checkRole(role, _msgSender());
  }

  /**
   * @dev Reverts with an {AccessControlUnauthorizedAccount} error if `account`
   * is missing `role`.
   */
  function _checkRole(bytes32 role, address account) internal view virtual {
    if (!hasRole(role, account)) {
      revert AccessControlUnauthorizedAccount(account, role);
    }
  }

  /**
   * @dev Returns the admin role that controls `role`. See {grantRole} and
   * {revokeRole}.
   *
   * To change a role's admin, use {_setRoleAdmin}.
   */
  function getRoleAdmin(bytes32 role) public view virtual returns (bytes32) {
    return _roles[role].adminRole;
  }

  /**
   * @dev Grants `role` to `account`.
   *
   * If `account` had not been already granted `role`, emits a {RoleGranted}
   * event.
   *
   * Requirements:
   *
   * - the caller must have ``role``'s admin role.
   *
   * May emit a {RoleGranted} event.
   */
  function grantRole(
    bytes32 role,
    address account
  ) public virtual onlyRole(getRoleAdmin(role)) {
    _grantRole(role, account);
  }

  /**
   * @dev Revokes `role` from `account`.
   *
   * If `account` had been granted `role`, emits a {RoleRevoked} event.
   *
   * Requirements:
   *
   * - the caller must have ``role``'s admin role.
   *
   * May emit a {RoleRevoked} event.
   */
  function revokeRole(
    bytes32 role,
    address account
  ) public virtual onlyRole(getRoleAdmin(role)) {
    _revokeRole(role, account);
  }

  /**
   * @dev Revokes `role` from the calling account.
   *
   * Roles are often managed via {grantRole} and {revokeRole}: this function's
   * purpose is to provide a mechanism for accounts to lose their privileges
   * if they are compromised (such as when a trusted device is misplaced).
   *
   * If the calling account had been revoked `role`, emits a {RoleRevoked}
   * event.
   *
   * Requirements:
   *
   * - the caller must be `callerConfirmation`.
   *
   * May emit a {RoleRevoked} event.
   */
  function renounceRole(
    bytes32 role,
    address callerConfirmation
  ) public virtual {
    if (callerConfirmation != _msgSender()) {
      revert AccessControlBadConfirmation();
    }

    _revokeRole(role, callerConfirmation);
  }

  /**
   * @dev Sets `adminRole` as ``role``'s admin role.
   *
   * Emits a {RoleAdminChanged} event.
   */
  function _setRoleAdmin(bytes32 role, bytes32 adminRole) internal virtual {
    bytes32 previousAdminRole = getRoleAdmin(role);
    _roles[role].adminRole = adminRole;
    emit RoleAdminChanged(role, previousAdminRole, adminRole);
  }

  /**
   * @dev Attempts to grant `role` to `account` and returns a boolean indicating if `role` was granted.
   *
   * Internal function without access restriction.
   *
   * May emit a {RoleGranted} event.
   */
  function _grantRole(
    bytes32 role,
    address account
  ) internal virtual returns (bool) {
    if (!hasRole(role, account)) {
      _roles[role].hasRole[account] = true;
      emit RoleGranted(role, account, _msgSender());
      return true;
    } else {
      return false;
    }
  }

  /**
   * @dev Attempts to revoke `role` to `account` and returns a boolean indicating if `role` was revoked.
   *
   * Internal function without access restriction.
   *
   * May emit a {RoleRevoked} event.
   */
  function _revokeRole(
    bytes32 role,
    address account
  ) internal virtual returns (bool) {
    if (hasRole(role, account)) {
      _roles[role].hasRole[account] = false;
      emit RoleRevoked(role, account, _msgSender());
      return true;
    } else {
      return false;
    }
  }
}

// File: mainnet/0x4983DB49336fD4f95e864aB6DA9135e057EF0be1/contracts/libs/Errors.sol

pragma solidity 0.8.20;

error DeadlineExceeded();
error DeadlineNotExceeded();
error DeadlineTooShort();
error InsufficientOutputAmount();

error MsgValueIncorrect();
error PoolNotFound();
error TokenAddressMismatch();
error TokenNotContract();
error TokenNotETH();
error TokensIdentical();

error ChainIncorrect();
error AmountIncorrect();
error ZeroAddress();

error DisputePeriodNotPassed();
error DisputePeriodPassed();
error SenderIncorrect();
error StatusIncorrect();
error TransactionIdIncorrect();
error TransactionRelayed();

// File: @openzeppelin/contracts/utils/Address.sol

// OpenZeppelin Contracts (last updated v5.0.0) (utils/Address.sol)

pragma solidity ^0.8.20;

/**
 * @dev Collection of functions related to the address type
 */
library Address {
  /**
   * @dev The ETH balance of the account is not enough to perform the operation.
   */
  error AddressInsufficientBalance(address account);

  /**
   * @dev There's no code at `target` (it is not a contract).
   */
  error AddressEmptyCode(address target);

  /**
   * @dev A call to an address target failed. The target may have reverted.
   */
  error FailedInnerCall();

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
   * https://solidity.readthedocs.io/en/v0.8.20/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
   */
  function sendValue(address payable recipient, uint256 amount) internal {
    if (address(this).balance < amount) {
      revert AddressInsufficientBalance(address(this));
    }

    (bool success, ) = recipient.call{value: amount}('');
    if (!success) {
      revert FailedInnerCall();
    }
  }

  /**
   * @dev Performs a Solidity function call using a low level `call`. A
   * plain `call` is an unsafe replacement for a function call: use this
   * function instead.
   *
   * If `target` reverts with a revert reason or custom error, it is bubbled
   * up by this function (like regular Solidity function calls). However, if
   * the call reverted with no returned reason, this function reverts with a
   * {FailedInnerCall} error.
   *
   * Returns the raw returned data. To convert to the expected return value,
   * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
   *
   * Requirements:
   *
   * - `target` must be a contract.
   * - calling `target` with `data` must not revert.
   */
  function functionCall(
    address target,
    bytes memory data
  ) internal returns (bytes memory) {
    return functionCallWithValue(target, data, 0);
  }

  /**
   * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
   * but also transferring `value` wei to `target`.
   *
   * Requirements:
   *
   * - the calling contract must have an ETH balance of at least `value`.
   * - the called Solidity function must be `payable`.
   */
  function functionCallWithValue(
    address target,
    bytes memory data,
    uint256 value
  ) internal returns (bytes memory) {
    if (address(this).balance < value) {
      revert AddressInsufficientBalance(address(this));
    }
    (bool success, bytes memory returndata) = target.call{value: value}(data);
    return verifyCallResultFromTarget(target, success, returndata);
  }

  /**
   * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
   * but performing a static call.
   */
  function functionStaticCall(
    address target,
    bytes memory data
  ) internal view returns (bytes memory) {
    (bool success, bytes memory returndata) = target.staticcall(data);
    return verifyCallResultFromTarget(target, success, returndata);
  }

  /**
   * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
   * but performing a delegate call.
   */
  function functionDelegateCall(
    address target,
    bytes memory data
  ) internal returns (bytes memory) {
    (bool success, bytes memory returndata) = target.delegatecall(data);
    return verifyCallResultFromTarget(target, success, returndata);
  }

  /**
   * @dev Tool to verify that a low level call to smart-contract was successful, and reverts if the target
   * was not a contract or bubbling up the revert reason (falling back to {FailedInnerCall}) in case of an
   * unsuccessful call.
   */
  function verifyCallResultFromTarget(
    address target,
    bool success,
    bytes memory returndata
  ) internal view returns (bytes memory) {
    if (!success) {
      _revert(returndata);
    } else {
      // only check if target is a contract if the call was successful and the return data is empty
      // otherwise we already know that it was a contract
      if (returndata.length == 0 && target.code.length == 0) {
        revert AddressEmptyCode(target);
      }
      return returndata;
    }
  }

  /**
   * @dev Tool to verify that a low level call was successful, and reverts if it wasn't, either by bubbling the
   * revert reason or with a default {FailedInnerCall} error.
   */
  function verifyCallResult(
    bool success,
    bytes memory returndata
  ) internal pure returns (bytes memory) {
    if (!success) {
      _revert(returndata);
    } else {
      return returndata;
    }
  }

  /**
   * @dev Reverts with returndata if present. Otherwise reverts with {FailedInnerCall}.
   */
  function _revert(bytes memory returndata) private pure {
    // Look for revert reason and bubble it up if present
    if (returndata.length > 0) {
      // The easiest way to bubble the revert reason is using memory via assembly
      /// @solidity memory-safe-assembly
      assembly {
        let returndata_size := mload(returndata)
        revert(add(32, returndata), returndata_size)
      }
    } else {
      revert FailedInnerCall();
    }
  }
}

// File: @openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol

// OpenZeppelin Contracts (last updated v5.0.0) (token/ERC20/extensions/IERC20Permit.sol)

pragma solidity ^0.8.20;

/**
 * @dev Interface of the ERC20 Permit extension allowing approvals to be made via signatures, as defined in
 * https://eips.ethereum.org/EIPS/eip-2612[EIP-2612].
 *
 * Adds the {permit} method, which can be used to change an account's ERC20 allowance (see {IERC20-allowance}) by
 * presenting a message signed by the account. By not relying on {IERC20-approve}, the token holder account doesn't
 * need to send a transaction, and thus is not required to hold Ether at all.
 *
 * ==== Security Considerations
 *
 * There are two important considerations concerning the use of `permit`. The first is that a valid permit signature
 * expresses an allowance, and it should not be assumed to convey additional meaning. In particular, it should not be
 * considered as an intention to spend the allowance in any specific way. The second is that because permits have
 * built-in replay protection and can be submitted by anyone, they can be frontrun. A protocol that uses permits should
 * take this into consideration and allow a `permit` call to fail. Combining these two aspects, a pattern that may be
 * generally recommended is:
 *
 * ```solidity
 * function doThingWithPermit(..., uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) public {
 *     try token.permit(msg.sender, address(this), value, deadline, v, r, s) {} catch {}
 *     doThing(..., value);
 * }
 *
 * function doThing(..., uint256 value) public {
 *     token.safeTransferFrom(msg.sender, address(this), value);
 *     ...
 * }
 * ```
 *
 * Observe that: 1) `msg.sender` is used as the owner, leaving no ambiguity as to the signer intent, and 2) the use of
 * `try/catch` allows the permit to fail and makes the code tolerant to frontrunning. (See also
 * {SafeERC20-safeTransferFrom}).
 *
 * Additionally, note that smart contract wallets (such as Argent or Safe) are not able to produce permit signatures, so
 * contracts should have entry points that don't rely on permit.
 */
interface IERC20Permit {
  /**
   * @dev Sets `value` as the allowance of `spender` over ``owner``'s tokens,
   * given ``owner``'s signed approval.
   *
   * IMPORTANT: The same issues {IERC20-approve} has related to transaction
   * ordering also apply here.
   *
   * Emits an {Approval} event.
   *
   * Requirements:
   *
   * - `spender` cannot be the zero address.
   * - `deadline` must be a timestamp in the future.
   * - `v`, `r` and `s` must be a valid `secp256k1` signature from `owner`
   * over the EIP712-formatted function arguments.
   * - the signature must use ``owner``'s current nonce (see {nonces}).
   *
   * For more information on the signature format, see the
   * https://eips.ethereum.org/EIPS/eip-2612#specification[relevant EIP
   * section].
   *
   * CAUTION: See Security Considerations above.
   */
  function permit(
    address owner,
    address spender,
    uint256 value,
    uint256 deadline,
    uint8 v,
    bytes32 r,
    bytes32 s
  ) external;

  /**
   * @dev Returns the current nonce for `owner`. This value must be
   * included whenever a signature is generated for {permit}.
   *
   * Every successful call to {permit} increases ``owner``'s nonce by one. This
   * prevents a signature from being used multiple times.
   */
  function nonces(address owner) external view returns (uint256);

  /**
   * @dev Returns the domain separator used in the encoding of the signature for {permit}, as defined by {EIP712}.
   */
  // solhint-disable-next-line func-name-mixedcase
  function DOMAIN_SEPARATOR() external view returns (bytes32);
}

// File: @openzeppelin/contracts/token/ERC20/IERC20.sol

// OpenZeppelin Contracts (last updated v5.0.0) (token/ERC20/IERC20.sol)

pragma solidity ^0.8.20;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
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

  /**
   * @dev Returns the value of tokens in existence.
   */
  function totalSupply() external view returns (uint256);

  /**
   * @dev Returns the value of tokens owned by `account`.
   */
  function balanceOf(address account) external view returns (uint256);

  /**
   * @dev Moves a `value` amount of tokens from the caller's account to `to`.
   *
   * Returns a boolean value indicating whether the operation succeeded.
   *
   * Emits a {Transfer} event.
   */
  function transfer(address to, uint256 value) external returns (bool);

  /**
   * @dev Returns the remaining number of tokens that `spender` will be
   * allowed to spend on behalf of `owner` through {transferFrom}. This is
   * zero by default.
   *
   * This value changes when {approve} or {transferFrom} are called.
   */
  function allowance(
    address owner,
    address spender
  ) external view returns (uint256);

  /**
   * @dev Sets a `value` amount of tokens as the allowance of `spender` over the
   * caller's tokens.
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
  function approve(address spender, uint256 value) external returns (bool);

  /**
   * @dev Moves a `value` amount of tokens from `from` to `to` using the
   * allowance mechanism. `value` is then deducted from the caller's
   * allowance.
   *
   * Returns a boolean value indicating whether the operation succeeded.
   *
   * Emits a {Transfer} event.
   */
  function transferFrom(
    address from,
    address to,
    uint256 value
  ) external returns (bool);
}

// File: @openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol

// OpenZeppelin Contracts (last updated v5.0.0) (token/ERC20/utils/SafeERC20.sol)

pragma solidity ^0.8.20;

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

  /**
   * @dev An operation with an ERC20 token failed.
   */
  error SafeERC20FailedOperation(address token);

  /**
   * @dev Indicates a failed `decreaseAllowance` request.
   */
  error SafeERC20FailedDecreaseAllowance(
    address spender,
    uint256 currentAllowance,
    uint256 requestedDecrease
  );

  /**
   * @dev Transfer `value` amount of `token` from the calling contract to `to`. If `token` returns no value,
   * non-reverting calls are assumed to be successful.
   */
  function safeTransfer(IERC20 token, address to, uint256 value) internal {
    _callOptionalReturn(token, abi.encodeCall(token.transfer, (to, value)));
  }

  /**
   * @dev Transfer `value` amount of `token` from `from` to `to`, spending the approval given by `from` to the
   * calling contract. If `token` returns no value, non-reverting calls are assumed to be successful.
   */
  function safeTransferFrom(
    IERC20 token,
    address from,
    address to,
    uint256 value
  ) internal {
    _callOptionalReturn(
      token,
      abi.encodeCall(token.transferFrom, (from, to, value))
    );
  }

  /**
   * @dev Increase the calling contract's allowance toward `spender` by `value`. If `token` returns no value,
   * non-reverting calls are assumed to be successful.
   */
  function safeIncreaseAllowance(
    IERC20 token,
    address spender,
    uint256 value
  ) internal {
    uint256 oldAllowance = token.allowance(address(this), spender);
    forceApprove(token, spender, oldAllowance + value);
  }

  /**
   * @dev Decrease the calling contract's allowance toward `spender` by `requestedDecrease`. If `token` returns no
   * value, non-reverting calls are assumed to be successful.
   */
  function safeDecreaseAllowance(
    IERC20 token,
    address spender,
    uint256 requestedDecrease
  ) internal {
    unchecked {
      uint256 currentAllowance = token.allowance(address(this), spender);
      if (currentAllowance < requestedDecrease) {
        revert SafeERC20FailedDecreaseAllowance(
          spender,
          currentAllowance,
          requestedDecrease
        );
      }
      forceApprove(token, spender, currentAllowance - requestedDecrease);
    }
  }

  /**
   * @dev Set the calling contract's allowance toward `spender` to `value`. If `token` returns no value,
   * non-reverting calls are assumed to be successful. Meant to be used with tokens that require the approval
   * to be set to zero before setting it to a non-zero value, such as USDT.
   */
  function forceApprove(IERC20 token, address spender, uint256 value) internal {
    bytes memory approvalCall = abi.encodeCall(token.approve, (spender, value));

    if (!_callOptionalReturnBool(token, approvalCall)) {
      _callOptionalReturn(token, abi.encodeCall(token.approve, (spender, 0)));
      _callOptionalReturn(token, approvalCall);
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
    // we're implementing it ourselves. We use {Address-functionCall} to perform this call, which verifies that
    // the target address contains contract code and also asserts for success in the low-level call.

    bytes memory returndata = address(token).functionCall(data);
    if (returndata.length != 0 && !abi.decode(returndata, (bool))) {
      revert SafeERC20FailedOperation(address(token));
    }
  }

  /**
   * @dev Imitates a Solidity high-level call (i.e. a regular function call to a contract), relaxing the requirement
   * on the return value: the return value is optional (but if data is returned, it must not be false).
   * @param token The token targeted by the call.
   * @param data The call data (encoded using abi.encode or one of its variants).
   *
   * This is a variant of {_callOptionalReturn} that silents catches all reverts and returns a bool instead.
   */
  function _callOptionalReturnBool(
    IERC20 token,
    bytes memory data
  ) private returns (bool) {
    // We need to perform a low level call here, to bypass Solidity's return data size checking mechanism, since
    // we're implementing it ourselves. We cannot use {Address-functionCall} here since this should return false
    // and not revert is the subcall reverts.

    (bool success, bytes memory returndata) = address(token).call(data);
    return
      success &&
      (returndata.length == 0 || abi.decode(returndata, (bool))) &&
      address(token).code.length > 0;
  }
}

// File: mainnet/0x4983DB49336fD4f95e864aB6DA9135e057EF0be1/contracts/libs/UniversalToken.sol

pragma solidity 0.8.20;

library UniversalTokenLib {
  using SafeERC20 for IERC20;

  address internal constant ETH_ADDRESS =
    0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;

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
    // Don't do anything, if trying to send zero value
    if (value == 0) return;
    if (token == ETH_ADDRESS) {
      /// @dev Note: this can potentially lead to executing code in `to`.
      // solhint-disable-next-line avoid-low-level-calls
      (bool success, ) = to.call{value: value}('');
      require(success, 'ETH transfer failed');
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
    if (allowance > 0) IERC20(token).safeDecreaseAllowance(spender, allowance);
    IERC20(token).safeIncreaseAllowance(spender, type(uint256).max);
  }

  /// @notice Returns the balance of the given token (or native ETH) for the given account.
  function universalBalanceOf(
    address token,
    address account
  ) internal view returns (uint256) {
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

// File: mainnet/0x4983DB49336fD4f95e864aB6DA9135e057EF0be1/contracts/Admin.sol

pragma solidity 0.8.20;

contract Admin is IAdmin, AccessControl {
  using UniversalTokenLib for address;

  bytes32 public constant RELAYER_ROLE = keccak256('RELAYER_ROLE');
  bytes32 public constant GUARD_ROLE = keccak256('GUARD_ROLE');
  bytes32 public constant GOVERNOR_ROLE = keccak256('GOVERNOR_ROLE');

  uint256 public constant FEE_BPS = 1e6;
  uint256 public constant FEE_RATE_MAX = 0.01e6; // max 1% on origin amount

  /// @notice Protocol fee rate taken on origin amount deposited in origin chain
  uint256 public protocolFeeRate;

  /// @notice Protocol fee amounts accumulated
  mapping(address => uint256) public protocolFees;

  /// @notice Chain gas amount to forward as rebate if requested
  uint256 public chainGasAmount;

  modifier onlyGuard() {
    require(hasRole(GUARD_ROLE, msg.sender), 'Caller is not a guard');
    _;
  }

  modifier onlyRelayer() {
    require(hasRole(RELAYER_ROLE, msg.sender), 'Caller is not a relayer');
    _;
  }

  modifier onlyGovernor() {
    require(hasRole(GOVERNOR_ROLE, msg.sender), 'Caller is not a governor');
    _;
  }

  constructor(address _owner) {
    _grantRole(DEFAULT_ADMIN_ROLE, _owner);
  }

  function addRelayer(address _relayer) external {
    require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
    _grantRole(RELAYER_ROLE, _relayer);
    emit RelayerAdded(_relayer);
  }

  function removeRelayer(address _relayer) external {
    require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
    _revokeRole(RELAYER_ROLE, _relayer);
    emit RelayerRemoved(_relayer);
  }

  function addGuard(address _guard) external {
    require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
    _grantRole(GUARD_ROLE, _guard);
    emit GuardAdded(_guard);
  }

  function removeGuard(address _guard) external {
    require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
    _revokeRole(GUARD_ROLE, _guard);
    emit GuardRemoved(_guard);
  }

  function addGovernor(address _governor) external {
    require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
    _grantRole(GOVERNOR_ROLE, _governor);
    emit GovernorAdded(_governor);
  }

  function removeGovernor(address _governor) external {
    require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
    _revokeRole(GOVERNOR_ROLE, _governor);
    emit GovernorRemoved(_governor);
  }

  function setProtocolFeeRate(uint256 newFeeRate) external onlyGovernor {
    require(newFeeRate <= FEE_RATE_MAX, 'newFeeRate > max');
    uint256 oldFeeRate = protocolFeeRate;
    protocolFeeRate = newFeeRate;
    emit FeeRateUpdated(oldFeeRate, newFeeRate);
  }

  function sweepProtocolFees(
    address token,
    address recipient
  ) external onlyGovernor {
    uint256 feeAmount = protocolFees[token];
    if (feeAmount == 0) return; // skip if no accumulated fees

    protocolFees[token] = 0;
    token.universalTransfer(recipient, feeAmount);
    emit FeesSwept(token, recipient, feeAmount);
  }

  function setChainGasAmount(uint256 newChainGasAmount) external onlyGovernor {
    uint256 oldChainGasAmount = chainGasAmount;
    chainGasAmount = newChainGasAmount;
    emit ChainGasAmountUpdated(oldChainGasAmount, newChainGasAmount);
  }
}

// File: mainnet/0x4983DB49336fD4f95e864aB6DA9135e057EF0be1/contracts/FastBridge.sol

pragma solidity 0.8.20;

contract FastBridge is IFastBridge, Admin {
  using SafeERC20 for IERC20;
  using UniversalTokenLib for address;

  /// @notice Dispute period for relayed transactions
  uint256 public constant DISPUTE_PERIOD = 30 minutes;

  /// @notice Prove period added to deadline period for proven transactions
  uint256 public constant PROVE_PERIOD = 60 minutes;

  /// @notice Minimum deadline period to relay a requested bridge transaction
  uint256 public constant MIN_DEADLINE_PERIOD = 30 minutes;

  enum BridgeStatus {
    NULL, // doesn't exist yet
    REQUESTED,
    RELAYER_PROVED,
    RELAYER_CLAIMED,
    REFUNDED
  }

  /// @notice Status of the bridge tx on origin chain
  mapping(bytes32 => BridgeStatus) public bridgeStatuses;
  /// @notice Proof of relayed bridge tx on origin chain
  mapping(bytes32 => BridgeProof) public bridgeProofs;
  /// @notice Whether bridge has been relayed on destination chain
  mapping(bytes32 => bool) public bridgeRelays;

  /// @dev to prevent replays
  uint256 public nonce;
  // @dev the block the contract was deployed at
  uint256 public immutable deployBlock;

  constructor(address _owner) Admin(_owner) {
    deployBlock = block.number;
  }

  /// @notice Pulls a requested token from the user to the requested recipient.
  /// @dev Be careful of re-entrancy issues when msg.value > 0 and recipient != address(this)
  function _pullToken(
    address recipient,
    address token,
    uint256 amount
  ) internal returns (uint256 amountPulled) {
    if (token != UniversalTokenLib.ETH_ADDRESS) {
      token.assertIsContract();
      // Record token balance before transfer
      amountPulled = IERC20(token).balanceOf(recipient);
      // Token needs to be pulled only if msg.value is zero
      // This way user can specify WETH as the origin asset
      IERC20(token).safeTransferFrom(msg.sender, recipient, amount);
      // Use the difference between the recorded balance and the current balance as the amountPulled
      amountPulled = IERC20(token).balanceOf(recipient) - amountPulled;
    } else {
      // Otherwise, we need to check that ETH amount matches msg.value
      if (amount != msg.value) revert MsgValueIncorrect();
      // Transfer value to recipient if not this address
      if (recipient != address(this))
        token.universalTransfer(recipient, amount);
      // We will forward msg.value in the external call later, if recipient is not this contract
      amountPulled = msg.value;
    }
  }

  /// @inheritdoc IFastBridge
  function getBridgeTransaction(
    bytes memory request
  ) public pure returns (BridgeTransaction memory) {
    return abi.decode(request, (BridgeTransaction));
  }

  /// @inheritdoc IFastBridge
  function bridge(BridgeParams memory params) external payable {
    // check bridge params
    if (params.dstChainId == block.chainid) revert ChainIncorrect();
    if (params.originAmount == 0 || params.destAmount == 0)
      revert AmountIncorrect();
    if (params.originToken == address(0) || params.destToken == address(0))
      revert ZeroAddress();
    if (params.deadline < block.timestamp + MIN_DEADLINE_PERIOD)
      revert DeadlineTooShort();

    // transfer tokens to bridge contract
    // @dev use returned originAmount in request in case of transfer fees
    uint256 originAmount = _pullToken(
      address(this),
      params.originToken,
      params.originAmount
    );

    // track amount of origin token owed to protocol
    uint256 originFeeAmount;
    if (protocolFeeRate > 0)
      originFeeAmount = (originAmount * protocolFeeRate) / FEE_BPS;
    originAmount -= originFeeAmount; // remove from amount used in request as not relevant for relayers

    // set status to requested
    bytes memory request = abi.encode(
      BridgeTransaction({
        originChainId: uint32(block.chainid),
        destChainId: params.dstChainId,
        originSender: params.sender,
        destRecipient: params.to,
        originToken: params.originToken,
        destToken: params.destToken,
        originAmount: originAmount,
        destAmount: params.destAmount,
        originFeeAmount: originFeeAmount,
        sendChainGas: params.sendChainGas,
        deadline: params.deadline,
        nonce: nonce++ // increment nonce on every bridge
      })
    );
    bytes32 transactionId = keccak256(request);
    bridgeStatuses[transactionId] = BridgeStatus.REQUESTED;

    emit BridgeRequested(
      transactionId,
      params.sender,
      request,
      params.dstChainId,
      params.originToken,
      params.destToken,
      originAmount,
      params.destAmount,
      params.sendChainGas
    );
  }

  /// @inheritdoc IFastBridge
  function relay(bytes memory request) external payable onlyRelayer {
    bytes32 transactionId = keccak256(request);
    BridgeTransaction memory transaction = getBridgeTransaction(request);
    if (transaction.destChainId != uint32(block.chainid))
      revert ChainIncorrect();

    // check haven't exceeded deadline for relay to happen
    if (block.timestamp > transaction.deadline) revert DeadlineExceeded();

    // mark bridge transaction as relayed
    if (bridgeRelays[transactionId]) revert TransactionRelayed();
    bridgeRelays[transactionId] = true;

    // transfer tokens to recipient on destination chain and gas rebate if requested
    address to = transaction.destRecipient;
    address token = transaction.destToken;
    uint256 amount = transaction.destAmount;

    uint256 rebate = chainGasAmount;
    if (!transaction.sendChainGas) {
      // forward erc20
      rebate = 0;
      _pullToken(to, token, amount);
    } else if (token == UniversalTokenLib.ETH_ADDRESS) {
      // lump in gas rebate into amount in native gas token
      _pullToken(to, token, amount + rebate);
    } else {
      // forward erc20 then forward gas rebate in native gas token
      _pullToken(to, token, amount);
      _pullToken(to, UniversalTokenLib.ETH_ADDRESS, rebate);
    }

    emit BridgeRelayed(
      transactionId,
      msg.sender,
      to,
      transaction.originChainId,
      transaction.originToken,
      transaction.destToken,
      transaction.originAmount,
      transaction.destAmount,
      rebate
    );
  }

  /// @inheritdoc IFastBridge
  function prove(
    bytes memory request,
    bytes32 destTxHash
  ) external onlyRelayer {
    bytes32 transactionId = keccak256(request);
    BridgeTransaction memory transaction = getBridgeTransaction(request);

    // check haven't exceeded deadline for prove to happen
    if (block.timestamp > transaction.deadline + PROVE_PERIOD)
      revert DeadlineExceeded();

    // update bridge tx status given proof provided
    if (bridgeStatuses[transactionId] != BridgeStatus.REQUESTED)
      revert StatusIncorrect();
    bridgeStatuses[transactionId] = BridgeStatus.RELAYER_PROVED;
    bridgeProofs[transactionId] = BridgeProof({
      timestamp: uint96(block.timestamp),
      relayer: msg.sender
    }); // overflow ok

    emit BridgeProofProvided(transactionId, msg.sender, destTxHash);
  }

  /// @notice Calculates time since proof submitted
  /// @dev proof.timestamp stores casted uint96(block.timestamp) block timestamps for gas optimization
  ///      _timeSince(proof) can accomodate rollover case when block.timestamp > type(uint96).max but
  ///      proof.timestamp < type(uint96).max via unchecked statement
  /// @param proof The bridge proof
  /// @return delta Time delta since proof submitted
  function _timeSince(
    BridgeProof memory proof
  ) internal view returns (uint256 delta) {
    unchecked {
      delta = uint96(block.timestamp) - proof.timestamp;
    }
  }

  /// @inheritdoc IFastBridge
  function canClaim(
    bytes32 transactionId,
    address relayer
  ) external view returns (bool) {
    if (bridgeStatuses[transactionId] != BridgeStatus.RELAYER_PROVED)
      revert StatusIncorrect();
    BridgeProof memory proof = bridgeProofs[transactionId];
    if (proof.relayer != relayer) revert SenderIncorrect();
    return _timeSince(proof) > DISPUTE_PERIOD;
  }

  /// @inheritdoc IFastBridge
  function claim(bytes memory request, address to) external onlyRelayer {
    bytes32 transactionId = keccak256(request);
    BridgeTransaction memory transaction = getBridgeTransaction(request);

    // update bridge tx status if able to claim origin collateral
    if (bridgeStatuses[transactionId] != BridgeStatus.RELAYER_PROVED)
      revert StatusIncorrect();

    BridgeProof memory proof = bridgeProofs[transactionId];
    if (proof.relayer != msg.sender) revert SenderIncorrect();
    if (_timeSince(proof) <= DISPUTE_PERIOD) revert DisputePeriodNotPassed();

    bridgeStatuses[transactionId] = BridgeStatus.RELAYER_CLAIMED;

    // update protocol fees if origin fee amount exists
    if (transaction.originFeeAmount > 0)
      protocolFees[transaction.originToken] += transaction.originFeeAmount;

    // transfer origin collateral less fee to specified address
    address token = transaction.originToken;
    uint256 amount = transaction.originAmount;
    token.universalTransfer(to, amount);

    emit BridgeDepositClaimed(transactionId, msg.sender, to, token, amount);
  }

  /// @inheritdoc IFastBridge
  function dispute(bytes32 transactionId) external onlyGuard {
    if (bridgeStatuses[transactionId] != BridgeStatus.RELAYER_PROVED)
      revert StatusIncorrect();
    if (_timeSince(bridgeProofs[transactionId]) > DISPUTE_PERIOD)
      revert DisputePeriodPassed();

    // @dev relayer gets slashed effectively if dest relay has gone thru
    bridgeStatuses[transactionId] = BridgeStatus.REQUESTED;
    delete bridgeProofs[transactionId];

    emit BridgeProofDisputed(transactionId, msg.sender);
  }

  /// @inheritdoc IFastBridge
  function refund(bytes memory request) external {
    bytes32 transactionId = keccak256(request);
    BridgeTransaction memory transaction = getBridgeTransaction(request);

    // check exceeded deadline for prove to happen
    if (block.timestamp <= transaction.deadline + PROVE_PERIOD)
      revert DeadlineNotExceeded();

    // set status to refunded if still in requested state
    if (bridgeStatuses[transactionId] != BridgeStatus.REQUESTED)
      revert StatusIncorrect();
    bridgeStatuses[transactionId] = BridgeStatus.REFUNDED;

    // transfer origin collateral back to original sender
    address to = transaction.originSender;
    address token = transaction.originToken;
    uint256 amount = transaction.originAmount + transaction.originFeeAmount;
    token.universalTransfer(to, amount);

    emit BridgeDepositRefunded(transactionId, to, token, amount);
  }
}

contract TestFastBridge is FastBridge {
  constructor(address _owner) FastBridge(_owner) {}

  function testBridge(
    uint32 dstChainId,
    address sender,
    address to,
    address originToken,
    address destToken,
    uint256 originAmount,
    uint256 destAmount,
    bool sendChainGas
  ) external {
    emit BridgeRequested(
      keccak256('dummyTransactionId'),
      sender,
      abi.encodePacked('dummyRequest'),
      dstChainId,
      originToken,
      destToken,
      originAmount,
      destAmount,
      sendChainGas
    );
  }

  function testRelay(
    bytes32 transactionId,
    address relayer,
    address to,
    uint32 originChainId,
    address originToken,
    address destToken,
    uint256 originAmount,
    uint256 destAmount,
    uint256 rebate
  ) external {
    emit BridgeRelayed(
      transactionId,
      relayer,
      to,
      originChainId,
      originToken,
      destToken,
      originAmount,
      destAmount,
      rebate
    );
  }
}
