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
    'https://www.synapseprotocol.com/_next/static/media/ethereum.ced7bf7d.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/arbitrum.9337c5cb.svg',
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
  imgUrl: 'https://www.synapseprotocol.com/_next/static/media/bnb.c7f1e06b.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/avalanche.80034115.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/canto.0da5392e.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/optimism.22de1334.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/polygon.8d188d92.svg',
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
  imgUrl: 'https://www.synapseprotocol.com/_next/static/media/dfk.1e8f3773.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/klaytn.0d2394a9.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/fantom.72c20b34.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/cronos.805a6c46.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/boba.837d03c1.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/metis.98d36df7.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/aurora.46aaee98.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/harmony.1c4d7731.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/moonbeam.902e3681.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/moonriver.3e23a112.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/dogechain.04c2007a.svg',
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
    'https://www.synapseprotocol.com/_next/static/media/base.1b390997.svg',
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
}

export const CHAINS_ARRAY = Object.values(CHAINS)

export const CHAINS_BY_ID = _.keyBy(CHAINS, 'id')
