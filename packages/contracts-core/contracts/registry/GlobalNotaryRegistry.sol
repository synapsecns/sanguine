// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { AbstractNotaryRegistry } from "./AbstractNotaryRegistry.sol";

/**
 * @notice A Registry to keep track of Notaries on all domains.
 *
 * @dev Modified OZ's EnumerableSet. This enables mapping into EnumerableSet.
 * https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/structs/EnumerableSet.sol
 */
contract GlobalNotaryRegistry is AbstractNotaryRegistry {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Array of active Notaries for every domain
    // [domain => [notaries]]
    mapping(uint32 => address[]) internal domainNotaries;

    // [domain => [notary => position in the above array plus 1]]
    // (index 0 means notary is not in the array)
    mapping(uint32 => mapping(address => uint256)) private notariesIndexes;

    // gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Emitted when a new Notary is added.
     * @param domain    Domain where a Notary was added
     * @param notary    Address of the added notary
     */
    event NotaryAdded(uint32 indexed domain, address notary);

    /**
     * @notice Emitted when a new Notary is removed.
     * @param domain    Domain where a Notary was removed
     * @param notary    Address of the removed notary
     */
    event NotaryRemoved(uint32 indexed domain, address notary);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Tries to add a new notary, emits an event only if notary was added.
     */
    function _addNotary(uint32 _domain, address _notary) internal override returns (bool) {
        if (_isNotary(_domain, _notary)) return false;
        domainNotaries[_domain].push(_notary);
        notariesIndexes[_domain][_notary] = domainNotaries[_domain].length;
        emit NotaryAdded(_domain, _notary);
        return true;
    }

    /**
     * @notice Tries to remove a notary, emits an event only if notary was removed.
     */
    function _removeNotary(uint32 _domain, address _notary) internal override returns (bool) {
        uint256 valueIndex = notariesIndexes[_domain][_notary];
        if (valueIndex == 0) return false;
        // To delete a Notary from the array in O(1),
        // we swap the Notary to delete with the last one in the array,
        // and then remove the last Notary (sometimes called as 'swap and pop').
        address[] storage notaries = domainNotaries[_domain];
        uint256 toDeleteIndex = valueIndex - 1;
        uint256 lastIndex = notaries.length - 1;
        if (lastIndex != toDeleteIndex) {
            address lastNotary = notaries[lastIndex];
            // Move the last Notary to the index where the Notary to delete is
            notaries[toDeleteIndex] = lastNotary;
            // Update the index for the moved Notary
            notariesIndexes[_domain][lastNotary] = valueIndex;
        }
        // Delete the slot where the moved Notary was stored
        notaries.pop();
        // Delete the index for the deleted slot
        delete notariesIndexes[_domain][_notary];
        emit NotaryRemoved(_domain, _notary);
        return true;
    }

    /**
     * @notice Returns whether a given address is a notary for a given domain.
     */
    function _isNotary(uint32 _domain, address _account) internal view override returns (bool) {
        return notariesIndexes[_domain][_account] != 0;
    }
}
