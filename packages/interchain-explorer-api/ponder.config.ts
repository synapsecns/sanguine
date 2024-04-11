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
          address: '0x5a81Dfa5885058ED838fa750060C804B49F69991',
          startBlock: 5628137,
        },
        arbSepolia: {
          address: '0xBdC170214726a994D9A837b15C0039bdeE090878',
          startBlock: 30278491,
        },
      },
      abi: InterchainClientV1Abi,
    },
  },
})
