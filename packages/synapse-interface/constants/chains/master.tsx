import arbitrumImg from '@assets/chains/arbitrum.svg'
import auroraImg from '@assets/chains/aurora.svg'
import avalancheImg from '@assets/chains/avalanche.svg'
import baseImg from '@assets/chains/base.svg'
import blastImg from '@assets/chains/blast.svg'
import bobaImg from '@assets/chains/boba.svg'
import bscImg from '@assets/chains/bnb.svg'
import cantoImg from '@assets/chains/canto.svg'
import cronosImg from '@assets/chains/cronos.svg'
import dfkImg from '@assets/chains/dfk.svg'
import dogechainImg from '@assets/chains/dogechain.svg'
import ethImg from '@assets/chains/ethereum.svg'
import fantomImg from '@assets/chains/fantom.svg'
import harmonyImg from '@assets/chains/harmony.svg'
import klaytnImg from '@assets/chains/klaytn.svg'
import metisImg from '@assets/chains/metis.svg'
import moonbeamImg from '@assets/chains/moonbeam.svg'
import moonriverImg from '@assets/chains/moonriver.svg'
import optimismImg from '@assets/chains/optimism.svg'
import polygonImg from '@assets/chains/polygon.svg'
import scrollImg from '@assets/chains/scroll.svg'
import lineaImg from '@assets/chains/linea.svg'
import worldchainImg from '@assets/chains/worldchain.svg'

import ethExplorerImg from '@assets/explorer/etherscan.svg'
import arbitrumExplorerImg from '@assets/explorer/arbiscan.svg'
import bnbExplorerImg from '@assets/explorer/bscscan.svg'
import baseExplorerImg from '@assets/explorer/basescan.svg'
import avalancheExplorerImg from '@assets/explorer/snowscan.svg'
import fantomExplorerImg from '@assets/explorer/ftmscan.svg'

import { zeroAddress } from 'viem'

import { Chain } from '@types'

export const ETH: Chain = {
  priorityRank: 100,
  id: 1,
  chainSymbol: 'ETH',
  name: 'Ethereum',
  chainImg: ethImg,
  layer: 1,
  rpcUrls: {
    primary:
      'https://eth-mainnet.g.alchemy.com/v2/rJ3f0IWjZbpgEwnzrRS6yYO3WNH0jGle',
    fallback: 'https://eth.llamarpc.com',
  },
  explorerUrl: 'https://etherscan.io',
  explorerName: 'Etherscan',
  explorerImg: ethExplorerImg,
  blockTime: 12000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: ethImg,
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
  blockTime: 300,
  rpcUrls: {
    primary:
      'https://arb-mainnet.g.alchemy.com/v2/7kjdkqKTh1zQ1mRYGi4nJJbxbyJXHkef',
    fallback: 'https://arb1.arbitrum.io/rpc',
  },
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: ethImg,
  },
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
  blockTime: 3000,
  rpcUrls: {
    primary: 'https://bsc-dataseed1.ninicoin.io/',
    fallback: 'https://bsc-dataseed2.ninicoin.io',
  },
  nativeCurrency: {
    name: 'Binance Coin',
    symbol: 'BNB',
    decimals: 18,
    address: zeroAddress,
    icon: bscImg,
  },
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
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://api.avax.network/ext/bc/C/rpc',
    fallback: 'https://1rpc.io/avax/c',
  },
  nativeCurrency: {
    name: 'Avax',
    symbol: 'AVAX',
    decimals: 18,
    address: zeroAddress,
    icon: avalancheImg,
  },
  explorerUrl: 'https://snowscan.xyz/',
  explorerName: 'SnowScan',
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
  blockTime: 6000,
  rpcUrls: {
    primary: 'https://mainnode.plexnode.org:8545',
    fallback: 'https://canto.slingshot.finance',
  },
  nativeCurrency: {
    name: 'Canto',
    symbol: 'CANTO',
    decimals: 18,
    address: zeroAddress,
    icon: cantoImg,
  },
  explorerUrl: 'https://tuber.build/',
  explorerName: 'Canto Explorer',
  explorerImg: cantoImg,
  color: 'green',
}

export const OPTIMISM: Chain = {
  priorityRank: 91,
  id: 10,
  chainSymbol: 'OPTIMISM',
  name: 'Optimism',
  chainImg: optimismImg,
  layer: 2,
  blockTime: 2000,
  rpcUrls: {
    primary:
      'https://opt-mainnet.g.alchemy.com/v2/x1--EvECmiLCc4IRpo1granp1S7xxbxQ',
    fallback: 'https://gateway.tenderly.co/public/optimism',
  },
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: ethImg,
  },
  explorerUrl: 'https://optimistic.etherscan.io',
  explorerName: 'Optimism Explorer',
  explorerImg: optimismImg,
  color: 'red',
}

export const POLYGON: Chain = {
  priorityRank: 80,
  id: 137,
  chainSymbol: 'POLYGON',
  name: 'Polygon',
  chainImg: polygonImg,
  layer: 2,
  blockTime: 2000,
  rpcUrls: {
    primary:
      'https://polygon-mainnet.g.alchemy.com/v2/mN1t8Oc6E912QF28iPHaRvVEmv6EpYSs',
    fallback: 'https://polygon.llamarpc.com',
  },
  nativeCurrency: {
    name: 'Matic',
    symbol: 'MATIC',
    decimals: 18,
    address: zeroAddress,
    icon: polygonImg,
  },
  explorerUrl: 'https://polygonscan.com',
  explorerName: 'PolygonScan',
  explorerImg: polygonImg,
  color: 'purple',
}

export const DFK: Chain = {
  priorityRank: 75,
  id: 53935,
  chainSymbol: 'DFK',
  name: 'DFK Chain',
  chainImg: dfkImg,
  layer: 1,
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
    fallback: 'https://dfkchain.rpc.defikingdoms.com/api=654302102',
  },
  nativeCurrency: {
    name: 'Jewel',
    symbol: 'JEWEL',
    decimals: 18,
    address: zeroAddress,
    icon: dfkImg,
  },
  explorerUrl: 'https://subnets.avax.network/defi-kingdoms',
  explorerName: 'DFK Subnet Explorer',
  explorerImg: dfkImg,
  color: 'lime',
}

export const KLAYTN: Chain = {
  priorityRank: 70,
  id: 8217,
  chainSymbol: 'KLAYTN',
  name: 'Klaytn',
  chainImg: klaytnImg,
  layer: 1,
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://klaytn.blockpi.network/v1/rpc/public',
    fallback: 'https://internal.klaytn.rpc.defikingdoms.com/api=654302102',
  },
  nativeCurrency: {
    name: 'Klaytn',
    symbol: 'KLAY',
    decimals: 18,
    address: zeroAddress,
    icon: klaytnImg,
  },
  explorerUrl: 'https://scope.klaytn.com',
  explorerName: 'Klaytn Explorer',
  explorerImg: klaytnImg,
  color: 'orange',
}

export const FANTOM: Chain = {
  priorityRank: 70,
  id: 250,
  chainSymbol: 'FANTOM',
  name: 'Fantom',
  chainImg: fantomImg,
  layer: 1,
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://rpc.ftm.tools',
    fallback: 'https://fantom-rpc.gateway.pokt.network/',
  },
  nativeCurrency: {
    name: 'Fantom',
    symbol: 'FTM',
    decimals: 18,
    address: zeroAddress,
    icon: fantomImg,
  },
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
  blockTime: 6000,
  rpcUrls: {
    primary: 'https://evm-cronos.crypto.org',
    fallback: 'https://cronos.blockpi.network/v1/rpc/public',
  },
  nativeCurrency: {
    name: 'Cronos',
    symbol: 'CRO',
    decimals: 18,
    address: zeroAddress,
    icon: cronosImg,
  },
  explorerUrl: 'https://cronoscan.com',
  explorerName: 'CronoScan',
  explorerImg: cronosImg,
  color: 'gray',
}

export const BOBA: Chain = {
  priorityRank: 10,
  id: 288,
  chainSymbol: 'BOBA',
  name: 'Boba Chain',
  chainImg: bobaImg,
  layer: 2,
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://mainnet.boba.network',
    fallback: 'https://replica.boba.network',
  },
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: bobaImg,
  },
  explorerUrl: 'https://bobascan.com',
  explorerName: 'Boba Explorer',
  explorerImg: bobaImg,
  color: 'lime',
}

export const METIS: Chain = {
  priorityRank: 10,
  id: 1088,
  chainSymbol: 'METIS',
  name: 'Metis',
  chainImg: metisImg,
  layer: 2,
  blockTime: 4000,
  rpcUrls: {
    primary: 'https://andromeda.metis.io/?owner=1088',
    fallback: 'https://metis-mainnet.public.blastapi.io',
  },
  nativeCurrency: {
    name: 'Metis',
    symbol: 'Metis',
    decimals: 18,
    address: zeroAddress,
    icon: metisImg,
  },
  explorerUrl: 'https://andromeda-explorer.metis.io',
  explorerName: 'Metis Explorer',
  explorerImg: metisImg,
  color: 'teal',
}

export const AURORA: Chain = {
  priorityRank: 10,
  id: 1313161554,
  chainSymbol: 'AURORA',
  name: 'Aurora',
  chainImg: auroraImg,
  layer: 1,
  blockTime: 1000,
  rpcUrls: {
    primary: 'https://mainnet.aurora.dev',
    fallback: 'https://1rpc.io/aurora',
  },
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: auroraImg,
  },
  explorerUrl: 'https://explorer.mainnet.aurora.dev',
  explorerName: 'Aurora Explorer',
  explorerImg: auroraImg,
  color: 'lime',
}

export const HARMONY: Chain = {
  priorityRank: 10,
  id: 1666600000,
  chainSymbol: 'HARMONY',
  name: 'Harmony',
  chainImg: harmonyImg,
  layer: 1,
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://api.harmony.one',
    fallback: 'https://api.s0.t.hmny.io',
  },
  nativeCurrency: {
    name: 'Harmony One',
    symbol: 'ONE',
    decimals: 18,
    address: zeroAddress,
    icon: harmonyImg,
  },
  explorerUrl: 'https://explorer.harmony.one',
  explorerName: 'Harmony Explorer',
  explorerImg: harmonyImg,
  color: 'cyan',
}

export const MOONBEAM: Chain = {
  priorityRank: 0,
  id: 1284,
  chainSymbol: 'MOONBEAM',
  name: 'Moonbeam',
  chainImg: moonbeamImg,
  layer: 1,
  blockTime: 12000,
  rpcUrls: {
    primary: 'https://rpc.api.moonbeam.network',
    fallback: 'https://moonbeam.public.blastapi.io',
  },
  nativeCurrency: {
    name: 'Glimmer',
    symbol: 'GLMR',
    decimals: 18,
    address: zeroAddress,
    icon: moonbeamImg,
  },
  explorerUrl: 'https://moonbeam.moonscan.io',
  explorerName: 'Moonbeam Explorer',
  explorerImg: moonbeamImg,
  color: 'teal',
}

export const MOONRIVER: Chain = {
  priorityRank: 0,
  id: 1285,
  chainSymbol: 'MOONRIVER',
  name: 'Moonriver',
  chainImg: moonriverImg,
  layer: 1,
  blockTime: 12000,
  rpcUrls: {
    primary: 'https://rpc.api.moonriver.moonbeam.network',
    fallback: 'https://moonriver.public.blastapi.io',
  },
  nativeCurrency: {
    name: 'Moonriver',
    symbol: 'MOVR',
    decimals: 18,
    address: zeroAddress,
    icon: moonriverImg,
  },
  explorerUrl: 'https://moonriver.moonscan.io',
  explorerName: 'Moonriver Explorer',
  explorerImg: moonriverImg,
  color: 'purple',
}

export const DOGE: Chain = {
  priorityRank: 0,
  id: 2000,
  chainSymbol: 'DOGE',
  name: 'Dogechain',
  chainImg: dogechainImg,
  layer: 1,
  blockTime: 2000,
  rpcUrls: {
    primary: 'https://rpc.dogechain.dog',
    fallback: 'https://rpc01-sg.dogechain.dog',
  },
  nativeCurrency: {
    name: 'DOGE',
    symbol: 'DOGE',
    decimals: 18,
    address: zeroAddress,
    icon: dogechainImg,
  },
  explorerUrl: 'https://explorer.dogechain.dog',
  explorerName: 'Dogechain Explorer',
  explorerImg: dogechainImg,
  color: 'purple',
}

export const BASE: Chain = {
  priorityRank: 90,
  id: 8453,
  chainSymbol: 'BASE',
  name: 'Base',
  chainImg: baseImg,
  layer: 2,
  rpcUrls: {
    primary:
      'https://base-mainnet.g.alchemy.com/v2/_YKy-Vm3LsknT8JKSa2ZTSmKu9Qp01Vd',
    fallback: 'https://developer-access-mainnet.base.org',
  },
  explorerUrl: 'https://basescan.org',
  explorerName: 'BaseScan',
  explorerImg: baseExplorerImg,
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: baseImg,
  },
  color: 'blue',
}

export const BLAST: Chain = {
  priorityRank: 90,
  id: 81457,
  chainSymbol: 'BLAST',
  name: 'Blast',
  chainImg: blastImg,
  layer: 2,
  rpcUrls: {
    primary:
      'https://lingering-indulgent-replica.blast-mainnet.quiknode.pro/6667a8f4be701cb6549b415d567bc706fb2f13a8/',
    fallback: 'https://blast.blockpi.network/v1/rpc/publicChain',
  },
  explorerUrl: 'https://blastscan.io',
  explorerName: 'Blastscan',
  explorerImg: blastImg,
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: blastImg,
  },
  color: 'yellow',
}

export const SCROLL: Chain = {
  priorityRank: 90,
  id: 534352,
  chainSymbol: 'SCROLL',
  name: 'Scroll',
  chainImg: scrollImg,
  layer: 2,
  rpcUrls: {
    primary: 'https://rpc.scroll.io/',
    fallback: 'https://scroll.blockpi.network/v1/rpc/public',
  },
  explorerUrl: 'https://scrollscan.com',
  explorerName: 'Scrollscan',
  explorerImg: scrollImg,
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: scrollImg,
  },
  color: 'orange',
  isNew: false,
}

export const LINEA: Chain = {
  priorityRank: 90,
  id: 59144,
  chainSymbol: 'LINEA',
  name: 'Linea',
  chainImg: lineaImg,
  layer: 2,
  rpcUrls: {
    primary: 'https://rpc.linea.build',
    fallback: 'https://linea.blockpi.network/v1/rpc/public',
  },
  explorerUrl: 'https://lineascan.build',
  explorerName: 'LineaScan',
  explorerImg: lineaImg,
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: lineaImg,
  },
  color: 'black',
  isNew: false,
}

export const WORLDCHAIN: Chain = {
  priorityRank: 99,
  id: 480,
  chainSymbol: 'WORLDCHAIN',
  name: 'World Chain',
  chainImg: worldchainImg,
  layer: 2,
  rpcUrls: {
    primary: 'https://worldchain-mainnet.g.alchemy.com/public',
    fallback: 'https://worldchain-mainnet.g.alchemy.com/public',
  },
  explorerUrl: 'https://worldchain-mainnet.explorer.alchemy.com',
  explorerName: 'World Chain Explorer',
  explorerImg: worldchainImg,
  blockTime: 3000,
  nativeCurrency: {
    name: 'Ethereum',
    symbol: 'ETH',
    decimals: 18,
    address: zeroAddress,
    icon: ethImg,
  },
  color: 'black',
  isNew: true,
}
