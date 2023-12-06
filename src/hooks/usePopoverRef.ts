import { useState, useRef, useEffect } from 'react'

function usePopover() {
  const [isOpen, setIsOpen] = useState(false)
  const popoverRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    function handleClickOutside(event) {
      if (popoverRef.current && !popoverRef.current.contains(event.target)) {
        setIsOpen(false)
      }
    }

    document.addEventListener('mousedown', handleClickOutside)
    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
    }
  }, [isOpen])

  const togglePopover = () => {
    setIsOpen(!isOpen)
  }

  const closePopover = () => {
    setIsOpen(false)
  }

  return { popoverRef, isOpen, togglePopover, closePopover }
}

export default usePopover
