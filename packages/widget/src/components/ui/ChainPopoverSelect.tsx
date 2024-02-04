import { useState, useEffect, useContext } from 'react'
import _ from 'lodash'
import { Chain } from 'types'

import usePopover from '@/hooks/usePopoverRef'
import { DownArrow } from '@/components/icons/DownArrow'
import { SearchInput } from './SearchInput'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { Web3Context } from '@/providers/Web3Provider'
import { ConnectedIndicator } from '@/components/icons/ConnectedIndicator'

type PopoverSelectProps = {
  options: Chain[]
  remaining: Chain[]
  targets: Chain[]
  onSelect: (selected: Chain) => void
  selected: Chain
  label: string
  isOrigin: boolean
}

export const ChainPopoverSelect = ({
  options,
  remaining,
  targets,
  onSelect,
  selected,
  label,
  isOrigin,
}: PopoverSelectProps) => {
  const { popoverRef, isOpen, togglePopover, closePopover } = usePopover()
  const [activeTab, setActiveTab] = useState<TabOption>(
    isOrigin ? 'All' : 'Target'
  )

  useEffect(() => {
    if (!targets || _.isEmpty(targets)) {
      setActiveTab('All')
    } else if (isOrigin) {
      setActiveTab('All')
    } else if (!isOrigin) {
      setActiveTab('Target')
    }
  }, [targets, isOrigin])

  const handleSelect = (option: Chain) => {
    onSelect(option)
    closePopover()
  }

  const handleTabSelect = (tab: TabOption) => {
    setActiveTab(tab)
  }

  const {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    filteredTargets,
    hasFilteredRemaining,
    hasFilteredResults,
    hasFilteredTargets,
  } = useChainInputFilter(options, remaining, targets, isOpen)

  return (
    <div
      data-test-id="chain-popover-select"
      className="relative w-min col-span-full"
      ref={popoverRef}
    >
      <div
        onClick={() => togglePopover()}
        style={{ background: 'var(--synapse-select-bg)' }}
        className={`
          flex px-2.5 py-1.5 gap-2 items-center rounded
          text-[--synapse-select-text] whitespace-nowrap
          border border-solid border-[--synapse-select-border]
          cursor-pointer hover:border-[--synapse-focus]
        `}
      >
        {selected?.name || 'Network'}
        <DownArrow />
      </div>
      {isOpen && (
        <div
          style={{ background: 'var(--synapse-select-bg)' }}
          className={`
            absolute z-50 mt-1 max-h-60 min-w-48 rounded
            shadow popover text-left list-none overflow-y-auto
            border border-solid border-[--synapse-select-border]
          `}
        >
          <SearchInput
            inputValue={filterValue}
            setInputValue={setFilterValue}
            placeholder="Search Chains"
            isActive={isOpen}
          />
          {targets && targets.length > 0 && (
            <ToggleTabs
              selectedTab={activeTab}
              onTabSelect={handleTabSelect}
              isOrigin={isOrigin}
            />
          )}
          {activeTab === 'All' ? (
            hasFilteredResults ? (
              <ul className="p-0 m-0">
                {filteredOptions.map((option, i) => (
                  <ChainOption
                    key={i}
                    option={option}
                    isSelected={option?.name === selected?.name}
                    onSelect={() => handleSelect(option)}
                  />
                ))}
                {hasFilteredRemaining && (
                  <div
                    className={`
                    sticky top-0 px-2.5 py-2 mt-2 text-sm
                    text-[--synapse-secondary] bg-[--synapse-surface]
                  `}
                  >
                    Other chains
                  </div>
                )}
                {filteredRemaining.map((option, i) => (
                  <ChainOption
                    key={i}
                    option={option}
                    isSelected={option?.name === selected?.name}
                    onSelect={() => handleSelect(option)}
                  />
                ))}
              </ul>
            ) : (
              <div className="p-2 text-sm break-all">
                No chains found
                <br />
                matching '{filterValue}'.
              </div>
            )
          ) : null}
          {activeTab === 'Target' ? (
            hasFilteredTargets ? (
              <ul className="p-0 m-0">
                {filteredTargets.map((option, i) => (
                  <ChainOption
                    key={i}
                    option={option}
                    isSelected={option?.name === selected?.name}
                    onSelect={() => handleSelect(option)}
                  />
                ))}
              </ul>
            ) : (
              <div className="p-2 text-sm break-all">
                No chains found
                <br />
                matching '{filterValue}'.
              </div>
            )
          ) : null}
        </div>
      )}
    </div>
  )
}

const ChainOption = ({
  option,
  isSelected,
  onSelect,
}: {
  option: Chain
  isSelected: boolean
  onSelect: (option: Chain) => void
}) => {
  const web3Context = useContext(Web3Context)

  const { networkId } = web3Context.web3Provider

  return (
    <li
      key={option.id}
      className={`
      pl-2.5 pr-2.5 py-2.5 rounded-[.1875rem] border border-solid
      hover:border-[--synapse-focus] active:opacity-40
      cursor-pointer whitespace-nowrap
      ${
        isSelected
          ? 'border-[--synapse-focus] hover:opacity-70'
          : 'border-transparent'
      }
    `}
      onClick={() => onSelect(option)}
    >
      <div className="flex justify-between">
        <div>{option.name}</div>
        {option.id === networkId && <ConnectedIndicator />}
      </div>
    </li>
  )
}

const useChainInputFilter = (
  options: Chain[],
  remaining: Chain[],
  targets: Chain[],
  isActive: boolean
) => {
  const [filterValue, setFilterValue] = useState('')

  useEffect(() => {
    if (!isActive) {
      setFilterValue('')
    }
  }, [isActive])

  const filterChains = (chains: Chain[], filter: string) => {
    const lowerFilter = filter.toLowerCase()
    return _.filter(chains, (option) => {
      const name = option.name.toLowerCase()
      return name.includes(lowerFilter) || name === lowerFilter
    })
  }

  const filteredOptions = filterChains(options, filterValue)
  const filteredRemaining = filterChains(remaining, filterValue)
  const filteredTargets = filterChains(targets, filterValue)

  const hasFilteredOptions = !_.isEmpty(filteredOptions)
  const hasFilteredRemaining = !_.isEmpty(filteredRemaining)
  const hasFilteredResults = hasFilteredOptions || hasFilteredRemaining
  const hasFilteredTargets = !_.isEmpty(filteredTargets)

  return {
    filterValue,
    setFilterValue,
    filteredOptions,
    filteredRemaining,
    filteredTargets,
    hasFilteredOptions,
    hasFilteredRemaining,
    hasFilteredResults,
    hasFilteredTargets,
  }
}

type TabOption = 'All' | 'Target'

type ToggleTabsProps = {
  selectedTab: TabOption
  onTabSelect: (tab: TabOption) => void
  isOrigin: boolean
}

const ToggleTabs: React.FC<ToggleTabsProps> = ({
  selectedTab,
  onTabSelect,
  isOrigin,
}) => {
  const { protocolName } = useBridgeState()
  const baseTabClass =
    'flex-grow text-sm font-medium text-center text-[--synapse-primary] rounded-sm p-1 '

  const activeTabClass = 'bg-[var(--synapse-surface)]'
  const inactiveTabClass =
    'bg-[var(--synapse-select-bg)] hover:bg-[var-(--synapse-surface)] hover:cursor-pointer'

  return (
    <div className="flex mt-2 mb-2" role="group">
      {isOrigin ? (
        <>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'All' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('All')}
          >
            All
          </div>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'Target' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('Target')}
          >
            {protocolName ?? 'Target'}
          </div>
        </>
      ) : (
        <>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'Target' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('Target')}
          >
            {protocolName ?? 'Target'}
          </div>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'All' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('All')}
          >
            All
          </div>
        </>
      )}
    </div>
  )
}
