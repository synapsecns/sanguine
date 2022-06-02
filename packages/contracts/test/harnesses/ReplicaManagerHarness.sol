// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { ReplicaManager } from "../../contracts/ReplicaManager.sol";

contract ReplicaManagerHarness is ReplicaManager {
    constructor(
        uint32 _localDomain,
        uint256 _processGas,
        uint256 _reserveGas
    ) ReplicaManager(_localDomain, _processGas, _reserveGas) {}
}
