// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "../libs/TypedMemView.sol";
import { Report } from "../libs/Report.sol";
import { Auth } from "../libs/Auth.sol";

abstract contract AbstractGuardRegistry {
    using Report for bytes;
    using Report for bytes29;
    using TypedMemView for bytes29;

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

    function _isGuard(address _guard) internal view virtual returns (bool);
}
