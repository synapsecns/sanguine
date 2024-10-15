import { useRef, useState } from 'react'

import { DownArrow } from '@/components/icons/DownArrow'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { useCloseOnEscape } from '@/utils/hooks/useCloseOnEscape'

export const DropdownMenu = ({ menuTitleElement, children }) => {
  const [open, setOpen] = useState<boolean>(false)
  const ref = useRef(null)

  const handleClick = () => {
    setOpen(!open)
  }

  const closeDropdown = () => setOpen(false)

  useCloseOnOutsideClick(ref, closeDropdown)
  useCloseOnEscape(closeDropdown)

  return (
    <div className="relative" ref={ref}>
      <div
        onClick={handleClick}
        className={`
          flex place-items-center justify-center
          px-2 py-1 rounded space-x-2 cursor-pointer
        `}
      >
        {menuTitleElement}
        <DownArrow />
      </div>

      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-0 bg-zinc-100 dark:bg-bgBase rounded shadow popover -right-1 list-none text-left text-sm
          `}
        >
          {children}
        </ul>
      )}
    </div>
  )
}
