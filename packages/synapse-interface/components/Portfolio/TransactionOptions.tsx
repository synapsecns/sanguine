import { Fragment } from 'react'
import { Popover, Transition } from '@headlessui/react'
import Image from 'next/image'
import { MoreButton } from '../layouts/LandingPageWrapper/MoreButton'
import { PopoverPanelContainer } from '../layouts/LandingPageWrapper'
import DiscordIcon from '../icons/DiscordIcon'
import { Chain } from '@/utils/types'

export const TransactionOptions = ({ originChain }: { originChain: Chain }) => {
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
            <MoreButton open={open} />
          </Popover.Button>
          <TransactionPopoverContainer className="-translate-x-full border border-[#3D3D5C] bg-[#252537] top- ">
            <OptionButton
              icon={
                <Image
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
          absolute z-10 top-[-74px] left-[42px] transform-gpu
          ${className ?? '-translate-x-1/2'}
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
      <div className="my-auto">{icon}</div>
      <div>{text}</div>
    </div>
  )
}
