import { BigNumber } from "@ethersproject/bignumber";

export type RawQuery = [string, string, BigNumber, BigNumber, string] & {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

export type RawRouterQuery = [string, string, BigNumber, BigNumber, string] & {
  routerAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

export type Query = {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

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

/// TODO: figure out how to fix this struct
export const convertRouterQuery = (rawQuery: RawRouterQuery): Query => {
  const { routerAdapter, tokenOut, minAmountOut, deadline, rawParams } =
    rawQuery
  return {
    swapAdapter: routerAdapter,
    tokenOut,
    minAmountOut,
    deadline,
    rawParams,
  }
}

export const convertFeeConfig = (rawFeeConfig: RawFeeConfig): FeeConfig => {
  const { bridgeFee, minFee, maxFee } = rawFeeConfig
  return { bridgeFee, minFee, maxFee }
}
