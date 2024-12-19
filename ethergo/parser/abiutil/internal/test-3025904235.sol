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

    // ============ Modifiers ============

    function doSomethingWithoutParams() external  {
        // I have no params!
    }

    event SomethingHappened(int a, int b);

    /**
    * @dev Returns the signature of the function `testSignatureOverload(int, int)`.
    *
    */
    function doSomething(int a, int b) external {
        emit SomethingHappened(a, b);
    }


    event SomethingHappenedOverload0(int a, int b);

    /**
    * @dev Overload of `doSomething(int, int)`.
    *
    */
    function doSomethingOverload(int a, int b) external {
        emit SomethingHappenedOverload0(a, b);
    }

    event SomethingHappenedOverload1(int a, address b);

    /**
    * @dev Overload of `doSomethingOverload(int, int)`.
    *
    */
    function doSomethingOverload(int a, address b) external {
        emit SomethingHappenedOverload1(a, b);
    }

    event SomethingHappenedManyTypes(
        int a,
        address b,
        uint256 c,
        bytes32 d,
        bytes e,
        bool f
    );

    function doSomethingManyTypes(
        int a,
        address b,
        uint256 c,
        bytes32 d,
        bytes memory e,
        bool f
    ) external  {
        emit SomethingHappenedManyTypes(a, b, c, d, e, f);
    }
}
