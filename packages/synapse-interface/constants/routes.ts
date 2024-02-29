import {
  EXPLORER_PATH,
  SWAP_PATH,
  STAKE_PATH,
  POOLS_PATH,
  POOL_PATH,
  LANDING_PATH,
  BRIDGE_PATH,
  INTERCHAIN_LINK,
  SOLANA_BRIDGE_LINK,
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
    match: LANDING_PATH,
  },
  Bridge: {
    path: BRIDGE_PATH,
    text: 'Bridge',
    match: BRIDGE_PATH,
  },
  Swap: {
    path: SWAP_PATH,
    text: 'Swap',
    match: SWAP_PATH,
  },
  Pools: {
    path: POOLS_PATH,
    text: 'Pools',
    match: {
      startsWith: POOL_PATH,
    },
  },
  Stake: {
    path: STAKE_PATH,
    text: 'Stake',
    match: {
      startsWith: STAKE_PATH,
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
    path: SOLANA_BRIDGE_LINK,
    text: 'Solana Bridge',
    match: null,
  },
}
