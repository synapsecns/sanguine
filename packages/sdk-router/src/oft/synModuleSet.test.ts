import { BigNumber, providers, utils } from 'ethers'

import {
  HYPERCORE_CHAIN_ID,
  isChainIdSupported,
  SYN_ADDRESS_MAP,
  SYN_COMPOSER_ADDRESS,
  SYN_OFT_ADDRESS_MAP,
  TOKEN_ZAP_V1_ADDRESS_MAP,
} from '../constants'
import { SynapseSDK } from '../sdk'
import { SynapseIntentRouterSet } from '../sir/synapseIntentRouterSet'
import { AMOUNT_NOT_PRESENT, decodeZapData, NoOpEngine } from '../swap'
import { SYN_COMPOSER_ABI } from './synModule'
import { UsdtModule } from './usdtModule'

const sender = '0x1111111111111111111111111111111111111111'
const recipient = '0x2222222222222222222222222222222222222222'
const unit = BigNumber.from(10).pow(12)
const coreUnit = BigNumber.from(10).pow(10)
const amount = utils.parseEther('1.123456789')
const rounded = amount.div(unit).mul(unit)
const nativeFee = BigNumber.from(12345)
const composerInterface = new utils.Interface(SYN_COMPOSER_ABI)
const oftInterface = UsdtModule.oftInterface

// Mock the RPC boundary, preserving ABI encoding/decoding, routing, SIR and zap construction.
const setup = () => {
  let active = true
  let capacity = true
  const quoteSendParams: any[] = []
  const providersByChain = [1, 999].map((chainId) => {
    const provider = new providers.StaticJsonRpcProvider(
      'http://localhost:1',
      chainId
    )
    jest.spyOn(provider, 'call').mockImplementation(async (tx) => {
      if (
        (await tx.to)?.toLowerCase() ===
          SYN_ADDRESS_MAP[chainId].toLowerCase() &&
        (await tx.data) === '0x313ce567'
      ) {
        return utils.defaultAbiCoder.encode(['uint8'], [18])
      }
      if ((await tx.to)?.toLowerCase() === SYN_COMPOSER_ADDRESS.toLowerCase()) {
        const parsed = composerInterface.parseTransaction({
          data: String(await tx.data),
        })
        if (parsed.name === 'coreUserExists') {
          return composerInterface.encodeFunctionResult(parsed.name, [[active]])
        }
        if (parsed.name === 'ERC20_ASSET_BRIDGE') {
          return composerInterface.encodeFunctionResult(parsed.name, [
            recipient,
          ])
        }
        if (parsed.name === 'quoteHyperCoreAmount') {
          if (!capacity) {
            throw new Error('TransferAmtExceedsAssetBridgeBalance')
          }
          const value = BigNumber.from(parsed.args[3])
          return composerInterface.encodeFunctionResult(parsed.name, [
            [value, value.div(coreUnit), 1000000000],
          ])
        }
      }
      if (
        (await tx.to)?.toLowerCase() ===
        SYN_OFT_ADDRESS_MAP[chainId].toLowerCase()
      ) {
        const parsed = oftInterface.parseTransaction({
          data: String(await tx.data),
        })
        if (parsed.name === 'quoteOFT') {
          const value = BigNumber.from(parsed.args[0][2]).div(unit).mul(unit)
          return oftInterface.encodeFunctionResult(parsed.name, [
            [0, utils.parseEther('100')],
            [],
            [value, value],
          ])
        }
        if (parsed.name === 'quoteSend') {
          quoteSendParams.push(parsed.args[0])
          return oftInterface.encodeFunctionResult(parsed.name, [
            [nativeFee, 0],
          ])
        }
      }
      throw new Error(`Unexpected RPC call ${String(await tx.data)}`)
    })
    return provider
  })
  const sdk = new SynapseSDK([1, 999], providersByChain)
  expect(sdk.allModuleSets).toContain(sdk.synModuleSet)
  sdk.allModuleSets = [sdk.synModuleSet]
  // Other swap engines cannot improve a SYN->SYN identity route.
  const noOp = new NoOpEngine()
  ;(sdk.swapEngineSet as any).engines = { [noOp.id]: noOp }
  const params = (fromChainId: number, toChainId: number) => ({
    fromChainId,
    toChainId,
    fromToken: SYN_ADDRESS_MAP[fromChainId],
    toToken: SYN_ADDRESS_MAP[toChainId],
    fromAmount: amount.toString(),
    fromSender: sender,
    toRecipient: recipient,
  })
  return {
    sdk,
    params,
    quoteSendParams,
    setActive: (value: boolean) => {
      active = value
    },
    setCapacity: (value: boolean) => {
      capacity = value
    },
  }
}

afterEach(() => jest.restoreAllMocks())

it.each([
  [1, 999],
  [999, 1],
  [1, HYPERCORE_CHAIN_ID],
])(
  'quotes and builds executable SYN route %i -> %i through public bridgeV2',
  async (from, to) => {
    const { sdk, params, quoteSendParams } = setup()
    const [quote] = await sdk.bridgeV2(params(from, to))
    expect(quote.moduleNames).toEqual(['SYN'])
    expect(quote.toChainId).toBe(to)
    expect(quote.toToken).toBe(SYN_ADDRESS_MAP[to])
    const destinationAmount =
      to === HYPERCORE_CHAIN_ID ? rounded.div(coreUnit) : rounded
    expect(quote.expectedToAmount).toBe(destinationAmount.toString())
    expect(quote.minToAmount).toBe(destinationAmount.toString())
    expect(quote.nativeFee).toBe(nativeFee.toString())
    expect(quote.tx!.value).toBe(nativeFee.toString())
    const intent = SynapseIntentRouterSet.sirInterface.parseTransaction({
      data: quote.tx!.data,
    })
    expect(intent.name).toBe('completeIntentWithBalanceChecks')
    expect(intent.args.zapRecipient).toBe(TOKEN_ZAP_V1_ADDRESS_MAP[from])
    expect(intent.args.amountIn).toEqual(amount)
    const zap = decodeZapData(intent.args.steps[0].zapData)
    expect(zap.amountPosition).toBe(AMOUNT_NOT_PRESENT)
    expect(zap.forwardTo?.toLowerCase()).toBe(sender)
    expect(zap.finalToken?.toLowerCase()).toBe(
      params(from, to).fromToken.toLowerCase()
    )
    const send = oftInterface.parseTransaction({ data: zap.payload! })
    expect(send.args[0]).toEqual(quoteSendParams[0])
    expect(send.args[0][2]).toEqual(amount)
    expect(send.args[0][3]).toEqual(rounded)
    expect(send.args[1][0]).toEqual(nativeFee)
    expect(send.args[2]).toBe(sender)
    // Quotes and sends rely on the adapter's enforced SEND / SEND_AND_CALL options.
    expect(send.args[0][4]).toBe('0x')
    expect(send.args[0][1].toLowerCase()).toBe(
      utils
        .hexZeroPad(
          to === HYPERCORE_CHAIN_ID ? SYN_COMPOSER_ADDRESS : recipient,
          32
        )
        .toLowerCase()
    )
    expect(send.args[0][5]).toBe(
      to === HYPERCORE_CHAIN_ID
        ? utils.defaultAbiCoder.encode(['uint256', 'address'], [0, recipient])
        : '0x'
    )
  }
)

it('returns estimates without calldata when disconnected, including HyperCore metadata without a virtual RPC', async () => {
  const { sdk, params } = setup()
  const [quote] = await sdk.bridgeV2({
    ...params(1, HYPERCORE_CHAIN_ID),
    fromSender: undefined,
    toRecipient: undefined,
  })
  expect(quote.tx).toBeUndefined()
  expect(quote.expectedToAmount).toBe(rounded.div(coreUnit).toString())
  expect(quote.toToken).toBe('0xf5f05eb8b9aa92365465f06daf5889c9')
  expect(isChainIdSupported(HYPERCORE_CHAIN_ID)).toBe(false)
  expect(sdk.providers[HYPERCORE_CHAIN_ID]).toBeUndefined()
  await expect(
    sdk.tokenMetadataFetcher.getTokenDecimals(
      HYPERCORE_CHAIN_ID,
      SYN_ADDRESS_MAP[HYPERCORE_CHAIN_ID]
    )
  ).resolves.toBe(8)
})

it('returns native HyperCore token units and decimals through public intent without changing OFT calldata', async () => {
  const { sdk, params } = setup()
  const [quote] = await sdk.intent(params(1, HYPERCORE_CHAIN_ID))
  expect(quote.toChainId).toBe(1337)
  expect(quote.toToken).toBe('0xf5f05eb8b9aa92365465f06daf5889c9')
  expect(quote.fromAmount).toBe(amount.toString())
  expect(quote.expectedToAmount).toBe('112345600')
  expect(quote.minToAmount).toBe('112345600')
  expect(quote.steps).toHaveLength(1)
  const [step] = quote.steps
  expect(step.fromTokenDecimals).toBe(18)
  expect(step.toTokenDecimals).toBe(8)
  expect(step.toToken).toBe(quote.toToken)
  expect(step.expectedToAmount).toBe(quote.expectedToAmount)
  expect(step.minToAmount).toBe(quote.minToAmount)
  const intent = SynapseIntentRouterSet.sirInterface.parseTransaction({
    data: step.tx!.data,
  })
  expect(intent.args.amountIn).toEqual(amount)
  const zap = decodeZapData(intent.args.steps[0].zapData)
  const send = oftInterface.parseTransaction({ data: zap.payload! })
  expect(send.args[0][2]).toEqual(amount)
  expect(send.args[0][3]).toEqual(rounded)
  expect(send.args[0][4]).toBe('0x')
  expect(sdk.providers[HYPERCORE_CHAIN_ID]).toBeUndefined()
})

it('rejects the former 998 destination and the linked EVM token as a Core token ID', async () => {
  const { sdk, params } = setup()
  await expect(
    sdk.bridgeV2({ ...params(1, HYPERCORE_CHAIN_ID), toChainId: 998 })
  ).resolves.toEqual([])
  await expect(
    sdk.bridgeV2({
      ...params(1, HYPERCORE_CHAIN_ID),
      toToken: SYN_ADDRESS_MAP[999],
    })
  ).resolves.toEqual([])
})

it('rejects an inactive HyperCore recipient with an actionable error before building calldata', async () => {
  const { sdk, params, setActive, quoteSendParams } = setup()
  setActive(false)
  await expect(sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID))).rejects.toThrow(
    'Activate the recipient account on HyperCore'
  )
  expect(quoteSendParams).toHaveLength(0)
})

it('omits zero-after-dust amounts and capacity failures', async () => {
  const { sdk, params, setCapacity } = setup()
  await expect(
    sdk.bridgeV2({ ...params(1, 999), fromAmount: unit.sub(1).toString() })
  ).resolves.toEqual([])
  setCapacity(false)
  await expect(sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID))).resolves.toEqual([])
})

it('does not request a dust refund when the complete input is bridged', async () => {
  const { sdk, params } = setup()
  const [quote] = await sdk.bridgeV2({
    ...params(1, 999),
    fromAmount: rounded.toString(),
  })
  const intent = SynapseIntentRouterSet.sirInterface.parseTransaction({
    data: quote.tx!.data,
  })
  const zap = decodeZapData(intent.args.steps[0].zapData)
  expect(zap.forwardTo).toBe(utils.getAddress('0x' + '0'.repeat(40)))
})

it.each([
  [999, HYPERCORE_CHAIN_ID],
  [HYPERCORE_CHAIN_ID, 999],
  [HYPERCORE_CHAIN_ID, 1],
])('excludes unsupported path %i -> %i', async (from, to) => {
  const { sdk, params } = setup()
  await expect(sdk.bridgeV2(params(from, to))).resolves.toEqual([])
})

it('does not expose non-SYN tokens, origin swaps, or legacy V1 quotes', async () => {
  const { sdk, params } = setup()
  await expect(
    sdk.bridgeV2({ ...params(1, 999), fromToken: sender })
  ).resolves.toEqual([])
  await expect(
    sdk.bridgeV2({ ...params(1, 999), toToken: recipient })
  ).resolves.toEqual([])
  await expect(
    sdk.allBridgeQuotes(
      1,
      999,
      SYN_ADDRESS_MAP[1],
      SYN_ADDRESS_MAP[999],
      amount
    )
  ).resolves.toEqual([])
})
