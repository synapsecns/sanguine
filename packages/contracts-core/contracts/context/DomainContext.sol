// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract DomainContext {
    /**
     * @notice Ensures that a domain matches the local domain.
     */
    modifier onlyLocalDomain(uint32 _domain) {
        _assertLocalDomain(_domain);
        _;
    }

    function localDomain() external view returns (uint32) {
        return _localDomain();
    }

    function _assertLocalDomain(uint32 _domain) internal view {
        require(_domain == _localDomain(), "!localDomain");
    }

    function _localDomain() internal view virtual returns (uint32);
}
