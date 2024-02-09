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
          w-5 h-[21px] rounded
          border border-solid border-white/10
          cursor-pointer hover:border-white/80
        `}
      >
        <DownArrow />
      </div>

      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-0 -right-1
            border border-white/20 bg-bgBase/10 backdrop-blur-lg
            rounded-md overflow-hidden shadow-md
            popover list-none text-left text-sm
          `}
        >
          {children}
        </ul>
      )}
    </div>
  )
}
