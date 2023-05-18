import _ from 'lodash'

/**
 * gives undefined if hash invalid/ not 64 len
 *
 * @param {string} hash
 * */
export function validateAndParseHash(hash) {
  try {
    if (/^0x([A-Fa-f0-9]{64})$/.test(hash)) {
      return _.toLower(hash)
    }
  } catch (error) {
    console.error(error)
    console.error(`${hash} is not a valid txnhash/hash.`)
  }
}
