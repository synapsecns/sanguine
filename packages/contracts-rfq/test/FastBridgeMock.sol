// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Admin} from "../contracts/Admin.sol";
import {IFastBridge} from "../contracts/interfaces/IFastBridge.sol";
import { FastBridge } from "../contracts/FastBridge.sol";

contract FastBridgeMock is IFastBridge, Admin {
    // @dev the block the contract was deployed at
    uint256 public immutable deployBlock;

    constructor(address _owner) Admin(_owner) {
        deployBlock = block.number;
    }

    /// @dev to prevent replays
    uint256 public nonce;

    function getBridgeTransaction(bytes memory request) public pure returns (BridgeTransaction memory) {
        return abi.decode(request, (BridgeTransaction));
    }

    // used for testing in go.
    // see: https://ethereum.stackexchange.com/questions/21155/how-to-expose-enum-in-solidity-contract
    // make sure to update fastbridge/status.go if this changes
    // or underliyng enum changes.
    //
    // TODO: a foundry test should be added to ensure this is always in sync.
    function getEnumKeyByValue (FastBridge.BridgeStatus keyValue) public pure returns (string memory) {
        if (FastBridge.BridgeStatus.NULL == keyValue) return "NULL";
        if (FastBridge.BridgeStatus.REQUESTED == keyValue) return "REQUESTED";
        if (FastBridge.BridgeStatus.RELAYER_PROVED == keyValue) return "RELAYER_PROVED";
        if (FastBridge.BridgeStatus.RELAYER_CLAIMED == keyValue) return "RELAYER_CLAIMED";
        if (FastBridge.BridgeStatus.REFUNDED == keyValue) return "REFUNDED";
        return "";
    }

    function mockBridgeRequest(bytes32 transactionId, address sender, BridgeParams memory params) external {
        uint256 originFeeAmount = (params.originAmount * protocolFeeRate) / FEE_BPS;
        params.originAmount -= originFeeAmount;

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
                originFeeAmount: originFeeAmount,
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
        revert("not implemented");
    }

    function relay(bytes memory request) external payable {
        revert("not implemented");
    }

    function prove(bytes memory request, bytes32 destTxHash) external {
        revert("not implemented");
    }
    
    function canClaim(bytes32 transactionid, address relayer) external view returns (bool) {
        revert("not implemented");
    }

    function claim(bytes memory request, address to) external {
        revert("not implemented");
    }

    function dispute(bytes32 transactionId) external {
        revert("not implemented");
    }

    function refund(bytes memory request, address to) external {
        revert("not implemented");
    }
}
