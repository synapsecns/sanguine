import reducer from '@/state/slices/bridgeQuote/reducer'
import { fetchBridgeQuote } from '@/state/slices/bridgeQuote/hooks'

const token = {
  addresses: { 1: '0x1', 2: '0x2' },
  decimals: { 1: 6, 2: 6 },
  symbol: 'USDC',
  name: 'USD Coin',
  swapableType: 'token',
  color: '#fff',
  priorityRank: 1,
  routeSymbol: 'USDC',
  imgUrl: '',
}

const createSdkQuote = (nativeFee: unknown) => ({
  expectedToAmount: '1500000',
  routerAddress: '0xrouter',
  estimatedTime: 45,
  moduleNames: ['SynapseRFQ'],
  nativeFee,
  tx: undefined,
})

const createFetchBridgeQuoteArgs = (nativeFee: unknown) => ({
  originChainId: 1,
  destinationChainId: 2,
  originToken: token as any,
  destinationToken: token as any,
  amount: 1_000_000n,
  debouncedInputAmount: '1',
  synapseSDK: {
    bridgeV2: jest.fn().mockResolvedValue([createSdkQuote(nativeFee)]),
  },
  requestId: 7,
  pausedModules: [],
  timestamp: 7,
  connectedAddress: undefined,
})

const createPendingAction = (
  args: ReturnType<typeof createFetchBridgeQuoteArgs>
) =>
  ({
    type: fetchBridgeQuote.pending.type,
    meta: {
      arg: args,
      requestId: `thunk-${args.requestId}`,
    },
  } as any)

const runFetchBridgeQuote = async (nativeFee: unknown) => {
  const args = createFetchBridgeQuoteArgs(nativeFee)
  const dispatch = jest.fn()

  let state = reducer(undefined, { type: '@@INIT' })
  state = reducer(state, createPendingAction(args))

  const action = await fetchBridgeQuote(args)(
    dispatch,
    () => ({} as any),
    undefined
  )

  expect(fetchBridgeQuote.fulfilled.match(action)).toBe(true)

  state = reducer(state, action as any)

  return { action, state, bridgeV2: args.synapseSDK.bridgeV2 }
}

describe('fetchBridgeQuote nativeFee normalization', () => {
  it.each([
    ['zero string fees', '0', 0n],
    ['positive string fees', '77', 77n],
    ['missing fees', undefined, 0n],
    ['malformed fees', 'abc', 0n],
  ])(
    'stores %s as the expected bigint',
    async (_scenario, nativeFee, expected) => {
      const { state, bridgeV2 } = await runFetchBridgeQuote(nativeFee)

      expect(bridgeV2).toHaveBeenCalledTimes(1)
      expect(state.bridgeQuote.nativeFee).toBe(expected)
    }
  )
})
