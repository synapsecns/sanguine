import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'

import { Ticker } from './ticker'

export type FastBridgeQuote = {
  ticker: Ticker
  destAmount: BigNumber
  maxOriginAmount: BigNumber
  fixedFee: BigNumber
  originFastBridge: string
  destFastBridge: string
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
  OriginFastBridgeAddress: string
  DestFastBridgeAddress: string
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
    originFastBridge: quote.OriginFastBridgeAddress,
    destFastBridge: quote.DestFastBridgeAddress,
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
    OriginFastBridgeAddress: quote.originFastBridge,
    DestFastBridgeAddress: quote.destFastBridge,
    RelayerAddr: quote.relayerAddr,
    UpdatedAt: new Date(quote.updatedAt).toISOString(),
  }
}

export const applyQuote = (
  quote: FastBridgeQuote,
  originAmount: BigNumber
): BigNumber => {
  // Check that the origin amount covers the fixed fee
  if (originAmount.lte(quote.fixedFee)) {
    return Zero
  }
  // Check that the Relayer is able to process the origin amount (post fixed fee)
  const amountAfterFee = originAmount.sub(quote.fixedFee)
  if (amountAfterFee.gt(quote.maxOriginAmount)) {
    return Zero
  }
  // After these checks: 0 < amountAfterFee <= quote.maxOriginAmount
  // Solve (amountAfterFee -> ?) using (maxOriginAmount -> destAmount) pricing ratio
  return amountAfterFee.mul(quote.destAmount).div(quote.maxOriginAmount)
}
