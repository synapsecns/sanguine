import { getAddress, Address, InvalidAddressError } from 'viem'

export const isValidAddress = (address: string): boolean => {
  if (address && address !== '') {
    try {
      const validatedAddress: string = getAddress(address)
      return true
    } catch (e: InvalidAddressError | any) {
      console.error('isValidAddress error: ', e)
      return false
    }
  } else {
    return false
  }
}

export const getValidAddress = (address: string): Address | any => {
  if (address && address !== '') {
    try {
      const validatedAddress: Address = getAddress(address)
      return validatedAddress
    } catch (e: InvalidAddressError | any) {
      console.error('getValidAddress error: ', e)
      return null
    }
  } else {
    return null
  }
}
