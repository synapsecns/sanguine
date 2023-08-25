import { Popover } from '@headlessui/react'
import Image from 'next/image'
import { MoreButton } from '../layouts/LandingPageWrapper/MoreButton'
import { PopoverPanelContainer } from '../layouts/LandingPageWrapper'

export const TransactionOptions = () => {
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
          <PopoverPanelContainer className="-translate-x-full">
            <button> Check on Block Scanner</button>
            <button> Contact Support </button>
          </PopoverPanelContainer>
        </>
      )}
    </Popover>
  )
}

export const OptionButton = ({
  logoSrc,
  text,
  onClick,
}: {
  logoSrc: any
  text: string
  onClick: () => void
}) => {
  return (
    <div data-test-id="option-button" className="flex" onClick={onClick}>
      <Image width={16} height={16} src={logoSrc} alt="option-button-logo" />
      <div>{text}</div>
    </div>
  )
}
