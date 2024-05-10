import { createPublicClient, http } from 'viem'
import { sepolia, arbitrumSepolia } from 'viem/chains'

export const publicClient = {
  11155111: createPublicClient({
    chain: sepolia,
    transport: http(),
  }),
  421614: createPublicClient({
    chain: arbitrumSepolia,
    transport: http(),
  }),
} as const
