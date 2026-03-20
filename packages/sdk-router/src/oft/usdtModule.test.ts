import { providers } from 'ethers'
import { mock } from 'jest-mock-extended'

import { MEDIAN_TIME_BLOCK, SupportedChainId, USDT_OFT_ADDRESS_MAP } from '../constants'
import { UsdtModule } from './usdtModule'

describe('UsdtModule', () => {
  const mockProvider = mock<providers.Provider>()
  const module = new UsdtModule(
    SupportedChainId.ETH,
    mockProvider,
    USDT_OFT_ADDRESS_MAP[SupportedChainId.ETH]!
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

  it('queries LayerZero pathway confirmations with a 60 second execution buffer', async () => {
    fetchMock.mockResolvedValue({
      ok: true,
      json: async () => ({
        data: [{ config: { outboundConfig: { confirmations: 2 } } }],
      }),
    })
    const expectedEstimatedTime =
      Math.ceil((2 * MEDIAN_TIME_BLOCK[SupportedChainId.ETH] + 60) / 30) * 30
    await expect(
      module.getEstimatedTime(SupportedChainId.ARBITRUM)
    ).resolves.toBe(expectedEstimatedTime)
  })

  it('returns undefined estimated time when querying the same chain', async () => {
    await expect(module.getEstimatedTime(SupportedChainId.ETH)).resolves.toBeUndefined()
  })
})
