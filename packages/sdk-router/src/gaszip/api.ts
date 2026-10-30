import { Zero } from '@ethersproject/constants'
import { BigNumber, BigNumberish } from 'ethers'

import { GASZIP_SUPPORTED_CHAIN_IDS } from '../constants'
import { getWithTimeout, logger } from '../utils'

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
    speed?: number
    usd: number
  }[]
}

export type GasZipQuote = {
  amountOut: BigNumber
  speed: number
  usd: number
}

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

export const getGasZipBlockHeightMap = async (): Promise<
  Map<number, number>
> => {
  const response = await getWithTimeout(
    'Gas.Zip API',
    `${GAS_ZIP_API_URL}/admin/indexer`,
    GAS_ZIP_API_TIMEOUT
  )
  if (!response) {
    return new Map()
  }
  const data: { chain: any; head: any }[] = await response.json()
  return new Map(
    data
      .filter((item) => {
        const chainNum = Number(item.chain)
        const headNum = Number(item.head)
        return Number.isInteger(chainNum) && Number.isInteger(headNum)
      })
      .map((item) => [Number(item.chain), Number(item.head)])
  )
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
  try {
    // Get response as text to handle large numbers properly
    const text = await response.text()
    // Wrap large numbers (> MAX_SAFE_INTEGER) in quotes to preserve precision
    const processedText = text.replace(/(:\s*)(\d{16,})(?=\s*[,}])/g, '$1"$2"')
    const data: QuoteResponse | null = JSON.parse(processedText)
    const quotes = data?.quotes
    // Gas.zip can return HTTP 200 with an error object when quoting is disabled.
    if (!Array.isArray(quotes) || !quotes[0]?.expected) {
      return EMPTY_GAS_ZIP_QUOTE
    }
    const quote = quotes[0]
    // Zero lets quote finalization use the module's default estimated time.
    const speed = quote.speed === undefined ? 0 : quote.speed
    if (!Number.isFinite(speed) || speed < 0) {
      return EMPTY_GAS_ZIP_QUOTE
    }
    return {
      amountOut: BigNumber.from(quote.expected),
      speed,
      usd: quote.usd,
    }
  } catch (error) {
    logger.warn('Gas.Zip API: invalid quote response', error)
    return EMPTY_GAS_ZIP_QUOTE
  }
}
