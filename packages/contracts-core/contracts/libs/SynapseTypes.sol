// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import { TypedMemView } from "./TypedMemView.sol";

/**
 * @dev The goal of this library is to assign a type for every bytes29 memory pointer
 * and enforce strict type checking for every bytes29 operation. This will prevent
 * a misuse of libraries, i.e. using Attestation functions on a Report pointer.
 */
library SynapseTypes {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          0X00: BYTE STRINGS                          ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    /**
     * 1. RAW_BYTES refers to a generic byte string, that is not supposed to be parsed
     * by the messaging contracts. RAW_BYTES is set to uint40(0) so that
     * the "default zero" type would represent a generic byte string.
     * 2. SIGNATURE refers to 65 bytes string that is an off-chain agent signature for some data.
     * 3. CALL_PAYLOAD refers to the payload, that is supposed to be used for an external call, i.e.
     * recipient.call(CALL_PAYLOAD). Its length is always (4 + 32 * N) bytes:
     *      - First 4 bytes represent the function selector.
     *      - 32 * N bytes represent N function arguments.
     */
    // prettier-ignore
    uint40 internal constant RAW_BYTES                  = 0x00_00_00_00_00;
    // prettier-ignore
    uint40 internal constant SIGNATURE                  = 0x00_01_00_00_00;
    // prettier-ignore
    uint40 internal constant CALL_PAYLOAD               = 0x00_02_00_00_00;

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

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             0X04: SYSTEM                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant SYSTEM_CALL                = 0x04_00_00_00_00;
}
