import { Address, zeroAddress } from 'viem'

import { ALL_TOKENS } from '@/constants/tokens/master'
import { ETHEREUM_ADDRESS } from '@/constants'

export const findTokenByAddressAndChain = (
  address: Address | string,
  chainId: number | string
) => {
  const searchAddress =
    address.toLowerCase() === ETHEREUM_ADDRESS.toLowerCase()
      ? zeroAddress
      : address

  for (const [, token] of Object.entries(ALL_TOKENS)) {
    const chainAddresses = token.addresses
    if (
      chainAddresses &&
      Object.keys(chainAddresses).length > 0 &&
      chainAddresses[chainId] &&
      chainAddresses[chainId].toLowerCase() === searchAddress.toLowerCase()
    ) {
      return token
    }
  }
  return null
}
