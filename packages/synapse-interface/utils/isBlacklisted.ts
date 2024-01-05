import { Address, getAddress } from 'viem'

import { EXCLUDED_ADDRESSES } from '@/constants/blacklist'

export const isBlacklisted = (address: Address | string) => {
  try {
    const checksummedAddress = getAddress(address)

    const normalizedExcludedAddresses = EXCLUDED_ADDRESSES.map((addr) =>
      getAddress(addr)
    )

    return normalizedExcludedAddresses.includes(checksummedAddress)
  } catch (error) {
    console.error('Invalid address:', error)
    return false
  }
}
