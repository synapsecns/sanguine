import { getRequestContext } from '@cloudflare/next-on-pages'

export const config = {
  runtime: 'edge',
}

const ALLOWED_DOMAINS = [
  'synapseprotocol.com',
  'cortexprotocol.com',
  'hypercall.xyz',
]

function isDomainAllowed(headerValue: string | null): boolean {
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

export default async function handler(req: Request) {
  if (req.method !== 'POST') {
    return new Response('Method not allowed', { status: 405 })
  }

  const origin = req.headers.get('origin')
  const referer = req.headers.get('referer')
  const host = req.headers.get('host')
  const bypassKey = req.headers.get('x-admin-bypass')

  const { env } = getRequestContext()

  const adminBypass =
    env.ADMIN_RPC_BYPASS && bypassKey === env.ADMIN_RPC_BYPASS
  const isSameOrigin = origin === `https://${host}`

  if (
    !adminBypass &&
    !isSameOrigin &&
    !isDomainAllowed(origin) &&
    !isDomainAllowed(referer)
  ) {
    return new Response('Forbidden', { status: 403 })
  }

  const secret = env.GOLDSKY_RPC_SECRET as string
  if (!secret) {
    return new Response('RPC proxy not configured', { status: 500 })
  }

  const chainId = new URL(req.url).pathname.split('/').pop()

  const body = await req.text()

  const resp = await fetch(
    `https://edge.goldsky.com/standard/evm/${chainId}`,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'X-ERPC-Secret-Token': secret,
      },
      body,
    }
  )

  return new Response(resp.body, {
    status: resp.status,
    headers: { 'Content-Type': 'application/json' },
  })
}
