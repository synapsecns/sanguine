import _ from 'lodash'

import type { Token } from '@/utils/types'
import type { NetworkTokenBalances } from '@/utils/actions/fetchPortfolioBalances'

export const hasBalance = (
  t: Token,
  chainId: number,
  portfolioBalances: NetworkTokenBalances
) => {
  if (!chainId) {
    return false
  }
  const pb = portfolioBalances[chainId]
  if (!pb) {
    return false
  }
  const token = _.pickBy(pb, (value, _key) => value.token === t)

  const tokenWithPb = Object.values(token)[0]

  return tokenWithPb && tokenWithPb.balance !== 0n
}
