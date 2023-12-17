// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IAdmin {
    // ============ Events ============

    event RelayerAdded(address relayer);
    event RelayerRemoved(address relayer);

    event GuardAdded(address guard);
    event GuardRemoved(address guard);

    event GovernorAdded(address governor);
    event GovernorRemoved(address governor);

    event FeeRateUpdated(uint256 oldFeeRate, uint256 newFeeRate);
    event FeesSwept(address token, address recipient, uint256 amount);

    // ============ Methods ============

    function addRelayer(address _relayer) external;

    function removeRelayer(address _relayer) external;

    function addGuard(address _guard) external;

    function removeGuard(address _guard) external;

    function addGovernor(address _governor) external;

    function removeGovernor(address _governor) external;

    function setProtocolFeeRate(uint256 newFeeRate) external;

    function sweepProtocolFees(address token, address recipient) external;
}
