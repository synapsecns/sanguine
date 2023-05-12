// SPDX-License-Identifier: MIT
pragma solidity 0.8.17;

import {SynapseLibraryTest} from "../../../utils/SynapseLibraryTest.t.sol";
import {Number, NumberHarness} from "../../../harnesses/libs/stack/NumberHarness.t.sol";

// solhint-disable func-name-mixedcase
contract NumberLibraryTest is SynapseLibraryTest {
    NumberHarness internal libHarness;

    function setUp() public {
        libHarness = new NumberHarness();
    }

    function test_compress(uint256 value) public {
        // decompress(compress(number)) could lead to precision loss
        Number num = libHarness.compress(value);
        uint256 decompressed = libHarness.decompress(num);
        // Should never exceed the initial value
        assertLe(decompressed, value, "decompressed too big");
        // The difference should be less the initial value / 256 (under 0.4%)
        if (decompressed < value) {
            assertLt((value - decompressed) * 256, value, "error too big");
        }
    }

    function test_parity_exponentNotFF(uint256 mantissa, uint256 exponent) public {
        mantissa = bound(mantissa, 0, 255);
        exponent = bound(exponent, 0, 247);
        uint256 value = (256 + mantissa) << exponent;
        require(value >> exponent == 256 + mantissa, "setup overflow");
        check_decompress({mantissa: mantissa, exponent: exponent, value: value});
    }

    function test_parity_exponentFF(uint256 value) public {
        value = bound(value, 0, 255);
        check_decompress({mantissa: value, exponent: 0xFF, value: value});
    }

    function check_decompress(uint256 mantissa, uint256 exponent, uint256 value) public {
        Number compressed = libHarness.compress(value);
        assertEq(uint8(Number.unwrap(compressed) >> 8), mantissa);
        assertEq(uint8(Number.unwrap(compressed)), exponent);
        assertEq(libHarness.decompress(compressed), value);
    }

    function test_mostSignificantBit(uint256 value) public {
        uint256 msb = libHarness.mostSignificantBit(value);
        if (value == 0) {
            assertEq(msb, 0);
        } else {
            // Check `2**msb <= value < 2**(msb + 1)`
            assertLe(msb, 255);
            assertLe(2 ** msb, value);
            if (msb != 255) {
                assertGt(2 ** (msb + 1), value);
            }
        }
    }
}
