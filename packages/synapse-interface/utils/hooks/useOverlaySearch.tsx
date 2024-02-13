import { useCallback, useEffect, useState } from "react"
import { useKeyPress } from "@/utils/hooks/useKeyPress"
import useCloseOnOutsideClick from "./useCloseOnOutsideClick"
import { useCloseOutsideRef } from "@/utils/hooks/useCloseOutsideRef"


export function useOverlaySearch(masterListLength, closeOverlayDispatchFunc) {
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')


//   const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setSearchStr('')
    closeOverlayDispatchFunc()
  }, [closeOverlayDispatchFunc])

  const overlayRef = useCloseOutsideRef(onClose)

//   const escFunc = () => {
//     if (escPressed) {
//       onClose()
//     }
//   }
  const arrowDownFunc = () => {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < masterListLength) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(arrowDownFunc, [arrowDown])
//   useEffect(escFunc, [escPressed])
  useEffect(arrowUpFunc, [arrowUp])
//   useCloseOnOutsideClick(overlayRef, onClose)
  return {
    overlayRef,
    onSearch,
    // arrowDownFunc,
    // arrowUpFunc,
    currentIdx,
    searchStr,
    onClose
  }
}