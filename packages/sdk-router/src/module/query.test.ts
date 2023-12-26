import { BigNumber } from '@ethersproject/bignumber'

import {
  RouterQuery,
  CCTPRouterQuery,
  Query,
  reduceToQuery,
  narrowToRouterQuery,
  narrowToCCTPRouterQuery,
  modifyDeadline,
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
})
