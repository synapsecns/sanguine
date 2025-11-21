import {
  CCTP_SUPPORTED_CHAIN_IDS,
  GASZIP_SUPPORTED_CHAIN_IDS,
  INTENTS_SUPPORTED_CHAIN_IDS,
  RELAY_SUPPORTED_CHAIN_IDS,
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

const RELAY_ADDRESS = '0xBBbfD134E9b44BfB5123898BA36b01dE7ab93d98'
const RELAY_EXCEPTION_MAP: AddressMap = {}
export const RELAY_ADDRESS_MAP: AddressMap = generateAddressMap(
  RELAY_SUPPORTED_CHAIN_IDS,
  RELAY_ADDRESS,
  RELAY_EXCEPTION_MAP
)

export const USDC_ADDRESS_MAP: AddressMap = {
  [SupportedChainId.ARBITRUM]: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
  [SupportedChainId.AVALANCHE]: '0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E',
  [SupportedChainId.BASE]: '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913',
  [SupportedChainId.ETH]: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
  [SupportedChainId.HYPEREVM]: '0xb88339CB7199b77E23DB6E890353E22632Ba630f',
  [SupportedChainId.LINEA]: '0x176211869cA2b568f2A7D4EE941E073a821EE1ff',
  [SupportedChainId.OPTIMISM]: '0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85',
  [SupportedChainId.POLYGON]: '0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359',
  [SupportedChainId.UNICHAIN]: '0x078D782b760474a361dDA0AF3839290b0EF57AD6',
  [SupportedChainId.WORLDCHAIN]: '0x79A02482A880bCE3F13e09Da970dC34db4CD24d1',
}

export const CCTP_V2_EXECUTOR_ADDRESS_MAP: AddressMap = {
  [SupportedChainId.ARBITRUM]: '0x8442d68524217601ed126f6859694E4B0C7c66A1',
  [SupportedChainId.AVALANCHE]: '0x3952914628650Ca510404872D84DfF10A844C5B5',
  [SupportedChainId.BASE]: '0xbd8d42f40a11b37bD1b3770D754f9629F7cd5679',
  [SupportedChainId.ETH]: '0x2cCf230467FE7387674BAa657747F0B5485c7fEC',
  [SupportedChainId.HYPEREVM]: '0xACD054f83c0b852d02503191e2c26527A7E72B1f',
  [SupportedChainId.LINEA]: '0xc48c126468BE919068dE1983F00F65af759a4E87',
  [SupportedChainId.OPTIMISM]: '0xd0A8940b2e743E33b682dAEc4D52b46713606C9D',
  [SupportedChainId.POLYGON]: '0xc8A8E6D760dCBd5d6746E2F66cd2fFA722dd1E59',
  [SupportedChainId.UNICHAIN]: '0xD5D5D640D8b758672Cc7A078734175c4433866d5',
  [SupportedChainId.WORLDCHAIN]: '0x789f2b91f7B35D5B890983328340c4600339B354',
}

/**
 * USDT0 contract address for supported chains.
 * Note: for Ethereum this is actually native USDT.
 */
export const USDT0_ADDRESS_MAP: AddressMap = {
  [SupportedChainId.ARBITRUM]: '0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9',
  [SupportedChainId.BERACHAIN]: '0x779Ded0c9e1022225f8E0630b35a9b54bE713736',
  [SupportedChainId.ETH]: '0xdAC17F958D2ee523a2206206994597C13D831ec7',
  [SupportedChainId.HYPEREVM]: '0xB8CE59FC3717ada4C02eaDF9682A9e934F625ebb',
  [SupportedChainId.OPTIMISM]: '0x01bFF41798a0BcF287b996046Ca68b395DbC1071',
  [SupportedChainId.POLYGON]: '0xc2132D05D31c914a87C6611C10748AEb04B58e8F',
  [SupportedChainId.UNICHAIN]: '0x9151434b16b9763660705744891fA906F660EcC5',
}

/**
 * USDT OFT contract address for supported chains.
 * This is the point of entry for USDT bridge on supported chains, including Ethereum.
 */
export const USDT_OFT_ADDRESS_MAP: AddressMap = {
  [SupportedChainId.ARBITRUM]: '0x14E4A1B13bf7F943c8ff7C51fb60FA964A298D92',
  [SupportedChainId.BERACHAIN]: '0x3Dc96399109df5ceb2C226664A086140bD0379cB',
  [SupportedChainId.ETH]: '0x6C96dE32CEa08842dcc4058c14d3aaAD7Fa41dee',
  [SupportedChainId.HYPEREVM]: '0x904861a24F30EC96ea7CFC3bE9EA4B476d237e98',
  [SupportedChainId.OPTIMISM]: '0xF03b4d9AC1D5d1E7c4cEf54C2A313b9fe051A0aD',
  [SupportedChainId.POLYGON]: '0x6BA10300f0DC58B7a1e4c0e41f5daBb7D7829e13',
  [SupportedChainId.UNICHAIN]: '0xc07bE8994D035631c36fb4a89C918CeFB2f03EC3',
}

export const LZ_EID_MAP: Record<number, number> = {
  [SupportedChainId.ARBITRUM]: 30110,
  [SupportedChainId.BERACHAIN]: 30362,
  [SupportedChainId.ETH]: 30101,
  [SupportedChainId.HYPEREVM]: 30367,
  [SupportedChainId.OPTIMISM]: 30111,
  [SupportedChainId.POLYGON]: 30109,
  [SupportedChainId.UNICHAIN]: 30320,
}

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
