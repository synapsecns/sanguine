import { useState } from 'react'
import { DownArrow } from './icons/DownArrow'
import { getTxBlockExplorerLink } from '@/utils/getTxBlockExplorerLink'
import { getTxSynapseExplorerLink } from '@/utils/getTxSynapseExplorerLink'

const TransactionStatus = ({ string }) => {
  return <div className="px-1">{string}</div>
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
        flex flex-row justify-between items-center px-2 py-1
        bg-[--synapse-bg-surface]
        border border-solid border-[--synapse-border] rounded-md
      `}
    >
      <TransactionStatus string="Pending" />
      <div className="flex flex-row items-center space-x-2">
        <div>5-7 min</div>
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
          p-1 cursor-pointer
          bg-[--synapse-bg-select]
          border border-solid border-[--synapse-border] hover:border-[--synapse-border-hover] rounded-md w-4 h-4 flex place-items-center justify-center
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
            border border-solid border-[--synapse-border]
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
