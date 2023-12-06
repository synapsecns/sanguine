import { TokenMetaData, Chain } from 'types'
import usePopover from '@/hooks/usePopoverRef'
import { DownArrow } from '../icons/DownArrow'

type PopoverSelectProps = {
  selectedChain: Chain
  options: TokenMetaData[]
  onSelect: (selected: TokenMetaData) => void
  selected: TokenMetaData
  label: string
}

export function TokenPopoverSelect({
  selectedChain,
  options,
  onSelect,
  selected,
  label,
}: PopoverSelectProps) {
  const { popoverRef, isOpen, togglePopover, closePopover } = usePopover()

  const handleSelect = (option: TokenMetaData) => {
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
          {selected.symbol}
        </div>
      </div>
      {isOpen && (
        <div className="absolute z-50 mt-1 bg-[--synapse-bg-select] rounded shadow popover">
          {options.map((option, index) => {
            return (
              <div
                data-test-id="token-option"
                key={index}
                className={`w-full cursor-pointer px-2 py-2.5 ${
                  option.symbol === selected.symbol
                    ? 'border border-zinc-300 rounded-md'
                    : 'border border-transparent active:border-zinc-300 rounded-md hover:bg-[--synapse-bg-surface] active:opacity-40'
                }`}
                onClick={() => handleSelect(option)}
              >
                <div className="flex items-center gap-2">
                  <div className="mr-3">{option.symbol}</div>
                  <div className="flex gap-1 ml-auto text-xs">
                    <div className="text-[--synapse-text-secondary]">
                      {option.symbol}
                    </div>
                  </div>
                </div>
              </div>
            )
          })}
        </div>
      )}
    </div>
  )
}
