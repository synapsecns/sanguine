import { getAddress } from '@ethersproject/address'

export type ChainToken = {
  chainId: number
  token: string
}

export type Ticker = {
  originToken: ChainToken
  destToken: ChainToken
}

export const marshallChainToken = (chainToken: ChainToken): string => {
  return `${chainToken.chainId}:${chainToken.token}`
}

export const unmarshallChainToken = (chainTokenStr: string): ChainToken => {
  // The unmashalled string should be in the format of "chainId:token"
  const items = chainTokenStr.split(':')
  if (items.length !== 2) {
    throw new Error(`Can not unmarshall "${chainTokenStr}": invalid format`)
  }
  // Check if the chain ID is a number
  const chainId = Number(items[0])
  if (isNaN(chainId)) {
    throw new Error(
      `Can not unmarshall "${chainTokenStr}": ${items[0]} is not a chain ID`
    )
  }
  const token = getAddress(items[1])
  return {
    chainId,
    token,
  }
}

export const marshallTicker = (ticker: Ticker): string => {
  return `${marshallChainToken(ticker.originToken)}-${marshallChainToken(
    ticker.destToken
  )}`
}

export const unmarshallTicker = (tickerStr: string): Ticker => {
  // The unmashalled string should be in the format of "originChainId:originToken-destChainId:destToken"
  const items = tickerStr.split('-')
  if (items.length !== 2) {
    throw new Error(`Can not unmarshall "${tickerStr}": invalid format`)
  }
  const originToken = unmarshallChainToken(items[0])
  const destToken = unmarshallChainToken(items[1])
  return {
    originToken,
    destToken,
  }
}
