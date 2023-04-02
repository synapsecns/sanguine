// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface IStateHub {
    /**
     * @notice Check that a given state is valid: matches the historical state of Origin contract.
     * Note: any Agent including an invalid state in their snapshot will be slashed
     * upon providing the snapshot and agent signature for it to Origin contract.
     * @dev Will revert if any of these is true:
     *  - State payload is not properly formatted.
     * @param statePayload      Raw payload with state data
     * @return isValid          Whether the provided state is valid
     */
    function isValidState(bytes memory statePayload) external view returns (bool isValid);

    /**
     * @notice Returns the amount of saved states so far.
     * @dev This includes the initial state of "empty Origin Merkle Tree".
     */
    function statesAmount() external view returns (uint256);

    /**
     * @notice Suggest the data (state after latest sent message) to sign for an Agent.
     * Note: signing the suggested state data will will never lead to slashing of the actor,
     * assuming they have confirmed that the block, which number is included in the data,
     * is not subject to reorganization (which is different for every observed chain).
     * @return statePayload     Raw payload with the latest state data
     */
    function suggestLatestState() external view returns (bytes memory statePayload);

    /**
     * @notice Given the historical nonce, suggest the state data to sign for an Agent.
     * Note: signing the suggested state data will will never lead to slashing of the actor,
     * assuming they have confirmed that the block, which number is included in the data,
     * is not subject to reorganization (which is different for every observed chain).
     * @param nonce             Historical nonce to form a state
     * @return statePayload     Raw payload with historical state data
     */
    function suggestState(uint32 nonce) external view returns (bytes memory statePayload);
}
