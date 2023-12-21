import { JsonRpcApiProvider, BrowserProvider, Provider } from 'ethers'
// import { Provider } from '@ethersproject/abstract-provider';

export interface BridgeableToken {
  addresses: {}
  decimals: number | {}
  symbol: string
  name: string
  swapableType: string
  color: string
  priorityRank: number
  routeSymbol: string
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

export interface WidgetProps {
  /** Supported ChainIds to Bridge defined by Consumer */
  chainIds: number[]

  /** Consumer Web3 Provider */
  web3Provider?: JsonRpcApiProvider | BrowserProvider

  /** Respective Network Providers */
  // TO DO: Pass proper ethers type
  networkProviders?: any[]

  /** Selected Day/Night Theme */
  theme?: 'light' | 'dark'

  /** Apply Custom Themes */
  customTheme?: CustomThemeVariables
  container?: Boolean

  /** Supported Tokens Metadata defined by Consumer */
  tokens: BridgeableToken[]

  /** Destination chain selected by Consumer */
  toChainId: number
}

export interface Chain {
  id: number
  name: string
}

export declare function Bridge(props: any): JSX.Element

export declare const USDC: BridgeableToken
export declare const USDT: BridgeableToken
export declare const DAI: BridgeableToken
export declare const ETH: BridgeableToken
