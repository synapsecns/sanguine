// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract DomainContext {
    /// @notice Domain of the local chain, set once upon contract creation
    uint32 public immutable localDomain;

    /**
     * @notice Ensures that a domain matches the local domain.
     */
    modifier onlyLocalDomain(uint32 domain) {
        _assertLocalDomain(domain);
        _;
    }

    constructor(uint32 domain) {
        localDomain = domain;
    }

    function _assertLocalDomain(uint32 domain) internal view {
        require(domain == localDomain, "!localDomain");
    }
}
