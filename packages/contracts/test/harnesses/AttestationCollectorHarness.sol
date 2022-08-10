// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { AttestationCollector } from "../../contracts/AttestationCollector.sol";

contract AttestationCollectorHarness is AttestationCollector {
    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }
}
