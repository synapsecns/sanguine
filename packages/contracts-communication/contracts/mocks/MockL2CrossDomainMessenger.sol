pragma solidity 0.8.20;

contract MockL2CrossDomainMessenger {
    address internal xDomainMsgSender;

    event MessageRelayed(
        uint256 nonce, address sender, address target, uint256 value, uint256 minGasLimit, bytes message
    );

    function relayMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _minGasLimit,
        bytes calldata _message
    )
        external
        payable
    {
        xDomainMsgSender = _sender;

        _target.call{value: _value, gas: _minGasLimit}(_message);

        emit MessageRelayed(_nonce, _sender, _target, _value, _minGasLimit, _message);
        xDomainMsgSender = address(0);
    }

    function xDomainMessageSender() public view returns (address) {
        return xDomainMsgSender;
    }
}
