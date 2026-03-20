import reducer, {
  EMPTY_BRIDGE_QUOTE,
  FetchState,
} from '@/state/slices/bridgeQuote/reducer'
import { fetchBridgeQuote } from '@/state/slices/bridgeQuote/hooks'

const token = {
  addresses: { 1: '0x1', 2: '0x2' },
  decimals: { 1: 6, 2: 6 },
}

const createFetchBridgeQuoteArgs = (requestId: number) => ({
  originChainId: 1,
  destinationChainId: 2,
  originToken: token as any,
  destinationToken: token as any,
  amount: 1n,
  debouncedInputAmount: '1',
  synapseSDK: {},
  requestId,
  pausedModules: [],
  timestamp: requestId,
  connectedAddress: undefined,
})

const createStoredQuote = (requestId: number) => ({
  ...EMPTY_BRIDGE_QUOTE,
  outputAmount: 1n,
  outputAmountString: '1.00',
  delta: 1n,
  requestId,
  timestamp: requestId,
})

const createPendingAction = (requestId: number) =>
  ({
    type: fetchBridgeQuote.pending.type,
    meta: {
      arg: createFetchBridgeQuoteArgs(requestId),
      requestId: `thunk-${requestId}`,
    },
  } as any)

const createFulfilledAction = (requestId: number) =>
  ({
    type: fetchBridgeQuote.fulfilled.type,
    payload: createStoredQuote(requestId),
    meta: {
      arg: createFetchBridgeQuoteArgs(requestId),
      requestId: `thunk-${requestId}`,
    },
  } as any)

const createRejectedAction = (
  requestId: number,
  errorMessage: string = `request ${requestId} failed`
) =>
  ({
    type: fetchBridgeQuote.rejected.type,
    payload: errorMessage,
    error: { message: errorMessage },
    meta: {
      arg: createFetchBridgeQuoteArgs(requestId),
      requestId: `thunk-${requestId}`,
      rejectedWithValue: true,
      requestStatus: 'rejected',
      aborted: false,
      condition: false,
    },
  } as any)

describe('bridgeQuote reducer request de-duping', () => {
  it('keeps the newer request pending when an older request fulfills late', () => {
    let state = reducer(undefined, { type: '@@INIT' })

    state = reducer(state, createPendingAction(1))
    state = reducer(state, createPendingAction(2))
    state = reducer(state, createFulfilledAction(1))

    expect(state.currentRequestId).toBe(2)
    expect(state.status).toBe(FetchState.LOADING)
    expect(state.isLoading).toBe(true)
    expect(state.error).toBeNull()
    expect(state.bridgeQuote).toEqual(EMPTY_BRIDGE_QUOTE)

    state = reducer(state, createFulfilledAction(2))

    expect(state.currentRequestId).toBeNull()
    expect(state.status).toBe(FetchState.VALID)
    expect(state.isLoading).toBe(false)
    expect(state.error).toBeNull()
    expect(state.bridgeQuote).toEqual(createStoredQuote(2))
  })

  it('keeps the newer request pending when an older request rejects late', () => {
    let state = reducer(undefined, { type: '@@INIT' })

    state = reducer(state, createPendingAction(1))
    state = reducer(state, createPendingAction(2))
    state = reducer(state, createRejectedAction(1))

    expect(state.currentRequestId).toBe(2)
    expect(state.status).toBe(FetchState.LOADING)
    expect(state.isLoading).toBe(true)
    expect(state.error).toBeNull()
    expect(state.bridgeQuote).toEqual(EMPTY_BRIDGE_QUOTE)

    state = reducer(state, createFulfilledAction(2))

    expect(state.currentRequestId).toBeNull()
    expect(state.status).toBe(FetchState.VALID)
    expect(state.isLoading).toBe(false)
    expect(state.error).toBeNull()
    expect(state.bridgeQuote).toEqual(createStoredQuote(2))
  })

  it('preserves the newer quote when an older request fulfills after the newer one', () => {
    let state = reducer(undefined, { type: '@@INIT' })

    state = reducer(state, createPendingAction(1))
    state = reducer(state, createPendingAction(2))
    state = reducer(state, createFulfilledAction(2))
    state = reducer(state, createFulfilledAction(1))

    expect(state.currentRequestId).toBeNull()
    expect(state.status).toBe(FetchState.VALID)
    expect(state.isLoading).toBe(false)
    expect(state.error).toBeNull()
    expect(state.bridgeQuote).toEqual(createStoredQuote(2))
  })

  it('preserves the newer quote when an older request rejects after the newer one', () => {
    let state = reducer(undefined, { type: '@@INIT' })

    state = reducer(state, createPendingAction(1))
    state = reducer(state, createPendingAction(2))
    state = reducer(state, createFulfilledAction(2))
    state = reducer(state, createRejectedAction(1))

    expect(state.currentRequestId).toBeNull()
    expect(state.status).toBe(FetchState.VALID)
    expect(state.isLoading).toBe(false)
    expect(state.error).toBeNull()
    expect(state.bridgeQuote).toEqual(createStoredQuote(2))
  })
})
