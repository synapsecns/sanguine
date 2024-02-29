import _ from 'lodash'

import { Token } from '@/utils/types'
import { NetworkTokenBalances } from '@/utils/actions/fetchPortfolioBalances'

export const hasBalance = (
  t: Token,
  chainId: number,
  portfolioBalances: NetworkTokenBalances
) => {
  if (!chainId) {
    return false
  }
  const pb = portfolioBalances[chainId]
  const token = _(pb)
    .pickBy((value, _key) => value.token === t)
    .value()
  const tokenWithPb = Object.values(token)[0]

  return tokenWithPb?.balance !== 0n
}
