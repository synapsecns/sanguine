import { useState, useEffect } from 'react'

/**
 * Checks when app has mounted
 */
export const useIsMounted = (): boolean => {
  const [mounted, setMounted] = useState<boolean>(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  return mounted
}
