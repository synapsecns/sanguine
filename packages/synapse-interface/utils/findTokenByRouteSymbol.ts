import _ from 'lodash'
import * as ALL_TOKENS from '@constants/tokens/master'

export const findTokenByRouteSymbol = (routeSymbol: string) => {
  return _.find(ALL_TOKENS, (token) => token.routeSymbol === routeSymbol)
}
