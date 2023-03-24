import { useState, useEffect } from 'react'

/**
 * @param {string} targetKey
 */
export function useKeyPress(targetKey: string) {
  const [keyPressed, setKeyPressed] = useState(false)
  // Add event listeners
  useEffect(() => {
    if (window) {
      window.addEventListener('keydown', downHandler)
      window.addEventListener('keyup', upHandler)

      // Remove event listeners on cleanup
      return cleanupFunc
    }
  }, [])

  // If pressed key is our target key then set to true
  const downHandler = ({ key }: { key: string }) => {
    if (key === targetKey) {
      setKeyPressed(true)
    }
  }

  // If released key is our target key then set to false
  const upHandler = ({ key }: { key: string }) => {
    if (key === targetKey) {
      setKeyPressed(false)
    }
  }

  const cleanupFunc = () => {
    window.removeEventListener('keydown', downHandler)
    window.removeEventListener('keyup', upHandler)
  }

  // Empty array should ensure that effect is only run on mount and unmount
  // State for keeping track of whether key is pressed
  return keyPressed
}
