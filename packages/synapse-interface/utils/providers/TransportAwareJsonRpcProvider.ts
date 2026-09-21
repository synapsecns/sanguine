import { StaticJsonRpcProvider } from '@ethersproject/providers'

type RpcError = Error & {
  code?: string
  status?: number
  body?: string
  error?: RpcError
}

// ethers v5 converts an HTTP failure during eth_call into CALL_EXCEPTION.
// FallbackProvider forwards that error immediately at quorum 1, so the next
// RPC is never tried. Keep actual JSON-RPC errors on the normal ethers path.
export class TransportAwareJsonRpcProvider extends StaticJsonRpcProvider {
  async perform(method: string, params: any): Promise<any> {
    try {
      return await super.perform(method, params)
    } catch (error) {
      const rpcError = error as RpcError
      const transportError = rpcError.error
      if (
        method === 'call' &&
        rpcError.code === 'CALL_EXCEPTION' &&
        transportError?.code === 'SERVER_ERROR' &&
        typeof transportError.status === 'number' &&
        !hasJsonRpcErrorResponse(transportError.body)
      ) {
        throw transportError
      }
      throw error
    }
  }
}

const hasJsonRpcErrorResponse = (body?: string): boolean => {
  if (!body) return false
  try {
    const response = JSON.parse(body)
    return (
      response !== null &&
      typeof response === 'object' &&
      response.error !== null &&
      typeof response.error === 'object' &&
      typeof response.error.code === 'number'
    )
  } catch {
    return false
  }
}
