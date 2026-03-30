import { act, renderHook, waitFor } from '@testing-library/react'

import { useNativeSafeMax } from '@/hooks/useNativeSafeMax'

jest.mock('ethers', () => ({
  ZeroAddress: '0x0000000000000000000000000000000000000000',
}))

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000'

const nativeToken = {
  addresses: { 1: ZERO_ADDRESS },
  decimals: { 1: 18 },
  symbol: 'ETH',
  name: 'Ether',
  swapableType: 'token',
  color: '#fff',
  priorityRank: 1,
  routeSymbol: 'ETH',
  imgUrl: '',
}

const destinationToken = {
  addresses: { 2: '0x2' },
  decimals: { 2: 18 },
  symbol: 'nETH',
  name: 'nETH',
  swapableType: 'token',
  color: '#fff',
  priorityRank: 1,
  routeSymbol: 'NETH',
  imgUrl: '',
}

const createQuote = ({
  moduleNames = ['SynapseRFQ'],
  nativeFee = '0',
  tx = {
    data: '0x1234',
    to: '0xrouter',
    value: '0',
  },
}: {
  moduleNames?: string[]
  nativeFee?: string
  tx?: {
    data?: string
    to?: string
    value?: string | null
  } | null
} = {}) => ({
  moduleNames,
  nativeFee,
  tx,
})

const createProvider = ({
  estimateGasValues = [100_000n, 80_000n],
  feeData = {
    gasPrice: 10_000_000_000n,
    maxFeePerGas: 1_000_000_000n,
  },
}: {
  estimateGasValues?: bigint[]
  feeData?: {
    gasPrice?: bigint | null
    maxFeePerGas?: bigint | null
  }
} = {}) => ({
  estimateGas: jest
    .fn()
    .mockImplementation(() => Promise.resolve(estimateGasValues.shift() ?? 0n)),
  getFeeData: jest.fn().mockResolvedValue(feeData),
})

const createProps = (
  overrides: Partial<Parameters<typeof useNativeSafeMax>[0]> = {}
) => ({
  amountKey: '',
  connectedAddress: '0xabc',
  destinationChainId: 2,
  destinationToken: destinationToken as any,
  isWalletPending: false,
  originChainId: 1,
  originChainProvider: createProvider() as any,
  originToken: nativeToken as any,
  pausedModules: [],
  rawBalance: '1000000000000000000',
  synapseSDK: {
    bridgeV2: jest
      .fn()
      .mockResolvedValue([
        createQuote(),
        createQuote({ moduleNames: ['SynapseBridge'] }),
      ]),
  },
  ...overrides,
})

const createDeferred = <T,>() => {
  let resolve!: (value: T | PromiseLike<T>) => void
  let reject!: (reason?: unknown) => void
  const promise = new Promise<T>((res, rej) => {
    resolve = res
    reject = rej
  })

  return { promise, reject, resolve }
}

describe('useNativeSafeMax', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })

  it('bootstraps with half the balance, refines once, and prefers maxFeePerGas', async () => {
    const bridgeV2 = jest
      .fn()
      .mockResolvedValueOnce([
        createQuote({ moduleNames: ['SynapseBridge'], nativeFee: '1' }),
        createQuote({ nativeFee: '100000000000000000' }),
      ])
      .mockResolvedValueOnce([createQuote({ nativeFee: '200000000000000000' })])
    const originChainProvider = createProvider()
    const initialProps = createProps({
      amountKey: '1',
      originChainProvider: originChainProvider as any,
      synapseSDK: { bridgeV2 },
    })
    const { result } = renderHook(
      (props: Parameters<typeof useNativeSafeMax>[0]) =>
        useNativeSafeMax(props),
      {
        initialProps,
      }
    )

    await waitFor(() => {
      expect(result.current.status).toBe('ready')
    })

    expect(bridgeV2).toHaveBeenNthCalledWith(
      1,
      expect.objectContaining({
        fromAmount: '500000000000000000',
        fromSender: '0xabc',
        toRecipient: '0xabc',
      })
    )
    expect(bridgeV2).toHaveBeenNthCalledWith(
      2,
      expect.objectContaining({
        fromAmount: '899850000000000000',
      })
    )
    expect(originChainProvider.getFeeData).toHaveBeenCalledTimes(2)
    expect(originChainProvider.estimateGas).toHaveBeenNthCalledWith(
      1,
      expect.objectContaining({
        data: '0x1234',
        from: '0xabc',
        to: '0xrouter',
      })
    )
    expect(result.current.amountWei).toBe(799880000000000000n)
    expect(result.current.fillAmount).toBe('0.799880000000000000')
    expect(result.current.isClickable).toBe(true)
    expect(result.current.labelAmount).toBe('0.7998')
  })

  it('does not fall back to the raw balance when quote selection cannot produce an executable transaction', async () => {
    const initialProps = createProps({
      synapseSDK: {
        bridgeV2: jest.fn().mockResolvedValue([createQuote({ tx: null })]),
      },
    })
    const { result } = renderHook(
      (props: Parameters<typeof useNativeSafeMax>[0]) =>
        useNativeSafeMax(props),
      {
        initialProps,
      }
    )

    await waitFor(() => {
      expect(result.current.status).toBe('unavailable')
    })

    expect(result.current.fillAmount).toBeNull()
    expect(result.current.isClickable).toBe(false)
    expect(result.current.labelAmount).toBeNull()
  })

  it('becomes unavailable when gas estimation cannot complete', async () => {
    const originChainProvider = {
      estimateGas: jest.fn().mockRejectedValue(new Error('estimateGas failed')),
      getFeeData: jest.fn().mockResolvedValue({
        maxFeePerGas: 1_000_000_000n,
      }),
    }
    const initialProps = createProps({
      originChainProvider: originChainProvider as any,
    })
    const { result } = renderHook(
      (props: Parameters<typeof useNativeSafeMax>[0]) =>
        useNativeSafeMax(props),
      {
        initialProps,
      }
    )

    await waitFor(() => {
      expect(result.current.status).toBe('unavailable')
    })

    expect(originChainProvider.getFeeData).toHaveBeenCalledTimes(1)
  })

  it('clamps a negative result to zero and keeps the control non-clickable', async () => {
    const originChainProvider = createProvider({
      estimateGasValues: [300_000n],
      feeData: {
        maxFeePerGas: 1_000_000_000n,
      },
    })
    const initialProps = createProps({
      originChainProvider: originChainProvider as any,
      rawBalance: '100000000000000000',
      synapseSDK: {
        bridgeV2: jest
          .fn()
          .mockResolvedValue([createQuote({ nativeFee: '99900000000000000' })]),
      },
    })
    const { result } = renderHook(
      (props: Parameters<typeof useNativeSafeMax>[0]) =>
        useNativeSafeMax(props),
      {
        initialProps,
      }
    )

    await waitFor(() => {
      expect(result.current.status).toBe('ready')
    })

    expect(result.current.amountWei).toBe(0n)
    expect(result.current.fillAmount).toBe('0.0')
    expect(result.current.isClickable).toBe(false)
    expect(result.current.labelAmount).toBe('0.0')
  })

  it('falls back to gasPrice when maxFeePerGas is unavailable', async () => {
    const originChainProvider = createProvider({
      feeData: {
        gasPrice: 2_000_000_000n,
        maxFeePerGas: null,
      },
    })
    const initialProps = createProps({
      originChainProvider: originChainProvider as any,
    })
    const { result } = renderHook(
      (props: Parameters<typeof useNativeSafeMax>[0]) =>
        useNativeSafeMax(props),
      {
        initialProps,
      }
    )

    await waitFor(() => {
      expect(result.current.status).toBe('ready')
    })

    expect(originChainProvider.getFeeData).toHaveBeenCalledTimes(2)
    expect(result.current.amountWei).toBe(999760000000000000n)
    expect(result.current.fillAmount).toBe('0.999760000000000000')
    expect(result.current.labelAmount).toBe('0.9997')
  })

  it('ignores stale async results after the amount key changes', async () => {
    const staleBootstrap = createDeferred<any[]>()
    const bridgeV2 = jest
      .fn()
      .mockImplementationOnce(() => staleBootstrap.promise)
      .mockResolvedValueOnce([createQuote({ nativeFee: '100000000000000000' })])
      .mockResolvedValueOnce([createQuote({ nativeFee: '200000000000000000' })])
    const originChainProvider = createProvider()
    const { result, rerender } = renderHook(
      (props: Parameters<typeof useNativeSafeMax>[0]) =>
        useNativeSafeMax(props),
      {
        initialProps: createProps({
          amountKey: '1',
          originChainProvider: originChainProvider as any,
          synapseSDK: { bridgeV2 },
        }),
      }
    )

    await waitFor(() => {
      expect(result.current.status).toBe('loading')
    })

    rerender(
      createProps({
        amountKey: '2',
        originChainProvider: originChainProvider as any,
        synapseSDK: { bridgeV2 },
      })
    )

    await waitFor(() => {
      expect(result.current.status).toBe('ready')
    })

    expect(result.current.amountWei).toBe(799880000000000000n)
    expect(originChainProvider.estimateGas).toHaveBeenCalledTimes(2)

    await act(async () => {
      staleBootstrap.resolve([createQuote({ nativeFee: '1' })])
      await Promise.resolve()
    })

    expect(result.current.amountWei).toBe(799880000000000000n)
    expect(originChainProvider.estimateGas).toHaveBeenCalledTimes(2)
  })
})
