/**
 * removes everything except digits and decimals
 * swaps first `.` for an `x`, removes rest of decimals
 * swaps back `x` for `.`
 */
export const cleanNumberInput = (value: string): string => {
  return value === ''
    ? ''
    : value.replace(/[^0-9.]/g, '').replace(/(\..*)\./g, '$1') ?? ''
}
