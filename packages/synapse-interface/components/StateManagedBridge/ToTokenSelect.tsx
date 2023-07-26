import _ from 'lodash'
import { setSelectToToken } from '@/slices/tokenSelectorSlice'
import { RootState } from '@/store/store'
import { useSelector } from 'react-redux'
import { useDispatch } from 'react-redux'
import Select from 'react-select'
import * as ALL_TOKENS from '@/constants/tokens/master'
import { Token } from '@/utils/types'

import { setToToken } from '@/slices/bridge/reducer'
import { coinSelectStyles } from './styles/coinSelectStyles'
import { useId } from 'react'
import { flattenPausedTokens } from '@/utils/flattenPausedTokens'

const ToTokenSelect = () => {
  const { toToken, toTokens } = useSelector(
    (state: RootState) => state.tokenSelector
  )

  const dispatch = useDispatch()

  const toTokenOptions = _.difference(toTokens, flattenPausedTokens()).map(
    (option) => ({
      label: option,
      value: option,
    })
  )

  const handleToTokenChange = (selectedOption) => {
    if (selectedOption) {
      dispatch(setSelectToToken(selectedOption.value))

      const symbol = selectedOption.value.split('-')[0]

      const token: Token = ALL_TOKENS[symbol]
      dispatch(setToToken(token))
    } else {
      dispatch(setSelectToToken(null))
    }
  }

  return (
    <Select
      instanceId={useId()}
      styles={coinSelectStyles}
      key={toToken}
      options={toTokenOptions}
      onChange={handleToTokenChange}
      isSearchable={true}
      placeholder={<span className="text-xl text-white">Out</span>}
      value={toTokenOptions.find((option) => option.value === toToken)}
    />
  )
}

export default ToTokenSelect
