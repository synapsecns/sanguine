import {
  ANALYTICS_PATH,
  SWAP_PATH,
  STAKE_PATH,
  POOLS_PATH,
  LANDING_PATH,
  BRIDGE_PATH,
} from './urls'

export interface RouteObject {
  [key: string]: {
    path: string
    text: string
    match: string | RegExp | { startsWith: string; endsWith: string }
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
    match: {
      startsWith: '/',
      endsWith: '/',
    },
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
}
