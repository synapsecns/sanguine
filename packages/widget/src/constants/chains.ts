import _ from 'lodash'
import { Chain } from 'types'

export const ETHEREUM: Chain = {
  id: 1,
  name: 'Ethereum',
  rpcUrls: {
    primary: 'https://ethereum.blockpi.network/v1/rpc/public',
    fallback: 'https://rpc.ankr.com/eth',
  },
  explorerUrl: 'https://etherscan.com',
  explorerName: 'Etherscan',
  blockTime: 12000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/ethereum.4a372106.svg',
  networkName: 'Ethereum Mainnet',
  networkUrl: 'https://eth.llamarpc.com',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const ARBITRUM: Chain = {
  id: 42161,
  name: 'Arbitrum',
  rpcUrls: {
    primary: 'https://arbitrum.blockpi.network/v1/rpc/public',
    fallback: 'https://arb1.arbitrum.io/rpc',
  },
  explorerUrl: 'https://arbiscan.io',
  explorerName: 'Arbiscan',
  blockTime: 300,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/arbitrum.8ddb1b22.svg',
  networkName: 'Arbitrum One',
  networkUrl: 'https://arb1.arbitrum.io/rpc',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const BNBCHAIN: Chain = {
  id: 56,
  name: 'BNB Chain',
  rpcUrls: {
    primary: 'https://bsc-dataseed1.ninicoin.io/',
    fallback: 'https://bsc-dataseed2.ninicoin.io',
  },
  explorerUrl: 'https://bscscan.com',
  explorerName: 'BscScan',
  blockTime: 3000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/bnb.5948fe5e.svg',
  networkName: 'BNB Smart Chain Mainnet',
  networkUrl: 'https://bsc-dataseed1.bnbchain.org',
  nativeCurrency: {
    name: 'Binance Coin',
    symbol: 'BNB',
    decimals: 18,
  },
}

export const AVALANCHE: Chain = {
  id: 43114,
  name: 'Avalanche',
  rpcUrls: {
    primary: 'https://api.avax.network/ext/bc/C/rpc',
    fallback: 'https://1rpc.io/avax/c',
  },
  explorerUrl: 'https://snowtrace.io/',
  explorerName: 'Snowtrace',
  blockTime: 2000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/avalanche.9d53cbf0.svg',
  networkName: 'Avalanche C-Chain',
  networkUrl: 'https://api.avax.network/ext/bc/C/rpc',
  nativeCurrency: {
    name: 'Avax',
    symbol: 'AVAX',
    decimals: 18,
  },
}

export const CANTO: Chain = {
  id: 7700,
  name: 'Canto',
  rpcUrls: {
    primary: 'https://mainnode.plexnode.org:8545',
    fallback: 'https://canto.slingshot.finance',
  },
  explorerUrl: 'https://tuber.build/',
  explorerName: 'Canto Explorer',
  blockTime: 6000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/canto.cb85e14f.svg',
  networkName: 'Canto',
  networkUrl: 'https://canto.slingshot.finance',
  nativeCurrency: {
    name: 'Canto',
    symbol: 'CANTO',
    decimals: 18,
  },
}

export const OPTIMISM: Chain = {
  id: 10,
  name: 'Optimism',
  rpcUrls: {
    primary: 'https://mainnet.optimism.io',
    fallback: 'https://1rpc.io/op',
  },
  explorerUrl: 'https://optimistic.etherscan.io',
  explorerName: 'Optimism Explorer',
  blockTime: 2000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/optimism.84d4f0ef.svg',
  networkName: 'OP Mainnet',
  networkUrl: 'https://mainnet.optimism.io',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const POLYGON: Chain = {
  id: 137,
  name: 'Polygon',
  rpcUrls: {
    primary: 'https://polygon-bor.publicnode.com',
    fallback: 'https://polygon.llamarpc.com',
  },
  explorerUrl: 'https://polygonscan.com',
  explorerName: 'PolygonScan',
  blockTime: 2000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/polygon.237cd2b6.svg',
  networkName: 'Polygon Mainnet',
  networkUrl: 'https://polygon-rpc.com',
  nativeCurrency: {
    name: 'Matic',
    symbol: 'MATIC',
    decimals: 18,
  },
}

export const DFK: Chain = {
  id: 53935,
  name: 'DFK Chain',
  rpcUrls: {
    primary: 'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
    fallback: 'https://dfkchain.api.onfinality.io/public',
  },
  explorerUrl: 'https://subnets.avax.network/defi-kingdoms',
  explorerName: 'DFK Subnet Explorer',
  blockTime: 2000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/dfk.2bd1f0e4.svg',
  networkName: 'DFK Chain',
  networkUrl: 'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
  nativeCurrency: {
    name: 'Jewel',
    symbol: 'JEWEL',
    decimals: 18,
  },
}

export const KLAYTN: Chain = {
  id: 8217,
  name: 'Klaytn',
  rpcUrls: {
    primary: 'https://klaytn.blockpi.network/v1/rpc/public',
    fallback: 'https://klaytn.api.onfinality.io/public',
  },
  explorerUrl: 'https://scope.klaytn.com',
  explorerName: 'Klaytn Explorer',
  blockTime: 1000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/klaytn.59495fbb.svg',
  networkName: 'Klaytn Mainnet Cypress',
  networkUrl: 'https://public-en-cypress.klaytn.net',
  nativeCurrency: {
    name: 'Klaytn',
    symbol: 'KLAY',
    decimals: 18,
  },
}

export const FANTOM: Chain = {
  id: 250,
  name: 'Fantom',
  rpcUrls: {
    primary: 'https://rpc.ftm.tools',
    fallback: 'https://fantom-rpc.gateway.pokt.network/',
  },
  explorerUrl: 'https://ftmscan.com',
  explorerName: 'FTMScan',
  blockTime: 1000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/fantom.1e444dad.svg',
  networkName: 'Fantom Opera',
  networkUrl: 'https://rpc.ftm.tools',
  nativeCurrency: {
    name: 'Fantom',
    symbol: 'FTM',
    decimals: 18,
  },
}

export const CRONOS: Chain = {
  id: 25,
  name: 'Cronos',
  rpcUrls: {
    primary: 'https://evm-cronos.crypto.org',
    fallback: 'https://cronos.blockpi.network/v1/rpc/public',
  },
  explorerUrl: 'https://cronoscan.com',
  explorerName: 'CronoScan',
  blockTime: 6000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/cronos.b06f8311.svg',
  networkName: 'Cronos Mainnet',
  networkUrl: 'https://evm.cronos.org',
  nativeCurrency: {
    name: 'Cronos',
    symbol: 'CRO',
    decimals: 18,
  },
}

export const BOBA: Chain = {
  id: 288,
  name: 'Boba Chain',
  rpcUrls: {
    primary: 'https://mainnet.boba.network',
    fallback: 'https://replica.boba.network',
  },
  explorerUrl: 'https://bobascan.com',
  explorerName: 'Boba Explorer',
  blockTime: 1000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/boba.2072e50b.svg',
  networkName: 'Boba Network',
  networkUrl: 'https://mainnet.boba.network',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const METIS: Chain = {
  id: 1088,
  name: 'Metis',
  rpcUrls: {
    primary: 'https://andromeda.metis.io/?owner=1088',
    fallback: 'https://metis-mainnet.public.blastapi.io',
  },
  explorerUrl: 'https://andromeda-explorer.metis.io',
  explorerName: 'Metis Explorer',
  blockTime: 4000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/metis.3114f675.svg',
  networkName: 'Metis Andromeda Mainnet',
  networkUrl: 'https://andromeda.metis.io/?owner=1088',
  nativeCurrency: {
    name: 'Metis',
    symbol: 'METIS',
    decimals: 18,
  },
}

export const AURORA: Chain = {
  id: 1313161554,
  name: 'Aurora',
  rpcUrls: {
    primary: 'https://mainnet.aurora.dev',
    fallback: 'https://1rpc.io/aurora',
  },
  explorerUrl: 'https://explorer.mainnet.aurora.dev',
  explorerName: 'Aurora Explorer',
  blockTime: 1000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/aurora.5a46037d.svg',
  networkName: 'Aurora Mainnet',
  networkUrl: 'https://mainnet.aurora.dev',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const HARMONY: Chain = {
  id: 1666600000,
  name: 'Harmony',
  rpcUrls: {
    primary: 'https://api.harmony.one',
    fallback: 'https://api.s0.t.hmny.io',
  },
  explorerUrl: 'https://explorer.harmony.one',
  explorerName: 'Harmony Explorer',
  blockTime: 2000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/harmony.af12f77e.svg',
  networkName: 'Harmony Mainnet Shard 0',
  networkUrl: 'https://api.harmony.one',
  nativeCurrency: {
    name: 'Harmony One',
    symbol: 'ONE',
    decimals: 18,
  },
}

export const MOONBEAM: Chain = {
  id: 1284,
  name: 'Moonbeam',
  rpcUrls: {
    primary: 'https://rpc.api.moonbeam.network',
    fallback: 'https://moonbeam.public.blastapi.io',
  },
  explorerUrl: 'https://moonbeam.moonscan.io',
  explorerName: 'Moonbeam Explorer',
  blockTime: 12000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/moonbeam.284ab9b4.svg',
  networkName: 'Moonbeam',
  networkUrl: 'https://rpc.api.moonbeam.network',
  nativeCurrency: {
    name: 'Glimmer',
    symbol: 'GLMR',
    decimals: 18,
  },
}

export const MOONRIVER: Chain = {
  id: 1285,
  name: 'Moonriver',
  rpcUrls: {
    primary: 'https://rpc.api.moonriver.moonbeam.network',
    fallback: 'https://moonriver.public.blastapi.io',
  },
  explorerUrl: 'https://moonriver.moonscan.io',
  explorerName: 'Moonriver Explorer',
  blockTime: 12000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/moonriver.3fb35010.svg',
  networkName: 'Moonriver',
  networkUrl: 'https://rpc.api.moonriver.moonbeam.network',
  nativeCurrency: {
    name: 'Moonriver',
    symbol: 'MOVR',
    decimals: 18,
  },
}

export const DOGE: Chain = {
  id: 2000,
  name: 'Dogechain',
  rpcUrls: {
    primary: 'https://rpc.dogechain.dog',
    fallback: 'https://rpc01-sg.dogechain.dog',
  },
  explorerUrl: 'https://explorer.dogechain.dog',
  explorerName: 'Dogechain Explorer',
  blockTime: 2000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/dogechain.36935650.svg',
  networkName: 'Dogechain Mainnet',
  networkUrl: 'https://rpc.dogechain.dog',
  nativeCurrency: {
    name: 'DOGE',
    symbol: 'DOGE',
    decimals: 18,
  },
}

export const BASE: Chain = {
  id: 8453,
  name: 'Base',
  rpcUrls: {
    primary: 'https://base.blockpi.network/v1/rpc/public',
    fallback: 'https://developer-access-mainnet.base.org',
  },
  explorerUrl: 'https://basescan.org',
  explorerName: 'BaseScan',
  blockTime: 3000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/base.d919fbef.svg',
  networkName: 'Base',
  networkUrl: 'https://mainnet.base.org',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const BLAST: Chain = {
  id: 81457,
  name: 'Blast',
  rpcUrls: {
    primary:
      'https://lingering-indulgent-replica.blast-mainnet.quiknode.pro/6667a8f4be701cb6549b415d567bc706fb2f13a8/',
    fallback: 'https://blast.blockpi.network/v1/rpc/publicChain',
  },
  explorerUrl: 'https://blastscan.io',
  explorerName: 'Blastscan',
  blockTime: 3000,
  imgUrl:
    'https://45a97b3d.sanguine-fe.pages.dev/_next/static/media/blast.e39807f8.svg',
  networkName: 'Blast',
  networkUrl: 'https://rpc.blast.io',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const SCROLL: Chain = {
  id: 534352,
  name: 'Scroll',
  rpcUrls: {
    primary: 'https://rpc.scroll.io/',
    fallback: 'https://scroll.blockpi.network/v1/rpc/public',
  },
  explorerUrl: 'https://scrollscan.com',
  explorerName: 'Scrollscan',
  blockTime: 3000,
  imgUrl:
    'https://fe-adds-scroll.sanguine-fe.pages.dev/_next/static/media/scroll.a805c122.svg',
  networkName: 'Scroll',
  networkUrl: 'https://rpc.scroll.io',
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
  },
}

export const CHAINS = {
  ETHEREUM,
  ARBITRUM,
  BNBCHAIN,
  AVALANCHE,
  CANTO,
  OPTIMISM,
  POLYGON,
  DFK,
  KLAYTN,
  FANTOM,
  CRONOS,
  BOBA,
  METIS,
  AURORA,
  HARMONY,
  MOONBEAM,
  MOONRIVER,
  DOGE,
  BASE,
  BLAST,
  SCROLL,
}

export const CHAINS_ARRAY = Object.values(CHAINS)

export const CHAINS_BY_ID = _.keyBy(CHAINS, 'id')
