// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import "../DestinationTools.t.sol";
import "../../../contracts/interfaces/ISystemRouter.sol";

abstract contract SystemRouterTools is DestinationTools {
    using ByteString for bytes;

    // domain => (list of all domain's system contracts)
    mapping(uint32 => address[]) internal systemContracts;

    // System caller information
    uint32 internal systemCallOrigin;
    address internal systemSender;
    ISystemRouter.SystemEntity internal systemCallSender;
    // System recipient information
    address internal systemRecipient;
    uint32 internal systemCallDestination;
    // Optimistic period for the system call
    uint32 internal systemCallOptimisticSeconds;
    // Saved system calls
    bytes4[] internal systemCallSelectors;
    // Recipient + data is used for systemRouter.systemCall()
    ISystemRouter.SystemEntity[] internal systemCallRecipients;
    bytes[] internal systemCallDataArray;
    // Formatted system calls
    bytes[] internal formattedSystemCalls;

    // Save all deployed system contracts
    function saveSystemContracts() public {
        for (uint256 d = 0; d < DOMAINS; ++d) {
            uint32 domain = domains[d];
            address[] storage domainContracts = systemContracts[domain];
            // Don't save twice
            if (domainContracts.length != 0) return;
            // Save all the system contracts in the same order as ISystemRouter.SystemEntity enum
            domainContracts.push(address(suiteOrigin(domain)));
            domainContracts.push(address(suiteDestination(domain)));
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                           CREATE TEST DATA                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Creates a dispatched SystemCall message
    function createDispatchedSystemCallMessage(uint32 origin, uint32 destination) public {
        // Context for the cross-chain message
        MessageContext memory context = MessageContext({
            origin: origin,
            sender: address(suiteSystemRouter(origin)),
            destination: destination
        });
        // System message is sent without the tips, message body is formatted system calls
        createEmptyTips();
        createDispatchedMessage({
            context: context,
            mockTips: false,
            body: abi.encode(formattedSystemCalls),
            recipient: SystemCall.SYSTEM_ROUTER,
            optimisticSeconds: 0
        });
        // Save dispatched message for later execution
        rawMessages.push(messageRaw);
        messageHashes.push(keccak256(messageRaw));
    }

    function createSystemCall(MessageContext memory context, address recipient) public {
        createSystemCall({
            context: context,
            index: 0,
            selector: SystemContractHarness.setSensitiveValue.selector,
            optimisticSeconds: 0,
            recipient: recipient,
            deleteCalls: true,
            formatCalls: true
        });
    }

    // Saves test data for a system call with given parameters
    function createSystemCall(
        MessageContext memory context,
        uint256 index,
        bytes4 selector,
        uint32 optimisticSeconds,
        address recipient,
        bool deleteCalls,
        bool formatCalls
    ) public {
        // Clear system calls arrays. Needed for systemCall() tests, so that
        // `formattedSystemCalls` has exactly one element: constructed system call
        if (deleteCalls) deleteSystemCalls();
        systemCallOrigin = context.origin;
        systemCallDestination = context.destination;
        systemSender = context.sender;
        systemRecipient = recipient;
        systemCallOptimisticSeconds = optimisticSeconds;
        if (_isSystemEntity({ domain: context.origin, addr: context.sender })) {
            systemCallSender = _getSystemEntity({ domain: context.origin, addr: context.sender });
        } else {
            // Use default value for "unauthorized access" tests
            systemCallSender = ISystemRouter.SystemEntity(0);
        }
        // Recipient and data is used, when calling systemRouter.systemCall()
        systemCallSelectors.push(selector);
        systemCallRecipients.push(
            _getSystemEntity({ domain: context.destination, addr: recipient })
        );
        // Use (0, 0, 0) for security arguments
        systemCallDataArray.push(
            abi.encodeWithSelector(selector, 0, 0, 0, _createMockSensitiveValue(index))
        );
        // Format calls. Abi encoded array of formatted calls forms a cross-chain message body.
        if (formatCalls) formatSystemCalls();
    }

    // Clear all system calls arrays
    function deleteSystemCalls() public {
        delete systemCallSelectors;
        delete systemCallRecipients;
        delete systemCallDataArray;
        // There's no need to delete formattedSystemCalls, it is reconstructed
        require(systemCallSelectors.length == 0, "!systemCallSelectors: clear");
        require(systemCallRecipients.length == 0, "!systemCallRecipients: clear");
        require(systemCallDataArray.length == 0, "!systemCallDataArray: clear");
    }

    // Format system calls using the saved data
    function formatSystemCalls() public {
        uint256 amount = systemCallRecipients.length;
        formattedSystemCalls = new bytes[](amount);
        for (uint256 i = 0; i < amount; ++i) {
            // Reconstruct payload created by System Router.
            // It adjusts three parameters to the passed calldata arguments:
            // Root timestamp, domain of origin chain, and system entity that sent the message
            bytes memory prefix = abi.encode(block.timestamp, systemCallOrigin, systemCallSender);
            // Save formatted system call
            formattedSystemCalls[i] = SystemCall.formatSystemCall({
                _systemRecipient: uint8(systemCallRecipients[i]),
                _payload: systemCallDataArray[i].castToCallPayload(),
                _prefix: prefix.castToRawBytes()
            });
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            EXPECT EVENTS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Expect N*N system call test events, assuming every sender called
    // every recipient one by one.
    function expectSystemCalls(uint32 domain, bytes4 selector) public {
        uint256 amount = systemContracts[domain].length;
        uint256 index = 0;
        for (uint256 s = 0; s < amount; ++s) {
            for (uint256 r = 0; r < amount; ++r) {
                address recipient = systemContracts[domain][r];
                expectSystemCallEvent({
                    recipient: recipient,
                    selector: selector,
                    newValue: _createMockSensitiveValue(index)
                });
                ++index;
            }
        }
        // Sanity check
        require(index > 0, "No events are expected");
    }

    function expectSystemCallsWrapperTests(
        uint32 domain,
        bytes4 selector,
        bool sameRecipient
    ) public {
        uint256 amount = systemContracts[domain].length;
        uint256 index = 0;
        for (uint256 s = 0; s < amount; ++s) {
            for (uint256 r = 0; r < amount; ++r) {
                // Wrapper tests use either same recipient
                address recipient = sameRecipient
                    ? systemContracts[domain][0]
                    : systemContracts[domain][r];
                // Or they use the same data
                uint256 newValue = sameRecipient
                    ? _createMockSensitiveValue(index)
                    : _createMockSensitiveValue(s * amount);
                expectSystemCallEvent(recipient, selector, newValue);
                ++index;
            }
        }
        // Sanity check
        require(index > 0, "No events are expected");
    }

    // Expect a system call test event with a given index
    function expectSystemCall(uint256 index) public {
        expectSystemCallEvent({
            recipient: _getSystemAddress(systemCallDestination, systemCallRecipients[index]),
            selector: systemCallSelectors[index],
            newValue: _createMockSensitiveValue(index)
        });
    }

    // Expect all system call test events
    function expectSystemMultiCall() public {
        for (uint256 i = 0; i < systemCallRecipients.length; ++i) {
            expectSystemCall(i);
        }
    }

    // Figure out the test event name by selector and expect such an event
    // solhint-disable-next-line code-complexity
    function expectSystemCallEvent(
        address recipient,
        bytes4 selector,
        uint256 newValue
    ) public {
        vm.expectEmit(true, true, true, true);
        if (selector == SystemContractHarness.setSensitiveValue.selector) {
            emit UsualCall(recipient, newValue);
        } else if (selector == SystemContractHarness.setSensitiveValueOnlyLocal.selector) {
            emit OnlyLocalCall(recipient, newValue);
        } else if (selector == SystemContractHarness.setSensitiveValueOnlyOrigin.selector) {
            emit OnlyOriginCall(recipient, newValue);
        } else if (selector == SystemContractHarness.setSensitiveValueOnlyDestination.selector) {
            emit OnlyDestinationCall(recipient, newValue);
        } else if (
            selector == SystemContractHarness.setSensitiveValueOnlyOriginDestination.selector
        ) {
            emit OnlyOriginDestinationCall(recipient, newValue);
        } else if (selector == SystemContractHarness.setSensitiveValueOnlyTwoHours.selector) {
            emit OnlyTwoHoursCall(recipient, newValue);
        } else if (selector == SystemContractHarness.setSensitiveValueOnlySynapseChain.selector) {
            emit OnlySynapseChainCall(recipient, newValue);
        } else {
            revert("Unknown selector");
        }
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                     TRIGGER FUNCTIONS (REVERTS)                      ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger systemRouter.systemCall() and expect a revert
    function systemRouterCall(uint256 index, bytes memory revertMessage) public {
        vm.expectRevert(revertMessage);
        systemRouterCall(index);
    }

    // Trigger systemRouter.systemMultiCall() and expect a revert
    function systemRouterMultiCall(bytes memory revertMessage) public {
        vm.expectRevert(revertMessage);
        systemRouterMultiCall();
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                          TRIGGER FUNCTIONS                           ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger systemRouter.systemCall() with the saved data
    function systemRouterCall(uint256 index) public {
        vm.prank(systemSender);
        suiteSystemRouter(systemCallOrigin).systemCall({
            _destination: systemCallDestination,
            _optimisticSeconds: systemCallOptimisticSeconds,
            _recipient: systemCallRecipients[index],
            _data: systemCallDataArray[index]
        });
    }

    // Trigger systemRouter.systemMultiCall() with the saved data
    function systemRouterMultiCall() public {
        vm.prank(systemSender);
        suiteSystemRouter(systemCallOrigin).systemMultiCall({
            _destination: systemCallDestination,
            _optimisticSeconds: systemCallOptimisticSeconds,
            _recipients: systemCallRecipients,
            _dataArray: systemCallDataArray
        });
    }

    function systemRouterMultiCallSameRecipient() public {
        vm.prank(systemSender);
        suiteSystemRouter(systemCallOrigin).systemMultiCall({
            _destination: systemCallDestination,
            _optimisticSeconds: systemCallOptimisticSeconds,
            _recipient: systemCallRecipients[0],
            _dataArray: systemCallDataArray
        });
    }

    function systemRouterMultiCallSameData() public {
        vm.prank(systemSender);
        suiteSystemRouter(systemCallOrigin).systemMultiCall({
            _destination: systemCallDestination,
            _optimisticSeconds: systemCallOptimisticSeconds,
            _recipients: systemCallRecipients,
            _data: systemCallDataArray[0]
        });
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                             TEST HELPERS                             ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Trigger a bunch of tets system calls.
    // For systemCall(): N*N calls, every sender calls every recipient
    // For systemMultiCall(): N calls, every sender does a multicall to every recipient
    // solhint-disable-next-line code-complexity
    function triggerTestSystemCalls(
        uint32 origin,
        uint32 destination,
        bytes4 selector,
        bool isMultiCall
    ) public {
        uint256 amount = systemContracts[origin].length;
        uint256 index = 0;
        for (uint256 s = 0; s < amount; ++s) {
            address sender = systemContracts[origin][s];
            // Context for a on-chain call
            MessageContext memory contextLocal = MessageContext(origin, sender, destination);
            deleteSystemCalls();
            for (uint256 r = 0; r < amount; (++r, ++index)) {
                createSystemCall({
                    context: contextLocal,
                    index: index,
                    selector: selector,
                    optimisticSeconds: 0,
                    recipient: systemContracts[destination][r],
                    deleteCalls: !isMultiCall,
                    formatCalls: true
                });
                if (!isMultiCall) {
                    // For a "single call" test, do a system call to a given recipient
                    if (origin != destination) {
                        // Cross-chain: construct a dispatched message and expect a Dispatch event
                        createDispatchedSystemCallMessage(origin, destination);
                        expectDispatch();
                    }
                    systemRouterCall({ index: 0 });
                }
            }
            if (isMultiCall) {
                // For a multi call test, do a system multi call to all the recipients
                formatSystemCalls();
                if (origin != destination) {
                    // For a cross-chain call:
                    // Construct a dispatched message and expect a Dispatch event
                    createDispatchedSystemCallMessage(origin, destination);
                    expectDispatch();
                }
                systemRouterMultiCall();
            }
        }
        if (origin != destination) {
            // Create proofs for later execution
            proofGen.createTree(messageHashes);
        }
    }

    function triggerTestWrapperMultiCall(
        uint32 domain,
        bytes4 selector,
        bool sameRecipient
    ) public {
        uint256 amount = systemContracts[domain].length;
        uint256 index = 0;
        for (uint256 s = 0; s < amount; ++s) {
            address sender = systemContracts[domain][s];
            // Context for a on-chain call
            MessageContext memory context = MessageContext(domain, sender, domain);
            deleteSystemCalls();
            for (uint256 r = 0; r < amount; (++r, ++index)) {
                createSystemCall({
                    context: context,
                    index: index,
                    selector: selector,
                    optimisticSeconds: 0,
                    recipient: systemContracts[domain][r],
                    deleteCalls: false,
                    formatCalls: false
                });
            }
            if (sameRecipient) {
                systemRouterMultiCallSameRecipient();
            } else {
                systemRouterMultiCallSameData();
            }
        }
    }

    // Prepares a test, where System Router is misconfigured (Origin address is missing).
    // Creates a data for system call to Origin, which should revert upon execution,
    // as SystemRouter is unaware of the Origin address.
    function prepareMisconfiguredTest(uint32 origin, uint32 destination) public {
        // Deploy a test router with Origin not configured
        SystemRouterHarness systemRouter = new SystemRouterHarness({
            _localDomain: DOMAIN_LOCAL,
            _origin: address(0),
            _destination: address(suiteDestination(destination))
        });
        vm.prank(owner);
        suiteDestination(destination).setSystemRouter(systemRouter);
        chains[destination].systemRouter = systemRouter;
        // Create test payloads:
        // Destination (from "origin" domain) is the caller
        MessageContext memory context = MessageContext({
            origin: origin,
            sender: address(suiteDestination(origin)),
            destination: destination
        });
        // Origin (from "destination" domain) is the recipient
        createSystemCall({ context: context, recipient: address(suiteOrigin(destination)) });
        // SystemRouter doesn't have Origin configured, so
        // both on-chain and cross-chain calls to Origin will fail
    }

    /*╔══════════════════════════════════════════════════════════════════════╗*\
    ▏*║                            INTERNAL VIEWS                            ║*▕
    \*╚══════════════════════════════════════════════════════════════════════╝*/

    // Get address of a given system entity
    function _getSystemAddress(uint32 domain, ISystemRouter.SystemEntity entity)
        internal
        view
        returns (address)
    {
        if (entity == ISystemRouter.SystemEntity.Origin) {
            return address(suiteOrigin(domain));
        } else if (entity == ISystemRouter.SystemEntity.Destination) {
            return address(suiteDestination(domain));
        } else {
            revert("Unknown system entity");
        }
    }

    // Get system entity value for a given address
    function _getSystemEntity(uint32 domain, address addr)
        internal
        view
        returns (ISystemRouter.SystemEntity)
    {
        if (addr == address(suiteOrigin(domain))) {
            return ISystemRouter.SystemEntity.Origin;
        } else if (addr == address(suiteDestination(domain))) {
            return ISystemRouter.SystemEntity.Destination;
        } else {
            revert("Unknown system entity");
        }
    }

    // Check if address is a system entity
    function _isSystemEntity(uint32 domain, address addr) internal view returns (bool) {
        return addr == address(suiteOrigin(domain)) || addr == address(suiteDestination(domain));
    }

    // Create a mock data for system call tests. Data is different for every new call.
    function _createMockSensitiveValue(uint256 index) internal pure returns (uint256) {
        return 100 * (index + 1);
    }
}
