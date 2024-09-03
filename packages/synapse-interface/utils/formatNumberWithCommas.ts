export const formatNumberWithCommas = (value: string): string => {
  if (!value) return ''

  const [integerPart, decimalPart] = value.split('.')

  const formattedInteger = integerPart.replace(/\B(?=(\d{3})+(?!\d))/g, ',')

  return decimalPart !== undefined
    ? `${formattedInteger}.${decimalPart}`
    : formattedInteger
}
