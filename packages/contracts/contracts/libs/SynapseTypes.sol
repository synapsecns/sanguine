// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import { TypedMemView } from "./TypedMemView.sol";

/**
 * @dev The goal of this library is to assign a type for every bytes29 memory pointer
 * and enforce strict type checking for every bytes29 operation. This will prevent
 * a misuse of libraries, i.e. using Attestation functions on a Report pointer.
 */
library SynapseTypes {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X00: SIGNATURE                              ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant SIGNATURE                  = 0x00_01_00_00_00;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X01: ATTESTATION                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant ATTESTATION                = 0x01_01_00_00_00;
    // prettier-ignore
    uint40 internal constant ATTESTATION_DATA           = 0x01_01_01_00_00;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X02: REPORT                                 ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant REPORT                     = 0x02_01_00_00_00;
    // prettier-ignore
    uint40 internal constant REPORT_DATA                = 0x02_01_01_00_00;

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X03: MESSAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant MESSAGE                    = 0x03_01_00_00_00;
    // prettier-ignore
    uint40 internal constant MESSAGE_HEADER             = 0x03_01_01_00_00;
    // prettier-ignore
    uint40 internal constant MESSAGE_TIPS               = 0x03_01_02_00_00;
    // prettier-ignore
    uint40 internal constant MESSAGE_BODY               = 0x03_01_03_00_00;

    // prettier-ignore
    uint40 internal constant SYSTEM_MESSAGE             = 0x03_02_00_00_00;
    // prettier-ignore
    uint40 internal constant SYSTEM_MESSAGE_CALL        = 0x03_02_01_00_00;
    // prettier-ignore
    uint40 internal constant SYSTEM_MESSAGE_ADJUST      = 0x03_02_02_00_00;
}
