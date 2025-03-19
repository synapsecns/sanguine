import { WeiPerEther } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

export type Slippage = {
  numerator: number
  denominator: number
}

export const SlippageZero: Slippage = {
  numerator: 0,
  denominator: 1,
}

// Max slippage that can be used by the swap engines, 100 bips (1%)
// TODO: consider lowering this value
export const SlippageMax: Slippage = {
  numerator: 100,
  denominator: 10000,
}

export const toBasisPoints = (slippage: Slippage): number => {
  return Math.round((slippage.numerator * 10000) / slippage.denominator)
}

export const toPercentFloat = (slippage: Slippage): number => {
  return (slippage.numerator * 100) / slippage.denominator
}

export const toFloat = (slippage: Slippage): number => {
  return slippage.numerator / slippage.denominator
}

export const toWei = (slippage: Slippage): BigNumber => {
  return BigNumber.from(slippage.numerator)
    .mul(WeiPerEther)
    .div(slippage.denominator)
}

export const applySlippage = (
  amount: BigNumber,
  slippage: Slippage
): BigNumber => {
  return amount.sub(amount.mul(slippage.numerator).div(slippage.denominator))
}
