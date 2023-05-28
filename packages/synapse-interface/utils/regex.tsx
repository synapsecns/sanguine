/*
  Helper Function: checkStringIfOnlyZeroes
  - regex function to determine if user input is only zeroes
  */
export function checkStringIfOnlyZeroes(str: string): boolean {
  const regex = /^0*\.?0*$|^$/
  return regex.test(str)
}
