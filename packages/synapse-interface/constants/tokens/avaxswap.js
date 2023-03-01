import synapseLogo from '@assets/icons/synapse.svg'

import { Token } from '@utils/classes/Token'
import { SYNAPSE_DOCS_URL } from '@urls'
import { ChainId } from '@constants/networks'

import { SYNAVAX, MULTIAVAX } from '@constants/tokens/basic'

/**
 * Harmony AVAX Swap
 */
export const HARMONY_AVAX_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.HARMONY]: '0x02f7D17f126BD54573c8EbAD9e05408A56f46452',
  },
  decimals: 18,
  symbol: 'AVAXLP',
  name: 'AVAX LP Token Harmony ',
  logo: synapseLogo,
  poolName: 'Harmony AVAX Swap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'avax2pool',
  poolId: 0,
  poolType: 'AVAX',
  swapAddresses: {
    [ChainId.HARMONY]: '0x00A4F57D926781f62D09bb05ec76e6D8aE4268da',
  },
  poolTokens: [SYNAVAX, MULTIAVAX],
  description: "Synapse's 2pool AVAX LP token on Harmony",
  docUrl: SYNAPSE_DOCS_URL,
})
