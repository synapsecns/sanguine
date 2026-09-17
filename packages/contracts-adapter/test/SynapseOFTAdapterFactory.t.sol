// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {SynapseOFTAdapter} from "../src/SynapseOFTAdapter.sol";
import {SynapseOFTAdapterFactory} from "../src/SynapseOFTAdapterFactory.sol";
import {ISynapseOFTAdapterFactory} from "../src/interfaces/ISynapseOFTAdapterFactory.sol";

import {EndpointMock} from "./mocks/EndpointMock.sol";
import {TestToken} from "./mocks/TestToken.sol";

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Test} from "forge-std/Test.sol";

contract SynapseOFTAdapterFactoryTest is Test {
    SynapseOFTAdapterFactory internal factory;
    EndpointMock internal endpoint;
    TestToken internal token;
    TestToken internal otherToken;

    address internal owner = makeAddr("Owner");
    address internal newOwner = makeAddr("New Owner");
    address internal user = makeAddr("User");

    event AdapterDeployed(address indexed token, address indexed adapter, bytes32 salt);
    event EndpointInitialized(address indexed endpoint);

    function setUp() public {
        endpoint = new EndpointMock();
        token = new TestToken();
        otherToken = new TestToken();
        factory = new SynapseOFTAdapterFactory(owner);
        vm.prank(owner);
        factory.initialize(address(endpoint));
    }

    function testFactoryCreate2AddressIncludesOnlyOwnerConstructorArgument() public {
        bytes32 salt = keccak256("Factory salt");
        bytes memory initCode = abi.encodePacked(type(SynapseOFTAdapterFactory).creationCode, abi.encode(owner));
        bytes32 digest = keccak256(abi.encodePacked(bytes1(0xff), address(this), salt, keccak256(initCode)));
        address predictedFactory = address(uint160(uint256(digest)));

        SynapseOFTAdapterFactory deployedFactory = new SynapseOFTAdapterFactory{salt: salt}(owner);

        assertEq(address(deployedFactory), predictedFactory);
        assertEq(deployedFactory.owner(), owner);
        assertEq(deployedFactory.endpoint(), address(0));

        vm.prank(owner);
        deployedFactory.initialize(address(endpoint));
        assertEq(deployedFactory.endpoint(), address(endpoint));
    }

    function testInitializeSetsEndpointAndEmitsEvent() public {
        SynapseOFTAdapterFactory uninitializedFactory = new SynapseOFTAdapterFactory(owner);

        vm.expectEmit(true, false, false, true, address(uninitializedFactory));
        emit EndpointInitialized(address(endpoint));
        vm.prank(owner);
        uninitializedFactory.initialize(address(endpoint));

        assertEq(uninitializedFactory.endpoint(), address(endpoint));
    }

    function testInitializeRevertsForNonOwner() public {
        SynapseOFTAdapterFactory uninitializedFactory = new SynapseOFTAdapterFactory(owner);

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user));
        vm.prank(user);
        uninitializedFactory.initialize(address(endpoint));

        assertEq(uninitializedFactory.endpoint(), address(0));
    }

    function testInitializeRevertsForZeroEndpoint() public {
        SynapseOFTAdapterFactory uninitializedFactory = new SynapseOFTAdapterFactory(owner);

        vm.expectRevert(SynapseOFTAdapterFactory.InvalidEndpoint.selector);
        vm.prank(owner);
        uninitializedFactory.initialize(address(0));

        assertEq(uninitializedFactory.endpoint(), address(0));
    }

    function testInitializeRevertsWhenAlreadyInitialized() public {
        EndpointMock replacementEndpoint = new EndpointMock();

        vm.expectRevert(SynapseOFTAdapterFactory.AlreadyInitialized.selector);
        vm.prank(owner);
        factory.initialize(address(replacementEndpoint));
    }

    function testDeployRevertsBeforeInitialization() public {
        SynapseOFTAdapterFactory uninitializedFactory = new SynapseOFTAdapterFactory(owner);

        vm.expectRevert(SynapseOFTAdapterFactory.NotInitialized.selector);
        vm.prank(owner);
        uninitializedFactory.deploy(address(token), keccak256("Uninitialized salt"));
    }

    function testDeployUsesExactCreate2AddressAndConstructorParams() public {
        bytes32 salt = keccak256("Adapter salt");
        address predictedAdapter = _predictAdapter(salt);

        vm.expectCall(address(endpoint), abi.encodeCall(EndpointMock.setDelegate, (owner)));
        vm.expectEmit(true, true, false, true, address(factory));
        emit AdapterDeployed(address(token), predictedAdapter, salt);
        vm.prank(owner);
        address deployedAdapter = factory.deploy(address(token), salt);

        assertEq(deployedAdapter, predictedAdapter);
        assertEq(SynapseOFTAdapter(deployedAdapter).token(), address(token));
        assertEq(address(SynapseOFTAdapter(deployedAdapter).endpoint()), address(endpoint));
        assertEq(SynapseOFTAdapter(deployedAdapter).owner(), owner);
    }

    function testDeployRevertsForNonOwner() public {
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user));
        vm.prank(user);
        factory.deploy(address(token), keccak256("Unauthorized salt"));
    }

    function testDeployRevertsForDuplicateSaltEvenWithDifferentToken() public {
        bytes32 salt = keccak256("Duplicate salt");

        vm.prank(owner);
        address firstAdapter = factory.deploy(address(token), salt);

        vm.expectRevert();
        vm.prank(owner);
        factory.deploy(address(otherToken), salt);

        assertEq(SynapseOFTAdapter(firstAdapter).token(), address(token));
    }

    function testDeployClearsConstructorParams() public {
        vm.prank(owner);
        factory.deploy(address(token), keccak256("Cleared params salt"));

        ISynapseOFTAdapterFactory.ConstructorParams memory params = factory.getConstructorParams();
        assertEq(params.token, address(0));
        assertEq(params.lzEndpoint, address(0));
        assertEq(params.owner, address(0));
    }

    function testTransferOwnershipAffectsOnlyNewAdapters() public {
        vm.prank(owner);
        address firstAdapter = factory.deploy(address(token), keccak256("First owner salt"));

        vm.prank(owner);
        factory.transferOwnership(newOwner);

        vm.prank(newOwner);
        address secondAdapter = factory.deploy(address(otherToken), keccak256("Second owner salt"));

        assertEq(SynapseOFTAdapter(firstAdapter).owner(), owner);
        assertEq(SynapseOFTAdapter(secondAdapter).owner(), newOwner);
    }

    function testDeploymentFailureRollsBackParamsAndDoesNotConsumeSalt() public {
        bytes32 salt = keccak256("Failed deployment salt");

        vm.expectRevert();
        vm.prank(owner);
        factory.deploy(address(0), salt);

        ISynapseOFTAdapterFactory.ConstructorParams memory params = factory.getConstructorParams();
        assertEq(params.token, address(0));
        assertEq(params.lzEndpoint, address(0));
        assertEq(params.owner, address(0));

        vm.prank(owner);
        address adapter = factory.deploy(address(token), salt);
        assertEq(adapter, _predictAdapter(salt));
        assertEq(SynapseOFTAdapter(adapter).token(), address(token));
    }

    function testSameFactoryAndAdapterAddressesAcrossChainsWithDifferentEndpointsAndTokens() public {
        EndpointMock firstEndpoint = new EndpointMock();
        EndpointMock secondEndpoint = new EndpointMock();
        TestToken firstToken = new TestToken();
        TestToken secondToken = new TestToken();
        bytes32 factorySalt = keccak256("Cross-chain factory salt");
        bytes32 adapterSalt = keccak256("Cross-chain adapter salt");
        uint256 snapshot = vm.snapshotState();

        SynapseOFTAdapterFactory firstFactory = new SynapseOFTAdapterFactory{salt: factorySalt}(owner);
        vm.startPrank(owner);
        firstFactory.initialize(address(firstEndpoint));
        address firstAdapter = firstFactory.deploy(address(firstToken), adapterSalt);
        vm.stopPrank();

        assertEq(address(SynapseOFTAdapter(firstAdapter).endpoint()), address(firstEndpoint));
        assertEq(SynapseOFTAdapter(firstAdapter).token(), address(firstToken));

        assertTrue(vm.revertToState(snapshot));

        SynapseOFTAdapterFactory secondFactory = new SynapseOFTAdapterFactory{salt: factorySalt}(owner);
        vm.startPrank(owner);
        secondFactory.initialize(address(secondEndpoint));
        address secondAdapter = secondFactory.deploy(address(secondToken), adapterSalt);
        vm.stopPrank();

        assertNotEq(address(firstEndpoint), address(secondEndpoint));
        assertNotEq(address(firstToken), address(secondToken));
        assertEq(address(secondFactory), address(firstFactory));
        assertEq(secondAdapter, firstAdapter);
        assertEq(address(SynapseOFTAdapter(secondAdapter).endpoint()), address(secondEndpoint));
        assertEq(SynapseOFTAdapter(secondAdapter).token(), address(secondToken));
    }

    function testConstructorSetsOwnerAndInitializeSetsEndpoint() public view {
        assertEq(factory.owner(), owner);
        assertEq(factory.endpoint(), address(endpoint));
    }

    function _predictAdapter(bytes32 salt) internal view returns (address) {
        bytes32 initCodeHash = keccak256(type(SynapseOFTAdapter).creationCode);
        bytes32 digest = keccak256(abi.encodePacked(bytes1(0xff), address(factory), salt, initCodeHash));
        return address(uint160(uint256(digest)));
    }
}
