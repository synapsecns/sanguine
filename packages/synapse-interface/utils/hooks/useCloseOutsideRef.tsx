import { useEffect, useRef } from 'react'
import { useKeyPress } from '@/utils/hooks/useKeyPress'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'

/**
 * Hook to handle closing a component when clicking outside of it or pressing the Escape key.
 *
 * This hook integrates two hooks that are commonly paired: `useKeyPress` to detect when the Escape key is pressed,
 * and `useCloseOnOutsideClick` to handle closing the component when a click occurs outside of it.
 * It returns a ref that should be attached to the component you want to apply this behavior to.
 *
 * @param {Function} onClose - A callback function that is executed when the Escape key is pressed or a click outside the component is detected.
 * @returns {React.MutableRefObject<null>} A ref object that should be attached to the component to monitor for outside clicks or Escape key presses.
 */
export function useCloseOutsideRef(onClose: () => void) {
  const ref = useRef(null)
  const escPressed = useKeyPress('Escape')

  function escFunc() {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escFunc, [escPressed])
  useCloseOnOutsideClick(ref, onClose)

  return ref //as React.MutableRefObject<null>
}

