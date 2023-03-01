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
  UST,
  METISETH,
  CANTOETH,
  SYNAVAX,
  JEWEL,
  WJEWEL,
  XJEWEL,
  SYNJEWEL,
  DFK_USDC,
  MULTIAVAX,
  WETH,
  WBTC,
  KLAYTN_USDC,
  KLAYTN_USDT,
  KLAYTN_DAI,

  KLAYTN_WETH,
  
  DOGECHAIN_BUSD,

  NOTE
} from '@constants/tokens/basic'

import { DOG, LINK, HIGHSTREET, JUMP, NFD, GOHM, SOLAR, NEWO, SDT, USDB, GMX, VSTA, SFI, H2O, L2DAO, PLS, AGEUR, UNIDX  } from '@constants/tokens/mintable'

import { ChainId } from '@constants/networks'

/**
 * Underlying bridge addresses utilized by zaps
 *
 * abi specified by {@link `@abis/synapseBridge.json`}
 */
export const SYNAPSE_BRIDGE_ADDRESSES = {
  [ChainId.ETH]:       '0x2796317b0fF8538F253012862c06787Adfb8cEb6',
  [ChainId.BSC]:       '0xd123f70AE324d34A9E76b67a27bf77593bA8749f',
  [ChainId.POLYGON]:   '0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280',     // test--> '0x5F06745ee8a2001198a379BAfBd0361475F3cFc3'   prod--> ,
  [ChainId.FANTOM]:    '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.BOBA]:      '0x432036208d2717394d2614d6697c46DF3Ed69540',
  [ChainId.ARBITRUM]:  '0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9',
  [ChainId.AVALANCHE]: '0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE',
  [ChainId.DFK]:       '0xE05c976d3f045D0E6E7A6f61083d98A15603cF6A',
  [ChainId.HARMONY]:   '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.MOONRIVER]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [ChainId.OPTIMISM]:  '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.AURORA]:    '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [ChainId.MOONBEAM]:  '0x84A420459cd31C3c34583F67E0f0fB191067D32f',
  [ChainId.CRONOS]:    '0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9',
  [ChainId.METIS]:     '0x06Fea8513FF03a0d3f61324da709D4cf06F42A5c',
  [ChainId.KLAYTN]:    '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.DOGECHAIN]: '0x9508BF380c1e6f751D97604732eF1Bae6673f299',
  [ChainId.CANTO]:     '0xDde5BEC4815E1CeCf336fb973Ca578e8D83606E0',
  [ChainId.TERRA]:     'terra1qwzdua7928ugklpytdzhua92gnkxp9z4vhelq8',
}

/**
 * ETH Only Bridge Config used to calculate swap fees
 *
 * abi specified by {@link `@abis/bridgeConfig.json`}
 */
export const BRIDGE_CONFIG_ADDRESSES = {
  [ChainId.ETH]:     '0x5217c83ca75559B1f8a8803824E5b7ac233A12a1',
  [ChainId.POLYGON]: '0xd69229f223a8fc84998e1361ae7b4ff724cf4a49',   // TESTING ADDRESS
}
// '0x3ee02f08B801B1990AC844d8CD2F119BA6Fb9bcF', // old config addr => '0x7fd806049608b7d04076b8187dd773343e0589e6',

/**
 * abi for ETH specified by {@link `@abis/bridgeZap.json`}
 * Handles stables -> add liquidity, get nerveUSD-LP -> Mint nUSD -> Swap to stables on non-ETH chain
 * abi for others specified by {@link `@abis/l2bridgezap.json`}
 * Handles stables -> Swap to nUSD -> Redeem for nerveUSD-LP -> Remove liquidity
 */
export const BRIDGE_ZAP_ADDRESSES = {
  [ChainId.ETH]:       '0x6571d6be3d8460CF5F7d6711Cd9961860029D85F',
  [ChainId.BSC]:       '0x749F37Df06A99D6A8E065dd065f8cF947ca23697',
  [ChainId.POLYGON]:   '0x1c6aE197fF4BF7BA96c66C5FD64Cb22450aF9cC8',
  [ChainId.FANTOM]:    '0xB003e75f7E0B5365e814302192E99b4EE08c0DEd',
  [ChainId.BOBA]:      '0x64B4097bCCD27D49BC2A081984C39C3EeC427a2d',
  [ChainId.ARBITRUM]:  '0x37f9aE2e0Ea6742b9CAD5AbCfB6bBC3475b3862B',
  [ChainId.AVALANCHE]: '0x0EF812f4c68DC84c22A4821EF30ba2ffAB9C2f3A',
  [ChainId.DFK]:       '0x75224b0f245Fe51d5bf47A898DbB6720D4150BA7', // '0x33d90B6ce7e0bFC42BCD35d05c443c6915296987',
  [ChainId.HARMONY]:   '0xB003e75f7E0B5365e814302192E99b4EE08c0DEd',
  [ChainId.MOONRIVER]: '0xfA28DdB74b08B2b6430f5F61A1Dd5104268CC29e',
  [ChainId.OPTIMISM]:  '0x470f9522ff620eE45DF86C58E54E6A645fE3b4A7',
  [ChainId.AURORA]:    '0xEbE30C0ADc8970344a9Ed4C8B2b3F6Ec3c9759d0', // '0x2D8Ee8d6951cB4Eecfe4a79eb9C2F973C02596Ed',
  [ChainId.MOONBEAM]:  '0xadA10A7474f4c71A829b55D2cB4232C281383fd5',
  [ChainId.CRONOS]:    '0x991adb00eF4c4a6D1eA6036811138Db4379377C2',
  [ChainId.METIS]:     '0x6571D58b3BF2469DF5878e213453E28dC1A4DA81',
  [ChainId.KLAYTN]:    '0x911766fA1a425Cb7cCCB0377BC152f37F276f8d6',
  [ChainId.DOGECHAIN]: '0x544450Ffdfa5EA20528F21918E8aAC7B2C733381',
  [ChainId.CANTO]:     '0xDe4da25B3e3FCA88EeA04d9C0012Fd9a71E18a5D'
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
  [ChainId.ETH]:       [USDC, USDT, DAI, ETH, SYN, WBTC, LINK, AGEUR, DOG, FRAX, GOHM, H2O, HIGHSTREET, NEWO, SDT, SFI, USDB, UNIDX, VSTA, NUSD],
  [ChainId.BSC]:       [BUSD, USDC, USDT, SYN, DOG, GOHM, H2O, HIGHSTREET, JUMP, NFD, USDB, NUSD],
  [ChainId.FANTOM]:    [USDC, USDT, FTMETH, SYN, GOHM, JUMP, SDT, USDB, NETH, NUSD],                                            //  FRAX,
  [ChainId.POLYGON]:   [USDC, USDT, DAI, SYN, DOG, GOHM, H2O, NFD, USDB, NUSD],
  [ChainId.BOBA]:      [USDC, USDT, DAI, ETH, SYN, GOHM, NETH, NUSD],
  [ChainId.MOONBEAM]:  [SYN, WMOVR, WAVAX, GOHM, H2O, SOLAR],                                                                   // FRAX, , WETHBEAM #temp
  [ChainId.MOONRIVER]: [SYN, MOVR, FRAX, GOHM, H2O, SOLAR, USDB],
  [ChainId.ARBITRUM]:  [USDC, USDT, ETH, SYN, AGEUR, GOHM, GMX, H2O, L2DAO, PLS, NEWO, SDT, UST, VSTA, UNIDX, NETH, NUSD],
  [ChainId.AVALANCHE]: [USDC, USDT, DAI, WETHE, SYN, AVAX, SYNJEWEL, GMX, GOHM, H2O, NEWO, NFD, SFI, SDT, USDB, NETH, NUSD],
  [ChainId.DFK]:       [DFK_USDC, JEWEL, XJEWEL, WAVAX, UST],
  [ChainId.HARMONY]:   [SYN, WJEWEL, XJEWEL, SYNAVAX, MULTIAVAX, FRAX, GOHM, SDT, NETH, NUSD],
  [ChainId.OPTIMISM]:  [USDC, ETH, SYN, AGEUR, GOHM, H2O, L2DAO, PLS, UST, NETH, NUSD],
  [ChainId.AURORA]:    [USDC, USDT, SYN, NUSD],
  [ChainId.CRONOS]:    [USDC, SYN, GOHM, NUSD],                                                                                 // USDT, DAI,
  [ChainId.METIS]:     [USDC, METISETH, SYN, GOHM, JUMP, UST, NETH, NUSD],
  [ChainId.KLAYTN]:    [KLAYTN_USDC, KLAYTN_USDT, KLAYTN_DAI, KLAYTN_WETH, LINK, WBTC],
  [ChainId.DOGECHAIN]: [KLAYTN_WETH, KLAYTN_USDC, KLAYTN_USDT, KLAYTN_DAI, FRAX, DOGECHAIN_BUSD, WBTC, SYN, NFD],
  [ChainId.CANTO]:     [USDC, USDT, CANTOETH, NOTE, NETH, NUSD, SYN],                                                                  
  [ChainId.TERRA]:     [UST],
}
// metis, cronos, arbitrum, optimism, boba

/**
 * number of required confirmations from bridge
 */
export const BRIDGE_REQUIRED_CONFIRMATIONS = {
  [ChainId.ETH]:       33,
  [ChainId.BSC]:       14,
  [ChainId.POLYGON]:   128,
  [ChainId.FANTOM]:    5,
  [ChainId.BOBA]:      1,     // rewrite
  [ChainId.OPTIMISM]:  1,     // rewrite 
  [ChainId.MOONBEAM]:  21,
  [ChainId.MOONRIVER]: 21,    // 5,
  [ChainId.ARBITRUM]:  40,
  [ChainId.AVALANCHE]: 5,
  [ChainId.DFK]:       6,
  [ChainId.HARMONY]:   1,     // rewrite
  [ChainId.AURORA]:    5,
  [ChainId.CRONOS]:    6,
  [ChainId.METIS]:     6,
  [ChainId.TERRA]:     1,
  [ChainId.DOGECHAIN]: 20,
  [ChainId.CANTO]:     20,
}

