import bscImg from '@assets/chains/bnb.svg'
import { zeroAddress } from 'viem'

import { Token } from '@/utils/types'
import * as CHAINS from '@/constants/chains/master'

export type GasToken = {
  addresses: { [x: number]: string }
  chainId: number
  decimals: number
  symbol: string
  name: string
  icon: any
}

// export const BNB = new Token({
//   addresses: {
//     [CHAINS.BNB.id]: zeroAddress,
//   },
//   decimals: 18,
//   symbol: 'BNB',
//   name: 'Binance Coin',
//   priorityRank: 1,
//   logo: bscImg,
// })

export const BNB_Gas: GasToken = {
  addresses: {
    [CHAINS.BNB.id]: zeroAddress,
  },
  chainId: 56,
  decimals: 18,
  symbol: 'BNB',
  name: 'Binance Coin',
  icon: bscImg,
}

// export const METIS: GasToken = {
//   chainId: 1088,
//   chainName: 'Metis',
//   name: 'Metis',
//   symbol: 'METIS',
//   decimals: 18,
// }

// export const NOTE: GasToken = {
//   chainId: 7700,
//   chainName: 'Canto',
//   name: 'Canto',
//   symbol: 'NOTE',
//   decimals: 18,
// }

// export const MOVR: GasToken = {
//   chainId: 1285,
//   chainName: 'Moonriver',
//   name: 'Moonriver',
//   symbol: 'MOVR',
//   decimals: 18,
// }

// export const GLMR: GasToken = {
//   chainId: 1284,
//   chainName: 'Moonbeam',
//   name: 'Glimmer',
//   symbol: 'GLMR',
//   decimals: 18,
// }

// export const ONE: GasToken = {
//   chainId: 1666600000,
//   chainName: 'Harmony',
//   name: 'Harmony One',
//   symbol: 'ONE',
//   decimals: 18,
// }
