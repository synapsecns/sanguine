// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../tools/AttestationCollectorTools.t.sol";

// solhint-disable func-name-mixedcase
contract AttestationCollectorTest is AttestationCollectorTools {
    uint256 internal savedSignatures;
    mapping(uint64 => mapping(address => uint32)) internal latestAgentNonce;
    mapping(uint64 => mapping(address => uint256)) internal savedAttestations;

    function setUp() public override {
        super.setUp();
        setupCollector();
    }

    function test_setup() public {
        assertEq(collector.owner(), owner, "!owner");
        for (uint256 i = 0; i < GUARDS; ++i) {
            address guard = suiteGuard(i);
            assertTrue(collector.isGuard(guard), "!isGuard");
        }
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < NOTARIES_PER_CHAIN; ++i) {
                address notary = suiteNotary(domain, i);
                assertTrue(collector.isNotary(domain, notary), "!isNotary");
            }
        }
    }

    function test_initialize() public {
        collector = new AttestationCollectorHarness();
        assertEq(collector.owner(), address(0), "!owner: pre init");
        collector.initialize();
        assertEq(collector.owner(), address(this), "!owner: post init");
    }

    function test_initialize_revert_onlyOnce() public {
        expectRevertAlreadyInitialized();
        collector.initialize();
    }

    function test_addRemoveAgent(uint32 domain, address agent) public {
        collector = new AttestationCollectorHarness();
        collector.initialize();
        collector.addAgent(domain, agent);
        assertTrue(collector.isActiveAgent(domain, agent), "!added");
        collector.removeAgent(domain, agent);
        assertFalse(collector.isActiveAgent(domain, agent), "!removed");
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                      TESTS: SUBMIT ATTESTATION                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line code-complexity
    function test_submitAttestation(uint256 guardSigs, uint256 notarySigs) public {
        guardSigs = guardSigs % GUARDS;
        notarySigs = notarySigs % NOTARIES_PER_CHAIN;
        (address[] memory guardSigners, address[] memory notarySigners) = _createSigners({
            destination: DOMAIN_REMOTE,
            guardSigs: guardSigs,
            notarySigs: notarySigs
        });
        saveAttestationMetadata();
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            guardSigners: guardSigners,
            notarySigners: notarySigners
        });
        bool expectedStored = false;
        // Expect events re: guard attestation
        for (uint256 i = 0; i < guardSigs; ++i) {
            address guard = attestationGuards[i];
            if (_isFreshAttestation(guard, attestationDomains, ra.nonce)) {
                expectAttestationSaved({
                    signatureIndex: savedSignatures,
                    isGuard: true,
                    agentIndex: i
                });
                ++savedSignatures;
                ++savedAttestations[attestationDomains][guard];
                expectedStored = true;
                latestAgentNonce[attestationDomains][guard] = ra.nonce;
            }
        }
        // Expect events re: notary attestation
        for (uint256 i = 0; i < notarySigs; ++i) {
            address notary = attestationNotaries[i];
            if (_isFreshAttestation(notary, attestationDomains, ra.nonce)) {
                expectAttestationSaved({
                    signatureIndex: savedSignatures,
                    isGuard: false,
                    agentIndex: i
                });
                ++savedSignatures;
                ++savedAttestations[attestationDomains][notary];
                expectedStored = true;
                latestAgentNonce[attestationDomains][notary] = ra.nonce;
            }
        }
        if (expectedStored) {
            expectAttestationAccepted();
        }
        if (guardSigs + notarySigs == 0) {
            vm.expectRevert("Not an attestation");
        }
        bool stored = collector.submitAttestation(attestationRaw);
        assertEq(stored, expectedStored, "!returnValue");
    }

    function test_submitAttestation_sameSigners_rejectsConflicts() public {
        // Submit a guard-only attestation
        test_submitAttestation({ guardSigs: 1, notarySigs: 0 });
        // Create conflict attestation signed by the same guard (no notary sig)
        createAttestation({
            origin: ra.origin,
            destination: ra.destination,
            nonce: ra.nonce,
            root: "Conflict root",
            guardIndex: 0,
            notaryIndex: NOTARIES_PER_CHAIN
        });
        vm.recordLogs();
        assertFalse(collector.submitAttestation(attestationRaw), "Saved attestation with conflict");
        assertEq(vm.getRecordedLogs().length, 0, "Emitted logs");
    }

    function test_submitAttestation_sameSigners_rejectsSameNonce(
        uint256 guardSigs,
        uint256 notarySigs
    ) public {
        guardSigs = guardSigs % GUARDS;
        notarySigs = notarySigs % NOTARIES_PER_CHAIN;
        vm.assume(guardSigs + notarySigs != 0);
        // Submit attestation with the given agents
        test_submitAttestation(guardSigs, notarySigs);
        vm.recordLogs();
        // Resubmit attestation with the same agents and nonce
        assertFalse(
            collector.submitAttestation(attestationRaw),
            "Saved attestation with same nonce"
        );
        assertEq(vm.getRecordedLogs().length, 0, "Emitted logs");
    }

    function test_submitAttestation_sameSigners_rejectsOlderNonce(
        uint256 guardSigs,
        uint256 notarySigs
    ) public {
        guardSigs = guardSigs % GUARDS;
        notarySigs = notarySigs % NOTARIES_PER_CHAIN;
        vm.assume(guardSigs + notarySigs != 0);
        // Submit attestation with the given agents
        test_submitAttestation(guardSigs, notarySigs);
        // Create attestation with the same agents and lower nonce
        mockNonce = ra.nonce - 1;
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            guardSigners: attestationGuards,
            notarySigners: attestationNotaries
        });
        vm.recordLogs();
        assertFalse(
            collector.submitAttestation(attestationRaw),
            "Saved attestation with older nonce"
        );
        assertEq(vm.getRecordedLogs().length, 0, "Emitted logs");
    }

    function test_submitAttestation_otherSigners_rejectsConflicts() public {
        // Submit a guard-only attestation
        test_submitAttestation({ guardSigs: 1, notarySigs: 0 });
        // Create conflict attestation signed by another guard (no notary sig)
        createAttestation({
            origin: ra.origin,
            destination: ra.destination,
            nonce: ra.nonce,
            root: "Conflict root",
            guardIndex: 1,
            notaryIndex: NOTARIES_PER_CHAIN
        });
        vm.recordLogs();
        assertFalse(collector.submitAttestation(attestationRaw), "Saved attestation with conflict");
        assertEq(vm.getRecordedLogs().length, 0, "Emitted logs");
    }

    function test_submitAttestation_otherSigners_acceptsSameNonce() public {
        // Submit a guard-only attestation
        test_submitAttestation({ guardSigs: 1, notarySigs: 0 });
        // Create conflict attestation signed by the same guard and a notary
        createAttestation({
            origin: ra.origin,
            destination: ra.destination,
            nonce: ra.nonce,
            root: ra.root,
            guardIndex: 0,
            notaryIndex: 0
        });
        expectAttestationSaved({ signatureIndex: 1, isGuard: false, agentIndex: 0 });
        expectAttestationAccepted();
        assertTrue(
            collector.submitAttestation(attestationRaw),
            "Did not save another agent attestation: same nonce"
        );
    }

    function test_submitAttestation_otherSigners_acceptsOlderNonce() public {
        // Submit a guard-only attestation
        test_submitAttestation({ guardSigs: 1, notarySigs: 0 });
        // Create conflict attestation signed by the same guard and a notary
        createAttestation({
            origin: ra.origin,
            destination: ra.destination,
            nonce: ra.nonce - 1,
            root: "Another root",
            guardIndex: 0,
            notaryIndex: 0
        });
        expectAttestationSaved({ signatureIndex: 1, isGuard: false, agentIndex: 0 });
        expectAttestationAccepted();
        assertTrue(
            collector.submitAttestation(attestationRaw),
            "Did not save another agent attestation: older nonce"
        );
    }

    function test_submitAttestation_revert_zeroRoot() public {
        createAttestation({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            nonce: 1,
            root: bytes32(0)
        });
        vm.expectRevert("Root is zero");
        collector.submitAttestation(attestationRaw);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            TESTS: GETTERS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_agentAttestations_savedAttestations() public {
        // Step1: Submit an attestation: guard0, notary0
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        _checkTotalGetters();
        // Step2: Submit an attestation: notary0, notary1
        test_submitAttestation({ guardSigs: 0, notarySigs: 2 });
        _checkTotalGetters();
        // Step3: Submit an attestation: guard0, guard1, guard2
        test_submitAttestation({ guardSigs: 3, notarySigs: 0 });
        _checkTotalGetters();
        // Step4: Submit an attestation: all agents
        test_submitAttestation({ guardSigs: GUARDS, notarySigs: NOTARIES_PER_CHAIN });
        _checkTotalGetters();
    }

    function test_getAgentAttestation_getSavedAttestation() public {
        address guard0 = suiteGuard(0);
        address guard1 = suiteGuard(1);
        address notary0 = suiteNotary(DOMAIN_REMOTE, 0);
        address notary1 = suiteNotary(DOMAIN_REMOTE, 1);
        // Step1: Submit an attestation: guard0, notary0
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        // Reconstruct guard-only attestation
        bytes memory guard00 = createSameAttestation({ isGuard: true, agent: guard0 });
        // Reconstruct notary-only attestation
        bytes memory notary00 = createSameAttestation({ isGuard: false, agent: notary0 });
        // Step2: Submit an attestation: guard0, guard1
        test_submitAttestation({ guardSigs: 2, notarySigs: 0 });
        // Reconstruct guard-only attestations
        bytes memory guard01 = createSameAttestation({ isGuard: true, agent: guard0 });
        bytes memory guard10 = createSameAttestation({ isGuard: true, agent: guard1 });
        // Step3: Submit an attestation: guard0, notary0, notary1
        test_submitAttestation({ guardSigs: 1, notarySigs: 2 });
        // Reconstruct guard-only attestation
        bytes memory guard02 = createSameAttestation({ isGuard: true, agent: guard0 });
        // Reconstruct notary-only attestations
        bytes memory notary01 = createSameAttestation({ isGuard: false, agent: notary0 });
        bytes memory notary10 = createSameAttestation({ isGuard: false, agent: notary1 });
        // Check: guard0
        assertEq(
            collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard0, 0),
            guard00,
            "!guard00"
        );
        assertEq(
            collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard0, 1),
            guard01,
            "!guard01"
        );
        assertEq(
            collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard0, 2),
            guard02,
            "!guard02"
        );
        // Check: guard1
        assertEq(
            collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard1, 0),
            guard10,
            "!guard10"
        );
        // Check: notary0
        assertEq(
            collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, notary0, 0),
            notary00,
            "!notary00"
        );
        assertEq(
            collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, notary0, 1),
            notary01,
            "!notary01"
        );
        // Check: notary1
        assertEq(
            collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, notary1, 0),
            notary10,
            "!notary10"
        );
        // Check: step1 (guard0, notary0)
        assertEq(collector.getSavedAttestation(0), guard00, "!getSavedAttestation(0)");
        assertEq(collector.getSavedAttestation(1), notary00, "!getSavedAttestation(1)");
        // Check: step2 (guard0, guard1)
        assertEq(collector.getSavedAttestation(2), guard01, "!getSavedAttestation(2)");
        assertEq(collector.getSavedAttestation(3), guard10, "!getSavedAttestation(3)");
        // Check: step3 (guard0, notary0, notary1)
        assertEq(collector.getSavedAttestation(4), guard02, "!getSavedAttestation(4)");
        assertEq(collector.getSavedAttestation(5), notary01, "!getSavedAttestation(5)");
        assertEq(collector.getSavedAttestation(6), notary10, "!getSavedAttestation(6)");
    }

    function test_getAgentAttestation_getSavedAttestation_revert_outOfRange() public {
        // Step1: Submit an attestation: guard0, notary0
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        // Step2: Submit an attestation: guard0, guard1
        test_submitAttestation({ guardSigs: 2, notarySigs: 0 });
        // Step3: Submit an attestation: guard0, notary0, notary1
        test_submitAttestation({ guardSigs: 1, notarySigs: 2 });
        address guard0 = suiteGuard(0);
        address guard1 = suiteGuard(1);
        address guard2 = suiteGuard(2);
        // Attestations for guard0: 3
        vm.expectRevert("Out of range");
        collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard0, 3);
        // Attestations for guard1: 1
        vm.expectRevert("Out of range");
        collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard1, 1);
        // Attestations for guard2: 0
        vm.expectRevert("Out of range");
        collector.getAgentAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard2, 0);
        // Saved attestations: 7
        vm.expectRevert("Out of range");
        collector.getSavedAttestation(7);
    }

    function test_getAttestation_guardThenNotary() public {
        address guard1 = suiteGuard(1);
        // Step1: Submit an attestation: guard0
        test_submitAttestation({ guardSigs: 1, notarySigs: 0 });
        bytes memory expectedAtt = attestationRaw;
        assertEq(
            collector.getAttestation({
                _origin: ra.origin,
                _destination: ra.destination,
                _nonce: ra.nonce
            }),
            expectedAtt,
            "!step1"
        );
        // Step2: Submit the same data signed by guard1
        createSameAttestation({ isGuard: true, agent: guard1 });
        assertTrue(collector.submitAttestation(attestationRaw), "!step2: guard1");
        // Should return the same guard0 attestation
        assertEq(
            collector.getAttestation({
                _origin: ra.origin,
                _destination: ra.destination,
                _nonce: ra.nonce
            }),
            expectedAtt,
            "!step2"
        );
        // Step3: Submit the same data signed by notary0, notary1
        createSameAttestation({ guardSigs: 0, notarySigs: 2 });
        assertTrue(collector.submitAttestation(attestationRaw), "!step3: notary0+notary1");
        // Expect to receive attestation signed by guard0 and notary0
        expectedAtt = createSameAttestation({ guardSigs: 1, notarySigs: 1 });
        assertEq(
            collector.getAttestation({
                _origin: ra.origin,
                _destination: ra.destination,
                _nonce: ra.nonce
            }),
            expectedAtt,
            "!step3"
        );
    }

    function test_getAttestation_notaryThenGuard() public {
        // Step1: Submit an attestation: notary0, notary1
        test_submitAttestation({ guardSigs: 0, notarySigs: 2 });
        bytes memory expectedAtt = createSameAttestation({ guardSigs: 0, notarySigs: 1 });
        assertEq(
            collector.getAttestation({
                _origin: ra.origin,
                _destination: ra.destination,
                _nonce: ra.nonce
            }),
            expectedAtt,
            "!step1"
        );
        // Step2: submit an attestation: all notaries
        createSameAttestation({ guardSigs: 0, notarySigs: 4 });
        assertTrue(collector.submitAttestation(attestationRaw), "!step2: all notaries");
        // Should return the same notary0 attestation
        assertEq(
            collector.getAttestation({
                _origin: ra.origin,
                _destination: ra.destination,
                _nonce: ra.nonce
            }),
            expectedAtt,
            "!step2"
        );
        // Step3: submit an attestation: guard0
        createSameAttestation({ guardSigs: 1, notarySigs: 0 });
        assertTrue(collector.submitAttestation(attestationRaw), "!step3: guard0");
        // Expect to receive attestation signed by guard0 and notary0
        expectedAtt = createSameAttestation({ guardSigs: 1, notarySigs: 1 });
        assertEq(
            collector.getAttestation({
                _origin: ra.origin,
                _destination: ra.destination,
                _nonce: ra.nonce
            }),
            expectedAtt,
            "!step3"
        );
    }

    function test_getAttestation_revert_unknownNonce() public {
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        vm.expectRevert("Unknown nonce");
        collector.getAttestation(ra.origin, ra.destination, ra.nonce - 1);
        vm.expectRevert("Unknown nonce");
        collector.getAttestation(ra.origin, ra.destination, ra.nonce + 1);
    }

    function test_getRoot() public {
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        uint32 nonce0 = ra.nonce;
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        uint32 nonce1 = ra.nonce;
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        uint32 nonce2 = ra.nonce;
        assertEq(
            collector.getRoot(ra.origin, ra.destination, nonce0),
            _createMockRoot(nonce0, 0),
            "!nonce0"
        );
        assertEq(
            collector.getRoot(ra.origin, ra.destination, nonce1),
            _createMockRoot(nonce1, 0),
            "!nonce1"
        );
        assertEq(
            collector.getRoot(ra.origin, ra.destination, nonce2),
            _createMockRoot(nonce2, 0),
            "!nonce2"
        );
    }

    function test_getLatest() public {
        address guard0 = suiteGuard(0);
        address guard1 = suiteGuard(1);
        address notary0 = suiteNotary(DOMAIN_REMOTE, 0);
        address notary1 = suiteNotary(DOMAIN_REMOTE, 1);
        // Step1: guard0, notary0
        test_submitAttestation({ guardSigs: 1, notarySigs: 1 });
        uint32 nonce0 = ra.nonce;
        bytes memory notaryLatest0 = createSameAttestation({ isGuard: false, agent: notary0 });
        // Step1: guard0, guard1
        test_submitAttestation({ guardSigs: 2, notarySigs: 0 });
        uint32 nonce1 = ra.nonce;
        bytes memory guardLatest0 = createSameAttestation({ isGuard: true, agent: guard0 });
        bytes memory guardLatest1 = createSameAttestation({ isGuard: true, agent: guard1 });
        // Check guard0
        uint32 latestNonce = collector.getLatestNonce(DOMAIN_LOCAL, DOMAIN_REMOTE, guard0);
        assertEq(latestNonce, nonce1, "!guard0: nonce");
        assertEq(
            collector.getLatestAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard0),
            guardLatest0,
            "!guard0: attestation"
        );
        // Check guard1
        latestNonce = collector.getLatestNonce(DOMAIN_LOCAL, DOMAIN_REMOTE, guard1);
        assertEq(latestNonce, nonce1, "!guard1: nonce");
        assertEq(
            collector.getLatestAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, guard1),
            guardLatest1,
            "!guard1: attestation"
        );
        // Check notary0
        latestNonce = collector.getLatestNonce(DOMAIN_LOCAL, DOMAIN_REMOTE, notary0);
        assertEq(latestNonce, nonce0, "!notary0: nonce");
        assertEq(
            collector.getLatestAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, notary0),
            notaryLatest0,
            "!notary0: attestation"
        );
        // Check notary1
        latestNonce = collector.getLatestNonce(DOMAIN_LOCAL, DOMAIN_REMOTE, notary1);
        assertEq(latestNonce, 0, "!notary1: nonce");
        vm.expectRevert("No attestations found");
        collector.getLatestAttestation(DOMAIN_LOCAL, DOMAIN_REMOTE, notary1);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // solhint-disable-next-line code-complexity
    function _checkTotalGetters() internal {
        for (uint32 o = 0; o < DOMAINS; ++o) {
            uint32 _origin = domains[o];
            for (uint32 d = 0; d < DOMAINS; ++d) {
                uint32 _destination = domains[d];
                uint64 attDomains = Attestation.attestationDomains(_origin, _destination);
                for (uint256 i = 0; i < GUARDS; ++i) {
                    address guard = suiteGuard(i);
                    assertEq(
                        collector.agentAttestations(_origin, _destination, guard),
                        savedAttestations[attDomains][guard],
                        "!agentAttestations: guard"
                    );
                }
                for (uint256 n = 0; n < DOMAINS; ++n) {
                    for (uint256 i = 0; i < NOTARIES_PER_CHAIN; ++i) {
                        address notary = suiteNotary(domains[n], i);
                        assertEq(
                            collector.agentAttestations(_origin, _destination, notary),
                            savedAttestations[attDomains][notary],
                            "!agentAttestations: notary"
                        );
                    }
                }
            }
        }
        assertEq(collector.savedAttestations(), savedSignatures, "!savedAttestations");
    }

    function _isFreshAttestation(
        address agent,
        uint64 attDomains,
        uint32 nonce
    ) internal view returns (bool) {
        return latestAgentNonce[attDomains][agent] < nonce;
    }
}
