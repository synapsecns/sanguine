import { useState } from 'react'

import { DownArrow } from '@/components/icons/DownArrow'

export const DropdownMenu = ({ menuTitleElement, children }) => {
  const [open, setOpen] = useState<boolean>(false)

  const handleClick = () => {
    setOpen(!open)
  }

  return (
    <div className="relative">
      <div
        onClick={handleClick}
        className={`
          flex place-items-center justify-center
          px-2 py-1 rounded space-x-2 cursor-pointer
          hover:border-[--synapse-focus] hover:bg-[--synapse-select-bg]
        `}
        // style={{ background: 'var(--synapse-select-bg' }}
      >
        {menuTitleElement}
        <DownArrow />
      </div>

      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-0 border border-solid border-[--synapse-select-border] rounded shadow popover -right-1 list-none text-left text-sm
          `}
          style={{ background: 'var(--synapse-select-bg)' }}
        >
          {children}
        </ul>
      )}
    </div>
  )
}
