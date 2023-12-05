export interface TokenMetaData {
  tokenAddress: string
  symbol: string
  chainId: number
  decimals: number
}

export interface CustomThemeVariables {
  bgColor?: string
  brandColor?: string
  '--synapse-text-primary'?: string
  '--synapse-text-secondary'?: string
  '--synapse-bg-select'?: string
  '--synapse-bg-surface'?: string
  '--synapse-bg-root'?: string
  '--synapse-border'?: string
  '--synapse-border-hover'?: string
  '--synapse-brand'?: string
}

export interface WidgetProps {
  /** Supported ChainIds to Bridge defined by Consumer */
  chainIds: number[]

  /** Consumer Web3 Provider */
  web3Provider?: any

  /** Respective Network Providers */
  networkProviders?: any[]

  /** Selected Day/Night Theme */
  theme?: 'day' | 'night'

  /** Apply Custom Themes */
  customTheme?: CustomThemeVariables

  /** Supported Tokens Metadata defined by Consumer */
  tokens: TokenMetaData[]
}
