import { useSwr } from '@hooks/useSwr'

const ID_MAP = {
  ETH: 'ethereum',
  AVAX: 'avalanche-2',
  JEWEL: 'defi-kingdoms',
  MOVR: 'moonriver',
  GLMR: 'moonbeam',
  CANTO: 'canto',
  FTM: 'fantom',
  Metis: 'metis-token',
  BNB: 'binancecoin',
  MATIC: 'matic-network',
  KLAY: 'klay-token',
}

export const useCoingeckoPrice = (symbol: string) => {
  const id = ID_MAP[symbol]

  const apiUrl = `https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=${id}`

  const { data } = useSwr(apiUrl)

  if (data) {
    const coin = data[0]
    return coin?.current_price
  }
}
