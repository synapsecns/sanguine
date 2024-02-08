// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainFactory} from "../../src/factories/InterchainFactory.sol";
import {Create2} from "../../src/libs/Create2.sol";
import {InterchainERC20} from "../../src/tokens/InterchainERC20.sol";
import {BurningProcessor} from "../../src/processors/BurningProcessor.sol";
import {LockingProcessor} from "../../src/processors/LockingProcessor.sol";

import {MockERC20Decimals} from "../mocks/MockERC20Decimals.sol";
import {MockInterchainFactory} from "../mocks/MockInterchainFactory.sol";

import {Test} from "forge-std/Test.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract InterchainFactoryTest is Test {
    struct TokenMetadata {
        string name;
        string symbol;
        uint8 decimals;
    }

    struct InterchainTokenParams {
        address initialAdmin;
        address processor;
    }

    event InterchainTokenDeployed(address indexed interchainToken, address indexed underlyingToken);
    event ProcessorDeployed(address indexed processor);

    InterchainFactory public factory;
    MockInterchainFactory public mockFactory;
    MockERC20Decimals public underlyingToken;

    address public deployerA;
    address public deployerB;

    function setUp() public {
        factory = new InterchainFactory();
        mockFactory = new MockInterchainFactory();
        underlyingToken = new MockERC20Decimals("Token", "TKN", 42);

        deployerA = makeAddr("Deployer A");
        deployerB = makeAddr("Deployer B");
    }

    function checkInterchainERC20Data(address deployed, InterchainERC20 identicalMock) internal {
        assertEq(InterchainERC20(deployed).name(), identicalMock.name());
        assertEq(InterchainERC20(deployed).symbol(), identicalMock.symbol());
        assertEq(InterchainERC20(deployed).decimals(), identicalMock.decimals());
        assertEq(deployed.code, address(identicalMock).code);
    }

    function checkInterchainERC20Params(address deployed, InterchainTokenParams memory params) internal {
        assertEq(InterchainERC20(deployed).PROCESSOR(), params.processor);
        assertTrue(InterchainERC20(deployed).hasRole(0, params.initialAdmin));
    }

    function predictInterchainERC20Address(
        address deployer,
        TokenMetadata memory metadata
    )
        internal
        view
        returns (address)
    {
        return factory.predictInterchainERC20Address(
            deployer, metadata.name, metadata.symbol, metadata.decimals, type(InterchainERC20).creationCode
        );
    }

    function deployStandalone(
        address deployer,
        TokenMetadata memory metadata,
        InterchainTokenParams memory params
    )
        internal
    {
        vm.prank(deployer);
        factory.deployInterchainERC20Standalone(
            metadata.name, metadata.symbol, metadata.decimals, params.initialAdmin, type(InterchainERC20).creationCode
        );
    }

    function deployStandaloneAndCheck(
        address deployer,
        TokenMetadata memory metadata,
        InterchainTokenParams memory params
    )
        internal
    {
        require(params.processor == address(0), "Processor should be zero address in this test");
        address predicted = predictInterchainERC20Address(deployer, metadata);
        vm.expectEmit(address(factory));
        emit InterchainTokenDeployed(predicted, address(0));
        deployStandalone(deployer, metadata, params);
        checkInterchainERC20Params(predicted, params);
        InterchainERC20 identicalMock = InterchainERC20(
            mockFactory.deployInterchainToken(
                metadata.name, metadata.symbol, metadata.decimals, params.initialAdmin, address(0)
            )
        );
        checkInterchainERC20Data(predicted, identicalMock);
    }

    function test_deployStandalone_deployerA() public {
        deployStandaloneAndCheck(
            deployerA, TokenMetadata("TokenA", "AAA", 10), InterchainTokenParams(address(1), address(0))
        );
    }

    function test_deployStandalone_deployerB() public {
        deployStandaloneAndCheck(
            deployerB, TokenMetadata("TokenB", "BBB", 20), InterchainTokenParams(address(2), address(0))
        );
    }

    function test_deployStandalone_revert_sameMetadata() public {
        TokenMetadata memory metadata = TokenMetadata("TokenA", "AAA", 10);
        address predicted = predictInterchainERC20Address(deployerA, metadata);
        deployStandalone(deployerA, metadata, InterchainTokenParams(address(1), address(0)));
        vm.expectRevert(abi.encodeWithSelector(Create2.Create2__AlreadyDeployed.selector, predicted));
        deployStandalone(deployerA, metadata, InterchainTokenParams(address(2), address(0)));
    }

    function test_deployStandalone_slightlyDifferentMetadata() public {
        InterchainTokenParams memory params = InterchainTokenParams(address(1), address(0));
        deployStandalone(deployerA, TokenMetadata("TokenA", "AAA", 10), params);
        // Should succeed, as the name is different
        deployStandaloneAndCheck(deployerA, TokenMetadata("TokenB", "AAA", 10), params);
        // Should succeed, as the symbol is different
        deployStandaloneAndCheck(deployerA, TokenMetadata("TokenA", "AAB", 10), params);
        // Should succeed, as the decimals are different
        deployStandaloneAndCheck(deployerA, TokenMetadata("TokenA", "AAA", 11), params);
    }

    function test_deployStandalone_sameMetadata_differentDeployers() public {
        TokenMetadata memory metadata = TokenMetadata("TokenA", "AAA", 10);
        InterchainTokenParams memory params = InterchainTokenParams(address(1), address(0));
        deployStandalone(deployerA, metadata, params);
        // Should succeed, as the deployer is different
        deployStandaloneAndCheck(deployerB, metadata, params);
    }
}
