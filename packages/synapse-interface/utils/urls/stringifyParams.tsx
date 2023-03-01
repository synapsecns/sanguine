
/**
 * @param {Object} obj
 * @return {string}
 */
export function stringifyParams(obj) {
  return new URLSearchParams(obj).toString()
}
