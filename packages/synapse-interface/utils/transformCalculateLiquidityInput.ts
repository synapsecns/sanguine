import _ from 'lodash'
import { zeroAddress } from 'viem'

import { WETH } from '@/constants/tokens/swapMaster'
import { Token } from './types'

const replaceKey = (
  obj: Record<string, bigint>,
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
  filteredInputValue?: Record<string, bigint>
): Record<string, bigint> => {
  const wethIndex = _.findIndex(
    pool.poolTokens,
    (t) => t.symbol === WETH.symbol
  )
  const poolHasWeth: boolean = wethIndex > 0
  const nativeEthAddress = zeroAddress
  const wethAddress = poolHasWeth
    ? pool.poolTokens[wethIndex].addresses[chainId]
    : null

  const transformedInput = poolHasWeth
    ? replaceKey(filteredInputValue, nativeEthAddress, wethAddress)
    : filteredInputValue

  return transformedInput
}
