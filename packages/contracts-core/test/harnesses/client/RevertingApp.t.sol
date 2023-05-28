// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {IMessageRecipient} from "../../../contracts/interfaces/IMessageRecipient.sol";

contract RevertingApp is IMessageRecipient {
    bool private willRevert = true;

    receive() external payable {
        if (willRevert) {
            revert("GM");
        }
    }

    function toggleRevert(bool willRevert_) external {
        willRevert = willRevert_;
    }

    function receiveBaseMessage(uint32, uint32, bytes32, uint256, uint32, bytes memory) external payable {
        if (willRevert) {
            revert("GM");
        }
    }
}
