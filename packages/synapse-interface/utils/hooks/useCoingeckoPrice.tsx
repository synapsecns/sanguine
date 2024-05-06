import { useAppSelector } from '@/store/hooks'

export const useCoingeckoPrice = (symbol: string) => {
  const { prices } = useAppSelector((state) => state.gasAirdrop)

  const data = prices.find((price) => price.symbol === symbol)

  return data?.price
}
