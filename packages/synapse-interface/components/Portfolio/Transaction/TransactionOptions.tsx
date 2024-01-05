import { Fragment, useCallback } from 'react'
import Image from 'next/image'
import { Address } from 'viem'
import { Popover, Transition } from '@headlessui/react'
import { Chain } from '@/utils/types'
import { TransactionStatus } from './Transaction'
import {
  getTransactionExplorerLink,
  getTransactionHashExplorerLink,
} from './components/TransactionExplorerLink'
import { getExplorerAddressUrl, getExplorerTxUrl } from '@/constants/urls'
import { DISCORD_URL } from '@/constants/urls'
import Button from '../../ui/tailwind/Button'
import SynapseLogo from '@assets/icons/syn.svg'
import DiscordIcon from '../../icons/DiscordIcon'
import { DownArrow } from '@/components/icons/DownArrow'

export const TransactionOptions = ({
  connectedAddress,
  originChain,
  destinationChain,
  kappa,
  transactionHash,
  transactionStatus,
  isDelayed,
}: {
  connectedAddress: Address
  originChain: Chain
  destinationChain: Chain
  kappa?: string
  transactionHash?: string
  transactionStatus: TransactionStatus
  isDelayed: boolean
}) => {
  const handleOriginExplorerLink: () => void = useCallback(() => {
    if (originChain && connectedAddress) {
      const explorerLink: string = getExplorerAddressUrl({
        chainId: originChain.id,
        address: connectedAddress,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
  }, [originChain, connectedAddress])

  const handleDestinationExplorerLink: () => void = useCallback(() => {
    if (destinationChain && connectedAddress) {
      const explorerLink: string = getExplorerAddressUrl({
        chainId: destinationChain.id,
        address: connectedAddress,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
  }, [destinationChain, connectedAddress])

  const handleSynapseExplorerLink: () => void = useCallback(() => {
    if (kappa) {
      const explorerLink: string = getTransactionExplorerLink({
        kappa,
        fromChainId: originChain?.id,
        toChainId: destinationChain?.id,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    } else if (transactionHash) {
      const explorerLink: string = getTransactionHashExplorerLink({
        transactionHash: transactionHash,
      })
      window.open(explorerLink, '_blank', 'noopener,noreferrer')
    }
  }, [kappa, originChain, transactionStatus, isDelayed, transactionHash])

  const handleSupportClick: () => void = () => {
    window.open(DISCORD_URL, '_blank', 'noopener,noreferrer')
  }

  return (
    <Popover className="relative inline-block pt-1 mb-auto">
      {({ open }) => (
        <>
          <Popover.Button
            as="div"
            onMouseEnter={() => {}}
            className={`
              group rounded-md inline-flex items-center
              hover:text-gray-900 focus:outline-none
              ${open ? 'text-gray-900' : 'text-purple-800'}
            `}
          >
            <DropdownButton open={open} isDelayed={isDelayed} />
          </Popover.Button>
          <TransactionPopoverContainer>
            <OptionButton
              icon={
                <Image
                  className="rounded-full"
                  height={20}
                  src={originChain.explorerImg}
                  alt={`${originChain.explorerName} logo`}
                />
              }
              text={`${originChain.explorerName}`}
              onClick={handleOriginExplorerLink}
            />
            <OptionButton
              icon={
                <Image
                  className="rounded-full"
                  height={20}
                  src={destinationChain.explorerImg}
                  alt={`${destinationChain.explorerName} logo`}
                />
              }
              text={`${destinationChain.explorerName}`}
              onClick={handleDestinationExplorerLink}
            />
            <OptionButton
              icon={
                <Image
                  className="rounded-full"
                  height={20}
                  src={SynapseLogo}
                  alt="Synapse Logo"
                />
              }
              text={`Synapse Explorer`}
              onClick={handleSynapseExplorerLink}
            />
            <OptionButton
              icon={<DiscordIcon height={20} />}
              text={'Contact Support'}
              onClick={handleSupportClick}
            />
          </TransactionPopoverContainer>
        </>
      )}
    </Popover>
  )
}

export function TransactionPopoverContainer({
  children,
  className,
}: {
  children: any
  className?: string
}) {
  return (
    <Transition
      as={Fragment}
      enter="transition ease-out duration-200"
      enterFrom="opacity-0 translate-y-1"
      enterTo="opacity-100 translate-y-0"
      leave="transition ease-in duration-150"
      leaveFrom="opacity-100 translate-y-0"
      leaveTo="opacity-0 translate-y-1"
    >
      <Popover.Panel
        style={{ boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25)' }}
        className={`
          absolute z-10 top-[25px] left-[20px] transform-gpu
          -translate-x-full border border-separator bg-surface
          w-screen max-w-xs rounded-md overflow-hidden shadow-md
        `}
      >
        <div className="shadow-xl">
          <div className="relative grid gap-1 p-1">{children}</div>
        </div>
      </Popover.Panel>
    </Transition>
  )
}

export function DropdownButton({
  open,
  onClick,
  isDelayed,
  className,
  ...props
}: {
  open: boolean
  onClick?: () => void
  isDelayed?: boolean
  className?: string
  props?: any
}) {
  return (
    <Button
      data-test-id="dropdown-button"
      onClick={onClick ? onClick : () => {}}
      className={`
        w-full
        group rounded-sm p-1
        border border-[#2F2F2F] hover:border-[#101018]
        bg-surface hover:bg-[#101018]
        focus:bg-transparent
        ${className}
      `}
      {...props}
    >
      <div className="space-x-2">
        <div className="flex items-center space-x-1 rounded-md">
          <DownArrow className="!w-3 !h-3 text-secondary" />
        </div>
      </div>
    </Button>
  )
}

export const OptionButton = ({
  icon,
  text,
  onClick,
}: {
  icon: any
  text: string
  onClick: () => void
}) => {
  return (
    <div
      data-test-id="option-button"
      onClick={onClick}
      className="flex hover:cursor-pointer hover:bg-[#0A415C] rounded-sm p-1"
    >
      <div className="my-auto mr-1">{icon}</div>
      <div>{text}</div>
    </div>
  )
}
