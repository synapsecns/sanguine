import ethImg from '@assets/chains/eth.jpg'
import bscImg from '@assets/chains/bsc.jpg'
import polygonImg from '@assets/chains/polygon.jpg'
import fantomImg from '@assets/chains/fantom.jpg'
import arbitrumImg from '@assets/chains/arbitrum.jpg'
import avalancheImg from '@assets/chains/avalanche.jpg'
import dfkImg from '@assets/chains/dfk.png'
import auroraImg from '@assets/chains/aurora.png'
import harmonyImg from '@assets/chains/harmonyone.jpg'
import optimismImg from '@assets/chains/optimism.png'
import bobaImg from '@assets/chains/boba.png'
import moonbeamImg from '@assets/chains/moonbeam.jpg'
import moonriverImg from '@assets/chains/moonriver.jpeg'
import cronosImg from '@assets/chains/cronos.png'
import metisImg from '@assets/chains/metis.png'
import klaytnImg from '@assets/chains/klaytn.jpeg'
import dogechainImg from '@assets/chains/dogechain.png'
import cantoImg from '@assets/chains/canto.svg'
import baseImg from '@assets/chains/base.svg'

import { Chain } from '@types'

export const ETH: Chain = {
  priorityRank: 100,
  id: 1,
  chainSymbol: 'ETH',
  name: 'Ethereum',
  codeName: 'Optimism',
  chainImg: ethImg,
  layer: 1,
  rpcUrls: {
    primary: 'https://rpc.ankr.com/eth',
    fallback: 'https://eth.llamarpc.com',
  },
  explorerUrl: 'https://etherscan.com',
  explorerName: 'Etherscan',
  blockTime: 12000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'eth',
}

export const ARBITRUM: Chain = {
  priorityRank: 90,
  id: 42161,
  chainSymbol: 'ARBITRUM',
  name: 'Arbitrum',
  chainImg: arbitrumImg,
  layer: 2,
  codeName: 'arbitrum',
  blockTime: 300,
  rpcUrls: {
    primary: 'https://arb1.arbitrum.io/rpc',
    fallback: 'https://arbitrum-one.publicnode.com',
  },
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://arbiscan.io',
  explorerName: 'Arbiscan',
  color: 'gray',
}

export const BNB: Chain = {
  priorityRank: 90,
  id: 56,
  chainSymbol: 'BNB',
  name: 'BNB Chain',
  chainImg: bscImg,
  altName: 'BNB',
  layer: 1,
  codeName: 'bsc',
  blockTime: 3000,
  rpcUrls: {
    primary: 'https://bsc-dataseed1.ninicoin.io/',
    fallback: 'https://bsc-dataseed2.ninicoin.io',
  },
  nativeCurrency: { name: 'Binance Coin', symbol: 'BNB', decimals: 18 },
  explorerUrl: 'https://bscscan.com',
  explorerName: 'BscScan',
  color: 'yellow',
}

export const AVALANCHE: Chain = {
  priorityRank: 90,
  id: 43114,
  chainSymbol: 'AVALANCHE',
  name: 'Avalanche',
  chainImg: avalancheImg,
  layer: 1,
  codeName: 'avalanche',
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://api.avax.network/ext/bc/C/rpc',
    fallback: 'https://rpc.ankr.com/avalanche',
  },
  nativeCurrency: { name: 'Avax', symbol: 'AVAX', decimals: 18 },
  explorerUrl: 'https://snowtrace.io',
  explorerName: 'Snowtrace',
  color: 'red',
}

export const CANTO: Chain = {
  priorityRank: 70,
  id: 7700,
  chainSymbol: 'CANTO',
  name: 'Canto',
  chainImg: cantoImg,
  layer: 1,
  codeName: 'canto',
  blockTime: 6000,
  rpcUrls: {
    primary: 'https://mainnode.plexnode.org:8545',
    fallback: 'https://canto.slingshot.finance',
  },
  nativeCurrency: { name: 'Canto', symbol: 'CANTO', decimals: 18 },
  explorerUrl: 'https://tuber.build/',
  explorerName: 'Canto Explorer',
  color: 'teal',
}

export const OPTIMISM: Chain = {
  priorityRank: 80,
  id: 10,
  chainSymbol: 'OPTIMISM',
  name: 'Optimism',
  chainImg: optimismImg,
  layer: 2,
  codeName: 'optimism',
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://mainnet.optimism.io',
    fallback: 'https://rpc.ankr.com/optimism',
  },
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://optimistic.etherscan.io',
  explorerName: 'Optimism Explorer',
  color: 'red',
}

export const POLYGON: Chain = {
  priorityRank: 80,
  id: 137,
  chainSymbol: 'POLYGON',
  name: 'Polygon',
  chainImg: polygonImg,
  layer: 2,
  codeName: 'polygon',
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://polygon-bor.publicnode.com',
    fallback: 'https://polygon.llamarpc.com',
  },
  nativeCurrency: { name: 'Matic', symbol: 'MATIC', decimals: 18 },
  explorerUrl: 'https://polygonscan.com',
  explorerName: 'PolygonScan',
  color: 'purple',
}

export const DFK: Chain = {
  priorityRank: 75,
  id: 53935,
  chainSymbol: 'DFK',
  name: 'DFK Chain',
  chainImg: dfkImg,
  layer: 1,
  codeName: 'dfk',
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
    fallback: 'https://dfkchain.api.onfinality.io/public',
  },
  nativeCurrency: { name: 'Jewel', symbol: 'JEWEL', decimals: 18 },
  explorerUrl: 'https://subnets.avax.network/defi-kingdoms',
  explorerName: 'DFK Subnet Explorer',
  color: 'lime',
}

export const KLAYTN: Chain = {
  priorityRank: 70,
  id: 8217,
  chainSymbol: 'KLAYTN',
  name: 'Klaytn',
  chainImg: klaytnImg,
  layer: 1,
  codeName: 'klaytn',
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://klaytn.blockpi.network/v1/rpc/public',
    fallback: 'https://klaytn.api.onfinality.io/public',
  },
  nativeCurrency: { name: 'Klaytn', symbol: 'KLAY', decimals: 18 },
  explorerUrl: 'https://scope.klaytn.com',
  explorerName: 'Klaytn Explorer',
  color: 'orange',
}

export const FANTOM: Chain = {
  priorityRank: 70,
  id: 250,
  chainSymbol: 'FANTOM',
  name: 'Fantom',
  chainImg: fantomImg,
  layer: 1,
  codeName: 'fantom',
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://rpc.ftm.tools',
    fallback: 'https://fantom-rpc.gateway.pokt.network/',
  },
  nativeCurrency: { name: 'Fantom', symbol: 'FTM', decimals: 18 },
  explorerUrl: 'https://ftmscan.com',
  explorerName: 'FTMScan',
  color: 'blue',
}

export const CRONOS: Chain = {
  priorityRank: 10,
  id: 25,
  chainSymbol: 'CRONOS',
  name: 'Cronos',
  chainImg: cronosImg,
  layer: 1,
  codeName: 'cronos',
  blockTime: 6000,
  rpcUrls: {
    primary: 'https://evm-cronos.crypto.org',
    fallback: 'https://cronos.blockpi.network/v1/rpc/public',
  },
  nativeCurrency: { name: 'Cronos', symbol: 'CRO', decimals: 18 },
  explorerUrl: 'https://cronoscan.com',
  explorerName: 'CronoScan',
  color: 'blue',
}

export const BOBA: Chain = {
  priorityRank: 10,
  id: 288,
  chainSymbol: 'BOBA',
  name: 'Boba Chain',
  chainImg: bobaImg,
  layer: 2,
  codeName: 'boba',
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://mainnet.boba.network',
    fallback: 'https://replica.boba.network',
  },
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://bobascan.com',
  explorerName: 'Boba Explorer',
  color: 'lime',
}

export const METIS: Chain = {
  priorityRank: 10,
  id: 1088,
  chainSymbol: 'METIS',
  name: 'Metis',
  chainImg: metisImg,
  layer: 2,
  codeName: 'metis',
  blockTime: 4000,
  rpcUrls: {
    primary: 'https://andromeda.metis.io/?owner=1088',
    fallback: 'https://metis-mainnet.public.blastapi.io',
  },
  nativeCurrency: { name: 'Metis', symbol: 'METIS', decimals: 18 },
  explorerUrl: 'https://andromeda-explorer.metis.io',
  explorerName: 'Metis Explorer',
  color: 'teal',
}

export const AURORA: Chain = {
  priorityRank: 10,
  id: 1313161554,
  chainSymbol: 'AURORA',
  name: 'Aurora',
  chainImg: auroraImg,
  layer: 1,
  codeName: 'aurora',
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://mainnet.aurora.dev',
    fallback: 'https://1rpc.io/aurora',
  },
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://explorer.mainnet.aurora.dev',
  explorerName: 'Aurora Explorer',
  color: 'lime',
}

export const HARMONY: Chain = {
  priorityRank: 10,
  id: 1666600000,
  chainSymbol: 'HARMONY',
  name: 'Harmony',
  chainImg: harmonyImg,
  layer: 1,
  codeName: 'harmony',
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://api.harmony.one',
    fallback: 'https://api.s0.t.hmny.io',
  },
  nativeCurrency: { name: 'Harmony One', symbol: 'ONE', decimals: 18 },
  explorerUrl: 'https://explorer.harmony.one',
  explorerName: 'Harmony Explorer',
  color: 'cyan',
}

export const MOONBEAM: Chain = {
  priorityRank: 0,
  id: 1284,
  chainSymbol: 'MOONBEAM',
  name: 'Moonbeam',
  chainImg: moonbeamImg,
  layer: 1,
  codeName: 'moonbeam',
  blockTime: 12000,
  rpcUrls: {
    primary: 'https://rpc.api.moonbeam.network',
    fallback: 'https://moonbeam.public.blastapi.io',
  },
  nativeCurrency: { name: 'Glimmer', symbol: 'GLMR', decimals: 18 },
  explorerUrl: 'https://moonbeam.moonscan.io',
  explorerName: 'Moonbeam Explorer',
  color: 'teal',
}

export const MOONRIVER: Chain = {
  priorityRank: 0,
  id: 1285,
  chainSymbol: 'MOONRIVER',
  name: 'Moonriver',
  chainImg: moonriverImg,
  layer: 1,
  codeName: 'moonriver',
  blockTime: 12000,
  rpcUrls: {
    primary: 'https://rpc.api.moonriver.moonbeam.network',
    fallback: 'https://moonriver.public.blastapi.io',
  },
  nativeCurrency: { name: 'Moonriver', symbol: 'MOVR', decimals: 18 },
  explorerUrl: 'https://moonriver.moonscan.io',
  explorerName: 'Moonriver Explorer',
  color: 'purple',
}

export const DOGE: Chain = {
  priorityRank: 0,
  id: 2000,
  chainSymbol: 'DOGE',
  name: 'Dogechain',
  chainImg: dogechainImg,
  layer: 1,
  codeName: 'dogechain',
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://rpc.ankr.com/dogechain',
    fallback: 'https://rpc.dogechain.dog',
  },
  nativeCurrency: { name: 'DOGE', symbol: 'DOGE', decimals: 18 },
  explorerUrl: 'https://explorer.dogechain.dog',
  explorerName: 'Dogechain Explorer',
  color: 'purple',
}

export const BASE: Chain = {
  priorityRank: 1,
  id: 8453,
  chainSymbol: 'ETH',
  name: 'Base',
  codeName: 'base',
  chainImg: baseImg,
  layer: 2,
  rpcUrls: {
    primary: 'https://base.blockpi.network/v1/rpc/public',
    fallback: 'https://developer-access-mainnet.base.org',
  },
  explorerUrl: 'https://basescan.org',
  explorerName: 'BaseScan',
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ether',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'blue',
}
