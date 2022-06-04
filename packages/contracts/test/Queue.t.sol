// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

// ============ Internal Imports ============
import { QueueLib } from "../contracts/libs/Queue.sol";
import { QueueManager } from "../contracts/Queue.sol";
// ============ External Imports ============
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "forge-std/Test.sol";

contract QueueTest is QueueManager, Test {
    using QueueLib for QueueLib.Queue;

    function initialize() public initializer {
        __QueueManager_initialize();
    }

    function setUp() public {
        initialize();
    }

    function test_enqueue(uint8 _items) public {
        vm.assume(_items > 1);
        for (uint8 i = 0; i < _items; i++) {
            if (i == _items - 1) {
                queue.enqueue(bytes32("Last test"));
            } else {
                queue.enqueue(bytes32("Test"));
            }
        }
        assertEq(queue.length(), _items);
        assertFalse(queue.isEmpty());
        assertEq(queue.peek(), bytes32("Test"));
        assertEq(queue.lastItem(), bytes32("Last test"));
    }

    function test_isEmpty() public {
        assert(queue.isEmpty());
    }

    function test_queueFirst() public {
        queue.enqueue(bytes32("Test 1"));
        queue.enqueue(bytes32("Test 2"));
        assertTrue(queue.contains(bytes32("Test 1")));
        assertTrue(queue.contains(bytes32("Test 2")));
        queue.dequeue();
        assertFalse(queue.contains(bytes32("Test 1")));
        assertTrue(queue.contains(bytes32("Test 2")));
        queue.dequeue();
        assertFalse(queue.contains(bytes32("Test 2")));
        assertEq(queue.length(), 0);
        assertTrue(queue.isEmpty());
    }

    function test_dequeueMultiple() public {
        queue.enqueue(bytes32("Test 1"));
        queue.enqueue(bytes32("Test 2"));
        queue.enqueue(bytes32("Test 3"));
        queue.enqueue(bytes32("Test 4"));
        assertTrue(queue.contains(bytes32("Test 1")));
        queue.dequeue(1);
        assertFalse((queue.contains(bytes32("Test 1"))));
        assertTrue((queue.contains(bytes32("Test 2"))));
        queue.dequeue(2);
        assertFalse((queue.contains(bytes32("Test 2"))));
        assertFalse((queue.contains(bytes32("Test 3"))));
        assertTrue((queue.contains(bytes32("Test 4"))));
        assertEq(queue.length(), 1);
        queue.dequeue(1);
        assertEq(queue.length(), 0);
    }
}
