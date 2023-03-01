// CONVENTION isFoo -> boolean

/**
 * Returns true if the string value is zero in hex
 * @param hexNumberString
 */
export function isZero(hexNumberString) {
  return /^0x0*$/.test(hexNumberString)
}