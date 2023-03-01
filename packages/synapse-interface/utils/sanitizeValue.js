/**
 * removes common copy/pasted financial characters
 * @param {string} rawVal
 */
export function sanitizeValue(rawVal) {
  if (rawVal) {
    const val = rawVal.replace(/[$,]/g, '')

    if ([".", "0.", ""].includes(val)) {
      return "0"
    } else {
      return val
    }
  } else {
    return "0"
  }
}
