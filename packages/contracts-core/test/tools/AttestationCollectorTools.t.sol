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
    uint32 internal currentOrigin;
    uint32 internal currentDestination;
    // uint32 internal curDomain;

    // Saved attestation data
    // notary => [nonces]
    // mapping(address => uint32[]) internal notaryNonces;
    // notary => [roots]
    // mapping(address => bytes32[]) internal notaryRoots;

    // attestationKey => [attestations]
    mapping(uint96 => bytes[]) internal keyAttestations;
    // attestationKey => [roots]
    mapping(uint96 => bytes32[]) internal keyRoots;

    // Latest attestation data
    // (attestationDomains => (notary => [attestation+nonce]))
    mapping(uint64 => mapping(address => AttestationNonce)) internal notaryLatestAttestation;
    // (attestationDomains => [attestation+nonce])
    mapping(uint64 => AttestationNonce) internal domainsLatestAttestation;

    function setupAttestationCollector() public {
        attestationCollector = new AttestationCollectorHarness();
        attestationCollector.initialize();
        attestationCollector.transferOwnership(owner);
    }

    // solhint-disable-next-line code-complexity
    function checkLatestAttestations() public {
        for (uint256 o = 0; o < DOMAINS; ++o) {
            uint32 origin = domains[o];
            for (uint256 d = 0; d < DOMAINS; ++d) {
                // Ignore origin == destination
                if (o == d) continue;
                uint32 destination = domains[d];
                uint64 _domains = Attestation.attestationDomains(origin, destination);
                AttestationNonce memory domainLatest = domainsLatestAttestation[_domains];
                if (domainLatest.nonce == 0) {
                    attestationCollectorGetLatestDomainAttestation({
                        origin: origin,
                        destination: destination,
                        revertMessage: "No attestations found"
                    });
                } else {
                    assertEq(
                        attestationCollector.getLatestAttestation(origin, destination),
                        domainLatest.attestation,
                        "!getLatestAttestation(domains)"
                    );
                }
                for (uint256 index = 0; index < NOTARIES_PER_CHAIN; ++index) {
                    address notary = suiteNotary(origin, index);
                    AttestationNonce memory notaryLatest = notaryLatestAttestation[_domains][
                        notary
                    ];
                    if (notaryLatest.nonce == 0) {
                        attestationCollectorGetLatestNotaryAttestation({
                            origin: origin,
                            destination: destination,
                            notaryIndex: index,
                            revertMessage: "No attestations found"
                        });
                    } else {
                        assertEq(
                            attestationCollector.getLatestAttestation(origin, destination, notary),
                            notaryLatest.attestation,
                            "!getLatestAttestation(domains,notary)"
                        );
                    }
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
            origin: currentOrigin,
            destination: currentDestination,
            nonce: nonce,
            notaryIndex: notaryIndex,
            salt: salt
        });
        if (isUnique) {
            expectAttestationAccepted();
        } else {
            vm.expectRevert("Duplicated attestation");
        }
        attestationCollectorSubmitAttestation({ returnValue: isUnique });
        if (isUnique) {
            // Save accepted attestation
            // notaryNonces[attestationNotary].push(attestationNonce);
            // notaryRoots[attestationNotary].push(attestationRoot);
            keyRoots[attestationKey].push(attestationRoot);
            keyAttestations[attestationKey].push(attestationRaw);
            // Update latest Notary attestation if needed
            if (
                attestationNonce >
                notaryLatestAttestation[attestationDomains][attestationNotary].nonce
            ) {
                notaryLatestAttestation[attestationDomains][attestationNotary] = AttestationNonce(
                    attestationRaw,
                    attestationNonce
                );
            }
            // Update latest domain attestation if needed
            if (attestationNonce > domainsLatestAttestation[attestationDomains].nonce) {
                domainsLatestAttestation[attestationDomains] = AttestationNonce(
                    attestationRaw,
                    attestationNonce
                );
            }
            // Check getLatestAttestation()
            checkLatestAttestations();
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectAttestationAccepted() public {
        vm.expectEmit(true, true, true, true);
        emit AttestationAccepted(attestationNotary, attestationRaw);
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
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        uint256 attestationIndex,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getAttestation(origin, destination, nonce, attestationIndex);
    }

    function attestationCollectorGetAttestationByRoot(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        bytes32 root,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getAttestation(origin, destination, nonce, root);
    }

    function attestationCollectorGetLatestDomainAttestation(
        uint32 origin,
        uint32 destination,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getLatestAttestation(origin, destination);
    }

    function attestationCollectorGetLatestNotaryAttestation(
        uint32 origin,
        uint32 destination,
        uint256 notaryIndex,
        bytes memory revertMessage
    ) public {
        address notary = suiteNotary(origin, notaryIndex);
        vm.expectRevert(revertMessage);
        attestationCollector.getLatestAttestation(origin, destination, notary);
    }

    function attestationCollectorGetRoot(
        uint32 origin,
        uint32 destination,
        uint32 nonce,
        uint256 attestationIndex,
        bytes memory revertMessage
    ) public {
        vm.expectRevert(revertMessage);
        attestationCollector.getRoot(origin, destination, nonce, attestationIndex);
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
                attestationOrigin,
                attestationDestination,
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
        return
            attestationCollector.getAttestation(
                attestationOrigin,
                attestationDestination,
                attestationNonce,
                root
            );
    }

    // Calls getLatestAttestation(domain)
    function attestationCollectorGetLatestDomainAttestation() public view returns (bytes memory) {
        return attestationCollector.getLatestAttestation(attestationOrigin, attestationDestination);
    }

    // Calls getLatestAttestation(domain, notary)
    function attestationCollectorGetLatestNotaryAttestation(uint256 notaryIndex)
        public
        view
        returns (bytes memory)
    {
        return
            attestationCollector.getLatestAttestation(
                attestationOrigin,
                attestationDestination,
                suiteNotary(attestationOrigin, notaryIndex)
            );
    }

    // Calls getRoot(domain, nonce, index)
    function attestationCollectorGetRoot(uint256 attestationIndex) public view returns (bytes32) {
        return
            attestationCollector.getRoot(
                attestationOrigin,
                attestationDestination,
                attestationNonce,
                attestationIndex
            );
    }
}
