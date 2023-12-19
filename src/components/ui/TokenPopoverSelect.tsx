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
        className={`
          cursor-pointer flex px-3 py-1 gap-1.5 items-center rounded-full
          text-[--synapse-text-primary]
          bg-[--synapse-bg-select]
          border border-solid border-[--synapse-border]
          hover:border-[--synapse-border-hover]
        `}
        onClick={() => togglePopover()}
      >
        {selected?.symbol || 'Token'}
        <DownArrow />
      </div>
      {isOpen && (
        <ul className="absolute z-50 mt-1 p-0 bg-[--synapse-bg-surface] border border-solid border-[--synapse-border] rounded shadow popover list-none -right-0">
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
        </ul>
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
  console.log(option)
  return (
    <li
      data-test-id="token-option"
      key={index}
      className={`cursor-pointer rounded border border-solid hover:border-[--synapse-border-hover] active:opacity-40 flex gap-4 items-center justify-between ${
        option.symbol === selected?.symbol
          ? 'border-[--synapse-border-hover] hover:opacity-70'
          : 'border-transparent'
      }`}
      onClick={() => onSelect(option)}
    >
      <abbr title={option.name} className="no-underline p-2">
        {option.symbol}
      </abbr>
      <data value={parsedBalance} className={`
        text-sm p-2
        ${parsedBalance
          ? 'text-[--synapse-text-secondary]' 
          : 'text-[--synapse-border-hover]'
        }
      `}>
        {parsedBalance ?? '−'}
      </data>
    </li>
  )
}
