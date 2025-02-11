// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
// solhint-disable

import "forge-std/Test.sol";
import "forge-std/console2.sol";

import "../contracts/FastBridge.sol";
import "../contracts/interfaces/IFastBridge.sol";
import "../contracts/libs/Errors.sol";
import "../contracts/libs/UniversalToken.sol";

import "./mocks/MockERC20.sol";

import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

contract FastBridgeTest is Test {
    FastBridge public fastBridge;

    uint256 public constant TX_DEADLINE = 60 minutes;

    address owner = address(1);
    address relayer = address(2);
    address guard = address(3);
    address user = address(4);
    address dstUser = address(5);
    address governor = address(6);
    address refunder = address(7);
    MockERC20 arbUSDC;
    MockERC20 ethUSDC;

    function setUp() public virtual {
        vm.chainId(42_161);
        fastBridge = FastBridge(deployFastBridge());
        arbUSDC = new MockERC20("arbUSDC", 6);
        ethUSDC = new MockERC20("ethUSDC", 6);
        _mintTokensToActors();
    }

    function deployFastBridge() internal virtual returns (address) {
        return address(new FastBridge(owner));
    }

    function _mintTokensToActors() internal virtual {
        arbUSDC.mint(relayer, 100 * 10 ** 6);
        arbUSDC.mint(guard, 100 * 10 ** 6);
        arbUSDC.mint(user, 100 * 10 ** 6);
        arbUSDC.mint(dstUser, 100 * 10 ** 6);
        ethUSDC.mint(relayer, 100 * 10 ** 6);
        ethUSDC.mint(guard, 100 * 10 ** 6);
        ethUSDC.mint(user, 100 * 10 ** 6);
        ethUSDC.mint(dstUser, 100 * 10 ** 6);
    }

    function assertCorrectProof(
        bytes32 transactionId,
        uint256 expectedTimestamp,
        address expectedRelayer
    )
        internal
        virtual
    {
        (uint96 proofTimestamp, address proofRelayer) = fastBridge.bridgeProofs(transactionId);
        assertEq(proofTimestamp, uint96(expectedTimestamp));
        assertEq(proofRelayer, expectedRelayer);
    }

    function _getBridgeRequestAndId(
        uint256 chainId,
        uint256 currentNonce,
        uint256 protocolFeeRate
    )
        internal
        view
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = address(arbUSDC);
        address destToken = address(ethUSDC);
        uint256 originAmount = 11 * 10 ** 6;
        uint256 destAmount = 10.97e6;
        uint256 deadline = block.timestamp + TX_DEADLINE;

        uint256 originFeeAmount;
        if (protocolFeeRate > 0) originFeeAmount = (originAmount * protocolFeeRate) / 1e6;
        originAmount -= originFeeAmount;

        // Calculate the expected transaction ID
        request = abi.encode(
            IFastBridge.BridgeTransaction({
                originChainId: originChainId,
                destChainId: dstChainId,
                originSender: user,
                destRecipient: to,
                originToken: originToken,
                destToken: destToken,
                originAmount: originAmount,
                destAmount: destAmount,
                originFeeAmount: originFeeAmount,
                sendChainGas: false,
                deadline: deadline,
                nonce: currentNonce
            })
        );
        transactionId = keccak256(request);
    }

    function _getBridgeRequestAndIdWithETH(
        uint256 chainId,
        uint256 currentNonce,
        uint256 protocolFeeRate
    )
        internal
        view
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = UniversalTokenLib.ETH_ADDRESS;
        address destToken = UniversalTokenLib.ETH_ADDRESS;
        uint256 originAmount = 11 * 10 ** 18;
        uint256 destAmount = 10.97e18;
        uint256 deadline = block.timestamp + TX_DEADLINE;

        uint256 originFeeAmount;
        if (protocolFeeRate > 0) originFeeAmount = (originAmount * protocolFeeRate) / 1e6;
        originAmount -= originFeeAmount;

        // Calculate the expected transaction ID
        request = abi.encode(
            IFastBridge.BridgeTransaction({
                originChainId: originChainId,
                destChainId: dstChainId,
                originSender: user,
                destRecipient: to,
                originToken: originToken,
                destToken: destToken,
                originAmount: originAmount,
                destAmount: destAmount,
                originFeeAmount: originFeeAmount,
                sendChainGas: false,
                deadline: deadline,
                nonce: currentNonce
            })
        );
        transactionId = keccak256(request);
    }

    function _getBridgeRequestAndIdWithChainGas(
        uint256 chainId,
        uint256 currentNonce,
        uint256 protocolFeeRate
    )
        internal
        view
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = address(arbUSDC);
        address destToken = address(ethUSDC);
        uint256 originAmount = 11 * 10 ** 6;
        uint256 destAmount = 10.97e6;
        uint256 deadline = block.timestamp + TX_DEADLINE;

        uint256 originFeeAmount;
        if (protocolFeeRate > 0) originFeeAmount = (originAmount * protocolFeeRate) / 1e6;
        originAmount -= originFeeAmount;

        // Calculate the expected transaction ID
        request = abi.encode(
            IFastBridge.BridgeTransaction({
                originChainId: originChainId,
                destChainId: dstChainId,
                originSender: user,
                destRecipient: to,
                originToken: originToken,
                destToken: destToken,
                originAmount: originAmount,
                destAmount: destAmount,
                originFeeAmount: originFeeAmount,
                sendChainGas: true,
                deadline: deadline,
                nonce: currentNonce
            })
        );
        transactionId = keccak256(request);
    }

    function _getBridgeRequestAndIdWithETHAndChainGas(
        uint256 chainId,
        uint256 currentNonce,
        uint256 protocolFeeRate
    )
        internal
        view
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = UniversalTokenLib.ETH_ADDRESS;
        address destToken = UniversalTokenLib.ETH_ADDRESS;
        uint256 originAmount = 11 * 10 ** 18;
        uint256 destAmount = 10.97e18;
        uint256 deadline = block.timestamp + TX_DEADLINE;

        uint256 originFeeAmount;
        if (protocolFeeRate > 0) originFeeAmount = (originAmount * protocolFeeRate) / 1e6;
        originAmount -= originFeeAmount;

        // Calculate the expected transaction ID
        request = abi.encode(
            IFastBridge.BridgeTransaction({
                originChainId: originChainId,
                destChainId: dstChainId,
                originSender: user,
                destRecipient: to,
                originToken: originToken,
                destToken: destToken,
                originAmount: originAmount,
                destAmount: destAmount,
                originFeeAmount: originFeeAmount,
                sendChainGas: true,
                deadline: deadline,
                nonce: currentNonce
            })
        );
        transactionId = keccak256(request);
    }

    function expectUnauthorized(address caller, bytes32 role) internal {
        vm.expectRevert(abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, caller, role));
    }

    function setUpRoles() public {
        vm.startPrank(owner);
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayer);
        fastBridge.grantRole(fastBridge.GUARD_ROLE(), guard);
        fastBridge.grantRole(fastBridge.GOVERNOR_ROLE(), governor);
        fastBridge.grantRole(fastBridge.REFUNDER_ROLE(), refunder);
        assertTrue(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        assertTrue(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        assertTrue(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        assertTrue(fastBridge.hasRole(fastBridge.REFUNDER_ROLE(), refunder));
        vm.stopPrank();
    }

    /// @notice Test to check if the owner is correctly set
    function test_owner() public view {
        assertTrue(fastBridge.hasRole(fastBridge.DEFAULT_ADMIN_ROLE(), owner));
    }

    /// @notice Test to check if a relayer can be successfully added
    function test_successfulAddRelayer() public {
        vm.startPrank(owner);
        assertFalse(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), relayer);
        assertTrue(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
    }

    /// @notice Test to check if only an admin can add a relayer
    function test_onlyAdminCanAddRelayer() public {
        bytes32 relayerRole = fastBridge.RELAYER_ROLE();
        vm.startPrank(relayer);
        assertFalse(fastBridge.hasRole(relayerRole, relayer));
        vm.expectRevert();
        fastBridge.grantRole(relayerRole, relayer);
    }

    /// @notice Test to check if a relayer can be successfully removed
    function test_successfulRemoveRelayer() public {
        test_successfulAddRelayer();
        assertTrue(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        vm.startPrank(owner);
        fastBridge.revokeRole(fastBridge.RELAYER_ROLE(), relayer);
        assertFalse(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
    }

    /// @notice Test to check if only an admin can remove a relayer
    function test_onlyAdminCanRemoveRelayer() public {
        bytes32 relayerRole = fastBridge.RELAYER_ROLE();
        test_successfulAddRelayer();
        vm.startPrank(relayer);
        assertTrue(fastBridge.hasRole(relayerRole, relayer));
        vm.expectRevert();
        fastBridge.revokeRole(relayerRole, relayer);
    }

    /// @notice Test to check if a guard can be successfully added
    function test_successfulAddGuard() public {
        vm.startPrank(owner);
        assertFalse(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        fastBridge.grantRole(fastBridge.GUARD_ROLE(), guard);
        assertTrue(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
    }

    /// @notice Test to check if only an admin can add a guard
    function test_onlyAdminCanAddGuard() public {
        bytes32 guardRole = fastBridge.GUARD_ROLE();
        vm.startPrank(guard);
        assertFalse(fastBridge.hasRole(guardRole, guard));
        vm.expectRevert();
        fastBridge.grantRole(guardRole, guard);
    }

    /// @notice Test to check if a guard can be successfully removed
    function test_successfulRemoveGuard() public {
        test_successfulAddGuard();
        assertTrue(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        vm.startPrank(owner);
        fastBridge.revokeRole(fastBridge.GUARD_ROLE(), guard);
        assertFalse(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
    }

    /// @notice Test to check if only an admin can remove a guard
    function test_onlyAdminCanRemoveGuard() public {
        bytes32 guardRole = fastBridge.GUARD_ROLE();
        test_successfulAddGuard();
        vm.startPrank(guard);
        assertTrue(fastBridge.hasRole(guardRole, guard));
        vm.expectRevert();
        fastBridge.revokeRole(guardRole, guard);
    }

    /// @notice Tests to check governor add, remove
    function test_successfulAddGovernor() public {
        vm.startPrank(owner);
        assertFalse(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        fastBridge.grantRole(fastBridge.GOVERNOR_ROLE(), governor);
        assertTrue(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
    }

    function test_onlyAdminCanAddGovernor() public {
        bytes32 governorRole = fastBridge.GOVERNOR_ROLE();
        vm.startPrank(governor);
        assertFalse(fastBridge.hasRole(governorRole, governor));
        vm.expectRevert();
        fastBridge.grantRole(governorRole, governor);
    }

    function test_successfulRemoveGovernor() public {
        test_successfulAddGovernor();
        assertTrue(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        vm.startPrank(owner);
        fastBridge.revokeRole(fastBridge.GOVERNOR_ROLE(), governor);
        assertFalse(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
    }

    function test_onlyAdminCanRemoveGovernor() public {
        bytes32 governorRole = fastBridge.GOVERNOR_ROLE();
        test_successfulAddGovernor();
        vm.startPrank(governor);
        assertTrue(fastBridge.hasRole(governorRole, governor));
        vm.expectRevert();
        fastBridge.revokeRole(governorRole, governor);
    }

    function test_successfulAddRefunder() public {
        vm.startPrank(owner);
        assertFalse(fastBridge.hasRole(fastBridge.REFUNDER_ROLE(), refunder));
        fastBridge.grantRole(fastBridge.REFUNDER_ROLE(), refunder);
        assertTrue(fastBridge.hasRole(fastBridge.REFUNDER_ROLE(), refunder));
    }

    function test_onlyAdminCanAddRefunder() public {
        bytes32 refunderRole = fastBridge.REFUNDER_ROLE();
        vm.startPrank(refunder);
        assertFalse(fastBridge.hasRole(refunderRole, refunder));
        vm.expectRevert();
        fastBridge.grantRole(refunderRole, refunder);
    }

    function test_successfulRemoveRefunder() public {
        test_successfulAddRefunder();
        assertTrue(fastBridge.hasRole(fastBridge.REFUNDER_ROLE(), refunder));
        vm.startPrank(owner);
        fastBridge.revokeRole(fastBridge.REFUNDER_ROLE(), refunder);
        assertFalse(fastBridge.hasRole(fastBridge.REFUNDER_ROLE(), refunder));
    }

    function test_onlyAdminCanRemoveRefunder() public {
        bytes32 refunderRole = fastBridge.REFUNDER_ROLE();
        test_successfulAddRefunder();
        vm.startPrank(refunder);
        assertTrue(fastBridge.hasRole(refunderRole, refunder));
        vm.expectRevert();
        fastBridge.revokeRole(refunderRole, refunder);
    }

    function test_successfulSetProtocolFeeRate() public {
        test_successfulAddGovernor();
        assertEq(fastBridge.protocolFeeRate(), 0);

        vm.startPrank(governor);
        uint256 protocolFeeRate = 0.001e6;
        fastBridge.setProtocolFeeRate(protocolFeeRate);
        assertEq(fastBridge.protocolFeeRate(), protocolFeeRate);
        vm.stopPrank();
    }

    function test_onlyGovernorCanSetProtocolFeeRate() public {
        test_successfulAddGovernor();
        uint256 protocolFeeRate = 0.001e6;
        vm.expectRevert();
        fastBridge.setProtocolFeeRate(protocolFeeRate);
    }

    function test_failedSetProtocolFeeRateWhenGreaterThanMax() public {
        test_successfulAddGovernor();
        uint256 protocolFeeRate = fastBridge.FEE_RATE_MAX() + 1;
        vm.startPrank(governor);
        vm.expectRevert("newFeeRate > max");
        fastBridge.setProtocolFeeRate(protocolFeeRate);
        vm.stopPrank();
    }

    // Tests to set chain gas amount
    function test_successfulSetChainGasAmount() public {
        test_successfulAddGovernor();
        assertEq(fastBridge.chainGasAmount(), 0);

        vm.startPrank(governor);
        uint256 chainGasAmount = 0.005e18;
        fastBridge.setChainGasAmount(chainGasAmount);
        assertEq(fastBridge.chainGasAmount(), chainGasAmount);
        vm.stopPrank();
    }

    function test_onlyGovernorCanSetChainGasAmount() public {
        test_successfulSetChainGasAmount();

        uint256 newChainGasAmount = 1e18;
        vm.expectRevert();
        fastBridge.setChainGasAmount(newChainGasAmount);
    }

    event BridgeRequested(
        bytes32 indexed transactionId,
        address indexed sender,
        bytes request,
        uint32 destChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        bool sendChainGas
    );

    // This test checks the successful execution of a bridge transaction
    function test_successfulBridge() public {
        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // get expected bridge request and tx id
        uint256 currentNonce = fastBridge.nonce();
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, currentNonce, 0);

        vm.expectEmit();
        emit BridgeRequested(
            transactionId,
            user,
            request,
            1,
            address(arbUSDC),
            address(ethUSDC),
            11 * 10 ** 6, // originAmount
            10.97e6, // destAmount
            false // sendChainGas
        );

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        fastBridge.bridge(params);
        // Check the state of the tokens after the bridge transaction
        // The fastBridge should have 11 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(address(fastBridge)), 11 * 10 ** 6);
        // The user should have 89 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(user), 89 * 10 ** 6);

        // Get the information of the bridge transaction
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulBridgeWithETH() public {
        // setup eth
        deal(user, 100 * 10 ** 18);
        uint256 userBalanceBefore = user.balance;
        uint256 bridgeBalanceBefore = address(fastBridge).balance;

        // Start a prank with the user
        vm.startPrank(user);

        // get expected bridge request and tx id
        uint256 currentNonce = fastBridge.nonce();
        (, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(block.chainid, currentNonce, 0);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: UniversalTokenLib.ETH_ADDRESS,
            destToken: UniversalTokenLib.ETH_ADDRESS,
            originAmount: 11 * 10 ** 18,
            destAmount: 10.97e18,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        fastBridge.bridge{value: params.originAmount}(params);

        // Check the state of the tokens after the bridge transaction
        uint256 userBalanceAfter = user.balance;
        uint256 bridgeBalanceAfter = address(fastBridge).balance;

        assertEq(userBalanceBefore - userBalanceAfter, 11 * 10 ** 18);
        assertEq(bridgeBalanceAfter - bridgeBalanceBefore, 11 * 10 ** 18);

        // Get the information of the bridge transaction
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulBridgeWithProtocolFeeOn() public {
        setUpRoles();

        // set protocol fee with governor to 10 bps
        vm.startPrank(governor);
        uint256 protocolFeeRate = 0.001e6;
        fastBridge.setProtocolFeeRate(protocolFeeRate);
        assertEq(fastBridge.protocolFeeRate(), protocolFeeRate);
        vm.stopPrank();

        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // get expected bridge request and tx id
        uint256 currentNonce = fastBridge.nonce();
        (bytes memory request, bytes32 transactionId) =
            _getBridgeRequestAndId(block.chainid, currentNonce, protocolFeeRate);

        vm.expectEmit();
        emit BridgeRequested(
            transactionId,
            user,
            request,
            1,
            address(arbUSDC),
            address(ethUSDC),
            10.989e6, // originAmount
            10.97e6, // destAmount
            false // sendChainGas
        );

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        fastBridge.bridge(params);
        // Check the state of the tokens after the bridge transaction
        // The fastBridge should have 11 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(address(fastBridge)), 11 * 10 ** 6);
        // The user should have 89 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(user), 89 * 10 ** 6);

        // Get the information of the bridge transaction
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulBridgeWithETHAndProtocolFeeOn() public {
        setUpRoles();

        // set protocol fee with governor to 10 bps
        vm.startPrank(governor);
        uint256 protocolFeeRate = 0.001e6;
        fastBridge.setProtocolFeeRate(protocolFeeRate);
        assertEq(fastBridge.protocolFeeRate(), protocolFeeRate);
        vm.stopPrank();

        // setup eth
        deal(user, 100 * 10 ** 18);
        uint256 userBalanceBefore = user.balance;
        uint256 bridgeBalanceBefore = address(fastBridge).balance;

        // Start a prank with the user
        vm.startPrank(user);

        // get expected bridge request and tx id
        uint256 currentNonce = fastBridge.nonce();
        (bytes memory request, bytes32 transactionId) =
            _getBridgeRequestAndIdWithETH(block.chainid, currentNonce, protocolFeeRate);

        vm.expectEmit();
        emit BridgeRequested(
            transactionId,
            user,
            request,
            1,
            UniversalTokenLib.ETH_ADDRESS,
            UniversalTokenLib.ETH_ADDRESS,
            10.989e18, // originAmount
            10.97e18, // destAmount
            false // sendChainGas
        );

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: UniversalTokenLib.ETH_ADDRESS,
            destToken: UniversalTokenLib.ETH_ADDRESS,
            originAmount: 11 * 10 ** 18,
            destAmount: 10.97e18,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        fastBridge.bridge{value: params.originAmount}(params);

        // Check the state of the tokens after the bridge transaction
        uint256 userBalanceAfter = user.balance;
        uint256 bridgeBalanceAfter = address(fastBridge).balance;

        assertEq(userBalanceBefore - userBalanceAfter, 11 * 10 ** 18);
        assertEq(bridgeBalanceAfter - bridgeBalanceBefore, 11 * 10 ** 18);

        // Get the information of the bridge transaction
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulBridgeWithChainGas() public {
        setUpRoles();

        // Start prank with governor to set chain gas (irrelevant for this test on origin)
        vm.startPrank(governor);
        uint256 chainGasAmount = 0.005e18;
        fastBridge.setChainGasAmount(chainGasAmount);
        assertEq(fastBridge.chainGasAmount(), chainGasAmount);
        vm.stopPrank();

        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // get expected bridge request and tx id
        uint256 currentNonce = fastBridge.nonce();
        (bytes memory request, bytes32 transactionId) =
            _getBridgeRequestAndIdWithChainGas(block.chainid, currentNonce, 0);

        vm.expectEmit();
        emit BridgeRequested(
            transactionId,
            user,
            request,
            1,
            address(arbUSDC),
            address(ethUSDC),
            11 * 10 ** 6, // originAmount
            10.97e6, // destAmount
            true // sendChainGas
        );

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: true,
            deadline: block.timestamp + TX_DEADLINE
        });
        fastBridge.bridge(params);
        // Check the state of the tokens after the bridge transaction
        // The fastBridge should have 11 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(address(fastBridge)), 11 * 10 ** 6);
        // The user should have 89 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(user), 89 * 10 ** 6);

        // Get the information of the bridge transaction
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulBridgeWithETHAndChainGas() public {
        setUpRoles();

        // setup eth
        deal(user, 100 * 10 ** 18);
        uint256 userBalanceBefore = user.balance;
        uint256 bridgeBalanceBefore = address(fastBridge).balance;

        // Start prank with governor to set chain gas (irrelevant for this test on origin)
        vm.startPrank(governor);
        uint256 chainGasAmount = 0.005e18;
        fastBridge.setChainGasAmount(chainGasAmount);
        assertEq(fastBridge.chainGasAmount(), chainGasAmount);
        vm.stopPrank();

        // Start a prank with the user
        vm.startPrank(user);

        // get expected bridge request and tx id
        uint256 currentNonce = fastBridge.nonce();
        (, bytes32 transactionId) = _getBridgeRequestAndIdWithETHAndChainGas(block.chainid, currentNonce, 0);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: UniversalTokenLib.ETH_ADDRESS,
            destToken: UniversalTokenLib.ETH_ADDRESS,
            originAmount: 11 * 10 ** 18,
            destAmount: 10.97e18,
            sendChainGas: true,
            deadline: block.timestamp + TX_DEADLINE
        });
        fastBridge.bridge{value: params.originAmount}(params);

        // Check the state of the tokens after the bridge transaction
        uint256 userBalanceAfter = user.balance;
        uint256 bridgeBalanceAfter = address(fastBridge).balance;

        assertEq(userBalanceBefore - userBalanceAfter, 11 * 10 ** 18);
        assertEq(bridgeBalanceAfter - bridgeBalanceBefore, 11 * 10 ** 18);

        // Get the information of the bridge transaction
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulBridgeWithSenderNotUser() public {
        // Start a prank with the user
        vm.startPrank(dstUser);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // get expected bridge request and tx id
        uint256 currentNonce = fastBridge.nonce();
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, currentNonce, 0);

        vm.expectEmit();
        emit BridgeRequested(
            transactionId,
            user,
            request,
            1,
            address(arbUSDC),
            address(ethUSDC),
            11 * 10 ** 6, // originAmount
            10.97e6, // destAmount
            false // sendChainGas
        );

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        fastBridge.bridge(params);
        // Check the state of the tokens after the bridge transaction
        // The fastBridge should have 11 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(address(fastBridge)), 11 * 10 ** 6);
        // The user should have 89 * 10 ** 6 of arbUSDC
        assertEq(arbUSDC.balanceOf(dstUser), 89 * 10 ** 6);

        // Get the information of the bridge transaction
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedBridgeSameChainId() public {
        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: uint32(block.chainid),
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        vm.expectRevert(abi.encodeWithSelector(ChainIncorrect.selector));
        fastBridge.bridge(params);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedBridgeOriginAmountZero() public {
        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 0,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        vm.expectRevert(abi.encodeWithSelector(AmountIncorrect.selector));
        fastBridge.bridge(params);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedBridgeDestAmountZero() public {
        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 0,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        vm.expectRevert(abi.encodeWithSelector(AmountIncorrect.selector));
        fastBridge.bridge(params);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedBridgeOriginTokenZero() public {
        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(0),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        vm.expectRevert(abi.encodeWithSelector(ZeroAddress.selector));
        fastBridge.bridge(params);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedBridgeDestTokenZero() public {
        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(0),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + TX_DEADLINE
        });
        vm.expectRevert(abi.encodeWithSelector(ZeroAddress.selector));
        fastBridge.bridge(params);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedBridgeDeadlineTooShort() public {
        // Start a prank with the user
        vm.startPrank(user);
        // Approve the fastBridge to spend 100 * 10 ** 6 of arbUSDC from the user
        arbUSDC.approve(address(fastBridge), 100 * 10 ** 6);

        // Execute the bridge transaction
        IFastBridge.BridgeParams memory params = IFastBridge.BridgeParams({
            dstChainId: 1,
            sender: user,
            to: user,
            originToken: address(arbUSDC),
            destToken: address(ethUSDC),
            originAmount: 11 * 10 ** 6,
            destAmount: 10.97e6,
            sendChainGas: false,
            deadline: block.timestamp + 1800 - 1
        });
        vm.expectRevert(abi.encodeWithSelector(DeadlineTooShort.selector));
        fastBridge.bridge(params);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    event BridgeRelayed(
        bytes32 indexed transactionId,
        address indexed relayer,
        address indexed to,
        uint32 originChainId,
        address originToken,
        address destToken,
        uint256 originAmount,
        uint256 destAmount,
        uint256 chainGasAmount
    );

    // This test checks the successful relaying of a destination bridge
    function test_successfulRelayDestination() public {
        // Set up the roles for the test
        setUpRoles();

        // Start prank with governor to set chain gas on dest chain
        vm.startPrank(governor);
        uint256 chainGasAmount = 0.005e18;
        fastBridge.setChainGasAmount(chainGasAmount);
        assertEq(fastBridge.chainGasAmount(), chainGasAmount);
        vm.stopPrank();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(42_161, 0, 0);

        // Get the initial information of the bridge transaction; make sure not relayed
        assertEq(fastBridge.bridgeRelays(transactionId), false);

        // Start a prank with the relayer
        vm.startPrank(relayer);
        // Approve the fastBridge to spend the maximum amount of ethUSDC from the relayer
        ethUSDC.approve(address(fastBridge), type(uint256).max);
        // Check the initial balances of the relayer and the user
        assertEq(ethUSDC.balanceOf(relayer), 100 * 10 ** 6);
        assertEq(ethUSDC.balanceOf(user), 100 * 10 ** 6);
        // Expect the BridgeRelayed event to be emitted
        vm.expectEmit();
        emit BridgeRelayed(
            transactionId, relayer, user, 42_161, address(arbUSDC), address(ethUSDC), 11 * 10 ** 6, 10.97e6, 0
        );
        // Expect not doing any calls to user address
        vm.expectCall(user, "", 0);
        // Relay the destination bridge
        vm.chainId(1); // set to dest chain
        fastBridge.relay(request);
        // Check the balances of the relayer and the user after relaying the destination bridge
        assertEq(ethUSDC.balanceOf(relayer), 89.03e6);
        assertEq(ethUSDC.balanceOf(user), 110.97e6);

        // Get the returned information of the bridge transaction relays status
        assertEq(fastBridge.bridgeRelays(transactionId), true);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRelayDestinationWithETH() public {
        // Set up the roles for the test
        setUpRoles();

        // deal some dest ETH to relayer
        deal(relayer, 100e18);

        // Start prank with governor to set chain gas on dest chain
        vm.startPrank(governor);
        uint256 chainGasAmount = 0.005e18;
        fastBridge.setChainGasAmount(chainGasAmount);
        assertEq(fastBridge.chainGasAmount(), chainGasAmount);
        vm.stopPrank();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(42_161, 0, 0);

        // Get the initial information of the bridge transaction; make sure not relayed
        assertEq(fastBridge.bridgeRelays(transactionId), false);

        // Start a prank with the relayer
        vm.startPrank(relayer);
        // Get the initial balances of the relayer and the user
        uint256 userBalanceBefore = user.balance;
        uint256 relayerBalanceBefore = relayer.balance;

        // Expect the BridgeRelayed event to be emitted
        vm.expectEmit();
        emit BridgeRelayed(
            transactionId,
            relayer,
            user,
            42_161,
            UniversalTokenLib.ETH_ADDRESS,
            UniversalTokenLib.ETH_ADDRESS,
            11 * 10 ** 18,
            10.97e18,
            0
        );
        // Expect exactly one call to user address
        vm.expectCall(user, "", 1);

        // Relay the destination bridge
        vm.chainId(1); // set to dest chain
        uint256 value = 10.97e18;
        fastBridge.relay{value: value}(request);

        // Check the balances of the relayer and the user after relaying the destination bridge
        uint256 userBalanceAfter = user.balance;
        uint256 relayerBalanceAfter = relayer.balance;

        assertEq(userBalanceAfter - userBalanceBefore, value);
        assertEq(relayerBalanceBefore - relayerBalanceAfter, value);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRelayDestinationWithETHAndChainGas() public {
        // Set up the roles for the test
        setUpRoles();

        // deal some dest ETH to relayer
        deal(relayer, 100e18);

        // Start prank with governor to set chain gas on dest chain
        vm.startPrank(governor);
        uint256 chainGasAmount = 0.005e18;
        fastBridge.setChainGasAmount(chainGasAmount);
        assertEq(fastBridge.chainGasAmount(), chainGasAmount);
        vm.stopPrank();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETHAndChainGas(42_161, 0, 0);

        // Get the initial information of the bridge transaction; make sure not relayed
        assertEq(fastBridge.bridgeRelays(transactionId), false);

        // Start a prank with the relayer
        vm.startPrank(relayer);
        // Get the initial balances of the relayer and the user
        uint256 userBalanceBefore = user.balance;
        uint256 relayerBalanceBefore = relayer.balance;

        // Expect the BridgeRelayed event to be emitted
        vm.expectEmit();
        emit BridgeRelayed(
            transactionId,
            relayer,
            user,
            42_161,
            UniversalTokenLib.ETH_ADDRESS,
            UniversalTokenLib.ETH_ADDRESS,
            11 * 10 ** 18,
            10.97e18,
            0.005e18
        );
        // Expect exactly one call to user address
        vm.expectCall(user, "", 1);

        // Relay the destination bridge
        vm.chainId(1); // set to dest chain
        uint256 value = 10.975e18;
        fastBridge.relay{value: value}(request);

        // Check the balances of the relayer and the user after relaying the destination bridge
        uint256 userBalanceAfter = user.balance;
        uint256 relayerBalanceAfter = relayer.balance;

        assertEq(userBalanceAfter - userBalanceBefore, value);
        assertEq(relayerBalanceBefore - relayerBalanceAfter, value);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRelayDestinationWithChainGas() public {
        // Set up the roles for the test
        setUpRoles();

        // deal some dest ETH to relayer
        deal(relayer, 100e18);

        // Start prank with governor to set chain gas on dest chain
        vm.startPrank(governor);
        uint256 chainGasAmount = 0.005e18;
        fastBridge.setChainGasAmount(chainGasAmount);
        assertEq(fastBridge.chainGasAmount(), chainGasAmount);
        vm.stopPrank();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithChainGas(42_161, 0, 0);

        // Get the initial information of the bridge transaction; make sure not relayed
        assertEq(fastBridge.bridgeRelays(transactionId), false);

        // Start a prank with the relayer
        vm.startPrank(relayer);
        // Approve the fastBridge to spend the maximum amount of ethUSDC from the relayer
        ethUSDC.approve(address(fastBridge), type(uint256).max);
        // Check the initial balances of the relayer and the user
        assertEq(ethUSDC.balanceOf(relayer), 100 * 10 ** 6);
        assertEq(ethUSDC.balanceOf(user), 100 * 10 ** 6);
        // Expect the BridgeRelayed event to be emitted
        vm.expectEmit();
        emit BridgeRelayed(
            transactionId, relayer, user, 42_161, address(arbUSDC), address(ethUSDC), 11 * 10 ** 6, 10.97e6, 0.005e18
        );
        // Expect exactly one call to user address
        vm.expectCall(user, "", 1);
        // Relay the destination bridge
        vm.chainId(1); // set to dest chain
        fastBridge.relay{value: chainGasAmount}(request);
        // Check the balances of the relayer and the user after relaying the destination bridge
        assertEq(ethUSDC.balanceOf(relayer), 89.03e6);
        assertEq(ethUSDC.balanceOf(user), 110.97e6);
        assertEq(relayer.balance, 99.995e18);
        assertEq(user.balance, 0.005e18);

        // Get the returned information of the bridge transaction relays status
        assertEq(fastBridge.bridgeRelays(transactionId), true);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRelayDestinationWithChainGasAmountZero() public {
        // Set up the roles for the test
        setUpRoles();

        // deal some dest ETH to relayer
        deal(relayer, 100e18);

        // get bridge request and tx id
        // chain gas param should be true but we forward 0 amount since not set by governor
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithChainGas(42_161, 0, 0);

        // Get the initial information of the bridge transaction; make sure not relayed
        assertEq(fastBridge.bridgeRelays(transactionId), false);

        // Start a prank with the relayer
        vm.startPrank(relayer);
        // Approve the fastBridge to spend the maximum amount of ethUSDC from the relayer
        ethUSDC.approve(address(fastBridge), type(uint256).max);
        // Check the initial balances of the relayer and the user
        assertEq(ethUSDC.balanceOf(relayer), 100 * 10 ** 6);
        assertEq(ethUSDC.balanceOf(user), 100 * 10 ** 6);
        // Expect the BridgeRelayed event to be emitted
        vm.expectEmit();
        emit BridgeRelayed(
            transactionId, relayer, user, 42_161, address(arbUSDC), address(ethUSDC), 11 * 10 ** 6, 10.97e6, 0
        );
        // Expect not doing any calls to user address
        vm.expectCall(user, "", 0);
        // Relay the destination bridge
        vm.chainId(1); // set to dest chain
        fastBridge.relay(request);
        // Check the balances of the relayer and the user after relaying the destination bridge
        assertEq(ethUSDC.balanceOf(relayer), 89.03e6);
        assertEq(ethUSDC.balanceOf(user), 110.97e6);
        assertEq(relayer.balance, 100e18);
        assertEq(user.balance, 0);

        // Get the returned information of the bridge transaction relays status
        assertEq(fastBridge.bridgeRelays(transactionId), true);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedRelayNotDestChain() public {
        // Set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(42_161, 0, 0);

        // Start a prank with the relayer
        vm.startPrank(relayer);

        // Approve the fastBridge to spend the maximum amount of ethUSDC from the relayer
        ethUSDC.approve(address(fastBridge), type(uint256).max);

        // Relay the destination bridge
        vm.expectRevert(abi.encodeWithSelector(ChainIncorrect.selector));
        vm.chainId(2); // wrong dest chain id
        fastBridge.relay(request);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedRelayTimeExceeded() public {
        // Set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(42_161, 0, 0);

        // Start a prank with the relayer
        vm.startPrank(relayer);

        // Approve the fastBridge to spend the maximum amount of ethUSDC from the relayer
        ethUSDC.approve(address(fastBridge), type(uint256).max);

        // deadline of 60 min on relay
        vm.warp(block.timestamp + 61 minutes);

        // Relay the destination bridge
        vm.expectRevert(abi.encodeWithSelector(DeadlineExceeded.selector));
        vm.chainId(1); // set to dest chain
        fastBridge.relay(request);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    // This test checks if the destination bridge has already been relayed
    function test_alreadyRelayedDestination() public {
        // First, we successfully relay the destination
        test_successfulRelayDestination();

        // Then, we set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(42_161, 0, 0);
        assertEq(fastBridge.bridgeRelays(transactionId), true);

        // We start a prank with the relayer
        vm.startPrank(relayer);
        // We expect a revert because the destination bridge has already been relayed
        vm.expectRevert(abi.encodeWithSelector(TransactionRelayed.selector));
        vm.chainId(1); // set to dest chain
        // We try to relay the destination bridge again
        fastBridge.relay(request);
        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedRelayNotRelayer() public virtual {
        // Set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(42_161, 0, 0);

        // Start a prank with the relayer
        vm.startPrank(guard);
        // Approve the fastBridge to spend the maximum amount of ethUSDC from the relayer
        ethUSDC.approve(address(fastBridge), type(uint256).max);

        // Relay the destination bridge
        vm.chainId(1); // set to dest chain
        expectUnauthorized(guard, fastBridge.RELAYER_ROLE());
        fastBridge.relay(request);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash);

    // This test checks the successful provision of relay proof
    function test_successfulRelayProof() public {
        // First, we successfully initiate the original bridge tx
        test_successfulBridge();

        // Then, we set up the roles for the test
        setUpRoles();

        // We start a prank with the relayer
        vm.startPrank(relayer);

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // We define a fake transaction hash to be the proof from dest chain
        bytes32 fakeTxnHash = bytes32("0x01");
        // We expect an event to be emitted
        vm.expectEmit();
        // We emit the BridgeProofProvided event to test again
        emit BridgeProofProvided(transactionId, relayer, fakeTxnHash);

        // We provide the relay proof
        fastBridge.prove(request, fakeTxnHash);

        // We check if the bridge transaction proof timestamp is set to the timestamp at which the proof was provided
        assertCorrectProof(transactionId, block.timestamp, relayer);

        // We check if the bridge status is RELAYER_PROVED
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.RELAYER_PROVED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulProveWithProofTimestampOverflow() public {
        // sets block timestamp to just before overflow of uint96
        vm.warp(uint256(type(uint96).max) + 1 minutes);

        // First, we successfully initiate the original bridge tx
        test_successfulBridge();

        // Then, we set up the roles for the test
        setUpRoles();

        // We start a prank with the relayer
        vm.startPrank(relayer);

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // We define a fake transaction hash to be the proof from dest chain
        bytes32 fakeTxnHash = bytes32("0x01");

        // We provide the relay proof
        fastBridge.prove(request, fakeTxnHash);

        // We check if the bridge transaction proof timestamp is set to the timestamp at which the proof was provided
        assertCorrectProof(transactionId, block.timestamp, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_proveWithHugeDelay() public {
        // First, we successfully initiate the original bridge tx
        test_successfulBridge();

        // Then, we set up the roles for the test
        setUpRoles();

        // We start a prank with the relayer
        vm.startPrank(relayer);

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // We define a fake transaction hash to be the proof from dest chain
        bytes32 fakeTxnHash = bytes32("0x01");

        // At this point the permissionless refund is available
        skip(30 days);

        // We provide the relay proof
        fastBridge.prove(request, fakeTxnHash);

        // We check if the bridge transaction proof timestamp is set to the timestamp at which the proof was provided
        assertCorrectProof(transactionId, block.timestamp, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedProveNotRequested() public {
        // Then, we set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // We start a prank with the relayer
        vm.startPrank(relayer);

        // We provide the relay proof
        vm.expectRevert(abi.encodeWithSelector(StatusIncorrect.selector));
        fastBridge.prove(request, bytes32("0x01"));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedProveNotRelayer() public {
        // First, we successfully initiate the original bridge tx
        test_successfulBridge();

        // Then, we set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // We provide the relay proof
        expectUnauthorized(address(this), fastBridge.RELAYER_ROLE());
        fastBridge.prove(request, bytes32("0x01"));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    event BridgeDepositClaimed(
        bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount
    );

    function test_successfulClaimOriginTokens() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.warp(block.timestamp + 31 minutes);

        vm.expectEmit();
        emit BridgeDepositClaimed(transactionId, relayer, relayer, address(arbUSDC), 11 * 10 ** 6);

        uint256 preClaimBalanceRelayer = arbUSDC.balanceOf(relayer);
        uint256 preClaimBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        fastBridge.claim(request, relayer);

        // check balance changes
        uint256 postClaimBalanceRelayer = arbUSDC.balanceOf(relayer);
        uint256 postClaimBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        assertEq(postClaimBalanceRelayer - preClaimBalanceRelayer, 11 * 10 ** 6);
        assertEq(preClaimBalanceBridge - postClaimBalanceBridge, 11 * 10 ** 6);

        // check status changed
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.RELAYER_CLAIMED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulClaimWithETH() public {
        setUpRoles();
        test_successfulBridgeWithETH();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndIdWithETH(block.chainid, 0, 0);

        vm.startPrank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.warp(block.timestamp + 31 minutes);

        uint256 preClaimBalanceRelayer = relayer.balance;
        uint256 preClaimBalanceBridge = address(fastBridge).balance;

        fastBridge.claim(request, relayer);

        // check balance changes
        uint256 postClaimBalanceRelayer = relayer.balance;
        uint256 postClaimBalanceBridge = address(fastBridge).balance;

        assertEq(postClaimBalanceRelayer - preClaimBalanceRelayer, 11 * 10 ** 18);
        assertEq(preClaimBalanceBridge - postClaimBalanceBridge, 11 * 10 ** 18);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulClaimOriginTokensWithProtocolFeeOn() public {
        setUpRoles();
        test_successfulBridgeWithProtocolFeeOn();

        // get bridge request and tx id
        uint256 protocolFeeRate = fastBridge.protocolFeeRate();
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, protocolFeeRate);

        vm.startPrank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.warp(block.timestamp + 31 minutes);

        uint256 amountToProtocol = ((11 * 10 ** 6) * protocolFeeRate) / 1e6;
        uint256 amountClaimed = 11 * 10 ** 6 - amountToProtocol;

        vm.expectEmit();
        emit BridgeDepositClaimed(transactionId, relayer, relayer, address(arbUSDC), amountClaimed);

        uint256 preClaimBalanceRelayer = arbUSDC.balanceOf(relayer);
        uint256 preClaimBalanceBridge = arbUSDC.balanceOf(address(fastBridge));
        uint256 preClaimProtocolFees = fastBridge.protocolFees(address(arbUSDC));

        fastBridge.claim(request, relayer);

        // check balance changes
        uint256 postClaimBalanceRelayer = arbUSDC.balanceOf(relayer);
        uint256 postClaimBalanceBridge = arbUSDC.balanceOf(address(fastBridge));
        uint256 postClaimProtocolFees = fastBridge.protocolFees(address(arbUSDC));

        assertEq(postClaimBalanceRelayer - preClaimBalanceRelayer, amountClaimed);
        assertEq(preClaimBalanceBridge - postClaimBalanceBridge, amountClaimed);

        assertEq(postClaimProtocolFees - preClaimProtocolFees, amountToProtocol);
        assertEq(postClaimBalanceBridge, amountToProtocol); // @dev assumes only one bridge tx in test

        // check status changed
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.RELAYER_CLAIMED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulClaimWithETHAndProtocolFeeOn() public {
        setUpRoles();
        test_successfulBridgeWithETHAndProtocolFeeOn();

        // get bridge request and tx id
        uint256 protocolFeeRate = fastBridge.protocolFeeRate();
        (bytes memory request,) = _getBridgeRequestAndIdWithETH(block.chainid, 0, protocolFeeRate);

        vm.startPrank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.warp(block.timestamp + 31 minutes);

        uint256 amountToProtocol = ((11 * 10 ** 18) * protocolFeeRate) / 1e6;
        uint256 amountClaimed = 11 * 10 ** 18 - amountToProtocol;

        uint256 preClaimBalanceRelayer = relayer.balance;
        uint256 preClaimBalanceBridge = address(fastBridge).balance;
        uint256 preClaimProtocolFees = fastBridge.protocolFees(UniversalTokenLib.ETH_ADDRESS);

        fastBridge.claim(request, relayer);

        // check balance changes
        uint256 postClaimBalanceRelayer = relayer.balance;
        uint256 postClaimBalanceBridge = address(fastBridge).balance;
        uint256 postClaimProtocolFees = fastBridge.protocolFees(UniversalTokenLib.ETH_ADDRESS);

        assertEq(postClaimBalanceRelayer - preClaimBalanceRelayer, amountClaimed);
        assertEq(preClaimBalanceBridge - postClaimBalanceBridge, amountClaimed);

        assertEq(postClaimProtocolFees - preClaimProtocolFees, amountToProtocol);
        assertEq(postClaimBalanceBridge, amountToProtocol); // @dev assumes only one bridge tx in test

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulClaimWithProofTimestampOverflow() public {
        // sets block timestamp to just before overflow of uint96
        vm.warp(type(uint96).max - 1 minutes);

        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.warp(block.timestamp + 31 minutes);

        fastBridge.claim(request, relayer);

        // check status changed
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.RELAYER_CLAIMED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedClaimNoProof() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(relayer);

        vm.warp(block.timestamp + 31 minutes);

        vm.expectRevert(abi.encodeWithSelector(StatusIncorrect.selector));
        fastBridge.claim(request, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedClaimNotOldRelayer() public virtual {
        setUpRoles();
        test_successfulBridge();

        vm.startPrank(owner);
        fastBridge.grantRole(fastBridge.RELAYER_ROLE(), address(this));
        vm.stopPrank();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.warp(block.timestamp + 31 minutes);

        vm.prank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.expectRevert(abi.encodeWithSelector(SenderIncorrect.selector));
        fastBridge.claim(request, relayer);
    }

    function test_failedClaimNotEnoughTime() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.expectRevert(abi.encodeWithSelector(DisputePeriodNotPassed.selector));
        fastBridge.claim(request, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedClaimNotRelayer() public virtual {
        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.warp(block.timestamp + 31 minutes);

        expectUnauthorized(address(this), fastBridge.RELAYER_ROLE());
        fastBridge.claim(request, relayer);
    }

    event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer);

    function test_successfulDisputeProof() public {
        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(guard);

        vm.expectEmit();
        emit BridgeProofDisputed(transactionId, guard);

        fastBridge.dispute(transactionId);

        // check status and proofs updated
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));
        assertCorrectProof(transactionId, 0, address(0));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulDisputeProofWithProofTimestampOverflow() public {
        // sets block timestamp to just before overflow of uint96
        vm.warp(type(uint96).max - 1 minutes);

        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.warp(block.timestamp + 25 minutes);

        vm.startPrank(guard);

        fastBridge.dispute(transactionId);

        // check status and proofs updated
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));
        assertCorrectProof(transactionId, 0, address(0));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedDisputeNoProof() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(guard);

        vm.expectRevert(abi.encodeWithSelector(StatusIncorrect.selector));
        fastBridge.dispute(transactionId);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedDisputeEnoughTime() public {
        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(guard);

        vm.warp(block.timestamp + 31 minutes);

        vm.expectRevert(abi.encodeWithSelector(DisputePeriodPassed.selector));
        fastBridge.dispute(transactionId);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedDisputeNotGuard() public {
        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        expectUnauthorized(address(this), fastBridge.GUARD_ROLE());
        fastBridge.dispute(transactionId);
    }

    event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount);

    function test_successfulRefund() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(refunder);

        skip(TX_DEADLINE + 1 seconds);

        vm.expectEmit();
        emit BridgeDepositRefunded(transactionId, user, address(arbUSDC), 11 * 10 ** 6);

        uint256 preRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 preRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        fastBridge.refund(request);

        // check balance changes
        uint256 postRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 postRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        assertEq(postRefundBalanceUser - preRefundBalanceUser, 11 * 10 ** 6);
        assertEq(preRefundBalanceBridge - postRefundBalanceBridge, 11 * 10 ** 6);

        // check bridge status updated
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REFUNDED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRefundWithETH() public {
        setUpRoles();
        test_successfulBridgeWithETH();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndIdWithETH(block.chainid, 0, 0);

        vm.startPrank(refunder);

        skip(TX_DEADLINE + 1 seconds);

        uint256 preRefundBalanceUser = user.balance;
        uint256 preRefundBalanceBridge = address(fastBridge).balance;

        fastBridge.refund(request);

        // check balance changes
        uint256 postRefundBalanceUser = user.balance;
        uint256 postRefundBalanceBridge = address(fastBridge).balance;

        assertEq(postRefundBalanceUser - preRefundBalanceUser, 11 * 10 ** 18);
        assertEq(preRefundBalanceBridge - postRefundBalanceBridge, 11 * 10 ** 18);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRefundWithProtocolFee() public {
        setUpRoles();
        test_successfulBridgeWithProtocolFeeOn();

        // get bridge request and tx id
        uint256 protocolFeeRate = fastBridge.protocolFeeRate();
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, protocolFeeRate);

        vm.startPrank(refunder);

        skip(TX_DEADLINE + 1 seconds);

        vm.expectEmit();
        emit BridgeDepositRefunded(transactionId, user, address(arbUSDC), 11 * 10 ** 6);

        uint256 preRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 preRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        fastBridge.refund(request);

        // check balance changes
        uint256 postRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 postRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        assertEq(postRefundBalanceUser - preRefundBalanceUser, 11 * 10 ** 6);
        assertEq(preRefundBalanceBridge - postRefundBalanceBridge, 11 * 10 ** 6);

        // check bridge status updated
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REFUNDED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRefundWithETHAndProtocolFeeOn() public {
        setUpRoles();
        test_successfulBridgeWithETHAndProtocolFeeOn();

        // get bridge request and tx id
        uint256 protocolFeeRate = fastBridge.protocolFeeRate();
        (bytes memory request,) = _getBridgeRequestAndIdWithETH(block.chainid, 0, protocolFeeRate);

        vm.startPrank(refunder);

        skip(TX_DEADLINE + 1 seconds);

        uint256 preRefundBalanceUser = user.balance;
        uint256 preRefundBalanceBridge = address(fastBridge).balance;

        fastBridge.refund(request);

        // check balance changes
        uint256 postRefundBalanceUser = user.balance;
        uint256 postRefundBalanceBridge = address(fastBridge).balance;

        assertEq(postRefundBalanceUser - preRefundBalanceUser, 11 * 10 ** 18);
        assertEq(preRefundBalanceBridge - postRefundBalanceBridge, 11 * 10 ** 18);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRefundWithSenderNotUser() public {
        setUpRoles();
        test_successfulBridgeWithSenderNotUser();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(refunder);

        skip(TX_DEADLINE + 1 seconds);

        vm.expectEmit();
        emit BridgeDepositRefunded(transactionId, user, address(arbUSDC), 11 * 10 ** 6);

        uint256 preRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 preRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        fastBridge.refund(request);

        // check balance changes
        uint256 postRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 postRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        assertEq(postRefundBalanceUser - preRefundBalanceUser, 11 * 10 ** 6);
        assertEq(preRefundBalanceBridge - postRefundBalanceBridge, 11 * 10 ** 6);

        // check bridge status updated
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REFUNDED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedRefundNotEnoughTime() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(refunder);

        // user relay deadline is 60 minutes
        skip(TX_DEADLINE);

        vm.expectRevert(abi.encodeWithSelector(DeadlineNotExceeded.selector));
        fastBridge.refund(request);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulRefundNotRefunder() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(user);

        // Should be able to refund permissionlessly 7 days after deadline
        skip(TX_DEADLINE + 7 days + 1 seconds);

        vm.expectEmit();
        emit BridgeDepositRefunded(transactionId, user, address(arbUSDC), 11 * 10 ** 6);

        uint256 preRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 preRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        fastBridge.refund(request);

        // check balance changes
        uint256 postRefundBalanceUser = arbUSDC.balanceOf(user);
        uint256 postRefundBalanceBridge = arbUSDC.balanceOf(address(fastBridge));

        assertEq(postRefundBalanceUser - preRefundBalanceUser, 11 * 10 ** 6);
        assertEq(preRefundBalanceBridge - postRefundBalanceBridge, 11 * 10 ** 6);

        // check bridge status updated
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REFUNDED));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedRefundNotRefunderNotEnoughTime() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(user);

        // Permissionless refund is available after 7 days
        skip(TX_DEADLINE + 7 days);

        vm.expectRevert(abi.encodeWithSelector(DeadlineNotExceeded.selector));
        fastBridge.refund(request);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedRefundNoBridge() public {
        setUpRoles();

        vm.startPrank(refunder);

        // get bridge request and tx id
        (bytes memory request,) = _getBridgeRequestAndId(block.chainid, 0, 0);

        skip(TX_DEADLINE + 1 seconds);

        vm.expectRevert(abi.encodeWithSelector(StatusIncorrect.selector));
        fastBridge.refund(request);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulSweepProtocolFees() public {
        setUpRoles();
        test_successfulClaimOriginTokensWithProtocolFeeOn();
        assertTrue(fastBridge.protocolFees(address(arbUSDC)) > 0);

        vm.startPrank(governor);

        uint256 preSweepProtocolFees = fastBridge.protocolFees(address(arbUSDC));
        uint256 preSweepBalanceUser = arbUSDC.balanceOf(user);

        fastBridge.sweepProtocolFees(address(arbUSDC), user);

        uint256 postSweepProtocolFees = fastBridge.protocolFees(address(arbUSDC));
        uint256 postSweepBalanceUser = arbUSDC.balanceOf(user);

        assertEq(postSweepProtocolFees, 0);
        assertEq(postSweepBalanceUser - preSweepBalanceUser, preSweepProtocolFees);

        vm.stopPrank();
    }

    function test_successfulSweepProtocolFeesWhenETH() public {
        setUpRoles();
        test_successfulClaimWithETHAndProtocolFeeOn();
        assertTrue(fastBridge.protocolFees(UniversalTokenLib.ETH_ADDRESS) > 0);

        vm.startPrank(governor);

        uint256 preSweepProtocolFees = fastBridge.protocolFees(UniversalTokenLib.ETH_ADDRESS);
        uint256 preSweepBalanceUser = user.balance;

        fastBridge.sweepProtocolFees(UniversalTokenLib.ETH_ADDRESS, user);

        uint256 postSweepProtocolFees = fastBridge.protocolFees(UniversalTokenLib.ETH_ADDRESS);
        uint256 postSweepBalanceUser = user.balance;

        assertEq(postSweepProtocolFees, 0);
        assertEq(postSweepBalanceUser - preSweepBalanceUser, preSweepProtocolFees);

        vm.stopPrank();
    }

    function test_failedSweepProtocolFeesWhenNotGovernor() public {
        setUpRoles();
        test_successfulClaimOriginTokensWithProtocolFeeOn();
        assertTrue(fastBridge.protocolFees(address(arbUSDC)) > 0);

        vm.expectRevert();
        fastBridge.sweepProtocolFees(address(arbUSDC), user);
    }

    function test_passedSweepProtocolFeesWhenNoFees() public {
        setUpRoles();
        assertTrue(fastBridge.protocolFees(address(arbUSDC)) == 0);

        vm.startPrank(governor);

        uint256 preSweepBalanceUser = arbUSDC.balanceOf(user);

        fastBridge.sweepProtocolFees(address(arbUSDC), user);

        uint256 postSweepProtocolFees = fastBridge.protocolFees(address(arbUSDC));
        uint256 postSweepBalanceUser = arbUSDC.balanceOf(user);

        assertEq(postSweepProtocolFees, 0);
        assertEq(postSweepBalanceUser - preSweepBalanceUser, 0);

        vm.stopPrank();
    }
}
