import { BigNumber } from 'ethers'
import { Zero } from '@ethersproject/constants'
import { hexlify } from '@ethersproject/bytes'

import { decodeZapData, encodeZapData, StepParams, ZapDataV1 } from '../core'

export const getLastStepZapData = (steps: StepParams[]): Partial<ZapDataV1> => {
  if (steps.length === 0) {
    throw new Error('No steps provided')
  }
  return decodeZapData(hexlify(steps[steps.length - 1].zapData))
}

export const setLastStepZapData = (
  steps: StepParams[],
  zapData: Partial<ZapDataV1>
): StepParams[] => {
  if (steps.length === 0) {
    throw new Error('No steps provided')
  }
  steps[steps.length - 1].zapData = encodeZapData(zapData)
  return steps
}

export const getMinFinalAmount = (steps: StepParams[]): BigNumber => {
  return getLastStepZapData(steps).minFinalAmount ?? Zero
}

export const setMinFinalAmount = (
  steps: StepParams[],
  minFinalAmount: BigNumber
): StepParams[] => {
  return setLastStepZapData(steps, {
    ...getLastStepZapData(steps),
    minFinalAmount,
  })
}
