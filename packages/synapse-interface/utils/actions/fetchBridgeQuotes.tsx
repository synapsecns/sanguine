import { BridgeQuote, Token } from '@/utils/types'

export interface BridgeQuoteResponse extends BridgeQuote {
  destinationToken: Token
  destinationChainId: number
}
