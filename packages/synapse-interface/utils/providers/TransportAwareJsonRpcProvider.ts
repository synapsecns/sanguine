import { StaticJsonRpcProvider } from '@ethersproject/providers'

type RpcError = {
  code?: string | number
  message?: string
  reason?: string
  status?: number
  body?: unknown
  error?: unknown
  data?: unknown
  originalError?: unknown
}

const asError = (value: unknown): RpcError | undefined =>
  value !== null && typeof value === 'object' && !Array.isArray(value)
    ? (value as RpcError)
    : undefined

const getResponseError = (body: unknown): RpcError | undefined => {
  if (typeof body !== 'string') return undefined
  try {
    return asError(asError(JSON.parse(body))?.error)
  } catch {
    return undefined
  }
}

const isExecutionError = (error: RpcError, depth = 0): boolean => {
  if (
    error.code === 3 ||
    (typeof error.message === 'string' &&
      /^(?:VM Exception while processing transaction: )?(?:(?:execution )?revert(?:ed)?|VM execution error|invalid opcode|out of gas)(?:$|[:.\s])/i.test(
        error.message
      ))
  ) {
    return true
  }
  if (depth >= 4) return false

  // Some RPCs wrap a data-less execution failure in an internal-error envelope.
  // Only inspect known error paths; never search request payloads or arbitrary
  // strings, and bound traversal for malformed/circular error objects.
  return [error.error, error.data, error.originalError].some((value) => {
    const nested = asError(value)
    return nested !== undefined && isExecutionError(nested, depth + 1)
  })
}

const isRpcError = (error: RpcError): boolean =>
  typeof error.code === 'number' &&
  Number.isInteger(error.code) &&
  typeof error.message === 'string'

// These codes describe endpoint capabilities/availability, not EVM outcomes.
const RETRYABLE_RPC_CODES = new Set([-32601, -32603, -32002, -32005])
const RETRYABLE_HTTP_STATUSES = new Set([401, 403, 404, 408, 429])

const isEndpointFailure = (error: RpcError): boolean => {
  if (error.code === 'TIMEOUT') return true
  if (error.code !== 'SERVER_ERROR') return false

  // ethers puts HTTP-200 RPC errors in error.error; non-2xx bodies remain text.
  // A protocol marker alone cannot distinguish service errors from reverts.
  const errors = [asError(error.error), getResponseError(error.body)].filter(
    (item): item is RpcError => item !== undefined
  )
  if (errors.some((item) => isExecutionError(item))) return false

  // Preserve malformed requests/parameters rather than masking caller bugs.
  if (errors.some(({ code }) => code === -32700 || code === -32602))
    return false

  const status = error.status
  if (
    typeof status === 'number' &&
    (RETRYABLE_HTTP_STATUSES.has(status) || (status >= 500 && status <= 599))
  ) {
    return true
  }
  if (
    errors.some(
      ({ code }) => typeof code === 'number' && RETRYABLE_RPC_CODES.has(code)
    )
  ) {
    return true
  }

  // Without an explicit endpoint failure, preserve unknown RPC errors
  // (including -32000 and invalid-request -32600). Hex data alone is not a revert.
  if (errors.some(isRpcError)) return false

  // Known ethers transport failures can lack both an HTTP status and a body:
  // disconnected sockets, exhausted HTTP-429 retries, and malformed responses.
  return (
    (typeof status === 'number' && status >= 300 && status <= 599) ||
    error.reason === 'missing response' ||
    error.reason === 'failed response' ||
    error.reason === 'processing response error'
  )
}

// ethers v5 wraps rejected eth_call requests in CALL_EXCEPTION, which is
// terminal at quorum one. Restore endpoint errors so FallbackProvider can try
// another RPC, while leaving execution results and all other methods alone.
export class TransportAwareJsonRpcProvider extends StaticJsonRpcProvider {
  async perform(method: string, params: any): Promise<any> {
    try {
      return await super.perform(method, params)
    } catch (error) {
      const rpcError = asError(error)
      const cause = asError(rpcError?.error)
      if (
        method === 'call' &&
        rpcError?.code === 'CALL_EXCEPTION' &&
        cause &&
        isEndpointFailure(cause)
      ) {
        throw cause
      }
      throw error
    }
  }
}
