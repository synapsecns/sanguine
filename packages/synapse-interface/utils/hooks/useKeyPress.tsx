import { useState, useEffect } from "react"

/**
 * @param {string} targetKey
 */
export function useKeyPress(targetKey: string ) {
  const [keyPressed, setKeyPressed] = useState(false)

  // If pressed key is our target key then set to true
  function downHandler({ key }: {key: string}) {
    if (key === targetKey) {
      setKeyPressed(true)
    }
  }

  // If released key is our target key then set to false
  function upHandler({ key }: {key: string}) {
    if (key === targetKey) {
      setKeyPressed(false)
    }
  }

  function cleanupFunc() {
    window.removeEventListener("keydown", downHandler)
    window.removeEventListener("keyup", upHandler)
  }

  function func() {
    if (window) {
      window.addEventListener("keydown", downHandler)
      window.addEventListener("keyup", upHandler)

      // Remove event listeners on cleanup
      return cleanupFunc
    }
  }
  // Add event listeners
  useEffect(func, [])
  // Empty array should ensure that effect is only run on mount and unmount
  // State for keeping track of whether key is pressed
  return keyPressed
}


