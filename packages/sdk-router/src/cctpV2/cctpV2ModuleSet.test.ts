import { BigNumber } from 'ethers'

import {
  CCTP_V2_DOMAIN_MAP,
  CCTP_V2_FORWARD_SERVICE_FEE_USDC,
  CCTP_V2_SUPPORTED_CHAIN_IDS,
  CCTP_V2_USDC_ADDRESS_MAP,
  SupportedChainId,
} from '../constants'
import { decodeZapData, EngineID } from '../swap'
import { getBurnUSDCFees } from './api'
import { CircleCCTPV2Module } from './cctpV2Module'
import { CircleCCTPV2ModuleSet } from './cctpV2ModuleSet'

jest.mock('./api', () => ({
  getBurnUSDCFees: jest.fn(),
}))

const mockGetBurnUSDCFees = getBurnUSDCFees as jest.MockedFunction<
  typeof getBurnUSDCFees
>

const ORIGIN_CHAIN_ID = SupportedChainId.ETH
const DEST_CHAIN_ID = SupportedChainId.ARBITRUM
const ORIGIN_TOKEN = CCTP_V2_USDC_ADDRESS_MAP[ORIGIN_CHAIN_ID]
const DEST_TOKEN = CCTP_V2_USDC_ADDRESS_MAP[DEST_CHAIN_ID]
const ETH_STANDARD_ESTIMATED_TIME = 1020
const ETH_FAST_ESTIMATED_TIME = 600

const makeModuleSet = () =>
  new CircleCCTPV2ModuleSet([
    { chainId: ORIGIN_CHAIN_ID, provider: {} as any },
    { chainId: DEST_CHAIN_ID, provider: {} as any },
  ])

const makeRouteParams = (
  expectedAmountIn = 1_000_000,
  minAmountIn = expectedAmountIn
) =>
  ({
    bridgeToken: {
      originChainId: ORIGIN_CHAIN_ID,
      destChainId: DEST_CHAIN_ID,
      originToken: ORIGIN_TOKEN,
      destToken: DEST_TOKEN,
    },
    originSwapRoute: {
      engineID: EngineID.NoOp,
      engineName: EngineID[EngineID.NoOp],
      chainId: ORIGIN_CHAIN_ID,
      fromToken: ORIGIN_TOKEN,
      fromAmount: BigNumber.from(expectedAmountIn),
      toToken: ORIGIN_TOKEN,
      expectedToAmount: BigNumber.from(expectedAmountIn),
      minToAmount: BigNumber.from(minAmountIn),
      steps: [],
    },
    toToken: DEST_TOKEN,
    allowMultipleTxs: false,
    fromSender: '0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
    toRecipient: '0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb',
  } as any)

const decodeBurnCalldata = (zapData: string) => {
  const decodedZapData = decodeZapData(zapData)
  return CircleCCTPV2Module.tokenMessengerV2Interface.decodeFunctionData(
    'depositForBurnWithHook',
    decodedZapData.payload!
  )
}

const getBurnMaxFeeWithSlippage = (maxFee: BigNumber): BigNumber =>
  maxFee.mul(11).add(9).div(10)

describe('CircleCCTPV2ModuleSet', () => {
  beforeEach(() => {
    mockGetBurnUSDCFees.mockReset()
  })

  it('filters candidates to supported CCTP V2 pairs', async () => {
    const moduleSet = makeModuleSet()

    const candidates = await moduleSet.getBridgeTokenCandidates({
      fromChainId: ORIGIN_CHAIN_ID,
      toChainId: DEST_CHAIN_ID,
      fromToken: ORIGIN_TOKEN,
      toToken: DEST_TOKEN,
    })

    expect(candidates).toEqual([
      {
        originChainId: ORIGIN_CHAIN_ID,
        destChainId: DEST_CHAIN_ID,
        originToken: ORIGIN_TOKEN,
        destToken: DEST_TOKEN,
      },
    ])

    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: ORIGIN_CHAIN_ID,
        toChainId: SupportedChainId.BSC,
        fromToken: ORIGIN_TOKEN,
        toToken: DEST_TOKEN,
      })
    ).resolves.toEqual([])

    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: ORIGIN_CHAIN_ID,
        toChainId: DEST_CHAIN_ID,
        fromToken: ORIGIN_TOKEN,
        toToken: '0x9999999999999999999999999999999999999999',
      })
    ).resolves.toEqual([])
  })

  it('selects max finality threshold from fee response', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 5,
        forwardFee: { high: 50 },
      },
      {
        finalityThreshold: 3000,
        minimumFee: 20,
        forwardFee: { high: 150 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams())

    expect(route).toBeDefined()
    expect(mockGetBurnUSDCFees).toHaveBeenCalledWith(
      CCTP_V2_DOMAIN_MAP[ORIGIN_CHAIN_ID],
      CCTP_V2_DOMAIN_MAP[DEST_CHAIN_ID]
    )
    expect(decodeBurnCalldata(route!.zapData!).minFinalityThreshold).toBe(3000)
  })

  it('uses standard estimated time for standard finality threshold', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 2000,
        minimumFee: 20,
        forwardFee: { high: 150 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams())

    expect(route).toBeDefined()
    expect(route!.estimatedTime).toBe(ETH_STANDARD_ESTIMATED_TIME)
    expect(route!.estimatedTime).toBeGreaterThan(0)
  })

  it('uses fast estimated time for fast finality threshold', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 20,
        forwardFee: { high: 150 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams())

    expect(route).toBeDefined()
    expect(route!.estimatedTime).toBe(ETH_FAST_ESTIMATED_TIME)
    expect(route!.estimatedTime).toBeGreaterThan(0)
  })

  it('uses fast estimated time when finality threshold is exactly 1000', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
        forwardFee: { high: 150 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams())

    expect(route).toBeDefined()
    expect(route!.estimatedTime).toBe(ETH_FAST_ESTIMATED_TIME)
    expect(route!.estimatedTime).toBeGreaterThan(0)
  })

  it('uses fast estimated time when finality threshold is 1999', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1999,
        minimumFee: 20,
        forwardFee: { high: 150 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams())

    expect(route).toBeDefined()
    expect(route!.estimatedTime).toBe(ETH_FAST_ESTIMATED_TIME)
    expect(route!.estimatedTime).toBeGreaterThan(0)
  })

  it('falls back to chain estimated time when finality threshold is unmappable', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 999,
        minimumFee: 20,
        forwardFee: { high: 150 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams())

    expect(route).toBeDefined()
    expect(route!.estimatedTime).toBe(ETH_STANDARD_ESTIMATED_TIME)
    expect(route!.estimatedTime).toBeGreaterThan(0)
  })

  it('returns non-zero estimated times for all supported CCTP V2 chain ids', () => {
    const moduleSet = makeModuleSet()

    CCTP_V2_SUPPORTED_CHAIN_IDS.forEach((chainId) => {
      expect(moduleSet.getEstimatedTime(chainId)).toBeGreaterThan(0)
    })
  })

  it('prefers med forwarding tier over high when both are present', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 25,
        forwardFee: { med: 150, high: 350, low: 120 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))

    expect(route).toBeDefined()
    expect(route!.expectedToAmount).toEqual(BigNumber.from(997_350))
    expect(route!.minToAmount).toEqual(BigNumber.from(997_350))
    const quoteMaxFee = BigNumber.from(2650)
    expect(decodeBurnCalldata(route!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(quoteMaxFee)
    )
  })

  it('prefers medium forwarding tier when med is absent', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 25,
        forwardFee: { medium: 170, high: 350, low: 120 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))

    expect(route).toBeDefined()
    expect(route!.expectedToAmount).toEqual(BigNumber.from(997_330))
    expect(route!.minToAmount).toEqual(BigNumber.from(997_330))
    const quoteMaxFee = BigNumber.from(2670)
    expect(decodeBurnCalldata(route!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(quoteMaxFee)
    )
  })

  it('falls back to high, then low, when middle tier is missing', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 25,
        forwardFee: { high: 350, low: 120 },
      },
    ])

    const highRoute = await moduleSet.getBridgeRouteV2(
      makeRouteParams(1_000_000)
    )

    expect(highRoute).toBeDefined()
    expect(highRoute!.expectedToAmount).toEqual(BigNumber.from(997_150))
    expect(highRoute!.minToAmount).toEqual(BigNumber.from(997_150))
    const highQuoteMaxFee = BigNumber.from(2850)
    expect(decodeBurnCalldata(highRoute!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(highQuoteMaxFee)
    )

    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 25,
        forwardFee: { high: 350.5, low: 120 },
      },
    ])

    const lowRoute = await moduleSet.getBridgeRouteV2(
      makeRouteParams(1_000_000)
    )

    expect(lowRoute).toBeDefined()
    expect(lowRoute!.expectedToAmount).toEqual(BigNumber.from(997_380))
    expect(lowRoute!.minToAmount).toEqual(BigNumber.from(997_380))
    const lowQuoteMaxFee = BigNumber.from(2620)
    expect(decodeBurnCalldata(lowRoute!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(lowQuoteMaxFee)
    )
  })

  it('derives destination minToAmount from origin minToAmount', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 25,
        forwardFee: { low: 120, high: 350 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(
      makeRouteParams(1_000_000, 900_000)
    )

    expect(route).toBeDefined()
    expect(route!.expectedToAmount).toEqual(BigNumber.from(997_150))
    expect(route!.minToAmount).toEqual(BigNumber.from(897_150))
  })

  it('supports fractional minimumFee bps with deterministic integer math', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 1.3,
        forwardFee: { low: 120, high: 350 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams(1_000_001))

    expect(route).toBeDefined()
    const quoteMaxFee = BigNumber.from(481)
    expect(route!.expectedToAmount).toEqual(BigNumber.from(999_520))
    expect(route!.minToAmount).toEqual(BigNumber.from(999_520))
    expect(decodeBurnCalldata(route!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(quoteMaxFee)
    )
  })

  it('applies ceiling to fractional minimumFee protocol fee', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 1.5,
        forwardFee: { high: 0 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams(6_667))

    expect(route).toBeDefined()
    const quoteMaxFee = BigNumber.from(2)
    expect(route!.expectedToAmount).toEqual(BigNumber.from(6_665))
    expect(route!.minToAmount).toEqual(BigNumber.from(6_665))
    expect(decodeBurnCalldata(route!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(quoteMaxFee)
    )
  })

  it('fails closed for scientific notation minimumFee values', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 1e-7,
        forwardFee: { high: 0 },
      },
    ])

    await expect(
      moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))
    ).resolves.toBeUndefined()
  })

  it('returns no quote when origin min amount is not greater than burn maxFee', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 25,
        forwardFee: { low: 120, high: 350 },
      },
    ])

    await expect(
      moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000, 2_000))
    ).resolves.toBeUndefined()
  })

  it('returns no quote when fee API returns no data', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce(null)

    await expect(
      moduleSet.getBridgeRouteV2(makeRouteParams())
    ).resolves.toBeUndefined()
  })

  it('applies forwarding fee fallback only when forwardFee is missing', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
      },
    ])

    const fallbackRoute = await moduleSet.getBridgeRouteV2(
      makeRouteParams(1_000_000)
    )
    expect(fallbackRoute).toBeDefined()

    const expectedMaxFeeWithFallback = BigNumber.from(2000).add(
      CCTP_V2_FORWARD_SERVICE_FEE_USDC.defaultFee
    )
    expect(fallbackRoute!.expectedToAmount).toEqual(
      BigNumber.from(1_000_000).sub(expectedMaxFeeWithFallback)
    )
    expect(decodeBurnCalldata(fallbackRoute!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(expectedMaxFeeWithFallback)
    )

    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
        forwardFee: {} as Record<string, number>,
      },
    ])

    await expect(
      moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))
    ).resolves.toBeUndefined()

    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
        forwardFee: { med: -1, high: 250, low: 100 },
      },
    ])

    await expect(
      moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))
    ).resolves.toBeUndefined()

    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
        forwardFee: { medium: -1, high: 250, low: 100 },
      },
    ])

    await expect(
      moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))
    ).resolves.toBeUndefined()

    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
        forwardFee: null as unknown as Record<string, number>,
      },
    ])

    await expect(
      moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))
    ).resolves.toBeUndefined()
  })

  it('allows destination token fallback in allowMultipleTxs mode', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
        forwardFee: { low: 100, high: 200 },
      },
    ])
    const params = makeRouteParams()
    params.allowMultipleTxs = true
    params.toToken = '0x9999999999999999999999999999999999999999'

    const route = await moduleSet.getBridgeRouteV2(params)

    expect(route).toBeDefined()
    expect(route!.toToken).toBe(DEST_TOKEN)
  })

  it('uses per-chain forwarding fee override for ETH destination', async () => {
    const originChainId = SupportedChainId.ARBITRUM
    const destChainId = SupportedChainId.ETH
    const originToken = CCTP_V2_USDC_ADDRESS_MAP[originChainId]
    const destToken = CCTP_V2_USDC_ADDRESS_MAP[destChainId]
    const moduleSet = new CircleCCTPV2ModuleSet([
      { chainId: originChainId, provider: {} as any },
      { chainId: destChainId, provider: {} as any },
    ])
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1000,
        minimumFee: 20,
      },
    ])

    const route = await moduleSet.getBridgeRouteV2({
      bridgeToken: {
        originChainId,
        destChainId,
        originToken,
        destToken,
      },
      originSwapRoute: {
        engineID: EngineID.NoOp,
        engineName: EngineID[EngineID.NoOp],
        chainId: originChainId,
        fromToken: originToken,
        fromAmount: BigNumber.from(2_000_000),
        toToken: originToken,
        expectedToAmount: BigNumber.from(2_000_000),
        minToAmount: BigNumber.from(2_000_000),
        steps: [],
      },
      toToken: destToken,
      allowMultipleTxs: false,
      fromSender: '0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
      toRecipient: '0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb',
    } as any)

    const quoteMaxFee = BigNumber.from(4000).add(
      CCTP_V2_FORWARD_SERVICE_FEE_USDC.perChainOverrides[SupportedChainId.ETH]
    )
    expect(route).toBeDefined()
    expect(route!.expectedToAmount).toEqual(
      BigNumber.from(2_000_000).sub(quoteMaxFee)
    )
    expect(decodeBurnCalldata(route!.zapData!).maxFee).toEqual(
      getBurnMaxFeeWithSlippage(quoteMaxFee)
    )
  })
})
