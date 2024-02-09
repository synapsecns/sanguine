import { useState, useEffect } from 'react'
import _ from 'lodash'
import { BridgeableToken } from 'types'

import usePopover from '@/hooks/usePopoverRef'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import { DownArrow } from '@/components/icons/DownArrow'
import { SearchInput } from '@/components/ui/SearchInput'
import { TokenOption } from '@/components/ui/TokenOption'
import { useBridgeState } from '@/state/slices/bridge/hooks'

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

  const nudge: boolean = (() => {
    const { originChainId, destinationChainId, originToken, destinationToken } =
      useBridgeState()

    return !!(
      !selected &&
      originChainId &&
      destinationChainId &&
      (balances.length || (!balances.length && originToken))
    )
  })()

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
          flex px-2.5 py-1.5 gap-2 items-center rounded
          text-[--synapse-select-text] whitespace-nowrap
          border border-solid border-[--synapse-select-border]
          cursor-pointer border border-solid
          ${
            nudge
              ? 'border-[--synapse-progress] hover:shadow-lg'
              : 'border-[--synapse-select-border] hover:border-[--synapse-focus]'
          }
        `}
      >
        {selected?.imgUrl && (
          <img
            src={selected?.imgUrl}
            alt={`${selected.symbol} token icon`}
            className="inline w-4 h-4 -ml-1"
          />
        )}
        {selected?.symbol || 'Token'}
        <DownArrow />
      </div>
      {isOpen && (
        <div
          style={{ background: 'var(--synapse-select-bg)' }}
          className={`
            absolute right-0 z-50 mt-1 max-h-80 min-w-48 rounded
            shadow popover text-left list-none overflow-y-auto
            border border-solid border-[--synapse-select-border]
            animate-slide-down origin-top
          `}
        >
          <div className="p-1">
            <SearchInput
              inputValue={filterValue}
              setInputValue={setFilterValue}
              placeholder="Search Tokens"
              isActive={isOpen}
            />
          </div>
          {hasFilteredResults ? (
            <ul className="p-0 m-0">
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
                  className={`
                    sticky top-0 px-2.5 py-2 mt-2 text-sm
                    text-[--synapse-secondary] bg-[--synapse-surface]
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
            <div className="p-2 text-sm break-all">
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
