/** @jest-environment node */

import reducer, { resetBridgeQuote } from '@/slices/bridgeQuote/reducer'
import { fetchBridgeQuote } from '@/slices/bridgeQuote/thunks'
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge'

jest.mock('../actions/getErc20TokenAllowance', () => ({
  getErc20TokenAllowance: jest.fn(),
}))
jest.mock('../contexts/SegmentAnalyticsProvider', () => ({
  segmentAnalyticsEvent: jest.fn(),
}))

const args = {} as any
const completedQuote = { ...EMPTY_BRIDGE_QUOTE, id: 'completed' } as any

describe('bridge quote request ordering', () => {
  it('ignores fulfilled and rejected results from stale requests', () => {
    const pending = reducer(
      undefined,
      fetchBridgeQuote.pending('current-request', args)
    )
    const afterFulfilled = reducer(
      pending,
      fetchBridgeQuote.fulfilled(completedQuote, 'stale-request', args)
    )
    const afterRejected = reducer(
      afterFulfilled,
      fetchBridgeQuote.rejected(
        new Error('stale failure'),
        'stale-request',
        args
      )
    )

    expect(afterRejected).toMatchObject({
      bridgeQuote: EMPTY_BRIDGE_QUOTE,
      currentRequestId: 'current-request',
      isLoading: true,
    })
  })

  it('ignores a request result after the quote is reset', () => {
    const pending = reducer(
      undefined,
      fetchBridgeQuote.pending('cancelled-request', args)
    )
    const reset = reducer(pending, resetBridgeQuote())
    const afterFulfilled = reducer(
      reset,
      fetchBridgeQuote.fulfilled(completedQuote, 'cancelled-request', args)
    )
    const afterRejected = reducer(
      reset,
      fetchBridgeQuote.rejected(
        new Error('cancelled failure'),
        'cancelled-request',
        args
      )
    )

    expect(afterFulfilled).toMatchObject({
      bridgeQuote: EMPTY_BRIDGE_QUOTE,
      currentRequestId: undefined,
      isLoading: false,
    })
    expect(afterRejected).toEqual(afterFulfilled)
  })
})
