import { useCallback, useEffect, useState } from "react"
import { useKeyPress } from "@/utils/hooks/useKeyPress"
import { useCloseOutsideRef } from "@/utils/hooks/useCloseOutsideRef"

/**
 * Custom hook for managing overlay search functionality.
 *
 * This hook provides functionality to navigate through a list using arrow keys, search within the list,
 * and close the overlay using custom hooks and state management. It manages the current index for navigation,
 * the search string for filtering, and provides a ref to handle closing the overlay when clicking outside or pressing Escape.
 *
 * @param {number} masterListLength - The total number of items in the list to navigate through.
 * @param {Function} closeOverlayDispatchFunc - A callback function to be called when the overlay needs to be closed.
 *
 * @returns {Object} An object containing:
 * - overlayRef: A ref to be attached to the overlay component for detecting outside clicks.
 * - onSearch: A function to update the search string and reset the current index.
 * - currentIdx: The current index in the list based on arrow key navigation.
 * - searchStr: The current search string for filtering the list.
 * - onClose: A function to close the overlay and reset the state.
 */
export function useOverlaySearch(masterListLength, closeOverlayDispatchFunc) {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')

  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setSearchStr('')
    closeOverlayDispatchFunc()
  }, [closeOverlayDispatchFunc])

  const overlayRef = useCloseOutsideRef(onClose)

  function arrowDownFunc() {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < masterListLength) {
      setCurrentIdx(nextIdx)
    }
  }

  function arrowUpFunc() {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  function onSearch(str: string) {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(arrowUpFunc, [arrowUp])

  return {
    overlayRef,
    onSearch,
    currentIdx,
    searchStr,
    onClose
  }
}