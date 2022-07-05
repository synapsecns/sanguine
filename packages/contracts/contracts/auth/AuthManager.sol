// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { HomeUpdate } from "../libs/HomeUpdate.sol";
import { Auth } from "../libs/Auth.sol";

abstract contract AuthManager {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                              LIBRARIES                               ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    using HomeUpdate for bytes29;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             UPGRADE GAP                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    uint256[50] private __GAP;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          INTERNAL FUNCTIONS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * @notice  Checks if the passed payload is a valid HomeUpdate message,
     *          if the signature is valid and if the signer is an authorized updater.
     * @param _updater      Signer of the message, needs to be authorized as updater, revert otherwise.
     * @param _payload      Message with update of Home merkle root. Needs to be valid, revert otherwise.
     * @param _signature    `_payload` signed by `_updater`. Needs to be valid, revert otherwise.
     */
    function _checkUpdaterAuth(
        address _updater,
        bytes memory _payload,
        bytes memory _signature
    ) internal view returns (bytes29 _view) {
        // This will revert if signature is invalid
        _view = Auth.checkSignature(_updater, _payload, _signature);
        require(_view.isValidUpdate(), "Message is not a valid update");
        require(_isUpdater(_view.homeDomain(), _updater), "Signer is not an updater");
    }

    function _checkWatchtowerAuth(
        address _watchtower,
        bytes memory _payload,
        bytes memory _signature
    ) internal view returns (bytes29 _view) {
        require(_isWatchtower(_watchtower), "Signer is not a watchtower");
        _view = Auth.checkSignature(_watchtower, _payload, _signature);
        // TODO: check if _payload is valid, once watchtower message standard is finalized
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          VIRTUAL FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function _isUpdater(uint32 _homeDomain, address _updater) internal view virtual returns (bool);

    function _isWatchtower(address _watchtower) internal view virtual returns (bool);
}
