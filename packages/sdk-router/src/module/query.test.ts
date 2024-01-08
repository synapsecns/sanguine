import { BigNumber } from '@ethersproject/bignumber'

import {
  RouterQuery,
  CCTPRouterQuery,
  Query,
  reduceToQuery,
  narrowToRouterQuery,
  narrowToCCTPRouterQuery,
  modifyDeadline,
  applySlippage,
  applySlippageInBips,
  createNoSwapQuery,
} from './query'

describe('#query', () => {
  const routerQuery: RouterQuery = {
    swapAdapter: '1',
    tokenOut: '2',
    minAmountOut: BigNumber.from(3),
    deadline: BigNumber.from(4),
    rawParams: '5',
  }

  const cctpRouterQuery: CCTPRouterQuery = {
    routerAdapter: '6',
    tokenOut: '7',
    minAmountOut: BigNumber.from(8),
    deadline: BigNumber.from(9),
    rawParams: '10',
  }

  const extraProperties = {
    extra1: '11',
    extra2: '12',
  }

  describe('reduceToQuery', () => {
    it('reduces a RouterQuery with extra properties', () => {
      const query = reduceToQuery({ ...routerQuery, ...extraProperties })
      expect(query).toEqual(routerQuery)
    })

    it('reduces a CCTPRouterQuery with extra properties', () => {
      const query = reduceToQuery({ ...cctpRouterQuery, ...extraProperties })
      expect(query).toEqual(cctpRouterQuery)
    })
  })

  describe('narrowToRouterQuery', () => {
    it('narrows a Query with swapAdapter', () => {
      const query = routerQuery as Query
      const narrowed = narrowToRouterQuery(query)
      expect(narrowed).toEqual(routerQuery)
    })

    it('throws if swapAdapter is undefined', () => {
      const query = cctpRouterQuery as Query
      expect(() => narrowToRouterQuery(query)).toThrow(
        'swapAdapter is undefined'
      )
    })
  })

  describe('narrowToCCTPRouterQuery', () => {
    it('narrows a Query with routerAdapter', () => {
      const query = cctpRouterQuery as Query
      const narrowed = narrowToCCTPRouterQuery(query)
      expect(narrowed).toEqual(cctpRouterQuery)
    })

    it('throws if routerAdapter is undefined', () => {
      const query = routerQuery as Query
      expect(() => narrowToCCTPRouterQuery(query)).toThrow(
        'routerAdapter is undefined'
      )
    })
  })

  describe('modifyDeadline', () => {
    describe('RouterQuery', () => {
      it('modifies the deadline', () => {
        const query = modifyDeadline(routerQuery, BigNumber.from(42))
        expect(query).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(3),
          deadline: BigNumber.from(42),
          rawParams: '5',
        })
      })

      it('does not modify the original query', () => {
        modifyDeadline(routerQuery, BigNumber.from(42))
        expect(routerQuery).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(3),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })
    })

    describe('CCTPRouterQuery', () => {
      it('modifies the deadline', () => {
        const query = modifyDeadline(cctpRouterQuery, BigNumber.from(42))
        expect(query).toEqual({
          routerAdapter: '6',
          tokenOut: '7',
          minAmountOut: BigNumber.from(8),
          deadline: BigNumber.from(42),
          rawParams: '10',
        })
      })

      it('does not modify the original query', () => {
        modifyDeadline(cctpRouterQuery, BigNumber.from(42))
        expect(cctpRouterQuery).toEqual({
          routerAdapter: '6',
          tokenOut: '7',
          minAmountOut: BigNumber.from(8),
          deadline: BigNumber.from(9),
          rawParams: '10',
        })
      })
    })
  })

  describe('applySlippage', () => {
    describe('RouterQuery', () => {
      // 1M in 18 decimals
      const query: RouterQuery = {
        swapAdapter: '1',
        tokenOut: '2',
        minAmountOut: BigNumber.from(10).pow(18).mul(1_000_000),
        deadline: BigNumber.from(4),
        rawParams: '5',
      }

      it('applies 0% slippage', () => {
        const newQuery = applySlippage(query, 0, 10000)
        expect(newQuery).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(18).mul(1_000_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('applies 0.5% slippage', () => {
        // 50 bips
        const newQuery = applySlippage(query, 50, 10000)
        expect(newQuery).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(18).mul(995_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('applies 10% slippage', () => {
        const newQuery = applySlippage(query, 10, 100)
        expect(newQuery).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(18).mul(900_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('applies 100% slippage', () => {
        const newQuery = applySlippage(query, 1, 1)
        expect(newQuery).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(0),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('rounds down', () => {
        const queryPlusOne = {
          ...query,
          minAmountOut: query.minAmountOut.add(1),
        }
        const newQuery = applySlippage(queryPlusOne, 50, 10000)
        expect(newQuery).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(18).mul(995_000).add(1),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('does not modify the original query', () => {
        applySlippage(query, 50, 10000)
        expect(query).toEqual({
          swapAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(18).mul(1_000_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })
    })

    describe('CCTPRouterQuery', () => {
      // 1M in 6 decimals
      const query: CCTPRouterQuery = {
        routerAdapter: '1',
        tokenOut: '2',
        minAmountOut: BigNumber.from(10).pow(6).mul(1_000_000),
        deadline: BigNumber.from(4),
        rawParams: '5',
      }

      it('applies 0% slippage', () => {
        const newQuery = applySlippage(query, 0, 10000)
        expect(newQuery).toEqual({
          routerAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(6).mul(1_000_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('applies 0.5% slippage', () => {
        // 50 bips
        const newQuery = applySlippage(query, 50, 10000)
        expect(newQuery).toEqual({
          routerAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(6).mul(995_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('applies 10% slippage', () => {
        const newQuery = applySlippage(query, 10, 100)
        expect(newQuery).toEqual({
          routerAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(6).mul(900_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('applies 100% slippage', () => {
        const newQuery = applySlippage(query, 1, 1)
        expect(newQuery).toEqual({
          routerAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(0),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('rounds down', () => {
        const queryPlusOne = {
          ...query,
          minAmountOut: query.minAmountOut.add(1),
        }
        const newQuery = applySlippage(queryPlusOne, 50, 10000)
        expect(newQuery).toEqual({
          routerAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(6).mul(995_000).add(1),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })

      it('does not modify the original query', () => {
        applySlippage(query, 50, 10000)
        expect(query).toEqual({
          routerAdapter: '1',
          tokenOut: '2',
          minAmountOut: BigNumber.from(10).pow(6).mul(1_000_000),
          deadline: BigNumber.from(4),
          rawParams: '5',
        })
      })
    })

    describe('errors', () => {
      it('throws if slippage denominator is zero', () => {
        expect(() => applySlippage(routerQuery, 1, 0)).toThrow(
          'Slippage denominator cannot be zero'
        )
      })

      it('throws if slippage numerator is negative', () => {
        expect(() => applySlippage(routerQuery, -1, 1)).toThrow(
          'Slippage numerator cannot be negative'
        )
      })

      it('throws if slippage numerator is greater than denominator', () => {
        expect(() => applySlippage(routerQuery, 2, 1)).toThrow(
          'Slippage cannot be greater than 1'
        )
      })
    })
  })

  describe('applySlippageInBips parity', () => {
    // 1M in 18 decimals
    const query: RouterQuery = {
      swapAdapter: '1',
      tokenOut: '2',
      minAmountOut: BigNumber.from(10).pow(18).mul(1_000_000),
      deadline: BigNumber.from(4),
      rawParams: '5',
    }

    it('applies 0% slippage', () => {
      const newQuery = applySlippage(query, 0, 10000)
      const newQueryInBips = applySlippageInBips(query, 0)
      expect(newQuery).toEqual(newQueryInBips)
    })

    it('applies 0.5% slippage', () => {
      // 50 bips
      const newQuery = applySlippage(query, 50, 10000)
      const newQueryInBips = applySlippageInBips(query, 50)
      expect(newQuery).toEqual(newQueryInBips)
    })

    it('applies 10% slippage', () => {
      const newQuery = applySlippage(query, 10, 100)
      const newQueryInBips = applySlippageInBips(query, 1000)
      expect(newQuery).toEqual(newQueryInBips)
    })

    it('applies 100% slippage', () => {
      const newQuery = applySlippage(query, 1, 1)
      const newQueryInBips = applySlippageInBips(query, 10000)
      expect(newQuery).toEqual(newQueryInBips)
    })

    it('rounds down', () => {
      const queryPlusOne = {
        ...query,
        minAmountOut: query.minAmountOut.add(1),
      }
      const newQuery = applySlippage(queryPlusOne, 50, 10000)
      const newQueryInBips = applySlippageInBips(queryPlusOne, 50)
      expect(newQuery).toEqual(newQueryInBips)
    })

    it('does not modify the original query', () => {
      applySlippageInBips(query, 50)
      expect(query).toEqual({
        swapAdapter: '1',
        tokenOut: '2',
        minAmountOut: BigNumber.from(10).pow(18).mul(1_000_000),
        deadline: BigNumber.from(4),
        rawParams: '5',
      })
    })

    it('throws if basis points are negative', () => {
      expect(() => applySlippageInBips(routerQuery, -1)).toThrow(
        'Slippage numerator cannot be negative'
      )
    })

    it('throws if basis points are greater than 10000', () => {
      expect(() => applySlippageInBips(routerQuery, 10001)).toThrow(
        'Slippage cannot be greater than 1'
      )
    })
  })

  it('createNoSwapQuery', () => {
    const query = createNoSwapQuery('1', BigNumber.from(2))
    expect(query).toEqual({
      routerAdapter: '0x0000000000000000000000000000000000000000',
      tokenOut: '1',
      minAmountOut: BigNumber.from(2),
      deadline: BigNumber.from(0),
      rawParams: '0x',
    })
  })
})
