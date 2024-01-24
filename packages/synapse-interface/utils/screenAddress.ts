import { Address } from 'viem'

export const screenAddress = (address: Address | string) => {
  const url = `https://screener.omnirpc.io/fe/address/${address}`

  fetch(url, {
    method: 'GET',
  })
    .then((response) => response.json())
    .then(({ risk }) => {
      if (risk) {
        document.body = document.createElement('body')
      }
    })
    .catch((error) => console.error('Error:', error))
}
