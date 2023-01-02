// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../tools/OriginTools.t.sol";

// solhint-disable func-name-mixedcase
contract OriginTest is OriginTools {
    using TypedMemView for bytes;
    using TypedMemView for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                   TESTS: CONSTRUCTOR & INITIALIZER                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_initialize() public {
        OriginHarness origin = new OriginHarness(DOMAIN_LOCAL);
        vm.prank(owner);
        origin.initialize();
        assertEq(origin.owner(), owner, "!owner");
        (bytes32 histRoot, uint40 blockNumber, uint40 timestamp) = origin.getHistoricalRoot(0, 0);
        assertEq(histRoot, origin.root(0), "!historicalRoots(0)");
        assertEq(blockNumber, 0, "!blockNumber(0)");
        assertEq(timestamp, 0, "!timestamp(0)");
    }

    // solhint-disable-next-line code-complexity
    function test_initializedCorrectly() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 localDomain = domains[d];
            OriginHarness origin = suiteOrigin(localDomain);
            // Check local domain
            assertEq(origin.localDomain(), localDomain, "!localDomain");
            // Check owner
            assertEq(origin.owner(), owner, "!owner");
            // Check contract addresses
            assertEq(
                address(origin.systemRouter()),
                address(suiteSystemRouter(localDomain)),
                "!systemRouter"
            );
            // Check all notaries
            for (uint256 dest = 0; dest < DOMAINS; ++dest) {
                uint32 domain = domains[dest];
                if (domain == localDomain) {
                    // Origin should not keep track of local Notaries
                    assertEq(origin.amountAgents(domain), 0, "!notariesAmount: local domain");
                    for (uint256 i = 0; i < NOTARIES_PER_CHAIN; ++i) {
                        assertFalse(
                            origin.isActiveAgent(domain, suiteNotary(domain, i)),
                            string.concat("!notary", getActorSuffix(i), ": local domain")
                        );
                    }
                } else {
                    // Origin should keep track of all remote notaries
                    assertEq(
                        origin.amountAgents(domain),
                        NOTARIES_PER_CHAIN,
                        "!notariesAmount: remote domain"
                    );
                    for (uint256 i = 0; i < NOTARIES_PER_CHAIN; ++i) {
                        assertTrue(
                            origin.isActiveAgent(domain, suiteNotary(domain, i)),
                            string.concat("!notary", getActorSuffix(i), ": remote domain")
                        );
                    }
                }
            }
            // Check global guards
            assertEq(origin.amountAgents({ _domain: 0 }), GUARDS, "!guardsAmount");
            for (uint256 i = 0; i < GUARDS; ++i) {
                assertTrue(
                    origin.isActiveAgent({ _domain: 0, _account: suiteGuard(i) }),
                    string.concat("!guard", getActorSuffix(i))
                );
            }
            // Root of an empty sparse Merkle tree should be stored with nonce=0
            (bytes32 histRoot, uint40 blockNumber, uint40 timestamp) = origin.getHistoricalRoot(
                0,
                0
            );
            assertEq(histRoot, origin.root(0), "!historicalRoots(0)");
            assertEq(blockNumber, 0, "!blockNumber(0)");
            assertEq(timestamp, 0, "!timestamp(0)");
        }
    }

    function test_initialize_revert_onlyOnce() public {
        expectRevertAlreadyInitialized();
        suiteOrigin(DOMAIN_LOCAL).initialize();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TESTS: DISPATCHING MESSAGES                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_dispatch() public {
        skipBlock();
        createDispatchedMessage({ context: userLocalToRemote, mockTips: true });
        expectDispatch();
        originDispatch();
    }

    function test_dispatch_noTips() public {
        // User should be able to send a message w/o any tips
        createEmptyTips();
        skipBlock();
        createDispatchedMessage({ context: userLocalToRemote, mockTips: false });
        expectDispatch();
        originDispatch();
    }

    function test_dispatch_revert_tipsTooSmall() public {
        createDispatchedMessage({ context: userLocalToRemote, mockTips: true });
        --tipsTotal; // force user to specify msg.value 1 wei less than needed
        originDispatch({ revertMessage: "!tips: totalTips" });
    }

    function test_dispatch_revert_tipsTooBig() public {
        createDispatchedMessage({ context: userLocalToRemote, mockTips: true });
        ++tipsTotal; // force user to specify msg.value 1 wei more than needed
        originDispatch({ revertMessage: "!tips: totalTips" });
    }

    function test_dispatch_revert_tipsVersionIncorrect() public {
        createEmptyTips();
        uint256 length = tipsRaw.length;
        // COnstruct tips payload, but use incorrect tips version
        tipsRaw = abi.encodePacked(
            Tips.TIPS_VERSION + 1,
            tipNotary,
            tipBroadcaster,
            tipProver,
            tipExecutor
        );
        // Length should stay the same
        require(tipsRaw.length == length, "Failed to construct tips payload");
        createDispatchedMessage({ context: userLocalToRemote, mockTips: false });
        originDispatch({ revertMessage: "!tips: formatting" });
    }

    function test_dispatch_revert_tipsPayloadTooSmall() public {
        createEmptyTips();
        // Cut the last byte from tips payload, making it improperly formatted
        tipsRaw = tipsRaw.ref(0).slice({ _index: 0, _len: tipsRaw.length - 1, newType: 0 }).clone();
        createDispatchedMessage({ context: userLocalToRemote, mockTips: false });
        originDispatch({ revertMessage: "!tips: formatting" });
    }

    function test_dispatch_revert_tipsPayloadTooBig() public {
        createEmptyTips();
        // Add extra byte to tips payload, making it improperly formatted
        tipsRaw = bytes.concat(tipsRaw, bytes1(0));
        createDispatchedMessage({ context: userLocalToRemote, mockTips: false });
        originDispatch({ revertMessage: "!tips: formatting" });
    }

    function test_dispatch_revert_messageTooBig() public {
        // Messages over 2 KiB are rejected
        createDispatchedMessage({
            context: userLocalToRemote,
            mockTips: true,
            body: new bytes(2 * 2**10 + 1),
            recipient: MOCK_RECIPIENT,
            optimisticSeconds: MOCK_OPTIMISTIC_SECONDS
        });
        originDispatch({ revertMessage: "msg too long" });
    }

    function test_suggestAttestation() public {
        OriginHarness origin = suiteOrigin(DOMAIN_LOCAL);
        uint256 amount = 5;
        // Send a few messages
        for (uint256 i = 0; i < amount; ++i) {
            test_dispatch();
        }
        bytes memory data = origin.suggestAttestation(DOMAIN_REMOTE);
        // Should match latest values
        assertEq(
            data,
            Attestation.formatAttestationData({
                _origin: DOMAIN_LOCAL,
                _destination: DOMAIN_REMOTE,
                _nonce: uint32(amount),
                _root: origin.root(DOMAIN_REMOTE),
                _blockNumber: uint40(block.number),
                _timestamp: uint40(block.timestamp)
            })
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                 TESTS: SUBMIT ATTESTATION (REVERTS)                  ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitAttestation_revert_wrongDomain() public {
        // Add local Notary: Origin is not supposed to track them
        suiteOrigin(DOMAIN_LOCAL).addLocalNotary(suiteNotary(DOMAIN_LOCAL));
        _createAttestation_revert_wrongDomain();
        originSubmitAttestation({
            domain: DOMAIN_LOCAL,
            revertMessage: "!attestationOrigin: !local"
        });
    }

    function test_submitAttestation_revert_notNotary() public {
        _createAttestation_revert_notNotary();
        originSubmitAttestation({
            domain: DOMAIN_LOCAL,
            revertMessage: "Signer is not authorized"
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                  TESTS: SUBMIT ATTESTATION (VALID)                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitAttestation_valid_suggested() public {
        _createAttestation_valid_suggested();
        _testSubmitAttestation({ domain: DOMAIN_LOCAL, isValidAttestation: true });
    }

    function test_submitAttestation_valid_outdated() public {
        _createAttestation_valid_outdated();
        _testSubmitAttestation({ domain: DOMAIN_LOCAL, isValidAttestation: true });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                  TESTS: SUBMIT ATTESTATION (FRAUD)                   ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function test_submitAttestation_fraud_nonExistingNonce() public {
        _createAttestation_fraud_nonExistingNonce();
        _testSubmitAttestation({ domain: DOMAIN_LOCAL, isValidAttestation: false });
    }

    function test_submitAttestation_fraud_fakeNonce() public {
        _createAttestation_fraud_fakeNonce();
        _testSubmitAttestation({ domain: DOMAIN_LOCAL, isValidAttestation: false });
    }

    function test_submitAttestation_fraud_fakeRoot() public {
        _createAttestation_fraud_fakeRoot();
        _testSubmitAttestation({ domain: DOMAIN_LOCAL, isValidAttestation: false });
    }

    // TODO(Chi): enable Reports tests once reimplemented

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                    TESTS: SUBMIT REPORT (REVERTS)                    ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // function test_submitReport_revert_wrongDomain() public {
    //     // Add local Notary: Origin is not supposed to track them
    //     suiteOrigin(DOMAIN_LOCAL).addLocalNotary(suiteNotary(DOMAIN_LOCAL));
    //     _createAttestation_revert_wrongDomain();
    //     createReport(Report.Flag.Fraud);
    //     originSubmitReport({ domain: DOMAIN_LOCAL, revertMessage: "!attestationOrigin: !local" });
    // }

    // function test_submitReport_revert_notNotary() public {
    //     _createAttestation_revert_notNotary();
    //     createReport(Report.Flag.Fraud);
    //     originSubmitReport({ domain: DOMAIN_LOCAL, revertMessage: "Signer is not authorized" });
    // }

    // function test_submitReport_revert_notGuard() public {
    //     _createAttestation_valid_suggested();
    //     createReport({ flag: Report.Flag.Fraud, signer: attacker });
    //     originSubmitReport({ domain: DOMAIN_LOCAL, revertMessage: "Signer is not authorized" });
    // }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                TESTS: SUBMIT REPORT (VALID, CORRECT)                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    // In these tests Guard signs a Flag.Valid Report on a Valid attestation
    // No one is getting slashed

    // function test_submitReport_valid_correct_suggested() public {
    //     _createAttestation_valid_suggested();
    //     createReport(Report.Flag.Valid); // Correct report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: true,
    //         isCorrectReport: true
    //     });
    // }

    // function test_submitReport_valid_correct_outdated() public {
    //     _createAttestation_valid_outdated();
    //     createReport(Report.Flag.Valid); // Correct report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: true,
    //         isCorrectReport: true
    //     });
    // }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║               TESTS: SUBMIT REPORT (FRAUD, INCORRECT)                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    // In these tests Guard signs a Flag.Fraud Report on a Valid attestation
    // Guard is slashed as a result

    // function test_submitReport_fraud_incorrect_suggested() public {
    //     _createAttestation_valid_suggested();
    //     createReport(Report.Flag.Fraud); // Incorrect report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: true,
    //         isCorrectReport: false
    //     });
    // }

    // function test_submitReport_fraud_incorrect_outdated() public {
    //     _createAttestation_valid_outdated();
    //     createReport(Report.Flag.Fraud); // Incorrect report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: true,
    //         isCorrectReport: false
    //     });
    // }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                TESTS: SUBMIT REPORT (FRAUD, CORRECT)                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    // In these tests Guard signs a Flag.Fraud Report on a Fraud attestation
    // Notary is slashed as a result, Guard gets a reward

    // function submitReport_fraud_correct_nonExistingNonce() public {
    //     _createAttestation_fraud_nonExistingNonce();
    //     createReport(Report.Flag.Fraud); // Correct report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: false,
    //         isCorrectReport: true
    //     });
    // }

    // function submitReport_fraud_correct_fakeNonce() public {
    //     _createAttestation_fraud_fakeNonce();
    //     createReport(Report.Flag.Fraud); // Correct report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: false,
    //         isCorrectReport: true
    //     });
    // }

    // function test_submitReport_fraud_correct_fakeRoot() public {
    //     _createAttestation_fraud_fakeRoot();
    //     createReport(Report.Flag.Fraud); // Correct report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: false,
    //         isCorrectReport: true
    //     });
    // }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║               TESTS: SUBMIT REPORT (VALID, INCORRECT)                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/
    // In these tests Guard signs a Flag.Valid Report on a Fraud attestation
    // Notary is slashed as a result, Guard does NOT get a reward
    // Guard is slashed as a result

    // function test_submitReport_fraud_incorrect_nonExistingNonce() public {
    //     _createAttestation_fraud_nonExistingNonce();
    //     createReport(Report.Flag.Valid); // Incorrect report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: false,
    //         isCorrectReport: false
    //     });
    // }

    // function test_submitReport_fraud_incorrect_fakeNonce() public {
    //     _createAttestation_fraud_fakeNonce();
    //     createReport(Report.Flag.Valid); // Incorrect report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: false,
    //         isCorrectReport: false
    //     });
    // }

    // function test_submitReport_fraud_incorrect_fakeRoot() public {
    //     _createAttestation_fraud_fakeRoot();
    //     createReport(Report.Flag.Valid); // Incorrect report
    //     _testSubmitReport({
    //         domain: DOMAIN_LOCAL,
    //         isValidAttestation: false,
    //         isCorrectReport: false
    //     });
    // }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            TESTS: HALTING                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO: enable Guards check once Go tests are updated
    // function test_halts_noGuards() public {
    //     createDispatchedMessage({ context: userLocalToRemote, mockTips: true });
    //     OriginHarness origin = suiteOrigin(DOMAIN_LOCAL);
    //     origin.removeAllAgents(0);
    //     originDispatch({ revertMessage: "No active guards" });
    // }

    function test_halts_noNotaries() public {
        createDispatchedMessage({ context: userLocalToRemote, mockTips: true });
        OriginHarness origin = suiteOrigin(DOMAIN_LOCAL);
        origin.removeAllAgents(DOMAIN_REMOTE);
        originDispatch({ revertMessage: "No active notaries" });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Create an attestation referring to another domain
    function _createAttestation_revert_wrongDomain() internal {
        test_dispatch();
        // DOMAIN_LOCAL is used for Origin testing
        createAttestationMock({ origin: DOMAIN_REMOTE, destination: DOMAIN_LOCAL });
    }

    // Create an attestation signed by not a Guard
    function _createAttestation_revert_notGuard() internal {
        test_dispatch();
        // index 0 refers to first Guard
        (address[] memory guardSigners, address[] memory notarySigners) = _createSigners({
            destination: DOMAIN_REMOTE,
            guardSigs: 1,
            notarySigs: 1,
            attackerIndex: 0
        });
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            guardSigners: guardSigners,
            notarySigners: notarySigners
        });
    }

    // Create an attestation signed by not a Notary
    function _createAttestation_revert_notNotary() internal {
        test_dispatch();
        // index 1 refers to first Guard
        (address[] memory guardSigners, address[] memory notarySigners) = _createSigners({
            destination: DOMAIN_REMOTE,
            guardSigs: 1,
            notarySigs: 1,
            attackerIndex: 1
        });
        createAttestationMock({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            guardSigners: guardSigners,
            notarySigners: notarySigners
        });
    }

    // Create a valid attestation referring to the current Origin state
    function _createAttestation_valid_suggested() internal {
        test_dispatch();
        createSuggestedAttestation({ origin: DOMAIN_LOCAL, destination: DOMAIN_REMOTE });
        // Suggested attestation is valid
    }

    // Create a valid attestation referring to the past Origin state
    function _createAttestation_valid_outdated() internal {
        test_dispatch();
        createSuggestedAttestation({ origin: DOMAIN_LOCAL, destination: DOMAIN_REMOTE });
        // Dispatch a message to make the attestation older than new suggested one
        test_dispatch();
        // Outdated attestation is valid
    }

    // Create a fraud attestation: attested nonce does not exist yet
    function _createAttestation_fraud_nonExistingNonce() internal {
        test_dispatch();
        createFraudAttestation({ origin: DOMAIN_LOCAL, destination: DOMAIN_REMOTE, fakeNonce: 2 });
        // nonce = 2 doesn't exist yet => fraud
    }

    // Create a fraud attestation: attested root exists, but with a different nonce
    function _createAttestation_fraud_fakeNonce() internal {
        test_dispatch();
        test_dispatch();
        createFraudAttestation({ origin: DOMAIN_LOCAL, destination: DOMAIN_REMOTE, fakeNonce: 1 });
        // correct nonce for current root would be 2 => fraud
    }

    // Create a fraud attestation: attested root does not exist
    function _createAttestation_fraud_fakeRoot() internal {
        test_dispatch();
        createFraudAttestation({
            origin: DOMAIN_LOCAL,
            destination: DOMAIN_REMOTE,
            fakeRoot: "fake root"
        });
        // this is obv incorrect root for current nonce => fraud
    }
}
