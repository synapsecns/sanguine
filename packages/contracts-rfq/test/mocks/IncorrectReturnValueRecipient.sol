// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IZapRecipient} from "../../contracts/interfaces/IZapRecipient.sol";

// solhint-disable no-empty-blocks
/// @notice Incorrectly implemented recipient mock for testing purposes. DO NOT USE IN PRODUCTION.
contract IncorrectReturnValueRecipient is IZapRecipient {
    /// @notice Mock needs to accept ETH
    receive() external payable {}

    /// @notice We include an empty "test" function so that this contract does not appear in the coverage report.
    function testIncorrectReturnValueRecipient() external {}

    /// @notice Incorrectly implemented - method returns incorrect value.
    function zap(address, uint256, bytes memory) external payable returns (bytes4) {
        // Flip the last bit
        return IZapRecipient.zap.selector ^ 0x00000001;
    }
}
