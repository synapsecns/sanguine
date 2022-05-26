// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Replica } from "../../contracts/Replica.sol";

contract ReplicaHarness is Replica {
    constructor(
        uint32 _localDomain,
        uint256 _processGas,
        uint256 _reserveGas
    ) Replica(_localDomain, _processGas, _reserveGas) {}
}
