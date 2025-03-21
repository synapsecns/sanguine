import {
  CCTP_SUPPORTED_CHAIN_IDS,
  INTENTS_SUPPORTED_CHAIN_IDS,
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
 */
const FAST_BRIDGE_ROUTER_ADDRESS = '0x00cD000000003f7F682BE4813200893d4e690000'
const FAST_BRIDGE_ROUTER_EXCEPTION_MAP: AddressMap = {}
export const FAST_BRIDGE_ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  RFQ_SUPPORTED_CHAIN_IDS,
  FAST_BRIDGE_ROUTER_ADDRESS,
  FAST_BRIDGE_ROUTER_EXCEPTION_MAP
)

/**
 * SynapseIntentRouter contract address for all chains except ones from SYNAPSE_INTENT_ROUTER_EXCEPTION_MAP.
 */
const SIR_ADDRESS = '0x512000a034E154908Efb1eC48579F4ffDb000512'
const SYNAPSE_INTENT_ROUTER_EXCEPTION_MAP: AddressMap = {}
export const SYNAPSE_INTENT_ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  INTENTS_SUPPORTED_CHAIN_IDS,
  SIR_ADDRESS,
  SYNAPSE_INTENT_ROUTER_EXCEPTION_MAP
)

/**
 * TokenZapV1 contract address for all chains except ones from TOKEN_ZAP_V1_EXCEPTION_MAP.
 */
const TOKEN_ZAP_V1_ADDRESS = '0x2aAaa9b71E479e6e2De7E091b09D61C25D2AAAa9'
const TOKEN_ZAP_V1_EXCEPTION_MAP: AddressMap = {}
export const TOKEN_ZAP_V1_ADDRESS_MAP: AddressMap = generateAddressMap(
  INTENTS_SUPPORTED_CHAIN_IDS,
  TOKEN_ZAP_V1_ADDRESS,
  TOKEN_ZAP_V1_EXCEPTION_MAP
)

/**
 * SynapseIntentPreviewer contract address for all chains except ones from SYNAPSE_INTENT_PREVIEWER_EXCEPTION_MAP.
 */
const SIP_ADDRESS = '0x519519a57a6Ea930f87e3436b6ea113A990fF519'
const SYNAPSE_INTENT_PREVIEWER_EXCEPTION_MAP: AddressMap = {}
export const SYNAPSE_INTENT_PREVIEWER_ADDRESS_MAP: AddressMap =
  generateAddressMap(
    INTENTS_SUPPORTED_CHAIN_IDS,
    SIP_ADDRESS,
    SYNAPSE_INTENT_PREVIEWER_EXCEPTION_MAP
  )

/**
 * SwapQuoterV2 contract address on Ethereum. Addresses for other chains are defined in the
 * SWAP_QUOTER_V2_EXCEPTION_MAP.
 * TODO: unified deployments
 */
const SWAP_QUOTER_V2_ADDRESS = '0x5682dC851C33adb48F6958a963A5d3Aa31F6f184'
const SWAP_QUOTER_V2_EXCEPTION_MAP: AddressMap = {
  [SupportedChainId.ARBITRUM]: '0xE402cC7826dD835FCe5E3cFb61D56703fEbc2642',
  [SupportedChainId.BASE]: '0x9FBFf54b967654B0c76b174D2B95614060Dd6B07',
  [SupportedChainId.BERACHAIN]: '0xc5269d5143B37877A1041DfaF2C21a76E709AF27',
  [SupportedChainId.BLAST]: '0xbAD189BDF6a05FDaFA33CA917d094A64954093c4',
  [SupportedChainId.BSC]: '0x1Db5a1d5D80fDEfc098635d3869Fa94d6fA44F5a',
  [SupportedChainId.LINEA]: '0x55DEc55aDbd9a2102438339A294CB921A5248285',
  [SupportedChainId.OPTIMISM]: '0xd6Bdb96b356F4F51bf491297DF03F25DCd0cBf6D',
  [SupportedChainId.SCROLL]: '0x55DEc55aDbd9a2102438339A294CB921A5248285',
  [SupportedChainId.UNICHAIN]: '0x55DEc55aDbd9a2102438339A294CB921A5248285',
  [SupportedChainId.WORLDCHAIN]: '0xc5269d5143B37877A1041DfaF2C21a76E709AF27',
}
export const SWAP_QUOTER_V2_ADDRESS_MAP: AddressMap = generateAddressMap(
  INTENTS_SUPPORTED_CHAIN_IDS,
  SWAP_QUOTER_V2_ADDRESS,
  SWAP_QUOTER_V2_EXCEPTION_MAP
)
