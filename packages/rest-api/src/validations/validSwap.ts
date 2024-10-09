import { tokenAddressToToken } from '../utils/tokenAddressToToken'

export const validSwap = (
  chain: number | string,
  fromToken: string,
  toToken: string
) => {
  const fromTokenInfo = tokenAddressToToken(chain.toString(), fromToken)
  const toTokenInfo = tokenAddressToToken(chain.toString(), toToken)

  if (!fromTokenInfo || !toTokenInfo) {
    return false
  }

  return fromTokenInfo.swappable.includes(toToken)
}
