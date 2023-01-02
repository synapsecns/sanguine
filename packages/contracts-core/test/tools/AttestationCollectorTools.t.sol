// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./libs/AttestationTools.t.sol";
import { AttestationCollectorHarness } from "../harnesses/AttestationCollectorHarness.t.sol";

abstract contract AttestationCollectorTools is AttestationTools {
    AttestationCollectorHarness internal collector;

    function setupCollector() public {
        collector = new AttestationCollectorHarness();
        collector.initialize();
        for (uint256 i = 0; i < GUARDS; ++i) {
            collector.addAgent({ _domain: 0, _account: suiteGuard(i) });
        }
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            for (uint256 i = 0; i < NOTARIES_PER_CHAIN; ++i) {
                collector.addAgent(domain, suiteNotary({ domain: domain, index: i }));
            }
        }
        collector.transferOwnership(owner);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectAttestationSaved(
        uint256 signatureIndex,
        bool isGuard,
        uint256 agentIndex
    ) public {
        bytes memory attData = ra.data;
        bytes memory signature = signMessage({
            signer: isGuard ? attestationGuards[agentIndex] : attestationNotaries[agentIndex],
            message: attData
        });
        bytes memory attestation = Attestation.formatAttestation({
            _data: attData,
            _guardSignatures: isGuard ? signature : bytes(""),
            _notarySignatures: isGuard ? bytes("") : signature
        });
        vm.expectEmit(true, true, true, true);
        emit AttestationSaved(signatureIndex, attestation);
    }
}
