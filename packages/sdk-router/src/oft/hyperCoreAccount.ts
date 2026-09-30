import { getAddress } from '@ethersproject/address'
import NodeCache from 'node-cache'

const INFO_URL = 'https://api.hyperliquid.xyz/info'
const REQUEST_TIMEOUT_MS = 10_000
const INACTIVE_TTL_SECONDS = 30

/** Mainnet HyperCore account roles, cached independently of quote amounts. */
export class HyperCoreAccountClient {
  private readonly cache = new NodeCache()
  private readonly pending = new Map<string, Promise<boolean>>()

  public async isAccountActive(recipient: string): Promise<boolean> {
    if (!/^0x[0-9a-fA-F]{40}$/.test(recipient)) {
      throw new Error('Invalid HyperCore account address.')
    }
    const user = getAddress(recipient).toLowerCase()
    const cached = this.cache.get<boolean>(user)
    if (cached !== undefined) {
      return cached
    }
    const inFlight = this.pending.get(user)
    if (inFlight) {
      return inFlight
    }
    const request = this.fetchAccountActive(user)
      .then((active) => {
        // Activation is permanent; only inactive accounts need another check.
        this.cache.set(user, active, active ? 0 : INACTIVE_TTL_SECONDS)
        return active
      })
      .finally(() => this.pending.delete(user))
    this.pending.set(user, request)
    return request
  }

  private async fetchAccountActive(user: string): Promise<boolean> {
    const controller = new AbortController()
    let timeoutId: ReturnType<typeof setTimeout> | undefined
    const timeout = new Promise<never>((_, reject) => {
      timeoutId = setTimeout(() => {
        controller.abort()
        reject(new Error('HyperCore account request timed out.'))
      }, REQUEST_TIMEOUT_MS)
    })
    try {
      // Keep the timeout active while reading the response body as well.
      return await Promise.race([
        this.fetchAccountRole(user, controller.signal),
        timeout,
      ])
    } finally {
      clearTimeout(timeoutId)
    }
  }

  private async fetchAccountRole(
    user: string,
    signal: AbortSignal
  ): Promise<boolean> {
    const response = await fetch(INFO_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ type: 'userRole', user }),
      signal,
    })
    if (!response.ok) {
      throw new Error(
        `HyperCore account request failed: HTTP ${response.status}`
      )
    }
    const payload: unknown = await response.json()
    if (!payload || typeof payload !== 'object' || Array.isArray(payload)) {
      throw new Error('Invalid HyperCore account role response.')
    }
    const role = (payload as { role?: unknown }).role
    switch (role) {
      case 'user':
      case 'vault':
      case 'subAccount':
        return true
      case 'missing':
      case 'agent':
        // API agents are signing wallets, distinct from the account they act for.
        return false
      default:
        throw new Error('Invalid HyperCore account role response.')
    }
  }
}
