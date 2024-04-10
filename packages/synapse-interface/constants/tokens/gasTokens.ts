import { zeroAddress } from 'viem'

import * as CHAINS from '@/constants/chains/master'

export type GasToken = {
  addresses: { [x: number]: string }
  chainId: number
  decimals: number
  symbol: string
  name: string
  icon: any
}

// export const BNB: GasToken = {
//   addresses: {
//     [CHAINS.BNB.id]: zeroAddress,
//   },
//   chainId: CHAINS.BNB.id,
//   decimals: CHAINS.BNB.nativeCurrency.decimals,
//   symbol: CHAINS.BNB.nativeCurrency.symbol,
//   name: CHAINS.BNB.nativeCurrency.name,
//   icon: CHAINS.BNB.chainImg,
// }

// export const METIS: GasToken = {
//   addresses: {
//     [CHAINS.METIS.id]: zeroAddress,
//   },
//   chainId: CHAINS.METIS.id,
//   name: CHAINS.METIS.nativeCurrency.name,
//   symbol: CHAINS.METIS.nativeCurrency.symbol,
//   decimals: CHAINS.METIS.nativeCurrency.decimals,
//   icon: CHAINS.METIS.chainImg,
// }

// export const CANTO: GasToken = {
//   addresses: {
//     [CHAINS.CANTO.id]: zeroAddress,
//   },
//   chainId: CHAINS.CANTO.id,
//   name: CHAINS.CANTO.nativeCurrency.name,
//   symbol: CHAINS.CANTO.nativeCurrency.symbol,
//   decimals: CHAINS.CANTO.nativeCurrency.decimals,
//   icon: CHAINS.CANTO.chainImg,
// }

// export const MOVR: GasToken = {
//   addresses: {
//     [CHAINS.MOONRIVER.id]: zeroAddress,
//   },
//   chainId: CHAINS.MOONRIVER.id,
//   name: CHAINS.MOONRIVER.nativeCurrency.name,
//   symbol: CHAINS.MOONRIVER.nativeCurrency.symbol,
//   decimals: CHAINS.MOONRIVER.nativeCurrency.decimals,
//   icon: CHAINS.MOONRIVER.chainImg,
// }

// export const GLMR: GasToken = {
//   addresses: {
//     [CHAINS.MOONBEAM.id]: zeroAddress,
//   },
//   chainId: CHAINS.MOONBEAM.id,
//   name: CHAINS.MOONBEAM.nativeCurrency.name,
//   symbol: CHAINS.MOONBEAM.nativeCurrency.symbol,
//   decimals: CHAINS.MOONBEAM.nativeCurrency.decimals,
//   icon: CHAINS.MOONBEAM.chainImg,
// }

// export const ONE: GasToken = {
//   addresses: {
//     [CHAINS.HARMONY.id]: zeroAddress,
//   },
//   chainId: CHAINS.HARMONY.id,
//   name: CHAINS.HARMONY.nativeCurrency.name,
//   symbol: CHAINS.HARMONY.nativeCurrency.symbol,
//   decimals: CHAINS.HARMONY.nativeCurrency.decimals,
//   icon: CHAINS.HARMONY.chainImg,
// }

// export const CRO: GasToken = {
//   addresses: {
//     [CHAINS.CRONOS.id]: zeroAddress,
//   },
//   chainId: CHAINS.CRONOS.id,
//   name: CHAINS.CRONOS.nativeCurrency.name,
//   symbol: CHAINS.CRONOS.nativeCurrency.symbol,
//   decimals: CHAINS.CRONOS.nativeCurrency.decimals,
//   icon: CHAINS.CRONOS.chainImg,
// }

// export const DOGE: GasToken = {
//   addresses: {
//     [CHAINS.DOGE.id]: zeroAddress,
//   },
//   chainId: CHAINS.DOGE.id,
//   name: CHAINS.DOGE.nativeCurrency.name,
//   symbol: CHAINS.DOGE.nativeCurrency.symbol,
//   decimals: CHAINS.DOGE.nativeCurrency.decimals,
//   icon: CHAINS.DOGE.chainImg,
// }
