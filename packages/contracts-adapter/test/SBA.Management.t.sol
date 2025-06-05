// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ISynapseBridgeAdapter} from "../src/interfaces/ISynapseBridgeAdapter.sol";
import {SynapseBridgeAdapter, SynapseBridgeAdapterTest} from "./SBA.t.sol";

// solhint-disable func-name-mixedcase, ordering
contract SynapseBridgeAdapterManagementTest is SynapseBridgeAdapterTest {
    address internal bridge = makeAddr("Bridge");
    address internal owner = makeAddr("Owner");
    address internal token = makeAddr("Token");
    address internal anotherToken = makeAddr("AnotherToken");
    bytes31 internal symbol = "SYMBOL";
    bytes31 internal anotherSymbol = "ANOTHERSYMBOL";
    string internal readableSymbol = "SYMBOL";

    mapping(uint32 eid => address token) internal mockRemoteAddressMap;
    ISynapseBridgeAdapter.RemoteToken[] internal allRemoteTokens;
    ISynapseBridgeAdapter.RemoteToken[] internal firstRemoteToken;
    ISynapseBridgeAdapter.RemoteToken[] internal secondRemoteToken;

    function afterAdapterDeployed() internal virtual override {
        mockRemoteAddressMap[DST_EID] = makeAddr("DST Token");
        mockRemoteAddressMap[OTHER_DST_EID] = makeAddr("Other DST Token");
        allRemoteTokens.push(ISynapseBridgeAdapter.RemoteToken(DST_EID, mockRemoteAddressMap[DST_EID]));
        allRemoteTokens.push(ISynapseBridgeAdapter.RemoteToken(OTHER_DST_EID, mockRemoteAddressMap[OTHER_DST_EID]));
        firstRemoteToken.push(allRemoteTokens[0]);
        secondRemoteToken.push(allRemoteTokens[1]);
    }

    function deployAdapter() internal virtual override returns (SynapseBridgeAdapter) {
        return new SynapseBridgeAdapter(endpoint, owner);
    }

    function checkTokenAdded(
        address token_,
        ISynapseBridgeAdapter.TokenType tokenType_,
        ISynapseBridgeAdapter.RemoteToken[] memory remoteTokens_
    )
        internal
        view
    {
        // Check token type by address
        ISynapseBridgeAdapter.TokenType adapterTokenType = adapter.getTokenType(token_);
        assertEq(uint8(adapterTokenType), uint8(tokenType_));
        for (uint256 i = 0; i < remoteTokens_.length; i++) {
            ISynapseBridgeAdapter.RemoteToken memory remoteToken = remoteTokens_[i];
            // Check remote address by local address
            address adapterRemoteAddress = adapter.getRemoteAddress(remoteToken.eid, token_);
            assertEq(adapterRemoteAddress, remoteToken.addr);
            // Check local address by remote address
            address adapterLocalAddress = adapter.getLocalAddress(remoteToken.eid, remoteToken.addr);
            assertEq(adapterLocalAddress, token_);
        }
    }

    function test_constructor() public view {
        assertEq(address(adapter.endpoint()), endpoint);
        assertEq(adapter.owner(), owner);
        assertEq(adapter.bridge(), address(0));
        assertEq(adapter.MIN_GAS_LIMIT(), 200_000);
    }

    // ═════════════════════════════════════════════════ ADD TOKEN ═════════════════════════════════════════════════════

    function addToken(
        address token_,
        ISynapseBridgeAdapter.TokenType tokenType_,
        ISynapseBridgeAdapter.RemoteToken[] memory remoteTokens_
    )
        internal
    {
        vm.prank(owner);
        adapter.addToken(token_, tokenType_, remoteTokens_);
    }

    function test_addToken_mintBurn_oneByOne() public {
        // Add first token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        // Add second token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_allAtOnce() public {
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_tokenAlreadyAddedAsWithdrawDeposit() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
    }

    function test_addToken_mintBurn_revert_remotePairAlreadySet_sameRemoteAddress() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_remotePairAlreadySet_diffRemoteAddress() public {
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.MintBurn,
            toArray(ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, anotherToken))
        );
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_localPairAlreadyExists() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as single entry
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.MintBurn, secondRemoteToken);
        // Add as a batch
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_tokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        // Add as a batch
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.MintBurn, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_tokenTypeUnknown() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        // Add as single entry
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, secondRemoteToken);
        // Add as a batch
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, allRemoteTokens);
    }

    function test_addToken_mintBurn_revert_remoteTokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.MintBurn,
            toArray(ISynapseBridgeAdapter.RemoteToken(firstRemoteToken[0].eid, address(0)))
        );
        // Add as a batch
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.MintBurn,
            toArray(firstRemoteToken[0], ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, address(0)))
        );
    }

    function test_addToken_mintBurn_revert_existingTokenEmptyArray() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_mintBurn_revert_newTokenEmptyArray() public {
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_withdrawDeposit_oneByOne() public {
        // Add first token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        // Add second token
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_allAtOnce() public {
        expectEventTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
        checkTokenAdded(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_tokenAlreadyAddedAsMint() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.MintBurn, firstRemoteToken);
        expectRevertTokenAlreadyAdded(token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
    }

    function test_addToken_withdrawDeposit_revert_remotePairAlreadySet_sameRemoteAddress() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_remotePairAlreadySet_diffRemoteAddress() public {
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.WithdrawDeposit,
            toArray(ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, anotherToken))
        );
        // Add as single entry
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as a batch
        expectRevertRemotePairAlreadySet(secondRemoteToken[0].eid, token);
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_localPairAlreadyExists() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as single entry
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, secondRemoteToken);
        // Add as a batch
        expectRevertLocalPairAlreadyExists(secondRemoteToken[0].eid, secondRemoteToken[0].addr);
        addToken(anotherToken, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_tokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        // Add as a batch
        expectRevertZeroAddress();
        addToken(address(0), ISynapseBridgeAdapter.TokenType.WithdrawDeposit, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_tokenTypeUnknown() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        // Add as single entry
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, secondRemoteToken);
        // Add as a batch
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, allRemoteTokens);
    }

    function test_addToken_withdrawDeposit_revert_remoteTokenAddressZero() public {
        // Add as single entry
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.WithdrawDeposit,
            toArray(ISynapseBridgeAdapter.RemoteToken(firstRemoteToken[0].eid, address(0)))
        );
        // Add as a batch
        expectRevertZeroAddress();
        addToken(
            token,
            ISynapseBridgeAdapter.TokenType.WithdrawDeposit,
            toArray(firstRemoteToken[0], ISynapseBridgeAdapter.RemoteToken(secondRemoteToken[0].eid, address(0)))
        );
    }

    function test_addToken_withdrawDeposit_revert_existingTokenEmptyArray() public {
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, firstRemoteToken);
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_withdrawDeposit_revert_newTokenEmptyArray() public {
        expectRevertZeroAmount();
        addToken(token, ISynapseBridgeAdapter.TokenType.WithdrawDeposit, new ISynapseBridgeAdapter.RemoteToken[](0));
    }

    function test_addToken_revert_tokenTypeUnknown() public {
        // Add as single entry
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, secondRemoteToken);
        // Add as a batch
        expectRevertTokenTypeUnknown();
        addToken(token, ISynapseBridgeAdapter.TokenType.Unknown, allRemoteTokens);
    }

    // ════════════════════════════════════════════════ SET BRIDGE ═════════════════════════════════════════════════════

    function test_setBridge() public {
        expectEventBridgeSet(bridge);
        vm.prank(owner);
        adapter.setBridge(bridge);
        assertEq(adapter.bridge(), bridge);
    }

    function test_setBridge_revert_notOwner(address caller) public {
        vm.assume(caller != owner);
        expectRevertCallerNotOwner(caller);
        vm.prank(caller);
        adapter.setBridge(bridge);
    }

    function test_setBridge_revert_bridgeAlreadySet() public {
        vm.prank(owner);
        adapter.setBridge(bridge);
        expectRevertBridgeAlreadySet();
        vm.prank(owner);
        adapter.setBridge(bridge);
    }

    function test_setBridge_revert_zeroAddress() public {
        expectRevertZeroAddress();
        vm.prank(owner);
        adapter.setBridge(address(0));
    }
}
