import { Chain, createPublicClient, http } from 'viem'
import {
  mainnet,
  arbitrum,
  optimism,
  base,
  scroll,
  linea,
  bsc,
  blast,
} from 'viem/chains'

const _worldchain: Chain = {
  id: 480,
  name: 'World Chain',
  nativeCurrency: {
    decimals: 18,
    name: 'Ether',
    symbol: 'ETH',
  },
  rpcUrls: {
    default: { http: ['https://worldchain-mainnet.g.alchemy.com/public'] },
    public: { http: ['https://worldchain-mainnet.g.alchemy.com/public'] },
  },
  blockExplorers: {
    default: {
      name: 'World Chain Explorer',
      url: 'https://worldchain-mainnet.explorer.alchemy.com',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 1517589,
    },
  },
}

const _unichain: Chain = {
  id: 130,
  name: 'Unichain',
  nativeCurrency: {
    decimals: 18,
    name: 'Ether',
    symbol: 'ETH',
  },
  rpcUrls: {
    default: { http: ['https://mainnet.unichain.org'] },
    public: { http: ['https://mainnet.unichain.org'] },
  },
  blockExplorers: {
    default: {
      name: 'Uniscan',
      url: 'https://uniscan.xyz',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 1,
    },
  },
}

const _berachain: Chain = {
  id: 80094,
  name: 'Berachain',
  nativeCurrency: {
    decimals: 18,
    name: 'Bera',
    symbol: 'BERA',
  },
  rpcUrls: {
    default: { http: ['https://rpc.berachain.com'] },
    public: { http: ['https://rpc.berachain.com'] },
  },
  blockExplorers: {
    default: {
      name: 'Berascan',
      url: 'https://berascan.com',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 1,
    },
  },
}

const _hyperEVM: Chain = {
  id: 999,
  name: 'HyperEVM',
  nativeCurrency: {
    decimals: 18,
    name: 'HYPE',
    symbol: 'HYPE',
  },
  rpcUrls: {
    default: { http: ['https://rpc.hyperliquid.xyz/evm'] },
    public: { http: ['https://rpc.hyperliquid.xyz/evm'] },
  },
  blockExplorers: {
    default: {
      name: 'Purrsec',
      url: 'https://purrsec.com',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 13051,
    },
  },
}

import { FastBridgeV2Abi } from './abis/FastBridgeV2'

interface NetworkEntry {
  name: string
  FastBridgeV2: {
    address: string
    abi: any
  }
  client: any
}

type NetworkConfig = {
  [chainId: number]: NetworkEntry
}

export const networkConfig: NetworkConfig = {
  1: {
    name: 'Ethereum',
    FastBridgeV2: {
      address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: mainnet,
      transport: http(),
    }),
  },
  42161: {
    name: 'Arbitrum',
    FastBridgeV2: {
      address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: arbitrum,
      transport: http(),
    }),
  },
  10: {
    name: 'Optimism',
    FastBridgeV2: {
      address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: optimism,
      transport: http(),
    }),
  },
  8453: {
    name: 'Base',
    FastBridgeV2: {
      address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: base,
      transport: http(),
    }),
  },
  534352: {
    name: 'Scroll',
    FastBridgeV2: {
      address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: scroll,
      transport: http(),
    }),
  },
  59144: {
    name: 'Linea',
    FastBridgeV2: {
      address: '0x34F52752975222d5994C206cE08C1d5B329f24dD',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: linea,
      transport: http(),
    }),
  },
  56: {
    name: 'BNB Chain',
    FastBridgeV2: {
      address: '0x34F52752975222d5994C206cE08C1d5B329f24dD',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: bsc,
      transport: http(),
    }),
  },
  81457: {
    name: 'Blast',
    FastBridgeV2: {
      address: '0x34F52752975222d5994C206cE08C1d5B329f24dD',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: blast,
      transport: http(),
    }),
  },
  480: {
    name: 'Worldchain',
    FastBridgeV2: {
      address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: _worldchain,
      transport: http(),
    }),
  },
  130: {
    name: 'Unichain',
    FastBridgeV2: {
      address: '0x63c3211257CcE0c12c7c7A6DBb75960fEaBF45Be',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: _unichain,
      transport: http(),
    }),
  },
  80094: {
    name: 'Berachain',
    FastBridgeV2: {
      address: '0x63c3211257CcE0c12c7c7A6DBb75960fEaBF45Be',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: _berachain,
      transport: http(),
    }),
  },
  999: {
    name: 'HyperEVM',
    FastBridgeV2: {
      address: '0x63c3211257CcE0c12c7c7A6DBb75960fEaBF45Be',
      abi: FastBridgeV2Abi,
    },
    client: createPublicClient({
      chain: _hyperEVM,
      transport: http(),
    }),
  },
} as const
