import { BigNumber, parseFixed } from '@ethersproject/bignumber'

import { getTestProvider } from '../constants/testProviders'
import { SupportedChainId, FAST_BRIDGE_ADDRESS_MAP } from '../constants'
import { FastBridgeSet } from './fastBridgeSet'
import { FastBridgeQuoteAPI } from './quote'
import { ChainToken } from './ticker'

type Pricing = {
  originAmount: number
  fixedFee: number
  destAmount: number
  originDecimals: number
  destDecimals: number
}

const createQuoteTokenFragment = (
  originToken: ChainToken,
  destToken: ChainToken
): {
  origin_chain_id: number
  origin_token_addr: string
  dest_chain_id: number
  dest_token_addr: string
  origin_fast_bridge_address: string
  dest_fast_bridge_address: string
} => {
  return {
    origin_chain_id: originToken.chainId,
    origin_token_addr: originToken.token,
    dest_chain_id: destToken.chainId,
    dest_token_addr: destToken.token,
    origin_fast_bridge_address: FAST_BRIDGE_ADDRESS_MAP[originToken.chainId],
    dest_fast_bridge_address: FAST_BRIDGE_ADDRESS_MAP[destToken.chainId],
  }
}

const createQuotePricingFragment = (
  price: Pricing
): {
  max_origin_amount: string
  fixed_fee: string
  dest_amount: string
} => {
  return {
    max_origin_amount: parseFixed(
      price.originAmount.toString(),
      price.originDecimals
    ).toString(),
    fixed_fee: parseFixed(
      price.fixedFee.toString(),
      price.originDecimals
    ).toString(),
    dest_amount: parseFixed(
      price.destAmount.toString(),
      price.destDecimals
    ).toString(),
  }
}

const createBridgeRouteTest = (
  fastBridgeSet: FastBridgeSet,
  originToken: ChainToken,
  destToken: ChainToken,
  originAmount: BigNumber,
  expectedDestAmounts: BigNumber[]
) => {
  it(`Should return ${expectedDestAmounts.length} routes for amount=${originAmount}`, async () => {
    const routes = await fastBridgeSet.getBridgeRoutes(
      originToken.chainId,
      destToken.chainId,
      originToken.token,
      destToken.token,
      originAmount
    )
    expect(routes.length).toEqual(expectedDestAmounts.length)
    routes.forEach((route, index) => {
      expect(route.destQuery.minAmountOut).toEqual(expectedDestAmounts[index])
    })
  })
}

const createBridgeRoutesTests = (
  fastBridgeSet: FastBridgeSet,
  originDecimals: number,
  destDecimals: number
) => {
  const tokenA = '0x000000000000000000000000000000000000000A'
  const tokenB = '0x000000000000000000000000000000000000000b'

  const arbA: ChainToken = {
    chainId: SupportedChainId.ARBITRUM,
    token: tokenA,
  }
  const arbB: ChainToken = {
    chainId: SupportedChainId.ARBITRUM,
    token: tokenB,
  }
  const opA: ChainToken = {
    chainId: SupportedChainId.OPTIMISM,
    token: tokenA,
  }
  const opB: ChainToken = {
    chainId: SupportedChainId.OPTIMISM,
    token: tokenB,
  }

  const price1: Pricing = {
    originAmount: 10000,
    fixedFee: 10,
    destAmount: 10000,
    originDecimals,
    destDecimals,
  }

  // Better price with higher fixed fee and lower liquidity
  const price2: Pricing = {
    originAmount: 1000,
    fixedFee: 100,
    destAmount: 2000,
    originDecimals,
    destDecimals,
  }

  // Use following combinations of tokens and prices:
  // - ARB_A -> OP_A: []
  // - ARB_A -> OP_B: [price1]
  // - ARB_B -> OP_A: [price2]
  // - ARB_B -> OP_B: [price1, price2]
  const mockedQuotesAPI: FastBridgeQuoteAPI[] = [
    {
      ...createQuoteTokenFragment(arbA, opB),
      ...createQuotePricingFragment(price1),
      relayer_addr: '0x0',
      updated_at: '2021-01-01T00:00:00.000Z',
    },
    {
      ...createQuoteTokenFragment(arbB, opA),
      ...createQuotePricingFragment(price2),
      relayer_addr: '0x0',
      updated_at: '2021-01-01T00:00:00.000Z',
    },
    {
      ...createQuoteTokenFragment(arbB, opB),
      ...createQuotePricingFragment(price1),
      relayer_addr: '0x0',
      updated_at: '2021-01-01T00:00:00.000Z',
    },
    {
      ...createQuoteTokenFragment(arbB, opB),
      ...createQuotePricingFragment(price2),
      relayer_addr: '0x0',
      updated_at: '2021-01-01T00:00:00.000Z',
    },
  ]

  beforeEach(() => {
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 200,
        ok: true,
        json: () => Promise.resolve(mockedQuotesAPI),
      })
    ) as any
    // Use UpdatedAt + 1 minute as the current time
    Date.now = jest.fn(() => Date.parse('2021-01-01T00:01:00.000Z'))
  })

  describe('arbA -> opA [no routes]', () => {
    createBridgeRouteTest(
      fastBridgeSet,
      arbA,
      opA,
      parseFixed('20', originDecimals),
      []
    )

    createBridgeRouteTest(
      fastBridgeSet,
      arbA,
      opA,
      parseFixed('500', originDecimals),
      []
    )

    createBridgeRouteTest(
      fastBridgeSet,
      arbA,
      opA,
      parseFixed('10011', originDecimals),
      []
    )
  })

  describe('arbA -> opB [(1:1, 10 fee, up to 10k)]', () => {
    createBridgeRouteTest(
      fastBridgeSet,
      arbA,
      opB,
      parseFixed('10', originDecimals),
      []
    )

    createBridgeRouteTest(
      fastBridgeSet,
      arbA,
      opB,
      parseFixed('500', originDecimals),
      [parseFixed('490', destDecimals)]
    )

    // Higher than available liquidity
    createBridgeRouteTest(
      fastBridgeSet,
      arbA,
      opB,
      parseFixed('10011', originDecimals),
      []
    )
  })

  describe('arbB -> opA [(1:2, 100 fee, up to 1k)]', () => {
    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opA,
      parseFixed('100', originDecimals),
      []
    )

    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opA,
      parseFixed('500', originDecimals),
      [parseFixed('800', destDecimals)]
    )

    // Higher than available liquidity
    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opA,
      parseFixed('1101', originDecimals),
      []
    )
  })

  describe('arbB -> opB [(1:1, 10 fee, up to 10k), (1:2, 100 fee, up to 1k)]', () => {
    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opB,
      parseFixed('10', originDecimals),
      []
    )

    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opB,
      parseFixed('100', originDecimals),
      [parseFixed('90', destDecimals)]
    )

    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opB,
      parseFixed('500', originDecimals),
      [parseFixed('490', destDecimals), parseFixed('800', destDecimals)]
    )

    // Higher than available liquidity for second price
    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opB,
      parseFixed('1101', originDecimals),
      [parseFixed('1091', destDecimals)]
    )

    // Higher than available liquidity for both prices
    createBridgeRouteTest(
      fastBridgeSet,
      arbB,
      opB,
      parseFixed('11011', originDecimals),
      []
    )
  })

  describe('timestamps', () => {
    afterEach(() => {
      // Use UpdatedAt + 1 minute as the current time
      Date.now = jest.fn(() => Date.parse('2021-01-01T00:01:00.000Z'))
    })

    it('ignores quotes with negative age', async () => {
      Date.now = jest.fn(() => Date.parse('2021-01-01T00:00:00.000Z') - 1)
      // arbB -> opB should have two quotes for 500 by default
      // But we expect zero quotes because the quotes are outdated
      const routes = await fastBridgeSet.getBridgeRoutes(
        arbB.chainId,
        opB.chainId,
        arbB.token,
        opB.token,
        parseFixed('500', originDecimals)
      )
      expect(routes.length).toEqual(0)
    })

    it('ignores quotes with age of 5 minutes', async () => {
      Date.now = jest.fn(() => Date.parse('2021-01-01T00:05:00.000Z'))
      // arbB -> opB should have two quotes for 500 by default
      // But we expect zero quotes because the quotes are outdated
      const routes = await fastBridgeSet.getBridgeRoutes(
        arbB.chainId,
        opB.chainId,
        arbB.token,
        opB.token,
        parseFixed('500', originDecimals)
      )
      expect(routes.length).toEqual(0)
    })

    it('includes quotes with age of 5 minutes - 1 millisecond', async () => {
      Date.now = jest.fn(() => Date.parse('2021-01-01T00:04:59.999Z'))
      // arbB -> opB should have two quotes for 500 by default
      // We expect two quotes because the quotes are not outdated
      const routes = await fastBridgeSet.getBridgeRoutes(
        arbB.chainId,
        opB.chainId,
        arbB.token,
        opB.token,
        parseFixed('500', originDecimals)
      )
      expect(routes.length).toEqual(2)
    })
  })
}

describe('FastBridgeSet', () => {
  const chainIds = [
    SupportedChainId.ARBITRUM,
    SupportedChainId.OPTIMISM,
    SupportedChainId.DOGECHAIN,
  ]
  const fastBridgeSet = new FastBridgeSet(
    chainIds.map((chainId) => ({
      chainId,
      provider: getTestProvider(chainId),
    }))
  )

  describe('getModule', () => {
    it('Returns correct module', () => {
      const module = fastBridgeSet.getModule(SupportedChainId.ARBITRUM)
      expect(module).toBeDefined()
      expect(module?.address).toEqual(
        FAST_BRIDGE_ADDRESS_MAP[SupportedChainId.ARBITRUM]
      )
    })

    it('Returns undefined for chain without module', () => {
      const module = fastBridgeSet.getModule(SupportedChainId.DOGECHAIN)
      expect(module).toBeUndefined()
    })

    it('Returns undefined for undefined chain', () => {
      const module = fastBridgeSet.getModule(SupportedChainId.BSC)
      expect(module).toBeUndefined()
    })
  })

  describe('getEstimatedTime', () => {
    it('Returns correct estimated time', () => {
      const estimatedTime = fastBridgeSet.getEstimatedTime(
        SupportedChainId.ARBITRUM
      )
      expect(estimatedTime).toEqual(30)
    })

    it('Throws error for chain without estimated time', () => {
      expect(() =>
        fastBridgeSet.getEstimatedTime(SupportedChainId.DOGECHAIN)
      ).toThrow()
    })

    it('Throws error for undefined chain', () => {
      expect(() =>
        fastBridgeSet.getEstimatedTime(SupportedChainId.BSC)
      ).toThrow()
    })
  })

  describe('getBridgeRoutes', () => {
    describe('6:6 decimals', () => {
      createBridgeRoutesTests(fastBridgeSet, 6, 6)
    })

    describe('18:18 decimals', () => {
      createBridgeRoutesTests(fastBridgeSet, 18, 18)
    })

    describe('6:18 decimals', () => {
      createBridgeRoutesTests(fastBridgeSet, 6, 18)
    })

    describe('18:6 decimals', () => {
      createBridgeRoutesTests(fastBridgeSet, 18, 6)
    })
  })
})
