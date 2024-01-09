import arbitrumImg from '../assets/chains/arbitrum.svg'
import auroraImg from '../assets/chains/aurora.svg'
import avalancheImg from '../assets/chains/avalanche.svg'
import baseImg from '../assets/chains/base.svg'
import bobaImg from '../assets/chains/boba.svg'
import bscImg from '../assets/chains/bnb.svg'
import cantoImg from '../assets/chains/canto.svg'
import cronosImg from '../assets/chains/cronos.svg'
import dfkImg from '../assets/chains/dfk.svg'
import dogechainImg from '../assets/chains/dogechain.svg'
import ethImg from '../assets/chains/ethereum.svg'
import fantomImg from '../assets/chains/fantom.svg'
import harmonyImg from '../assets/chains/harmony.svg'
import klaytnImg from '../assets/chains/klaytn.svg'
import metisImg from '../assets/chains/metis.svg'
import moonbeamImg from '../assets/chains/moonbeam.svg'
import moonriverImg from '../assets/chains/moonriver.svg'
import optimismImg from '../assets/chains/optimism.svg'
import polygonImg from '../assets/chains/polygon.svg'
import ethExplorerImg from '../assets/explorer/etherscan.svg'
import arbitrumExplorerImg from '../assets/explorer/arbitrum.svg'
import bnbExplorerImg from '../assets/explorer/bscscan.svg'
import avalancheExplorerImg from '../assets/explorer/avalanche.svg'
import cantoExplorerImg from '../assets/explorer/canto.svg'
import optimismExplorerImg from '../assets/explorer/optimism.svg'
import polygonExplorerImg from '../assets/explorer/polygon.svg'
import dfkExplorerImg from '../assets/explorer/dfk-chain.svg'
import klaytynExplorerImg from '../assets/explorer/klaytn.svg'
import fantomExplorerImg from '../assets/explorer/fantom.svg'
import cronosExplorerImg from '../assets/explorer/cronos.svg'
import bobaExplorerImg from '../assets/explorer/boba.svg'
import metisExplorerImg from '../assets/explorer/metis.svg'
import auroraExplorerImg from '../assets/explorer/aurora.svg'
import harmonyExplorerImg from '../assets/explorer/harmony.svg'
import moonbeamExplorerImg from '../assets/explorer/moonbeam.svg'
import moonriverExplorerImg from '../assets/explorer/moonriver.svg'
import dogeExplorerImg from '../assets/explorer/dogecoin.svg'
import baseExplorerImg from '../assets/explorer/basescan.svg'
import { Chain } from '../types'

export const ETH: Chain = {
  priorityRank: 100,
  id: 1,
  chainSymbol: 'ETH',
  name: 'Ethereum',
  codeName: 'ethereum',
  chainImg: ethImg,
  layer: 1,
  rpcUrls: {
    primary: 'https://rpc.ankr.com/eth',
    fallback: 'https://1rpc.io/eth',
  },
  explorerUrl: 'https://etherscan.com',
  explorerName: 'Etherscan',
  explorerImg: ethExplorerImg,
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
    primary: 'https://arbitrum.blockpi.network/v1/rpc/public',
    fallback: 'https://arb1.arbitrum.io/rpc',
  },
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://arbiscan.io',
  explorerName: 'Arbiscan',
  explorerImg: arbitrumExplorerImg,
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
  explorerImg: bnbExplorerImg,
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
    fallback: 'https://1rpc.io/avax/c',
  },
  nativeCurrency: { name: 'Avax', symbol: 'AVAX', decimals: 18 },
  explorerUrl: 'https://avascan.info/blockchain/c/',
  explorerName: 'Avascan',
  explorerImg: avalancheExplorerImg,
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
  explorerImg: cantoExplorerImg,
  color: 'green',
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
    fallback: 'https://1rpc.io/op',
  },
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://optimistic.etherscan.io',
  explorerName: 'Optimism Explorer',
  explorerImg: optimismExplorerImg,
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
  explorerImg: polygonExplorerImg,
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
  explorerImg: dfkExplorerImg,
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
  explorerImg: klaytynExplorerImg,
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
  explorerImg: fantomExplorerImg,
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
  explorerImg: cronosExplorerImg,
  color: 'gray',
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
  explorerImg: bobaExplorerImg,
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
  explorerImg: metisExplorerImg,
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
  explorerImg: auroraExplorerImg,
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
  explorerImg: harmonyExplorerImg,
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
  explorerImg: moonbeamExplorerImg,
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
  explorerImg: moonriverExplorerImg,
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
    primary: 'https://rpc.dogechain.dog',
    fallback: 'https://rpc01-sg.dogechain.dog',
  },
  nativeCurrency: { name: 'DOGE', symbol: 'DOGE', decimals: 18 },
  explorerUrl: 'https://explorer.dogechain.dog',
  explorerName: 'Dogechain Explorer',
  explorerImg: dogeExplorerImg,
  color: 'purple',
}

export const BASE: Chain = {
  priorityRank: 90,
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
  explorerImg: baseExplorerImg,
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ether',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'blue',
}
