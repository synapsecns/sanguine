import { zeroAddress } from 'viem'
import _ from 'lodash'

import * as BRIDGEABLE from '@/constants/tokens/bridgeable'
import { BRIDGE_MAP } from '@/constants/bridgeMap'
import { ETHEREUM_ADDRESS } from '@/constants'

export const generateChainIdAddressMapping = (routeSymbol: string) => {
  const result: { [key: number]: string } = {}

  Object.entries(BRIDGE_MAP).forEach(([chainId, tokens]) => {
    Object.entries(tokens).forEach(([address, token]) => {
      if (token.symbol === routeSymbol) {
        result[Number(chainId)] =
          address === ETHEREUM_ADDRESS ? zeroAddress : address
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
    _.map(BRIDGEABLE, (token) => {
      return token.routeSymbol
    })
  )
}

/*
Looks for discrepancies in tokens master file vs. the bridgeMap
*/
export const testerMasterVsJson = () => {
  Object.entries(BRIDGEABLE).forEach(([_key, token]) => {
    const map = generateChainIdAddressMapping(token.routeSymbol)

    if (!caseInsensitiveDeepEqual(token.addresses, map)) {
      console.log(token.routeSymbol)
      console.log(`token.addresses count`, Object.keys(token.addresses).length)
      console.log(`map`, Object.keys(map).length)
      console.log(`token.addresses`, token.addresses)
      console.log(`map`, map)
      console.log(
        'same addresses',
        caseInsensitiveDeepEqual(token.addresses, map)
      )
      console.log('\n')
    }
  })
}

const caseInsensitiveDeepEqual = (obj1, obj2) => {
  const formattedObj1 = _.mapValues(obj1, (v) => v.toLowerCase())
  const formattedObj2 = _.mapValues(obj2, (v) => v.toLowerCase())

  return _.isEqual(formattedObj1, formattedObj2)
}
