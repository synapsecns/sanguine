import { QuestionMarkCircleIcon } from '@heroicons/react/outline'

import { CHAIN_INFO_MAP } from '@constants/networks'
import { getNetworkTextColor } from '@styles/networks'
import Image from 'next/image'

export function ChainInfo({
  chainId,
  imgClassName = 'w-4 h-4',
  textClassName = getNetworkTextColor(chainId),
}) {
  const { chainName, chainImg } = CHAIN_INFO_MAP[chainId] ?? {}

  if (chainName) {
    return (
      <div className="flex items-center">
        <Image
          className={`inline mr-2 rounded-lg ${imgClassName}`}
          src={chainImg}
          alt={chainImg}
        />
        <span className={textClassName}>{chainName}</span>
      </div>
    )
  } else {
    return (
      <div className="flex items-center">
        <QuestionMarkCircleIcon
          className={`inline mr-2 rounded-lg ${imgClassName}`}
          strokeWidth={1}
        />
        <span>--</span>
      </div>
    )
  }
}
