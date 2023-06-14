import _ from 'lodash'
import { BigNumber } from 'ethers'

import { WETH } from '@/constants/tokens/swapMaster'
import { Token } from './types'

const replaceKey = (
  obj: Record<string, BigNumber>,
  oldKey: string,
  newKey: string
) => {
  if (obj.hasOwnProperty(oldKey)) {
    obj[newKey] = obj[oldKey]
    delete obj[oldKey]
  }
  return obj
}

export const transformCalculateLiquidityInput = (
  chainId: number,
  pool: Token,
  filteredInputValue?: Record<string, BigNumber>
): Record<string, BigNumber> => {
  const wethIndex = _.findIndex(
    pool.poolTokens,
    (t) => t.symbol === WETH.symbol
  )
  const poolHasWeth: boolean = wethIndex > 0
  const nativeEthAddress = '0x0000000000000000000000000000000000000000'
  const wethAddress = poolHasWeth
    ? pool.poolTokens[wethIndex].addresses[chainId]
    : null

  const transformedInput = poolHasWeth
    ? replaceKey(filteredInputValue, nativeEthAddress, wethAddress)
    : filteredInputValue

  return transformedInput
}
