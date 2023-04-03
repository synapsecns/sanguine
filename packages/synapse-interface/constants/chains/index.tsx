import * as all from './master'
import { Chain } from '@/utils/types/index'

type ChainsByChainID = {
  [cID: number]: Chain
}
export const CHAINS_ARR = Object.values(all).sort(
  (a, b) => b.visibilityRank - a.visibilityRank
)

const getChainEnumById = () => {
  let outObj: Record<number, string> = {}
  CHAINS_ARR.map((chain) => {
    outObj[chain.id] = chain.codeName
  })
  return outObj
}

const getids = () => {
  let outObj = {}
  CHAINS_ARR.map((chain) => {
    outObj[chain.chainSymbol] = chain.id
  })
  return outObj
}
const getChainsByID = (): ChainsByChainID => {
  let outObj: ChainsByChainID = {}
  CHAINS_ARR.map((chain) => {
    outObj[chain.id] = chain
  })
  return outObj
}

export const CHAIN_ENUM_BY_ID = getChainEnumById()
export const CHAIN_IDS = getids() // used to be ids
export const CHAINS_BY_ID = getChainsByID()
export const ORDERED_CHAINS_BY_ID = CHAINS_ARR.map((chain) => chain.id)

// export const INVERTED_CHAIN_IDS =  Object.fromEntries(CHAIN_IDS).map((k, v) => v, k)// used to be INVERTED_CHAIN_ID_MAP
