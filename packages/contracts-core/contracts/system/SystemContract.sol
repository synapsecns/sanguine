// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { ISystemRouter } from "../interfaces/ISystemRouter.sol";
import { DomainContext } from "../context/DomainContext.sol";

import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @notice Shared utilities between Synapse System Contracts: Origin, Destination, etc.
 */
abstract contract SystemContract is DomainContext, OwnableUpgradeable {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              CONSTANTS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // domain of the Synapse Chain
    // Answer to the Ultimate Question of Life, the Universe, and Everything
    // And answer to less important questions wink wink
    uint32 public constant SYNAPSE_DOMAIN = 4269;
    // TODO: replace the placeholder with actual value

    uint256 internal constant ORIGIN = 1 << uint8(ISystemRouter.SystemEntity.Origin);
    uint256 internal constant DESTINATION = 1 << uint8(ISystemRouter.SystemEntity.Destination);
    uint256 internal constant BONDING_MANAGER =
        1 << uint8(ISystemRouter.SystemEntity.BondingManager);

    // TODO: reevaluate optimistic period for staking/unstaking bonds
    uint32 internal constant BONDING_OPTIMISTIC_PERIOD = 1 days;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    ISystemRouter public systemRouter;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              MODIFIERS                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * System Contracts on all chains (either local or remote).
     * Note: any function protected by this modifier should have last three params:
     * - uint32 _callOrigin
     * - SystemEntity _caller
     * - uint256 _rootSubmittedAt
     * Make sure to check domain/caller, if a function should be only called
     * from a given domain / by a given caller.
     * Make sure to check that a needed amount of time has passed since
     * root submission for the cross-chain calls.
     */
    modifier onlySystemRouter() {
        _assertSystemRouter();
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * System Contracts on Synapse chain.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     */
    modifier onlySynapseChain(uint32 _callOrigin) {
        _assertSynapseChain(_callOrigin);
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * a set of System Contracts on any chain.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     * Note: check constants section for existing mask constants
     * E.g. to restrict the set of callers to three allowed system callers:
     *  onlyCallers(MASK_0 | MASK_1 | MASK_2, _caller)
     */
    modifier onlyCallers(uint256 _allowedMask, ISystemRouter.SystemEntity _caller) {
        _assertEntityAllowed(_allowedMask, _caller);
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * BondingManager on their local chain.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     */
    modifier onlyLocalBondingManager(uint32 _callOrigin, ISystemRouter.SystemEntity _caller) {
        _assertLocalDomain(_callOrigin);
        _assertEntityAllowed(BONDING_MANAGER, _caller);
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * BondingManager on Synapse Chain.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     */
    modifier onlySynapseChainBondingManager(
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller
    ) {
        _assertSynapseChain(_callOrigin);
        _assertEntityAllowed(BONDING_MANAGER, _caller);
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * System Contracts on remote chain with a defined minimum optimistic period.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     * Note: message could be sent with a period lower than that, but will be executed
     * only when `_optimisticSeconds` have passed.
     * Note: _optimisticSeconds=0 will allow calls from a local chain as well
     */
    modifier onlyOptimisticPeriodOver(uint256 _rootSubmittedAt, uint256 _optimisticSeconds) {
        _assertOptimisticPeriodOver(_rootSubmittedAt, _optimisticSeconds);
        _;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line func-name-mixedcase
    function __SystemContract_initialize() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              OWNER ONLY                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line ordering
    function setSystemRouter(ISystemRouter _systemRouter) external onlyOwner {
        systemRouter = _systemRouter;
    }

    /**
     * @dev Should be impossible to renounce ownership;
     * we override OpenZeppelin OwnableUpgradeable's
     * implementation of renounceOwnership to make it a no-op
     */
    function renounceOwnership() public override onlyOwner {} //solhint-disable-line no-empty-blocks

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          SYSTEM ROUTER ONLY                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Receive a system call indicating that a new Notary staked a bond.
     * @param _domain           Domain where the new Notary will be active
     * @param _notary           New Notary that staked a bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function bondNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external virtual;

    /**
     * @notice Receive a system call indicating that an active Notary unstaked their bond.
     * @param _domain           Domain where the Notary was active
     * @param _notary           Active Notary that unstaked their bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function unbondNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external virtual;

    /**
     * @notice Receive a system call indicating that an active Notary was slashed.
     * @param _domain           Domain where the slashed Notary was active
     * @param _notary           Active Notary that was slashed
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function slashNotary(
        uint32 _domain,
        address _notary,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external virtual;

    /**
     * @notice Receive a system call indicating that a new Guard staked a bond.
     * @param _guard            New Guard that staked a bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function bondGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external virtual;

    /**
     * @notice Receive a system call indicating that an active Guard unstaked their bond.
     * @param _guard            Active Guard that unstaked their bond
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function unbondGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external virtual;

    /**
     * @notice Receive a system call indicating that an active Guard was slashed.
     * @param _guard            Active Guard that was slashed
     * @param _callOrigin       Domain where the system call originated
     * @param _caller           Entity which performed the system call
     * @param _rootSubmittedAt  Time when merkle root (used for proving this message) was submitted
     */
    function slashGuard(
        address _guard,
        uint32 _callOrigin,
        ISystemRouter.SystemEntity _caller,
        uint256 _rootSubmittedAt
    ) external virtual;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _assertSystemRouter() internal view {
        require(msg.sender == address(systemRouter), "!systemRouter");
    }

    function _assertOptimisticPeriodOver(uint256 _rootSubmittedAt, uint256 _optimisticSeconds)
        internal
        view
    {
        require(block.timestamp >= _rootSubmittedAt + _optimisticSeconds, "!optimisticPeriod");
    }

    function _assertEntityAllowed(uint256 _allowedMask, ISystemRouter.SystemEntity _caller)
        internal
        pure
    {
        require(_entityAllowed(_allowedMask, _caller), "!allowedCaller");
    }

    function _assertSynapseChain(uint32 _domain) internal pure {
        require(_domain == SYNAPSE_DOMAIN, "!synapseDomain");
    }

    /**
     * @notice Checks if a given entity is allowed to call a function using a _systemMask
     * @param _systemMask a mask of allowed entities
     * @param _entity a system entity to check
     * @return true if _entity is allowed to call a function
     *
     * @dev this function works by converting the enum value to a non-zero bit mask
     * we then use a bitwise AND operation to check if permission bits allow the entity
     * to perform this operation, more details can be found here:
     * https://en.wikipedia.org/wiki/Bitwise_operation#AND
     */
    function _entityAllowed(uint256 _systemMask, ISystemRouter.SystemEntity _entity)
        internal
        pure
        returns (bool)
    {
        return _systemMask & _getSystemMask(_entity) != 0;
    }

    /**
     * @notice Returns a mask for a given system entity
     * @param _entity System entity
     * @return a non-zero mask for a given system entity
     *
     * Converts an enum value into a non-zero bit mask used for a bitwise AND check
     * E.g. for Origin (0) returns 1, for Destination (1) returns 2
     */
    function _getSystemMask(ISystemRouter.SystemEntity _entity) internal pure returns (uint256) {
        return 1 << uint8(_entity);
    }
}
