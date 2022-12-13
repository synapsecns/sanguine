// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { AttestationCollector } from "../../contracts/AttestationCollector.sol";

contract AttestationCollectorHarness is AttestationCollector {
    function isGuard(address _guard) external view returns (bool) {
        return _isActiveAgent(0, _guard);
    }

    function isNotary(uint32 _domain, address _notary) external view returns (bool) {
        return _isActiveAgent(_domain, _notary);
    }
}
