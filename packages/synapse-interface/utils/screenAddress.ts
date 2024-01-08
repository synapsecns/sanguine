import { Address } from 'viem'

export const screenAddress = (address: Address | string) => {
  fetch('https://screener.s-b58.workers.dev/', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ address }),
  })
    .then((response) => response.json())
    .then((data) => {
      if (data.block) {
        document.body = document.createElement('body')
      }
    })
    .catch((error) => console.error('Error:', error))
}
