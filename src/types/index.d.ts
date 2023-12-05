export interface TokenMetaData {
  tokenAddress: string
  symbol: string
  chainId: number
  decimals: number
}

export interface CustomTheme {
  primary?: string
  secondary?: string
  accent?: string
  small?: string
  separator?: string
  background?: string
  surface?: string
}

export interface CustomThemeVariables {
  '--h'?: string
  '--s'?: string
  '--primary'?: string
  '--secondary'?: string
  '--accent'?: string
  '--small'?: string
  '--separator'?: string
  '--background'?: string
  '--surface'?: string
  '--brand'?: string
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
  customTheme?: CustomTheme

  /** Supported Tokens Metadata defined by Consumer */
  tokens: TokenMetaData[]
}
