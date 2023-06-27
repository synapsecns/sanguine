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

export const CCTP_ROUTER_ADDRESS: { [chainId: number]: string } = {
  [SupportedChainId.ETH]: '0xd359bc471554504f683fbd4f6e36848612349ddf',
  [SupportedChainId.ARBITRUM]: '0xd359bc471554504f683fbd4f6e36848612349ddf',
  [SupportedChainId.AVALANCHE]: '0xd359bc471554504f683fbd4f6e36848612349ddf',
}

export const ROUTER_ADDRESS: { [chainId: number]: string } = {
  [SupportedChainId.BSC]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.ETH]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.POLYGON]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.BOBA]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.MOONBEAM]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.MOONRIVER]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.ARBITRUM]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.OPTIMISM]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.AVALANCHE]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.DFK]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.FANTOM]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.HARMONY]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.AURORA]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.CRONOS]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.METIS]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.KLAYTN]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.DOGECHAIN]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
  [SupportedChainId.CANTO]: '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a',
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
