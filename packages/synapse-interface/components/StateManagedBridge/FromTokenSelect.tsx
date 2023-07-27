import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Select from 'react-select'

import { Token } from '@/utils/types'
import { setFromToken } from '@/slices/bridge/reducer'
import { coinSelectStyles } from './styles/coinSelectStyles'
import { useId } from 'react'

import { useAppSelector } from '@/store/hooks'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'

const ImageAndCoin = ({ option }: { option: Token }) => {
  const { fromChainId } = useAppSelector((state) => state.bridge)

  const balancesAndAllowances = usePortfolioBalances()

  if (!option) {
    return null
  }
  const { icon, name, symbol, routeSymbol } = option
  const parsedBalance = balancesAndAllowances[fromChainId]?.find(
    (token) => token.tokenAddress === option.addresses[fromChainId]
  )?.parsedBalance

  return (
    <div className="flex items-center justify-between">
      <div className="flex items-center space-x-2" key={option.symbol}>
        <img src={icon.src} className="w-5 h-5" />
        <div className="">
          <div className="text-primaryTextColor">{symbol}</div>
          <div className="text-xs select-hidden text-secondaryTextColor">
            {name}
          </div>
        </div>
      </div>
      <div className="select-hidden text-primaryTextColor">
        {parsedBalance !== '0.0' ? `${parsedBalance} ${symbol}` : ''}
      </div>
    </div>
  )
}

const FromTokenSelect = () => {
  const { fromChainId, fromToken, fromTokens } = useAppSelector(
    (state) => state.bridge
  )
  const balancesAndAllowances = usePortfolioBalances()

  const dispatch = useDispatch()

  const fromTokenOptions = [
    {
      label: 'Wallet',
      options: fromTokens
        .filter(
          (token) =>
            balancesAndAllowances[fromChainId]?.find(
              (tokenBalance) =>
                tokenBalance.tokenAddress === token.addresses[fromChainId]
            )?.balance !== 0n
        )
        .map((token) => ({
          label: <ImageAndCoin option={token} />,
          value: token,
        })),
    },
    {
      label: 'All tokens',
      options: fromTokens
        .filter(
          (token) =>
            balancesAndAllowances[fromChainId]?.find(
              (tokenBalance) =>
                tokenBalance.tokenAddress === token.addresses[fromChainId]
            )?.balance === 0n
        )
        .map((token) => ({
          label: <ImageAndCoin option={token} />,
          value: token,
        })),
    },
  ]

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
        option.value?.symbol.toLowerCase().includes(searchTerm) ||
        option.value?.name.toLowerCase().includes(searchTerm) ||
        option.value?.routeSymbol.toLowerCase().includes(searchTerm) ||
        (fromChainId &&
          option.value.addresses[fromChainId]
            .toLowerCase()
            .includes(searchTerm))
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
      value={fromTokenOptions
        .flatMap((group) => group.options)
        .find((option) => option.value === fromToken)}
    />
  )
}

export default FromTokenSelect
