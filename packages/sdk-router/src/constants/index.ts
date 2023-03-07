import JSBI from 'jsbi'
import { BigNumber } from '@ethersproject/bignumber'

export enum SupportedChainId {
  ARBITRUM = 42161,
  AVALANCHE = 43114,
}

export const ROUTER_ADDRESS: object = {
  [SupportedChainId.ARBITRUM]: '0x59aA81DF6C4d400A68f99d850f69c8A6494eB0E8',
  [SupportedChainId.AVALANCHE]: '0x59aA81DF6C4d400A68f99d850f69c8A6494eB0E8',
}

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
