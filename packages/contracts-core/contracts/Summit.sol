// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;
// ══════════════════════════════ LIBRARY IMPORTS ══════════════════════════════
import "./libs/State.sol";
// ═════════════════════════════ INTERNAL IMPORTS ══════════════════════════════
import { SummitEvents } from "./events/SummitEvents.sol";
import { InterfaceSummit } from "./interfaces/InterfaceSummit.sol";
import { SnapshotHub } from "./hubs/SnapshotHub.sol";
import { Attestation, Snapshot, StatementHub } from "./hubs/StatementHub.sol";
// ═════════════════════════════ EXTERNAL IMPORTS ══════════════════════════════
import {
    OwnableUpgradeable
} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @notice Accepts snapshots signed by Guards and Notaries. Verifies Notaries attestations.
 */
contract Summit is StatementHub, SnapshotHub, OwnableUpgradeable, SummitEvents, InterfaceSummit {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             INITIALIZER                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function initialize() external initializer {
        __Ownable_init_unchained();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            ADDING AGENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // TODO (Chi): merge Summit with BondingPrimary
    function addAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _addAgent(_domain, _account);
    }

    function removeAgent(uint32 _domain, address _account) external onlyOwner returns (bool) {
        return _removeAgent(_domain, _account);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          ACCEPT STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function submitSnapshot(bytes memory _snapPayload, bytes memory _snapSignature)
        external
        returns (bool wasAccepted)
    {
        // This will revert if payload is not a snapshot, or signer is not an active Agent
        (Snapshot snapshot, uint32 domain, address agent) = _verifySnapshot(
            _snapPayload,
            _snapSignature
        );
        if (domain == 0) {
            // This will revert if Guard has previously submitted
            // a fresher state than one in the snapshot.
            _acceptGuardSnapshot(snapshot, agent);
        } else {
            // This will revert if any of the states from the Notary snapshot
            // haven't been submitted by any of the Guards before.
            _acceptNotarySnapshot(snapshot, agent);
        }
        emit SnapshotAccepted(domain, agent, _snapPayload, _snapSignature);
        return true;
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VERIFY STATEMENTS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function verifyAttestation(bytes memory _attPayload, bytes memory _attSignature)
        external
        returns (bool isValid)
    {
        // This will revert if payload is not an attestation, or signer is not an active Notary
        (Attestation att, uint32 domain, address notary) = _verifyAttestation(
            _attPayload,
            _attSignature
        );
        isValid = _isValidAttestation(att);
        if (!isValid) {
            emit InvalidAttestation(_attPayload, _attSignature);
            _slashAgent(domain, notary);
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                                VIEWS                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /// @inheritdoc InterfaceSummit
    function getLatestState(uint32 _origin) external view returns (bytes memory statePayload) {
        uint256 guardsAmount = amountAgents(0);
        SummitState memory latestState;
        for (uint256 i = 0; i < guardsAmount; ++i) {
            address guard = getAgent(0, i);
            SummitState memory state = _latestState(_origin, guard);
            if (state.nonce > latestState.nonce) {
                latestState = state;
            }
        }
        // Check if we found anything
        if (latestState.nonce != 0) {
            statePayload = latestState.formatSummitState();
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL LOGIC                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _slashAgent(uint32 _domain, address _account) internal {
        // TODO: Move somewhere else?
        // TODO: send a system call indicating agent was slashed
        _removeAgent(_domain, _account);
    }

    function _isIgnoredAgent(uint32, address) internal view virtual override returns (bool) {
        // Summit keeps track of every agent
        return false;
    }
}
