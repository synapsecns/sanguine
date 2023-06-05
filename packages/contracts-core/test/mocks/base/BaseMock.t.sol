// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

abstract contract BaseMock {
    bytes private _mockReturnValue;

    /// @notice Prevents this contract from being included in the coverage report
    function testBaseMock() external {}

    function deleteMockReturnValue() external {
        delete _mockReturnValue;
    }

    function setMockReturnValue(uint256 value) external {
        _mockReturnValue = abi.encode(value);
    }

    function setMockReturnValue(address value) external {
        _mockReturnValue = abi.encode(value);
    }

    function setMockReturnValue(bytes memory value) external {
        _mockReturnValue = value;
    }

    function getReturnValueBytes() public view returns (bytes memory) {
        return _mockReturnValue;
    }

    function getReturnValueAddress() public view returns (address) {
        return _mockReturnValue.length == 0 ? address(0) : abi.decode(_mockReturnValue, (address));
    }

    function getReturnValueBool() public view returns (bool) {
        return _mockReturnValue.length == 0 ? false : abi.decode(_mockReturnValue, (bool));
    }

    function getReturnValueUint() public view returns (uint256) {
        return _mockReturnValue.length == 0 ? 0 : abi.decode(_mockReturnValue, (uint256));
    }

    function getReturnValueUint32() public view returns (uint32) {
        return _mockReturnValue.length == 0 ? 0 : abi.decode(_mockReturnValue, (uint32));
    }
}
