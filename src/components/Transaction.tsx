import { useState } from 'react'
import { DownArrow } from './icons/DownArrow'

const TransactionStatus = ({ string }) => {
  return <div>{string}</div>
}

export const Transaction = () => {
  return (
    <div
      data-test-id="transaction"
      className={`
        flex flex-row justify-between px-3 py-1
        bg-[--synapse-bg-surface]
        border border-[--synapse-border] rounded-md
      `}
    >
      <TransactionStatus string="Pending" />
      <div className="flex flex-row items-center space-x-2">
        <div>5-7 min</div>
        <DropdownMenu>
          <MenuItem text="Etherscan" link="" />
          <MenuItem text="Arbiscan" link="" />
          <MenuItem text="Synapse Explorer" link="" />
          <MenuItem text="Contact Support" link="" />
        </DropdownMenu>
      </div>
    </div>
  )
}

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
          p-1 cursor-pointer
          bg-[--synapse-bg-select]
          border border-[--synapse-border] rounded-md
        `}
      >
        <DownArrow />
      </div>

      {open && (
        <div
          className={`
            flex flex-col 
            absolute right-0 z-10
            bg-[--synapse-bg-select] 
            border border-[--synapse-border]
          `}
        >
          {children}
        </div>
      )}
    </div>
  )
}

export const MenuItem = ({ text, link }: { text: string; link: string }) => {
  return (
    <a
      href={link ?? ''}
      target="_blank"
      rel="noreferrer"
      className="flex whitespace-nowrap"
    >
      {text}
    </a>
  )
}
