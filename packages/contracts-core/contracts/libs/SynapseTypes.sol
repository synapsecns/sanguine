// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

/**
 * @dev
 * 1. `bytes29` is supposed to be used as a memory view over a generic byte string.
 * 2. CallData is a memory view over the payload, that is supposed to be used for an external call, i.e.
 * recipient.call(callData). Its length is always (4 + 32 * N) bytes:
 * - First 4 bytes represent the function selector.
 * - 32 * N bytes represent N words that function arguments occupy.
 * 3. Signature is a memory view over a "65 bytes" string
 * that is an off-chain agent signature for some data.
 */

type CallData is bytes29;
type Signature is bytes29;

/// @dev SystemMessage is a memory view over the message with instructions for a system call.
type SystemMessage is bytes29;

/// @dev Attestation is a memory view over a formatted attestation payload.
type Attestation is bytes29;
/// @dev AttestationData is a memory view over a formatted attestation data.
type AttestationData is bytes29;

/// @dev Report is a memory view over a formatted report payload.
type Report is bytes29;
/// @dev ReportData is a memory view over a formatted report data.
type ReportData is bytes29;

/**
 * @dev The goal of this library is to assign a type for every bytes29 memory pointer
 * and enforce strict type checking for every bytes29 operation. This will prevent
 * a misuse of libraries, i.e. using Attestation functions on a Report pointer.
 */
library SynapseTypes {
    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                         0X03: MESSAGE                                ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // prettier-ignore
    uint40 internal constant MESSAGE                    = 0x03_01_00_00_00;
    // prettier-ignore
    uint40 internal constant MESSAGE_HEADER             = 0x03_01_01_00_00;
    // prettier-ignore
    uint40 internal constant MESSAGE_TIPS               = 0x03_01_02_00_00;
}
