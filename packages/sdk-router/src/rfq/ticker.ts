import { getAddress } from '@ethersproject/address'

export type ChainToken = {
  chainId: number
  token: string
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
