import { BigNumber, providers } from 'ethers'
import { mock } from 'jest-mock-extended'

import {
  MEDIAN_TIME_BLOCK,
  SupportedChainId,
  SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP,
  SYNAPSE_INTENT_ROUTER_ADDRESS_MAP,
  TOKEN_ZAP_V1_ADDRESS_MAP,
} from '../constants'
import { BridgeTokenCandidate, GetBridgeRouteV2Parameters } from '../module'
import { EngineID, decodeZapData } from '../swap'
import { getSbaRemoteToken, getSbaSupportedTokens } from './supportedTokens'
import { SwapEngineRoute } from '../swap/models'
import { ETH_NATIVE_TOKEN_ADDRESS } from '../utils'
import { SynapseBridgeAdapterModule } from './synapseBridgeAdapterModule'
import { SynapseBridgeAdapterModuleSet } from './synapseBridgeAdapterModuleSet'

const ETH_AGEUR = '0x1a7e4e63778B4f12a199C062f3eFdD288afCBce8'
const OP_AGEUR = '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22'
const ETH_DOG = '0xBAac2B4491727D78D2b78815144570b9f2Fe8899'
const BSC_DOG = '0xaA88C603d142C371eA0eAC8756123c5805EdeE03'
const BSC_BUSD = '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56'
const ARB_GMX = '0xfc5A1A6EB076a2C7aD06eD22C90d7E710E35ad0a'
const ETH_NETH = '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'
const OP_NETH = '0x809DC529f07651bD43A172e8dB6f4a7a0d771036'
const OP_SYN = '0x5A5fFf6F753d7C11A56A52FE47a177a87e431655'
const HARMONY_NETH = '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB'
const KLAYTN_NETH = '0xCD6f29dC9Ca217d0973d3D21bF58eDd3CA871a86'
const DFK_KLAY = '0x97855Ba65aa7ed2F65Ed832a776537268158B78a'
const KLAYTN_KLAY = '0x5819b6af194A78511c79C85Ea68D2377a7e9335f'
const OTHER_TOKEN = '0x00000000000000000000000000000000000000f1'
const SENDER = '0x00000000000000000000000000000000000000a1'
const RECIPIENT = '0x00000000000000000000000000000000000000b1'

const createNoOpRoute = (
  token: string,
  amount: BigNumber = BigNumber.from(1000),
  chainId: SupportedChainId = SupportedChainId.ETH
): SwapEngineRoute => ({
  engineID: EngineID.NoOp,
  engineName: EngineID[EngineID.NoOp],
  chainId,
  fromToken: token,
  fromAmount: amount,
  toToken: token,
  expectedToAmount: amount,
  minToAmount: amount,
  steps: [],
})

describe('SynapseBridgeAdapterModuleSet', () => {
  const ethProvider = mock<providers.Provider>()
  const opProvider = mock<providers.Provider>()
  const baseProvider = mock<providers.Provider>()
  const dfkProvider = mock<providers.Provider>()
  const harmonyProvider = mock<providers.Provider>()
  const klaytnProvider = mock<providers.Provider>()
  const moduleSet = new SynapseBridgeAdapterModuleSet([
    {
      chainId: SupportedChainId.ETH,
      provider: ethProvider,
    },
    {
      chainId: SupportedChainId.OPTIMISM,
      provider: opProvider,
    },
    {
      chainId: SupportedChainId.BASE,
      provider: baseProvider,
    },
    {
      chainId: SupportedChainId.DFK,
      provider: dfkProvider,
    },
    {
      chainId: SupportedChainId.HARMONY,
      provider: harmonyProvider,
    },
    {
      chainId: SupportedChainId.KLAYTN,
      provider: klaytnProvider,
    },
  ])
  const originModule = moduleSet.modules[
    SupportedChainId.ETH
  ] as SynapseBridgeAdapterModule
  const dfkModule = moduleSet.modules[
    SupportedChainId.DFK
  ] as SynapseBridgeAdapterModule

  const bridgeToken: BridgeTokenCandidate = {
    originChainId: SupportedChainId.ETH,
    destChainId: SupportedChainId.OPTIMISM,
    originToken: ETH_NETH,
    destToken: OP_NETH,
  }

  const harmonyToKlaytnBridgeToken: BridgeTokenCandidate = {
    originChainId: SupportedChainId.HARMONY,
    destChainId: SupportedChainId.KLAYTN,
    originToken: HARMONY_NETH,
    destToken: KLAYTN_NETH,
  }

  const dfkToKlaytnWrappedNativeBridgeToken: BridgeTokenCandidate = {
    originChainId: SupportedChainId.DFK,
    destChainId: SupportedChainId.KLAYTN,
    originToken: DFK_KLAY,
    destToken: ETH_NATIVE_TOKEN_ADDRESS,
  }

  const dfkToKlaytnWrappedBridgeToken: BridgeTokenCandidate = {
    originChainId: SupportedChainId.DFK,
    destChainId: SupportedChainId.KLAYTN,
    originToken: DFK_KLAY,
    destToken: KLAYTN_KLAY,
  }

  const getRouteParams = (
    overrides: Partial<GetBridgeRouteV2Parameters> = {}
  ): GetBridgeRouteV2Parameters => ({
    originSwapRoute: createNoOpRoute(ETH_NETH),
    bridgeToken,
    toToken: OP_NETH,
    allowMultipleTxs: false,
    ...overrides,
  })

  beforeEach(() => {
    jest.restoreAllMocks()
    jest.clearAllMocks()
  })

  it('uses normalized chain IDs for representative mappings and excludes BUSD and GMX', () => {
    expect(
      getSbaRemoteToken(
        SupportedChainId.ETH,
        ETH_AGEUR,
        SupportedChainId.OPTIMISM
      )
    ).toEqual(OP_AGEUR)
    expect(
      getSbaRemoteToken(SupportedChainId.BSC, BSC_DOG, SupportedChainId.ETH)
    ).toEqual(ETH_DOG)
    expect(
      getSbaRemoteToken(SupportedChainId.DFK, DFK_KLAY, SupportedChainId.KLAYTN)
    ).toEqual(ETH_NATIVE_TOKEN_ADDRESS)
    expect(
      getSbaRemoteToken(
        SupportedChainId.HARMONY,
        HARMONY_NETH,
        SupportedChainId.KLAYTN
      )
    ).toEqual(KLAYTN_NETH)
    expect(
      getSbaRemoteToken(
        SupportedChainId.HARMONY,
        HARMONY_NETH,
        SupportedChainId.ETH
      )
    ).toEqual(ETH_NATIVE_TOKEN_ADDRESS)
    expect(
      getSbaRemoteToken(SupportedChainId.BSC, BSC_BUSD, SupportedChainId.ETH)
    ).toBeUndefined()
    expect(
      getSbaRemoteToken(
        SupportedChainId.ARBITRUM,
        ARB_GMX,
        SupportedChainId.AVALANCHE
      )
    ).toBeUndefined()
  })

  it('returns no candidates when the SBA module is missing on either chain', async () => {
    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.AVALANCHE,
        fromToken: HARMONY_NETH,
      })
    ).resolves.toEqual([])
  })

  it('returns no candidates when the temporary SBA bridge allowlist excludes either chain', async () => {
    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.ETH,
        toChainId: SupportedChainId.OPTIMISM,
        fromToken: ETH_NETH,
      })
    ).resolves.toEqual([])
  })

  it('returns all artifact-backed candidates for supported pairs and ignores fromToken as a filter', async () => {
    const directCandidates = await moduleSet.getBridgeTokenCandidates({
      fromChainId: SupportedChainId.HARMONY,
      toChainId: SupportedChainId.KLAYTN,
      fromToken: OTHER_TOKEN,
    })
    const nativeCandidates = await moduleSet.getBridgeTokenCandidates({
      fromChainId: SupportedChainId.HARMONY,
      toChainId: SupportedChainId.KLAYTN,
      fromToken: ETH_NATIVE_TOKEN_ADDRESS,
    })

    expect(directCandidates).toEqual(nativeCandidates)
    expect(directCandidates).toEqual(
      getSbaSupportedTokens(SupportedChainId.HARMONY, SupportedChainId.KLAYTN)
    )
    expect(directCandidates).toContainEqual(harmonyToKlaytnBridgeToken)
  })

  it('filters candidates by the exact artifact-mapped destination token', async () => {
    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.KLAYTN,
        fromToken: OTHER_TOKEN,
        toToken: KLAYTN_NETH,
      })
    ).resolves.toEqual([harmonyToKlaytnBridgeToken])
    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.HARMONY,
        toChainId: SupportedChainId.KLAYTN,
        fromToken: OTHER_TOKEN,
        toToken: OTHER_TOKEN,
      })
    ).resolves.toEqual([])
  })

  it('treats destination wrapped-native mappings as effective native outputs', async () => {
    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.DFK,
        toChainId: SupportedChainId.KLAYTN,
        fromToken: OTHER_TOKEN,
        toToken: ETH_NATIVE_TOKEN_ADDRESS,
      })
    ).resolves.toEqual([dfkToKlaytnWrappedNativeBridgeToken])
    await expect(
      moduleSet.getBridgeTokenCandidates({
        fromChainId: SupportedChainId.DFK,
        toChainId: SupportedChainId.KLAYTN,
        fromToken: OTHER_TOKEN,
        toToken: KLAYTN_KLAY,
      })
    ).resolves.toEqual([])
  })

  it('matches the helper lookup for candidate discovery', () => {
    expect(
      getSbaSupportedTokens(
        SupportedChainId.HARMONY,
        SupportedChainId.KLAYTN,
        KLAYTN_NETH
      )
    ).toEqual([harmonyToKlaytnBridgeToken])
  })

  it('accepts one-step origin routes and preserves their minimum amount', async () => {
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    jest.spyOn(originModule, 'getEstimatedTime').mockResolvedValue(undefined)

    const route = await moduleSet.getBridgeRouteV2({
      ...getRouteParams(),
      originSwapRoute: {
        ...createNoOpRoute(ETH_NETH, BigNumber.from(1100)),
        fromToken: OTHER_TOKEN,
        expectedToAmount: BigNumber.from(1000),
        minToAmount: BigNumber.from(925),
        steps: [
          {
            token: OTHER_TOKEN,
            amount: BigNumber.from(1100),
            msgValue: BigNumber.from(0),
            zapData: '0x1234',
          },
        ],
      },
    })

    expect(route).toMatchObject({
      toToken: OP_NETH,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(925),
      nativeFee: BigNumber.from(77),
    })
  })

  it('accepts multi-step origin routes and preserves their minimum amount', async () => {
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    jest.spyOn(originModule, 'getEstimatedTime').mockResolvedValue(undefined)

    const route = await moduleSet.getBridgeRouteV2({
      ...getRouteParams(),
      originSwapRoute: {
        ...createNoOpRoute(ETH_NETH, BigNumber.from(1200)),
        fromToken: OTHER_TOKEN,
        expectedToAmount: BigNumber.from(1000),
        minToAmount: BigNumber.from(900),
        steps: [
          {
            token: OTHER_TOKEN,
            amount: BigNumber.from(1200),
            msgValue: BigNumber.from(0),
            zapData: '0x1111',
          },
          {
            token: ETH_NETH,
            amount: BigNumber.from(1000),
            msgValue: BigNumber.from(0),
            zapData: '0x2222',
          },
        ],
      },
    })

    expect(route).toMatchObject({
      toToken: OP_NETH,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(900),
      nativeFee: BigNumber.from(77),
    })
  })

  it('accepts native-origin swap routes through the generic pipeline and preserves their minimum amount', async () => {
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    jest.spyOn(originModule, 'getEstimatedTime').mockResolvedValue(undefined)

    const route = await moduleSet.getBridgeRouteV2({
      ...getRouteParams(),
      originSwapRoute: {
        ...createNoOpRoute(ETH_NETH, BigNumber.from(1100)),
        engineID: EngineID.DefaultPools,
        engineName: EngineID[EngineID.DefaultPools],
        fromToken: ETH_NATIVE_TOKEN_ADDRESS,
        expectedToAmount: BigNumber.from(1000),
        minToAmount: BigNumber.from(950),
        steps: [
          {
            token: ETH_NATIVE_TOKEN_ADDRESS,
            amount: BigNumber.from(1100),
            msgValue: BigNumber.from(1100),
            zapData: '0x1234',
          },
        ],
      },
    })

    expect(route).toMatchObject({
      toToken: OP_NETH,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(950),
      nativeFee: BigNumber.from(77),
    })
  })

  it('drops routes when the selected bridge token disagrees with the artifact snapshot', async () => {
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))

    await expect(
      moduleSet.getBridgeRouteV2(
        getRouteParams({
          bridgeToken: {
            ...bridgeToken,
            destToken: OP_SYN,
          },
          toToken: OP_SYN,
        })
      )
    ).resolves.toBeUndefined()
  })

  it('surfaces native output when the artifact destination token is wrapped native', async () => {
    jest.spyOn(dfkModule, 'getNativeFee').mockResolvedValue(BigNumber.from(33))
    jest.spyOn(dfkModule, 'getEstimatedTime').mockResolvedValue(undefined)

    const route = await moduleSet.getBridgeRouteV2({
      originSwapRoute: createNoOpRoute(
        DFK_KLAY,
        BigNumber.from(1000),
        SupportedChainId.DFK
      ),
      bridgeToken: dfkToKlaytnWrappedNativeBridgeToken,
      toToken: ETH_NATIVE_TOKEN_ADDRESS,
      allowMultipleTxs: false,
    })

    expect(route).toMatchObject({
      bridgeToken: dfkToKlaytnWrappedNativeBridgeToken,
      toToken: ETH_NATIVE_TOKEN_ADDRESS,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(1000),
      nativeFee: BigNumber.from(33),
    })
  })

  it('does not quote wrapped-native destinations as ERC20 outputs', async () => {
    jest.spyOn(dfkModule, 'getNativeFee').mockResolvedValue(BigNumber.from(33))

    await expect(
      moduleSet.getBridgeRouteV2({
        originSwapRoute: createNoOpRoute(
          DFK_KLAY,
          BigNumber.from(1000),
          SupportedChainId.DFK
        ),
        bridgeToken: dfkToKlaytnWrappedNativeBridgeToken,
        toToken: dfkToKlaytnWrappedBridgeToken.destToken,
        allowMultipleTxs: false,
      })
    ).resolves.toBeUndefined()
  })

  it('forwards the native fee and keeps the SBA bridge step 1:1', async () => {
    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    jest.spyOn(originModule, 'getEstimatedTime').mockResolvedValue(undefined)

    const route = await moduleSet.getBridgeRouteV2(getRouteParams())

    expect(route).toMatchObject({
      toToken: OP_NETH,
      expectedToAmount: BigNumber.from(1000),
      minToAmount: BigNumber.from(1000),
      nativeFee: BigNumber.from(77),
    })
  })

  it('encodes bridgeERC20 zap data with the correct amount position', async () => {
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
    expect(decodedCall.token.toLowerCase()).toEqual(ETH_NETH.toLowerCase())
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
      64 * MEDIAN_TIME_BLOCK[SupportedChainId.ETH] + 60
    )
    expect(
      moduleSet.getEstimatedTime(
        SupportedChainId.ETH,
        SupportedChainId.OPTIMISM
      )
    ).toEqual(expectedFallbackEta)

    jest
      .spyOn(originModule, 'getNativeFee')
      .mockResolvedValue(BigNumber.from(77))
    const estimatedTimeSpy = jest
      .spyOn(originModule, 'getEstimatedTime')
      .mockResolvedValue(42)

    const firstRoute = await moduleSet.getBridgeRouteV2(getRouteParams())
    const secondRoute = await moduleSet.getBridgeRouteV2(getRouteParams())

    expect(
      moduleSet.getEstimatedTime(
        SupportedChainId.ETH,
        SupportedChainId.OPTIMISM
      )
    ).toEqual(42)
    expect(firstRoute?.estimatedTime).toEqual(42)
    expect(secondRoute?.estimatedTime).toEqual(42)
    expect(estimatedTimeSpy).toHaveBeenCalledTimes(1)
  })

  it('includes DFK, Harmony, and Klaytn in the shared-intent address maps', () => {
    expect(SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.DFK]).toEqual(
      '0x512000a034E154908Efb1eC48579F4ffDb000512'
    )
    expect(SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.HARMONY]).toEqual(
      '0x512000a034E154908Efb1eC48579F4ffDb000512'
    )
    expect(SYNAPSE_INTENT_ROUTER_ADDRESS_MAP[SupportedChainId.KLAYTN]).toEqual(
      '0x512000a034E154908Efb1eC48579F4ffDb000512'
    )
    expect(TOKEN_ZAP_V1_ADDRESS_MAP[SupportedChainId.DFK]).toEqual(
      '0x2aAaa9b71E479e6e2De7E091b09D61C25D2AAAa9'
    )
    expect(
      SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP[SupportedChainId.HARMONY]
    ).toEqual('0x519519a57a6Ea930f87e3436b6ea113A990fF519')
  })
})
