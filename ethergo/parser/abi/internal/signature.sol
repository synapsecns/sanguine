// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

contract TestSignature {
    /**
     * @dev Returns the signature of the function `testSignature()`.
     *
     * This is a public function that is not present in the contract ABI. It is
     * used to test that the signature of a function is correctly computed.
     */
    function testSignature() public pure returns (bytes4) {
        return this.testSignature.selector;
    }


    /**
    * @dev Returns the signature of the function `testSignatureArgs(int, int)`.
    *
    * This is a public function that is not present in the contract ABI. It is
    * used to test that the signature of a function is correctly computed.
    */
    function testSignatureArgs(int a, int b) public pure returns (bytes4) {
        return this.testSignatureArgs.selector;
    }

    /**
    * @dev Returns the signature of the function `testSignatureOverload(int, int)`.
    *
    */
    function testSignatureOverload(int a, int b) public pure returns (bytes4) {
        return bytes4(keccak256("testSignatureOverload(int256,int256)"));
    }

    /**
    * @dev Returns the signature of the function `testSignatureOverload(int, int, int)`.
    *
    */
    function testSignatureOverload(int a, int b, int c) public pure returns (bytes4) {
        return bytes4(keccak256("testSignatureOverload(int256,int256,int256)"));
    }
}
