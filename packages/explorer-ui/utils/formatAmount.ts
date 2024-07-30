import numeral from 'numeral'

export const formatAmount = (value) => {
  numeral.nullFormat('--')

  // Round up if the value is less than 0.001
  if (value > 0 && value < 0.001) {
    value = 0.001
  }

  let format
  if (value >= 1000) {
    format = '0,0' // No decimal places for values 1000 or greater
  } else if (value < 0.01) {
    format = '0,0.000' // Four decimal places for values less than 0.01
  } else {
    format = '0,0.00' // Two decimal places for all other values
  }

  return numeral(value).format(format)
}
