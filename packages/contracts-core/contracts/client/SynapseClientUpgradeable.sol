// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { Client } from "./Client.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

abstract contract SynapseClientUpgradeable is Client, OwnableUpgradeable {
    // ============ Internal Variables ============

    /**
     * @dev Contracts addresses on the remote chains, which can:
     *      (1) send messages to this contract
     *      (2) receive messages from this contract
     */
    mapping(uint32 => bytes32) internal _trustedSenders;

    // ============ Upgrade gap ============

    // gap for upgrade safety
    uint256[49] private __GAP; //solhint-disable-line var-name-mixedcase

    // ============ Constructor ============

    // solhint-disable-next-line no-empty-blocks
    constructor(address origin_, address destination_) Client(origin_, destination_) {}

    // ============ Initializer ============

    // solhint-disable-next-line func-name-mixedcase
    function __SynapseClient_init() internal onlyInitializing {
        __Ownable_init_unchained();
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks
    function __SynapseClient_init_unchained() internal onlyInitializing {}

    // ============ Restricted Functions  ============

    /**
     * @notice  Sets the trusted sender for the given remote chain.
     * @dev     Only callable by owner (Governance).
     * @param remoteDomain      The domain of the remote chain
     * @param trustedSender_    The trusted sender
     */
    // solhint-disable-next-line ordering
    function setTrustedSender(uint32 remoteDomain, bytes32 trustedSender_) external onlyOwner {
        _setTrustedSender(remoteDomain, trustedSender_);
    }

    /**
     * @notice  Sets the trusted sender for a bunch of remote chains.
     * @dev     Only callable by owner (Governance).
     * @param remoteDomains     List of domains for the remote chains
     * @param trustedSenders    List of trusted senders for given chains
     */
    function setTrustedSenders(uint32[] calldata remoteDomains, bytes32[] calldata trustedSenders)
        external
        onlyOwner
    {
        uint256 length = trustedSenders.length;
        require(remoteDomains.length == length, "!arrays");
        for (uint256 i = 0; i < length; ) {
            _setTrustedSender(remoteDomains[i], trustedSenders[i]);
            unchecked {
                ++i;
            }
        }
    }

    // ============ Public Functions  ============

    /// @notice Returns the trusted sender for the given remote chain.
    function trustedSender(uint32 remoteDomain) public view override returns (bytes32) {
        return _trustedSenders[remoteDomain];
    }

    // ============ Internal Functions  ============

    /// @dev Checks both domain and trusted sender, then updates the records.
    function _setTrustedSender(uint32 remoteDomain, bytes32 trustedSender_) internal {
        require(remoteDomain != 0, "!domain");
        require(trustedSender_ != bytes32(0), "!sender");
        _trustedSenders[remoteDomain] = trustedSender_;
    }
}
