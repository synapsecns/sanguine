import { getWithTimeout } from '../utils/api'

const GAS_ZIP_API_URL = 'https://backend.gas.zip/v2'
const GAS_ZIP_API_TIMEOUT = 2000

interface Transaction {
  status: string
}

interface TransactionStatusData {
  txs: Transaction[]
}

interface Chains {
  chains: [
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
  return data.chains.map((chain) => chain.chain)
}
