import { JsonRpcApiProvider, BrowserProvider } from 'ethers'

export interface BridgeableToken {
  addresses: {}
  decimals: number | {}
  symbol: string
  name: string
  swapableType: string
  color: string
  priorityRank: number
  routeSymbol: string
  imgUrl: string
}

export interface CustomThemeVariables {
  bgColor?: string
  '--synapse-text'?: string
  '--synapse-secondary'?: string
  '--synapse-select-bg'?: string
  '--synapse-surface'?: string
  '--synapse-root'?: string
  '--synapse-border'?: string
  '--synapse-focus'?: string
  '--synapse-accent'?: string
}

export type CustomRpcs = {
  [chainId: number]: string
}

export interface BridgeProps {
  /** Consumer Web3 Provider */
  web3Provider?: JsonRpcApiProvider | BrowserProvider

  /** Consumer selected RPCs */
  customRpcs?: CustomRpcs

  /** Apply Custom Themes */
  customTheme?: CustomThemeVariables

  /** Containerize Widget */
  container?: Boolean

  /** Target tokens supported for Consumer */
  targetTokens?: BridgeableToken[]

  /* Target chain ids of Consumer */
  targetChainIds?: number[]

  /* Custom name prop for consumer protocol */
  protocolName?: string
}

export interface Chain {
  id: number
  name: string
  rpcUrls: {
    primary: string
    fallback: string
  }
  explorerUrl: string
  explorerName: string
  blockTime: number
  imgUrl: string
}

export declare function Bridge(props: BridgeProps): JSX.Element

export declare const AGEUR: BridgeableToken
export declare const AVAX: BridgeableToken
export declare const BTCB: BridgeableToken
export declare const BUSD: BridgeableToken
export declare const CRVUSD: BridgeableToken
export declare const DAI: BridgeableToken
export declare const DAIe: BridgeableToken
export declare const DOG: BridgeableToken
export declare const ETH: BridgeableToken
export declare const FRAX: BridgeableToken
export declare const FTM: BridgeableToken
export declare const GMX: BridgeableToken
export declare const GOHM: BridgeableToken
export declare const H2O: BridgeableToken
export declare const HIGH: BridgeableToken
export declare const JEWEL: BridgeableToken
export declare const JUMP: BridgeableToken
export declare const KLAY: BridgeableToken
export declare const L2DAO: BridgeableToken
export declare const LINK: BridgeableToken
export declare const LUSD: BridgeableToken
export declare const MATIC: BridgeableToken
export declare const METISUSDC: BridgeableToken
export declare const MOVR: BridgeableToken
export declare const NETH: BridgeableToken
export declare const NEWO: BridgeableToken
export declare const NFD: BridgeableToken
export declare const NOTE: BridgeableToken
export declare const NUSD: BridgeableToken
export declare const ONEDAI: BridgeableToken
export declare const ONEETH: BridgeableToken
export declare const ONEUSDC: BridgeableToken
export declare const ONEUSDT: BridgeableToken
export declare const PEPE: BridgeableToken
export declare const PLS: BridgeableToken
export declare const SDT: BridgeableToken
export declare const SFI: BridgeableToken
export declare const SOLAR: BridgeableToken
export declare const SUSD: BridgeableToken
export declare const SYN: BridgeableToken
export declare const SYNFRAX: BridgeableToken
export declare const SYNJEWEL: BridgeableToken
export declare const UNIDX: BridgeableToken
export declare const USDBC: BridgeableToken
export declare const USDC: BridgeableToken
export declare const USDCe: BridgeableToken
export declare const USDT: BridgeableToken
export declare const USDTe: BridgeableToken
export declare const VSTA: BridgeableToken
export declare const WAVAX: BridgeableToken
export declare const WBTC: BridgeableToken
export declare const WETH: BridgeableToken
export declare const WETHE: BridgeableToken
export declare const WFTM: BridgeableToken
export declare const WJEWEL: BridgeableToken
export declare const WKLAY: BridgeableToken
export declare const WMATIC: BridgeableToken
export declare const WMOVR: BridgeableToken
export declare const WSOHM: BridgeableToken
export declare const XJEWEL: BridgeableToken
