import { createPublicClient, http } from 'viem'
import {
  mainnet,
  arbitrum,
  optimism,
  base,
  scroll,
  linea,
  bsc,
  blast,
  worldchain,
} from 'viem/chains'

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
      chain: worldchain,
      transport: http(),
    }),
  },
} as const
