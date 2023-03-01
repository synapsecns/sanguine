
/**
 * @param {string} rawVal
 * checks if input string can be converted into a valid number
 */
export function checkValidNumberInput(rawVal) {
  return ((rawVal === "") || rawVal.match(/^([0-9]*(\.)?)[0-9]*$/g) )
}
