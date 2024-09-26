import { Address } from 'viem'

import { DISCORD_URL } from '@/constants/urls'
import { GlobalEventEmitter } from '@/utils/globalEventEmitter'
import { isBlacklisted } from '@/utils/isBlacklisted'

const createRiskDetectedEvent = (address: Address | string) => {
  return new CustomEvent('riskDetected', {
    detail: {
      message: `This address ${address} has been flagged for violating the terms of service. Find out more at https://explorer.synapseprotocol.com/terms. If you think this is a mistake, please contact <a style="text-decoration: underline; text-underline-offset: 0.2em;" href=${DISCORD_URL} target="_blank" rel="noopener noreferrer">support</a>.`,
    },
  })
}

const dispatchRiskDetectedEvent = (address: Address | string) => {
  const event = createRiskDetectedEvent(address)
  GlobalEventEmitter.dispatchEvent(event)
}

export const screenAddress = async (
  address: Address | string
): Promise<boolean> => {
  const url = `https://screener.omnirpc.io/fe/address/${address}`

  if (isBlacklisted(address)) {
    dispatchRiskDetectedEvent(address)
    return true
  }

  try {
    const response = await fetch(url, { method: 'GET' })
    const { risk } = await response.json()

    if (risk) {
      dispatchRiskDetectedEvent(address)
      return true
    } else {
      return false
    }
  } catch (error) {
    dispatchRiskDetectedEvent(address)
    return true
  }
}
