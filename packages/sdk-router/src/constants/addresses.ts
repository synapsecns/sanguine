import {
  CCTP_SUPPORTED_CHAIN_IDS,
  RFQ_SUPPORTED_CHAIN_IDS,
  SUPPORTED_CHAIN_IDS,
  SupportedChainId,
} from './chainIds'

export type AddressMap = {
  [chainId: number]: string
}

/**
 * Generates an address map for a given address and list of chain ids.
 * Will use the same address for all chain ids unless an exception map is provided.
 * In which case, the exception map will be used to override the address for the
 * specified chain ids.
 *
 * @param chainIds list of chain ids
 * @param address address to use for all chain ids unless overridden by exception map
 * @param exceptionMap optional map of chain ids to addresses to override the address param
 * @returns
 */
const generateAddressMap = (
  chainIds: number[],
  address: string,
  exceptionMap?: AddressMap
): AddressMap => {
  return Object.fromEntries(
    chainIds.map((chainId) => [chainId, exceptionMap?.[chainId] ?? address])
  )
}

/**
 * SynapseRouter contract address for all chains except ones from ROUTER_EXCEPTION_MAP.
 */
const ROUTER_ADDRESS = '0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a'
const ROUTER_EXCEPTION_MAP: AddressMap = {
  [SupportedChainId.BLAST]: '0x0000000000365b1d5B142732CF4d33BcddED21Fc',
}
export const ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  SUPPORTED_CHAIN_IDS,
  ROUTER_ADDRESS,
  ROUTER_EXCEPTION_MAP
)

/**
 * SynapseCCTP contract address for all chains except ones from CCTP_ROUTER_EXCEPTION_MAP.
 */
const CCTP_ROUTER_ADDRESS = '0xd5a597d6e7ddf373a92C8f477DAAA673b0902F48'
const CCTP_ROUTER_EXCEPTION_MAP: AddressMap = {}
export const CCTP_ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  CCTP_SUPPORTED_CHAIN_IDS,
  CCTP_ROUTER_ADDRESS,
  CCTP_ROUTER_EXCEPTION_MAP
)

/**
 * FastBridgeRouter contract address for all chains except ones from FAST_BRIDGE_ROUTER_EXCEPTION_MAP.
 * TODO: this is a temporary Router for FastBridgeV2, revert to FastBridgeV1 when a dedicated Router Set is implemented
 */
const FAST_BRIDGE_ROUTER_ADDRESS = '0x5849dC2fe58EcCB9EA76aA8D77dc127d89dE9b4d'
const FAST_BRIDGE_ROUTER_EXCEPTION_MAP: AddressMap = {}
export const FAST_BRIDGE_ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  RFQ_SUPPORTED_CHAIN_IDS,
  FAST_BRIDGE_ROUTER_ADDRESS,
  FAST_BRIDGE_ROUTER_EXCEPTION_MAP
)

/**
 * FastBridgeV2 contract address for all chains except ones from FAST_BRIDGE_V2_EXCEPTION_MAP.
 * TODO: this is a staging FastBridgeV2 deployment, update to the production deployment when ready.
 */
const FAST_BRIDGE_V2_ADDRESS = '0xEb1eb846342274d5d652e068DD189EfCEd256332'
const FAST_BRIDGE_V2_EXCEPTION_MAP: AddressMap = {}
export const FAST_BRIDGE_V2_ADDRESS_MAP: AddressMap = generateAddressMap(
  RFQ_SUPPORTED_CHAIN_IDS,
  FAST_BRIDGE_V2_ADDRESS,
  FAST_BRIDGE_V2_EXCEPTION_MAP
)

/**
 * TokenZapV1 contract address for all chains except ones from TOKEN_ZAP_V1_EXCEPTION_MAP.
 * TODO: this is a staging TokenZapV1 deployment, update to the production deployment when ready.
 */
const TOKEN_ZAP_V1_ADDRESS = '0x6327797F149a75D506aFda46D5fCE6E74fC409D5'
const TOKEN_ZAP_V1_EXCEPTION_MAP: AddressMap = {}
export const TOKEN_ZAP_V1_ADDRESS_MAP: AddressMap = generateAddressMap(
  RFQ_SUPPORTED_CHAIN_IDS,
  TOKEN_ZAP_V1_ADDRESS,
  TOKEN_ZAP_V1_EXCEPTION_MAP
)

/**
 * SynapseIntentRouter contract address for all chains except ones from SYNAPSE_INTENT_ROUTER_EXCEPTION_MAP.
 * TODO: this is a staging SynapseIntentRouter deployment, update to the production deployment when ready.
 */
const SYNAPSE_INTENT_ROUTER_ADDRESS =
  '0x57203c65DeA2ded4EE4E303a9494bee04df030BF'
const SYNAPSE_INTENT_ROUTER_EXCEPTION_MAP: AddressMap = {}
export const SYNAPSE_INTENT_ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  RFQ_SUPPORTED_CHAIN_IDS,
  SYNAPSE_INTENT_ROUTER_ADDRESS,
  SYNAPSE_INTENT_ROUTER_EXCEPTION_MAP
)

/**
 * SynapseIntentPreviewer contract address for all chains except ones from SYNAPSE_INTENT_PREVIEWER_EXCEPTION_MAP.
 * TODO: this is a staging SynapseIntentPreviewer deployment, update to the production deployment when ready.
 */
const SYNAPSE_INTENT_PREVIEWER_ADDRESS =
  '0xfC2352150681A96E591F3a1ea511970FEF005A55'
const SYNAPSE_INTENT_PREVIEWER_EXCEPTION_MAP: AddressMap = {}
export const SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP: AddressMap =
  generateAddressMap(
    RFQ_SUPPORTED_CHAIN_IDS,
    SYNAPSE_INTENT_PREVIEWER_ADDRESS,
    SYNAPSE_INTENT_PREVIEWER_EXCEPTION_MAP
  )

/**
 * SwapQuoterV2 contract address on Ethereum. Addresses for other chains are defined in the
 * SWAP_QUOTER_V2_EXCEPTION_MAP.
 */
const SWAP_QUOTER_V2_ADDRESS = '0x5682dC851C33adb48F6958a963A5d3Aa31F6f184'
const SWAP_QUOTER_V2_EXCEPTION_MAP: AddressMap = {
  [SupportedChainId.ARBITRUM]: '0xE402cC7826dD835FCe5E3cFb61D56703fEbc2642',
  [SupportedChainId.OPTIMISM]: '0xd6Bdb96b356F4F51bf491297DF03F25DCd0cBf6D',
}
export const SWAP_QUOTER_V2_ADDRESS_MAP: AddressMap = generateAddressMap(
  SUPPORTED_CHAIN_IDS,
  SWAP_QUOTER_V2_ADDRESS,
  SWAP_QUOTER_V2_EXCEPTION_MAP
)
