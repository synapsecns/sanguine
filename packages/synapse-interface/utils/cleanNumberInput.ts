/**
 * removes everything except digits and decimals
 * swaps first `.` for an `x`, removes rest of decimals
 * swaps back `x` for `.`
 */
export const cleanNumberInput = (value: string): string => {
  console.log('cleanNumberInput', value)
  const jo =
    value === '' ? '' : value.replace(/[^0-9.]/g, '').replace(/(\..*)\./g, '$1')
  console.log(jo)
  return jo ?? ''
}
