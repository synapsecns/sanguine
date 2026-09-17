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

describe('SYN bridge quote state', () => {
  it('preserves the inactive HyperCore recipient error for the existing toast', async () => {
    const result = await fetchBridgeQuote({
      synapseSDK: {
        bridgeV2: jest
          .fn()
          .mockRejectedValue(
            Object.assign(
              new Error(
                'Activate the recipient account on HyperCore before bridging SYN.'
              ),
              { code: 'HYPERCORE_ACCOUNT_INACTIVE' }
            )
          ),
      },
      fromChainId: 1,
      toChainId: 1337,
      fromToken: SYN,
      toToken: SYN,
      debouncedFromValue: '1',
      requestId: 1,
      currentTimestamp: 123,
      address: sender,
      destinationAddress: recipient,
      pausedModulesList: [],
    })(jest.fn(), jest.fn(), undefined)
    expect(fetchBridgeQuote.rejected.match(result)).toBe(true)
    if (fetchBridgeQuote.rejected.match(result)) {
      expect(result.error.code).toBe('HYPERCORE_ACCOUNT_INACTIVE')
      expect(result.error.message).toContain('Activate the recipient account')
    }
  })

  it.each([
    [1, 999],
    [999, 1],
    [1, 1337],
  ])(
    'preserves the SDK destination and transaction for %s -> %s',
    async (fromChainId, toChainId) => {
      const tx = { to: routerAddress, value: '100', data: '0x1234' }
      const synapseSDK = {
        bridgeV2: jest.fn().mockResolvedValue([
          {
            id: 'syn-quote',
            fromChainId,
            toChainId,
            routerAddress,
            expectedToAmount:
              toChainId === 1337 ? '100000000' : '1000000000000000000',
            minToAmount:
              toChainId === 1337 ? '100000000' : '1000000000000000000',
            moduleNames: ['SYN'],
            estimatedTime: 840,
            gasDropAmount: '0',
            tx,
          },
        ]),
      }
      const result = await fetchBridgeQuote({
        synapseSDK,
        fromChainId,
        toChainId,
        fromToken: SYN,
        toToken: SYN,
        debouncedFromValue: '1',
        requestId: 1,
        currentTimestamp: 123,
        address: sender,
        destinationAddress: recipient,
        pausedModulesList: [],
      })(jest.fn(), jest.fn(), undefined)

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
