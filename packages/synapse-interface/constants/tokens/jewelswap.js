import synapseLogo from '@assets/icons/synapse.svg'

import { Token } from '@utils/classes/Token'
import { SYNAPSE_DOCS_URL } from '@urls'
import { ChainId } from '@constants/networks'

import { WJEWEL, SYNJEWEL } from '@constants/tokens/basic'

/**
 * Harmony AVAX Swap
 */
export const HARMONY_JEWEL_SWAP_TOKEN = new Token({
  addresses: {
    [ChainId.HARMONY]: '0x0000000000000000000000000000000000000000',
  },
  decimals: 18,
  symbol: 'JEWELP',
  name: 'Jewel LP Token Harmony ',
  logo: synapseLogo,
  poolName: 'Harmony Jewel Swap Pool ', // DONT GET RID OF SPACE AFTER POOL
  routerIndex: 'jewel2pool',
  poolId: 0,
  poolType: 'JEWEL',
  swapAddresses: {
    [ChainId.HARMONY]: '0x7bE461cce1501f07969BCE24Ccb2140fCA0a35b3',
  },
  poolTokens: [WJEWEL, SYNJEWEL],
  description: "Synapse's 2pool JEWEL swapper psuedotoken on Harmony",
  docUrl: SYNAPSE_DOCS_URL,
})
