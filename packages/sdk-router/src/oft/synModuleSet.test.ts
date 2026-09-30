import { BigNumber, constants, providers, utils } from 'ethers'

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
import { quoteSynHyperCore } from './synModule'
import { UsdtModule } from './usdtModule'

const sender = '0x1111111111111111111111111111111111111111'
const recipient = '0x2222222222222222222222222222222222222222'
const unit = BigNumber.from(10).pow(12)
const coreUnit = BigNumber.from(10).pow(10)
const amount = utils.parseEther('1.123456789')
const rounded = amount.div(unit).mul(unit)
const nativeFee = BigNumber.from(12345)
const oftInterface = UsdtModule.oftInterface

// Mock the RPC boundary, preserving ABI encoding/decoding, routing, SIR and zap construction.
const setup = (chainIds = [1, 999]) => {
  const quoteSendParams: any[] = []
  const providersByChain = chainIds.map((chainId) => {
    const provider = new providers.StaticJsonRpcProvider(
      'http://localhost:1',
      chainId
    )
    jest.spyOn(provider, 'call').mockImplementation(async (tx) => {
      if (
        (await tx.to)?.toLowerCase() ===
        SYN_OFT_ADDRESS_MAP[chainId].toLowerCase()
      ) {
        const parsed = oftInterface.parseTransaction({
          data: String(await tx.data),
        })
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
  const sdk = new SynapseSDK(chainIds, providersByChain)
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
    providersByChain,
  }
}

afterEach(() => jest.restoreAllMocks())

it('quotes HyperCore with only an Ethereum provider and one fee read per refresh', async () => {
  const { sdk, params, providersByChain } = setup([1])
  const fetchMock = jest
    .spyOn(global, 'fetch')
    .mockRejectedValue(
      new Error('Quoting must not fetch Core account or balance data')
    )
  for (let i = 0; i < 2; i++) {
    const [quote] = await sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID))
    expect(quote.expectedToAmount).toBe('112345600')
    expect(quote.tx).toBeDefined()
  }
  expect(providersByChain[0].call).toHaveBeenCalledTimes(2)
  expect(sdk.providers[999]).toBeUndefined()
  expect(fetchMock).not.toHaveBeenCalled()
  // Keep destination-provider requirements for the other existing SYN paths.
  await expect(sdk.bridgeV2(params(1, 999))).resolves.toEqual([])
})

it('shares identical concurrent fee quotes without caching settled fees', async () => {
  const { sdk, params, providersByChain } = setup([1])
  const quotes = await Promise.all([
    sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID)),
    sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID)),
  ])
  expect(quotes[0][0].expectedToAmount).toBe(quotes[1][0].expectedToAmount)
  expect(providersByChain[0].call).toHaveBeenCalledTimes(1)
  await sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID))
  expect(providersByChain[0].call).toHaveBeenCalledTimes(2)
})

it('keeps concurrent fee requests separate when the amount or recipient changes', async () => {
  const { sdk, params, providersByChain } = setup([1])
  const input = params(1, HYPERCORE_CHAIN_ID)
  const quotes = await Promise.all([
    sdk.bridgeV2(input),
    sdk.bridgeV2({ ...input, fromAmount: amount.add(unit).toString() }),
    sdk.bridgeV2({ ...input, toRecipient: sender }),
  ])
  expect(quotes.every((routes) => routes.length === 1)).toBe(true)
  expect(providersByChain[0].call).toHaveBeenCalledTimes(3)
})

it('retries fee reads after a shared request fails', async () => {
  const { sdk, params, providersByChain } = setup([1])
  ;(providersByChain[0].call as jest.Mock).mockRejectedValueOnce(
    new Error('Fee RPC unavailable')
  )
  const failed = await Promise.all([
    sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID)),
    sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID)),
  ])
  expect(failed).toEqual([[], []])
  expect(providersByChain[0].call).toHaveBeenCalledTimes(1)
  expect(await sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID))).toHaveLength(1)
  expect(providersByChain[0].call).toHaveBeenCalledTimes(2)
})

it.each([
  [unit, '100'],
  [unit.add(1), '100'],
  [unit.mul(2).sub(1), '100'],
  [unit.mul(2), '200'],
])(
  'preserves dust rounding at the OFT boundary for %s',
  async (input, core) => {
    const { sdk, params, quoteSendParams } = setup([1])
    const [quote] = await sdk.bridgeV2({
      ...params(1, HYPERCORE_CHAIN_ID),
      fromAmount: input.toString(),
    })
    expect(quote.expectedToAmount).toBe(core)
    expect(quoteSendParams[0][2]).toEqual(input)
    expect(quoteSendParams[0][3]).toEqual(input.div(unit).mul(unit))
  }
)

it('rejects uint64 overflow in Core units before quoting a fee', async () => {
  const { sdk, params, providersByChain } = setup([1])
  await expect(
    sdk.bridgeV2({
      ...params(1, HYPERCORE_CHAIN_ID),
      fromAmount: BigNumber.from(2).pow(64).mul(coreUnit).add(unit).toString(),
    })
  ).resolves.toEqual([])
  expect(providersByChain[0].call).not.toHaveBeenCalled()
})

it('rejects unrepresentable input and minimum amounts during local OFT calculation', async () => {
  const { sdk } = setup([1])
  const module = sdk.synModuleSet.modules[1]
  const sendParams = {
    toEid: 30367,
    toRecipient: recipient,
    fromSender: sender,
  }
  for (const invalid of [BigNumber.from(-1), constants.MaxUint256.add(1)]) {
    await expect(
      module.getDestinationQuote({ ...sendParams, amount: invalid })
    ).rejects.toThrow('uint256')
    await expect(
      module.getDestinationQuote({ ...sendParams, amount, minAmount: invalid })
    ).rejects.toThrow('uint256')
  }
  expect(() => quoteSynHyperCore(BigNumber.from(-1))).toThrow(
    'HyperCore amount range'
  )
  await expect(
    module.getDestinationQuote({
      ...sendParams,
      amount: BigNumber.from(2).pow(64).mul(unit),
    })
  ).rejects.toThrow('shared amount range')
  await expect(
    module.getDestinationQuote({ ...sendParams, amount, minAmount: amount })
  ).rejects.toThrow('minimum received amount')
})

it.each([
  [1, 999],
  [999, 1],
  [1, HYPERCORE_CHAIN_ID],
])(
  'quotes and builds executable SYN route %i -> %i through public bridgeV2',
  async (from, to) => {
    const { sdk, params, quoteSendParams, providersByChain } = setup()
    const [quote] = await sdk.bridgeV2(params(from, to))
    expect(sdk.providers[from].call).toHaveBeenCalledTimes(1)
    for (const provider of providersByChain) {
      if (provider !== sdk.providers[from]) {
        expect(provider.call).not.toHaveBeenCalled()
      }
    }
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
  const { sdk, params, providersByChain } = setup()
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
  expect(providersByChain[0].call).toHaveBeenCalledTimes(1)
  expect(providersByChain[1].call).not.toHaveBeenCalled()
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

it('exposes recipient activation while allowing a quote with the automatic HyperEVM fallback', async () => {
  const { sdk, params, quoteSendParams, providersByChain } = setup([1])
  const fetchMock = jest.spyOn(global, 'fetch').mockResolvedValue({
    ok: true,
    json: async () => ({ role: 'missing' }),
  } as Response)
  await expect(
    sdk.synModuleSet.isHyperCoreAccountActive(recipient)
  ).resolves.toBe(false)
  const [quote] = await sdk.bridgeV2(params(1, HYPERCORE_CHAIN_ID))
  expect(quote.tx).toBeDefined()
  expect(quoteSendParams[0][5]).toBe(
    utils.defaultAbiCoder.encode(['uint256', 'address'], [0, recipient])
  )
  expect(fetchMock).toHaveBeenCalledTimes(1)
  expect(providersByChain[0].call).toHaveBeenCalledTimes(1)
})

it('rejects unavailable recipient API checks without requiring a HyperEVM provider', async () => {
  const { sdk } = setup([1])
  jest.spyOn(global, 'fetch').mockRejectedValue(new Error('API unavailable'))
  await expect(
    sdk.synModuleSet.isHyperCoreAccountActive(recipient)
  ).rejects.toThrow('API unavailable')
})

it('omits zero-after-dust amounts without any RPC read', async () => {
  const { sdk, params, providersByChain } = setup()
  await expect(
    sdk.bridgeV2({ ...params(1, 999), fromAmount: unit.sub(1).toString() })
  ).resolves.toEqual([])
  await expect(
    sdk.bridgeV2({
      ...params(1, HYPERCORE_CHAIN_ID),
      fromAmount: unit.sub(1).toString(),
    })
  ).resolves.toEqual([])
  providersByChain.forEach((provider) =>
    expect(provider.call).not.toHaveBeenCalled()
  )
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
