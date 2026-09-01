import * as all from './master'
import { Chain } from '@/utils/types/index'

export type ChainsByChainID = {
  [cID: number]: Chain
}

// Retired chains remain in CHAINS_BY_ID so historical activity can still render.
export const RETIRED_CHAIN_IDS = new Set([all.DFK.id])

export const isActiveChainId = (chainId: number | string) =>
  !RETIRED_CHAIN_IDS.has(Number(chainId))

export const sortChains = (chains: Chain[]) =>
  Object.values(chains).sort((a, b) => b.priorityRank - a.priorityRank)

const ALL_CHAINS_ARR = Object.values(all)

export const CHAINS_ARR = ALL_CHAINS_ARR.filter(({ id }) =>
  isActiveChainId(id)
).sort((a, b) => b.priorityRank - a.priorityRank)

const getids = (chains: Chain[]) => {
  const outObj = {}
  chains.map((chain) => {
    outObj[chain.chainSymbol] = chain.id
  })
  return outObj
}
const getChainsByID = (chains: Chain[]): ChainsByChainID => {
  const outObj: ChainsByChainID = {}
  chains.map((chain) => {
    outObj[chain.id] = chain
  })
  return outObj
}

export const CHAIN_IDS = getids(CHAINS_ARR) // used to be ids
export const CHAINS_BY_ID = getChainsByID(ALL_CHAINS_ARR)
export const ACTIVE_CHAINS_BY_ID = getChainsByID(CHAINS_ARR)
export const ORDERED_CHAINS_BY_ID = CHAINS_ARR.map((chain) => String(chain.id))

export const ChainId = {
  ETH: 1,
  ROPSTEN: 3,
  RINKEBY: 4,
  GÖRLI: 5,
  OPTIMISM: 10,
  CRONOS: 25,
  KOVAN: 42,
  BSC: 56,
  UNICHAIN: 130,
  POLYGON: 137,
  FANTOM: 250,
  BOBA: 288,
  WORLDCHAIN: 480,
  HYPEREVM: 999,
  METIS: 1088,
  MOONBEAM: 1284,
  MOONRIVER: 1285,
  DOGECHAIN: 2000,
  CANTO: 7700,
  KLAYTN: 8217,
  HARDHAT: 31337,
  ARBITRUM: 42161,
  BASE: 8453,
  BERACHAIN: 80094,
  BLAST: 81457,
  LINEA: 59144,
  SCROLL: 534352,
  AVALANCHE: 43114,
  AURORA: 1313161554,
  HARMONY: 1666600000,
  TERRA: 121014925, //"columbus-5", the day columbus reportedly landed in america followed by 5
}

export const AcceptedChainId = Object.fromEntries(
  Object.entries(ChainId).map(([key, value]) => [value, key])
)
