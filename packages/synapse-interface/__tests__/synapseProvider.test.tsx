import { renderHook } from '@testing-library/react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { FallbackProvider } from '@ethersproject/providers'

import {
  SynapseProvider,
  useSynapseContext,
} from '@/utils/providers/SynapseProvider'

jest.mock('@synapsecns/sdk-router', () => ({ SynapseSDK: jest.fn() }))
jest.mock('../utils/providers/TransportAwareJsonRpcProvider', () => ({
  TransportAwareJsonRpcProvider: jest
    .fn()
    .mockImplementation((url, id) => ({ url, network: { chainId: id } })),
}))
jest.mock('@ethersproject/providers', () => ({
  FallbackProvider: jest.fn().mockImplementation((configs) => ({
    configs,
    network: configs[0].provider.network,
  })),
}))

const initial = [
  { id: 1, configRpc: 'https://primary', fallbackRpc: 'https://fallback' },
]

describe('interface SDK provider lifetime', () => {
  beforeEach(() => {
    jest.clearAllMocks()
    ;(SynapseSDK as jest.Mock).mockImplementation(() => ({}))
  })

  it('preserves SDK and providers across children and equivalent chain prop changes', () => {
    let chains = initial
    const wrapper = ({ children }) => (
      <SynapseProvider chains={chains}>{children}</SynapseProvider>
    )
    const { result, rerender } = renderHook(useSynapseContext, { wrapper })
    const context = result.current
    chains = initial.map((chain) => ({ ...chain }))
    rerender()
    expect(result.current).toBe(context)
    expect(SynapseSDK).toHaveBeenCalledTimes(1)
    expect(FallbackProvider).toHaveBeenCalledTimes(1)
    chains = [{ ...initial[0], fallbackRpc: 'https://replacement' }]
    rerender()
    expect(result.current).not.toBe(context)
    expect(SynapseSDK).toHaveBeenCalledTimes(2)
    expect(FallbackProvider).toHaveBeenCalledTimes(2)
  })
})
