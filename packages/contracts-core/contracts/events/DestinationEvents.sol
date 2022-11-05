// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract DestinationEvents {
    /**
     * @notice Emitted when message is executed
     * @param remoteDomain  Remote domain where message originated
     * @param messageHash   The keccak256 hash of the message that was executed
     */
    event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash);

    /**
     * @notice Emitted when a root's confirmation is modified by governance
     * @param remoteDomain      The domain for which root's confirmAt has been set
     * @param root              The root for which confirmAt has been set
     * @param previousConfirmAt The previous value of confirmAt
     * @param newConfirmAt      The new value of confirmAt
     */
    event SetConfirmation(
        uint32 indexed remoteDomain,
        bytes32 indexed root,
        uint256 previousConfirmAt,
        uint256 newConfirmAt
    );

    /**
     * @notice Emitted when a Notary is blacklisted due to a submitted Guard's fraud Report
     * @param notary    The notary that was blacklisted
     * @param guard     The guard that signed the fraud report
     * @param reporter  The actor who submitted signed fraud report
     * @param report    Raw bytes of fraud report
     */
    event NotaryBlacklisted(
        address indexed notary,
        address indexed guard,
        address indexed reporter,
        bytes report
    );
}
