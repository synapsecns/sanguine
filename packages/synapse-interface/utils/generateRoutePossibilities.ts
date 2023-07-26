import _ from 'lodash'

import { EXISTING_BRIDGE_ROUTES } from '@/constants/existing-bridge-routes'

const getTokenAndChainId = (tokenAndChainId: string) => {
  const [symbol, chainId] = tokenAndChainId.split('-')

  return { symbol, chainId: Number(chainId) }
}

// get fromChainIds
const getPossibleFromChainIds = () => {
  return _.uniq(
    Object.keys(EXISTING_BRIDGE_ROUTES).map(
      (token) => getTokenAndChainId(token).chainId
    )
  )
}

const getPossibleFromChainIdsByFromToken = (fromToken: string) => {
  const symbol = getTokenAndChainId(fromToken).symbol

  return _.uniq(
    Object.keys(EXISTING_BRIDGE_ROUTES)
      .filter((key) => key.startsWith(symbol))
      .map((token) => getTokenAndChainId(token).chainId)
  )
}

const getPossibleFromChainIdsByToChainId = (toChainId: number) => {
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

const getPossibleFromChainIdsByToToken = (toToken: string) => {
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
const getPossibleFromTokens = () => {
  return Object.keys(EXISTING_BRIDGE_ROUTES)
}

export const getPossibleFromTokensByFromChainId = (fromChainId: number) => {
  return Object.keys(EXISTING_BRIDGE_ROUTES).filter((token) =>
    token.endsWith(`-${fromChainId}`)
  )
}

const getPossibleFromTokensByFromToken = (fromToken: string) => {
  const chainId = getTokenAndChainId(fromToken).chainId

  return getPossibleFromTokensByFromChainId(chainId)
}

const getPossibleFromTokensByFromChainIdAndToChainId = (
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

const getPossibleFromTokensByToChainId = (toChainId: number) => {
  return _.uniq(
    Object.keys(EXISTING_BRIDGE_ROUTES).filter((key) =>
      EXISTING_BRIDGE_ROUTES[key].some((value) =>
        value.endsWith(`-${toChainId}`)
      )
    )
  )
}

const getPossibleFromTokensByToToken = (toToken) => {
  return _.uniq(
    Object.entries(EXISTING_BRIDGE_ROUTES)
      .filter(([, values]) => values.includes(toToken))
      .map(([key]) => key)
  )
}

// get toChainIds
const getPossibleToChainIds = () => {
  return _(EXISTING_BRIDGE_ROUTES)
    .values()
    .flatten()
    .map((token) => getTokenAndChainId(token).chainId)
    .uniq()
    .value()
}

const getPossibleToChainIdsByFromChainId = (fromChainId: number) => {
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

const getPossibleToChainIdsByToToken = (toToken: string) => {
  return getPossibleToChainIds()
}

// get toTokens
const getPossibleToTokens = () => {
  return _(EXISTING_BRIDGE_ROUTES).values().flatten().uniq().value()
}

const getPossibleToTokensByFromChainId = (fromChainId: number) => {
  return _(EXISTING_BRIDGE_ROUTES)
    .pickBy((_value, key) => key.endsWith(`-${fromChainId}`))
    .values()
    .flatten()
    .value()
}

const getPossibleToTokensByFromToken = (fromToken: string) => {
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

const getPossibleToTokensByToChainId = (toChainId) => {
  return _(EXISTING_BRIDGE_ROUTES)
    .values()
    .flatten()
    .filter((value) => value.endsWith(`-${toChainId}`))
    .uniq()
    .value()
}

const getPossibleToTokensByToToken = (toToken) => {
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
    toTokens = getPossibleToTokensByToChainId(toToken)

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

  if (
    fromChainId === null &&
    fromToken === null &&
    toChainId &&
    toToken === null
  ) {
    fromChainIds = getPossibleFromChainIdsByToChainId(toChainId)
    fromTokens = getPossibleFromTokensByToChainId(toChainId)
    toChainIds = getPossibleToChainIds()
    toTokens = getPossibleToTokensByToChainId(toToken)

    newFromChainId = null
    newFromToken = null
    newToChainId = toChainId
    newToToken = null
  }

  return {
    fromChainId: newFromChainId,
    fromToken: newFromToken,
    toChainId: newToChainId,
    toToken: newToToken,
    fromChainIds,
    fromTokens,
    toChainIds,
    toTokens,
  }
}
