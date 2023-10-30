// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

contract CalleeMock {
    uint256 public secret;

    /// @notice Prevents this contract from being included in the coverage report
    function testCalleeMock() external {}

    function setSecret(uint256 _secret) external payable {
        secret = _secret;
    }
}

contract CalleeReturnDataMock {
    uint256 public secret;

    /// @notice Prevents this contract from being included in the coverage report
    function testCalleeReturnDataMock() external {}

    function setSecret(uint256 _secret) external payable returns (bool) {
        secret = _secret;
        return true;
    }
}
