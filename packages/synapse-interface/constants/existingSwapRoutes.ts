import _ from 'lodash'

import { BRIDGE_MAP } from './bridgeMap'
import { findTokenByAddressAndChain } from '@/utils/findTokenByAddressAndChainId'

export const FILTERED = _(BRIDGE_MAP)
  .mapValues((chainObj) => {
    return _(chainObj)
      .pickBy(
        (tokenObj: any) =>
          Array.isArray(tokenObj.swappable) && tokenObj.swappable.length > 0
      )
      .value()
  })
  .pickBy((value, _key) => Object.values(value).length > 0)
  .value()

export const SWAP_CHAIN_IDS = Object.keys(FILTERED).map(Number)

export const EXISTING_SWAP_ROUTES = _(FILTERED)
  .map((tokens, chainId) => {
    return _(tokens)
      .map((info, tokenAddress) => {
        const symbol = findTokenByAddressAndChain(
          tokenAddress,
          chainId
        )?.routeSymbol
        const key = `${symbol}-${chainId}`
        const swappable = info.swappable.map((address) => {
          const symbol = findTokenByAddressAndChain(
            address,
            chainId
          )?.routeSymbol
          return `${symbol}-${chainId}`
        })
        return [key, swappable]
      })
      .value()
  })
  .flatten()
  .fromPairs()
  .value()
