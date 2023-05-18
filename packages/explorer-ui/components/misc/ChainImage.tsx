import { QuestionMarkCircleIcon } from '@heroicons/react/outline'
import { CHAIN_INFO_MAP } from '@constants/networks'
import Image from 'next/image'

export function ChainImage({ chainId, imgSize = 'w-4 h-4', className }) {
  if (chainId) {
    const chainImg = CHAIN_INFO_MAP[chainId]
    return (
      <Image
        src={chainImg.chainImg}
        className={`${imgSize} rounded-full mr-2 inline ${className}`}
        alt={chainImg.chainName}
      />
    )
  } else {
    return (
      <QuestionMarkCircleIcon
        className={`${imgSize} rounded-full mr-2 inline`}
        strokeWidth={1}
      />
    )
  }
}
