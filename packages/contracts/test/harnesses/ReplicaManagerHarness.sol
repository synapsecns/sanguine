// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { ReplicaManager } from "../../contracts/ReplicaManager.sol";

import { ReplicaLib } from "../../contracts/libs/Replica.sol";
import { Tips } from "../../contracts/libs/Tips.sol";

contract ReplicaManagerHarness is ReplicaManager {
    using ReplicaLib for ReplicaLib.Replica;

    using Tips for bytes29;

    event LogTips(uint96 updaterTip, uint96 relayerTip, uint96 proverTip, uint96 processorTip);

    constructor(
        uint32 _localDomain,
        uint256 _processGas,
        uint256 _reserveGas
    ) ReplicaManager(_localDomain, _processGas, _reserveGas) {}

    function setMessageStatus(
        uint32 _remoteDomain,
        bytes32 _messageHash,
        bytes32 _status
    ) external {
        allReplicas[activeReplicas[_remoteDomain]].setMessageStatus(_messageHash, _status);
    }

    function _storeTips(bytes29 _tips) internal override {
        emit LogTips(
            _tips.updaterTip(),
            _tips.relayerTip(),
            _tips.proverTip(),
            _tips.processorTip()
        );
    }
}
