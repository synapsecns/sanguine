// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { AbstractNotaryRegistry } from "./AbstractNotaryRegistry.sol";

// solhint-disable max-line-length
/**
 * @notice A Registry to keep track of Notaries on all domains.
 *
 * @dev Modified OZ's EnumerableSet.
 * This enables mapping(uint32 => EnumerableSet), which is not supported by Solidity natively.
 * https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/structs/EnumerableSet.sol
 *
 * It is assumed that the Notary is active on a single chain.
 */
// solhint-enable max-line-length
contract GlobalNotaryRegistry is AbstractNotaryRegistry {
    /**
     * @notice Information about an active Notary, optimized for fir in one word of storage.
     * @dev Since we're storing both domain and index, we can store the actual notary position
     * in domainNotaries[domain] instead of using the OZ approach of storing (position + 1).
     * @param domain    Domain where Notary is active (domain 0 means notary is not active)
     * @param index     Notary position in domainNotaries[domain] array
     */
    struct NotaryInfo {
        uint32 domain;
        uint224 index;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               STORAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Array of active Notaries for every domain
    // [domain => [notaries]]
    mapping(uint32 => address[]) internal domainNotaries;

    // [notary => notary info]
    mapping(address => NotaryInfo) internal notariesInfo;

    // gap for upgrade safety
    uint256[48] private __GAP; // solhint-disable-line var-name-mixedcase

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Returns addresses of all Notaries for a given domain.
     * @dev This copies storage into memory, so can consume a lof of gas, if
     * amount of notaries is large (see EnumerableSet.values())
     */
    function allNotaries(uint32 _domain) external view returns (address[] memory) {
        return domainNotaries[_domain];
    }

    /**
     * @notice Returns i-th Notary for a given domain. O(1)
     * @dev Will revert if index is out of range
     */
    function getNotary(uint32 _domain, uint256 _index) public view returns (address) {
        return domainNotaries[_domain][_index];
    }

    /**
     * @notice Returns amount of active notaries for a given domain. O(1)
     */
    function notariesAmount(uint32 _domain) public view returns (uint256) {
        return domainNotaries[_domain].length;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice Tries to add a new notary, emits an event only if notary was added.
     * @dev Notary will not be added, if it's active on ANY domain.
     */
    function _addNotary(uint32 _domain, address _notary) internal override returns (bool) {
        if (notariesInfo[_notary].domain != 0) return false;
        notariesInfo[_notary] = NotaryInfo({
            domain: _domain,
            index: uint224(domainNotaries[_domain].length)
        });
        domainNotaries[_domain].push(_notary);
        emit NotaryAdded(_domain, _notary);
        return true;
    }

    /**
     * @notice Tries to remove a notary, emits an event only if notary was removed.
     */
    function _removeNotary(uint32 _domain, address _notary) internal override returns (bool) {
        NotaryInfo memory info = notariesInfo[_notary];
        if (info.domain != _domain) return false;
        // To delete a Notary from the array in O(1),
        // we swap the Notary to delete with the last one in the array,
        // and then remove the last Notary (sometimes called as 'swap and pop').
        address[] storage notaries = domainNotaries[_domain];
        uint256 lastIndex = notaries.length - 1;
        if (lastIndex != info.index) {
            address lastNotary = notaries[lastIndex];
            // Move the last Notary to the index where the Notary to delete is
            notaries[info.index] = lastNotary;
            // Update the index for the moved Notary
            notariesInfo[lastNotary].index = info.index;
        }
        // Delete the slot where the moved Notary was stored
        notaries.pop();
        // Delete the index for the deleted slot
        delete notariesInfo[_notary];
        emit NotaryRemoved(_domain, _notary);
        return true;
    }

    /**
     * @notice Returns whether a given address is a notary for a given domain.
     */
    function _isNotary(uint32 _domain, address _account) internal view override returns (bool) {
        return notariesInfo[_account].domain == _domain;
    }
}
