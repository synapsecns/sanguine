// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { Message } from "./libs/Message.sol";
// ============ External Imports ============
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title UpdaterStorage
 * @author Illusory Systems Inc.
 * @notice Shared utilities between Home and Replica.
 */
abstract contract UpdaterStorage is Initializable, OwnableUpgradeable {
    // ============ Immutable Variables ============

    // Domain of chain on which the contract is deployed
    uint32 public immutable localDomain;

    // ============ Public Variables ============

    // Address of bonded Updater
    address public updater;

    // ============ Upgrade Gap ============

    // gap for upgrade safety
    uint256[49] private __GAP;

    // ============ Events ============

    /**
     * @notice Emitted when update is made on Home
     * or unconfirmed update root is submitted on Replica
     * @param homeDomain Domain of home contract
     * @param oldRoot Old merkle root
     * @param newRoot New merkle root
     * @param signature Updater's signature on `oldRoot` and `newRoot`
     */
    event Update(
        uint32 indexed homeDomain,
        bytes32 indexed oldRoot,
        bytes32 indexed newRoot,
        bytes signature
    );

    /**
     * @notice Emitted when Updater is rotated
     * @param oldUpdater The address of the old updater
     * @param newUpdater The address of the new updater
     */
    event NewUpdater(address oldUpdater, address newUpdater);

    // ============ Constructor ============

    constructor(uint32 _localDomain) {
        localDomain = _localDomain;
    }

    // ============ Initializer ============

    function __SynapseBase_initialize(address _updater) internal onlyInitializing {
        __Ownable_init();
        _setUpdater(_updater);
    }

    // ============ Internal Functions ============

    /**
     * @notice Hash of domain concatenated with "SYN"
     * @param _domain The domain to hash
     */
    function _domainHash(uint32 _domain) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_domain, "SYN"));
    }

    /**
     * @notice Set the Updater
     * @param _newUpdater Address of the new Updater
     */
    function _setUpdater(address _newUpdater) internal {
        address _oldUpdater = updater;
        updater = _newUpdater;
        emit NewUpdater(_oldUpdater, _newUpdater);
    }

    /**
     * @notice Checks that signature was signed by Updater
     * @param _homeDomain Domain of Home contract where the signing was done
     * @param _oldRoot Old merkle root
     * @param _newRoot New merkle root
     * @param _signature Signature on `_oldRoot` and `_newRoot`
     * @return TRUE if signature is valid signed by updater
     **/
    function _isUpdaterSignature(
        uint32 _homeDomain,
        bytes32 _oldRoot,
        bytes32 _newRoot,
        bytes memory _signature
    ) internal view returns (bool) {
        bytes32 _digest = keccak256(abi.encodePacked(_domainHash(_homeDomain), _oldRoot, _newRoot));
        _digest = ECDSA.toEthSignedMessageHash(_digest);
        return (ECDSA.recover(_digest, _signature) == updater);
    }

    /**
     * @dev should be impossible to renounce ownership;
     * we override OpenZeppelin OwnableUpgradeable's
     * implementation of renounceOwnership to make it a no-op
     */
    function renounceOwnership() public override onlyOwner {
        // do nothing
    }
}
