// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract OriginHubEvents {
    /**
     * @notice Emitted when a correct report on a fraud attestation is submitted.
     * @param guard     Guard who signed the fraud report
     * @param report    Report data and signature
     */
    event CorrectFraudReport(address indexed guard, bytes report);

    /**
     * @notice Emitted when proof of an incorrect report is submitted.
     * @param guard     Guard who signed the incorrect report
     * @param report    Report data and signature
     */
    event IncorrectReport(address indexed guard, bytes report);

    /**
     * @notice Emitted when proof of a fraud attestation is submitted.
     * @param guards        Guards who signed the fraud attestation
     * @param notaries      Notaries who signed the fraud attestation
     * @param attestation   Attestation data and signature
     */
    event FraudAttestation(address[] guards, address[] notaries, bytes attestation);
}
