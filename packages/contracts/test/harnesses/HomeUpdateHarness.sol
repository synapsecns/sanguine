// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { HomeUpdate } from "../../contracts/libs/HomeUpdate.sol";
import { TypedMemView } from "../../contracts/libs/TypedMemView.sol";

contract HomeUpdateHarness {
    using HomeUpdate for bytes29;
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    function formatHomeUpdate(
        uint32 _domain,
        uint32 _nonce,
        bytes32 _root
    ) public pure returns (bytes memory) {
        return HomeUpdate.formatHomeUpdate(_domain, _nonce, _root);
    }

    function isValid(bytes memory _update) public pure returns (bool) {
        return _update.ref(0).isValidUpdate();
    }

    function domain(bytes memory _update) public pure returns (uint32) {
        return _update.ref(0).updateDomain();
    }

    function nonce(bytes memory _update) public pure returns (uint32) {
        return _update.ref(0).updateNonce();
    }

    function root(bytes memory _update) public pure returns (bytes32) {
        return _update.ref(0).updateRoot();
    }
}
