import { useDispatch, useSelector } from 'react-redux'
import Select from 'react-select'
import { useAccount } from 'wagmi'

import { CHAINS_BY_ID } from '@/constants/chains'
import { RootState } from '@/store/store'

import { setFromChainId } from '@/slices/bridge/reducer'
import { networkSelectStyles } from './styles/networkSelectStyles'
import { useEffect, useId, useState } from 'react'

const FromChainSelect = () => {
  const { fromChainId, fromChainIds } = useSelector(
    (state: RootState) => state.bridge
  )

  const dispatch = useDispatch()
  const { isConnected: isConnectedOriginal } = useAccount()
  const [isConnected, setIsConnected] = useState(false)

  useEffect(() => {
    setIsConnected(isConnectedOriginal)
  }, [isConnectedOriginal])

  const fromChainOptions = fromChainIds.map((option) => ({
    label: (
      <span className="flex items-center space-x-1">
        <img src={CHAINS_BY_ID[option].chainImg.src} className="w-5 h-5" />
        <div>
          {CHAINS_BY_ID[option].name} [{option}]
        </div>
      </span>
    ),
    value: option,
  }))

  const handleFromChainChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setFromChainId(Number(selectedOption.value)))
    } else {
      dispatch(setFromChainId(null))
    }
  }

  const customFilterOption = (option, rawInput) => {
    const searchTerm = rawInput.toLowerCase()

    return (
      option.data.label.props.children[1].props.children[0]
        .toLowerCase()
        .includes(searchTerm) ||
      option.value.toString().toLowerCase().includes(searchTerm)
    )
  }

  return (
    <div className="flex items-center justify-between">
      <div className="pl-1">
        <div className="pl-2 -mb-1 text-xs text-secondaryTextColor">From</div>
        <Select
          instanceId={useId()}
          styles={networkSelectStyles}
          key={fromChainId}
          options={fromChainOptions}
          onChange={handleFromChainChange}
          isClearable={true}
          isSearchable={true}
          filterOption={customFilterOption}
          placeholder={<div className="text-secondaryTextColor">Network</div>}
          value={fromChainOptions.find(
            (option) => Number(option.value) === fromChainId
          )}
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
