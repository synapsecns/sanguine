import _ from 'lodash'
import { zeroAddress } from 'viem'

import { BRIDGE_MAP } from './bridgeMap'
import { findTokenByAddressAndChain } from '@/utils/findTokenByAddressAndChainId'
import { ETHEREUM_ADDRESS } from '.'

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
        if (tokenAddress.toLowerCase() === ETHEREUM_ADDRESS.toLowerCase()) {
          tokenAddress = zeroAddress
        }

        const symbol = findTokenByAddressAndChain(
          tokenAddress,
          chainId
        )?.routeSymbol
        const key = `${symbol}-${chainId}`
        const swappable = info.swappable.map((address) => {
          if (address.toLowerCase() === ETHEREUM_ADDRESS.toLowerCase()) {
            address = zeroAddress
          }

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
