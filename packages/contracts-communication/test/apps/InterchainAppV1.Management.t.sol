// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {InterchainAppV1Test} from "./InterchainAppV1.t.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
abstract contract InterchainAppV1ManagementTest is InterchainAppV1Test {
    address public newClient = makeAddr("New Client");
    address public newModule = makeAddr("New Module");
    address public newExecutionService = makeAddr("New Execution Service");

    function addInterchainClient(address client, bool updateLatest) public {
        vm.prank(governor);
        appHarness.addInterchainClient({client: client, updateLatest: updateLatest});
    }

    function removeInterchainClient(address client) public {
        vm.prank(governor);
        appHarness.removeInterchainClient(client);
    }

    function setLatestInterchainClient(address client) public {
        vm.prank(governor);
        appHarness.setLatestInterchainClient(client);
    }

    function expectRevertUnauthorizedGovernor(address caller) internal virtual;

    function test_addInterchainClient_dontUpdateLatest_noClients() public {
        expectEventInterchainClientAdded(newClient);
        vm.recordLogs();
        addInterchainClient({client: newClient, updateLatest: false});
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(newClient));
        assertEq(appHarness.getLatestInterchainClient(), address(0));
    }

    function test_addInterchainClient_dontUpdateLatest_hasClients_noLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: false});
        expectEventInterchainClientAdded(newClient);
        vm.recordLogs();
        addInterchainClient({client: newClient, updateLatest: false});
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(icClient, newClient));
        assertEq(appHarness.getLatestInterchainClient(), address(0));
    }

    function test_addInterchainClient_dontUpdateLatest_hasClients_hasLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: true});
        expectEventInterchainClientAdded(newClient);
        vm.recordLogs();
        addInterchainClient({client: newClient, updateLatest: false});
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(icClient, newClient));
        assertEq(appHarness.getLatestInterchainClient(), icClient);
    }

    function test_addInterchainClient_dontUpdateLatest_hasClients_revert_sameClient() public {
        addInterchainClient({client: icClient, updateLatest: true});
        expectRevertClientAlreadyAdded(icClient);
        addInterchainClient({client: icClient, updateLatest: false});
    }

    function test_addInterchainClient_updateLatest_noClients() public {
        expectEventInterchainClientAdded(newClient);
        expectEventLatestClientSet(newClient);
        vm.recordLogs();
        addInterchainClient({client: newClient, updateLatest: true});
        assertEq(vm.getRecordedLogs().length, 2);
        assertEq(appHarness.getInterchainClients(), toArray(newClient));
        assertEq(appHarness.getLatestInterchainClient(), newClient);
    }

    function test_addInterchainClient_updateLatest_hasClients_noLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: false});
        expectEventInterchainClientAdded(newClient);
        expectEventLatestClientSet(newClient);
        vm.recordLogs();
        addInterchainClient({client: newClient, updateLatest: true});
        assertEq(vm.getRecordedLogs().length, 2);
        assertEq(appHarness.getInterchainClients(), toArray(icClient, newClient));
        assertEq(appHarness.getLatestInterchainClient(), newClient);
    }

    function test_addInterchainClient_updateLatest_hasClients_hasLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: true});
        expectEventInterchainClientAdded(newClient);
        expectEventLatestClientSet(newClient);
        vm.recordLogs();
        addInterchainClient({client: newClient, updateLatest: true});
        assertEq(vm.getRecordedLogs().length, 2);
        assertEq(appHarness.getInterchainClients(), toArray(icClient, newClient));
        assertEq(appHarness.getLatestInterchainClient(), newClient);
    }

    function test_addInterchainClient_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.addInterchainClient({client: newClient, updateLatest: false});
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.addInterchainClient({client: newClient, updateLatest: true});
    }

    function test_addInterchainClient_revert_zeroAddress() public {
        expectRevertInterchainClientZeroAddress();
        addInterchainClient({client: address(0), updateLatest: false});
        expectRevertInterchainClientZeroAddress();
        addInterchainClient({client: address(0), updateLatest: true});
    }

    function test_removeInterchainClient_singleClient_noLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: false});
        expectEventInterchainClientRemoved(icClient);
        vm.recordLogs();
        removeInterchainClient(icClient);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), new address[](0));
        assertEq(appHarness.getLatestInterchainClient(), address(0));
    }

    function test_removeInterchainClient_singleClient_hasLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: true});
        expectEventInterchainClientRemoved(icClient);
        expectEventLatestClientSet(address(0));
        vm.recordLogs();
        removeInterchainClient(icClient);
        assertEq(vm.getRecordedLogs().length, 2);
        assertEq(appHarness.getInterchainClients(), new address[](0));
        assertEq(appHarness.getLatestInterchainClient(), address(0));
    }

    function test_removeInterchainClient_twoClients_noLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: false});
        addInterchainClient({client: newClient, updateLatest: false});
        expectEventInterchainClientRemoved(icClient);
        vm.recordLogs();
        removeInterchainClient(icClient);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(newClient));
        assertEq(appHarness.getLatestInterchainClient(), address(0));
    }

    function test_removeInterchainClient_twoClients_hasLatestSet_removeNotLatest() public {
        addInterchainClient({client: icClient, updateLatest: true});
        addInterchainClient({client: newClient, updateLatest: false});
        expectEventInterchainClientRemoved(newClient);
        vm.recordLogs();
        removeInterchainClient(newClient);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(icClient));
        assertEq(appHarness.getLatestInterchainClient(), icClient);
    }

    function test_removeInterchainClient_twoClients_hasLatestSet_removeLatest() public {
        addInterchainClient({client: icClient, updateLatest: true});
        addInterchainClient({client: newClient, updateLatest: false});
        expectEventInterchainClientRemoved(icClient);
        expectEventLatestClientSet(address(0));
        vm.recordLogs();
        removeInterchainClient(icClient);
        assertEq(vm.getRecordedLogs().length, 2);
        assertEq(appHarness.getInterchainClients(), toArray(newClient));
        assertEq(appHarness.getLatestInterchainClient(), address(0));
    }

    function test_removeInterchainClient_revert_clientNotAdded() public {
        addInterchainClient({client: icClient, updateLatest: true});
        expectRevertCallerNotInterchainClient(newClient);
        removeInterchainClient(newClient);
    }

    function test_removeInterchainClient_revert_notGovernor(address caller) public {
        addInterchainClient({client: icClient, updateLatest: true});
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.removeInterchainClient(icClient);
    }

    function test_setLatestInterchainClient_noLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: false});
        addInterchainClient({client: newClient, updateLatest: false});
        expectEventLatestClientSet(icClient);
        vm.recordLogs();
        setLatestInterchainClient(icClient);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(icClient, newClient));
        assertEq(appHarness.getLatestInterchainClient(), icClient);
    }

    function test_setLatestInterchainClient_noLatestSet_revert_alreadyLatest() public {
        addInterchainClient({client: icClient, updateLatest: false});
        addInterchainClient({client: newClient, updateLatest: false});
        expectRevertAlreadyLatestClient(address(0));
        setLatestInterchainClient(address(0));
    }

    function test_setLatestInterchainClient_hasLatestSet() public {
        addInterchainClient({client: icClient, updateLatest: true});
        addInterchainClient({client: newClient, updateLatest: false});
        expectEventLatestClientSet(newClient);
        vm.recordLogs();
        setLatestInterchainClient(newClient);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(icClient, newClient));
        assertEq(appHarness.getLatestInterchainClient(), newClient);
    }

    function test_setLatestInterchainClient_hasLatestSet_setToZero() public {
        addInterchainClient({client: icClient, updateLatest: true});
        addInterchainClient({client: newClient, updateLatest: false});
        expectEventLatestClientSet(address(0));
        vm.recordLogs();
        setLatestInterchainClient(address(0));
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getInterchainClients(), toArray(icClient, newClient));
        assertEq(appHarness.getLatestInterchainClient(), address(0));
    }

    function test_setLatestInterchainClient_revert_alreadyLatest() public {
        addInterchainClient({client: icClient, updateLatest: true});
        expectRevertAlreadyLatestClient(icClient);
        setLatestInterchainClient(icClient);
    }

    function test_setLatestInterchainClient_revert_CallerNotInterchainClient() public {
        addInterchainClient({client: icClient, updateLatest: true});
        expectRevertCallerNotInterchainClient(newClient);
        setLatestInterchainClient(newClient);
    }

    function test_setLatestInterchainClient_revert_notGovernor(address caller) public {
        addInterchainClient({client: icClient, updateLatest: true});
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.setLatestInterchainClient(icClient);
    }

    function test_linkRemoteApp() public {
        expectEventAppLinked(REMOTE_CHAIN_ID, linkedAppMockBytes32);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.linkRemoteApp({chainId: REMOTE_CHAIN_ID, remoteApp: linkedAppMockBytes32});
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getLinkedApp(REMOTE_CHAIN_ID), linkedAppMockBytes32);
        assertEq(appHarness.getLinkedAppEVM(REMOTE_CHAIN_ID), linkedAppMock);
    }

    function test_linkRemoteApp_revert_ChainIdNotRemote() public {
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        vm.prank(governor);
        appHarness.linkRemoteApp({chainId: LOCAL_CHAIN_ID, remoteApp: linkedAppMockBytes32});
    }

    function test_linkRemoteApp_revert_zeroAddress() public {
        expectRevertRemoteAppZeroAddress();
        vm.prank(governor);
        appHarness.linkRemoteApp({chainId: REMOTE_CHAIN_ID, remoteApp: 0});
    }

    function test_linkRemoteApp_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.linkRemoteApp({chainId: REMOTE_CHAIN_ID, remoteApp: linkedAppMockBytes32});
    }

    function test_linkRemoteAppEVM() public {
        expectEventAppLinked(REMOTE_CHAIN_ID, linkedAppMockBytes32);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.linkRemoteAppEVM(REMOTE_CHAIN_ID, linkedAppMock);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getLinkedApp(REMOTE_CHAIN_ID), linkedAppMockBytes32);
        assertEq(appHarness.getLinkedAppEVM(REMOTE_CHAIN_ID), linkedAppMock);
    }

    function test_linkRemoteAppEVM_revert_ChainIdNotRemote() public {
        expectRevertChainIdNotRemote(LOCAL_CHAIN_ID);
        vm.prank(governor);
        appHarness.linkRemoteAppEVM(LOCAL_CHAIN_ID, linkedAppMock);
    }

    function test_linkRemoteAppEVM_revert_zeroAddress() public {
        expectRevertRemoteAppZeroAddress();
        vm.prank(governor);
        appHarness.linkRemoteAppEVM(REMOTE_CHAIN_ID, address(0));
    }

    function test_linkRemoteAppEVM_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.linkRemoteAppEVM(REMOTE_CHAIN_ID, linkedAppMock);
    }

    function test_getLinkedAppEVM_revert_notEVMAddress() public {
        bytes32 nonEVM = keccak256("nonEVM");
        vm.prank(governor);
        appHarness.linkRemoteApp({chainId: REMOTE_CHAIN_ID, remoteApp: nonEVM});
        expectRevertNotEVMLinkedApp(nonEVM);
        appHarness.getLinkedAppEVM(REMOTE_CHAIN_ID);
    }

    function test_addTrustedModule_noModules() public {
        expectEventTrustedModuleAdded(moduleMock);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getModules(), toArray(moduleMock));
    }

    function test_addTrustedModule_hasModules() public {
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
        expectEventTrustedModuleAdded(newModule);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.addTrustedModule(newModule);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getModules(), toArray(moduleMock, newModule));
    }

    function test_addTrustedModule_revert_sameModule() public {
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
        expectRevertModuleAlreadyAdded(moduleMock);
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
    }

    function test_addTrustedModule_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.addTrustedModule(moduleMock);
    }

    function test_addTrustedModule_revert_zeroAddress() public {
        expectRevertModuleZeroAddress();
        vm.prank(governor);
        appHarness.addTrustedModule(address(0));
    }

    function test_removeTrustedModule_singleModule() public {
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
        expectEventTrustedModuleRemoved(moduleMock);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.removeTrustedModule(moduleMock);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getModules(), new address[](0));
    }

    function test_removeTrustedModule_twoModules() public {
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
        vm.prank(governor);
        appHarness.addTrustedModule(newModule);
        expectEventTrustedModuleRemoved(moduleMock);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.removeTrustedModule(moduleMock);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getModules(), toArray(newModule));
    }

    function test_removeTrustedModule_revert_moduleNotAdded() public {
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
        expectRevertModuleNotAdded(newModule);
        vm.prank(governor);
        appHarness.removeTrustedModule(newModule);
    }

    function test_removeTrustedModule_revert_notGovernor(address caller) public {
        vm.prank(governor);
        appHarness.addTrustedModule(moduleMock);
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.removeTrustedModule(moduleMock);
    }

    function test_setAppConfigV1_whenNotSet() public {
        expectEventAppConfigV1Set(APP_REQUIRED_RESPONSES, APP_OPTIMISTIC_PERIOD);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.setAppConfigV1(APP_REQUIRED_RESPONSES, APP_OPTIMISTIC_PERIOD);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getAppConfigV1().requiredResponses, APP_REQUIRED_RESPONSES);
        assertEq(appHarness.getAppConfigV1().optimisticPeriod, APP_OPTIMISTIC_PERIOD);
    }

    function test_setAppConfigV1_whenSet() public {
        vm.prank(governor);
        appHarness.setAppConfigV1(APP_REQUIRED_RESPONSES, APP_OPTIMISTIC_PERIOD);
        expectEventAppConfigV1Set(APP_REQUIRED_RESPONSES + 1, APP_OPTIMISTIC_PERIOD + 1);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.setAppConfigV1(APP_REQUIRED_RESPONSES + 1, APP_OPTIMISTIC_PERIOD + 1);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getAppConfigV1().requiredResponses, APP_REQUIRED_RESPONSES + 1);
        assertEq(appHarness.getAppConfigV1().optimisticPeriod, APP_OPTIMISTIC_PERIOD + 1);
    }

    function test_setAppConfigV1_revert_notGovernor(address caller) public {
        vm.assume(caller != governor);
        expectRevertUnauthorizedGovernor(caller);
        vm.prank(caller);
        appHarness.setAppConfigV1(APP_REQUIRED_RESPONSES, APP_OPTIMISTIC_PERIOD);
    }

    function test_setAppConfigV1_revert_zeroConfirmations() public {
        expectRevertInvalidAppConfig(0, APP_OPTIMISTIC_PERIOD);
        vm.prank(governor);
        appHarness.setAppConfigV1(0, APP_OPTIMISTIC_PERIOD);
    }

    function test_setAppConfigV1_revert_zeroOptimisticPeriod() public {
        expectRevertInvalidAppConfig(APP_REQUIRED_RESPONSES, 0);
        vm.prank(governor);
        appHarness.setAppConfigV1(APP_REQUIRED_RESPONSES, 0);
    }

    function test_setAppConfigV1_revert_zeroedAppConfig() public {
        expectRevertInvalidAppConfig(0, 0);
        vm.prank(governor);
        appHarness.setAppConfigV1(0, 0);
    }

    function test_setExecutionService_whenNotSet() public {
        expectEventExecutionServiceSet(execServiceMock);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.setExecutionService(execServiceMock);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getExecutionService(), execServiceMock);
    }

    function test_setExecutionService_whenSet() public {
        vm.prank(governor);
        appHarness.setExecutionService(execServiceMock);
        expectEventExecutionServiceSet(newExecutionService);
        vm.recordLogs();
        vm.prank(governor);
        appHarness.setExecutionService(newExecutionService);
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getExecutionService(), newExecutionService);
    }

    function test_setExecutionService_whenSet_setToZero() public {
        vm.prank(governor);
        appHarness.setExecutionService(execServiceMock);
        expectEventExecutionServiceSet(address(0));
        vm.recordLogs();
        vm.prank(governor);
        appHarness.setExecutionService(address(0));
        assertEq(vm.getRecordedLogs().length, 1);
        assertEq(appHarness.getExecutionService(), address(0));
    }
}
