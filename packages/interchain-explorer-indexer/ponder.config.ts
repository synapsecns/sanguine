import { createConfig } from '@ponder/core'
import { http } from 'viem'

import { InterchainClientV1Abi } from '@/abis/InterchainClientV1Abi'
import { InterchainDBAbi } from './abis/InterchainDBAbi'
import { SynapseModuleAbi } from './abis/SynapseModuleAbi'

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
          address: '0xfcb988E117dbAa8c8b48047afd7c22d4a8321bCA',
          startBlock: 5628137,
        },
        arbSepolia: {
          address: '0x4AdfEb01C090e14BFA84411b74D0d03dDE12e39b',
          startBlock: 30278491,
        },
      },
      abi: InterchainClientV1Abi,
    },
    InterchainDB: {
      network: {
        ethSepolia: {
          address: '0x1Ce3a9d87A26Ae56a43a5BB1D5e9A8D14550D0a4',
          startBlock: 5628137,
        },
        arbSepolia: {
          address: '0x4361F461c5Df0DCf109BA2CF0E46dfA26e73f54f',
          startBlock: 30278491,
        },
      },
      abi: InterchainDBAbi,
    },
    SynapseModule: {
      network: {
        ethSepolia: {
          address: '0x95f2e2fAFE38f2aAdC9F9cBef98785809cc4bb6B',
          startBlock: 5628137,
        },
        arbSepolia: {
          address: '0xC13e2b478f6531Ef096FF05733Ed65E3bc7fC5AF',
          startBlock: 30278491,
        },
      },
      abi: SynapseModuleAbi,
    },
  },
})

/* TODO: Refactor */
export const networkDetails: any = {
  11155111: {
    name: 'ethSepolia',
    InterchainClientV1: {
      address: '0xfcb988E117dbAa8c8b48047afd7c22d4a8321bCA',
      startBlock: 5628137,
      abi: InterchainClientV1Abi,
    },
    InterchainDB: {
      address: '0x1Ce3a9d87A26Ae56a43a5BB1D5e9A8D14550D0a4',
      startBlock: 5628137,
      abi: InterchainDBAbi,
    },
    SynapseModule: {
      address: '0x95f2e2fAFE38f2aAdC9F9cBef98785809cc4bb6B',
      startBlock: 5628137,
      abi: SynapseModuleAbi,
    },
  },
  421614: {
    name: 'arbSepolia',
    InterchainClientV1: {
      address: '0x4AdfEb01C090e14BFA84411b74D0d03dDE12e39b',
      startBlock: 30278491,
      abi: InterchainClientV1Abi,
    },
    InterchainDB: {
      address: '0x4361F461c5Df0DCf109BA2CF0E46dfA26e73f54f',
      startBlock: 30278491,
      abi: InterchainDBAbi,
    },
    SynapseModule: {
      address: '0xC13e2b478f6531Ef096FF05733Ed65E3bc7fC5AF',
      startBlock: 30278491,
      abi: SynapseModuleAbi,
    },
  },
}
