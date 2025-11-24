import {
  EXPLORER_PATH,
  SWAP_PATH,
  STAKE_PATH,
  POOLS_PATH,
  POOL_PATH,
  LANDING_PATH,
  BRIDGE_PATH,
  SOLANA_BRIDGE_LINK,
  SYN_TOKEN_LINK,
  STAKE_SYN_FOR_CX_URL,
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
  SYN: {
    path: SYN_TOKEN_LINK,
    text: '$SYN',
    match: null,
  },
  Solana: {
    path: SOLANA_BRIDGE_LINK,
    text: 'Solana Bridge',
    match: null,
  },
}
