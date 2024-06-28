import React, { useState, useEffect } from 'react'

export const useIsFocused = (ref: React.RefObject<HTMLInputElement>) => {
  const [isFocused, setIsFocused] = useState<boolean>(false)

  useEffect(() => {
    const handleFocus = () => setIsFocused(true)
    const handleBlur = () => setIsFocused(false)
    const currentRef = ref.current

    if (currentRef) {
      currentRef.addEventListener('focus', handleFocus)
      currentRef.addEventListener('blur', handleBlur)
      return () => {
        currentRef.removeEventListener('focus', handleFocus)
        currentRef.removeEventListener('blur', handleBlur)
      }
    }
  }, [ref])

  return isFocused
}
