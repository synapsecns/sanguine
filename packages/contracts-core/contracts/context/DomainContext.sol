// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

abstract contract DomainContext {
    /**
     * @notice Ensures that a domain matches the local domain.
     */
    modifier onlyLocalDomain(uint32 _domain) {
        require(_domain == _localDomain(), "!localDomain");
        _;
    }

    function localDomain() external view returns (uint32) {
        return _localDomain();
    }

    function _localDomain() internal view virtual returns (uint32);
}
