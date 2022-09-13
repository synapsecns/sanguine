// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

abstract contract DomainContext {
    function localDomain() external view returns (uint32) {
        return _localDomain();
    }

    function _localDomain() internal view virtual returns (uint32);
}
