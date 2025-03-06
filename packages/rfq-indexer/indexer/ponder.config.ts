import { createConfig } from '@ponder/core'
import { http } from 'viem'

import { FastBridgeV2Abi } from '@/abis/FastBridgeV2'
import { AddressConfig } from '@/indexer/src/types'

// Mainnets
const ethereumChainId = 1
const optimismChainId = 10
const arbitrumChainId = 42161
const baseChainId = 8453
const blastChainId = 81457
const scrollChainId = 534352
const lineaChainId = 59144
const bnbChainId = 56
const worldchainChainId = 480
const unichainChainId = 130
const berachainChainId = 80094
const hyperevmChainId = 999

const configByChainId = {
  [1]: {
    transport: http(process.env.ETH_MAINNET_RPC),
    chainName: 'ethereum',
    FastBridgeV2Address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
    //   FastBridgeV2StartBlock: 19420718, first block
    FastBridgeV2StartBlock: 20426589, // new block
  },
  [10]: {
    transport: http(process.env.OPTIMISM_MAINNET_RPC),
    chainName: 'optimism',
    FastBridgeV2Address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
    //   FastBridgeV2StartBlock: 117334308, first block
    FastBridgeV2StartBlock: 123416470, // new block
  },
  [42161]: {
    transport: http(process.env.ARBITRUM_MAINNET_RPC),
    chainName: 'arbitrum',
    FastBridgeV2Address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
    //   FastBridgeV2StartBlock: 189700328, first block
    FastBridgeV2StartBlock: 237979967, // new block
  },
  [8453]: {
    transport: http(process.env.BASE_MAINNET_RPC),
    chainName: 'base',
    FastBridgeV2Address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
    //   FastBridgeV2StartBlock: 12478374, first block
    FastBridgeV2StartBlock: 17821292, // new block
  },
  [81457]: {
    transport: http(process.env.BLAST_MAINNET_RPC),
    chainName: 'blast',
    FastBridgeV2Address: '0x34F52752975222d5994C206cE08C1d5B329f24dD',
    //   FastBridgeV2StartBlock: 6378234, first block
    FastBridgeV2StartBlock: 6811045, // new block
  },
  [534352]: {
    transport: http(process.env.SCROLL_MAINNET_RPC),
    chainName: 'scroll',
    FastBridgeV2Address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
    //   FastBridgeV2StartBlock: 5357000, first block
    FastBridgeV2StartBlock: 7941653, // new block
  },
  [59144]: {
    transport: http(process.env.LINEA_MAINNET_RPC),
    chainName: 'linea',
    FastBridgeV2Address: '0x34F52752975222d5994C206cE08C1d5B329f24dD',
    FastBridgeV2StartBlock: 7124666, // first block and new block
  },
  [56]: {
    transport: http(process.env.BNB_MAINNET_RPC),
    chainName: 'bnb',
    FastBridgeV2Address: '0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E',
    FastBridgeV2StartBlock: 40497843, // first block and new block
  },
  [480]: {
    transport: http(process.env.WORLDCHAIN_MAINNET_RPC),
    chainName: 'worldchain',
    FastBridgeV2Address: '0x05C62156C7C47E76223A560210EA648De5e6B53B',
    FastBridgeV2StartBlock: 4598830, // first block and new block
  },
  [130]: {
    transport: http(process.env.UNICHAIN_MAINNET_RPC),
    chainName: 'unichain',
    FastBridgeV2Address: '0x63c3211257CcE0c12c7c7A6DBb75960fEaBF45Be',
    FastBridgeV2StartBlock: 8435535, // first block and new block
  },
  [80094]: {
    transport: http(process.env.BERACHAIN_MAINNET_RPC),
    chainName: 'berachain',
    FastBridgeV2Address: '0x63c3211257CcE0c12c7c7A6DBb75960fEaBF45Be',
    FastBridgeV2StartBlock: 953594, // first block and new block
  },
  [999]: {
    transport: http(process.env.HYPEREVM_MAINNET_RPC),
    chainName: 'hyperevm',
    FastBridgeV2Address: '0x63c3211257CcE0c12c7c7A6DBb75960fEaBF45Be',
    FastBridgeV2StartBlock: 256939, // first block and new block
  },
  disableCache: true,
}

export const networkDetails = {
  [ethereumChainId]: {
    name: configByChainId[ethereumChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[ethereumChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[ethereumChainId].FastBridgeV2StartBlock,
    },
  },
  [optimismChainId]: {
    name: configByChainId[optimismChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[optimismChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[optimismChainId].FastBridgeV2StartBlock,
    },
  },
  [arbitrumChainId]: {
    name: configByChainId[arbitrumChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[arbitrumChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[arbitrumChainId].FastBridgeV2StartBlock,
    },
  },
  [baseChainId]: {
    name: configByChainId[baseChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[baseChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[baseChainId].FastBridgeV2StartBlock,
    },
  },
  [blastChainId]: {
    name: configByChainId[blastChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[blastChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[blastChainId].FastBridgeV2StartBlock,
    },
  },
  [scrollChainId]: {
    name: configByChainId[scrollChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[scrollChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[scrollChainId].FastBridgeV2StartBlock,
    },
  },
  [lineaChainId]: {
    name: configByChainId[lineaChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[lineaChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[lineaChainId].FastBridgeV2StartBlock,
    },
  },
  [bnbChainId]: {
    name: configByChainId[bnbChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[bnbChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[bnbChainId].FastBridgeV2StartBlock,
    },
  },
  [worldchainChainId]: {
    name: configByChainId[worldchainChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[worldchainChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[worldchainChainId].FastBridgeV2StartBlock,
    },
  },
  [unichainChainId]: {
    name: configByChainId[unichainChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[unichainChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[unichainChainId].FastBridgeV2StartBlock,
    },
  },
  [berachainChainId]: {
    name: configByChainId[berachainChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[berachainChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[berachainChainId].FastBridgeV2StartBlock,
    },
  },
  [hyperevmChainId]: {
    name: configByChainId[hyperevmChainId].chainName,
    FastBridgeV2: {
      address: configByChainId[hyperevmChainId].FastBridgeV2Address,
      abi: FastBridgeV2Abi,
      startBlock: configByChainId[hyperevmChainId].FastBridgeV2StartBlock,
      maxBlockRange: 50,
    },
  },
} as Record<number, AddressConfig>

const config = createConfig({
  networks: {
    [configByChainId[ethereumChainId].chainName]: {
      chainId: ethereumChainId,
      transport: configByChainId[ethereumChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[optimismChainId].chainName]: {
      chainId: optimismChainId,
      transport: configByChainId[optimismChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[arbitrumChainId].chainName]: {
      chainId: arbitrumChainId,
      transport: configByChainId[arbitrumChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[baseChainId].chainName]: {
      chainId: baseChainId,
      transport: configByChainId[baseChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[blastChainId].chainName]: {
      chainId: blastChainId,
      transport: configByChainId[blastChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[scrollChainId].chainName]: {
      chainId: scrollChainId,
      transport: configByChainId[scrollChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[lineaChainId].chainName]: {
      chainId: lineaChainId,
      transport: configByChainId[lineaChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[bnbChainId].chainName]: {
      chainId: bnbChainId,
      transport: configByChainId[bnbChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[worldchainChainId].chainName]: {
      chainId: worldchainChainId,
      transport: configByChainId[worldchainChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[unichainChainId].chainName]: {
      chainId: unichainChainId,
      transport: configByChainId[unichainChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[berachainChainId].chainName]: {
      chainId: berachainChainId,
      transport: configByChainId[berachainChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
    [configByChainId[hyperevmChainId].chainName]: {
      chainId: hyperevmChainId,
      transport: configByChainId[hyperevmChainId].transport,
      //   disableCache: configByChainId.disableCache,
    },
  },
  contracts: {
    FastBridgeV2: {
      network: {
        [configByChainId[ethereumChainId].chainName]: {
          address: networkDetails[ethereumChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[ethereumChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[optimismChainId].chainName]: {
          address: networkDetails[optimismChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[optimismChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[arbitrumChainId].chainName]: {
          address: networkDetails[arbitrumChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[arbitrumChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[baseChainId].chainName]: {
          address: networkDetails[baseChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[baseChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[blastChainId].chainName]: {
          address: networkDetails[blastChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[blastChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[scrollChainId].chainName]: {
          address: networkDetails[scrollChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[scrollChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[lineaChainId].chainName]: {
          address: networkDetails[lineaChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[lineaChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[bnbChainId].chainName]: {
          address: networkDetails[bnbChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[bnbChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[worldchainChainId].chainName]: {
          address: networkDetails[worldchainChainId]?.FastBridgeV2.address,
          startBlock:
            networkDetails[worldchainChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[unichainChainId].chainName]: {
          address: networkDetails[unichainChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[unichainChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[berachainChainId].chainName]: {
          address: networkDetails[berachainChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[berachainChainId]?.FastBridgeV2.startBlock,
        },
        [configByChainId[hyperevmChainId].chainName]: {
          address: networkDetails[hyperevmChainId]?.FastBridgeV2.address,
          startBlock: networkDetails[hyperevmChainId]?.FastBridgeV2.startBlock,
        },
      },
      abi: FastBridgeV2Abi,
    },
  },
})

export default config
