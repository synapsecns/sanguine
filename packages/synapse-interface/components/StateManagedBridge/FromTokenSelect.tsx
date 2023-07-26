import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Select from 'react-select'

import { Token } from '@/utils/types'
import { setFromToken } from '@/slices/bridge/reducer'
import { coinSelectStyles } from './styles/coinSelectStyles'
import { useId } from 'react'

import { useAppSelector } from '@/store/hooks'

const ImageAndCoin = ({ option }: { option: Token }) => {
  const { fromChainId } = useAppSelector((state) => state.bridge)
  const { balancesAndAllowances } = useAppSelector((state) => state.portfolio)

  const { icon, symbol } = option
  const parsedBalance = balancesAndAllowances[fromChainId]?.find(
    (token) => token.tokenAddress === option.addresses[fromChainId]
  )?.parsedBalance

  return (
    <div className="flex items-center justify-between">
      <div className="flex items-center space-x-2" key={option.symbol}>
        <img src={icon.src} className="w-6 h-6" />
        <div className="text-xl">{symbol}</div>
      </div>
      <div className="select-hidden">
        {parsedBalance !== '0.0' ? parsedBalance : ''}
      </div>
    </div>
  )
}

const FromTokenSelect = () => {
  const { fromChainId, fromToken, fromTokens } = useAppSelector(
    (state) => state.bridge
  )

  const dispatch = useDispatch()

  const fromTokenOptions = fromTokens.map((option, i) => {
    return {
      label: <ImageAndCoin option={option} />,
      value: option,
    }
  })

  const handleFromTokenChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setFromToken(selectedOption.value))
    } else {
      dispatch(setFromToken(null))
    }
  }

  const customFilter = (option, searchInput) => {
    if (searchInput) {
      const searchTerm = searchInput.toLowerCase()
      return (
        option.value.symbol.toLowerCase().includes(searchTerm) ||
        option.value.name.toLowerCase().includes(searchTerm) ||
        option.value.addresses[fromChainId].toLowerCase().includes(searchTerm)
      )
    }
    return true
  }

  return (
    <Select
      styles={coinSelectStyles}
      classNamePrefix="mySelect"
      instanceId={useId()}
      key={fromToken?.symbol}
      options={fromTokenOptions}
      filterOption={customFilter}
      onChange={handleFromTokenChange}
      isSearchable={true}
      placeholder={<span className="text-xl text-white">In</span>}
      value={fromTokenOptions.find((option) => option.value === fromToken)}
    />
  )
}

export default FromTokenSelect
