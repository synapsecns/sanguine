export {
  FAST_BRIDGE_ROUTER_ADDRESS_MAP,
  INTENTS_SUPPORTED_CHAIN_IDS,
} from '@synapsecns/sdk-router'

import { CHAINS } from './chains'

export * from './slippage'

export const EXPLORER_GRAPHQL_URL = 'https://explorer.omnirpc.io/graphql'

export const VALID_BRIDGE_MODULES = [
  'SynapseBridge',
  'SynapseCCTP',
  'SynapseRFQ',
  'Gas.zip',
]

export const ZeroAddress = '0x0000000000000000000000000000000000000000'
export const NativeGasAddress = '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'

export const SUPPORTED_SWAP_CHAIN_IDS = [
  CHAINS.ARBITRUM.id,
  CHAINS.AURORA.id,
  CHAINS.AVALANCHE.id,
  CHAINS.BASE.id,
  CHAINS.BLAST.id,
  CHAINS.BNBCHAIN.id,
  CHAINS.BOBA.id,
  CHAINS.CANTO.id,
  CHAINS.CRONOS.id,
  CHAINS.ETHEREUM.id,
  CHAINS.FANTOM.id,
  CHAINS.HARMONY.id,
  CHAINS.KLAYTN.id,
  CHAINS.METIS.id,
  CHAINS.OPTIMISM.id,
  CHAINS.POLYGON.id,
]
