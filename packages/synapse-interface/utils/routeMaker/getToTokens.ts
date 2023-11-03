import _ from 'lodash'

import { EXISTING_BRIDGE_ROUTES } from '@/constants/existingBridgeRoutes'
import { RouteQueryFields } from './generateRoutePossibilities'

export const getToTokens = ({
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
    return _(EXISTING_BRIDGE_ROUTES).values().flatten().uniq().value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
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
    return _(EXISTING_BRIDGE_ROUTES)
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
    return EXISTING_BRIDGE_ROUTES[`${fromTokenRouteSymbol}-${fromChainId}`]
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol === null
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
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
    return _(EXISTING_BRIDGE_ROUTES)
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
    return _(EXISTING_BRIDGE_ROUTES)
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
    return EXISTING_BRIDGE_ROUTES[
      `${fromTokenRouteSymbol}-${fromChainId}`
    ]?.filter((token) => token.endsWith(`-${toChainId}`))
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES).values().flatten().uniq().value()
  }

  if (
    fromChainId &&
    fromTokenRouteSymbol === null &&
    toChainId === null &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
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
    return _(EXISTING_BRIDGE_ROUTES)
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
    return EXISTING_BRIDGE_ROUTES[`${fromTokenRouteSymbol}-${fromChainId}`]
  }

  if (
    fromChainId === null &&
    fromTokenRouteSymbol === null &&
    toChainId &&
    toTokenRouteSymbol
  ) {
    return _(EXISTING_BRIDGE_ROUTES)
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
    return _(EXISTING_BRIDGE_ROUTES)
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
    return _(EXISTING_BRIDGE_ROUTES)
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
      EXISTING_BRIDGE_ROUTES[`${fromTokenRouteSymbol}-${fromChainId}`]?.filter(
        (value) => value.endsWith(`-${toChainId}`)
      )
    )
  }
}
