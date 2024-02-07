// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {DISPUTE_TIMEOUT_NOTARY} from "../../contracts/libs/Constants.sol";
import {CallerNotInbox, DisputeTimeoutNotOver, NotaryInDispute} from "../../contracts/libs/Errors.sol";
import {ChainGas, GasData, InterfaceDestination} from "../../contracts/interfaces/InterfaceDestination.sol";

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

    function test_setupCorrectly() public {
        Destination dst = Destination(localDestination());
        // Check Agents: all Agents are known in BondingManager
        for (uint256 d = 0; d < allDomains.length; ++d) {
            uint32 domain = allDomains[d];
            for (uint256 i = 0; i < domains[domain].agents.length; ++i) {
                address agent = domains[domain].agents[i];
                checkAgentStatus(agent, dst.agentStatus(agent), AgentFlag.Active);
            }
        }
        // Check AgentManager and Inbox
        assertEq(dst.agentManager(), localAgentManager(), "!agentManager");
        assertEq(dst.inbox(), localInbox(), "!inbox");
        // Check version
        assertEq(dst.version(), LATEST_VERSION, "!version");
        // Check pending Agent Merkle Root
        bool rootPending = dst.passAgentRoot();
        assertFalse(rootPending);
    }

    function test_cleanSetup(Random memory random) public override {
        uint32 domain = DOMAIN_SYNAPSE;
        vm.chainId(domain);
        address caller = random.nextAddress();
        BondingManager manager = new BondingManager(domain);
        address inbox_ = random.nextAddress();
        bytes32 agentRoot = random.next();
        Destination cleanContract = new Destination(DOMAIN_SYNAPSE, address(manager), inbox_);
        manager.initialize(address(0), address(cleanContract), inbox_, address(0));
        // agentRoot should be ignored on Synapse Chain
        vm.prank(caller);
        cleanContract.initialize(agentRoot);
        assertEq(cleanContract.owner(), caller, "!owner");
        assertEq(cleanContract.localDomain(), domain, "!localDomain");
        assertEq(cleanContract.agentManager(), address(manager), "!agentManager");
        assertEq(cleanContract.inbox(), inbox_, "!inbox");
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
            inbox.submitSnapshot(snapPayload, guardSignature);
            inbox.submitSnapshot(snapPayload, notarySignature);
            skipBlock();
        }
        for (uint32 index = 0; index < amount; ++index) {
            (bytes memory attPayload, bytes memory attSignature) =
                InterfaceDestination(localDestination()).getAttestation(index);
            assertEq(attPayload, attPayloads[index], "!payload");
            assertEq(attSignature, "", "!signature");
        }
    }

    function test_acceptAttestation_revert_notInbox(address caller) public {
        vm.assume(caller != localInbox());
        vm.expectRevert(CallerNotInbox.selector);
        vm.prank(caller);
        InterfaceDestination(localDestination()).acceptAttestation(0, 0, "", 0, new ChainGas[](0));
    }

    function test_acceptAttestation_revert_notaryInDispute(uint256 domainId, uint256 notaryId) public {
        address notary = getNotary(domainId, notaryId);
        openDispute({guard: domains[0].agent, notary: notary});
        vm.prank(address(inbox));
        vm.expectRevert(NotaryInDispute.selector);
        InterfaceDestination(localDestination()).acceptAttestation(agentIndex[notary], 0, "", 0, new ChainGas[](0));
    }

    // ═════════════════════════════════════════════════ GAS DATA ══════════════════════════════════════════════════════

    function test_getGasData(Random memory random) public {
        RawSnapshot memory firstSnap;
        firstSnap.states = new RawState[](2);
        firstSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: 1});
        firstSnap.states[1] = random.nextState({origin: DOMAIN_LOCAL, nonce: 1});
        address firstNotary = domains[DOMAIN_LOCAL].agents[0];
        (bytes memory firstSnapPayload, bytes memory firstNotarySignature) = signSnapshot(firstNotary, firstSnap);
        (, bytes memory firstGuardSignature) = signSnapshot(domains[0].agent, firstSnap);
        inbox.submitSnapshot(firstSnapPayload, firstGuardSignature);
        inbox.submitSnapshot(firstSnapPayload, firstNotarySignature);
        uint256 firstSkipTime = random.nextUint32();
        skip(firstSkipTime);
        RawSnapshot memory secondSnap;
        secondSnap.states = new RawState[](1);
        secondSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: 2});
        vm.assume(
            GasData.unwrap(firstSnap.states[0].castToState().gasData())
                != GasData.unwrap(secondSnap.states[0].castToState().gasData())
        );
        address secondNotary = domains[DOMAIN_REMOTE].agents[0];
        (bytes memory secondSnapPayload, bytes memory secondNotarySignature) = signSnapshot(secondNotary, secondSnap);
        (, bytes memory secondGuardSignature) = signSnapshot(domains[0].agent, secondSnap);
        inbox.submitSnapshot(secondSnapPayload, secondGuardSignature);
        inbox.submitSnapshot(secondSnapPayload, secondNotarySignature);
        uint256 secondSkipTime = random.nextUint32();
        skip(secondSkipTime);
        // Check getGasData
        GasData firstRemoteGasData = firstSnap.states[0].castToState().gasData();
        GasData firstLocalGasData = firstSnap.states[1].castToState().gasData();
        GasData secondRemoteGasData = secondSnap.states[0].castToState().gasData();
        emit log_named_uint("Remote gasData: first", GasData.unwrap(firstRemoteGasData));
        emit log_named_uint("Remote gasData: second", GasData.unwrap(secondRemoteGasData));
        emit log_named_uint("Local gasData: first", GasData.unwrap(firstLocalGasData));
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(secondRemoteGasData), "!remoteGasData");
        assertEq(dataMaturity, secondSkipTime, "!remoteDataMaturity");
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_LOCAL);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(firstLocalGasData), "!localGasData");
        assertEq(dataMaturity, firstSkipTime + secondSkipTime, "!localDataMaturity");
    }

    function test_getGasData_localDomain(Random memory random) public {
        RawSnapshot memory firstSnap;
        firstSnap.states = new RawState[](2);
        firstSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: 1});
        firstSnap.states[1] = random.nextState({origin: localDomain(), nonce: 1});
        address firstNotary = domains[DOMAIN_LOCAL].agents[0];
        (bytes memory firstSnapPayload, bytes memory firstNotarySignature) = signSnapshot(firstNotary, firstSnap);
        (, bytes memory firstGuardSignature) = signSnapshot(domains[0].agent, firstSnap);
        inbox.submitSnapshot(firstSnapPayload, firstGuardSignature);
        inbox.submitSnapshot(firstSnapPayload, firstNotarySignature);
        uint256 firstSkipTime = random.nextUint32();
        skip(firstSkipTime);
        // Check getGasData
        GasData firstRemoteGasData = firstSnap.states[0].castToState().gasData();
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(firstRemoteGasData), "!remoteGasData");
        assertEq(dataMaturity, firstSkipTime, "!remoteDataMaturity");
        // Should not save data for local domain
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(localDomain());
        assertEq(GasData.unwrap(gasData), 0, "!localGasData");
        assertEq(dataMaturity, 0, "!localDataMaturity");
    }

    function test_getGasData_noDataForDomain(Random memory random, uint32 domain) public {
        vm.assume(domain != DOMAIN_REMOTE && domain != DOMAIN_LOCAL);
        RawSnapshot memory firstSnap;
        firstSnap.states = new RawState[](2);
        firstSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: 1});
        firstSnap.states[1] = random.nextState({origin: DOMAIN_LOCAL, nonce: 1});
        address firstNotary = domains[DOMAIN_LOCAL].agents[0];
        (bytes memory firstSnapPayload, bytes memory firstNotarySignature) = signSnapshot(firstNotary, firstSnap);
        (, bytes memory firstGuardSignature) = signSnapshot(domains[0].agent, firstSnap);
        inbox.submitSnapshot(firstSnapPayload, firstGuardSignature);
        inbox.submitSnapshot(firstSnapPayload, firstNotarySignature);
        uint256 firstSkipTime = random.nextUint32();
        skip(firstSkipTime);
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(domain);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    function prepareGasDataDisputeTest() internal returns (GasData remoteGasData, GasData localGasData) {
        address notary = domains[DOMAIN_LOCAL].agent;
        address reportGuard = domains[0].agent;
        address snapshotGuard = domains[0].agents[1];

        Random memory random = Random("salt");
        RawSnapshot memory rawSnap = RawSnapshot(new RawState[](2));
        rawSnap.states[0] = random.nextState({origin: DOMAIN_REMOTE, nonce: 1});
        rawSnap.states[1] = random.nextState({origin: DOMAIN_LOCAL, nonce: 2});
        remoteGasData = rawSnap.states[0].castToState().gasData();
        localGasData = rawSnap.states[1].castToState().gasData();

        // Another Guard signs the snapshot
        (bytes memory snapPayload, bytes memory guardSignature) = signSnapshot(snapshotGuard, rawSnap);
        inbox.submitSnapshot(snapPayload, guardSignature);
        // Notary signs the snapshot
        (, bytes memory notarySignature) = signSnapshot(notary, rawSnap);
        inbox.submitSnapshot(snapPayload, notarySignature);
        // Sanity checks
        {
            (GasData gasData,) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
            assert(GasData.unwrap(gasData) == GasData.unwrap(remoteGasData));
        }
        {
            (GasData gasData,) = InterfaceDestination(localDestination()).getGasData(DOMAIN_LOCAL);
            assert(GasData.unwrap(gasData) == GasData.unwrap(localGasData));
        }
        openTestDispute({guardIndex: agentIndex[reportGuard], notaryIndex: agentIndex[notary]});
    }

    function prepareNotaryWonDisputeTest() internal {
        address notary = domains[DOMAIN_LOCAL].agent;
        address guard = domains[0].agent;
        resolveTestDispute({slashedIndex: agentIndex[guard], rivalIndex: agentIndex[notary]});
    }

    function test_getGasData_notaryInDispute() public {
        prepareGasDataDisputeTest();
        skip(7 days);
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_LOCAL);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    function test_getGasData_notaryWonDisputeTimeout() public {
        prepareGasDataDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY - 1);
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_LOCAL);
        assertEq(GasData.unwrap(gasData), 0);
        assertEq(dataMaturity, 0);
    }

    function test_getGasData_afterNotaryDisputeTimeout() public {
        (GasData remoteGasData, GasData localGasData) = prepareGasDataDisputeTest();
        skip(7 days);
        prepareNotaryWonDisputeTest();
        skip(DISPUTE_TIMEOUT_NOTARY);
        // Check getGasData
        (GasData gasData, uint256 dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_REMOTE);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(remoteGasData));
        assertEq(dataMaturity, 7 days + DISPUTE_TIMEOUT_NOTARY);
        (gasData, dataMaturity) = InterfaceDestination(localDestination()).getGasData(DOMAIN_LOCAL);
        assertEq(GasData.unwrap(gasData), GasData.unwrap(localGasData));
        assertEq(dataMaturity, 7 days + DISPUTE_TIMEOUT_NOTARY);
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
        inbox.submitSnapshot(snapPayload, guardSignature);
        inbox.submitSnapshot(snapPayload, notarySignature);
    }

    /// @notice Returns local domain for the tested contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
