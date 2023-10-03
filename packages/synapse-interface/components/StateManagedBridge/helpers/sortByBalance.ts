import _ from 'lodash'

import { Token } from '@/utils/types'
import { NetworkTokenBalancesAndAllowances } from '@/utils/actions/fetchPortfolioBalances'

export const hasBalance = (
  t: Token,
  chainId: number,
  portfolioBalances: NetworkTokenBalancesAndAllowances
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
