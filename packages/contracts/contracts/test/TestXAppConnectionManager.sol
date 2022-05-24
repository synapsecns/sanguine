// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../XAppConnectionManager.sol";
import "../../libs/TypeCasts.sol";

contract TestXAppConnectionManager is XAppConnectionManager {
    constructor() XAppConnectionManager() {} // solhint-disable-line no-empty-blocks

    function testRecoverWatcherFromSig(
        uint32 _domain,
        address _replica,
        address _updater,
        bytes memory _signature
    ) external view returns (address) {
        return
            _recoverWatcherFromSig(
                _domain,
                TypeCasts.addressToBytes32(_replica),
                TypeCasts.addressToBytes32(_updater),
                _signature
            );
    }
}
