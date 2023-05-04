import Card from '@/components/ui/tailwind/Card'
import { getNetworkTextColor } from '@/styles/chains'
import { Chain } from '@/utils/types'

const NoStakeCard = ({ chain }: { chain?: Chain }) => {
  const chainName = chain?.name ?? 'current network'
  const networkColor = chain?.color
  return (
    <Card
      divider={false}
      className={`
        transform transition-all duration-100
        rounded-xl max-w-[420px]
      `}
    >
      <div className="w-full pt-4 text-center text-gray-400">
        No stakes available on{' '}
        <span className={`${getNetworkTextColor(networkColor)} font-medium`}>
          {chainName}
        </span>
      </div>
    </Card>
  )
}

export default NoStakeCard
