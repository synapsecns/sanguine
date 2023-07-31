import _ from 'lodash'

import { EXISTING_BRIDGE_ROUTES } from '@/constants/existing-bridge-routes'
import { flattenPausedTokens } from './flattenPausedTokens'
import { Token } from './types'
import * as ALL_TOKENS from '@constants/tokens/master'

const getTokenAndChainId = (tokenAndChainId: string) => {
  const [symbol, chainId] = tokenAndChainId.split('-')

  return { symbol, chainId: Number(chainId) }
}

interface RouteQueryFields {
  fromChainId?: number
  fromTokenRouteSymbol?: string
  toChainId?: number
  toTokenRouteSymbol?: string
}

export const getFromChainIds = ({
  fromChainId,
  fromTokenRouteSymbol,
  toChainId,
  toTokenRouteSymbol,
}: RouteQueryFields) => {
  return _.chain(EXISTING_BRIDGE_ROUTES)
    .keys()
    .filter((key) => {
      const tokenAndChainId = getTokenAndChainId(key)
      const keySymbol = tokenAndChainId.symbol
      const keyChainId = tokenAndChainId.chainId

      // Check the key against fromChainId and fromToken
      if (fromTokenRouteSymbol && keySymbol !== fromTokenRouteSymbol)
        return false
      if (fromChainId && keyChainId === fromChainId) return true

      // If toChainId or toToken is provided, check the values for the key
      if (toChainId || toTokenRouteSymbol) {
        const values = EXISTING_BRIDGE_ROUTES[key]

        // Check the values against toChainId and toToken
        if (
          !values.some((value) => {
            const valueTokenAndChainId = getTokenAndChainId(value)
            const valueSymbol = valueTokenAndChainId.symbol
            const valueChainId = valueTokenAndChainId.chainId

            if (toChainId && valueChainId !== toChainId) return false
            if (toTokenRouteSymbol && valueSymbol !== toTokenRouteSymbol)
              return false

            return true
          })
        )
          return false
      }

      return true
    })
    .map((token) => getTokenAndChainId(token).chainId)
    .uniq()
    .value()
}

export const getFromTokens = ({
  fromChainId,
  fromTokenRouteSymbol,
  toChainId,
  toTokenRouteSymbol,
}: RouteQueryFields) => {
  return _.chain(EXISTING_BRIDGE_ROUTES)
    .keys()
    .filter((key) => {
      const tokenAndChainId = getTokenAndChainId(key)
      const keySymbol = tokenAndChainId.symbol
      const keyChainId = tokenAndChainId.chainId

      // Check the key against fromChainId and fromToken
      if (fromChainId && keyChainId !== fromChainId) return false

      // If toChainId or toToken is provided, check the values for the key
      if (toChainId || toTokenRouteSymbol) {
        const values = EXISTING_BRIDGE_ROUTES[key]

        // Check the values against toChainId and toToken
        return values.some((value) => {
          const valueTokenAndChainId = getTokenAndChainId(value)
          const valueSymbol = valueTokenAndChainId.symbol
          const valueChainId = valueTokenAndChainId.chainId

          if (toTokenRouteSymbol && valueSymbol !== toTokenRouteSymbol)
            return false
          if (toChainId && valueChainId !== toChainId) return false

          return true
        })
      }

      return true
    })
    .uniq()
    .value()
}

export const getToChainIds = ({
  fromChainId,
  fromTokenRouteSymbol,
  toChainId,
  toTokenRouteSymbol,
}: RouteQueryFields) => {
  if (!fromChainId && !fromTokenRouteSymbol) {
    if (toTokenRouteSymbol) {
      // Find keys with matching toTokenRouteSymbol and return their associated chainIds
      return _.chain(EXISTING_BRIDGE_ROUTES)
        .keys()
        .filter((key) => {
          const keySymbol = getTokenAndChainId(key).symbol
          return keySymbol === toTokenRouteSymbol
        })
        .map((key) => {
          // Include the chainId from the key
          return getTokenAndChainId(key).chainId
        })
        .uniq()
        .value()
    }

    if (toChainId) {
      return _.chain(EXISTING_BRIDGE_ROUTES)
        .values()
        .flatten()
        .filter((value) => {
          const valueChainId = getTokenAndChainId(value).chainId
          return valueChainId === toChainId
        })
        .map((value) => {
          const keyChainId = getTokenAndChainId(value).chainId
          return keyChainId
        })
        .uniq()
        .value()
    }
  }

  return _.chain(EXISTING_BRIDGE_ROUTES)
    .entries()
    .filter(([key, values]) => {
      const tokenAndChainId = getTokenAndChainId(key)
      const keySymbol = tokenAndChainId.symbol
      const keyChainId = tokenAndChainId.chainId

      // Check the key against fromChainId and fromToken
      if (fromTokenRouteSymbol && keySymbol !== fromTokenRouteSymbol)
        return false
      if (fromChainId && keyChainId !== fromChainId) return false

      // Check the values against toChainId and toToken
      if (
        !values.some((value) => {
          const valueTokenAndChainId = getTokenAndChainId(value)
          const valueSymbol = valueTokenAndChainId.symbol
          const valueChainId = valueTokenAndChainId.chainId

          if (toChainId && valueChainId !== toChainId) return false
          if (toTokenRouteSymbol && valueSymbol !== toTokenRouteSymbol)
            return false

          return true
        })
      )
        return false

      return true
    })
    .map(([_, values]) =>
      values.map((value) => getTokenAndChainId(value).chainId)
    )
    .flatten()
    .uniq()
    .value()
}

export const getToTokens = ({
  fromChainId,
  fromTokenRouteSymbol,
  toChainId,
  toTokenRouteSymbol,
}: RouteQueryFields) => {
  return _.chain(EXISTING_BRIDGE_ROUTES)
    .flatMap((values, key) => {
      const keyTokenAndChainId = getTokenAndChainId(key)
      const keySymbol = keyTokenAndChainId.symbol
      const keyChainId = keyTokenAndChainId.chainId

      // only consider the combination of fromToken and fromChainId if both are provided
      if (fromChainId && fromTokenRouteSymbol) {
        if (`${fromTokenRouteSymbol}-${fromChainId}` !== key) return []
      } else if (fromTokenRouteSymbol && keySymbol !== fromTokenRouteSymbol)
        return []
      else if (fromChainId && keyChainId !== fromChainId) return []

      // Filter out the destinations based on toChainId and toToken
      return values.filter((value) => {
        const valueTokenAndChainId = getTokenAndChainId(value)
        const valueSymbol = valueTokenAndChainId.symbol
        const valueChainId = valueTokenAndChainId.chainId

        // only consider the combination of toToken and toChainId if both are provided
        if (toChainId && toTokenRouteSymbol) {
          if (`${toTokenRouteSymbol}-${toChainId}` !== value) return false
        } else if (
          toTokenRouteSymbol &&
          toChainId &&
          valueSymbol !== toTokenRouteSymbol &&
          valueChainId !== toChainId
        )
          return false
        else if (toChainId && valueChainId !== toChainId) return false

        return true
      })
    })
    .uniq()
    .value()
}

// isBridgeableTo(fromToken, toChainId): true or false
// isBridgeableFrom(toToken, fromChainId): true or false

export const getRoutePossibilities = ({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: {
  fromChainId: number
  fromToken: Token
  toChainId: number
  toToken: Token
}) => {
  // check if fromToken is allowed in fromTokens
  // check if toToken is allowed in toTokens

  const fromTokenRouteSymbol = fromToken && fromToken.routeSymbol
  const toTokenRouteSymbol = toToken && toToken.routeSymbol

  const fromChainIds = getFromChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  }).filter((chainId) => chainId !== toChainId)

  const fromTokens = _.uniq(
    _.difference(
      getFromTokens({
        fromChainId,
        fromTokenRouteSymbol,
        toChainId,
        toTokenRouteSymbol,
      }),
      flattenPausedTokens()
    ).map(getSymbol)
  ).map((symbol) => ALL_TOKENS[symbol])

  const toChainIds = getToChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  }).filter((chainId) => chainId !== fromChainId)

  const toTokens = _.uniq(
    _.difference(
      getToTokens({
        fromChainId,
        fromTokenRouteSymbol,
        toChainId,
        toTokenRouteSymbol,
      }),
      flattenPausedTokens()
    ).map(getSymbol)
  ).map((symbol) => ALL_TOKENS[symbol])

  return {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    fromTokens,
    toChainIds,
    toTokens,
  }
}

const getSymbol = (tokenAndChainId: string): string => {
  return tokenAndChainId.split('-')[0]
}

/*NOTES

  * Set intelligent defaults when fromTokens/toTokens change
  * check if from/to token is allowed in list


  -- To strip out --
  * paused tokens [x]
  * swap exceptions?

*/
