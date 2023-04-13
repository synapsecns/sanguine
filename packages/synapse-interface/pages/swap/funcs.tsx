import { BigNumber } from '@ethersproject/bignumber'
import { formatUnits } from '@ethersproject/units'
import { Zero, Two } from '@ethersproject/constants'

const ERROR_THRESHOLD = 1/5000 // 10000

const MAX_LOOP_COUNT = 20

export async function estimateAmountToGive({ targetAmountToRecieve, toCoin, fromCoin, chainId, ...props}) {
  let fromCoinDecimalAdjustment = BigNumber.from(10).pow(18 - fromCoin.decimals[chainId])
  let toCoinDecimalAdjustment = BigNumber.from(10).pow(18 - toCoin.decimals[chainId])

  targetAmountToRecieve = targetAmountToRecieve.mul(toCoinDecimalAdjustment)

  const sharedCalcArgs = {
    ...props,
    fromCoinDecimalAdjustment,
    toCoinDecimalAdjustment,
  }

  let currentAmtToRecieve = await calcDecimalAdjustedAmountToRecieve({
    ...sharedCalcArgs,
    amountToGive: targetAmountToRecieve.div(fromCoinDecimalAdjustment)
  })


  let newAmountToGive = targetAmountToRecieve

  if (targetAmountToRecieve.isZero()) {
    newAmountToGive = Zero
  } else {
    let error = targetAmountToRecieve.sub(currentAmtToRecieve)

    let errorFrac = getErrorFrac({ error, targetAmountToRecieve })
    let loopCount = 0
    while ((errorFrac > ERROR_THRESHOLD) && (loopCount < MAX_LOOP_COUNT)) {
      error = targetAmountToRecieve.sub(currentAmtToRecieve)
      errorFrac = getErrorFrac({ error, targetAmountToRecieve })
      const errorCorrection = error.div(Two)
      newAmountToGive = newAmountToGive.add(errorCorrection)

      currentAmtToRecieve = await calcDecimalAdjustedAmountToRecieve({
        ...sharedCalcArgs,
        amountToGive: newAmountToGive
      })

      loopCount += 1
    }
  }

  return newAmountToGive.div(fromCoinDecimalAdjustment)
}


function getErrorFrac({ error, targetAmountToRecieve}) {
  return (
    Math.abs(Number(formatUnits(error)) / Number(formatUnits(targetAmountToRecieve)))
  )
}


async function calcDecimalAdjustedAmountToRecieve({ amountToGive, fromCoinDecimalAdjustment, toCoinDecimalAdjustment, ...props}) {
  let currentAmtToRecieve = await calcAmountToRecieve({
    ...props,
    amountToGive: amountToGive.div(fromCoinDecimalAdjustment)
  })

  return currentAmtToRecieve.mul(toCoinDecimalAdjustment)

}



export async function calcAmountToRecieve({ swapContract, tokenIndexFrom, tokenIndexTo, amountToGive }) {

  let amountToReceive
  const swapArgs = [tokenIndexFrom, tokenIndexTo, amountToGive]
  if (amountToGive.isZero()) {
    amountToReceive = Zero
  } else {
    amountToReceive = await swapContract.calculateSwap(...swapArgs)
  }
  return amountToReceive
}



/**
 * @param {BigNumber} amountFrom
 * @param {number} tokenPrecisionFrom
 * @param {BigNumber} amountTo
 * @param {number} tokenPrecisionTo
 */
export function calculateExchangeRate(
  amountFrom,
  tokenPrecisionFrom,
  amountTo,
  tokenPrecisionTo
) {
  if (amountFrom.gt('0')) {
    return amountTo
      .mul(BigNumber.from(10).pow(36 - tokenPrecisionTo)) // convert to standard 1e18 precision
      .div(amountFrom.mul(BigNumber.from(10).pow(18 - tokenPrecisionFrom)))
  } else {
    return Zero
  }
}

