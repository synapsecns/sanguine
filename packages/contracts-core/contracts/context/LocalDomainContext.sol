// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { DomainContext } from "./DomainContext.sol";

contract LocalDomainContext is DomainContext {
    uint32 private immutable __localDomain;

    constructor(uint32 localDomain_) {
        __localDomain = localDomain_;
    }

    function _localDomain() internal view override returns (uint32) {
        return __localDomain;
    }
}
