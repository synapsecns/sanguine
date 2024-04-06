import { createConfig } from '@ponder/core'
import { http } from 'viem'

import { InterchainClientV1Abi } from './abis/InterchainClientV1Abi'

const arbSepoliaTransport = http(process.env.ARB_SEPOLIA_RPC)
const ethSepoliaTransport = http(process.env.ETH_SEPOLIA_RPC)

export default createConfig({
  networks: {
    ethSepolia: {
      chainId: 11155111,
      transport: ethSepoliaTransport,
    },
    arbSepolia: {
      chainId: 421614,
      transport: arbSepoliaTransport,
    },
  },
  contracts: {
    InterchainClientV1: {
      network: {
        ethSepolia: {
          address: '0x28f0B8E80e5afd62A20C23BD5098237006634318',
          startBlock: 5628137,
        },
        arbSepolia: {
          address: '0x188cA7f9615042654e483Ed840582208009A9ADF',
          startBlock: 30278491,
        },
      },
      abi: InterchainClientV1Abi,
    },
  },
})
