
import { Contract } from '@ethersproject/contracts'


import { isAddress } from '@utils/isAddress'


/**
 * account is not optional
 * @param library
 * @param {string} account
 */
export function getSigner(library, account) {
  return library.getSigner(account).connectUnchecked()
}

/**
 * @param library
 * @param {string?} account
 */
export function getProviderOrSigner(library, account) {
  return account ? getSigner(library, account) : library
}

/**
 * @param {string} address
 * @param abi
 * @param library
 * @param {string?} account
 */
export function getContract(address, abi, library, account) {
  if (!isAddress(address)) {
    throw Error(`Invalid 'address' parameter '${address}'.`)
  }

  return new Contract(address, abi, getProviderOrSigner(library, account))
}
