import { Address } from 'viem'

import { ALL_TOKENS } from '@/constants/tokens/master'

export const findTokenByAddressAndChain = (
  address: Address | string,
  chainId: number | string
) => {
  for (const [, token] of Object.entries(ALL_TOKENS)) {
    const chainAddresses = token.addresses
    if (
      chainAddresses &&
      Object.keys(chainAddresses).length > 0 &&
      chainAddresses[chainId] &&
      chainAddresses[chainId].toLowerCase() === address.toLowerCase()
    ) {
      return token
    }
  }
  return null
}
