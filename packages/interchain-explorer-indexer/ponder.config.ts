import { createConfig } from '@ponder/core'
import { http } from 'viem'

import { InterchainClientV1Abi } from '@/abis/InterchainClientV1Abi'
import { InterchainDBAbi } from '@/abis/InterchainDBAbi'

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
          address: '0x6bAb7426099ba52ac37F309903169C4c0A5f7534',
          startBlock: 5628137,
        },
        arbSepolia: {
          address: '0x15ACDFd1F2027aE084B4d92da20D22cc945d07Ec',
          startBlock: 30278491,
        },
      },
      abi: InterchainClientV1Abi,
    },
    InterchainDB: {
      network: {
        ethSepolia: {
          address: '0x8d50e833331A0D01d6F286881ce2C3A5DAD12e26',
          startBlock: 5628137,
        },
        arbSepolia: {
          address: '0x943257aE5037f5997ab302c4E158EFe48BBCE89d',
          startBlock: 30278491,
        },
      },
      abi: InterchainDBAbi,
    },
  },
})

/* TODO: Refactor */
export const networkDetails: any = {
  11155111: {
    name: 'ethSepolia',
    InterchainClientV1: {
      address: '0x6bAb7426099ba52ac37F309903169C4c0A5f7534',
      startBlock: 5628137,
      abi: InterchainClientV1Abi,
    },
    InterchainDB: {
      address: '0x8d50e833331A0D01d6F286881ce2C3A5DAD12e26',
      startBlock: 5628137,
      abi: InterchainDBAbi,
    },
  },
  421614: {
    name: 'arbSepolia',
    InterchainClientV1: {
      address: '0x15ACDFd1F2027aE084B4d92da20D22cc945d07Ec',
      startBlock: 30278491,
      abi: InterchainClientV1Abi,
    },
    InterchainDB: {
      address: '0x943257aE5037f5997ab302c4E158EFe48BBCE89d',
      startBlock: 30278491,
      abi: InterchainDBAbi,
    },
  },
}
