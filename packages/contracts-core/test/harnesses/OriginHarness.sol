// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Origin } from "../../contracts/Origin.sol";
import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";

contract OriginHarness is Origin, GuardRegistryHarness {
    uint256 public sensitiveValue;

    event LogSystemCall(uint32 origin, uint8 caller);

    constructor(uint32 _domain) Origin(_domain) {}

    function isNotary(address _notary) public view returns (bool) {
        return _isNotary(localDomain, _notary);
    }

    function setSensitiveValue(
        uint256 _newValue,
        uint32 _origin,
        uint8 _caller
    ) external onlySystemRouter {
        sensitiveValue = _newValue;
        emit LogSystemCall(_origin, _caller);
    }

    function setFailed(address _notary) public {
        _fail(_notary, address(0));
    }

    function destinationAndNonce(uint32 _destination, uint32 _nonce) public pure returns (uint64) {
        return _destinationAndNonce(_destination, _nonce);
    }
}
