import { NETWORK_CONNECTOR_MAP } from '@connectors'

/**
 * @param {number} chainId
 */
export function getProvider(chainId) {
  return NETWORK_CONNECTOR_MAP[chainId].providers[chainId]
}
