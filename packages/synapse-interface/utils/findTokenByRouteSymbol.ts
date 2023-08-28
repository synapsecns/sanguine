import _ from 'lodash'
import * as ALL_TOKENS from '@constants/tokens/bridgeable'

export const findTokenByRouteSymbol = (routeSymbol: string) => {
  return _.find(ALL_TOKENS, (token) => token.routeSymbol === routeSymbol)
}
