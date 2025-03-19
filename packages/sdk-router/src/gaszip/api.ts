import { Zero } from '@ethersproject/constants'
import { BigNumber } from 'ethers'

import { BigintIsh } from '../constants'
import { getWithTimeout } from '../utils'

const GAS_ZIP_API_URL = 'https://backend.gas.zip/v2'
const GAS_ZIP_API_TIMEOUT = 2000

interface Transaction {
  status: string
}

interface TransactionStatusData {
  txs: Transaction[]
}

interface Chains {
  chains?: [
    {
      name: string
      chain: number // native chain id
      short: number // unique Gas.zip id
      gas: string // gas usage of a simple transfer
      gwei: string // current gas price
      bal: string // balance of the Gas.zip reloader
      rpcs: string[]
      symbol: string
      price: number
    }
  ]
}

interface CalldataQuoteResponse {
  calldata: string
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
  calldata: string
}

const EMPTY_GAS_ZIP_QUOTE: GasZipQuote = {
  amountOut: Zero,
  calldata: '0x',
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

export const getChainIds = async (): Promise<number[]> => {
  const response = await getWithTimeout(
    'Gas.Zip API',
    `${GAS_ZIP_API_URL}/chains`,
    GAS_ZIP_API_TIMEOUT
  )
  if (!response) {
    return []
  }
  const data: Chains = await response.json()
  return data.chains?.map((chain) => chain.chain) ?? []
}

export const getGasZipQuote = async (
  originChainId: number,
  destChainId: number,
  amount: BigintIsh,
  to: string,
  from?: string
): Promise<GasZipQuote> => {
  const response = await getWithTimeout(
    'Gas.Zip API',
    `${GAS_ZIP_API_URL}/quotes/${originChainId}/${amount}/${destChainId}`,
    GAS_ZIP_API_TIMEOUT,
    {
      from,
      to,
    }
  )
  if (!response) {
    return EMPTY_GAS_ZIP_QUOTE
  }
  const data: CalldataQuoteResponse = await response.json()
  if (data.quotes.length === 0 || !data.quotes[0].expected) {
    return EMPTY_GAS_ZIP_QUOTE
  }
  return {
    amountOut: BigNumber.from(data.quotes[0].expected.toString()),
    calldata: data.calldata,
  }
}
