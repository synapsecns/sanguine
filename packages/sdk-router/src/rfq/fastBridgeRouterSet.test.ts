import { Provider } from '@ethersproject/abstract-provider'
import { BigNumber, parseFixed } from '@ethersproject/bignumber'

import { SupportedChainId } from '../constants'
import { getTestProvider } from '../constants/testProviders'
import { FastBridgeRouterSet } from './fastBridgeRouterSet'
import { ChainProvider } from '../router'
import { CCTPRouterQuery } from '../module'
import { createSlippageTests } from '../router/synapseCCTPRouterSet.test'

describe('FastBridgeRouterSet', () => {
  const ethProvider: Provider = getTestProvider(SupportedChainId.ETH)
  const arbProvider: Provider = getTestProvider(SupportedChainId.ARBITRUM)

  const testProviders: ChainProvider[] = [
    {
      chainId: SupportedChainId.ETH,
      provider: ethProvider,
    },
    {
      chainId: SupportedChainId.ARBITRUM,
      provider: arbProvider,
    },
  ]

  const routerSet = new FastBridgeRouterSet(testProviders)

  describe('applySlippage', () => {
    const fixedFee = parseFixed('40', 18)
    const originQuery: CCTPRouterQuery = {
      routerAdapter: '1',
      tokenOut: '2',
      minAmountOut: parseFixed('1000', 18),
      deadline: BigNumber.from(3),
      rawParams: '4',
    }

    const destQuery: CCTPRouterQuery = {
      routerAdapter: '5',
      tokenOut: '6',
      minAmountOut: originQuery.minAmountOut.sub(fixedFee),
      deadline: BigNumber.from(8),
      rawParams: '9',
    }

    describe('0% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        parseFixed('1000', 18),
        destQuery.minAmountOut,
        0,
        10000
      )
    })

    // Destination amount should be not modified by slippage
    describe('0.1% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        parseFixed('999', 18),
        destQuery.minAmountOut,
        10,
        10000
      )
    })

    // Origin slippage should be capped at 5% of the fixed fee
    // Destination amount should be not modified by slippage
    describe('1% slippage', () => {
      createSlippageTests(
        routerSet,
        originQuery,
        destQuery,
        // 1000 - 40 * 0.05 = 998
        parseFixed('998', 18),
        destQuery.minAmountOut,
        100,
        10000
      )
    })
  })
})
