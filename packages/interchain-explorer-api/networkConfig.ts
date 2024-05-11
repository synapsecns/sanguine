import { InterchainClientV1Abi } from '@/abis/InterchainClientV1Abi'
import { createPublicClient, http } from 'viem'
import { sepolia, arbitrumSepolia } from 'viem/chains'

interface NetworkEntry {
  name: string
  InterchainClientV1: {
    address: string
    abi: any
  }
  client: any
}

type NetworkConfig = {
  [chainId: number]: NetworkEntry
}

export const networkConfig: NetworkConfig = {
  11155111: {
    name: 'ethSepolia',
    InterchainClientV1: {
      address: '0x6bAb7426099ba52ac37F309903169C4c0A5f7534',
      abi: InterchainClientV1Abi,
    },
    client: createPublicClient({
      chain: sepolia,
      transport: http(),
    }),
  },
  421614: {
    name: 'arbSepolia',
    InterchainClientV1: {
      address: '0x15ACDFd1F2027aE084B4d92da20D22cc945d07Ec',
      abi: InterchainClientV1Abi,
    },
    client: createPublicClient({
      chain: arbitrumSepolia,
      transport: http(),
    }),
  },
} as const
