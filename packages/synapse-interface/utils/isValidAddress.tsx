import { getAddress, Address, InvalidAddressError } from 'viem'

export const isValidAddress = (address: string): boolean => {
  try {
    const validatedAddress: string = getAddress(address)
    return true
  } catch (e: InvalidAddressError | any) {
    return false
  }
}

export const getValidAddress = (address: string): Address | any => {
  try {
    const validatedAddress: Address = getAddress(address)
    return validatedAddress
  } catch (e: InvalidAddressError | any) {
    return null
  }
}
