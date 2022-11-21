// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./libs/AttestationTools.t.sol";
import "./registry/GlobalNotaryRegistryTools.t.sol";
import { AttestationCollectorHarness } from "../harnesses/AttestationCollectorHarness.t.sol";

abstract contract AttestationCollectorTools is GlobalNotaryRegistryTools, AttestationTools {
    struct AttestationNonce {
        bytes attestation;
        uint32 nonce;
    }

    uint32 internal constant NONCE_TEST = 12;

    AttestationCollectorHarness internal attestationCollector;
    uint32 internal curDomain;

    // Saved attestation data
    // notary => [nonces]
    mapping(address => uint32[]) internal notaryNonces;
    // notary => [roots]
    mapping(address => bytes32[]) internal notaryRoots;
    // domain => (nonce => [roots])
    mapping(uint32 => mapping(uint32 => bytes[])) internal domainAttestations;
    // domain => (nonce => [roots])
    mapping(uint32 => mapping(uint32 => bytes32[])) internal domainRoots;

    // Latest attestation data
    mapping(address => AttestationNonce) internal notaryLatestAttestation;
    mapping(uint32 => AttestationNonce) internal domainLatestAttestation;

    function setupAttestationCollector() public {
        attestationCollector = new AttestationCollectorHarness();
        attestationCollector.initialize();
        attestationCollector.transferOwnership(owner);
    }

    function checkLatestAttestations() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            AttestationNonce memory domainLatest = domainLatestAttestation[domain];
            if (domainLatest.nonce == 0) {
                attestationCollectorGetLatestDomainAttestation({
                    domain: domain,
                    revertMessage: "No attestations found"
                });
            } else {
                assertEq(
                    attestationCollector.getLatestAttestation(domain),
                    domainLatest.attestation,
                    "!getLatestAttestation(domain)"
                );
            }
            for (uint256 index = 0; index < NOTARIES_PER_CHAIN; ++index) {
                address notary = suiteNotary(domain, index);
                AttestationNonce memory notaryLatest = notaryLatestAttestation[notary];
                if (notaryLatest.nonce == 0) {
                    attestationCollectorGetLatestNotaryAttestation({
                        domain: domain,
                        notaryIndex: index,
                        revertMessage: "No attestations found"
                    });
                } else {
                    assertEq(
                        attestationCollector.getLatestAttestation(domain, notary),
                        notaryLatest.attestation,
                        "!getLatestAttestation(domain,notary)"
                    );
                }
            }
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function submitTestAttestation(
        uint32 nonce,
        uint256 notaryIndex,
        bool isUnique
    ) public {
        // Use notaryIndex as salt by default
        submitTestAttestation(nonce, notaryIndex, notaryIndex, isUnique);
    }

    function submitTestAttestation(
        uint32 nonce,
        uint256 notaryIndex,
        uint256 salt,
        bool isUnique
    ) public {
        createAttestationMock({
            domain: curDomain,
            nonce: nonce,
            notaryIndex: notaryIndex,
            salt: salt
        });
        if (!isUnique) vm.expectRevert("Duplicated attestation");
        attestationCollectorSubmitAttestation({ returnValue: isUnique });
        if (isUnique) {
            // Save accepted attestation
            notaryNonces[attestationNotary].push(attestationNonce);
            notaryRoots[attestationNotary].push(attestationRoot);
            domainRoots[attestationDomain][attestationNonce].push(attestationRoot);
            domainAttestations[attestationDomain][attestationNonce].push(attestationRaw);
            // Update latest Notary attestation if needed
            if (attestationNonce > notaryLatestAttestation[attestationNotary].nonce) {
                notaryLatestAttestation[attestationNotary].nonce = attestationNonce;
                notaryLatestAttestation[attestationNotary].attestation = attestationRaw;
            }
            // Update latest domain attestation if needed
            if (attestationNonce > domainLatestAttestation[attestationDomain].nonce) {
                domainLatestAttestation[attestationDomain].nonce = attestationNonce;
                domainLatestAttestation[attestationDomain].attestation = attestationRaw;
            }
            // Check getLatestAttestation()
            checkLatestAttestations();
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectAttestationSubmitted() public {
        vm.expectEmit(true, true, true, true);
        emit AttestationSubmitted(attestationNotary, attestationRaw);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function attestationCollectorAddNotary(
        uint32 domain,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(domain, notaryIndex);
        vm.expectRevert(revertMessage);
        attestationCollector.addNotary(domain, notary);
    }

    function attestationCollectorRemoveNotary(
        uint32 domain,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(domain, notaryIndex);
        vm.expectRevert(revertMessage);
        attestationCollector.removeNotary(domain, notary);
    }

    function attestationCollectorSubmitAttestation(bytes memory revertMessage) public {
        vm.expectRevert(revertMessage);
        vm.prank(broadcaster);
        attestationCollector.submitAttestation(attestationRaw);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function attestationCollectorAddNotary(
        uint32 domain,
        uint256 notaryIndex,
        bool returnValue
    ) public {
        assertEq(
            attestationCollector.addNotary(domain, suiteNotary(domain, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    function attestationCollectorRemoveNotary(
        uint32 domain,
        uint256 notaryIndex,
        bool returnValue
    ) public {
        assertEq(
            attestationCollector.removeNotary(domain, suiteNotary(domain, notaryIndex)),
            returnValue,
            "!returnValue"
        );
    }

    function attestationCollectorSubmitAttestation(bool returnValue) public {
        vm.prank(broadcaster);
        assertEq(
            attestationCollector.submitAttestation(attestationRaw),
            returnValue,
            "!returnValue"
        );
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       TRIGGER VIEWS (REVERTS)                        ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function attestationCollectorGetAttestationByIndex(
        uint32 domain,
        uint32 nonce,
        uint256 attestationIndex,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getAttestation(domain, nonce, attestationIndex);
    }

    function attestationCollectorGetAttestationByRoot(
        uint32 domain,
        uint32 nonce,
        bytes32 root,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getAttestation(domain, nonce, root);
    }

    function attestationCollectorGetLatestDomainAttestation(
        uint32 domain,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getLatestAttestation(domain);
    }

    function attestationCollectorGetLatestNotaryAttestation(
        uint32 domain,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(domain, notaryIndex);
        vm.expectRevert(revertMessage);
        attestationCollector.getLatestAttestation(domain, notary);
    }

    function attestationCollectorGetRoot(
        uint32 domain,
        uint32 nonce,
        uint256 attestationIndex,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getRoot(domain, nonce, attestationIndex);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            TRIGGER VIEWS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Calls getAttestation(domain, nonce, index)
    function attestationCollectorGetAttestationByIndex(uint256 attestationIndex)
        public
        view
        returns (bytes memory)
    {
        return
            attestationCollector.getAttestation(
                attestationDomain,
                attestationNonce,
                attestationIndex
            );
    }

    // Calls getAttestation(domain, nonce, root)
    function attestationCollectorGetAttestationByRoot(bytes32 root)
        public
        view
        returns (bytes memory)
    {
        return attestationCollector.getAttestation(attestationDomain, attestationNonce, root);
    }

    // Calls getLatestAttestation(domain)
    function attestationCollectorGetLatestDomainAttestation() public view returns (bytes memory) {
        return attestationCollector.getLatestAttestation(attestationDomain);
    }

    // Calls getLatestAttestation(domain, notary)
    function attestationCollectorGetLatestNotaryAttestation(uint256 notaryIndex)
        public
        view
        returns (bytes memory)
    {
        return
            attestationCollector.getLatestAttestation(
                attestationDomain,
                suiteNotary(attestationDomain, notaryIndex)
            );
    }

    // Calls getRoot(domain, nonce, index)
    function attestationCollectorGetRoot(uint256 attestationIndex) public view returns (bytes32) {
        return attestationCollector.getRoot(attestationDomain, attestationNonce, attestationIndex);
    }
}
