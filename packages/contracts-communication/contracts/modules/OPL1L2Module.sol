pragma solidity 0.8.20;

import "./IL1CrossDomainMessenger.sol";
import "../IInterchain.sol";

contract OPL1L2Module {
    address public L1CrossDomainMessenger;
    address public L2CrossDomainMessenger;
    address public L1OPModule;
    address public L2OPModule;
    address public Interchain;

    uint256 defaultGasLimit = 1_000_000;

    event MessageReceived(bytes transaction);

    constructor(
        address _l1CrossDomainMessenger,
        address _l1OPModule,
        address _l2CrossDomainMessenger,
        address _l2OPModule,
        address _interchain
    ) {
        L1CrossDomainMessenger = _l1CrossDomainMessenger;
        L1OPModule = _l1OPModule;
        L2CrossDomainMessenger = _l2CrossDomainMessenger;
        L2OPModule = _l2OPModule;
    }

    function setInterchain(address _interchain) public {
        Interchain = _interchain;
    }

    function setL1OPModule(address _l1OPModule) public {
        L1OPModule = _l1OPModule;
    }

    function setL2OPModule(address _l2OPModule) public {
        L2OPModule = _l2OPModule;
    }

    // TODO: Implemented for tests until we have OP-Specific Fee Estimation
    function estimateFee(uint256 dstChainId) public view returns (uint256) {
        return 0;
    }

    function sendModuleMessage(bytes calldata transaction) external payable {
        IL1CrossDomainMessenger(L1CrossDomainMessenger).sendMessage(
            L2OPModule,
            abi.encodeWithSelector(bytes4(keccak256("receiveMessageModule(bytes)")), transaction),
            uint32(defaultGasLimit)
        );
    }

    function receiveModuleMessage(bytes calldata transaction) external {
        // Check that message is sent from Optimism L2 Messenger
        require(
            msg.sender == L2CrossDomainMessenger, "OPL1L2Module: Only L2CrossDomainMessenger can call this function"
        );

        // Check that original message sender is OPL1Module
        // This trusts Optimism Messenger to be honest
        require(
            IL1CrossDomainMessenger(L2CrossDomainMessenger).xDomainMessageSender() == L1OPModule,
            "OPL1L2Module: Only L1OPModule can call this function"
        );

        // Now proceed to do things with the message
        // Mainly to notify Interchain of the transaction verification
        IInterchain(Interchain).interchainReceive(transaction);
        emit MessageReceived(transaction);
    }
}
