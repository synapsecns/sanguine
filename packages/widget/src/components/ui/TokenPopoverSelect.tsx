import { useState, useEffect } from 'react'
import _ from 'lodash'
import { BridgeableToken } from 'types'

import usePopover from '@/hooks/usePopoverRef'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import { DownArrow } from '@/components/icons/DownArrow'
import { SearchInput } from './SearchInput'

type PopoverSelectProps = {
  options: BridgeableToken[]
  remaining: BridgeableToken[]
  balances: TokenBalance[]
  onSelect: (selected: BridgeableToken) => void
  selected: BridgeableToken
}

export const TokenPopoverSelect = ({
  options,
  remaining,
  balances,
  onSelect,
  selected,
}: PopoverSelectProps) => {
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

  const {
    filterValue,
    setFilterValue,
    filteredOptions: filteredSortedOptionsWithBalances,
    filteredRemaining: filteredSortedRemainingWithBalances,
    hasFilteredRemaining,
    hasFilteredResults,
  } = useTokenInputFilter(
    sortedOptionsWithBalances,
    sortedRemainingWithBalances,
    isOpen
  )

  return (
    <div
      data-test-id="token-popover-select"
      className="relative w-min justify-self-end align-self-center"
      ref={popoverRef}
    >
      <div
        onClick={() => togglePopover()}
        style={{ background: 'var(--synapse-select-bg)' }}
        className={`
          flex px-2.5 py-1.5 gap-2 items-center rounded-lg
          text-[--synapse-select-text] whitespace-nowrap
          border border-solid border-[--synapse-select-border]
          cursor-pointer hover:border-[--synapse-focus]
        `}
      >
        {selected?.symbol || 'Token'}
        <DownArrow />
      </div>
      {isOpen && (
        <div
          style={{ background: 'var(--synapse-select-bg)' }}
          className={`
            absolute right-0 z-50 mt-1 p-1 max-h-80 min-w-48 rounded-lg
            shadow popover text-left list-none overflow-y-auto
            border border-solid border-[--synapse-select-border]
          `}
        >
          <SearchInput
            inputValue={filterValue}
            setInputValue={setFilterValue}
            placeholder="Search Tokens"
          />
          {hasFilteredResults ? (
            <ul className="p-0 mt-px mb-0 space-y-px">
              {filteredSortedOptionsWithBalances?.map(
                (option: TokenBalance, index) => (
                  <TokenOption
                    option={option?.token}
                    key={index}
                    onSelect={handleSelect}
                    selected={selected}
                    parsedBalance={option?.parsedBalance}
                  />
                )
              )}
              {hasFilteredRemaining && (
                <div
                  style={{ background: 'var(--synapse-select-bg)' }}
                  className={`
                    sticky top-0 px-2.5 py-2 mt-2
                    text-sm text-[--synapse-secondary]
                  `}
                >
                  Other tokens
                </div>
              )}
              {filteredSortedRemainingWithBalances?.map(
                (option: TokenBalance, index) => (
                  <TokenOption
                    option={option?.token}
                    key={index}
                    onSelect={handleSelect}
                    selected={selected}
                    parsedBalance={option?.parsedBalance}
                  />
                )
              )}
            </ul>
          ) : (
            <div className="p-2 break-all">
              No tokens found
              <br />
              matching '{filterValue}'.
            </div>
          )}
        </div>
      )}
    </div>
  )
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
      className={`
        flex gap-4 items-center justify-between
        cursor-pointer rounded-lg border border-solid
        hover:border-[--synapse-focus] active:opacity-40
        ${
          option?.symbol === selected?.symbol
            ? 'border-[--synapse-focus] hover:opacity-70'
            : 'border-transparent'
        }
      `}
      onClick={() => onSelect(option)}
    >
      <abbr title={option?.name} className="p-2.5 no-underline">
        {option?.symbol}
      </abbr>
      <data
        value={parsedBalance}
        className={`
          text-sm p-2.5
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

const useTokenInputFilter = (
  options: TokenBalance[],
  remaining: TokenBalance[],
  isActive: boolean
) => {
  const [filterValue, setFilterValue] = useState('')

  useEffect(() => {
    if (!isActive) {
      setFilterValue('')
    }
  }, [isActive])

  const filterTokens = (tokens: TokenBalance[], filter: string) => {
    const lowerFilter = filter.toLowerCase()
    return _.filter(tokens, (option) => {
      const symbol = option.token.symbol.toLowerCase()
      return symbol.includes(lowerFilter) || symbol === lowerFilter
    })
  }

  const filteredOptions = filterTokens(options, filterValue)
  const filteredRemaining = filterTokens(remaining, filterValue)

  const hasFilteredOptions = !_.isEmpty(filteredOptions)
  const hasFilteredRemaining = !_.isEmpty(filteredRemaining)
  const hasFilteredResults = hasFilteredOptions || hasFilteredRemaining

  return {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    hasFilteredOptions,
    hasFilteredRemaining,
    hasFilteredResults,
  }
}

const mergeTokenOptionsWithBalances = (
  tokens: BridgeableToken[],
  balances: TokenBalance[]
) => {
  return tokens?.map((token) => {
    /** If token balance does not exist, set balance to null */
    if (_.isArray(balances) && _.isEmpty(balances)) {
      return {
        token,
        balance: null,
        parsedBalance: null,
      }
    } else {
      const matchedTokenBalance: TokenBalance = balances?.find(
        (currentToken: TokenBalance) => currentToken.token === token
      )
      return {
        token,
        balance: matchedTokenBalance?.balance,
        parsedBalance: matchedTokenBalance?.parsedBalance,
      }
    }
  })
}
