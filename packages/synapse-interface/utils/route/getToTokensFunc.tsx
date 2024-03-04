import _ from 'lodash'

import type { RouteQueryFields } from '@/utils/route/types'

export const getToTokensFunc = (routes) => {
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
      return _(routes).values().flatten().uniq().value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .pickBy((_values, key) => key.endsWith(`-${fromChainId}`))
        .values()
        .flatten()
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .pickBy((_values, key) => key.startsWith(`${fromTokenRouteSymbol}-`))
        .values()
        .flatten()
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return routes[`${fromTokenRouteSymbol}-${fromChainId}`]
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol === null &&
      toChainId &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .values()
        .flatten()
        .filter((token) => token.endsWith(`-${toChainId}`))
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
        .pickBy((_values, key) => key.endsWith(`-${fromChainId}`))
        .values()
        .flatten()
        .filter((value) => value.endsWith(`-${toChainId}`))
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
        .pickBy((_values, key) => key.startsWith(`${fromTokenRouteSymbol}-`))
        .values()
        .flatten()
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
      return routes[
        `${fromTokenRouteSymbol}-${fromChainId}`
      ]?.filter((token) => token.endsWith(`-${toChainId}`))
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return _(routes).values().flatten().uniq().value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .pickBy((_values, key) => key.endsWith(`-${fromChainId}`))
        .values()
        .flatten()
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .pickBy((_values, key) => key.startsWith(`${fromTokenRouteSymbol}-`))
        .values()
        .flatten()
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return routes[`${fromTokenRouteSymbol}-${fromChainId}`]
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol === null &&
      toChainId &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .mapValues((values) =>
          values.filter((token) => token === `${toTokenRouteSymbol}-${toChainId}`)
        )
        .values()
        .flatten()
        .uniq()
        .value()
    }
    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .mapValues((values) =>
          values.filter((token) => token === `${toTokenRouteSymbol}-${toChainId}`)
        )
        .pickBy((_values, key) => key.endsWith(`-${fromChainId}`))
        .values()
        .flatten()
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol &&
      toChainId &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .mapValues((values) =>
          values.filter((token) => token === `${toTokenRouteSymbol}-${toChainId}`)
        )
        .pickBy((_values, key) => key.startsWith(`${fromTokenRouteSymbol}-`))
        .values()
        .flatten()
        .uniq()
        .value()
    }

    if (fromChainId && fromTokenRouteSymbol && toChainId && toTokenRouteSymbol) {
      return _.uniq(
        routes[`${fromTokenRouteSymbol}-${fromChainId}`]?.filter(
          (value) => value.endsWith(`-${toChainId}`)
        )
      )
    }
  }

}

