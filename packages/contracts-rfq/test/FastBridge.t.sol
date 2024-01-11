// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "forge-std/console2.sol";

import "../contracts/FastBridge.sol";
import "../contracts/interfaces/IFastBridge.sol";
import "../contracts/libs/Errors.sol";
import "../contracts/libs/UniversalToken.sol";

import "./MockERC20.sol";

contract FastBridgeTest is Test {
    FastBridge public fastBridge;

    address owner = address(1);
    address relayer = address(2);
    address guard = address(3);
    address user = address(4);
    address dstUser = address(5);
    address governor = address(6);
    MockERC20 arbUSDC;
    MockERC20 ethUSDC;

    function setUp() public {
        vm.chainId(42161);
        fastBridge = new FastBridge(owner);
        arbUSDC = new MockERC20("arbUSDC", 6);
        ethUSDC = new MockERC20("ethUSDC", 6);
        _mintTokensToActors();
    }

    function _mintTokensToActors() internal {
        arbUSDC.mint(relayer, 100 * 10 ** 6);
        arbUSDC.mint(guard, 100 * 10 ** 6);
        arbUSDC.mint(user, 100 * 10 ** 6);
        arbUSDC.mint(dstUser, 100 * 10 ** 6);
        ethUSDC.mint(relayer, 100 * 10 ** 6);
        ethUSDC.mint(guard, 100 * 10 ** 6);
        ethUSDC.mint(user, 100 * 10 ** 6);
        ethUSDC.mint(dstUser, 100 * 10 ** 6);
    }

    function _getBridgeRequestAndId(uint256 chainId, uint256 currentNonce, uint256 protocolFeeRate)
        internal
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        address oldRelayer = relayer;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = address(arbUSDC);
        address destToken = address(ethUSDC);
        uint256 originAmount = 11 * 10 ** 6;
        uint256 destAmount = 10.97e6;
        uint256 deadline = block.timestamp + 3600;

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

    function _getBridgeRequestAndIdWithETH(uint256 chainId, uint256 currentNonce, uint256 protocolFeeRate)
        internal
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        address oldRelayer = relayer;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = UniversalTokenLib.ETH_ADDRESS;
        address destToken = UniversalTokenLib.ETH_ADDRESS;
        uint256 originAmount = 11 * 10 ** 18;
        uint256 destAmount = 10.97e18;
        uint256 deadline = block.timestamp + 3600;

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

    function _getBridgeRequestAndIdWithChainGas(uint256 chainId, uint256 currentNonce, uint256 protocolFeeRate)
        internal
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        address oldRelayer = relayer;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = address(arbUSDC);
        address destToken = address(ethUSDC);
        uint256 originAmount = 11 * 10 ** 6;
        uint256 destAmount = 10.97e6;
        uint256 deadline = block.timestamp + 3600;

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

    function _getBridgeRequestAndIdWithETHAndChainGas(uint256 chainId, uint256 currentNonce, uint256 protocolFeeRate)
        internal
        returns (bytes memory request, bytes32 transactionId)
    {
        // Define input variables for the bridge transaction
        address to = user;
        address oldRelayer = relayer;
        uint32 originChainId = uint32(chainId);
        uint32 dstChainId = 1;
        address originToken = UniversalTokenLib.ETH_ADDRESS;
        address destToken = UniversalTokenLib.ETH_ADDRESS;
        uint256 originAmount = 11 * 10 ** 18;
        uint256 destAmount = 10.97e18;
        uint256 deadline = block.timestamp + 3600;

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

    function setUpRoles() public {
        vm.startPrank(owner);
        fastBridge.addRelayer(relayer);
        fastBridge.addGuard(guard);
        fastBridge.addGovernor(governor);
        assertTrue(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        assertTrue(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        assertTrue(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        vm.stopPrank();
    }

    /// @notice Test to check if the owner is correctly set
    function test_owner() public {
        assertTrue(fastBridge.hasRole(fastBridge.DEFAULT_ADMIN_ROLE(), owner));
    }

    /// @notice Test to check if a relayer can be successfully added
    function test_successfulAddRelayer() public {
        vm.startPrank(owner);
        assertFalse(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        fastBridge.addRelayer(relayer);
        assertTrue(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
    }

    /// @notice Test to check if only an admin can add a relayer
    function test_onlyAdminCanAddRelayer() public {
        vm.startPrank(relayer);
        assertFalse(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        vm.expectRevert();
        fastBridge.addRelayer(relayer);
    }

    /// @notice Test to check if a relayer can be successfully removed
    function test_successfulRemoveRelayer() public {
        test_successfulAddRelayer();
        assertTrue(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        vm.startPrank(owner);
        fastBridge.removeRelayer(relayer);
        assertFalse(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
    }

    /// @notice Test to check if only an admin can remove a relayer
    function test_onlyAdminCanRemoveRelayer() public {
        test_successfulAddRelayer();
        vm.startPrank(relayer);
        assertTrue(fastBridge.hasRole(fastBridge.RELAYER_ROLE(), relayer));
        vm.expectRevert();
        fastBridge.removeRelayer(relayer);
    }

    /// @notice Test to check if a guard can be successfully added
    function test_successfulAddGuard() public {
        vm.startPrank(owner);
        assertFalse(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        fastBridge.addGuard(guard);
        assertTrue(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
    }

    /// @notice Test to check if only an admin can add a guard
    function test_onlyAdminCanAddGuard() public {
        vm.startPrank(guard);
        assertFalse(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        vm.expectRevert();
        fastBridge.addGuard(guard);
    }

    /// @notice Test to check if a guard can be successfully removed
    function test_successfulRemoveGuard() public {
        test_successfulAddGuard();
        assertTrue(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        vm.startPrank(owner);
        fastBridge.removeGuard(guard);
        assertFalse(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
    }

    /// @notice Test to check if only an admin can remove a guard
    function test_onlyAdminCanRemoveGuard() public {
        test_successfulAddGuard();
        vm.startPrank(guard);
        assertTrue(fastBridge.hasRole(fastBridge.GUARD_ROLE(), guard));
        vm.expectRevert();
        fastBridge.removeGuard(guard);
    }

    /// @notice Tests to check governor add, remove
    function test_successfulAddGovernor() public {
        vm.startPrank(owner);
        assertFalse(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        fastBridge.addGovernor(governor);
        assertTrue(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
    }

    function test_onlyAdminCanAddGovernor() public {
        vm.startPrank(governor);
        assertFalse(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        vm.expectRevert();
        fastBridge.addGovernor(governor);
    }

    function test_successfulRemoveGovernor() public {
        test_successfulAddGovernor();
        assertTrue(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        vm.startPrank(owner);
        fastBridge.removeGovernor(governor);
        assertFalse(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
    }

    function test_onlyAdminCanRemoveGovernor() public {
        test_successfulAddGovernor();
        vm.startPrank(governor);
        assertTrue(fastBridge.hasRole(fastBridge.GOVERNOR_ROLE(), governor));
        vm.expectRevert();
        fastBridge.removeGovernor(governor);
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
            deadline: block.timestamp + 3600
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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(block.chainid, currentNonce, 0);

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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
        (bytes memory request, bytes32 transactionId) =
            _getBridgeRequestAndIdWithETHAndChainGas(block.chainid, currentNonce, 0);

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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
            deadline: block.timestamp + 3600
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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(42161, 0, 0);

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
            transactionId, relayer, user, 42161, address(arbUSDC), address(ethUSDC), 11 * 10 ** 6, 10.97e6, 0
        );
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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(42161, 0, 0);

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
            42161,
            UniversalTokenLib.ETH_ADDRESS,
            UniversalTokenLib.ETH_ADDRESS,
            11 * 10 ** 18,
            10.97e18,
            0
        );

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETHAndChainGas(42161, 0, 0);

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
            42161,
            UniversalTokenLib.ETH_ADDRESS,
            UniversalTokenLib.ETH_ADDRESS,
            11 * 10 ** 18,
            10.97e18,
            0.005e18
        );

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithChainGas(42161, 0, 0);

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
            transactionId, relayer, user, 42161, address(arbUSDC), address(ethUSDC), 11 * 10 ** 6, 10.97e6, 0.005e18
        );
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

    function test_failedRelayNotDestChain() public {
        // Set up the roles for the test
        setUpRoles();

        vm.prank(owner);
        fastBridge.addRelayer(address(this));

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(42161, 0, 0);

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

        vm.prank(owner);
        fastBridge.addRelayer(address(this));

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(42161, 0, 0);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(42161, 0, 0);
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

    function test_failedRelayNotRelayer() public {
        // Set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(42161, 0, 0);

        // Start a prank with the relayer
        vm.startPrank(guard);
        // Approve the fastBridge to spend the maximum amount of ethUSDC from the relayer
        ethUSDC.approve(address(fastBridge), type(uint256).max);

        // Relay the destination bridge
        vm.expectRevert("Caller is not a relayer");
        vm.chainId(1); // set to dest chain
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
        (uint96 _timestamp, address _oldRelayer) = fastBridge.bridgeProofs(transactionId);
        assertEq(_timestamp, uint96(block.timestamp));
        assertEq(_oldRelayer, relayer);

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
        (uint96 _timestamp, address _oldRelayer) = fastBridge.bridgeProofs(transactionId);
        assertEq(_timestamp, uint96(block.timestamp));
        assertEq(_oldRelayer, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedProveTimeExceeded() public {
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

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

        // We provide the relay proof
        vm.expectRevert(abi.encodeWithSelector(DeadlineExceeded.selector));
        fastBridge.prove(request, fakeTxnHash);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedProveNotRequested() public {
        // Then, we set up the roles for the test
        setUpRoles();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // We provide the relay proof
        vm.expectRevert("Caller is not a relayer");
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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(block.chainid, 0, 0);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(block.chainid, 0, protocolFeeRate);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(relayer);

        vm.warp(block.timestamp + 31 minutes);

        vm.expectRevert(abi.encodeWithSelector(StatusIncorrect.selector));
        fastBridge.claim(request, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedClaimNotoldRelayer() public {
        setUpRoles();
        test_successfulBridge();

        vm.prank(owner);
        fastBridge.addRelayer(address(this));

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(relayer);
        fastBridge.prove(request, bytes32("0x04"));

        vm.expectRevert(abi.encodeWithSelector(DisputePeriodNotPassed.selector));
        fastBridge.claim(request, relayer);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedClaimNotRelayer() public {
        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.warp(block.timestamp + 31 minutes);

        vm.expectRevert("Caller is not a relayer");
        fastBridge.claim(request, relayer);
    }

    event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer);

    function test_successfulDisputeProof() public {
        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(guard);

        vm.expectEmit();
        emit BridgeProofDisputed(transactionId, guard);

        fastBridge.dispute(transactionId);

        // check status and proofs updated
        (uint96 _timestamp, address _oldRelayer) = fastBridge.bridgeProofs(transactionId);
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));
        assertEq(_timestamp, 0);
        assertEq(_oldRelayer, address(0));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_successfulDisputeProofWithProofTimestampOverflow() public {
        // sets block timestamp to just before overflow of uint96
        vm.warp(type(uint96).max - 1 minutes);

        setUpRoles();
        test_successfulRelayProof();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.warp(block.timestamp + 25 minutes);

        vm.startPrank(guard);

        fastBridge.dispute(transactionId);

        // check status and proofs updated
        (uint96 _timestamp, address _oldRelayer) = fastBridge.bridgeProofs(transactionId);
        assertEq(uint256(fastBridge.bridgeStatuses(transactionId)), uint256(FastBridge.BridgeStatus.REQUESTED));
        assertEq(_timestamp, 0);
        assertEq(_oldRelayer, address(0));

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedDisputeNoProof() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.expectRevert("Caller is not a guard");
        fastBridge.dispute(transactionId);
    }

    event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount);

    function test_successfulRefund() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(user);

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(block.chainid, 0, 0);

        vm.startPrank(user);

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

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

        vm.startPrank(user);

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

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
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndIdWithETH(block.chainid, 0, protocolFeeRate);

        vm.startPrank(user);

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

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

        vm.startPrank(user);

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

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

    function test_successfulRefundNotSender() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

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
    }

    function test_failedRefundNotEnoughTime() public {
        setUpRoles();
        test_successfulBridge();

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        vm.startPrank(user);

        vm.warp(block.timestamp + 119 minutes);

        vm.expectRevert(abi.encodeWithSelector(DeadlineNotExceeded.selector));
        fastBridge.refund(request);

        // We stop a prank to contain within test
        vm.stopPrank();
    }

    function test_failedRefundNoBridge() public {
        setUpRoles();

        vm.startPrank(user);

        // get bridge request and tx id
        (bytes memory request, bytes32 transactionId) = _getBridgeRequestAndId(block.chainid, 0, 0);

        // user relay deadline is 60 minutes, so add prove period of 60 minutes to that
        vm.warp(block.timestamp + 121 minutes);

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

        uint256 preSweepProtocolFees = fastBridge.protocolFees(address(arbUSDC));
        uint256 preSweepBalanceUser = arbUSDC.balanceOf(user);

        fastBridge.sweepProtocolFees(address(arbUSDC), user);

        uint256 postSweepProtocolFees = fastBridge.protocolFees(address(arbUSDC));
        uint256 postSweepBalanceUser = arbUSDC.balanceOf(user);

        assertEq(postSweepProtocolFees, 0);
        assertEq(postSweepBalanceUser - preSweepBalanceUser, 0);

        vm.stopPrank();
    }
}
