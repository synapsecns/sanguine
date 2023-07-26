import _ from 'lodash'
import { setSelectFromToken } from '@/slices/tokenSelectorSlice'
import { RootState } from '@/store/store'
import { useSelector } from 'react-redux'
import { useDispatch } from 'react-redux'
import Select from 'react-select'

import * as ALL_TOKENS from '@/constants/tokens/master'
import { Token } from '@/utils/types'
import { setFromToken } from '@/slices/bridge/reducer'
import { coinSelectStyles } from './styles/coinSelectStyles'
import { useId } from 'react'
import { flattenPausedTokens } from '@/utils/flattenPausedTokens'

const ImageAndCoin = ({ symbol }) => {
  const { icon } = ALL_TOKENS[symbol]
  return (
    <div className="flex items-center space-x-2" key={symbol}>
      <img src={icon.src} className="w-6 h-6" />
      <div className="text-xl">{symbol}</div>
    </div>
  )
}

const FromTokenSelect = () => {
  const { fromToken, fromTokens } = useSelector(
    (state: RootState) => state.tokenSelector
  )

  const dispatch = useDispatch()

  const fromTokenOptions = _.difference(fromTokens, flattenPausedTokens()).map(
    (option) => {
      const symbol = option.split('-')[0]
      return {
        label: <ImageAndCoin symbol={symbol} />,
        value: option,
      }
    }
  )

  const handleFromTokenChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setSelectFromToken(selectedOption.value))

      const symbol = selectedOption.value.split('-')[0]

      const token: Token = ALL_TOKENS[symbol]

      dispatch(setFromToken(token))
    } else {
      dispatch(setFromToken(null))
      dispatch(setSelectFromToken(null))
    }
  }

  return (
    <Select
      instanceId={useId()}
      styles={coinSelectStyles}
      key={fromToken}
      options={fromTokenOptions}
      onChange={handleFromTokenChange}
      isSearchable={true}
      placeholder={<span className="text-xl text-white">In</span>}
      value={fromTokenOptions.find((option) => option.value === fromToken)}
    />
  )
}

export default FromTokenSelect
