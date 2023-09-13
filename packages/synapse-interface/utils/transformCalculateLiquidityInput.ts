import _ from 'lodash'
import { zeroAddress } from 'viem'

import { WETH } from '@/constants/tokens/bridgeable'
import { Token } from './types'

const replaceKey = (
  obj: Record<string, bigint>,
  oldKey: string,
  newKey: string
) => {
  if (obj.hasOwnProperty(oldKey)) {
    const newObj = { ...obj, [newKey]: obj[oldKey] }
    delete newObj[oldKey]

    return newObj
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

  const wethAddress = poolHasWeth
    ? pool.poolTokens[wethIndex].addresses[chainId]
    : null

  const transformedInput = poolHasWeth
    ? replaceKey(filteredInputValue, zeroAddress, wethAddress)
    : filteredInputValue

  return transformedInput
}
