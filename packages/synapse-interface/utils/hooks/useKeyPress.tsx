import { useState, useEffect } from 'react'

/**
 * Hook that detects when a specific key is pressed, with focus awareness.
 * @param {string} targetKey - The key to detect.
 * @param {boolean} isActive - Whether the hook should be active (e.g., the component is focused).
 */
export function useKeyPress(targetKey: string, isActive: boolean) {
  const [keyPressed, setKeyPressed] = useState(false)

  useEffect(() => {
    // Only add event listeners if the hook is active
    if (isActive && window) {
      const downHandler = ({ key }: { key: string }) => {
        if (key === targetKey) {
          setKeyPressed(true)
        }
      }

      const upHandler = ({ key }: { key: string }) => {
        if (key === targetKey) {
          setKeyPressed(false)
        }
      }

      window.addEventListener('keydown', downHandler)
      window.addEventListener('keyup', upHandler)

      // Cleanup function to remove event listeners
      return () => {
        window.removeEventListener('keydown', downHandler)
        window.removeEventListener('keyup', upHandler)
      }
    }
  }, [targetKey, isActive]) // Re-run the effect if targetKey or isActive changes

  return keyPressed
}
