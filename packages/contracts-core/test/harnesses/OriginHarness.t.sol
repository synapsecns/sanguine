// SPDX-License-Identifier: MIT

pragma solidity 0.8.17;

import { OriginHub } from "../../contracts/hubs/OriginHub.sol";
import { AgentSet } from "../../contracts/libs/AgentSet.sol";
import "../../contracts/libs/Message.sol";
import { Origin } from "../../contracts/Origin.sol";

import { AgentRegistryExtended } from "./system/AgentRegistryExtended.t.sol";
import { SystemContractHarness } from "./system/SystemContractHarness.t.sol";

contract OriginHarness is Origin, AgentRegistryExtended, SystemContractHarness {
    using AgentSet for AgentSet.DomainAddressSet;

    //solhint-disable-next-line no-empty-blocks
    constructor(uint32 _domain) Origin(_domain) {}

    function addLocalNotary(address _notary) external {
        agents[_currentEpoch()].add(_localDomain(), _notary);
    }

    function removeAllAgents(uint32 _domain) public {
        uint256 amount = amountAgents(_domain);
        // Remove every Agent to halt the contract
        for (uint256 i = 0; i < amount; ++i) {
            _removeAgent(_domain, getAgent({ _domain: _domain, _agentIndex: 0 }));
        }
    }

    function getNextMessage(
        uint32 _destination,
        bytes32 _recipientAddress,
        uint32 _optimisticSeconds,
        bytes memory _tips,
        bytes memory _messageBody
    ) public view returns (bytes memory message) {
        message = MessageLib.formatMessage(
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
        returns (
            uint32 latestNonce,
            bytes32 latestRoot,
            uint40 blockNumber,
            uint40 timestamp
        )
    {
        latestNonce = nonce(_destination);
        if (latestNonce == 0) {
            latestRoot = EMPTY_TREE_ROOT;
        } else {
            latestRoot = historicalRoots[_destination][latestNonce];
            blockNumber = historicalMetadata[_destination][latestNonce].blockNumber;
            timestamp = historicalMetadata[_destination][latestNonce].timestamp;
        }
    }

    function _isIgnoredAgent(uint32 _domain, address _account)
        internal
        view
        override(AgentRegistryExtended, OriginHub)
        returns (bool)
    {
        return OriginHub._isIgnoredAgent(_domain, _account);
    }
}
