// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {ChainGas, InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";

import {Random} from "../utils/libs/Random.t.sol";
import {RawAttestation, RawSnapshot, RawState} from "../utils/libs/SynapseStructs.t.sol";
import {AgentFlag, AgentStatus, BondingManager, Destination, SynapseTest} from "../utils/SynapseTest.t.sol";
import {ExecutionHubTest} from "./hubs/ExecutionHub.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract DestinationSynapseTest is ExecutionHubTest {
    // Deploy Production version of Destination and Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_DESTINATION_SYNAPSE | DEPLOY_PROD_SUMMIT) {}

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = DOMAIN_SYNAPSE;
        address caller = random.nextAddress();
        BondingManager manager = new BondingManager(domain);
        bytes32 agentRoot = random.next();
        Destination cleanContract = new Destination(domain, address(manager));
        manager.initialize(address(0), address(cleanContract), address(0));
        // agentRoot should be ignored on Synapse Chain
        vm.prank(caller);
        cleanContract.initialize(agentRoot);
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.agentManager(), address(manager), "!agentManager");
        assertEq(cleanContract.nextAgentRoot(), 0, "!nextAgentRoot");
    }

    function initializeLocalContract() public override {
        Destination(localContract()).initialize(0);
    }

    function test_getAttestation(Random memory random) public {
        uint256 amount = 10;
        bytes[] memory attPayloads = new bytes[](amount);
        uint256 statesAmount = allDomains.length - 1;
        for (uint32 index = 0; index < amount; ++index) {
            RawSnapshot memory rs;
            rs.states = new RawState[](statesAmount);
            for (uint256 d = 0; d < statesAmount; ++d) {
                rs.states[d] = random.nextState({origin: allDomains[d + 1], nonce: index + 1});
            }
            RawAttestation memory ra = rs.castToRawAttestation({
                agentRoot: getAgentRoot(),
                nonce: index + 1,
                blockNumber: uint40(block.number),
                timestamp: uint40(block.timestamp)
            });
            attPayloads[index] = ra.formatAttestation();
            address guard = domains[0].agents[index % DOMAIN_AGENTS];
            address notary = domains[DOMAIN_LOCAL].agents[index % DOMAIN_AGENTS];
            (bytes memory snapPayload, bytes memory guardSignature) = signSnapshot(guard, rs);
            (, bytes memory notarySignature) = signSnapshot(notary, rs);
            bondingManager.submitSnapshot(snapPayload, guardSignature);
            bondingManager.submitSnapshot(snapPayload, notarySignature);
            skipBlock();
        }
        for (uint32 index = 0; index < amount; ++index) {
            (bytes memory attPayload, bytes memory attSignature) =
                InterfaceDestination(destinationSynapse).getAttestation(index);
            assertEq(attPayload, attPayloads[index], "!payload");
            assertEq(attSignature, "", "!signature");
        }
    }

    function test_acceptAttestation_revert_notAgentManager(address caller) public {
        vm.assume(caller != localAgentManager());
        vm.expectRevert("!agentManager");
        vm.prank(caller);
        InterfaceDestination(destinationSynapse).acceptAttestation(0, 0, "", new ChainGas[](0));
    }

    // ══════════════════════════════════════════════════ HELPERS ══════════════════════════════════════════════════════

    /// @notice Prepares execution of the created messages
    function prepareExecution(SnapshotMock memory sm)
        public
        override
        returns (bytes32 snapRoot, bytes32[] memory snapProof)
    {
        RawAttestation memory ra;
        (ra, snapProof) = createSnapshotProof(sm);
        snapRoot = ra.snapRoot;
        (bytes memory snapPayload, bytes memory guardSignature) = createSignedSnapshot(domains[0].agent, sm.rs, sm.rsi);
        (, bytes memory notarySignature) = createSignedSnapshot(domains[DOMAIN_LOCAL].agent, sm.rs, sm.rsi);
        bondingManager.submitSnapshot(snapPayload, guardSignature);
        bondingManager.submitSnapshot(snapPayload, notarySignature);
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
