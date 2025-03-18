import { BigNumber, BigNumberish } from 'ethers'

import { IntentStep } from '../module'
import { SynapseSDK } from '../sdk'
import { Slippage } from '../swap'

export type IntentParameters = {
  fromChainId: number
  fromToken: string
  fromAmount: BigNumberish
  fromSender?: string
  toChainId: number
  toToken: string
  toRecipient?: string
  slippage?: Slippage
  deadline?: number
  allowMultipleTxs?: boolean
}

export type IntentQuote = {
  id: string
  fromChainId: number
  fromToken: string
  fromAmount: BigNumber
  toChainId: number
  toToken: string
  expectedToAmount: BigNumber
  minToAmount: BigNumber
  estimatedTime: number
  steps: IntentStep[]
}

export async function intent(
  this: SynapseSDK,
  params: IntentParameters
): Promise<IntentQuote[]> {
  // TODO: implement
}
