import {
  CCTP_SUPPORTED_CHAIN_IDS,
  GASZIP_SUPPORTED_CHAIN_IDS,
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
 * FastBridgeInterceptor contract address for all chains except ones from FBI_EXCEPTION_MAP.
 */
const FBI_ADDRESS = '0xFb1fb1060C550A9b274C64f70dadF16f2aD34fB1'
const FBI_EXCEPTION_MAP: AddressMap = {}
export const FAST_BRIDGE_INTERCEPTOR_ADDRESS_MAP: AddressMap =
  generateAddressMap(RFQ_SUPPORTED_CHAIN_IDS, FBI_ADDRESS, FBI_EXCEPTION_MAP)

/**
 * GasZip contract address for all chains except ones from GASZIP_EXCEPTION_MAP.
 */
const GAS_ZIP_ADDRESS = '0x2a37D63EAdFe4b4682a3c28C1c2cD4F109Cc2762'
const GASZIP_EXCEPTION_MAP: AddressMap = {
  [SupportedChainId.FANTOM]: '0xA60768b03eB14d940F6c9a8553329B7F9037C91b',
  [SupportedChainId.LINEA]: '0xA60768b03eB14d940F6c9a8553329B7F9037C91b',
}
export const GASZIP_ADDRESS_MAP: AddressMap = generateAddressMap(
  GASZIP_SUPPORTED_CHAIN_IDS,
  GAS_ZIP_ADDRESS,
  GASZIP_EXCEPTION_MAP
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
 * Empty address is used to avoid filling the address map with incorrect addresses.
 * TODO: unified deployments
 */
const SWAP_QUOTER_V2_ADDRESS = ''
const SWAP_QUOTER_V2_EXCEPTION_MAP: AddressMap = {
  [SupportedChainId.AVALANCHE]: '0x40d9dDE17D776bF057083E156578d2443685851C',
  [SupportedChainId.ARBITRUM]: '0xE402cC7826dD835FCe5E3cFb61D56703fEbc2642',
  [SupportedChainId.BASE]: '0x9FBFf54b967654B0c76b174D2B95614060Dd6B07',
  [SupportedChainId.BERACHAIN]: '0xc5269d5143B37877A1041DfaF2C21a76E709AF27',
  [SupportedChainId.BLAST]: '0xbAD189BDF6a05FDaFA33CA917d094A64954093c4',
  [SupportedChainId.BSC]: '0x1Db5a1d5D80fDEfc098635d3869Fa94d6fA44F5a',
  [SupportedChainId.ETH]: '0x5682dC851C33adb48F6958a963A5d3Aa31F6f184',
  [SupportedChainId.HYPEREVM]: '0xc5269d5143B37877A1041DfaF2C21a76E709AF27',
  [SupportedChainId.LINEA]: '0x55DEc55aDbd9a2102438339A294CB921A5248285',
  [SupportedChainId.OPTIMISM]: '0xd6Bdb96b356F4F51bf491297DF03F25DCd0cBf6D',
  [SupportedChainId.POLYGON]: '0x7443C01542f4913f276AED4D617117f3E8FAD21e',
  [SupportedChainId.SCROLL]: '0x55DEc55aDbd9a2102438339A294CB921A5248285',
  [SupportedChainId.UNICHAIN]: '0x55DEc55aDbd9a2102438339A294CB921A5248285',
  [SupportedChainId.WORLDCHAIN]: '0xc5269d5143B37877A1041DfaF2C21a76E709AF27',
}
export const SWAP_QUOTER_V2_ADDRESS_MAP: AddressMap = generateAddressMap(
  INTENTS_SUPPORTED_CHAIN_IDS,
  SWAP_QUOTER_V2_ADDRESS,
  SWAP_QUOTER_V2_EXCEPTION_MAP
)
