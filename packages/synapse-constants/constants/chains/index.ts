import { Chain } from '../types'
import * as all from './master'

export * from './master'

export const hello = 2

export type ChainsByChainID = {
  [cID: number]: Chain
}
export const sortChains = (chains: Chain[]) =>
  Object.values(chains).sort(
    (a: Chain, b: Chain) => (b.priorityRank ?? 0) - (a.priorityRank ?? 0)
  )

export const CHAINS_ARR = Object.values(all)
  .filter((item): item is Chain => typeof item !== 'number')
  .sort((a: Chain, b: Chain) => (b.priorityRank ?? 0) - (a.priorityRank ?? 0))

const getChainEnumById = () => {
  const outObj: Record<number, string> = {}
  CHAINS_ARR.map((chain: any) => {
    outObj[chain.id] = chain.codeName
  })
  return outObj
}

const getids = () => {
  const outObj: { [key: string]: any } = {}
  CHAINS_ARR.map((chain: any) => {
    outObj[chain.chainSymbol] = chain.id
  })
  return outObj
}
const getChainsByID = (): ChainsByChainID => {
  const outObj: ChainsByChainID = {}
  CHAINS_ARR.map((chain: any) => {
    outObj[chain.id] = chain
  })
  return outObj
}

export const CHAIN_ENUM_BY_ID = getChainEnumById()
export const CHAIN_IDS = getids() // used to be ids
export const CHAINS_BY_ID = getChainsByID()
export const ORDERED_CHAINS_BY_ID = CHAINS_ARR.map((chain: Chain) =>
  String(chain.id)
)

export const CHAIN_ID_NAMES_REVERSE = Object.fromEntries(
  Object.entries(CHAIN_ENUM_BY_ID).map(([k, v]) => [v, k])
)

export const PAUSED_FROM_CHAIN_IDS = []
export const PAUSED_TO_CHAIN_IDS = [all.DOGE.id]

export const ChainId = {
  ETH: 1,
  ROPSTEN: 3,
  RINKEBY: 4,
  GÖRLI: 5,
  OPTIMISM: 10,
  CRONOS: 25,
  KOVAN: 42,
  BSC: 56,
  POLYGON: 137,
  FANTOM: 250,
  BOBA: 288,
  METIS: 1088,
  MOONBEAM: 1284,
  MOONRIVER: 1285,
  DOGECHAIN: 2000,
  CANTO: 7700,
  KLAYTN: 8217,
  HARDHAT: 31337,
  ARBITRUM: 42161,
  BASE: 8453,
  AVALANCHE: 43114,
  DFK: 53935,
  AURORA: 1313161554,
  HARMONY: 1666600000,
  TERRA: 121014925, //"columbus-5", the day columbus reportedly landed in america followed by 5
}

export const AcceptedChainId = Object.fromEntries(
  Object.entries(ChainId).map(([key, value]) => [value, key])
)

export const BRIDGE_CONTRACTS = {
  [ChainId.ETH]: '0x2796317b0fF8538F253012862c06787Adfb8cEb6',
  [ChainId.OPTIMISM]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.CRONOS]: '0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9',
  [ChainId.BSC]: '0xd123f70AE324d34A9E76b67a27bf77593bA8749f',
  [ChainId.POLYGON]: '0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280',
  [ChainId.FANTOM]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.BOBA]: '0x432036208d2717394d2614d6697c46DF3Ed69540',
  [ChainId.METIS]: '0x06Fea8513FF03a0d3f61324da709D4cf06F42A5c',
  [ChainId.MOONBEAM]: '0x84A420459cd31C3c34583F67E0f0fB191067D32f',
  [ChainId.MOONRIVER]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [ChainId.KLAYTN]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.ARBITRUM]: '0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9',
  [ChainId.AVALANCHE]: '0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE',
  [ChainId.DFK]: '0xE05c976d3f045D0E6E7A6f61083d98A15603cF6A',
  [ChainId.AURORA]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [ChainId.HARMONY]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.CANTO]: '0xDde5BEC4815E1CeCf336fb973Ca578e8D83606E0',
  [ChainId.DOGECHAIN]: '0x9508BF380c1e6f751D97604732eF1Bae6673f299',
  [ChainId.BASE]: '0xf07d1C752fAb503E47FEF309bf14fbDD3E867089',
}

export const CCTP_CONTRACTS = {
  [ChainId.ARBITRUM]: '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E',
  [ChainId.AVALANCHE]: '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E',
  [ChainId.BASE]: '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E',
  [ChainId.ETH]: '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E',
  [ChainId.POLYGON]: '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E',
  [ChainId.OPTIMISM]: '0x12715a66773BD9C54534a01aBF01d05F6B4Bd35E',
}
