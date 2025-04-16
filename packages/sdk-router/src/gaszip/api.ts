import { Zero } from '@ethersproject/constants'
import { BigNumber, BigNumberish } from 'ethers'

import { GASZIP_SUPPORTED_CHAIN_IDS } from '../constants'
import { getWithTimeout } from '../utils'

const GAS_ZIP_API_URL = 'https://backend.gas.zip/v2'
const GAS_ZIP_API_TIMEOUT = 2000

interface Transaction {
  status: string
}

interface TransactionStatusData {
  txs: Transaction[]
}

export interface Chains {
  chains?: {
    name: string
    chain: number // native chain id
    short: number // unique Gas.zip id
    gas: string // gas usage of a simple transfer
    gwei: string // current gas price
    bal: string // balance of the Gas.zip reloader
    rpcs: string[]
    symbol: string
    price: number
  }[]
}

interface QuoteResponse {
  quotes: {
    chain: number
    expected: string
    gas: string
    speed: number
    usd: number
  }[]
}

export type GasZipQuote = {
  amountOut: BigNumber
  speed: number
  usd: number
}

const QUOTE_MIN_USD = 0.1
const QUOTE_MAX_USD = 200

const EMPTY_GAS_ZIP_QUOTE: GasZipQuote = {
  amountOut: Zero,
  speed: 0,
  usd: 0,
}

export const getGasZipTxStatus = async (txHash: string): Promise<boolean> => {
  const response = await getWithTimeout(
    'Gas.Zip API',
    `${GAS_ZIP_API_URL}/search/${txHash}`,
    GAS_ZIP_API_TIMEOUT
  )
  if (!response) {
    return false
  }
  const data: TransactionStatusData = await response.json()
  return data.txs.length > 0 && data.txs[0].status === 'CONFIRMED'
}

export const getChains = async (): Promise<Chains> => {
  const response = await getWithTimeout(
    'Gas.Zip API',
    `${GAS_ZIP_API_URL}/chains`,
    GAS_ZIP_API_TIMEOUT
  )
  if (!response) {
    return {}
  }
  const data: Chains = await response.json()
  // Filter out unsupported chains
  return {
    chains: data.chains?.filter((chain) =>
      GASZIP_SUPPORTED_CHAIN_IDS.includes(chain.chain)
    ),
  }
}

export const getGasZipQuote = async (
  originChainId: number,
  destChainId: number,
  amount: BigNumberish
): Promise<GasZipQuote> => {
  const response = await getWithTimeout(
    'Gas.Zip API',
    `${GAS_ZIP_API_URL}/quotes/${originChainId}/${amount}/${destChainId}`,
    GAS_ZIP_API_TIMEOUT
  )
  if (!response) {
    return EMPTY_GAS_ZIP_QUOTE
  }
  const data: QuoteResponse = await response.json()
  if (data.quotes.length === 0 || !data.quotes[0].expected) {
    return EMPTY_GAS_ZIP_QUOTE
  }
  const quote = data.quotes[0]
  if (quote.usd < QUOTE_MIN_USD || quote.usd > QUOTE_MAX_USD) {
    return EMPTY_GAS_ZIP_QUOTE
  }
  return {
    amountOut: BigNumber.from(quote.expected.toString()),
    speed: quote.speed,
    usd: quote.usd,
  }
}
