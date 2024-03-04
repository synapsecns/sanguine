import _ from 'lodash'

import type { RouteQueryFields } from '@/utils/route/types'
import { getTokenAndChainId } from '@utils/route/getTokenAndChainId'

export const getFromTokensFunc = (routes) => {
  return ({
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
      return _(routes).keys().uniq().value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
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
      return _(routes).keys().uniq().value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
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
      return _(routes)
        .pickBy((values, _key) => {
          return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
        })
        .keys()
        .value()
    }

    if (fromChainId && fromTokenRouteSymbol && toChainId && toTokenRouteSymbol) {
      return _(routes)
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
}

