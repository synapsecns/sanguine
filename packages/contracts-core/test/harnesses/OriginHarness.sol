// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Origin } from "../../contracts/Origin.sol";
import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";
import { SystemContractHarness } from "./SystemContractHarness.sol";

contract OriginHarness is Origin, SystemContractHarness, GuardRegistryHarness {
    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) Origin(_domain) {}

    function removeAllNotaries() public {
        uint256 amount = notariesAmount();
        // Remove every Notary to halt the contract
        for (uint256 i = 0; i < amount; ++i) {
            _removeNotary(getNotary(0));
        }
    }

    function isNotary(address _notary) public view returns (bool) {
        return _isNotary(_localDomain(), _notary);
    }
}
