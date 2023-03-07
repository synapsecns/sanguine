import synapseLogo from '@assets/icons/synapse.svg'

import { Token } from '@utils/classes/Token'
import { SYNAPSE_DOCS_URL } from '@/constants/urls'
import { ChainId } from '@constants/networks'

import {
  BUSD,
  USDC,
  USDT,
  DAI,
  MIM,
  NUSD,
  KLAYTN_USDT,
  KLAYTN_oUSDT,
  NOTE
} from '@constants/tokens/basic'

/**
 * Eth Stablecoin Swap
 */
export const ETH_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.ETH]: NUSD.addresses[ChainId.ETH],
  },
  decimals: 18,
  symbol: 'nUSD',
  name: 'Synapse nUSD LP Token Ethereum',
  logo: synapseLogo,
  poolName: 'Ethereum Stableswap Pool',
  routerIndex: 'eth3pool',
  poolId: 420,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.ETH]: '0x1116898DdA4015eD8dDefb84b6e8Bc24528Af2d8',
  },
  poolTokens: [DAI, USDC, USDT],
  description: "Synapse's 3pool stableswap LP token on ETH",
  docUrl: SYNAPSE_DOCS_URL,
})

// Stablecoin Swap
export const BSC_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.BSC]: '0xa4b7Bc06EC817785170C2DbC1dD3ff86CDcdcc4C',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token',
  logo: synapseLogo,
  poolName: 'BSC Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'bscnusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.BSC]: '0x28ec0B36F0819ecB5005cAB836F4ED5a2eCa4D13',
  },
  poolTokens: [NUSD, BUSD, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Optimism Stablecoin Swap
 */
export const OPTIMISM_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.OPTIMISM]: '0x2c6d91accC5Aa38c84653F28A80AEC69325BDd12',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Optimism ',
  logo: synapseLogo,
  poolName: 'Optimism Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'optimismnusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.OPTIMISM]: '0xF44938b0125A6662f9536281aD2CD6c499F22004',
  },
  poolTokens: [NUSD, USDC],
  description: "Synapse's 2pool stableswap LP token on Optimism",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Optimism Stablecoin Swap
 */
export const CRONOS_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.CRONOS]: '0xd5Da35646A79C42a0BAB148656192A22e8CBfad6', //'0x8415D4EB17F0949e2388fdF52909db4cC0a2B082',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Cronos ',
  logo: synapseLogo,
  poolName: 'Cronos Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'cronosnusd',
  poolId: 0,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.CRONOS]: '0xCb6674548586F20ca39C97A52A0ded86f48814De', //'0x3b1601c386247A127287b094F9CCB585D4D0B99b',
  },
  poolTokens: [NUSD, USDC], // [NUSD, DAI, USDC, USDT],
  description: "Synapse's 2pool stableswap LP token on Cronos",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Polygon Stablecoin Swap
 */
export const POLYGON_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.POLYGON]: '0x7479e1bc2f2473f9e78c89b4210eb6d55d33b645',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Polygon ',
  logo: synapseLogo,
  poolName: 'Polygon Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'polygonnusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.POLYGON]: '0x85fCD7Dd0a1e1A9FCD5FD886ED522dE8221C3EE5',
  },
  poolTokens: [NUSD, DAI, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Polygon/Matic",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Avalanche Stablecoin Swap
 */
export const AVALANCHE_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0xCA87BF3ec55372D9540437d7a86a7750B42C02f4',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Avalanche',
  logo: synapseLogo,
  poolName: 'Avalanche Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'avalanchenusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.AVALANCHE]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
  },
  poolTokens: [NUSD, DAI, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Avalanche",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Avalanche Stablecoin Swap
 */
export const ARBITRUM_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.ARBITRUM]: '0xADeac0343C2Ac62DFE5A5f51E896AefFF5Ab513E',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Arbitrum',
  logo: synapseLogo,
  poolName: 'Legacy Arbitrum Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'arbitrumnusd',
  poolId: 2,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.ARBITRUM]: '0x0Db3FE3B770c95A0B99D1Ed6F2627933466c0Dd8',
  },
  poolTokens: [NUSD, MIM, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Arbitrum",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Fantom Stablecoin Swap
 */
export const FANTOM_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.FANTOM]: '0x464d121D3cA63cEEfd390D76f19364D3Bd024cD2',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Fantom',
  logo: synapseLogo,
  poolName: 'Legacy Fantom Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'fantomnusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.FANTOM]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
  },
  poolTokens: [NUSD, MIM, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Fantom",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Harmony Stablecoin Swap
 */
export const HARMONY_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.HARMONY]: '0xE269abBFAF52b26D2632F55B6b223A5223088B96',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Harmony',
  logo: synapseLogo,
  poolName: 'Harmony Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'harmonynusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.HARMONY]: '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e',
  },
  poolTokens: [NUSD, DAI, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Harmony",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Boba Stablecoin Swap
 */
export const BOBA_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.BOBA]: '0x9D7283A6AeeD9BCd4Ac70876fEA2b69a63DD8cb9',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Boba',
  logo: synapseLogo,
  poolName: 'Boba Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'bobanusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.BOBA]: '0x75FF037256b36F15919369AC58695550bE72fead',
  },
  poolTokens: [NUSD, DAI, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Boba",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Aurora Stablecoin Swap
 */
export const AURORA_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.AURORA]: '0xEAdC3524f3F007cdC5104BF28663b1141D3e3127',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Aurora',
  logo: synapseLogo,
  poolName: 'Legacy Aurora Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'auroranusd',
  poolId: 0,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.AURORA]: '0xcEf6C2e20898C2604886b888552CA6CcF66933B0',
  },
  poolTokens: [NUSD, USDC, USDT],
  description: "Synapse's 3pool stableswap LP token on Aurora",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Aurora Stablecoin Swap (Trisolaris)
 */
export const AURORA_TS_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.AURORA]: '0xffb69779f14E851A8c550Bf5bB1933c44BBDE129',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Trisolaris nUSD LP Token Aurora',
  logo: synapseLogo,
  poolName: 'Aurora Trisolaris Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'auroratrisolarisnusd',
  poolId: 0,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.AURORA]: '0xCCd87854f58773fe75CdDa542457aC48E46c2D65',
  },
  poolTokens: [NUSD, USDC, USDT],
  description: "Trisolaris's 3pool stableswap LP token on Aurora",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Avalanche Stablecoin Swap
 */
export const ARBITRUM_3POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.ARBITRUM]: '0xcFd72be67Ee69A0dd7cF0f846Fc0D98C33d60F16',
  },
  decimals: 18,
  symbol: 'nUSDLP',
  name: 'Synapse nUSD LP Token Arbitrum',
  logo: synapseLogo,
  poolName: 'Arbitrum 3Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'arbitrum3pool',
  poolId: 3,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.ARBITRUM]: '0x9Dd329F5411466d9e0C488fF72519CA9fEf0cb40',
  },
  poolTokens: [NUSD, USDC, USDT],
  description: "Synapse's 3pool stableswap LP token on Arbitrum",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Fantom Stablecoin Swap
 */
export const FANTOM_3POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.FANTOM]: '0x2DC777ff99058a12844A33D9B1AE6c8AB4701F66',
  },
  decimals: 18,
  symbol: 'nUSDLP',
  name: 'Synapse nUSD LP Token Fantom',
  logo: synapseLogo,
  poolName: 'Fantom 3Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'fantom3pool',
  poolId: 3,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.FANTOM]: '0x85662fd123280827e11C59973Ac9fcBE838dC3B4',
  },
  poolTokens: [NUSD, USDC, USDT],
  description: "Synapse's 3pool stableswap LP token on Fantom",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Metis Stablecoin Swap
 */
export const METIS_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.METIS]: '0xC6f684aE516480A35f337a4dA8b40EB6550e07E0',
  },
  decimals: 18,
  symbol: 'nUSDLP',
  name: 'Synapse nUSD LP Token Metis',
  logo: synapseLogo,
  poolName: 'Metis Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'metis2pool',
  poolId: 0,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.METIS]: '0x555982d2E211745b96736665e19D9308B615F78e',
  },
  poolTokens: [NUSD, USDC],
  description: "Synapse's 2pool stableswap LP token on Metis",
  docUrl: SYNAPSE_DOCS_URL,
})


/**
 * Metis Stablecoin Swap
 */
 export const CANTO_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.CANTO]: '0x830c377036a61911aaE49D61c70cb2926823d292',
  },
  decimals: 18,
  name: 'Synapse nUSD LP Token Canto',
  symbol: 'nUSDLP',
  logo: synapseLogo,
  poolName: 'Canto NOTE Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'canto2pool',
  poolId: 0,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.CANTO]: '0x07379565cD8B0CaE7c60Dc78e7f601b34AF2A21c',
  },
  poolTokens: [NUSD, NOTE],
  description: "Synapse's nUSD & NOTE stableswap LP token on Canto",
  docUrl: SYNAPSE_DOCS_URL,
})

export const CANTO_USDC_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.CANTO]: '0xB690FCA5bDc6Ae32c7316dF9B7B7963B7103Fc95',
  },
  decimals: 18,
  name: 'Synapse nUSD LP Token Canto',
  symbol: 'nUSD-LP',
  logo: synapseLogo,
  poolName: 'Canto USDC Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'cantousdcpool',
  poolId: 2,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.CANTO]: '0x273508478e099Fdf953349e6B3704E7c3dEE91a5',
  },
  poolTokens: [NUSD, USDC],
  description: "Synapse's nUSD & USDC stableswap LP token on Canto",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Metis Stablecoin Swap
 */
 export const CANTO_WRAPPER_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.CANTO]: '0x0271984e4cfA2A0f02664baACD551CcFCC9920E8',
  },
  decimals: 18,
  name: 'Synapse nUSD LP Token Canto',
  symbol: 'Wrapper',
  logo: synapseLogo,
  poolName: 'Canto Wrapper Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'cantowrapper',
  poolId: 420,
  poolType: 'USD',
  swapAddresses: {
    [ChainId.CANTO]: '0x0271984e4cfA2A0f02664baACD551CcFCC9920E8',
  },
  poolTokens: [NUSD, NOTE, USDC, USDT],
  description: "",
  docUrl: SYNAPSE_DOCS_URL,
})



/**
 * Klaytn Stablecoin Swap
 */
export const KLAYTN_ORBIT_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.KLAYTN]: '0x656757eb1e04011f6862688957ebA9eA12881244',
  },
  decimals: 18,
  symbol: 'oUSDT-LP',
  name: 'Synapse Orbit UST LP Token Klaytn',
  logo: synapseLogo,
  poolName: 'Klaytn Synapse & Orbit USDT Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'klaytn2pool',
  poolId: 0,
  poolType: 'KLAYTN_USDT',
  swapAddresses: {
    [ChainId.KLAYTN]: '0xfDbaD1699A550F933EFebF652a735F2f89d3833c',
  },
  poolTokens: [KLAYTN_USDT, KLAYTN_oUSDT],
  description: "Synapse's 2pool stableswap LP token on Klaytn",
  docUrl: SYNAPSE_DOCS_URL,
})
