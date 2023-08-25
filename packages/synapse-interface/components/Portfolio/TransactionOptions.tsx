import { Fragment, useCallback } from 'react'
import { Popover, Transition } from '@headlessui/react'
import { ChevronDownIcon } from '@heroicons/react/outline'
import Image from 'next/image'
import { Chain } from '@/utils/types'
import { MoreButton } from '../layouts/LandingPageWrapper/MoreButton'
import { TransactionStatus } from './Transaction'
import { PopoverPanelContainer } from '../layouts/LandingPageWrapper'
import DiscordIcon from '../icons/DiscordIcon'
import Button from '../ui/tailwind/Button'
import { getTransactionExplorerLink } from './Activity'
import { getExplorerTxUrl } from '@/constants/urls'

export const TransactionOptions = ({
  originChain,
  destinationChain,
  kappa,
  transactionHash,
  transactionStatus,
}: {
  originChain: Chain
  destinationChain: Chain
  kappa?: string
  transactionHash?: string
  transactionStatus: TransactionStatus
}) => {
  const handleExplorerClick: () => void = useCallback(() => {
    if (
      kappa &&
      originChain &&
      transactionStatus === TransactionStatus.COMPLETED
    ) {
      const explorerLink: string = getTransactionExplorerLink({
        kappa,
        fromChainId: originChain.id,
        toChainId: destinationChain.id,
      })
      window.open(explorerLink, '_blank')
    } else {
      const explorerLink: string = getExplorerTxUrl({
        chainId: originChain.id,
        hash: transactionHash,
      })
      window.open(explorerLink, '_blank')
    }
  }, [kappa, originChain, destinationChain, transactionStatus, transactionHash])

  return (
    <Popover className="relative inline-block">
      {({ open }) => (
        <>
          <Popover.Button
            as="div"
            onMouseEnter={() => {}}
            className={`
            ${open ? 'text-gray-900' : 'text-purple-800'}
            group  rounded-md inline-flex items-center  hover:text-gray-900 focus:outline-none
            `}
          >
            <DropdownButton open={open} />
          </Popover.Button>
          <TransactionPopoverContainer>
            <OptionButton
              icon={
                <Image
                  className="rounded-full"
                  height={20}
                  src={originChain.chainImg}
                  alt={`${originChain.explorerName} logo`}
                />
              }
              text={`Check on ${originChain.explorerName}`}
              onClick={null}
            />
            <OptionButton
              icon={<DiscordIcon height={20} />}
              text={'Contact Support'}
              onClick={null}
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
        className={`
          absolute z-10 top-[-74px] left-[30px] transform-gpu
          -translate-x-full border border-[#3D3D5C] bg-[#252537]
          w-screen max-w-xs rounded-md overflow-hidden
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
  className,
  ...props
}: {
  open: boolean
  onClick?: () => void
  className?: string
  props?: any
}) {
  return (
    <Button
      onClick={onClick ? onClick : () => {}}
      className={`
        w-full
        group rounded-lg p-1
        border border-[#2F2F2F] hover:border-[#101018]
        bg-transparent hover:bg-[#101018]
        focus:bg-transparent
        ${className}
      `}
      {...props}
    >
      <div className="space-x-2">
        <div className="rounded-md">
          <ChevronDownIcon className="w-5 h-5 text-white" />
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
