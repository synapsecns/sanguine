// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { LightManager } from "../../contracts/manager/LightManager.sol";
import { SystemMessage } from "../../contracts/libs/SystemMessage.sol";
import { AppHarness, SynapseTestStorage } from "./SynapseTestStorage.t.sol";
import { SynapseUtilities } from "./SynapseUtilities.t.sol";

// solhint-disable no-empty-blocks
// solhint-disable ordering
contract SynapseTestSuite is SynapseUtilities, SynapseTestStorage {
    /// @notice Prevents this contract from being included in the coverage report
    function testSynapseTestSuite() external {}

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                SETUP                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function setUp() public virtual {
        setupAgents();
        for (uint256 d = 0; d < DOMAINS; ++d) {
            setupChain(domains[d], domainNames[d]);
        }
        setupContext();
    }

    function setupAgents() public {
        domains = new uint32[](DOMAINS);
        domainNames = new string[](DOMAINS);
        domains[0] = DOMAIN_SYNAPSE;
        domains[1] = DOMAIN_LOCAL;
        domains[2] = DOMAIN_REMOTE;
        domainNames[0] = "SynChain";
        domainNames[1] = "Local";
        domainNames[2] = "Remote";
        // Create notaries
        for (uint256 domainIndex = 0; domainIndex < DOMAINS; ++domainIndex) {
            for (uint256 notaryIndex = 0; notaryIndex < NOTARIES_PER_CHAIN; ++notaryIndex) {
                uint32 domain = domains[domainIndex];
                uint256 notaryPK = generatePrivateKey(
                    abi.encode("notary", domainIndex, notaryIndex)
                );
                address notary = registerPK(notaryPK);
                vm.label(
                    notary,
                    string.concat("Notary ", domainNames[domainIndex], getActorSuffix(notaryIndex))
                );
                chains[domain].notaries.push(notary);
            }
        }
        // Create guards
        for (uint256 guardIndex = 0; guardIndex < GUARDS; ++guardIndex) {
            uint256 guardPK = generatePrivateKey(abi.encode("guard", guardIndex));
            address guard = registerPK(guardPK);
            vm.label(guard, string.concat("Guard", getActorSuffix(guardIndex)));
            guards.push(guard);
        }
        // Create owner
        owner = registerActor("Owner");
        // Create proxy admin
        proxyAdmin = registerActor("Proxy admin");
        // Create attacker
        attacker = registerActor("The Attackooor");
        // Create user
        user = registerActor("The Bridgooor");
        // Create broadcaster
        broadcaster = registerActor("Broadcaster");
        // Deploy proof generator
        clearMerkleTree();
    }

    // All contracts are deployed by this contract, the ownership is then transferred to `owner`
    // solhint-disable-next-line code-complexity
    function setupChain(uint32 domain, string memory chainName) public {
        // Deploy messaging contracts
        // LightManager agentManager = new LightManager(domain);
        // TODO: Setup destination
        // TODO: Setup origin
        // Setup AgentManager
        // agentManager.initialize();
        // agentManager.setSystemRouter(systemRouter);
        // Add global notaries via AgentManager
        for (uint256 i = 0; i < DOMAINS; ++i) {
            // uint32 domainToAdd = domains[i];
            // Origin and Destination will filter our agents themselves
            for (uint256 j = 0; j < NOTARIES_PER_CHAIN; ++j) {
                // address notary = suiteNotary(domainToAdd, j);
                // agentManager.addAgent(domainToAdd, notary);
            }
        }
        // Add guards  via AgentManager
        for (uint256 i = 0; i < GUARDS; ++i) {
            // agentManager.addAgent({ domain: 0, _account: guards[i] });
        }
        // Deploy app
        AppHarness app = new AppHarness(APP_OPTIMISTIC_SECONDS);
        // Transfer ownership everywhere
        // agentManager.transferOwnership(owner);
        // Label deployments
        // vm.label(address(destination), string.concat("Destination ", chainName));
        // vm.label(address(origin), string.concat("Origin ", chainName));
        // vm.label(address(agentManager), string.concat("AgentManager ", chainName));
        // vm.label(address(systemRouter), string.concat("SystemRouter ", chainName));
        vm.label(address(app), string.concat("App ", chainName));
        // Save deployments
        // chains[domain].destination = destination;
        // chains[domain].origin = origin;
        // chains[domain].agentManager = agentManager_;
        // chains[domain].systemRouter = systemRouter_;
        chains[domain].app = app;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                               SIGNING                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function signMessage(uint256 privKey, bytes memory message)
        public
        pure
        returns (bytes memory signature)
    {
        bytes32 digest = keccak256(message);
        digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", digest));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privKey, digest);
        signature = abi.encodePacked(r, s, v);
    }

    function signMessage(address signer, bytes memory message)
        public
        view
        returns (bytes memory signature)
    {
        uint256 privKey = privKeys[signer];
        require(privKey != 0, "Unknown account");
        return signMessage(privKey, message);
    }

    function signMessage(uint256[] memory keys, bytes memory message)
        public
        pure
        returns (bytes memory signatures)
    {
        for (uint256 i = 0; i < keys.length; ++i) {
            // There probably exists a more efficient way to do this without relying on TypedMemView
            signatures = bytes.concat(signatures, signMessage(keys[i], message));
        }
    }

    function signMessage(address[] memory signers, bytes memory message)
        public
        view
        returns (bytes memory signatures)
    {
        for (uint256 i = 0; i < signers.length; ++i) {
            // There probably exists a more efficient way to do this without relying on TypedMemView
            signatures = bytes.concat(signatures, signMessage(signers[i], message));
        }
    }

    function registerPK(uint256 privKey) public returns (address account) {
        account = vm.addr(privKey);
        // Save priv key for later usage
        privKeys[account] = privKey;
    }

    function registerActor(string memory actorName) public returns (address account) {
        account = registerPK(generatePrivateKey(abi.encode(actorName)));
        vm.label(account, actorName);
    }
}
