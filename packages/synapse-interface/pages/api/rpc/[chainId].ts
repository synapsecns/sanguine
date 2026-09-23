// @ts-ignore - types not shipped, but available at runtime on CF Pages
import { getRequestContext } from '@cloudflare/next-on-pages'

export const config = {
  runtime: 'edge',
}

const ALLOWED_DOMAINS = [
  'synapseprotocol.com',
  'cortexprotocol.com',
  'hypercall.xyz',
]

const isDomainAllowed = (headerValue: string | null): boolean => {
  if (!headerValue) return false
  try {
    const { hostname } = new URL(headerValue)
    return ALLOWED_DOMAINS.some(
      (domain) => hostname === domain || hostname.endsWith(`.${domain}`)
    )
  } catch {
    return false
  }
}

const handler = async (req: Request) => {
  if (req.method !== 'POST') {
    return new Response('Method not allowed', { status: 405 })
  }

  const origin = req.headers.get('origin')
  const referer = req.headers.get('referer')
  const host = req.headers.get('host')
  const bypassKey = req.headers.get('x-admin-bypass')

  const { env } = getRequestContext()

  const adminBypass = env.ADMIN_RPC_BYPASS && bypassKey === env.ADMIN_RPC_BYPASS
  const isSameOrigin = origin === `https://${host}`

  if (
    !adminBypass &&
    !isSameOrigin &&
    !isDomainAllowed(origin) &&
    !isDomainAllowed(referer)
  ) {
    return new Response('Forbidden', { status: 403 })
  }

  const requestUrl = new URL(req.url)
  const safeChainId = requestUrl.pathname.split('/').pop()
  if (!safeChainId || !/^\d+$/.test(safeChainId)) {
    return new Response('Invalid chainId', { status: 400 })
  }

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  }
  let upstreamUrl: string

  if (safeChainId === '999') {
    // Use QuickNode's /evm endpoint for live HyperCore precompile reads.
    // The full URL includes its auth token and must remain server-side.
    upstreamUrl = env.QUICKNODE_HYPEREVM_RPC_URL as string
    if (!upstreamUrl) {
      return new Response('RPC proxy not configured', { status: 500 })
    }
  } else {
    const secret = env.GOLDSKY_RPC_SECRET as string
    if (!secret) {
      return new Response('RPC proxy not configured', { status: 500 })
    }
    upstreamUrl = `https://edge.goldsky.com/standard/evm/${safeChainId}`
    headers['X-ERPC-Secret-Token'] = secret
  }

  const body = await req.text()
  const resp = await fetch(upstreamUrl, {
    method: 'POST',
    headers,
    body,
  })

  return new Response(resp.body, {
    status: resp.status,
    headers: { 'Content-Type': 'application/json' },
  })
}

export default handler
