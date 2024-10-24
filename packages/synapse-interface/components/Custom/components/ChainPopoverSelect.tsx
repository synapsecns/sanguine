import _ from 'lodash'
import { type Chain } from '@/utils/types'

import usePopover from '../hooks/usePopoverRef'
import { DownArrow } from '@/components/icons/DownArrow'
import { ChainOption } from './ChainOption'

type PopoverSelectProps = {
  options: Chain[]
  onSelect: (selected: Chain) => void
  selected: Chain
  label: string
  isOrigin: boolean
}

export const ChainPopoverSelect = ({
  options,
  onSelect,
  selected,
  label,
  isOrigin,
}: PopoverSelectProps) => {
  const { popoverRef, isOpen, togglePopover, closePopover } = usePopover()

  const handleSelect = (option: Chain) => {
    onSelect(option)
    closePopover()
  }

  return (
    <div
      data-test-id="chain-popover-select"
      className="relative w-min col-span-full"
      ref={popoverRef}
    >
      <div
        id={`${isOrigin ? 'origin' : 'destination'}-chain-select`}
        onClick={() => togglePopover()}
        className={`
          flex px-2.5 py-1.5 gap-2 items-center rounded
           whitespace-nowrap
          border border-solid border-zinc-200 dark:border-zinc-700
          cursor-pointer 
          w-[150px]
        `}
      >
        {selected?.chainImg && (
          <img
            src={selected?.chainImg.src}
            alt={`${selected?.name} chain icon`}
            className="inline w-4 h-4"
          />
        )}
        {selected?.name || 'Network'}
        <DownArrow />
      </div>
      {isOpen && (
        <div
          className={`
            bg-zinc-100 dark:bg-bgBase
            absolute z-50 mt-1 max-h-60 min-w-48 rounded
            shadow popover text-left list-none overflow-y-auto
            border border-solid border-zinc-200 dark:border-zinc-700
            animate-slide-down origin-top
          `}
        >
          <ul className="p-0 m-0">
            {options.map((option, i) => (
              <ChainOption
                key={i}
                option={option}
                isSelected={option?.name === selected?.name}
                onSelect={() => handleSelect(option)}
                isOrigin={isOrigin}
              />
            ))}
          </ul>
        </div>
      )}
    </div>
  )
}
