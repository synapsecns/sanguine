import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Select from 'react-select'
import { useAccount } from 'wagmi'

import { CHAINS_BY_ID } from '@/constants/chains'

import { setFromChainId, setFromToken } from '@/slices/bridge/reducer'
import { networkSelectStyles } from './styles/networkSelectStyles'
import { useEffect, useId, useState } from 'react'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { filterPortfolioBalancesWithBalances } from '../Portfolio/Portfolio'

const FromChainSelect = () => {
  const { fromChainId, fromChainIds } = useBridgeState()

  const dispatch = useDispatch()
  const { isConnected: isConnectedOriginal } = useAccount()
  const [isConnected, setIsConnected] = useState(false)

  const balancesAndAllowances = usePortfolioBalances()

  const chainIdsWithBalances = Object.keys(
    filterPortfolioBalancesWithBalances(balancesAndAllowances)
  ).map(Number)

  const sortedFromChainIds = _.orderBy(fromChainIds, [
    (id) => !chainIdsWithBalances.includes(id),
  ])

  useEffect(() => {
    setIsConnected(isConnectedOriginal)
  }, [isConnectedOriginal])

  const fromChainOptions = [
    {
      label: 'Wallet',
      options: sortedFromChainIds
        .filter((option) => chainIdsWithBalances.includes(option))
        .map((option) => ({
          label: (
            <span className="flex items-center space-x-2">
              <img
                src={CHAINS_BY_ID[option].chainImg.src}
                className="w-5 h-5"
              />
              <div className="text-primaryTextColor">
                {CHAINS_BY_ID[option].name}
              </div>
            </span>
          ),
          value: option,
        })),
    },
    {
      label: 'All other chains',
      options: sortedFromChainIds
        .filter((option) => !chainIdsWithBalances.includes(option))
        .map((option) => ({
          label: (
            <span className="flex items-center space-x-2">
              <img
                src={CHAINS_BY_ID[option].chainImg.src}
                className="w-5 h-5"
              />
              <div className="text-primaryTextColor">
                {CHAINS_BY_ID[option].name}
              </div>
            </span>
          ),
          value: option,
        })),
    },
  ]

  const handleFromChainChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setFromChainId(Number(selectedOption.value)))
    } else {
      dispatch(setFromChainId(null))
      dispatch(setFromToken(null))
    }
  }

  const customFilterOption = (option, rawInput) => {
    const chainId = option.value
    const name = CHAINS_BY_ID[chainId].name
    const searchTerm = rawInput.toLowerCase()

    return (
      name.toLowerCase().includes(searchTerm) ||
      chainId.toString().includes(searchTerm)
    )
  }

  return (
    <div className="flex items-center justify-between">
      <div className="pl-1">
        <div className="pl-2 -mb-1 text-xs text-secondaryTextColor">From</div>
        <Select
          instanceId={useId()}
          classNamePrefix="react-select"
          styles={networkSelectStyles}
          key={fromChainId}
          options={fromChainOptions}
          onChange={handleFromChainChange}
          isClearable={true}
          isSearchable={true}
          filterOption={customFilterOption}
          placeholder={<div className="text-secondaryTextColor">Network</div>}
          value={fromChainOptions
            .flatMap((group) => group.options)
            .find((option) => Number(option.value) === fromChainId)}
        />
      </div>
      <div>
        {isConnected ? (
          <ConnectedButton />
        ) : (
          <button style={{ display: 'none' }}></button>
        )}
      </div>
    </div>
  )
}

export default FromChainSelect

const ConnectedButton = () => {
  return (
    <button
      data-test-id="connected-button"
      className={`
      flex items-center justify-center
      text-base text-white px-3 py-1 rounded-3xl
      text-center transform-gpu transition-all duration-75
      border border-solid border-transparent
      `}
    >
      <div className="flex flex-row text-sm">
        <div
          className={`
            my-auto ml-auto mr-2 w-2 h-2
            bg-green-500 rounded-full
            `}
        />
        Connected
      </div>
    </button>
  )
}
