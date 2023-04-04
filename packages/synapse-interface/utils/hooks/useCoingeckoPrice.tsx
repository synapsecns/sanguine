import { useSwr } from '@hooks/useSwr'

export const useCoingeckoPrice = (symbol: string) => {
  const apiUrl = `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&symbols=${symbol}`

  const { data } = useSwr(apiUrl)

  if (data) {
    const coin = data[0]
    return coin?.current_price
  }
}
