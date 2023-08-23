import _ from 'lodash'

import { EXISTING_BRIDGE_ROUTES } from '@/constants/existingBridgeRoutes'
import { RouteQueryFields } from './generateRoutePossibilities'
import { getTokenAndChainId } from './getTokenAndChainId'

export const getAllFromChainIds = () => {
  return _(EXISTING_BRIDGE_ROUTES)
    .keys()
    .map((token) => getTokenAndChainId(token).chainId)
    .uniq()
    .value()
}

export const getFromChainIds = ({
  fromChainId,
  fromTokenRouteSymbol,
  toChainId,
  toTokenRouteSymbol,
}: RouteQueryFields) => {
  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .keys()
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .keys()
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .keys()
      .filter((key) => {
        const { symbol } = getTokenAndChainId(key)
        return symbol === fromTokenRouteSymbol
      })
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .keys()
      .filter((key) => {
        const { symbol } = getTokenAndChainId(key)
        return symbol === fromTokenRouteSymbol
      })
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .entries()
      .filter(([_key, values]) =>
        values.some((v) => v.endsWith(`-${toChainId}`))
      )
      .map(([key]) => getTokenAndChainId(key).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .entries()
      .filter(([_key, values]) =>
        values.some((v) => v.endsWith(`-${toChainId}`))
      )
      .map(([key]) => getTokenAndChainId(key).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .entries()
      .filter(([key, _values]) => key.startsWith(`${fromTokenRouteSymbol}-`))
      .filter(([_key, values]) =>
        values.some((v) => getTokenAndChainId(v).chainId === toChainId)
      )
      .map(([key, _values]) => key)
      .filter((token) => token.endsWith(`-${toChainId}`))
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .pickBy((_values, key) => key.startsWith(`${fromTokenRouteSymbol}-`))
      .keys()
      .map((token) => getTokenAndChainId(token).chainId)
      .filter((chainId) => chainId !== toChainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .chain()
      .filter((values, _key) => {
        return values.some((v) => {
          const { symbol } = getTokenAndChainId(v)
          return symbol === toTokenRouteSymbol
        })
      })
      .flatten()
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .chain()
      .filter((values, _key) => {
        return values.some((v) => {
          const { symbol } = getTokenAndChainId(v)
          return symbol === toTokenRouteSymbol
        })
      })
      .flatten()
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .entries()
      .filter(([key, _values]) => key.startsWith(`${fromTokenRouteSymbol}-`))
      .map(([key, _values]) => key)
      .flatten()
      .filter((token) => token.startsWith(`${toTokenRouteSymbol}-`))
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .chain()
      .filter((values, key) => {
        return (
          values.some((v) => {
            const { symbol } = getTokenAndChainId(v)
            return symbol === toTokenRouteSymbol
          }) && key.startsWith(`${fromTokenRouteSymbol}-`)
        )
      })
      .flatten()
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .pickBy((values, _key) => {
        return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
      })
      .keys()
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .pickBy((values, _key) => {
        return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
      })
      .keys()
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
      .pickBy((values, _key) => {
        return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
      })
      .keys()
      .filter((k) => k.startsWith(`${fromTokenRouteSymbol}-`))
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }

  if (fromChainId && fromTokenRouteSymbol && toChainId && toTokenRouteSymbol) {
    return _(EXISTING_BRIDGE_ROUTES)
      .pickBy((values, _key) => {
        return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
      })
      .keys()
      .filter((k) => k.startsWith(`${fromTokenRouteSymbol}-`))
      .map((token) => getTokenAndChainId(token).chainId)
      .uniq()
      .value()
  }
}
