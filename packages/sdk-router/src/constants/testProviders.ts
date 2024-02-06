import {
  JsonRpcProvider,
  FallbackProvider,
  FallbackProviderConfig,
} from '@ethersproject/providers'

import { getTestProviderUrl } from './testValues'
import { SupportedChainId } from './chainIds'

const PUBLIC_PROVIDER_URLS: Partial<Record<SupportedChainId, string[]>> = {
  [SupportedChainId.ETH]: [
    // 'https://eth.llamarpc.com',
    'https://ethereum.publicnode.com',
    'https://eth-pokt.nodies.app',
  ],
  [SupportedChainId.OPTIMISM]: [
    // 'https://optimism.llamarpc.com',
    'https://optimism.publicnode.com',
    'https://op-pokt.nodies.app',
  ],
  [SupportedChainId.BSC]: [
    // 'https://binance.llamarpc.com',
    'https://bsc.publicnode.com',
    'https://bsc-pokt.nodies.app',
  ],
  [SupportedChainId.MOONBEAM]: [
    'https://rpc.api.moonbeam.network',
    'https://moonbeam.publicnode.com',
    'https://moonbeam.public.blastapi.io',
  ],
  [SupportedChainId.ARBITRUM]: [
    // 'https://arbitrum.llamarpc.com',
    'https://arbitrum-one.publicnode.com',
    'https://arb-pokt.nodies.app',
  ],
  [SupportedChainId.AVALANCHE]: [
    'https://api.avax.network/ext/bc/C/rpc',
    'https://avalanche-c-chain.publicnode.com',
    'https://avax-pokt.nodies.app/ext/bc/C/rpc',
  ],
}

export const getTestProvider = (
  chainId: SupportedChainId
): FallbackProvider => {
  // Merge the test provider URL with the list of public providers
  const providerUrls = [
    getTestProviderUrl(chainId),
    ...(PUBLIC_PROVIDER_URLS[chainId] || []),
  ]
  // Use index of URL in the list as the priority, this way the "main" test
  // provider will be used first
  const providerConfigs: FallbackProviderConfig[] = providerUrls.map(
    (url, index) => ({
      provider: new JsonRpcProvider(url),
      priority: index,
      stallTimeout: 2000,
    })
  )
  // Use quorum of 1 for tests
  return new FallbackProvider(providerConfigs, 1)
}
