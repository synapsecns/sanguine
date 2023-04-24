// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import {Versioned} from "../Version.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

abstract contract SystemBase is Versioned, OwnableUpgradeable {
    // ══════════════════════════════════════════ CONSTANTS & IMMUTABLES ═══════════════════════════════════════════════

    // domain of the Synapse Chain
    // For MVP this is Optimism chainId
    // TODO: replace the placeholder with actual value
    uint32 public constant SYNAPSE_DOMAIN = 10;

    // TODO: reevaluate optimistic period for staking/unstaking bonds
    uint32 internal constant BONDING_OPTIMISTIC_PERIOD = 1 days;

    /// @notice Domain of the local chain, set once upon contract creation
    uint32 public immutable localDomain;

    // ══════════════════════════════════════════════════ STORAGE ══════════════════════════════════════════════════════

    /// @dev gap for upgrade safety
    uint256[50] private __GAP; // solhint-disable-line var-name-mixedcase

    constructor(uint32 domain) {
        localDomain = domain;
    }

    /**
     * @dev Should be impossible to renounce ownership;
     * we override OpenZeppelin OwnableUpgradeable's
     * implementation of renounceOwnership to make it a no-op
     */
    function renounceOwnership() public override onlyOwner {} //solhint-disable-line no-empty-blocks
}
