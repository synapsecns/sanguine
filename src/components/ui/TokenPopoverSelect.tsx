import { BridgeableToken, Chain } from 'types'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import usePopover from '@/hooks/usePopoverRef'
import { DownArrow } from '../icons/DownArrow'

type PopoverSelectProps = {
  selectedChainId: number
  options: BridgeableToken[]
  remaining: BridgeableToken[]
  balances: TokenBalance[]
  onSelect: (selected: BridgeableToken) => void
  selected: BridgeableToken
  label: string
}

export function TokenPopoverSelect({
  selectedChainId,
  options,
  remaining,
  balances,
  onSelect,
  selected,
  label,
}: PopoverSelectProps) {
  const { popoverRef, isOpen, togglePopover, closePopover } = usePopover()

  const handleSelect = (option: BridgeableToken) => {
    onSelect(option)
    closePopover()
  }

  return (
    <div
      data-test-id="token-popover-select"
      className="relative w-min"
      ref={popoverRef}
    >
      <div
        className="cursor-pointer items-center grid rounded-full bg-[--synapse-bg-select] border border-solid border-[--synapse-border] hover:border-[--synapse-border-hover]"
        onClick={() => togglePopover()}
      >
        <span className="col-start-1 row-start-1 pr-3 text-xs h-min justify-self-end">
          <DownArrow />
        </span>
        <div className="col-start-1 row-start-1 py-1 pl-3 bg-transparent outline-none appearance-none cursor-pointer pr-7">
          {selected?.symbol || 'Token'}
        </div>
      </div>
      {isOpen && (
        <div className="absolute z-50 mt-1 bg-[--synapse-bg-surface] border border-solid border-[--synapse-border] rounded shadow popover right-0">
          {options.map((option: BridgeableToken, index) => {
            const matchedTokenBalance: TokenBalance = balances?.find(
              (token: TokenBalance) => token.token === option
            )
            const parsedBalance: string = matchedTokenBalance?.parsedBalance
            return (
              <TokenOption
                option={option}
                index={index}
                onSelect={handleSelect}
                selected={selected}
                parsedBalance={parsedBalance}
              />
            )
          })}
          {remaining?.map((option: BridgeableToken, index) => {
            const matchedTokenBalance: TokenBalance = balances?.find(
              (token: TokenBalance) => token.token === option
            )
            const parsedBalance: string = matchedTokenBalance?.parsedBalance
            return (
              <TokenOption
                option={option}
                index={index}
                onSelect={handleSelect}
                selected={selected}
                parsedBalance={parsedBalance}
              />
            )
          })}
        </div>
      )}
    </div>
  )
}

const TokenOption = ({
  option,
  index,
  onSelect,
  selected,
  parsedBalance,
}: {
  option: BridgeableToken
  index: number
  onSelect: (option: BridgeableToken) => void
  selected: BridgeableToken
  parsedBalance: string
}) => {
  return (
    <div
      data-test-id="token-option"
      key={index}
      className={`cursor-pointer px-2 py-2.5 ${
        option.symbol === selected?.symbol
          ? 'border border-solid border-[--synapse-border-hover] rounded-md hover:border-[--synapse-border-hover] hover:opacity-70 active:opacity-40'
          : 'border border-solid border-transparent rounded hover:bg-[--synapse-bg-select] hover:border-[--synapse-border-hover] active:opacity-40'
      }`}
      onClick={() => onSelect(option)}
    >
      <div className="flex items-center gap-2">
        <div className="mr-3">{option.symbol}</div>
        <div className="flex gap-1 ml-auto text-xs">
          {parsedBalance && <div>{parsedBalance}</div>}
          <div className="text-[--synapse-text-secondary]">{option.symbol}</div>
        </div>
      </div>
    </div>
  )
}
