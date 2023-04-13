import {
  ANALYTICS_PATH,
  SWAP_PATH,
  STAKE_PATH,
  POOLS_PATH,
  LANDING_PATH,
  BRIDGE_PATH,
} from './urls'

interface RouteObject {
  [name: string]: {
    path: string
    text: string
  }
}

export const NAVIGATION: RouteObject = {
  About: {
    path: LANDING_PATH,
    text: 'About',
  },
  Bridge: {
    path: BRIDGE_PATH,
    text: 'Bridge',
  },
  Swap: {
    path: SWAP_PATH,
    text: 'Swap',
  },
  Pools: {
    path: POOLS_PATH,
    text: 'Pools',
  },
  Stake: {
    path: STAKE_PATH,
    text: 'Stake',
  },
  Analytics: {
    path: ANALYTICS_PATH,
    text: 'Explorer',
  },
}
