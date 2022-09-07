// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { Report } from "../libs/Report.sol";
import { Auth } from "../libs/Auth.sol";

abstract contract AbstractGuardRegistry {
    using Report for bytes;
    using Report for bytes29;
    using TypedMemView for bytes29;

    /**
     * @notice  Checks all following statements are true:
     *          - `_report` is a formatted Report payload
     *          - `_report` contains a signature
     *          - such signature belongs to an authorized Guard
     * @param _report   Report on a Attestation of Origin merkle root
     * @return _guard   Notary that signed the Attestation
     * @return _view    Memory view on report
     */
    function _checkGuardAuth(bytes memory _report)
        internal
        view
        returns (address _guard, bytes29 _view)
    {
        _view = _report.castToReport();
        require(_view.isReport(), "Not a report");
        _guard = Auth.recoverSigner(_view.reportData(), _view.guardSignature().clone());
        require(_isGuard(_guard), "Signer is not a guard");
    }

    /**
     * @notice Checks whether a given account in an authorized Guard.
     * @dev Child contracts should implement this depending on how Guards are stored.
     * @param _account  Address to check for being a Guard
     * @return TRUE if the account is an authorized Guard.
     */
    function _isGuard(address _account) internal view virtual returns (bool);
}
