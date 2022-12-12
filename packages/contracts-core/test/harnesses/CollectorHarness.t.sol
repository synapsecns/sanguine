// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Collector } from "../../contracts/Collector.sol";

contract CollectorHarness is Collector {
    function isGuard(address _guard) external view returns (bool) {
        return _isActiveAgent(0, _guard);
    }

    function isNotary(uint32 _domain, address _notary) external view returns (bool) {
        return _isActiveAgent(_domain, _notary);
    }
}
