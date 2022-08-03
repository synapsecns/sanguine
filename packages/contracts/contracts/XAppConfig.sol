// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { Home } from "./Home.sol";
import { ReplicaManager } from "./ReplicaManager.sol";
import { TypeCasts } from "./libs/TypeCasts.sol";
// ============ External Imports ============
import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title XAppConnectionManager
 * @author Illusory Systems Inc.
 * @notice Manages a registry of local Replica contracts
 * for remote Home domains. Acepts Guard signatures
 * to un-enroll Replicas attached to fraudulent remote Homes
 */
contract XAppConfig is Ownable {
    // ============ Public Storage ============

    // Home contract
    Home public home;
    // local Replica address => remote Home domain
    mapping(address => uint32) public replicaToDomain;
    // remote Home domain => local Replica address
    mapping(uint32 => address) public domainToReplica;
    // guard address => replica remote domain => has/doesn't have permission
    mapping(address => mapping(uint32 => bool)) private guardPermissions;

    // ============ Events ============

    /**
     * @notice Emitted when a new Replica is enrolled / added
     * @param domain the remote domain of the Home contract for the Replica
     * @param replica the address of the Replica
     */
    event ReplicaEnrolled(uint32 indexed domain, address replica);

    /**
     * @notice Emitted when a new Replica is un-enrolled / removed
     * @param domain the remote domain of the Home contract for the Replica
     * @param replica the address of the Replica
     */
    event ReplicaUnenrolled(uint32 indexed domain, address replica);

    /**
     * @notice Emitted when Guard permissions are changed
     * @param domain the remote domain of the Home contract for the Replica
     * @param guard the address of the Guard
     * @param access TRUE if the Guard was given permissions, FALSE if permissions were removed
     */
    event GuardPermissionSet(uint32 indexed domain, address guard, bool access);

    // ============ Constructor ============

    // solhint-disable-next-line no-empty-blocks
    constructor() Ownable() {}

    // ============ External Functions ============

    /**
     * @notice Un-Enroll a replica contract
     * in the case that fraud was detected on the Home
     * @dev in the future, if fraud occurs on the Home contract,
     * the Guard will submit their signature directly to the Home
     * and it can be relayed to all remote chains to un-enroll the Replicas
     * @param _domain the remote domain of the Home contract for the Replica
     * @param _notary the address of the Notary for the Home contract (also stored on Replica)
     * @param _signature signature of guard on (domain, replica address, notary address)
     */
    function unenrollReplica(
        uint32 _domain,
        bytes32 _notary,
        bytes memory _signature
    ) external {
        // ensure that the replica is currently set
        address _replica = domainToReplica[_domain];
        require(_replica != address(0), "!replica exists");
        // ensure that the signature is on the proper notary
        require(
            ReplicaManager(_replica).notary() == TypeCasts.bytes32ToAddress(_notary),
            "!current notary"
        );
        // get the guard address from the signature
        // and ensure that the guard has permission to un-enroll this replica
        address _guard = _recoverGuardFromSig(
            _domain,
            TypeCasts.addressToBytes32(_replica),
            _notary,
            _signature
        );
        require(guardPermissions[_guard][_domain], "!valid guard");
        // remove the replica from mappings
        _unenrollReplica(_replica);
    }

    /**
     * @notice Set the address of the local Home contract
     * @param _home the address of the local Home contract
     */
    function setHome(address _home) external onlyOwner {
        home = Home(_home);
    }

    /**
     * @notice Allow Owner to enroll Replica contract
     * @param _replica the address of the Replica
     * @param _domain the remote domain of the Home contract for the Replica
     */
    function ownerEnrollReplica(address _replica, uint32 _domain) external onlyOwner {
        // un-enroll any existing replica
        _unenrollReplica(_replica);
        // add replica and domain to two-way mapping
        replicaToDomain[_replica] = _domain;
        domainToReplica[_domain] = _replica;
        emit ReplicaEnrolled(_domain, _replica);
    }

    /**
     * @notice Allow Owner to un-enroll Replica contract
     * @param _replica the address of the Replica
     */
    function ownerUnenrollReplica(address _replica) external onlyOwner {
        _unenrollReplica(_replica);
    }

    /**
     * @notice Allow Owner to set Guard permissions for a Replica
     * @param _guard the address of the Guard
     * @param _domain the remote domain of the Home contract for the Replica
     * @param _access TRUE to give the Guard permissions, FALSE to remove permissions
     */
    function setGuardPermission(
        address _guard,
        uint32 _domain,
        bool _access
    ) external onlyOwner {
        guardPermissions[_guard][_domain] = _access;
        emit GuardPermissionSet(_domain, _guard, _access);
    }

    /**
     * @notice Query local domain from Home
     * @return local domain
     */
    function localDomain() external view returns (uint32) {
        return home.localDomain();
    }

    /**
     * @notice Get access permissions for the guard on the domain
     * @param _guard the address of the guard
     * @param _domain the domain to check for guard permissions
     * @return TRUE iff _guard has permission to un-enroll replicas on _domain
     */
    function guardPermission(address _guard, uint32 _domain) external view returns (bool) {
        return guardPermissions[_guard][_domain];
    }

    // ============ Public Functions ============

    /**
     * @notice Check whether _replica is enrolled
     * @param _replica the replica to check for enrollment
     * @return TRUE iff _replica is enrolled
     */
    function isReplica(address _replica) public view returns (bool) {
        return replicaToDomain[_replica] != 0;
    }

    // ============ Internal Functions ============

    /**
     * @notice Remove the replica from the two-way mappings
     * @param _replica replica to un-enroll
     */
    function _unenrollReplica(address _replica) internal {
        uint32 _currentDomain = replicaToDomain[_replica];
        domainToReplica[_currentDomain] = address(0);
        replicaToDomain[_replica] = 0;
        emit ReplicaUnenrolled(_currentDomain, _replica);
    }

    /**
     * @notice Get the Guard address from the provided signature
     * @return address of guard that signed
     */
    function _recoverGuardFromSig(
        uint32 _domain,
        bytes32 _replica,
        bytes32 _notary,
        bytes memory _signature
    ) internal view returns (address) {
        bytes32 _homeDomainHash = ReplicaManager(TypeCasts.bytes32ToAddress(_replica))
            .homeDomainHash(_domain);
        bytes32 _digest = keccak256(abi.encodePacked(_homeDomainHash, _domain, _notary));
        _digest = ECDSA.toEthSignedMessageHash(_digest);
        return ECDSA.recover(_digest, _signature);
    }

    /**
     * @dev should be impossible to renounce ownership;
     * we override OpenZeppelin Ownable implementation
     * of renounceOwnership to make it a no-op
     */
    function renounceOwnership() public override onlyOwner {
        // do nothing
    }
}
