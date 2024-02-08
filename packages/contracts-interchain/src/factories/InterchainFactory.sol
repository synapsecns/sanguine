// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {IInterchainFactory} from "../interfaces/IInterchainFactory.sol";
import {Create2} from "../libs/Create2.sol";

contract InterchainFactory is IInterchainFactory {
    address internal _firstArg;
    address internal _secondArg;

    /// @notice Emitted when a new InterchainERC20 is deployed
    /// @param interchainToken  The address of the deployed InterchainERC20
    /// @param underlyingToken  The address of the underlying token
    event InterchainTokenDeployed(address indexed interchainToken, address indexed underlyingToken);

    /// @notice Deploys a standalone InterchainERC20 with the given metadata.
    /// @dev Will revert if the metadata has been used by the deployer before.
    /// @param name             The name of the token
    /// @param symbol           The symbol of the token
    /// @param decimals         The number of decimals for the token
    /// @param initialAdmin     The address of the initial admin for the token
    /// @param tokenCode        The creation code for the InterchainERC20 contract, without the constructor arguments
    function deployInterchainERC20Standalone(
        string memory name,
        string memory symbol,
        uint8 decimals,
        address initialAdmin,
        bytes memory tokenCode
    )
        external
    {
        bytes memory tokenCreationCode = abi.encodePacked(tokenCode, abi.encode(name, symbol, decimals));
        // Use processor = address(0) to deploy a standalone InterchainERC20
        address interchainToken = _deployContract(tokenCreationCode, initialAdmin, address(0));
        emit InterchainTokenDeployed(interchainToken, address(0));
    }

    /// @inheritdoc IInterchainFactory
    function getInterchainTokenDeployParameters() external view returns (address initialAdmin, address processor) {
        initialAdmin = _firstArg;
        processor = _secondArg;
    }

    /// @inheritdoc IInterchainFactory
    function getProcessorDeployParameters() external view returns (address interchainToken, address underlyingToken) {
        interchainToken = _firstArg;
        underlyingToken = _secondArg;
    }

    /// @notice Predicts the address of the InterchainERC20 contract that will be deployed.
    /// @param deployer      The address of the deployer
    /// @param name          The name of the token
    /// @param symbol        The symbol of the token
    /// @param decimals      The number of decimals for the token
    /// @param tokenCode     The creation code for the InterchainERC20 contract, without the constructor arguments
    function predictInterchainERC20Address(
        address deployer,
        string memory name,
        string memory symbol,
        uint8 decimals,
        bytes memory tokenCode
    )
        external
        view
        returns (address)
    {
        bytes memory tokenCreationCode = abi.encodePacked(tokenCode, abi.encode(name, symbol, decimals));
        return Create2.predictDeployment(address(this), _getDeployerSalt(deployer), tokenCreationCode);
    }

    /// @dev Deploys a contract with the given creation code. Two extra arguments will be passed in
    /// the following callback to this factory.
    function _deployContract(
        bytes memory creationCode,
        address firstArg,
        address secondArg
    )
        internal
        returns (address deployedAt)
    {
        _firstArg = firstArg;
        _secondArg = secondArg;
        deployedAt = Create2.deploy(0, _getDeployerSalt(msg.sender), creationCode);
        delete _firstArg;
        delete _secondArg;
    }

    /// @dev Returns the deployer-dependent salt for the Create2 deployment
    function _getDeployerSalt(address deployer) internal pure returns (bytes32) {
        return bytes32(bytes20(deployer));
    }
}
