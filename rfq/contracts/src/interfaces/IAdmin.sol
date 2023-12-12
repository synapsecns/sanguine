// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IAdmin {
    // ============ Events ============

    event RelayerAdded(address relayer);
    event RelayerRemoved(address relayer);

    event GuardAdded(address guard);
    event GuardRemoved(address guard);

    // ============ Methods ============

    function addRelayer(address _relayer) external;

    function removeRelayer(address _relayer) external;

    function addGuard(address _guard) external;

    function removeGuard(address _guard) external;
}
