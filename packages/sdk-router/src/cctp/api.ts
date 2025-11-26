import { getWithTimeout, postWithTimeout } from '../utils'

const WH_API_URL = 'https://executor.labsapis.com'
const WH_API_TIMEOUT = 3000

const CCTP_API_URL = 'https://iris-api.circle.com'
const CCTP_API_TIMEOUT = 3000

export interface ExecutorQuoteRequest {
  srcChain: number
  dstChain: number
  relayInstructions: string
}

export interface ExecutorQuoteResponse {
  signedQuote: string
  estimatedCost: string
}

export interface ExecutorTxStatusRequest {
  txHash: string
  chainId: number
}

export type ExecutorTxStatusResponse = Array<{
  status: string
  txHash: string
  failureCause?: string
}>

export interface CircleFeesRequest {
  sourceDomainId: number
  destDomainId: number
}

export type CircleFeesResponse = Array<{
  finalityThreshold: number
  minimumFee: number
}>

export interface CircleFastAllowanceResponse {
  allowance: number
  lastUpdated: string
}

export const getExecutorQuote = async (
  request: ExecutorQuoteRequest
): Promise<ExecutorQuoteResponse | null> => {
  const response = await postWithTimeout(
    'Wormhole API',
    `${WH_API_URL}/v0/quote`,
    WH_API_TIMEOUT,
    request
  )
  if (!response) {
    return null
  }
  const data: ExecutorQuoteResponse = await response.json()
  return data
}

export const getExecutorTxStatus = async (
  request: ExecutorTxStatusRequest
): Promise<ExecutorTxStatusResponse | null> => {
  const response = await postWithTimeout(
    'Wormhole API',
    `${WH_API_URL}/v0/status/tx`,
    WH_API_TIMEOUT,
    request
  )
  if (!response) {
    return null
  }
  const data: ExecutorTxStatusResponse = await response.json()

  return data
}

export const getCircleFees = async (
  request: CircleFeesRequest
): Promise<CircleFeesResponse | null> => {
  const response = await getWithTimeout(
    'CCTP API',
    `${CCTP_API_URL}/v2/burn/USDC/fees/${request.sourceDomainId}/${request.destDomainId}`,
    CCTP_API_TIMEOUT
  )
  if (!response) {
    return null
  }
  const data: CircleFeesResponse = await response.json()
  return data
}

export const getCircleFastAllowance =
  async (): Promise<CircleFastAllowanceResponse | null> => {
    const response = await getWithTimeout(
      'CCTP API',
      `${CCTP_API_URL}/v2/fastBurn/USDC/allowance`,
      CCTP_API_TIMEOUT
    )
    if (!response) {
      return null
    }
    const data: CircleFastAllowanceResponse = await response.json()
    return data
  }
