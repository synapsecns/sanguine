pragma solidity 0.8.20;

interface IL2CrossDomainMessenger {
    function relayMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _minGasLimit,
        bytes calldata _message
    )
        external
        payable;

    function xDomainMessageSender() external returns (address);
}
