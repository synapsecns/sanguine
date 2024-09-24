import { tokenAddressToToken } from '../utils/tokenAddressToToken'
import { BRIDGE_ROUTE_MAPPING_SYMBOLS } from '../utils/bridgeRouteMapping'

export const validateRouteExists = (fromChain, fromToken, toChain, toToken) => {
  const fromTokenInfo = tokenAddressToToken(fromChain.toString(), fromToken)
  const toTokenInfo = tokenAddressToToken(toChain.toString(), toToken)

  if (!fromTokenInfo || !toTokenInfo) {
    return false
  }

  const key = `${fromTokenInfo.symbol}-${fromChain}`
  const routes = BRIDGE_ROUTE_MAPPING_SYMBOLS[key]

  if (!routes) {
    return false
  }

  return routes.includes(`${toTokenInfo.symbol}-${toChain}`)
}
