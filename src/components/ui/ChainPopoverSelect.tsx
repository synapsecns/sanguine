import usePopover from '@/hooks/usePopoverRef'
import { DownArrow } from '@/components/icons/DownArrow'
import { Chain } from 'types'

type PopoverSelectProps = {
  options: Chain[]
  onSelect: (selected: Chain) => void
  selected: Chain
  label: string
}

export function ChainPopoverSelect({
  options,
  onSelect,
  selected,
  label,
}: PopoverSelectProps) {
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
        className={`
          cursor-pointer flex px-3 py-1 gap-1.5 items-center rounded-full
          text-[--synapse-select-text]
          bg-[--synapse-select-bg]
          border border-solid border-[--synapse-select-border]
          hover:border-[--synapse-focus]
        `}
        onClick={() => togglePopover()}
      >
        {selected?.name || 'Network'}
        <DownArrow />
      </div>
      {isOpen && (
        <ul className="absolute z-50 mt-1 p-0 bg-[--synapse-surface] border border-solid border-[--synapse-border] rounded shadow popover text-left list-none">
          {options.map((option, index) => (
            <li
              key={index}
              className={`cursor-pointer pl-2 pr-3 py-2 rounded border border-solid hover:border-[--synapse-focus] active:opacity-40 ${
                option?.name === selected?.name
                  ? 'border-[--synapse-focus] hover:opacity-70'
                  : 'border-transparent'
              }`}
              onClick={() => handleSelect(option)}
            >
              {option?.name}
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}
