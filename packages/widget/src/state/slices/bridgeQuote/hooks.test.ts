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

const createSdkQuote = ({
  moduleNames = ['SynapseRFQ'],
  nativeFee,
}: {
  moduleNames?: string[]
  nativeFee: unknown
}) => ({
  expectedToAmount: '1500000',
  routerAddress: '0xrouter',
  estimatedTime: 45,
  moduleNames,
  nativeFee,
  tx: undefined,
})

const createFetchBridgeQuoteArgs = ({
  nativeFee,
  pausedModules = [],
  quotes = [createSdkQuote({ nativeFee })],
}: {
  nativeFee: unknown
  pausedModules?: any[]
  quotes?: Array<ReturnType<typeof createSdkQuote>>
}) => ({
  originChainId: 1,
  destinationChainId: 2,
  originToken: token as any,
  destinationToken: token as any,
  amount: 1_000_000n,
  debouncedInputAmount: '1',
  synapseSDK: {
    bridgeV2: jest.fn().mockResolvedValue(quotes),
  },
  requestId: 7,
  pausedModules,
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

const runFetchBridgeQuote = async ({
  nativeFee,
  pausedModules,
  quotes,
}: {
  nativeFee: unknown
  pausedModules?: any[]
  quotes?: Array<ReturnType<typeof createSdkQuote>>
}) => {
  const args = createFetchBridgeQuoteArgs({
    nativeFee,
    pausedModules,
    quotes,
  })
  const dispatch = jest.fn()

  let state = reducer(undefined, { type: '@@INIT' })
  state = reducer(state, createPendingAction(args))

  const action = await fetchBridgeQuote(args)(
    dispatch,
    () => ({} as any),
    undefined
  )

  state = reducer(state, action as any)

  return { action, state, bridgeV2: args.synapseSDK.bridgeV2 }
}

describe('fetchBridgeQuote quote filtering', () => {
  it.each([
    ['zero string fees', '0', 0n],
    ['positive string fees', '77', 77n],
  ])(
    'stores %s as the expected bigint',
    async (_scenario, nativeFee, expected) => {
      const { action, state, bridgeV2 } = await runFetchBridgeQuote({
        nativeFee,
      })

      expect(fetchBridgeQuote.fulfilled.match(action)).toBe(true)
      expect(bridgeV2).toHaveBeenCalledTimes(1)
      expect(state.bridgeQuote.id).toBeNull()
      expect(state.bridgeQuote.nativeFee).toBe(expected)
    }
  )

  it.each([
    ['missing fees', undefined],
    ['malformed fees', 'abc'],
    ['negative fees', '-1'],
  ])('rejects quotes with %s', async (_scenario, nativeFee) => {
    const { action, state, bridgeV2 } = await runFetchBridgeQuote({
      nativeFee,
    })

    expect(fetchBridgeQuote.rejected.match(action)).toBe(true)
    expect(action.payload).toBe('No active bridge quotes available')
    expect(bridgeV2).toHaveBeenCalledTimes(1)
    expect(state.status).toBe('invalid')
    expect(state.bridgeQuote.nativeFee).toBe(0n)
  })

  it('keeps the widget RFQ-first selection behavior after filtering invalid quotes', async () => {
    const { action, state } = await runFetchBridgeQuote({
      nativeFee: '77',
      quotes: [
        createSdkQuote({
          moduleNames: ['SynapseBridge'],
          nativeFee: '55',
        }),
        createSdkQuote({
          moduleNames: ['SynapseRFQ'],
          nativeFee: '77',
        }),
      ],
    })

    expect(fetchBridgeQuote.fulfilled.match(action)).toBe(true)
    expect(state.bridgeQuote.bridgeModuleName).toBe('SynapseRFQ')
    expect(state.bridgeQuote.nativeFee).toBe(77n)
  })

  it.each([
    ['malformed nativeFee', 'abc'],
    ['negative nativeFee', '-1'],
  ])(
    'skips an invalid RFQ quote with %s and selects the next valid active quote',
    async (_scenario, nativeFee) => {
      const { action, state } = await runFetchBridgeQuote({
        nativeFee,
        quotes: [
          createSdkQuote({
            moduleNames: ['SynapseRFQ'],
            nativeFee,
          }),
          createSdkQuote({
            moduleNames: ['SynapseBridge'],
            nativeFee: '12',
          }),
        ],
      })

      expect(fetchBridgeQuote.fulfilled.match(action)).toBe(true)
      expect(state.bridgeQuote.bridgeModuleName).toBe('SynapseBridge')
      expect(state.bridgeQuote.nativeFee).toBe(12n)
    }
  )

  it('filters paused bridge modules before selecting the active quote', async () => {
    const { action, state } = await runFetchBridgeQuote({
      nativeFee: '12',
      pausedModules: [
        {
          bridgeModuleName: 'SynapseRFQ',
          chainId: 1,
        },
      ],
      quotes: [
        createSdkQuote({
          moduleNames: ['SynapseRFQ'],
          nativeFee: '77',
        }),
        createSdkQuote({
          moduleNames: ['SynapseBridge'],
          nativeFee: '12',
        }),
      ],
    })

    expect(fetchBridgeQuote.fulfilled.match(action)).toBe(true)
    expect(state.bridgeQuote.bridgeModuleName).toBe('SynapseBridge')
    expect(state.bridgeQuote.nativeFee).toBe(12n)
  })

  it('filters destination chain-specific paused bridge modules', async () => {
    const { action, state } = await runFetchBridgeQuote({
      nativeFee: '12',
      pausedModules: [
        {
          bridgeModuleName: 'SynapseRFQ',
          toChainId: 2,
        },
      ],
      quotes: [
        createSdkQuote({
          moduleNames: ['SynapseRFQ'],
          nativeFee: '77',
        }),
        createSdkQuote({
          moduleNames: ['SynapseBridge'],
          nativeFee: '12',
        }),
      ],
    })

    expect(fetchBridgeQuote.fulfilled.match(action)).toBe(true)
    expect(state.bridgeQuote.bridgeModuleName).toBe('SynapseBridge')
    expect(state.bridgeQuote.nativeFee).toBe(12n)
  })
})
