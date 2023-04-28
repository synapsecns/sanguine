import { useState, useEffect } from 'react'

export const useMouseClickNavigation = () => {
  const [clickPosition, setClickPosition] = useState(null)

  useEffect(() => {
    if (clickPosition) {
      window.scrollTo({
        top: clickPosition,
        behavior: 'smooth',
      })
    }
  }, [clickPosition])

  const handleClick = (event: MouseEvent) => {
    setClickPosition(event.clientY)
  }

  return handleClick
}
