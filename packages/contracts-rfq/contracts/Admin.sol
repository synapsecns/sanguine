// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Ownable2Step, Ownable} from "@openzeppelin/contracts/access/Ownable2Step.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

import {IAdmin} from "./interfaces/IAdmin.sol";

contract Admin is IAdmin, Ownable2Step, AccessControl {
    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");
    bytes32 public constant GUARD_ROLE = keccak256("GUARD_ROLE");

    modifier onlyGuard() {
        require(hasRole(GUARD_ROLE, msg.sender), "Caller is not a guard");
        _;
    }

    modifier onlyRelayer() {
        require(hasRole(RELAYER_ROLE, msg.sender), "Caller is not a relayer");
        _;
    }

    constructor(address _owner) Ownable(_owner) {
        _grantRole(DEFAULT_ADMIN_ROLE, _owner);
    }

    function addRelayer(address _relayer) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _grantRole(RELAYER_ROLE, _relayer);
        emit RelayerAdded(_relayer);
    }

    function removeRelayer(address _relayer) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _revokeRole(RELAYER_ROLE, _relayer);
        emit RelayerRemoved(_relayer);
    }

    function addGuard(address _guard) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _grantRole(GUARD_ROLE, _guard);
        emit GuardAdded(_guard);
    }

    function removeGuard(address _guard) external {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender));
        _revokeRole(GUARD_ROLE, _guard);
        emit GuardRemoved(_guard);
    }
}
