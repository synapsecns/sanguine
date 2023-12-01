import { BigNumber } from '@ethersproject/bignumber'
import * as CHAINS from '@synapse-constants/constants/chains/master'
import { getAddress } from '@ethersproject/address'

export type Chain = {
  id: number
  chainSymbol: string
  name: string
  altName?: string
  codeName: string
  chainImg: any
  layer: number
  rpcUrls: { primary: string; fallback: string }
  explorerUrl: string
  explorerName: string
  explorerImg: any
  blockTime: number
  nativeCurrency: { name: string; symbol: string; decimals: number }
  priorityRank?: number
  color?: string
}
export type PoolToken = {
  symbol: string
  percent: string
  balance: string
  balanceStr: string
  token: Token
  isLP: boolean
  rawBalance: bigint
}
export type Query = [string, string, BigNumber, BigNumber, string] & {
  swapAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}
export type PoolUserData = {
  name: string
  tokens: PoolToken[]
  lpTokenBalance: bigint
  nativeTokens?: any
}
export type PoolData = {
  name: string
  tokens: PoolToken[]
  totalLocked: number
  totalLockedUSD: number
  virtualPrice?: bigint
  nativeTokens?: any
  swapFee?: bigint
}

interface TokensByChain {
  [cID: string]: Token[]
}
export type PoolCardInfo = {
  index: number
  label: string
  poolsByChain: TokensByChain
}
export enum WalletId {
  MetaMask = 'metaMask',
  WalletConnect = 'walletConnect',
  CoinbaseWallet = 'coinbaseWallet',
}
export interface IconProps {
  walletId?: string
  className?: string
}
export type PoolTokenObject = {
  token: Token
  balance: string
  rawBalance: bigint
  isLP: boolean
}
/**
 * Represents an ERC20-like token with a unique address, chainId, and some metadata.
 */
export class Token {
  addresses: { [x: number]: string } // list of token addresses on each chain
  wrapperAddresses?: Record<number, string> // list of wrapper addresses on each chain like gmx
  decimals: number | Record<number, number> = {} // list of decimals on each chain
  symbol?: string // token symbol
  name?: string // token name
  logo?: any // token logo
  icon?: any // token icon
  poolName?: string // token pool name
  swapAddresses?: Record<number, string> // list of swap addresses on each chain
  swapWrapperAddresses?: Record<number, string> // list of swap wrapper addresses on each chain
  swapDepositAddresses?: Record<number, string> // list of swap deposit addresses on each chain
  swapEthAddresses?: Record<number, string> // list of swap eth addresses on each chain
  routerIndex?: string // router index
  poolId: number | Record<number, number> = {} // list of pool ids on each chain
  poolType?: string // pool type
  poolTokens?: Token[] // list of pool tokens
  depositTokens?: Token[] // list of deposit tokens
  nativeTokens?: Token[] // list of native tokens
  description?: string // token description
  docUrl = '' // token doc url
  forceMeta?: boolean // force meta
  swapableType?: string // swapable type
  isNative = false // is native
  swapExceptions: number | Record<number, number[]> = {} // for specifying tokens where limited dest chains are available.
  visibilityRank = 0 // rank in which token is displayed, least visible is 0, there is no max
  isMeta = false // is meta
  isEthSwap = false // is eth swap
  category: { bridge: boolean; swap: boolean; pool: boolean } = {
    bridge: true,
    swap: true,
    pool: true,
  } // list of categories on each chain
  swapableOn: number[] = [] // list of chains where token is swapable
  display = true // display token
  legacy = false // legacy token
  priorityRank: number // priority token ordering
  chainId?: number // chain id of swap pool
  incentivized?: boolean // pool is incentivized or not
  customRewardToken?: string // reward token symbol when pool staking rewards are in something other than SYN
  miniChefAddress: string // mini chef address
  priorityPool?: boolean = false // priority pool
  color?:
    | 'gray'
    | 'yellow'
    | 'green'
    | 'lime'
    | 'sky'
    | 'blue'
    | 'orange'
    | 'purple'
    | 'indigo'
    | 'cyan'
    | 'red'
  priceUnits?: string
  notStake?: boolean
  routeSymbol?: string
  constructor({
    addresses,
    wrapperAddresses,
    decimals,
    symbol,
    name,
    logo,
    poolName,
    swapAddresses,
    swapWrapperAddresses,
    swapDepositAddresses,
    swapEthAddresses,
    routerIndex,
    poolId,
    poolType,
    poolTokens,
    depositTokens,
    nativeTokens,
    description,
    docUrl = '',
    forceMeta,
    swapableType,
    isNative = false,
    swapExceptions,
    visibilityRank,
    isMeta,
    isEthSwap,
    category,
    swapableOn,
    display,
    legacy,
    priorityRank,
    chainId,
    incentivized,
    customRewardToken,
    miniChefAddress,
    priorityPool,
    color,
    priceUnits,
    notStake,
    routeSymbol,
  }: {
    addresses: { [x: number]: string }
    wrapperAddresses?: Record<number, string>
    decimals?: number | Record<number, number>
    symbol?: string
    name?: string
    logo?: any
    poolName?: string
    swapAddresses?: Record<number, string>
    swapWrapperAddresses?: Record<number, string>
    swapDepositAddresses?: Record<number, string>
    swapEthAddresses?: Record<number, string>
    routerIndex?: string
    poolId?: number | Record<number, number>
    poolType?: string
    poolTokens?: Token[]
    depositTokens?: Token[]
    nativeTokens?: Token[]
    description?: string
    docUrl?: string
    forceMeta?: boolean
    swapableType?: string
    isNative?: boolean
    swapExceptions?: number | Record<number, number[]>
    visibilityRank?: number
    isMeta?: boolean
    isEthSwap?: boolean
    category?: { bridge: boolean; swap: boolean; pool: boolean }
    swapableOn?: number[]
    display?: boolean
    legacy?: boolean
    priorityRank: number
    chainId?: number
    incentivized?: boolean
    customRewardToken?: string
    miniChefAddress?: string
    priorityPool?: boolean
    color?:
      | 'gray'
      | 'yellow'
      | 'green'
      | 'lime'
      | 'sky'
      | 'blue'
      | 'orange'
      | 'purple'
      | 'indigo'
      | 'cyan'
      | 'red'
    priceUnits?: string
    notStake?: boolean
    routeSymbol?: string
  }) {
    const isMetaVar = Boolean(swapDepositAddresses || forceMeta)
    this.addresses = validateAddresses(addresses)
    this.wrapperAddresses = wrapperAddresses
    // this.decimals             = decimals
    this.decimals = makeMultiChainObj(decimals)
    this.symbol = symbol
    this.name = name
    this.icon = logo
    this.poolName = poolName
    this.swapAddresses = swapAddresses
    this.swapWrapperAddresses = swapWrapperAddresses
    this.swapDepositAddresses = swapDepositAddresses
    this.swapEthAddresses = swapEthAddresses
    this.routerIndex = routerIndex

    this.poolTokens = poolTokens
    this.nativeTokens = nativeTokens ?? poolTokens
    this.depositTokens = depositTokens ?? this.nativeTokens
    this.description = description
    this.docUrl = docUrl ?? ''

    this.poolId = makeMultiChainObj(poolId)
    this.poolType = poolType
    this.visibilityRank = visibilityRank ?? 0
    this.isMeta = isMeta
    this.isEthSwap = swapEthAddresses ? true : false
    this.isNative = isNative ?? false
    this.swapableType = swapableType
    this.swapExceptions = swapExceptions ?? []
    this.category = category ?? { bridge: true, swap: true, pool: true }
    this.swapableOn = swapableOn ?? []
    this.display = display ?? true
    this.legacy = legacy ?? false
    this.priorityRank = priorityRank
    this.chainId = chainId
    this.incentivized = incentivized
    this.customRewardToken = customRewardToken
    this.miniChefAddress = miniChefAddress
    this.priorityPool = priorityPool ?? false
    this.color = color ?? 'gray'
    this.priceUnits = priceUnits ?? 'USD'
    this.notStake = notStake ?? false
    this.routeSymbol = routeSymbol
  }
}

const makeMultiChainObj = (valOrObj) => {
  if (typeof valOrObj === 'object') {
    return valOrObj
  } else {
    const obj = {}
    for (const chain of Object.values(CHAINS)) {
      obj[chain.id] = valOrObj
    }
    return obj
  }
}

const validateAddresses = (addresses: {
  [x: number]: string
}): { [x: number]: string } => {
  const reformatted: { [x: number]: string } = {}
  for (const chainId in addresses) {
    reformatted[chainId] = addresses[chainId]
      ? getAddress(addresses[chainId])
      : ''
  }
  return reformatted
}
