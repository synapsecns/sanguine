import _ from 'lodash'

import { Token } from '../types'
import { ARBITRUM, AVALANCHE, ETH } from '@/constants/chains/master'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'
import { useBridgeState } from '@/slices/bridge/hooks'
import { stringToBigInt } from '@/utils/bigint/format'
import { getUnderlyingBridgeTokens } from '@/utils/getUnderlyingBridgeTokens'
import { useAppSelector } from '@/store/hooks'
import { findTokenByAddressAndChain } from '../findTokenByAddressAndChainId'

export enum BridgeModules {
  SYNAPSE_RFQ = 'SynapseRFQ',
  SYNAPSE_CCTP = 'SynapseCCTP',
  SYNAPSE_BRIDGE = 'SynapseBridge',
}

export const ELIGIBILITY_DEFAULT_TEXT = 'Fee rebate until March 29th'

const TO_FROM_FEES_AND_REBATE_BPS = {
  42161: {
    anyFromChain: {
      SynapseBridge: {
        nETH: { fee: 4, rebate: 6 },
        nUSD: { fee: 4, rebate: 6 },
        GMX: { fee: 5, rebate: 6 },
      },
      SynapseCCTP: {
        USDC: { fee: 4, rebate: 5 },
      },
      SynapseRFQ: {
        USDC: { fee: 4, rebate: 5 },
      },
    },
  },
  1: {
    42161: {
      SynapseBridge: {
        nETH: { fee: 10, rebate: 12 },
        nUSD: { fee: 12, rebate: 14 },
      },
      SynapseCCTP: {
        USDC: { fee: 4, rebate: 5 },
      },
      SynapseRFQ: {
        USDC: { fee: 4, rebate: 5 },
      },
    },
  },
  43114: {
    42161: {
      SynapseBridge: {
        GMX: { fee: 5, rebate: 6 },
      },
      SynapseCCTP: {},
      SynapseRFQ: {},
    },
  },
}

/*
const ARB = {
  name: 'Arbitrum',
  addresses: {
    [ARBITRUM.id]: '0x912CE59144191C1204E64559FE8253a0e49E6548',
  },
  decimals: {
    [ARBITRUM.id]: 6,
  },
  routeSymbol: 'ARB',
}
*/

export const useStipEligibility = () => {
  const {
    ethPrice,
    gmxPrice,
    arbPrice,
    fraxPrice,
    usdtPrice,
    usdcPrice,
    crvUsdPrice,
    daiPrice,
    lusdPrice,
    notePrice,
    susdPrice,
    usdbcPrice,
    usdcePrice,
    usdtePrice,
    musdcPrice,
    daiePrice,
  } = useAppSelector((state) => state.priceData)

  const prices = {
    ETH: ethPrice,
    nETH: ethPrice,
    'WETH.e': ethPrice,
    USDC: usdcPrice,
    USDT: usdtPrice,
    GMX: gmxPrice,
    ARB: arbPrice,
    DAI: daiPrice,
    crvUSD: crvUsdPrice,
    FRAX: fraxPrice,
    NOTE: notePrice,
    sUSD: susdPrice,
    LUSD: lusdPrice,
    USDbC: usdbcPrice,
    'USDC.e': usdcePrice,
    'USDT.e': usdtePrice,
    'DAI.e': daiePrice,
    'm.USDC': musdcPrice,
    nUSD: 1,
  }

  const { fromChainId, fromToken, toChainId, bridgeQuote, debouncedFromValue } =
    useBridgeState()

  const isRouteEligible =
    isRfqEligible(fromToken, fromChainId, toChainId, bridgeQuote) ||
    isCctpEligible(fromToken, fromChainId, toChainId, bridgeQuote) ||
    isSynapseEligible(fromToken, fromChainId, toChainId, bridgeQuote)

  const isActiveRouteEligible =
    isRouteEligible &&
    bridgeQuote &&
    bridgeQuote.outputAmount !== EMPTY_BRIDGE_QUOTE.outputAmount

  const rebate = calculateRebate(
    bridgeQuote,
    fromChainId,
    toChainId,
    fromToken,
    stringToBigInt(debouncedFromValue, fromToken?.decimals[fromChainId]),
    prices[fromToken?.routeSymbol],
    arbPrice
  )

  return {
    isRouteEligible,
    isActiveRouteEligible,
    rebate,
  }
}

export const isChainEligible = (
  fromChainId: number,
  toChainId: number,
  fromToken?: Token
) => {
  // if no from Token
  if (!fromToken) {
    return (
      toChainId === ARBITRUM.id ||
      (fromChainId === ARBITRUM.id && toChainId === ETH.id) ||
      (fromChainId === ARBITRUM.id && toChainId === AVALANCHE.id)
    )
  }

  // if fromToken
  return (
    toArbitrum(fromChainId, toChainId, fromToken) ||
    fromArbitrumToEthereum(fromChainId, toChainId, fromToken) ||
    (fromChainId === ARBITRUM.id &&
      toChainId === AVALANCHE.id &&
      fromToken?.routeSymbol === 'GMX')
  )
}

const toArbitrum = (
  fromChainId: number,
  toChainId: number,
  fromToken: Token
) => {
  const underlyingBridgeTokens = getUnderlyingBridgeTokens(
    fromToken,
    fromChainId
  )

  return (
    toChainId === ARBITRUM.id &&
    _.some(underlyingBridgeTokens, (value) =>
      _.includes(['USDC', 'nUSD', 'CCTP.USDC', 'nETH', 'nUSD', 'GMX'], value)
    )
  )
}

const fromArbitrumToEthereum = (
  fromChainId: number,
  toChainId: number,
  fromToken: Token
) => {
  const underlyingBridgeTokens = getUnderlyingBridgeTokens(
    fromToken,
    fromChainId
  )

  return (
    fromChainId === ARBITRUM.id &&
    toChainId === ETH.id &&
    _.some(underlyingBridgeTokens, (value) =>
      _.includes(['USDC', 'nUSD', 'CCTP.USDC'], value)
    )
  )
}

const calculateRebate = (
  bridgeQuote,
  fromChainId: number,
  toChainId: number,
  fromToken: Token,
  bigIntTokenAmount: bigint,
  tokenPriceInDollars: number,
  arbPrice: number
) => {
  if (
    !bigIntTokenAmount ||
    tokenPriceInDollars === 0 ||
    !tokenPriceInDollars ||
    !arbPrice ||
    bridgeQuote?.outputAmount === EMPTY_BRIDGE_QUOTE.outputAmount
  )
    return

  const {
    bridgeModuleName,
    quotes: {
      originQuery: { tokenOut },
    },
  } = bridgeQuote

  const bridgeToken = findTokenByAddressAndChain(tokenOut, fromChainId)

  if (!bridgeToken) {
    return
  }

  let rebateBps
  if (toChainId === ARBITRUM.id) {
    rebateBps =
      TO_FROM_FEES_AND_REBATE_BPS[toChainId]?.anyFromChain?.[bridgeModuleName][
        bridgeToken?.routeSymbol
      ]?.rebate
  } else {
    rebateBps =
      TO_FROM_FEES_AND_REBATE_BPS[toChainId]?.[fromChainId]?.[bridgeModuleName][
        bridgeToken?.routeSymbol
      ]?.rebate
  }

  if (!rebateBps) {
    return
  }

  const tokenDecimals = fromToken?.decimals[fromChainId]

  const normalizationFactor = BigInt(10 ** tokenDecimals)
  const normalizedTokenAmount =
    Number(bigIntTokenAmount) / Number(normalizationFactor)

  const totalValueInDollars = normalizedTokenAmount * tokenPriceInDollars

  const rebate = (totalValueInDollars * (rebateBps / 10000)) / arbPrice

  return rebate
}

export const isRfqEligible = (
  token: Token,
  fromChainId: number,
  toChainId: number,
  bridgeQuote: any
) => {
  const underlyingBridgeTokens = getUnderlyingBridgeTokens(token, fromChainId)

  if (!underlyingBridgeTokens || !token) {
    return false
  }

  return (
    (bridgeQuote.bridgeModuleName === BridgeModules.SYNAPSE_RFQ &&
      token.swapableType === 'USD' &&
      toChainId === ARBITRUM.id) ||
    (bridgeQuote.bridgeModuleName === BridgeModules.SYNAPSE_RFQ &&
      token.swapableType === 'USD' &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id)
  )
}

export const isCctpEligible = (
  token: Token,
  fromChainId: number,
  toChainId: number,
  bridgeQuote: any
) => {
  const underlyingBridgeTokens = getUnderlyingBridgeTokens(token, fromChainId)

  if (!underlyingBridgeTokens || !token) {
    return false
  }

  return (
    (bridgeQuote.bridgeModuleName === BridgeModules.SYNAPSE_CCTP &&
      _.some(underlyingBridgeTokens, (value) =>
        _.includes(['USDC', 'nUSD', 'CCTP.USDC'], value)
      ) &&
      toChainId === ARBITRUM.id) ||
    (bridgeQuote.bridgeModuleName === BridgeModules.SYNAPSE_CCTP &&
      _.some(underlyingBridgeTokens, (value) =>
        _.includes(['USDC', 'nUSD', 'CCTP.USDC'], value)
      ) &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id)
  )
}

export const isSynapseEligible = (
  token: Token,
  fromChainId: number,
  toChainId: number,
  bridgeQuote: any
) => {
  const underlyingBridgeTokens = getUnderlyingBridgeTokens(token, fromChainId)

  if (!underlyingBridgeTokens || !token) {
    return false
  }

  return (
    (bridgeQuote.bridgeModuleName === BridgeModules.SYNAPSE_BRIDGE &&
      _.some(underlyingBridgeTokens, (value) =>
        _.includes(['nETH', 'nUSD', 'GMX'], value)
      ) &&
      toChainId === ARBITRUM.id) ||
    (bridgeQuote.bridgeModuleName === BridgeModules.SYNAPSE_BRIDGE &&
      _.some(underlyingBridgeTokens, (value) =>
        _.includes(['nETH', 'nUSD'], value)
      ) &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (bridgeQuote.bridgeModuleName === BridgeModules.SYNAPSE_BRIDGE &&
      _.some(underlyingBridgeTokens, (value) => _.includes(['GMX'], value)) &&
      fromChainId === ARBITRUM.id &&
      toChainId === AVALANCHE.id)
  )
}
