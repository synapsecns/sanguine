import { getAddress } from '@ethersproject/address'

/**
 * returns the checksummed address if the address is valid, otherwise returns false
 * @param {string} value
 */
export function isAddress(value) {
  try {
    return getAddress(value)
  } catch {
    return false
  }
}