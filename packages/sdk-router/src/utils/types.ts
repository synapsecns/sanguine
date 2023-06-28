import { BigNumber } from '@ethersproject/bignumber'

export type RawQuery = [string, string, BigNumber, BigNumber, string] & {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

export interface SynapseRouterQuery {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

export interface SynapseCCTPRouterQuery {
  routerAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

export type Query = SynapseRouterQuery | SynapseCCTPRouterQuery

// export type Query = {
//   swapAdapter: string
//   tokenOut: string
//   minAmountOut: BigNumber
//   deadline: BigNumber
//   rawParams: string
// }

export type RawFeeConfig = [number, BigNumber, BigNumber] & {
  bridgeFee: number
  minFee: BigNumber
  maxFee: BigNumber
}

export type FeeConfig = {
  bridgeFee: number
  minFee: BigNumber
  maxFee: BigNumber
}

export type PoolToken = { isWeth: boolean | undefined; token: string }

export const convertQuery = (rawQuery: RawQuery): Query => {
  const { swapAdapter, tokenOut, minAmountOut, deadline, rawParams } = rawQuery
  return { swapAdapter, tokenOut, minAmountOut, deadline, rawParams }
}

export const convertFeeConfig = (rawFeeConfig: RawFeeConfig): FeeConfig => {
  const { bridgeFee, minFee, maxFee } = rawFeeConfig
  return { bridgeFee, minFee, maxFee }
}
