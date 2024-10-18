import { JsonRpcProvider } from '@ethersproject/providers'
import { SynapseSDK } from '@synapsecns/sdk-router'

import { CHAINS_ARRAY } from '../constants/chains'

const providers = CHAINS_ARRAY.map(
  (chain) =>
    new JsonRpcProvider(chain.rpcUrls.primary || chain.rpcUrls.fallback)
)
const chainIds = CHAINS_ARRAY.map((chain) => chain.id)

export const Synapse = new SynapseSDK(chainIds, providers)
