import { useDispatch, useSelector } from 'react-redux'
import Select from 'react-select'
import { Address, useAccount } from 'wagmi'

import { CHAINS_BY_ID } from '@/constants/chains'
import { setToChainId } from '@/slices/bridge/reducer'
import { RootState } from '@/store/store'
import { networkSelectStyles } from './styles/networkSelectStyles'
import { shortenAddress } from '@/utils/shortenAddress'
import { useEffect, useId, useState } from 'react'

const ToChainSelect = () => {
  const { toChainId, toChainIds } = useSelector(
    (state: RootState) => state.bridge
  )

  const { address: isConnectedAddress } = useAccount()
  const [address, setAddress] = useState<Address>()

  useEffect(() => {
    setAddress(isConnectedAddress)
  }, [isConnectedAddress])

  const dispatch = useDispatch()

  const toChainOptions = toChainIds.map((option) => ({
    label: (
      <span className="flex items-center space-x-2">
        <img src={CHAINS_BY_ID[option].chainImg.src} className="w-5 h-5" />
        <div className="text-primaryTextColor">
          {CHAINS_BY_ID[option].name} [{option}]
        </div>
      </span>
    ),
    value: option,
  }))

  const handleToChainChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setToChainId(Number(selectedOption.value)))
    } else {
      dispatch(setToChainId(null))
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

  // update DisplayAddress for destinationAddress

  return (
    <div className="flex items-center justify-between">
      <div className="pl-1">
        <div className="pl-2 -mb-1 text-xs text-secondaryTextColor">To</div>
        <Select
          instanceId={useId()}
          styles={networkSelectStyles}
          key={toChainId}
          options={toChainOptions}
          filterOption={customFilterOption}
          onChange={handleToChainChange}
          isClearable={true}
          isSearchable={true}
          placeholder={<div className="text-secondaryTextColor">Network</div>}
          value={toChainOptions.find((option) => option.value === toChainId)}
        />
      </div>
      {address && <DisplayAddress address={address} />}
    </div>
  )
}

const DisplayAddress = ({ address }) => {
  return (
    <div className="border-[0.5px] border-secondaryTextColor rounded-xl pt-1 pb-1 pl-3 pr-3 text-secondaryTextColor text-xxs">
      {shortenAddress(address, 4)}
    </div>
  )
}

export default ToChainSelect
