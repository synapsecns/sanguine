import { defaultAbiCoder } from '@ethersproject/abi'
import { BigNumber } from '@ethersproject/bignumber'
import { hexlify } from '@ethersproject/bytes'

import { ISynapseIntentRouter } from '../typechain/SynapseIntentRouter'

export type StepParams = ISynapseIntentRouter.StepParamsStruct
const stepParamsArray = ['tuple(address,uint256,uint256,bytes)[]']

export const encodeStepParams = (steps: StepParams[]): string => {
  // Unwrap every struct into a tuple
  return defaultAbiCoder.encode(stepParamsArray, [
    steps.map((step) => [step.token, step.amount, step.msgValue, step.zapData]),
  ])
}

export const decodeStepParams = (data: string): StepParams[] => {
  const decoded = defaultAbiCoder.decode(stepParamsArray, data)
  // decoded is [[[token0, amount0, msgValue0, zapData0], [token1, amount1, msgValue1, zapData1], ...]]
  return decoded[0].map(
    ([token, amount, msgValue, zapData]: [
      string,
      BigNumber,
      BigNumber,
      string
    ]) => ({
      token,
      amount,
      msgValue,
      zapData,
    })
  )
}

export const extractSingleZapData = (steps: StepParams[]): string => {
  if (steps.length > 1) {
    throw new Error('extractSingleZapData: more than one step')
  }
  return steps.length === 0 ? '0x' : hexlify(steps[0].zapData)
}
