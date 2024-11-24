// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BridgeTransactionV2Lib, FastBridgeV2SrcTest, IFastBridgeV2} from "./FastBridgeV2.Src.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract FastBridgeV2SrcRefundV1Test is FastBridgeV2SrcTest {
    function cancel(address caller, IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) public virtual override {
        vm.prank({msgSender: caller, txOrigin: caller});
        fastBridge.refund(BridgeTransactionV2Lib.encodeV2(bridgeTx));
    }
}
