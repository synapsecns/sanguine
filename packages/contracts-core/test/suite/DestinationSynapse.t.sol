// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {RawAttestation} from "../utils/libs/SynapseStructs.t.sol";
import {AgentFlag, AgentStatus, SynapseTest} from "../utils/SynapseTest.t.sol";
import {ExecutionHubTest} from "./hubs/ExecutionHub.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable no-empty-blocks
// solhint-disable ordering
contract DestinationSynapseTest is ExecutionHubTest {
    // Deploy Production version of Destination and Summit and mocks for everything else
    constructor() SynapseTest(DEPLOY_PROD_DESTINATION_SYNAPSE | DEPLOY_PROD_SUMMIT) {}

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

    /// @notice Returns local domain for the tested system contract
    function localDomain() public pure override returns (uint32) {
        return DOMAIN_SYNAPSE;
    }
}
