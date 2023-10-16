import { getAddress } from '@ethersproject/address'
import { getAddress as getAddressViem, Address } from 'viem'

export const isValidAddress = (address: string): boolean => {
  try {
    const validatedAddress: string = getAddress(address)
    return true
  } catch (e: any) {
    return false
  }
}

export const getValidAddress = (address: string): Address | any => {
  try {
    const validatedAddress: Address = getAddressViem(address)
    return validatedAddress
  } catch (e: any) {
    console.error('getValidAddress error: ', e)
    return null
  }
}
