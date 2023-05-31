// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Request, RequestLib} from "../../../../contracts/libs/stack/Request.sol";

// solhint-disable ordering
contract RequestHarness {
    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    function encodeRequest(uint96 gasDrop_, uint64 gasLimit_, uint32 version_) public pure returns (uint192) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Request request = RequestLib.encodeRequest(gasDrop_, gasLimit_, version_);
        return Request.unwrap(request);
    }

    function wrapPadded(uint256 paddedRequest) public pure returns (uint192) {
        return Request.unwrap(RequestLib.wrapPadded(paddedRequest));
    }

    function gasLimit(uint256 paddedRequest) public pure returns (uint64) {
        return RequestLib.wrapPadded(paddedRequest).gasLimit();
    }

    function gasDrop(uint256 paddedRequest) public pure returns (uint96) {
        return RequestLib.wrapPadded(paddedRequest).gasDrop();
    }

    function version(uint256 paddedRequest) public pure returns (uint32) {
        return RequestLib.wrapPadded(paddedRequest).version();
    }
}
