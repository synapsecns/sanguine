import { ETH, SYN } from '@constants/tokens/bridgeable'
import { Token } from '@types'
import * as CHAINS from '@constants/chains/master'
import { SYN_ETH_SUSHI_TOKEN } from '@constants/tokens/sushiMaster'
// Hardcoding this shit for now until actual plan around routing
let SYNAPSE_BASE_URL = ''
if (process?.env?.NODE_ENV === 'development') {
  SYNAPSE_BASE_URL = 'http://localhost:3000'
} else {
  SYNAPSE_BASE_URL = 'https://synapseprotocol.com'
}

export { SYNAPSE_BASE_URL }

export const BASE_PATH = '/'
export const EXPLORER_KAPPA = 'https://explorer.synapseprotocol.com/tx/'
export const EXPLORER_PATH = 'https://explorer.synapseprotocol.com/'
export const AIRDROP_PATH = '/claim'
export const SWAP_PATH = '/swap'
export const STAKE_PATH = '/stake'
export const POOLS_PATH = '/pools'
export const POOL_PATH = '/pool'
export const BRIDGE_PATH = '/'
export const CONTRACTS_PATH =
  'https://docs.synapseprotocol.com/reference/contract-addresses'
export const INTERCHAIN_LINK = 'https://interchain.synapseprotocol.com/'
export const SOLANA_LINK = 'https://solana.synapseprotocol.com/'
export const STATISTICS_PATH = '/statistics'
export const LANDING_PATH = '/landing'
export const TERMS_OF_SERVICE_PATH =
  'https://explorer.synapseprotocol.com/terms'
export const PRIVACY_POLICY_PATH =
  'https://explorer.synapseprotocol.com/privacy'
export const SYNAPSE_PFP_PATH = '/returntomonke'

export const getPoolUrl = (token: Token) => {
  if (token?.symbol && token.symbol === SYN_ETH_SUSHI_TOKEN.symbol) {
    return getSushiSwapUrl({
      fromCoin: ETH,
      toCoin: SYN,
      chainId: CHAINS.ETH.id,
    })
  } else {
    return `${POOL_PATH}/${token?.routerIndex}`
  }
}

export const getExplorerTxUrl = ({
  hash,
  data,
  chainId = 1,
  type = 'tx',
}: {
  hash?: string
  data?: string
  chainId?: number
  type?: string
}) => {
  let explorerUrl = Object.values(CHAINS).filter(
    (chain) => chain.id === chainId
  )[0].explorerUrl
  let baseUrl = explorerUrl ?? CHAINS.ETH.explorerUrl

  return `${baseUrl}/${type}/${hash ?? data}`
}

export const getExplorerAddressUrl = ({ address, chainId }) => {
  let explorerUrl = Object.values(CHAINS).filter(
    (chain) => chain.id === chainId
  )[0].explorerUrl
  let baseUrl = explorerUrl ?? CHAINS.ETH.explorerUrl

  return `${baseUrl}/address/${address}`
}

export const getCompleteUrl = (uriPath: string) => {
  return `${SYNAPSE_BASE_URL}${uriPath}`
}

export const DOCS_URL = 'https://docs.synapseprotocol.com'
export const DISCORD_URL = 'https://discord.gg/synapseprotocol'
export const TELEGRAM_URL = 'https://t.me/synapseprotocol'
export const FORUM_URL = 'https://forum.synapseprotocol.com/'
export const TWITTER_URL = 'https://twitter.com/SynapseProtocol'

export const BUILD_ON_URL =
  'https://docs.synapseprotocol.com/synapse-interchain-network-sin/build-on-the-synapse-interchain-network'
export const GITHUB_URL = 'https://github.com/synapsecns'
export const MEDIUM_URL = 'https://synapseprotocol.medium.com/'
export const MIRROR_URL = 'https://synapse.mirror.xyz/'

export const HOW_TO_BRIDGE_URL =
  'https://docs.synapseprotocol.com/how-to/bridge'
export const HOW_TO_SWAP_URL = 'https://docs.synapseprotocol.com/how-to/swap'
export const HOW_TO_STAKE_URL =
  'https://docs.synapseprotocol.com/how-to/provide-liquidity'
// Patching this as docs for now need to swap w/ git link

export const SYNAPSE_DOCS_URL = 'https://docs.synapseprotocol.com'

const SUSHISWAP_BASE_URL = 'https://app.sushi.com'
// Need to switch this with fei url
const getSushiSwapUrl = ({
  fromCoin,
  toCoin,
  chainId,
}: {
  fromCoin?: any
  toCoin?: any
  chainId: number
}) => {
  const inputCurrency = fromCoin?.addresses?.[chainId] ?? ''
  const outputCurrency = toCoin?.addresses?.[chainId] ?? ''
  return `${SUSHISWAP_BASE_URL}/swap?inputCurrency=${inputCurrency}&outputCurrency=${outputCurrency}`
}

/** Thanks @0xngmi for building defillama as a whole, it may be a thankless job but we appreciate it */
export const LLAMA_API_URL = 'https://api.llama.fi/protocol/synapse'

export const BRIDGESYN_ANALYTICS_API = 'https://explorer.omnirpc.io/graphql'
