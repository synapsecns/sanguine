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
     * - SystemEntity _systemCaller
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
    modifier onlySynapseChain(uint32 _originDomain) {
        require(_originDomain == SYNAPSE_DOMAIN, "!synapseDomain");
        _;
    }

    /**
     * @dev Modifier for functions that are supposed to be called only from
     * a set of System Contracts on any chain.
     * Note: has to be used alongside with `onlySystemRouter`
     * See `onlySystemRouter` for details about the functions protected by such modifiers.
     * Note: check constants section for existing mask constants
     * E.g. to restrict the set of callers to three allowed system callers:
     *  onlyCallers(MASK_0 | MASK_1 | MASK_2, _systemCaller)
     */
    modifier onlyCallers(uint256 _allowedMask, ISystemRouter.SystemEntity _systemCaller) {
        require(_entityAllowed(_allowedMask, _systemCaller), "!allowedCaller");
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
