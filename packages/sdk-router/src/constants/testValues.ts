import { SupportedChainId } from './supportedChains'

export const PUBLIC_PROVIDER_URLS: { [chainId: number]: string } = {
  [SupportedChainId.ETH]: 'https://eth.llamarpc.com',
  [SupportedChainId.OPTIMISM]: 'https://mainnet.optimism.io',
  [SupportedChainId.CRONOS]: 'https://evm.cronos.org',
  [SupportedChainId.BSC]: 'https://binance.llamarpc.com',
  [SupportedChainId.POLYGON]: 'https://polygon.llamarpc.com',
  [SupportedChainId.FANTOM]: 'https://rpc3.fantom.network',
  [SupportedChainId.BOBA]: 'https://boba-ethereum.gateway.tenderly.co',
  [SupportedChainId.METIS]: 'https://andromeda.metis.io/?owner=1088',
  [SupportedChainId.MOONBEAM]: 'https://moonbeam.public.blastapi.io',
  [SupportedChainId.MOONRIVER]: 'https://moonriver.public.blastapi.io',
  [SupportedChainId.DOGECHAIN]: 'https://rpc.ankr.com/dogechain',
  [SupportedChainId.CANTO]: 'https://mainnode.plexnode.org:8545',
  [SupportedChainId.KLAYTN]: 'https://klaytn.api.onfinality.io/public',
  [SupportedChainId.BASE]: 'https://developer-access-mainnet.base.org',
  [SupportedChainId.ARBITRUM]: 'https://arbitrum.llamarpc.com',
  [SupportedChainId.AVALANCHE]: 'https://api.avax.network/ext/bc/C/rpc',
  [SupportedChainId.DFK]:
    'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
  [SupportedChainId.AURORA]: 'https://mainnet.aurora.dev',
  [SupportedChainId.HARMONY]: 'https://api.s0.t.hmny.io',
}

// Token addresses on Ethereum mainnet
export const ETH_DAI = '0x6B175474E89094C44Da98b954EedeAC495271d0F'
export const ETH_NUSD = '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F'
export const ETH_USDC = '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48'
export const ETH_USDT = '0xdAC17F958D2ee523a2206206994597C13D831ec7'
// Token addresses on Arbitrum mainnet
export const ARB_NETH = '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e'
export const ARB_NUSD = '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688'
export const ARB_USDC = '0xaf88d065e77c8cC2239327C5EDb3A432268e5831'
export const ARB_USDC_E = '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8'
export const ARB_USDT = '0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9'
export const ARB_WETH = '0x82aF49447D8a07e3bd95BD0d56f35241523fBab1'
// Token addresses on Avalanche mainnet
export const AVAX_GOHM = '0x321E7092a180BB43555132ec53AaA65a5bF84251'
export const AVAX_USDC = '0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E'
export const AVAX_USDC_E = '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664'
// Token addresses on BSC mainnet
export const BSC_GOHM = '0x88918495892BAF4536611E38E75D771Dc6Ec0863'
export const BSC_USDC = '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d'

// Pool addresses on Ethereum mainnet
export const ETH_POOL_NUSD = '0x1116898DdA4015eD8dDefb84b6e8Bc24528Af2d8'
// Pool addresses on Arbitrum mainnet
export const ARB_POOL_ETH_WRAPPER = '0x1c3fe783a7c06bfAbd124F2708F5Cc51fA42E102'
export const ARB_POOL_NETH = '0xa067668661C84476aFcDc6fA5D758C4c01C34352'
export const ARB_POOL_NUSD = '0x9Dd329F5411466d9e0C488fF72519CA9fEf0cb40'
