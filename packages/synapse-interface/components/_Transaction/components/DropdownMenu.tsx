import { useState } from 'react'
import { DownArrow } from '@/components/icons/DownArrow'

export const DropdownMenu = ({ menuTitleElement, children }) => {
  const [open, setOpen] = useState<boolean>(false)

  const handleClick = () => {
    setOpen(!open)
  }

  return (
    <div id="dropdown-menu" className="relative">
      <div
        onClick={handleClick}
        className={`
          flex w-fit px-2 py-0.5 space-x-1
          relative place-items-center justify-center
          rounded cursor-pointer
        hover:bg-zinc-700
        `}
      >
        {menuTitleElement}
        <DownArrow />
      </div>

      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-0 -right-1 bg-surface
            border border-solid border-tint rounded shadow
            popover list-none text-left text-sm
          `}
        >
          {children}
        </ul>
      )}
    </div>
  )
}
