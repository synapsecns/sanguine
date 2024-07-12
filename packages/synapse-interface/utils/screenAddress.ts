import { Address } from 'viem'

import { GlobalEventEmitter } from '@/utils/globalEventEmitter'
import { DISCORD_URL } from '@/constants/urls'

export const screenAddress = async (
  address: Address | string
): Promise<boolean> => {
  const url = `https://screener.omnirpc.io/fe/address/${address}`
  try {
    const response = await fetch(url, { method: 'GET' })
    const { risk } = await response.json()
    if (risk) {
      const event = new CustomEvent('riskDetected', {
        detail: {
          message: `This address ${address} has been flagged for being associated with illicit activities. If you think this is a mistake, please contact <a style="text-decoration: underline; text-underline-offset: 0.2em;" href=${DISCORD_URL} target="_blank" rel="noopener noreferrer">support</a>.`,
        },
      })
      GlobalEventEmitter.dispatchEvent(event)
      return true
    }
    return false
  } catch (error) {
    console.error('Error:', error)
    return false
  }
}
