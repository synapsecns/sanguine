// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

import {InterchainFactory, IInterchainFactory} from "../../src/factories/InterchainFactory.sol";
import {Create2} from "../../src/libs/Create2.sol";
import {InterchainERC20} from "../../src/tokens/InterchainERC20.sol";
import {AbstractProcessor, BurningProcessor} from "../../src/processors/BurningProcessor.sol";
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

    address public deployerA;
    address public deployerB;

    function setUp() public {
        factory = new InterchainFactory();
        mockFactory = new MockInterchainFactory();

        deployerA = makeAddr("Deployer A");
        deployerB = makeAddr("Deployer B");
    }

    function checkInterchainTokenDeployParameters() internal {
        vm.expectRevert(IInterchainFactory.InterchainFactory__NoActiveDeployment.selector);
        factory.getInterchainTokenDeployParameters();
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

    function checkInterchainERC20AgainstMock(
        address deployed,
        TokenMetadata memory metadata,
        InterchainTokenParams memory params
    )
        internal
    {
        InterchainERC20 identicalMock = InterchainERC20(
            mockFactory.deployInterchainToken(
                metadata.name, metadata.symbol, metadata.decimals, params.initialAdmin, params.processor
            )
        );
        checkInterchainERC20Data(deployed, identicalMock);
        checkInterchainERC20Params(deployed, params);
    }

    function checkProcessorDeployParameters() internal {
        vm.expectRevert(IInterchainFactory.InterchainFactory__NoActiveDeployment.selector);
        factory.getProcessorDeployParameters();
    }

    function checkProcessor(address deployed, AbstractProcessor identicalMock) internal {
        assertEq(AbstractProcessor(deployed).INTERCHAIN_TOKEN(), AbstractProcessor(identicalMock).INTERCHAIN_TOKEN());
        assertEq(AbstractProcessor(deployed).UNDERLYING_TOKEN(), AbstractProcessor(identicalMock).UNDERLYING_TOKEN());
        assertEq(deployed.code, address(identicalMock).code);
    }

    function checkProcessorAgainstMock(
        address deployed,
        address interchainToken,
        address underlyingToken,
        function (address, address) external returns (address) deployProcessorMock
    )
        internal
    {
        AbstractProcessor identicalMock = AbstractProcessor(deployProcessorMock(interchainToken, underlyingToken));
        checkProcessor(deployed, identicalMock);
    }

    function predictInterchainERC20StandaloneAddress(
        address deployer,
        TokenMetadata memory metadata
    )
        internal
        view
        returns (address)
    {
        return factory.predictInterchainERC20StandaloneAddress(
            deployer, metadata.name, metadata.symbol, metadata.decimals, type(InterchainERC20).creationCode
        );
    }

    function predictInterchainERC20Address(address deployer, address underlyingToken) internal view returns (address) {
        return factory.predictInterchainERC20Address(deployer, underlyingToken, type(InterchainERC20).creationCode);
    }

    function predictProcessorAddress(
        address deployer,
        address underlyingToken,
        bytes memory processorCreationCode
    )
        internal
        view
        returns (address)
    {
        return factory.predictProcessorAddress(deployer, underlyingToken, processorCreationCode);
    }

    function deployStandalone(address deployer, TokenMetadata memory metadata, address initialAdmin) internal {
        vm.prank(deployer);
        factory.deployInterchainERC20Standalone(
            metadata.name, metadata.symbol, metadata.decimals, initialAdmin, type(InterchainERC20).creationCode
        );
    }

    function deployWithProcessor(
        address deployer,
        address underlyingToken,
        address initialAdmin,
        bytes memory processorCreationCode
    )
        internal
    {
        vm.prank(deployer);
        factory.deployInterchainERC20WithProcessor(
            underlyingToken, initialAdmin, type(InterchainERC20).creationCode, processorCreationCode
        );
    }

    function deployStandaloneAndCheck(
        address deployer,
        TokenMetadata memory metadata,
        InterchainTokenParams memory params
    )
        internal
    {
        // solhint-disable-next-line
        require(params.processor == address(0), "Processor should be zero address in this test");
        address predicted = predictInterchainERC20StandaloneAddress(deployer, metadata);
        checkInterchainTokenDeployParameters();
        vm.expectEmit(address(factory));
        emit InterchainTokenDeployed(predicted, address(0));
        deployStandalone(deployer, metadata, params.initialAdmin);
        checkInterchainERC20AgainstMock(predicted, metadata, params);
        checkInterchainTokenDeployParameters();
    }

    function deployWithProcessorAndCheck(
        address deployer,
        address underlyingToken,
        TokenMetadata memory interchainMetadata,
        address initialAdmin,
        bytes memory processorCreationCode,
        function (address, address) external returns (address) deployProcessorMock
    )
        internal
    {
        address predictedInterchainERC20 = predictInterchainERC20Address(deployer, underlyingToken);
        address predictedProcessor = predictProcessorAddress(deployer, underlyingToken, processorCreationCode);
        checkInterchainTokenDeployParameters();
        checkProcessorDeployParameters();
        vm.expectEmit(address(factory));
        emit InterchainTokenDeployed(predictedInterchainERC20, underlyingToken);
        vm.expectEmit(address(factory));
        emit ProcessorDeployed(predictedProcessor);
        deployWithProcessor(deployer, underlyingToken, initialAdmin, processorCreationCode);
        checkInterchainERC20AgainstMock(
            predictedInterchainERC20, interchainMetadata, InterchainTokenParams(initialAdmin, predictedProcessor)
        );
        checkProcessorAgainstMock(predictedProcessor, predictedInterchainERC20, underlyingToken, deployProcessorMock);
        checkInterchainTokenDeployParameters();
        checkProcessorDeployParameters();
    }

    // ═══════════════════════════════════════ TESTS: STANDALONE DEPLOYMENT ════════════════════════════════════════════

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
        address predicted = predictInterchainERC20StandaloneAddress(deployerA, metadata);
        deployStandalone(deployerA, metadata, address(1));
        vm.expectRevert(abi.encodeWithSelector(Create2.Create2__AlreadyDeployed.selector, predicted));
        deployStandalone(deployerA, metadata, address(2));
    }

    function test_deployStandalone_slightlyDifferentMetadata() public {
        InterchainTokenParams memory params = InterchainTokenParams(address(1), address(0));
        deployStandalone(deployerA, TokenMetadata("TokenA", "AAA", 10), params.initialAdmin);
        // Should succeed, as the name is different
        deployStandaloneAndCheck(deployerA, TokenMetadata("TokenB", "AAA", 10), params);
        // Should succeed, as the symbol is different
        deployStandaloneAndCheck(deployerA, TokenMetadata("TokenA", "AAB", 10), params);
        // Should succeed, as the decimals are different
        deployStandaloneAndCheck(deployerA, TokenMetadata("TokenA", "AAA", 11), params);
    }

    function test_deployStandalone_sameMetadata_diffDeployer() public {
        TokenMetadata memory metadata = TokenMetadata("TokenA", "AAA", 10);
        InterchainTokenParams memory params = InterchainTokenParams(address(1), address(0));
        deployStandalone(deployerA, metadata, params.initialAdmin);
        // Should succeed, as the deployer is different
        deployStandaloneAndCheck(deployerB, metadata, params);
    }

    // ════════════════════════════════════ TESTS: INTERCHAIN ERC20 + PROCESSOR ════════════════════════════════════════

    function test_deployWithBurningProcessor_deployerA() public {
        MockERC20Decimals underlyingToken = new MockERC20Decimals("TokenA", "AAA", 10);
        deployWithProcessorAndCheck(
            deployerA,
            address(underlyingToken),
            TokenMetadata("InterchainTokenA", "icAAA", 10),
            address(1),
            type(BurningProcessor).creationCode,
            mockFactory.deployBurningProcessor
        );
    }

    function test_deployWithLockingProcessor_deployerA() public {
        MockERC20Decimals underlyingToken = new MockERC20Decimals("TokenA", "AAA", 10);
        deployWithProcessorAndCheck(
            deployerA,
            address(underlyingToken),
            TokenMetadata("InterchainTokenA", "icAAA", 10),
            address(1),
            type(LockingProcessor).creationCode,
            mockFactory.deployLockingProcessor
        );
    }

    function test_deployWithBurningProcessor_deployerB() public {
        MockERC20Decimals underlyingToken = new MockERC20Decimals("TokenB", "BBB", 20);
        deployWithProcessorAndCheck(
            deployerB,
            address(underlyingToken),
            TokenMetadata("InterchainTokenB", "icBBB", 20),
            address(2),
            type(BurningProcessor).creationCode,
            mockFactory.deployBurningProcessor
        );
    }

    function test_deployWithLockingProcessor_deployerB() public {
        MockERC20Decimals underlyingToken = new MockERC20Decimals("TokenB", "BBB", 20);
        deployWithProcessorAndCheck(
            deployerB,
            address(underlyingToken),
            TokenMetadata("InterchainTokenB", "icBBB", 20),
            address(2),
            type(LockingProcessor).creationCode,
            mockFactory.deployLockingProcessor
        );
    }

    function test_deployWithProcessor_revert_sameUnderlying_sameProcessor() public {
        address underlyingToken = address(new MockERC20Decimals("TokenA", "AAA", 10));
        address predictedInterchainERC20 = predictInterchainERC20Address(deployerA, underlyingToken);
        deployWithProcessor(deployerA, underlyingToken, address(1), type(BurningProcessor).creationCode);
        vm.expectRevert(abi.encodeWithSelector(Create2.Create2__AlreadyDeployed.selector, predictedInterchainERC20));
        deployWithProcessor(deployerA, underlyingToken, address(2), type(BurningProcessor).creationCode);
    }

    function test_deployWithProcessor_revert_sameUnderlying_diffProcessor() public {
        address underlyingToken = address(new MockERC20Decimals("TokenA", "AAA", 10));
        address predictedInterchainERC20 = predictInterchainERC20Address(deployerA, underlyingToken);
        deployWithProcessor(deployerA, underlyingToken, address(1), type(BurningProcessor).creationCode);
        vm.expectRevert(abi.encodeWithSelector(Create2.Create2__AlreadyDeployed.selector, predictedInterchainERC20));
        deployWithProcessor(deployerA, underlyingToken, address(2), type(LockingProcessor).creationCode);
    }

    function test_deployWithProcessor_diffUnderlying_sameProcessor_sameDeployer() public {
        MockERC20Decimals underlyingTokenA = new MockERC20Decimals("TokenA", "AAA", 10);
        MockERC20Decimals underlyingTokenB = new MockERC20Decimals("TokenB", "BBB", 20);
        deployWithProcessor(deployerA, address(underlyingTokenA), address(1), type(BurningProcessor).creationCode);
        deployWithProcessorAndCheck(
            deployerA,
            address(underlyingTokenB),
            TokenMetadata("InterchainTokenB", "icBBB", 20),
            address(1),
            type(BurningProcessor).creationCode,
            mockFactory.deployBurningProcessor
        );
    }

    function test_deployWithProcessor_diffUnderlying_sameProcessor_sameDeployer_sameMetadata() public {
        MockERC20Decimals underlyingTokenA = new MockERC20Decimals("TokenA", "AAA", 10);
        MockERC20Decimals underlyingTokenCopy = new MockERC20Decimals("TokenA", "AAA", 10);
        deployWithProcessor(deployerA, address(underlyingTokenA), address(1), type(BurningProcessor).creationCode);
        deployWithProcessorAndCheck(
            deployerA,
            address(underlyingTokenCopy),
            TokenMetadata("InterchainTokenA", "icAAA", 10),
            address(1),
            type(BurningProcessor).creationCode,
            mockFactory.deployBurningProcessor
        );
    }

    function test_deployWithProcessor_diffUnderlying_diffProcessor_sameDeployer() public {
        MockERC20Decimals underlyingTokenA = new MockERC20Decimals("TokenA", "AAA", 10);
        MockERC20Decimals underlyingTokenB = new MockERC20Decimals("TokenB", "BBB", 20);
        deployWithProcessor(deployerA, address(underlyingTokenA), address(1), type(BurningProcessor).creationCode);
        deployWithProcessorAndCheck(
            deployerA,
            address(underlyingTokenB),
            TokenMetadata("InterchainTokenB", "icBBB", 20),
            address(1),
            type(LockingProcessor).creationCode,
            mockFactory.deployLockingProcessor
        );
    }

    function test_deployWithProcessor_diffUnderlying_diffProcessor_sameDeployer_sameMetadata() public {
        MockERC20Decimals underlyingTokenA = new MockERC20Decimals("TokenA", "AAA", 10);
        MockERC20Decimals underlyingTokenCopy = new MockERC20Decimals("TokenA", "AAA", 10);
        deployWithProcessor(deployerA, address(underlyingTokenA), address(1), type(BurningProcessor).creationCode);
        deployWithProcessorAndCheck(
            deployerA,
            address(underlyingTokenCopy),
            TokenMetadata("InterchainTokenA", "icAAA", 10),
            address(1),
            type(LockingProcessor).creationCode,
            mockFactory.deployLockingProcessor
        );
    }

    function test_deployWithProcessor_sameUnderlying_sameProcessor_diffDeployer() public {
        MockERC20Decimals underlyingToken = new MockERC20Decimals("TokenA", "AAA", 10);
        deployWithProcessor(deployerA, address(underlyingToken), address(1), type(BurningProcessor).creationCode);
        deployWithProcessorAndCheck(
            deployerB,
            address(underlyingToken),
            TokenMetadata("InterchainTokenA", "icAAA", 10),
            address(1),
            type(BurningProcessor).creationCode,
            mockFactory.deployBurningProcessor
        );
    }

    function test_deployWithProcessor_sameUnderlying_diffProcessor_diffDeployer() public {
        MockERC20Decimals underlyingToken = new MockERC20Decimals("TokenA", "AAA", 10);
        deployWithProcessor(deployerA, address(underlyingToken), address(1), type(BurningProcessor).creationCode);
        deployWithProcessorAndCheck(
            deployerB,
            address(underlyingToken),
            TokenMetadata("InterchainTokenA", "icAAA", 10),
            address(1),
            type(LockingProcessor).creationCode,
            mockFactory.deployLockingProcessor
        );
    }
}
