import _ from 'lodash'

import { Token } from '@/utils/types'

export const sortByBalances = (
  t: Token,
  chainId: number,
  portfolioBalances: any
) => {
  const pb = portfolioBalances[chainId]
  const token = _(pb)
    .pickBy((value, _key) => value.token === t)
    .value()

  const tokenWithPb = Object.values(token)[0]

  return Object.keys(token).length > 0 && tokenWithPb?.balance !== 0n ? -1 : 1
}
