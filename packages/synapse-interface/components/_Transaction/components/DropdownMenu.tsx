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
          rounded w-5 h-[21px] flex place-items-center justify-center
          bg-surface
          border border-solid border-[--synapse-border]
          hover:border-[--synapse-focus]
          cursor-pointer
        `}
      >
        <DownArrow />
      </div>

      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-0 bg-surface border border-solid border-tint rounded shadow popover -right-1 list-none text-left text-sm
          `}
        >
          {children}
        </ul>
      )}
    </div>
  )
}

export const MenuItem = ({
  text,
  link,
  onClick,
}: {
  text: string
  link: string
  onClick?: () => any
}) => {
  return (
    <li
      className={`
      rounded cursor-pointer
      border border-solid border-transparent
      hover:border-[--synapse-focus]
      active:opacity-40
    `}
    >
      {onClick ? (
        <div
          onClick={onClick}
          className={`
            block pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline after:content-['_↗'] after:text-xs after:text-[--synapse-secondary]
          `}
        >
          {text}
        </div>
      ) : (
        <a
          href={link ?? ''}
          onClick={onClick}
          target="_blank"
          rel="noreferrer"
          className={`
            block pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline after:content-['_↗'] after:text-xs after:text-[--synapse-secondary]
          `}
        >
          {text}
        </a>
      )}
    </li>
  )
}
