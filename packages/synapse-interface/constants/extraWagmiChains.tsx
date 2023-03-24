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
} as const satisfies Chain

export const boba = {
  id: 288,
  name: 'Boba Network',
  network: 'boba',
  nativeCurrency: {
    decimals: 18,
    name: 'Boba',
    symbol: 'BOBA',
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
    default: { http: ['https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc'] },
    public: { http: ['https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc'] },
  },
  blockExplorers: {
    etherscan: { name: 'DFKSubnetScan', url: 'https://subnets.avax.network/defi-kingdoms' },
    default: { name: 'DFKSubnetScan', url: 'https://subnets.avax.network/defi-kingdoms' },
  },
} as const satisfies Chain

// export const moonbeam = {
//   id: 1284,
//   name: 'Moonbeam',
//   network: 'moonbeam',
//   nativeCurrency: {
//     decimals: 18,
//     name: 'Moonbeam',
//     symbol: 'GLMR',
//   },
//   rpcUrls: {
//     default: { http: ['https://rpc.api.moonbeam.network'] },
//     public: { http: ['https://rpc.api.moonbeam.network'] },
//   },
//   blockExplorers: {
//     etherscan: { name: 'MoonBeamMoonScan', url: 'https://moonbeam.moonscan.io/' },
//     default: { name: 'MoonBeamMoonScan', url: 'https://moonbeam.moonscan.io/' },
//   },
//   contracts: {
//     multicall3: {
//       address: '0xcA11bde05977b3631167028862bE2a173976CA11',
//       blockCreated: 609002,
//     },
//   },
// } as const satisfies Chain

// export const moonriver = {
//   id: 1285,
//   name: 'Moonriver',
//   network: 'moonriver',
//   nativeCurrency: {
//     decimals: 18,
//     name: 'Moonriver',
//     symbol: 'MOVR',
//   },
//   rpcUrls: {
//     default: { http: ['https://rpc.api.moonriver.moonbeam.network'] },
//     public: { http: ['https://rpc.api.moonriver.moonbeam.network'] },
//   },
//   blockExplorers: {
//     etherscan: { name: 'MoonRiverMoonScan', url: 'https://moonriver.moonscan.io/' },
//     default: { name: 'MoonRiverMoonScan', url: 'https://moonriver.moonscan.io/' },
//   },
//   contracts: {
//     multicall3: {
//       address: '0xcA11bde05977b3631167028862bE2a173976CA11',
//       blockCreated: 1597904,
//     },
//   },
// } as const satisfies Chain

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
    etherscan: { name: 'DogeChainExplorer', url: 'https://explorer.dogechain.dog' },
    default: { name: 'DogeChainExplorer', url: 'https://explorer.dogechain.dog' },
  },

} as const satisfies Chain
