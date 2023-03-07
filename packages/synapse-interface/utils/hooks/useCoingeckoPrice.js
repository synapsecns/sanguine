import { useSwr } from '@hooks/useSwr'

export function useCoingeckoPrice(symbol) {
  let apiUrl = `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&symbols=${symbol}`

  const { data } = useSwr(apiUrl)

  if (data) {
    let coin = data[0]
    return coin?.current_price
  }
}
