// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

/**
 * @notice Based on OpenZeppelin's implementation:
 * https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/structs/DoubleEndedQueue.sol
 */
library QueueLib {
    /**
     * @dev An operation (e.g. {front}) couldn't be completed due to the queue being empty.
     */
    error Empty();

    /**
     * @dev An operation (e.g. {at}) couldn't be completed due to an index being out of bounds.
     */
    error OutOfBounds();

    /**
     * @dev Indices are unsigned integers because the queue can grow only in one direction. They are 128 bits so begin and end
     * are packed in a single storage slot for efficient access. Since the items are added one at a time we can safely
     * assume that these 128-bit indices will not overflow, and use unchecked arithmetic.
     *
     * Struct members have an underscore prefix indicating that they are "private" and should not be read or written to
     * directly. Use the functions provided below instead. Modifying the struct manually may violate assumptions and
     * lead to unexpected behavior.
     *
     * Indices are in the range (begin, end] which means the first item is at data[begin + 1] and the last item is at
     * data[end].
     *
     * This also means that indices are always greater than zero.
     */
    struct Queue {
        uint128 _begin;
        uint128 _end;
        mapping(uint256 => bytes32) _data;
    }

    /**
     * @dev Inserts an item at the end of the queue.
     *      OZ analog: pushBack
     */
    function enqueue(Queue storage queue, bytes32 value) internal {
        uint128 backIndex;
        unchecked {
            // (begin; end] -> last item is queue._end (insert after that)
            backIndex = queue._end + 1;
        }
        queue._data[backIndex] = value;
        queue._end = backIndex;
    }

    /**
     * @dev Removes the item at the beginning of the queue and returns it.
     *      OZ analog: popFront
     * Reverts with `Empty` if the queue is empty.
     */
    function dequeue(Queue storage queue) internal returns (bytes32 value) {
        if (isEmpty(queue)) revert Empty();
        uint128 frontIndex;
        unchecked {
            // (begin; end] -> first item is queue._begin + 1
            frontIndex = queue._begin + 1;
        }
        value = queue._data[frontIndex];
        delete queue._data[frontIndex];
        queue._begin = frontIndex;
    }

    /**
     * @dev Batch inserts several items at the end of the queue.
     *      OZ analog: pushBack
     */
    function enqueue(Queue storage queue, bytes32[] memory values) internal {
        uint128 backIndex = queue._end;
        uint256 len = values.length;
        for (uint256 i = 0; i < len; ) {
            // (begin; end] -> last item is backIndex (insert after that)
            unchecked {
                ++backIndex;
            }
            queue._data[backIndex] = values[i];
            unchecked {
                ++i;
            }
        }
        // backIndex is now pointing to the last inserted item
        queue._end = backIndex;
    }

    /**
     * @dev Batch removes `number` items at the beginning of the queue and returns them.
     *      OZ analog: popFront
     * Reverts with `Empty` if `number` > queue length
     */
    function dequeue(Queue storage queue, uint256 number)
        internal
        returns (bytes32[] memory values)
    {
        uint128 frontIndex = queue._begin;
        unchecked {
            // This will not underflow assuming all queue operations were done through interface
            if (number > queue._end - frontIndex) revert Empty();
        }
        values = new bytes32[](number);
        for (uint256 i = 0; i < number; ) {
            // (begin; end] -> first item is frontIndex + 1
            unchecked {
                ++frontIndex;
            }
            values[i] = queue._data[frontIndex];
            delete queue._data[frontIndex];
            unchecked {
                ++i;
            }
        }
        // firstIndex is now pointing to the last deleted item
        queue._begin = frontIndex;
    }

    function contains(Queue storage queue, bytes32 item) internal view returns (bool) {
        // Most of the time we'll be checking a merkle root that has been recently added,
        // so checking from back to front is likely to find it faster.
        uint128 backIndex = queue._end;
        uint128 frontIndex = queue._begin;
        // elements are stored at (begin, end] range,
        // so we need to start from queue._end
        // and finish at queue._begin + 1
        for (; backIndex > frontIndex; ) {
            if (queue._data[backIndex] == item) return true;
            unchecked {
                --backIndex;
            }
        }
        return false;
    }

    /// @notice Returns last item in queue
    /// @dev Returns bytes32(0) if queue is empty
    function lastItem(Queue storage queue) internal view returns (bytes32 item) {
        // (begin; end] -> last item is queue._end
        return queue._data[queue._end];
    }

    /// @notice Returns element at front of queue without removing element
    /// @dev Reverts if queue is empty
    function peek(Queue storage queue) internal view returns (bytes32 item) {
        if (isEmpty(queue)) revert Empty();
        unchecked {
            // (begin; end] -> first item is queue._begin + 1
            item = queue._data[queue._begin + 1];
        }
    }

    /// @notice Returns true if queue is empty and false if otherwise
    function isEmpty(Queue storage queue) internal view returns (bool) {
        // (begin; end] -> (end <= begin) is empty queue
        return queue._end <= queue._begin;
    }

    function length(Queue storage queue) internal view returns (uint256) {
        // The interface preserves the invariant that begin <= end so we assume this will not overflow.
        unchecked {
            return queue._end - queue._begin;
        }
    }
}
