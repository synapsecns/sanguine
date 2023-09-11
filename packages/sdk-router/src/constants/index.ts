import JSBI from 'jsbi'
import { BigNumber } from '@ethersproject/bignumber'

export enum SupportedChainId {
  ETH = 1,
  OPTIMISM = 10,
  CRONOS = 25,
  BSC = 56,
  POLYGON = 137,
  FANTOM = 250,
  BOBA = 288,
  METIS = 1088,
  MOONBEAM = 1284,
  MOONRIVER = 1285,
  DOGECHAIN = 2000,
  CANTO = 7700,
  KLAYTN = 8217,
  BASE = 8453,
  ARBITRUM = 42161,
  AVALANCHE = 43114,
  DFK = 53935,
  AURORA = 1313161554,
  HARMONY = 1666600000,
}

export const PUBLIC_PROVIDER_URLS: { [chainId: number]: string } = {
  [SupportedChainId.ETH]: 'https://eth.llamarpc.com',
  [SupportedChainId.OPTIMISM]: 'https://mainnet.optimism.io',
  [SupportedChainId.CRONOS]: 'https://evm.cronos.org',
  [SupportedChainId.BSC]: 'https://bsc-dataseed1.ninicoin.io/',
  [SupportedChainId.POLYGON]: 'https://polygon.llamarpc.com',
  [SupportedChainId.FANTOM]: 'https://rpc3.fantom.network',
  [SupportedChainId.BOBA]: 'https://boba-ethereum.gateway.tenderly.co',
  [SupportedChainId.METIS]: 'https://andromeda.metis.io/?owner=1088',
  [SupportedChainId.MOONBEAM]: 'https://moonbeam.public.blastapi.io',
  [SupportedChainId.MOONRIVER]: 'https://moonriver.public.blastapi.io',
  [SupportedChainId.DOGECHAIN]: 'https://rpc.ankr.com/dogechain',
  [SupportedChainId.CANTO]: 'https://mainnode.plexnode.org:8545',
  [SupportedChainId.KLAYTN]: 'https://klaytn.api.onfinality.io/public',
  [SupportedChainId.BASE]: 'https://developer-access-mainnet.base.org',
  [SupportedChainId.ARBITRUM]: 'https://arb1.arbitrum.io/rpc',
  [SupportedChainId.AVALANCHE]: 'https://api.avax.network/ext/bc/C/rpc',
  [SupportedChainId.DFK]:
    'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc',
  [SupportedChainId.AURORA]: 'https://mainnet.aurora.dev',
  [SupportedChainId.HARMONY]: 'https://api.s0.t.hmny.io',
}

/**
 * List of supported chain ids, where SynapseBridge is deployed.
 */
export const SUPPORTED_CHAIN_IDS: number[] = Object.values(SupportedChainId)
  .map((chainId) => Number(chainId))
  .filter((chainId) => !isNaN(chainId))

/**
 * List of chain ids where SynapseCCTP is deployed.
 *
 * Note: This is a subset of SUPPORTED_CHAIN_IDS.
 */
export const CCTP_SUPPORTED_CHAIN_IDS: number[] = [
  SupportedChainId.ETH,
  SupportedChainId.ARBITRUM,
  SupportedChainId.AVALANCHE,
  SupportedChainId.OPTIMISM,
]

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
const CCTP_ROUTER_ADDRESS = '0xd359bc471554504f683fbd4f6e36848612349ddf'
const CCTP_ROUTER_EXCEPTION_MAP: AddressMap = {}

export const CCTP_ROUTER_ADDRESS_MAP: AddressMap = generateAddressMap(
  CCTP_SUPPORTED_CHAIN_IDS,
  CCTP_ROUTER_ADDRESS,
  CCTP_ROUTER_EXCEPTION_MAP
)

// exports for external consumption
export type BigintIsh = JSBI | BigNumber | string | number

export enum TradeType {
  EXACT_INPUT,
  EXACT_OUTPUT,
}

export enum Rounding {
  ROUND_DOWN,
  ROUND_HALF_UP,
  ROUND_UP,
}

export const MaxUint256 = JSBI.BigInt(
  '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
)
