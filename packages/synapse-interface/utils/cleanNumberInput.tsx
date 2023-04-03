/**
 * removes everything except digits and decimals
 * swaps first `.` for an `x`, removes rest of decimals
 * swaps back `x` for `.`
 */
export function cleanNumberInput(value: string) {
  if (value === '') {
    return ''
  } else {
    let val = value.replace(/[^\d.]/g, '')
    return val.replace(/\./, 'x').replace(/\./g, '').replace(/x/, '.')
  }
}
