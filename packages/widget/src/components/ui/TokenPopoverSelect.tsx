import { useState, useEffect } from 'react'
import _ from 'lodash'
import { BridgeableToken } from 'types'

import usePopover from '@/hooks/usePopoverRef'
import { TokenBalance } from '@/utils/actions/fetchTokenBalances'
import { DownArrow } from '@/components/icons/DownArrow'
import { InputFilter } from './InputFilter'

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
    noFilteredResults,
  } = useTokenInputFilter(
    sortedOptionsWithBalances,
    sortedRemainingWithBalances,
    isOpen
  )

  /** Filters tokens based on User Input */
  // const [filterValue, setFilterValue] = useState('')

  // useEffect(() => {
  //   if (!isOpen) {
  //     setFilterValue('')
  //   }
  // }, [isOpen])

  // const filteredSortedOptionsWithBalances = _.filter(
  //   sortedOptionsWithBalances,
  //   (option) => {
  //     const symbol = option.token.symbol
  //     const lowerSymbol = symbol.toLowerCase()
  //     const lowerFilter = filterValue.toLowerCase()
  //     return lowerSymbol.includes(lowerFilter) || lowerSymbol === lowerFilter
  //   }
  // )
  // const filteredSortedRemainingWithBalances = _.filter(
  //   sortedRemainingWithBalances,
  //   (option) => {
  //     const symbol = option.token.symbol
  //     const lowerSymbol = symbol.toLowerCase()
  //     const lowerFilter = filterValue.toLowerCase()
  //     return lowerSymbol.includes(lowerFilter) || lowerSymbol === lowerFilter
  //   }
  // )
  // const noFilteredOptions = _.isEmpty(filteredSortedOptionsWithBalances)
  // const noFilteredRemaining = _.isEmpty(filteredSortedRemainingWithBalances)
  // const noFilteredResults = noFilteredOptions && noFilteredRemaining

  return (
    <div
      data-test-id="token-popover-select"
      className="relative w-min justify-self-end align-self-center"
      ref={popoverRef}
    >
      <div
        className={`
          cursor-pointer flex px-2.5 py-1.5 gap-2 items-center rounded-[.1875rem]
          text-[--synapse-select-text]
          border border-solid border-[--synapse-select-border]
          hover:border-[--synapse-focus]
        `}
        style={{ background: 'var(--synapse-select-bg)' }}
        onClick={() => togglePopover()}
      >
        {selected?.symbol || 'Token'}
        <DownArrow />
      </div>
      {isOpen && (
        <div
          className="absolute z-50 mt-1 p-0 border border-solid border-[--synapse-select-border] rounded shadow popover list-none right-0 overflow-y-auto max-h-80 min-w-48"
          style={{ background: 'var(--synapse-select-bg)' }}
        >
          <InputFilter
            inputValue={filterValue}
            setInputValue={setFilterValue}
            placeholder="Search Tokens"
          />
          {noFilteredResults ? (
            <div className="p-2 break-all">
              No tokens found matching '{filterValue}'.
            </div>
          ) : (
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

              {remaining?.length > 0 && (
                <div
                  className="px-2.5 py-2 mt-2 text-sm text-[--synapse-secondary] cursor-default sticky top-0"
                  style={{ background: 'var(--synapse-select-bg)' }}
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
          )}
        </div>
      )}
    </div>
  )
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

  const noFilteredOptions = _.isEmpty(filteredOptions)
  const noFilteredRemaining = _.isEmpty(filteredRemaining)
  const noFilteredResults = noFilteredOptions && noFilteredRemaining

  return {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    noFilteredResults,
  }
}
