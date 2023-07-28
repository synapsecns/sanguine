import { useEffect, useState, useRef } from 'react'

const useElementWidth = (ref) => {
  const [elementWidth, setElementWidth] = useState(0)

  const updateElementWidth = () => {
    console.log('got hit')
    if (ref.current) {
      setElementWidth(ref.current.clientWidth)
    }
  }

  useEffect(() => {
    updateElementWidth() // Initial update
    window.addEventListener('resize', updateElementWidth)
    window.addEventListener('click', updateElementWidth)

    // Clean up the event listener on component unmount
    return () => {
      window.removeEventListener('resize', updateElementWidth)
      window.removeEventListener('click', updateElementWidth)
    }
  }, [ref])

  return elementWidth
}

export default useElementWidth
