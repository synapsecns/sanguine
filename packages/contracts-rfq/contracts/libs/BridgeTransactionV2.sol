// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {IFastBridgeV2} from "../interfaces/IFastBridgeV2.sol";

library BridgeTransactionV2Lib {
    function encodeV2(IFastBridgeV2.BridgeTransactionV2 memory bridgeTx) internal pure returns (bytes memory) {
        return abi.encode(bridgeTx);
    }

    function decodeV2(bytes memory encodedTx) internal pure returns (IFastBridgeV2.BridgeTransactionV2 memory) {
        return abi.decode(encodedTx, (IFastBridgeV2.BridgeTransactionV2));
    }
}
