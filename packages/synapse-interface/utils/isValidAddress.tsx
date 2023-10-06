import { getAddress } from '@ethersproject/address'

export const isValidAddress = (address: string): boolean => {
  try {
    const validatedAddress: string = getAddress(address)
    return true
  } catch (e: any) {
    return false
  }
}
