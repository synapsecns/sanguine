// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Auth } from "../../contracts/libs/Auth.sol";

contract AuthHarness {
    function checkSignature(
        address _signer,
        bytes memory _payload,
        bytes memory _signature
    ) public pure returns (bytes29) {
        return Auth.checkSignature(_signer, _payload, _signature);
    }
}
