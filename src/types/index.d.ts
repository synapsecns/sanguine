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
  accentColor?: string
  '--synapse-text-primary'?: string
  '--synapse-text-secondary'?: string
  '--synapse-bg-select'?: string
  '--synapse-bg-surface'?: string
  '--synapse-bg-root'?: string
  '--synapse-border'?: string
  '--synapse-border-hover'?: string
  '--synapse-accent'?: string
}

export interface WidgetProps {
  /** Supported ChainIds to Bridge defined by Consumer */
  chainIds: number[]

  /** Consumer Web3 Provider */
  web3Provider?: any

  /** Respective Network Providers */
  networkProviders?: any[]

  /** Selected Day/Night Theme */
  theme?: 'light' | 'dark'

  /** Apply Custom Themes */
  customTheme?: CustomThemeVariables

  /** Supported Tokens Metadata defined by Consumer */
  tokens: BridgeableToken[]
}

export interface Chain {
  id: number
  name: string
}

export declare function Bridge(props: WidgetProps): JSX.Element

export declare const USDC: BridgeableToken
export declare const USDT: BridgeableToken
export declare const DAI: BridgeableToken
