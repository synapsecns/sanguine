// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {IInterchainFactory} from "../interfaces/IInterchainFactory.sol";
import {Create2} from "../libs/Create2.sol";

import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

contract InterchainFactory is IInterchainFactory {
    address internal _firstArg;
    address internal _secondArg;

    /// @notice Emitted when a new InterchainERC20 is deployed
    /// @param interchainToken  The address of the deployed InterchainERC20
    /// @param underlyingToken  The address of the underlying token
    event InterchainTokenDeployed(address indexed interchainToken, address indexed underlyingToken);

    /// @notice Emitted when a new Processor is deployed
    /// @param processor  The address of the deployed Processor
    event ProcessorDeployed(address indexed processor);

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
        // Use underlyingToken = address(0) to deploy a standalone InterchainERC20
        bytes32 salt = _getDeploymentSalt(msg.sender, address(0));
        bytes memory tokenCreationCode = abi.encodePacked(tokenCode, abi.encode(name, symbol, decimals));
        // Use processor = address(0) to deploy a standalone InterchainERC20
        address interchainToken = _deployContract(tokenCreationCode, salt, initialAdmin, address(0));
        emit InterchainTokenDeployed(interchainToken, address(0));
    }

    /// @notice Deploys an InterchainERC20 for the given underlying token, and a Processor for the token pair.
    /// @dev The underlying token will be used to derive metadata for the InterchainERC20.
    /// Will revert if the underlying token has been used by the deployer before.
    /// @param underlyingToken          The address of the underlying token
    /// @param initialAdmin             The address of the initial admin for the token
    /// @param tokenCode                The creation code for the InterchainERC20 contract, without the constructor args
    /// @param processorCreationCode    The creation code for the Processor contract, without the constructor args
    function deployInterchainERC20WithProcessor(
        address underlyingToken,
        address initialAdmin,
        bytes memory tokenCode,
        bytes memory processorCreationCode
    )
        external
    {
        bytes32 salt = _getDeploymentSalt(msg.sender, underlyingToken);
        bytes memory tokenCreationCode = abi.encodePacked(tokenCode, _deriveMetadata(underlyingToken));
        // Predict deployment addresses for InterchainERC20 and Processor
        address interchainToken = Create2.predictDeployment(address(this), salt, tokenCreationCode);
        address processor = Create2.predictDeployment(address(this), salt, processorCreationCode);
        // Deploy InterchainERC20. We use assert statement as a sanity check for create2 prediction.
        address deployment = _deployContract(tokenCreationCode, salt, initialAdmin, processor);
        assert(deployment == interchainToken);
        emit InterchainTokenDeployed(interchainToken, underlyingToken);
        // Deploy Processor. We use assert statement as a sanity check for create2 prediction.
        // Note: Processor contract doesn't require any constructor arguments.
        deployment = _deployContract(processorCreationCode, salt, interchainToken, underlyingToken);
        assert(deployment == processor);
        emit ProcessorDeployed(processor);
    }

    /// @inheritdoc IInterchainFactory
    function getInterchainTokenDeployParameters() external view returns (address initialAdmin, address processor) {
        initialAdmin = _firstArg;
        processor = _secondArg;
        if (initialAdmin == address(0) && processor == address(0)) {
            revert InterchainFactory__NoActiveDeployment();
        }
    }

    /// @inheritdoc IInterchainFactory
    function getProcessorDeployParameters() external view returns (address interchainToken, address underlyingToken) {
        interchainToken = _firstArg;
        underlyingToken = _secondArg;
        if (interchainToken == address(0) && underlyingToken == address(0)) {
            revert InterchainFactory__NoActiveDeployment();
        }
    }

    /// @notice Predicts the address of the InterchainERC20 contract that will be deployed.
    /// @dev The underlying token will be used to derive metadata for the InterchainERC20.
    /// @param deployer        The address of the deployer
    /// @param underlyingToken The address of the underlying token
    /// @param tokenCode       The creation code for the InterchainERC20 contract, without the constructor arguments
    function predictInterchainERC20Address(
        address deployer,
        address underlyingToken,
        bytes memory tokenCode
    )
        external
        view
        returns (address)
    {
        bytes32 salt = _getDeploymentSalt(deployer, underlyingToken);
        bytes memory tokenCreationCode = abi.encodePacked(tokenCode, _deriveMetadata(underlyingToken));
        return Create2.predictDeployment(address(this), salt, tokenCreationCode);
    }

    /// @notice Predicts the address of the InterchainERC20 contract that will be deployed.
    /// @param deployer      The address of the deployer
    /// @param name          The name of the token
    /// @param symbol        The symbol of the token
    /// @param decimals      The number of decimals for the token
    /// @param tokenCode     The creation code for the InterchainERC20 contract, without the constructor arguments
    function predictInterchainERC20StandaloneAddress(
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
        bytes32 salt = _getDeploymentSalt(deployer, address(0));
        bytes memory tokenCreationCode = abi.encodePacked(tokenCode, abi.encode(name, symbol, decimals));
        return Create2.predictDeployment(address(this), salt, tokenCreationCode);
    }

    /// @notice Predicts the address of the Processor contract that will be deployed.
    /// @param deployer                 The address of the deployer
    /// @param processorCreationCode    The creation code for the Processor contract, without the constructor arguments
    function predictProcessorAddress(
        address deployer,
        address underlyingToken,
        bytes memory processorCreationCode
    )
        external
        view
        returns (address)
    {
        bytes32 salt = _getDeploymentSalt(deployer, underlyingToken);
        return Create2.predictDeployment(address(this), salt, processorCreationCode);
    }

    /// @dev Deploys a contract with the given creation code. Two extra arguments will be passed in
    /// the following callback to this factory.
    function _deployContract(
        bytes memory creationCode,
        bytes32 salt,
        address firstArg,
        address secondArg
    )
        internal
        returns (address deployedAt)
    {
        _firstArg = firstArg;
        _secondArg = secondArg;
        deployedAt = Create2.deploy(0, salt, creationCode);
        delete _firstArg;
        delete _secondArg;
    }

    /// @dev Derives metadata for InterchainERC20 from the underlying token
    /// Note: same amount of decimals will be used, and "Interchain" prefix will be added to the name and symbol
    function _deriveMetadata(address underlyingToken) internal view returns (bytes memory encodedMetadata) {
        string memory name = string.concat("Interchain", IERC20Metadata(underlyingToken).name());
        string memory symbol = string.concat("ic", IERC20Metadata(underlyingToken).symbol());
        uint8 decimals = IERC20Metadata(underlyingToken).decimals();
        return abi.encode(name, symbol, decimals);
    }

    /// @dev Returns the salt for InterchainERC20 and Processor Crate2 deployments.
    function _getDeploymentSalt(address deployer, address underlyingToken) internal pure returns (bytes32) {
        return keccak256(abi.encode(deployer, underlyingToken));
    }
}
