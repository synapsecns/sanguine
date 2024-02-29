import { USDC } from '@constants/tokens/bridgeable'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import * as CHAINS from '@constants/chains/master'

export const QUOTE_POLLING_INTERVAL = 10000

export const EMPTY_BRIDGE_QUOTE = {
  outputAmount: 0n,
  outputAmountString: '',
  routerAddress: '',
  allowance: 0n,
  exchangeRate: 0n,
  feeAmount: 0n,
  delta: 0n,
  originQuery: null,
  destQuery: null,
  estimatedTime: null,
  bridgeModuleName: null,
  gasDropAmount: 0n,
}

export const EMPTY_BRIDGE_QUOTE_ZERO = {
  outputAmount: 0n,
  outputAmountString: '0',
  routerAddress: '',
  allowance: 0n,
  exchangeRate: 0n,
  feeAmount: 0n,
  delta: 0n,
  originQuery: null,
  destQuery: null,
  estimatedTime: null,
  bridgeModuleName: null,
  gasDropAmount: 0n,
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

/**
 * number of required confirmations from bridge
 */
export const BRIDGE_REQUIRED_CONFIRMATIONS = {
  [CHAINS.ETH.id]: 33,
  [CHAINS.BNB.id]: 33,
  [CHAINS.POLYGON.id]: 128,
  [CHAINS.FANTOM.id]: 80,
  [CHAINS.BOBA.id]: 33,
  [CHAINS.OPTIMISM.id]: 750,
  [CHAINS.MOONBEAM.id]: 33,
  [CHAINS.MOONRIVER.id]: 33,
  [CHAINS.ARBITRUM.id]: 200,
  [CHAINS.AVALANCHE.id]: 80,
  [CHAINS.DFK.id]: 33,
  [CHAINS.HARMONY.id]: 33,
  [CHAINS.AURORA.id]: 33,
  [CHAINS.CRONOS.id]: 33,
  [CHAINS.METIS.id]: 33,
  [CHAINS.DOGE.id]: 33,
  [CHAINS.CANTO.id]: 20,
  [CHAINS.BASE.id]: 750,
  [CHAINS.KLAYTN.id]: 20,
}

export const DEFAULT_FROM_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_TO_TOKEN_SYMBOL = 'USDC'
export const DEFAULT_FROM_TOKEN = USDC
export const DEFAULT_TO_TOKEN = USDC
export const DEFAULT_SWAPABLE_TYPE = 'USD'
export const DEFAULT_FROM_CHAIN = CHAINS.ETH.id
export const DEFAULT_TO_CHAIN = CHAINS.ARBITRUM.id

export const TRANSITIONS_PROPS = {
  ...COIN_SLIDE_OVER_PROPS,
  className: `
    origin-bottom absolute
    w-full h-full
    md:w-[95%] md:h-[95%]
    -ml-0 md:-ml-3
    md:mt-3
    bg-bgBase
    z-20 rounded-lg
  `,
}

export const SETTINGS_TRANSITIONS_PROPS = {
  ...COIN_SLIDE_OVER_PROPS,
  className: `
  origin-bottom absolute
  w-full h-full
  md:w-[95%]
  -ml-0 md:-ml-3
  md:-mt-3
  bg-bgBase
  z-20 rounded-lg
  `,
}

export const BRIDGE_CONTRACTS = {
  [CHAINS.ETH.id]: '0x2796317b0fF8538F253012862c06787Adfb8cEb6',
  [CHAINS.OPTIMISM.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.CRONOS.id]: '0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9',
  [CHAINS.BNB.id]: '0xd123f70AE324d34A9E76b67a27bf77593bA8749f',
  [CHAINS.POLYGON.id]: '0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280',
  [CHAINS.FANTOM.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.BOBA.id]: '0x432036208d2717394d2614d6697c46DF3Ed69540',
  [CHAINS.METIS.id]: '0x06Fea8513FF03a0d3f61324da709D4cf06F42A5c',
  [CHAINS.MOONBEAM.id]: '0x84A420459cd31C3c34583F67E0f0fB191067D32f',
  [CHAINS.MOONRIVER.id]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [CHAINS.KLAYTN.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.ARBITRUM.id]: '0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9',
  [CHAINS.AVALANCHE.id]: '0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE',
  [CHAINS.DFK.id]: '0xE05c976d3f045D0E6E7A6f61083d98A15603cF6A',
  [CHAINS.AURORA.id]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [CHAINS.HARMONY.id]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [CHAINS.CANTO.id]: '0xDde5BEC4815E1CeCf336fb973Ca578e8D83606E0',
  [CHAINS.DOGE.id]: '0x9508BF380c1e6f751D97604732eF1Bae6673f299',
  [CHAINS.BASE.id]: '0xf07d1C752fAb503E47FEF309bf14fbDD3E867089',
}

export const SYNAPSE_CCTP_CONTRACTS = {
  [CHAINS.ETH.id]: '0xfB2Bfc368a7edfD51aa2cbEC513ad50edEa74E84',
  [CHAINS.ARBITRUM.id]: '0xfB2Bfc368a7edfD51aa2cbEC513ad50edEa74E84',
  [CHAINS.AVALANCHE.id]: '0xfB2Bfc368a7edfD51aa2cbEC513ad50edEa74E84',
  [CHAINS.BASE.id]: '0xfB2Bfc368a7edfD51aa2cbEC513ad50edEa74E84',
}

export const ROLE_EVENTS = ['RoleGranted', 'RoleRevoked', 'RoleAdminChanged']

export const INCLUDED_BRIDGE_EVENTS = [
  'TokenMintAndSwap',
  'TokenRedeem',
  'TokenRedeemV2',
  'TokenRedeemAndSwap',
  'TokenRedeemAndRemove',
  'TokenDepositAndSwap',
  'TokenDeposit',
  'TokenMint',
  'TokenWithdraw',
  'TokenWithdrawAndRemove',
]
