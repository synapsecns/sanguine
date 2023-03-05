// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./BondingManager.t.sol";

// solhint-disable func-name-mixedcase
contract BondingSecondaryTest is BondingManagerTest {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TESTS: SETUP                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_constructor_revert_onSynapseChain() public {
        // Should not be able to deploy on Synapse Chain
        vm.expectRevert("Can't be deployed on SynChain");
        new BondingSecondary(DOMAIN_SYNAPSE);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SLASH AGENT REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // slashAgent(): localDomain_notOrigin is tested in BondingManager.t.sol

    function test_slashAgent_revert_synapseDomain_notBondingManager() public {
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject system calls from Synapse domain, if caller is not BondingManager
            SystemEntity caller = SystemEntity(c);
            if (caller == SystemEntity.BondingManager) continue;
            vm.expectRevert("!allowedCaller");
            // Use mocked agent info
            _mockSlashAgentCall({
                callOrigin: DOMAIN_SYNAPSE,
                systemCaller: caller,
                info: guardInfo({ guard: address(0), bonded: false })
            });
        }
    }

    function test_slashAgent_revert_remoteNotSynapseDomain(uint32 callOrigin) public {
        // Exclude local calls and calls from Synapse Chain
        vm.assume(callOrigin != DOMAIN_LOCAL && callOrigin != DOMAIN_SYNAPSE);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject cross-chain system calls from domains other than Synapse domain
            SystemEntity caller = SystemEntity(c);
            vm.expectRevert("!synapseDomain");
            // Use mocked agent info
            _mockSlashAgentCall({
                callOrigin: callOrigin,
                systemCaller: caller,
                info: guardInfo({ guard: address(0), bonded: false })
            });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SYNC AGENTS REVERTS                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgents_revert_localDomain() public {
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject all system calls from local domain
            SystemEntity caller = SystemEntity(c);
            // Calls from local domain never pass the optimistic period check
            vm.expectRevert("!optimisticPeriod");
            // Use mocked list of agents
            _mockSyncAgentsCall({
                callOrigin: localDomain,
                systemCaller: caller,
                requestID: 0,
                removeExisting: false,
                infos: new AgentInfo[](0)
            });
        }
    }

    function test_syncAgents_revert_remoteNotSynapseDomain(uint32 callOrigin) public {
        // Exclude local calls and calls from Synapse Chain
        vm.assume(callOrigin != localDomain && callOrigin != DOMAIN_SYNAPSE);
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            // Should reject all system calls from remote domains other than Synapse domain
            SystemEntity caller = SystemEntity(c);
            vm.expectRevert("!synapseDomain");
            // Use mocked list of agents
            _mockSyncAgentsCall({
                callOrigin: callOrigin,
                systemCaller: caller,
                requestID: 0,
                removeExisting: false,
                infos: new AgentInfo[](0)
            });
        }
    }

    function test_syncAgents_revert_synapseDomain_notBondingManager() public {
        _skipBondingOptimisticPeriod();
        for (uint256 c = 0; c < uint8(type(SystemEntity).max); ++c) {
            SystemEntity caller = SystemEntity(c);
            // Should reject system calls from Synapse domain, if caller is not BondingManager
            if (caller == SystemEntity.BondingManager) continue;
            vm.expectRevert("!allowedCaller");
            // Use mocked list of agents
            _mockSyncAgentsCall({
                callOrigin: DOMAIN_SYNAPSE,
                systemCaller: caller,
                requestID: 0,
                removeExisting: false,
                infos: new AgentInfo[](0)
            });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║              TESTS: RECEIVE SYSTEM CALLS (LOCAL DOMAIN)              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_slashAgent_localDomain_origin(uint32 domain, address account) public {
        AgentInfo memory info = agentInfo({ domain: domain, account: account, bonded: false });
        bytes memory data = abi.encodeWithSelector(
            ISystemContract.slashAgent.selector,
            0, // rootSubmittedAt
            0, // callOrigin
            0, // systemCaller
            info
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            // (_rootSubmittedAt, _callOrigin, _caller, _info)
            vm.expectCall(
                systemRegistries[r],
                abi.encodeWithSelector(
                    SystemContractMock.slashAgent.selector,
                    block.timestamp,
                    DOMAIN_LOCAL,
                    SystemEntity.BondingManager,
                    info
                )
            );
        }
        // data should be forwarded to Synapse Chain
        vm.expectCall(
            address(systemRouter),
            abi.encodeWithSelector(
                SystemRouter.systemCall.selector,
                DOMAIN_SYNAPSE, // destination
                BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                SystemEntity.BondingManager, //recipient
                data
            )
        );
        // Mock a local system call: [Local Origin] -> [Local BondingManager].slashAgent
        systemRouter.mockSystemCall({
            _recipient: SystemEntity.BondingManager,
            _rootSubmittedAt: block.timestamp,
            _callOrigin: localDomain,
            _systemCaller: SystemEntity.Origin,
            _data: data
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║             TESTS: RECEIVE SYSTEM CALLS (SYNAPSE CHAIN)              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_syncAgents_synapseDomain_bondingManager(
        uint256 requestID,
        bool removeExisting,
        uint32 domain,
        address account,
        bool bonded
    ) public {
        _skipBondingOptimisticPeriod();
        AgentInfo[] memory infos = infoToArray(agentInfo(domain, account, bonded));
        // Data for the system call
        bytes memory data = abi.encodeWithSelector(
            ISystemContract.syncAgents.selector,
            0, // rootSubmittedAt
            0, // callOrigin
            0, // systemCaller
            requestID,
            removeExisting,
            infos
        );
        // Empty array should be passed back
        bytes memory forwardedData = abi.encodeWithSelector(
            ISystemContract.syncAgents.selector,
            0, // rootSubmittedAt
            0, // callOrigin
            0, // systemCaller
            requestID,
            removeExisting,
            new AgentInfo[](0)
        );
        // All system registries should be system called
        for (uint256 r = 0; r < systemRegistries.length; ++r) {
            // (_rootSubmittedAt, _callOrigin, _caller, _requestID, _removeExisting, _infos)
            vm.expectCall(
                systemRegistries[r],
                abi.encodeWithSelector(
                    SystemContractMock.syncAgents.selector,
                    block.timestamp,
                    DOMAIN_LOCAL,
                    SystemEntity.BondingManager,
                    requestID,
                    removeExisting,
                    infos
                )
            );
        }
        // data should be forwarded to Synapse Chain
        vm.expectCall(
            address(systemRouter),
            abi.encodeWithSelector(
                SystemRouter.systemCall.selector,
                DOMAIN_SYNAPSE, // destination
                BONDING_OPTIMISTIC_PERIOD, // optimisticSeconds
                SystemEntity.BondingManager, //recipient
                forwardedData
            )
        );
        // Mock a system call: [SynapseChain BondingManager] -> [Local BondingManager].syncAgents
        systemRouter.mockSystemCall({
            _recipient: SystemEntity.BondingManager,
            _rootSubmittedAt: rootSubmittedAt,
            _callOrigin: DOMAIN_SYNAPSE,
            _systemCaller: SystemEntity.BondingManager,
            _data: data
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _deployBondingManager(uint32 domain) internal override returns (BondingManager) {
        return new BondingSecondary(domain);
    }

    function _getTestLocalDomain() internal pure override returns (uint32) {
        return DOMAIN_LOCAL;
    }
}
