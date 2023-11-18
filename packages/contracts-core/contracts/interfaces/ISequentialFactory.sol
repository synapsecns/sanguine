// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

interface ISequentialFactory {
    /**
     * @notice Deploys an arbitrary contract at the predicted address. The predicted address is
     * calculated using the nonce of the factory and the address of the factory itself.
     * The context of `code` does not influence the address of the deployment.
     * @dev Will revert if nonce is not equal to the nonce of the factory.
     * This is done to enable a multi-transaction deployment, where one unsuccessful deployment
     * will lead to the nonce not being incremented and the next deployments will also fail.
     * This guarantees the sequentiality of the deployments, and that one failed deployment will
     * not lead to the predicted address being occupied by a different contract.
     * Only the contract owner can deploy.
     * @param nonce     The nonce of the factory
     * @param code      The bytecode of the contract to deploy
     */
    function deploy(uint256 nonce, bytes memory code) external returns (address deployedAt);

    /**
     * @notice Returns the nonce of the factory. Using this value for the next deployment will
     * guarantee the successful deployment.
     */
    function getNonce() external view returns (uint256 nonce);

    /**
     * @notice Returns the address of the contract that would be deployed using `deploy(nonce, *)`
     * regardless of the `code` passed.
     * @param nonce     The nonce of the factory
     */
    function predictDeployment(uint256 nonce) external view returns (address deployedAt);

    /**
     * @notice Returns the next address of the contract that would be deployed using `deploy(getNonce())`
     * regardless of the `code` passed.
     */
    function predictNextDeployment() external view returns (address deployedAt);
}
