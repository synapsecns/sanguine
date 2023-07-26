import React, { useEffect, useState } from 'react'

const useIsMobileScreen = (): boolean => {
  const [width, setWidth] = useState<number>(window.innerWidth)

  const handleWindowSizeChange = () => {
    setWidth(window.innerWidth)
  }

  useEffect(() => {
    window.addEventListener('resize', handleWindowSizeChange)
    return () => {
      window.removeEventListener('resize', handleWindowSizeChange)
    }
  }, [])

  return width <= 768
}

export default useIsMobileScreen
