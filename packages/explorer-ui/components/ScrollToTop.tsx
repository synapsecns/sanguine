import { useEffect } from 'react'

export const ScrollToTop = () => {
  const pathname = ''

  useEffect(() => {
    window.scrollTo(0, 0)
  }, [pathname])

  return null
}
