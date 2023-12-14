// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Admin} from "../contracts/Admin.sol";
import {IFastBridge} from "../contracts/interfaces/IFastBridge.sol";

contract FastBridgeMock is IFastBridge, Admin {
    // @dev the block the contract was deployed at
    uint256 public immutable deployBlock;

    constructor(address _owner) Admin(_owner) {
        deployBlock = block.number;
    }

    /// @dev to prevent replays
    uint256 public nonce;

    function mockBridgeRequest(bytes32 transactionId, address sender, BridgeParams memory params) external {
        bytes memory request = abi.encode(
            BridgeTransaction({
                originChainId: uint32(block.chainid),
                destChainId: params.dstChainId,
                originSender: msg.sender,
                destRecipient: params.to,
                originToken: params.originToken,
                destToken: params.destToken,
                originAmount: params.originAmount, // includes relayer fee
                destAmount: params.destAmount,
                deadline: params.deadline,
                nonce: nonce++ // increment nonce on every bridge
            })
        );

        emit BridgeRequested(transactionId, msg.sender, request);
    }

    function mockBridgeRequestRaw(bytes32  transactionId, address sender, bytes memory request) external  {
        emit BridgeRequested(transactionId, sender, request);
    }

    function mockBridgeRelayer(bytes32 transactionId, address relayer, address to, address token, uint256 amount) external  {
        emit BridgeRelayed(transactionId, relayer, to, token, amount);
    }

    function bridge(BridgeParams memory params) external payable{
        // do nothing
    }

    function relay(bytes memory request) external payable {
        // do nothing
    }

    function prove(bytes memory request, bytes32 destTxHash) external {
        // do nothing
    }

    function claim(bytes memory request, address to) external {
        // do nothing
    }

    function dispute(bytes32 transactionId) external {
        // do nothing
    }

    function refund(bytes memory request, address to) external {
        // do nothing
    }
}
