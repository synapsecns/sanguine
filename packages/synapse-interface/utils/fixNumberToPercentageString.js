export function fixNumberToPercentageString(num, numDecimals=2) {
  return `${num?.toFixed(numDecimals)}%`
}
