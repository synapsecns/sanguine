import { SupportedChainId } from './chainIds'

export const PUBLIC_PROVIDER_URLS: { [chainId: number]: string } = {
  [SupportedChainId.ETH]: 'https://eth.llamarpc.com',
  [SupportedChainId.OPTIMISM]: 'https://mainnet.optimism.io',
  [SupportedChainId.CRONOS]: 'https://evm.cronos.org',
  [SupportedChainId.BSC]: 'https://bsc-dataseed1.ninicoin.io/',
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
  [SupportedChainId.ARBITRUM]: 'https://arb1.arbitrum.io/rpc',
  [SupportedChainId.AVALANCHE]: 'https://api.avax.network/ext/bc/C/rpc',
  [SupportedChainId.DFK]:
    'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
  [SupportedChainId.AURORA]: 'https://mainnet.aurora.dev',
  [SupportedChainId.HARMONY]: 'https://api.s0.t.hmny.io',
}
