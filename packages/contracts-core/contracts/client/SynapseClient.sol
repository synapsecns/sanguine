// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { Client } from "./Client.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

abstract contract SynapseClient is Client, Ownable {
    // ============ Internal Variables ============

    /**
     * @dev Contracts addresses on the remote chains, which can:
     *      (1) send messages to this contract
     *      (2) receive messages from this contract
     */
    mapping(uint32 => bytes32) internal trustedSenders;

    // ============ Constructor ============

    // solhint-disable-next-line no-empty-blocks
    constructor(address _origin, address _destination) Client(_origin, _destination) {}

    // ============ Restricted Functions  ============

    /**
     * @notice  Sets the trusted sender for the given remote chain.
     * @dev     Only callable by owner (Governance).
     * @param _remoteDomain     The domain of the remote chain
     * @param _trustedSender    The trusted sender
     */
    function setTrustedSender(uint32 _remoteDomain, bytes32 _trustedSender) external onlyOwner {
        _setTrustedSender(_remoteDomain, _trustedSender);
    }

    /**
     * @notice  Sets the trusted sender for a bunch of remote chains.
     * @dev     Only callable by owner (Governance).
     * @param _remoteDomains    List of domains for the remote chains
     * @param _trustedSenders   List of trusted senders for given chains
     */
    function setTrustedSenders(uint32[] calldata _remoteDomains, bytes32[] calldata _trustedSenders)
        external
        onlyOwner
    {
        uint256 length = _trustedSenders.length;
        require(_remoteDomains.length == length, "!arrays");
        for (uint256 i = 0; i < length; ) {
            _setTrustedSender(_remoteDomains[i], _trustedSenders[i]);
            unchecked {
                ++i;
            }
        }
    }

    // ============ Public Functions  ============

    /// @notice Returns the trusted sender for the given remote chain.
    function trustedSender(uint32 _remoteDomain) public view override returns (bytes32) {
        return trustedSenders[_remoteDomain];
    }

    // ============ Internal Functions  ============

    /// @dev Checks both domain and trusted sender, then updates the records.
    function _setTrustedSender(uint32 _remoteDomain, bytes32 _trustedSender) internal {
        require(_remoteDomain != 0, "!domain");
        require(_trustedSender != bytes32(0), "!sender");
        trustedSenders[_remoteDomain] = _trustedSender;
    }
}
