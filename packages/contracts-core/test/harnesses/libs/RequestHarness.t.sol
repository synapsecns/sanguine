// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {Request, RequestLib, TypedMemView} from "../../../contracts/libs/Request.sol";

// solhint-disable ordering
contract RequestHarness {
    using RequestLib for bytes;
    using RequestLib for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    // Note: we don't add an empty test() function here, as it currently leads
    // to zero coverage on the corresponding library.

    // ══════════════════════════════════════════════════ GETTERS ══════════════════════════════════════════════════════

    function castToRequest(bytes memory payload) public view returns (bytes memory) {
        // Walkaround to get the forge coverage working on libraries, see
        // https://github.com/foundry-rs/foundry/pull/3128#issuecomment-1241245086
        Request request = RequestLib.castToRequest(payload);
        return request.unwrap().clone();
    }

    function isRequest(bytes memory payload) public pure returns (bool) {
        return payload.ref(0).isRequest();
    }

    function gasLimit(bytes memory payload) public pure returns (uint64) {
        return payload.castToRequest().gasLimit();
    }

    // ════════════════════════════════════════════════ FORMATTERS ═════════════════════════════════════════════════════

    function formatRequest(uint64 gasLimit_) public pure returns (bytes memory) {
        return RequestLib.formatRequest(gasLimit_);
    }
}
