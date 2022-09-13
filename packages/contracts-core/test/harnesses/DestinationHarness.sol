// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import { Destination } from "../../contracts/Destination.sol";

import { MirrorLib } from "../../contracts/libs/Mirror.sol";
import { Tips } from "../../contracts/libs/Tips.sol";
import { ISystemRouter } from "../../contracts/interfaces/ISystemRouter.sol";

import { GuardRegistryHarness } from "./GuardRegistryHarness.sol";
import { SystemContractHarness } from "./SystemContractHarness.sol";

contract DestinationHarness is Destination, SystemContractHarness, GuardRegistryHarness {
    using MirrorLib for MirrorLib.Mirror;
    using Tips for bytes29;

    event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip);

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _localDomain) Destination(_localDomain) {}

    function addNotary(uint32 _domain, address _notary) public {
        _addNotary(_domain, _notary);
    }

    function removeNotary(uint32 _domain, address _notary) public {
        _removeNotary(_domain, _notary);
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
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
