import JSBI from 'jsbi'
import { BigNumber } from '@ethersproject/bignumber'

export * from './addresses'
export * from './chainIds'
export * from './medianTime'
export * from './testValues'

// exports for external consumption
export type BigintIsh = JSBI | BigNumber | string | number

export enum TradeType {
  EXACT_INPUT,
  EXACT_OUTPUT,
}

export enum Rounding {
  ROUND_DOWN,
  ROUND_HALF_UP,
  ROUND_UP,
}

export const MaxUint256 = JSBI.BigInt(
  '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
)
