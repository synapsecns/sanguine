import * as CHAINS from '@constants-new/chains/master'
import {
  BUSD,
  USDC,
  USDT,
  DAI,
  NUSD,
  NOTE,
  NETH,
  ETH,
  WETHE,
  ONEETH,
  USDCe,
  USDTe,
  WJEWEL,
  SYNJEWEL,
  AVAX,
  METISUSDC,
  DAIE,
  ONEDAI,
  ONEUSDC,
  ONEUSDT,
  WETH,
} from '@constants-new/tokens/bridgeable'
import {
  AVWETH,
  FANTOMETH,
  FANTOMUSDC,
  FANTOMUSDT,
  KLAYTN_oUSDT,
  MIM,
  MULTIAVAX,
} from '@constants-new/tokens/auxilliary'
import synapseLogo from '@assets/icons/syn.svg'
import { MINICHEF_ADDRESSES } from '@constants-new/minichef'

import { Token } from '../../utils/types/index'

// @dev: Reassign correct priorityRanking

export const ETH_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.ETH.id]: NUSD.addresses[CHAINS.ETH.id],
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
    [CHAINS.ETH.id]: '0x1116898DdA4015eD8dDefb84b6e8Bc24528Af2d8',
  },
  poolTokens: [DAI, USDC, USDT],
  description: "Synapse's 3pool stableswap LP token on ETH",
  display: true,
  priorityPool: true,
  color: 'gray',
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.ETH.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.ETH.id],
})

export const BSC_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.BNB.id]: '0xa4b7Bc06EC817785170C2DbC1dD3ff86CDcdcc4C',
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
    [CHAINS.BNB.id]: '0x28ec0B36F0819ecB5005cAB836F4ED5a2eCa4D13',
  },
  poolTokens: [NUSD, BUSD, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token",
  display: true,
  priorityPool: true,
  color: 'purple',
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.BNB.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.BNB.id],
})

export const OPTIMISM_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.OPTIMISM.id]: '0x2c6d91accC5Aa38c84653F28A80AEC69325BDd12',
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
    [CHAINS.OPTIMISM.id]: '0xF44938b0125A6662f9536281aD2CD6c499F22004',
  },
  poolTokens: [NUSD, USDCe],
  description: "Synapse's 2pool stableswap LP token on Optimism",
  display: true,
  priorityPool: true,
  priorityRank: 6,
  chainId: CHAINS.OPTIMISM.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.OPTIMISM.id],
})

export const CRONOS_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.CRONOS.id]: '0xd5Da35646A79C42a0BAB148656192A22e8CBfad6', //'0x8415D4EB17F0949e2388fdF52909db4cC0a2B082',
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
    [CHAINS.CRONOS.id]: '0xCb6674548586F20ca39C97A52A0ded86f48814De', //'0x3b1601c386247A127287b094F9CCB585D4D0B99b',
  },
  poolTokens: [NUSD, USDC], // [NUSD, DAI, USDC, USDT],
  description: "Synapse's 2pool stableswap LP token on Cronos",
  display: true,
  priorityPool: true,
  priorityRank: 6,
  chainId: CHAINS.CRONOS.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.CRONOS.id],
})

export const POLYGON_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.POLYGON.id]: '0x7479e1bc2f2473f9e78c89b4210eb6d55d33b645',
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
    [CHAINS.POLYGON.id]: '0x85fCD7Dd0a1e1A9FCD5FD886ED522dE8221C3EE5',
  },
  poolTokens: [NUSD, DAI, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Polygon/Matic",
  display: true,
  priorityPool: true,
  color: 'purple',
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.POLYGON.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.POLYGON.id],
})

export const AVALANCHE_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0xa5C7E254b77e9401639Bd5f261dae6D5E4441A8F',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Avalanche',
  logo: synapseLogo,
  poolName: 'Avalanche Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'avalanchenusd',
  // poolId: 3,
  // 420 pool id sets pool to not show in staking page
  notStake: true,
  poolId: 3,
  poolType: 'USD',
  swapAddresses: {
    [CHAINS.AVALANCHE.id]: '0xA196a03653f6cc5cA0282A8BD7Ec60e93f620afc',
  },
  poolTokens: [NUSD, USDC, USDT],
  description: "Synapse's 3pool stableswap LP token on Avalanche",
  priorityRank: 6,
  chainId: CHAINS.AVALANCHE.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.AVALANCHE.id],
})

export const LEGACY_AVALANCHE_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0xCA87BF3ec55372D9540437d7a86a7750B42C02f4',
  },
  decimals: 18,
  symbol: 'nUSD-LP',
  name: 'Synapse nUSD LP Token Avalanche',
  logo: synapseLogo,
  poolName: 'Legacy Avalanche Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'legacyavalanchenusd',
  poolId: 1,
  poolType: 'USD',
  swapAddresses: {
    [CHAINS.AVALANCHE.id]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
  },
  poolTokens: [NUSD, DAIE, USDCe, USDTe],
  description: "Synapse's 4pool stableswap LP token on Avalanche",
  priorityRank: 6,
  chainId: CHAINS.AVALANCHE.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.AVALANCHE.id],
})

export const ARBITRUM_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.ARBITRUM.id]: '0xADeac0343C2Ac62DFE5A5f51E896AefFF5Ab513E',
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
    [CHAINS.ARBITRUM.id]: '0x0Db3FE3B770c95A0B99D1Ed6F2627933466c0Dd8',
  },
  poolTokens: [NUSD, MIM, USDCe, USDT],
  description: "Synapse's 4pool stableswap LP token on Arbitrum",
  display: true,
  legacy: true,
  color: 'purple',
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.ARBITRUM.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.ARBITRUM.id],
})

export const FANTOM_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x464d121D3cA63cEEfd390D76f19364D3Bd024cD2',
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
    [CHAINS.FANTOM.id]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
  },
  poolTokens: [NUSD, MIM, FANTOMUSDC, FANTOMUSDT],
  description: "Synapse's 4pool stableswap LP token on Fantom",
  display: true,
  legacy: true,
  color: 'blue',
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.FANTOM.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.FANTOM.id],
})

export const HARMONY_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0xE269abBFAF52b26D2632F55B6b223A5223088B96',
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
    [CHAINS.HARMONY.id]: '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e',
  },
  poolTokens: [NUSD, ONEDAI, ONEUSDC, ONEUSDT],
  description: "Synapse's 4pool stableswap LP token on Harmony",
  display: true,
  priorityPool: true,
  color: 'purple',
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.HARMONY.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.HARMONY.id],
})

export const BOBA_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.BOBA.id]: '0x9D7283A6AeeD9BCd4Ac70876fEA2b69a63DD8cb9',
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
    [CHAINS.BOBA.id]: '0x75FF037256b36F15919369AC58695550bE72fead',
  },
  poolTokens: [NUSD, DAI, USDC, USDT],
  description: "Synapse's 4pool stableswap LP token on Boba",
  display: true,
  priorityPool: true,
  color: 'green',
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.BOBA.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.BOBA.id],
})

export const AURORA_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.AURORA.id]: '0xEAdC3524f3F007cdC5104BF28663b1141D3e3127',
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
    [CHAINS.AURORA.id]: '0xcEf6C2e20898C2604886b888552CA6CcF66933B0',
  },
  poolTokens: [NUSD, USDCe, USDTe],
  description: "Synapse's 3pool stableswap LP token on Aurora",
  display: true,
  legacy: true,
  priorityRank: 6,
  chainId: CHAINS.AURORA.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.AURORA.id],
})

export const AURORA_TS_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.AURORA.id]: '0xffb69779f14E851A8c550Bf5bB1933c44BBDE129',
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
    [CHAINS.AURORA.id]: '0xCCd87854f58773fe75CdDa542457aC48E46c2D65',
  },
  poolTokens: [NUSD, USDC, USDT],
  description: "Trisolaris's 3pool stableswap LP token on Aurora",
  display: false,
  priorityPool: true,
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.AURORA.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.AURORA.id],
})

export const ARBITRUM_3POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.ARBITRUM.id]: '0xcFd72be67Ee69A0dd7cF0f846Fc0D98C33d60F16',
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
    [CHAINS.ARBITRUM.id]: '0x9Dd329F5411466d9e0C488fF72519CA9fEf0cb40',
  },
  poolTokens: [NUSD, USDCe, USDT],
  description: "Synapse's 3pool stableswap LP token on Arbitrum",
  display: true,
  priorityPool: true,
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.ARBITRUM.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.ARBITRUM.id],
})

export const FANTOM_3POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x2DC777ff99058a12844A33D9B1AE6c8AB4701F66',
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
    [CHAINS.FANTOM.id]: '0x85662fd123280827e11C59973Ac9fcBE838dC3B4',
  },
  poolTokens: [NUSD, FANTOMUSDC, FANTOMUSDT],
  description: "Synapse's 3pool stableswap LP token on Fantom",
  display: true,
  priorityPool: true,
  priceUnits: 'USD',
  priorityRank: 6,
  chainId: CHAINS.FANTOM.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.FANTOM.id],
})

export const METIS_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.METIS.id]: '0xC6f684aE516480A35f337a4dA8b40EB6550e07E0',
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
    [CHAINS.METIS.id]: '0x555982d2E211745b96736665e19D9308B615F78e',
  },
  poolTokens: [NUSD, METISUSDC],
  description: "Synapse's 2pool stableswap LP token on Metis",
  display: true,
  priorityPool: true,
  priorityRank: 6,
  chainId: CHAINS.METIS.id,
  incentivized: true,
  customRewardToken: 'METIS',
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.METIS.id],
})

export const CANTO_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.CANTO.id]: '0x830c377036a61911aaE49D61c70cb2926823d292',
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
    [CHAINS.CANTO.id]: '0x07379565cD8B0CaE7c60Dc78e7f601b34AF2A21c',
  },
  poolTokens: [NUSD, NOTE],
  description: "Synapse's nUSD & NOTE stableswap LP token on Canto",
  display: true,
  priorityRank: 6,
  chainId: CHAINS.CANTO.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.CANTO.id],
})

// export const CANTO_USDC_SWAP_TOKEN = new Token({
//   addresses: {
//     [CHAINS.CANTO.id]: '0xB690FCA5bDc6Ae32c7316dF9B7B7963B7103Fc95',
//   },
//   decimals: 18,
//   name: 'Synapse nUSD LP Token Canto',
//   symbol: 'nUSD-LP',
//   logo: synapseLogo,
//   poolName: 'Canto USDC Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
//   routerIndex: 'cantousdcpool',
//   poolId: 2,
//   poolType: 'USD',
//   swapAddresses: {
//     [CHAINS.CANTO.id]: '0x273508478e099Fdf953349e6B3704E7c3dEE91a5',
//   },
//   poolTokens: [NUSD, USDC],
//   description: "Synapse's nUSD & USDC stableswap LP token on Canto",
//   display: true,
//   priorityRank: 6,
//   chainId: CHAINS.CANTO.id,
//   miniChefAddress: MINICHEF_ADDRESSES[CHAINS.CANTO.id],
// })

export const CANTO_WRAPPER_POOL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.CANTO.id]: '0x0271984e4cfA2A0f02664baACD551CcFCC9920E8',
  },
  decimals: 18,
  name: 'Synapse nUSD LP Token Canto',
  symbol: 'Wrapper',
  logo: synapseLogo,
  poolName: 'Canto Wrapper Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'cantowrapper',
  poolId: 420,
  notStake: true,
  poolType: 'USD',
  swapAddresses: {
    [CHAINS.CANTO.id]: '0x0271984e4cfA2A0f02664baACD551CcFCC9920E8',
  },
  poolTokens: [NUSD, NOTE, USDC, USDT],
  description: '',
  display: false,
  priorityPool: true,
  priorityRank: 6,
  chainId: CHAINS.CANTO.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.CANTO.id],
})

export const KLAYTN_ORBIT_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.KLAYTN.id]: '0x656757eb1e04011f6862688957ebA9eA12881244',
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
    [CHAINS.KLAYTN.id]: '0xfDbaD1699A550F933EFebF652a735F2f89d3833c',
  },
  poolTokens: [USDT, KLAYTN_oUSDT],
  description: "Synapse's 2pool stableswap LP token on Klaytn",
  display: true,
  priorityPool: true,
  priorityRank: 6,
  chainId: CHAINS.KLAYTN.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.KLAYTN.id],
})

export const HARMONY_AVAX_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0x02f7D17f126BD54573c8EbAD9e05408A56f46452',
  },
  decimals: 18,
  symbol: 'AVAXLP',
  name: 'AVAX LP Token Harmony ',
  logo: synapseLogo,
  poolName: 'Harmony AVAX Swap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'avax2pool',
  poolId: 0,
  poolType: 'AVAX',
  swapAddresses: {
    [CHAINS.HARMONY.id]: '0x00A4F57D926781f62D09bb05ec76e6D8aE4268da',
  },
  poolTokens: [AVAX, MULTIAVAX],
  description: "Synapse's 2pool AVAX LP token on Harmony",
  display: true,
  priorityPool: true,
  color: 'red',
  priceUnits: 'AVAX',
  priorityRank: 6,
  chainId: CHAINS.HARMONY.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.HARMONY.id],
})

export const ARBITRUM_ETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.ARBITRUM.id]: '0xD70A52248e546A3B260849386410C7170c7BD1E9',
  },
  decimals: 18,
  symbol: 'nETH-LP', // make sure this gets update to match conytract
  name: 'Synapse Eth LP Token Arbitrum',
  logo: synapseLogo,
  poolName: 'Arbitrum ETH Pool',
  routerIndex: 'arbitrumethpool',
  poolId: 0,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.ARBITRUM.id]: '0xa067668661C84476aFcDc6fA5D758C4c01C34352',
  },
  swapEthAddresses: {
    [CHAINS.ARBITRUM.id]: '0x1c3fe783a7c06bfAbd124F2708F5Cc51fA42E102',
  },
  poolTokens: [NETH, WETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, ETH],
  description: "Synapse's eth swap LP token on Arbitrum",
  display: true,
  priorityPool: true,
  color: 'sky',
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.ARBITRUM.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.ARBITRUM.id],
})

export const OPTIMISM_ETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.OPTIMISM.id]: '0x4619a06ddd3b8f0f951354ec5e75c09cd1cd1aef',
  },
  decimals: 18,
  symbol: 'nETH-LP', // make sure this gets update to match conytract
  name: 'Synapse Eth LP Token Optimism',
  logo: synapseLogo,
  poolName: 'Optimism ETH Pool',
  routerIndex: 'optimismethpool',
  poolId: 0,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.OPTIMISM.id]: '0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9',
  },
  swapEthAddresses: {
    [CHAINS.OPTIMISM.id]: '0x8c7d5f8A8e154e1B59C92D8FB71314A43F32ef7B',
  },
  poolTokens: [NETH, WETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, ETH],
  description: "Synapse's eth swap LP token on Optimism",
  display: true,
  priorityPool: true,
  color: 'sky',
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.OPTIMISM.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.OPTIMISM.id],
})

export const BOBA_ETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.BOBA.id]: '0x498657f2AF18D525049dE520dD86ee376Db9c67c',
  },
  decimals: 18,
  symbol: 'nETH-LP', // make sure this gets update to match conytract
  name: 'Synapse Eth LP Token Boba',
  logo: synapseLogo,
  poolName: 'Boba ETH Pool',
  routerIndex: 'bobaethpool',
  poolId: 2,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.BOBA.id]: '0x753bb855c8fe814233d26Bb23aF61cb3d2022bE5',
  },
  swapEthAddresses: {
    [CHAINS.BOBA.id]: '0x4F4f66964335D7bef23C16a62Fcd3d1E89f02959',
  },
  poolTokens: [NETH, WETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, ETH],
  description: "Synapse's eth swap LP token on Boba",
  display: true,
  priorityPool: true,
  color: 'sky',
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.BOBA.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.BOBA.id],
})

export const AVALANCHE_AVETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0x5dF1dB940dd8fEE0e0eB0C8917cb50b4dfaDF98c',
  },
  decimals: 18,
  symbol: 'nETH-LP', // make sure this gets update to match conytract
  name: 'Synapse Eth LP Token Avalanche',
  logo: synapseLogo,
  poolName: 'Avalanche ETH Pool',
  routerIndex: 'avalancheethpool',
  poolId: 2,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.AVALANCHE.id]: '0x77a7e60555bC18B4Be44C181b2575eee46212d44',
  },
  swapWrapperAddresses: {
    [CHAINS.AVALANCHE.id]: '0xdd60483Ace9B215a7c019A44Be2F22Aa9982652E', // '0xf7e6214E1f2b03b54f1594ECfa3834148aB26888',
  },
  swapEthAddresses: {
    [CHAINS.AVALANCHE.id]: '0xdd60483Ace9B215a7c019A44Be2F22Aa9982652E', // '0xf7e6214E1f2b03b54f1594ECfa3834148aB26888',
  },
  poolTokens: [NETH, AVWETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, WETHE],
  depositTokens: [NETH, WETHE],
  description: "Synapse's ETH swap LP token on Avalanche",
  display: true,
  priorityPool: true,
  color: 'sky',
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.AVALANCHE.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.AVALANCHE.id],
})

export const HARMONY_ONEETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0x464d121D3cA63cEEfd390D76f19364D3Bd024cD2',
  },
  decimals: 18,
  symbol: 'nETH-LP',
  name: 'Synapse 1ETH LP Token Harmony',
  logo: synapseLogo,
  poolName: 'Harmony 1ETH Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'harmonyethpool',
  poolId: 2,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.HARMONY.id]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
  },
  poolTokens: [NETH, ONEETH],
  description: "Synapse's ETH swap LP token on Harmony",
  display: true,
  priorityPool: true,
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.HARMONY.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.HARMONY.id],
})

export const FANTOM_WETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x0e3dD3403ee498694A8f61B04AFed8919F747f77',
  },
  decimals: 18,
  symbol: 'nETH-LP',
  name: 'Synapse ETH LP Token Fantom',
  logo: synapseLogo,
  poolName: 'Fantom ETH Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'fantomethpool',
  poolId: 2,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.FANTOM.id]: '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
  },
  poolTokens: [NETH, FANTOMETH],
  description: "Synapse's ETH swap LP token on Fantom",
  display: true,
  priorityPool: true,
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.FANTOM.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.FANTOM.id],
})

export const METIS_WETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.METIS.id]: '0x9C1340Bf093d057fA29819575517fb9fE2f04AcE',
  },
  decimals: 18,
  symbol: 'nETH-LP',
  name: 'Synapse ETH LP Token Metis',
  logo: synapseLogo,
  poolName: 'Metis ETH Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'metisethpool',
  poolId: 1,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.METIS.id]: '0x09fEC30669d63A13c666d2129230dD5588E2e240',
  },
  poolTokens: [NETH, WETH],
  description: "Synapse's ETH swap LP token on Metis",
  display: true,
  priorityPool: true,
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.METIS.id,
  incentivized: true,
  customRewardToken: 'METIS',
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.METIS.id],
})

export const CANTO_WETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.CANTO.id]: '0xE7002d7Ee2C2aC7A4286F3C075950CcAc2DB3401',
  },
  decimals: 18,
  symbol: 'nETH-LP',
  name: 'Synapse ETH LP Token Canto',
  logo: synapseLogo,
  poolName: 'Canto ETH Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'cantoethpool',
  poolId: 1,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.CANTO.id]: '0xF60F88bA0CB381b8D8A662744fF93486273c22F9',
  },
  poolTokens: [NETH, ETH],
  description: "Synapse's ETH swap LP token on Canto",
  display: true,
  priorityPool: true,
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.CANTO.id,
  incentivized: false,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.CANTO.id],
})

export const HARMONY_JEWEL_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0x0000000000000000000000000000000000000000',
  },
  decimals: 18,
  symbol: 'JEWELP',
  name: 'Jewel LP Token Harmony ',
  logo: synapseLogo,
  poolName: 'Harmony Jewel Swap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'jewel2pool',
  poolId: 0,
  poolType: 'JEWEL',
  swapAddresses: {
    [CHAINS.HARMONY.id]: '0x7bE461cce1501f07969BCE24Ccb2140fCA0a35b3',
  },
  poolTokens: [WJEWEL, SYNJEWEL],
  description: "Synapse's 2pool JEWEL swapper psuedotoken on Harmony",
  display: false,
  priorityPool: true,
  priorityRank: 6,
  chainId: CHAINS.HARMONY.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.HARMONY.id],
})

export const BASE_ETH_SWAP_TOKEN = new Token({
  addresses: {
    [CHAINS.BASE.id]: '0xC35b09c8fdefc90ec580B327d32DeAAda4f581Da',
  },
  decimals: 18,
  symbol: 'nETH-LP',
  name: 'Synapse Eth LP Token Base',
  logo: synapseLogo,
  poolName: 'Base ETH Pool',
  routerIndex: 'baseethpool',
  poolId: 0,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.BASE.id]: '0x6223bD82010E2fB69F329933De20897e7a4C225f',
  },
  swapEthAddresses: {
    [CHAINS.BASE.id]: '0xa9E90579eb086bcdA910dD94041ffE041Fb4aC89',
  },
  poolTokens: [NETH, WETH],
  nativeTokens: [NETH, ETH],
  description: "Synapse's eth swap LP token on Base",
  display: true,
  priorityPool: true,
  color: 'sky',
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.BASE.id,
  incentivized: true,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.BASE.id],
})

// MIGRATED POOLS

export const METIS_WETH_SWAP_TOKEN_MIGRATED = new Token({
  addresses: {
    [CHAINS.METIS.id]: '0x9C1340Bf093d057fA29819575517fb9fE2f04AcE',
  },
  decimals: 18,
  symbol: 'nETH-LP',
  name: 'Synapse ETH LP Token Metis',
  logo: synapseLogo,
  poolName: 'Metis ETH Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'metisethpool-migrated',
  poolId: 1,
  poolType: 'ETH',
  swapAddresses: {
    [CHAINS.METIS.id]: '0x09fEC30669d63A13c666d2129230dD5588E2e240',
  },
  poolTokens: [NETH, WETH],
  description: "Synapse's ETH swap LP token on Metis",
  display: true,
  priorityPool: true,
  priceUnits: 'ETH',
  priorityRank: 6,
  chainId: CHAINS.METIS.id,
  incentivized: false,
  miniChefAddress: '0xaB0D8Fc46249DaAcd5cB36c5F0bC4f0DAF34EBf5',
})

export const METIS_POOL_SWAP_TOKEN_MIGRATED = new Token({
  addresses: {
    [CHAINS.METIS.id]: '0xC6f684aE516480A35f337a4dA8b40EB6550e07E0',
  },
  decimals: 18,
  symbol: 'nUSDLP',
  name: 'Synapse nUSD LP Token Metis',
  logo: synapseLogo,
  poolName: 'Metis Stableswap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'metis2pool-migrated',
  poolId: 0,
  poolType: 'USD',
  swapAddresses: {
    [CHAINS.METIS.id]: '0x555982d2E211745b96736665e19D9308B615F78e',
  },
  poolTokens: [NUSD, METISUSDC],
  description: "Synapse's 2pool stableswap LP token on Metis",
  display: true,
  priorityPool: true,
  priorityRank: 6,
  chainId: CHAINS.METIS.id,
  incentivized: false,
  miniChefAddress: '0xaB0D8Fc46249DaAcd5cB36c5F0bC4f0DAF34EBf5',
})
