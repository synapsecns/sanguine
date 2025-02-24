import { Chain } from '../../types'

export const ETH: Chain = {
  priorityRank: 100,
  id: 1,
  chainSymbol: 'ETH',
  name: 'Ethereum',
  codeName: 'ethereum',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/ethereum.4a372106.svg',
  layer: 1,
  rpcUrls: {
    primary: 'https://rpc.ankr.com/eth',
    fallback: 'https://1rpc.io/eth',
  },
  explorerUrl: 'https://etherscan.io',
  explorerName: 'Etherscan',
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/ethereum.4a372106.svg',
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
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/arbitrum.8ddb1b22.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/arbitrum.8ddb1b22.svg',
  color: 'gray',
}

export const BNB: Chain = {
  priorityRank: 90,
  id: 56,
  chainSymbol: 'BNB',
  name: 'BNB Chain',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/bnb.5948fe5e.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/bscscan.a50e7cfb.svg',
  color: 'yellow',
}

export const AVALANCHE: Chain = {
  priorityRank: 90,
  id: 43114,
  chainSymbol: 'AVALANCHE',
  name: 'Avalanche',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/avalanche.9d53cbf0.svg',
  layer: 1,
  codeName: 'avalanche',
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://api.avax.network/ext/bc/C/rpc',
    fallback: 'https://1rpc.io/avax/c',
  },
  nativeCurrency: { name: 'Avax', symbol: 'AVAX', decimals: 18 },
  explorerUrl: 'https://snowscan.xyz/',
  explorerName: 'SnowScan',
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/snowscan.1d03dfbf.svg',
  color: 'red',
}

export const CANTO: Chain = {
  priorityRank: 70,
  id: 7700,
  chainSymbol: 'CANTO',
  name: 'Canto',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/canto.cb85e14f.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/snowscan.1d03dfbf.svg',
  color: 'green',
}

export const OPTIMISM: Chain = {
  priorityRank: 80,
  id: 10,
  chainSymbol: 'OPTIMISM',
  name: 'Optimism',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/optimism.84d4f0ef.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/optimism.84d4f0ef.svg',
  color: 'red',
}

export const POLYGON: Chain = {
  priorityRank: 80,
  id: 137,
  chainSymbol: 'POLYGON',
  name: 'Polygon',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/polygon.237cd2b6.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/polygon.237cd2b6.svg',
  color: 'purple',
}

export const DFK: Chain = {
  priorityRank: 75,
  id: 53935,
  chainSymbol: 'DFK',
  name: 'DFK Chain',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/dfk.2bd1f0e4.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/dfk.2bd1f0e4.svg',
  color: 'lime',
}

export const KLAYTN: Chain = {
  priorityRank: 70,
  id: 8217,
  chainSymbol: 'KLAYTN',
  name: 'Klaytn',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/klaytn.59495fbb.svg',
  layer: 1,
  codeName: 'klaytn',
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://kaia.blockpi.network/v1/rpc/public',
    fallback: 'https://klaytn.api.onfinality.io/public',
  },
  nativeCurrency: { name: 'Klaytn', symbol: 'KLAY', decimals: 18 },
  explorerUrl: 'https://scope.klaytn.com',
  explorerName: 'Klaytn Explorer',
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/klaytn.59495fbb.svg',
  color: 'orange',
}

export const FANTOM: Chain = {
  priorityRank: 70,
  id: 250,
  chainSymbol: 'FANTOM',
  name: 'Fantom',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/fantom.1e444dad.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/fantom.1e444dad.svg',
  color: 'blue',
}

export const CRONOS: Chain = {
  priorityRank: 10,
  id: 25,
  chainSymbol: 'CRONOS',
  name: 'Cronos',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/cronos.b06f8311.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/cronos.b06f8311.svg',
  color: 'gray',
}

export const BOBA: Chain = {
  priorityRank: 10,
  id: 288,
  chainSymbol: 'BOBA',
  name: 'Boba Chain',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/boba.2072e50b.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/boba.2072e50b.svg',
  color: 'lime',
}

export const METIS: Chain = {
  priorityRank: 10,
  id: 1088,
  chainSymbol: 'METIS',
  name: 'Metis',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/metis.90b6abf0.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/metis.90b6abf0.svg',
  color: 'teal',
}

export const AURORA: Chain = {
  priorityRank: 10,
  id: 1313161554,
  chainSymbol: 'AURORA',
  name: 'Aurora',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/aurora.5a46037d.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/aurora.5a46037d.svg',
  color: 'lime',
}

export const HARMONY: Chain = {
  priorityRank: 10,
  id: 1666600000,
  chainSymbol: 'HARMONY',
  name: 'Harmony',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/harmony.af12f77e.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/harmony.af12f77e.svg',
  color: 'cyan',
}

export const MOONBEAM: Chain = {
  priorityRank: 0,
  id: 1284,
  chainSymbol: 'MOONBEAM',
  name: 'Moonbeam',
  chainImg:
    'https://adf1cceb.sanguine-fe.pages.dev/_next/static/media/moonbeam.63f34507.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/moonbeam.284ab9b4.svg',
  color: 'purple',
}

export const MOONRIVER: Chain = {
  priorityRank: 0,
  id: 1285,
  chainSymbol: 'MOONRIVER',
  name: 'Moonriver',
  chainImg:
    'https://adf1cceb.sanguine-fe.pages.dev/_next/static/media/moonriver.275d190a.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/moonriver.3fb35010.svg',
  color: 'lime',
}

export const DOGE: Chain = {
  priorityRank: 0,
  id: 2000,
  chainSymbol: 'DOGE',
  name: 'Dogechain',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/dogechain.36935650.svg',
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
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/dogechain.36935650.svg',
  color: 'purple',
}

export const BASE: Chain = {
  priorityRank: 90,
  id: 8453,
  chainSymbol: 'ETH',
  name: 'Base',
  codeName: 'base',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/base.d919fbef.svg',
  layer: 2,
  rpcUrls: {
    primary: 'https://base.blockpi.network/v1/rpc/public',
    fallback: 'https://developer-access-mainnet.base.org',
  },
  explorerUrl: 'https://basescan.org',
  explorerName: 'BaseScan',
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/base.d919fbef.svg',
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ether',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'blue',
}

export const BLAST: Chain = {
  priorityRank: 90,
  id: 81457,
  chainSymbol: 'ETH',
  name: 'Blast',
  codeName: 'blast',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/blast.e39807f8.svg',
  layer: 2,
  rpcUrls: {
    primary: 'https://rpc.blast.io',
    fallback: 'https://blast.blockpi.network/v1/rpc/publicChain',
  },
  explorerUrl: 'https://blastscan.io',
  explorerName: 'Blastscan',
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/blast.e39807f8.svg',
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ether',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'yellow',
}

export const SCROLL: Chain = {
  priorityRank: 90,
  id: 534352,
  chainSymbol: 'SCROLL',
  name: 'Scroll',
  codeName: 'scroll',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/scroll.a805c122.svg',
  layer: 2,
  rpcUrls: {
    primary: 'https://rpc.scroll.io',
    fallback: 'https://scroll.blockpi.network/v1/rpc/public',
  },
  explorerUrl: 'https://scrollscan.com/',
  explorerName: 'Scrollscan',
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/scroll.a805c122.svg',
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ether',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'orange',
}

export const LINEA: Chain = {
  priorityRank: 90,
  id: 59144,
  chainSymbol: 'LINEA',
  name: 'Linea',
  codeName: 'linea',
  chainImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/linea.e476f2ad.svg',
  layer: 2,
  rpcUrls: {
    primary: 'https://linea.blockpi.network/v1/rpc/public',
    fallback: 'https://rpc.linea.build',
  },
  explorerUrl: 'https://lineascan.build/',
  explorerName: 'LineaScan',
  explorerImg:
    'https://8f3ea9f2.sanguine-fe.pages.dev/_next/static/media/scroll.a805c122.svg',
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ether',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'black',
}

export const WORLDCHAIN: Chain = {
  priorityRank: 99,
  id: 480,
  chainSymbol: 'WORLDCHAIN',
  name: 'World Chain',
  codeName: 'worldchain',
  chainImg:
    'https://synapse-interface-worldchain.sanguine-fe.pages.dev/_next/static/media/worldchain.62d1dfd2.svg',
  layer: 2,
  rpcUrls: {
    primary: 'https://worldchain-mainnet.g.alchemy.com/public',
    fallback: 'https://worldchain-mainnet.g.alchemy.com/public',
  },
  explorerUrl: 'https://worldchain-mainnet.explorer.alchemy.com',
  explorerName: 'World Chain Explorer',
  explorerImg:
    'https://synapse-interface-worldchain.sanguine-fe.pages.dev/_next/static/media/worldchain.62d1dfd2.svg',
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
  color: 'black',
}

export const BERACHAIN: Chain = {
  priorityRank: 99,
  id: 80094,
  chainSymbol: 'BERACHAIN',
  name: 'Berachain',
  codeName: 'berachain',
  chainImg:
    'https://e8ff9599.sanguine-fe.pages.dev/_next/static/media/berachain.57304c86.svg',
  layer: 1,
  rpcUrls: {
    primary: 'https://berachain.blockpi.network/v1/rpc/public',
    fallback: 'https://rpc.berachain.com',
  },
  explorerUrl: 'https://berascan.com',
  explorerName: 'Berascan',
  explorerImg:
    'https://e8ff9599.sanguine-fe.pages.dev/_next/static/media/berachain.57304c86.svg',
  blockTime: 3000,
  nativeCurrency: { name: 'Berachain', symbol: 'BERA', decimals: 18 },
  color: 'brown',
}

export const UNICHAIN: Chain = {
  priorityRank: 99,
  id: 130,
  chainSymbol: 'UNICHAIN',
  name: 'Unichain',
  codeName: 'unichain',
  chainImg:
    'https://e8ff9599.sanguine-fe.pages.dev/_next/static/media/unichain.02b81e3e.svg',
  layer: 2,
  rpcUrls: {
    primary: 'https://unichain-rpc.publicnode.com',
    fallback: 'https://rpc.unichain.io',
  },
  explorerUrl: 'https://uniscan.xyz',
  explorerName: 'Uniscan',
  explorerImg:
    'https://e8ff9599.sanguine-fe.pages.dev/_next/static/media/unichain.02b81e3e.svg',
  blockTime: 3000,
  nativeCurrency: { name: 'Ethereum', symbol: 'ETH', decimals: 18 },
  color: 'pink',
}
