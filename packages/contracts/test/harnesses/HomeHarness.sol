// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Home } from "../../contracts/Home.sol";

contract HomeHarness is Home {
    constructor(uint32 _domain) Home(_domain) {}

    function setFailed() public {
        _fail();
    }

    function destinationAndNonce(uint32 _destination, uint32 _nonce) public pure returns (uint64) {
        return _destinationAndNonce(_destination, _nonce);
    }

    function isUpdaterSignature(
        bytes32 _oldRoot,
        bytes32 _newRoot,
        bytes memory _signature
    ) external view returns (bool) {
        return _isUpdaterSignature(_oldRoot, _newRoot, _signature);
    }


}
