// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Origin } from "../../contracts/Origin.sol";
import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";
import { SystemContractHarness } from "./SystemContractHarness.sol";

contract OriginHarness is Origin, SystemContractHarness, GuardRegistryHarness {
    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) Origin(_domain) {}

    function isNotary(address _notary) public view returns (bool) {
        return _isNotary(_localDomain(), _notary);
    }

    function setFailed(address _notary) public {
        _fail(_notary, address(0));
    }

    function destinationAndNonce(uint32 _destination, uint32 _nonce) public pure returns (uint64) {
        return _destinationAndNonce(_destination, _nonce);
    }
}
