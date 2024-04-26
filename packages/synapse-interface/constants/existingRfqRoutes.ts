import { zeroAddress } from 'viem'

import * as ALL_BRIDGEABLE_TOKENS from '@/constants/tokens/bridgeable'
import { ETHEREUM_ADDRESS } from '@/constants'

const reduceToSymbol = (tokenAddressChainId: string) => {
  return tokenAddressToRouteSymbolMap[tokenAddressChainId.toLowerCase()]
}

const tokenAddressToRouteSymbolMap = Object.values(
  ALL_BRIDGEABLE_TOKENS
).reduce((acc, token) => {
  Object.entries(token.addresses).forEach(([chainId, address]) => {
    const effectiveAddress =
      address === zeroAddress ? ETHEREUM_ADDRESS : address
    const key = `${effectiveAddress.toLowerCase()}-${chainId}`
    acc[key] = `${token.routeSymbol}-${chainId}`
  })
  return acc
}, {})

export const transformRFQMap = (rfqMap) => {
  const transformedMap = {}

  Object.keys(rfqMap).forEach((key) => {
    const transformedKey = reduceToSymbol(key)
    if (!transformedKey) {
      console.error(`No match for key: ${key}`)
    }

    const transformedValues = rfqMap[key].map((value) => {
      const transformedValue = reduceToSymbol(value)
      if (!transformedValue) {
        console.error(`No match for value: ${value}`)
      }
      return transformedValue
    })

    transformedMap[transformedKey] = transformedValues.filter(
      (v) => v !== undefined
    )
  })

  return transformedMap
}
