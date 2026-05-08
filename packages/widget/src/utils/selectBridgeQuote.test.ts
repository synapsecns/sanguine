import { selectBridgeQuote } from '@/utils/selectBridgeQuote'

const createQuote = (id: string, moduleNames: string[]) => ({
  id,
  moduleNames,
})

describe('selectBridgeQuote', () => {
  it('prefers an RFQ quote after paused-module filtering', () => {
    const quote = selectBridgeQuote({
      quotes: [
        createQuote('bridge', ['SynapseBridge']),
        createQuote('rfq', ['SynapseRFQ']),
      ],
      originChainId: 1,
      destinationChainId: 10,
      pausedModules: [],
    })

    expect(quote?.id).toBe('rfq')
  })

  it('respects chain-specific and ALL paused-module entries', () => {
    const quote = selectBridgeQuote({
      quotes: [
        createQuote('bridge', ['SynapseBridge']),
        createQuote('rfq', ['SynapseRFQ']),
      ],
      originChainId: 1,
      destinationChainId: 10,
      pausedModules: [
        { bridgeModuleName: 'SynapseBridge', chainId: 1 },
        { bridgeModuleName: 'ALL' },
      ],
    })

    expect(quote).toBeNull()
  })

  it('respects destination chain-specific paused-module entries', () => {
    const quote = selectBridgeQuote({
      quotes: [
        createQuote('bridge', ['SynapseBridge']),
        createQuote('rfq', ['SynapseRFQ']),
      ],
      originChainId: 1,
      destinationChainId: 10,
      pausedModules: [{ bridgeModuleName: 'SynapseRFQ', toChainId: 10 }],
    })

    expect(quote?.id).toBe('bridge')
  })

  it('does not apply destination chain-specific entries to other destinations', () => {
    const quote = selectBridgeQuote({
      quotes: [
        createQuote('bridge', ['SynapseBridge']),
        createQuote('rfq', ['SynapseRFQ']),
      ],
      originChainId: 1,
      destinationChainId: 42161,
      pausedModules: [{ bridgeModuleName: 'SynapseRFQ', toChainId: 10 }],
    })

    expect(quote?.id).toBe('rfq')
  })
})
