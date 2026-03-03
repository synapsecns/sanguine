import NodeCache from 'node-cache'

import { CIRCLE_IRIS_API_HOST } from '../constants'
import { getWithTimeout } from '../utils'

const CIRCLE_API_TIMEOUT = 5000
const FEE_CACHE_TTL_SECONDS = 15
const STATUS_CACHE_TTL_SECONDS = 5

const circleApiCache = new NodeCache({
  useClones: false,
})

export type CctpV2Fee = {
  finalityThreshold: number
  minimumFee: number
  forwardFee?: Record<string, number>
}

export type CctpV2Message = {
  status?: string
  forwardState?: string
  decodedMessage?: {
    decodedMessageBody?: {
      hookData?: string
    }
  }
}

const isObject = (value: unknown): value is Record<string, unknown> => {
  return value !== null && typeof value === 'object'
}

const toNonNegativeInteger = (value: unknown): number | undefined => {
  return Number.isInteger(value) && (value as number) >= 0
    ? (value as number)
    : undefined
}

const parseForwardFee = (
  value: unknown
): Record<string, number> | undefined => {
  if (!isObject(value)) {
    return undefined
  }
  const parsedEntries = Object.entries(value).filter(
    ([, fee]) => Number.isInteger(fee) && (fee as number) >= 0
  ) as [string, number][]
  if (!parsedEntries.length) {
    return undefined
  }
  return Object.fromEntries(parsedEntries)
}

const parseFeeResponse = (data: unknown): CctpV2Fee[] | null => {
  if (!Array.isArray(data)) {
    return null
  }
  const fees = data
    .map((item): CctpV2Fee | null => {
      if (!isObject(item)) {
        return null
      }
      const finalityThreshold = toNonNegativeInteger(item.finalityThreshold)
      const minimumFee = toNonNegativeInteger(item.minimumFee)
      if (finalityThreshold === undefined || minimumFee === undefined) {
        return null
      }
      if (!Object.prototype.hasOwnProperty.call(item, 'forwardFee')) {
        return {
          finalityThreshold,
          minimumFee,
        }
      }
      const forwardFee = parseForwardFee(item.forwardFee)
      if (!forwardFee) {
        return null
      }
      return {
        finalityThreshold,
        minimumFee,
        forwardFee,
      }
    })
    .filter((entry): entry is CctpV2Fee => !!entry)
  return fees.length ? fees : null
}

const parseMessagesResponse = (data: unknown): CctpV2Message[] | null => {
  if (!isObject(data) || !Array.isArray(data.messages)) {
    return null
  }
  const messages = data.messages.filter((message): message is CctpV2Message =>
    isObject(message)
  )
  return messages
}

export const getBurnUSDCFees = async (
  sourceDomainId: number,
  destDomainId: number
): Promise<CctpV2Fee[] | null> => {
  const cacheKey = `cctp-v2-fee-${sourceDomainId}-${destDomainId}`
  const cached = circleApiCache.get<CctpV2Fee[] | null>(cacheKey)
  if (cached !== undefined) {
    return cached
  }
  const response = await getWithTimeout(
    'Circle CCTP API',
    `${CIRCLE_IRIS_API_HOST}/v2/burn/USDC/fees/${sourceDomainId}/${destDomainId}`,
    CIRCLE_API_TIMEOUT,
    { forward: true }
  )
  if (!response) {
    circleApiCache.set(cacheKey, null, FEE_CACHE_TTL_SECONDS)
    return null
  }
  try {
    const parsed = parseFeeResponse(await response.json())
    circleApiCache.set(cacheKey, parsed, FEE_CACHE_TTL_SECONDS)
    return parsed
  } catch {
    circleApiCache.set(cacheKey, null, FEE_CACHE_TTL_SECONDS)
    return null
  }
}

export const getMessages = async (
  sourceDomainId: number,
  transactionHash: string
): Promise<CctpV2Message[] | null> => {
  const cacheKey = `cctp-v2-status-${sourceDomainId}-${transactionHash.toLowerCase()}`
  const cached = circleApiCache.get<CctpV2Message[] | null>(cacheKey)
  if (cached !== undefined) {
    return cached
  }
  const response = await getWithTimeout(
    'Circle CCTP API',
    `${CIRCLE_IRIS_API_HOST}/v2/messages/${sourceDomainId}`,
    CIRCLE_API_TIMEOUT,
    { transactionHash }
  )
  if (!response) {
    circleApiCache.set(cacheKey, null, STATUS_CACHE_TTL_SECONDS)
    return null
  }
  try {
    const parsed = parseMessagesResponse(await response.json())
    circleApiCache.set(cacheKey, parsed, STATUS_CACHE_TTL_SECONDS)
    return parsed
  } catch {
    circleApiCache.set(cacheKey, null, STATUS_CACHE_TTL_SECONDS)
    return null
  }
}
