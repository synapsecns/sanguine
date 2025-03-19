import { hexlify } from '@ethersproject/bytes'
import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import {
  decodeZapData,
  encodeZapData,
  StepParams,
  PartialZapDataV1,
} from '../core'

export const getLastStepZapData = (steps: StepParams[]): PartialZapDataV1 => {
  if (steps.length === 0) {
    throw new Error('No steps provided')
  }
  return decodeZapData(hexlify(steps[steps.length - 1].zapData))
}

export const setLastStepZapData = (
  steps: StepParams[],
  zapData: PartialZapDataV1
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
