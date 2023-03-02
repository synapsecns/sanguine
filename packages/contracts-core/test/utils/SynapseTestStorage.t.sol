// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../../contracts/bonding/BondingManager.sol";
// ============ Harnesses ============
import "../harnesses/client/AppHarness.t.sol";
import "../harnesses/system/SystemRouterHarness.t.sol";
// ============ Utils ============
import "./SynapseConstants.t.sol";
import "./SynapseEvents.t.sol";
import "./proof/HistoricalProofGenerator.t.sol";

contract SynapseTestStorage is SynapseConstants, SynapseEvents {
    struct TestDeployments {
        BondingManager bondingManager;
        SystemRouterHarness systemRouter;
        AppHarness app;
        address[] notaries;
    }

    struct MessageContext {
        uint32 origin;
        address sender;
        uint32 destination;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           STORAGE: CHAINS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // domain => chain contracts
    mapping(uint32 => TestDeployments) internal chains;
    // All test domains
    uint32[] internal domains;
    // Names of test domains
    string[] internal domainNames;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           STORAGE: ACTORS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // address => private key
    mapping(address => uint256) internal privKeys;
    // All guards
    address[] internal guards;
    // Address that is set up as contract owner for all deployed contracts in the tests
    address internal owner;
    // Address that is set up as admin for all deployed proxies in the tests
    address internal proxyAdmin;
    // Attacker address
    address internal attacker;
    // User address
    address internal user;
    // Broadcaster address
    address internal broadcaster;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            STORAGE: MISC                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Merkle proof generator
    HistoricalProofGenerator internal proofGen;
    // Context for tests
    MessageContext internal userLocalToRemote;
    MessageContext internal userRemoteToLocal;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            CONTEXT SETUP                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function setupContext() public {
        // Context for tests, where user is sending a message
        userLocalToRemote = MessageContext({
            origin: DOMAIN_LOCAL,
            sender: user,
            destination: DOMAIN_REMOTE
        });
        userRemoteToLocal = MessageContext({
            origin: DOMAIN_REMOTE,
            sender: user,
            destination: DOMAIN_LOCAL
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            MERKLE PROOFS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function clearMerkleTree() public {
        proofGen = new HistoricalProofGenerator();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                       GETTERS: CHAIN CONTRACTS                       ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function suiteBondingManager(uint32 domain) public view returns (BondingManager) {
        return chains[domain].bondingManager;
    }

    function suiteSystemRouter(uint32 domain) public view returns (SystemRouterHarness) {
        return chains[domain].systemRouter;
    }

    function suiteApp(uint32 domain) public view returns (AppHarness) {
        return chains[domain].app;
    }

    function suiteAgent(uint32 domain) public view returns (address) {
        return suiteAgent(domain, 0);
    }

    function suiteAgent(uint32 domain, uint256 index) public view returns (address) {
        if (domain == 0) {
            return suiteGuard(index);
        } else {
            return suiteNotary(domain, index);
        }
    }

    function suiteGuard() public view returns (address) {
        return suiteGuard(0);
    }

    function suiteGuard(uint256 index) public view returns (address) {
        if (index >= guards.length) return address(0);
        return guards[index];
    }

    function suiteNotary(uint32 domain) public view returns (address) {
        return suiteNotary(domain, 0);
    }

    function suiteNotary(uint32 domain, uint256 index) public view returns (address) {
        if (index >= chains[domain].notaries.length) return address(0);
        return chains[domain].notaries[index];
    }

    // Return a default notary from domain other than a given one
    function suiteForeignNotary(uint32 domain) public view returns (address) {
        return suiteNotary(foreignDomain(domain));
    }

    // Return a notary with a given index from domain other than a given one
    function suiteForeignNotary(uint32 domain, uint256 index) public view returns (address) {
        return suiteNotary(foreignDomain(domain), index);
    }

    // Return test domain, that is different from a given domain
    function foreignDomain(uint32 domain) public pure returns (uint32) {
        if (domain == DOMAIN_LOCAL) return DOMAIN_REMOTE;
        if (domain == DOMAIN_REMOTE) return DOMAIN_SYNAPSE;
        if (domain == DOMAIN_SYNAPSE) return DOMAIN_LOCAL;
        revert("Unknown domain");
    }
}
