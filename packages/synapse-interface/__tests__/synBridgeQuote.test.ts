/** @jest-environment node */

import { fetchBridgeQuote } from '@/slices/bridgeQuote/thunks'
import { SYN } from '@/constants/tokens/bridgeable'
import { getErc20TokenAllowance } from '@/actions/getErc20TokenAllowance'

jest.mock('../actions/getErc20TokenAllowance', () => ({
  getErc20TokenAllowance: jest.fn().mockResolvedValue(10n ** 18n),
}))
jest.mock('../contexts/SegmentAnalyticsProvider', () => ({
  segmentAnalyticsEvent: jest.fn(),
}))

const sender = '0x1111111111111111111111111111111111111111'
const recipient = '0x2222222222222222222222222222222222222222'
const routerAddress = '0x3333333333333333333333333333333333333333'
const tx = { to: routerAddress, value: '100', data: '0x1234' }

const quote = (fromChainId = 1, toChainId = 1337) => ({
  id: 'syn-quote',
  fromChainId,
  toChainId,
  routerAddress,
  expectedToAmount: toChainId === 1337 ? '100000000' : '1000000000000000000',
  minToAmount: toChainId === 1337 ? '100000000' : '1000000000000000000',
  moduleNames: ['SYN'],
  estimatedTime: 840,
  gasDropAmount: '0',
  tx,
})

const runQuote = ({
  synapseSDK,
  fromChainId = 1,
  toChainId = 1337,
  address = sender,
  destinationAddress,
}: {
  synapseSDK: any
  fromChainId?: number
  toChainId?: number
  address?: string | null
  destinationAddress?: string
}) =>
  fetchBridgeQuote({
    synapseSDK,
    fromChainId,
    toChainId,
    fromToken: SYN,
    toToken: SYN,
    debouncedFromValue: '1',
    requestId: 1,
    currentTimestamp: 123,
    address: address as any,
    destinationAddress: destinationAddress as any,
    pausedModulesList: [],
  })(jest.fn(), jest.fn(), undefined)

describe('SYN bridge quote state', () => {
  beforeEach(() => jest.clearAllMocks())

  it.each([
    ['inactive custom recipient', recipient, false],
    ['active connected recipient', undefined, true],
  ])(
    'records the %s check against the actual recipient',
    async (_, custom, isActive) => {
      const isHyperCoreAccountActive = jest.fn().mockResolvedValue(isActive)
      const synapseSDK = {
        bridgeV2: jest.fn().mockResolvedValue([quote()]),
        synModuleSet: { isHyperCoreAccountActive },
      }

      const result = await runQuote({
        synapseSDK,
        destinationAddress: custom,
      })
      const expectedRecipient = custom ?? sender

      expect(isHyperCoreAccountActive).toHaveBeenCalledWith(expectedRecipient)
      expect(synapseSDK.bridgeV2).toHaveBeenCalledWith(
        expect.objectContaining({ toRecipient: expectedRecipient })
      )
      expect(fetchBridgeQuote.fulfilled.match(result)).toBe(true)
      expect(result.payload).toMatchObject({
        hyperCoreRecipient: { address: expectedRecipient, isActive },
      })
    }
  )

  it.each([
    ['a disconnected quote', 1, 1337, null],
    ['a non-HyperCore route', 1, 999, sender],
  ])(
    'skips the activation check for %s',
    async (_, fromChainId, toChainId, address) => {
      const isHyperCoreAccountActive = jest.fn()
      const synapseSDK = {
        bridgeV2: jest.fn().mockResolvedValue([quote(fromChainId, toChainId)]),
        synModuleSet: { isHyperCoreAccountActive },
      }

      const result = await runQuote({
        synapseSDK,
        fromChainId,
        toChainId,
        address,
      })

      expect(fetchBridgeQuote.fulfilled.match(result)).toBe(true)
      expect(isHyperCoreAccountActive).not.toHaveBeenCalled()
      expect(result.payload).toHaveProperty('hyperCoreRecipient', undefined)
    }
  )

  it('rejects the quote when the activation check fails', async () => {
    const isHyperCoreAccountActive = jest
      .fn()
      .mockRejectedValue(new Error('activation RPC unavailable'))
    const result = await runQuote({
      synapseSDK: {
        bridgeV2: jest.fn().mockResolvedValue([quote()]),
        synModuleSet: { isHyperCoreAccountActive },
      },
      destinationAddress: recipient,
    })

    expect(isHyperCoreAccountActive).toHaveBeenCalledWith(recipient)
    expect(fetchBridgeQuote.rejected.match(result)).toBe(true)
    if (fetchBridgeQuote.rejected.match(result)) {
      expect(result.error.message).toBe('activation RPC unavailable')
    }
  })

  it.each([
    [1, 999],
    [999, 1],
    [1, 1337],
  ])(
    'preserves the SDK destination and transaction for %s -> %s',
    async (fromChainId, toChainId) => {
      const synapseSDK = {
        bridgeV2: jest.fn().mockResolvedValue([quote(fromChainId, toChainId)]),
        synModuleSet: {
          isHyperCoreAccountActive: jest.fn().mockResolvedValue(true),
        },
      }
      const result = await runQuote({
        synapseSDK,
        fromChainId,
        toChainId,
        destinationAddress: recipient,
      })

      expect(synapseSDK.bridgeV2).toHaveBeenCalledWith(
        expect.objectContaining({
          fromChainId,
          toChainId,
          fromToken: SYN.addresses[fromChainId],
          toToken: SYN.addresses[toChainId],
          fromSender: sender,
          toRecipient: recipient,
        })
      )
      expect(fetchBridgeQuote.fulfilled.match(result)).toBe(true)
      expect(result.payload).toMatchObject({
        bridgeModuleName: 'SYN',
        originChainId: fromChainId,
        destChainId: toChainId,
        outputAmountString: '1.0',
        exchangeRate: 10n ** 18n,
        tx,
      })
      expect(getErc20TokenAllowance).toHaveBeenCalledWith(
        expect.objectContaining({
          chainId: fromChainId,
          tokenAddress: SYN.addresses[fromChainId],
          spender: routerAddress,
        })
      )
    }
  )
})
