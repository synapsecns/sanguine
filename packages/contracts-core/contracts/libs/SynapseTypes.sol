// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// The goal of this library is to assign a type for every generic memory view (bytes29)
// and enforce compile-time strict type checking for every operation with the memory views.
// This will prevent a misuse of libraries, i.e. using Attestation functions on a Report view.
// Every type is supposed to define a method to wrap a generic memory view into a given type,
// while checking that the view is over a properly formatted payload, i.e. `castToAttestation()`.
// Different types may define methods with the same name without any issues:
//      Message msg;
//      msg.nonce();    // gets a message nonce
//      AttestationData data;
//      data.nonce();   // gets an attestation nonce

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

/// @dev Message is a memory over over a formatted message payload.
type Message is bytes29;
/// @dev Header is a memory over over a formatted message header payload.
type Header is bytes29;
/// @dev Tips is a memory over over a formatted message tips payload.
type Tips is bytes29;
