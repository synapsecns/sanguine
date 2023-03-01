import synapseLogo from '@assets/icons/synapse.svg'

import { Token } from '@utils/classes/Token'
import { SYNAPSE_DOCS_URL } from '@urls'
import { ChainId } from '@constants/networks'

import {
  NETH,
  WETH,
  ETH,
  AVWETH,
  WETHE,
  ONEETH,
  FTMETH,
  METISETH,
  CANTOETH
} from '@constants/tokens/basic'

/**
 * Avalanche Stablecoin Swap
 */

export const ARBITRUM_ETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.ARBITRUM]: '0xD70A52248e546A3B260849386410C7170c7BD1E9',
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
    [ChainId.ARBITRUM]: '0xa067668661C84476aFcDc6fA5D758C4c01C34352',
  },
  swapEthAddresses: {
    [ChainId.ARBITRUM]: '0x1c3fe783a7c06bfAbd124F2708F5Cc51fA42E102',
  },
  poolTokens: [NETH, WETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, ETH],
  description: "Synapse's eth swap LP token on Arbitrum",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Avalanche Stablecoin Swap
 */
export const OPTIMISM_ETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.OPTIMISM]: '0x4619a06ddd3b8f0f951354ec5e75c09cd1cd1aef',
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
    [ChainId.OPTIMISM]: '0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9',
  },
  swapEthAddresses: {
    [ChainId.OPTIMISM]: '0x8c7d5f8A8e154e1B59C92D8FB71314A43F32ef7B',
  },
  poolTokens: [NETH, WETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, ETH],
  description: "Synapse's eth swap LP token on Optimism",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Avalanche Stablecoin Swap
 */
export const BOBA_ETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.BOBA]: '0x498657f2AF18D525049dE520dD86ee376Db9c67c',
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
    [ChainId.BOBA]: '0x753bb855c8fe814233d26Bb23aF61cb3d2022bE5',
  },
  swapEthAddresses: {
    [ChainId.BOBA]: '0x4F4f66964335D7bef23C16a62Fcd3d1E89f02959',
  },
  poolTokens: [NETH, WETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, ETH],
  description: "Synapse's eth swap LP token on Boba",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Avalanche Stablecoin Swap
 */
export const AVALANCHE_AVETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0x5dF1dB940dd8fEE0e0eB0C8917cb50b4dfaDF98c',
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
    [ChainId.AVALANCHE]: '0x77a7e60555bC18B4Be44C181b2575eee46212d44',
  },
  swapWrapperAddresses: {
    [ChainId.AVALANCHE]: '0xdd60483Ace9B215a7c019A44Be2F22Aa9982652E', // '0xf7e6214E1f2b03b54f1594ECfa3834148aB26888',
  },
  swapEthAddresses: {
    [ChainId.AVALANCHE]: '0xdd60483Ace9B215a7c019A44Be2F22Aa9982652E', // '0xf7e6214E1f2b03b54f1594ECfa3834148aB26888',
  },
  poolTokens: [NETH, AVWETH], // add eth token whether eth or weth here
  nativeTokens: [NETH, WETHE],
  depositTokens: [NETH, WETHE],
  description: "Synapse's ETH swap LP token on Avalanche",
  docUrl: SYNAPSE_DOCS_URL,
})

/**
 * Harmony ETH Swap
 */
export const HARMONY_ONEETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.HARMONY]: '0x464d121D3cA63cEEfd390D76f19364D3Bd024cD2',
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
    [ChainId.HARMONY]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
  },
  poolTokens: [NETH, ONEETH],
  description: "Synapse's ETH swap LP token on Harmony",
  docUrl: SYNAPSE_DOCS_URL,
})

export const FANTOM_WETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.FANTOM]: '0x0e3dD3403ee498694A8f61B04AFed8919F747f77',
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
    [ChainId.FANTOM]: '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
  },
  poolTokens: [NETH, FTMETH],
  description: "Synapse's ETH swap LP token on Fantom",
  docUrl: SYNAPSE_DOCS_URL,
})

export const METIS_WETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.METIS]: '0x9C1340Bf093d057fA29819575517fb9fE2f04AcE',
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
    [ChainId.METIS]: '0x09fEC30669d63A13c666d2129230dD5588E2e240',
  },
  poolTokens: [NETH, METISETH],
  description: "Synapse's ETH swap LP token on Metis",
  docUrl: SYNAPSE_DOCS_URL,
})

export const CANTO_WETH_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.CANTO]: '0xE7002d7Ee2C2aC7A4286F3C075950CcAc2DB3401',
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
    [ChainId.CANTO]: '0xF60F88bA0CB381b8D8A662744fF93486273c22F9',
  },
  poolTokens: [NETH, CANTOETH],
  description: "Synapse's ETH swap LP token on Canto",
  docUrl: SYNAPSE_DOCS_URL,
})