import { useState } from 'react'
import { DownArrow } from '@/components/icons/DownArrow'

export const DropdownMenu = ({ children }) => {
  const [open, setOpen] = useState<boolean>(false)

  const handleClick = () => {
    setOpen(!open)
  }

  return (
    <div className="relative">
      <div
        onClick={handleClick}
        className={`
          rounded w-5 h-[21px] flex place-items-center justify-center
          border border-solid border-[--synapse-border]
          hover:border-[--synapse-focus]
          cursor-pointer
        `}
        style={{background: 'var(--synapse-select-bg'}}
      >
        <DownArrow />
      </div>

      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-0 border border-solid border-[--synapse-select-border] rounded shadow popover -right-1 list-none text-left text-sm
          `}
          style={{background: 'var(--synapse-select-bg)'}}
        >
          {children}
        </ul>
      )}
    </div>
  )
}
