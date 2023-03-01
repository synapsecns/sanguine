export type Price = number
export type Size = number

export type Bid = Array<[Price, Size]>
export type Ask = Array<[Price, Size]>

export interface Trade {
  id: number
  price: number
  side: string
  size: number
  time: number
}

export interface Order {
  createdAt: number
  sizeFilled: number
  id: number
  symbol: string
  price: number
  avgFillPrice: number
  sizeRemaining: number
  side: string
  size: number
  status: string
  type: string
  tif: string
  postOnly: boolean
  clientId: string | null
}

export interface Market {
  symbol: string
  marketType: string
  tickSize: number
  lotSize: number
  maxPrice: number
  minPrice: number
  minSize: number
  maxSize: number
  quoteVolume24h: number
  baseVolume24h: number
  bid: number
  ask: number
  last: number
  change1h: number
  change24h: number
  enabled: boolean
  postOnly: boolean
}

export interface SpotData {
  symbol: string
  change24h: number
  quoteVolume24h: number
  last: number
}

export interface Balance {
  asset: string
  total: number
  available: number
  totalUSDValue: number
}

// TODO: Better name for this interface?
export interface Asset {
  text: string
  address: string
  logo: string
  amount: string
  tickSize: string
}

export interface Chain {
  chainId: number
  chainSymbol: string
  chainName: string
  chainLogo: string
  chainImg: string
  layer: number
}

export interface ThemeContextInterface {
  theme: string | null
  toggleFunction: () => void
}

export interface Account {
  makerFee: number
  takerFee: number
  address: string
  balances: Balance[]
}

export interface OrderSubmission {
  symbol: string
  side: string
  size: number
  type: string
  price?: number
  postOnly: boolean
  tif: string
  clientId: string
  sizeSelected: boolean
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

export interface Palette {
  light: number
  warmth: number
  lang: string
}
