import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { Ticker } from './ticker'

export type FastBridgeQuote = {
  ticker: Ticker
  destAmount: BigNumber
  maxOriginAmount: BigNumber
  fixedFee: BigNumber
  relayerAddr: string
  updatedAt: number
}

export type FastBridgeQuoteAPI = {
  OriginChainID: number
  OriginTokenAddr: string
  DestChainID: number
  DestTokenAddr: string
  DestAmount: string
  MaxOriginAmount: string
  FixedFee: string
  RelayerAddr: string
  UpdatedAt: string
}

export const unmarshallFastBridgeQuote = (
  quote: FastBridgeQuoteAPI
): FastBridgeQuote => {
  return {
    ticker: {
      originToken: {
        chainId: quote.OriginChainID,
        token: quote.OriginTokenAddr,
      },
      destToken: {
        chainId: quote.DestChainID,
        token: quote.DestTokenAddr,
      },
    },
    destAmount: BigNumber.from(quote.DestAmount),
    maxOriginAmount: BigNumber.from(quote.MaxOriginAmount),
    fixedFee: BigNumber.from(quote.FixedFee),
    relayerAddr: quote.RelayerAddr,
    updatedAt: Date.parse(quote.UpdatedAt),
  }
}

export const marshallFastBridgeQuote = (
  quote: FastBridgeQuote
): FastBridgeQuoteAPI => {
  return {
    OriginChainID: quote.ticker.originToken.chainId,
    OriginTokenAddr: quote.ticker.originToken.token,
    DestChainID: quote.ticker.destToken.chainId,
    DestTokenAddr: quote.ticker.destToken.token,
    DestAmount: quote.destAmount.toString(),
    MaxOriginAmount: quote.maxOriginAmount.toString(),
    FixedFee: quote.fixedFee.toString(),
    RelayerAddr: quote.relayerAddr,
    UpdatedAt: new Date(quote.updatedAt).toISOString(),
  }
}

export const applyQuote = (
  quote: FastBridgeQuote,
  originAmount: BigNumber
): BigNumber => {
  // Check that the Relayer is able to process the origin amount
  if (originAmount.gt(quote.maxOriginAmount)) {
    return Zero
  }
  // Check that the origin amount covers the fixed fee
  if (originAmount.lte(quote.fixedFee)) {
    return Zero
  }
  // TODO: adjust this once the concept of price is introduced
  const quotedAmount = originAmount.sub(quote.fixedFee)
  // Check that the quoted amount is less than available destination amount
  return quotedAmount.lte(quote.destAmount) ? quotedAmount : Zero
}
