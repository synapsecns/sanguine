import Card from '@tw/Card'
import { CHAINS_BY_ID } from '@constants/chains'
import { getNetworkTextColor } from '@styles/chains'
import { useMemo } from 'react'

const NoSwapCard = ({ chainId }: { chainId: number }) => {
  const chain = useMemo(() => CHAINS_BY_ID[chainId], [chainId])
  return (
    <Card
      title="Swap"
      divider={false}
      className={`
        transform transition-all duration-100 rounded-xl min-w-[320px]
      `}
    >
      <div className="w-full pt-4 text-center text-gray-400">
        No swaps available on{' '}
        <span className={`${getNetworkTextColor(chain?.color)} font-medium`}>
          {chain?.name ?? 'current network'}
        </span>
      </div>
    </Card>
  )
}

export default NoSwapCard
