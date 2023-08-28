import { zeroAddress } from 'viem'
import _ from 'lodash'
import * as ALL_TOKENS from '@constants/tokens/bridgeable'
import { BRIDGE_MAP } from '@constants/bridgeMap'

export const generateChainIdAddressMapping = (routeSymbol: string) => {
  const result: { [key: number]: string } = {}

  Object.entries(BRIDGE_MAP).forEach(([chainId, tokens]) => {
    Object.entries(tokens).forEach(([address, token]) => {
      if (token.symbol === routeSymbol) {
        result[Number(chainId)] =
          address === '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'
            ? zeroAddress
            : address
      }
    })
  })

  return result
}

export const getAllSymbols = () => {
  return _.sortBy(
    _(BRIDGE_MAP)
      .values()
      .flatMap((tokens) => _.map(tokens, 'symbol'))
      .uniq()
      .value()
  )
}

export const getAllInternalSymbols = () => {
  return _.sortBy(
    _.map(ALL_TOKENS, (token) => {
      return token.routeSymbol
    })
  )
}

/*
Looks for discrepancies in tokens master file vs. the bridgeMap
*/
export const testerMasterVsJson = () => {
  Object.entries(ALL_TOKENS).forEach(([_key, token]) => {
    const map = generateChainIdAddressMapping(token.routeSymbol)

    if (!_.isEqual(token.addresses, map)) {
      console.log(token.routeSymbol)
      console.log(`token.addresses count`, Object.keys(token.addresses).length)
      console.log(`map`, Object.keys(map).length)
      console.log(`token.addresses`, token.addresses)
      console.log(`map`, map)
      console.log('same addresses', _.isEqual(token.addresses, map))
      console.log('\n')
    }
  })
}
