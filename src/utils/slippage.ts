import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits } from '@ethersproject/units'

export const Slippages = {
  One: 'ONE',
  OneTenth: 'ONE_TENTH',
  TwoTenth: 'TWO_TENTH',
  Quarter: 'QUARTER',
  OnePercent: 'ONE_PERCENT',
  Custom: 'CUSTOM',
  OneHundredth: 'ONE_HUNDREDTH',
}

/**
 * Given an input value and slippage redux state values, do the math.
 *
 * @param {BigNumber} inputValue
 * @param {Slippages} slippageSelected
 * @param {NumberInputState} slippageCustom
 * @param {boolean} add
 * @return {BigNumber}
 */
export const _applySlippage = (
  inputValue,
  slippageSelected,
  slippageCustom,
  add = false
) => {
  let numerator
  let denominator
  if (slippageSelected === Slippages.Custom && !!slippageCustom) {
    denominator = BigNumber.from(10).pow(slippageCustom.precision + 2)
    numerator = add
      ? denominator.add(slippageCustom.valueSafe)
      : denominator.sub(slippageCustom.valueSafe)
  } else if (slippageSelected === Slippages.OneTenth) {
    denominator = 1000
    numerator = denominator + (add ? 1 : -1)
  } else if (slippageSelected === Slippages.OneHundredth) {
    denominator = 10000
    numerator = denominator + (add ? 1 : -1)
  } else if (slippageSelected === Slippages.TwoTenth) {
    denominator = 500
    numerator = denominator + (add ? 1 : -1)
  } else if (slippageSelected === Slippages.Quarter) {
    denominator = 50
    numerator = denominator + (add ? 1 : -1)
  } else if (slippageSelected === Slippages.OnePercent) {
    denominator = 100
    numerator = denominator + (add ? 1 : -1)
  } else {
    // default to 1%
    denominator = 100
    numerator = denominator + (add ? 1 : -1)
  }
  return inputValue.mul(numerator).div(denominator)
}

export const addSlippage = (inputValue, slippageSelected, slippageCustom) => {
  return _applySlippage(inputValue, slippageSelected, slippageCustom, true)
}

export const subtractSlippage = (
  inputValue,
  slippageSelected,
  slippageCustom
) => {
  return _applySlippage(inputValue, slippageSelected, slippageCustom, false)
}

export const formatSlippageToString = (slippageSelected, slippageCustom) => {
  if (slippageSelected === Slippages.Custom && !!slippageCustom) {
    return formatUnits(slippageCustom.valueSafe, slippageCustom?.precision)
  } else if (slippageSelected === Slippages.OneTenth) {
    return formatUnits(BigNumber.from(100), 3)
  } else if (slippageSelected === Slippages.TwoTenth) {
    return formatUnits(BigNumber.from(200), 3)
  } else if (slippageSelected === Slippages.Quarter) {
    return formatUnits(BigNumber.from(2000), 3)
  } else if (slippageSelected === Slippages.One) {
    return formatUnits(BigNumber.from(100), 2)
  } else {
    return 'N/A'
  }
}

// Below functions are dupes of above functions except with input and output values of bigint
// When app completely removes need for BigNumber, we can deprecate the BigNumber versions

/**
 * Given an input value and slippage redux state values, do the math.
 *
 * @param {bigint} inputValue
 * @param {Slippages} slippageSelected
 * @param {NumberInputState} slippageCustom
 * @param {boolean} add
 * @return {bigint}
 */
export const _applySlippageBigInt = (
  inputValue,
  slippageSelected,
  slippageCustom,
  add = false
) => {
  let numerator
  let denominator
  if (slippageSelected === Slippages.Custom && !!slippageCustom) {
    denominator = 10n ** BigInt(slippageCustom.precision + 2)
    numerator = add
      ? denominator + slippageCustom.valueSafe
      : denominator - slippageCustom.valueSafe
  } else if (slippageSelected === Slippages.OneTenth) {
    denominator = 1000n
    numerator = add ? denominator + 1n : denominator - 1n
  } else if (slippageSelected === Slippages.OneHundredth) {
    denominator = 10000n
    numerator = add ? denominator + 1n : denominator - 1n
  } else if (slippageSelected === Slippages.TwoTenth) {
    denominator = 500n
    numerator = add ? denominator + 1n : denominator - 1n
  } else if (slippageSelected === Slippages.Quarter) {
    denominator = 400n
    numerator = add ? denominator + 1n : denominator - 1n
  } else if (slippageSelected === Slippages.OnePercent) {
    denominator = 100n
    numerator = add ? denominator + 1n : denominator - 1n
  } else {
    // default to 1%
    denominator = 100n
    numerator = add ? denominator + 1n : denominator - 1n
  }
  return (inputValue * numerator) / denominator
}

export const subtractSlippageBigInt = (
  inputValue,
  slippageSelected,
  slippageCustom
) => {
  return _applySlippageBigInt(
    inputValue,
    slippageSelected,
    slippageCustom,
    false
  )
}
