// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {Create2} from "../../src/libs/Create2.sol";

contract Create2Harness {
    function deploy(
        uint256 value,
        bytes32 salt,
        bytes memory creationCode
    )
        external
        payable
        returns (address deployedAt)
    {
        return Create2.deploy(value, salt, creationCode);
    }

    function predictDeployment(
        address deployer,
        bytes32 salt,
        bytes memory creationCode
    )
        external
        pure
        returns (address deployedAt)
    {
        return Create2.predictDeployment(deployer, salt, creationCode);
    }
}
