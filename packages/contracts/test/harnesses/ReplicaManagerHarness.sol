// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { ReplicaManager } from "../../contracts/ReplicaManager.sol";

import { ReplicaLib } from "../../contracts/libs/Replica.sol";
import { Tips } from "../../contracts/libs/Tips.sol";

contract ReplicaManagerHarness is ReplicaManager {
    using ReplicaLib for ReplicaLib.Replica;

    uint256 public sensitiveValue;
    using Tips for bytes29;

    event LogTips(uint96 updaterTip, uint96 relayerTip, uint96 proverTip, uint96 processorTip);

    constructor(uint32 _localDomain) ReplicaManager(_localDomain) {}

    function addNotary(uint32 _domain, address _notary) public {
        _addNotary(_domain, _notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }

    function setSensitiveValue(uint256 _newValue) external onlySystemMessenger {
        sensitiveValue = _newValue;
    }

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
