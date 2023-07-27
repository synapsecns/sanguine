import React, { useEffect, useState } from 'react'

const isClient = typeof window === 'object'

export default function useIsMobileScreen(): boolean {
  const [width, setWidth] = useState<number>(
    isClient ? window.innerWidth : 1000
  )

  const handleWindowSizeChange = () => {
    setWidth(window.innerWidth)
  }

  useEffect(() => {
    if (isClient) {
      window.addEventListener('resize', handleWindowSizeChange)

      return () => {
        window.removeEventListener('resize', handleWindowSizeChange)
      }
    }
  }, [])

  return isClient ? width <= 768 : false
}
