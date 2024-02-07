import { useState } from 'react'
import { DownArrow } from '@/components/icons/DownArrow'

export const DropdownMenu = ({ children }) => {
  const [open, setOpen] = useState<boolean>(false)

  const handleClick = () => {
    setOpen(!open)
  }

  return (
    <div id="dropdown-menu" className="relative">
      <div
        onClick={handleClick}
        className={`
          flex place-items-center justify-center
          w-5 h-[21px] bg-surface rounded
          border border-zinc-700
          cursor-pointer hover:border-zinc-500
        `}
      >
        Waiting…
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
