import { BigNumber } from 'ethers'

import {
  CCTP_V2_DOMAIN_MAP,
  CCTP_V2_FORWARD_SERVICE_FEE_USDC,
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
        forwardFee: { basic: 50 },
      },
      {
        finalityThreshold: 3000,
        minimumFee: 20,
        forwardFee: { basic: 150 },
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

  it('computes maxFee as minimumFee budget plus forwarding fee budget', async () => {
    const moduleSet = makeModuleSet()
    mockGetBurnUSDCFees.mockResolvedValueOnce([
      {
        finalityThreshold: 1200,
        minimumFee: 25,
        forwardFee: { low: 120, high: 350 },
      },
    ])

    const route = await moduleSet.getBridgeRouteV2(makeRouteParams(1_000_000))

    expect(route).toBeDefined()
    expect(route!.expectedToAmount).toEqual(BigNumber.from(997_150))
    expect(route!.minToAmount).toEqual(BigNumber.from(997_150))
    expect(decodeBurnCalldata(route!.zapData!).maxFee).toEqual(
      BigNumber.from(2850)
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

  it('returns no quote when origin min amount is not greater than maxFee', async () => {
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
      CCTP_V2_FORWARD_SERVICE_FEE_USDC.NON_ETH
    )
    expect(fallbackRoute!.expectedToAmount).toEqual(
      BigNumber.from(1_000_000).sub(expectedMaxFeeWithFallback)
    )
    expect(decodeBurnCalldata(fallbackRoute!.zapData!).maxFee).toEqual(
      expectedMaxFeeWithFallback
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
})
