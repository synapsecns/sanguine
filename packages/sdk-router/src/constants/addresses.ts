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
const ROUTER_EXCEPTION_MAP: AddressMap = {}

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

// https://developers.circle.com/stablecoins/docs/evm-smart-contracts#tokenminter-mainnet
export const CCTP_TOKEN_MINTER_MAP: AddressMap = {
  [SupportedChainId.ETH]: '0xc4922d64a24675E16e1586e3e3Aa56C06fABe907',
  [SupportedChainId.AVALANCHE]: '0x420F5035fd5dC62a167E7e7f08B604335aE272b8',
  [SupportedChainId.OPTIMISM]: '0x33E76C5C31cb928dc6FE6487AB3b2C0769B1A1e3',
  [SupportedChainId.ARBITRUM]: '0xE7Ed1fa7f45D05C508232aa32649D89b73b8bA48',
  [SupportedChainId.BASE]: '0xe45B133ddc64bE80252b0e9c75A8E74EF280eEd6',
  [SupportedChainId.POLYGON]: '0x10f7835F827D6Cf035115E10c50A853d7FB2D2EC',
}

/**
 * FastBridgeRouter contract address for all chains except ones from FAST_BRIDGE_ROUTER_EXCEPTION_MAP.
 */
const FAST_BRIDGE_ROUTER_ADDRESS = '0x0000000000489d89D2B233D3375C045dfD05745F'
const FAST_BRIDGE_ROUTER_EXCEPTION_MAP: AddressMap = {}
export const FAST_BRIDGE_ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  RFQ_SUPPORTED_CHAIN_IDS,
  FAST_BRIDGE_ROUTER_ADDRESS,
  FAST_BRIDGE_ROUTER_EXCEPTION_MAP
)
