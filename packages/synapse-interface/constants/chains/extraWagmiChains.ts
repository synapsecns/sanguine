export const dfk = {
  id: 53935,
  name: 'DFK Chain',
  network: 'dfk',
  nativeCurrency: {
    decimals: 18,
    name: 'Jewel',
    symbol: 'JEWEL',
  },
  rpcUrls: {
    default: {
      http: ['https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc'],
    },
    public: {
      http: ['https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc'],
    },
  },
  blockExplorers: {
    etherscan: {
      name: 'DFKSubnetScan',
      url: 'https://subnets.avax.network/defi-kingdoms',
    },
    default: {
      name: 'DFKSubnetScan',
      url: 'https://subnets.avax.network/defi-kingdoms',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 14790551,
    },
  },
}

export const dogechain = {
  id: 2000,
  name: 'Dogechain',
  network: 'dogechain',
  nativeCurrency: {
    decimals: 18,
    name: 'Dogechain',
    symbol: 'DC',
  },
  rpcUrls: {
    default: { http: ['https://rpc.dogechain.dog'] },
    public: { http: ['https://rpc.dogechain.dog'] },
  },
  blockExplorers: {
    etherscan: {
      name: 'DogeChainExplorer',
      url: 'https://explorer.dogechain.dog',
    },
    default: {
      name: 'DogeChainExplorer',
      url: 'https://explorer.dogechain.dog',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 13882887,
    },
  },
}

export const worldchain = {
  id: 480,
  name: 'World Chain',
  network: 'worldchain',
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

export const unichain = {
  id: 130,
  name: 'Unichain',
  network: 'unichain',
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

export const berachain = {
  id: 80094,
  name: 'Berachain',
  network: 'berachain',
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

export const hyperEVM = {
  id: 999,
  name: 'HyperEVM',
  network: 'hyperevm',
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
