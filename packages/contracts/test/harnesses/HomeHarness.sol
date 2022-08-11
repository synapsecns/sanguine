// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Home } from "../../contracts/Home.sol";

contract HomeHarness is Home {
    uint256 public sensitiveValue;

    constructor(uint32 _domain) Home(_domain) {}

    function isNotary(address _notary) public view returns (bool) {
        return _isNotary(localDomain, _notary);
    }

    function setSensitiveValue(uint256 _newValue) external onlySystemMessenger {
        sensitiveValue = _newValue;
    }

    function setFailed(address _notary) public {
        _fail(_notary);
    }

    function destinationAndNonce(uint32 _destination, uint32 _nonce) public pure returns (uint64) {
        return _destinationAndNonce(_destination, _nonce);
    }
}
