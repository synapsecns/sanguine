import {
  BUSD,
  USDC,
  USDT,
  DAI,
  NUSD,
  NETH,
  ETH,
  SYN,
  MIM,
  FRAX,
  WETHE,
  ONEETH,
  WETHBEAM,
  MOVR,
  AVAX,
  WAVAX,
  WMOVR,
  FTMETH,
  METISETH,
  CANTOETH,
  SYNAVAX,
  WJEWEL,
  XJEWEL,
  SYNJEWEL,
  DFK_USDC,
  MULTIAVAX,
  WETH,
  WBTC,
  KLAYTN_WETH,
  NOTE,
  DOG,
  LINK,
  HIGHSTREET,
  JUMP,
  NFD,
  GOHM,
  SOLAR,
  NEWO,
  SDT,
  USDB,
  GMX,
  VSTA,
  SFI,
  H2O,
  L2DAO,
  PLS,
  AGEUR,
  UNIDX,
} from '@constants/tokens/master'

import * as CHAINS from '@constants/chains/master'

/**
 * Underlying bridge addresses utilized by zaps
 *
 * abi specified by {@link `@abis/synapseBridge.json`}
 */
export const SYNAPSE_BRIDGE_ADDRESSES = {
  [CHAINS.ETH.id]: '0x2796317b0fF8538F253012862c06787Adfb8cEb6',
  [CHAINS.BNB.id]: '0xd123f70AE324d34A9E76b67a27bf77593bA8749f',
  [CHAINS.POLYGON.id]: '0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280', // test--> '0x5F06745ee8a2001198a379BAfBd0361475F3cFc3'   prod--> ,
  [CHAINS.FANTOM.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.BOBA.id]: '0x432036208d2717394d2614d6697c46DF3Ed69540',
  [CHAINS.ARBITRUM.id]: '0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9',
  [CHAINS.AVALANCHE.id]: '0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE',
  [CHAINS.DFK.id]: '0xE05c976d3f045D0E6E7A6f61083d98A15603cF6A',
  [CHAINS.HARMONY.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.MOONRIVER.id]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [CHAINS.OPTIMISM.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.AURORA.id]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [CHAINS.MOONBEAM.id]: '0x84A420459cd31C3c34583F67E0f0fB191067D32f',
  [CHAINS.CRONOS.id]: '0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9',
  [CHAINS.METIS.id]: '0x06Fea8513FF03a0d3f61324da709D4cf06F42A5c',
  [CHAINS.KLAYTN.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.DOGE.id]: '0x9508BF380c1e6f751D97604732eF1Bae6673f299',
  [CHAINS.CANTO.id]: '0xDde5BEC4815E1CeCf336fb973Ca578e8D83606E0',
  // [ChainId.TERRA]: 'terra1qwzdua7928ugklpytdzhua92gnkxp9z4vhelq8',
}

/**
 * ETH Only Bridge Config used to calculate swap fees
 *
 * abi specified by {@link `@abis/bridgeConfig.json`}
 */
export const BRIDGE_CONFIG_ADDRESSES = {
  [CHAINS.ETH.id]: '0x5217c83ca75559B1f8a8803824E5b7ac233A12a1',
  [CHAINS.POLYGON.id]: '0xd69229f223a8fc84998e1361ae7b4ff724cf4a49', // TESTING ADDRESS
}
// '0x3ee02f08B801B1990AC844d8CD2F119BA6Fb9bcF', // old config addr => '0x7fd806049608b7d04076b8187dd773343e0589e6',

/**
 * abi for ETH specified by {@link `@abis/bridgeZap.json`}
 * Handles stables -> add liquidity, get nerveUSD-LP -> Mint nUSD -> Swap to stables on non-ETH chain
 * abi for others specified by {@link `@abis/l2bridgezap.json`}
 * Handles stables -> Swap to nUSD -> Redeem for nerveUSD-LP -> Remove liquidity
 */
export const BRIDGE_ZAP_ADDRESSES = {
  [CHAINS.ETH.id]: '0x6571d6be3d8460CF5F7d6711Cd9961860029D85F',
  [CHAINS.BNB.id]: '0x749F37Df06A99D6A8E065dd065f8cF947ca23697',
  [CHAINS.POLYGON.id]: '0x1c6aE197fF4BF7BA96c66C5FD64Cb22450aF9cC8',
  [CHAINS.FANTOM.id]: '0xB003e75f7E0B5365e814302192E99b4EE08c0DEd',
  [CHAINS.BOBA.id]: '0x64B4097bCCD27D49BC2A081984C39C3EeC427a2d',
  [CHAINS.ARBITRUM.id]: '0x37f9aE2e0Ea6742b9CAD5AbCfB6bBC3475b3862B',
  [CHAINS.AVALANCHE.id]: '0x0EF812f4c68DC84c22A4821EF30ba2ffAB9C2f3A',
  [CHAINS.DFK.id]: '0x75224b0f245Fe51d5bf47A898DbB6720D4150BA7', // '0x33d90B6ce7e0bFC42BCD35d05c443c6915296987',
  [CHAINS.HARMONY.id]: '0xB003e75f7E0B5365e814302192E99b4EE08c0DEd',
  [CHAINS.MOONRIVER.id]: '0xfA28DdB74b08B2b6430f5F61A1Dd5104268CC29e',
  [CHAINS.OPTIMISM.id]: '0x470f9522ff620eE45DF86C58E54E6A645fE3b4A7',
  [CHAINS.AURORA.id]: '0xEbE30C0ADc8970344a9Ed4C8B2b3F6Ec3c9759d0', // '0x2D8Ee8d6951cB4Eecfe4a79eb9C2F973C02596Ed',
  [CHAINS.MOONBEAM.id]: '0xadA10A7474f4c71A829b55D2cB4232C281383fd5',
  [CHAINS.CRONOS.id]: '0x991adb00eF4c4a6D1eA6036811138Db4379377C2',
  [CHAINS.METIS.id]: '0x6571D58b3BF2469DF5878e213453E28dC1A4DA81',
  [CHAINS.KLAYTN.id]: '0x911766fA1a425Cb7cCCB0377BC152f37F276f8d6',
  [CHAINS.DOGE.id]: '0x544450Ffdfa5EA20528F21918E8aAC7B2C733381',
  [CHAINS.CANTO.id]: '0xDe4da25B3e3FCA88EeA04d9C0012Fd9a71E18a5D',
}

/**
 * Handles getting which tokens are shown on bridge card dropdown.
 * Ordering:
 * 1. Stables from nUSD pool: (BUSD) USDC (USDT) (DAI)
 * 2. ETH, SYN
 * 3. Other chains gas tokens in order of usage (incl WBTC)
 * 4. Partner tokens alphabetically
 * 5. nETH, nUSD
 */
// TODO: Revert LINK back to alphabetical post launch weeks.
export const BRIDGABLE_TOKENS = {
  [CHAINS.ETH.id]: [
    USDC,
    USDT,
    DAI,
    ETH,
    SYN,
    WBTC,
    LINK,
    AGEUR,
    DOG,
    FRAX,
    GOHM,
    H2O,
    HIGHSTREET,
    NEWO,
    SDT,
    SFI,
    USDB,
    UNIDX,
    VSTA,
    NUSD,
  ],
  [CHAINS.BNB.id]: [
    BUSD,
    USDC,
    USDT,
    SYN,
    DOG,
    GOHM,
    H2O,
    HIGHSTREET,
    JUMP,
    NFD,
    USDB,
    NUSD,
  ],
  [CHAINS.FANTOM.id]: [
    USDC,
    USDT,
    FTMETH,
    SYN,
    GOHM,
    JUMP,
    SDT,
    USDB,
    NETH,
    NUSD,
  ], //  FRAX,
  [CHAINS.POLYGON.id]: [USDC, USDT, DAI, SYN, DOG, GOHM, H2O, NFD, USDB, NUSD],
  [CHAINS.BOBA.id]: [USDC, USDT, DAI, ETH, SYN, GOHM, NETH, NUSD],
  [CHAINS.MOONBEAM.id]: [SYN, WMOVR, WAVAX, GOHM, H2O, SOLAR], // FRAX, , WETHBEAM #temp
  [CHAINS.MOONRIVER.id]: [SYN, MOVR, FRAX, GOHM, H2O, SOLAR, USDB],
  [CHAINS.ARBITRUM.id]: [
    USDC,
    USDT,
    ETH,
    SYN,
    AGEUR,
    GOHM,
    GMX,
    H2O,
    L2DAO,
    PLS,
    NEWO,
    SDT,
    VSTA,
    UNIDX,
    NETH,
    NUSD,
  ],
  [CHAINS.AVALANCHE.id]: [
    USDC,
    USDT,
    DAI,
    WETHE,
    SYN,
    AVAX,
    SYNJEWEL,
    GMX,
    GOHM,
    H2O,
    NEWO,
    NFD,
    SFI,
    SDT,
    USDB,
    NETH,
    NUSD,
  ],
  [CHAINS.DFK.id]: [DFK_USDC, XJEWEL, WAVAX],
  [CHAINS.HARMONY.id]: [
    SYN,
    WJEWEL,
    XJEWEL,
    SYNAVAX,
    MULTIAVAX,
    FRAX,
    GOHM,
    SDT,
    NETH,
    NUSD,
  ],
  [CHAINS.OPTIMISM.id]: [
    USDC,
    ETH,
    SYN,
    AGEUR,
    GOHM,
    H2O,
    L2DAO,
    PLS,
    NETH,
    NUSD,
  ],
  [CHAINS.AURORA.id]: [USDC, USDT, SYN, NUSD],
  [CHAINS.CRONOS.id]: [USDC, SYN, GOHM, NUSD], // USDT, DAI,
  [CHAINS.METIS.id]: [USDC, METISETH, SYN, GOHM, JUMP, NETH, NUSD],
  [CHAINS.KLAYTN.id]: [KLAYTN_WETH, LINK, WBTC],
  [CHAINS.DOGE.id]: [KLAYTN_WETH, FRAX, WBTC, SYN, NFD],
  [CHAINS.CANTO.id]: [USDC, USDT, CANTOETH, NOTE, NETH, NUSD, SYN],
  // [ChainId.TERRA]: [UST],
}
// metis, cronos, arbitrum, optimism, boba

/**
 * number of required confirmations from bridge
 */
export const BRIDGE_REQUIRED_CONFIRMATIONS = {
  [CHAINS.ETH.id]: 33,
  [CHAINS.BNB.id]: 14,
  [CHAINS.POLYGON.id]: 128,
  [CHAINS.FANTOM.id]: 5,
  [CHAINS.BOBA.id]: 1, // rewrite
  [CHAINS.OPTIMISM.id]: 1, // rewrite
  [CHAINS.MOONBEAM.id]: 21,
  [CHAINS.MOONRIVER.id]: 21, // 5,
  [CHAINS.ARBITRUM.id]: 40,
  [CHAINS.AVALANCHE.id]: 5,
  [CHAINS.DFK.id]: 6,
  [CHAINS.HARMONY.id]: 1, // rewrite
  [CHAINS.AURORA.id]: 5,
  [CHAINS.CRONOS.id]: 6,
  [CHAINS.METIS.id]: 6,
  // [ChainId.TERRA]: 1,
  [CHAINS.DOGE.id]: 20,
  [CHAINS.CANTO.id]: 20,
}

export const DEFAULT_FROM_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_TO_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_FROM_TOKEN = USDC
export const DEFAULT_TO_TOKEN = USDC
export const DEFAULT_SWAPABLE_TYPE = 'USD'
export const DEFAULT_FROM_CHAIN = CHAINS.ETH.id
export const DEFAULT_TO_CHAIN = CHAINS.ARBITRUM.id
