// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {MockERC20} from "./MockERC20.sol";

// solhint-disable no-empty-blocks
contract WeirdERC20Mock is MockERC20 {
    uint256 public transferToFastBridgeValue;
    address public fastBridge;

    constructor(string memory name_, uint8 decimals_) MockERC20(name_, decimals_) {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testWeirdERC20Mock() external {}

    function setFastBridge(address fastBridge_) public {
        fastBridge = fastBridge_;
    }

    function setTransferToFastBridgeValue(uint256 value) public {
        transferToFastBridgeValue = value;
    }

    function transfer(address to, uint256 value) public virtual override returns (bool) {
        if (to == fastBridge) {
            return super.transfer(to, transferToFastBridgeValue);
        } else {
            return super.transfer(to, value);
        }
    }

    function transferFrom(address from, address to, uint256 value) public virtual override returns (bool) {
        if (to == fastBridge) {
            return super.transferFrom(from, to, transferToFastBridgeValue);
        } else {
            return super.transferFrom(from, to, value);
        }
    }
}
