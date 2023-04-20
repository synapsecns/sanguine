// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Request, RequestLib} from "../../../contracts/libs/Request.sol";

// solhint-disable ordering
contract RequestHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    function encodeRequest(uint96 gasDrop_, uint64 gasLimit_) public pure returns (uint256) {
        return Request.unwrap(RequestLib.encodeRequest(gasDrop_, gasLimit_));
    }

    function wrapPadded(uint256 paddedRequest) public pure returns (uint160) {
        return Request.unwrap(RequestLib.wrapPadded(paddedRequest));
    }

    function gasLimit(uint256 paddedRequest) public pure returns (uint64) {
        return RequestLib.wrapPadded(paddedRequest).gasLimit();
    }

    function gasDrop(uint256 paddedRequest) public pure returns (uint96) {
        return RequestLib.wrapPadded(paddedRequest).gasDrop();
    }
}
