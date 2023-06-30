import bscLogo from '@assets/icons/binance.svg'
import ethLogo from '@assets/icons/eth.svg'
import polygonLogo from '@assets/icons/polygon.svg'
import fantomLogo from '@assets/icons/fantom.svg'
import arbitrumLogo from '@assets/icons/arbitrum.svg'
import avalancheLogo from '@assets/icons/avalanche.svg'
import auroraLogo from '@assets/icons/aurora.svg'
import harmonyLogo from '@assets/icons/harmonyone.svg'
import optimismLogo from '@assets/icons/optimism.svg'
import bobaLogo from '@assets/icons/boba.svg'
import cronosLogo from '@assets/icons/cronos.svg'
import metisLogo from '@assets/icons/metis.svg'

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

import { Chain } from '@types'

export const ETH: Chain = {
  priorityRank: 100,
  id: 1,
  chainSymbol: 'ETH',
  name: 'Ethereum',
  codeName: 'Optimism',
  chainLogo: ethLogo,
  chainImg: ethImg,
  layer: 1,
  rpc: 'https://rpc.ankr.com/eth',
  writeRpc: 'https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161',
  explorerUrl: 'https://etherscan.com',
  blockTime: 10000,
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
  chainLogo: arbitrumLogo,
  chainImg: arbitrumImg,
  layer: 2,
  codeName: 'arbitrum',
  blockTime: 5000,
  rpc: 'https://arbitrum.blockpi.network/v1/rpc/public',
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://arbiscan.io',
  color: 'gray',
}
export const BNB: Chain = {
  priorityRank: 90,
  id: 56,
  chainSymbol: 'BNB',
  name: 'BNB Chain',
  chainImg: bscImg,
  chainLogo: bscLogo,
  altName: 'BNB',
  layer: 1,
  codeName: 'bsc',
  blockTime: 10000,
  rpc: 'https://bsc-dataseed1.ninicoin.io/',
  nativeCurrency: { name: 'Binance Coin', symbol: 'BNB', decimals: 18 },
  explorerUrl: 'https://bscscan.com',
  color: 'yellow',
}
export const AVALANCHE: Chain = {
  priorityRank: 90,
  id: 43114,
  chainSymbol: 'AVALANCHE',
  name: 'Avalanche',
  chainLogo: avalancheLogo,
  chainImg: avalancheImg,
  layer: 1,
  codeName: 'avalanche',
  blockTime: 5000,
  rpc: 'https://api.avax.network/ext/bc/C/rpc',
  nativeCurrency: { name: 'Avax', symbol: 'AVAX', decimals: 18 },
  explorerUrl: 'https://snowtrace.io',
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
  blockTime: 50000,
  rpc: 'https://mainnode.plexnode.org:8545',
  nativeCurrency: { name: 'Canto', symbol: 'CANTO', decimals: 18 },
  explorerUrl: 'https://evm.explorer.canto.io',
  color: 'teal',
}

export const OPTIMISM: Chain = {
  priorityRank: 80,
  id: 10,
  chainSymbol: 'OPTIMISM',
  name: 'Optimism',
  chainLogo: optimismLogo,
  chainImg: optimismImg,
  layer: 2,
  codeName: 'optimism',
  blockTime: 10000,
  rpc: 'https://rpc.ankr.com/optimism',
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://optimistic.etherscan.io',
  color: 'red',
}

export const POLYGON: Chain = {
  priorityRank: 80,
  id: 137,
  chainSymbol: 'POLYGON',
  name: 'Polygon',
  chainLogo: polygonLogo,
  chainImg: polygonImg,
  layer: 2,
  codeName: 'polygon',
  blockTime: 10000,
  rpc: 'https://polygon-bor.publicnode.com',
  nativeCurrency: { name: 'Matic', symbol: 'MATIC', decimals: 18 },
  explorerUrl: 'https://polygonscan.com',
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
  blockTime: 10000,
  rpc: 'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
  nativeCurrency: { name: 'Jewel', symbol: 'JEWEL', decimals: 18 },
  explorerUrl: 'https://subnets.avax.network/defi-kingdoms',
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
  blockTime: 10000,
  rpc: 'https://klaytn.blockpi.network/v1/rpc/public',
  nativeCurrency: { name: 'Klaytn', symbol: 'KLAY', decimals: 18 },
  explorerUrl: 'https://scope.klaytn.com',
  color: 'orange',
}
export const FANTOM: Chain = {
  priorityRank: 70,
  id: 250,
  chainSymbol: 'FANTOM',
  name: 'Fantom',
  chainLogo: fantomLogo,
  chainImg: fantomImg,
  layer: 1,
  codeName: 'fantom',
  blockTime: 5000,
  rpc: 'https://rpc.ftm.tools',
  nativeCurrency: { name: 'Fantom', symbol: 'FTM', decimals: 18 },
  explorerUrl: 'https://ftmscan.com',
  color: 'blue',
}
export const CRONOS: Chain = {
  priorityRank: 10,
  id: 25,
  chainSymbol: 'CRONOS',
  name: 'Cronos',
  chainLogo: cronosLogo,
  chainImg: cronosImg,
  layer: 1,
  codeName: 'cronos',
  blockTime: 10000,
  rpc: 'https://evm-cronos.crypto.org',
  nativeCurrency: { name: 'Cronos', symbol: 'CRO', decimals: 18 },
  explorerUrl: 'https://cronoscan.com',
  color: 'blue',
}
export const BOBA: Chain = {
  priorityRank: 10,
  id: 288,
  chainSymbol: 'BOBA',
  name: 'Boba Chain',
  chainLogo: bobaLogo,
  chainImg: bobaImg,
  layer: 2,
  codeName: 'boba',
  blockTime: 20000,
  rpc: 'https://mainnet.boba.network',
  writeRpc: 'https://mainnet.boba.network',
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://blockexplorer.boba.network',
  color: 'lime',
}
export const METIS: Chain = {
  priorityRank: 10,
  id: 1088,
  chainSymbol: 'METIS',
  name: 'Metis',
  chainLogo: metisLogo,
  chainImg: metisImg,
  layer: 2,
  codeName: 'metis',
  blockTime: 10000,
  rpc: 'https://andromeda.metis.io/?owner=1088',
  nativeCurrency: { name: 'Metis', symbol: 'METIS', decimals: 18 },
  explorerUrl: 'https://andromeda-explorer.metis.io',
  color: 'teal',
}

export const AURORA: Chain = {
  priorityRank: 10,
  id: 1313161554,
  chainSymbol: 'AURORA',
  name: 'Aurora',
  chainLogo: auroraLogo,
  chainImg: auroraImg,
  layer: 1,
  codeName: 'aurora',
  blockTime: 10000,
  rpc: 'https://mainnet.aurora.dev',
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  explorerUrl: 'https://explorer.mainnet.aurora.dev',
  color: 'lime',
}
export const HARMONY: Chain = {
  priorityRank: 10,
  id: 1666600000,
  chainSymbol: 'HARMONY',
  name: 'Harmony',
  chainLogo: harmonyLogo,
  chainImg: harmonyImg,
  layer: 1,
  codeName: 'harmony',
  blockTime: 10000,
  rpc: 'https://api.harmony.one',
  nativeCurrency: { name: 'Harmony One', symbol: 'ONE', decimals: 18 },
  explorerUrl: 'https://explorer.harmony.one',
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
  blockTime: 10000,
  rpc: 'https://rpc.api.moonbeam.network',
  nativeCurrency: { name: 'Glimmer', symbol: 'GLMR', decimals: 18 },
  explorerUrl: 'https://moonbeam.moonscan.io',
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
  blockTime: 5000,
  rpc: 'https://rpc.api.moonriver.moonbeam.network',
  nativeCurrency: { name: 'Moonriver', symbol: 'MOVR', decimals: 18 },
  explorerUrl: 'https://moonriver.moonscan.io',
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
  blockTime: 10000,
  rpc: 'https://1rpc.io/avax/c',
  nativeCurrency: { name: 'DOGE', symbol: 'DOGE', decimals: 18 },
  explorerUrl: 'https://explorer.dogechain.dog',
  color: 'purple',
}
