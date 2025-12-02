// Mock @wormhole-foundation/sdk-connect to avoid ESM/CJS compatibility issues
// The Wormhole SDK has nested ESM dependencies (axios) that tsdx's Jest 25 cannot transform
jest.mock('@wormhole-foundation/sdk-connect', () => ({
  Network: {},
  Chain: {},
  ChainId: {},
  chain: {
    chainToChainId: jest.fn((chain) => 0),
  },
  nativeChainIds: {
    platformNativeChainIdToNetworkChain: jest.fn((platform, chainId) => [
      'Mainnet',
      'Ethereum',
    ]),
  },
  relayInstructionsLayout: {},
  serializeLayout: jest.fn(() => new Uint8Array([0])),
}))
