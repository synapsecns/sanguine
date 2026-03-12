import { BigNumber, providers } from 'ethers'
import { mock } from 'jest-mock-extended'

import { MEDIAN_TIME_BLOCK, SupportedChainId } from '../constants'
import { BridgeTokenCandidate, GetBridgeRouteV2Parameters } from '../module'
import { EngineID, decodeZapData } from '../swap'
import { SynapseBridgeAdapterModule } from './synapseBridgeAdapterModule'
import { SynapseBridgeAdapterModuleSet } from './synapseBridgeAdapterModuleSet'
import { SwapEngineRoute } from '../swap/models'

const ETH_TOKEN = '0x00000000000000000000000000000000000000e1'
const OP_TOKEN = '0x00000000000000000000000000000000000000b2'
const OTHER_TOKEN = '0x00000000000000000000000000000000000000f1'
const SENDER = '0x00000000000000000000000000000000000000a1'
const RECIPIENT = '0x00000000000000000000000000000000000000b1'

const createNoOpRoute = (
  amount: BigNumber = BigNumber.from(1000)
): SwapEngineRoute => ({
  engineID: EngineID.NoOp,
  engineName: EngineID[EngineID.NoOp],
  chainId: SupportedChainId.ETH,
  fromToken: ETH_TOKEN,
  fromAmount: amount,
  toToken: ETH_TOKEN,
  expectedToAmount: amount,
  minToAmount: amount,
  steps: [],
})

describe('SynapseBridgeAdapterModuleSet', () => {
  const ethProvider = mock<providers.Provider>()
  const opProvider = mock<providers.Provider>()
  const moduleSet = new SynapseBridgeAdapterModuleSet([
    {
      chainId: SupportedChainId.ETH,
      provider: ethProvider,
    },
    {
      chainId: SupportedChainId.OPTIMISM,
      provider: opProvider,
    },
  ])
  const originModule = moduleSet.modules[
    SupportedChainId.ETH
  ] as SynapseBridgeAdapterModule

  const bridgeToken: BridgeTokenCandidate = {
    originChainId: SupportedChainId.ETH,
    destChainId: SupportedChainId.OPTIMISM,
    originToken: ETH_TOKEN,
    destToken: OP_TOKEN,
  }

  const getRouteParams = (
    overrides: Partial<GetBridgeRouteV2Parameters> = {}
  ): GetBridgeRouteV2Parameters => ({
    originSwapRoute: createNoOpRoute(),
    bridgeToken,
    toToken: OP_TOKEN,
    allowMultipleTxs: false,
    ...overrides,
  })

  beforeEach(() => {
    jest.restoreAllMocks()
    jest.clearAllMocks()
  })

  it('returns no candidates when the SBA module is missing on either chain', async () => {
    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.BASE,
        fromToken: ETH_TOKEN,
      })
    ).resolves.toEqual([])
  })

  it('returns no candidates when getRemoteAddress returns zero address', async () => {
    jest
      .spyOn(originModule, 'getRemoteAddress')
      .mockResolvedValue('0x0000000000000000000000000000000000000000')

    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: ETH_TOKEN,
      })
    ).resolves.toEqual([])
  })

  it('returns no candidates when the requested destination token mismatches', async () => {
    jest.spyOn(originModule, 'getRemoteAddress').mockResolvedValue(OP_TOKEN)

    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: ETH_TOKEN,
        toToken: OTHER_TOKEN,
      })
    ).resolves.toEqual([])
  })

  it('returns the direct origin and destination token candidate for supported pairs', async () => {
    jest.spyOn(originModule, 'getRemoteAddress').mockResolvedValue(OP_TOKEN)

    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: ETH_TOKEN,
      })
    ).resolves.toEqual([bridgeToken])
  })

  it('rejects bridge routes that require origin swap steps', async () => {
    jest.spyOn(originModule, 'getRemoteAddress').mockResolvedValue(OP_TOKEN)

    await expect(
      moduleSet.getBridgeRouteV2({
        ...getRouteParams(),
        originSwapRoute: {
          ...createNoOpRoute(),
          steps: [
            {
              token: ETH_TOKEN,
              amount: BigNumber.from(1000),
              msgValue: BigNumber.from(0),
              zapData: '0x',
            },
          ],
        },
      })
    ).resolves.toBeUndefined()
  })

  it('forwards the native fee and keeps the SBA bridge step 1:1', async () => {
    jest.spyOn(originModule, 'getRemoteAddress').mockResolvedValue(OP_TOKEN)
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    jest.spyOn(originModule, 'getEstimatedTime').mockResolvedValue(undefined)

    const route = await moduleSet.getBridgeRouteV2(getRouteParams())

    expect(route).toMatchObject({
      toToken: OP_TOKEN,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(1000),
      nativeFee: BigNumber.from(77),
    })
  })

  it('encodes bridgeERC20 zap data with the correct amount position', async () => {
    jest.spyOn(originModule, 'getRemoteAddress').mockResolvedValue(OP_TOKEN)
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    jest.spyOn(originModule, 'getEstimatedTime').mockResolvedValue(undefined)

    const route = await moduleSet.getBridgeRouteV2(
      getRouteParams({
        fromSender: SENDER,
        toRecipient: RECIPIENT,
      })
    )
    const decodedZap = decodeZapData(route!.zapData!)
    const decodedCall =
      SynapseBridgeAdapterModule.sbaInterface.decodeFunctionData(
        'bridgeERC20',
        decodedZap.payload!
      )

    expect(decodedZap.target).toEqual(originModule.address.toLowerCase())
    expect(decodedZap.amountPosition).toEqual(100)
    expect(decodedCall.dstEid).toEqual(30111)
    expect(decodedCall.to.toLowerCase()).toEqual(RECIPIENT.toLowerCase())
    expect(decodedCall.token.toLowerCase()).toEqual(ETH_TOKEN.toLowerCase())
    expect(decodedCall.amount).toEqual(BigNumber.from(1000))
    expect(decodedCall.gasLimit).toEqual(BigNumber.from(200000))
  })

  it('passes tx id through from the concrete module', async () => {
    const txHash = '0x1234'
    await expect(
      moduleSet.getSynapseTxId(SupportedChainId.ETH, txHash)
    ).resolves.toEqual(txHash)
  })

  it('reports LayerZero completion false on routine failure', async () => {
    global.fetch = jest.fn().mockRejectedValue(new Error('network')) as any
    await expect(
      moduleSet.getBridgeTxStatus(SupportedChainId.OPTIMISM, '0x1234')
    ).resolves.toBe(false)
  })

  it('uses fallback ETA until a live pathway estimate is cached', async () => {
    const expectedFallbackEta = Math.ceil(
      64 * MEDIAN_TIME_BLOCK[SupportedChainId.ETH] +
        3 * MEDIAN_TIME_BLOCK[SupportedChainId.OPTIMISM]
    )
    expect(
      moduleSet.getEstimatedTime(
        SupportedChainId.ETH,
        SupportedChainId.OPTIMISM
      )
    ).toEqual(expectedFallbackEta)

    jest.spyOn(originModule, 'getRemoteAddress').mockResolvedValue(OP_TOKEN)
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    const estimatedTimeSpy = jest
      .spyOn(originModule, 'getEstimatedTime')
      .mockResolvedValue(42)

    await moduleSet.getBridgeRouteV2(getRouteParams())
    await moduleSet.getBridgeRouteV2(getRouteParams())

    expect(
      moduleSet.getEstimatedTime(
        SupportedChainId.ETH,
        SupportedChainId.OPTIMISM
      )
    ).toEqual(42)
    expect(estimatedTimeSpy).toHaveBeenCalledTimes(1)
  })
})
