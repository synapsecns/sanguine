import usePopover from '@/hooks/usePopoverRef'
import { DownArrow } from '@/components/icons/DownArrow'
import { Chain } from 'types'

type PopoverSelectProps = {
  options: Chain[]
  remaining: Chain[]
  onSelect: (selected: Chain) => void
  selected: Chain
  label: string
}

export function ChainPopoverSelect({
  options,
  remaining,
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
          cursor-pointer flex px-2.5 py-1.5 gap-2 items-center rounded-[.1875rem]
          text-[--synapse-select-text]
          bg-[--synapse-select-bg]
          border border-solid border-[--synapse-select-border]
          hover:border-[--synapse-focus]
          whitespace-nowrap
        `}
        onClick={() => togglePopover()}
      >
        {selected?.name || 'Network'}
        <DownArrow />
      </div>
      {isOpen && (
        <ul className="absolute z-50 mt-1 p-0 bg-[--synapse-surface] border border-solid border-[--synapse-border] rounded shadow popover text-left list-none overflow-y-auto max-h-60">
          {options.map((option) => (
            <li
              key={option.id}
              className={`cursor-pointer pl-2.5 pr-8 py-2.5 rounded border border-solid hover:border-[--synapse-focus] active:opacity-40 whitespace-nowrap ${
                option?.name === selected?.name
                  ? 'border-[--synapse-focus] hover:opacity-70'
                  : 'border-transparent'
              }`}
              onClick={() => handleSelect(option)}
            >
              {option?.name}
            </li>
          ))}
          <div className="px-2.5 py-2 mt-2 text-sm text-[--synapse-secondary] cursor-default sticky top-0 bg-[--synapse-surface]">Other chains</div>
          {remaining.map((option) => (
            <li
              key={option.id}
              className={`cursor-pointer pl-2.5 pr-8 py-2.5 rounded border border-solid hover:border-[--synapse-focus] active:opacity-40 whitespace-nowrap ${
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
