import _ from 'lodash'

import usePopover from '@/hooks/usePopoverRef'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import { DownArrow } from '@/components/icons/DownArrow'
import { BridgeableToken } from 'types'

type PopoverSelectProps = {
  options: BridgeableToken[]
  remaining: BridgeableToken[]
  balances: TokenBalance[]
  onSelect: (selected: BridgeableToken) => void
  selected: BridgeableToken
}

export function TokenPopoverSelect({
  options,
  remaining,
  balances,
  onSelect,
  selected,
}: PopoverSelectProps) {
  const { popoverRef, isOpen, togglePopover, closePopover } = usePopover()

  const handleSelect = (option: BridgeableToken) => {
    onSelect(option)
    closePopover()
  }

  /** Merge options with balances to sort */
  const optionsWithBalances: TokenBalance[] = mergeTokenOptionsWithBalances(
    options,
    balances
  )

  /** Merge remaining with balances to sort */
  const remainingWithBalances: TokenBalance[] = mergeTokenOptionsWithBalances(
    remaining,
    balances
  )

  const sortedOptionsWithBalances: TokenBalance[] = _.orderBy(
    optionsWithBalances,
    ['parsedBalance', 'token.priorityRank'],
    ['desc', 'asc']
  )

  const sortedRemainingWithBalances: TokenBalance[] = _.orderBy(
    remainingWithBalances,
    ['parsedBalance', 'token.priorityRank'],
    ['desc', 'asc']
  )

  return (
    <div
      data-test-id="token-popover-select"
      className="relative w-min justify-self-end align-self-center"
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
        {selected?.symbol || 'Token'}
        <DownArrow />
      </div>
      {isOpen && (
        <ul className="absolute z-50 mt-1 p-0 bg-[--synapse-surface] border border-solid border-[--synapse-border] rounded shadow popover list-none right-0 overflow-y-auto max-h-80">
          {sortedOptionsWithBalances?.map((option: TokenBalance, index) => (
            <TokenOption
              option={option?.token}
              key={index}
              onSelect={handleSelect}
              selected={selected}
              parsedBalance={option?.parsedBalance}
            />
          ))}
          <div className="pl-2 text-sm underline">Other tokens</div>
          {sortedRemainingWithBalances?.map((option: TokenBalance, index) => (
            <TokenOption
              option={option?.token}
              key={index}
              onSelect={handleSelect}
              selected={selected}
              parsedBalance={option?.parsedBalance}
            />
          ))}
        </ul>
      )}
    </div>
  )
}

function mergeTokenOptionsWithBalances(
  tokens: BridgeableToken[],
  balances: TokenBalance[]
) {
  return tokens?.map((token) => {
    /** If token balance does not exist, set balance to null */
    if (_.isArray(balances) && _.isEmpty(balances)) {
      return {
        token: token,
        balance: null,
        parsedBalance: null,
      }
    } else {
      const matchedTokenBalance: TokenBalance = balances?.find(
        (currentToken: TokenBalance) => currentToken.token === token
      )
      return {
        token: token,
        balance: matchedTokenBalance?.balance,
        parsedBalance: matchedTokenBalance?.parsedBalance,
      }
    }
  })
}

const TokenOption = ({
  option,
  onSelect,
  selected,
  parsedBalance,
}: {
  option: BridgeableToken
  onSelect: (option: BridgeableToken) => void
  selected: BridgeableToken
  parsedBalance: string
}) => {
  return (
    <li
      data-test-id="token-option"
      className={`cursor-pointer rounded border border-solid hover:border-[--synapse-focus] active:opacity-40 flex gap-4 items-center justify-between ${
        option?.symbol === selected?.symbol
          ? 'border-[--synapse-focus] hover:opacity-70'
          : 'border-transparent'
      }`}
      onClick={() => onSelect(option)}
    >
      <abbr title={option?.name} className="p-2 no-underline">
        {option?.symbol}
      </abbr>
      <data
        value={parsedBalance}
        className={`
        text-sm p-2
        ${
          parsedBalance
            ? 'text-[--synapse-secondary]'
            : 'text-[--synapse-focus]'
        }
      `}
      >
        {parsedBalance ?? '−'}
      </data>
    </li>
  )
}
