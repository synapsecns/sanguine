import JSBI from 'jsbi'
import { BigNumber } from '@ethersproject/bignumber'

export enum SupportedChainId {
  ETH = 1,
  OPTIMISM = 10,
  CRONOS = 25,
  BSC = 56,
  POLYGON = 137,
  FANTOM = 250,
  BOBA = 288,
  METIS = 1088,
  MOONBEAM = 1284,
  MOONRIVER = 1285,
  DOGECHAIN = 2000,
  CANTO = 7700,
  KLAYTN = 8217,
  ARBITRUM = 42161,
  AVALANCHE = 43114,
  DFK = 53935,
  AURORA = 1313161554,
  HARMONY = 1666600000,
}

export const ROUTER_ADDRESS: object = {
  [SupportedChainId.BSC]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.ETH]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.POLYGON]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.BOBA]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.MOONBEAM]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.MOONRIVER]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.ARBITRUM]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.OPTIMISM]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.AVALANCHE]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.DFK]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.FANTOM]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.HARMONY]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.AURORA]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.CRONOS]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.METIS]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.KLAYTN]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.DOGECHAIN]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
  [SupportedChainId.CANTO]: '0x25f8fA4917180FF308883e37Ea27CEaEB68c1f19',
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
