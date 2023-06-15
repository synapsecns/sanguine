import {
  ANALYTICS_PATH,
  SWAP_PATH,
  STAKE_PATH,
  POOLS_PATH,
  LANDING_PATH,
  BRIDGE_PATH,
  CONTRACTS_PATH,
} from './urls'

interface RouteObject {
  [name: string]: {
    path: string
    text: string
    match: string | null
  }
}

export const NAVIGATION: RouteObject = {
  About: {
    path: LANDING_PATH,
    text: 'About',
    match: '/landing',
  },
  Bridge: {
    path: BRIDGE_PATH,
    text: 'Bridge',
    match: '/?outputChain',
  },
  Swap: {
    path: SWAP_PATH,
    text: 'Swap',
    match: '/swap',
  },
  Pools: {
    path: POOLS_PATH,
    text: 'Pools',
    match: '/pool',
  },
  Stake: {
    path: STAKE_PATH,
    text: 'Stake',
    match: '/stake',
  },
  Analytics: {
    path: ANALYTICS_PATH,
    text: 'Explorer',
    match: null,
  },
  Contracts: {
    path: CONTRACTS_PATH,
    text: 'Contracts',
    match: '/contracts',
  },
}
