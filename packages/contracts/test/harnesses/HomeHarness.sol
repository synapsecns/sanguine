// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Home } from "../../contracts/Home.sol";

contract HomeHarness is Home {
    uint256 public sensitiveValue;

    constructor(uint32 _domain) Home(_domain) {}

    function setSensitiveValue(uint256 _newValue) external onlySystemMessenger {
        sensitiveValue = _newValue;
    }

    function historicalRoots(uint256 _nonce) public view returns (bytes32) {
        return _historicalRoots[_nonce];
    }

    function setFailed() public {
        _fail();
    }

    function destinationAndNonce(uint32 _destination, uint32 _nonce) public pure returns (uint64) {
        return _destinationAndNonce(_destination, _nonce);
    }
}
