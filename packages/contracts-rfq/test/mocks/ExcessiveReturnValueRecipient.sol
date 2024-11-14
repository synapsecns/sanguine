// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import {IZapRecipient} from "../../contracts/interfaces/IZapRecipient.sol";

// solhint-disable no-empty-blocks
/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract ExcessiveReturnValueRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testExcessiveReturnValueRecipient() external {}

    /// @notice Incorrectly implemented - method returns excessive bytes.
    function zap(address, uint256, bytes memory) external payable returns (bytes4, uint256) {
        return (IZapRecipient.zap.selector, 1337);
    }
}
