import _ from 'lodash'
import { ChainId } from '@constants/networks'


import {
  OPTIMISM_ETH_SWAP_TOKEN,
  BOBA_ETH_SWAP_TOKEN,
  ARBITRUM_ETH_SWAP_TOKEN,
  AVALANCHE_AVETH_SWAP_TOKEN,
  HARMONY_ONEETH_SWAP_TOKEN,
  FANTOM_WETH_SWAP_TOKEN,
  METIS_WETH_SWAP_TOKEN,
  CANTO_WETH_SWAP_TOKEN
} from '@constants/tokens/ethswap'

import {
  ETH_POOL_SWAP_TOKEN,
  BSC_POOL_SWAP_TOKEN,
  OPTIMISM_POOL_SWAP_TOKEN,
  POLYGON_POOL_SWAP_TOKEN,
  AVALANCHE_POOL_SWAP_TOKEN,
  HARMONY_POOL_SWAP_TOKEN,
  BOBA_POOL_SWAP_TOKEN,
  AURORA_TS_POOL_SWAP_TOKEN,
  FANTOM_3POOL_SWAP_TOKEN,
  ARBITRUM_3POOL_SWAP_TOKEN,
  METIS_POOL_SWAP_TOKEN,
  CRONOS_POOL_SWAP_TOKEN,
  CANTO_POOL_SWAP_TOKEN,
  CANTO_USDC_SWAP_TOKEN,
  CANTO_WRAPPER_POOL_SWAP_TOKEN
} from '@constants/tokens/poolswap'

import { WETH, ETH, NETH, NUSD, SYN, FRAX, WETHBEAM, AVAX, MOVR, WAVAX, WMOVR, JEWEL, SYNJEWEL, WJEWEL, XJEWEL, SYNAVAX, DFK_USDC, UST, WBTC, KLAYTN_USDC, KLAYTN_USDT, KLAYTN_DAI, KLAYTN_WETH, DOGECHAIN_BUSD } from '@constants/tokens/basic'
import { DOG, LINK, HIGHSTREET, JUMP, NFD, GOHM, SOLAR, NEWO, SDT, USDB, GMX, VSTA, SFI, H2O, L2DAO, AGEUR, PLS, UNIDX } from '@constants/tokens/mintable'
import { HARMONY_AVAX_SWAP_TOKEN } from '@constants/tokens/avaxswap'







export const BRIDGE_ZAP_USD_SWAP_TOKEN_BY_CHAIN = {
  // [ChainId.OPTIMISM]:  OPTIMISM_POOL_SWAP_TOKEN, // USD pool not active
  [ChainId.BSC]:       BSC_POOL_SWAP_TOKEN,
  [ChainId.ETH]:       ETH_POOL_SWAP_TOKEN,
  [ChainId.OPTIMISM]:  OPTIMISM_POOL_SWAP_TOKEN,
  [ChainId.POLYGON]:   POLYGON_POOL_SWAP_TOKEN,
  [ChainId.FANTOM]:    FANTOM_3POOL_SWAP_TOKEN,
  [ChainId.BOBA]:      BOBA_POOL_SWAP_TOKEN,
  [ChainId.ARBITRUM]:  ARBITRUM_3POOL_SWAP_TOKEN,
  [ChainId.AVALANCHE]: AVALANCHE_POOL_SWAP_TOKEN,
  [ChainId.AURORA]:    AURORA_TS_POOL_SWAP_TOKEN,
  [ChainId.HARMONY]:   HARMONY_POOL_SWAP_TOKEN,
  [ChainId.METIS]:     METIS_POOL_SWAP_TOKEN,
  [ChainId.CRONOS]:    CRONOS_POOL_SWAP_TOKEN,
  [ChainId.DFK]:       { poolTokens: [DFK_USDC], isNative: false },
  [ChainId.DOGECHAIN]:    undefined,
  [ChainId.CANTO]:     CANTO_WRAPPER_POOL_SWAP_TOKEN
}

export const USD_SWAP_TOKENS = _.values(BRIDGE_ZAP_USD_SWAP_TOKEN_BY_CHAIN)

export const BRIDGE_ZAP_ETH_SWAP_TOKEN_BY_CHAIN = {
  [ChainId.OPTIMISM]:  OPTIMISM_ETH_SWAP_TOKEN,
  [ChainId.BSC]:       undefined,
  [ChainId.ETH]:       {poolTokens:[ETH], isNative: true},
  [ChainId.POLYGON]:   undefined,
  [ChainId.FANTOM]:    FANTOM_WETH_SWAP_TOKEN,
  [ChainId.BOBA]:      BOBA_ETH_SWAP_TOKEN,
  [ChainId.MOONBEAM]:  undefined,                            // {poolTokens: [WETHBEAM]},
  [ChainId.MOONRIVER]: undefined,
  [ChainId.ARBITRUM]:  ARBITRUM_ETH_SWAP_TOKEN,
  [ChainId.AVALANCHE]: AVALANCHE_AVETH_SWAP_TOKEN,
  [ChainId.AURORA]:    undefined,
  [ChainId.HARMONY]:   HARMONY_ONEETH_SWAP_TOKEN,
  [ChainId.CRONOS]:    undefined,                            // {poolTokens: [WETHBEAM]},
  [ChainId.METIS]:     METIS_WETH_SWAP_TOKEN,
  [ChainId.DOGECHAIN]: undefined,
  [ChainId.CANTO]:     CANTO_WETH_SWAP_TOKEN

}


export const BRIDGE_CHAINS_BY_TYPE = {
  USD:        [ChainId.ETH, ChainId.AVALANCHE, ChainId.ARBITRUM, ChainId.OPTIMISM, ChainId.BSC, ChainId.POLYGON, ChainId.FANTOM, ChainId.HARMONY, ChainId.BOBA, ChainId.AURORA, ChainId.METIS, ChainId.DFK, ChainId.CRONOS, ChainId.CANTO],
  ETH:        [ChainId.ETH, ChainId.AVALANCHE, ChainId.ARBITRUM, ChainId.OPTIMISM, ChainId.FANTOM, ChainId.BOBA, ChainId.HARMONY, ChainId.METIS, ChainId.KLAYTN, ChainId.DOGECHAIN, ChainId.CANTO],
  SYN:        [ChainId.ETH, ChainId.AVALANCHE, ChainId.ARBITRUM, ChainId.BSC, ChainId.POLYGON, ChainId.FANTOM, ChainId.HARMONY, ChainId.BOBA, ChainId.MOONRIVER, ChainId.OPTIMISM, ChainId.AURORA, ChainId.MOONBEAM, ChainId.CRONOS, ChainId.METIS, ChainId.DOGECHAIN],
  OHM:        [ChainId.ETH, ChainId.AVALANCHE, ChainId.ARBITRUM, ChainId.POLYGON, ChainId.FANTOM, ChainId.OPTIMISM, ChainId.MOONRIVER, ChainId.MOONBEAM, ChainId.CRONOS, ChainId.METIS],
  HIGHSTREET: [ChainId.ETH, ChainId.BSC],
  DOG:        [ChainId.ETH, ChainId.BSC, ChainId.POLYGON],
  JUMP:       [ChainId.FANTOM, ChainId.BSC, ChainId.METIS],
  SFI:        [ChainId.ETH, ChainId.AVALANCHE],
  FRAX:       [ChainId.ETH, ChainId.HARMONY, ChainId.MOONRIVER],
  NFD:        [ChainId.POLYGON, ChainId.BSC, ChainId.AVALANCHE],
  SOLAR:      [ChainId.MOONRIVER, ChainId.MOONBEAM],
  GMX:        [ChainId.ARBITRUM, ChainId.AVALANCHE],
  AVAX:       [ChainId.AVALANCHE, ChainId.DFK, ChainId.HARMONY, ChainId.MOONBEAM],
  MOVR:       [ChainId.MOONRIVER, ChainId.MOONBEAM],
  VSTA:       [ChainId.ETH, ChainId.ARBITRUM],
  H2O:        [ChainId.ETH, ChainId.ARBITRUM, ChainId.AVALANCHE, ChainId.BSC, ChainId.MOONBEAM, ChainId.MOONRIVER, ChainId.OPTIMISM, ChainId.POLYGON],
  UST:        [ChainId.TERRA, ChainId.ARBITRUM, ChainId.OPTIMISM, ChainId.METIS, ChainId.DFK],
  NEWO:       [ChainId.ETH, ChainId.AVALANCHE, ChainId.ARBITRUM],
  USDB:       [ChainId.ETH, ChainId.AVALANCHE, ChainId.BSC, ChainId.POLYGON, ChainId.FANTOM, ChainId.MOONRIVER],
  SDT:        [ChainId.ETH, ChainId.AVALANCHE, ChainId.ARBITRUM, ChainId.FANTOM, ChainId.HARMONY],
  JEWEL:      [ChainId.DFK, ChainId.AVALANCHE, ChainId.HARMONY],
  XJEWEL:     [ChainId.DFK, ChainId.HARMONY],
  KLAYTN_USDC: [ChainId.ETH, ChainId.KLAYTN, ChainId.DOGECHAIN],
  KLAYTN_USDT: [ChainId.ETH, ChainId.KLAYTN, ChainId.DOGECHAIN],
  KLAYTN_DAI:  [ChainId.ETH, ChainId.KLAYTN, ChainId.DOGECHAIN],
  WBTC:       [ChainId.ETH, ChainId.KLAYTN, ChainId.DOGECHAIN],
  L2DAO:      [ChainId.ARBITRUM, ChainId.OPTIMISM],
  PLS:        [ChainId.ARBITRUM, ChainId.OPTIMISM],
  AGEUR:      [ChainId.ETH, ChainId.ARBITRUM, ChainId.OPTIMISM],
  DOGECHAIN_BUSD: [ChainId.BSC, ChainId.DOGECHAIN],
  LINK:       [ChainId.ETH, ChainId.KLAYTN],
  UNIDX:      [ChainId.ETH, ChainId.ARBITRUM]
}

//   - Arbitrum
//   - Optimism
//   - Metis

// Add:
// - Aurora(https://aurorascan.dev/token/0xb1Da21B0531257a7E5aEfa0cd3CbF23AfC674cE1) - In BridgeConfig

//   - DFK Chain(https://subnets.avax.network/defi-kingdoms/dfk-chain/explorer/address/0x360d6DD540E3448371876662FBE7F1aCaf08c5Ab) - Adding now

let BRIDGE_SWAPABLE_TYPES_BY_CHAIN = {}
for (const [coinSymbol, chainIdArr] of _.entries(BRIDGE_CHAINS_BY_TYPE)) {
  for (const cid of chainIdArr) {
    if (BRIDGE_SWAPABLE_TYPES_BY_CHAIN[cid]) {
      BRIDGE_SWAPABLE_TYPES_BY_CHAIN[cid].push(coinSymbol)
    } else {
      BRIDGE_SWAPABLE_TYPES_BY_CHAIN[cid] = [coinSymbol]
    }
  }
}

export { BRIDGE_SWAPABLE_TYPES_BY_CHAIN }


function moveFirstToLast(arr) {
  return [
    ...arr.slice(1, arr.length),
    arr[0]
  ]
}



export const BRIDGE_SWAPABLE_TOKENS_BY_TYPE = {
  [ChainId.ETH]: {
    USD:        [...ETH_POOL_SWAP_TOKEN.poolTokens, NUSD],
    ETH:        [ETH],
    SYN:        [SYN],
    OHM:        [GOHM],
    HIGHSTREET: [HIGHSTREET],
    LINK:       [LINK],
    DOG:        [DOG],
    FRAX:       [FRAX],
    NEWO:       [NEWO],
    USDB:       [USDB],
    SDT:        [SDT],
    VSTA:       [VSTA],
    SFI:        [SFI],
    H2O:        [H2O],
    AGEUR:      [AGEUR],
    WBTC:       [WBTC],
    KLAYTN_USDC:[KLAYTN_USDC],
    KLAYTN_USDT:[KLAYTN_USDT],
    KLAYTN_DAI: [KLAYTN_DAI],
    UNIDX:      [UNIDX],
  },
  [ChainId.OPTIMISM]: {
    USD: moveFirstToLast(OPTIMISM_POOL_SWAP_TOKEN.poolTokens),
    ETH: moveFirstToLast(OPTIMISM_ETH_SWAP_TOKEN.nativeTokens),
    SYN: [SYN],
    H2O: [H2O],
    OHM: [GOHM],
    UST: [UST],
    H2O: [H2O],
    L2DAO:  [L2DAO],
    PLS:  [PLS],
    AGEUR:  [AGEUR],
  },
  [ChainId.BSC]: {
    USD:        moveFirstToLast(BSC_POOL_SWAP_TOKEN.poolTokens),
    SYN:        [SYN],
    OHM:        [GOHM],
    HIGHSTREET: [HIGHSTREET],
    DOG:        [DOG],
    JUMP:       [JUMP],
    NFD:        [NFD],
    USDB:       [USDB],
    H2O:        [H2O],
    DOGECHAIN_BUSD: [DOGECHAIN_BUSD],
  },
  [ChainId.POLYGON]: {
    USD:  moveFirstToLast(POLYGON_POOL_SWAP_TOKEN.poolTokens),
    SYN:  [SYN],
    OHM:  [GOHM],
    DOG:  [DOG],
    NFD:  [NFD],
    USDB: [USDB],
    H2O:  [H2O],
  },
  [ChainId.FANTOM]: {
    USD:  moveFirstToLast(FANTOM_3POOL_SWAP_TOKEN .poolTokens),
    ETH:  moveFirstToLast(FANTOM_WETH_SWAP_TOKEN.poolTokens),
    SYN:  [SYN],
    OHM:  [GOHM],
    JUMP: [JUMP],
    SDT:  [SDT],
    USDB: [USDB],
    // FRAX: [FRAX],
  },
  [ChainId.BOBA]: {
    USD: moveFirstToLast(BOBA_POOL_SWAP_TOKEN.poolTokens),
    ETH: moveFirstToLast(BOBA_ETH_SWAP_TOKEN.poolTokens),
    SYN: [SYN],
    OHM: [GOHM],
    // UST: [UST],
  },
  [ChainId.MOONBEAM]: {
    // ETH:   [WETHBEAM], #temp
    // FRAX:  [FRAX], #temp
    AVAX:  [WAVAX],
    MOVR:  [WMOVR],
    SOLAR: [SOLAR],
    SYN:   [SYN],
    OHM:   [GOHM],
    H2O:   [H2O],
  },
  [ChainId.MOONRIVER]: {
    MOVR:  [MOVR],
    FRAX:  [FRAX],
    SOLAR: [SOLAR],
    SYN:   [SYN],
    OHM:   [GOHM],
    USDB:  [USDB],
    H2O:   [H2O],
  },
  [ChainId.ARBITRUM]: {
    USD:  moveFirstToLast(ARBITRUM_3POOL_SWAP_TOKEN.poolTokens),
    ETH:  moveFirstToLast(ARBITRUM_ETH_SWAP_TOKEN.nativeTokens),
    SYN:  [SYN],
    SDT:  [SDT],
    OHM:  [GOHM],
    GMX:  [GMX],
    NEWO: [NEWO],
    VSTA: [VSTA],
    UST:  [UST],
    H2O:  [H2O],
    // DOG: [DOG],
    H2O:  [H2O],
    L2DAO:  [L2DAO],
    PLS:  [PLS],
    AGEUR:  [AGEUR],
    UNIDX:      [UNIDX],
  },
  [ChainId.AVALANCHE]: {
    USD:   moveFirstToLast(AVALANCHE_POOL_SWAP_TOKEN.poolTokens),
    ETH:   moveFirstToLast(AVALANCHE_AVETH_SWAP_TOKEN.depositTokens),
    SYN:   [SYN],
    NFD:   [NFD],
    H2O:   [H2O],
    OHM:   [GOHM],
    GMX:   [GMX],
    AVAX:  [AVAX],
    SDT:   [SDT],
    USDB:  [USDB],
    NEWO:  [NEWO],
    SFI:   [SFI],
    JEWEL: [SYNJEWEL],
    H2O:   [H2O],
  },
  [ChainId.HARMONY]: {
    USD:    [NUSD],
    ETH:    [NETH],
    FRAX:   [FRAX],
    SYN:    [SYN],
    OHM:    [GOHM],
    SDT:    [SDT],
    AVAX:   HARMONY_AVAX_SWAP_TOKEN.poolTokens,     //[SYNAVAX],
    JEWEL:  [WJEWEL],
    XJEWEL: [XJEWEL],
  },
  [ChainId.DFK]: {
    USD:    [DFK_USDC],
    JEWEL:  [JEWEL],
    XJEWEL: [XJEWEL],
    AVAX:   [WAVAX],
    UST:    [UST],
  },
  [ChainId.AURORA]: {
    USD: moveFirstToLast(AURORA_TS_POOL_SWAP_TOKEN.poolTokens),
    // UST: [UST],
    SYN: [SYN],
  },
  [ChainId.CRONOS]: {
    SYN: [SYN],
    OHM: [GOHM],
    // UST: [UST],
  },
  [ChainId.KLAYTN]: {
    KLAYTN_USDT:[KLAYTN_USDT],
    KLAYTN_USDC:[KLAYTN_USDC],
    KLAYTN_DAI: [KLAYTN_DAI],
    WBTC: [WBTC],
    ETH: [KLAYTN_WETH],
    LINK:       [LINK],
    // UST: [UST],
  },
  [ChainId.METIS]: {
    USD:  moveFirstToLast(METIS_POOL_SWAP_TOKEN.poolTokens),
    ETH:  moveFirstToLast(METIS_WETH_SWAP_TOKEN.poolTokens),
    SYN:  [SYN],
    OHM:  [GOHM],
    JUMP: [JUMP],
    UST:  [UST],
  },
  [ChainId.DOGECHAIN]: {
    SYN: [SYN],
    FRAX: [FRAX],
    KLAYTN_USDT:[KLAYTN_USDT],
    KLAYTN_USDC:[KLAYTN_USDC],
    KLAYTN_DAI: [KLAYTN_DAI],
    WBTC: [WBTC],
    ETH: [KLAYTN_WETH],
    DOGECHAIN_BUSD: [DOGECHAIN_BUSD],
  },
  [ChainId.CANTO]: {
    USD:  moveFirstToLast(CANTO_WRAPPER_POOL_SWAP_TOKEN.poolTokens),
    ETH:  moveFirstToLast(CANTO_WETH_SWAP_TOKEN.poolTokens),
    SYN: [SYN]
  },
  [ChainId.TERRA]: {
    UST: [UST]
  },
}

let BRIDGE_SWAPABLE_TOKENS_BY_CHAIN = {}
for (const [chainId, typeObj] of _.entries(BRIDGE_SWAPABLE_TOKENS_BY_TYPE)) {
  BRIDGE_SWAPABLE_TOKENS_BY_CHAIN[chainId] = _.flattenDeep(
    _.entries(typeObj).map( ([swapableType, poolTokens]) => poolTokens)
  )
}

export { BRIDGE_SWAPABLE_TOKENS_BY_CHAIN }

const AVAX_SPOOFED_SWAP_TOKEN  = { poolTokens: [AVAX] }
const WAVAX_SPOOFED_SWAP_TOKEN = { poolTokens: [WAVAX] }
const SYNAVAX_SPOOFED_SWAP_TOKEN = { poolTokens: [SYNAVAX]}

const JEWEL_SPOOFED_SWAP_TOKEN    = { poolTokens: [JEWEL]}
const WJEWEL_SPOOFED_SWAP_TOKEN   = { poolTokens: [WJEWEL]}
const XJEWEL_SPOOFED_SWAP_TOKEN   = { poolTokens: [XJEWEL]}
const SYNJEWEL_SPOOFED_SWAP_TOKEN = { poolTokens: [SYNJEWEL]}

const DFK_USDC_SPOOFED_SWAP_TOKEN = { poolTokens: [DFK_USDC] }

const MOVR_SPOOFED_SWAP_TOKEN  = { poolTokens: [MOVR] }
const WMOVR_SPOOFED_SWAP_TOKEN = { poolTokens: [WMOVR] }

const GMX_SPOOFED_SWAP_TOKEN        = { poolTokens: [GMX] }
const SOLAR_SPOOFED_SWAP_TOKEN      = { poolTokens: [SOLAR] }
const OHM_SPOOFED_SWAP_TOKEN        = { poolTokens: [GOHM] }
const NEWO_SPOOFED_SWAP_TOKEN       = { poolTokens: [NEWO] }
const VSTA_SPOOFED_SWAP_TOKEN       = { poolTokens: [VSTA] }
const SDT_SPOOFED_SWAP_TOKEN        = { poolTokens: [SDT] }
const USDB_SPOOFED_SWAP_TOKEN       = { poolTokens: [USDB] }
const NFD_SPOOFED_SWAP_TOKEN        = { poolTokens: [NFD] }
const FRAX_SPOOFED_SWAP_TOKEN       = { poolTokens: [FRAX] }
const JUMP_SPOOFED_SWAP_TOKEN       = { poolTokens: [JUMP] }
const SFI_SPOOFED_SWAP_TOKEN        = { poolTokens: [SFI] }
const HIGHSTREET_SPOOFED_SWAP_TOKEN = { poolTokens: [HIGHSTREET] }
const LINK_SPOOFED_SWAP_TOKEN       = { poolTokens: [LINK] }
const DOG_SPOOFED_SWAP_TOKEN        = { poolTokens: [DOG] }
const SYN_SPOOFED_SWAP_TOKEN        = { poolTokens: [SYN] }
const H2O_SPOOFED_SWAP_TOKEN        = { poolTokens: [H2O] }
const ETH_SPOOFED_SWAP_TOKEN        = { poolTokens: [WETH]}
const WETHBEAM_SPOOFED_SWAP_TOKEN   = { poolTokens: [WETHBEAM]}
const L2DAO_SPOOFED_SWAP_TOKEN      = { poolTokens: [L2DAO] }
const PLS_SPOOFED_SWAP_TOKEN        = { poolTokens: [PLS]}
const AGEUR_SPOOFED_SWAP_TOKEN      = { poolTokens: [AGEUR] }
const UNIDX_SPOOFED_SWAP_TOKEN      = { poolTokens: [UNIDX] }

const UST_SPOOFED_SWAP_TOKEN = { poolTokens: [UST] }


const KLAYTN_USDC_SPOOFED_SWAP_TOKEN        = { poolTokens: [KLAYTN_USDC]}
const KLAYTN_USDT_SPOOFED_SWAP_TOKEN        = { poolTokens: [KLAYTN_USDT]}
const KLAYTN_DAI_SPOOFED_SWAP_TOKEN        = { poolTokens: [KLAYTN_DAI]}
const KLAYTN_WETH_SPOOFED_SWAP_TOKEN        = { poolTokens: [KLAYTN_WETH]}
const WBTC_SPOOFED_SWAP_TOKEN        =       { poolTokens: [WBTC]}
const DOGECHAIN_BUSD_SPOOFED_SWAP_TOKEN      = { poolTokens: [DOGECHAIN_BUSD]}




export const BRIDGE_SWAPABLE_TYPE_POOLS_BY_CHAIN = {
  [ChainId.ETH]: {
    ETH:        ETH_SPOOFED_SWAP_TOKEN,
    USD:        ETH_POOL_SWAP_TOKEN,
    SYN:        SYN_SPOOFED_SWAP_TOKEN,
    OHM:        OHM_SPOOFED_SWAP_TOKEN,
    HIGHSTREET: HIGHSTREET_SPOOFED_SWAP_TOKEN,
    LINK:       LINK_SPOOFED_SWAP_TOKEN,
    DOG:        DOG_SPOOFED_SWAP_TOKEN,
    NEWO:       NEWO_SPOOFED_SWAP_TOKEN,
    SDT:        SDT_SPOOFED_SWAP_TOKEN,
    USDB:       USDB_SPOOFED_SWAP_TOKEN,
    FRAX:       FRAX_SPOOFED_SWAP_TOKEN,
    VSTA:       VSTA_SPOOFED_SWAP_TOKEN,
    SFI:        SFI_SPOOFED_SWAP_TOKEN,
    H2O:        H2O_SPOOFED_SWAP_TOKEN,
    AGEUR:      AGEUR_SPOOFED_SWAP_TOKEN,
    WBTC:       WBTC_SPOOFED_SWAP_TOKEN,
    KLAYTN_USDC: KLAYTN_USDC_SPOOFED_SWAP_TOKEN,
    KLAYTN_USDT: KLAYTN_USDT_SPOOFED_SWAP_TOKEN,
    KLAYTN_DAI: KLAYTN_DAI_SPOOFED_SWAP_TOKEN,
    UNIDX:  UNIDX_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.BSC]: {
    USD:        BSC_POOL_SWAP_TOKEN,
    SYN:        SYN_SPOOFED_SWAP_TOKEN,
    OHM:        OHM_SPOOFED_SWAP_TOKEN,
    HIGHSTREET: HIGHSTREET_SPOOFED_SWAP_TOKEN,
    DOG:        DOG_SPOOFED_SWAP_TOKEN,
    JUMP:       JUMP_SPOOFED_SWAP_TOKEN,
    NFD:        NFD_SPOOFED_SWAP_TOKEN,
    USDB:       USDB_SPOOFED_SWAP_TOKEN,
    H2O:        H2O_SPOOFED_SWAP_TOKEN,
    DOGECHAIN_BUSD: DOGECHAIN_BUSD_SPOOFED_SWAP_TOKEN
  },
  [ChainId.POLYGON]: {
    USD:  POLYGON_POOL_SWAP_TOKEN,
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
    OHM:  OHM_SPOOFED_SWAP_TOKEN,
    DOG:  DOG_SPOOFED_SWAP_TOKEN,
    NFD:  NFD_SPOOFED_SWAP_TOKEN,
    USDB: USDB_SPOOFED_SWAP_TOKEN,
    H2O:  H2O_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.FANTOM]: {
    ETH:  FANTOM_WETH_SWAP_TOKEN,
    USD:  FANTOM_3POOL_SWAP_TOKEN,
    // FRAX: FRAX_SPOOFED_SWAP_TOKEN,
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
    OHM:  OHM_SPOOFED_SWAP_TOKEN,
    JUMP: JUMP_SPOOFED_SWAP_TOKEN,
    SDT:  SDT_SPOOFED_SWAP_TOKEN,
    USDB: USDB_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.BOBA]: {
    ETH: BOBA_ETH_SWAP_TOKEN,
    USD: BOBA_POOL_SWAP_TOKEN,
    SYN: SYN_SPOOFED_SWAP_TOKEN,
    OHM: OHM_SPOOFED_SWAP_TOKEN,
    // UST: UST_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.MOONBEAM]: {
    // FRAX:  FRAX_SPOOFED_SWAP_TOKEN, #temp
    AVAX:  WAVAX_SPOOFED_SWAP_TOKEN,
    MOVR:  WMOVR_SPOOFED_SWAP_TOKEN,
    SYN:   SYN_SPOOFED_SWAP_TOKEN,
    SOLAR: SOLAR_SPOOFED_SWAP_TOKEN,
    OHM:   OHM_SPOOFED_SWAP_TOKEN,
    // ETH:   WETHBEAM_SPOOFED_SWAP_TOKEN #temp
    H2O:   H2O_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.MOONRIVER]: {
    MOVR:  MOVR_SPOOFED_SWAP_TOKEN,
    FRAX:  FRAX_SPOOFED_SWAP_TOKEN,
    SYN:   SYN_SPOOFED_SWAP_TOKEN,
    SOLAR: SOLAR_SPOOFED_SWAP_TOKEN,
    OHM:   OHM_SPOOFED_SWAP_TOKEN,
    USDB:  USDB_SPOOFED_SWAP_TOKEN,
    H2O:   H2O_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.ARBITRUM]: {
    ETH:  ARBITRUM_ETH_SWAP_TOKEN,
    USD:  ARBITRUM_3POOL_SWAP_TOKEN,
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
    OHM:  OHM_SPOOFED_SWAP_TOKEN,
    GMX:  GMX_SPOOFED_SWAP_TOKEN,
    SDT:  SDT_SPOOFED_SWAP_TOKEN,
    NEWO: NEWO_SPOOFED_SWAP_TOKEN,
    UST:  UST_SPOOFED_SWAP_TOKEN,
    VSTA: VSTA_SPOOFED_SWAP_TOKEN,
    H2O:  H2O_SPOOFED_SWAP_TOKEN,
    // DOG: DOG_SPOOFED_SWAP_TOKEN,
    L2DAO:  L2DAO_SPOOFED_SWAP_TOKEN,
    PLS:    PLS_SPOOFED_SWAP_TOKEN,
    UNIDX:  UNIDX_SPOOFED_SWAP_TOKEN,
    AGEUR:  AGEUR_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.AVALANCHE]: {
    USD:  AVALANCHE_POOL_SWAP_TOKEN,
    ETH:  AVALANCHE_AVETH_SWAP_TOKEN,
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
    NFD:  NFD_SPOOFED_SWAP_TOKEN,
    OHM:  OHM_SPOOFED_SWAP_TOKEN,
    NEWO: NEWO_SPOOFED_SWAP_TOKEN,
    SDT:  SDT_SPOOFED_SWAP_TOKEN,
    USDB: USDB_SPOOFED_SWAP_TOKEN,
    GMX:  GMX_SPOOFED_SWAP_TOKEN,
    AVAX: AVAX_SPOOFED_SWAP_TOKEN,
    SFI:  SFI_SPOOFED_SWAP_TOKEN,
    JEWEL: SYNJEWEL_SPOOFED_SWAP_TOKEN,
    H2O:   H2O_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.DFK]: {
    USD:    DFK_USDC_SPOOFED_SWAP_TOKEN,
    JEWEL:  JEWEL_SPOOFED_SWAP_TOKEN,
    XJEWEL: XJEWEL_SPOOFED_SWAP_TOKEN,
    AVAX:   WAVAX_SPOOFED_SWAP_TOKEN,
    UST:    UST_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.AURORA]: {
    USD: AURORA_TS_POOL_SWAP_TOKEN,
    // UST: UST_SPOOFED_SWAP_TOKEN,
    SYN: SYN_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.HARMONY]: {
    USD:  HARMONY_POOL_SWAP_TOKEN,
    ETH:  HARMONY_ONEETH_SWAP_TOKEN,
    FRAX: FRAX_SPOOFED_SWAP_TOKEN,
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
    OHM:  OHM_SPOOFED_SWAP_TOKEN,
    SDT:  SDT_SPOOFED_SWAP_TOKEN,

    AVAX:   HARMONY_AVAX_SWAP_TOKEN,
    JEWEL:  WJEWEL_SPOOFED_SWAP_TOKEN,
    XJEWEL: XJEWEL_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.OPTIMISM]: {
    USD: OPTIMISM_POOL_SWAP_TOKEN,
    ETH: OPTIMISM_ETH_SWAP_TOKEN,
    SYN: SYN_SPOOFED_SWAP_TOKEN,
    OHM: OHM_SPOOFED_SWAP_TOKEN,
    UST: UST_SPOOFED_SWAP_TOKEN,
    H2O: H2O_SPOOFED_SWAP_TOKEN,
    L2DAO:  L2DAO_SPOOFED_SWAP_TOKEN,
    PLS: PLS_SPOOFED_SWAP_TOKEN,
    AGEUR:  AGEUR_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.CRONOS]: {
    USD: CRONOS_POOL_SWAP_TOKEN,
    SYN: SYN_SPOOFED_SWAP_TOKEN,
    OHM: OHM_SPOOFED_SWAP_TOKEN,
    // UST: UST_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.KLAYTN]: {
    LINK:        LINK_SPOOFED_SWAP_TOKEN,
    KLAYTN_USDC: KLAYTN_USDC_SPOOFED_SWAP_TOKEN,
    KLAYTN_USDT: KLAYTN_USDT_SPOOFED_SWAP_TOKEN,
    KLAYTN_DAI: KLAYTN_DAI_SPOOFED_SWAP_TOKEN,
    WBTC: WBTC_SPOOFED_SWAP_TOKEN,
    ETH: KLAYTN_WETH_SPOOFED_SWAP_TOKEN
  },
  [ChainId.METIS]: {
    USD:  METIS_POOL_SWAP_TOKEN,
    ETH:  METIS_WETH_SWAP_TOKEN,
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
    OHM:  OHM_SPOOFED_SWAP_TOKEN,
    JUMP: JUMP_SPOOFED_SWAP_TOKEN,
    UST:  UST_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.CANTO]: {
    USD:  CANTO_WRAPPER_POOL_SWAP_TOKEN,
    ETH: CANTO_WETH_SWAP_TOKEN,
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.DOGECHAIN]: {
    SYN:  SYN_SPOOFED_SWAP_TOKEN,
    FRAX: FRAX_SPOOFED_SWAP_TOKEN,
    KLAYTN_USDC: KLAYTN_USDC_SPOOFED_SWAP_TOKEN,
    KLAYTN_USDT: KLAYTN_USDT_SPOOFED_SWAP_TOKEN,
    KLAYTN_DAI: KLAYTN_DAI_SPOOFED_SWAP_TOKEN,
    WBTC: WBTC_SPOOFED_SWAP_TOKEN,
    ETH: KLAYTN_WETH_SPOOFED_SWAP_TOKEN,
    DOGECHAIN_BUSD: DOGECHAIN_BUSD_SPOOFED_SWAP_TOKEN,
    NFD:  NFD_SPOOFED_SWAP_TOKEN,
  },
  [ChainId.TERRA]: {
    UST: UST_SPOOFED_SWAP_TOKEN
  },
}





export const MINT_BURN_TOKENS = [
  DFK_USDC,
  NUSD,
  SYN,
  NETH,
  WETHBEAM,
  GOHM,
  HIGHSTREET,
  LINK,
  DOG,
  JUMP,
  FRAX,
  NFD,
  SOLAR,
  GMX,
  AVAX,
  WAVAX,
  SYNAVAX,
  MOVR,
  WMOVR,
  UST,
  NEWO,
  VSTA,
  SDT,
  SFI,
  H2O,
  USDB,
  SYNJEWEL,
  WJEWEL,
  XJEWEL,
  JEWEL,
  KLAYTN_USDC,
  KLAYTN_DAI,
  KLAYTN_USDT,
  WBTC,
  L2DAO,
  PLS,
  UNIDX,
  AGEUR,
  DOGECHAIN_BUSD
]
