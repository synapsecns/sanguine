import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import Card from '@tw/Card'
import { CHAIN_INFO_MAP } from '@constants/networks'
import { getNetworkTextColor } from '@utils/styles/networks'

export default function NoSwapCard() {
  const { chainId } = useActiveWeb3React()

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
        <span className={`${getNetworkTextColor(chainId)} font-medium`}>
          {CHAIN_INFO_MAP[chainId].chainName}
        </span>
      </div>
    </Card>
  )
}
