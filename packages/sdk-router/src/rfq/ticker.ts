import { getAddress } from '@ethersproject/address'

export type ChainToken = {
  chainId: number
  token: string
}

export type Ticker = {
  originToken: ChainToken
  destToken: ChainToken
}

/**
 * Marshalls a ChainToken object into a string. Follows the format of "chainId:token".
 *
 * @param chainToken - The ChainToken object to marshall.
 * @returns The marshalled string.
 */
export const marshallChainToken = (chainToken: ChainToken): string => {
  return `${chainToken.chainId}:${chainToken.token}`
}

/**
 * Unmarshalls a string into a ChainToken object. Follows the format of "chainId:token".
 *
 * @param chainTokenStr - The string to unmarshall.
 * @returns The unmarshalled ChainToken object.
 * @throws Will throw an error if the string is not in the correct format.
 */
export const unmarshallChainToken = (chainTokenStr: string): ChainToken => {
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

/**
 * Marshalls a Ticker object into a string. Follows the format of "originChainId:originToken-destChainId:destToken".
 *
 * @param ticker - The Ticker object to marshall.
 * @returns The marshalled string.
 */
export const marshallTicker = (ticker: Ticker): string => {
  return `${marshallChainToken(ticker.originToken)}-${marshallChainToken(
    ticker.destToken
  )}`
}

/**
 * Unmarshalls a string into a Ticker object. Follows the format of "originChainId:originToken-destChainId:destToken".
 *
 * @param tickerStr - The string to unmarshall.
 * @returns The unmarshalled Ticker object.
 * @throws Will throw an error if the string is not in the correct format.
 */
export const unmarshallTicker = (tickerStr: string): Ticker => {
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
