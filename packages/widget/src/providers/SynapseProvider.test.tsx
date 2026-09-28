import { renderHook } from '@testing-library/react'
import { SynapseSDK } from '@synapsecns/sdk-router'
import { FallbackProvider } from '@ethersproject/providers'
import { SynapseProvider, useSynapseContext } from './SynapseProvider'

jest.mock('@synapsecns/sdk-router', () => ({ SynapseSDK: jest.fn() }))
jest.mock('@ethersproject/providers', () => ({
  StaticJsonRpcProvider: jest
    .fn()
    .mockImplementation((url, id) => ({ url, network: { chainId: id } })),
  FallbackProvider: jest.fn().mockImplementation((configs) => ({
    configs,
    network: configs[0].provider.network,
  })),
}))

const initial = [
  {
    id: 1,
    rpcUrls: { primary: 'https://primary', fallback: 'https://fallback' },
  },
] as any

describe('widget SDK provider lifetime', () => {
  beforeEach(() => {
    jest.clearAllMocks()
    ;(SynapseSDK as jest.Mock).mockImplementation(() => ({}))
  })

  it('retains equivalent providers and SDK, but rebuilds when custom RPC changes', () => {
    let chains = initial
    let customRpcs = { 1: 'https://custom' }
    const wrapper = ({ children }) => (
      <SynapseProvider chains={chains} customRpcs={customRpcs}>
        {children}
      </SynapseProvider>
    )
    const { result, rerender } = renderHook(useSynapseContext, { wrapper })
    const context = result.current
    chains = initial.map((chain) => ({
      ...chain,
      rpcUrls: { ...chain.rpcUrls },
    }))
    customRpcs = { ...customRpcs }
    rerender()
    expect(result.current).toBe(context)
    expect(SynapseSDK).toHaveBeenCalledTimes(1)
    expect(FallbackProvider).toHaveBeenCalledTimes(1)
    customRpcs = { 1: 'https://replacement' }
    rerender()
    expect(result.current).not.toBe(context)
    expect(SynapseSDK).toHaveBeenCalledTimes(2)
    expect(FallbackProvider).toHaveBeenCalledTimes(2)
    expect(
      (FallbackProvider as unknown as jest.Mock).mock.calls[1][0][0].provider
        .url
    ).toBe('https://replacement')
  })
})
