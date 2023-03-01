import invariant from 'tiny-invariant'
import warning from 'tiny-warning'
import { getAddress } from '@ethersproject/address'

/**
 * warns if addresses are not checksummed dev only, dont show on prod
 * @param {string} address
 * */
export function validateAndParseAddress(address) {
  try {
    const checksummedAddress = getAddress(address)
    warning(address === checksummedAddress, `${address} is not checksummed.`)
    return checksummedAddress
  } catch (error) {
    console.error(error)
    console.error(`${address} is not a valid address.`)

    // invariant(false, `${address} is not a valid address.`)
  }
}
