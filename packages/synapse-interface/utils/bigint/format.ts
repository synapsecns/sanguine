import { txErrorHandler } from '../txErrorHandler'

export const formatBigIntToString = (
  bi: bigint,
  nativePrecision: number,
  decimalPlaces?: number
) => {
  if (typeof bi !== 'bigint' && !bi) {
    return
  }
  try {
    // Check if input is zero
    if (bi === 0n) {
      return '0.0'
    }

    // Check if the input is negative
    const isNegative = bi < 0n
    if (isNegative) {
      // Convert to positive for the calculation
      bi = -bi
    }
    // Convert to string and add padding zeros if necessary
    let str = bi.toString().padStart(nativePrecision, '0')

    // Insert decimal point
    const idx = str.length - nativePrecision
    str = `${str.slice(0, idx)}.${str.slice(idx)}`

    // Handle values below zero by adding a '0' before the decimal point
    if (str.startsWith('.')) {
      str = '0' + str
    }

    // Trim to desired number of decimal places
    if (decimalPlaces !== undefined) {
      const decimalIdx = str.indexOf('.')
      str = str.slice(0, decimalIdx + decimalPlaces + 1)
    }

    // Add the negative sign back if necessary
    if (isNegative) {
      str = '-' + str
    }

    return str
  } catch (error) {
    console.log(`error`, error)
    txErrorHandler(error)
  }
}

export const powBigInt = (base, exponent) => {
  let result = 1n
  for (let i = 0; i < exponent; i++) {
    result *= base
  }
  return result
}

export const formatBigIntToPercentString = (
  bn: bigint,
  nativePrecison: number,
  decimalPlaces = 2,
  convert = true
) => {
  try {
    // Calculate the conversion factor based on the native precision and required decimal places
    const conversionFactor = powBigInt(
      10n,
      BigInt(nativePrecison - 2 + decimalPlaces)
    )

    // Convert the bigint to a floating-point number, preserving the requested number of decimal places
    const percentConvert = convert ? 100 : 1
    const num = (Number(bn) * percentConvert) / Number(conversionFactor)

    // Format the number as a percentage string
    return `${num.toFixed(decimalPlaces)}%`
  } catch (error) {
    console.log(`error`, error)
    txErrorHandler(error)
  }
}

// Some environments have issues with RegEx that contain back-tracking, so we cannot
// use them.
export const commify = (value: string | number): string => {
  const comps = String(value).split('.')

  if (
    comps.length > 2 ||
    !comps[0].match(/^-?[0-9]*$/) ||
    (comps[1] && !comps[1].match(/^[0-9]*$/)) ||
    value === '.' ||
    value === '-.'
  ) {
    console.log('invalid value', 'value', value)
  }

  // Make sure we have at least one whole digit (0 if none)
  let whole = comps[0]

  let negative = ''
  if (whole.substring(0, 1) === '-') {
    negative = '-'
    whole = whole.substring(1)
  }

  // Make sure we have at least 1 whole digit with no leading zeros
  while (whole.substring(0, 1) === '0') {
    whole = whole.substring(1)
  }
  if (whole === '') {
    whole = '0'
  }

  let suffix = ''
  if (comps.length === 2) {
    suffix = '.' + (comps[1] || '0')
  }
  while (suffix.length > 2 && suffix[suffix.length - 1] === '0') {
    suffix = suffix.substring(0, suffix.length - 1)
  }

  const formatted = []
  while (whole.length) {
    if (whole.length <= 3) {
      formatted.unshift(whole)
      break
    } else {
      const index = whole.length - 3
      formatted.unshift(whole.substring(index))
      whole = whole.substring(0, index)
    }
  }

  return negative + formatted.join(',') + suffix
}

export const commifyBigIntToString = (
  big: bigint,
  precision: number,
  decimals = 2
) => {
  return commify(formatBigIntToString(big, precision, decimals))
}

export const commifyBigIntWithDefault = (big: bigint, decimals: number) => {
  return big ? commifyBigIntToString(big, decimals) : '0'
}

export const stringToBigInt = (rawVal: string, rawDecimals: number) => {
  if (typeof rawVal !== 'string' && !rawVal) {
    return 0n
  }

  try {
    let value = rawVal.replace(/[$,]/g, '')
    if (['.', '0.', '', '.0'].includes(value)) {
      return 0n
    }
    if (value.startsWith('.')) {
      value = `0${value}`
    }

    // Scale the decimal number up by the appropriate number of decimals.
    const scaleFactor = BigInt(10 ** rawDecimals)

    // Split the input into whole and fractional parts.
    const [wholePart, fractionalPart = '0'] = value.split('.')

    // Convert the whole part directly to a BigInt.
    const wholeBigInt = BigInt(wholePart) * scaleFactor

    // For the fractional part, first scale it up to the right size, then trim any excess decimal places.
    const fractionalBigInt = BigInt(
      fractionalPart.padEnd(rawDecimals, '0').slice(0, rawDecimals)
    )

    return wholeBigInt + fractionalBigInt
  } catch (error) {
    // console.log(error)
    txErrorHandler(error)
  }
}
