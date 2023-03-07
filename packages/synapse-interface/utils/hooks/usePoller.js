import { useEffect, useRef } from 'react'

export const usePoller = (fn, delay) => {
  const savedCallback = useRef()

  // Remember the latest fn.
  useEffect(() => {
    savedCallback.current = fn
  }, [fn])

  function tick() {
    savedCallback.current?.()
  }

  // Set up the interval.
  useEffect(() => {
    if (delay !== null) {
      const id = setInterval(tick, delay)

      return () => clearInterval(id)
    }
  }, [delay])

  // run at start too
  useEffect(() => {
    fn()
  }, [])
}

