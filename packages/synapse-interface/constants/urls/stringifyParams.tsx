/**
 * @param {Object} obj
 * @return {string}
 */
export function stringifyParams(obj: any) {
  return new URLSearchParams(obj).toString()
}
