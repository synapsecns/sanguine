import binanceLogo from '@assets/icons/binance.svg'
import ethLogo from '@assets/icons/eth.svg'
import polygonLogo from '@assets/icons/polygon.svg'
import fantomLogo from '@assets/icons/fantom.svg'
import arbitrumLogo from '@assets/icons/arbitrum.svg'
import avalancheLogo from '@assets/icons/avalanche.svg'
import auroraLogo from '@assets/icons/aurora.svg'
import harmonyLogo from '@assets/icons/harmonyone.svg'
import optimismLogo from '@assets/icons/optimism.svg'
import bobaLogo from '@assets/icons/boba.svg'
import cronosLogo from '@assets/icons/cronos.svg'
import metisLogo from '@assets/icons/metis.svg'
import cantoLogo from '@assets/icons/canto.svg'
import dogeLogo from '@assets/icons/doge.svg'
import ethImg from '@assets/networks/eth.jpg'
import bscImg from '@assets/networks/bsc.jpg'
import polygonImg from '@assets/networks/polygon.jpg'
import fantomImg from '@assets/networks/fantom.jpg'
import arbitrumImg from '@assets/networks/arbitrum.jpg'
import avalancheImg from '@assets/networks/avalanche.jpg'
import dfkImg from '@assets/networks/dfk.png'
import auroraImg from '@assets/networks/aurora.png'
import harmonyImg from '@assets/networks/harmonyone.jpg'
import optimismImg from '@assets/networks/optimism.png'
import bobaImg from '@assets/networks/boba.png'
import moonbeamImg from '@assets/networks/moonbeam.jpg'
import moonriverImg from '@assets/networks/moonriver.jpeg'
import cronosImg from '@assets/networks/cronos.png'
import metisImg from '@assets/networks/metis.png'
import klaytnImg from '@assets/networks/klaytn.jpeg'
import cantoImg from '@assets/networks/canto.png'
import dogeImg from '@assets/networks/doge.png'
import terraImg from '@assets/networks/terra.png'
import baseImg from '@assets/networks/base.svg'
import { toHexStr } from '@utils/toHexStr'

export const NetworkContextName = 'DEFAULT_NETWORK'

export const ChainId = {
  ETH: 1,
  ROPSTEN: 3,
  RINKEBY: 4,
  GÖRLI: 5,
  OPTIMISM: 10,
  CRONOS: 25,
  KOVAN: 42,
  BSC: 56,
  POLYGON: 137,
  FANTOM: 250,
  BOBA: 288,
  METIS: 1088,
  MOONBEAM: 1284,
  MOONRIVER: 1285,
  KLAYTN: 8217,
  HARDHAT: 31337,
  ARBITRUM: 42161,
  AVALANCHE: 43114,
  DFK: 53935,
  AURORA: 1313161554,
  HARMONY: 1666600000,
  CANTO: 7700,
  DOGECHAIN: 2000,
  BASE: 8453,
  TERRA: 121014925, //"columbus-5", the day columbus reportedly landed in america followed by 5
}

export const INVERTED_CHAIN_ID_MAP = Object.fromEntries(
  Object.entries(ChainId).map(([k, v]) => [v, k])
)

export const CHAIN_INFO_MAP = {
  [ChainId.ETH]: {
    chainId: ChainId.ETH,
    chainSymbol: 'ETH',
    chainName: 'Ethereum',
    chainLogo: ethLogo,
    chainImg: ethImg,
    layer: 1,
  },
  [ChainId.OPTIMISM]: {
    chainId: ChainId.OPTIMISM,
    chainSymbol: 'OPTIMISM',
    chainName: 'Optimism',
    chainLogo: optimismLogo,
    chainImg: optimismImg,
    layer: 2,
  },
  [ChainId.CRONOS]: {
    chainId: ChainId.CRONOS,
    chainSymbol: 'CRONOS',
    chainName: 'Cronos',
    chainLogo: cronosLogo,
    chainImg: cronosImg,
    layer: 1,
  },
  [ChainId.BSC]: {
    chainId: ChainId.BSC,
    chainSymbol: 'BNB',
    chainName: 'BNB Chain',
    chainShortName: 'BNB',
    chainLogo: binanceLogo,
    chainImg: bscImg,
    layer: 1,
  },
  [ChainId.POLYGON]: {
    chainId: ChainId.POLYGON,
    chainSymbol: 'POLYGON',
    chainName: 'Polygon',
    chainLogo: polygonLogo,
    chainImg: polygonImg,
    layer: 2,
  },
  [ChainId.FANTOM]: {
    chainId: ChainId.FANTOM,
    chainSymbol: 'FANTOM',
    chainName: 'Fantom',
    chainLogo: fantomLogo,
    chainImg: fantomImg,
    layer: 1,
  },
  [ChainId.BOBA]: {
    chainId: ChainId.BOBA,
    chainSymbol: 'BOBA',
    chainName: 'Boba Network',
    chainLogo: bobaLogo,
    chainImg: bobaImg,
    layer: 2,
  },
  [ChainId.METIS]: {
    chainId: ChainId.METIS,
    chainSymbol: 'METIS',
    chainName: 'Metis',
    chainLogo: metisLogo,
    chainImg: metisImg,
    layer: 2,
  },
  [ChainId.MOONBEAM]: {
    chainId: ChainId.MOONBEAM,
    chainSymbol: 'MOONBEAM',
    chainName: 'Moonbeam',
    chainLogo: moonbeamImg,
    chainImg: moonbeamImg,
    layer: 1,
  },
  [ChainId.MOONRIVER]: {
    chainId: ChainId.MOONRIVER,
    chainSymbol: 'MOONRIVER',
    chainName: 'Moonriver',
    chainLogo: moonriverImg,
    chainImg: moonriverImg,
    layer: 1,
  },
  [ChainId.ARBITRUM]: {
    chainId: ChainId.ARBITRUM,
    chainSymbol: 'ARBITRUM',
    chainName: 'Arbitrum',
    chainLogo: arbitrumLogo,
    chainImg: arbitrumImg,
    layer: 2,
  },
  [ChainId.AVALANCHE]: {
    chainId: ChainId.AVALANCHE,
    chainSymbol: 'AVALANCHE',
    chainName: 'Avalanche',
    chainLogo: avalancheLogo,
    chainImg: avalancheImg,
    layer: 1,
  },
  [ChainId.DFK]: {
    chainId: ChainId.DFK,
    chainSymbol: 'DFK',
    chainName: 'DFK Chain',
    chainLogo: dfkImg,
    chainImg: dfkImg,
    layer: 1,
  },
  [ChainId.AURORA]: {
    chainId: ChainId.AURORA,
    chainSymbol: 'AURORA',
    chainName: 'Aurora',
    chainLogo: auroraLogo,
    chainImg: auroraImg,
    layer: 1,
  },
  [ChainId.HARMONY]: {
    chainId: ChainId.HARMONY,
    chainSymbol: 'HARMONY',
    chainName: 'Harmony',
    chainLogo: harmonyLogo,
    chainImg: harmonyImg,
    layer: 1,
  },
  [ChainId.KLAYTN]: {
    chainId: ChainId.KLAYTN,
    chainSymbol: 'KLAYTN',
    chainName: 'Klaytn',
    chainLogo: harmonyLogo,
    chainImg: klaytnImg,
    layer: 1,
  },
  [ChainId.CANTO]: {
    chainId: ChainId.CANTO,
    chainSymbol: 'CANTO',
    chainName: 'Canto',
    chainLogo: cantoLogo,
    chainImg: cantoImg,
    layer: 2,
  },
  [ChainId.DOGECHAIN]: {
    chainId: ChainId.DOGECHAIN,
    chainSymbol: 'DOGE',
    chainName: 'Dogechain',
    chainLogo: dogeLogo,
    chainImg: dogeImg,
    layer: 2,
  },
  [ChainId.BASE]: {
    chainId: ChainId.BASE,
    chainSymbol: 'BASE',
    chainName: 'Base',
    chainLogo: baseImg,
    chainImg: baseImg,
    layer: 2,
  },
  // NON_EVM
  [ChainId.TERRA]: {
    chainId: ChainId.TERRA,
    chainSymbol: 'TERRA',
    chainName: 'Terra',
    chainLogo: terraImg,
    chainImg: terraImg,
    layer: 1,
  },
}

// export const SUPPORTED_CHAINS = Object.keys(CHAIN_INFO_MAP).filter(({chainId}) => _.isNumber(chainId))

export const CHAIN_ID_DISPLAY_ORDER = [
  ChainId.ETH,
  ChainId.TERRA,
  ChainId.AVALANCHE,
  ChainId.ARBITRUM,
  ChainId.AURORA,
  ChainId.BSC,
  ChainId.BASE,
  ChainId.BOBA,
  ChainId.CRONOS,
  ChainId.KLAYTN,
  ChainId.DFK,
  ChainId.FANTOM,
  ChainId.HARMONY,
  ChainId.POLYGON,
  ChainId.METIS,
  ChainId.MOONBEAM,
  ChainId.MOONRIVER,
  ChainId.OPTIMISM,
]

// main read rpcs, for write, it may be overridden
export const CHAIN_RPC = {
  [ChainId.ETH]:
    'https://eth-mainnet.alchemyapi.io/v2/2KGc5anzcDyuKCRT3EXPd7QanER_vvJq', // 'https://eth-mainnet.alchemyapi.io/v2/0AovFRYl9L7l4YUf6nPaMrs7H2_pj_Pf',
  [ChainId.OPTIMISM]: 'https://mainnet.optimism.io', //'https://mainnet.optimism.io',
  [ChainId.BSC]: 'https://bsc-dataseed1.ninicoin.io/', //https://bscrpc.com/  // 'https://bsc-dataseed1.ninicoin.io/' 'https://bsc-mainnet.gateway.pokt.network/v1/lb/6136201a7bad1500343e248d',//'https://bsc-dataseed1.binance.org/',                                      // this will prob backfire but hey 'https://bsc-dataseed.binance.org/',
  [ChainId.FANTOM]: 'https://rpc.ftm.tools/',
  [ChainId.BOBA]: 'https://replica-oolong.boba.network/',
  [ChainId.MOONBEAM]: 'https://moonbeam.api.onfinality.io/public', // 'https://replica-boba-synapse.boba.network/', /** read only rpc */
  [ChainId.MOONRIVER]: 'https://rpc.api.moonriver.moonbeam.network/', // 'https://moonriver.api.onfinality.io/public',
  [ChainId.POLYGON]: 'https://polygon-rpc.com/', // NEED TO TEST POLYGON PRIOR TO DEPLOY
  [ChainId.AVALANCHE]: 'https://api.avax.network/ext/bc/C/rpc',
  [ChainId.DFK]: 'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
  [ChainId.ARBITRUM]: 'https://arb1.arbitrum.io/rpc',
  [ChainId.AURORA]: 'https://mainnet.aurora.dev',
  [ChainId.HARMONY]: 'https://harmony-0-rpc.gateway.pokt.network/', // 'https://api.harmony.one',
  [ChainId.CRONOS]: 'https://evm-cronos.crypto.org',
  [ChainId.METIS]: 'https://andromeda.metis.io/?owner=1088',
  [ChainId.KLAYTN]: 'https://cypress.chain.thebifrost.io/',
  [ChainId.BASE]: 'https://mainnet.base.org',

  // [ChainId.XDAI]: 'https://rpc.xdaichain.com',
}

export const NON_EVM_CHAIN_RPC = {
  [ChainId.TERRA]: 'http://public-node.terra.dev:26657/',
}

export const CHAIN_EXPLORER_URLS = {
  [ChainId.BSC]: 'https://bscscan.com',
  [ChainId.ETH]: 'https://etherscan.com',
  [ChainId.POLYGON]: 'https://polygonscan.com',
  [ChainId.BOBA]: 'https://blockexplorer.boba.network',
  [ChainId.MOONBEAM]: 'https://moonbeam.moonscan.io',
  [ChainId.MOONRIVER]: 'https://moonriver.moonscan.io',
  [ChainId.ARBITRUM]: 'https://arbiscan.io',
  [ChainId.OPTIMISM]: 'https://optimistic.etherscan.io',
  [ChainId.AVALANCHE]: 'https://avascan.info/blockchain/c/',
  [ChainId.DFK]:
    'https://subnets.avax.network/defi-kingdoms/dfk-chain/explorer',
  [ChainId.FANTOM]: 'https://ftmscan.com',
  [ChainId.HARMONY]: 'https://explorer.harmony.one',
  [ChainId.AURORA]: 'https://explorer.mainnet.aurora.dev',
  [ChainId.CRONOS]: 'https://cronoscan.com',
  [ChainId.METIS]: 'https://andromeda-explorer.metis.io',
  [ChainId.KLAYTN]: 'https://scope.klaytn.com',
  [ChainId.TERRA]: 'https://terrasco.pe/mainnet', //'https://finder.terra.money/mainnet',
  [ChainId.DOGECHAIN]: 'https://explorer.dogechain.dog',
  [ChainId.CANTO]: 'https://tuber.build',
  [ChainId.BASE]: 'https://basescan.org',
}

/** write rpcs */
const WRITE_CHAIN_RPC = {
  ...CHAIN_RPC,
  [ChainId.ETH]:
    'https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161',
  [ChainId.BOBA]: 'https://mainnet.boba.network/',
}

const ETH_NATIVE_CURRENCY = {
  name: 'Ethereum',
  symbol: 'ETH',
  decimals: 18,
}

/**
 * The below need to be MetaMask compatible keys/objects.
 * extra keys can cause MetaMask to cause really unexpected errors
 * - The order is the order displayed in the chain selector, this is stupid but deal with it
 */
export const CHAIN_PARAMS = {
  [ChainId.ETH]: {
    chainId: toHexStr(ChainId.ETH),
    chainName: CHAIN_INFO_MAP[ChainId.ETH].chainName,
    nativeCurrency: ETH_NATIVE_CURRENCY,
    rpcUrls: [WRITE_CHAIN_RPC[ChainId.ETH]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.ETH]], // might need to add / after explorer url
  },
  [ChainId.OPTIMISM]: {
    chainId: toHexStr(ChainId.OPTIMISM),
    chainName: CHAIN_INFO_MAP[ChainId.OPTIMISM].chainName,
    nativeCurrency: ETH_NATIVE_CURRENCY,
    rpcUrls: [CHAIN_RPC[ChainId.OPTIMISM]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.OPTIMISM]],
  },
  [ChainId.CRONOS]: {
    chainId: toHexStr(ChainId.CRONOS),
    chainName: CHAIN_INFO_MAP[ChainId.CRONOS].chainName,
    nativeCurrency: {
      name: 'Cronos',
      symbol: 'CRO',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.CRONOS]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.CRONOS]],
  },
  [ChainId.BSC]: {
    chainId: toHexStr(ChainId.BSC),
    chainName: CHAIN_INFO_MAP[ChainId.BSC].chainName,
    nativeCurrency: {
      name: 'Binance Coin',
      symbol: 'BNB',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.BSC]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.BSC]],
  },
  [ChainId.POLYGON]: {
    chainId: toHexStr(ChainId.POLYGON),
    chainName: CHAIN_INFO_MAP[ChainId.POLYGON].chainName,
    nativeCurrency: {
      name: 'Matic',
      symbol: 'MATIC',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.POLYGON]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.POLYGON]],
  },
  [ChainId.FANTOM]: {
    chainId: toHexStr(ChainId.FANTOM),
    chainName: CHAIN_INFO_MAP[ChainId.FANTOM].chainName,
    nativeCurrency: {
      name: 'Fantom',
      symbol: 'FTM',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.FANTOM]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.FANTOM]],
  },
  [ChainId.BOBA]: {
    chainId: toHexStr(ChainId.BOBA),
    chainName: CHAIN_INFO_MAP[ChainId.BOBA].chainName,
    nativeCurrency: ETH_NATIVE_CURRENCY,
    rpcUrls: [WRITE_CHAIN_RPC[ChainId.BOBA]], // NOTE: this one uses a dedicated write RPC
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.BOBA]],
  },
  [ChainId.METIS]: {
    chainId: toHexStr(ChainId.METIS),
    chainName: CHAIN_INFO_MAP[ChainId.METIS].chainName,
    nativeCurrency: {
      name: 'Metis',
      symbol: 'METIS',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.METIS]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.METIS]],
  },
  [ChainId.MOONBEAM]: {
    chainId: toHexStr(ChainId.MOONBEAM),
    chainName: CHAIN_INFO_MAP[ChainId.MOONBEAM].chainName,
    nativeCurrency: {
      name: 'Glimmer',
      symbol: 'GLMR',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.MOONBEAM]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.MOONBEAM]],
  },
  [ChainId.MOONRIVER]: {
    chainId: toHexStr(ChainId.MOONRIVER),
    chainName: CHAIN_INFO_MAP[ChainId.MOONRIVER].chainName,
    nativeCurrency: {
      name: 'Moonriver',
      symbol: 'MOVR',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.MOONRIVER]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.MOONRIVER]],
  },
  [ChainId.ARBITRUM]: {
    chainId: toHexStr(ChainId.ARBITRUM),
    chainName: CHAIN_INFO_MAP[ChainId.ARBITRUM].chainName,
    nativeCurrency: ETH_NATIVE_CURRENCY,
    rpcUrls: [CHAIN_RPC[ChainId.ARBITRUM]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.ARBITRUM]],
  },
  [ChainId.AVALANCHE]: {
    chainId: toHexStr(ChainId.AVALANCHE),
    chainName: CHAIN_INFO_MAP[ChainId.AVALANCHE].chainName,
    nativeCurrency: {
      name: 'Avax',
      symbol: 'AVAX',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.AVALANCHE]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.AVALANCHE]],
  },
  [ChainId.DFK]: {
    chainId: toHexStr(ChainId.DFK),
    chainName: CHAIN_INFO_MAP[ChainId.DFK].chainName,
    nativeCurrency: {
      name: 'Jewel',
      symbol: 'JEWEL',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.DFK]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.DFK]],
  },
  [ChainId.AURORA]: {
    chainId: toHexStr(ChainId.AURORA),
    chainName: CHAIN_INFO_MAP[ChainId.AURORA].chainName,
    nativeCurrency: ETH_NATIVE_CURRENCY,
    rpcUrls: [CHAIN_RPC[ChainId.AURORA]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.AURORA]],
  },
  [ChainId.HARMONY]: {
    chainId: toHexStr(ChainId.HARMONY),
    chainName: CHAIN_INFO_MAP[ChainId.HARMONY].chainName,
    nativeCurrency: {
      name: 'Harmony One',
      symbol: 'ONE',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.HARMONY]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.HARMONY]],
  },
  [ChainId.KLAYTN]: {
    chainId: toHexStr(ChainId.KLAYTN),
    chainName: CHAIN_INFO_MAP[ChainId.KLAYTN].chainName,
    nativeCurrency: {
      name: 'Klaytn',
      symbol: 'KLAY',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.KLAYTN]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.KLAYTN]],
  },
  [ChainId.CANTO]: {
    chainId: toHexStr(ChainId.CANTO),
    chainName: CHAIN_INFO_MAP[ChainId.CANTO].chainName,
    nativeCurrency: {
      name: 'Canto',
      symbol: 'CANTO',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.CANTO]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.CANTO]],
  },
  [ChainId.DOGECHAIN]: {
    chainId: toHexStr(ChainId.DOGECHAIN),
    chainName: CHAIN_INFO_MAP[ChainId.DOGECHAIN].chainName,
    nativeCurrency: {
      name: 'Doge Chain',
      symbol: 'DC',
      decimals: 18,
    },
    rpcUrls: [CHAIN_RPC[ChainId.DOGECHAIN]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.DOGECHAIN]],
  },
  [ChainId.BASE]: {
    chainId: toHexStr(ChainId.BASE),
    chainName: CHAIN_INFO_MAP[ChainId.BASE].chainName,
    nativeCurrency: ETH_NATIVE_CURRENCY,
    rpcUrls: [CHAIN_RPC[ChainId.BASE]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.BASE]],
  },
  [ChainId.TERRA]: {
    chainId: toHexStr(ChainId.TERRA),
    chainName: CHAIN_INFO_MAP[ChainId.TERRA].chainName,
    nativeCurrency: {
      name: 'Luna',
      symbol: 'LUNA',
      decimals: 18,
    },
    rpcUrls: [NON_EVM_CHAIN_RPC[ChainId.TERRA]],
    blockExplorerUrls: [CHAIN_EXPLORER_URLS[ChainId.TERRA]],
  },
}

/**
 * NOTE: this is currently set to a far lower value than what it normally is.  normally 2500
 */
export const BLOCK_TIME = 5000 // 30000   // 5000

export const CHAIN_BLOCK_TIME = {
  [ChainId.ETH]: 10000, // BLOCK_TIME,   // 15000,
  [ChainId.OPTIMISM]: 10000,
  [ChainId.CRONOS]: 10000,
  [ChainId.BSC]: 10000, // 5000,
  [ChainId.POLYGON]: 10000, // 5000,
  [ChainId.FANTOM]: BLOCK_TIME,
  [ChainId.BOBA]: 20000,
  [ChainId.METIS]: 10000,
  [ChainId.MOONBEAM]: 10000,
  [ChainId.MOONRIVER]: BLOCK_TIME,
  [ChainId.ARBITRUM]: BLOCK_TIME,
  [ChainId.AVALANCHE]: BLOCK_TIME,
  [ChainId.DFK]: 10000,
  [ChainId.AURORA]: 10000,
  [ChainId.HARMONY]: 10000,
  [ChainId.KLAYTN]: 10000,
  [ChainId.BASE]: 10000,
}

export const CHAIN_ENUM_BY_ID = {
  [ChainId.ETH]: 'ethereum',
  [ChainId.OPTIMISM]: 'optimism',
  [ChainId.CRONOS]: 'cronos',
  [ChainId.BSC]: 'bsc',
  [ChainId.POLYGON]: 'polygon',
  [ChainId.FANTOM]: 'fantom',
  [ChainId.BOBA]: 'boba',
  [ChainId.METIS]: 'metis',
  [ChainId.MOONBEAM]: 'moonbeam',
  [ChainId.MOONRIVER]: 'moonriver',
  [ChainId.KLAYTN]: 'klaytn',
  [ChainId.ARBITRUM]: 'arbitrum',
  [ChainId.AVALANCHE]: 'avalanche',
  [ChainId.DFK]: 'dfk',
  [ChainId.AURORA]: 'aurora',
  [ChainId.HARMONY]: 'harmony',
  [ChainId.CANTO]: 'canto',
  [ChainId.DOGECHAIN]: 'dogechain',
  [ChainId.BASE]: 'base',
}

export const CHAIN_ID_NAMES_REVERSE = Object.fromEntries(
  Object.entries(CHAIN_ENUM_BY_ID).map(([k, v]) => [v, k])
)

export const BRIDGE_CONTRACTS = {
  [ChainId.ETH]: '0x2796317b0fF8538F253012862c06787Adfb8cEb6',
  [ChainId.OPTIMISM]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.CRONOS]: '0xE27BFf97CE92C3e1Ff7AA9f86781FDd6D48F5eE9',
  [ChainId.BSC]: '0xd123f70AE324d34A9E76b67a27bf77593bA8749f',
  [ChainId.POLYGON]: '0x8F5BBB2BB8c2Ee94639E55d5F41de9b4839C1280',
  [ChainId.FANTOM]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.BOBA]: '0x432036208d2717394d2614d6697c46DF3Ed69540',
  [ChainId.METIS]: '0x06Fea8513FF03a0d3f61324da709D4cf06F42A5c',
  [ChainId.MOONBEAM]: '0x84A420459cd31C3c34583F67E0f0fB191067D32f',
  [ChainId.MOONRIVER]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [ChainId.KLAYTN]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.ARBITRUM]: '0x6F4e8eBa4D337f874Ab57478AcC2Cb5BACdc19c9',
  [ChainId.AVALANCHE]: '0xC05e61d0E7a63D27546389B7aD62FdFf5A91aACE',
  [ChainId.DFK]: '0xE05c976d3f045D0E6E7A6f61083d98A15603cF6A',
  [ChainId.AURORA]: '0xaeD5b25BE1c3163c907a471082640450F928DDFE',
  [ChainId.HARMONY]: '0xAf41a65F786339e7911F4acDAD6BD49426F2Dc6b',
  [ChainId.CANTO]: '0xDde5BEC4815E1CeCf336fb973Ca578e8D83606E0',
  [ChainId.DOGECHAIN]: '0x9508BF380c1e6f751D97604732eF1Bae6673f299',
  [ChainId.BASE]: '0xf07d1C752fAb503E47FEF309bf14fbDD3E867089',
}

export const CCTP_CONTRACTS = {
  [ChainId.ETH]: '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48',
  [ChainId.ARBITRUM]: '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48',
  [ChainId.AVALANCHE]: '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48',
  [ChainId.BASE]: '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48',
}
