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
    <div className="relative w-min" ref={popoverRef}>
      <div
        className="cursor-pointer items-center grid rounded-full bg-[--synapse-bg-select] border border-[--synapse-border] hover:border-[--synapse-border-hover]"
        onClick={() => togglePopover()}
      >
        <span className="col-start-1 row-start-1 pr-3 text-xs h-min justify-self-end">
          <DownArrow />
        </span>
        <div className="col-start-1 row-start-1 py-1 pl-3 bg-transparent outline-none appearance-none cursor-pointer pr-7">
          {selected?.name || 'Network'}
        </div>
      </div>
      {isOpen && (
        <div className="absolute z-50 mt-1 bg-[--synapse-bg-surface] border border-[--synapse-border] rounded shadow popover">
          {options.map((option, index) => (
            <div
              key={index}
              className={`cursor-pointer pl-2 pr-4 py-2.5 ${
                option?.name === selected?.name
                  ? 'border border-[--synapse-border-hover] rounded hover:border-[--synapse-border-hover] hover:opacity-70 active:opacity-40'
                  : 'border border-transparent rounded hover:bg-[--synapse-bg-select] hover:border-[--synapse-border-hover] active:opacity-40'
              }`}
              onClick={() => handleSelect(option)}
            >
              <div className="flex gap-2">
                <div>{option?.name}</div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
