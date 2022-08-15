// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SystemMessenger } from "../../contracts/system/SystemMessenger.sol";

contract SystemMessengerHarness is SystemMessenger {
    constructor(
        address _origin,
        address _destination,
        uint32 _optimisticSeconds
    ) SystemMessenger(_origin, _destination, _optimisticSeconds) {}
}
