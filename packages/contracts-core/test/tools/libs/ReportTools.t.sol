// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "./AttestationTools.t.sol";

abstract contract ReportTools is AttestationTools {
    // Saved report data
    address internal reportGuard;
    // Saved report payloads
    bytes internal reportRaw;
    bytes internal signatureGuard;

    // Default Guard's Report with the given flag
    function createReport(Report.Flag flag) public {
        _createReport(flag, suiteGuard());
    }

    // Given Guards's Report with the given flag
    function createReport(Report.Flag flag, uint256 guardIndex) public {
        _createReport(flag, suiteGuard(guardIndex));
    }

    // Signer's Report with the given flag
    function createReport(Report.Flag flag, address signer) public {
        _createReport(flag, signer);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    function expectLogReport() public {
        // vm.expectEmit(true, true, true, true);
        // emit LogReport(reportGuard, attestationNotary, attestationRaw, reportRaw, reportRaw);
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           INTERNAL HELPERS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Create report using a given flag and saved data
    function _createReport(Report.Flag flag, address signer) internal {
        reportGuard = signer;
        (reportRaw, signatureGuard) = signReport(flag, attestationRaw, signer);
    }
}
