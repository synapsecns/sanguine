// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {FastBridgeV2, IFastBridgeV2} from "../../contracts/FastBridgeV2.sol";
import {BridgeTransactionV2Lib} from "../../contracts/libs/BridgeTransactionV2.sol";

import {IFastBridge, MulticallTargetIntegrationTest} from "./MulticallTarget.t.sol";

contract FastBridgeV2MulticallTargetTest is MulticallTargetIntegrationTest {
    function deployAndConfigureFastBridge() public override returns (address) {
        FastBridgeV2 fb = new FastBridgeV2(address(this));
        fb.addProver(relayer);
        return address(fb);
    }

    function getEncodedBridgeTx(IFastBridge.BridgeTransaction memory bridgeTx)
        public
        pure
        override
        returns (bytes memory)
    {
        IFastBridgeV2.BridgeTransactionV2 memory bridgeTxV2;
        bridgeTxV2.originChainId = bridgeTx.originChainId;
        bridgeTxV2.destChainId = bridgeTx.destChainId;
        bridgeTxV2.originSender = bridgeTx.originSender;
        bridgeTxV2.destRecipient = bridgeTx.destRecipient;
        bridgeTxV2.originToken = bridgeTx.originToken;
        bridgeTxV2.destToken = bridgeTx.destToken;
        bridgeTxV2.originAmount = bridgeTx.originAmount;
        bridgeTxV2.destAmount = bridgeTx.destAmount;
        bridgeTxV2.originFeeAmount = bridgeTx.originFeeAmount;
        bridgeTxV2.deadline = bridgeTx.deadline;
        bridgeTxV2.nonce = bridgeTx.nonce;
        return BridgeTransactionV2Lib.encodeV2(bridgeTxV2);
    }
}
