import { Chain } from 'wagmi'

export const klaytn = {
  id: 8217,
  name: 'Klaytn',
  network: 'klaytn',
  nativeCurrency: {
    decimals: 18,
    name: 'Klaytn',
    symbol: 'KLAY',
  },
  rpcUrls: {
    default: { http: ['https://public-node-api.klaytnapi.com/v1/cypress'] },
    public: { http: ['https://public-node-api.klaytnapi.com/v1/cypress'] },
  },
  blockExplorers: {
    etherscan: { name: 'KlaytnScope', url: 'https://scope.klaytn.com/' },
    default: { name: 'KlaytnScope', url: 'https://scope.klaytn.com/' },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 96002415,
    },
  },
} as const satisfies Chain

export const boba = {
  id: 288,
  name: 'Boba Network',
  network: 'boba',
  nativeCurrency: {
    decimals: 18,
    name: 'Ether',
    symbol: 'ETH',
  },
  rpcUrls: {
    default: { http: ['https://mainnet.boba.network'] },
    public: { http: ['https://mainnet.boba.network'] },
  },
  blockExplorers: {
    etherscan: { name: 'BOBAScan', url: 'https://bobascan.com' },
    default: { name: 'BOBAScan', url: 'https://bobascan.com' },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 446859,
    },
  },
} as const satisfies Chain

export const cronos = {
  id: 25,
  name: 'Cronos',
  network: 'cronos',
  nativeCurrency: {
    decimals: 18,
    name: 'Cronos',
    symbol: 'CRO',
  },
  rpcUrls: {
    default: { http: ['https://node.croswap.com/rpc'] },
    public: { http: ['https://node.croswap.com/rpc'] },
  },
  blockExplorers: {
    etherscan: { name: 'CronosScan', url: 'https://cronoscan.com/' },
    default: { name: 'CronosScan', url: 'https://cronoscan.com/' },
  },
  contracts: {
    multicall3: {
      address: '0xcA11bde05977b3631167028862bE2a173976CA11',
      blockCreated: 1963112,
    },
  },
} as const satisfies Chain

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
} as const satisfies Chain

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
} as const satisfies Chain

export const metis = {
  id: 1088,
  name: 'Metis',
  network: 'andromeda',
  nativeCurrency: {
    decimals: 18,
    name: 'Metis',
    symbol: 'METIS',
  },
  rpcUrls: {
    default: {
      http: ['https://andromeda.metis.io/?owner=1088'],
    },
    public: {
      http: ['https://andromeda.metis.io/?owner=1088'],
    },
  },
  blockExplorers: {
    default: {
      name: 'Andromeda Explorer',
      url: 'https://andromeda-explorer.metis.io',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 2338552,
    },
  },
} as const satisfies Chain

export const canto = {
  id: 7700,
  name: 'Canto',
  network: 'canto',
  nativeCurrency: {
    decimals: 18,
    name: 'Canto',
    symbol: 'CANTO',
  },
  rpcUrls: {
    default: {
      http: ['https://canto.slingshot.finance'],
    },
    public: {
      http: ['https://canto.slingshot.finance'],
    },
  },
  blockExplorers: {
    default: {
      name: 'Canto EVM Explorer (Blockscout)',
      url: 'https://tuber.build',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 4876481, //update this when cantoscan is working
    },
  },
} as const satisfies Chain

export const aurora = {
  id: 1313161554,
  name: 'Aurora',
  network: 'aurora',
  nativeCurrency: {
    decimals: 18,
    name: 'Ether',
    symbol: 'ETH',
  },
  rpcUrls: {
    infura: {
      http: ['https://aurora-mainnet.infura.io/v3'],
    },
    default: {
      http: ['https://mainnet.aurora.dev'],
    },
    public: {
      http: ['https://mainnet.aurora.dev'],
    },
  },
  blockExplorers: {
    etherscan: {
      name: 'Aurorascan',
      url: 'https://aurorascan.dev',
    },
    default: {
      name: 'Aurorascan',
      url: 'https://aurorascan.dev',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 62907816,
    },
  },
} as const satisfies Chain

export const base = {
  id: 8453,
  name: 'Base',
  network: 'base',
  nativeCurrency: {
    decimals: 18,
    name: 'Ether',
    symbol: 'ETH',
  },
  rpcUrls: {
    default: {
      http: ['https://developer-access-mainnet.base.org'],
    },
    public: {
      http: ['https://developer-access-mainnet.base.org'],
    },
  },
  blockExplorers: {
    // blockscout: {
    //   name: 'Basescout',
    //   url: 'https://base.blockscout.com',
    // },
    default: {
      name: 'Basescan',
      url: 'https://basescan.org',
    },
    etherscan: {
      name: 'Basescan',
      url: 'https://basescan.org',
    },
  },
  contracts: {
    multicall3: {
      address: '0xca11bde05977b3631167028862be2a173976ca11',
      blockCreated: 5022,
    },
  },
} as const satisfies Chain
