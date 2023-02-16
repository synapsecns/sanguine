import numeral from 'numeral'

export function formatAmount(value) {
  numeral.nullFormat('--')
  numeral.zeroFormat('--')

  return numeral(value).format(`0,0.00`)
}
