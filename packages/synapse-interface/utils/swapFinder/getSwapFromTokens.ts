import _ from 'lodash'

import { EXISTING_SWAP_ROUTES } from '@/constants/existingSwapRoutes'
import { RouteQueryFields } from './generateSwapPossibilities'
import { getTokenAndChainId } from './getTokenAndChainId'

export const getSwapFromTokens = ({
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
    return _(EXISTING_SWAP_ROUTES).keys().uniq().value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .keys()
      .filter((token) => token.endsWith(`-${fromChainId}`))
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_SWAP_ROUTES).keys().uniq().value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .keys()
      .filter((key) => getTokenAndChainId(key).chainId === fromChainId)
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) => values.some((v) => v.endsWith(`-${toChainId}`)))
      .keys()
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) => values.some((v) => v.endsWith(`-${toChainId}`)))
      .keys()
      .filter((key) => key.endsWith(`-${fromChainId}`))
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .entries()
      .filter(([key, _values]) => key.startsWith(`${fromTokenRouteSymbol}-`))
      .filter(([_key, values]) =>
        values.some((v) => getTokenAndChainId(v).chainId === toChainId)
      )
      .map(([key, _values]) => key)
      .filter((token) => token.endsWith(`-${toChainId}`))
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) => values.some((v) => v.endsWith(`-${toChainId}`)))
      .pickBy((_values, key) => key.endsWith(`-${fromChainId}`))
      .keys()
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .chain()
      .filter((values, _key) =>
        values.some((v) => getTokenAndChainId(v).symbol === toTokenRouteSymbol)
      )
      .flatten()
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) =>
        values.some((v) => getTokenAndChainId(v).symbol === toTokenRouteSymbol)
      )
      .keys()
      .filter((k) => k.endsWith(`-${fromChainId}`))
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) =>
        values.some((v) => v.startsWith(`${toTokenRouteSymbol}-`))
      )
      .keys()
      .uniq()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .chain()
      .filter((values, _key) => {
        return values.some((v) => {
          const { symbol } = getTokenAndChainId(v)
          return symbol === toTokenRouteSymbol
        })
      })
      .flatten()
      .uniq()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) => {
        return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
      })
      .keys()
      .value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((_values, key) => key.endsWith(`-${fromChainId}`))
      .pickBy((values, _key) => {
        return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
      })
      .keys()
      .value()
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol &&
    toChainId &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) => {
        return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
      })
      .keys()
      .value()
  }

  if (fromChainId && fromTokenRouteSymbol && toChainId && toTokenRouteSymbol) {
    return _(EXISTING_SWAP_ROUTES)
      .pickBy((values, _key) =>
        values.some((v) => {
          return v === `${toTokenRouteSymbol}-${toChainId}`
        })
      )
      .keys()
      .filter((key) => key.endsWith(`-${fromChainId}`))
      .value()
  }
}
