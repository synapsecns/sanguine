import bech32 from 'bech32'

/**
 * terra address validation, it verify also the bech32 checksum
 * @param {string} address
 */
export function validateTerraAddress(address) {
  try {
    const { prefix: decodedPrefix } = bech32.decode(address) // throw error if checksum is invalid
    // verify address prefix
    if (decodedPrefix === 'terra') {
      return address
    } else {
      return false
    }
  } catch {
    // invalid checksum
    return false
  }
}