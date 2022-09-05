// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Destination } from "../../contracts/Destination.sol";

import { MirrorLib } from "../../contracts/libs/Mirror.sol";
import { Tips } from "../../contracts/libs/Tips.sol";

import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";

contract DestinationHarness is Destination, GuardRegistryHarness {
    using MirrorLib for MirrorLib.Mirror;

    uint256 public sensitiveValue;
    using Tips for bytes29;

    event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip);

    constructor(uint32 _localDomain) Destination(_localDomain) {}

    function addNotary(uint32 _domain, address _notary) public {
        _addNotary(_domain, _notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }

    function setSensitiveValue(uint256 _newValue) external onlySystemRouter {
        sensitiveValue = _newValue;
    }

    function setMessageStatus(
        uint32 _remoteDomain,
        bytes32 _messageHash,
        bytes32 _status
    ) external {
        allMirrors[activeMirrors[_remoteDomain]].setMessageStatus(_messageHash, _status);
    }

    function _storeTips(bytes29 _tips) internal override {
        emit LogTips(
            _tips.notaryTip(),
            _tips.broadcasterTip(),
            _tips.proverTip(),
            _tips.executorTip()
        );
    }
}
