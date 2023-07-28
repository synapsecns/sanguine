import { useEffect, useState, useRef } from 'react'

const useElementWidth = (ref) => {
  const [elementWidth, setElementWidth] = useState(0)

  const updateElementWidth = () => {
    if (ref.current) {
      setElementWidth(ref.current.clientWidth)
    }
  }

  useEffect(() => {
    updateElementWidth() // Initial update
    window.addEventListener('resize', updateElementWidth)

    // Clean up the event listener on component unmount
    return () => {
      window.removeEventListener('resize', updateElementWidth)
    }
  }, [ref])

  return elementWidth
}

export default useElementWidth
