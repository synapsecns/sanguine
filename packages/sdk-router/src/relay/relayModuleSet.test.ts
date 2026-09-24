import { BigNumber, providers } from 'ethers'

import {
  HYPERCORE_CHAIN_ID,
  HYPERCORE_USDC_ADDRESS,
  SupportedChainId,
  TOKEN_ZAP_V1_ADDRESS_MAP,
} from '../constants'
import { EngineID, decodeZapData, getEmptyRoute } from '../swap'
import { TokenMetadataFetcher } from '../utils'
import * as relayApi from './api'
import { RelayModuleSet } from './relayModuleSet'

const base = 8453
const sourceToken = '0x50c5725949A6F0c72E6C4a641F24049A917DB0Cb'
const sender = '0x1111111111111111111111111111111111111111'
const recipient = '0x2222222222222222222222222222222222222222'
const relayDeposit = '0x4cd00e387622c35bddb9b4c962c136462338bc31'
const relayCalldata = '0xe8017952'

const setup = () =>
  new RelayModuleSet(
    [base, 1, SupportedChainId.HYPEREVM].map((chainId) => ({
      chainId,
      provider: new providers.StaticJsonRpcProvider(
        'http://localhost:1',
        chainId
      ),
    }))
  )

afterEach(() => jest.restoreAllMocks())

it('passes the selected source token to Relay for HyperCore and other destinations', async () => {
  const relay = setup()
  const coreCandidates = await relay.getBridgeTokenCandidates({
    fromChainId: base,
    toChainId: HYPERCORE_CHAIN_ID,
    fromToken: sourceToken,
    toToken: HYPERCORE_USDC_ADDRESS,
  })
  expect(coreCandidates).toEqual([
    {
      originChainId: base,
      destChainId: HYPERCORE_CHAIN_ID,
      originToken: sourceToken,
      destToken: HYPERCORE_USDC_ADDRESS,
    },
  ])
  expect(
    await relay.getBridgeTokenCandidates({
      fromChainId: base,
      toChainId: HYPERCORE_CHAIN_ID,
      fromToken: sourceToken,
      toToken: sourceToken,
    })
  ).toEqual([])
  expect(
    await relay.getBridgeTokenCandidates({
      fromChainId: base,
      toChainId: 1,
      fromToken: sourceToken,
    })
  ).toEqual([expect.objectContaining({ originToken: sourceToken })])
  expect(
    await relay.getBridgeTokenCandidates({
      fromChainId: SupportedChainId.HYPEREVM,
      toChainId: HYPERCORE_CHAIN_ID,
      fromToken: '0xB8CE59FC3717ada4C02eaDF9682A9e934F625ebb',
      toToken: HYPERCORE_USDC_ADDRESS,
    })
  ).toEqual([
    expect.objectContaining({
      originChainId: SupportedChainId.HYPEREVM,
      destChainId: HYPERCORE_CHAIN_ID,
      originToken: '0xB8CE59FC3717ada4C02eaDF9682A9e934F625ebb',
    }),
  ])
})

it('quotes one HyperCore deposit step for the requested recipient', async () => {
  const relay = setup()
  const [bridgeToken] = await relay.getBridgeTokenCandidates({
    fromChainId: base,
    toChainId: HYPERCORE_CHAIN_ID,
    fromToken: sourceToken,
    toToken: HYPERCORE_USDC_ADDRESS,
  })
  const quoteSpy = jest.spyOn(relayApi, 'getQuote').mockResolvedValue({
    steps: [
      {
        id: 'deposit',
        kind: 'transaction',
        items: [
          {
            data: {
              from: TOKEN_ZAP_V1_ADDRESS_MAP[base],
              to: relayDeposit,
              data: relayCalldata,
              value: '0',
              chainId: base,
            },
          },
        ],
      },
    ],
    requestId: 'test',
    details: {
      currencyIn: {} as relayApi.QuoteResponse['details']['currencyIn'],
      currencyOut: {
        currency: {
          chainId: HYPERCORE_CHAIN_ID,
          address: HYPERCORE_USDC_ADDRESS,
          symbol: 'USDC',
          name: 'USDC (Perps)',
          decimals: 8,
        },
        amount: '9996999900',
        amountFormatted: '99.969999',
        amountUsd: '99.969999',
        minimumAmount: '9987002900',
      },
      timeEstimate: 60,
    },
  } as relayApi.QuoteResponse)
  const originSwapRoute = {
    ...getEmptyRoute(EngineID.NoOp),
    chainId: base,
    fromToken: sourceToken,
    toToken: bridgeToken.originToken,
    fromAmount: BigNumber.from('100000000000000000000'),
    expectedToAmount: BigNumber.from('100000000000000000000'),
  }
  const route = await relay.getBridgeRouteV2({
    bridgeToken,
    originSwapRoute,
    toToken: HYPERCORE_USDC_ADDRESS,
    fromSender: sender,
    toRecipient: recipient,
  })
  expect(quoteSpy).toHaveBeenCalledWith(
    expect.objectContaining({
      user: TOKEN_ZAP_V1_ADDRESS_MAP[base],
      originCurrency: bridgeToken.originToken,
      destinationChainId: HYPERCORE_CHAIN_ID,
      destinationCurrency: HYPERCORE_USDC_ADDRESS,
      recipient,
      refundTo: sender,
      useReceiver: true,
      explicitDeposit: true,
      usePermit: false,
    })
  )
  expect(route?.expectedToAmount.toString()).toBe('9996999900')
  expect(route?.minToAmount.toString()).toBe('9987002900')
  expect(decodeZapData(route?.zapData ?? '0x')).toEqual(
    expect.objectContaining({
      target: relayDeposit,
      payload: relayCalldata,
    })
  )
})

it('uses Relay status and 8-decimal native USDC without a HyperCore provider', async () => {
  const relay = setup()
  expect(relay.getModule(HYPERCORE_CHAIN_ID)).toBeDefined()
  const fetcher = new TokenMetadataFetcher({})
  expect(
    await fetcher.getTokenDecimals(HYPERCORE_CHAIN_ID, HYPERCORE_USDC_ADDRESS)
  ).toBe(8)
  jest.spyOn(relayApi, 'getRequests').mockResolvedValue({
    requests: [
      {
        id: 'test',
        status: relayApi.Status.Success,
        data: { inTxs: [{ hash: '0x1234' }] },
      },
    ],
  })
  expect(await relay.getBridgeTxStatus(HYPERCORE_CHAIN_ID, '0x1234')).toBe(true)
})
