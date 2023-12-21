import { useState } from 'react'
import { DownArrow } from './icons/DownArrow'
import { getTxBlockExplorerLink } from '@/utils/getTxBlockExplorerLink'
import { getTxSynapseExplorerLink } from '@/utils/getTxSynapseExplorerLink'

const TransactionStatus = ({ string }) => {
  return <>{string}</>
}

export const Transaction = ({
  originChainId,
  destinationChainId,
  originTxHash,
  destinationTxHash,
  kappa,
  timestamp,
}: {
  originChainId: number
  destinationChainId: number
  originTxHash?: string
  destinationTxHash?: string
  kappa?: string
  timestamp?: string
}) => {
  const originTxExplorerLink = getTxBlockExplorerLink(
    originChainId,
    originTxHash
  )
  const destTxExplorerLink = getTxBlockExplorerLink(
    destinationChainId,
    destinationTxHash
  )
  const synapseExplorerLink = getTxSynapseExplorerLink({
    originChainId,
    destinationChainId,
    txHash: originTxHash,
    kappa,
  })

  return (
    <div
      data-test-id="transaction"
      className={`
        flex flex-wrap-reverse gap-1 justify-end items-center pl-2.5 pr-1.5 py-1
        bg-[--synapse-surface]
        border border-solid border-[--synapse-border] rounded-md
      `}
    >
      <TransactionStatus string="Pending" />
      <div className="flex gap-2 items-center grow justify-end">
        <div className="whitespace-nowrap">5-7 min</div>
        <DropdownMenu>
          <MenuItem text="Arbiscan" link={originTxExplorerLink} />
          <MenuItem text="Polyscan" link={destTxExplorerLink} />
          <MenuItem text="Synapse Explorer" link={synapseExplorerLink} />
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
          rounded w-5 h-[21px] flex place-items-center justify-center
          bg-[--synapse-select-bg]
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
            absolute z-50 mt-1 p-0 bg-[--synapse-surface] border border-solid border-[--synapse-border] rounded shadow popover -right-1 list-none text-left text-sm
          `}
        >
          {children}
        </ul>
      )}
    </div>
  )
}

export const MenuItem = ({ text, link }: { text: string; link: string }) => {
  return (
    <li className={`
      rounded cursor-pointer
      border border-solid border-transparent
      hover:border-[--synapse-focus]
      active:opacity-40
    `}>
      <a
        href={link ?? ''}
        target="_blank"
        rel="noreferrer"
        className={`
          block pl-2 pr-3 py-2 whitespace-nowrap text-[--synapse-text-primary] no-underline after:content-['_↗'] after:text-xs after:text-[--synapse-secondary]
        `}
      >
        {text}
      </a>
    </li>
  )
}
