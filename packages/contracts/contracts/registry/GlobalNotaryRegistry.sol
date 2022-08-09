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

    // [domain => [notaries]]
    mapping(uint32 => address[]) internal domainNotaries;

    // [domain => [notary => position in the above array plus 1]]
    // (index 0 means notary is not in the array)
    mapping(uint32 => mapping(address => uint256)) private notariesIndexes;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line var-name-mixedcase
    uint256[48] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                EVENTS                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    event NotaryAdded(uint32 indexed domain, address notary);

    event NotaryRemoved(uint32 indexed domain, address notary);

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _addNotary(uint32 _domain, address _notary) internal returns (bool) {
        if (_isNotary(_domain, _notary)) return false;
        domainNotaries[_domain].push(_notary);
        notariesIndexes[_domain][_notary] = domainNotaries[_domain].length;
        emit NotaryAdded(_domain, _notary);
        return true;
    }

    function _removeNotary(uint32 _domain, address _notary) internal returns (bool) {
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

    function _isNotary(uint32 _domain, address _notary) internal view override returns (bool) {
        return notariesIndexes[_domain][_notary] != 0;
    }
}
