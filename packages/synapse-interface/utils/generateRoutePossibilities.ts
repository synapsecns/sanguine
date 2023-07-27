import _ from 'lodash'

import { EXISTING_BRIDGE_ROUTES } from '@/constants/existing-bridge-routes'
import { flattenPausedTokens } from './flattenPausedTokens'

const getTokenAndChainId = (tokenAndChainId: string) => {
  const [symbol, chainId] = tokenAndChainId.split('-')

  return { symbol, chainId: Number(chainId) }
}

// generates object that matches each bridgeable fromToken to
// the first chainId that shows up in existing bridge routes
export const extractFirstChainIdBySymbol = () => {
  return _.chain(EXISTING_BRIDGE_ROUTES)
    .keys()
    .reduce((result, key) => {
      const [symbol, chainId] = key.split('-')
      if (!result[symbol]) {
        result[symbol] = Number(chainId)
      }
      return result
    }, {})
    .value()
}

// get fromChainIds
export const getPossibleFromChainIds = () => {
  return _.uniq(
    Object.keys(EXISTING_BRIDGE_ROUTES).map(
      (token) => getTokenAndChainId(token).chainId
    )
  )
}

export const getPossibleFromChainIdsByFromToken = (fromToken: string) => {
  const symbol = getTokenAndChainId(fromToken).symbol

  return _.chain(EXISTING_BRIDGE_ROUTES)
    .keys()
    .filter((key) => _.startsWith(key, symbol))
    .map((token) => getTokenAndChainId(token).chainId)
    .uniq()
    .value()
}

export const getPossibleFromChainIdsByToChainId = (toChainId: number) => {
  return _.uniq(
    Object.keys(EXISTING_BRIDGE_ROUTES)
      .filter((key) =>
        EXISTING_BRIDGE_ROUTES[key].some((value) =>
          value.endsWith(`-${toChainId}`)
        )
      )
      .map((token) => getTokenAndChainId(token).chainId)
  )
}

export const getPossibleFromChainIdsByToToken = (toToken: string) => {
  return _.uniq(
    Object.entries(EXISTING_BRIDGE_ROUTES)
      .filter(([, values]) => values.includes(toToken))
      .map(([key]) => {
        const chainId = getTokenAndChainId(key).chainId
        return Number(chainId)
      })
  )
}

// get fromTokens
export const getPossibleFromTokens = () => {
  return Object.keys(EXISTING_BRIDGE_ROUTES)
}

export const getPossibleFromTokensByFromChainId = (fromChainId: number) => {
  return Object.keys(EXISTING_BRIDGE_ROUTES).filter((token) =>
    token.endsWith(`-${fromChainId}`)
  )
}

export const getPossibleFromTokensByFromToken = (fromToken: string) => {
  const chainId = getTokenAndChainId(fromToken).chainId

  return getPossibleFromTokensByFromChainId(chainId)
}

export const getPossibleFromTokensByFromChainIdAndToChainId = (
  fromChainId: number,
  toChainId: number
) => {
  const fromOptions = Object.keys(EXISTING_BRIDGE_ROUTES).filter((token) =>
    token.endsWith(`-${fromChainId}`)
  )

  return _.uniq(
    fromOptions.filter((key) =>
      EXISTING_BRIDGE_ROUTES[key].some((value) =>
        value.endsWith(`-${toChainId}`)
      )
    )
  )
}

export const getPossibleFromTokensByToChainId = (toChainId: number) => {
  return _.uniq(
    Object.keys(EXISTING_BRIDGE_ROUTES).filter((key) =>
      EXISTING_BRIDGE_ROUTES[key].some((value) =>
        value.endsWith(`-${toChainId}`)
      )
    )
  )
}

export const getPossibleFromTokensByToToken = (toToken) => {
  return _.uniq(
    Object.entries(EXISTING_BRIDGE_ROUTES)
      .filter(([, values]) => values.includes(toToken))
      .map(([key]) => key)
  )
}

// get toChainIds
export const getPossibleToChainIds = () => {
  return _(EXISTING_BRIDGE_ROUTES)
    .values()
    .flatten()
    .map((token) => getTokenAndChainId(token).chainId)
    .uniq()
    .value()
}

export const getPossibleToChainIdsByFromChainId = (fromChainId: number) => {
  return _(EXISTING_BRIDGE_ROUTES)
    .pickBy((_value, key) => key.endsWith(`-${fromChainId}`))
    .values()
    .flatten()
    .map((token) => getTokenAndChainId(token).chainId)
    .uniq()
    .value()
}

export const getPossibleToChainIdsByFromToken = (fromToken: string) => {
  return _.uniq(
    EXISTING_BRIDGE_ROUTES[fromToken].map(
      (token) => getTokenAndChainId(token).chainId
    )
  )
}

export const getPossibleToChainIdsByToToken = (toToken: string) => {
  return _.uniq(
    Object.entries(EXISTING_BRIDGE_ROUTES)
      .filter(([, values]) => values.includes(toToken))
      .map(([key]) => Number(key.split('-')[1]))
  )
}

// get toTokens
export const getPossibleToTokens = () => {
  return _(EXISTING_BRIDGE_ROUTES).values().flatten().uniq().value()
}

export const getPossibleToTokensByFromChainId = (fromChainId: number) => {
  return _(EXISTING_BRIDGE_ROUTES)
    .pickBy((_value, key) => key.endsWith(`-${fromChainId}`))
    .values()
    .flatten()
    .value()
}

export const getPossibleToTokensByFromToken = (fromToken: string) => {
  return EXISTING_BRIDGE_ROUTES[fromToken]
}

export const getPossibleToTokensByFromTokenAndToChainId = (
  fromToken: string,
  toChainId: number
) => {
  return getPossibleToTokensByFromToken(fromToken).filter((token) =>
    token.endsWith(`-${toChainId}`)
  )
}

export const getPossibleToTokensByToChainId = (toChainId: number) => {
  return _(EXISTING_BRIDGE_ROUTES)
    .values()
    .flatten()
    .filter((value) => value.endsWith(`-${toChainId}`))
    .uniq()
    .value()
}

// is this right??
export const getPossibleToTokensByToToken = (toToken: string) => {
  const chainId = getTokenAndChainId(toToken).chainId

  return getPossibleToTokensByToChainId(chainId)
}

export const generateRoutePossibilities = ({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: {
  fromChainId: number
  fromToken: string
  toChainId: number
  toToken: string
}) => {
  let newFromChainId
  let newFromToken
  let newToChainId
  let newToToken
  let fromChainIds
  let fromTokens
  let toChainIds
  let toTokens

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    fromChainIds = getPossibleFromChainIds()
    fromTokens = getPossibleFromTokens()
    toChainIds = getPossibleToChainIds()
    toTokens = getPossibleToTokens()

    newFromChainId = null
    newFromToken = null
    newToChainId = null
    newToToken = null
  }

  if (fromChainId === null && fromToken === null && toChainId && toToken) {
    fromChainIds = getPossibleFromChainIdsByToToken(toToken)
    fromTokens = getPossibleFromTokensByToToken(toToken)
    toChainIds = getPossibleToChainIdsByToToken(toToken)
    toTokens = getPossibleToTokensByToChainId(toChainId)

    newFromChainId = null
    newFromToken = null
    newToChainId = toChainId
    newToToken = toToken
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId &&
    toToken === null
  ) {
    fromChainIds = getPossibleFromChainIdsByToChainId(toChainId)
    fromTokens = getPossibleFromTokensByToChainId(toChainId)
    toChainIds = getPossibleToChainIds()
    toTokens = getPossibleToTokensByToChainId(toChainId)

    newFromChainId = null
    newFromToken = null
    newToChainId = toChainId
    newToToken = null
  }

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId === null &&
    toToken
  ) {
    fromChainIds = getPossibleFromChainIdsByToToken(toToken)
    fromTokens = getPossibleFromTokensByToToken(toToken)
    toChainIds = getPossibleToChainIdsByToToken(toToken)
    toTokens = getPossibleToTokensByToToken(toToken)

    newFromChainId = null
    newFromToken = null
    newToChainId = null
    newToToken = toToken
  }

  if (fromChainId && fromToken && toChainId === null && toToken === null) {
    fromChainIds = getPossibleFromChainIdsByFromToken(fromToken)
    fromTokens = getPossibleFromTokensByFromToken(fromToken)
    toChainIds = getPossibleToChainIdsByFromToken(fromToken)
    toTokens = getPossibleToTokensByFromToken(fromToken)

    newFromChainId = fromChainId
    newFromToken = fromToken
    newToChainId = null
    newToToken = null
  }

  if (fromChainId && fromToken && toChainId && toToken) {
    fromChainIds = getPossibleFromChainIds()
    fromTokens = getPossibleFromTokensByFromChainId(fromChainId)
    toChainIds = getPossibleToChainIdsByFromToken(fromToken)
    toTokens = getPossibleToTokensByFromTokenAndToChainId(fromToken, toChainId)

    newFromChainId = fromChainId
    newFromToken = fromToken
    newToChainId = toChainId
    newToToken = toToken
  }

  if (fromChainId && fromToken && toChainId && toToken === null) {
    fromChainIds = getPossibleFromChainIds()
    fromTokens = getPossibleFromTokensByFromChainId(fromChainId)
    toChainIds = getPossibleToChainIdsByFromToken(fromToken)
    toTokens = getPossibleToTokensByFromTokenAndToChainId(fromToken, toChainId)

    newFromChainId = fromChainId
    newFromToken = fromToken
    newToChainId = toChainId
    newToToken = null
  }

  if (fromChainId === null && fromToken && toChainId && toToken) {
    fromChainIds = getPossibleFromChainIds()
    fromTokens = getPossibleFromTokensByFromToken(fromToken)
    toChainIds = getPossibleToChainIdsByFromToken(fromToken)
    toTokens = getPossibleToTokensByFromTokenAndToChainId(fromToken, toChainId)

    newFromChainId = null
    newFromToken = fromToken
    newToChainId = toChainId
    newToToken = toToken
  }

  if (
    fromChainId &&
    fromToken === null &&
    toChainId === null &&
    toToken === null
  ) {
    fromChainIds = getPossibleFromChainIds()
    fromTokens = getPossibleFromTokensByFromChainId(fromChainId)
    toChainIds = getPossibleToChainIdsByFromChainId(fromChainId)
    toTokens = getPossibleToTokensByFromChainId(fromChainId)

    newFromChainId = fromChainId
    newFromToken = null
    newToChainId = null
    newToToken = null
  }

  if (fromChainId && fromToken === null && toChainId === null && toToken) {
    fromChainIds = getPossibleFromChainIdsByToToken(toToken)
    fromTokens = getPossibleFromTokensByToToken(toToken)
    toChainIds = getPossibleToChainIdsByToToken(toToken)
    toTokens = getPossibleToTokensByToToken(toToken)

    newFromChainId = fromChainId
    newFromToken = null
    newToChainId = null
    newToToken = toToken
  }

  if (fromChainId && fromToken === null && toChainId && toToken === null) {
    fromChainIds = getPossibleFromChainIdsByToChainId(toChainId)
    fromTokens = getPossibleFromTokensByFromChainIdAndToChainId(
      fromChainId,
      toChainId
    )
    toChainIds = getPossibleToChainIdsByFromChainId(fromChainId)
    toTokens = getPossibleToTokensByToChainId(toChainId)

    newFromChainId = fromChainId
    newFromToken = null
    newToChainId = toChainId
    newToToken = null
  }

  if (fromChainId && fromToken === null && toChainId && toToken) {
    fromChainIds = getPossibleFromChainIdsByToToken(toToken)
    fromTokens = getPossibleFromTokensByFromChainIdAndToChainId(
      fromChainId,
      toChainId
    )

    toChainIds = getPossibleToChainIdsByToToken(toToken)
    toTokens = getPossibleToTokensByToToken(toToken)

    newFromChainId = fromChainId
    newFromToken = null
    newToChainId = toChainId
    newToToken = toToken
  }

  if (
    fromChainId === null &&
    fromToken &&
    toChainId === null &&
    toToken === null
  ) {
    fromChainIds = getPossibleFromChainIdsByFromToken(fromToken)
    fromTokens = getPossibleFromTokensByFromToken(fromToken)
    toChainIds = getPossibleToChainIdsByFromToken(fromToken)
    toTokens = getPossibleToTokensByFromToken(fromToken)

    newFromChainId = null
    newFromToken = fromToken
    newToChainId = null
    newToToken = null
  }

  return {
    fromChainId: newFromChainId,
    fromToken: newFromToken,
    toChainId: newToChainId,
    toToken: newToToken,
    fromChainIds,
    fromTokens: _.difference(fromTokens, flattenPausedTokens()),
    toChainIds,
    toTokens: _.difference(toTokens, flattenPausedTokens()),
  }
}

/*NOTES

  * Set intelligent defaults when fromTokens/toTokens change
  * check if from/to token is allowed in list


  -- To strip out --
  * paused tokens
  * swap exceptions?

*/
