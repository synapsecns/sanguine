import { Popover } from '@headlessui/react'
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
          <PopoverPanelContainer className="-translate-x-full border border-[#3D3D5C] bg-[#252537] ">
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
          </PopoverPanelContainer>
        </>
      )}
    </Popover>
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
      className="flex hover:cursor-pointer hover:bg-[#0A415C]"
    >
      <div className="my-auto">{icon}</div>
      <div>{text}</div>
    </div>
  )
}
