import { AddressZero } from '@ethersproject/constants'

import { getWithTimeout, isNativeToken, postWithTimeout } from '../utils'

const API_URL = 'https://api.relay.link'
// TODO: revisit timeout
const API_TIMEOUT = 5000

export enum TradeType {
  ExactInput = 'EXACT_INPUT',
  ExactOutput = 'EXACT_OUTPUT',
  ExpectedOutput = 'EXPECTED_OUTPUT',
}

export interface QuoteRequest {
  user: string
  originChainId: number
  destinationChainId: number
  originCurrency: string
  destinationCurrency: string
  amount: string
  tradeType: TradeType
  recipient: string
  refundTo: string
  refundOnOrigin: boolean
  useReceiver: boolean
  explicitDeposit: boolean
  usePermit: boolean
  slippageTolerance?: string
}

enum StepId {
  Deposit = 'deposit',
  Approve = 'approve',
  Authorize = 'authorize',
  Authorize1 = 'authorize1',
  Authorize2 = 'authorize2',
  Swap = 'swap',
  Send = 'send',
}

enum StepKind {
  Transaction = 'transaction',
  Signature = 'signature',
}

export interface StepData {
  from: string
  to: string
  data: string
  value: string
  chainId: number
}

export interface QuoteStep {
  id: StepId
  kind: StepKind
  items: {
    data: StepData
  }[]
}

interface Currency {
  currency: {
    chainId: number
    address: string
    symbol: string
    name: string
    decimals: number
  }
  amount: string
  amountFormatted: string
  amountUsd: string
  minimumAmount: string
}

export interface QuoteResponse {
  steps: QuoteStep[]
  requestId: string
  details: {
    currencyIn: Currency
    currencyOut: Currency
    timeEstimate: number
  }
}

export enum Status {
  Refund = 'refund',
  Delayed = 'delayed',
  Waiting = 'waiting',
  Failure = 'failure',
  Pending = 'pending',
  Success = 'success',
}

interface RequestsRequest {
  hash: string
}

export interface RequestsResponse {
  requests: {
    id: string
    status: Status
    data: {
      inTxs: {
        hash: string
      }[]
    }
  }[]
}

export const isStepActionable = (step: QuoteStep): boolean => {
  // Note: Approval is done by TokenZap automatically, so this step doesn't require an action
  return [StepId.Deposit, StepId.Send, StepId.Swap].includes(step.id)
}

export const addressToCurrency = (address: string): string => {
  return isNativeToken(address) ? AddressZero : address
}

export const getQuote = async (
  quoteRequest: QuoteRequest
): Promise<QuoteResponse | null> => {
  const response = await postWithTimeout(
    'Relay API',
    `${API_URL}/quote`,
    API_TIMEOUT,
    quoteRequest
  )
  if (!response) {
    return null
  }
  const data: QuoteResponse = await response.json()
  return data
}

export const getRequests = async (
  requestsRequest: RequestsRequest
): Promise<RequestsResponse | null> => {
  const response = await getWithTimeout(
    'Relay API',
    `${API_URL}/requests/v2`,
    API_TIMEOUT,
    requestsRequest
  )
  if (!response) {
    return null
  }
  const data: RequestsResponse = await response.json()
  return data
}
