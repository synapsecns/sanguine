import { tokenAddressToToken } from '../utils/tokenAddressToToken'

export const validSwap = (chain, fromToken, toToken) => {
  const fromTokenInfo = tokenAddressToToken(chain.toString(), fromToken)
  const toTokenInfo = tokenAddressToToken(chain.toString(), toToken)

  if (!fromTokenInfo || !toTokenInfo) {
    return false
  }

  return fromTokenInfo.swappable.includes(toToken)
}
