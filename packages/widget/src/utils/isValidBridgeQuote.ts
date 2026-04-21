import { parseBigIntValue } from '@/utils/parseBigIntValue'

type BridgeQuoteLike = {
  nativeFee?: unknown
}

export const isValidBridgeQuote = <T extends BridgeQuoteLike>(quote: T) => {
  const nativeFeeWei = parseBigIntValue(quote.nativeFee)

  return nativeFeeWei !== null && nativeFeeWei >= 0n
}
