import { Address } from 'viem'

import { GlobalEventEmitter } from '@/utils/globalEventEmitter'
import { DISCORD_URL } from '@/constants/urls'

export const screenAddress = (address: Address | string) => {
  const url = `https://screener.omnirpc.io/fe/address/${address}`

  fetch(url, {
    method: 'GET',
  })
    .then((response) => response.json())
    .then(({ risk }) => {
      if (risk) {
        const event = new CustomEvent('riskDetected', {
          detail: {
            message: `This address has been flagged for being associated with illicit activities. If you think this is a mistake, please contact <a style="text-decoration: underline; text-underline-offset: 0.2em;" href=${DISCORD_URL} target="_blank" rel="noopener noreferrer">support</a>.`,
          },
        })

        GlobalEventEmitter.dispatchEvent(event)
      }
    })
    .catch((error) => console.error('Error:', error))
}
