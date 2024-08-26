import warning from 'tiny-warning'
import { getAddress } from '@ethersproject/address'

/**
 * warns if addresses are not checksummed
 *
 * @param {string} address
 * */
export const validateAndParseAddress = (address) => {
  try {
    const checksummedAddress = getAddress(address)
    warning(address === checksummedAddress, `${address} is not checksummed.`)
    return checksummedAddress
  } catch (error) {
    console.error(error)
    console.error(`${address} is not a valid address.`)
  }
}
