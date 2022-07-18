// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { SystemMessenger } from "../../contracts/system/SystemMessenger.sol";

contract SystemMessengerHarness is SystemMessenger {
    constructor(
        address _home,
        address _replicaManager,
        uint32 _optimisticSeconds
    ) SystemMessenger(_home, _replicaManager, _optimisticSeconds) {}
}
