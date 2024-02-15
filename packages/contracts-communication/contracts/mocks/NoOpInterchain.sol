pragma solidity 0.8.20;

import {IInterchain} from "../IInterchain.sol";

contract NoOpInterchain is IInterchain {
    function interchainSend(
        bytes32 receiver,
        uint256 dstChainId,
        bytes calldata message,
        address[] calldata modules
    ) external payable override {
        // Do nothing
    }

    function interchainReceive(bytes calldata transaction) external override {
        // Do nothing
    }
}
