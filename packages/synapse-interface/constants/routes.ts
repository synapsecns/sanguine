import {
  EXPLORER_PATH,
  SWAP_PATH,
  STAKE_PATH,
  POOLS_PATH,
  LANDING_PATH,
  BRIDGE_PATH,
  INTERCHAIN_LINK,
  SOLANA_LINK,
} from './urls'

export interface RouteObject {
  [key: string]: {
    path: string
    text: string
    match: string | { startsWith: string }
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
    match: '/',
  },
  Swap: {
    path: SWAP_PATH,
    text: 'Swap',
    match: '/swap',
  },
  Pools: {
    path: POOLS_PATH,
    text: 'Pools',
    match: {
      startsWith: '/pool',
    },
  },
  Stake: {
    path: STAKE_PATH,
    text: 'Stake',
    match: {
      startsWith: '/stake',
    },
  },
  Analytics: {
    path: EXPLORER_PATH,
    text: 'Explorer',
    match: null,
  },
  Contracts: {
    path: INTERCHAIN_LINK,
    text: 'Interchain Network',
    match: null,
  },
  Solana: {
    path: SOLANA_LINK,
    text: 'Solana Bridge',
    match: null,
  },
}
