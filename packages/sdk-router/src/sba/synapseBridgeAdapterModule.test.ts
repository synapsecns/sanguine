import { providers } from 'ethers'
import { mock } from 'jest-mock-extended'

import { MEDIAN_TIME_BLOCK, SupportedChainId } from '../constants'
import { SBA_CHAIN_METADATA } from './metadata'
import { SynapseBridgeAdapterModule } from './synapseBridgeAdapterModule'

describe('SynapseBridgeAdapterModule', () => {
  const mockProvider = mock<providers.Provider>()
  const module = new SynapseBridgeAdapterModule(
    SupportedChainId.ETH,
    mockProvider,
    SBA_CHAIN_METADATA[SupportedChainId.ETH]!.adapterAddress
  )
  let fetchMock: jest.Mock

  beforeEach(() => {
    fetchMock = jest.fn()
    global.fetch = fetchMock as any
  })

  afterEach(() => {
    jest.restoreAllMocks()
    jest.clearAllMocks()
  })

  it('bridge() throws because bridge V1 is not supported', async () => {
    await expect(module.bridge()).rejects.toThrow('bridge V1 not supported')
  })

  it('passes tx hash through as the synapse tx id', async () => {
    const txHash = '0x1234'
    await expect(module.getSynapseTxId(txHash)).resolves.toEqual(txHash)
  })

  it.each(['CONFIRMING', 'DELIVERED'])(
    'treats LayerZero status %s as complete',
    async (status) => {
      fetchMock.mockResolvedValue({
        ok: true,
        json: async () => ({
          data: [{ status: { name: status } }],
        }),
      })
      await expect(module.getBridgeTxStatus('0x1234')).resolves.toBe(true)
    }
  )

  it('returns false for incomplete LayerZero status', async () => {
    fetchMock.mockResolvedValue({
      ok: true,
      json: async () => ({
        data: [{ status: { name: 'INFLIGHT' } }],
      }),
    })
    await expect(module.getBridgeTxStatus('0x1234')).resolves.toBe(false)
  })

  it('returns false when the LayerZero API response is not ok', async () => {
    fetchMock.mockResolvedValue({
      ok: false,
      text: async () => 'boom',
    })
    await expect(module.getBridgeTxStatus('0x1234')).resolves.toBe(false)
  })

  it('returns false when fetch rejects during status polling', async () => {
    fetchMock.mockRejectedValue(new Error('network error'))
    await expect(module.getBridgeTxStatus('0x1234')).resolves.toBe(false)
  })

  it('queries LayerZero pathway confirmations for estimated time', async () => {
    fetchMock.mockResolvedValue({
      ok: true,
      json: async () => ({
        data: [{ config: { outboundConfig: { confirmations: 2 } } }],
      }),
    })
    const expectedEstimatedTime =
      2 * MEDIAN_TIME_BLOCK[SupportedChainId.ETH] + 60
    await expect(
      module.getEstimatedTime(SupportedChainId.OPTIMISM)
    ).resolves.toBeCloseTo(expectedEstimatedTime)
  })

  it('returns undefined estimated time when querying the same chain', async () => {
    await expect(
      module.getEstimatedTime(SupportedChainId.ETH)
    ).resolves.toBeUndefined()
  })

  it('discovers the encoded amount position for bridgeERC20', () => {
    expect(module.getAmountPosition()).toEqual(100)
    expect(module.getAmountPosition()).toEqual(100)
  })
})
