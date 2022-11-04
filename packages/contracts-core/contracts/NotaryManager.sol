// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

// ============ Internal Imports ============
import { INotaryManager } from "./interfaces/INotaryManager.sol";
import { Destination } from "./Destination.sol";
// ============ External Imports ============
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title NotaryManager
 * @author Illusory Systems Inc.
 * @notice MVP / centralized version of contract
 * that will manage Notary bonding, slashing,
 * selection and rotation
 */
contract NotaryManager is INotaryManager, Ownable {
    // ============ Public Storage ============

    // address of destination contract
    address public destination;

    // ============ Private Storage ============

    // address of the current notary
    address private _notary;

    // ============ Events ============

    /**
     * @notice Emitted when a new destination is set
     * @param origin The address of the new origin contract
     */
    event NewDestination(address destination);

    /**
     * @notice Emitted when a new notary is set
     * @param notary The address of the new notary
     */
    event NewNotary(address notary);

    /**
     * @notice Emitted when slashNotary is called
     */
    event FakeSlashed(address reporter);

    // ============ Modifiers ============

    /**
     * @notice Require that the function is called
     * by the Origin contract
     */
    modifier onlyDestination() {
        require(msg.sender == destination, "!destination");
        _;
    }

    // ============ Constructor ============

    constructor(address _notaryAddress) payable Ownable() {
        _notary = _notaryAddress;
    }

    // ============ External Functions ============

    /**
     * @notice Set the address of the a new destination contract
     * @dev only callable by trusted owner
     * @param _destination The address of the new destination contract
     */
    function setDestination(address _destination) external onlyOwner {
        require(Address.isContract(_destination), "!contract destination");
        destination = _destination;

        emit NewDestination(_destination);
    }

    /**
     * @notice Set the address of a new notary
     * @dev only callable by trusted owner
     * @param _notaryAddress The address of the new notary
     */
    function setNotary(address _notaryAddress) external onlyOwner {
        _notary = _notaryAddress;
        Destination(destination).setNotary(_notaryAddress);
        emit NewNotary(_notaryAddress);
    }

    /**
     * @notice Slashes the notary
     * @dev Currently does nothing, functionality will be implemented later
     * when notary bonding and rotation are also implemented
     * @param _reporter The address of the entity that reported the notary fraud
     */
    function slashNotary(address payable _reporter) external override onlyDestination {
        emit FakeSlashed(_reporter);
    }

    /**
     * @notice Get address of current notary
     * @return the notary address
     */
    function notary() external view override returns (address) {
        return _notary;
    }

    /**
     * @dev should be impossible to renounce ownership;
     * we override OpenZeppelin Ownable implementation
     * of renounceOwnership to make it a no-op
     */
    // solhint-disable-next-line no-empty-blocks
    function renounceOwnership() public override onlyOwner {
        // do nothing
    }
}
