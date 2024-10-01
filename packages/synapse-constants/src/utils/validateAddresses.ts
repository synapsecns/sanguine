import { getAddress } from '@ethersproject/address'

export const validateAddresses = (addresses: {
  [x: number]: string
}): { [x: number]: string } => {
  const reformatted: { [x: number]: string } = {}
  for (const chainId in addresses) {
    reformatted[chainId] = addresses[chainId]
      ? getAddress(addresses[chainId])
      : ''
  }
  return reformatted
}
