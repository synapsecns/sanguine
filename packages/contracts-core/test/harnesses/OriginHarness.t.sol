// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { Message } from "../../contracts/libs/Message.sol";
import { Origin } from "../../contracts/Origin.sol";
import { GuardRegistryHarness } from "./registry/GuardRegistryHarness.t.sol";
import { SystemContractHarness } from "./system/SystemContractHarness.t.sol";

contract OriginHarness is Origin, SystemContractHarness, GuardRegistryHarness {
    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) Origin(_domain) {}

    function removeAllNotaries(uint32 _domain) public {
        uint256 amount = notariesAmount(_domain);
        // Remove every Notary to halt the contract
        for (uint256 i = 0; i < amount; ++i) {
            _removeNotary(_domain, getNotary({ _domain: _domain, _index: 0 }));
        }
    }

    function isNotary(uint32 _domain, address _notary) public view returns (bool) {
        return _isNotary(_domain, _notary);
    }

    function getNextMessage(
        uint32 _destination,
        bytes32 _recipientAddress,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) public view returns (bytes memory message) {
        message = Message.formatMessage(
            _localDomain(),
            _checkForSystemRouter(_recipientAddress),
            nonce(_destination) + 1,
            _destination,
            _recipientAddress,
            _optimisticSeconds,
            _tips,
            _messageBody
        );
    }

    function suggestNonceRoot(uint32 _destination)
        public
        view
        returns (uint32 latestNonce, bytes32 latestRoot)
    {
        latestNonce = nonce(_destination);
        latestRoot = getHistoricalRoot(_destination, latestNonce);
    }
}
